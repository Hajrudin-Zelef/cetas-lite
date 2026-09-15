// Panneaux de la Configuration : Fonctionnalités, Recherche Web, Sessions,
// Apparence. Module dédié (dépendances légères) pour rester testable sous
// jsdom sans charger tout modals.js.
import { api, getPrefs, putPrefs } from "./api.js";
import { applyPalette, THEME_PALETTES } from "./model-select.js";
import { confirmDialog } from "./dialogs.js";

function closeConfigModal() {
  const el = document.getElementById("apikeys-modal-overlay");
  if (el) el.style.display = "none";
}

// --- Panneaux de la Configuration (Fonctionnalités / Recherche / Sessions / Apparence) ---

// Petit helper DOM.
function cfgEl(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

// Six lignes de Fonctionnalités (uniquement dans les Paramètres CETAS,
// jamais exposées dans la vue Agents).
const FEATURE_DEFS = [
  { key: "tts", label: "Synthèse vocale", kind: "system", fallback: "system" },
  { key: "transcription", label: "Transcription", kind: "system", fallback: "system" },
  { key: "prompt_enhance", label: "Amélioration de prompts/rôles", kind: "provider", fallback: "none" },
  { key: "summarizer", label: "Résumé IA", kind: "provider", fallback: "none" },
  { key: "title_gen", label: "Génération du titre", kind: "provider-conv", fallback: "conversation" },
  { key: "error_analysis", label: "Analyse des erreurs", kind: "provider", fallback: "none" },
];

let featuresLoaded = false;
async function loadFeaturesPanel() {
  const wrap = document.getElementById("features-rows");
  if (!wrap || featuresLoaded) return;
  featuresLoaded = true;
  wrap.innerHTML = "";
  wrap.appendChild(cfgEl("div", "sb-tree-empty", "Chargement…"));
  let prefs = null, provData = null;
  try { prefs = await getPrefs(); } catch (e) {}
  try { provData = await api("/api/providers"); } catch (e) {}
  const providers = (provData && provData.providers) || [];
  wrap.innerHTML = "";
  for (const def of FEATURE_DEFS) {
    const row = cfgEl("div", "feat-row");
    row.appendChild(cfgEl("span", "feat-name", def.label));
    const sel = cfgEl("select", "feat-select");
    sel.dataset.feature = def.key;
    const opts = [];
    if (def.kind === "system") {
      opts.push(["system", "Système (navigateur)"], ["none", "Aucun"]);
    } else {
      if (def.kind === "provider-conv") opts.push(["conversation", "Modèle de la conversation"]);
      opts.push(["none", "Aucun"]);
      for (const p of providers) {
        opts.push([p.id, p.label + (p.configured ? "" : " — non configuré")]);
      }
    }
    for (const [v, l] of opts) {
      const o = document.createElement("option");
      o.value = v; o.textContent = l;
      sel.appendChild(o);
    }
    const cur = (prefs && prefs[def.key]) || def.fallback;
    sel.value = Array.from(sel.options).some((o) => o.value === cur) ? cur : def.fallback;
    sel.addEventListener("change", async () => {
      try {
        await putPrefs({ [def.key]: sel.value });
        window.dispatchEvent(new CustomEvent("cetas:features-changed"));
      } catch (e) {}
    });
    row.appendChild(sel);
    wrap.appendChild(row);
  }
}

let searchLoaded = false;
let searchState = null; // { mode, providers: [...] } — réponse GET /api/search/settings

async function refreshSearchSettings() {
  try {
    searchState = await api("/api/search/settings");
  } catch (e) {
    searchState = null;
  }
  return searchState;
}

async function loadSearchPanel() {
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl || searchLoaded) return;
  searchLoaded = true;
  const modeSel = document.getElementById("search-mode");
  contentEl.innerHTML = "";
  contentEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement…"));
  const data = await refreshSearchSettings();
  if (!data) {
    contentEl.innerHTML = "";
    contentEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement impossible."));
    return;
  }
  if (modeSel) {
    modeSel.value = data.mode === "priority" ? "priority" : "race";
    modeSel.onchange = async () => {
      try {
        await api("/api/search/settings", { method: "PUT", body: { mode: modeSel.value } });
        if (searchState) searchState.mode = modeSel.value;
      } catch (e) {}
    };
  }
  renderSearchTabs();
}

