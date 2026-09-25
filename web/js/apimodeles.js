/* Panneau "API et Modèles" façon référence : onglets fournisseurs à icônes,
   clé API (masquée + Valider + œil), catalogue de modèles avec toggle
   Textes/Images, recherche et sélection persistée côté serveur. */

import { api } from "./api.js";

// --- Config statique (labels/licônes/liens) ---
const PROVIDERS_META = [
  { id: "openrouter", label: "OpenRouter", icon: "OpenRouter.svg", keyUrl: "https://openrouter.ai/settings/keys", keyLabel: "Obtenir une clé API OpenRouter", llmImages: true, catalog: "openrouter" },
  { id: "deepseek", label: "DeepSeek", icon: "DeepSeek.svg", keyUrl: "https://platform.deepseek.com/api_keys", keyLabel: "Obtenir une clé API DeepSeek", catalog: "static" },
  { id: "opencode", label: "OpenCode Zen", icon: "Opencode.svg", keyUrl: "https://opencode.ai/", keyLabel: "Obtenir une clé API OpenCode", catalog: "static" },
  { id: "opencode-go", label: "OpenCode Go", icon: "Opencode.svg", keyUrl: "https://opencode.ai/", keyLabel: "Obtenir une clé API OpenCode", catalog: "static" },
];
// Les moteurs SamGen (llama.cpp, Ollama, LM Studio) ont leur propre panneau
// "IA locale" (samgen.js, namespace /api/local/engines) : logique separee
// des providers cloud, ils ne sont plus listes ici.

const MAKER_LABELS = {
  "openai": "OpenAI", "anthropic": "Anthropic", "google": "Google",
  "meta-llama": "Meta", "mistralai": "Mistral", "deepseek": "DeepSeek",
  "qwen": "Qwen", "z-ai": "Z.ai", "poolside": "Poolside", "x-ai": "xAI",
  "cohere": "Cohere", "microsoft": "Microsoft", "nvidia": "Nvidia",
  "perplexity": "Perplexity", "openrouter": "OpenRouter", "amazon": "Amazon",
  "bytedance": "ByteDance", "minimax": "MiniMax", "moonshotai": "Moonshot",
};

const EYE_SHOW = '<svg class="apikey-eye-show" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>';
const EYE_HIDE = '<svg class="apikey-eye-hide" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/><path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/><line x1="1" y1="1" x2="23" y2="23"/></svg>';
const SEARCH_SVG = '<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="11" cy="11" r="7"/><line x1="21" y1="21" x2="16.5" y2="16.5"/></svg>';
const REFRESH_SVG = '<svg class="catalog-spin" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="23 4 23 10 17 10"/><path d="M20.49 15a9 9 0 1 1-.07-8.3"/></svg>';

function escHtml(s) {
  return String(s == null ? "" : s).replace(/[&<>"']/g, (c) => ({
    "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;",
  })[c]);
}

function makerLabel(id) {
  const prefix = String(id).split("/")[0] || "";
  if (MAKER_LABELS[prefix]) return MAKER_LABELS[prefix];
  return prefix ? prefix.charAt(0).toUpperCase() + prefix.slice(1) : "";
}

function priceStr(promptPer1M, completionPer1M) {
  const p = Number(promptPer1M) || 0;
  const c = Number(completionPer1M) || 0;
  if (!p && !c) return "Gratuit";
  const f = (v) => "$" + (v >= 100 ? v.toFixed(0) : v >= 1 ? v.toFixed(2) : v.toFixed(v >= 0.01 ? 3 : 4));
  return f(p) + " — " + f(c) + " /M";
}

// --- État ---
let _providers = [];
let _activeProvider = "openrouter";
let _catalogSel = {};          // provider -> Set(ids désactivés)
let _staticCatalog = null;     // GET /api/catalog
let _orModels = { text: null, image: null };  // cache session par type
let _orType = "text";
let _saveTimer = null;
let _initialized = false;

function metaFor(id) {
  return PROVIDERS_META.find((m) => m.id === id) || { id, label: id, icon: "" };
}

function providerState(id) {
  return _providers.find((p) => p.id === id) || { id, configured: false, local: false, url: "" };
}

export function isDisabled(provider, modelId) {
  const set = _catalogSel[provider];
  return !!(set && set.has(modelId));
}

