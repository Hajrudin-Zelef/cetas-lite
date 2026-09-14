import { api } from "./api.js";
import { ThreadView, el } from "./thread-view.js";
import { estimateTokens } from "./turn-tokens.js";
import { logout } from "./auth.js";

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

        <div class="sb-group-label">Projets</div>
        <button class="sb-section-toggle" id="mx-toggle-active"><span>Projet actif</span>${I.chevSm}</button>
        <div class="sb-collapsible" id="mx-panel-active">
          <div class="sb-project-header">${I.folder}<span id="mx-active-name">Espace partagé</span></div>
          <div class="sb-tree-empty">Aucun fichier dans le workspace.</div>
        </div>
        <button class="sb-section-toggle" id="mx-toggle-projects"><span>Mes projets</span>${I.chevSm}</button>
        <div class="sb-collapsible" id="mx-panel-projects">
          <div class="sb-tree-empty" id="mx-projects-empty">Aucun projet importé.</div>
          <div id="mx-projects-list"></div>
          <button id="mx-new-project" class="sb-new-project-btn">+ Nouveau projet</button>
        </div>

        <div class="sb-divider"></div>
        <div class="sb-group-label">Discussions</div>
        <button class="sb-section-toggle" id="mx-toggle-disc"><span>Toutes les discussions</span>${I.chevSm}</button>
        <div class="sb-collapsible" id="mx-panel-disc">
          <div class="sb-tree-empty" id="mx-disc-empty">Aucune discussion.</div>
          <div id="mx-disc-list"></div>
        </div>
      </div>

      <div class="sb-user-wrap" id="mx-user-wrap">
        <div class="sb-user-menu" id="mx-user-menu">
          <button class="sb-user-menu-item danger" id="mx-logout">${I.logout} Se déconnecter</button>
        </div>
        <button class="sb-user-btn" id="mx-user-btn">
          <span class="sb-avatar" id="mx-user-avatar">U</span>
          <span id="mx-user-name">Utilisateur</span>
        </button>
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

        <div class="mx-composer" id="mx-composer">
          <div class="composer-selectors">
            <div class="cdrop" id="mx-dd-project">
              <button class="selector" id="mx-btn-project">${I.folder13}<span id="mx-label-project">Espace partagé</span>${I.chevron}</button>
              <div class="cdrop-menu" id="mx-menu-project"></div>
            </div>
          </div>
          <textarea id="mx-input" rows="1" placeholder="Décrivez la tâche de code à réaliser... (glissez des images ici)"></textarea>
          <div class="composer-info-bar"><span id="mx-token-counter"></span><span id="mx-queue" style="display:none"></span></div>
          <div class="composer-footer">
            <div class="composer-footer-left">
              <div class="cdrop" id="mx-dd-plus">
                <button class="plus-btn" id="mx-plus" aria-label="Ajouter" title="Ajouter">${I.plus}</button>
                <div class="cdrop-menu" id="mx-menu-plus"></div>
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
                <button class="model-select" id="mx-btn-model"><span id="mx-label-model">Choisir un modèle</span>${I.chevron}</button>
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
  let repo = LS.get("repo", "");
  let useWorktree = LS.get("worktree", "1") !== "0";
  let projects = LS.getJSON("projects", []);
  let favorites = new Set(LS.getJSON("favs", []));
  let missionTitles = LS.getJSON("titles", {});
  let agents = [];
  let currentId = null;
  let thread = null;
  let opened = false;
  let metricsTimer = null;
  let discTimer = null;
  let prevNet = null;
  let prevNetTs = 0;
  let spUserClosed = false;
  const undoStack = [];
  const redoStack = [];
  let lastPushed = "";

  // ---------------- dropdowns ----------------
  function closeAllDrops() {
    view.querySelectorAll(".cdrop-menu.open").forEach((m) => m.classList.remove("open"));
    const um = $("#mx-user-menu");
    if (um) um.classList.remove("open");
  }
  document.addEventListener("click", (e) => {
    if (opened && !e.target.closest(".cdrop")) closeAllDrops();
  });
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") closeAllDrops();
  });
  function wireDrop(btnSel, menuSel) {
    const btn = $(btnSel),
      menu = $(menuSel);
    btn.addEventListener("click", (e) => {
      e.stopPropagation();
      const was = menu.classList.contains("open");
      closeAllDrops();
      if (!was) menu.classList.add("open");
    });
    menu.addEventListener("click", (e) => e.stopPropagation());
  }
  wireDrop("#mx-btn-model", "#mx-menu-model");
  wireDrop("#mx-btn-perm", "#mx-menu-perm");
  wireDrop("#mx-btn-project", "#mx-menu-project");
  wireDrop("#mx-plus", "#mx-menu-plus");

  function modeLabel(family, mode) {
    const f = families.find((x) => x.id === family);
    const m = f && (f.modes || []).find((x) => x.mode === mode);
    return (f ? f.label || family : family) + " · " + (m ? m.label || mode : mode);
  }

  // ---------------- menu modèle ----------------
  function buildModelMenu() {
    const menu = $("#mx-menu-model");
    menu.innerHTML = "";
    let html = "";
    for (const f of families) {
      const modes = (f.modes || []).filter((m) => m.agent);
      if (!modes.length) continue;
      html += '<div class="cdrop-section-label">' + esc(f.label || f.id) + "</div>";
      for (const m of modes) {
        html += cdropItemHTML(m.label || m.mode, m.rule || "", f.id === selFamily && m.mode === selMode, "",
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
    $("#mx-label-model").textContent = selFamily ? modeLabel(selFamily, selMode) : "Choisir un modèle";
    buildModelMenu();
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
        perm = it.dataset.perm;
        LS.set("perm", perm);
        $("#mx-label-perm").textContent = PERMS[perm].label;
        closeAllDrops();
        buildPermMenu();
      });
    });
  }

  // ---------------- menu + ----------------
  function buildPlusMenu() {
    const menu = $("#mx-menu-plus");
    menu.innerHTML =
      '<div class="cdrop-row"><span class="cdrop-row-label">' +
      I.folder +
      "<span>Worktree git isolé</span></span>" +
      '<label class="cdrop-toggle"><input type="checkbox" id="mx-wt-toggle"' +
      (useWorktree ? " checked" : "") +
      '><span class="cdrop-toggle-slider"></span></label></div>';
    const t = menu.querySelector("#mx-wt-toggle");
    t.addEventListener("change", () => {
      useWorktree = t.checked;
      LS.set("worktree", useWorktree ? "1" : "0");
    });
  }

  // ---------------- projets ----------------
  function activeProjectName() {
    return repo || "Espace partagé";
  }
  function refreshProjectLabels() {
    $("#mx-label-project").textContent = activeProjectName();
    $("#mx-active-name").textContent = activeProjectName();
  }
  function renderProjects() {
    const list = $("#mx-projects-list");
    list.innerHTML = "";
    $("#mx-projects-empty").style.display = projects.length ? "none" : "";
    for (const p of projects) {
      const b = el("button", "sb-tree-item" + (p === repo ? " active" : ""));
      b.type = "button";
      b.innerHTML = I.folder + "<span>" + esc(p) + "</span>";
      b.title = p;
      b.addEventListener("click", () => {
        repo = p;
        LS.set("repo", repo);
        refreshProjectLabels();
        renderProjects();
      });
      list.appendChild(b);
    }
  }
  function buildProjectMenu() {
    const menu = $("#mx-menu-project");
    let html = '<div class="cdrop-section-label">Projet</div>';
    html += cdropItemHTML("Espace partagé", "Dépôt courant, sans isolation git", !repo, I.folder, 'data-repo="__shared__"');
    for (const p of projects) {
      html += cdropItemHTML(p, "Dépôt importé", p === repo, I.folder, 'data-repo="' + esc(p) + '"');
    }
    html += '<div class="cdrop-divider"></div>';
    html += '<button class="cdrop-upload-btn" id="mx-pick-repo">' + I.upload + " Choisir un dépôt…</button>";
    menu.innerHTML = html;
    menu.querySelectorAll(".cdrop-item[data-repo]").forEach((it) => {
      it.addEventListener("click", () => {
        repo = it.dataset.repo === "__shared__" ? "" : it.dataset.repo;
        LS.set("repo", repo);
        refreshProjectLabels();
        renderProjects();
        closeAllDrops();
        buildProjectMenu();
      });
    });
    menu.querySelector("#mx-pick-repo").addEventListener("click", () => {
      const v = prompt("Chemin du dépôt (sous ton espace, ex. ~/workspace/mon-projet) :", repo || "~/workspace/");
      if (v == null) return;
      const path = v.trim();
      if (!path) return;
      repo = path;
      if (!projects.includes(path)) {
        projects.push(path);
        LS.setJSON("projects", projects);
      }
      LS.set("repo", repo);
      refreshProjectLabels();
      renderProjects();
      closeAllDrops();
      buildProjectMenu();
    });
  }

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
  renderDiscussions(searchInput.value.trim().toLowerCase());
  updateStatusBadge();
}

