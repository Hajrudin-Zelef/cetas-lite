import { api, getToken } from "./api.js";
import { ThreadView, el } from "./thread-view.js";
import {
  applyWebToggle,
  applyMCPToggle,
  applyThinkingToggle,
  setThinking,
  applyApproveToggle,
  applyPlanToggle,
  isAgentMode,
  persistPrefs,
} from "./model-select.js";

export function initChat() {
  const log = document.getElementById("chat-log");
  const input = document.getElementById("prompt-input");
  const form = document.getElementById("composer");
  const stopBtn = document.getElementById("stop-btn");
  const routeBadge = document.getElementById("route-badge");
  const statsBadge = document.getElementById("stats-badge");
  const webToggle = document.getElementById("web-toggle");

  const emptyHTML = document.getElementById("empty-chat").outerHTML;

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

  function selection() {
    const family = document.getElementById("family-select").value;
    const mode = document.getElementById("mode-select").value;
    const web = webToggle ? webToggle.getAttribute("aria-pressed") === "true" : false;
    const mcpToggle = document.getElementById("mcp-toggle");
    const mcp = mcpToggle ? mcpToggle.getAttribute("aria-pressed") === "true" : false;
    const thinkToggle = document.getElementById("thinking-toggle");
    const think = thinkToggle ? thinkToggle.getAttribute("aria-pressed") === "true" : false;
    const effortSel = document.getElementById("effort-select");
    const effort = effortSel ? effortSel.value : "default";
    const approveToggle = document.getElementById("approve-toggle");
    const approve = approveToggle && !approveToggle.hidden ? approveToggle.getAttribute("aria-pressed") === "true" : false;
    const planToggle = document.getElementById("plan-toggle");
    const plan = planToggle && !planToggle.hidden ? planToggle.getAttribute("aria-pressed") === "true" : false;
    return { family, mode, web, mcp, think, effort, approve: isAgentMode() && approve, plan: isAgentMode() && plan };
  }

  const view = new ThreadView({
    log,
    stopBtn,
    routeBadge,
    statsBadge,
    emptyHTML,
    streamURL: (from) => "/api/chat/stream?from=" + from,
    sendURL: "/api/chat/send",
    stopURL: "/api/chat/stop",
    approveURL: "/api/chat/approve",
    regenerateURL: "/api/chat/regenerate",
    getPayload: (text) => {
      const sel = selection();
      return {
        family: sel.family,
        mode: sel.mode,
        message: text,
        web: sel.web,
        mcp: sel.mcp,
        think: sel.think,
        effort: sel.effort,
        approve: sel.approve,
        plan: sel.plan,
        attachments: attachments.map((a) => a.id),
      };
    },
  });

  function addAttachment(file) {
    const chip = el("div", "attach-chip");
    const thumb = el("img", "attach-thumb");
    thumb.alt = file.name;
    chip.appendChild(thumb);
    chip.appendChild(el("span", "attach-name", file.name));
    const remove = el("button", "attach-remove", "×");
    remove.type = "button";
    remove.setAttribute("aria-label", "Retirer la pièce jointe");
    chip.appendChild(remove);
    attachChips.appendChild(chip);
    const entry = { id: file.id, name: file.name, chip };
    remove.addEventListener("click", () => removeAttachment(entry));
    thumbFor(file.id).then((url) => {
      if (url) thumb.src = url;
      else thumb.remove();
    });
    attachments.push(entry);
  }

  function removeAttachment(entry) {
    attachments = attachments.filter((a) => a !== entry);
    if (entry.chip && entry.chip.parentNode) entry.chip.remove();
    api("/api/chat/attach/" + encodeURIComponent(entry.id), { method: "DELETE" }).catch(() => {});
  }

  function renderChips() {
    attachChips.innerHTML = "";
    const keep = attachments;
    attachments = [];
    for (const a of keep) addAttachment({ id: a.id, name: a.name });
  }

  async function uploadFiles(files) {
    for (const f of files) {
      try {
        const fd = new FormData();
        fd.append("file", f);
        const resp = await fetch("/api/chat/attach", {
          method: "POST",
          headers: { Authorization: "Bearer " + getToken() },
          body: fd,
        });
        if (!resp.ok) throw new Error("échec du téléversement (" + resp.status + ")");
        const data = await resp.json();
        addAttachment({ id: data.id, name: data.name });
      } catch (e) {
        view.addError("Pièce jointe : " + e.message);
      }
    }
  }

  async function send() {
    const text = input.value.trim();
    if (!text || view.generating) return;
    const sel = selection();
    if (!sel.family || !sel.mode) {
      view.addError("Choisis une famille et un mode.");
      return;
    }
    input.value = "";
    autoGrow();
    const ok = await view.sendText(text);
    if (ok) {
      attachments = [];
      renderChips();
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
  stopBtn.addEventListener("click", () => view.stop());
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

  view.connect();
}
