// Panneaux de la Configuration : Fonctionnalités, Recherche Web, Apparence.
// Module dédié (dépendances légères) pour rester testable sous
// jsdom sans charger tout modals.js.
import { api, getPrefs, putPrefs } from "./api.js";
import { applyPalette, applyTheme, THEME_PALETTES } from "./model-select.js";
import { confirmDialog } from "./dialogs.js";

// --- Sauvegarde par module + sauvegarde générale ---
// Les panneaux « stagent » leurs modifications au lieu de sauvegarder
// immédiatement. Chaque panneau enregistre un ou plusieurs
// { isDirty(), save(), revert?() } sous son onglet ; la barre de sauvegarde
// générale (pied de la modale) et les boutons « Enregistrer » par onglet
// déclenchent la persistance.
const panelSavers = new Map(); // tab -> Array<{ isDirty, save, revert? }>

export function registerPanelSaver(tab, saver) {
  if (!panelSavers.has(tab)) panelSavers.set(tab, []);
  panelSavers.get(tab).push(saver);
  refreshConfigSavebar();
}
export function notifyConfigDirty() {
  refreshConfigSavebar();
}
function saversDirty(tab) {
  const list = panelSavers.get(tab) || [];
  return list.filter((s) => {
    try {
      return s.isDirty();
    } catch (e) {
      return false;
    }
  });
}
export function isConfigDirty() {
  for (const tab of panelSavers.keys()) if (saversDirty(tab).length) return true;
  return false;
}
export async function saveConfigPanel(tab) {
  for (const s of saversDirty(tab)) await s.save();
  refreshConfigSavebar();
  flashConfigSaved();
}
export async function saveAllConfigPanels() {
  for (const tab of panelSavers.keys()) await saveConfigPanel(tab);
}
// Repli visuel si la modale est fermée sans sauvegarde (ex. aperçu du thème).
export function revertUnsavedConfig() {
  for (const list of panelSavers.values()) {
    for (const s of list) {
      try {
        if (s.isDirty() && typeof s.revert === "function") s.revert();
      } catch (e) {}
    }
  }
  refreshConfigSavebar();
}
function refreshConfigSavebar() {
  const bar = document.getElementById("config-savebar");
  if (!bar) return;
  const hint = document.getElementById("config-savebar-hint");
  if (hint) {
    hint.textContent = "Modifications non enregistrées";
    hint.classList.remove("saved");
  }
  bar.style.display = isConfigDirty() ? "" : "none";
}
let savedFlashTimer = null;
function flashConfigSaved() {
  const bar = document.getElementById("config-savebar");
  const hint = document.getElementById("config-savebar-hint");
  if (!bar || !hint || isConfigDirty()) return;
  hint.textContent = "Enregistré ✓";
  hint.classList.add("saved");
  bar.style.display = "";
  if (savedFlashTimer) clearTimeout(savedFlashTimer);
  savedFlashTimer = setTimeout(() => {
    hint.textContent = "Modifications non enregistrées";
    hint.classList.remove("saved");
    bar.style.display = "none";
  }, 1500);
  // Node : ne pas retenir la sortie du processus de test.
  if (savedFlashTimer && typeof savedFlashTimer.unref === "function") savedFlashTimer.unref();
}
if (typeof window !== "undefined" && window.addEventListener) {
  window.addEventListener("cetas:config-closed", revertUnsavedConfig);
}

// --- Panneaux de la Configuration (Fonctionnalités / Recherche / Apparence) ---

