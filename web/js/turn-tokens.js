// Indicateurs de tokens : ligne au-dessus du champ de saisie
// ("↑ input ↓ output Tokens · Coût estimé") et compteur
// "tokens utilisés / contexte max" dans le champ.
import { api } from "./api.js";
import { currentSelection, getFamilies } from "./model-select.js";

let lastIn = 0;
let lastOut = 0;
let inputEst = 0;
const infoCache = {};

export function estimateTokens(text) {
  return Math.ceil(String(text || "").length / 4);
}

export function fmtCtx(n) {
  n = Number(n) || 0;
  if (n >= 1000) return (n / 1000).toFixed(n >= 10000 ? 0 : 1).replace(/\.0$/, "") + "K";
  return String(n);
}

function selectedMember() {
  try {
    const sel = currentSelection();
    const f = getFamilies().find((x) => x.id === sel.family);
    const m = f && f.modes.find((x) => x.mode === sel.mode);
    return (m && m.pool && m.pool[0]) || null;
  } catch (e) {
    return null;
  }
}

async function getModelInfo() {
  const member = selectedMember();
  if (!member) return { member: null, info: null };
  const key = member.provider + "/" + member.model;
  if (!(key in infoCache)) {
    try {
      infoCache[key] = await api(
        "/api/model-info?provider=" + encodeURIComponent(member.provider) + "&model=" + encodeURIComponent(member.model)
      );
    } catch (e) {
      infoCache[key] = null;
    }
  }
  return { member, info: infoCache[key] };
}

function formatCost(inTok, outTok, member, info) {
  let pin = member && member.input_per_1m != null ? member.input_per_1m : null;
  let pout = member && member.output_per_1m != null ? member.output_per_1m : null;
  if (pin == null && info) {
    pin = info.input_per_1m;
    pout = info.output_per_1m;
  }
  if (pin == null || pout == null) return "n/a";
  const cost = (inTok / 1e6) * pin + (outTok / 1e6) * pout;
  return "$" + cost.toFixed(4);
}

async function refreshCost() {
  const costEl = document.getElementById("turn-token-cost");
  if (!costEl) return;
  const { member, info } = await getModelInfo();
  costEl.textContent = "Coût estimé : " + formatCost(lastIn, lastOut, member, info);
}

export function getLastTokens() {
  return { in: lastIn, out: lastOut };
}

export function updateTokenLine() {
  const line = document.getElementById("turn-token-line");
  const ioEl = document.getElementById("turn-token-io");
  if (!line || !ioEl) return;
  ioEl.textContent = "↑ " + lastIn.toLocaleString("fr") + " (input) ↓ " + lastOut.toLocaleString("fr") + " (output) Tokens";
  line.style.display = "";
  refreshCost();
}

// Appelé à chaque frappe : estimation live des tokens d'entrée.
export function setInputEstimate(n) {
  inputEst = n;
  lastIn = n;
  updateTokenLine();
}

// Appelé à chaque événement stats du tour en cours.
export function setTurnStats(inTok, outTok) {
  if (inTok != null) lastIn = inTok;
  if (outTok != null) lastOut = outTok;
  updateTokenLine();
  refreshCtxCounter();
}

// Compteur "utilisés / contexte max" à droite dans le champ de saisie.
export async function refreshCtxCounter() {
  const ctr = document.getElementById("ctx-counter");
  if (!ctr) return;
  const { info } = await getModelInfo();
  const used = lastIn + lastOut;
  if (info && info.context_window) {
    const pct = Math.min(100, Math.round((used * 100) / info.context_window));
    const color = pct >= 90 ? "var(--danger,#ef4444)" : pct >= 70 ? "var(--warning,#eab308)" : "var(--success,#22c55e)";
    ctr.innerHTML =
      '<span style="color:' + color + '">' + fmtCtx(used) + " / " + fmtCtx(info.context_window) + "</span>";
    ctr.title = used.toLocaleString("fr") + " / " + info.context_window.toLocaleString("fr") + " tokens (" + pct + "%)";
  } else {
    ctr.textContent = fmtCtx(used) + " / —";
    ctr.title = "Contexte max inconnu pour ce modèle";
  }
  ctr.style.display = "";
}

// A appeler quand la sélection du modèle change.
export function refreshModelMeta() {
  // le cache est par modèle ; on rafraîchit coût + compteur
  lastIn = inputEst;
  updateTokenLine();
  refreshCtxCounter();
}

export function resetTurnTokens() {
  lastIn = 0;
  lastOut = 0;
  const line = document.getElementById("turn-token-line");
  if (line) line.style.display = "none";
}
