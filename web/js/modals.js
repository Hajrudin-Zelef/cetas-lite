import { api, getPrefs, putPrefs } from "./api.js";
import { applyTheme, applyPalette, persistPrefs, THEME_PALETTES } from "./model-select.js";
import { confirmDialog, alertDialog } from "./dialogs.js";
export { confirmDialog, alertDialog };
import { getCategories, saveCategories } from "./sidebar.js";
import { saveRoles } from "./right-panel.js";
import { renderConnectorsInto, openNewProjectModal, ProjectsAPI } from "./projects.js";
import { initApiModelesPanel } from "./apimodeles.js";
import { loadFeaturesPanel, loadSearchPanel, loadAppearancePanel, registerPanelSaver, notifyConfigDirty, saveConfigPanel, saveAllConfigPanels } from "./config-panels.js";
import { loadVaultPanel } from "./vault.js";
import { loadSelectorPanel, loadAgentSelectorPanel } from "./model-selector.js";
import { loadDeepThinkPanel } from "./deepthink.js";
import { loadAgenticPanel } from "./agentic-panel.js";

const PROMPTS_KEY = "cetas-lite-prompts";

function loadPrompts() {
  try {
    return JSON.parse(localStorage.getItem(PROMPTS_KEY) || "[]");
  } catch (e) {
    return [];
  }
}
function savePrompts(p) {
  try {
    localStorage.setItem(PROMPTS_KEY, JSON.stringify(p));
  } catch (e) {}
}

// --- Dialogue personnalisé (alert/confirm) : voir ./dialogs.js ---

function openOverlay(id) {
  const el = document.getElementById(id);
  if (el) el.style.display = "";
}
function closeOverlay(id) {
  const el = document.getElementById(id);
  if (el) el.style.display = "none";
  if (id === "apikeys-modal-overlay") {
    // Les aperçus non sauvegardés (thème/palette) sont repliés.
    window.dispatchEvent(new CustomEvent("cetas:config-closed"));
  }
}

