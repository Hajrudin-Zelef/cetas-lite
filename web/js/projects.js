// ============================================================================
// Projets de l'agent : upload local, dossier distant SFTP, arborescence,
// lecteur de fichiers (texte/code/PDF) et connecteurs (GitHub).
// Câblé sur /api/projects et /api/connectors.
// ============================================================================
import { api, getToken } from "./api.js";
import { confirmDialog, alertDialog } from "./dialogs.js";

function esc(s) {
  return String(s == null ? "" : s)
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;");
}

const I = {
  folder: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 2 2H5a2 2 0 0 1-2-2V7z"/></svg>',
  file: '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>',
  chev: '<svg width="11" height="11" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 9l6 6 6-6"/></svg>',
  close: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>',
  trash: '<svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/></svg>',
  server: '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="2" y="2" width="20" height="8" rx="2"/><rect x="2" y="14" width="20" height="8" rx="2"/><line x1="6" y1="6" x2="6" y2="6"/><line x1="6" y1="18" x2="6" y2="18"/></svg>',
};

// ---------------- API ----------------

export const ProjectsAPI = {
  list: () => api("/api/projects"),
  create: (body) => api("/api/projects", { method: "POST", body }),
  remove: (id) => api("/api/projects/" + encodeURIComponent(id), { method: "DELETE" }),
  tree: (id) => api("/api/projects/" + encodeURIComponent(id) + "/tree"),
  file: (id, path) =>
    api("/api/projects/" + encodeURIComponent(id) + "/file?path=" + encodeURIComponent(path)),
  removeFile: (id, path) =>
    api("/api/projects/" + encodeURIComponent(id) + "/file?path=" + encodeURIComponent(path), {
      method: "DELETE",
    }),
  sftpTest: (body) => api("/api/projects/sftp-test", { method: "POST", body }),
  getActive: () => api("/api/projects/active"),
  setActive: (id) => api("/api/projects/active", { method: "PUT", body: { id } }),
  connectors: () => api("/api/connectors"),
  githubConnect: (token) => api("/api/connectors/github", { method: "PUT", body: { token } }),
  githubDisconnect: () => api("/api/connectors/github", { method: "DELETE" }),

  // Upload multipart (FormData) : fetch brut avec le token.
  async upload(id, formData) {
    const token = getToken();
    const resp = await fetch("/api/projects/" + encodeURIComponent(id) + "/upload", {
      method: "POST",
      headers: token ? { Authorization: "Bearer " + token } : {},
      body: formData,
    });
    const data = await resp.json().catch(() => ({}));
    if (!resp.ok) throw new Error(data.error || "Échec de l'upload (" + resp.status + ")");
    return data;
  },

  // Récupère un PDF en blob (l'iframe ne peut pas envoyer le header Auth).
  async fileBlob(id, path) {
    const token = getToken();
    const resp = await fetch(
      "/api/projects/" + encodeURIComponent(id) + "/file?path=" + encodeURIComponent(path),
      { headers: token ? { Authorization: "Bearer " + token } : {} }
    );
    if (!resp.ok) throw new Error("Fichier introuvable (" + resp.status + ")");
    return await resp.blob();
  },
};

// ---------------- état ----------------

export const Projects = {
  list: [],
  activeId: "",
  onChange: null, // callback après refresh / changement actif

  get active() {
    return this.list.find((p) => p.id === this.activeId) || null;
  },

  async refresh() {
    try {
      const [l, a] = await Promise.all([ProjectsAPI.list(), ProjectsAPI.getActive()]);
      this.list = (l && l.projects) || [];
      this.activeId = (a && a.active) || "";
    } catch (e) {
      this.list = [];
      this.activeId = "";
    }
    if (this.onChange) this.onChange();
  },

  async setActive(id) {
    await ProjectsAPI.setActive(id || "");
    this.activeId = id || "";
    if (this.onChange) this.onChange();
  },
};

// ---------------- arborescence ----------------

function fileIcon(name) {
  return '<span class="mx-tree-file-dot"></span>';
}

