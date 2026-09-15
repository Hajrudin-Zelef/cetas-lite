import { api, getPrefs, putPrefs } from "./api.js";

let families = [];
let thinkingPref = false;
let appMode = "chat"; // "chat" | "agent" : mode top-level choisi dans l'UI

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
  syncGlobeButton();
}
export function applyMCPToggle(on) {
  setCheck("mcp-toggle", on);
}
export function applyThinkingToggle(on) {
  setCheck("thinking-toggle", on);
  setCheck("plus-reflection-toggle", on);
  syncThinkingButton();
}

// Bouton globe du composer (a cote du +) : reflete l'etat du toggle web.
function syncGlobeButton() {
  const btn = document.getElementById("web-search-btn");
  if (!btn) return;
  const on = getCheck("web-toggle");
  btn.classList.toggle("active", on);
  btn.setAttribute("aria-pressed", on ? "true" : "false");
  btn.title = on ? "Recherche web : activée" : "Recherche web : désactivée";
}

// Bouton Thinking du composer (a cote du globe) : en mode Agent la reflexion
// est obligatoire et le bouton est verrouille.
function syncThinkingButton() {
  const btn = document.getElementById("thinking-toggle-btn");
  if (!btn) return;
  const locked = isAgentActive();
  const on = locked || getCheck("thinking-toggle");
  btn.classList.toggle("active", on);
  btn.classList.toggle("locked", locked);
  btn.setAttribute("aria-pressed", on ? "true" : "false");
  btn.title = locked ? "Réflexion (toujours active en mode Agent)" : on ? "Réflexion : activée" : "Réflexion : désactivée";
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

export function getAppMode() {
  return appMode;
}

// Capacité agent de la famille/mode sélectionnée (indépendant du mode top-level).
export function isFamilyAgent() {
  const familySel = document.getElementById("family-select");
  return !!(familySel && familySel.dataset.agent === "1");
}

// Agent réellement actif : mode "Agent" choisi ET famille compatible.
export function isAgentActive() {
  return appMode === "agent" && isFamilyAgent();
}

export function currentSelection() {
  const family = document.getElementById("family-select");
  const mode = document.getElementById("mode-select");
  const effortSel = document.getElementById("effort-select");
  const agent = isAgentActive();
  return {
    appMode,
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
  body.app_mode = appMode;
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
  const active = appMode === "agent" && agent;
  const pillRow = document.getElementById("agent-pill-row");
  if (pillRow) pillRow.style.display = active ? "" : "none";
  const th = document.getElementById("thinking-toggle");
  if (th) {
    if (active) {
      th.checked = true;
      th.disabled = true;
      th.title = "Réflexion (active en mode Agent)";
    } else {
      th.disabled = false;
      th.title = "Réflexion";
      th.checked = thinkingPref;
    }
  }
  syncThinkingButton();
  updateAppModeUI();
}

// Première famille/mode compatible agent (pour basculer automatiquement).
function firstAgentFamily() {
  for (const f of families) {
    const m = (f.modes || []).find((x) => x.agent);
    if (m) return { family: f.id, mode: m.mode };
  }
  return null;
}

// Met à jour l'UI en fonction du mode top-level (sans changer appMode).
function updateAppModeUI() {
  const chatBtn = document.getElementById("mode-chat-btn");
  const agentBtn = document.getElementById("mode-agent-btn");
  for (const [btn, mode] of [[chatBtn, "chat"], [agentBtn, "agent"]]) {
    if (!btn) continue;
    const on = appMode === mode;
    btn.classList.toggle("active", on);
    btn.setAttribute("aria-selected", on ? "true" : "false");
  }
  const isAgent = appMode === "agent";
  for (const id of ["rp-approvals-section", "rp-plan-section"]) {
    const sec = document.getElementById(id);
    if (sec) sec.style.display = isAgent ? "" : "none";
  }
  const input = document.getElementById("prompt-input");
  if (input) {
    input.placeholder = isAgent
      ? "Décrivez la tâche à accomplir…"
      : "Écrivez votre message…";
  }
  document.body.classList.toggle("app-mode-agent", isAgent);
  document.body.classList.toggle("app-mode-chat", !isAgent);
  updateAgentPillOnly();
}

function updateAgentPillOnly() {
  const pillRow = document.getElementById("agent-pill-row");
  if (pillRow) pillRow.style.display = isAgentActive() ? "" : "none";
}

export function setAppMode(mode, opts = {}) {
  if (mode !== "chat" && mode !== "agent") return;
  if (appMode === mode && !opts.force) {
    updateAppModeUI();
    return;
  }
  appMode = mode;
  // En mode Agent, basculer sur une famille compatible si besoin.
  if (mode === "agent" && !isFamilyAgent()) {
    const target = firstAgentFamily();
    if (target) {
      const familySel = document.getElementById("family-select");
      const modeSel = document.getElementById("mode-select");
      if (familySel) familySel.value = target.family;
      renderModes();
      if (modeSel) modeSel.value = target.mode;
    }
  }
  updateAgent();
  updateAppModeUI();
  if (opts.persist !== false) {
    putPrefs({ app_mode: appMode }).catch(() => {});
  }
  window.dispatchEvent(new CustomEvent("cetas:app-mode-changed", { detail: appMode }));
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
  // Slider "Tokens max par réponse" : réglage serveur (300..32768)
  initMaxTokensSlider();

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

  // Boutons globe / Thinking du composer (a cote du bouton +).
  const globeBtn = document.getElementById("web-search-btn");
  if (globeBtn) {
    globeBtn.addEventListener("click", () => {
      applyWebToggle(!getCheck("web-toggle"));
      persistPrefs().catch(() => {});
      window.dispatchEvent(new CustomEvent("cetas:composer-toggles"));
    });
  }
  const thinkBtn = document.getElementById("thinking-toggle-btn");
  if (thinkBtn) {
    thinkBtn.addEventListener("click", () => {
      if (isAgentActive()) return; // reflexion obligatoire en mode Agent
      setThinking(!getCheck("thinking-toggle"));
      persistPrefs().catch(() => {});
      window.dispatchEvent(new CustomEvent("cetas:composer-toggles"));
    });
  }

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

  // Sélecteur top-level Chat / Agent
  const chatBtn = document.getElementById("mode-chat-btn");
  const agentBtn = document.getElementById("mode-agent-btn");
  if (chatBtn) chatBtn.addEventListener("click", () => setAppMode("chat"));
  if (agentBtn) agentBtn.addEventListener("click", () => setAppMode("agent"));

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
    // Restaure le mode top-level Chat / Agent enregistré.
    setAppMode(prefs && prefs.app_mode === "agent" ? "agent" : "chat", { persist: false, force: true });
    renderPlusModelList();
  }

  await load();
  return load;
}

export function getFamilies() {
  return families;
}

// --- Slider "Tokens max par réponse" (menu +) ---
let maxTokensVal = 4096;
let maxTokensInit = false;

function fmtMaxTok(v) {
  return v >= 1000 ? (v / 1000).toFixed(v >= 10000 ? 0 : 1).replace(/\.0$/, "") + "K" : String(v);
}

export function getMaxTokens() {
  return maxTokensVal;
}

function initMaxTokensSlider() {
  const slider = document.getElementById("plus-max-tok-slider");
  const valEl = document.getElementById("plus-max-tok-val");
  if (!slider || maxTokensInit) return;
  maxTokensInit = true;
  const apply = (v) => {
    maxTokensVal = v;
    slider.value = String(v);
    if (valEl) valEl.textContent = fmtMaxTok(v);
  };
  getPrefs()
    .then((prefs) => {
      const v = prefs && prefs.max_tokens ? parseInt(prefs.max_tokens, 10) : 4096;
      apply(Math.min(32768, Math.max(300, isNaN(v) ? 4096 : v)));
    })
    .catch(() => apply(4096));
  let saveTimer = null;
  slider.addEventListener("input", () => {
    const v = parseInt(slider.value, 10) || 4096;
    maxTokensVal = v;
    if (valEl) valEl.textContent = fmtMaxTok(v);
    clearTimeout(saveTimer);
    saveTimer = setTimeout(() => putPrefs({ max_tokens: v }).catch(() => {}), 500);
  });
}