// Petit helper DOM.
function cfgEl(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

// Six lignes de Fonctionnalités (uniquement dans les Paramètres CETAS,
// jamais exposées dans la vue Agents).
const FEATURE_DEFS = [
  { key: "tts", label: "Synthèse vocale", kind: "system", fallback: "system" },
  { key: "transcription", label: "Transcription", kind: "system", fallback: "system" },
  { key: "prompt_enhance", label: "Amélioration de prompts/rôles", kind: "provider", fallback: "none" },
  { key: "summarizer", label: "Résumé IA", kind: "provider", fallback: "none" },
  { key: "title_gen", label: "Génération du titre", kind: "provider-conv", fallback: "conversation" },
  { key: "error_analysis", label: "Analyse des erreurs", kind: "provider", fallback: "none" },
];

let featuresLoaded = false;
let featuresSaved = null;
let featuresStaged = null;

async function loadFeaturesPanel() {
  const wrap = document.getElementById("features-rows");
  if (!wrap || featuresLoaded) return;
  featuresLoaded = true;
  wrap.innerHTML = "";
  wrap.appendChild(cfgEl("div", "sb-tree-empty", "Chargement…"));
  let prefs = null, provData = null;
  try { prefs = await getPrefs(); } catch (e) {}
  try { provData = await api("/api/providers"); } catch (e) {}
  const providers = (provData && provData.providers) || [];
  wrap.innerHTML = "";
  featuresSaved = {};
  featuresStaged = {};
  for (const def of FEATURE_DEFS) {
    const row = cfgEl("div", "feat-row");
    row.appendChild(cfgEl("span", "feat-name", def.label));
    const sel = cfgEl("select", "feat-select");
    sel.dataset.feature = def.key;
    const opts = [];
    if (def.kind === "system") {
      opts.push(["system", "Système (navigateur)"], ["none", "Aucun"]);
    } else {
      if (def.kind === "provider-conv") opts.push(["conversation", "Modèle de la conversation"]);
      opts.push(["none", "Aucun"]);
      for (const p of providers) {
        opts.push([p.id, p.label + (p.configured ? "" : " — non configuré")]);
      }
    }
    for (const [v, l] of opts) {
      const o = document.createElement("option");
      o.value = v; o.textContent = l;
      sel.appendChild(o);
    }
    const cur = (prefs && prefs[def.key]) || def.fallback;
    const val = Array.from(sel.options).some((o) => o.value === cur) ? cur : def.fallback;
    sel.value = val;
    featuresSaved[def.key] = val;
    featuresStaged[def.key] = val;
    sel.addEventListener("change", () => {
      featuresStaged[def.key] = sel.value;
      notifyConfigDirty();
    });
    row.appendChild(sel);
    wrap.appendChild(row);
  }
  registerPanelSaver("models", {
    isDirty: () => FEATURE_DEFS.some((d) => featuresStaged[d.key] !== featuresSaved[d.key]),
    save: async () => {
      const body = {};
      for (const d of FEATURE_DEFS) {
        if (featuresStaged[d.key] !== featuresSaved[d.key]) body[d.key] = featuresStaged[d.key];
      }
      if (Object.keys(body).length) await putPrefs(body);
      Object.assign(featuresSaved, featuresStaged);
      window.dispatchEvent(new CustomEvent("cetas:features-changed"));
    },
  });
}

// --- Panneau Recherche Web : moteurs + mode (staged) ---

let searchLoaded = false;
let searchState = null; // { mode, providers: [...] } — réponse GET /api/search/settings
let searchStaged = null; // { mode, enabled: {id: bool}, keys: {id: string}, websearchMode }
let searchSavedWebsearchMode = "auto";
let activeSearchTab = null;

async function refreshSearchSettings() {
  try {
    searchState = await api("/api/search/settings");
  } catch (e) {
    searchState = null;
  }
  return searchState;
}

async function loadSearchPanel() {
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl || searchLoaded) return;
  searchLoaded = true;
  const modeSel = document.getElementById("search-mode");
  const wsmSel = document.getElementById("cfg-websearch-mode");
  contentEl.innerHTML = "";
  contentEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement…"));
  const data = await refreshSearchSettings();
  let prefs = null;
  try { prefs = await getPrefs(); } catch (e) {}
  if (!data) {
    contentEl.innerHTML = "";
    contentEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement impossible."));
    return;
  }
  searchStaged = {
    mode: data.mode === "priority" ? "priority" : "race",
    enabled: {},
    keys: {},
    websearchMode: (prefs && prefs.websearch_mode) || "auto",
  };
  for (const p of data.providers || []) searchStaged.enabled[p.id] = !!p.enabled;
  searchSavedWebsearchMode = searchStaged.websearchMode;
  if (modeSel) {
    modeSel.value = searchStaged.mode;
    modeSel.onchange = () => {
      searchStaged.mode = modeSel.value;
      notifyConfigDirty();
    };
  }
  if (wsmSel) {
    // NOTE : modals.js ne sauvegarde plus ce champ automatiquement ;
    // il est « stagé » ici avec le reste de l'onglet Recherche Web.
    wsmSel.value = searchStaged.websearchMode;
    wsmSel.onchange = () => {
      searchStaged.websearchMode = wsmSel.value;
      notifyConfigDirty();
    };
  }
  renderSearchTabs();
  registerPanelSaver("search", {
    isDirty: () => {
      if (!searchStaged || !searchState) return false;
      if (searchStaged.mode !== (searchState.mode === "priority" ? "priority" : "race")) return true;
      if (searchStaged.websearchMode !== searchSavedWebsearchMode) return true;
      if (Object.keys(searchStaged.keys).length) return true;
      return (searchState.providers || []).some((p) => !!searchStaged.enabled[p.id] !== !!p.enabled);
    },
    save: async () => {
      const provBody = {};
      for (const p of searchState.providers || []) {
        const entry = { enabled: !!searchStaged.enabled[p.id] };
        if (searchStaged.keys[p.id] !== undefined) entry.key = searchStaged.keys[p.id];
        provBody[p.id] = entry;
      }
      await api("/api/search/settings", {
        method: "PUT",
        body: { mode: searchStaged.mode, providers: provBody },
      });
      if (searchStaged.websearchMode !== searchSavedWebsearchMode) {
        await putPrefs({ websearch_mode: searchStaged.websearchMode });
        searchSavedWebsearchMode = searchStaged.websearchMode;
      }
      searchStaged.keys = {};
      await refreshSearchSettings();
      searchStaged.mode = searchState.mode === "priority" ? "priority" : "race";
      for (const p of searchState.providers || []) searchStaged.enabled[p.id] = !!p.enabled;
      renderSearchTabs();
    },
  });
}

