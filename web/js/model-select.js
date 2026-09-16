import { api, getPrefs, putPrefs } from "./api.js";
import { modeSubtitle, familyShortLabel, buildStaged } from "./model-selector.js";

let families = [];
let aliasDefaults = [];
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

// ===== Menu + : cascade multi-niveaux =====
// Famille › → [Auto (= fallback de l'alias), Mode › → …]
// Mode › → [Auto (= fallback du mode), Models › → modèles] — le fallback,
//           c'est le mode Auto ; en dessous, sélection manuelle du modèle.
// Nano (mode unique Free) : Free › → groupes par fournisseur
// (« Free de OpenRouter », « Zen Free »), chacun en cascade.
// SamGen (local) : routes directes, sans Auto.
//
// Un item : { label, sub, active, onPick } (feuille) ou
// { label, sub, active, children } (branche → sous-menu au niveau suivant).

// Groupes du mode Free de Nano : les gratuits OpenRouter, puis les
// gratuits OpenCode (aliasés « Zen Free »).
const NANO_FREE_GROUPS = [
  { provider: "openrouter", label: "Free de OpenRouter" },
  { provider: "opencode", label: "Zen Free" },
];

function effPoolOf(famId, modeId) {
  const f = families.find((x) => x.id === famId);
  const m = f && (f.modes || []).find((x) => x.mode === modeId);
  return (m && m.pool) || [];
}

// Tous les modèles proposables pour un mode : pool par défaut + effectifs
// (dédupliqués). Nécessaire car le pool effectif peut être réduit à
// « 1 modèle » alors que l'utilisateur veut choisir parmi tous.
function allModeModels(f, mode) {
  const def = (aliasDefaults.find((d) => d.id === f.id) || {}).modes || [];
  const dm = def.find((x) => x.mode === mode.mode);
  const seen = new Set();
  const out = [];
  for (const p of [...((dm && dm.pool) || []), ...(mode.pool || [])]) {
    const k = p.provider + "/" + p.model;
    if (!seen.has(k)) {
      seen.add(k);
      out.push(p);
    }
  }
  return out;
}

function pickFamilyMode(f, modeId) {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  if (familySel) familySel.value = f.id;
  renderModes();
  if (modeSel) modeSel.value = modeId;
  persistPrefs().catch(() => {});
  renderPlusModelList(); // referme aussi les sous-menus
  // Le hint du composer (chat.js) se rafraîchit sur cet événement.
  window.dispatchEvent(new CustomEvent("cetas:model-changed"));
}

// Persiste l'override d'un mode. Le PUT /api/aliases REMPLACE tous les
// overrides : on reconstruit donc l'ensemble depuis le serveur (GET +
// buildStaged) avant d'appliquer le changement, pour ne jamais écraser
// les réglages des autres modes.
async function persistModeOverride(famId, modeId, members) {
  const data = await api("/api/aliases");
  const staged = buildStaged(data.families || [], data.defaults || []);
  if (!staged[famId]) staged[famId] = {};
  staged[famId][modeId] = members.map((m) => ({ provider: m.provider, model: m.model }));
  const updated = await api("/api/aliases", { method: "PUT", body: staged });
  families = (updated && updated.families) || families;
  // renderModes() reconstruit les options : préserver la sélection courante.
  const sel = currentSelection();
  renderModes();
  const modeSel = document.getElementById("mode-select");
  if (modeSel && sel.mode && [...modeSel.options].some((o) => o.value === sel.mode)) {
    modeSel.value = sel.mode;
  }
  renderPlusModelList();
  // "cetas:model-changed" est émis par pickFamilyMode (sélection appliquée
  // avant la persistance) : pas de double émission ici.
  window.dispatchEvent(new CustomEvent("cetas:aliases-changed", { detail: { source: "plus" } }));
}

// « Auto » d'un mode : tout le pool en fallback.
// La sélection est appliquée d'abord (feedback immédiat), la persistance
// suit : le rafraîchissement async déclenché par "cetas:aliases-changed"
// fige sinon les anciennes valeurs et écrase la sélection.
async function pickModeAuto(f, mode) {
  const members = allModeModels(f, mode);
  if (!members.length) return;
  pickFamilyMode(f, mode.mode);
  try {
    await persistModeOverride(f.id, mode.mode, members);
  } catch (e) {
    // L'override n'a pas été persisté ; la sélection reste affichée.
  }
}

// Choix manuel d'un modèle : le mode passe en « 1 modèle » sur celui-ci.
async function pickModel(f, mode, m) {
  pickFamilyMode(f, mode.mode);
  try {
    await persistModeOverride(f.id, mode.mode, [m]);
  } catch (e) {
    // L'override n'a pas été persisté ; la sélection reste affichée.
  }
}