function renderDiscussions(filter) {
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
    getPayload: (text) => ({ message: text, approve: PERMS[perm].approve, plan: PERMS[perm].plan }),
    reasonHooks,
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

function newConversation() {
  disconnectThread();
  currentId = null;
  hero.classList.remove("has-chat");
  chatPanel.style.display = "none";
  chatLog.innerHTML = "";
  mxResetReason();
  input.value = "";
  autosize();
  updateCounter();
  updateFav();
  renderDiscussions(searchInput.value.trim().toLowerCase());
  updateStatusBadge();
}

function openDiscussion(id) {
  if (currentId === id && thread) return;
  currentId = id;
  enterChat();
  chatLog.innerHTML = "";
  mxResetReason();
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
        },
      });
      missionTitles[res.id] = text;
      LS.setJSON("titles", missionTitles);
      currentId = res.id;
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

// ---------------- ouverture / fermeture ----------------
function openView() {
  if (opened) return;
  opened = true;
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
  setTimeout(() => input.focus(), 60);
}
function closeView() {
  if (!opened) return;
  opened = false;
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
input.addEventListener("keydown", (e) => {
  if (e.key === "Enter" && !e.shiftKey) {
    e.preventDefault();
    send();
  }
});
input.addEventListener("input", () => {
  autosize();
  updateCounter();
  clearTimeout(undoTimer);
  undoTimer = setTimeout(() => pushUndo(input.value), 600);
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
toolbarBtn.addEventListener("click", () => (opened ? closeView() : openView()));
window.addEventListener("cetas:open-agents", () => {
  if (!opened) openView();
});

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
wireToggle("#mx-toggle-active", "#mx-panel-active");
wireToggle("#mx-toggle-projects", "#mx-panel-projects");
wireToggle("#mx-toggle-disc", "#mx-panel-disc");

$("#mx-new-project").addEventListener("click", () => {
  const v = prompt("Chemin du dépôt (sous ton espace, ex. ~/workspace/mon-projet) :", "~/workspace/");
  if (v == null) return;
  const path = v.trim();
  if (!path || projects.includes(path)) return;
  projects.push(path);
  LS.setJSON("projects", projects);
  renderProjects();
  buildProjectMenu();
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
$("#mx-logout").addEventListener("click", () => {
  closeView();
  logout();
});

// ---------------- init ----------------
$("#mx-label-perm").textContent = PERMS[perm].label;
buildPermMenu();
buildPlusMenu();
buildProjectMenu();
refreshProjectLabels();
renderProjects();
mxResetReason();
syncUndoBtns();
api("/api/me")
  .then((me) => {
    const name = (me && me.username) || "Utilisateur";
    $("#mx-user-name").textContent = name;
    $("#mx-user-avatar").textContent = name.charAt(0).toUpperCase();
  })
  .catch(() => {});

}