function renderTreeNodes(container, node, projectId, prefix) {
  const kids = node.children || [];
  for (const c of kids) {
    const path = c.path || (prefix ? prefix + "/" + c.name : c.name);
    if (c.is_dir) {
      const dirBtn = document.createElement("button");
      dirBtn.type = "button";
      dirBtn.className = "sb-tree-item sb-tree-dir";
      dirBtn.innerHTML = I.chev + I.folder + "<span>" + esc(c.name) + "</span>";
      const sub = document.createElement("div");
      sub.className = "sb-tree-sub sb-ws-tree collapsed";
      dirBtn.addEventListener("click", () => {
        const collapsed = sub.classList.toggle("collapsed");
        dirBtn.querySelector("svg").style.transform = collapsed ? "rotate(-90deg)" : "";
      });
      container.appendChild(dirBtn);
      container.appendChild(sub);
      renderTreeNodes(sub, c, projectId, path);
    } else {
      const fBtn = document.createElement("button");
      fBtn.type = "button";
      fBtn.className = "sb-tree-item";
      fBtn.innerHTML = '<span class="sb-tree-leaf-spacer"></span>' + I.file + "<span>" + esc(c.name) + "</span>";
      fBtn.title = path;
      fBtn.addEventListener("click", () => openFileReader(projectId, path));
      container.appendChild(fBtn);
    }
  }
}

// Affiche l'arborescence du projet actif dans le conteneur donné.
export async function renderActiveTree(container) {
  container.innerHTML = "";
  const p = Projects.active;
  if (!p) {
    container.innerHTML = '<div class="sb-tree-empty">Aucun projet actif.</div>';
    return;
  }
  const loading = document.createElement("div");
  loading.className = "sb-tree-empty";
  loading.textContent = "Chargement…";
  container.appendChild(loading);
  try {
    const data = await ProjectsAPI.tree(p.id);
    container.innerHTML = "";
    const root = (data && data.tree) || { children: [] };
    if (data && data.truncated) {
      const w = document.createElement("div");
      w.className = "sb-tree-empty";
      w.textContent = "Arborescence tronquée (2000 entrées max).";
      container.appendChild(w);
    }
    if (!(root.children || []).length) {
      container.innerHTML += '<div class="sb-tree-empty">Projet vide.</div>';
      return;
    }
    renderTreeNodes(container, root, p.id, "");
  } catch (e) {
    container.innerHTML =
      '<div class="sb-tree-empty">Lecture impossible : ' + esc(e.message || e) + "</div>";
  }
}

// ---------------- liste des projets ----------------

export function renderProjectsList(container, emptyEl, opts = {}) {
  container.innerHTML = "";
  const list = Projects.list;
  if (emptyEl) emptyEl.style.display = list.length ? "none" : "";
  for (const p of list) {
    const row = document.createElement("div");
    row.className = "mx-project-row" + (p.id === Projects.activeId ? " active" : "");
    const b = document.createElement("button");
    b.type = "button";
    b.className = "sb-tree-item";
    b.innerHTML =
      (p.mode === "sftp" ? I.server : I.folder) +
      "<span>" +
      esc(p.name) +
      '</span><span class="mx-project-mode">' +
      esc(p.mode === "sftp" ? "distant" : "local") +
      "</span>";
    b.title = p.mode === "sftp" ? p.user + "@" + p.host + ":" + p.remote_path : "Projet local";
    b.addEventListener("click", async () => {
      try {
        await Projects.setActive(p.id);
      } catch (e) {
        alert("Sélection impossible : " + (e.message || e));
      }
    });
    const del = document.createElement("button");
    del.type = "button";
    del.className = "mx-project-del";
    del.innerHTML = I.trash;
    del.title = "Supprimer le projet";
    del.addEventListener("click", async (ev) => {
      ev.stopPropagation();
      if (!confirm('Supprimer le projet "' + p.name + '" ?' + (p.mode === "local" ? " Les fichiers importés seront effacés." : "")))
        return;
      try {
        await ProjectsAPI.remove(p.id);
        await Projects.refresh();
      } catch (e) {
        alert("Suppression impossible : " + (e.message || e));
      }
    });
    row.appendChild(b);
    row.appendChild(del);
    container.appendChild(row);
  }
  if (opts.onRendered) opts.onRendered();
}

