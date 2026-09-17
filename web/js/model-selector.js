import { api } from "./api.js";

// --- Panneau Configuration > Sélecteur de modèles ---
//
// Pour chaque alias (famille/mode), l'utilisateur choisit :
// - "1 modèle" : un seul modèle fixe, pas de fallback ;
// - "Fallback" : un ensemble de modèles, tirage aléatoire à chaque requête
//   puis bascule séquentielle en cas d'échec (moteur : pool mélangé).
//
// La configuration est persistée via PUT /api/aliases (overrides) et
// restaurée au démarrage. Chaque changement émet "cetas:aliases-changed"
// pour rafraîchir le menu + et les sélecteurs.

export function poolKey(m) {
  return m.provider + "/" + m.model;
}

export function sameKeySet(a, b) {
  if (a.length !== b.length) return false;
  const s = new Set(a.map(poolKey));
  return b.every((m) => s.has(poolKey(m)));
}

const ENGINE_NAMES = { llamacpp: "llama.cpp", ollama: "Ollama", lmstudio: "LM Studio" };
export function engineName(id) {
  return ENGINE_NAMES[id] || id || "";
}

// Sous-titre propre pour le menu + et les selects : le modèle fixe ou
// "Fallback · N modèles". Jamais de texte technique.
export function modeSubtitle(fam, mode) {
  const pool = mode.pool || [];
  if (fam.local) {
    if (pool.length === 1) return pool[0].label || pool[0].model;
    if (pool.length > 1) return "Fallback · " + pool.length + " modèles";
    return engineName(mode.engine);
  }
  if (pool.length === 1) return pool[0].label || pool[0].model;
  if (pool.length > 1) return "Fallback · " + pool.length + " modèles";
  return "";
}

export function familyShortLabel(fam) {
  return { "samagent-nano": "Nano", "samagent-n4": "N4", "samagent-n8": "N8", code: "Code", samgen: "SamGen" }[fam.id] || fam.label || fam.id;
}

// Reconstruit les overrides à partir des pools effectifs vs défauts.
// Utilisé au chargement pour initialiser l'état "staged".
export function buildStaged(families, defaults) {
  const staged = {};
  const defByFam = {};
  for (const d of defaults) {
    defByFam[d.id] = {};
    for (const m of d.modes || []) defByFam[d.id][m.mode] = m.pool || [];
  }
  for (const f of families) {
    for (const m of f.modes || []) {
      const pool = m.pool || [];
      const def = (defByFam[f.id] || {})[m.mode] || [];
      const differs = f.local ? pool.length > 0 : !sameKeySet(pool, def);
      if (differs) {
        if (!staged[f.id]) staged[f.id] = {};
        staged[f.id][m.mode] = pool.map((p) => ({ provider: p.provider, model: p.model }));
      }
    }
  }
  return staged;
}

export function fmtPrice(m) {
  if (!m.input_per_1m && !m.output_per_1m) return "gratuit";
  const f = (v) => (v === 0 ? "0" : "$" + v);
  return f(m.input_per_1m) + " → " + f(m.output_per_1m) + " /1M";
}

const BODY_MAIN = "selector-body";
const BODY_AGENT = "selector-agent-body";
const SAVE_MAIN = "selector-save-state";
const SAVE_AGENT = "selector-agent-save-state";
// Onglet actif ("selector" = chat, "selector-agent" = vue Agents) : détermine
// le corps et l'indicateur de sauvegarde utilisés par les helpers.
let activeTab = "selector";
const loadedBodies = { [BODY_MAIN]: false, [BODY_AGENT]: false };
let families = [];
let defaults = [];
let staged = {};
let discovered = {}; // engine -> [ids]
let saveTimer = null;

function activeBodyId() {
  return activeTab === "selector-agent" ? BODY_AGENT : BODY_MAIN;
}

function setSaveState(txt, cls) {
  const el = document.getElementById(activeTab === "selector-agent" ? SAVE_AGENT : SAVE_MAIN);
  if (!el) return;
  el.textContent = txt;
  el.className = "ms-save-state" + (cls ? " " + cls : "");
}

