import { api, getToken } from "./api.js";
import { ThreadView, el } from "./thread-view.js";
import { estimateTokens } from "./turn-tokens.js";
import { logout } from "./auth.js";
import { openDocs } from "./docs.js";
import {
  Projects,
  renderActiveTree,
  renderProjectsList,
  renderProjectBar,
  openNewProjectModal,
} from "./projects.js";

// ============================================================================
// Vue "Agents" façon Marexcode.
// Interface plein écran dédiée aux agents parallèles, calquée sur le front
// Marexcode : sidebar (retour, métriques, nouvelle conversation, recherche,
// projets, discussions, utilisateur), hero (orbe, titre, composer), panneau
// latéral "Raisonnement". Toute la logique est câblée sur les backends
// cetas-lite (/api/aliases, /api/agents, /api/metrics, /api/settings).
// ============================================================================

function esc(s) {
  return String(s == null ? "" : s)
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;");
}

const I = {
  plus: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>',
  chevron: '<svg class="chevron" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 9l6 6 6-6"/></svg>',
  chevSm: '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 9l6 6 6-6"/></svg>',
  close: '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>',
  close14: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>',
  star: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 2l2.5 5.5L20 9l-4 4 1 6-5-3-5 3 1-6-4-4 5.5-1.5z"/></svg>',
  star20: '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2l2.5 5.5L20 9l-4 4 1 6-5-3-5 3 1-6-4-4 5.5-1.5z"/></svg>',
  folder: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 2 2H5a2 2 0 0 1-2-2V7z"/></svg>',
  folder13: '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 2 2H5a2 2 0 0 1-2-2V7z"/></svg>',
  eye: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>',
  edit: '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 20h9"/><path d="M16.5 3.5a2.1 2.1 0 0 1 3 3L7 19l-4 1 1-4z"/></svg>',
  ask: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/></svg>',
  undo: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 10h10a5 5 0 0 1 0 10H12"/><path d="M3 10l4-4M3 10l4 4"/></svg>',
  redo: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 10H11a5 5 0 0 0 0 10h1"/><path d="M21 10l-4-4M21 10l-4 4"/></svg>',
  send: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2"><path d="M12 19V5M5 12l7-7 7 7"/></svg>',
  check: '<svg class="check" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4"><path d="M20 6L9 17l-5-5"/></svg>',
  stop: '<svg width="12" height="12" viewBox="0 0 24 24" fill="currentColor" stroke="none"><rect x="5" y="5" width="14" height="14" rx="2"/></svg>',
  upload: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>',
  burger: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/></svg>',
  logout: '<svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" y1="12" x2="9" y2="12"/></svg>',
  gear: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 1 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 1 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 1 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 1 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>',
  book: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>',
  faq: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><line x1="12" y1="17" x2="12.01" y2="17"/></svg>',
  globe: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>',
  think: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9.5 2a2.5 2.5 0 0 1 2.5 2.5c2.5 0 4.5 1.5 5.3 3.6a3.5 3.5 0 0 1-.6 6.9 3.5 3.5 0 0 1-2.7 3.4 3.5 3.5 0 0 1-6.3.6A3.5 3.5 0 0 1 4 15.5 3.5 3.5 0 0 1 4.6 8.6 4.5 4.5 0 0 1 9.5 2z"/></svg>',
};

const VIEW_HTML = `
<div class="mx-frame">
  <div class="mx-app">
    <div class="mx-mobile-bar">
      <button class="hamburger" id="mx-hamburger" aria-label="Menu">${I.burger}</button>
      <span class="mobile-title">Agents</span>
      <button class="mobile-new" id="mx-mobile-new" aria-label="Nouvelle conversation">${I.plus}</button>
    </div>

    <div class="mx-sidebar" id="mx-sidebar">
      <button class="sb-close" id="mx-sb-close" aria-label="Fermer">${I.close14}</button>
      <div class="sb-scroll">
        <button class="sb-back-btn" id="mx-back">
          <img src="/images/Cetas42.png" alt="" width="18" height="18" style="border-radius:4px">
          Retour à Cetas
        </button>

        <div class="sb-group-label">Metriques</div>
        <div class="sb-metrics" id="mx-metrics">
          <div class="sb-metric-line"><span class="sb-metric-dot" id="mx-m-cpu-dot"></span><span class="sb-metric-label">CPU :</span><span class="sb-metric-value" id="mx-m-cpu">—</span></div>
          <div class="sb-metric-line"><span class="sb-metric-dot" id="mx-m-ram-dot"></span><span class="sb-metric-label">RAM :</span><span class="sb-metric-value" id="mx-m-ram">—</span></div>
          <div class="sb-metric-line"><span class="sb-metric-dot" id="mx-m-disk-dot"></span><span class="sb-metric-label">Disk :</span><span class="sb-metric-value" id="mx-m-disk">—</span></div>
          <div class="sb-metric-line"><span class="sb-metric-dot" id="mx-m-net-dot"></span><span class="sb-metric-label">Réseau :</span><span class="sb-metric-value" id="mx-m-net">—</span></div>
        </div>
        <div class="sb-divider"></div>

        <button class="new-session-btn" id="mx-new">${I.plus} Nouvelle conversation</button>
        <div class="mx-search-wrap"><input class="mx-search-input" id="mx-search" placeholder="Rechercher des conversations" autocomplete="off"></div>
        <div class="sb-divider"></div>

        <button class="sb-group-toggle" id="mx-toggle-projects"><span class="sb-group-label">Projets</span>${I.chevSm}</button>
        <div class="sb-collapsible" id="mx-panel-projects">
          <div class="sb-project-header">
            <span id="mx-active-icon">${I.folder}</span><span id="mx-active-name">Espace partagé</span><span class="sb-active-badge" id="mx-active-badge" hidden>Actif</span>
          </div>
          <div class="sb-active-sub" id="mx-active-sub" hidden></div>
          <div id="mx-active-tree"></div>
          <div class="sb-mini-label" id="mx-projects-label">Tous les projets</div>
          <div class="sb-tree-empty" id="mx-projects-empty">Aucun projet importé.</div>
          <div id="mx-projects-list"></div>
          <button id="mx-new-project" class="sb-new-project-btn">+ Nouveau projet</button>
        </div>

        <div class="sb-divider"></div>
        <button class="sb-group-toggle" id="mx-toggle-disc"><span class="sb-group-label">Discussions</span>${I.chevSm}</button>
        <div class="sb-collapsible" id="mx-panel-disc">
          <div class="sb-tree-empty" id="mx-disc-empty">Aucune discussion.</div>
          <div id="mx-disc-list"></div>
        </div>
      </div>

      <div class="sb-user-wrap" id="mx-user-wrap">
        <div class="sb-user-menu" id="mx-user-menu">
          <button class="sb-user-menu-item" id="mx-menu-settings">${I.gear} Paramètre</button>
          <button class="sb-user-menu-item" id="mx-menu-help">${I.ask} Obtenir de l'aide</button>
          <button class="sb-user-menu-item" id="mx-menu-about">${I.book} En savoir plus</button>
          <button class="sb-user-menu-item" id="mx-menu-faq">${I.faq} FAQ</button>
          <div class="sb-user-menu-divider"></div>
          <button class="sb-user-menu-item danger" id="mx-logout">${I.logout} Déconnexion</button>
        </div>
        <div class="sb-user-row">
          <button class="sb-user-btn" id="mx-user-btn">
            <span class="sb-avatar" id="mx-user-avatar">U</span>
            <span id="mx-user-name">Utilisateur</span>
          </button>
          <button class="sb-docs-btn" id="mx-docs-btn" title="Centre d'aide" aria-label="Centre d'aide">${I.ask}</button>
        </div>
      </div>
    </div>

    <div class="mx-main">
      <div class="ocean-fx" aria-hidden="true"><div class="ocean-particles">
        <span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span><span class="ocean-particle"></span>
      </div></div>

      <div class="mx-hero" id="mx-hero">
        <div class="hero-loader" aria-hidden="true"><div class="ripple-loader">
          <div class="box"></div><div class="box"></div><div class="box"></div><div class="box"></div><div class="box"></div>
          <div class="logo"><svg width="44" height="44" viewBox="0 0 24 24" fill="none" stroke="#5EC8F2" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2l2.5 5.5L20 9l-4 4 1 6-5-3-5 3 1-6-4-4 5.5-1.5z"/></svg></div>
        </div></div>

        <div class="hero-title">
          <span class="hero-title-mark">${I.star20}</span>
          <h1>Agents</h1>
          <span class="badge-pill">Preview</span>
        </div>
        <div class="mx-agent-status" id="mx-status" style="display:none"><span class="dot"></span><span id="mx-status-text"></span></div>

        <div class="mx-todos" id="mx-todos" style="display:none" aria-live="polite"></div>
        <div class="mx-composer mode-build" id="mx-composer">
          <div class="composer-selectors">
            <span class="mx-mode-badge build" id="mx-mode-badge" title="Tab : basculer Plan / Build">BUILD</span>
            <div class="cdrop" id="mx-dd-project">
              <button class="selector" id="mx-btn-project">${I.folder13}<span id="mx-label-project">Espace partagé</span>${I.chevron}</button>
              <div class="cdrop-menu" id="mx-menu-project"></div>
            </div>
          </div>
          <div class="mx-pbar" id="mx-project-bar"></div>
          <textarea id="mx-input" rows="1" placeholder="Décrivez la tâche de code à réaliser... (glissez des images ici)"></textarea>
          <div class="mx-attach-row" id="mx-attach-row" style="display:none"></div>
          <div class="mx-vision-warn" id="mx-vision-warn" style="display:none"></div>
          <input type="file" id="mx-file-input" multiple hidden>
          <div class="composer-info-bar"><span id="mx-token-counter"></span><span class="mx-mode-hint">Tab : Plan / Build</span><span id="mx-queue" style="display:none"></span></div>
          <div class="composer-footer">
            <div class="composer-footer-left">
              <div class="cdrop" id="mx-dd-plus">
                <button class="plus-btn" id="mx-plus" aria-label="Ajouter" title="Ajouter">${I.plus}</button>
                <div class="cdrop-menu" id="mx-menu-plus"></div>
              </div>
              <button class="icon-btn" id="mx-web" title="Recherche web" aria-label="Recherche web">${I.globe}</button>
              <div class="cdrop" id="mx-dd-effort">
                <button class="selector" id="mx-btn-effort" title="Niveau d'effort de réflexion">${I.think}<span id="mx-label-effort">Défaut</span>${I.chevron}</button>
                <div class="cdrop-menu" id="mx-menu-effort"></div>
              </div>
              <button class="icon-btn mx-fav-btn" id="mx-fav" title="Ajouter aux favoris">${I.star}</button>
              <div class="cdrop" id="mx-dd-perm">
                <button class="selector" id="mx-btn-perm" style="padding:0">${I.edit}<span id="mx-label-perm">Espace Write</span>${I.chevron}</button>
                <div class="cdrop-menu" id="mx-menu-perm"></div>
              </div>
            </div>
            <div class="composer-footer-right">
              <button class="icon-btn" id="mx-undo" disabled title="Annuler">${I.undo}</button>
              <button class="icon-btn" id="mx-redo" disabled title="Rétablir">${I.redo}</button>
              <div class="cdrop" id="mx-dd-model">
                <button class="model-select" id="mx-btn-model"><span id="mx-label-model">Choisir un niveau</span>${I.chevron}</button>
                <div class="cdrop-menu up-right" id="mx-menu-model"></div>
              </div>
              <button class="stop-btn" id="mx-stop" aria-label="Arrêter" hidden>${I.stop}</button>
              <button class="send-btn" id="mx-send" aria-label="Envoyer">${I.send}</button>
            </div>
          </div>
        </div>

        <div class="mx-chat-panel" id="mx-chat-panel" style="display:none">
          <div class="mx-chat-log" id="mx-chat-log"></div>
        </div>
      </div>
    </div>

    <div class="mx-side-panel" id="mx-side-panel">
      <div class="side-panel-header">
        <span class="side-panel-header-title"><span class="think-spinner-sm" id="mx-sp-spinner" style="display:none"></span><span>Raisonnement</span></span>
        <button class="side-panel-close" id="mx-sp-close" aria-label="Fermer">${I.close}</button>
      </div>
      <div class="side-panel-body" id="mx-sp-body"><div class="side-panel-empty">Le raisonnement de l'agent s'affichera ici.</div></div>
    </div>
    <button class="mx-think-fab" id="mx-think-fab">Raisonnement</button>
    <div class="mx-menu-layer" id="mx-menu-layer" aria-hidden="true"></div>
  </div>
</div>`;

