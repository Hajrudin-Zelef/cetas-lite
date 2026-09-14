import { api, getToken, readSSE } from "./api.js";
import { appendLinkified, renderInto } from "./markdown.js";
import { createStreamRenderer } from "./stream-render.js";
import { applyWebToggle, applyMCPToggle, persistPrefs, setThinking } from "./model-select.js";

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
    const effortSel = document.getElementById("effort-select");
    return {
      family: document.getElementById("family-select").value,
      mode: document.getElementById("mode-select").value,
      web: !!(webToggle && webToggle.getAttribute("aria-pressed") === "true"),
      mcp: !!(mcpToggle && !mcpToggle.hidden && mcpToggle.getAttribute("aria-pressed") === "true"),
      think: !!(thinkingToggle && thinkingToggle.getAttribute("aria-pressed") === "true"),
      effort: effortSel ? effortSel.value : "default",
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
    log.appendChild(el("div", "msg user", text));
    scroll(true);
  }

  function ensureAssistant() {
    if (!assistant) {
      clearEmpty();
      assistant = el("div", "msg assistant");
      const body = el("div", "msg-text");
      assistant.appendChild(body);
      log.appendChild(assistant);
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
      reasoningEl = el("details", "msg-reasoning");
      reasoningEl.appendChild(el("summary", null, "Raisonnement"));
      reasoningEl.appendChild(el("div", "reasoning-body"));
      log.insertBefore(reasoningEl, assistant);
    }
    return reasoningEl;
  }

  function appendReasoning(text, isReplace) {
    ensureReasoning();
    reasoningText = isReplace ? String(text) : reasoningText + String(text);
    reasoningEl.querySelector(".reasoning-body").textContent = reasoningText;
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
    log.appendChild(el("div", "msg error", text));
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

  function addActions(box, raw) {
    if (!box || box.querySelector(".msg-actions")) return;
    const bar = el("div", "msg-actions");
    const copy = el("button", "msg-action", "Copier");
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
    const regen = el("button", "msg-action", "Regenerer");
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
    box.appendChild(bar);
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

  function renderDiff(lines) {
    const wrap = el("div", "tool-diff");
    for (const line of lines) {
      const cls = line.kind === "+" ? "diff-add" : "diff-del";
      wrap.appendChild(el("span", cls, (line.kind === "+" ? "+ " : "- ") + line.text + "\n"));
    }
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
      body.appendChild(renderDiff(ev.diff));
    }
    if (typeof ev.result === "string" && ev.result) {
      const pre = el("pre", "tool-result");
      appendLinkified(pre, ev.result);
      body.appendChild(pre);
    }
    scroll();
  }

  function finishTurn() {
    if (textRenderer) textRenderer.flush();
    const box = assistant;
    const raw = textRenderer ? textRenderer.text() : "";
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
      await api("/api/chat/send", { method: "POST", body: { family: sel.family, mode: sel.mode, message: text, web: sel.web, mcp: sel.mcp, think: sel.think, effort: sel.effort, attachments: attachments.map((a) => a.id) } });
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
      setThinking(next);
      persistPrefs().catch(() => {});
    });
  }
  const effortSel = document.getElementById("effort-select");
  if (effortSel) {
    effortSel.addEventListener("change", () => persistPrefs().catch(() => {}));
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

  connect();
}
