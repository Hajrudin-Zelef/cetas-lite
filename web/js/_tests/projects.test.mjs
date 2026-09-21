import test from "node:test";
import assert from "node:assert/strict";
// jsdom : dependance de test declaree dans package.json (npm install).
import { JSDOM } from "jsdom";

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.window.confirm = () => true; // confirmDialog() y replie hors modale HTML.
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);

const PROJECTS = [
  { id: "p1", name: "Site vitrine", mode: "local", created_at: "2026-09-14T10:00:00Z" },
  { id: "p2", name: "API prod", mode: "sftp", host: "srv.ex", port: 22, user: "deploy", remote_path: "/srv/api", created_at: "2026-09-14T11:00:00Z" },
];
const TREE_ROOT = {
  name: "", path: "", is_dir: true, truncated: false,
  children: [
    { name: "src", path: "src", is_dir: true },
    { name: "README.md", path: "README.md", is_dir: false, size: 10 },
  ],
};
const TREE_SRC = {
  name: "src", path: "src", is_dir: true, truncated: false,
  children: [{ name: "main.go", path: "src/main.go", is_dir: false, size: 42 }],
};
const TREE_BIG = {
  name: "big", path: "big", is_dir: true, truncated: false,
  children: Array.from({ length: 250 }, (_, i) => ({
    name: "f" + i + ".txt", path: "big/f" + i + ".txt", is_dir: false, size: 1,
  })),
};
TREE_ROOT.children.push({ name: "big", path: "big", is_dir: true });
// NB : le backend renvoie le nœud vfs directement (pas d'enveloppe {tree}).
let activeId = "p1";
let ghConnected = false;
const calls = [];
globalThis.fetch = async (url, opts = {}) => {
  const u = String(url);
  const json = (body, ok = true) => ({
    ok,
    status: ok ? 200 : 400,
    json: async () => body,
    text: async () => JSON.stringify(body),
  });
  calls.push({ url: u, method: opts.method || "GET" });
  if (u === "/api/projects") return json({ projects: PROJECTS });
  if (u === "/api/projects/active") {
    if ((opts.method || "GET") === "PUT") {
      activeId = JSON.parse(opts.body).id;
      return json({ ok: true, active: activeId });
    }
    return json({ active: activeId });
  }
  if (u.startsWith("/api/projects/p1/tree")) {
    const qp = new URLSearchParams(u.split("?")[1] || "");
    const p = qp.get("path") || "";
    if (p === "src") return json(TREE_SRC);
    if (p === "big") return json(TREE_BIG);
    return json(TREE_ROOT);
  }
  if (u.startsWith("/api/projects/p1/file?path=logo.png")) {
    return { ok: true, status: 200, blob: async () => new Blob(["fakepng"]) };
  }
  if (u.startsWith("/api/projects/p1/file?path=src%2Fmain.go")) return json({ content: "package main", size: 12 });
  if (u === "/api/connectors") return json({ github: { connected: ghConnected, login: ghConnected ? "tester" : "" } });
  if (u === "/api/connectors/github" && (opts.method || "GET") === "PUT") { ghConnected = true; return json({ ok: true }); }
  if (u === "/api/connectors/github" && (opts.method || "GET") === "DELETE") { ghConnected = false; return json({ ok: true }); }
  return json({});
};

const { Projects, renderProjectsList, renderActiveTree, renderProjectBar, openFileReader, renderConnectorsInto, fileIcon } =
  await import("../projects.js");

const $ = (s) => document.querySelector(s);

await test("Projects.refresh charge la liste et l'actif", async () => {
  await Projects.refresh();
  assert.equal(Projects.list.length, 2);
  assert.equal(Projects.activeId, "p1");
  assert.equal(Projects.active.name, "Site vitrine");
});

await test("renderProjectsList affiche les projets avec leur mode", async () => {
  const box = document.createElement("div");
  const empty = document.createElement("div");
  document.body.append(box, empty);
  renderProjectsList(box, empty);
  assert.equal(box.querySelectorAll(".mx-project-row").length, 2);
  assert.match(box.innerHTML, /distant/);
  assert.match(box.innerHTML, /local/);
  assert.equal(empty.style.display, "none");
});

await test("setActive change le projet actif", async () => {
  await Projects.setActive("p2");
  assert.equal(Projects.activeId, "p2");
  assert.equal(Projects.active.host, "srv.ex");
  await Projects.setActive("p1");
});

await test("renderActiveTree affiche l'arborescence repliable (lazy)", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  assert.match(box.innerHTML, /src/);
  assert.match(box.innerHTML, /README\.md/);
  const dirBtn = box.querySelector(".sb-tree-dir");
  assert.ok(dirBtn, "dossier cliquable");
  const sub = box.querySelector(".sb-tree-sub");
  assert.ok(sub.classList.contains("collapsed"));
  assert.doesNotMatch(sub.innerHTML, /main\.go/, "enfants non chargés avant expansion");
  dirBtn.click();
  assert.ok(!sub.classList.contains("collapsed"), "déplié au clic");
  await new Promise((r) => setTimeout(r, 50));
  assert.match(sub.innerHTML, /main\.go/, "enfants chargés paresseusement");
  // Replie pour ne pas polluer les tests suivants (état déplié conservé).
  dirBtn.click();
});

