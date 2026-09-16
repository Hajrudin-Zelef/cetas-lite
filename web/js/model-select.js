import { api, getPrefs, putPrefs } from "./api.js";
import { modeSubtitle, familyShortLabel } from "./model-selector.js";

let families = [];
let thinkingPref = false;

// --- Préférences Fonctionnalités (panneau Configuration) ---
// Cache local des choix tts / transcription / prompt_enhance / summarizer /
// title_gen / error_analysis. Rafraîchi au chargement des prefs et sur
// l'événement "cetas:features-changed" émis par le panneau.
const featurePrefs = {};
export function getFeaturePref(key, fallback) {
  const v = featurePrefs[key];
  return v === undefined ? fallback : v;
}
function refreshFeaturePrefs(prefs) {
  if (!prefs) return;
  for (const k of ["tts", "transcription", "prompt_enhance", "summarizer", "title_gen", "error_analysis"]) {
    if (prefs[k] !== undefined && prefs[k] !== "") featurePrefs[k] = prefs[k];
  }
}
if (typeof window !== "undefined") {
  window.addEventListener("cetas:features-changed", () => {
    getPrefs().then(refreshFeaturePrefs).catch(() => {});
  });
}

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

// Bouton Thinking du composer (a cote du globe) : reflete l'etat du toggle.
function syncThinkingButton() {
  const btn = document.getElementById("thinking-toggle-btn");
  if (!btn) return;
  const on = getCheck("thinking-toggle");
  btn.classList.toggle("active", on);
  btn.setAttribute("aria-pressed", on ? "true" : "false");
  btn.title = on ? "Réflexion : activée" : "Réflexion : désactivée";
}
export function setThinking(on) {
  thinkingPref = !!on;
  applyThinkingToggle(thinkingPref);
}

