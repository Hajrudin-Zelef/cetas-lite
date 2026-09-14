import { api, getPrefs, putPrefs } from "./api.js";
import { applyTheme, persistPrefs } from "./model-select.js";
import { getCategories, saveCategories } from "./sidebar.js";
import { saveRoles } from "./right-panel.js";

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

// --- Dialogue personnalisé (alert/confirm) ---
export function confirmDialog(message, opts = {}) {
  return new Promise((resolve) => {
    const overlay = document.getElementById("custom-dialog-overlay");
    const icon = document.getElementById("custom-dialog-icon");
    const msg = document.getElementById("custom-dialog-message");
    const okBtn = document.getElementById("custom-dialog-ok");
    const cancelBtn = document.getElementById("custom-dialog-cancel");
    if (!overlay) {
      resolve(window.confirm(message));
      return;
    }
    msg.textContent = message;
    icon.textContent = opts.danger ? "⚠️" : "❓";
    okBtn.textContent = opts.okLabel || "OK";
    cancelBtn.style.display = opts.hideCancel ? "none" : "";
    const done = (v) => {
      overlay.style.display = "none";
      okBtn.removeEventListener("click", onOk);
      cancelBtn.removeEventListener("click", onCancel);
      overlay.removeEventListener("click", onBg);
      resolve(v);
    };
    const onOk = () => done(true);
    const onCancel = () => done(false);
    const onBg = (e) => {
      if (e.target === overlay) done(false);
    };
    okBtn.addEventListener("click", onOk);
    cancelBtn.addEventListener("click", onCancel);
    overlay.addEventListener("click", onBg);
    overlay.style.display = "";
  });
}

export function alertDialog(message) {
  return confirmDialog(message, { okLabel: "Compris", hideCancel: true });
}