function renderSearchTabs() {
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl || !searchState) return;
  const providers = searchState.providers || [];
  tabsEl.innerHTML = "";
  contentEl.innerHTML = "";
  providers.forEach((p, i) => {
    const sec = cfgEl("div", "provider-section" + (i === 0 ? " active" : ""));
    sec.dataset.provider = p.id;
    renderSearchProviderPane(p, sec);
    contentEl.appendChild(sec);
    const b = cfgEl("button", "provider-tab" + (i === 0 ? " active" : ""), p.label);
    b.type = "button";
    b.dataset.provider = p.id;
    b.addEventListener("click", () => {
      tabsEl.querySelectorAll(".provider-tab").forEach((t) => t.classList.toggle("active", t === b));
      contentEl.querySelectorAll(".provider-section").forEach((s) => s.classList.toggle("active", s.dataset.provider === p.id));
    });
    tabsEl.appendChild(b);
  });
}

function renderSearchProviderPane(p, sec) {
  sec.innerHTML = "";
  const wrap = cfgEl("div", "search-provider-pane");
  // Ligne d'activation
  const trow = cfgEl("div", "audio-setting-row search-toggle-row");
  trow.appendChild(cfgEl("span", "audio-setting-label", "Activer " + p.label));
  const tlabel = cfgEl("label", "plus-menu-toggle");
  const cb = document.createElement("input");
  cb.type = "checkbox";
  cb.checked = !!p.enabled;
  cb.setAttribute("aria-label", "Activer " + p.label);
  const slider = cfgEl("span", "plus-menu-toggle-slider");
  tlabel.appendChild(cb);
  tlabel.appendChild(slider);
  trow.appendChild(tlabel);
  wrap.appendChild(trow);
  const status = cfgEl("p", "search-status", searchStatusText(p));
  // Zone clé
  if (p.keyless) {
    wrap.appendChild(cfgEl("p", "apikey-intro", "Aucune clé requise : ce moteur est utilisé en repli, sans authentification."));
  } else {
    const labelRow = cfgEl("div", "apikey-label-row");
    const lab = cfgEl("label", "sp-modal-label", p.key_label || ("Clé API " + p.label));
    lab.htmlFor = "search-key-" + p.id;
    labelRow.appendChild(lab);
    if (p.key_url) {
      const a = cfgEl("a", "apikey-get-link", "Obtenir une clé");
      a.href = p.key_url;
      a.target = "_blank";
      a.rel = "noopener noreferrer";
      labelRow.appendChild(a);
    }
    wrap.appendChild(labelRow);
    const field = cfgEl("div", "apikey-field");
    const iwrap = cfgEl("div", "apikey-input-wrap");
    const input = document.createElement("input");
    input.type = "password";
    input.id = "search-key-" + p.id;
    input.className = "sp-modal-input apikey-input";
    input.autocomplete = "off";
    input.spellcheck = false;
    input.placeholder = p.configured ? "•••••••• (configurée)" : "Coller la clé API…";
    const eye = cfgEl("button", "search-eye-btn", "👁");
    eye.type = "button";
    eye.title = "Afficher la clé";
    eye.setAttribute("aria-label", "Afficher la clé");
    eye.addEventListener("click", () => {
      const show = input.type === "password";
      input.type = show ? "text" : "password";
      eye.classList.toggle("on", show);
    });
    iwrap.appendChild(input);
    iwrap.appendChild(eye);
    field.appendChild(iwrap);
    wrap.appendChild(field);
    const row = cfgEl("div", "search-key-actions");
    const saveBtn = cfgEl("button", "models-save-btn", p.configured ? "Mettre à jour" : "Enregistrer");
    saveBtn.type = "button";
    saveBtn.addEventListener("click", async () => {
      const key = input.value.trim();
      if (!key && !p.configured) {
        status.textContent = "Collez d'abord la clé.";
        status.classList.add("err");
        return;
      }
      saveBtn.disabled = true;
      try {
        await api("/api/search/settings", {
          method: "PUT",
          body: { providers: { [p.id]: { enabled: cb.checked, key: key || undefined } } },
        });
        await refreshSearchSettings();
        renderSearchTabs();
        selectSearchTab(p.id);
      } catch (e) {
        status.textContent = "Échec : " + (e.message || e);
        status.classList.add("err");
        saveBtn.disabled = false;
      }
    });
    row.appendChild(saveBtn);
    const delBtn = cfgEl("button", "apikey-delete-link", "Supprimer");
    delBtn.type = "button";
    delBtn.title = "Supprimer la clé";
    delBtn.style.display = p.configured ? "" : "none";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer la clé " + p.label + " ?", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      try {
        await api("/api/search/settings", {
          method: "PUT",
          body: { providers: { [p.id]: { enabled: cb.checked, key: "" } } },
        });
        await refreshSearchSettings();
        renderSearchTabs();
        selectSearchTab(p.id);
      } catch (e) {
        status.textContent = "Échec : " + (e.message || e);
        status.classList.add("err");
      }
    });
    row.appendChild(delBtn);
    wrap.appendChild(row);
  }
  wrap.appendChild(status);
  sec.appendChild(wrap);
  cb.addEventListener("change", async () => {
    try {
      await api("/api/search/settings", {
        method: "PUT",
        body: { providers: { [p.id]: { enabled: cb.checked } } },
      });
      await refreshSearchSettings();
      const cur = currentSearchProvider(p.id);
      status.textContent = searchStatusText(cur);
      status.classList.remove("err");
      if (cur) p.enabled = cur.enabled;
    } catch (e) {
      cb.checked = !cb.checked;
      status.textContent = "Échec : " + (e.message || e);
      status.classList.add("err");
    }
  });
}

