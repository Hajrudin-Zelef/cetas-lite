// Panneau lateral droit "Raisonnement" : affiche les informations de la
// requete (modele, temps de generation, tokens, contexte, cout) puis le
// raisonnement du modele en streaming. HISTORIQUE : le panneau conserve UN
// BLOC PAR TOUR (.reason-turn) — le raisonnement d'un nouveau tour s'ajoute
// EN DESSOUS des precedents, il ne les remplace jamais. Seule une nouvelle
// conversation vide le panneau. S'ouvre automatiquement au premier delta
// sauf si l'utilisateur l'a ferme manuellement pendant le tour en cours.
// Quand le raisonnement est termine, le panneau se masque automatiquement ;
// chaque reponse garde son bouton "Raisonnement" pour reconsulter son tour
// (defilement vers le bloc dans l'historique).

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

// --- Historique des tours --------------------------------------------------
// Le panneau conserve UN BLOC PAR TOUR (.reason-turn) : le raisonnement d'un
// nouveau tour s'ajoute EN DESSOUS des précédents, il ne les remplace jamais.
// Seule une nouvelle conversation (reset) vide le panneau.
let turnSeq = 0;

// Dernier bloc de tour, scellé ou non.
function lastTurnEl() {
  const b = body();
  const turns = b ? b.querySelectorAll(".reason-turn") : [];
  return turns.length ? turns[turns.length - 1] : null;
}

// Bloc du tour en cours (non scellé), ou null s'il n'y en a pas.
function currentTurnEl() {
  const te = lastTurnEl();
  return te && !te.hasAttribute("data-sealed") ? te : null;
}

// --- Traduction DeepThink ---
// Le bouton "Traduire" ne remplace JAMAIS le raisonnement : la traduction
// s'affiche SOUS l'original (bloc .reason-translation), dans le bloc du tour
// concerné. Chaque tour a son propre bouton : un clic traduit CE tour
// immédiatement ET active le mode auto ; chaque raisonnement suivant est
// alors traduit automatiquement 5 secondes après sa fin. Un second clic
// désactive le mode auto. Le mode auto persiste (localStorage).
const DT_AUTO_KEY = "cetas.deepthink.auto";
let dtAutoDelayMs = 5000;

// Réglable pour les tests (évite d'attendre 5s réelles).
export function _setAutoTranslateDelayForTests(ms) {
  dtAutoDelayMs = ms;
}

// Tours en cours de traduction (un par bloc : deux tours peuvent traduire
// en parallèle sans se bloquer mutuellement).
const translatingTurns = new Set();

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

// État des boutons selon le mode auto (à la création d'un bloc et à chaque
// changement de mode). Les boutons des tours en cours de traduction
// affichent "…" et restent désactivés.
function refreshTranslateBtns() {
  const auto = isAutoTranslate();
  document.querySelectorAll("#reason-panel-body .reason-translate-btn").forEach((btn) => {
    const te = btn.closest(".reason-turn");
    if (te && translatingTurns.has(te)) return;
    if (auto) {
      btn.textContent = "🌐 Auto ✓";
      btn.classList.add("on");
      btn.title = "Traduction automatique activée — cliquer pour désactiver";
    } else {
      btn.textContent = "🌐 Traduire";
      btn.classList.remove("on");
      btn.title = "Traduire le raisonnement + activer la traduction auto (DeepThink Global)";
    }
    btn.disabled = false;
  });
}