// ---------------- barre projet au-dessus du textarea ----------------

export function renderProjectBar(barEl) {
  const p = Projects.active;
  if (!p) {
    barEl.innerHTML =
      '<span class="mx-pbar-icon">' + I.folder + '</span><span class="mx-pbar-name">Espace partagé</span>';
    barEl.title = "Aucun projet : l'agent travaille dans l'espace partagé.";
    return;
  }
  const sub =
    p.mode === "sftp"
      ? esc(p.user) + "@" + esc(p.host) + ":" + esc(p.remote_path)
      : "Projet local importé";
  barEl.innerHTML =
    '<span class="mx-pbar-icon">' +
    (p.mode === "sftp" ? I.server : I.folder) +
    '</span><span class="mx-pbar-name">' +
    esc(p.name) +
    '</span><span class="mx-pbar-sub">' +
    sub +
    "</span>";
  barEl.title = p.mode === "sftp" ? "Dossier distant via SFTP" : "Fichiers importés localement";
}

// ---------------- modal lecteur de fichier ----------------

function langOf(name) {
  const m = /\.([a-z0-9]+)$/i.exec(name || "");
  return (m && m[1].toLowerCase()) || "";
}

export async function openFileReader(projectId, path) {
  const overlay = document.createElement("div");
  overlay.className = "mx-modal-overlay";
  const name = path.split("/").pop();
  overlay.innerHTML =
    '<div class="mx-modal mx-reader-modal" role="dialog" aria-label="Lecteur de fichier">' +
    '<div class="mx-modal-head"><span class="mx-modal-title">' +
    esc(name) +
    '</span><span class="mx-modal-path">' +
    esc(path) +
    '</span><button class="mx-modal-close" type="button">' +
    I.close +
    "</button></div>" +
    '<div class="mx-reader-body"><div class="sb-tree-empty">Chargement…</div></div>' +
    "</div>";
  document.body.appendChild(overlay);
  const close = () => overlay.remove();
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) close();
  });
  overlay.querySelector(".mx-modal-close").addEventListener("click", close);
  document.addEventListener("keydown", function esc2(e) {
    if (e.key === "Escape") {
      close();
      document.removeEventListener("keydown", esc2);
    }
  });
  const body = overlay.querySelector(".mx-reader-body");
  try {
    if (/\.pdf$/i.test(name)) {
      const blob = await ProjectsAPI.fileBlob(projectId, path);
      const url = URL.createObjectURL(blob);
      body.innerHTML = '<iframe class="mx-reader-pdf" src="' + url + '"></iframe>';
      const iv = setInterval(() => {
        if (!document.body.contains(overlay)) {
          URL.revokeObjectURL(url);
          clearInterval(iv);
        }
      }, 2000);
      return;
    }
    const data = await ProjectsAPI.file(projectId, path);
    if (data.binary) {
      body.innerHTML =
        '<div class="sb-tree-empty">Fichier binaire (' +
        (data.size || 0) +
        " octets) — aperçu non disponible.</div>";
      return;
    }
    const pre = document.createElement("pre");
    pre.className = "mx-reader-code lang-" + esc(langOf(name));
    pre.textContent = data.content || "";
    body.innerHTML = "";
    body.appendChild(pre);
  } catch (e) {
    body.innerHTML = '<div class="sb-tree-empty">Lecture impossible : ' + esc(e.message || e) + "</div>";
  }
}

// ---------------- modal nouveau projet ----------------

