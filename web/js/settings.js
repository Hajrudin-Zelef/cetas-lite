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
  let catalog = [];
  let draft = {};
  let activeFamily = "";
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

  function modelInfo(provider, model) {
    for (const p of catalog) {
      if (p.id !== provider) continue;
      for (const m of p.models) {
        if (m.id === model) return m;
      }
    }
    return null;
  }

  function hasCloudModes(f) {
    return (f.modes || []).some((m) => !m.local);
  }

  function buildDraft() {
    draft = {};
    for (const f of families) {
      if (f.local) continue;
      draft[f.id] = {};
      for (const m of f.modes || []) {
        if (m.local) continue;
        draft[f.id][m.mode] = (m.pool || []).map((p) => ({ provider: p.provider, model: p.model }));
      }
    }
    if (!draft[activeFamily]) {
      activeFamily = families.find((f) => !f.local && hasCloudModes(f))?.id || "";
    }
  }

  function modelLabel(provider, model) {
    const info = modelInfo(provider, model);
    return info ? info.label : model;
  }

  function modelPrice(provider, model) {
    const info = modelInfo(provider, model);
    if (!info) return "";
    if (!info.input_per_1m && !info.output_per_1m) return "gratuit";
    return "$" + info.input_per_1m + " / $" + info.output_per_1m;
  }

  function modeBlock(family, mode) {
    const list = draft[family.id][mode.mode] || [];
    const block = document.createElement("div");
    block.className = "mode-block";

    const head = document.createElement("div");
    head.className = "mode-head";
    const title = document.createElement("span");
    title.className = "mode-title";
    title.textContent = mode.label;
    head.appendChild(title);
    if (mode.agent) {
      const pill = document.createElement("span");
      pill.className = "agent-pill";
      pill.textContent = "Agent";
      head.appendChild(pill);
    }
    if (mode.rule) {
      const rule = document.createElement("span");
      rule.className = "mode-rule";
      rule.textContent = mode.rule;
      head.appendChild(rule);
    }
    block.appendChild(head);

    const ul = document.createElement("ol");
    ul.className = "pool-list";
    if (!list.length) {
      const empty = document.createElement("li");
      empty.className = "pool-empty";
      empty.textContent = "Aucun modele : ajoute-en un ci-dessous.";
      ul.appendChild(empty);
    }
    list.forEach((entry, idx) => {
      ul.appendChild(poolItem(family, mode, idx));
    });
    block.appendChild(ul);
    block.appendChild(addRow(family, mode));
    return block;
  }

  function poolItem(family, mode, idx) {
    const entry = draft[family.id][mode.mode][idx];
    const li = document.createElement("li");
    li.className = "pool-item";

    const rank = document.createElement("span");
    rank.className = "pool-rank";
    rank.textContent = "#" + (idx + 1);
    li.appendChild(rank);

    const name = document.createElement("span");
    name.className = "pool-name";
    name.textContent = modelLabel(entry.provider, entry.model);
    name.title = entry.provider + " " + entry.model;
    li.appendChild(name);

    const prov = document.createElement("span");
    prov.className = "pool-provider";
    prov.textContent = entry.provider;
    li.appendChild(prov);

    const price = modelPrice(entry.provider, entry.model);
    if (price) {
      const pr = document.createElement("span");
      pr.className = "pool-price";
      pr.textContent = price;
      li.appendChild(pr);
    }

    const actions = document.createElement("div");
    actions.className = "pool-actions";
    actions.appendChild(poolButton("↑", "Monter", idx > 0, () => move(family.id, mode.mode, idx, -1)));
    actions.appendChild(poolButton("↓", "Descendre", idx < draft[family.id][mode.mode].length - 1, () => move(family.id, mode.mode, idx, 1)));
    actions.appendChild(poolButton("×", "Retirer", true, () => removeEntry(family.id, mode.mode, idx), true));
    li.appendChild(actions);
    return li;
  }

  function poolButton(label, title, enabled, onClick, danger) {
    const b = document.createElement("button");
    b.type = "button";
    b.className = "pool-btn" + (danger ? " danger" : "");
    b.textContent = label;
    b.title = title;
    b.setAttribute("aria-label", title);
    b.disabled = !enabled;
    b.addEventListener("click", onClick);
    return b;
  }

  function move(familyId, modeId, idx, delta) {
    const list = draft[familyId][modeId];
    const next = idx + delta;
    if (next < 0 || next >= list.length) return;
    [list[idx], list[next]] = [list[next], list[idx]];
    render();
  }

  function removeEntry(familyId, modeId, idx) {
    draft[familyId][modeId].splice(idx, 1);
    render();
  }

  function addRow(family, mode) {
    const row = document.createElement("div");
    row.className = "pool-add";

    const select = document.createElement("select");
    select.className = "pool-model-select";
    const placeholder = document.createElement("option");
    placeholder.value = "";
    placeholder.textContent = "Choisir un modele...";
    select.appendChild(placeholder);
    for (const p of catalog) {
      const group = document.createElement("optgroup");
      group.label = p.label;
      for (const m of p.models) {
        const opt = document.createElement("option");
        opt.value = p.id + "|" + m.id;
        const price = !m.input_per_1m && !m.output_per_1m ? "gratuit" : "$" + m.input_per_1m + " / $" + m.output_per_1m;
        opt.textContent = m.label + "  ·  " + price;
        group.appendChild(opt);
      }
      select.appendChild(group);
    }
    row.appendChild(select);

    const add = document.createElement("button");
    add.type = "button";
    add.className = "small-btn";
    add.textContent = "Ajouter";
    add.addEventListener("click", () => {
      const value = select.value;
      if (!value) return;
      const i = value.indexOf("|");
      const provider = value.slice(0, i);
      const model = value.slice(i + 1);
      const list = draft[family.id][mode.mode];
      if (list.some((e) => e.provider === provider && e.model === model)) return;
      list.push({ provider, model });
      render();
    });
    row.appendChild(add);
    return row;
  }

  function render() {
    editor.innerHTML = "";
    const tabs = document.createElement("div");
    tabs.className = "fam-tabs";
    for (const f of families) {
      if (f.local || !hasCloudModes(f)) continue;
      const b = document.createElement("button");
      b.type = "button";
      b.className = "fam-tab" + (f.id === activeFamily ? " active" : "");
      b.textContent = f.label;
      b.addEventListener("click", () => {
        activeFamily = f.id;
        render();
      });
      tabs.appendChild(b);
    }
    editor.appendChild(tabs);

    const content = document.createElement("div");
    content.className = "provider-content";
    const fam = families.find((f) => f.id === activeFamily);
    if (!fam) {
      const p = document.createElement("p");
      p.className = "settings-status";
      p.textContent = "Aucune famille cloud configuree.";
      content.appendChild(p);
    } else {
      for (const m of fam.modes || []) {
        if (m.local) continue;
        content.appendChild(modeBlock(fam, m));
      }
    }
    editor.appendChild(content);
  }

  async function load() {
    const [aliasData, catalogData] = await Promise.all([
      api("/api/aliases"),
      api("/api/catalog").catch(() => ({ providers: [] })),
    ]);
    families = aliasData.families || [];
    catalog = catalogData.providers || [];
    buildDraft();
    render();
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
    for (const [familyId, modes] of Object.entries(draft)) {
      for (const [modeId, list] of Object.entries(modes)) {
        if (!list.length) continue;
        if (!ov[familyId]) ov[familyId] = {};
        ov[familyId][modeId] = list.map((e) => ({ provider: e.provider, model: e.model }));
      }
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