function currentSearchProvider(id) {
  const list = (searchState && searchState.providers) || [];
  return list.find((x) => x.id === id) || null;
}

function selectSearchTab(id) {
  const tabsEl = document.getElementById("search-providers-tabs");
  const contentEl = document.getElementById("search-provider-content");
  if (!tabsEl || !contentEl) return;
  tabsEl.querySelectorAll(".provider-tab").forEach((t) => t.classList.toggle("active", t.dataset.provider === id));
  contentEl.querySelectorAll(".provider-section").forEach((s) => s.classList.toggle("active", s.dataset.provider === id));
}

function searchStatusText(p) {
  if (!p) return "";
  if (!p.enabled) return "Désactivé — ce moteur ne sera pas utilisé.";
  if (p.keyless) return "Activé — utilisé sans clé.";
  return p.configured ? "Activé — clé configurée." : "Activé — aucune clé : ce moteur sera ignoré.";
}

// --- Panneau Sessions : historique unifié CETAS + Agents ---
let sessionsLoaded = false;
let sessionsData = null; // { chat: [...], agents: [...] }

function fmtSessionDate(ts) {
  if (!ts) return "";
  try {
    const ms = ts < 1e12 ? ts * 1000 : ts;
    return new Date(ms).toLocaleDateString(undefined, { day: "numeric", month: "short" }) +
      " " + new Date(ms).toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
  } catch (e) {
    return "";
  }
}

async function loadSessionsPanel() {
  const chatEl = document.getElementById("sessions-chat");
  const agentsEl = document.getElementById("sessions-agents");
  const filterEl = document.getElementById("sessions-filter");
  if (!chatEl || !agentsEl) return;
  if (!sessionsLoaded) {
    sessionsLoaded = true;
    if (filterEl) filterEl.addEventListener("input", () => renderSessions(filterEl.value.trim().toLowerCase()));
    await refreshSessionsPanel();
  }
}

async function refreshSessionsPanel() {
  const chatEl = document.getElementById("sessions-chat");
  const agentsEl = document.getElementById("sessions-agents");
  if (!chatEl || !agentsEl) return;
  chatEl.innerHTML = "";
  chatEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement…"));
  agentsEl.innerHTML = "";
  try {
    sessionsData = await api("/api/sessions");
  } catch (e) {
    chatEl.innerHTML = "";
    chatEl.appendChild(cfgEl("div", "sb-tree-empty", "Chargement impossible."));
    return;
  }
  const filterEl = document.getElementById("sessions-filter");
  renderSessions(filterEl ? filterEl.value.trim().toLowerCase() : "");
}

