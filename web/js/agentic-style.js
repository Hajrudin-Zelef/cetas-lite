// Module "Agentic" : choix du style d'affichage de la vue Agents.
//
// Harness  : rendu DeepSeek Harness (lignes d'outils compactes, icônes
//            sobres par outil, loader Marex, statuts tours/outils/tokens).
// OpenCode : reproduction du TUI opencode (blocs à bordure gauche épaisse,
//            en-têtes "Nom: paramètres", résultats bornés à 10 lignes,
//            statuts exacts, diffs aux couleurs du TUI).
// Codex    : reproduction du TUI Codex (lignes "• Ran/Running <cmd>",
//            pastille verte/rouge selon le code de sortie, sorties en
//            5 premières + 5 dernières lignes, trace des décisions).
//
// Concerne UNIQUEMENT la vue Agents (ThreadView construit avec
// opts.agentic === true). Le chat général Cetas n'utilise jamais ce module.
//
// Persistance : localStorage "mx.agentic.style" (même préfixe que les autres
// réglages Agents). Défaut : "harness". Tout changement émet
// AGENTIC_STYLE_EVENT (appliqué immédiatement, sans rechargement).

export const AGENTIC_STYLE_EVENT = "cetas:agentic-style-changed";
export const AGENTIC_STYLE_HARNESS = "harness";
export const AGENTIC_STYLE_OPENCODE = "opencode";
export const AGENTIC_STYLE_CODEX = "codex";
export const AGENTIC_STYLES = [
  AGENTIC_STYLE_HARNESS,
  AGENTIC_STYLE_OPENCODE,
  AGENTIC_STYLE_CODEX,
];

const KEY = "mx.agentic.style";

export function getAgenticStyle() {
  try {
    const v = localStorage.getItem(KEY);
    if (AGENTIC_STYLES.includes(v)) return v;
  } catch (e) {
    /* stockage indisponible : défaut */
  }
  return AGENTIC_STYLE_HARNESS;
}

export function setAgenticStyle(style) {
  const v = AGENTIC_STYLES.includes(style) ? style : AGENTIC_STYLE_HARNESS;
  try {
    localStorage.setItem(KEY, v);
  } catch (e) {
    /* stockage indisponible : on propage quand même le choix */
  }
  window.dispatchEvent(
    new CustomEvent(AGENTIC_STYLE_EVENT, { detail: { style: v } })
  );
  return v;
}