function renderSearchTabs() {
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl || !searchState) return;
  const providers = searchState.providers || [];
  if (activeSearchTab && !providers.some((p) => p.id === activeSearchTab)) activeSearchTab = null;
  const firstId = providers.length ? providers[0].id : null;
  const activeId = activeSearchTab || firstId;
  tabsEl.innerHTML = "";
  contentEl.innerHTML = "";
  providers.forEach((p) => {
    const isActive = p.id === activeId;
    const sec = cfgEl("div", "provider-section" + (isActive ? " active" : ""));
    sec.dataset.provider = p.id;
    renderSearchProviderPane(p, sec);
    contentEl.appendChild(sec);
    const b = cfgEl("button", "provider-tab" + (isActive ? " active" : ""), p.label);
    b.type = "button";
    b.dataset.provider = p.id;
    b.addEventListener("click", () => {
      activeSearchTab = p.id;
      tabsEl.querySelectorAll(".provider-tab").forEach((t) => t.classList.toggle("active", t === b));
      contentEl.querySelectorAll(".provider-section").forEach((s) => s.classList.toggle("active", s.dataset.provider === p.id));
    });
    tabsEl.appendChild(b);
  });
}

function searchProviderStaged(p) {
  if (!searchStaged) return false;
  if (searchStaged.keys[p.id] !== undefined) return true;
  return !!searchStaged.enabled[p.id] !== !!p.enabled;
}