export function openNewProjectModal(onCreated, opts) {
  const overlay = document.createElement("div");
  overlay.className = "mx-modal-overlay";
  overlay.innerHTML =
    '<div class="mx-modal" role="dialog" aria-label="Nouveau projet">' +
    '<div class="mx-modal-head"><span class="mx-modal-title">Nouveau projet</span>' +
    '<button class="mx-modal-close" type="button">' + I.close + "</button></div>" +
    '<div class="mx-modal-tabs">' +
    '<button type="button" class="mx-modal-tab active" data-tab="upload">Importer des fichiers</button>' +
    '<button type="button" class="mx-modal-tab" data-tab="sftp">Serveur distant (SFTP)</button>' +
    "</div>" +
    '<div class="mx-modal-pane" data-pane="upload">' +
    '<label class="mx-field"><span>Nom du projet</span><input id="mx-np-name" type="text" placeholder="mon-projet" maxlength="80"></label>' +
    '<label class="mx-field"><span>Fichiers</span><input id="mx-np-files" type="file" multiple></label>' +
    '<label class="mx-field"><span>Dossier entier</span><input id="mx-np-dir" type="file" webkitdirectory></label>' +
    '<div class="mx-hint">Les fichiers sont copiés sur le serveur. 200 Mo max par import.</div>' +
    '<div class="mx-modal-error" id="mx-np-error"></div>' +
    '<div class="mx-modal-foot"><button type="button" class="mx-btn primary" id="mx-np-create-upload">Créer et importer</button></div>' +
    "</div>" +
    '<div class="mx-modal-pane hidden" data-pane="sftp">' +
    '<label class="mx-field"><span>Nom du projet</span><input id="mx-np-sftp-name" type="text" placeholder="serveur-prod" maxlength="80"></label>' +
    '<div class="mx-field-row">' +
    '<label class="mx-field"><span>Hôte</span><input id="mx-np-host" type="text" placeholder="exemple.com" autocomplete="off"></label>' +
    '<label class="mx-field mx-field-sm"><span>Port</span><input id="mx-np-port" type="number" value="22" min="1" max="65535"></label>' +
    "</div>" +
    '<div class="mx-field-row">' +
    '<label class="mx-field"><span>Utilisateur</span><input id="mx-np-user" type="text" placeholder="deploy" autocomplete="username"></label>' +
    '<label class="mx-field"><span>Dossier distant</span><input id="mx-np-rpath" type="text" placeholder="/srv/app"></label>' +
    "</div>" +
    '<label class="mx-field"><span>Authentification</span><select id="mx-np-auth"><option value="password">Mot de passe</option><option value="key">Clé privée</option></select></label>' +
    '<label class="mx-field" id="mx-np-pw-wrap"><span>Mot de passe</span><input id="mx-np-password" type="password" autocomplete="current-password"></label>' +
    '<label class="mx-field hidden" id="mx-np-key-wrap"><span>Clé privée (PEM)</span><textarea id="mx-np-key" rows="4" placeholder="-----BEGIN OPENSSH PRIVATE KEY-----"></textarea></label>' +
    '<label class="mx-field hidden" id="mx-np-pp-wrap"><span>Phrase secrète (optionnel)</span><input id="mx-np-passphrase" type="password"></label>' +
    '<div class="mx-hint">Le mot de passe / la clé sont chiffrés avant stockage. Premier contact : l\'empreinte de l\'hôte vous sera présentée (TOFU).</div>' +
    '<div class="mx-modal-error" id="mx-np-sftp-error"></div>' +
    '<div class="mx-modal-foot"><button type="button" class="mx-btn" id="mx-np-test">Tester</button>' +
    '<button type="button" class="mx-btn primary" id="mx-np-create-sftp">Créer</button></div>' +
    "</div>" +
    "</div>";
  document.body.appendChild(overlay);
  const close = () => overlay.remove();
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) close();
  });
  overlay.querySelector(".mx-modal-close").addEventListener("click", close);
  const $ = (sel) => overlay.querySelector(sel);

  // Onglets
  const switchTab = (name) => {
    overlay.querySelectorAll(".mx-modal-tab").forEach((x) => x.classList.remove("active"));
    overlay.querySelectorAll(".mx-modal-pane").forEach((x) => x.classList.add("hidden"));
    const t = overlay.querySelector('.mx-modal-tab[data-tab="' + name + '"]');
    const p = overlay.querySelector('[data-pane="' + name + '"]');
    if (t) t.classList.add("active");
    if (p) p.classList.remove("hidden");
  };
  overlay.querySelectorAll(".mx-modal-tab").forEach((t) =>
    t.addEventListener("click", () => switchTab(t.dataset.tab))
  );
  // Onglet initial (ex. "sftp" quand on vient du volet Remote).
  if (opts && opts.tab) switchTab(opts.tab);
  // Bascule mot de passe / clé
  $("#mx-np-auth").addEventListener("change", (e) => {
    const isKey = e.target.value === "key";
    $("#mx-np-pw-wrap").classList.toggle("hidden", isKey);
    $("#mx-np-key-wrap").classList.toggle("hidden", !isKey);
    $("#mx-np-pp-wrap").classList.toggle("hidden", !isKey);
  });

  const errBox = (pane) => overlay.querySelector(pane === "sftp" ? "#mx-np-sftp-error" : "#mx-np-error");
  const fail = (pane, msg) => {
    errBox(pane).textContent = msg;
  };

  // ---- import local ----
  $("#mx-np-create-upload").addEventListener("click", async () => {
    fail("upload", "");
    const name = $("#mx-np-name").value.trim();
    if (!name) return fail("upload", "Donnez un nom au projet.");
    const files = [...$("#mx-np-files").files, ...$("#mx-np-dir").files];
    if (!files.length) return fail("upload", "Sélectionnez des fichiers ou un dossier.");
    const btn = $("#mx-np-create-upload");
    btn.disabled = true;
    btn.textContent = "Import…";
    try {
      const created = await ProjectsAPI.create({ name, mode: "local" });
      const fd = new FormData();
      const paths = [];
      for (const f of files) {
        fd.append("files", f, f.name);
        paths.push(f.webkitRelativePath || f.name);
      }
      fd.append("paths", JSON.stringify(paths));
      const res = await ProjectsAPI.upload(created.id, fd);
      await Projects.refresh();
      await Projects.setActive(created.id);
      close();
      if (onCreated) onCreated(created);
      if (res.skipped)
        alert(res.count + " fichier(s) importé(s), " + res.skipped + " ignoré(s) (chemin invalide).");
    } catch (e) {
      fail("upload", e.message || String(e));
      btn.disabled = false;
      btn.textContent = "Créer et importer";
    }
  });

  // ---- SFTP ----
  const sftpBody = (trust) => {
    const isKey = $("#mx-np-auth").value === "key";
    const body = {
      name: $("#mx-np-sftp-name").value.trim(),
      host: $("#mx-np-host").value.trim(),
      port: parseInt($("#mx-np-port").value, 10) || 22,
      user: $("#mx-np-user").value.trim(),
      remote_path: $("#mx-np-rpath").value.trim(),
    };
    if (isKey) {
      body.private_key = $("#mx-np-key").value;
      body.passphrase = $("#mx-np-passphrase").value || undefined;
    } else {
      body.password = $("#mx-np-password").value;
    }
    if (trust) body.trust_host_key = trust;
    return body;
  };
  const validateSftp = (body) => {
    if (!body.name) return "Donnez un nom au projet.";
    if (!body.host) return "Indiquez l'hôte.";
    if (!body.user) return "Indiquez l'utilisateur.";
    if (!body.remote_path) return "Indiquez le dossier distant.";
    if (body.password === "" && !body.private_key) return "Mot de passe ou clé privée requis.";
    return "";
  };
  // Test avec gestion TOFU : si l'empreinte est inconnue, on la présente.
  async function sftpTestFlow(body) {
    try {
      return await ProjectsAPI.sftpTest(body);
    } catch (e) {
      const msg = String((e && e.message) || e);
      const m = /Nouvel h.te[^:]*: ([A-Za-z0-9+/=]+)/.exec(msg);
      if (m && confirm("Nouvel hôte SFTP.\nEmpreinte de la clé : " + m[1] + "\n\nValider cet hôte ?")) {
        return await ProjectsAPI.sftpTest(Object.assign({}, body, { trust_host_key: m[1] }));
      }
      throw e;
    }
  }
  $("#mx-np-test").addEventListener("click", async () => {
    fail("sftp", "");
    const body = sftpBody();
    const v = validateSftp(body);
    if (v) return fail("sftp", v);
    const btn = $("#mx-np-test");
    btn.disabled = true;
    btn.textContent = "Test…";
    try {
      const res = await sftpTestFlow(body);
      fail("sftp", "");
      alert("Connexion réussie (" + res.entries + " entrée(s) dans " + res.remote_path + ").");
    } catch (e) {
      fail("sftp", e.message || String(e));
    } finally {
      btn.disabled = false;
      btn.textContent = "Tester";
    }
  });
  $("#mx-np-create-sftp").addEventListener("click", async () => {
    fail("sftp", "");
    const body = sftpBody();
    const v = validateSftp(body);
    if (v) return fail("sftp", v);
    const btn = $("#mx-np-create-sftp");
    btn.disabled = true;
    btn.textContent = "Création…";
    try {
      await sftpTestFlow(body); // valide (et TOFU) avant de créer
      const created = await ProjectsAPI.create(Object.assign({ mode: "sftp" }, body));
      await Projects.refresh();
      await Projects.setActive(created.id);
      close();
      if (onCreated) onCreated(created);
    } catch (e) {
      fail("sftp", e.message || String(e));
      btn.disabled = false;
      btn.textContent = "Créer";
    }
  });
}

