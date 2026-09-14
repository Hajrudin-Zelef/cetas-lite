import { api, getPrefs, putPrefs } from "./api.js";

let families = [];
let thinkingPref = false;

function setCheck(id, on) {
  const el = document.getElementById(id);
  if (el) el.checked = !!on;
}
function getCheck(id) {
  const el = document.getElementById(id);
  return !!(el && el.checked);
}

export function applyWebToggle(on) {
  setCheck("web-toggle", on);
  setCheck("plus-websearch-toggle", on);
}
export function applyMCPToggle(on) {
  setCheck("mcp-toggle", on);
}
export function applyThinkingToggle(on) {
  setCheck("thinking-toggle", on);
  setCheck("plus-reflection-toggle", on);
}
export function setThinking(on) {
  thinkingPref = !!on;
  applyThinkingToggle(thinkingPref);
}
export function applyApproveToggle(on) {
  setCheck("approve-toggle", on);
}
export function applyPlanToggle(on) {
  setCheck("plan-toggle", on);
}

export function isAgentMode() {
  const familySel = document.getElementById("family-select");
  return !!(familySel && familySel.dataset.agent === "1");
}

export function currentSelection() {
  const family = document.getElementById("family-select");
  const mode = document.getElementById("mode-select");
  const effortSel = document.getElementById("effort-select");
  const agent = isAgentMode();
  return {
    family: family ? family.value : "",
    mode: mode ? mode.value : "",
    web: getCheck("web-toggle"),
    mcp: getCheck("mcp-toggle"),
    think: getCheck("thinking-toggle"),
    effort: effortSel ? effortSel.value : "default",
    approve: agent && getCheck("approve-toggle"),
    plan: agent && getCheck("plan-toggle"),
  };
}

export function persistPrefs() {
  const sel = currentSelection();
  const body = {
    theme: document.documentElement.dataset.theme || undefined,
    family: sel.family || undefined,
    mode: sel.mode || undefined,
    web_default: sel.web,
    thinking_default: sel.think,
    thinking_effort: sel.effort,
  };
  const mcpEl = document.getElementById("mcp-toggle");
  if (mcpEl && mcpEl.closest(".right-panel-section") && !mcpEl.disabled) {
    body.mcp_default = sel.mcp;
  }
  return putPrefs(body);
}

function modesFor(id) {
  const f = families.find((x) => x.id === id);
  return f ? f.modes : [];
}

function updateAgent() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  if (!familySel || !modeSel) return;
  const m = modesFor(familySel.value).find((x) => x.mode === modeSel.value);
  const agent = !!(m && m.agent);
  familySel.dataset.agent = agent ? "1" : "";
  const pillRow = document.getElementById("agent-pill-row");
  if (pillRow) pillRow.style.display = agent ? "" : "none";
  const th = document.getElementById("thinking-toggle");
  if (th) {
    if (agent) {
      th.checked = true;
      th.disabled = true;
      th.title = "Réflexion (active en mode Agent)";
    } else {
      th.disabled = false;
      th.title = "Réflexion";
      th.checked = thinkingPref;
    }
  }
}

function renderModes() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  if (!familySel || !modeSel) return;
  modeSel.innerHTML = "";
  for (const m of modesFor(familySel.value)) {
    const opt = document.createElement("option");
    opt.value = m.mode;
    opt.textContent = m.label + (m.rule ? " · " + m.rule : "");
    modeSel.appendChild(opt);
  }
  updateAgent();
}

function renderPlusModelList() {
  const list = document.getElementById("plus-model-list");
  if (!list) return;
  list.innerHTML = "";
  for (const f of families) {
    const group = document.createElement("div");
    group.className = "plus-model-group";
    const label = document.createElement("div");
    label.className = "plus-model-group-label";
    label.textContent = f.label;
    group.appendChild(label);
    for (const m of f.modes) {
      const btn = document.createElement("button");
      btn.type = "button";
      btn.className = "plus-model-option";
      btn.dataset.family = f.id;
      btn.dataset.mode = m.mode;
      const familySel = document.getElementById("family-select");
      const modeSel = document.getElementById("mode-select");
      if (familySel && modeSel && familySel.value === f.id && modeSel.value === m.mode) {
        btn.classList.add("active");
      }
      btn.innerHTML =
        '<span class="plus-model-option-name"></span>' +
        (m.agent ? '<span class="plus-model-option-badge">Agent</span>' : "");
      btn.querySelector(".plus-model-option-name").textContent = m.label;
      btn.title = f.label + " · " + m.label + (m.rule ? " — " + m.rule : "");
      btn.addEventListener("click", () => {
        if (familySel) familySel.value = f.id;
        renderModes();
        if (modeSel) modeSel.value = m.mode;
        updateAgent();
        persistPrefs().catch(() => {});
        renderPlusModelList();
      });
      group.appendChild(btn);
    }
    list.appendChild(group);
  }
}