function renderSearchProviderPane(p, sec) {
  sec.innerHTML = "";
  const wrap = cfgEl("div", "search-provider-pane");
  // Ligne d'activation (état « stagé »).
  const trow = cfgEl("div", "audio-setting-row search-toggle-row");
  trow.appendChild(cfgEl("span", "audio-setting-label", "Activer " + p.label));
  const tlabel = cfgEl("label", "plus-menu-toggle");
  const cb = document.createElement("input");
  cb.type = "checkbox";
  cb.checked = searchStaged ? !!searchStaged.enabled[p.id] : !!p.enabled;
  cb.setAttribute("aria-label", "Activer " + p.label);
  const slider = cfgEl("span", "plus-menu-toggle-slider");
  tlabel.appendChild(cb);
  tlabel.appendChild(slider);
  trow.appendChild(tlabel);
  cb.addEventListener("change", () => {
    searchStaged.enabled[p.id] = cb.checked;
    notifyConfigDirty();
    refreshStagedHint();
  });
  wrap.appendChild(trow);
  const status = cfgEl("p", "search-status", searchStatusText(p));
  const stagedHint = cfgEl("p", "search-key-staged", "");
  const refreshStagedHint = () => {
    stagedHint.textContent = searchProviderStaged(p) ? "Modifié — cliquez « Enregistrer »." : "";
  };
  // Zone clé
  if (p.keyless) {
    wrap.appendChild(cfgEl("p", "apikey-intro", "Aucune clé requise : ce moteur est utilisé en repli, sans authentification."));
  } else {
    const labelRow = cfgEl("div", "apikey-label-row");
    const lab = cfgEl("label", "sp-modal-label", p.key_label || ("Clé API " + p.label));
    lab.htmlFor = "search-key-" + p.id;
    labelRow.appendChild(lab);
    if (p.key_url) {
      const a = cfgEl("a", "apikey-get-link", "Obtenir une clé");
      a.href = p.key_url;
      a.target = "_blank";
      a.rel = "noopener noreferrer";
      labelRow.appendChild(a);
    }
    wrap.appendChild(labelRow);
    const field = cfgEl("div", "apikey-field");
    const iwrap = cfgEl("div", "apikey-input-wrap");
    const input = document.createElement("input");
    input.type = "password";
    input.id = "search-key-" + p.id;
    input.className = "sp-modal-input apikey-input";
    input.autocomplete = "off";
    input.spellcheck = false;
    input.placeholder = p.configured ? "•••••••• (configurée)" : "Coller la clé API…";
    const eye = cfgEl("button", "search-eye-btn", "👁");
    eye.type = "button";
    eye.title = "Afficher la clé";
    eye.setAttribute("aria-label", "Afficher la clé");
    eye.addEventListener("click", () => {
      const show = input.type === "password";
      input.type = show ? "text" : "password";
      eye.classList.toggle("on", show);
    });
    iwrap.appendChild(input);
    field.appendChild(iwrap);
    field.appendChild(eye);
    wrap.appendChild(field);
    input.addEventListener("input", () => {
      const v = input.value.trim();
      if (v) searchStaged.keys[p.id] = v;
      else if (p.configured) searchStaged.keys[p.id] = ""; // chaîne vide = suppression
      else delete searchStaged.keys[p.id];
      notifyConfigDirty();
      refreshStagedHint();
    });
    const row = cfgEl("div", "search-key-actions");
    const delBtn = cfgEl("button", "apikey-delete-link", "Supprimer la clé");
    delBtn.type = "button";
    delBtn.title = "Supprimer la clé (appliqué à l'enregistrement)";
    delBtn.style.display = p.configured ? "" : "none";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer la clé " + p.label + " ? (appliqué à l'enregistrement)", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      searchStaged.keys[p.id] = "";
      input.value = "";
      notifyConfigDirty();
      refreshStagedHint();
    });
    row.appendChild(delBtn);
    wrap.appendChild(row);
  }
  wrap.appendChild(stagedHint);
  wrap.appendChild(status);
  sec.appendChild(wrap);
  refreshStagedHint();
}