// --- Point d'entrée ---
export async function initApiModelesPanel() {
  const tabsEl = document.getElementById("providers-tabs");
  const contentEl = document.getElementById("provider-content");
  if (!tabsEl || !contentEl) return;

  tabsEl.innerHTML = "";
  contentEl.innerHTML = '<div class="catalog-loading">Chargement des fournisseurs…</div>';
  try {
    const data = await api("/api/providers");
    _providers = (data && data.providers) || [];
  } catch (e) {
    contentEl.innerHTML = '<div class="catalog-error">Impossible de charger les fournisseurs : ' + escHtml(e.message) + "</div>";
    return;
  }
  try {
    const sel = await api("/api/catalog/selection");
    _catalogSel = {};
    const dis = (sel && sel.disabled) || {};
    for (const k of Object.keys(dis)) _catalogSel[k] = new Set(dis[k] || []);
  } catch (_) { _catalogSel = {}; }

  // Ordre : meta d'abord (référence visuelle), puis éventuels supplémentaires.
  const ordered = PROVIDERS_META.map((m) => m.id)
    .filter((id) => _providers.some((p) => p.id === id));
  for (const p of _providers) {
    if (!ordered.includes(p.id)) ordered.push(p.id);
  }

  tabsEl.innerHTML = ordered.map((id, i) => {
    const m = metaFor(id);
    const icon = m.icon ? "images/providers/" + m.icon : "";
    return '<button type="button" class="provider-tab' + (i === 0 ? " active" : "") + '" data-provider="' + escHtml(id) + '" title="' + escHtml(m.label) + '">' +
      (icon ? '<img src="' + icon + '" class="provider-tab-icon" alt="' + escHtml(m.label) + '">' : "") +
      '<span class="provider-tab-label">' + escHtml(m.label) + "</span></button>";
  }).join("");

  contentEl.innerHTML = ordered.map((id, i) => _sectionHtml(id, i === 0)).join("");

  tabsEl.querySelectorAll(".provider-tab").forEach((btn) => {
    btn.addEventListener("click", () => selectProvider(btn.dataset.provider));
  });

  for (const id of ordered) _bindSection(id);

  _activeProvider = ordered[0] || "openrouter";
  if (ordered.includes("openrouter")) _activeProvider = "openrouter";
  selectProvider(_activeProvider);
  _initialized = true;
  _bindSaveAll();
}

export function selectProvider(id) {
  _activeProvider = id;
  document.querySelectorAll("#providers-tabs .provider-tab").forEach((b) =>
    b.classList.toggle("active", b.dataset.provider === id));
  document.querySelectorAll("#provider-content .provider-section").forEach((s) =>
    s.classList.toggle("active", s.dataset.provider === id));
  const m = metaFor(id);
  if (m.catalog) loadCatalog(id, false);
}