// --- Modale Configuration ---
function initConfigModal() {
  const overlay = document.getElementById("apikeys-modal-overlay");
  if (!overlay) return;

  overlay.querySelectorAll(".apikeys-tab").forEach((tab) => {
    tab.addEventListener("click", () => {
      overlay.querySelectorAll(".apikeys-tab").forEach((t) => t.classList.remove("active"));
      overlay.querySelectorAll(".apikeys-panel").forEach((p) => p.classList.remove("active"));
      tab.classList.add("active");
      const panel = document.getElementById("panel-" + tab.dataset.tab);
      if (panel) panel.classList.add("active");
      // Les connecteurs se chargent à l'ouverture de l'onglet.
      if (tab.dataset.tab === "connecteurs") {
        const body = document.getElementById("connectors-body");
        if (body) renderConnectorsInto(body);
      }
      // Onglet Remote : liste des serveurs SFTP.
      if (tab.dataset.tab === "remote") loadRemoteTab();
      // Onglet Compétences : éditeur de skills.
      if (tab.dataset.tab === "competences") loadSkillsTab();
      // Onglet Sélecteur de modèles : 1 modèle ou fallback par alias.
      if (tab.dataset.tab === "selector") loadSelectorPanel();
      // Onglet Sélecteur agent : idem, dédié aux familles agent.
      if (tab.dataset.tab === "selector-agent") loadAgentSelectorPanel();
      // Onglet Agentic : style d'affichage de la vue Agents.
      if (tab.dataset.tab === "agentic") loadAgenticPanel();
      // Onglet DeepThink Global : langue + modèle de traduction du raisonnement.
      if (tab.dataset.tab === "deepthink") loadDeepThinkPanel();
      // Onglet Fonctionnalités : six lignes de choix.
      if (tab.dataset.tab === "models") loadFeaturesPanel();
      // Onglet Recherche Web : moteurs + mode.
      if (tab.dataset.tab === "search") loadSearchPanel();
      // Onglet Apparence : palettes.
      if (tab.dataset.tab === "appearance") loadAppearancePanel();
      // Onglet Coffre : secrets chiffrés.
      if (tab.dataset.tab === "vault") loadVaultPanel();
    });
  });
  document.getElementById("apikeys-close-btn")?.addEventListener("click", () => closeOverlay("apikeys-modal-overlay"));
  // Sauvegarde : bouton par onglet + barre générale (modifs « stagées »).
  document.getElementById("config-save-all")?.addEventListener("click", async (e) => {
    const btn = e.currentTarget;
    btn.disabled = true;
    try {
      await saveAllConfigPanels();
    } catch (_) {}
    btn.disabled = false;
  });
  overlay.querySelectorAll(".config-panel-save-btn").forEach((b) => {
    b.addEventListener("click", async () => {
      b.disabled = true;
      try {
        await saveConfigPanel(b.dataset.tab);
      } catch (_) {}
      b.disabled = false;
    });
  });
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) closeOverlay("apikeys-modal-overlay");
  });

  // FAQ accordéon
  document.querySelectorAll("#faq-body .faq-question").forEach((q) => {
    q.addEventListener("click", () => {
      const item = q.closest(".faq-item");
      const ans = item ? item.querySelector(".faq-answer") : null;
      const open = item.classList.toggle("open");
      if (ans) ans.style.display = open ? "" : "none";
    });
    const ans = q.closest(".faq-item")?.querySelector(".faq-answer");
    if (ans) ans.style.display = "none";
  });


  async function loadFeatureDefaults() {
    const prefs = await getPrefs().catch(() => null);
    const set = (id, v) => {
      const el = document.getElementById(id);
      if (el) el.checked = !!v;
    };
    set("cfg-web-default", prefs && prefs.web_default);
    set("cfg-thinking-default", prefs && prefs.thinking_default);
    set("cfg-mcp-default", prefs && typeof prefs.mcp_default === "boolean" ? prefs.mcp_default : true);
    const eff = document.getElementById("cfg-thinking-effort");
    if (eff) eff.value = (prefs && prefs.thinking_effort) || "default";
    const wsm = document.getElementById("cfg-websearch-mode");
    if (wsm) wsm.value = (prefs && prefs.websearch_mode) || "auto";
    const theme = document.getElementById("cfg-theme");
    if (theme) theme.value = document.documentElement.dataset.theme || "clair";
    advPrefsSaved = readAdvancedPrefs(); // base de comparaison du staging
    const mcpStatus = document.getElementById("cfg-mcp-status");
    if (mcpStatus) {
      const mcp = await api("/api/mcp").catch(() => null);
      const servers = (mcp && mcp.servers) || [];
      mcpStatus.textContent = servers.length
        ? servers.length + " serveur(s) : " + servers.map((s) => s.name || s.id).join(", ")
        : "Aucun serveur configuré (mcp.json)";
    }
    await loadPluginsStatus();
  }

  async function loadPluginsStatus() {
    const el = document.getElementById("cfg-plugins-status");
    if (!el) return;
    const data = await api("/api/plugins").catch(() => null);
    const list = (data && data.plugins) || [];
    const errs = (data && data.errors) || [];
    const names = list.map((p) => p.name + (p.version ? " v" + p.version : ""));
    let txt = names.length ? names.length + " plugin(s) : " + names.join(", ") : "Aucun plugin (dossier plugins/)";
    if (errs.length) txt += " — " + errs.length + " erreur(s) : " + errs.join(" ; ");
    el.textContent = txt;
    el.title = txt;
  }

  // Réglages avancés de l'onglet Fonctionnalités : « stagés » comme le reste
  // (sauvegardés via le bouton Enregistrer de l'onglet ou la barre générale).
  // #cfg-websearch-mode est stagé par le panneau Recherche Web,
  // #cfg-theme par le panneau Apparence : ils ne sont plus ici.
  let advPrefsSaved = null;
  function readAdvancedPrefs() {
    const mcpEl = document.getElementById("cfg-mcp-default");
    return {
      web_default: !!document.getElementById("cfg-web-default")?.checked,
      thinking_default: !!document.getElementById("cfg-thinking-default")?.checked,
      thinking_effort: document.getElementById("cfg-thinking-effort")?.value || "default",
      mcp_default: mcpEl ? !!mcpEl.checked : true,
    };
  }

  function bindFeatureToggles() {
    const onAdvChange = () => notifyConfigDirty();
    ["cfg-web-default", "cfg-thinking-default", "cfg-mcp-default"].forEach((id) => {
      document.getElementById(id)?.addEventListener("change", onAdvChange);
    });
    document.getElementById("cfg-thinking-effort")?.addEventListener("change", onAdvChange);
    registerPanelSaver("models", {
      isDirty: () => {
        if (!advPrefsSaved) return false;
        const cur = readAdvancedPrefs();
        return Object.keys(cur).some((k) => cur[k] !== advPrefsSaved[k]);
      },
      save: async () => {
        const cur = readAdvancedPrefs();
        const body = {};
        for (const k of Object.keys(cur)) if (cur[k] !== advPrefsSaved[k]) body[k] = cur[k];
        if (Object.keys(body).length) {
          await putPrefs(body);
          try { persistPrefs(); } catch (e) {}
        }
        advPrefsSaved = cur;
      },
    });
    document.getElementById("cfg-plugins-reload")?.addEventListener("click", async (e) => {
      const btn = e.currentTarget;
      btn.disabled = true;
      try {
        await api("/api/plugins/reload", { method: "POST" });
      } catch (_) {}
      await loadPluginsStatus();
      btn.disabled = false;
    });
    document.getElementById("cfg-marex-save")?.addEventListener("click", async (e) => {
      const btn = e.currentTarget;
      const ta = document.getElementById("cfg-marex-content");
      const st = document.getElementById("cfg-marex-status");
      btn.disabled = true;
      try {
        await api("/api/marex", { method: "PUT", body: { content: ta ? ta.value : "" } });
        if (st) st.textContent = "Enregistré.";
      } catch (err) {
        if (st) st.textContent = "Échec : " + (err && err.message ? err.message : "erreur");
      }
      btn.disabled = false;
    });
  }
  bindFeatureToggles();

  async function loadMarex() {
    const ta = document.getElementById("cfg-marex-content");
    const st = document.getElementById("cfg-marex-status");
    if (!ta) return;
    try {
      const d = await api("/api/marex");
      ta.value = (d && d.content) || "";
      if (st) st.textContent = (d && d.path) ? "Fichier : " + d.path : "";
    } catch (_) {
      if (st) st.textContent = "Chargement impossible.";
    }
  }

  window.addEventListener("cetas:open-config", () => {
    openOverlay("apikeys-modal-overlay");
    initApiModelesPanel();
    loadFeatureDefaults();
    loadMarex();
  });
  window.addEventListener("cetas:open-config-faq", () => {
    openOverlay("apikeys-modal-overlay");
    overlay.querySelectorAll(".apikeys-tab").forEach((t) => t.classList.toggle("active", t.dataset.tab === "faq"));
    overlay.querySelectorAll(".apikeys-panel").forEach((p) => p.classList.toggle("active", p.id === "panel-faq"));
  });
  // Ouverture de la Configuration sur un onglet précis (depuis la vue Agents).
  window.addEventListener("cetas:open-config-tab", (e) => {
    const tab = (e && e.detail && e.detail.tab) || "apimodeles";
    openOverlay("apikeys-modal-overlay");
    const btn = overlay.querySelector('.apikeys-tab[data-tab="' + tab + '"]');
    if (btn) btn.click();
    else {
      overlay.querySelectorAll(".apikeys-tab").forEach((t) => t.classList.toggle("active", t.dataset.tab === tab));
      overlay.querySelectorAll(".apikeys-panel").forEach((p) => p.classList.toggle("active", p.id === "panel-" + tab));
    }
  });
  window.addEventListener("cetas:theme", (e) => {
    applyTheme(e.detail || "clair");
    putPrefs({ theme: document.documentElement.dataset.theme }).catch(() => {});
  });
}