// ---------------- connecteurs (Réglages) ----------------

// ---------------- connecteurs (Réglages) ----------------
// Registre des connecteurs : chaque entrée décrit le connecteur et les
// champs demandés par la modale « Associer ». Ajouter une entrée ici suffit
// pour exposer un nouveau connecteur.
const CONNECTOR_DEFS = [
  {
    id: "github",
    name: "GitHub",
    // Logo officiel (mark GitHub) en SVG inline, teinté via currentColor
    // pour suivre le thème clair/sombre.
    iconSvg:
      '<svg viewBox="0 0 98 96" fill="currentColor" aria-hidden="true"><path fill-rule="evenodd" clip-rule="evenodd" d="M48.854 0C21.839 0 0 22 0 49.217c0 21.756 13.993 40.172 33.405 46.69 2.427.49 3.316-1.059 3.316-2.362 0-1.141-.08-5.052-.08-9.127-13.59 2.934-16.42-5.867-16.42-5.867-2.184-5.704-5.42-7.17-5.42-7.17-4.448-3.015.324-3.015.324-3.015 4.934.326 7.523 5.052 7.523 5.052 4.367 7.496 11.404 5.378 14.235 4.074.404-3.178 1.699-5.378 3.074-6.6-10.839-1.141-22.243-5.378-22.243-24.283 0-5.378 1.94-9.778 5.014-13.2-.485-1.222-2.184-6.275.486-13.038 0 0 4.125-1.304 13.426 5.052a46.97 46.97 0 0 1 12.214-1.63c4.125 0 8.33.571 12.213 1.63 9.302-6.356 13.427-5.052 13.427-5.052 2.67 6.763.97 11.816.485 13.038 3.155 3.422 5.015 7.822 5.015 13.2 0 18.905-11.404 23.06-22.324 24.283 1.78 1.548 3.316 4.481 3.316 9.126 0 6.6-.08 11.897-.08 13.526 0 1.304.89 2.853 3.316 2.364 19.412-6.52 33.405-24.935 33.405-46.691C97.707 22 75.788 0 48.854 0z"/></svg>',
    desc: "Permet à l'agent d'utiliser git (commit, diff, push) sur vos dépôts. Le token est chiffré avant stockage.",
    fields: [
      { key: "token", label: "Token personnel (scope repo)", type: "password", placeholder: "ghp_…" },
    ],
    connectLabel: "Connecter",
    // state: { connected, login } — issu de GET /api/connectors
    statusOf: (st) => st || { connected: false },
    badgeOf: (st) => (st.connected ? "Connecté · " + (st.login || "") : "Déconnecté"),
    onConnect: async (values) => {
      const token = (values.token || "").trim();
      if (!token) throw new Error("Collez votre token GitHub.");
      await ProjectsAPI.githubConnect(token);
    },
    onDisconnect: async () => {
      await ProjectsAPI.githubDisconnect();
    },
  },
];