function defaultPool(famId, modeId) {
  const f = defaults.find((x) => x.id === famId);
  const m = f && (f.modes || []).find((x) => x.mode === modeId);
  return (m && m.pool) || [];
}

function effectivePool(famId, modeId) {
  const f = families.find((x) => x.id === famId);
  const m = f && (f.modes || []).find((x) => x.mode === modeId);
  return (m && m.pool) || [];
}

// Modèles proposés pour un mode : le pool par défaut (+ les éventuels
// modèles effectifs hors défaut), ou les modèles découverts (SamGen).
function availableModels(fam, mode) {
  if (fam.local) {
    const ids = discovered[mode.engine] || [];
    const cur = effectivePool(fam.id, mode.mode).map((p) => p.model);
    const all = [...new Set([...ids, ...cur])];
    return all.map((id) => ({ provider: mode.engine, model: id, label: id, input_per_1m: 0, output_per_1m: 0 }));
  }
  const def = defaultPool(fam.id, mode.mode);
  const seen = new Set(def.map(poolKey));
  const extra = effectivePool(fam.id, mode.mode).filter((p) => !seen.has(poolKey(p)));
  return [...def, ...extra];
}

function currentKind(fam, mode) {
  const pool = effectivePool(fam.id, mode.mode);
  if (fam.local && pool.length === 0) return "fallback";
  return pool.length <= 1 ? "single" : "fallback";
}

function el(tag, cls, txt) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (txt !== undefined) e.textContent = txt;
  return e;
}

async function persist() {
  setSaveState("Enregistrement…", "ms-saving");
  try {
    const data = await api("/api/aliases", { method: "PUT", body: staged });
    families = data.families || families;
    setSaveState("Enregistré ✓", "ms-ok");
    setTimeout(() => setSaveState("", ""), 2500);
    window.dispatchEvent(new CustomEvent("cetas:aliases-changed", { detail: { source: activeTab } }));
  } catch (e) {
    setSaveState("Erreur : " + (e.message || e), "ms-err");
  }
}

function schedulePersist() {
  clearTimeout(saveTimer);
  // Fige l'onglet d'origine : l'utilisateur peut changer d'onglet pendant
  // le debounce, la sauvegarde doit rester rattachée au bon panneau.
  const tab = activeTab;
  saveTimer = setTimeout(() => persistAs(tab), 450);
}

async function persistAs(tab) {
  const prev = activeTab;
  activeTab = tab;
  try {
    await persist();
  } finally {
    activeTab = prev;
  }
}

function setModeSelection(famId, modeId, members) {
  if (!members || members.length === 0) {
    if (staged[famId]) {
      delete staged[famId][modeId];
      if (Object.keys(staged[famId]).length === 0) delete staged[famId];
    }
  } else {
    if (!staged[famId]) staged[famId] = {};
    staged[famId][modeId] = members.map((m) => ({ provider: m.provider, model: m.model }));
  }
  schedulePersist();
}