await test("clic sur un fichier ouvre le lecteur", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  const dirBtn = box.querySelector(".sb-tree-dir");
  dirBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  const fileBtn = [...box.querySelectorAll(".sb-tree-item")].find((b) => b.title === "src/main.go");
  assert.ok(fileBtn, "bouton fichier trouvé");
  fileBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  const modal = document.querySelector(".mx-reader-modal");
  assert.ok(modal, "modale lecteur ouverte");
  assert.match(modal.innerHTML, /main\.go/);
  assert.match(document.querySelector(".mx-reader-code").textContent, /package main/);
  modal.closest(".mx-modal-overlay").remove();
  dirBtn.click(); // replie
});

await test("fileIcon distingue les types de fichiers", () => {
  const go = fileIcon("main.go");
  const md = fileIcon("README.md");
  const png = fileIcon("logo.png");
  const zip = fileIcon("a.zip");
  const unknown = fileIcon("fichier_sans_ext");
  assert.notEqual(go, md, "code != markdown");
  assert.notEqual(png, zip, "image != archive");
  assert.notEqual(go, unknown, "code != générique");
  assert.match(png, /<svg/, "icône image en SVG");
});

await test("clic sur une image ouvre l'aperçu image", async () => {
  globalThis.URL.createObjectURL = globalThis.URL.createObjectURL || (() => "blob:faux");
  globalThis.URL.revokeObjectURL = globalThis.URL.revokeObjectURL || (() => {});
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  // Ajoute temporairement une image à la racine mockée.
  TREE_ROOT.children.push({ name: "logo.png", path: "logo.png", is_dir: false, size: 7 });
  await renderActiveTree(box);
  const fileBtn = [...box.querySelectorAll(".sb-tree-item")].find((b) => b.title === "logo.png");
  assert.ok(fileBtn, "bouton image trouvé");
  fileBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  const img = document.querySelector(".mx-reader-img");
  assert.ok(img, "aperçu <img> affiché");
  document.querySelector(".mx-modal-overlay").remove();
  TREE_ROOT.children.pop();
});

await test("refresh conserve les dossiers dépliés", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  const dirBtn = box.querySelector(".sb-tree-dir");
  dirBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  // Second rendu (comme après un tour agent) : l'état déplié est restauré.
  await renderActiveTree(box);
  await new Promise((r) => setTimeout(r, 100));
  const sub = box.querySelector(".sb-tree-sub");
  assert.ok(!sub.classList.contains("collapsed"), "toujours déplié après refresh");
  assert.match(sub.innerHTML, /main\.go/, "enfants rechargés");
  box.querySelector(".sb-tree-dir").click(); // replie
});

await test("dossier volumineux : affichage par lots", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  const bigBtn = [...box.querySelectorAll(".sb-tree-dir")].find((b) => b.dataset.dirPath === "big");
  assert.ok(bigBtn, "dossier big trouvé");
  bigBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  const sub = bigBtn.nextElementSibling;
  assert.equal(
    sub.querySelectorAll(".sb-tree-item:not(.sb-tree-more)").length,
    100,
    "premier lot de 100"
  );
  const more = sub.querySelector(".sb-tree-more");
  assert.ok(more, "bouton + N autres présent");
  assert.match(more.textContent, /150 autres/);
  more.click();
  assert.equal(sub.querySelectorAll(".sb-tree-item:not(.sb-tree-more)").length, 200);
  assert.match(sub.querySelector(".sb-tree-more").textContent, /50 autres/);
  bigBtn.click(); // replie
});

await test("renderProjectBar montre le projet et le workspace", async () => {
  const bar = document.createElement("div");
  renderProjectBar(bar);
  assert.match(bar.innerHTML, /Site vitrine/);
  assert.match(bar.innerHTML, /Projet local/);
  await Projects.setActive("p2");
  renderProjectBar(bar);
  assert.match(bar.innerHTML, /deploy@srv\.ex:\/srv\/api/);
  await Projects.setActive("p1");
});

await test("renderConnectorsInto affiche les sections Connecté/Disponible", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderConnectorsInto(box);
  assert.match(box.innerHTML, /GitHub/);
  assert.match(box.innerHTML, /Disponible/);
  assert.ok(box.querySelector(".conn-assoc"), "bouton Associer présent");
  assert.ok(box.querySelector(".conn-menu-btn"), "menu ⋮ présent");
  const iconSvg = box.querySelector(".conn-icon svg");
  assert.ok(iconSvg, "logo GitHub en SVG (pas d'emoji)");
  assert.equal(iconSvg.getAttribute("viewBox"), "0 0 98 96");
});