// --- Section par fournisseur ---
function _sectionHtml(id, active) {
  const m = metaFor(id);
  const st = providerState(id);
  const maskedRow = st.configured
      ? '<div class="apikey-masked-row" id="apikey-masked-' + escHtml(id) + '">' +
        '<span class="apikey-masked-key">••••••••</span>' +
        '<button type="button" class="apikey-validate-btn" data-provider="' + escHtml(id) + '">Valider</button>' +
        '<button type="button" class="apikey-delete-link" data-provider="' + escHtml(id) + '" title="Supprimer la clé">Supprimer</button>' +
        "</div>"
      : '<div class="apikey-masked-row" id="apikey-masked-' + escHtml(id) + '" style="display:none">' +
        '<span class="apikey-masked-key">••••••••</span>' +
        '<button type="button" class="apikey-validate-btn" data-provider="' + escHtml(id) + '">Valider</button>' +
        '<button type="button" class="apikey-delete-link" data-provider="' + escHtml(id) + '" title="Supprimer la clé">Supprimer</button>' +
        "</div>";
  const keyBlock =
      '<div class="apikey-label-row"><label class="sp-modal-label" for="apikey-' + escHtml(id) + '">Clé API ' + escHtml(m.label) +
      (m.llmImages ? ' <span class="apikey-local-hint">(LLM &amp; Images)</span>' : "") + "</label>" +
      (m.keyUrl ? '<a class="apikey-get-link" href="' + escHtml(m.keyUrl) + '" target="_blank" rel="noopener noreferrer">' + escHtml(m.keyLabel) + "</a>" : "") +
      "</div>" +
      maskedRow +
      '<div class="apikey-field"><div class="apikey-input-wrap">' +
      '<input type="password" id="apikey-' + escHtml(id) + '" class="sp-modal-input apikey-input" placeholder="' + (st.configured ? "Nouvelle clé (laisser vide pour conserver)" : "Coller la clé API…") + '" autocomplete="off" spellcheck="false">' +
      '<button type="button" class="apikey-eye-btn" data-target="apikey-' + escHtml(id) + '" title="Afficher la clé" aria-label="Afficher la clé">' + EYE_SHOW + EYE_HIDE + "</button>" +
      "</div></div>";

  const catalogBlock = m.catalog
    ? '<div class="provider-models-header"><span class="provider-models-title">Sélectionnez les modèles à utiliser</span>' +
      (m.catalog === "openrouter"
        ? '<div class="catalog-type-btns" role="tablist">' +
          '<button type="button" class="catalog-type-btn active" data-type="text">Textes</button>' +
          '<button type="button" class="catalog-type-btn" data-type="image">Images</button></div>'
        : "") +
      "</div>" +
      '<div class="catalog-toolbar">' +
      '<select class="catalog-category-select" id="catalog-cat-' + escHtml(id) + '"><option value="all">Toutes les catégories</option></select>' +
      '<div class="catalog-search-wrap"><span class="catalog-search-icon">' + SEARCH_SVG + "</span>" +
      '<input type="text" class="catalog-search-input" id="catalog-search-' + escHtml(id) + '" placeholder="Rechercher un modèle…" autocomplete="off"></div>' +
      '<button type="button" class="catalog-refresh-btn" id="catalog-refresh-' + escHtml(id) + '" title="Actualiser">' + REFRESH_SVG + "</button>" +
      "</div>" +
      '<div class="catalog-list provider-catalog-list" id="catalog-list-' + escHtml(id) + '" data-provider="' + escHtml(id) + '"></div>'
    : "";

  return '<div class="provider-section' + (active ? " active" : "") + '" data-provider="' + escHtml(id) + '">' +
    keyBlock + catalogBlock + "</div>";
}

function _bindSection(id) {
  const sec = document.querySelector('#provider-content .provider-section[data-provider="' + id + '"]');
  if (!sec) return;

  // Œil afficher/masquer.
  sec.querySelectorAll(".apikey-eye-btn").forEach((eye) => {
    eye.addEventListener("click", () => {
      const input = document.getElementById(eye.dataset.target);
      if (!input) return;
      const show = input.type === "password";
      input.type = show ? "text" : "password";
      eye.classList.toggle("shown", show);
    });
  });

  // Valider : enregistre la clé saisie.
  sec.querySelectorAll(".apikey-validate-btn").forEach((btn) => {
    btn.addEventListener("click", () => _validateKey(id));
  });
  const input = sec.querySelector("#apikey-" + id);
  if (input) input.addEventListener("keydown", (e) => {
    if (e.key === "Enter") _validateKey(id);
  });

  // Supprimer.
  sec.querySelectorAll(".apikey-delete-link").forEach((btn) => {
    btn.addEventListener("click", async () => {
      const m2 = metaFor(id);
      if (!window.confirm("Supprimer la clé du fournisseur « " + m2.label + " » ?")) return;
      try {
        await api("/api/providers/" + encodeURIComponent(id), { method: "DELETE" });
        const st = providerState(id);
        st.configured = false;
        const row = sec.querySelector("#apikey-masked-" + id);
        if (row) row.style.display = "none";
        if (input) { input.value = ""; input.placeholder = "Coller la clé API…"; }
      } catch (e) {
        alert("Échec de la suppression : " + e.message);
      }
    });
  });

  // Catalogue : toggle Textes/Images, recherche, refresh.
  sec.querySelectorAll(".catalog-type-btn").forEach((btn) => {
    btn.addEventListener("click", () => {
      _orType = btn.dataset.type;
      sec.querySelectorAll(".catalog-type-btn").forEach((b) => b.classList.toggle("active", b === btn));
      loadCatalog(id, false);
    });
  });
  const search = sec.querySelector("#catalog-search-" + id);
  if (search) search.addEventListener("input", () => _renderCatalogList(id));
  const cat = sec.querySelector("#catalog-cat-" + id);
  if (cat) cat.addEventListener("change", () => _renderCatalogList(id));
  const refresh = sec.querySelector("#catalog-refresh-" + id);
  if (refresh) refresh.addEventListener("click", () => loadCatalog(id, true));
}