const PERMS = {
  write: { label: "Espace Write", approve: false, plan: false, icon: "edit", desc: "L'agent lit et modifie librement le workspace" },
  ask: { label: "Ask permission", approve: true, plan: false, icon: "ask", desc: "Demande ta validation avant chaque modification" },
  read: { label: "Read only", approve: false, plan: true, icon: "eye", desc: "L'agent peut lire les fichiers, jamais les modifier" },
};

const LS = {
  get(k, d) { try { const v = localStorage.getItem("mx." + k); return v == null ? d : v; } catch (e) { return d; } },
  set(k, v) { try { localStorage.setItem("mx." + k, v); } catch (e) {} },
  del(k) { try { localStorage.removeItem("mx." + k); } catch (e) {} },
  getJSON(k, d) { try { const v = localStorage.getItem("mx." + k); return v == null ? d : JSON.parse(v); } catch (e) { return d; } },
  setJSON(k, v) { try { localStorage.setItem("mx." + k, JSON.stringify(v)); } catch (e) {} },
};

// ============================================================================
// Corps de la vue (2/3) : dropdowns, modèle, permissions, projets, discussions
// ============================================================================

function cdropItemHTML(t, d, selected, icon, attrs) {
  return '<button class="cdrop-item' + (selected ? " selected" : "") + '"' + (attrs ? " " + attrs : "") + ">" +
    '<span class="cdrop-item-left">' + (icon || "") +
    '<span class="cdrop-item-text"><span class="t">' + esc(t) + "</span>" +
    (d ? '<span class="d">' + esc(d) + "</span>" : "") + "</span></span>" + I.check + "</button>";
}

// ---------------- panneau Todos (gestion des tâches de l'agent) ----------------
// Fonctions pures au niveau module (testables) ; le panneau vit dans #marex-view.
export function renderMxTodos(todos) {
  const panel = document.querySelector("#mx-todos");
  if (!panel) return;
  if (!Array.isArray(todos) || !todos.length) {
    panel.style.display = "none";
    panel.innerHTML = "";
    return;
  }
  const done = todos.filter((t) => t.status === "completed").length;
  const inprog = todos.filter((t) => t.status === "in_progress").length;
  const pending = todos.length - done - inprog;
  const marks = { completed: "✓", in_progress: "◐", pending: "○" };
  // En-tête façon Harness : "≣ To-dos · 1 in progress · 1 pending".
  let html = '<div class="mx-todos-head"><span class="mx-todos-title">≣ To-dos</span>' +
    '<span class="mx-todos-progress">' + done + "/" + todos.length +
    " · " + inprog + " in progress · " + pending + " pending</span></div>" +
    '<ul class="mx-todos-list">';
  for (const t of todos) {
    const st = t.status || "pending";
    const content = String(t.content || "").replace(/</g, "&lt;");
    html += '<li class="mx-todo-item todo-' + st + '"><span class="mx-todo-mark">' +
      (marks[st] || "○") + "</span><span>" + content + "</span></li>";
  }
  panel.innerHTML = html + "</ul>";
  panel.style.display = "";
}
export function clearMxTodos() {
  const panel = document.querySelector("#mx-todos");
  if (panel) { panel.style.display = "none"; panel.innerHTML = ""; }
}

// Placement d'un menu du composer dans le calque #mx-menu-layer.
// Mathématiques pures (rects + dimensions) : testable sans DOM.
// - par défaut le menu s'ouvre VERS LE HAUT (boutons du pied de composer) ;
// - `down:true` : vers le bas (sélecteur projet, en haut du composer) ;
// - `right:true` : aligné à droite du bouton (menu modèle) ;
// - repli : si pas la place en haut, ouvre en bas ; clampé à la vue.
export function positionMenu(viewRect, btnRect, menuW, menuH, opts = {}) {
  const gap = 8;
  const vw = viewRect.width;
  const vh = viewRect.height;
  let left = btnRect.left - (viewRect.left || 0);
  if (opts.right) left = btnRect.right - (viewRect.left || 0) - menuW;
  left = Math.max(8, Math.min(left, vw - menuW - 8));
  let top;
  if (opts.down) {
    top = btnRect.bottom - (viewRect.top || 0) + gap;
  } else {
    top = btnRect.top - (viewRect.top || 0) - menuH - gap;
    if (top < 8) top = btnRect.bottom - (viewRect.top || 0) + gap;
  }
  if (top + menuH > vh - 8) top = Math.max(8, vh - menuH - 8);
  return { left: Math.round(left), top: Math.round(top) };
}