const CONN_FAVS_KEY = "cetas-lite-conn-favs";
function connFavs() {
  try {
    const v = JSON.parse(localStorage.getItem(CONN_FAVS_KEY) || "[]");
    return Array.isArray(v) ? v : [];
  } catch (e) {
    return [];
  }
}
function setConnFav(id, on) {
  let favs = connFavs();
  favs = on ? [...new Set([...favs, id])] : favs.filter((x) => x !== id);
  try {
    localStorage.setItem(CONN_FAVS_KEY, JSON.stringify(favs));
  } catch (e) {}
}

function connEl(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

// Modale « Associer » : saisie des champs du connecteur, validée par « Connecter ».
function openConnectorModal(def, done) {
  closeConnectorModal();
  const overlay = connEl("div", "sp-modal-overlay conn-modal-overlay");
  overlay.id = "conn-modal-overlay";
  const modal = connEl("div", "sp-modal conn-modal");
  modal.appendChild(connEl("h3", "conn-modal-title", "Associer " + def.name));
  modal.appendChild(connEl("p", "apikey-intro", def.desc));
  const fields = {};
  for (const f of def.fields || []) {
    const lab = connEl("label", "mx-field");
    lab.appendChild(connEl("span", null, f.label));
    const input = document.createElement("input");
    input.type = f.type || "text";
    input.placeholder = f.placeholder || "";
    input.autocomplete = "off";
    input.spellcheck = false;
    lab.appendChild(input);
    fields[f.key] = input;
    modal.appendChild(lab);
  }
  const errEl = connEl("div", "mx-modal-error");
  modal.appendChild(errEl);
  const row = connEl("div", "conn-modal-actions");
  const cancelBtn = connEl("button", "mx-btn", "Annuler");
  cancelBtn.type = "button";
  cancelBtn.addEventListener("click", closeConnectorModal);
  const okBtn = connEl("button", "mx-btn primary", def.connectLabel || "Connecter");
  okBtn.type = "button";
  okBtn.addEventListener("click", async () => {
    errEl.textContent = "";
    okBtn.disabled = true;
    okBtn.textContent = "Vérification…";
    try {
      const values = {};
      for (const k of Object.keys(fields)) values[k] = fields[k].value;
      await def.onConnect(values);
      closeConnectorModal();
      done();
    } catch (e) {
      errEl.textContent = e.message || String(e);
      okBtn.disabled = false;
      okBtn.textContent = def.connectLabel || "Connecter";
    }
  });
  row.appendChild(cancelBtn);
  row.appendChild(okBtn);
  modal.appendChild(row);
  overlay.appendChild(modal);
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) closeConnectorModal();
  });
  document.body.appendChild(overlay);
  const first = modal.querySelector("input");
  if (first) first.focus();
}

