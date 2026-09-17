import { api, getToken } from "./api.js";
import { ThreadView, el } from "./thread-view.js";
import { currentSelection } from "./model-select.js";
import { getFamilies, getMaxTokens, getFeaturePref, plusModelSubmenuContains, closePlusModelSubmenu } from "./model-select.js";
import { initReasonPanel } from "./reasoning-panel.js";
import { initRequestsPanel } from "./requests-panel.js";
import {
  estimateTokens,
  refreshCtxCounter,
  refreshModelMeta,
  resetTurnTokens,
  setInputEstimate,
} from "./turn-tokens.js";

export function initChat() {
  const log = document.getElementById("chat-container");
  const input = document.getElementById("prompt-input");
  const stopBtn = document.getElementById("stop-btn");
  const routeBadge = document.getElementById("token-info");
  const statsBadge = document.getElementById("cost-info");
  const modelAlert = document.getElementById("model-alert");
  const inputHint = document.getElementById("input-hint");

  const emptyHTML = document.getElementById("empty-chat-placeholder")
    ? document.getElementById("empty-chat-placeholder").outerHTML
    : "";

  let attachments = [];
  const attachPreview = document.getElementById("attach-preview");
  const fileInput = document.getElementById("file-input");
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
    actions: true,
    reasonPanel: true,
    trackTokens: true,
    getPayload: (text) => {
      const sel = currentSelection();
      return {
        family: sel.family,
        mode: sel.mode,
        message: text,
        web: sel.web,
        mcp: sel.mcp,
        think: sel.think,
        effort: sel.effort,
        approve: false,
        plan: false,
        max_tokens: getMaxTokens(),
        attachments: attachments.map((a) => a.id),
      };
    },
    onDone: () => refreshHint(),
  });

  function showModelAlert(show) {
    if (modelAlert) modelAlert.style.display = show ? "" : "none";
  }
  const alertClose = document.getElementById("model-alert-close");
  if (alertClose) alertClose.addEventListener("click", () => showModelAlert(false));

  function refreshHint() {
    if (!inputHint) return;
    const sel = currentSelection();
    const fams = getFamilies();
    const f = fams.find((x) => x.id === sel.family);
    const m = f ? f.modes.find((x) => x.mode === sel.mode) : null;
    const parts = [];
    if (f) parts.push(f.label);
    if (sel.mode === "auto") parts.push("Défaut");
    else if (m) parts.push(m.label);
    const extra = [];
    if (sel.web) extra.push("web");
    if (sel.think) extra.push("réflexion");
    inputHint.textContent = parts.join(" · ") + (extra.length ? " — " + extra.join(", ") : "");
    refreshModelMeta();
  }
  document.getElementById("family-select")?.addEventListener("change", refreshHint);
  document.getElementById("mode-select")?.addEventListener("change", refreshHint);
  window.addEventListener("cetas:composer-toggles", refreshHint);
  window.addEventListener("cetas:model-changed", refreshHint); // choix depuis le menu +

  function renderPreview() {
    if (!attachPreview) return;
    attachPreview.innerHTML = "";
    for (const a of attachments) {
      const chip = el("div", "attach-chip-preview");
      const img = el("img", "attach-thumb");
      img.alt = a.name;
      thumbFor(a.id).then((url) => {
        if (url) img.src = url;
        else img.remove();
      });
      chip.appendChild(img);
      chip.appendChild(el("span", "attach-name", a.name));
      const rm = el("button", "attach-thumb-remove", "×");
      rm.type = "button";
      rm.setAttribute("aria-label", "Retirer");
      rm.addEventListener("click", () => removeAttachment(a));
      chip.appendChild(rm);
      attachPreview.appendChild(chip);
    }
  }

  function addAttachment(file) {
    attachments.push({ id: file.id, name: file.name });
    renderPreview();
  }

  function removeAttachment(entry) {
    attachments = attachments.filter((a) => a !== entry);
    renderPreview();
    api("/api/chat/attach/" + encodeURIComponent(entry.id), { method: "DELETE" }).catch(() => {});
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
    const sel = currentSelection();
    if (!sel.family || !sel.mode) {
      showModelAlert(true);
      return;
    }
    showModelAlert(false);
    input.value = "";
    autoGrow();
    const ok = await view.sendText(text);
    if (ok) {
      attachments = [];
      renderPreview();
      window.dispatchEvent(new CustomEvent("cetas:chat-changed"));
    }
  }

  function autoGrow() {
    input.style.height = "auto";
    input.style.height = Math.min(input.scrollHeight, 200) + "px";
  }

  input.addEventListener("input", autoGrow);
  input.addEventListener("input", () => setInputEstimate(estimateTokens(input.value)));
  input.addEventListener("keydown", (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      send();
    }
  });
  if (stopBtn) stopBtn.addEventListener("click", () => view.stop());

  // Plus-menu
  const plusBtn = document.getElementById("plus-menu-btn");
  const plusMenu = document.getElementById("plus-menu-dropdown");
  if (plusBtn && plusMenu) {
    plusBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      const open = plusMenu.style.display === "none";
      plusMenu.style.display = open ? "" : "none";
    });
    document.addEventListener("click", (e) => {
      if (plusMenu.style.display !== "none" && !plusMenu.contains(e.target) && !plusModelSubmenuContains(e.target) && e.target !== plusBtn && !plusBtn.contains(e.target)) {
        plusMenu.style.display = "none";
        closePlusModelSubmenu();
      }
    });
    plusMenu.querySelectorAll(".plus-menu-item").forEach((item) => {
      item.addEventListener("click", () => {
        const action = item.dataset.action;
        plusMenu.style.display = "none";
        if (action === "attach" && fileInput) fileInput.click();
      });
    });
  }

  if (fileInput) {
    fileInput.addEventListener("change", () => {
      if (fileInput.files && fileInput.files.length) uploadFiles(Array.from(fileInput.files));
      fileInput.value = "";
    });
  }
  const inputWrapper = input.closest(".input-wrapper");
  if (inputWrapper) {
    inputWrapper.addEventListener("dragover", (e) => {
      e.preventDefault();
      inputWrapper.classList.add("drag-over");
    });
    inputWrapper.addEventListener("dragleave", () => inputWrapper.classList.remove("drag-over"));
    inputWrapper.addEventListener("drop", (e) => {
      e.preventDefault();
      inputWrapper.classList.remove("drag-over");
      if (e.dataTransfer && e.dataTransfer.files && e.dataTransfer.files.length) {
        uploadFiles(Array.from(e.dataTransfer.files));
      }
    });
  }

  // Micro — masqué si la transcription est désactivée dans Fonctionnalités.
  const micBtn = document.getElementById("mic-btn");
  const applyMicPref = () => {
    if (!micBtn) return;
    if (!(window.SpeechRecognition || window.webkitSpeechRecognition)) {
      micBtn.style.display = "none";
      return;
    }
    micBtn.style.display = getFeaturePref("transcription", "system") === "none" ? "none" : "";
  };
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
    applyMicPref();
    window.addEventListener("cetas:features-changed", applyMicPref);
  } else if (micBtn) {
    micBtn.style.display = "none";
  }

  // Export (menu partage)
  const shareBtn = document.getElementById("share-btn");
  const shareMenu = document.getElementById("share-menu");
  if (shareBtn && shareMenu) {
    shareBtn.style.display = "";
    shareBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      shareMenu.classList.toggle("open");
    });
    document.addEventListener("click", () => shareMenu.classList.remove("open"));
    const dl = async (format) => {
      try {
        const resp = await fetch("/api/chat/export?format=" + format, {
          headers: { Authorization: "Bearer " + getToken() },
        });
        if (!resp.ok) throw new Error("export impossible");
        const blob = await resp.blob();
        const a = document.createElement("a");
        a.href = URL.createObjectURL(blob);
        a.download = "conversation." + (format === "html" ? "html" : "md");
        document.body.appendChild(a);
        a.click();
        a.remove();
        setTimeout(() => URL.revokeObjectURL(a.href), 1000);
      } catch (e) {
        view.addError(e.message);
      }
    };
    document.getElementById("share-menu-md")?.addEventListener("click", () => dl("md"));
    document.getElementById("share-menu-html")?.addEventListener("click", () => dl("html"));
  }
  const summaryBtn = document.getElementById("summary-btn");
  if (summaryBtn) summaryBtn.style.display = "none";

  // Reset du chat (nouvelle conversation gérée par la sidebar)
  window.addEventListener("cetas:chat-reset", () => {
    view.reset();
    attachments = [];
    renderPreview();
    resetTurnTokens();
    refreshHint();
  });

  // Ouverture / création / suppression de session (sidebar) : le serveur a
  // déjà basculé la session courante. On déconnecte le stream de l'ancienne
  // session, on vide la vue, puis on reconnecte : la nouvelle session est
  // rejouée depuis 0 via SSE.
  window.addEventListener("cetas:session-open", () => {
    view.disconnect();
    view.reset();
    attachments = [];
    renderPreview();
    resetTurnTokens();
    refreshHint();
    refreshCtxCounter();
    view.connect();
  });

  initReasonPanel();
  initRequestsPanel();
  // Flèche "retour en bas" ancrée au-dessus du composer, centrée
  // (façon DeepSeek) plutôt qu'en bas à droite.
  const inputArea = document.querySelector(".input-area");
  if (view.toBottomBtn && inputArea) {
    if (getComputedStyle(inputArea).position === "static") inputArea.style.position = "relative";
    inputArea.appendChild(view.toBottomBtn);
  }
  refreshHint();
  refreshCtxCounter();
  view.connect();
  return view;
}