await test("Associer ouvre la modale puis Connecter associe le connecteur", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderConnectorsInto(box);
  box.querySelector(".conn-assoc").click();
  const modal = document.querySelector("#conn-modal-overlay .conn-modal");
  assert.ok(modal, "modale Associer ouverte");
  const input = modal.querySelector("input");
  assert.ok(input, "champ token présent");
  input.value = "ghp_test";
  const connectBtn = [...modal.querySelectorAll("button")].find((b) => b.textContent === "Connecter");
  assert.ok(connectBtn, "bouton Connecter présent");
  connectBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  assert.ok(!document.querySelector("#conn-modal-overlay"), "modale fermée après connexion");
  assert.match(box.innerHTML, /Connecté/, "section Connecté affichée");
  assert.ok(!box.querySelector(".conn-assoc"), "plus de bouton Associer une fois connecté");
});

await test("menu ⋮ : clic extérieur ferme le menu, même après re-rendu", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderConnectorsInto(box);
  box.querySelector(".conn-menu-btn").click();
  const menu = box.querySelector(".conn-menu");
  assert.notEqual(menu.style.display, "none", "menu ouvert");
  document.body.click();
  assert.equal(menu.style.display, "none", "menu fermé au clic extérieur");
  // Un clic extérieur « perdu » (sans menu ouvert) ne doit pas casser la
  // fermeture suivante — le listener n'est plus en { once:true }.
  document.body.click();
  await renderConnectorsInto(box);
  box.querySelector(".conn-menu-btn").click();
  const menu2 = box.querySelector(".conn-menu");
  assert.notEqual(menu2.style.display, "none", "menu rouvert après re-rendu");
  document.body.click();
  assert.equal(menu2.style.display, "none", "fermeture fiable après re-rendu");
});

await test("menu ⋮ : favoris puis Dissocier", async () => {
  localStorage.removeItem("cetas-lite-conn-favs");
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderConnectorsInto(box);
  // Ajouter aux favoris via le menu ⋮.
  box.querySelector(".conn-menu-btn").click();
  const favItem = [...box.querySelectorAll(".conn-menu-item")].find((b) => b.textContent === "Ajouter aux favoris");
  assert.ok(favItem, "item favori présent");
  favItem.click();
  await new Promise((r) => setTimeout(r, 20));
  assert.ok(box.querySelector(".conn-star"), "étoile affichée");
  // Dissocier via le menu ⋮ (le connecteur est associé depuis le test précédent).
  box.querySelector(".conn-menu-btn").click();
  const disItem = [...box.querySelectorAll(".conn-menu-item")].find((b) => b.textContent === "Dissocier");
  assert.ok(disItem, "item Dissocier présent");
  disItem.click();
  await new Promise((r) => setTimeout(r, 50));
  assert.ok(box.querySelector(".conn-assoc"), "bouton Associer de retour après dissociation");
  assert.ok(box.querySelector(".conn-star"), "favori conservé après dissociation");
});

await test("la modale projet définit sa palette hors #marex-view (anti texte noir)", async () => {
  // L'overlay est inséré dans document.body : les variables du thème Agents
  // (#marex-view) n'y sont pas héritées. Ce test verrouille leur redéfinition
  // locale — sans elle, le texte tombe en noir sur fond noir (bug constaté).
  const { readFileSync } = await import("node:fs");
  const { fileURLToPath } = await import("node:url");
  const { dirname, join } = await import("node:path");
  const css = readFileSync(
    join(dirname(fileURLToPath(import.meta.url)), "..", "..", "css", "features", "marex-agents.css"),
    "utf8"
  );
  const style = document.createElement("style");
  style.textContent = css;
  document.head.appendChild(style);
  const overlay = document.createElement("div");
  overlay.className = "mx-modal-overlay";
  document.body.appendChild(overlay);
  const cs = window.getComputedStyle(overlay);
  assert.equal(cs.getPropertyValue("--text-primary").trim(), "#f2f2f4");
  assert.equal(cs.getPropertyValue("--bg-input-btn").trim(), "#1f2024");
  assert.equal(cs.getPropertyValue("--border").trim(), "#25262b");
  assert.equal(cs.getPropertyValue("--text-secondary").trim(), "#9a9ba3");
  const card = document.createElement("div");
  card.className = "mx-conn-card";
  document.body.appendChild(card);
  const cs2 = window.getComputedStyle(card);
  assert.ok(cs2.getPropertyValue("--text-primary").trim(), "carte connecteurs : --text-primary défini");
  assert.ok(cs2.getPropertyValue("--bg-input-btn").trim(), "carte connecteurs : --bg-input-btn défini");
  overlay.remove();
  card.remove();
  style.remove();
});