// --- Onglet Remote (SFTP) de la Configuration ---
async function loadRemoteTab() {
  const body = document.getElementById("remote-body");
  if (!body) return;
  body.textContent = "Chargement…";
  let remotes = [];
  let activeId = "";
  try {
    const d = await ProjectsAPI.list();
    remotes = ((d && d.projects) || []).filter((p) => p.mode === "sftp");
    activeId = (d && d.active) || "";
  } catch (e) {
    body.textContent = "Erreur de chargement : " + (e.message || e);
    return;
  }
  body.innerHTML = "";
  const addBtn = document.createElement("button");
  addBtn.type = "button";
  addBtn.className = "mod-add-btn";
  addBtn.textContent = "+ Nouveau serveur";
  addBtn.addEventListener("click", () => {
    closeOverlay("apikeys-modal-overlay");
    // Après création : on rouvre la Configuration sur l'onglet Remote
    // (le clic sur l'onglet recharge la liste via le chargement paresseux).
    openNewProjectModal(() => {
      window.dispatchEvent(new CustomEvent("cetas:open-config-tab", { detail: { tab: "remote" } }));
    }, { tab: "sftp" });
  });
  body.appendChild(addBtn);
  if (!remotes.length) {
    const empty = document.createElement("div");
    empty.className = "mod-empty";
    const icon = document.createElement("span");
    icon.className = "mod-empty-icon";
    icon.textContent = "\u{1F5A5}\uFE0F";
    empty.appendChild(icon);
    empty.appendChild(document.createTextNode("Aucun serveur distant configuré."));
    body.appendChild(empty);
    return;
  }
  for (const r of remotes) {
    const card = document.createElement("div");
    card.className = "mod-card";
    const head = document.createElement("div");
    head.className = "mod-card-head";
    const title = document.createElement("div");
    title.className = "mod-card-title";
    title.textContent = r.name;
    head.appendChild(title);
    const isActive = r.id === activeId;
    const badge = document.createElement("span");
    badge.className = "mod-badge" + (isActive ? " on" : " off");
    badge.textContent = isActive ? "Actif" : "SFTP";
    head.appendChild(badge);
    card.appendChild(head);
    const desc = document.createElement("div");
    desc.className = "mod-card-desc";
    desc.textContent = r.user + "@" + r.host + ":" + r.remote_path;
    card.appendChild(desc);
    const actions = document.createElement("div");
    actions.className = "mod-card-actions";
    const useBtn = document.createElement("button");
    useBtn.type = "button";
    useBtn.className = "sess-btn";
    useBtn.textContent = isActive ? "Actif ✓" : "Utiliser";
    useBtn.title = "Définir comme projet actif";
    useBtn.disabled = isActive;
    useBtn.addEventListener("click", async () => {
      try {
        await ProjectsAPI.setActive(r.id);
        loadRemoteTab();
      } catch (e) {
        alertDialog("Sélection impossible : " + (e.message || e));
      }
    });
    const delBtn = document.createElement("button");
    delBtn.type = "button";
    delBtn.className = "sess-btn danger";
    delBtn.textContent = "Supprimer";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer le serveur « " + r.name + " » ?", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      try {
        await ProjectsAPI.remove(r.id);
        loadRemoteTab();
      } catch (e) {
        alertDialog("Suppression impossible : " + (e.message || e));
      }
    });
    actions.appendChild(useBtn);
    actions.appendChild(delBtn);
    card.appendChild(actions);
    body.appendChild(card);
  }
}