export function currentSelection() {
  const family = document.getElementById("family-select");
  const mode = document.getElementById("mode-select");
  const effortSel = document.getElementById("effort-select");
  return {
    family: family ? family.value : "",
    mode: mode ? mode.value : "",
    web: getCheck("web-toggle"),
    mcp: getCheck("mcp-toggle"),
    think: getCheck("thinking-toggle"),
    effort: effortSel ? effortSel.value : "default",
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

// Options de modes pour une famille : "Auto (fallback)" en tête pour les
// familles cloud (union des pools de l'alias, résolue côté moteur), puis les
// modes configurés. SamGen (local) n'a pas d'Auto : on choisit la route.
export function modeOptionsFor(fam) {
  const modes = fam ? fam.modes || [] : [];
  if (fam && !fam.local) return [{ mode: "auto", label: "Auto (fallback)", auto: true }, ...modes];
  return modes;
}

// Familles affichées dans le menu + : Nano, N4, N8, SamGen. Pas Code :
// l'agent vit désormais dans son propre module.
const PLUS_MENU_ORDER = ["samagent-nano", "samagent-n4", "samagent-n8", "samgen"];
export function plusMenuFamilies() {
  const byId = new Map(families.map((f) => [f.id, f]));
  return PLUS_MENU_ORDER.map((id) => byId.get(id)).filter(Boolean);
}

function renderModes() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  if (!familySel || !modeSel) return;
  const fam = families.find((x) => x.id === familySel.value);
  modeSel.innerHTML = "";
  for (const m of modeOptionsFor(fam)) {
    const opt = document.createElement("option");
    opt.value = m.mode;
    if (m.auto) {
      opt.textContent = m.label;
    } else {
      const sub = fam ? modeSubtitle(fam, m) : "";
      opt.textContent = m.label + (sub ? " · " + sub : "");
    }
    modeSel.appendChild(opt);
  }
}

function plusOptionRow(f, modeId, name, sub) {
  const btn = document.createElement("button");
  btn.type = "button";
  btn.className = "plus-model-option";
  btn.dataset.family = f.id;
  btn.dataset.mode = modeId;
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  if (familySel && modeSel && familySel.value === f.id && modeSel.value === modeId) {
    btn.classList.add("active");
  }
  btn.innerHTML =
    '<span class="plus-model-option-main">' +
    '<span class="plus-model-option-name"></span>' +
    (sub ? '<span class="plus-model-option-rule"></span>' : "") +
    "</span>";
  btn.querySelector(".plus-model-option-name").textContent = name;
  const ruleEl = btn.querySelector(".plus-model-option-rule");
  if (ruleEl) ruleEl.textContent = sub;
  btn.title = familyShortLabel(f) + " · " + name + (sub ? " — " + sub : "");
  btn.addEventListener("click", () => {
    if (familySel) familySel.value = f.id;
    renderModes();
    if (modeSel) modeSel.value = modeId;
    persistPrefs().catch(() => {});
    renderPlusModelList();
  });
  return btn;
}

function renderPlusModelList() {
  const list = document.getElementById("plus-model-list");
  if (!list) return;
  list.innerHTML = "";
  for (const f of plusMenuFamilies()) {
    const group = document.createElement("div");
    group.className = "plus-model-group";
    const label = document.createElement("div");
    label.className = "plus-model-group-label";
    label.textContent = familyShortLabel(f);
    group.appendChild(label);
    // "Auto (fallback)" : tout l'alias en fallback (familles cloud).
    if (!f.local) group.appendChild(plusOptionRow(f, "auto", "Auto (fallback)", ""));
    const modes = f.modes || [];
    for (const m of modes) {
      // Famille à un seul mode (Nano) : "Modèle — [sélection effective]".
      const name = modes.length === 1 ? "Modèle" : m.label;
      group.appendChild(plusOptionRow(f, m.mode, name, modeSubtitle(f, m)));
    }
    list.appendChild(group);
  }
}

export const THEME_PALETTES = [
  { id: "bleu", label: "Bleu", color: "#3b87ce" },
  { id: "violet", label: "Violet", color: "#8b5cf6" },
  { id: "vert", label: "Vert", color: "#22c55e" },
  { id: "vert_pur", label: "Vert pur", color: "#00e676" },
  { id: "bleu_ocean", label: "Bleu océan", color: "#0ea5e9" },
  { id: "jaune_or", label: "Jaune or", color: "#d4a017" },
  { id: "rouge", label: "Rouge", color: "#ef4444" },
];

export function applyPalette(palette) {
  const ids = THEME_PALETTES.map((p) => p.id);
  const p = ids.includes(palette) ? palette : "bleu";
  document.documentElement.dataset.palette = p;
  try {
    localStorage.setItem("cetas-lite-palette", p);
  } catch (e) {}
  const name = document.querySelector(".palette-name");
  const found = THEME_PALETTES.find((x) => x.id === p);
  if (name && found) name.textContent = found.label;
  document.querySelectorAll(".palette-swatch").forEach((s) => {
    s.classList.toggle("active", s.dataset.palette === p);
  });
}

export function applyTheme(theme) {
  const t = theme === "ocean" ? "clair" : theme; // ancien nom
  if (!["sombre", "hard_dark", "clair"].includes(t)) return;
  document.documentElement.dataset.theme = t;
  document.body.className = t === "sombre" ? "dark" : t === "hard_dark" ? "dark hard-dark" : "";
  const sel = document.getElementById("cfg-theme");
  if (sel) sel.value = t;
  try {
    localStorage.setItem("cetas-lite-theme", t);
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
  if (modeSel) modeSel.addEventListener("change", () => { persistPrefs().catch(() => {}); renderPlusModelList(); });

  async function load() {
    const data = await api("/api/aliases");
    families = data.families || [];
    const prefs = await getPrefs().catch(() => null);
    refreshFeaturePrefs(prefs);
    if (prefs && prefs.theme) applyTheme(prefs.theme);
    if (prefs && prefs.palette) applyPalette(prefs.palette);
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
    if (modeSel && prevM && [...modeSel.options].some((o) => o.value === prevM)) {
      modeSel.value = prevM;
        }
    renderPlusModelList();
  }

  await load();
  return load;
}

export function getFamilies() {
  return families;
}

// Recharge les alias (apres une sauvegarde du selecteur de modeles)
// et rafraichit les selects + le menu +.
export async function refreshFamilies() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  const prevF = familySel ? familySel.value : "";
  const prevM = modeSel ? modeSel.value : "";
  try {
    const data = await api("/api/aliases");
    families = data.families || [];
  } catch (e) {
    return;
  }
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
  if (modeSel && prevM && [...modeSel.options].some((o) => o.value === prevM)) {
    modeSel.value = prevM;
  }
  renderPlusModelList();
}

if (typeof window !== "undefined") {
  window.addEventListener("cetas:aliases-changed", () => {
    refreshFamilies().catch(() => {});
  });
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
