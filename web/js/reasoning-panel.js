// Panneau lateral droit "Raisonnement" : affiche les informations de la
// requete (modele, temps de generation, tokens, contexte, cout) puis le
// raisonnement du modele en streaming. S'ouvre automatiquement au premier
// delta sauf si l'utilisateur l'a ferme manuellement pendant le tour en
// cours. Quand le raisonnement est termine, le panneau se masque
// automatiquement ; chaque requete garde son bouton "Raisonnement"
// (au-dessus du message utilisateur) pour le reconsulter.

import { api } from "./api.js";

function el(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

const EMPTY_HTML = "Le raisonnement du modèle s'affichera ici pendant la génération.";

let userClosed = false;

// Etat du tour courant affiche dans le panneau.
const turn = {
  modelLabel: "",
  provider: "",
  model: "",
  startTs: 0,
  inputTok: null,
  outputTok: null,
  ctxMax: null,
  inputPer1M: null,
  outputPer1M: null,
  cost: null,
  elapsedMs: null,
  tickTimer: 0,
};

// Cache des infos modele (contexte + tarifs) par "provider/model".
const modelInfoCache = new Map();

// --- Traduction DeepThink ---
// Le bouton "Traduire" ne remplace JAMAIS le raisonnement : la traduction
// s'affiche SOUS l'original (bloc .reason-translation). Un clic traduit le
// tour affiché immédiatement ET active le mode auto : chaque raisonnement
// suivant est traduit automatiquement 5 secondes après sa fin. Un second
// clic désactive le mode auto. Le mode auto persiste (localStorage).
const DT_AUTO_KEY = "cetas.deepthink.auto";
let dtAutoDelayMs = 5000;

// Réglable pour les tests (évite d'attendre 5s réelles).
export function _setAutoTranslateDelayForTests(ms) {
  dtAutoDelayMs = ms;
}

let translating = false;
let pendingAutoTimer = 0;

function isAutoTranslate() {
  try {
    return localStorage.getItem(DT_AUTO_KEY) === "1";
  } catch (e) {
    return false;
  }
}

function setAutoTranslate(on) {
  try {
    if (on) localStorage.setItem(DT_AUTO_KEY, "1");
    else localStorage.removeItem(DT_AUTO_KEY);
  } catch (e) {}
}

// État du bouton selon le mode auto (appelé à la création du bloc et à
// chaque changement de mode).
function refreshTranslateBtn() {
  const b = body();
  const btn = b && b.querySelector(".reason-translate-btn");
  if (!btn) return;
  if (isAutoTranslate()) {
    btn.textContent = "🌐 Auto ✓";
    btn.classList.add("on");
    btn.title = "Traduction automatique activée — cliquer pour désactiver";
  } else {
    btn.textContent = "🌐 Traduire";
    btn.classList.remove("on");
    btn.title = "Traduire le raisonnement + activer la traduction auto (DeepThink Global)";
  }
  if (!translating) btn.disabled = false;
}

// Supprime le bloc de traduction affiché (nouveau contenu en streaming :
// la traduction serait périmée).
function clearTranslation() {
  const b = body();
  const box = b && b.querySelector(".reason-translation");
  if (box) box.remove();
}

function translationBlock(translated, lang, isError) {
  const box = el(
    "div",
    "reason-translation" + (isError ? " reason-translation-error" : "")
  );
  box.appendChild(
    el("div", "reason-trans-sec", "TRADUCTION" + (lang ? " · " + String(lang).toUpperCase() : ""))
  );
  box.appendChild(el("div", "reason-translation-text", (isError ? "⚠ " : "") + translated));
  return box;
}

// Traduit le raisonnement affiché et l'ajoute SOUS l'original (jamais de
// remplacement). Rouvre le panneau sauf fermeture manuelle.
async function translateCurrentTurn() {
  const b = body();
  const t = b && b.querySelector(".reason-text");
  if (!t || translating) return;
  const src = (t.textContent || "").trim();
  if (!src || b.querySelector(".reason-translation")) return;
  translating = true;
  const btn = b.querySelector(".reason-translate-btn");
  if (btn) {
    btn.disabled = true;
    btn.textContent = "…";
  }
  try {
    const data = await api("/api/deepthink/translate", { method: "POST", body: { text: src } });
    const translated = data && typeof data.translation === "string" ? data.translation.trim() : "";
    if (!translated) throw new Error("Traduction vide.");
    const b2 = body();
    // Le bloc a pu être reconstruit pendant l'appel : on ne touche que le
    // même nœud texte, sinon on annule.
    if (b2 && b2.querySelector(".reason-text") === t) {
      b2.appendChild(translationBlock(translated, (data && data.lang) || "", false));
      if (!userClosed) openReasonPanel();
      schedulePanelScroll();
    }
  } catch (e) {
    const b2 = body();
    if (b2 && b2.querySelector(".reason-text") === t) {
      b2.appendChild(translationBlock((e && e.message) || "Échec de la traduction", "", true));
    }
  } finally {
    translating = false;
    refreshTranslateBtn();
  }
}

// Clic sur le bouton : OFF -> traduit maintenant + active l'auto ;
// ON -> désactive l'auto.
async function toggleReasonTranslation() {
  const b = body();
  const btn = b && b.querySelector(".reason-translate-btn");
  if (!b || !btn || translating) return;
  if (isAutoTranslate()) {
    setAutoTranslate(false);
    refreshTranslateBtn();
    return;
  }
  setAutoTranslate(true);
  refreshTranslateBtn();
  await translateCurrentTurn();
}

// Fin d'un raisonnement : si l'auto est actif, traduction automatique
// 5 secondes après. Vérification d'identité au déclenchement : si le
// panneau affiche un autre texte entre-temps, on annule.
function scheduleAutoTranslate() {
  if (pendingAutoTimer) {
    clearTimeout(pendingAutoTimer);
    pendingAutoTimer = 0;
  }
  if (!isAutoTranslate()) return;
  const b = body();
  const t = b && b.querySelector(".reason-text");
  const src = t ? (t.textContent || "").trim() : "";
  if (!src) return;
  pendingAutoTimer = setTimeout(() => {
    pendingAutoTimer = 0;
    const b2 = body();
    const t2 = b2 && b2.querySelector(".reason-text");
    if (!t2 || (t2.textContent || "").trim() !== src) return;
    translateCurrentTurn();
  }, dtAutoDelayMs);
  if (pendingAutoTimer && typeof pendingAutoTimer.unref === "function") pendingAutoTimer.unref();
}

function panel() {
  return document.getElementById("reason-panel");
}

function body() {
  return document.getElementById("reason-panel-body");
}

function panelRaf(fn) {
  if (typeof requestAnimationFrame === "function") return requestAnimationFrame(fn);
  return setTimeout(fn, 16);
}

let panelScrollFrame = 0;

// Coalesce le scroll sur une frame : evite le reflow (scrollHeight) a
// chaque delta de raisonnement (technique Marexcode).
function schedulePanelScroll() {
  const b = body();
  if (!b || panelScrollFrame) return;
  panelScrollFrame = panelRaf(() => {
    panelScrollFrame = 0;
    const bb = body();
    if (bb) bb.scrollTop = bb.scrollHeight;
  });
}

function hasContent() {
  const b = body();
  const t = b && b.querySelector(".reason-text");
  return !!(t && t.textContent.trim());
}

// --- Bloc informations de la requete -------------------------------------

function fmtTok(n) {
  if (n == null) return "—";
  if (n >= 1e6) return (Math.round(n / 1e5) / 10) + "M";
  if (n >= 1000) return Math.round(n / 1000) + "K";
  return String(n);
}

function fmtSecs(ms) {
  const s = ms / 1000;
  return (s < 10 ? s.toFixed(1) : Math.round(s)) + "s";
}

// Construit le bloc infos + l'en-tete REASONING s'ils n'existent pas.
function ensureInfoBlock() {
  const b = body();
  if (!b || b.querySelector(".reason-info")) return;
  b.innerHTML = "";
  const info = el("div", "reason-info");
  info.appendChild(riRow("Modèle :", "ri-model", turn.modelLabel || "—", true));
  info.appendChild(riRow("Temps de génération :", "ri-time", "0s", false));
  info.appendChild(riRow("Input :", "ri-input", "—", false));
  info.appendChild(riRow("Output - t/s :", "ri-output", "—", false));
  info.appendChild(riRow("Contexte :", "ri-ctx", "—", false));
  info.appendChild(riRow("Coût :", "ri-cost", "—", false));
  b.appendChild(info);
  const sec = el("div", "reason-sec");
  sec.appendChild(el("span", "reason-sec-label", "REASONING"));
  const tbtn = el("button", "reason-translate-btn", "🌐 Traduire");
  tbtn.type = "button";
  tbtn.title = "Traduire le raisonnement (DeepThink Global)";
  tbtn.addEventListener("click", toggleReasonTranslation);
  sec.appendChild(tbtn);
  b.appendChild(sec);
  refreshTranslateBtn(); // reflète le mode auto persisté
  const t = el("div", "reason-text");
  b.appendChild(t);
  startTick();
  // Le modele et ses infos ont pu arriver avant le premier delta.
  renderUsage();
}

function riRow(label, valueCls, value, accent) {
  const row = el("div", "ri-row");
  row.appendChild(el("span", "ri-k", label + " "));
  const v = el("span", valueCls + (accent ? " ri-accent" : ""), value);
  row.appendChild(v);
  return row;
}

function setRow(cls, text, html) {
  const b = body();
  const e = b && b.querySelector("." + cls);
  if (!e) return;
  if (html !== undefined) e.innerHTML = html;
  else e.textContent = text;
}

function renderTime() {
  const ms = turn.elapsedMs != null ? turn.elapsedMs : Date.now() - turn.startTs;
  setRow("ri-time", fmtSecs(Math.max(0, ms)));
}

function startTick() {
  stopTick();
  if (!turn.startTs) turn.startTs = Date.now();
  renderTime();
  turn.tickTimer = setInterval(renderTime, 500);
  // Node (tests) : ne pas retenir la boucle d'evenements.
  if (turn.tickTimer && typeof turn.tickTimer.unref === "function") turn.tickTimer.unref();
}

function stopTick() {
  if (turn.tickTimer) {
    clearInterval(turn.tickTimer);
    turn.tickTimer = 0;
  }
}

function renderUsage() {
  if (turn.inputTok != null) setRow("ri-input", turn.inputTok.toLocaleString("fr") + " tok");
  if (turn.outputTok != null) {
    let txt = turn.outputTok.toLocaleString("fr") + " tok";
    const ms = turn.elapsedMs != null ? turn.elapsedMs : Date.now() - turn.startTs;
    if (ms > 0) txt += " (" + Math.round((turn.outputTok / ms) * 1000) + "/s)";
    setRow("ri-output", txt);
  }
  if (turn.inputTok != null || turn.outputTok != null) {
    const used = (turn.inputTok || 0) + (turn.outputTok || 0);
    if (turn.ctxMax) {
      const pct = Math.min(100, Math.round((used / turn.ctxMax) * 100));
      const color = pct >= 90 ? "#ef4444" : pct >= 70 ? "#eab308" : "#22c55e";
      setRow(
        "ri-ctx",
        undefined,
        '<span style="color:' + color + '">' + fmtTok(used) + " / " + fmtTok(turn.ctxMax) + "</span> · " + pct + "% utilisé"
      );
    } else {
      setRow("ri-ctx", fmtTok(used) + " tok");
    }
  }
  if (turn.cost != null) setRow("ri-cost", "$" + turn.cost.toFixed(4));
  else if (turn.inputTok != null && turn.inputPer1M != null) {
    const c = (turn.inputTok / 1e6) * turn.inputPer1M + ((turn.outputTok || 0) / 1e6) * (turn.outputPer1M || 0);
    setRow("ri-cost", "$" + c.toFixed(4));
  }
}

function fetchModelInfo() {
  const key = turn.provider + "/" + turn.model;
  if (!turn.provider || !turn.model) return;
  if (modelInfoCache.has(key)) {
    applyModelInfo(modelInfoCache.get(key));
    return;
  }
  modelInfoCache.set(key, null); // anti double-fetch
  api("/api/model-info?provider=" + encodeURIComponent(turn.provider) + "&model=" + encodeURIComponent(turn.model))
    .then((info) => {
      modelInfoCache.set(key, info || null);
      applyModelInfo(info);
    })
    .catch(() => {
      modelInfoCache.set(key, null);
    });
}

function applyModelInfo(info) {
  if (!info || info.unknown) return;
  if (info.context_window) turn.ctxMax = info.context_window;
  if (info.input_per_1m != null) turn.inputPer1M = info.input_per_1m;
  if (info.output_per_1m != null) turn.outputPer1M = info.output_per_1m;
  renderUsage();
}

// --- API publique ----------------------------------------------------------

export function initReasonPanel() {
  const closeBtn = document.getElementById("reason-panel-close");
  if (closeBtn) closeBtn.addEventListener("click", () => closeReasonPanel(true));
}

export function openReasonPanel() {
  const p = panel();
  if (!p) return;
  p.classList.add("open");
}

export function closeReasonPanel(manual) {
  if (manual) userClosed = true;
  const p = panel();
  if (p) p.classList.remove("open");
}

export function isReasonPanelOpen() {
  const p = panel();
  return !!p && p.classList.contains("open");
}

// Debut d'un tour : memorise le modele (si connu) pour le bloc infos.
export function beginReasonTurn(info) {
  info = info || {};
  turn.modelLabel = info.label || "";
  turn.provider = info.provider || "";
  turn.model = info.model || "";
  turn.startTs = Date.now();
  turn.inputTok = null;
  turn.outputTok = null;
  turn.ctxMax = null;
  turn.inputPer1M = null;
  turn.outputPer1M = null;
  turn.cost = null;
  turn.elapsedMs = null;
  clearTranslation();
}

// L'evenement route apporte le modele : met a jour le bloc infos et
// recupere fenetre de contexte + tarifs.
export function setReasonModel(label, provider, model) {
  turn.modelLabel = label || turn.modelLabel;
  turn.provider = provider || turn.provider;
  turn.model = model || turn.model;
  // Les infos modele (contexte, tarifs) sont chargees des que le modele est
  // connu, meme si le bloc n'existe pas encore (il les reprendra).
  fetchModelInfo();
  const b = body();
  if (b && b.querySelector(".reason-info")) setRow("ri-model", turn.modelLabel || "—");
}

export function resetReasonPanel() {
  userClosed = false;
  stopTick();
  turn.modelLabel = "";
  turn.provider = "";
  turn.model = "";
  turn.startTs = 0;
  turn.inputTok = null;
  turn.outputTok = null;
  turn.ctxMax = null;
  turn.inputPer1M = null;
  turn.outputPer1M = null;
  turn.cost = null;
  turn.elapsedMs = null;
  const b = body();
  if (b) {
    b.innerHTML = "";
    b.appendChild(el("div", "reason-panel-empty", EMPTY_HTML));
  }
  setReasoningStreaming(false);
}

export function appendReasoningPanel(text, replace) {
  ensureInfoBlock();
  // Nouveau contenu en streaming : la traduction affichée serait périmée.
  clearTranslation();
  const b = body();
  if (!b) return;
  let t = b.querySelector(".reason-text");
  if (!t) {
    t = el("div", "reason-text");
    b.appendChild(t);
  }
  // Rendu incremental : on n'ajoute QUE le nouveau morceau au noeud texte
  // (appendData). Jamais de textContent sur tout le texte -> pas de O(n)
  // par delta quand le raisonnement est long (technique Marexcode).
  if (replace === true) {
    t.textContent = String(text);
  } else {
    let node = t.firstChild;
    if (!node || node.nodeType !== 3) {
      t.textContent = "";
      node = document.createTextNode("");
      t.appendChild(node);
    }
    node.appendData(String(text));
  }
  if (!userClosed) openReasonPanel();
  schedulePanelScroll();
  setReasoningStreaming(true);
}

// Le backend envoie les tokens en fin de generation (stats). Ne cree jamais
// le bloc : sans raisonnement (Thinking desactive), le panneau reste vide.
export function updateReasonUsage(stats) {
  stats = stats || {};
  if (stats.prompt_tokens != null) turn.inputTok = stats.prompt_tokens;
  if (stats.completion_tokens != null) turn.outputTok = stats.completion_tokens;
  const b = body();
  if (b && b.querySelector(".reason-info")) renderUsage();
}

// Fin de la phase de raisonnement : le panneau se masque automatiquement.
// Si le mode auto DeepThink est actif, la traduction démarre 5 secondes après.
export function finishReasoning() {
  setReasoningStreaming(false);
  stopTick();
  renderTime();
  scheduleAutoTranslate();
  closeReasonPanel(false);
}

// Fin du tour : fige le temps et les stats, retourne un instantane
// reutilisable par le bouton "Raisonnement" de la requete.
export function finalizeReasonTurn(elapsedMs) {
  stopTick();
  if (elapsedMs != null) turn.elapsedMs = elapsedMs;
  renderTime();
  renderUsage();
  const b = body();
  const t = b && b.querySelector(".reason-text");
  // Le texte du raisonnement n'est jamais remplacé (la traduction s'affiche
  // dans un bloc séparé) : l'instantané contient toujours l'original.
  const snapText = t ? t.textContent : "";
  return {
    modelLabel: turn.modelLabel,
    startTs: turn.startTs,
    inputTok: turn.inputTok,
    outputTok: turn.outputTok,
    ctxMax: turn.ctxMax,
    inputPer1M: turn.inputPer1M,
    outputPer1M: turn.outputPer1M,
    elapsedMs: turn.elapsedMs,
    text: snapText,
  };
}

// Re-affiche un instantane precedent (bouton "Raisonnement").
export function restoreReasonSnapshot(snap) {
  if (!snap) return;
  resetReasonPanel();
  turn.modelLabel = snap.modelLabel || "";
  turn.inputTok = snap.inputTok;
  turn.outputTok = snap.outputTok;
  turn.ctxMax = snap.ctxMax;
  turn.inputPer1M = snap.inputPer1M;
  turn.outputPer1M = snap.outputPer1M;
  turn.elapsedMs = snap.elapsedMs;
  turn.startTs = snap.startTs || Date.now();
  ensureInfoBlock();
  renderUsage();
  renderTime();
  const b = body();
  const t = b && b.querySelector(".reason-text");
  if (t) t.textContent = snap.text || "";
  userClosed = false;
  openReasonPanel();
}

// Re-ouvre le panneau pour le tour en cours (annule une fermeture manuelle).
export function reopenReasonPanel() {
  userClosed = false;
  openReasonPanel();
}

export function setReasoningStreaming(on) {
  const sp = document.getElementById("reason-spinner");
  if (sp) sp.style.display = on ? "" : "none";
}