// Signature de la liste des discussions : permet de ne reconstruire le DOM
// que lorsque quelque chose a réellement changé (statut, titre, favori,
// sélection, filtre). Fonction pure, testable sans DOM.
export function discSignature(list, titles, favorites, currentId, filter) {
  const lbl = (a) => {
    const t = (titles && titles[a.id]) || (a && a.preview) || "";
    const s = String(t).replace(/\s+/g, " ").trim();
    return s.length > 42 ? s.slice(0, 42) + "…" : s;
  };
  return JSON.stringify({
    f: filter || "",
    items: (list || []).map((a) => [
      a.id,
      a.status || "",
      favorites && favorites.has(a.id) ? 1 : 0,
      a.id === currentId ? 1 : 0,
      lbl(a),
    ]),
  });
}

export function initAgents() {
  const toolbarBtn =
    document.getElementById("agents-btn") ||
    document.querySelector('.dev-module-btn[data-module="agents"]');
  if (!toolbarBtn) return;
  if (document.getElementById("marex-view")) return;

  const view = document.createElement("div");
  view.id = "marex-view";
  view.setAttribute("data-accent", "sky");
  view.innerHTML = VIEW_HTML;
  document.body.appendChild(view);
  const $ = (s) => view.querySelector(s);

  const hero = $("#mx-hero"),
    chatPanel = $("#mx-chat-panel"),
    chatLog = $("#mx-chat-log"),
    input = $("#mx-input"),
    sendBtn = $("#mx-send"),
    stopBtn = $("#mx-stop"),
    tokenCounter = $("#mx-token-counter"),
    statusWrap = $("#mx-status"),
    statusText = $("#mx-status-text"),
    sidePanel = $("#mx-side-panel"),
    spBody = $("#mx-sp-body"),
    spSpinner = $("#mx-sp-spinner"),
    thinkFab = $("#mx-think-fab"),
    sidebar = $("#mx-sidebar"),
    discList = $("#mx-disc-list"),
    discEmpty = $("#mx-disc-empty"),
    searchInput = $("#mx-search"),
    favBtn = $("#mx-fav");

  // ---------------- état ----------------
  let families = [];
  let selFamily = LS.get("family", null);
  let selMode = LS.get("mode", null);
  let perm = LS.get("perm", "write");
  if (!PERMS[perm]) perm = "write";
  // Mode Plan/Build : "plan" <=> perm "read" (lecture seule), "build" <=>
  // derniere permission d'ecriture ("write" ou "ask"). Tab bascule l'un
  // vers l'autre ; le payload agent envoie plan: PERMS[perm].plan.
  let lastBuildPerm = perm === "read" ? "write" : perm;
  const mxMode = () => (perm === "read" ? "plan" : "build");
  function setPerm(p) {
    if (!PERMS[p]) return;
    if (p !== "read") lastBuildPerm = p;
    perm = p;
    LS.set("perm", perm);
    $("#mx-label-perm").textContent = PERMS[perm].label;
    buildPermMenu();
    applyMxModeVisual();
  }
  function setMxMode(mode) {
    setPerm(mode === "plan" ? "read" : lastBuildPerm || "write");
  }
  function applyMxModeVisual() {
    const mode = mxMode();
    const composer = $("#mx-composer");
    if (composer) {
      composer.classList.toggle("mode-plan", mode === "plan");
      composer.classList.toggle("mode-build", mode === "build");
    }
    const badge = $("#mx-mode-badge");
    if (badge) {
      badge.textContent = mode === "plan" ? "PLAN" : "BUILD";
      badge.classList.toggle("plan", mode === "plan");
      badge.classList.toggle("build", mode === "build");
      badge.title = mode === "plan"
        ? "Mode Plan — lecture seule (Tab pour passer en Build)"
        : "Mode Build — l'agent peut modifier (Tab pour passer en Plan)";
    }
  }
  let repo = LS.get("repo", "");
  // Worktree DÉSACTIVÉ par défaut : l'option n'a de sens qu'avec un dépôt
  // renseigné, sinon chaque run affichait "Worktree indisponible…".
  let useWorktree = LS.get("worktree", "0") === "1";
  // Globe (recherche web) et effort de réflexion (thinking obligatoire :
  // l'agent réfléchit toujours, pas d'interrupteur).
  let mxWeb = LS.get("mx_web", "0") === "1";
  let mxEffort = LS.get("mx_effort", "default");
  // Profondeur de recherche web : "standard" | "deep".
  let mxWebDepth = LS.get("mx_webdepth", "standard") === "deep" ? "deep" : "standard";
  // Pièces jointes du tour en cours (ids renvoyés par /api/chat/attach).
  let mxAttachments = [];
  if (!["default", "low", "medium", "high"].includes(mxEffort)) mxEffort = "default";
  const EFFORT_LABELS = { default: "Défaut", low: "Faible", medium: "Moyen", high: "Max" };
  let favorites = new Set(LS.getJSON("favs", []));
  let missionTitles = LS.getJSON("titles", {});
  let agents = [];
  let currentId = null;
  let thread = null;
  let opened = false;
  let metricsTimer = null;
  let discTimer = null;
  let prevNet = null;
  // Dernière signature rendue de la liste des discussions : évite de
  // reconstruire le DOM à chaque polling quand rien n'a changé.
  let lastDiscSig = "";
  // Brouillon du composer, persisté par discussion ("new" si aucune).
  let draftTimer = null;
  const draftKey = () => "draft:" + (currentId || "new");
  function saveDraftNow() { LS.set(draftKey(), input.value); }
  function loadDraft() {
    input.value = LS.get(draftKey(), "");
    autosize();
    updateCounter();
  }
  let prevNetTs = 0;
  let spUserClosed = false;
  const undoStack = [];
  const redoStack = [];
  let lastPushed = "";

  // ---------------- dropdowns (portalés) ----------------
  // Les menus sont déplacés dans #mx-menu-layer (enfant direct de
  // #marex-view) : .mx-composer a overflow:hidden (coins arrondis) qui
  // rognait les menus, et le positionnement vers le haut n'existait
  // que sous 600px. Le calque n'est rogné par aucun ancêtre.
  function closeAllDrops() {
    view.querySelectorAll(".cdrop-menu.open").forEach((m) => m.classList.remove("open"));
    const um = $("#mx-user-menu");
    if (um) um.classList.remove("open");
  }
  document.addEventListener("click", (e) => {
    if (opened && !e.target.closest(".cdrop") && !e.target.closest(".mx-menu-layer")) closeAllDrops();
  });
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") closeAllDrops();
  });
  // Un menu ouvert se ferme si la vue défile ou est redimensionnée
  // (le positionnement absolu ne suivrait plus le bouton).
  view.addEventListener(
    "scroll",
    (e) => {
      if (e.target.closest && e.target.closest(".cdrop-menu")) return; // scroll interne au menu : garder ouvert
      closeAllDrops();
    },
    true
  );
  window.addEventListener("resize", closeAllDrops);
  function placeMenu(btn, menu) {
    const layer = $("#mx-menu-layer");
    if (menu.parentElement !== layer) layer.appendChild(menu);
    menu.classList.add("open");
    const viewRect = view.getBoundingClientRect();
    const r = btn.getBoundingClientRect();
    const p = positionMenu(viewRect, r, menu.offsetWidth, menu.offsetHeight, {
      down: menu.id === "mx-menu-project", // sélecteur projet (haut du composer) : ouvre vers le bas
      right: menu.classList.contains("up-right"), // menu modèle : aligné à droite
    });
    menu.style.left = p.left + "px";
    menu.style.top = p.top + "px";
    menu.style.right = "auto";
    menu.style.bottom = "auto";
  }
  function wireDrop(btnSel, menuSel) {
    const btn = $(btnSel),
      menu = $(menuSel);
    btn.addEventListener("click", (e) => {
      e.stopPropagation();
      const was = menu.classList.contains("open");
      closeAllDrops();
      if (!was) placeMenu(btn, menu);
    });
    menu.addEventListener("click", (e) => e.stopPropagation());
  }
  wireDrop("#mx-btn-model", "#mx-menu-model");
  wireDrop("#mx-btn-perm", "#mx-menu-perm");
  wireDrop("#mx-btn-project", "#mx-menu-project");
  wireDrop("#mx-plus", "#mx-menu-plus");
  wireDrop("#mx-btn-effort", "#mx-menu-effort");

  $("#mx-web").addEventListener("click", () => {
    mxWeb = !mxWeb;
    LS.set("mx_web", mxWeb ? "1" : "0");
    refreshMxWeb();
  });

  // ---------------- pièces jointes : input + glisser-déposer ----------------
  $("#mx-file-input").addEventListener("change", (e) => {
    if (e.target.files && e.target.files.length) mxUploadFiles(e.target.files);
    e.target.value = "";
  });
  const composerEl = $("#mx-composer");
  composerEl.addEventListener("dragover", (e) => e.preventDefault());
  composerEl.addEventListener("drop", (e) => {
    e.preventDefault();
    if (e.dataTransfer && e.dataTransfer.files && e.dataTransfer.files.length) {
      mxUploadFiles(e.dataTransfer.files);
    }
  });


  // Libellé du modèle : uniquement le mode (Flash, Standard, Elite…),
  // sans le préfixe de la famille.
  function modeLabel(family, mode) {
    const f = families.find((x) => x.id === family);
    const m = f && (f.modes || []).find((x) => x.mode === mode);
    return m ? m.label || mode : mode;
  }

  // Description du niveau à partir du pool EFFECTIF (overrides du
  // Sélecteur agent appliqués) : le choix du modèle se fait uniquement
  // dans Configuration → Sélecteur agent ; le composer n'affiche que le
  // modèle réellement configuré pour chaque niveau.
  function modeModelDesc(m) {
    const pool = (m && m.pool) || [];
    if (pool.length === 1) {
      const p = pool[0];
      return (p.label || p.model || "") + " · " + (p.provider || "");
    }
    if (pool.length > 1) return "Fallback · " + pool.length + " modèles";
    return (m && m.rule) || "";
  }

  // ---------------- menu niveau ----------------
  // Le composer choisit le NIVEAU (Flash / Standard / Elite). Le MODÈLE
  // de chaque niveau se choisit exclusivement dans Configuration →
  // Sélecteur agent : la description affiche le pool effectif.
  function buildModelMenu() {
    const menu = $("#mx-menu-model");
    menu.innerHTML = "";
    // Vue Agents : uniquement les modes agent (Flash / Standard / Elite),
    // liste plate sans préfixe ni en-tête de famille. Les identifiants
    // internes famille + mode restent intacts pour l'API.
    let html = "";
    for (const f of families) {
      for (const m of (f.modes || []).filter((m) => m.agent)) {
        html += cdropItemHTML(m.label || m.mode, modeModelDesc(m), f.id === selFamily && m.mode === selMode, "",
          'data-f="' + esc(f.id) + '" data-m="' + esc(m.mode) + '"');
      }
    }
    menu.innerHTML = html || '<div class="sb-tree-empty">Aucun modèle agent disponible.</div>';
    menu.querySelectorAll(".cdrop-item[data-f]").forEach((it) => {
      it.addEventListener("click", () => {
        selFamily = it.dataset.f;
        selMode = it.dataset.m;
        LS.set("family", selFamily);
        LS.set("mode", selMode);
        $("#mx-label-model").textContent = modeLabel(selFamily, selMode);
        closeAllDrops();
        buildModelMenu();
        updateCounter();
        checkVisionWarning();
      });
    });
  }

  async function loadFamilies() {
    try {
      const data = await api("/api/aliases");
      families = (data.families || []).filter((f) => (f.modes || []).some((m) => m.agent));
    } catch (e) {
      families = [];
    }
    const ok = families.some(
      (f) => f.id === selFamily && (f.modes || []).some((m) => m.mode === selMode && m.agent)
    );
    if (!ok) {
      const f0 = families[0];
      const m0 = f0 && (f0.modes || []).find((m) => m.agent);
      selFamily = f0 ? f0.id : null;
      selMode = m0 ? m0.mode : null;
      LS.set("family", selFamily || "");
      LS.set("mode", selMode || "");
    }
    $("#mx-label-model").textContent = selFamily ? modeLabel(selFamily, selMode) : "Choisir un niveau";
    buildModelMenu();
    checkVisionWarning();
  }

  // ---------------- menu permissions ----------------
  function buildPermMenu() {
    const menu = $("#mx-menu-perm");
    let html = '<div class="cdrop-section-label">Permissions</div>';
    for (const key of ["read", "write", "ask"]) {
      const p = PERMS[key];
      html += cdropItemHTML(p.label, p.desc, perm === key, I[p.icon], 'data-perm="' + key + '"');
    }
    menu.innerHTML = html;
    menu.querySelectorAll(".cdrop-item[data-perm]").forEach((it) => {
      it.addEventListener("click", () => {
        setPerm(it.dataset.perm);
        closeAllDrops();
      });
    });
  }

  // ---------------- menu + ----------------
  // Ouvre la modale Configuration sur un onglet précis (Remote,
  // Compétences, Fonctionnalités…).
  function openConfigTab(tab) {
    closeAllDrops();
    window.dispatchEvent(new CustomEvent("cetas:open-config-tab", { detail: { tab } }));
  }

  // Profondeur de recherche web : standard | deep.
  function setWebDepth(v) {
    mxWebDepth = v === "deep" ? "deep" : "standard";
    LS.set("mx_webdepth", mxWebDepth);
    // Choisir une profondeur implique la recherche web : on allume le globe.
    if (!mxWeb) {
      mxWeb = true;
      LS.set("mx_web", "1");
      refreshMxWeb();
    }
    buildPlusMenu();
  }

  function buildPlusMenu() {
    const menu = $("#mx-menu-plus");
    const skillCount = mxSkills.filter((s) => s.enabled).length;
    const skillDesc = mxSkills.length
      ? skillCount + " active" + (skillCount > 1 ? "s" : "") + " / " + mxSkills.length
      : "Aucune compétence définie";
    let html = '<div class="cdrop-section-label">Fichiers</div>';
    html +=
      '<button class="cdrop-upload-btn" data-action="files">' +
      I.upload +
      " Ajouter des fichiers…</button>" +
      '<div class="cdrop-item-desc">Texte, code, PDF, images — lus par l’agent</div>';
    html += '<div class="cdrop-section-label">Projet</div>';
    html +=
      '<button class="cdrop-upload-btn" data-action="project">' +
      I.folder +
      " Ajouter un projet…</button>" +
      '<div class="cdrop-item-desc">Rechercher ou créer un projet</div>';
    html += '<div class="cdrop-section-label">Compétences</div>';
    html +=
      '<button class="cdrop-upload-btn" data-action="skills">' +
      I.star +
      " Compétences…</button>" +
      '<div class="cdrop-item-desc">' +
      esc(skillDesc) +
      "</div>";
    html += '<div class="cdrop-section-label">Recherche web</div>';
    html += cdropItemHTML(
      "Standard",
      "Réponse rapide, quelques sources",
      mxWebDepth === "standard",
      I.globe,
      'data-action="webdepth" data-v="standard"'
    );
    html += cdropItemHTML(
      "Approfondie",
      "Requêtes multiples, pages lues en entier, sources recoupées",
      mxWebDepth === "deep",
      I.globe,
      'data-action="webdepth" data-v="deep"'
    );
    html += '<div class="cdrop-divider"></div>';
    html += '<div class="cdrop-section-label">Plugins</div>';
    html +=
      '<button class="cdrop-upload-btn" data-action="plugins">' +
      I.plus +
      " Ajouter des plugins…</button>" +
      '<div class="cdrop-item-desc">Dossier plugins/ du serveur</div>';
    html +=
      '<div class="cdrop-row"><span class="cdrop-row-label">' +
      I.folder +
      "<span>Worktree git isolé</span></span>" +
      '<label class="cdrop-toggle"><input type="checkbox" id="mx-wt-toggle"' +
      (useWorktree ? " checked" : "") +
      '><span class="cdrop-toggle-slider"></span></label></div>';
    menu.innerHTML = html;
    menu.querySelector("#mx-wt-toggle").addEventListener("change", (e) => {
      useWorktree = e.target.checked;
      LS.set("worktree", useWorktree ? "1" : "0");
    });
    menu.querySelectorAll("[data-action]").forEach((it) => {
      it.addEventListener("click", () => {
        const a = it.dataset.action;
        if (a === "files") {
          closeAllDrops();
          $("#mx-file-input").click();
        } else if (a === "project") {
          closeAllDrops();
          openNewProjectModal();
        } else if (a === "skills") {
          openConfigTab("competences");
        } else if (a === "plugins") {
          // Plugins externes : gérés dans l'onglet API et Modèles
          // (rechargement). Fonctionnalités reste inaccessible depuis Agents.
          openConfigTab("apimodeles");
        } else if (a === "webdepth") {
          setWebDepth(it.dataset.v);
          closeAllDrops();
        }
      });
    });
  }

  // ---------------- pièces jointes ----------------
  function renderMxAttachments() {
    const row = $("#mx-attach-row");
    if (!mxAttachments.length) {
      row.style.display = "none";
      row.innerHTML = "";
      return;
    }
    row.style.display = "flex";
    row.innerHTML = "";
    for (const a of mxAttachments) {
      const chip = document.createElement("span");
      chip.className = "mx-attach-chip";
      chip.title = a.name;
      const icon = a.kind === "image" ? "🖼️" : "📄";
      chip.innerHTML =
        '<span class="ic">' + icon + '</span><span class="n">' + esc(a.name) + "</span>";
      const x = document.createElement("button");
      x.type = "button";
      x.textContent = "✕";
      x.title = "Retirer";
      x.addEventListener("click", () => removeMxAttachment(a));
      chip.appendChild(x);
      row.appendChild(chip);
    }
  }

  function removeMxAttachment(entry) {
    mxAttachments = mxAttachments.filter((a) => a !== entry);
    renderMxAttachments();
    checkVisionWarning();
    api("/api/chat/attach/" + encodeURIComponent(entry.id), { method: "DELETE" }).catch(() => {});
  }

  // Vide le composer après envoi. Les fichiers restent côté serveur : le
  // tour de l'agent les lit de façon asynchrone et la régénération peut
  // les relire plus tard. Le nettoyage des orphelins est fait par TTL
  // côté serveur (attach.Store.CleanOlderThan). Ne PAS supprimer ici :
  // le DELETE partirait en course avec le chargement par l'agent.
  function clearAttachments() {
    mxAttachments = [];
    renderMxAttachments();
    checkVisionWarning();
  }

  const IMAGE_EXTS = [".png", ".jpg", ".jpeg", ".gif", ".webp"];
  function isImageName(name) {
    const n = String(name || "").toLowerCase();
    return IMAGE_EXTS.some((e) => n.endsWith(e));
  }

  async function mxUploadFiles(files) {
    for (const f of files) {
      try {
        const fd = new FormData();
        fd.append("file", f);
        const resp = await fetch("/api/chat/attach", {
          method: "POST",
          headers: { Authorization: "Bearer " + getToken() },
          body: fd,
        });
        if (!resp.ok) throw new Error("échec du téléversement (" + resp.status + ")");
        const data = await resp.json();
        mxAttachments.push({
          id: data.id,
          name: data.name || f.name,
          kind: isImageName(data.name || f.name) ? "image" : "file",
        });
      } catch (e) {
        tokenCounter.textContent = "Pièce jointe : " + e.message;
      }
    }
    renderMxAttachments();
    checkVisionWarning();
  }

  // ---------------- alerte vision ----------------
  // Si une image est jointe et que le modèle sélectionné ne sait pas la
  // lire, on l'écrit explicitement et on propose les modèles compatibles.
  let mxCaps = null;
  async function loadMxCaps() {
    if (mxCaps) return mxCaps;
    try {
      const d = await api("/api/capabilities");
      mxCaps = (d && d.caps) || {};
    } catch (_) {
      mxCaps = {};
    }
    return mxCaps;
  }
  function modePool(familyId, modeId) {
    const f = families.find((x) => x.id === familyId);
    const m = f && (f.modes || []).find((x) => x.mode === modeId);
    return (m && m.pool) || [];
  }
  function poolHasVision(pool, caps) {
    return pool.some((mb) => caps[mb.provider + "/" + mb.model] && caps[mb.provider + "/" + mb.model].vision);
  }
  function visionModes() {
    const out = [];
    for (const f of families) {
      for (const m of (f.modes || []).filter((x) => x.agent)) {
        if (poolHasVision(m.pool || [], mxCaps || {})) {
          out.push({ family: f.id, mode: m.mode, label: m.label || m.mode });
        }
      }
    }
    return out;
  }
  async function checkVisionWarning() {
    const box = $("#mx-vision-warn");
    const hasImage = mxAttachments.some((a) => a.kind === "image");
    if (!hasImage || !selFamily || !selMode) {
      box.style.display = "none";
      box.innerHTML = "";
      return;
    }
    const caps = await loadMxCaps();
    if (poolHasVision(modePool(selFamily, selMode), caps)) {
      box.style.display = "none";
      box.innerHTML = "";
      return;
    }
    const compat = visionModes().slice(0, 4);
    box.style.display = "block";
    box.innerHTML =
      "⚠️ <b>" +
      esc(modeLabel(selFamily, selMode)) +
      "</b> ne sait pas lire les images. " +
      (compat.length
        ? "Modèles compatibles : " +
          compat
            .map(
              (c) =>
                '<button type="button" data-f="' +
                esc(c.family) +
                '" data-m="' +
                esc(c.mode) +
                '">' +
                esc(c.label) +
                "</button>"
            )
            .join("")
        : "Déclarez un modèle vision dans Configuration → Capacités des modèles.");
    box.querySelectorAll("button[data-f]").forEach((b) => {
      b.addEventListener("click", () => {
        selFamily = b.dataset.f;
        selMode = b.dataset.m;
        LS.set("family", selFamily);
        LS.set("mode", selMode);
        $("#mx-label-model").textContent = modeLabel(selFamily, selMode);
        buildModelMenu();
        updateCounter();
        checkVisionWarning();
      });
    });
  }

  // ---------------- compétences (comptage pour le menu +) ----------------
  let mxSkills = [];
  async function loadMxSkills() {
    try {
      const d = await api("/api/skills");
      mxSkills = (d && d.skills) || [];
    } catch (_) {
      mxSkills = [];
    }
    buildPlusMenu();
  }
  // La Configuration notifie après chaque sauvegarde (PUT /api/skills).
  window.addEventListener("cetas:skills-changed", loadMxSkills);

  // ---------------- globe (web) + effort de réflexion ----------------
  function refreshMxWeb() {
    const b = $("#mx-web");
    b.classList.toggle("active", mxWeb);
    b.title = mxWeb ? "Recherche web : activée" : "Recherche web : désactivée";
    b.setAttribute("aria-pressed", mxWeb ? "true" : "false");
  }
  function buildEffortMenu() {
    const menu = $("#mx-menu-effort");
    const defs = [
      ["default", "Défaut", "Équilibré automatiquement"],
      ["low", "Faible", "Rapide"],
      ["medium", "Moyen", "Équilibré"],
      ["high", "Max", "Qualité maximale"],
    ];
    menu.innerHTML = defs
      .map(([v, label, desc]) => cdropItemHTML(label, desc, mxEffort === v, "", 'data-effort="' + v + '"'))
      .join("");
    menu.querySelectorAll(".cdrop-item[data-effort]").forEach((it) => {
      it.addEventListener("click", () => {
        mxEffort = it.dataset.effort;
        LS.set("mx_effort", mxEffort);
        $("#mx-label-effort").textContent = EFFORT_LABELS[mxEffort] || "Défaut";
        buildEffortMenu();
        closeAllDrops();
      });
    });
  }

  // ---------------- projets (via /api/projects) ----------------
  function activeProjectName() {
    const p = Projects.active;
    return p ? p.name : "Espace partagé";
  }
  function refreshProjectLabels() {
    $("#mx-label-project").textContent = activeProjectName();
    $("#mx-active-name").textContent = activeProjectName();
    const p = Projects.active;
    $("#mx-active-icon").innerHTML = p && p.mode === "sftp" ? I.upload : I.folder;
    $("#mx-active-badge").hidden = !p;
    const sub = $("#mx-active-sub");
    if (p) {
      sub.hidden = false;
      sub.textContent = p.mode === "sftp" ? p.user + "@" + p.host + ":" + p.remote_path : "Projet local";
    } else {
      sub.hidden = true;
      sub.textContent = "";
    }
    $("#mx-active-tree").style.display = p ? "" : "none";
    renderProjectBar($("#mx-project-bar"));
  }
  function renderProjects() {
    renderProjectsList($("#mx-projects-list"), $("#mx-projects-empty"));
    $("#mx-projects-label").style.display = Projects.list.length ? "" : "none";
    renderActiveTree($("#mx-active-tree"));
  }
  function buildProjectMenu(filter) {
    const menu = $("#mx-menu-project");
    const q = (filter || "").toLowerCase();
    const match = (name) => !q || String(name || "").toLowerCase().includes(q);
    let html = '<input class="cdrop-search" id="mx-project-filter" type="text" placeholder="Rechercher un projet…" autocomplete="off">';
    html += '<div class="cdrop-section-label">Projet</div>';
    if (match("Espace partagé")) {
      html += cdropItemHTML("Espace partagé", "Workspace partagé, sans projet", !Projects.activeId, I.folder, 'data-pid="__shared__"');
    }
    for (const p of Projects.list) {
      if (p.mode === "sftp" || !match(p.name)) continue;
      html += cdropItemHTML(
        p.name,
        "Projet local",
        p.id === Projects.activeId,
        I.folder,
        'data-pid="' + esc(p.id) + '"'
      );
    }
    // Volet Remote : dossiers distants via SFTP.
    html += '<div class="cdrop-section-label">Remote</div>';
    for (const p of Projects.list) {
      if (p.mode !== "sftp" || !match(p.name)) continue;
      html += cdropItemHTML(
        p.name,
        p.user + "@" + p.host + ":" + p.remote_path,
        p.id === Projects.activeId,
        I.upload,
        'data-pid="' + esc(p.id) + '"'
      );
    }
    html += '<button class="cdrop-upload-btn" data-action="sftp">' + I.upload + " Serveur distant (SFTP)…</button>";
    html += '<div class="cdrop-divider"></div>';
    html += '<button class="cdrop-upload-btn" id="mx-pick-repo">' + I.upload + " Nouveau projet…</button>";
    menu.innerHTML = html;
    const filterInput = menu.querySelector("#mx-project-filter");
    filterInput.value = filter || "";
    filterInput.addEventListener("input", () => {
      const v = filterInput.value;
      buildProjectMenu(v);
      const inp = menu.querySelector("#mx-project-filter");
      inp.focus();
      inp.setSelectionRange(v.length, v.length);
    });
    // Le clic dans le champ ne doit pas fermer le menu.
    filterInput.addEventListener("click", (e) => e.stopPropagation());
    menu.querySelectorAll(".cdrop-item[data-pid]").forEach((it) => {
      it.addEventListener("click", async () => {
        try {
          await Projects.setActive(it.dataset.pid === "__shared__" ? "" : it.dataset.pid);
        } catch (e) {
          alert("Sélection impossible : " + (e.message || e));
        }
        closeAllDrops();
      });
    });
    menu.querySelector('[data-action="sftp"]').addEventListener("click", () => {
      closeAllDrops();
      openNewProjectModal(null, { tab: "sftp" });
    });
    menu.querySelector("#mx-pick-repo").addEventListener("click", () => {
      closeAllDrops();
      openNewProjectModal();
    });
  }
  // Rafraîchit tout le volet projet quand l'état change.
  Projects.onChange = () => {
    refreshProjectLabels();
    renderProjects();
    buildProjectMenu();
  };

