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
  const capsPanel = document.getElementById("caps-panel");
  const capsSave = document.getElementById("caps-save");
  const capsAdd = document.getElementById("caps-add");
  const capsAddInput = document.getElementById("caps-add-input");
  const capsStatus = document.getElementById("caps-status");
  let families = [];
  let caps = {};
  let capKeys = [];

  function selectTab(tab) {
    for (const b of overlay.querySelectorAll(".apikeys-tab")) {
      b.classList.toggle("active", b.dataset.tab === tab);
    }
    for (const p of overlay.querySelectorAll(".apikeys-panel")) {
      p.classList.toggle("active", p.id === "panel-" + tab);
    }
  }

  overlay.querySelectorAll(".apikeys-tab").forEach((b) => {
    b.addEventListener("click", () => selectTab(b.dataset.tab));
  });

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

  function capKeysFromFamilies() {
    const set = new Set(Object.keys(caps));
    for (const f of families) {
      for (const mode of f.modes || []) {
        for (const p of mode.pool || []) {
          if (p && p.provider && p.model) set.add(p.provider + "/" + p.model);
        }
      }
    }
    capKeys = Array.from(set).sort();
  }

  function renderCaps() {
    capsPanel.innerHTML = "";
    if (!capKeys.length) {
      const p = document.createElement("p");
      p.className = "settings-status";
      p.textContent = "Aucun modele dans les alias ; ajoute-en un ci-dessous.";
      capsPanel.appendChild(p);
      return;
    }
    for (const key of capKeys) {
      const c = caps[key] || {};
      const row = document.createElement("div");
      row.className = "caps-row";
      const name = document.createElement("span");
      name.className = "caps-name";
      name.textContent = key;
      row.appendChild(name);
      for (const cap of ["vision", "tts", "stt"]) {
        const wrap = document.createElement("label");
        wrap.className = "caps-check";
        const cb = document.createElement("input");
        cb.type = "checkbox";
        cb.checked = !!c[cap];
        cb.addEventListener("change", () => {
          const cur = caps[key] || {};
          cur[cap] = cb.checked;
          if (cur.vision || cur.tts || cur.stt) caps[key] = cur;
          else delete caps[key];
        });
        wrap.appendChild(cb);
        wrap.appendChild(document.createTextNode(cap));
        row.appendChild(wrap);
      }
      capsPanel.appendChild(row);
    }
  }

  async function loadCaps() {
    capsPanel.textContent = "Chargement...";
    try {
      const data = await api("/api/capabilities");
      caps = (data && data.caps) || {};
    } catch (e) {
      capsPanel.textContent = e.message;
      return;
    }
    capKeysFromFamilies();
    renderCaps();
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
    loadCaps();
  });
  mcpRefresh.addEventListener("click", () => loadMCP(true));

  capsAdd.addEventListener("click", () => {
    const raw = (capsAddInput.value || "").trim();
    const i = raw.indexOf("/");
    if (i <= 0 || i >= raw.length - 1) {
      capsStatus.textContent = "Format attendu : provider/model";
      return;
    }
    capsStatus.textContent = "";
    capsAddInput.value = "";
    caps[raw] = caps[raw] || {};
    capKeysFromFamilies();
    renderCaps();
  });

  capsSave.addEventListener("click", async () => {
    capsStatus.textContent = "Enregistrement...";
    try {
      await api("/api/capabilities", { method: "PUT", body: { caps } });
      capsStatus.textContent = "Enregistre.";
    } catch (e) {
      capsStatus.textContent = e.message;
    }
  });
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
