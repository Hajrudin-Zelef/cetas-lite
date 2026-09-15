import test from "node:test";
import assert from "node:assert/strict";
// ESM : NODE_PATH ne s'applique pas aux imports, chemin absolu requis.
import { JSDOM } from "/tmp/node_modules/jsdom/lib/api.js";

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);

const PROJECTS = [
  { id: "p1", name: "Site vitrine", mode: "local", created_at: "2026-09-14T10:00:00Z" },
  { id: "p2", name: "API prod", mode: "sftp", host: "srv.ex", port: 22, user: "deploy", remote_path: "/srv/api", created_at: "2026-09-14T11:00:00Z" },
];
const TREE = {
  tree: {
    name: "/", path: "", is_dir: true,
    children: [
      {
        name: "src", path: "src", is_dir: true,
        children: [{ name: "main.go", path: "src/main.go", is_dir: false, size: 42 }],
      },
      { name: "README.md", path: "README.md", is_dir: false, size: 10 },
    ],
  },
  truncated: false,
};
let activeId = "p1";
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
  if (u === "/api/projects/p1/tree") return json(TREE);
  if (u.startsWith("/api/projects/p1/file?path=src%2Fmain.go")) return json({ content: "package main", size: 12 });
  if (u === "/api/connectors") return json({ github: { connected: false } });
  return json({});
};

const { Projects, renderProjectsList, renderActiveTree, renderProjectBar, openFileReader, renderConnectorsInto } =
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

await test("renderActiveTree affiche l'arborescence repliable", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  assert.match(box.innerHTML, /src/);
  assert.match(box.innerHTML, /README\.md/);
  const dirBtn = box.querySelector(".sb-tree-dir");
  assert.ok(dirBtn, "dossier cliquable");
  const sub = box.querySelector(".sb-tree-sub");
  assert.ok(sub.classList.contains("collapsed"));
  dirBtn.click();
  assert.ok(!sub.classList.contains("collapsed"), "déplié au clic");
});

await test("clic sur un fichier ouvre le lecteur", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderActiveTree(box);
  const fileBtn = [...box.querySelectorAll(".sb-tree-item")].find((b) => b.title === "src/main.go");
  assert.ok(fileBtn, "bouton fichier trouvé");
  fileBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  const modal = document.querySelector(".mx-reader-modal");
  assert.ok(modal, "modale lecteur ouverte");
  assert.match(modal.innerHTML, /main\.go/);
  assert.match(document.querySelector(".mx-reader-code").textContent, /package main/);
  modal.closest(".mx-modal-overlay").remove();
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

await test("renderConnectorsInto affiche la carte GitHub", async () => {
  const box = document.createElement("div");
  document.body.appendChild(box);
  await renderConnectorsInto(box);
  assert.match(box.innerHTML, /GitHub/);
  assert.ok(box.querySelector("#mx-gh-token"), "champ token présent");
  assert.ok(box.querySelector("#mx-gh-connect"), "bouton connecter présent");
});