// --- Onglet Compétences de la Configuration ---
let cfgSkills = null; // cache local édité, sauvegardé via PUT /api/skills

async function loadSkillsTab() {
  const body = document.getElementById("skills-body");
  if (!body) return;
  if (cfgSkills === null) {
    body.textContent = "Chargement…";
    try {
      const d = await api("/api/skills");
      cfgSkills = (d && d.skills) || [];
    } catch (e) {
      body.textContent = "Erreur de chargement : " + (e.message || e);
      return;
    }
  }
  renderSkillsTab(body);
}

function renderSkillsTab(body) {
  body.innerHTML = "";
  const bar = document.createElement("div");
  bar.className = "mod-card-actions";
  bar.style.marginTop = "0";
  bar.style.marginBottom = "14px";
  const addBtn = document.createElement("button");
  addBtn.type = "button";
  addBtn.className = "mod-add-btn";
  addBtn.style.marginTop = "0";
  addBtn.textContent = "+ Nouvelle compétence";
  addBtn.addEventListener("click", () => {
    cfgSkills.unshift({ id: "", name: "", description: "", instructions: "", enabled: true, _open: true });
    renderSkillsTab(body);
  });
  const saveBtn = document.createElement("button");
  saveBtn.type = "button";
  saveBtn.className = "sess-btn";
  saveBtn.style.alignSelf = "center";
  saveBtn.textContent = "Enregistrer";
  saveBtn.addEventListener("click", async () => {
    const payload = cfgSkills.map((s) => ({
      id: s.id,
      name: s.name,
      description: s.description,
      instructions: s.instructions,
      enabled: !!s.enabled,
    }));
    try {
      const d = await api("/api/skills", { method: "PUT", body: { skills: payload } });
      cfgSkills = (d && d.skills) || [];
      renderSkillsTab(body);
      // Notifie la vue Agents : elle recharge le compteur du menu +.
      window.dispatchEvent(new CustomEvent("cetas:skills-changed"));
      const ok = document.createElement("span");
      ok.className = "rp-mcp-status";
      ok.textContent = "Enregistré ✓";
      bar.appendChild(ok);
      setTimeout(() => ok.remove(), 2000);
    } catch (e) {
      alertDialog("Enregistrement impossible : " + (e.message || e));
    }
  });
  bar.appendChild(addBtn);
  bar.appendChild(saveBtn);
  body.appendChild(bar);

  if (!cfgSkills.length) {
    const empty = document.createElement("div");
    empty.className = "mod-empty";
    const icon = document.createElement("span");
    icon.className = "mod-empty-icon";
    icon.textContent = "\U0001F9E0";
    empty.appendChild(icon);
    empty.appendChild(document.createTextNode("Aucune compétence. Créez-en une : elle sera injectée dans le prompt système de l'agent."));
    body.appendChild(empty);
    return;
  }
  for (const s of cfgSkills) {
    const card = document.createElement("div");
    card.className = "mod-card";

    const head = document.createElement("div");
    head.className = "mod-card-head";
    const nameEl = document.createElement("div");
    nameEl.className = "mod-card-title";
    nameEl.textContent = s.name || "(sans nom)";
    head.appendChild(nameEl);
    const badge = document.createElement("span");
    badge.className = "mod-badge" + (s.enabled ? " on" : " off");
    badge.textContent = s.enabled ? "Activée" : "Désactivée";
    head.appendChild(badge);
    card.appendChild(head);

    if (s.description && !s._open) {
      const d = document.createElement("div");
      d.className = "mod-card-desc";
      d.textContent = s.description;
      card.appendChild(d);
    }

    const form = document.createElement("div");
    form.className = "skill-form";
    form.style.display = s._open ? "" : "none";
    const nameIn = document.createElement("input");
    nameIn.className = "sp-modal-input";
    nameIn.placeholder = "Nom (80 caractères max)";
    nameIn.maxLength = 80;
    nameIn.value = s.name || "";
    nameIn.setAttribute("aria-label", "Nom de la compétence");
    nameIn.addEventListener("input", () => { s.name = nameIn.value; nameEl.textContent = s.name || "(sans nom)"; });
    const descIn = document.createElement("input");
    descIn.className = "sp-modal-input";
    descIn.placeholder = "Description courte (optionnel)";
    descIn.maxLength = 300;
    descIn.value = s.description || "";
    descIn.setAttribute("aria-label", "Description");
    descIn.addEventListener("input", () => { s.description = descIn.value; });
    const instrIn = document.createElement("textarea");
    instrIn.className = "sp-modal-input";
    instrIn.placeholder = "Instructions injectées dans le prompt système de l'agent…";
    instrIn.rows = 4;
    instrIn.value = s.instructions || "";
    instrIn.setAttribute("aria-label", "Instructions");
    instrIn.addEventListener("input", () => { s.instructions = instrIn.value; });
    form.appendChild(nameIn);
    form.appendChild(descIn);
    form.appendChild(instrIn);
    card.appendChild(form);

    const actions = document.createElement("div");
    actions.className = "mod-card-actions";
    const toggleWrap = document.createElement("label");
    toggleWrap.className = "plus-menu-toggle";
    toggleWrap.title = "Compétence activée";
    const cb = document.createElement("input");
    cb.type = "checkbox";
    cb.checked = !!s.enabled;
    cb.setAttribute("aria-label", "Activer la compétence");
    cb.addEventListener("change", () => {
      s.enabled = cb.checked;
      badge.className = "mod-badge" + (s.enabled ? " on" : " off");
      badge.textContent = s.enabled ? "Activée" : "Désactivée";
    });
    const slider = document.createElement("span");
    slider.className = "plus-menu-toggle-slider";
    toggleWrap.appendChild(cb);
    toggleWrap.appendChild(slider);
    const editBtn = document.createElement("button");
    editBtn.type = "button";
    editBtn.className = "sess-btn";
    editBtn.textContent = s._open ? "Réduire" : "Modifier";
    editBtn.addEventListener("click", () => {
      s._open = !s._open;
      renderSkillsTab(body);
    });
    const delBtn = document.createElement("button");
    delBtn.type = "button";
    delBtn.className = "sess-btn danger";
    delBtn.textContent = "Supprimer";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer la compétence « " + (s.name || "(sans nom)") + " » ?", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      cfgSkills = cfgSkills.filter((x) => x !== s);
      renderSkillsTab(body);
    });
    actions.appendChild(toggleWrap);
    actions.appendChild(editBtn);
    actions.appendChild(delBtn);
    card.appendChild(actions);
    body.appendChild(card);
  }
}