export function applyTheme(theme) {
  if (!["ocean", "sombre", "clair"].includes(theme)) return;
  document.documentElement.dataset.theme = theme;
  document.body.className =
    theme === "sombre" ? "ocean-theme dark" : theme === "ocean" ? "ocean-theme" : "";
  const sel = document.getElementById("cfg-theme");
  if (sel) sel.value = theme;
  try {
    localStorage.setItem("cetas-lite-theme", theme);
  } catch (e) {}
}

export async function initModels() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");

  // Onglets du plus-menu : seul "Texte" est pertinent pour Cetas Lite
  document.querySelectorAll("#plus-model-tabs .plus-model-tab").forEach((t) => {
    if (t.dataset.tab !== "text") t.style.display = "none";
  });
  // Pas de réglage de tokens max côté serveur : masquer la section
  const maxTokSection = document.getElementById("plus-max-tok-slider");
  if (maxTokSection) {
    const sec = maxTokSection.closest(".plus-menu-section");
    if (sec) {
      sec.style.display = "none";
      const sep = sec.nextElementSibling;
      if (sep && sep.classList.contains("plus-menu-sep")) sep.style.display = "none";
    }
  }

  function bindCheck(id, apply, persist = true) {
    const elc = document.getElementById(id);
    if (elc) {
      elc.addEventListener("change", () => {
        apply(elc.checked);
        if (id === "thinking-toggle" || id === "plus-reflection-toggle") thinkingPref = elc.checked;
        if (persist) persistPrefs().catch(() => {});
      });
    }
  }
  bindCheck("web-toggle", applyWebToggle);
  bindCheck("plus-websearch-toggle", applyWebToggle);
  bindCheck("mcp-toggle", applyMCPToggle);
  bindCheck("thinking-toggle", applyThinkingToggle);
  bindCheck("plus-reflection-toggle", applyThinkingToggle);
  bindCheck("approve-toggle", applyApproveToggle, false);
  bindCheck("plan-toggle", applyPlanToggle, false);

  // Pills d'effort du plus-menu
  document.querySelectorAll("#plus-effort-pills .plus-menu-pill").forEach((p) => {
    p.addEventListener("click", () => {
      document.querySelectorAll("#plus-effort-pills .plus-menu-pill").forEach((x) => x.classList.remove("active"));
      p.classList.add("active");
      const effortSel = document.getElementById("effort-select");
      if (effortSel) effortSel.value = p.dataset.effort;
      persistPrefs().catch(() => {});
    });
  });
  const effortSel = document.getElementById("effort-select");
  if (effortSel) {
    effortSel.addEventListener("change", () => {
      document.querySelectorAll("#plus-effort-pills .plus-menu-pill").forEach((x) =>
        x.classList.toggle("active", x.dataset.effort === effortSel.value)
      );
      persistPrefs().catch(() => {});
    });
  }

  if (familySel) familySel.addEventListener("change", () => { renderModes(); persistPrefs().catch(() => {}); renderPlusModelList(); });
  if (modeSel) modeSel.addEventListener("change", () => { updateAgent(); persistPrefs().catch(() => {}); renderPlusModelList(); });

  async function load() {
    const data = await api("/api/aliases");
    families = data.families || [];
    const prefs = await getPrefs().catch(() => null);
    if (prefs && prefs.theme) applyTheme(prefs.theme);
    applyWebToggle(!!(prefs && prefs.web_default));
    thinkingPref = !!(prefs && prefs.thinking_default);
    applyThinkingToggle(thinkingPref);
    if (effortSel) {
      effortSel.value = (prefs && prefs.thinking_effort) || "default";
      document.querySelectorAll("#plus-effort-pills .plus-menu-pill").forEach((x) =>
        x.classList.toggle("active", x.dataset.effort === effortSel.value)
      );
    }

    const mcp = await api("/api/mcp").catch(() => null);
    const mcpSection = document.getElementById("rp-mcp-section");
    const mcpLine = document.getElementById("mcp-status-line");
    const servers = (mcp && mcp.servers) || [];
    if (mcpSection) {
      if (servers.length > 0) {
        mcpSection.style.display = "";
        applyMCPToggle(prefs && typeof prefs.mcp_default === "boolean" ? prefs.mcp_default : true);
        if (mcpLine) {
          mcpLine.style.display = "";
          mcpLine.textContent = servers.length + " serveur(s) : " + servers.map((s) => s.name || s.id).join(", ");
        }
      } else {
        mcpSection.style.display = "none";
        applyMCPToggle(false);
      }
    }

    const prevF = (prefs && prefs.family) || "";
    const prevM = (prefs && prefs.mode) || "";
    if (familySel) {
      familySel.innerHTML = "";
      for (const f of families) {
        const opt = document.createElement("option");
        opt.value = f.id;
        opt.textContent = f.label;
        familySel.appendChild(opt);
      }
      if (prevF && families.some((f) => f.id === prevF)) familySel.value = prevF;
    }
    renderModes();
    if (modeSel && prevM && modesFor(familySel.value).some((m) => m.mode === prevM)) {
      modeSel.value = prevM;
      updateAgent();
    }
    renderPlusModelList();
  }

  await load();
  return load;
}

export function getFamilies() {
  return families;
}