// ============================================================================
// Corps de la vue (3/3) : discussions, métriques, raisonnement, thread, envoi
// ============================================================================

function discLabel(a) {
  const t = missionTitles[a.id] || a.preview;
  if (t) {
    const s = String(t).replace(/\s+/g, " ").trim();
    return s.length > 42 ? s.slice(0, 42) + "…" : s;
  }
  return "#" + String(a.id).slice(0, 8);
}

function statusTextFr(status) {
  return { running: "En cours", done: "Terminé", stopped: "Arrêté", error: "Erreur" }[status] || status;
}

async function refreshAgents() {
  try {
    const data = await api("/api/agents");
    agents = data.agents || [];
  } catch (e) {
    return;
  }
  updateActivityDot();
  const filter = searchInput.value.trim().toLowerCase();
  const sig = discSignature(agents, missionTitles, favorites, currentId, filter);
  if (sig === lastDiscSig) return; // aucun changement visible : pas de reflow
  renderDiscussions(filter);
  updateStatusBadge();
}

// Pastille pulsante sur le bouton module quand au moins un agent travaille,
// visible même quand la vue Agents est fermée.
let dotTimer = null;
function updateActivityDot() {
  const running = agents.filter((a) => a.status === "running").length;
  let dot = toolbarBtn.querySelector(".mod-activity-dot");
  if (!dot) {
    dot = document.createElement("span");
    dot.className = "mod-activity-dot";
    dot.setAttribute("aria-hidden", "true");
    toolbarBtn.appendChild(dot);
  }
  dot.hidden = running === 0;
  toolbarBtn.classList.toggle("has-running", running > 0);
  // Vue fermée : on continue de surveiller en arrière-plan à cadence réduite
  // (uniquement tant qu'un agent travaille) pour éteindre la pastille à temps.
  // Vue ouverte : le polling des discussions (5 s) prend le relais.
  if (!opened) {
    if (running > 0 && !dotTimer) {
      dotTimer = setInterval(async () => {
        try {
          const data = await api("/api/agents");
          agents = data.agents || [];
        } catch (e) {
          return;
        }
        updateActivityDot();
      }, 15000);
    } else if (running === 0 && dotTimer) {
      clearInterval(dotTimer);
      dotTimer = null;
    }
  } else if (dotTimer) {
    clearInterval(dotTimer);
    dotTimer = null;
  }
}

