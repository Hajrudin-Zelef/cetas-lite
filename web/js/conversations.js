import { api } from "./api.js";

function fmtDate(ms) {
  if (!ms) return "";
  try {
    return new Date(ms).toLocaleString();
  } catch (e) {
    return "";
  }
}

async function restore(id, onChanged) {
  const state = await api("/api/chat/state").catch(() => ({}));
  if ((state.turns || 0) > 0 && !confirm("Restaurer cette conversation ? La conversation actuelle sera archivée.")) {
    return;
  }
  try {
    await api("/api/conversations/restore", { method: "POST", body: { id } });
    await onChanged();
  } catch (e) {}
}

async function remove(id, onChanged) {
  if (!confirm("Supprimer définitivement cette conversation ?")) return;
  try {
    await api("/api/conversations/" + encodeURIComponent(id), { method: "DELETE" });
    await onChanged();
  } catch (e) {}
}

function convItem(a, onChanged) {
  const row = document.createElement("div");
  row.className = "conv-item";

  const main = document.createElement("button");
  main.type = "button";
  main.className = "conv-main";
  main.title = a.title || "";
  const title = document.createElement("span");
  title.className = "conv-title";
  title.textContent = a.title || "(sans titre)";
  const meta = document.createElement("span");
  meta.className = "conv-meta";
  const n = a.messages || 0;
  meta.textContent = n + " message" + (n > 1 ? "s" : "") + (a.updated ? " · " + fmtDate(a.updated) : "");
  main.appendChild(title);
  main.appendChild(meta);
  main.addEventListener("click", () => restore(a.id, onChanged));

  const del = document.createElement("button");
  del.type = "button";
  del.className = "conv-del";
  del.title = "Supprimer";
  del.setAttribute("aria-label", "Supprimer la conversation");
  del.textContent = "\u00d7";
  del.addEventListener("click", (e) => {
    e.stopPropagation();
    remove(a.id, onChanged);
  });

  row.appendChild(main);
  row.appendChild(del);
  return row;
}

export function initConversations() {
  const list = document.getElementById("conv-list");
  const newBtn = document.getElementById("new-chat-btn");

  async function refresh() {
    let data;
    try {
      data = await api("/api/conversations");
    } catch (e) {
      return;
    }
    const archives = (data && data.archives) || [];
    list.innerHTML = "";
    if (!archives.length) {
      const empty = document.createElement("div");
      empty.className = "conv-empty";
      empty.textContent = "Aucune conversation archivée.";
      list.appendChild(empty);
      return;
    }
    for (const a of archives) list.appendChild(convItem(a, refresh));
  }

  newBtn.addEventListener("click", async () => {
    const state = await api("/api/chat/state").catch(() => ({}));
    if ((state.turns || 0) > 0 && !confirm("Démarrer une nouvelle conversation ? La conversation actuelle sera archivée.")) {
      return;
    }
    try {
      await api("/api/chat/reset", { method: "POST" });
    } catch (e) {}
    await refresh();
  });

  refresh();
}
