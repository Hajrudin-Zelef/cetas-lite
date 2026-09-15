// Centre d'aide Cetas Lite — vue plein écran façon "help center" :
// en-tête (marque + recherche), sidebar des collections, article avec fil
// d'Ariane, sommaire "Sur cette page". Contenu : docs-data.js (public).
import { CETAS_DOCS } from "./docs-data.js";

const ICONS = {
  rocket: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09z"/><path d="M12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"/><path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0"/><path d="M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"/></svg>',
  chat: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>',
  agents: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="4" y="4" width="16" height="16" rx="2"/><rect x="9" y="9" width="6" height="6"/><line x1="9" y1="1" x2="9" y2="4"/><line x1="15" y1="1" x2="15" y2="4"/><line x1="9" y1="20" x2="9" y2="23"/><line x1="15" y1="20" x2="15" y2="23"/><line x1="20" y1="9" x2="23" y2="9"/><line x1="20" y1="14" x2="23" y2="14"/><line x1="1" y1="9" x2="4" y2="9"/><line x1="1" y1="14" x2="4" y2="14"/></svg>',
  folder: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"/></svg>',
  terminal: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="4 17 10 11 4 5"/><line x1="12" y1="19" x2="20" y2="19"/></svg>',
  puzzle: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14.7 6.3a4 4 0 0 0-4.4 4.4 4 4 0 0 0-5.6 5.6 4 4 0 0 0 5.6 5.6 4 4 0 0 0 4.4-4.4 4 4 0 0 0 5.6-5.6 4 4 0 0 0-5.6-5.6z"/><path d="M14.7 6.3v3.2a1 1 0 0 0 1 1h3.2"/></svg>',
  settings: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 1 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 1 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 1 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 1 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>',
  user: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>',
  book: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>',
  search: '<svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>',
  close: '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>',
  chevR: '<svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M9 18l6-6-6-6"/></svg>',
  help: '<svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><line x1="12" y1="17" x2="12.01" y2="17"/></svg>',
};

// Index plat : ordre global, accès par id, recherche.
const FLAT = [];
const BY_ID = new Map();
for (const coll of CETAS_DOCS.collections) {
  for (const art of coll.articles) {
    const entry = { coll, art };
    FLAT.push(entry);
    BY_ID.set(art.id, entry);
  }
}
function stripTags(html) {
  return String(html).replace(/<[^>]*>/g, " ").replace(/\s+/g, " ").trim();
}
// Recherche : le titre compte triple, les intertitres double.
export function searchDocs(query) {
  const q = query.trim().toLowerCase();
  if (q.length < 2) return [];
  const out = [];
  for (const { coll, art } of FLAT) {
    let score = 0;
    if (art.title.toLowerCase().includes(q)) score += 30;
    if (coll.title.toLowerCase().includes(q)) score += 5;
    let snippet = "";
    for (const s of art.sections) {
      if (s.h.toLowerCase().includes(q)) { score += 12; snippet = snippet || stripTags(s.html).slice(0, 140); }
      const txt = stripTags(s.html).toLowerCase();
      const i = txt.indexOf(q);
      if (i >= 0) {
        score += 3;
        if (!snippet) snippet = "…" + stripTags(s.html).slice(Math.max(0, i - 40), i + 100) + "…";
      }
    }
    if (score > 0) out.push({ coll, art, score, snippet });
  }
  out.sort((a, b) => b.score - a.score);
  return out.slice(0, 12);
}

let root = null;
let els = {};
let state = { collId: null, artId: null, query: "" };
let spy = null;