function renderDiscussions(filter) {
  lastDiscSig = discSignature(agents, missionTitles, favorites, currentId, filter);
  discList.innerHTML = "";
  const items = agents.filter((a) => !filter || discLabel(a).toLowerCase().includes(filter));
  discEmpty.style.display = items.length ? "none" : "";
  // favoris d'abord, puis plus récents
  items.sort((x, y) => {
    const fx = favorites.has(x.id) ? 0 : 1,
      fy = favorites.has(y.id) ? 0 : 1;
    if (fx !== fy) return fx - fy;
    return (y.created || 0) - (x.created || 0);
  });
  for (const a of items) {
    const b = el("button", "sb-hist-item" + (a.id === currentId ? " active" : ""));
    b.type = "button";
    b.title = discLabel(a);
    const lab = el("span", "sb-hist-label", (favorites.has(a.id) ? "★ " : "") + discLabel(a));
    const del = el("button", "mx-hist-del", "×");
    del.type = "button";
    del.title = "Supprimer";
    del.setAttribute("aria-label", "Supprimer la discussion");
    del.addEventListener("click", (e) => {
      e.stopPropagation();
      deleteAgent(a.id);
    });
    b.appendChild(lab);
    b.appendChild(del);
    b.addEventListener("click", () => openDiscussion(a.id));
    discList.appendChild(b);
  }
}