// --- Modale Sauvegarde ---
function initSaveModal() {
  const overlay = document.getElementById("save-modal-overlay");
  if (!overlay) return;
  document.getElementById("save-modal-close")?.addEventListener("click", () => closeOverlay("save-modal-overlay"));
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) closeOverlay("save-modal-overlay");
  });

  document.getElementById("save-modal-export-btn")?.addEventListener("click", async () => {
    try {
      const [convs, prefs] = await Promise.all([
        api("/api/sessions").catch(() => ({ sessions: [] })),
        getPrefs().catch(() => ({})),
      ]);
      const includePrefs = document.getElementById("save-modal-include-keys")?.checked;
      const data = {
        app: "cetas-lite",
        version: 1,
        exported_at: new Date().toISOString(),
        conversations: convs.sessions || [],
        prefs: includePrefs ? prefs : undefined,
      };
      const blob = new Blob([JSON.stringify(data, null, 2)], { type: "application/json" });
      const a = document.createElement("a");
      a.href = URL.createObjectURL(blob);
      a.download = "cetas-lite-sauvegarde.json";
      document.body.appendChild(a);
      a.click();
      a.remove();
      setTimeout(() => URL.revokeObjectURL(a.href), 1000);
    } catch (e) {
      alertDialog("Export impossible : " + e.message);
    }
  });

  const fileInput = document.getElementById("import-file-input");
  document.getElementById("save-modal-import-btn")?.addEventListener("click", () => fileInput?.click());
  if (fileInput) {
    fileInput.addEventListener("change", async () => {
      const f = fileInput.files && fileInput.files[0];
      fileInput.value = "";
      if (!f) return;
      try {
        const data = JSON.parse(await f.text());
        if (data.app !== "cetas-lite") throw new Error("fichier non reconnu");
        // Les conversations sont restaurées une par une via l'API d'archive n'existe pas :
        // on informe l'utilisateur de la limite.
        await alertDialog(
          "Import : " + (data.conversations ? data.conversations.length : 0) + " conversation(s) dans le fichier.\n\n" +
          "La restauration automatique des archives n'est pas encore supportée : conservez ce fichier comme sauvegarde."
        );
      } catch (e) {
        alertDialog("Import impossible : " + e.message);
      }
    });
  }

  window.addEventListener("cetas:open-save-modal", () => openOverlay("save-modal-overlay"));
}