function esc(s) {
  return String(s).replace(/[&<>"']/g, (c) => ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[c]));
}

function build() {
  if (root) return;
  root = document.createElement("div");
  root.id = "cetas-docs";
  root.setAttribute("role", "dialog");
  root.setAttribute("aria-label", "Centre d'aide");
  root.innerHTML = `
    <header class="docs-header">
      <button class="docs-burger" id="docs-burger" aria-label="Menu">${ICONS.chevR}</button>
      <div class="docs-brand"><span class="docs-brand-mark">✳</span><span>Centre d'aide</span></div>
      <div class="docs-search-wrap">
        <span class="docs-search-icon">${ICONS.search}</span>
        <input id="docs-search" class="docs-search" type="search" placeholder="Rechercher dans l'aide…" autocomplete="off" aria-label="Rechercher dans l'aide">
        <kbd class="docs-search-kbd">Ctrl K</kbd>
      </div>
      <button class="docs-close" id="docs-close" aria-label="Fermer l'aide">${ICONS.close}</button>
    </header>
    <div class="docs-body">
      <aside class="docs-sidebar" id="docs-sidebar" aria-label="Collections"></aside>
      <main class="docs-main" id="docs-main" tabindex="-1"></main>
      <aside class="docs-toc" id="docs-toc" aria-label="Sur cette page"></aside>
    </div>`;
  document.body.appendChild(root);
  els = {
    sidebar: root.querySelector("#docs-sidebar"),
    main: root.querySelector("#docs-main"),
    toc: root.querySelector("#docs-toc"),
    search: root.querySelector("#docs-search"),
    close: root.querySelector("#docs-close"),
    burger: root.querySelector("#docs-burger"),
  };
  renderSidebar();
  els.close.addEventListener("click", closeDocs);
  els.burger.addEventListener("click", () => els.sidebar.classList.toggle("open"));
  els.search.addEventListener("input", () => {
    state.query = els.search.value;
    if (state.query.trim()) renderSearch();
    else renderArticle(state.collId, state.artId);
  });
  els.search.addEventListener("keydown", (e) => {
    if (e.key === "Escape") { els.search.value = ""; state.query = ""; renderArticle(state.collId, state.artId); }
  });
  root.addEventListener("click", (e) => {
    const go = e.target.closest("[data-docs-art]");
    if (go) { e.preventDefault(); openArticle(go.dataset.docsColl, go.dataset.docsArt); return; }
    const toc = e.target.closest("[data-docs-sec]");
    if (toc) {
      e.preventDefault();
      const t = els.main.querySelector("#" + CSS.escape(toc.dataset.docsSec));
      if (t) t.scrollIntoView({ block: "start", behavior: "smooth" });
    }
  });
}

/* Le listener clavier est armé dans openDocs (réarmé à chaque ouverture). */

function onKey(e) {
  if (!root || !root.classList.contains("open")) return;
  if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "k") {
    e.preventDefault();
    els.search.focus();
    els.search.select();
  } else if (e.key === "Escape" && document.activeElement !== els.search) {
    closeDocs();
  }
}

function renderSidebar() {
  els.sidebar.innerHTML = CETAS_DOCS.collections.map((c) => `
    <div class="docs-coll${c.id === state.collId ? " active" : ""}" data-coll="${esc(c.id)}">
      <button class="docs-coll-head" data-coll-toggle="${esc(c.id)}" aria-expanded="${c.id === state.collId}">
        <span class="docs-coll-icon">${ICONS[c.icon] || ICONS.book}</span>
        <span class="docs-coll-title">${esc(c.title)}</span>
        <span class="docs-coll-chev">${ICONS.chevR}</span>
      </button>
      <div class="docs-coll-arts">
        ${c.articles.map((a) => `
          <button class="docs-art-link${a.id === state.artId ? " active" : ""}"
                  data-docs-art="${esc(a.id)}" data-docs-coll="${esc(c.id)}">${esc(a.title)}</button>`).join("")}
      </div>
    </div>`).join("");
  els.sidebar.querySelectorAll("[data-coll-toggle]").forEach((b) => {
    b.addEventListener("click", () => {
      const id = b.dataset.collToggle;
      state.collId = (state.collId === id) ? null : id;
      renderSidebar();
    });
  });
}

function crumbs(coll, art) {
  return `
    <nav class="docs-crumbs" aria-label="Fil d'Ariane">
      <button data-docs-home>Toutes les collections</button>
      <span class="docs-crumb-sep">${ICONS.chevR}</span>
      <button data-docs-art="${esc(art.id)}" data-docs-coll="${esc(coll.id)}">${esc(coll.title)}</button>
      <span class="docs-crumb-sep">${ICONS.chevR}</span>
      <span class="docs-crumb-current">${esc(art.title)}</span>
    </nav>`;
}

function renderHome() {
  state.collId = null; state.artId = null;
  renderSidebar();
  els.toc.innerHTML = "";
  els.toc.style.display = "none";
  els.main.innerHTML = `
    <div class="docs-home">
      <h1 class="docs-h1">Comment pouvons-nous vous aider ?</h1>
      <p class="docs-lead">La documentation complète de Cetas Lite, organisée par modules.</p>
      <div class="docs-cards">
        ${CETAS_DOCS.collections.map((c) => `
          <button class="docs-card" data-docs-art="${esc(c.articles[0].id)}" data-docs-coll="${esc(c.id)}">
            <span class="docs-card-icon">${ICONS[c.icon] || ICONS.book}</span>
            <span class="docs-card-title">${esc(c.title)}</span>
            <span class="docs-card-count">${c.articles.length} article${c.articles.length > 1 ? "s" : ""}</span>
            <span class="docs-card-first">${esc(c.articles.slice(0, 3).map((a) => a.title).join(" · "))}</span>
          </button>`).join("")}
      </div>
    </div>`;
  els.main.scrollTop = 0;
}

function renderArticle(collId, artId) {
  const entry = (artId && BY_ID.get(artId)) || FLAT[0];
  const { coll, art } = entry;
  state.collId = coll.id; state.artId = art.id;
  renderSidebar();
  els.toc.style.display = "";
  els.main.innerHTML = `
    ${crumbs(coll, art)}
    <h1 class="docs-h1">${esc(art.title)}</h1>
    ${art.sections.map((s, i) => `
      <section class="docs-sec" id="sec-${i}">
        <h2 class="docs-h2">${esc(s.h)}</h2>
        <div class="docs-prose">${s.html}</div>
      </section>`).join("")}
    <nav class="docs-prevnext">${prevNext(entry)}</nav>`;
  els.toc.innerHTML = `
    <div class="docs-toc-title">Sur cette page</div>
    ${art.sections.map((s, i) => `<button class="docs-toc-link" data-docs-sec="sec-${i}">${esc(s.h)}</button>`).join("")}`;
  els.main.scrollTop = 0;
  setupSpy();
  // Le clic sur le fil d'Ariane "Toutes les collections" (sans data-docs-art).
  els.main.querySelector("[data-docs-home]")?.addEventListener("click", renderHome);
}

function prevNext(entry) {
  const i = FLAT.indexOf(entry);
  const prev = FLAT[i - 1], next = FLAT[i + 1];
  return `
    ${prev ? `<button class="docs-pn" data-docs-art="${esc(prev.art.id)}" data-docs-coll="${esc(prev.coll.id)}"><span class="docs-pn-label">← Précédent</span><span class="docs-pn-title">${esc(prev.art.title)}</span></button>` : "<span></span>"}
    ${next ? `<button class="docs-pn docs-pn-next" data-docs-art="${esc(next.art.id)}" data-docs-coll="${esc(next.coll.id)}"><span class="docs-pn-label">Suivant →</span><span class="docs-pn-title">${esc(next.art.title)}</span></button>` : "<span></span>"}`;
}

function renderSearch() {
  const q = state.query.trim();
  const results = searchDocs(q);
  els.toc.style.display = "none";
  els.main.innerHTML = `
    <div class="docs-results">
      <div class="docs-results-count">${results.length} résultat${results.length > 1 ? "s" : ""} pour « ${esc(q)} »</div>
      ${results.map(({ coll, art, snippet }) => `
        <button class="docs-result" data-docs-art="${esc(art.id)}" data-docs-coll="${esc(coll.id)}">
          <span class="docs-result-coll">${esc(coll.title)}</span>
          <span class="docs-result-title">${esc(art.title)}</span>
          ${snippet ? `<span class="docs-result-snippet">${esc(snippet)}</span>` : ""}
        </button>`).join("") || `<p class="docs-no-result">Aucun résultat. Essayez un autre mot-clé, ou parcourez les collections.</p>`}
    </div>`;
  els.main.scrollTop = 0;
}

function setupSpy() {
  if (spy) { spy.disconnect(); spy = null; }
  if (typeof IntersectionObserver === "undefined") return;
  const links = [...els.toc.querySelectorAll(".docs-toc-link")];
  const secs = [...els.main.querySelectorAll(".docs-sec")];
  spy = new IntersectionObserver((entries) => {
    for (const en of entries) {
      if (en.isIntersecting) {
        const id = en.target.id;
        links.forEach((l) => l.classList.toggle("active", l.dataset.docsSec === id));
      }
    }
  }, { root: els.main, rootMargin: "-20% 0px -70% 0px" });
  secs.forEach((s) => spy.observe(s));
}

export function openArticle(collId, artId) {
  build();
  els.search.value = ""; state.query = "";
  renderArticle(collId, artId);
  els.sidebar.classList.remove("open");
}

export function openDocs(articleId) {
  build();
  root.classList.add("open");
  document.body.classList.add("docs-open");
  // Un seul listener clavier, réarmé à chaque ouverture.
  document.removeEventListener("keydown", onKey);
  document.addEventListener("keydown", onKey);
  els.search.value = ""; state.query = "";
  if (articleId && BY_ID.has(articleId)) renderArticle(BY_ID.get(articleId).coll.id, articleId);
  else renderHome();
  els.sidebar.classList.remove("open");
}

export function closeDocs() {
  if (!root) return;
  root.classList.remove("open");
  document.body.classList.remove("docs-open");
  if (spy) { spy.disconnect(); spy = null; }
  document.removeEventListener("keydown", onKey);
}

// Pour les tests : réinitialise le singleton DOM.
export function __resetDocsForTests() {
  if (spy) { spy.disconnect(); spy = null; }
  document.removeEventListener("keydown", onKey);
  root?.remove();
  root = null; els = {};
  state = { collId: null, artId: null, query: "" };
}