async function deleteAgent(id) {
  if (!confirm("Supprimer cet agent et son worktree ?")) return;
  try {
    await api("/api/agents/" + encodeURIComponent(id), { method: "DELETE" });
  } catch (e) {}
  delete missionTitles[id];
  LS.setJSON("titles", missionTitles);
  favorites.delete(id);
  LS.setJSON("favs", [...favorites]);
  LS.del("draft:" + id); // brouillon orphelin
  if (currentId === id) newConversation();
  else refreshAgents();
}

// ---------------- métriques ----------------
function dotClass(pct) {
  if (pct == null || isNaN(pct)) return "";
  return pct >= 90 ? "crit" : pct >= 70 ? "warn" : "ok";
}
function setMetric(kind, text, pct) {
  const v = document.getElementById("mx-m-" + kind);
  const d = document.getElementById("mx-m-" + kind + "-dot");
  if (v) v.textContent = text;
  if (d) d.className = "sb-metric-dot pulse " + dotClass(pct);
}
function fmtRate(bps) {
  if (!isFinite(bps) || bps < 0) bps = 0;
  const units = ["B/s", "KB/s", "MB/s", "GB/s"];
  let i = 0;
  while (bps >= 1024 && i < units.length - 1) {
    bps /= 1024;
    i++;
  }
  return (i === 0 ? Math.round(bps) : bps.toFixed(bps < 10 ? 1 : 0)) + " " + units[i];
}
async function pollMetrics() {
  try {
    const m = await api("/api/metrics");
    const now = Date.now();
    if (m.cpu != null) setMetric("cpu", Math.round(m.cpu) + "%", m.cpu);
    else setMetric("cpu", "—", null);
    if (m.ram && m.ram.total) {
      const pct = (m.ram.used * 100) / m.ram.total;
      setMetric(
        "ram",
        (m.ram.used / 1073741824).toFixed(1) + " / " + (m.ram.total / 1073741824).toFixed(1) + " GiB",
        pct
      );
    } else setMetric("ram", "—", null);
    if (m.disk && m.disk.total) {
      const pct = (m.disk.used * 100) / m.disk.total;
      setMetric(
        "disk",
        Math.round(m.disk.used / 1073741824) + " / " + Math.round(m.disk.total / 1073741824) + " GB",
        pct
      );
    } else setMetric("disk", "—", null);
    if (m.net && m.net.rx_bytes != null) {
      if (prevNet && now > prevNetTs) {
        const dt = (now - prevNetTs) / 1000;
        setMetric(
          "net",
          "↓ " + fmtRate((m.net.rx_bytes - prevNet.rx) / dt) + " · ↑ " + fmtRate((m.net.tx_bytes - prevNet.tx) / dt),
          null
        );
      }
      prevNet = { rx: m.net.rx_bytes, tx: m.net.tx_bytes };
      prevNetTs = now;
    } else setMetric("net", "—", null);
  } catch (e) {}
}