function openOverlay(id) {
  const el = document.getElementById(id);
  if (el) el.style.display = "";
}
function closeOverlay(id) {
  const el = document.getElementById(id);
  if (el) el.style.display = "none";
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
    });
  });
  document.getElementById("apikeys-close-btn")?.addEventListener("click", () => closeOverlay("apikeys-modal-overlay"));
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

  async function loadProviders() {
    const list = document.getElementById("providers-list");
    if (!list) return;
    list.innerHTML = '<div class="conv-list-loading">Chargement…</div>';
    let providers;
    try {
      const data = await api("/api/providers");
      providers = (data && data.providers) || [];
    } catch (e) {
      list.innerHTML = '<div class="models-error">Impossible de charger les fournisseurs : ' + e.message + "</div>";
      return;
    }
    list.innerHTML = "";
    for (const p of providers) {
      const row = document.createElement("div");
      row.className = "provider-row" + (p.configured ? " configured" : "");
      const head = document.createElement("div");
      head.className = "provider-row-head";
      const name = document.createElement("span");
      name.className = "provider-row-name";
      name.textContent = p.label || p.id;
      const badge = document.createElement("span");
      badge.className = "provider-badge " + (p.configured ? "on" : "off");
      badge.textContent = p.configured ? "Configuré" : "Non configuré";
      head.appendChild(name);
      head.appendChild(badge);
      const body = document.createElement("div");
      body.className = "provider-row-body";
      const input = document.createElement("input");
      input.type = "password";
      input.className = "sp-modal-input";
      input.placeholder = p.configured ? "•••••••• (laisser vide pour conserver)" : "Coller la clé API…";
      input.autocomplete = "off";
      const actions = document.createElement("div");
      actions.className = "provider-row-actions";
      const save = document.createElement("button");
      save.type = "button";
      save.className = "models-save-btn";
      save.textContent = "Enregistrer";
      save.addEventListener("click", async () => {
        const key = input.value.trim();
        if (!key) return;
        save.disabled = true;
        try {
          await api("/api/providers/" + encodeURIComponent(p.id), { method: "PUT", body: { key } });
          input.value = "";
          await loadProviders();
        } catch (e) {
          alertDialog("Échec de l'enregistrement : " + e.message);
        } finally {
          save.disabled = false;
        }
      });
      actions.appendChild(save);
      if (p.configured) {
        const del = document.createElement("button");
        del.type = "button";
        del.className = "models-cancel-btn";
        del.textContent = "Supprimer";
        del.addEventListener("click", async () => {
          const ok = await confirmDialog("Supprimer la clé du fournisseur « " + (p.label || p.id) + " » ?", { okLabel: "Supprimer", danger: true });
          if (!ok) return;
          try {
            await api("/api/providers/" + encodeURIComponent(p.id), { method: "DELETE" });
            await loadProviders();
          } catch (e) {
            alertDialog("Échec de la suppression : " + e.message);
          }
        });
        actions.appendChild(del);
      }
      body.appendChild(input);
      body.appendChild(actions);
      row.appendChild(head);
      row.appendChild(body);
      list.appendChild(row);
    }
  }

  async function loadFamilies() {
    const list = document.getElementById("families-list");
    if (!list) return;
    try {
      const data = await api("/api/aliases");
      list.innerHTML = "";
      for (const f of data.families || []) {
        const row = document.createElement("div");
        row.className = "family-row";
        const name = document.createElement("span");
        name.className = "family-row-name";
        name.textContent = f.label;
        const modes = document.createElement("span");
        modes.className = "family-row-modes";
        modes.textContent = (f.modes || []).map((m) => m.label).join(" · ");
        row.appendChild(name);
        row.appendChild(modes);
        list.appendChild(row);
      }
    } catch (e) {
      list.innerHTML = '<div class="models-error">Chargement impossible.</div>';
    }
  }

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
    if (theme) theme.value = document.documentElement.dataset.theme || "ocean";
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

  function bindFeatureToggles() {
    const save = () => {
      const body = {
        web_default: document.getElementById("cfg-web-default")?.checked || false,
        websearch_mode: document.getElementById("cfg-websearch-mode")?.value || "auto",
        thinking_default: document.getElementById("cfg-thinking-default")?.checked || false,
        thinking_effort: document.getElementById("cfg-thinking-effort")?.value || "default",
        mcp_default: document.getElementById("cfg-mcp-default")?.checked || false,
      };
      putPrefs(body).then(() => persistPrefs().catch(() => {})).catch(() => {});
    };
    ["cfg-web-default", "cfg-thinking-default", "cfg-mcp-default"].forEach((id) => {
      document.getElementById(id)?.addEventListener("change", save);
    });
    document.getElementById("cfg-thinking-effort")?.addEventListener("change", save);
    document.getElementById("cfg-websearch-mode")?.addEventListener("change", save);
    document.getElementById("cfg-theme")?.addEventListener("change", (e) => {
      applyTheme(e.target.value);
      putPrefs({ theme: e.target.value }).catch(() => {});
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
    loadProviders();
    loadFamilies();
    loadFeatureDefaults();
    loadMarex();
  });
  window.addEventListener("cetas:open-config-faq", () => {
    openOverlay("apikeys-modal-overlay");
    overlay.querySelectorAll(".apikeys-tab").forEach((t) => t.classList.toggle("active", t.dataset.tab === "faq"));
    overlay.querySelectorAll(".apikeys-panel").forEach((p) => p.classList.toggle("active", p.id === "panel-faq"));
  });
  window.addEventListener("cetas:theme", (e) => {
    applyTheme(e.detail || "ocean");
    putPrefs({ theme: document.documentElement.dataset.theme }).catch(() => {});
  });
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
        api("/api/conversations").catch(() => ({ archives: [] })),
        getPrefs().catch(() => ({})),
      ]);
      const includePrefs = document.getElementById("save-modal-include-keys")?.checked;
      const data = {
        app: "cetas-lite",
        version: 1,
        exported_at: new Date().toISOString(),
        conversations: convs.archives || [],
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