function renderModeCard(fam, mode) {
  const card = el("div", "ms-mode-card");
  const head = el("div", "ms-mode-head");
  head.appendChild(el("span", "ms-mode-title", mode.label));
  const seg = el("div", "ms-seg");
  const kind = currentKind(fam, mode);
  const btnSingle = el("button", "ms-seg-btn" + (kind === "single" ? " active" : ""), "1 modèle");
  const btnFallback = el("button", "ms-seg-btn" + (kind === "fallback" ? " active" : ""), "Fallback");
  btnSingle.type = "button";
  btnFallback.type = "button";
  seg.appendChild(btnSingle);
  seg.appendChild(btnFallback);
  head.appendChild(seg);
  const isOverridden = !!(staged[fam.id] && staged[fam.id][mode.mode]);
  if (isOverridden) {
    const reset = el("button", "ms-reset-btn", "Réinitialiser");
    reset.type = "button";
    reset.title = "Revenir à la sélection par défaut";
    reset.addEventListener("click", () => {
      setModeSelection(fam.id, mode.mode, null);
      refreshModeCard(fam, mode);
    });
    head.appendChild(reset);
  }
  card.appendChild(head);

  const list = el("div", "ms-model-list");
  const avail = availableModels(fam, mode);
  const cur = effectivePool(fam.id, mode.mode);
  const curKeys = new Set(cur.map(poolKey));
  // SamGen sans sélection : le défaut = tous les modèles découverts.
  const allChecked = fam.local && cur.length === 0;
  if (avail.length === 0) {
    list.appendChild(el("div", "ms-empty", fam.local ? "Aucun modèle découvert — vérifiez que le moteur tourne." : "Aucun modèle disponible."));
  }
  const inputType = kind === "single" ? "radio" : "checkbox";
  const groupName = "ms-" + fam.id + "-" + mode.mode;
  // En mode single sans sélection (cas SamGen par défaut), présélectionner le 1er.
  let singleSel = cur.length === 1 ? poolKey(cur[0]) : null;
  if (kind === "single" && !singleSel && avail.length > 0) singleSel = poolKey(avail[0]);
  for (const m of avail) {
    const k = poolKey(m);
    const row = el("label", "ms-model-row");
    const input = document.createElement("input");
    input.type = inputType;
    if (inputType === "radio") input.name = groupName;
    input.checked = kind === "single" ? k === singleSel : allChecked || curKeys.has(k);
    input.dataset.key = k;
    input.dataset.provider = m.provider;
    input.dataset.model = m.model;
    row.appendChild(input);
    const main = el("span", "ms-model-main");
    main.appendChild(el("span", "ms-model-name", m.label || m.model));
    main.appendChild(el("span", "ms-model-prov", m.provider));
    row.appendChild(main);
    if (!fam.local) row.appendChild(el("span", "ms-model-price", fmtPrice(m)));
    list.appendChild(row);
  }
  card.appendChild(list);

  const hint = el("div", "ms-hint", kind === "single"
    ? "Modèle fixe — pas de fallback."
    : "Tirage aléatoire à chaque requête, bascule séquentielle en cas d'échec.");
  card.appendChild(hint);

  const onChange = () => {
    const checked = [...list.querySelectorAll("input:checked")];
    if (kind === "fallback" && checked.length === 0) {
      // Toujours au moins 1 modèle : on réactive celui qu'on vient de décocher.
      const last = list.querySelectorAll("input")[0];
      if (last) last.checked = true;
      return;
    }
    const members = (kind === "single" ? checked.slice(0, 1) : checked)
      .map((i) => ({ provider: i.dataset.provider, model: i.dataset.model }));
    setModeSelection(fam.id, mode.mode, members);
    hint.textContent = members.length <= 1
      ? "Modèle fixe — pas de fallback."
      : "Tirage aléatoire à chaque requête, bascule séquentielle en cas d'échec.";
  };
  list.addEventListener("change", onChange);

  const switchKind = (newKind) => {
    if (newKind === kind) return;
    const checked = [...list.querySelectorAll("input:checked")]
      .map((i) => ({ provider: i.dataset.provider, model: i.dataset.model }));
    // Bascule vers Fallback : tout l'ensemble disponible. On conserve une
    // sélection manuelle existante si elle couvre déjà plusieurs modèles ;
    // sinon (cas « 1 modèle »), on coche tout — sinon le pool resterait à
    // 1 modèle et l'onglet semblerait ne rien faire (bug constaté).
    const members = newKind === "single"
      ? checked.slice(0, 1)
      : (checked.length > 1 ? checked : avail.map((m) => ({ provider: m.provider, model: m.model })));
    setModeSelection(fam.id, mode.mode, members);
    refreshModeCard(fam, mode);
  };
  btnSingle.addEventListener("click", () => switchKind("single"));
  btnFallback.addEventListener("click", () => switchKind("fallback"));

  return card;
}

