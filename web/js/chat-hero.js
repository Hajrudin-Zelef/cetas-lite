// Hero façon Gemini affiché quand le chat est vide.
//
// Le VRAI composer (.input-area, avec son indicateur de sélection et son
// alerte modèle) est déplacé dans le hero quand la conversation est vide,
// et remis en bas dès le premier message : tous les boutons (menu +,
// micro, thinking, pièces jointes…) continuent de fonctionner car ce sont
// les mêmes nœuds DOM.
//
// Piloté par un MutationObserver sur #chat-container : robuste quel que
// soit le chemin (reset, nouvelle session, premier message, rejeu
// d'historique). Sécurité : si le hero est détruit (innerHTML) alors que le
// composer est dedans, les nœuds survivent via les références gardées ici
// et sont réinsérés dans le nouveau slot avant le prochain paint
// (microtask).

const HERO_PLACEHOLDER = "Demander à Cetas";

let movables = []; // [{ node, parent, next }]
let promptInput = null;
let defaultPlaceholder = null;
let heroOn = false;
let observer = null;
let observedLog = null;

// Le slot n'est pris en compte que s'il est dans le log (le hero affiché).
function slotInLog() {
  const log = document.getElementById("chat-container");
  const slot = document.getElementById("hero-composer-slot");
  return log && slot && log.contains(slot) ? slot : null;
}

export function isHeroMode() {
  return heroOn;
}

function placeIn(slot) {
  for (const m of movables) slot.appendChild(m.node);
}

function placeHome() {
  // Ordre inverse : chaque ancre `next` doit déjà être revenue dans le
  // parent au moment de l'insertion (sinon NotFoundError).
  for (let i = movables.length - 1; i >= 0; i--) {
    const m = movables[i];
    const ref = m.next && m.next.parentNode === m.parent ? m.next : null;
    m.parent.insertBefore(m.node, ref);
  }
}

function placedIn(slot) {
  return movables.length > 0 && movables.every((m) => slot.contains(m.node));
}

export function updateHeroMode() {
  if (movables.length === 0) return;
  const slot = slotInLog();
  // Après un reset() (innerHTML), le slot est nouveau et le composer a été
  // détaché avec l'ancien : il faut le réinsérer même si heroOn est vrai.
  const inSlot = !!slot && placedIn(slot);
  if (slot && (!heroOn || !inSlot)) {
    // Passage en mode hero : le composer remonte au centre.
    if (!inSlot) placeIn(slot);
    document.body.classList.add("hero-mode");
    if (promptInput) {
      if (defaultPlaceholder === null) {
        defaultPlaceholder = promptInput.getAttribute("placeholder") || "";
      }
      promptInput.setAttribute("placeholder", HERO_PLACEHOLDER);
    }
    heroOn = true;
  } else if (!slot && heroOn) {
    // Premier message : le composer redescend en bas.
    placeHome();
    document.body.classList.remove("hero-mode");
    if (promptInput && defaultPlaceholder !== null) {
      promptInput.setAttribute("placeholder", defaultPlaceholder);
    }
    heroOn = false;
  }
}

export function initChatHero() {
  const log = document.getElementById("chat-container");
  const area = document.querySelector(".input-area");
  if (!area || !log) return;
  if (observer && observedLog === log) {
    updateHeroMode();
    return;
  }
  // (Ré)init : nouveau log (tests) ou premier appel.
  if (observer) observer.disconnect();
  // Ordre dans le slot : alerte puis composer (#input-hint est déjà dans
  // .input-area et voyage avec lui).
  movables = [];
  const alert = document.getElementById("model-alert");
  if (alert) movables.push({ node: alert, parent: alert.parentNode, next: alert.nextSibling });
  movables.push({ node: area, parent: area.parentNode, next: area.nextSibling });
  promptInput = document.getElementById("prompt-input");
  heroOn = false;
  observer = new MutationObserver(() => updateHeroMode());
  observer.observe(log, { childList: true });
  observedLog = log;
  updateHeroMode();
}
