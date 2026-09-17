// Panneau flottant "Requêtes" façon DeepSeek : la liste des conversations
// affichée sur le côté du chat, scrollable haut/bas, clic pour basculer.
// Vue principale uniquement (la vue Agents reste identique à Marexcode).
import { api } from "./api.js";
import { confirmDialog } from "./dialogs.js";

export function initRequestsPanel() {
  const main = document.querySelector("main.main");
  if (!main || !document.getElementById("chat-container")) return null;
  if (main.querySelector(":scope > .requests-panel")) return null;
  if (getComputedStyle(main).position === "static") main.style.position = "relative";

  const panel = document.createElement("aside");
  panel.className = "requests-panel";
  panel.setAttribute("aria-label", "Requêtes");
  panel.hidden = true;
  main.appendChild(panel);

  async function currentId() {
    try {
      const st = await api("/api/chat/state");
      return st && st.id ? st.id : null;
    } catch (e) {
      return null;
    }
  }

  async function restore(id) {
    let state = {};
    try {
      state = await api("/api/chat/state");
    } catch (e) {}
    if ((state.turns || 0) > 0) {
      const ok = await confirmDialog(
        "Restaurer cette conversation ? La conversation actuelle sera archivée.",
        { okLabel: "Restaurer" }
      );
      if (!ok) return;
    }
    try {
      await api("/api/conversations/restore", { method: "POST", body: { id } });
      window.dispatchEvent(new CustomEvent("cetas:chat-reset"));
      await refresh();
    } catch (e) {}
  }

  async function refresh() {
    let archives = [];
    try {
      const data = await api("/api/conversations");
      archives = (data && data.archives) || [];
    } catch (e) {}
    const cur = await currentId();
    panel.innerHTML = "";
    panel.hidden = archives.length === 0;
    for (const a of archives) {
      const b = document.createElement("button");
      b.type = "button";
      b.className = "request-item" + (cur && a.id === cur ? " active" : "");
      b.title = a.title || "(sans titre)";
      b.textContent = a.title || "(sans titre)";
      b.dataset.id = a.id;
      b.addEventListener("click", () => restore(a.id));
      panel.appendChild(b);
    }
  }

  window.addEventListener("cetas:chat-reset", refresh);
  window.addEventListener("cetas:chat-changed", refresh);
  refresh();
  return { refresh, panel };
}