// Supprime le bloc de traduction du tour en cours (nouveau contenu en
// streaming : la traduction serait périmée). Ne touche jamais aux tours
// scellés : leur traduction fait partie de l'historique.
function clearTranslation() {
  const te = currentTurnEl();
  const box = te && te.querySelector(".reason-translation");
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

// Traduit le raisonnement d'un tour et ajoute la traduction SOUS l'original
// (jamais de remplacement), dans le bloc de ce tour. Rouvre le panneau sauf
// fermeture manuelle.
async function translateTurn(te) {
  if (!te || translatingTurns.has(te)) return;
  const t = te.querySelector(".reason-text");
  if (!t) return;
  const src = (t.textContent || "").trim();
  if (!src || te.querySelector(".reason-translation")) return;
  translatingTurns.add(te);
  const btn = te.querySelector(".reason-translate-btn");
  if (btn) {
    btn.disabled = true;
    btn.textContent = "…";
  }
  try {
    const data = await api("/api/deepthink/translate", { method: "POST", body: { text: src } });
    const translated = data && typeof data.translation === "string" ? data.translation.trim() : "";
    if (!translated) throw new Error("Traduction vide.");
    // Le bloc a pu être supprimé pendant l'appel (nouvelle conversation) :
    // on ne touche que le même nœud texte, sinon on annule.
    if (te.isConnected && te.querySelector(".reason-text") === t) {
      te.appendChild(translationBlock(translated, (data && data.lang) || "", false));
      if (!userClosed) openReasonPanel();
      schedulePanelScroll();
    }
  } catch (e) {
    if (te.isConnected && te.querySelector(".reason-text") === t) {
      te.appendChild(translationBlock((e && e.message) || "Échec de la traduction", "", true));
    }
  } finally {
    translatingTurns.delete(te);
    refreshTranslateBtns();
  }
}

// Clic sur le bouton d'un tour : OFF -> traduit CE tour maintenant + active
// l'auto ; ON -> désactive l'auto.
async function toggleReasonTranslation(e) {
  const btn = e && e.currentTarget;
  const te = btn && btn.closest(".reason-turn");
  if (!te || translatingTurns.has(te)) return;
  if (isAutoTranslate()) {
    setAutoTranslate(false);
    refreshTranslateBtns();
    return;
  }
  setAutoTranslate(true);
  refreshTranslateBtns();
  await translateTurn(te);
}

// Fin d'un raisonnement : si l'auto est actif, traduction automatique de CE
// tour 5 secondes après. Vérification d'identité au déclenchement : si le
// bloc a été supprimé ou son texte a changé entre-temps, on annule.
function scheduleAutoTranslate(te) {
  if (!te) return;
  if (te._autoTimer) {
    clearTimeout(te._autoTimer);
    te._autoTimer = 0;
  }
  if (!isAutoTranslate()) return;
  const t = te.querySelector(".reason-text");
  const src = t ? (t.textContent || "").trim() : "";
  if (!src || te.querySelector(".reason-translation")) return;
  te._autoTimer = setTimeout(() => {
    te._autoTimer = 0;
    if (!te.isConnected) return;
    const t2 = te.querySelector(".reason-text");
    if (!t2 || (t2.textContent || "").trim() !== src) return;
    translateTurn(te);
  }, dtAutoDelayMs);
  if (te._autoTimer && typeof te._autoTimer.unref === "function") te._autoTimer.unref();
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

// Construit le bloc du tour en cours (.reason-turn : infos + en-tête
// REASONING + texte) s'il n'existe pas, et le retourne. Chaque nouveau tour
// crée son propre bloc EN DESSOUS des précédents : l'historique n'est jamais
// écrasé. Retourne null si le panneau est absent.
function ensureTurnBlock() {
  let te = currentTurnEl();
  if (te) return te;
  const b = body();
  if (!b) return null;
  // Retire l'état vide ("Le raisonnement du modèle s'affichera ici...")
  const empty = b.querySelector(".reason-panel-empty");
  if (empty) empty.remove();
  turnSeq += 1;
  te = el("div", "reason-turn");
  te.dataset.turnId = String(turnSeq);
  const info = el("div", "reason-info");
  info.appendChild(riRow("Modèle :", "ri-model", turn.modelLabel || "—", true));
  info.appendChild(riRow("Temps de génération :", "ri-time", "0s", false));
  info.appendChild(riRow("Input :", "ri-input", "—", false));
  info.appendChild(riRow("Output - t/s :", "ri-output", "—", false));
  info.appendChild(riRow("Contexte :", "ri-ctx", "—", false));
  info.appendChild(riRow("Coût :", "ri-cost", "—", false));
  te.appendChild(info);
  const sec = el("div", "reason-sec");
  sec.appendChild(el("span", "reason-sec-label", "REASONING"));
  const tbtn = el("button", "reason-translate-btn", "🌐 Traduire");
  tbtn.type = "button";
  tbtn.title = "Traduire le raisonnement (DeepThink Global)";
  tbtn.addEventListener("click", toggleReasonTranslation);
  sec.appendChild(tbtn);
  te.appendChild(sec);
  te.appendChild(el("div", "reason-text"));
  b.appendChild(te);
  refreshTranslateBtns(); // reflète le mode auto persisté (après attachement au DOM)
  startTick();
  // Le modele et ses infos ont pu arriver avant le premier delta.
  renderUsage(te);
  return te;
}

function riRow(label, valueCls, value, accent) {
  const row = el("div", "ri-row");
  row.appendChild(el("span", "ri-k", label + " "));
  const v = el("span", valueCls + (accent ? " ri-accent" : ""), value);
  row.appendChild(v);
  return row;
}

function setRow(te, cls, text, html) {
  const e = te && te.querySelector("." + cls);
  if (!e) return;
  if (html !== undefined) e.innerHTML = html;
  else e.textContent = text;
}

function renderTime(te) {
  te = te || currentTurnEl() || lastTurnEl();
  const ms = turn.elapsedMs != null ? turn.elapsedMs : Date.now() - turn.startTs;
  setRow(te, "ri-time", fmtSecs(Math.max(0, ms)));
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

function renderUsage(te) {
  te = te || currentTurnEl() || lastTurnEl();
  if (turn.inputTok != null) setRow(te, "ri-input", turn.inputTok.toLocaleString("fr") + " tok");
  if (turn.outputTok != null) {
    let txt = turn.outputTok.toLocaleString("fr") + " tok";
    const ms = turn.elapsedMs != null ? turn.elapsedMs : Date.now() - turn.startTs;
    if (ms > 0) txt += " (" + Math.round((turn.outputTok / ms) * 1000) + "/s)";
    setRow(te, "ri-output", txt);
  }
  if (turn.inputTok != null || turn.outputTok != null) {
    const used = (turn.inputTok || 0) + (turn.outputTok || 0);
    if (turn.ctxMax) {
      const pct = Math.min(100, Math.round((used / turn.ctxMax) * 100));
      const color = pct >= 90 ? "#ef4444" : pct >= 70 ? "#eab308" : "#22c55e";
      setRow(
        te,
        "ri-ctx",
        undefined,
        '<span style="color:' + color + '">' + fmtTok(used) + " / " + fmtTok(turn.ctxMax) + "</span> · " + pct + "% utilisé"
      );
    } else {
      setRow(te, "ri-ctx", fmtTok(used) + " tok");
    }
  }
  if (turn.cost != null) setRow(te, "ri-cost", "$" + turn.cost.toFixed(4));
  else if (turn.inputTok != null && turn.inputPer1M != null) {
    const c = (turn.inputTok / 1e6) * turn.inputPer1M + ((turn.outputTok || 0) / 1e6) * (turn.outputPer1M || 0);
    setRow(te, "ri-cost", "$" + c.toFixed(4));
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
  // Panneau vide (seul l'espace reserve est present) => on ne l'ouvre pas.
  const b = body();
  if (!b || !b.querySelector(".reason-turn")) return;
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
// Le bloc du tour précédent a été scellé par sealReasonTurn() : le prochain
// delta créera un nouveau bloc EN DESSOUS.
export function beginReasonTurn(info) {
  info = info || {};
  stopTick();
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
  const te = currentTurnEl();
  if (te) setRow(te, "ri-model", turn.modelLabel || "—");
}

// Nouvelle conversation : vide TOUT le panneau (seul cas où l'historique
// est effacé).
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
    b.querySelectorAll(".reason-turn").forEach((te) => {
      if (te._autoTimer) {
        clearTimeout(te._autoTimer);
        te._autoTimer = 0;
      }
    });
    translatingTurns.clear();
    b.innerHTML = "";
    b.appendChild(el("div", "reason-panel-empty", EMPTY_HTML));
  }
  // Corps vidé => rien à y afficher : on referme (plus de placeholder
  // permanent à l'écran).
  const p = panel();
  if (p) p.classList.remove("open");
  setReasoningStreaming(false);
}

// Nouveau tour (nouveau message utilisateur) : fige le tour précédent
// (temps final, bloc scellé) SANS vider le panneau — le prochain
// raisonnement créera son bloc en dessous. Idempotent.
export function sealReasonTurn() {
  stopTick();
  const te = lastTurnEl();
  if (te) {
    renderTime(te);
    te.setAttribute("data-sealed", "1");
  }
  setReasoningStreaming(false);
}

// Le serveur demande d'abandonner le raisonnement du tour en cours
// (ex. régénération) : on retire le bloc de CE tour, pas tout l'historique.
export function dropCurrentReasonTurn() {
  stopTick();
  const te = currentTurnEl();
  if (te) {
    if (te._autoTimer) {
      clearTimeout(te._autoTimer);
      te._autoTimer = 0;
    }
    translatingTurns.delete(te);
    te.remove();
  }
  const b = body();
  if (b && !b.querySelector(".reason-turn")) {
    b.innerHTML = "";
    b.appendChild(el("div", "reason-panel-empty", EMPTY_HTML));
    const p = panel();
    if (p) p.classList.remove("open");
  }
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
  setReasoningStreaming(false);
}

export function appendReasoningPanel(text, replace) {
  const te = ensureTurnBlock();
  if (!te) return;
  // Nouveau contenu en streaming : la traduction affichée serait périmée.
  clearTranslation();
  let t = te.querySelector(".reason-text");
  if (!t) {
    t = el("div", "reason-text");
    te.appendChild(t);
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
  renderUsage();
}

// Fin de la phase de raisonnement : scelle le bloc du tour (il reste dans
// l'historique) et masque automatiquement le panneau.
// Si le mode auto DeepThink est actif, la traduction de CE tour démarre
// 5 secondes après.
export function finishReasoning() {
  const te = lastTurnEl();
  stopTick();
  if (te) {
    renderTime(te);
    te.setAttribute("data-sealed", "1");
  }
  setReasoningStreaming(false);
  scheduleAutoTranslate(te);
  closeReasonPanel(false);
}

// Fin du tour : fige le temps et les stats, retourne un instantane
// reutilisable par le bouton "Raisonnement" de la reponse.
export function finalizeReasonTurn(elapsedMs) {
  stopTick();
  if (elapsedMs != null) turn.elapsedMs = elapsedMs;
  const te = lastTurnEl();
  renderTime(te);
  renderUsage(te);
  const t = te && te.querySelector(".reason-text");
  // Le texte du raisonnement n'est jamais remplacé (la traduction s'affiche
  // dans un bloc séparé) : l'instantané contient toujours l'original.
  const snapText = t ? t.textContent : "";
  return {
    turnId: te && te.dataset ? te.dataset.turnId || null : null,
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

// Fait défiler le panneau jusqu'au bloc d'un tour et le signale brièvement.
function scrollTurnIntoView(te) {
  const b = body();
  try {
    if (typeof te.scrollIntoView === "function") te.scrollIntoView({ block: "start" });
    else if (b) b.scrollTop = te.offsetTop;
  } catch (e) {
    if (b) b.scrollTop = te.offsetTop;
  }
  te.classList.add("reason-flash");
  const flashTimer = setTimeout(() => te.classList.remove("reason-flash"), 1600);
  if (flashTimer && typeof flashTimer.unref === "function") flashTimer.unref();
}

// Re-affiche un tour precedent (bouton "Raisonnement"). Cas courant : le
// tour est déjà dans l'historique du panneau — on fait simplement défiler
// jusqu'à son bloc (l'historique n'est jamais vidé pour reconsulter un
// tour). Repli : le tour n'y est plus (ex. session restaurée) — on
// reconstruit un bloc scellé unique avec l'instantané.
export function restoreReasonSnapshot(snap) {
  if (!snap) return;
  const b = body();
  const te = snap.turnId && b ? b.querySelector('.reason-turn[data-turn-id="' + snap.turnId + '"]') : null;
  if (te) {
    userClosed = false;
    openReasonPanel();
    scrollTurnIntoView(te);
    return;
  }
  resetReasonPanel();
  turn.modelLabel = snap.modelLabel || "";
  turn.inputTok = snap.inputTok;
  turn.outputTok = snap.outputTok;
  turn.ctxMax = snap.ctxMax;
  turn.inputPer1M = snap.inputPer1M;
  turn.outputPer1M = snap.outputPer1M;
  turn.elapsedMs = snap.elapsedMs;
  turn.startTs = snap.startTs || Date.now();
  const te2 = ensureTurnBlock();
  if (!te2) return;
  renderUsage(te2);
  renderTime(te2);
  const t = te2.querySelector(".reason-text");
  if (t) t.textContent = snap.text || "";
  te2.setAttribute("data-sealed", "1");
  stopTick();
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
