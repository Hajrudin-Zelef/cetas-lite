// Panneau lateral droit "Raisonnement" : affiche le raisonnement du modele
// en streaming. S'ouvre automatiquement au premier delta sauf si
// l'utilisateur l'a ferme manuellement pendant le tour en cours.

function el(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

const EMPTY_HTML = "Le raisonnement du modèle s'affichera ici pendant la génération.";

let userClosed = false;

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

function reopenBtn() {
  return document.getElementById("reason-panel-reopen");
}

function hasContent() {
  const b = body();
  const t = b && b.querySelector(".reason-text");
  return !!(t && t.textContent.trim());
}

export function initReasonPanel() {
  const closeBtn = document.getElementById("reason-panel-close");
  if (closeBtn) closeBtn.addEventListener("click", () => closeReasonPanel(true));
  const ro = reopenBtn();
  if (ro) {
    ro.addEventListener("click", () => {
      userClosed = false;
      openReasonPanel();
    });
  }
}

export function openReasonPanel() {
  const p = panel();
  if (!p) return;
  p.classList.add("open");
  const ro = reopenBtn();
  if (ro) ro.style.display = "none";
}

export function closeReasonPanel(manual) {
  if (manual) userClosed = true;
  const p = panel();
  if (p) p.classList.remove("open");
  const ro = reopenBtn();
  if (ro) ro.style.display = hasContent() ? "" : "none";
}

export function isReasonPanelOpen() {
  const p = panel();
  return !!p && p.classList.contains("open");
}

export function resetReasonPanel() {
  userClosed = false;
  const b = body();
  if (b) {
    b.innerHTML = "";
    b.appendChild(el("div", "reason-panel-empty", EMPTY_HTML));
  }
  setReasoningStreaming(false);
  const ro = reopenBtn();
  if (ro) ro.style.display = "none";
}

export function appendReasoningPanel(text, replace) {
  const b = body();
  if (!b) return;
  let t = b.querySelector(".reason-text");
  if (!t) {
    b.innerHTML = "";
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
  else {
    const ro = reopenBtn();
    if (ro && hasContent()) ro.style.display = "";
  }
  schedulePanelScroll();
  setReasoningStreaming(true);
}

export function finishReasoningPanel() {
  setReasoningStreaming(false);
  const ro = reopenBtn();
  if (ro) ro.style.display = hasContent() && !isReasonPanelOpen() ? "" : "none";
}

export function setReasoningStreaming(on) {
  const sp = document.getElementById("reason-spinner");
  if (sp) sp.style.display = on ? "" : "none";
}
