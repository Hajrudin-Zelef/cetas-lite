import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";
// jsdom : résolution standard (node_modules du projet) puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);

import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT = join(__dirname, "..");
const { CETAS_DOCS } = await import(`${ROOT}/docs-data.js`);
const docs = await import(`${ROOT}/docs.js`);

// ---------- contenu ----------

test("docs-data : 9 collections, sommaire complet, aucun doublon d'article", () => {
  assert.equal(CETAS_DOCS.collections.length, 9);
  const ids = new Set();
  let articles = 0;
  for (const c of CETAS_DOCS.collections) {
    assert.ok(c.id && c.title && c.articles.length > 0, `collection ${c.id}`);
    for (const a of c.articles) {
      articles++;
      assert.ok(!ids.has(a.id), `doublon ${a.id}`);
      ids.add(a.id);
      assert.ok(a.title, `titre manquant ${a.id}`);
      assert.ok(a.sections.length >= 2, `article ${a.id} : au moins 2 sections`);
      for (const s of a.sections) {
        assert.ok(s.h && s.html && s.html.includes("<"), `section vide ${a.id}`);
      }
    }
  }
  assert.ok(articles >= 40, `au moins 40 articles, vu ${articles}`);
});

test("docs-data : contenu public — aucune info d'infrastructure sensible", () => {
  const raw = JSON.stringify(CETAS_DOCS);
  const banned = [
    "127.0.0.1", "localhost:8787", "CETAS_LITE_", "$CETAS_LITE_HOME",
    "bbolt", "scrypt", "AES-256-GCM", "sk-", "Bearer ",
  ];
  for (const b of banned) {
    assert.ok(!raw.includes(b), `contenu sensible détecté : ${b}`);
  }
});

// ---------- recherche ----------

test("searchDocs : trouve par titre, extrait, requête trop courte", () => {
  const r1 = docs.searchDocs("worktree");
  assert.ok(r1.length >= 1, "worktree trouvé");
  assert.ok(r1[0].art.title.toLowerCase().includes("worktree") || r1[0].score >= 30);
  assert.equal(docs.searchDocs("a").length, 0, "1 caractère = rien");
  assert.equal(docs.searchDocs("   ").length, 0, "vide = rien");
  const r2 = docs.searchDocs("GitHub");
  assert.ok(r2.some((r) => r.art.id === "outils-github"), "article outils-github trouvé");
  assert.ok(r2[0].snippet.length > 0, "extrait présent");
});

// ---------- vue ----------

test("openDocs : construit la vue, accueil avec 9 cartes", () => {
  docs.__resetDocsForTests();
  docs.openDocs();
  const root = document.getElementById("cetas-docs");
  assert.ok(root, "#cetas-docs créé");
  assert.ok(root.classList.contains("open"), "visible");
  assert.ok(document.body.classList.contains("docs-open"), "scroll fond bloqué");
  const cards = root.querySelectorAll(".docs-card");
  assert.equal(cards.length, 9, "9 cartes collections");
  assert.ok(root.querySelector("#docs-search"), "champ recherche présent");
  assert.ok(root.querySelector(".docs-sidebar .docs-coll"), "sidebar collections");
});

test("openDocs(article) : article, fil d'Ariane, sommaire", () => {
  docs.__resetDocsForTests();
  docs.openDocs("permissions");
  const root = document.getElementById("cetas-docs");
  const h1 = root.querySelector(".docs-h1");
  assert.equal(h1.textContent, "Permissions de l'agent");
  const crumbs = root.querySelector(".docs-crumbs").textContent;
  assert.ok(crumbs.includes("Toutes les collections"), "fil d'Ariane racine");
  assert.ok(crumbs.includes("Agents"), "fil d'Ariane collection");
  const tocLinks = root.querySelectorAll(".docs-toc-link");
  const secs = root.querySelectorAll(".docs-sec");
  assert.equal(tocLinks.length, secs.length, "sommaire = sections");
  assert.ok(tocLinks.length >= 2);
  // navigation précédent / suivant
  assert.ok(root.querySelector(".docs-prevnext"), "nav prev/next");
});

test("clic sidebar : navigue vers l'article", () => {
  docs.__resetDocsForTests();
  docs.openDocs();
  const root = document.getElementById("cetas-docs");
  // déplie la collection Agents puis clique "Modes Plan et Build"
  const head = root.querySelector('[data-coll-toggle="agents"]');
  head.click();
  const link = root.querySelector('[data-docs-art="plan-build"]');
  assert.ok(link, "lien plan-build visible");
  link.click();
  assert.equal(root.querySelector(".docs-h1").textContent, "Modes Plan et Build");
});

test("recherche dans la vue : résultats puis retour", () => {
  docs.__resetDocsForTests();
  docs.openDocs("permissions");
  const root = document.getElementById("cetas-docs");
  const input = root.querySelector("#docs-search");
  input.value = "sftp";
  input.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  const results = root.querySelectorAll(".docs-result");
  assert.ok(results.length >= 1, "au moins un résultat sftp");
  // clic sur le premier résultat -> article correspondant
  results[0].click();
  assert.ok(root.querySelector(".docs-h1"), "article affiché après clic");
  // Échap dans la recherche : vide et restaure l'article
  input.value = "";
  input.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  assert.equal(root.querySelectorAll(".docs-result").length, 0, "résultats effacés");
});

test("closeDocs : masque la vue et libère le scroll", () => {
  docs.__resetDocsForTests();
  docs.openDocs();
  docs.closeDocs();
  const root = document.getElementById("cetas-docs");
  assert.ok(!root.classList.contains("open"), "masquée");
  assert.ok(!document.body.classList.contains("docs-open"), "scroll restauré");
  // réouverture : le clavier (Ctrl+K / Échap) est réarmé
  docs.openDocs("bienvenue");
  assert.ok(document.getElementById("cetas-docs").classList.contains("open"));
  document.getElementById("cetas-docs").querySelector(".docs-close").click();
  assert.ok(!document.getElementById("cetas-docs").classList.contains("open"), "fermeture via bouton");
});

test("agents.js : menu avatar à 5 entrées + bouton docs", async () => {
  const src = await import("node:fs").then((fs) => fs.readFileSync(`${ROOT}/agents.js`, "utf8"));
  for (const id of ["mx-menu-settings", "mx-menu-help", "mx-menu-about", "mx-menu-faq", "mx-logout", "mx-docs-btn"]) {
    assert.ok(src.includes(`id="${id}"`), `élément ${id} présent`);
  }
  for (const label of ["Paramètre", "Obtenir de l'aide", "En savoir plus", "FAQ", "Déconnexion"]) {
    assert.ok(src.includes(label), `entrée "${label}" présente`);
  }
  assert.ok(src.includes('from "./docs.js"'), "agents.js importe docs.js");
  assert.ok(src.includes("openDocs()"), "aide / en savoir plus / bouton ? ouvrent le centre d'aide");
});