function refreshModeCard(fam, mode) {
  // Re-résout les pools effectifs depuis le staged local (sans attendre le serveur).
  const f = families.find((x) => x.id === fam.id);
  const m = f && (f.modes || []).find((x) => x.mode === mode.mode);
  if (m) {
    const st = (staged[fam.id] || {})[mode.mode];
    if (st) {
      m.pool = st.map((s) => {
        const a = availableModels(fam, mode).find((x) => poolKey(x) === s.provider + "/" + s.model);
        return a || { provider: s.provider, model: s.model, label: s.model };
      });
    } else {
      m.pool = defaultPool(fam.id, mode.mode);
      if (fam.local) m.pool = [];
    }
  }
  const body = document.getElementById(activeBodyId());
  if (body) {
    const idx = [...body.querySelectorAll(".ms-family-card")].findIndex((c) => c.dataset.fam === fam.id);
    if (idx >= 0) {
      const fresh = renderFamilyCard(filterFamilyForTab(fam));
      body.children[idx].replaceWith(fresh);
    }
  }
}

// Vue filtrée d'une famille pour l'onglet actif : l'onglet agent ne montre
// que les modes agent (m.agent), l'onglet principal montre tout.
function filterFamilyForTab(fam) {
  if (activeTab !== "selector-agent") return fam;
  return { ...fam, modes: (fam.modes || []).filter((m) => m.agent) };
}

function renderFamilyCard(fam) {
  const card = el("div", "ms-family-card");
  card.dataset.fam = fam.id;
  const head = el("div", "ms-family-head");
  head.appendChild(el("span", "ms-family-title", familyShortLabel(fam)));
  if (fam.local) head.appendChild(el("span", "ms-family-tag", "local"));
  card.appendChild(head);
  for (const mode of fam.modes || []) {
    card.appendChild(renderModeCard(fam, mode));
  }
  return card;
}

export async function loadSelectorPanel() {
  return loadInto("selector", BODY_MAIN);
}

// Onglet "Sélecteur agent" : mêmes réglages (1 modèle / Fallback) mais
// limités aux familles et modes agent (vue Agents).
export async function loadAgentSelectorPanel() {
  return loadInto("selector-agent", BODY_AGENT);
}

async function loadInto(tab, bodyId) {
  activeTab = tab;
  const body = document.getElementById(bodyId);
  if (!body || loadedBodies[bodyId]) return;
  loadedBodies[bodyId] = true;
  body.innerHTML = "";
  body.appendChild(el("div", "ms-loading", "Chargement…"));
  try {
    const data = await api("/api/aliases");
    families = data.families || [];
    defaults = data.defaults || [];
    staged = buildStaged(families, defaults);
    // Modèles locaux découverts (SamGen), en parallèle.
    const samgen = families.find((f) => f.id === "samgen");
    if (samgen) {
      await Promise.all((samgen.modes || []).map(async (m) => {
        try {
          const d = await api("/api/local/models?engine=" + encodeURIComponent(m.engine));
          discovered[m.engine] = d.models || [];
        } catch (e) {
          discovered[m.engine] = [];
        }
      }));
    }
  } catch (e) {
    body.innerHTML = "";
    body.appendChild(el("div", "ms-empty", "Impossible de charger : " + (e.message || e)));
    loadedBodies[bodyId] = false;
    return;
  }
  body.innerHTML = "";
  const order = ["samagent-nano", "samagent-n4", "samagent-n8", "code", "samgen"];
  const sorted = [...families].sort((a, b) => order.indexOf(a.id) - order.indexOf(b.id));
  for (const fam of sorted) {
    const view = filterFamilyForTab(fam);
    if ((view.modes || []).length === 0) continue;
    body.appendChild(renderFamilyCard(view));
  }
}

// Rafraîchit les panneaux après une sauvegarde externe.
export function resetSelectorPanel() {
  loadedBodies[BODY_MAIN] = false;
  loadedBodies[BODY_AGENT] = false;
}

if (typeof window !== "undefined") {
  // Le menu + ou l'autre onglet peut modifier les overrides : recharger le
  // panneau affiché. On ignore notre propre sauvegarde (déjà reflétée).
  window.addEventListener("cetas:aliases-changed", (e) => {
    const src = e.detail && e.detail.source;
    for (const [bodyId, tab] of [[BODY_MAIN, "selector"], [BODY_AGENT, "selector-agent"]]) {
      if (src === tab || !loadedBodies[bodyId]) continue;
      const body = document.getElementById(bodyId);
      if (body && body.isConnected) {
        loadedBodies[bodyId] = false;
        loadInto(tab, bodyId);
      }
    }
  });
}