function closeConnectorModal() {
  document.getElementById("conn-modal-overlay")?.remove();
}

export async function renderConnectorsInto(container) {
  container.innerHTML = "";
  container.appendChild(connEl("div", "sb-tree-empty", "Chargement…"));
  let state = { github: { connected: false } };
  try {
    state = await ProjectsAPI.connectors();
  } catch (e) {
    container.innerHTML = "";
    container.appendChild(connEl("div", "sb-tree-empty", "Connecteurs injoignables."));
    return;
  }
  const favs = connFavs();
  const items = CONNECTOR_DEFS.map((def) => {
    const st = def.statusOf(state[def.id]);
    return { def, st, connected: !!st.connected, fav: favs.includes(def.id) };
  });
  const sortItems = (arr) =>
    arr.sort((a, b) => (b.fav - a.fav) || a.def.name.localeCompare(b.def.name));
  const connected = sortItems(items.filter((x) => x.connected));
  const available = sortItems(items.filter((x) => !x.connected));

  container.innerHTML = "";
  const rerender = () => renderConnectorsInto(container);

  const makeCard = (item) => {
    const { def, st, connected: isConn, fav } = item;
    const card = connEl("div", "conn-card" + (fav ? " fav" : ""));
    // Menu ⋮ devant le nom.
    const menuWrap = connEl("div", "conn-menu-wrap");
    const menuBtn = connEl("button", "conn-menu-btn", "⋮");
    menuBtn.type = "button";
    menuBtn.title = "Options";
    menuBtn.setAttribute("aria-label", "Options " + def.name);
    const menu = connEl("div", "conn-menu");
    menu.style.display = "none";
    const addItem = (label, cls, fn) => {
      const b = connEl("button", "conn-menu-item" + (cls ? " " + cls : ""), label);
      b.type = "button";
      b.addEventListener("click", async (e) => {
        e.stopPropagation();
        menu.style.display = "none";
        await fn();
      });
      menu.appendChild(b);
    };
    if (isConn) {
      addItem("Dissocier", "", async () => {
        const ok = await confirmDialog("Dissocier " + def.name + " ?", { okLabel: "Dissocier" });
        if (!ok) return;
        try {
          await def.onDisconnect();
          rerender();
        } catch (e) {
          alertDialog("Dissociation impossible : " + (e.message || e));
        }
      });
    }
    addItem(fav ? "Retirer des favoris" : "Ajouter aux favoris", "", async () => {
      setConnFav(def.id, !fav);
      rerender();
    });
    // « Supprimer » : réinitialise la configuration du connecteur (dissocie +
    // retire des favoris). Pour un connecteur intégré comme GitHub, la
    // définition subsiste : il réapparaît dans « Disponible ».
    addItem("Supprimer", "danger", async () => {
      const ok = await confirmDialog(
        "Supprimer le connecteur " + def.name + " ?" + (isConn ? " Il sera dissocié." : ""),
        { okLabel: "Supprimer", danger: true }
      );
      if (!ok) return;
      try {
        if (isConn) await def.onDisconnect();
        setConnFav(def.id, false);
        rerender();
      } catch (e) {
        alertDialog("Suppression impossible : " + (e.message || e));
      }
    });
    menuBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      document.querySelectorAll(".conn-menu").forEach((m) => {
        if (m !== menu) m.style.display = "none";
      });
      menu.style.display = menu.style.display === "none" ? "" : "none";
    });
    menuWrap.appendChild(menuBtn);
    menuWrap.appendChild(menu);
    card.appendChild(menuWrap);
    // Icône + infos (logo SVG si fourni, sinon emoji).
    const iconBox = connEl("div", "conn-icon");
    if (def.iconSvg) iconBox.innerHTML = def.iconSvg;
    else iconBox.textContent = def.icon || "🔌";
    card.appendChild(iconBox);
    const info = connEl("div", "conn-info");
    const nameRow = connEl("div", "conn-name", def.name);
    if (fav) nameRow.appendChild(connEl("span", "conn-star", "★"));
    info.appendChild(nameRow);
    info.appendChild(connEl("div", "conn-desc", def.desc));
    card.appendChild(info);
    card.appendChild(connEl("span", "conn-badge " + (isConn ? "on" : "off"), def.badgeOf(st)));
    if (!isConn) {
      const assoc = connEl("button", "conn-assoc", "Associer");
      assoc.type = "button";
      assoc.addEventListener("click", () => openConnectorModal(def, rerender));
      card.appendChild(assoc);
    }
    return card;
  };

  if (connected.length) {
    container.appendChild(connEl("div", "conn-section-label", "Connecté"));
    for (const item of connected) container.appendChild(makeCard(item));
  }
  if (available.length) {
    container.appendChild(connEl("div", "conn-section-label", "Disponible"));
    for (const item of available) container.appendChild(makeCard(item));
  }
  if (!connected.length && !available.length) {
    container.appendChild(connEl("div", "sb-tree-empty", "Aucun connecteur."));
  }
  // Fermeture des menus ⋮ au clic extérieur : un seul listener document stable
  // (pas de { once:true } recréé à chaque rendu, qui pouvait être consommé
  // par un clic antérieur et laisser ensuite les menus ouverts).
  if (!connOutsideCloseInstalled) {
    connOutsideCloseInstalled = true;
    document.addEventListener("click", (e) => {
      if (!e.target.closest || !e.target.closest(".conn-menu-wrap")) {
        document.querySelectorAll(".conn-menu").forEach((m) => (m.style.display = "none"));
      }
    });
  }
}

// Installé une seule fois (voir renderConnectorsInto).
let connOutsideCloseInstalled = false;