function currentSearchProvider(id) {
  const list = (searchState && searchState.providers) || [];
  return list.find((x) => x.id === id) || null;
}

function selectSearchTab(id) {
  activeSearchTab = id;
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl) return;
  tabsEl.querySelectorAll(".provider-tab").forEach((t) => t.classList.toggle("active", t.dataset.provider === id));
  contentEl.querySelectorAll(".provider-section").forEach((s) => s.classList.toggle("active", s.dataset.provider === id));
}

function searchStatusText(p) {
  if (!p) return "";
  if (!p.enabled) return "Désactivé — ce moteur ne sera pas utilisé.";
  if (p.keyless) return "Activé — utilisé sans clé.";
  return p.configured ? "Activé — clé configurée." : "Activé — aucune clé : ce moteur sera ignoré.";
}

// --- Panneau Apparence : thème + palettes (aperçu immédiat, sauvegarde explicite) ---
let appearanceLoaded = false;
let appearanceSaved = null;
let appearanceStaged = null;

async function loadAppearancePanel() {
  if (appearanceLoaded) return;
  appearanceLoaded = true;
  const themeSel = document.getElementById("cfg-theme");
  appearanceSaved = {
    theme: document.documentElement.dataset.theme || "clair",
    palette: document.documentElement.dataset.palette || "bleu",
  };
  appearanceStaged = { ...appearanceSaved };
  if (themeSel) {
    themeSel.value = appearanceStaged.theme;
    themeSel.onchange = () => {
      appearanceStaged.theme = themeSel.value;
      applyTheme(themeSel.value); // aperçu immédiat
      notifyConfigDirty();
    };
  }
  renderPaletteSwatches();
  registerPanelSaver("appearance", {
    isDirty: () =>
      appearanceStaged.theme !== appearanceSaved.theme ||
      appearanceStaged.palette !== appearanceSaved.palette,
    save: async () => {
      const body = {};
      if (appearanceStaged.theme !== appearanceSaved.theme) body.theme = appearanceStaged.theme;
      if (appearanceStaged.palette !== appearanceSaved.palette) body.palette = appearanceStaged.palette;
      if (Object.keys(body).length) await putPrefs(body);
      appearanceSaved = { ...appearanceStaged };
    },
    revert: () => {
      appearanceStaged = { ...appearanceSaved };
      applyTheme(appearanceSaved.theme);
      applyPalette(appearanceSaved.palette);
      const sel = document.getElementById("cfg-theme");
      if (sel) sel.value = appearanceSaved.theme;
    },
  });
}

function renderPaletteSwatches() {
  const wrap = document.getElementById("palette-swatches");
  if (!wrap) return;
  wrap.innerHTML = "";
  const current = (appearanceStaged && appearanceStaged.palette) ||
    document.documentElement.dataset.palette || "bleu";
  for (const p of THEME_PALETTES) {
    const b = cfgEl("button", "palette-swatch" + (p.id === current ? " active" : ""), "");
    b.type = "button";
    b.dataset.palette = p.id;
    b.title = p.label;
    b.setAttribute("aria-label", "Palette " + p.label);
    b.style.setProperty("--sw", p.color);
    b.addEventListener("click", () => {
      applyPalette(p.id); // aperçu immédiat
      if (appearanceStaged) appearanceStaged.palette = p.id;
      notifyConfigDirty();
    });
    wrap.appendChild(b);
  }
}

// --- Helpers exportés pour les tests ---
export const _test = {
  FEATURE_DEFS,
  loadFeaturesPanel,
  loadSearchPanel,
  renderSearchTabs,
  selectSearchTab,
  currentSearchProvider,
  searchStatusText,
  loadAppearancePanel,
  renderPaletteSwatches,
  registerPanelSaver,
  notifyConfigDirty,
  isConfigDirty,
  saveConfigPanel,
  saveAllConfigPanels,
  revertUnsavedConfig,
};