// --- Modale Catégories ---
function initCatModal() {
  const overlay = document.getElementById("cat-modal-overlay");
  if (!overlay) return;
  const listView = document.getElementById("cat-manage-list-view");
  const editView = document.getElementById("cat-manage-edit-view");
  const listEl = document.getElementById("cat-manage-list");
  const nameInput = document.getElementById("cat-modal-nom");
  const emojiBtn = document.getElementById("cat-modal-icone-btn");
  const emojiPreview = document.getElementById("cat-modal-icone-preview");
  const titleEl = document.getElementById("cat-modal-title");
  let editingId = null;
  let emoji = "";

  function showList() {
    editView.style.display = "none";
    listView.style.display = "";
    render();
  }
  function render() {
    const cats = getCategories();
    listEl.innerHTML = "";
    if (!cats.length) {
      listEl.innerHTML = '<div class="manage-list-empty">Aucune catégorie</div>';
      return;
    }
    for (const c of cats) {
      const row = document.createElement("div");
      row.className = "cat-manage-row";
      const label = document.createElement("span");
      label.textContent = (c.emoji ? c.emoji + " " : "") + c.name;
      const edit = document.createElement("button");
      edit.type = "button";
      edit.className = "cat-manage-add-btn";
      edit.textContent = "Modifier";
      edit.addEventListener("click", () => openEdit(c));
      const del = document.createElement("button");
      del.type = "button";
      del.className = "cat-modal-delete-btn";
      del.textContent = "Supprimer";
      del.addEventListener("click", async () => {
        const ok = await confirmDialog('Supprimer la catégorie « ' + c.name + " » ?", { okLabel: "Supprimer", danger: true });
        if (!ok) return;
        saveCategories(getCategories().filter((x) => x.id !== c.id));
        render();
      });
      row.appendChild(label);
      row.appendChild(edit);
      row.appendChild(del);
      listEl.appendChild(row);
    }
  }
  function openEdit(c) {
    editingId = c ? c.id : null;
    emoji = c ? c.emoji || "" : "";
    if (titleEl) titleEl.textContent = c ? "Modifier la catégorie" : "Nouvelle catégorie";
    if (nameInput) nameInput.value = c ? c.name : "";
    if (emojiPreview) emojiPreview.textContent = emoji || "🏷️";
    listView.style.display = "none";
    editView.style.display = "";
  }

  document.getElementById("cat-manage-close")?.addEventListener("click", () => closeOverlay("cat-modal-overlay"));
  document.getElementById("cat-manage-add-btn")?.addEventListener("click", () => openEdit(null));
  document.getElementById("cat-modal-back")?.addEventListener("click", showList);
  document.getElementById("cat-modal-cancel")?.addEventListener("click", showList);
  if (emojiBtn) {
    emojiBtn.addEventListener("click", () => {
      const e = prompt("Emoji de la catégorie :", emoji || "🏷️");
      if (e !== null) {
        emoji = e.trim().slice(0, 4);
        if (emojiPreview) emojiPreview.textContent = emoji || "🏷️";
      }
    });
  }
  document.getElementById("cat-modal-save")?.addEventListener("click", () => {
    const name = (nameInput?.value || "").trim();
    if (!name) return;
    const cats = getCategories();
    if (editingId) {
      const c = cats.find((x) => x.id === editingId);
      if (c) {
        c.name = name;
        c.emoji = emoji;
      }
    } else {
      cats.push({ id: "cat-" + Date.now().toString(36), name, emoji });
    }
    saveCategories(cats);
    showList();
  });
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) closeOverlay("cat-modal-overlay");
  });

  window.addEventListener("cetas:open-cat-modal", () => {
    openOverlay("cat-modal-overlay");
    showList();
  });
}