function modelLeafItems(f, mode, models) {
  const sel = currentSelection();
  const pool = effPoolOf(f.id, mode.mode);
  return models.map((m) => ({
    label: m.label || m.model,
    sub: m.provider,
    active:
      sel.family === f.id &&
      sel.mode === mode.mode &&
      pool.length === 1 &&
      pool[0].provider + "/" + pool[0].model === m.provider + "/" + m.model,
    onPick: () => pickModel(f, mode, m),
  }));
}

// Contenu du sous-menu d'un mode : « Auto » (= le fallback) puis la
// sélection manuelle (groupes par fournisseur pour Nano/Free, « Models »
// sinon).
function modeChildren(f, mode) {
  const sel = currentSelection();
  const pool = effPoolOf(f.id, mode.mode);
  const models = allModeModels(f, mode);
  const items = [];
  const isNanoFree = f.id === "samagent-nano";
  if (!isNanoFree) {
    items.push({
      label: "Auto",
      sub: "",
      active: sel.family === f.id && sel.mode === mode.mode && pool.length > 1,
      onPick: () => pickModeAuto(f, mode),
    });
  }
  if (isNanoFree) {
    for (const g of NANO_FREE_GROUPS) {
      const gm = models.filter((m) => m.provider === g.provider);
      if (!gm.length) continue;
      items.push({ label: g.label, sub: "", children: modelLeafItems(f, mode, gm) });
    }
    const rest = models.filter((m) => !NANO_FREE_GROUPS.some((g) => g.provider === m.provider));
    if (rest.length) items.push({ label: "Autres", sub: "", children: modelLeafItems(f, mode, rest) });
  } else if (models.length) {
    items.push({ label: "Models", sub: "", children: modelLeafItems(f, mode, models) });
  }
  return items;
}

// Contenu du sous-menu d'une famille : « Auto » (= fallback de l'alias)
// puis un rang par mode (le libellé du mode : « Free » pour Nano).
function familyChildren(f) {
  const sel = currentSelection();
  const items = [];
  if (!f.local) {
    items.push({
      label: "Auto",
      sub: "",
      active: sel.family === f.id && sel.mode === "auto",
      onPick: () => pickFamilyMode(f, "auto"),
    });
  }
  for (const m of f.modes || []) {
    const isSel = sel.family === f.id && sel.mode === m.mode;
    if (f.local) {
      items.push({ label: m.label, sub: "", active: isSel, onPick: () => pickFamilyMode(f, m.mode) });
    } else {
      items.push({
        label: m.label,
        sub: modeSubtitle(f, m),
        active: isSel,
        children: modeChildren(f, m),
      });
    }
  }
  return items;
}

// --- Pile de sous-menus (un par niveau de cascade) ---
// Chaque niveau est un élément body-level en position:fixed : le dropdown
// parent a overflow-y:auto, un absolute y serait rogné.
const submenuLevels = []; // [{ el, timer, anchor }]

function submenuLevel(i) {
  while (submenuLevels.length <= i) {
    const entry = { el: null, timer: null, anchor: null };
    const el = document.createElement("div");
    el.className = "plus-model-submenu";
    el.style.display = "none";
    document.body.appendChild(el);
    el.addEventListener("mouseenter", () => clearTimeout(entry.timer));
    el.addEventListener("mouseleave", () => scheduleCloseFrom(submenuLevels.indexOf(entry)));
    entry.el = el;
    submenuLevels.push(entry);
  }
  const e = submenuLevels[i];
  if (!e.el.isConnected) document.body.appendChild(e.el);
  return e;
}

function scheduleCloseFrom(level) {
  for (let i = level; i < submenuLevels.length; i++) {
    const e = submenuLevels[i];
    clearTimeout(e.timer);
    e.timer = setTimeout(() => {
      e.el.style.display = "none";
      if (e.anchor) e.anchor.classList.remove("open");
      e.anchor = null;
    }, 140);
  }
}

function closeSubmenusImmediate(from) {
  for (let i = from; i < submenuLevels.length; i++) {
    const e = submenuLevels[i];
    clearTimeout(e.timer);
    e.timer = null;
    e.el.style.display = "none";
    if (e.anchor) e.anchor.classList.remove("open");
    e.anchor = null;
  }
}

export function closePlusModelSubmenu() {
  closeSubmenusImmediate(0);
}

export function plusModelSubmenuContains(el) {
  return !!el && submenuLevels.some((e) => e.el.contains(el));
}

