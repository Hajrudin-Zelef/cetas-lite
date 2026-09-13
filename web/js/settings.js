import { api } from "./api.js";

export function initSettings() {
  const overlay = document.getElementById("settings-overlay");
  const editor = document.getElementById("aliases-editor");
  const status = document.getElementById("aliases-status");
  const openBtn = document.getElementById("settings-btn");
  const closeBtn = document.getElementById("settings-close");
  const saveBtn = document.getElementById("aliases-save");
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

  openBtn.addEventListener("click", async () => {
    status.textContent = "";
    overlay.hidden = false;
    try {
      await load();
    } catch (e) {
      status.textContent = e.message;
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
    } catch (e) {
      status.textContent = e.message;
    }
  });
}