// --- Rôles & Prompts ---
function initRolesPrompts() {
  // Liste des rôles
  const rolesOverlay = document.getElementById("roles-manage-overlay");
  const rolesList = document.getElementById("roles-manage-list");
  const rolesEmpty = document.getElementById("roles-manage-empty");
  function renderRoles() {
    let roles = [];
    try {
      roles = JSON.parse(localStorage.getItem("cetas-lite-roles") || "[]");
    } catch (e) {}
    rolesList.innerHTML = "";
    rolesEmpty.style.display = roles.length ? "none" : "";
    roles.forEach((r, i) => {
      const row = document.createElement("div");
      row.className = "manage-list-row";
      const label = document.createElement("span");
      label.textContent = r.name;
      const edit = document.createElement("button");
      edit.type = "button";
      edit.className = "cat-manage-add-btn";
      edit.textContent = "Modifier";
      edit.addEventListener("click", () => openRoleModal(i));
      const del = document.createElement("button");
      del.type = "button";
      del.className = "cat-modal-delete-btn";
      del.textContent = "Supprimer";
      del.addEventListener("click", async () => {
        const ok = await confirmDialog('Supprimer le rôle « ' + r.name + " » ?", { okLabel: "Supprimer", danger: true });
        if (!ok) return;
        const all = JSON.parse(localStorage.getItem("cetas-lite-roles") || "[]");
        all.splice(i, 1);
        saveRoles(all);
        renderRoles();
      });
      row.appendChild(label);
      row.appendChild(edit);
      row.appendChild(del);
      rolesList.appendChild(row);
    });
  }
  document.getElementById("roles-manage-close")?.addEventListener("click", () => closeOverlay("roles-manage-overlay"));
  document.getElementById("roles-manage-add")?.addEventListener("click", () => openRoleModal(null));
  rolesOverlay?.addEventListener("click", (e) => {
    if (e.target === rolesOverlay) closeOverlay("roles-manage-overlay");
  });

  // Édition d'un rôle
  let roleIdx = null;
  function openRoleModal(i) {
    roleIdx = i;
    const roles = JSON.parse(localStorage.getItem("cetas-lite-roles") || "[]");
    const r = i != null ? roles[i] : null;
    document.getElementById("sp-modal-title").textContent = r ? "Modifier le rôle" : "Nouveau rôle";
    document.getElementById("sp-modal-nom").value = r ? r.name : "";
    document.getElementById("sp-modal-contenu").value = r ? r.content : "";
    document.getElementById("sp-modal-delete").style.display = r ? "" : "none";
    openOverlay("sp-modal-overlay");
  }
  document.getElementById("sp-modal-cancel")?.addEventListener("click", () => closeOverlay("sp-modal-overlay"));
  document.getElementById("sp-modal-save")?.addEventListener("click", () => {
    const name = document.getElementById("sp-modal-nom").value.trim();
    const content = document.getElementById("sp-modal-contenu").value.trim();
    if (!name || !content) return;
    const roles = JSON.parse(localStorage.getItem("cetas-lite-roles") || "[]");
    if (roleIdx != null) roles[roleIdx] = { name, content };
    else roles.push({ name, content });
    saveRoles(roles);
    closeOverlay("sp-modal-overlay");
    renderRoles();
  });
  document.getElementById("sp-modal-delete")?.addEventListener("click", async () => {
    const ok = await confirmDialog("Supprimer ce rôle ?", { okLabel: "Supprimer", danger: true });
    if (!ok) return;
    const roles = JSON.parse(localStorage.getItem("cetas-lite-roles") || "[]");
    if (roleIdx != null) roles.splice(roleIdx, 1);
    saveRoles(roles);
    closeOverlay("sp-modal-overlay");
    renderRoles();
  });
  document.getElementById("sp-modal-optimize")?.addEventListener("click", () => {
    alertDialog("L'amélioration par IA arrivera prochainement.");
  });

  // Liste des prompts
  const prOverlay = document.getElementById("prompts-manage-overlay");
  const prList = document.getElementById("prompts-manage-list");
  const prEmpty = document.getElementById("prompts-manage-empty");
  function renderPrompts() {
    const prompts = loadPrompts();
    prList.innerHTML = "";
    prEmpty.style.display = prompts.length ? "none" : "";
    prompts.forEach((p, i) => {
      const row = document.createElement("div");
      row.className = "manage-list-row";
      const label = document.createElement("span");
      label.textContent = p.name;
      const insert = document.createElement("button");
      insert.type = "button";
      insert.className = "cat-manage-add-btn";
      insert.textContent = "Insérer";
      insert.addEventListener("click", () => {
        const input = document.getElementById("prompt-input");
        if (input) {
          input.value = p.content + (input.value ? "\n\n" + input.value : "");
          input.dispatchEvent(new Event("input"));
          input.focus();
        }
        closeOverlay("prompts-manage-overlay");
      });
      const edit = document.createElement("button");
      edit.type = "button";
      edit.className = "cat-manage-add-btn";
      edit.textContent = "Modifier";
      edit.addEventListener("click", () => openPromptModal(i));
      const del = document.createElement("button");
      del.type = "button";
      del.className = "cat-modal-delete-btn";
      del.textContent = "Supprimer";
      del.addEventListener("click", async () => {
        const ok = await confirmDialog('Supprimer le prompt « ' + p.name + " » ?", { okLabel: "Supprimer", danger: true });
        if (!ok) return;
        const all = loadPrompts();
        all.splice(i, 1);
        savePrompts(all);
        renderPrompts();
      });
      row.appendChild(label);
      row.appendChild(insert);
      row.appendChild(edit);
      row.appendChild(del);
      prList.appendChild(row);
    });
  }
  document.getElementById("prompts-manage-close")?.addEventListener("click", () => closeOverlay("prompts-manage-overlay"));
  document.getElementById("prompts-manage-add")?.addEventListener("click", () => openPromptModal(null));
  prOverlay?.addEventListener("click", (e) => {
    if (e.target === prOverlay) closeOverlay("prompts-manage-overlay");
  });

  let promptIdx = null;
  function openPromptModal(i) {
    promptIdx = i;
    const prompts = loadPrompts();
    const p = i != null ? prompts[i] : null;
    document.getElementById("pr-modal-title").textContent = p ? "Modifier le prompt" : "Nouveau Prompt";
    document.getElementById("pr-modal-nom").value = p ? p.name : "";
    document.getElementById("pr-modal-contenu").value = p ? p.content : "";
    document.getElementById("pr-modal-delete").style.display = p ? "" : "none";
    openOverlay("pr-modal-overlay");
  }
  document.getElementById("pr-modal-cancel")?.addEventListener("click", () => closeOverlay("pr-modal-overlay"));
  document.getElementById("pr-modal-save")?.addEventListener("click", () => {
    const name = document.getElementById("pr-modal-nom").value.trim();
    const content = document.getElementById("pr-modal-contenu").value.trim();
    if (!name || !content) return;
    const prompts = loadPrompts();
    if (promptIdx != null) prompts[promptIdx] = { name, content };
    else prompts.push({ name, content });
    savePrompts(prompts);
    closeOverlay("pr-modal-overlay");
    renderPrompts();
  });
  document.getElementById("pr-modal-delete")?.addEventListener("click", async () => {
    const ok = await confirmDialog("Supprimer ce prompt ?", { okLabel: "Supprimer", danger: true });
    if (!ok) return;
    const prompts = loadPrompts();
    if (promptIdx != null) prompts.splice(promptIdx, 1);
    savePrompts(prompts);
    closeOverlay("pr-modal-overlay");
    renderPrompts();
  });
  document.getElementById("pr-modal-enhance")?.addEventListener("click", () => {
    alertDialog("L'amélioration par IA arrivera prochainement.");
  });

  // Bouton d'insertion rapide dans la zone de saisie
  const toolbarInsert = document.getElementById("toolbar-insert-btn");
  if (toolbarInsert) {
    toolbarInsert.style.display = "";
    toolbarInsert.addEventListener("click", () => {
      openOverlay("prompts-manage-overlay");
      renderPrompts();
    });
  }

  window.addEventListener("cetas:open-roles-modal", () => {
    openOverlay("roles-manage-overlay");
    renderRoles();
  });
  window.addEventListener("cetas:open-prompts-modal", () => {
    openOverlay("prompts-manage-overlay");
    renderPrompts();
  });
}

export function initModals() {
  initConfigModal();
  initSaveModal();
  initCatModal();
  initRolesPrompts();
}