function renderSessions(filter) {
  const chatEl = document.getElementById("sessions-chat");
  const agentsEl = document.getElementById("sessions-agents");
  if (!chatEl || !agentsEl) return;
  chatEl.innerHTML = "";
  agentsEl.innerHTML = "";
  const match = (t) => !filter || (t || "").toLowerCase().includes(filter);

  const chat = ((sessionsData && sessionsData.chat) || []).filter((s) => match(s.title));
  if (!chat.length) chatEl.appendChild(cfgEl("div", "sb-tree-empty", filter ? "Aucun résultat." : "Aucune conversation."));
  for (const s of chat) {
    const row = cfgEl("div", "sess-row");
    const info = cfgEl("div", "sess-info");
    info.appendChild(cfgEl("div", "sess-title", s.title || "Conversation sans titre"));
    const meta = [s.messages != null ? s.messages + " messages" : "", fmtSessionDate(s.updated)].filter(Boolean).join(" · ");
    info.appendChild(cfgEl("div", "sess-meta", meta));
    row.appendChild(info);
    const acts = cfgEl("div", "sess-actions");
    const openBtn = cfgEl("button", "sess-btn", "Ouvrir");
    openBtn.type = "button";
    openBtn.addEventListener("click", async () => {
      try {
        // Comme dans la sidebar : si la conversation courante n'est pas vide,
        // elle sera archivée — on demande confirmation.
        const state = await api("/api/chat/state").catch(() => ({}));
        if ((state.turns || 0) > 0) {
          const ok = await confirmDialog("Restaurer cette conversation ? La conversation actuelle sera archivée.", { okLabel: "Restaurer" });
          if (!ok) return;
        }
        await api("/api/conversations/restore", { method: "POST", body: { id: s.id } });
        window.dispatchEvent(new CustomEvent("cetas:chat-reset"));
        window.dispatchEvent(new CustomEvent("cetas:chat-changed"));
        closeConfigModal();
      } catch (e) {}
    });
    const delBtn = cfgEl("button", "sess-btn danger", "Supprimer");
    delBtn.type = "button";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer définitivement cette conversation ?", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      try {
        await api("/api/conversations/" + encodeURIComponent(s.id), { method: "DELETE" });
        window.dispatchEvent(new CustomEvent("cetas:chat-changed"));
        await refreshSessionsPanel();
      } catch (e) {}
    });
    acts.appendChild(openBtn);
    acts.appendChild(delBtn);
    row.appendChild(acts);
    chatEl.appendChild(row);
  }

  const agents = ((sessionsData && sessionsData.agents) || []).filter((s) => match(s.title));
  if (!agents.length) agentsEl.appendChild(cfgEl("div", "sb-tree-empty", filter ? "Aucun résultat." : "Aucun agent."));
  for (const s of agents) {
    const row = cfgEl("div", "sess-row");
    const info = cfgEl("div", "sess-info");
    info.appendChild(cfgEl("div", "sess-title", s.title || "Agent " + s.id));
    const parts = [];
    if (s.status) parts.push(agentStatusLabel(s.status));
    if (s.family) parts.push(s.family);
    if (s.mode) parts.push(s.mode);
    const d = fmtSessionDate(s.updated);
    if (d) parts.push(d);
    info.appendChild(cfgEl("div", "sess-meta", parts.join(" · ")));
    row.appendChild(info);
    const acts = cfgEl("div", "sess-actions");
    const openBtn = cfgEl("button", "sess-btn", "Ouvrir");
    openBtn.type = "button";
    openBtn.addEventListener("click", () => {
      closeConfigModal();
      window.dispatchEvent(new CustomEvent("cetas:open-agent", { detail: { id: s.id } }));
    });
    const delBtn = cfgEl("button", "sess-btn danger", "Supprimer");
    delBtn.type = "button";
    delBtn.addEventListener("click", async () => {
      const ok = await confirmDialog("Supprimer cet agent et son worktree ?", { okLabel: "Supprimer", danger: true });
      if (!ok) return;
      try {
        await api("/api/agents/" + encodeURIComponent(s.id), { method: "DELETE" });
        await refreshSessionsPanel();
      } catch (e) {}
    });
    acts.appendChild(openBtn);
    acts.appendChild(delBtn);
    row.appendChild(acts);
    agentsEl.appendChild(row);
  }
}

function agentStatusLabel(st) {
  if (st === "running") return "En cours";
  if (st === "done") return "Terminé";
  if (st === "stopped") return "Arrêté";
  return st;
}

// --- Panneau Apparence : swatches de palettes ---
let appearanceLoaded = false;
async function loadAppearancePanel() {
  if (appearanceLoaded) return;
  appearanceLoaded = true;
  renderPaletteSwatches();
}

function renderPaletteSwatches() {
  const wrap = document.getElementById("palette-swatches");
  if (!wrap) return;
  wrap.innerHTML = "";
  const current = document.documentElement.dataset.palette || "bleu";
  for (const p of THEME_PALETTES) {
    const b = cfgEl("button", "palette-swatch" + (p.id === current ? " active" : ""), "");
    b.type = "button";
    b.dataset.palette = p.id;
    b.title = p.label;
    b.setAttribute("aria-label", "Palette " + p.label);
    b.style.setProperty("--sw", p.color);
    b.addEventListener("click", async () => {
      applyPalette(p.id);
      wrap.querySelectorAll(".palette-swatch").forEach((s) => s.classList.toggle("active", s.dataset.palette === p.id));
      try {
        await putPrefs({ palette: p.id });
      } catch (e) {}
    });
    wrap.appendChild(b);
  }
}


// --- Helpers exportés pour les tests ---
export const _test = {
  FEATURE_DEFS,
  loadFeaturesPanel,
  loadSearchPanel,
  renderSearchTabs,
  selectSearchTab,
  currentSearchProvider,
  searchStatusText,
  loadSessionsPanel,
  refreshSessionsPanel,
  renderSessions,
  fmtSessionDate,
  agentStatusLabel,
  loadAppearancePanel,
  renderPaletteSwatches,
};
