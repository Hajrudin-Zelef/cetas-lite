// Panneau "Agentic" des paramètres : choix du style d'affichage de la vue
// Agents (Harness / OpenCode). Sauvegarde immédiate dans le localStorage
// (mx.agentic.style), application instantanée sans rechargement. Le chat
// général Cetas n'est pas concerné.

import { getAgenticStyle, setAgenticStyle } from "./agentic-style.js";

function el(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

const STYLE_LABELS = {
  harness: "Harness",
  opencode: "OpenCode",
};

const STYLE_DESCR = {
  harness:
    "Rendu DeepSeek Harness : lignes d'outils compactes « ✨ Tool call · " +
    "Nom · détail », loader Marex pendant l'attente, barre de statut " +
    "tours / outils / tokens.",
  opencode:
    "Reproduction du TUI opencode : messages et outils en blocs à bordure " +
    "gauche épaisse, en-têtes « Nom: paramètres », résultats bornés à " +
    "10 lignes, statuts exacts, diffs aux couleurs du TUI.",
};

function setSaveState(text) {
  const s = document.getElementById("agentic-save-state");
  if (s) s.textContent = text || "";
}

function render(bodyEl) {
  bodyEl.innerHTML = "";

  const title = el("div", "dt-row-title", "Style d'affichage de la vue Agents");
  bodyEl.appendChild(title);

  const seg = el("div", "ms-seg");
  seg.style.marginLeft = "0";
  const refresh = () => {
    const cur = getAgenticStyle();
    seg.querySelectorAll(".ms-seg-btn").forEach((b) => {
      b.classList.toggle("active", b.dataset.style === cur);
    });
  };
  for (const s of ["harness", "opencode"]) {
    const b = el("button", "ms-seg-btn", STYLE_LABELS[s]);
    b.type = "button";
    b.dataset.style = s;
    b.addEventListener("click", () => {
      if (getAgenticStyle() === s) return;
      setAgenticStyle(s);
      refresh();
      setSaveState("Enregistré ✓");
    });
    seg.appendChild(b);
  }
  bodyEl.appendChild(seg);
  refresh();

  for (const s of ["harness", "opencode"]) {
    const p = el("p", "apikey-intro", STYLE_LABELS[s] + " — " + STYLE_DESCR[s]);
    p.style.marginTop = "8px";
    bodyEl.appendChild(p);
  }

  const note = el(
    "p",
    "apikey-intro",
    "Ce réglage concerne uniquement la vue Agents : il s'applique aux " +
      "nouveaux messages sans rechargement. Le chat général Cetas garde " +
      "son affichage actuel."
  );
  note.style.marginTop = "14px";
  bodyEl.appendChild(note);
}

export function loadAgenticPanel() {
  const bodyEl = document.getElementById("agentic-body");
  if (!bodyEl || bodyEl.dataset.loaded) return;
  bodyEl.dataset.loaded = "1";
  render(bodyEl);
}