async function _validateKey(id) {
  const sec = document.querySelector('#provider-content .provider-section[data-provider="' + id + '"]');
  const input = sec && sec.querySelector("#apikey-" + id);
  const btn = sec && sec.querySelector('.apikey-validate-btn[data-provider="' + id + '"]');
  const key = input ? input.value.trim() : "";
  if (!key) { if (input) input.focus(); return; }
  if (btn) btn.disabled = true;
  try {
    await api("/api/providers/" + encodeURIComponent(id), { method: "PUT", body: { key } });
    const st = providerState(id);
    st.configured = true;
    const row = sec.querySelector("#apikey-masked-" + id);
    if (row) row.style.display = "";
    if (input) { input.value = ""; input.placeholder = "Nouvelle clé (laisser vide pour conserver)"; }
  } catch (e) {
    alert("Échec de l'enregistrement : " + e.message);
  } finally {
    if (btn) btn.disabled = false;
  }
}

// --- Catalogue ---
async function _ensureStaticCatalog() {
  if (_staticCatalog) return _staticCatalog;
  try {
    const data = await api("/api/catalog");
    _staticCatalog = (data && data.providers) || [];
  } catch (_) { _staticCatalog = []; }
  return _staticCatalog;
}

async function loadCatalog(id, refresh) {
  const m = metaFor(id);
  if (!m.catalog) return;
  const listEl = document.getElementById("catalog-list-" + id);
  if (!listEl) return;
  listEl.innerHTML = '<div class="catalog-loading">Chargement des modèles…</div>';

  let models = [];
  if (m.catalog === "openrouter") {
    const type = _orType;
    if (!refresh && _orModels[type]) {
      models = _orModels[type];
    } else {
      try {
        const data = await api("/api/openrouter/models?type=" + type + (refresh ? "&refresh=1" : ""));
        models = ((data && data.models) || []).map((x) => ({
          id: x.id, name: x.name, maker: makerLabel(x.id),
          promptPer1M: x.prompt_per_1m, completionPer1M: x.completion_per_1m,
          desc: x.description || "",
          tooltip: x.name + (x.context_length ? "\nContexte : " + Number(x.context_length).toLocaleString("fr-FR") + " tokens" : "") +
            (x.description ? "\n\n" + x.description : ""),
        }));
        _orModels[type] = models;
      } catch (_) {
        // Repli hors-ligne : catalogue statique embarqué côté serveur.
        const cat = await _ensureStaticCatalog();
        const prov = cat.find((p) => p.id === "openrouter");
        models = ((prov && prov.models) || []).map((x) => ({
          id: x.id, name: x.label, maker: makerLabel(x.id),
          promptPer1M: x.input_per_1m, completionPer1M: x.output_per_1m,
          desc: "", tooltip: x.label,
        }));
      }
    }
  } else {
    const cat = await _ensureStaticCatalog();
    const prov = cat.find((p) => p.id === id);
    models = ((prov && prov.models) || []).map((x) => ({
      id: x.id, name: x.label, maker: m.label,
      promptPer1M: x.input_per_1m, completionPer1M: x.output_per_1m,
      desc: "", tooltip: x.label,
    }));
  }

  listEl._models = models;
  _fillCategories(id, models);
  _renderCatalogList(id);
}

function _fillCategories(id, models) {
  const sel = document.getElementById("catalog-cat-" + id);
  if (!sel) return;
  const makers = [...new Set(models.map((x) => x.maker).filter(Boolean))].sort((a, b) =>
    a.localeCompare(b, "fr"));
  const cur = sel.value || "all";
  sel.innerHTML = '<option value="all">Toutes les catégories</option>' +
    makers.map((x) => '<option value="' + escHtml(x) + '">' + escHtml(x) + "</option>").join("");
  sel.value = makers.includes(cur) ? cur : "all";
}

