import { api, getToken, readSSE } from "./api.js";
import { appendLinkified, renderInto } from "./markdown.js";
import { createStreamRenderer } from "./stream-render.js";
import { applyWebToggle, applyMCPToggle, persistPrefs, setThinking, applyApproveToggle, applyPlanToggle, isAgentMode } from "./model-select.js";

const BRAILLE = ["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"];

export function initChat() {
  const log = document.getElementById("chat-log");
  const input = document.getElementById("prompt-input");
  const form = document.getElementById("composer");
  const stopBtn = document.getElementById("stop-btn");
  const routeBadge = document.getElementById("route-badge");
  const statsBadge = document.getElementById("stats-badge");
  const webToggle = document.getElementById("web-toggle");

  const emptyHTML = document.getElementById("empty-chat").outerHTML;
  let empty = document.getElementById("empty-chat");

  let lastSeq = 0;
  let assistant = null;
  let textRenderer = null;
  let reasoningEl = null;
  let reasoningText = "";
  let controller = null;
  let generating = false;
  let waitEl = null;
  let waitTimer = null;
  let waitIdx = 0;
  const toolBoxes = new Map();
  let attachments = [];
  const attachChips = document.getElementById("attach-chips");
  const attachInput = document.getElementById("attach-input");
  const attachBtn = document.getElementById("attach-btn");
  const thumbCache = {};

  async function thumbFor(id) {
    if (thumbCache[id]) return thumbCache[id];
    try {
      const resp = await fetch("/api/chat/attach/" + encodeURIComponent(id), {
        headers: { Authorization: "Bearer " + getToken() },
      });
      if (!resp.ok) return "";
      const url = URL.createObjectURL(await resp.blob());
      thumbCache[id] = url;
      return url;
    } catch (e) {
      return "";
    }
  }

  function el(tag, cls, text) {
    const e = document.createElement(tag);
    if (cls) e.className = cls;
    if (text !== undefined) e.textContent = text;
    return e;
  }

  function clearEmpty() {
    if (empty && empty.parentNode) empty.remove();
    empty = null;
  }

  function reAddEmpty() {
    log.innerHTML = emptyHTML;
    empty = document.getElementById("empty-chat");
  }

  function atBottom() {
    return log.scrollHeight - log.scrollTop - log.clientHeight < 80;
  }

  function scroll(force) {
    if (force || atBottom()) log.scrollTop = log.scrollHeight;
  }

  function selection() {
    const mcpToggle = document.getElementById("mcp-toggle");
    const thinkingToggle = document.getElementById("thinking-toggle");
    const approveToggle = document.getElementById("approve-toggle");
    const planToggle = document.getElementById("plan-toggle");
    const effortSel = document.getElementById("effort-select");
    const agent = isAgentMode();
    return {
      family: document.getElementById("family-select").value,
      mode: document.getElementById("mode-select").value,
      web: !!(webToggle && webToggle.getAttribute("aria-pressed") === "true"),
      mcp: !!(mcpToggle && !mcpToggle.hidden && mcpToggle.getAttribute("aria-pressed") === "true"),
      think: !!(thinkingToggle && thinkingToggle.getAttribute("aria-pressed") === "true"),
      effort: effortSel ? effortSel.value : "default",
      approve: agent && !!(approveToggle && !approveToggle.hidden && approveToggle.getAttribute("aria-pressed") === "true"),
      plan: agent && !!(planToggle && !planToggle.hidden && planToggle.getAttribute("aria-pressed") === "true"),
    };
  }

  function setBusy(v) {
    log.setAttribute("aria-busy", v ? "true" : "false");
  }

  function showWait() {
    if (waitEl) return;
    waitEl = el("div", "stream-waiting");
    waitEl.appendChild(el("span", "stream-spinner", BRAILLE[0]));
    log.appendChild(waitEl);
    waitTimer = setInterval(() => {
      waitIdx = (waitIdx + 1) % BRAILLE.length;
      if (waitEl && waitEl.firstChild) waitEl.firstChild.textContent = BRAILLE[waitIdx];
    }, 80);
    scroll(true);
  }

  function hideWait() {
    if (waitTimer) {
      clearInterval(waitTimer);
      waitTimer = null;
    }
    if (waitEl) {
      waitEl.remove();
      waitEl = null;
    }
  }

  function addUser(text) {
    clearEmpty();
    const wrapper = el("div", "message-wrapper message-wrapper-user");
    const bubble = el("div", "message message-user");
    bubble.appendChild(el("div", "message-text", text));
    wrapper.appendChild(bubble);
    log.appendChild(wrapper);
    scroll(true);
  }

  function ensureAssistant() {
    if (!assistant) {
      clearEmpty();
      const wrapper = el("div", "message-wrapper message-wrapper-assistant");
      assistant = el("div", "message message-assistant streaming");
      const body = el("div", "message-text");
      assistant.appendChild(body);
      wrapper.appendChild(assistant);
      log.appendChild(wrapper);
      textRenderer = createStreamRenderer({
        onRender: (t) => renderInto(body, t),
        onFirst: hideWait,
      });
    }
    return assistant;
  }

  function appendContent(text, isReplace) {
    ensureAssistant();
    if (isReplace) textRenderer.replace(text);
    else textRenderer.add(text);
    scroll();
  }

  function ensureReasoning() {
    if (!reasoningEl) {
      ensureAssistant();
      reasoningEl = el("details", "thinking-block");
      reasoningEl.open = true;
      reasoningEl.appendChild(el("summary", null, "Raisonnement"));
      reasoningEl.appendChild(el("div", "thinking-content"));
      assistant.insertBefore(reasoningEl, assistant.firstChild);
    }
    return reasoningEl;
  }

  function appendReasoning(text, isReplace) {
    ensureReasoning();
    reasoningText = isReplace ? String(text) : reasoningText + String(text);
    reasoningEl.querySelector(".thinking-content").textContent = reasoningText;
    scroll();
  }

  function removeReasoning() {
    if (reasoningEl) {
      reasoningEl.remove();
      reasoningEl = null;
      reasoningText = "";
    }
  }

  function addError(text) {
    hideWait();
    setBusy(false);
    clearEmpty();
    const wrapper = el("div", "message-wrapper message-wrapper-assistant");
    const bubble = el("div", "message message-assistant message-error");
    bubble.appendChild(el("div", "message-text", text));
    wrapper.appendChild(bubble);
    log.appendChild(wrapper);
    scroll();
  }

  function addSystem(text) {
    clearEmpty();
    log.appendChild(el("div", "msg-system", text));
    scroll();
  }

  function renderChips() {
    if (!attachChips) return;
    attachChips.innerHTML = "";
    attachChips.hidden = attachments.length === 0;
    for (const a of attachments) {
      const chip = el("span", "attach-chip");
      if (a.kind === "image") {
        const img = el("img", "attach-thumb");
        img.alt = a.name;
        thumbFor(a.id).then((u) => {
          if (u) img.src = u;
        });
        chip.appendChild(img);
      }
      chip.appendChild(el("span", "attach-name", a.name));
      const rm = el("button", "attach-remove", "\u00d7");
      rm.type = "button";
      rm.title = "Retirer";
      rm.setAttribute("aria-label", "Retirer " + a.name);
      rm.addEventListener("click", () => removeAttachment(a.id));
      chip.appendChild(rm);
      attachChips.appendChild(chip);
    }
  }

  async function uploadFiles(files) {
    for (const f of files) {
      const fd = new FormData();
      fd.append("file", f);
      let resp;
      try {
        resp = await fetch("/api/chat/attach", {
          method: "POST",
          headers: { Authorization: "Bearer " + getToken() },
          body: fd,
        });
      } catch (e) {
        addError("Envoi du fichier impossible.");
        continue;
      }
      const data = await resp.json().catch(() => ({}));
      if (!resp.ok) {
        addError(data.error || "Erreur " + resp.status);
        continue;
      }
      attachments.push({ id: data.id, name: data.name, kind: data.kind });
    }
    renderChips();
  }

  function removeAttachment(id) {
    attachments = attachments.filter((a) => a.id !== id);
    renderChips();
    fetch("/api/chat/attach/" + encodeURIComponent(id), {
      method: "DELETE",
      headers: { Authorization: "Bearer " + getToken() },
    }).catch(() => {});
  }

  function speakable(md) {
    return String(md || "")
      .replace(/```[\s\S]*?```/g, " (bloc de code) ")
      .replace(/`([^`]+)`/g, "$1")
      .replace(/^#{1,6}\s*/gm, "")
      .replace(/[*_>#|]/g, " ")
      .replace(/\s+/g, " ")
      .trim();
  }

  function addActions(box, raw) {
    const wrapper = box.closest(".message-wrapper") || box;
    if (!wrapper || wrapper.querySelector(".message-btn-row")) return;
    const bar = el("div", "message-btn-row");
    const copy = el("button", "message-copy-btn", "Copier");
    copy.type = "button";
    copy.addEventListener("click", () => {
      if (!navigator.clipboard) return;
      navigator.clipboard.writeText(raw).then(() => {
        copy.textContent = "Copie";
        setTimeout(() => {
          copy.textContent = "Copier";
        }, 1200);
      }).catch(() => {});
    });
    const regen = el("button", "regen-btn", "Regenerer");
    regen.type = "button";
    regen.addEventListener("click", async () => {
      if (generating) return;
      try {
        await api("/api/chat/regenerate", { method: "POST" });
      } catch (e) {
        addError(e.message);
      }
    });
    bar.appendChild(copy);
    bar.appendChild(regen);
    if (window.speechSynthesis) {
      const speak = el("button", "message-tts-btn", "Lire");
      speak.type = "button";
      speak.addEventListener("click", () => {
        window.speechSynthesis.cancel();
        const u = new SpeechSynthesisUtterance(speakable(raw));
        u.lang = document.documentElement.lang || "fr";
        window.speechSynthesis.speak(u);
      });
      bar.appendChild(speak);
    }
    wrapper.appendChild(bar);
  }

  function summarize(args) {
    if (!args) return "";
    return args.command || args.file_path || args.pattern || args.query || "";
  }

  function renderTodos(todos) {
    const list = el("ul", "chat-todo-list");
    for (const t of todos) {
      const status = t.status || "pending";
      list.appendChild(el("li", "chat-todo-item todo-" + status, t.content || ""));
    }
    return list;
  }

  function renderDiff(lines, filePath) {
    const wrap = el("div", "tool-diff");
    let adds = 0, dels = 0;
    for (const line of lines) {
      if (line.kind === "+") adds++;
      else if (line.kind === "-") dels++;
    }
    const head = el("div", "diff-head");
    head.appendChild(el("span", "diff-file", filePath ? String(filePath) : "modification"));
    head.appendChild(el("span", "diff-stats", "+" + adds + " / -" + dels));
    wrap.appendChild(head);
    const body = el("div", "diff-body");
    let oldN = 0, newN = 0;
    for (const line of lines) {
      const kind = line.kind === "+" ? "diff-add" : line.kind === "-" ? "diff-del" : "diff-ctx";
      const row = el("div", "diff-row " + kind);
      let gutter = "";
      if (line.kind === "-") { oldN++; gutter = String(oldN); }
      else if (line.kind === "+") { newN++; gutter = String(newN); }
      row.appendChild(el("span", "diff-gutter", gutter));
      row.appendChild(el("span", "diff-sign", line.kind === "…" ? "…" : line.kind || " "));
      row.appendChild(el("span", "diff-code", line.text != null ? String(line.text) : ""));
      body.appendChild(row);
    }
    wrap.appendChild(body);
    return wrap;
  }

  function addTool(ev) {
    const key = ev.name + "|" + JSON.stringify(ev.args || {});
    if (ev.phase === "start") {
      clearEmpty();
      hideWait();
      const box = el("details", "msg-tool");
      const summary = el("summary");
      summary.appendChild(el("span", "tool-name", ev.name));
      const hint = summarize(ev.args);
      if (hint) summary.appendChild(el("span", "tool-hint", " " + hint));
      box.appendChild(summary);
      const body = el("div", "tool-body");
      if (ev.name === "TodoWrite" && ev.args && Array.isArray(ev.args.todos)) {
        body.appendChild(renderTodos(ev.args.todos));
      }
      box.appendChild(body);
      log.appendChild(box);
      toolBoxes.set(key, body);
      if (textRenderer) textRenderer.flush();
      assistant = null;
      textRenderer = null;
      reasoningEl = null;
      reasoningText = "";
      scroll(true);
      return;
    }
    const body = toolBoxes.get(key);
    if (!body) return;
    if (Array.isArray(ev.diff) && ev.diff.length) {
      body.appendChild(renderDiff(ev.diff, ev.args && ev.args.file_path));
    }
    if (typeof ev.result === "string" && ev.result) {
      const pre = el("pre", "tool-result");
      appendLinkified(pre, ev.result);
      body.appendChild(pre);
    }
    scroll();
  }

  const approvalCards = new Map();

  function summarizeApprovalArgs(tool, args) {
    if (!args) return "";
    if (args.file_path) return String(args.file_path);
    if (args.command) return String(args.command);
    if (args.pattern) return String(args.pattern);
    if (args.query) return String(args.query);
    return "";
  }

  async function decideApproval(id, approved, always) {
    try {
      await api("/api/chat/approve", { method: "POST", body: { id, approved, always: !!always } });
    } catch (e) {
      addError("Approbation : " + e.message);
    }
  }

  function setApprovalResolved(id, approved, timeout) {
    const card = approvalCards.get(id);
    if (!card) return;
    const btns = card.querySelectorAll("button");
    btns.forEach((b) => { b.disabled = true; });
    const status = card.querySelector(".approval-status");
    if (status) {
      status.textContent = timeout ? "Expirée (10 min sans réponse)" : approved ? "Approuvé" : "Refusé";
      status.classList.add(approved && !timeout ? "approved" : "denied");
    }
    scroll();
  }

  function addApproval(ev) {
    const id = ev.id;
    if (ev.phase === "resolved") {
      setApprovalResolved(id, ev.approved === true, ev.timeout === true);
      return;
    }
    if (ev.phase !== "request" || !id || approvalCards.has(id)) return;
    clearEmpty();
    hideWait();
    const card = el("div", "msg-approval");
    const head = el("div", "approval-head");
    const isPlan = ev.kind === "plan";
    head.appendChild(el("span", "approval-title", isPlan ? "Plan à valider" : "Approbation requise"));
    if (!isPlan && ev.tool) head.appendChild(el("span", "tool-name", " " + ev.tool));
    card.appendChild(head);
    const hint = isPlan ? "" : summarizeApprovalArgs(ev.tool, ev.args);
    if (hint) card.appendChild(el("div", "approval-hint", hint));
    if (isPlan && ev.plan) {
      const pre = el("pre", "approval-plan");
      pre.textContent = String(ev.plan).slice(0, 4000);
      card.appendChild(pre);
    } else if (ev.args && typeof ev.args === "object") {
      const det = el("details", "approval-args");
      det.appendChild(el("summary", "", "Détails"));
      const pre = el("pre");
      const shown = Object.assign({}, ev.args);
      for (const k of ["content", "code", "new"]) {
        if (typeof shown[k] === "string" && shown[k].length > 600) shown[k] = shown[k].slice(0, 600) + "…";
      }
      pre.textContent = JSON.stringify(shown, null, 2);
      det.appendChild(pre);
      card.appendChild(det);
    }
    const status = el("div", "approval-status", "En attente de ta décision…");
    card.appendChild(status);
    const row = el("div", "approval-actions");
    const btnApprove = el("button", "btn-approve", isPlan ? "Valider le plan" : "Approuver");
    btnApprove.addEventListener("click", () => decideApproval(id, true, false));
    const btnDeny = el("button", "btn-deny", "Refuser");
    btnDeny.addEventListener("click", () => decideApproval(id, false, false));
    row.appendChild(btnApprove);
    row.appendChild(btnDeny);
    if (!isPlan) {
      const btnAlways = el("button", "btn-always", "Toujours approuver (ce tour)");
      btnAlways.addEventListener("click", () => decideApproval(id, true, true));
      row.appendChild(btnAlways);
    }
    card.appendChild(row);
    log.appendChild(card);
    approvalCards.set(id, card);
    if (textRenderer) textRenderer.flush();
    scroll(true);
  }

  function finishTurn() {    if (textRenderer) textRenderer.flush();
    const box = assistant;
    const raw = textRenderer ? textRenderer.text() : "";
    if (box) box.classList.remove("streaming");
    hideWait();
    setBusy(false);
    generating = false;
    stopBtn.hidden = true;
    if (box && raw.trim()) addActions(box, raw);
    assistant = null;
    textRenderer = null;
    reasoningEl = null;
    reasoningText = "";
  }

  function resetLocal() {
    hideWait();
    setBusy(false);
    log.innerHTML = emptyHTML;
    empty = document.getElementById("empty-chat");
    lastSeq = 0;
    assistant = null;
    textRenderer = null;
    reasoningEl = null;
    reasoningText = "";
    toolBoxes.clear();
    approvalCards.clear();
    generating = false;
    stopBtn.hidden = true;
    routeBadge.hidden = true;
    attachments = [];
    renderChips();
    if (statsBadge) {
      statsBadge.hidden = true;
      statsBadge.textContent = "";
    }
  }

  function handleEvent(ev) {
    if (typeof ev.seq === "number" && ev.seq > lastSeq) lastSeq = ev.seq;
    if (ev.reset) {
      resetLocal();
      return;
    }
    if (ev.pad !== undefined) return;
    if (ev.caught_up) {
      scroll(true);
      return;
    }
    if (ev.user !== undefined) {
      addUser(String(ev.user));
      assistant = null;
      textRenderer = null;
      reasoningEl = null;
      reasoningText = "";
      generating = true;
      stopBtn.hidden = false;
      showWait();
      setBusy(true);
      return;
    }
    if (ev.reasoning_content !== undefined) {
      appendReasoning(String(ev.reasoning_content), ev.replace === true);
      return;
    }
    if (ev.content !== undefined) {
      appendContent(String(ev.content), ev.replace === true);
      return;
    }
    if (ev.tool !== undefined) {
      addTool(ev.tool);
      return;
    }
    if (ev.approval !== undefined) {
      addApproval(ev.approval);
      return;
    }
    if (ev.stats !== undefined) {
      const s = ev.stats || {};
      if (statsBadge) {
        statsBadge.hidden = false;
        statsBadge.textContent = "\u2191" + (s.prompt_tokens || 0) + " \u2193" + (s.completion_tokens || 0);
      }
      return;
    }
    if (ev.compact) {
      addSystem("Contexte compacté pour rester dans la fenêtre du modèle.");
      return;
    }
    if (ev.route !== undefined) {
      const r = ev.route || {};
      routeBadge.hidden = false;
      const name = r.label || r.model || "";
      routeBadge.textContent =
        (r.provider ? r.provider + "/" : "") + name + (r.local ? " (local)" : "") + (r.fallback ? " (repli)" : "");
      return;
    }
    if (ev.drop_reasoning) {
      removeReasoning();
      return;
    }
    if (ev.error !== undefined) {
      addError(String(ev.error));
      return;
    }
    if (ev.turn_done !== undefined) {
      finishTurn();
    }
  }

  async function connect() {
    if (controller) controller.abort();
    controller = new AbortController();
    try {
      const resp = await fetch("/api/chat/stream?from=" + lastSeq, {
        headers: { Authorization: "Bearer " + getToken() },
        signal: controller.signal,
      });
      if (!resp.ok || !resp.body) return;
      await readSSE(resp, handleEvent);
    } catch (e) {
      if (e && e.name === "AbortError") return;
      setTimeout(connect, 1500);
    }
  }

  async function send() {
    const text = input.value.trim();
    if (!text || generating) return;
    const sel = selection();
    if (!sel.family || !sel.mode) {
      addError("Choisis une famille et un mode.");
      return;
    }
    input.value = "";
    autoGrow();
    try {
      await api("/api/chat/send", { method: "POST", body: { family: sel.family, mode: sel.mode, message: text, web: sel.web, mcp: sel.mcp, think: sel.think, effort: sel.effort, approve: sel.approve, plan: sel.plan, attachments: attachments.map((a) => a.id) } });
      attachments = [];
      renderChips();
      generating = true;
      stopBtn.hidden = false;
      setBusy(true);
      scroll(true);
    } catch (err) {
      if (/en cours/i.test(err.message)) {
        generating = true;
        stopBtn.hidden = false;
        setBusy(true);
      } else {
        addError(err.message);
      }
    }
  }

  function autoGrow() {
    input.style.height = "auto";
    input.style.height = Math.min(input.scrollHeight, 200) + "px";
  }

  form.addEventListener("submit", (e) => {
    e.preventDefault();
    send();
  });
  input.addEventListener("input", autoGrow);
  input.addEventListener("keydown", (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      send();
    }
  });
  stopBtn.addEventListener("click", () => {
    api("/api/chat/stop", { method: "POST" }).catch(() => {});
  });
  if (webToggle) {
    webToggle.addEventListener("click", () => {
      const next = webToggle.getAttribute("aria-pressed") !== "true";
      applyWebToggle(next);
      persistPrefs().catch(() => {});
    });
  }
  const mcpToggle = document.getElementById("mcp-toggle");
  if (mcpToggle) {
    mcpToggle.addEventListener("click", () => {
      const next = mcpToggle.getAttribute("aria-pressed") !== "true";
      applyMCPToggle(next);
      persistPrefs().catch(() => {});
    });
  }
  const thinkingToggle = document.getElementById("thinking-toggle");
  if (thinkingToggle) {
    thinkingToggle.addEventListener("click", () => {
      if (thinkingToggle.disabled) return;
      const next = thinkingToggle.getAttribute("aria-pressed") !== "true";
      setThinking(next);      persistPrefs().catch(() => {});
    });
  }
  const effortSel = document.getElementById("effort-select");
  if (effortSel) {
    effortSel.addEventListener("change", () => persistPrefs().catch(() => {}));
  }
  const approveToggle = document.getElementById("approve-toggle");
  if (approveToggle) {
    approveToggle.addEventListener("click", () => {
      const next = approveToggle.getAttribute("aria-pressed") !== "true";
      applyApproveToggle(next);
    });
  }
  const planToggle = document.getElementById("plan-toggle");
  if (planToggle) {
    planToggle.addEventListener("click", () => {
      const next = planToggle.getAttribute("aria-pressed") !== "true";
      applyPlanToggle(next);
    });
  }

  if (attachBtn && attachInput) {
    attachBtn.addEventListener("click", () => attachInput.click());
    attachInput.addEventListener("change", () => {
      if (attachInput.files && attachInput.files.length) {
        uploadFiles(Array.from(attachInput.files));
      }
      attachInput.value = "";
    });
  }
  form.addEventListener("dragover", (e) => {
    e.preventDefault();
    form.classList.add("dragover");
  });
  form.addEventListener("dragleave", () => form.classList.remove("dragover"));
  form.addEventListener("drop", (e) => {
    e.preventDefault();
    form.classList.remove("dragover");
    if (e.dataTransfer && e.dataTransfer.files && e.dataTransfer.files.length) {
      uploadFiles(Array.from(e.dataTransfer.files));
    }
  });

  const micBtn = document.getElementById("mic-btn");
  const SR = window.SpeechRecognition || window.webkitSpeechRecognition;
  if (micBtn && SR) {
    const rec = new SR();
    rec.lang = document.documentElement.lang || "fr";
    rec.interimResults = false;
    rec.continuous = false;
    micBtn.addEventListener("click", () => {
      try {
        rec.start();
        micBtn.classList.add("active");
      } catch (e) {}
    });
    rec.addEventListener("result", (e) => {
      const t = e.results && e.results[0] && e.results[0][0] ? e.results[0][0].transcript : "";
      if (t) {
        input.value = (input.value ? input.value + " " : "") + t;
        autoGrow();
      }
    });
    rec.addEventListener("end", () => micBtn.classList.remove("active"));
    rec.addEventListener("error", () => micBtn.classList.remove("active"));
  } else if (micBtn) {
    micBtn.hidden = true;
  }

  const plusBtn = document.getElementById("plus-menu-btn");
  const plusMenu = document.getElementById("plus-menu-dropdown");
  if (plusBtn && plusMenu) {
    plusMenu.hidden = true;
    plusBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      plusMenu.hidden = !plusMenu.hidden;
    });
    document.addEventListener("click", (e) => {
      if (!plusMenu.hidden && !plusMenu.contains(e.target) && e.target !== plusBtn) {
        plusMenu.hidden = true;
      }
    });
  }

  connect();
}