function positionSubmenu(el, anchor) {
  el.style.visibility = "hidden";
  el.style.display = "block";
  const r = anchor.getBoundingClientRect();
  const w = el.offsetWidth || 244;
  const h = el.offsetHeight || 120;
  let left = r.right + 8;
  if (left + w > window.innerWidth - 8) left = Math.max(8, r.left - 8 - w);
  let top = Math.max(8, r.top - 6);
  if (top + h > window.innerHeight - 8) top = Math.max(8, window.innerHeight - h - 8);
  el.style.left = left + "px";
  el.style.top = top + "px";
  el.style.visibility = "";
}

function touchOnly() {
  return !!(window.matchMedia && window.matchMedia("(hover: none)").matches);
}

function openSubmenu(level, items, anchor) {
  if (!items || !items.length) return;
  // Ferme les niveaux plus profonds, garde les parents.
  closeSubmenusImmediate(level + 1);
  const entry = submenuLevel(level);
  clearTimeout(entry.timer);
  entry.timer = null;
  // Bascule tactile : même ancrage déjà ouvert → on referme.
  if (entry.anchor === anchor && entry.el.style.display !== "none") {
    closeSubmenusImmediate(level);
    return;
  }
  if (entry.anchor) entry.anchor.classList.remove("open");
  entry.anchor = anchor;
  anchor.classList.add("open");
  fillSubmenu(entry.el, items, level);
  positionSubmenu(entry.el, anchor);
}

function fillSubmenu(el, items, level) {
  el.innerHTML = "";
  for (const it of items) {
    const btn = document.createElement("button");
    btn.type = "button";
    btn.className = "plus-model-option";
    if (it.active) btn.classList.add("active");
    btn.innerHTML =
      '<span class="plus-model-option-main">' +
      '<span class="plus-model-option-name"></span>' +
      (it.sub ? '<span class="plus-model-option-rule"></span>' : "") +
      "</span>" +
      (it.children ? '<span class="plus-model-family-chev">›</span>' : "");
    btn.querySelector(".plus-model-option-name").textContent = it.label;
    const ruleEl = btn.querySelector(".plus-model-option-rule");
    if (ruleEl) ruleEl.textContent = it.sub;
    btn.title = it.label + (it.sub ? " — " + it.sub : "");
    if (it.children) {
      btn.addEventListener("mouseenter", () => openSubmenu(level + 1, it.children, btn));
      btn.addEventListener("mouseleave", () => scheduleCloseFrom(level + 1));
      btn.addEventListener("focus", () => openSubmenu(level + 1, it.children, btn));
      // Tactile : le survol n'existe pas, la tape bascule le sous-menu.
      if (touchOnly()) btn.addEventListener("click", () => openSubmenu(level + 1, it.children, btn));
    } else if (it.onPick) {
      btn.addEventListener("click", () => it.onPick());
    }
    el.appendChild(btn);
  }
}

function renderPlusModelList() {
  const list = document.getElementById("plus-model-list");
  if (!list) return;
  closePlusModelSubmenu();
  list.innerHTML = "";
  const sel = currentSelection();
  for (const f of plusMenuFamilies()) {
    const wrap = document.createElement("div");
    wrap.className = "plus-model-family";
    const row = document.createElement("button");
    row.type = "button";
    row.className = "plus-model-family-row";
    row.dataset.family = f.id;
    if (sel.family === f.id) row.classList.add("active");
    row.innerHTML =
      '<span class="plus-model-family-name"></span>' +
      '<span class="plus-model-family-chev">›</span>';
    row.querySelector(".plus-model-family-name").textContent = familyShortLabel(f);
    row.title = familyShortLabel(f);
    row.addEventListener("mouseenter", () => openSubmenu(0, familyChildren(f), row));
    row.addEventListener("mouseleave", () => scheduleCloseFrom(0));
    row.addEventListener("focus", () => openSubmenu(0, familyChildren(f), row));
    // Tactile : le survol n'existe pas, la tape bascule le sous-menu.
    if (touchOnly()) row.addEventListener("click", () => openSubmenu(0, familyChildren(f), row));
    wrap.appendChild(row);
    list.appendChild(wrap);
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
    aliasDefaults = data.defaults || [];
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
    aliasDefaults = data.defaults || [];
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
  window.addEventListener("cetas:aliases-changed", (e) => {
    // Notre propre sauvegarde (source: "plus") a déjà rafraîchi depuis la
    // réponse PUT : éviter un second GET + un écrasement de sélection.
    if (e.detail && e.detail.source === "plus") return;
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