// ---------------- panneau raisonnement ----------------
let spScrollFrame = 0;
function spRaf(fn) {
  if (typeof requestAnimationFrame === "function") return requestAnimationFrame(fn);
  return setTimeout(fn, 16);
}
// Coalesce le scroll sur une frame : evite le reflow (scrollHeight) a
// chaque delta de raisonnement (technique Marexcode).
function spScheduleScroll() {
  if (spScrollFrame) return;
  spScrollFrame = spRaf(() => {
    spScrollFrame = 0;
    spBody.scrollTop = spBody.scrollHeight;
  });
}
function mxResetReason() {
  spUserClosed = false;
  spBody.innerHTML = '<div class="side-panel-empty">Le raisonnement de l\'agent s\'affichera ici.</div>';
  spSpinner.style.display = "none";
  thinkFab.classList.remove("show");
}
function mxOpenReason() {
  sidePanel.classList.add("open");
  thinkFab.classList.remove("show");
}
function mxCloseReason(manual) {
  if (manual) spUserClosed = true;
  sidePanel.classList.remove("open");
  const t = spBody.querySelector(".reason-text");
  thinkFab.classList.toggle("show", !!(t && t.textContent.trim()));
}
const reasonHooks = {
  append(text, replace) {
    let t = spBody.querySelector(".reason-text");
    if (!t) {
      spBody.innerHTML = "";
      t = el("div", "reason-text");
      spBody.appendChild(t);
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
    spSpinner.style.display = "";
    if (!spUserClosed) mxOpenReason();
    else thinkFab.classList.add("show");
    spScheduleScroll();
  },
  finish() {
    spSpinner.style.display = "none";
    const t = spBody.querySelector(".reason-text");
    thinkFab.classList.toggle("show", !!(t && t.textContent.trim() && !sidePanel.classList.contains("open")));
  },
  reset() {
    mxResetReason();
  },
};

// ---------------- cycle de vie du thread ----------------
function disconnectThread() {
  if (thread) {
    thread.disconnect();
    thread = null;
  }
}

function updateStatusBadge() {
  const a = agents.find((x) => x.id === currentId);
  if (!a) {
    statusWrap.style.display = "none";
    return;
  }
  statusWrap.style.display = "";
  statusWrap.className = "mx-agent-status " + a.status;
  statusText.textContent = statusTextFr(a.status) + " · " + modeLabel(a.family, a.mode);
}

function buildThread(id) {
  disconnectThread();
  const routeBadge = el("span", "route-badge");
  routeBadge.hidden = true;
  const statsBadge = el("span", "stats-badge");
  statsBadge.hidden = true;
  thread = new ThreadView({
    log: chatLog,
    stopBtn,
    routeBadge,
    statsBadge,
    emptyHTML: '<div class="agent-empty" data-empty>En attente des premiers messages de l’agent…</div>',
    streamURL: (from) => "/api/agents/" + encodeURIComponent(id) + "/stream?from=" + from,
    sendURL: "/api/agents/" + encodeURIComponent(id) + "/message",
    stopURL: "/api/agents/" + encodeURIComponent(id) + "/stop",
    approveURL: "/api/agents/" + encodeURIComponent(id) + "/approve",
    getPayload: (text) => {
      const p = {
        message: text,
        approve: PERMS[perm].approve,
        plan: PERMS[perm].plan,
        web: mxWeb,
        web_depth: mxWebDepth,
        think: true, // réflexion obligatoire pour l'agent
        effort: mxEffort,
      };
      if (mxAttachments.length) p.attachments = mxAttachments.map((a) => a.id);
      return p;
    },
    reasonHooks,
    onEvent: handleMxEvent,
    onDone: () => {
      if (statsBadge.textContent) tokenCounter.textContent = statsBadge.textContent;
      refreshAgents();
    },
  });
}

function enterChat() {
  hero.classList.add("has-chat");
  chatPanel.style.display = "flex";
}

function handleMxEvent(ev) {
  // Les événements d'historique rejouent aussi ici : le dernier TodoWrite gagne.
  if (ev && ev.tool && ev.tool.name === "TodoWrite" && ev.tool.args && Array.isArray(ev.tool.args.todos)) {
    renderMxTodos(ev.tool.args.todos);
  }
  if (ev && ev.reset) clearMxTodos();
}

function newConversation() {
  disconnectThread();
  clearMxTodos();
  clearTimeout(draftTimer);
  saveDraftNow(); // brouillon de la discussion quittée (currentId encore positionné)
  currentId = null;
  try {
    sessionStorage.removeItem("cetas.agents.current");
  } catch (e) {}
  hero.classList.remove("has-chat");
  chatPanel.style.display = "none";
  chatLog.innerHTML = "";
  mxResetReason();
  loadDraft(); // restaure le brouillon "new" (ou vide)
  updateFav();
  renderDiscussions(searchInput.value.trim().toLowerCase());
  updateStatusBadge();
}

function openDiscussion(id) {
  if (currentId === id && thread) return;
  clearMxTodos();
  clearTimeout(draftTimer);
  saveDraftNow(); // brouillon de la discussion quittée
  currentId = id;
  try {
    sessionStorage.setItem("cetas.agents.current", id);
  } catch (e) {}
  enterChat();
  chatLog.innerHTML = "";
  mxResetReason();
  loadDraft(); // restaure le brouillon de cette discussion
  buildThread(id);
  thread.connect();
  renderDiscussions(searchInput.value.trim().toLowerCase());
  updateStatusBadge();
  updateFav();
  if (window.innerWidth <= 1024) sidebar.classList.remove("open");
}

async function send() {
  const text = input.value.trim();
  if (!text) return;
  if (!selFamily || !selMode) {
    $("#mx-btn-model").style.color = "#e05252";
    setTimeout(() => ($("#mx-btn-model").style.color = ""), 1200);
    document.getElementById("mx-menu-model").classList.add("open");
    return;
  }
  const sentDraftKey = draftKey(); // clé du brouillon AVANT création éventuelle de l'agent
  clearTimeout(draftTimer);
  input.disabled = true;
  try {
    if (!currentId) {
      const res = await api("/api/agents", {
        method: "POST",
        body: {
          family: selFamily,
          mode: selMode,
          message: text,
          approve: PERMS[perm].approve,
          plan: PERMS[perm].plan,
          worktree: useWorktree,
          repo,
          project_id: Projects.activeId || undefined,
          web: mxWeb,
          web_depth: mxWebDepth,
          think: true, // réflexion obligatoire pour l'agent
          effort: mxEffort,
          attachments: mxAttachments.length ? mxAttachments.map((a) => a.id) : undefined,
        },
      });
      missionTitles[res.id] = text;
      LS.setJSON("titles", missionTitles);
      currentId = res.id;
      try {
        sessionStorage.setItem("cetas.agents.current", res.id);
      } catch (e) {}
      enterChat();
      chatLog.innerHTML = "";
      mxResetReason();
      buildThread(res.id);
      thread.connect();
      renderDiscussions(searchInput.value.trim().toLowerCase());
      updateStatusBadge();
      updateFav();
    } else if (thread) {
      const ok = await thread.sendText(text);
      if (!ok) {
        input.disabled = false;
        input.focus();
        return; // agent déjà en génération : on garde le texte
      }
    }
    input.value = "";
    LS.del(sentDraftKey); // brouillon envoyé : on le purge (clé d'avant l'envoi)
    clearAttachments(); // pièces jointes envoyées : on les purge
    pushUndo("");
    autosize();
    updateCounter();
    refreshAgents();
  } catch (e) {
    // message d'erreur visible dans le fil si un thread existe, sinon compteur
    if (thread) thread.addError(e.message);
    else tokenCounter.textContent = "Erreur : " + e.message;
  } finally {
    input.disabled = false;
    input.focus();
  }
}

// ---------------- composer : autosize, undo/redo, compteur ----------------
function autosize() {
  input.style.height = "24px";
  input.style.height = Math.min(input.scrollHeight, 200) + "px";
}
function pushUndo(v) {
  if (v === lastPushed) return;
  lastPushed = v;
  undoStack.push(v);
  if (undoStack.length > 60) undoStack.shift();
  redoStack.length = 0;
  syncUndoBtns();
}
function syncUndoBtns() {
  $("#mx-undo").disabled = undoStack.length === 0;
  $("#mx-redo").disabled = redoStack.length === 0;
}
let undoTimer = null;
function updateCounter() {
  const t = input.value.trim();
  tokenCounter.textContent = t ? "↑ ~" + estimateTokens(t) : "";
}
function updateFav() {
  const on = currentId && favorites.has(currentId);
  favBtn.classList.toggle("active", !!on);
  favBtn.title = on ? "Retirer des favoris" : "Ajouter aux favoris";
}

// ---------------- ouverture / fermeture (page complète plein écran) ----------------
// La vue Agents est TOUJOURS sombre (fond #000). La coloration hljs y est
// garantie par la feuille toujours active /css/features/hljs-agents-dark.css
// (scopée #marex-view), indépendante du thème applicatif et du JS : en thème
// clair, les blocs de code restent clairs sur sombre.

function openView() {
  if (opened) return;
  opened = true;
  // Persistance de navigation (session) : un rafraîchissement depuis la
  // vue Agents doit y revenir, pas retomber sur le chat général.
  try {
    sessionStorage.setItem("cetas.agents.open", "1");
  } catch (e) {}
  view.classList.add("open");
  toolbarBtn.classList.add("active");
  document.body.style.overflow = "hidden";
  loadFamilies();
  refreshAgents();
  pollMetrics();
  metricsTimer = setInterval(pollMetrics, 10000);
  const tickDisc = () => {
    if (opened && !document.hidden) refreshAgents();
  };
  discTimer = setInterval(tickDisc, 5000);
  loadDraft();
  setTimeout(() => input.focus(), 60);
}
function closeView() {
  if (!opened) return;
  opened = false;
  try {
    sessionStorage.removeItem("cetas.agents.open");
    sessionStorage.removeItem("cetas.agents.current");
  } catch (e) {}
  clearTimeout(draftTimer);
  saveDraftNow(); // ne jamais perdre le texte en cours
  view.classList.remove("open");
  toolbarBtn.classList.remove("active");
  document.body.style.overflow = "";
  clearInterval(metricsTimer);
  clearInterval(discTimer);
  metricsTimer = discTimer = null;
  closeAllDrops();
  sidebar.classList.remove("open");
}

// ---------------- câblage ----------------
sendBtn.addEventListener("click", send);
// Le bouton Stop n'était câblé que dans le chat général : ici il doit
// interrompre le tour de l'agent en cours.
stopBtn.addEventListener("click", () => {
  if (thread) thread.stop();
});
input.addEventListener("keydown", (e) => {
  if (e.key === "Tab") {
    e.preventDefault();
    setMxMode(mxMode() === "plan" ? "build" : "plan");
    return;
  }
  if (e.key === "Enter" && !e.shiftKey) {
    e.preventDefault();
    send();
  }
});
// Clic sur le badge : même bascule que Tab.
$("#mx-mode-badge").addEventListener("click", () => {
  setMxMode(mxMode() === "plan" ? "build" : "plan");
  input.focus();
});
input.addEventListener("input", () => {
  autosize();
  updateCounter();
  clearTimeout(undoTimer);
  undoTimer = setTimeout(() => pushUndo(input.value), 600);
  clearTimeout(draftTimer);
  draftTimer = setTimeout(saveDraftNow, 400);
});
$("#mx-undo").addEventListener("click", () => {
  if (!undoStack.length) return;
  redoStack.push(input.value);
  input.value = undoStack.pop();
  lastPushed = input.value;
  autosize();
  updateCounter();
  syncUndoBtns();
});
$("#mx-redo").addEventListener("click", () => {
  if (!redoStack.length) return;
  undoStack.push(input.value);
  input.value = redoStack.pop();
  lastPushed = input.value;
  autosize();
  updateCounter();
  syncUndoBtns();
});
input.addEventListener("drop", (e) => e.preventDefault());
input.addEventListener("dragover", (e) => e.preventDefault());

favBtn.addEventListener("click", () => {
  if (!currentId) return;
  if (favorites.has(currentId)) favorites.delete(currentId);
  else favorites.add(currentId);
  LS.setJSON("favs", [...favorites]);
  updateFav();
  renderDiscussions(searchInput.value.trim().toLowerCase());
});

$("#mx-new").addEventListener("click", newConversation);
$("#mx-mobile-new").addEventListener("click", newConversation);
$("#mx-back").addEventListener("click", closeView);
// Le bouton module de la sidebar pilote via "cetas:toggle-agents" (sidebar.js).
// Le clic direct n'est câblé que si le bouton n'est pas un bouton module
// (ex. #agents-btn historique / tests jsdom) — sinon double bascule.
if (!toolbarBtn.classList.contains("dev-module-btn")) {
  toolbarBtn.addEventListener("click", () => (opened ? closeView() : openView()));
}
window.addEventListener("cetas:open-agents", () => {
  if (!opened) openView();
});
// Ouverture d'un agent précis (panneau Sessions de la Configuration).
window.addEventListener("cetas:open-agent", (e) => {
  const id = e && e.detail && e.detail.id;
  if (!id) return;
  if (!opened) openView();
  openDiscussion(id);
});
window.addEventListener("cetas:toggle-agents", () => {
  if (opened) closeView();
  else openView();
});
// Le Sélecteur agent (Configuration) modifie les modèles : recharger les
// familles pour que le composer affiche le pool effectif de chaque niveau.
window.addEventListener("cetas:aliases-changed", () => {
  loadFamilies();
});
window.addEventListener("cetas:close-agents", closeView);

searchInput.addEventListener("input", () => {
  renderDiscussions(searchInput.value.trim().toLowerCase());
});

function wireToggle(btnSel, panelSel) {
  const btn = $(btnSel),
    panel = $(panelSel);
  btn.addEventListener("click", () => {
    panel.classList.toggle("hidden");
    btn.querySelector("svg").classList.toggle("collapsed", panel.classList.contains("hidden"));
  });
}
wireToggle("#mx-toggle-projects", "#mx-panel-projects");
wireToggle("#mx-toggle-disc", "#mx-panel-disc");

$("#mx-new-project").addEventListener("click", () => {
  openNewProjectModal();
});

$("#mx-sp-close").addEventListener("click", () => mxCloseReason(true));
thinkFab.addEventListener("click", () => {
  spUserClosed = false;
  mxOpenReason();
});

$("#mx-hamburger").addEventListener("click", () => sidebar.classList.toggle("open"));
$("#mx-sb-close").addEventListener("click", () => sidebar.classList.remove("open"));

$("#mx-user-btn").addEventListener("click", (e) => {
  e.stopPropagation();
  $("#mx-user-menu").classList.toggle("open");
});
// Bouton "?" aligné à l'avatar, à droite : ouvre le centre d'aide.
$("#mx-docs-btn").addEventListener("click", (e) => {
  e.stopPropagation();
  $("#mx-user-menu").classList.remove("open");
  openDocs();
});
// Menu utilisateur : Paramètre / Aide / En savoir plus / FAQ / Déconnexion.
$("#mx-menu-settings").addEventListener("click", () => {
  closeAllDrops();
  window.dispatchEvent(new CustomEvent("cetas:open-config"));
});
$("#mx-menu-help").addEventListener("click", () => {
  closeAllDrops();
  openDocs();
});
$("#mx-menu-about").addEventListener("click", () => {
  closeAllDrops();
  openDocs();
});
$("#mx-menu-faq").addEventListener("click", () => {
  closeAllDrops();
  window.dispatchEvent(new CustomEvent("cetas:open-config-faq"));
});
$("#mx-logout").addEventListener("click", () => {
  closeView();
  logout();
});

// ---------------- init ----------------
$("#mx-label-perm").textContent = PERMS[perm].label;
applyMxModeVisual();
$("#mx-label-effort").textContent = EFFORT_LABELS[mxEffort] || "Défaut";
buildPermMenu();
buildPlusMenu();
buildProjectMenu();
buildEffortMenu();
refreshProjectLabels();
renderProjects();
Projects.refresh(); // charge /api/projects (+ projet actif), re-rend via onChange
refreshMxWeb();
loadMxSkills(); // recharge aussi le menu + (compteur de compétences)
mxResetReason();
syncUndoBtns();
refreshAgents(); // état initial (pastille d'activité même vue fermée)
api("/api/me")
  .then((me) => {
    const name = (me && me.username) || "Utilisateur";
    $("#mx-user-name").textContent = name;
    $("#mx-user-avatar").textContent = name.charAt(0).toUpperCase();
  })
  .catch(() => {});

// Restauration après rafraîchissement : si la vue Agents était ouverte,
// la rouvrir (et la discussion en cours, si elle existe toujours).
try {
  if (sessionStorage.getItem("cetas.agents.open") === "1") {
    openView();
    const rid = sessionStorage.getItem("cetas.agents.current");
    if (rid) {
      api("/api/agents")
        .then((d) => {
          const ids = new Set((d.agents || []).map((a) => a.id));
          if (ids.has(rid)) openDiscussion(rid);
          else {
            try {
              sessionStorage.removeItem("cetas.agents.current");
            } catch (e) {}
          }
        })
        .catch(() => {});
    }
  }
} catch (e) {}

}
