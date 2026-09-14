import { api, getToken } from "./api.js";

async function download(url, name) {
  let resp;
  try {
    resp = await fetch(url, { headers: { Authorization: "Bearer " + getToken() } });
  } catch (e) {
    return;
  }
  if (!resp.ok) return;
  const blob = await resp.blob();
  const a = document.createElement("a");
  a.href = URL.createObjectURL(blob);
  a.download = name;
  document.body.appendChild(a);
  a.click();
  a.remove();
  setTimeout(() => URL.revokeObjectURL(a.href), 1000);
}

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

  const content = document.createElement("div");
  content.className = "conv-item-content";
  const title = document.createElement("div");
  title.className = "conv-item-title";
  title.textContent = a.title || "(sans titre)";
  title.title = a.title || "";
  const dateLine = document.createElement("div");
  dateLine.className = "conv-item-date-line";
  const date = document.createElement("span");
  date.className = "conv-item-date";
  const n = a.messages || 0;
  date.textContent = n + " message" + (n > 1 ? "s" : "") + (a.updated ? " · " + fmtDate(a.updated) : "");
  dateLine.appendChild(date);
  content.appendChild(title);
  content.appendChild(dateLine);
  content.addEventListener("click", () => restore(a.id, onChanged));

  const actions = document.createElement("div");
  actions.className = "conv-item-actions";

  const exp = document.createElement("button");
  exp.type = "button";
  exp.className = "conv-action-btn";
  exp.title = "Exporter (Markdown)";
  exp.setAttribute("aria-label", "Exporter la conversation");
  exp.textContent = "\u2913";
  exp.addEventListener("click", (e) => {
    e.stopPropagation();
    download("/api/conversations/" + encodeURIComponent(a.id) + "/export?format=md", a.id + ".md");
  });

  const del = document.createElement("button");
  del.type = "button";
  del.className = "conv-action-btn danger";
  del.title = "Supprimer";
  del.setAttribute("aria-label", "Supprimer la conversation");
  del.textContent = "\u00d7";
  del.addEventListener("click", (e) => {
    e.stopPropagation();
    remove(a.id, onChanged);
  });

  actions.appendChild(exp);
  actions.appendChild(del);
  row.appendChild(content);
  row.appendChild(actions);
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

  const exportBtn = document.getElementById("export-btn");
  if (exportBtn) {
    exportBtn.addEventListener("click", () => {
      download("/api/chat/export?format=md", "conversation.md");
    });
  }

  refresh();
}
