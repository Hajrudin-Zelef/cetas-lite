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
  b.appendChild(el("div", "reason-sec", "REASONING"));
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
  if (!info) return;
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
export function finishReasoning() {
  setReasoningStreaming(false);
  stopTick();
  renderTime();
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
  return {
    modelLabel: turn.modelLabel,
    startTs: turn.startTs,
    inputTok: turn.inputTok,
    outputTok: turn.outputTok,
    ctxMax: turn.ctxMax,
    inputPer1M: turn.inputPer1M,
    outputPer1M: turn.outputPer1M,
    elapsedMs: turn.elapsedMs,
    text: t ? t.textContent : "",
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
