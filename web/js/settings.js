import { api } from "./api.js";

export function initSettings({ reloadModels } = {}) {
  const overlay = document.getElementById("settings-overlay");
  const editor = document.getElementById("aliases-editor");
  const status = document.getElementById("aliases-status");
  const openBtn = document.getElementById("settings-btn");
  const closeBtn = document.getElementById("settings-close");
  const saveBtn = document.getElementById("aliases-save");
  const mcpPanel = document.getElementById("mcp-panel");
  const mcpRefresh = document.getElementById("mcp-refresh");
  let families = [];

  function render() {
    editor.innerHTML = "";
    for (const f of families) {
      const modes = (f.modes || []).filter((m) => !m.local && (m.pool || []).length > 0);
      if (!modes.length) continue;
      const group = document.createElement("div");
      group.className = "alias-group";
      const h = document.createElement("h3");
      h.textContent = f.label;
      group.appendChild(h);
      for (const m of modes) {
        const wrap = document.createElement("div");
        wrap.className = "alias-mode";
        const label = document.createElement("label");
        label.textContent = m.label;
        const ta = document.createElement("textarea");
        ta.dataset.family = f.id;
        ta.dataset.mode = m.mode;
        ta.value = (m.pool || []).map((p) => p.provider + " " + p.model).join("\n");
        wrap.appendChild(label);
        wrap.appendChild(ta);
        group.appendChild(wrap);
      }
      editor.appendChild(group);
    }
  }

  async function load() {
    const data = await api("/api/aliases");
    families = data.families || [];
    render();
  }

  function parsePool(value) {
    const pool = [];
    for (const raw of value.split("\n")) {
      const line = raw.trim();
      if (!line) continue;
      const i = line.indexOf(" ");
      if (i < 0) continue;
      const provider = line.slice(0, i).trim();
      const model = line.slice(i + 1).trim();
      if (provider && model) pool.push({ provider, model });
    }
    return pool;
  }

  function renderMCP(servers) {
    mcpPanel.innerHTML = "";
    if (!servers.length) {
      const p = document.createElement("p");
      p.className = "settings-status";
      p.textContent = "Aucun serveur MCP configure ($CETAS_LITE_HOME/mcp.json).";
      mcpPanel.appendChild(p);
      return;
    }
    for (const s of servers) {
      const row = document.createElement("div");
      row.className = "mcp-row";
      const name = document.createElement("span");
      name.className = "mcp-name";
      name.textContent = s.name || "";
      const meta = document.createElement("span");
      meta.className = "mcp-meta";
      const state = s.connected ? "connecte" : s.error ? "erreur" : "hors ligne";
      meta.textContent =
        (s.transport || "") + " · " + state + " · " + (s.tools || 0) + " outil(s)" + (s.error ? " · " + s.error : "");
      row.appendChild(name);
      row.appendChild(meta);
      mcpPanel.appendChild(row);
    }
  }

  async function loadMCP(probe) {
    mcpPanel.textContent = "Chargement...";
    try {
      const data = await api("/api/mcp" + (probe ? "?probe=1" : ""));
      renderMCP((data && data.servers) || []);
    } catch (e) {
      mcpPanel.textContent = e.message;
    }
  }

  openBtn.addEventListener("click", async () => {
    status.textContent = "";
    overlay.hidden = false;
    try {
      await load();
    } catch (e) {
      status.textContent = e.message;
    }
    loadMCP(false);
  });
  mcpRefresh.addEventListener("click", () => loadMCP(true));
  closeBtn.addEventListener("click", () => {
    overlay.hidden = true;
  });
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) overlay.hidden = true;
  });

  saveBtn.addEventListener("click", async () => {
    const ov = {};
    for (const ta of editor.querySelectorAll("textarea")) {
      const pool = parsePool(ta.value);
      if (!pool.length) continue;
      if (!ov[ta.dataset.family]) ov[ta.dataset.family] = {};
      ov[ta.dataset.family][ta.dataset.mode] = pool;
    }
    try {
      await api("/api/aliases", { method: "PUT", body: ov });
      status.textContent = "Enregistre.";
      await load();
      if (reloadModels) await reloadModels();
    } catch (e) {
      status.textContent = e.message;
    }
  });
}