function _renderCatalogList(id) {
  const listEl = document.getElementById("catalog-list-" + id);
  if (!listEl || !listEl._models) return;
  const models = listEl._models;
  const catSel = document.getElementById("catalog-cat-" + id);
  const searchEl = document.getElementById("catalog-search-" + id);
  const cat = catSel ? catSel.value : "all";
  const q = searchEl ? searchEl.value.trim().toLowerCase() : "";

  const rows = models.filter((x) => {
    if (cat !== "all" && x.maker !== cat) return false;
    if (q && !(x.name.toLowerCase().includes(q) || x.id.toLowerCase().includes(q))) return false;
    return true;
  });

  if (!rows.length) {
    listEl.innerHTML = '<div class="catalog-empty">Aucun modèle trouvé.</div>';
    return;
  }
  listEl.innerHTML = rows.map((x) => {
    const checked = !isDisabled(id, x.id);
    const info = x.tooltip
      ? '<span class="catalog-row-info" data-tooltip="' + escHtml(x.tooltip.slice(0, 600)) + '">i</span>'
      : '<span class="catalog-row-info-spacer" aria-hidden="true"></span>';
    return '<div class="catalog-row">' +
      '<label class="catalog-row-main">' +
      '<input type="checkbox" class="catalog-cb" data-provider="' + escHtml(id) + '" data-id="' + escHtml(x.id) + '"' + (checked ? " checked" : "") + ">" +
      '<span class="catalog-row-name">' + escHtml(x.name) + "</span>" +
      "</label>" +
      '<span class="catalog-row-maker">' + escHtml(x.maker || "") + "</span>" +
      '<span class="catalog-row-price">' + escHtml(priceStr(x.promptPer1M, x.completionPer1M)) + "</span>" +
      info + "</div>";
  }).join("");

  listEl.querySelectorAll(".catalog-cb").forEach((cb) => {
    cb.addEventListener("change", () => {
      const prov = cb.dataset.provider;
      const mid = cb.dataset.id;
      if (!_catalogSel[prov]) _catalogSel[prov] = new Set();
      if (cb.checked) _catalogSel[prov].delete(mid);
      else _catalogSel[prov].add(mid);
      _scheduleSaveSelection();
    });
  });
}

function _scheduleSaveSelection() {
  if (_saveTimer) clearTimeout(_saveTimer);
  _saveTimer = setTimeout(() => _saveSelection(), 800);
}

async function _saveSelection() {
  if (_saveTimer) { clearTimeout(_saveTimer); _saveTimer = null; }
  const disabled = {};
  for (const k of Object.keys(_catalogSel)) {
    if (_catalogSel[k].size) disabled[k] = [..._catalogSel[k]];
  }
  try {
    await api("/api/catalog/selection", { method: "PUT", body: { disabled } });
  } catch (_) { /* réessayé au prochain changement */ }
}

// --- Sauvegarder tout ---
function _bindSaveAll() {
  const btn = document.getElementById("apimodeles-save-btn");
  if (!btn || btn._bound) return;
  btn._bound = true;
  btn.addEventListener("click", async () => {
    btn.disabled = true;
    try {
      // Clés saisies mais non validées.
      for (const m of PROVIDERS_META) {
        const input = document.getElementById("apikey-" + m.id);
        const key = input ? input.value.trim() : "";
        if (key) {
          await api("/api/providers/" + encodeURIComponent(m.id), { method: "PUT", body: { key } });
          const st = providerState(m.id);
          st.configured = true;
          const row = document.getElementById("apikey-masked-" + m.id);
          if (row) row.style.display = "";
          input.value = "";
          input.placeholder = "Nouvelle clé (laisser vide pour conserver)";
        }
      }
      await _saveSelection();
      const label = btn.textContent;
      btn.classList.add("models-save-btn--saved");
      btn.textContent = "Sauvegardé ✓";
      setTimeout(() => {
        btn.classList.remove("models-save-btn--saved");
        btn.textContent = label;
        btn.disabled = false;
      }, 1300);
    } catch (e) {
      btn.disabled = false;
      alert("Échec de la sauvegarde : " + e.message);
    }
  });
}

// --- Helpers exportés pour les tests ---
export const _test = {
  PROVIDERS_META, escHtml, makerLabel, priceStr, isDisabled,
  _sectionHtml, _selectProviderForTest: (id) => { _activeProvider = id; },
};
