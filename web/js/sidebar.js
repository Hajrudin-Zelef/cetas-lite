import { api, getToken } from "./api.js";
import { logout } from "./auth.js";
import { confirmDialog } from "./modals.js";

const CATS_KEY = "cetas-lite-cats";
const CONV_CATS_KEY = "cetas-lite-conv-cats";
const FAVS_KEY = "cetas-lite-favs";

function loadJSON(key, fallback) {
  try {
    const v = JSON.parse(localStorage.getItem(key) || "null");
    return v == null ? fallback : v;
  } catch (e) {
    return fallback;
  }
}
function saveJSON(key, val) {
  try {
    localStorage.setItem(key, JSON.stringify(val));
  } catch (e) {}
}

export function getCategories() {
  return loadJSON(CATS_KEY, []);
}
export function saveCategories(cats) {
  saveJSON(CATS_KEY, cats);
  window.dispatchEvent(new CustomEvent("cetas:cats-changed"));
}
export function getConvCats() {
  return loadJSON(CONV_CATS_KEY, {});
}
export function setConvCat(convId, catId) {
  const m = getConvCats();
  if (catId) m[convId] = catId;
  else delete m[convId];
  saveJSON(CONV_CATS_KEY, m);
}
export function getFavs() {
  return loadJSON(FAVS_KEY, []);
}
export function toggleFav(convId) {
  let favs = getFavs();
  favs = favs.includes(convId) ? favs.filter((x) => x !== convId) : [...favs, convId];
  saveJSON(FAVS_KEY, favs);
  return favs.includes(convId);
}

function fmtDate(ms) {
  if (!ms) return "";
  try {
    return new Date(ms).toLocaleDateString(undefined, { day: "numeric", month: "short" }) +
      " " + new Date(ms).toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
  } catch (e) {
    return "";
  }
}

async function download(url, name) {
  let resp;
  try {
    resp = await fetch(url, { headers: { Authorization: "Bearer " + getToken() } });
  } catch (e) {
    return;
  }
  if (!resp.ok) return;
  const blob = await resp.blob();
  const a = document.createElement("a");
  a.href = URL.createObjectURL(blob);
  a.download = name;
  document.body.appendChild(a);
  a.click();
  a.remove();
  setTimeout(() => URL.revokeObjectURL(a.href), 1000);
}

export function initSidebar() {
  const list = document.getElementById("conv-list");
  const searchInput = document.getElementById("conv-search");
  const newBtn = document.getElementById("new-chat-btn");
  const catSelect = document.getElementById("cat-select");
  const catDropdown = document.getElementById("cat-select-dropdown");
  const catLabel = catSelect ? catSelect.querySelector(".cat-select-label") : null;
  let activeCat = "";
  let archives = [];

  function filtered() {
    const q = (searchInput ? searchInput.value : "").trim().toLowerCase();
    const convCats = getConvCats();
    return archives.filter((a) => {
      if (activeCat && convCats[a.id] !== activeCat) return false;
      if (q && !(a.title || "").toLowerCase().includes(q)) return false;
      return true;
    });
  }

  async function restore(id) {
    const state = await api("/api/chat/state").catch(() => ({}));
    if ((state.turns || 0) > 0) {
      const ok = await confirmDialog("Restaurer cette conversation ? La conversation actuelle sera archivée.", { okLabel: "Restaurer" });
      if (!ok) return;
    }
    try {
      await api("/api/conversations/restore", { method: "POST", body: { id } });
      window.dispatchEvent(new CustomEvent("cetas:chat-reset"));
      await refresh();
    } catch (e) {}
  }

  async function remove(id) {
    const ok = await confirmDialog("Supprimer définitivement cette conversation ?", { okLabel: "Supprimer", danger: true });
    if (!ok) return;
    try {
      await api("/api/conversations/" + encodeURIComponent(id), { method: "DELETE" });
      await refresh();
    } catch (e) {}
  }

  function convItem(a) {
    const row = document.createElement("div");
    row.className = "conv-item";
    const favs = getFavs();
    const isFav = favs.includes(a.id);
    if (isFav) row.classList.add("conv-item-fav");

    const content = document.createElement("div");
    content.className = "conv-item-content";
    const title = document.createElement("div");
    title.className = "conv-item-title";
    title.textContent = a.title || "(sans titre)";
    title.title = a.title || "";
    const dateLine = document.createElement("div");
    dateLine.className = "conv-item-date-line";
    const n = a.messages || 0;
    dateLine.textContent = n + " message" + (n > 1 ? "s" : "") + (a.updated ? " · " + fmtDate(a.updated) : "");
    content.appendChild(title);
    content.appendChild(dateLine);
    content.addEventListener("click", () => restore(a.id));

    const actions = document.createElement("div");
    actions.className = "conv-item-actions";

    const fav = document.createElement("button");
    fav.type = "button";
    fav.className = "conv-action-btn" + (isFav ? " active" : "");
    fav.title = isFav ? "Retirer des favoris" : "Ajouter aux favoris";
    fav.textContent = "★";
    fav.addEventListener("click", (e) => {
      e.stopPropagation();
      toggleFav(a.id);
      refresh();
    });

    const exp = document.createElement("button");
    exp.type = "button";
    exp.className = "conv-action-btn";
    exp.title = "Exporter (Markdown)";
    exp.innerHTML = '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>';
    exp.addEventListener("click", (e) => {
      e.stopPropagation();
      download("/api/conversations/" + encodeURIComponent(a.id) + "/export?format=md", (a.title || a.id) + ".md");
    });

    const del = document.createElement("button");
    del.type = "button";
    del.className = "conv-action-btn danger";
    del.title = "Supprimer";
    del.textContent = "×";
    del.addEventListener("click", (e) => {
      e.stopPropagation();
      remove(a.id);
    });

    actions.appendChild(fav);
    actions.appendChild(exp);
    actions.appendChild(del);
    row.appendChild(content);
    row.appendChild(actions);
    return row;
  }

  async function refresh() {
    if (!list) return;
    let data;
    try {
      data = await api("/api/conversations");
    } catch (e) {
      return;
    }
    archives = (data && data.archives) || [];
    render();
  }

  function render() {
    if (!list) return;
    list.innerHTML = "";
    const items = filtered();
    if (!items.length) {
      const empty = document.createElement("div");
      empty.className = "conv-empty";
      empty.textContent = archives.length ? "Aucun résultat." : "Aucune conversation archivée.";
      list.appendChild(empty);
      return;
    }
    for (const a of items) list.appendChild(convItem(a));
    renderFavs();
  }

  function renderFavs() {
    const favSection = document.getElementById("fav-section");
    const favList = document.getElementById("fav-list");
    if (!favSection || !favList) return;
    const favs = getFavs();
    const favItems = archives.filter((a) => favs.includes(a.id));
    favSection.style.display = favItems.length ? "" : "none";
    favList.innerHTML = "";
    for (const a of favItems.slice(0, 8)) {
      const b = document.createElement("button");
      b.type = "button";
      b.className = "fav-item";
      b.title = a.title || "";
      b.textContent = "★ " + (a.title || "(sans titre)");
      b.addEventListener("click", () => restore(a.id));
      favList.appendChild(b);
    }
  }

  function renderCatDropdown() {
    if (!catDropdown) return;
    catDropdown.innerHTML = "";
    const cats = getCategories();
    const mk = (id, label) => {
      const d = document.createElement("div");
      d.className = "cat-select-option" + (activeCat === id ? " active" : "");
      d.textContent = label;
      d.addEventListener("click", (e) => {
        e.stopPropagation();
        activeCat = id;
        if (catLabel) catLabel.textContent = label;
        catDropdown.style.display = "none";
        render();
      });
      catDropdown.appendChild(d);
    };
    mk("", "Toutes les catégories");
    for (const c of cats) mk(c.id, (c.emoji ? c.emoji + " " : "") + c.name);
  }

  if (catSelect && catDropdown) {
    catSelect.addEventListener("click", (e) => {
      e.stopPropagation();
      renderCatDropdown();
      catDropdown.style.display = catDropdown.style.display === "none" ? "" : "none";
    });
    document.addEventListener("click", () => {
      catDropdown.style.display = "none";
    });
    const manageBtn = document.getElementById("cat-manage-btn");
    if (manageBtn) manageBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      window.dispatchEvent(new CustomEvent("cetas:open-cat-modal"));
    });
  }

  if (searchInput) searchInput.addEventListener("input", render);

  if (newBtn) {
    newBtn.addEventListener("click", async () => {
      // "Nouvelle conversation" concerne le chat : on quitte le module Agent.
      window.dispatchEvent(new CustomEvent("cetas:close-agents"));
      const state = await api("/api/chat/state").catch(() => ({}));
      if ((state.turns || 0) > 0) {
        const ok = await confirmDialog("Démarrer une nouvelle conversation ? La conversation actuelle sera archivée.", { okLabel: "Nouvelle conversation" });
        if (!ok) return;
      }
      try {
        await api("/api/chat/reset", { method: "POST" });
      } catch (e) {}
      window.dispatchEvent(new CustomEvent("cetas:chat-reset"));
      await refresh();
    });
  }

  // Toggle sidebar
  const toggle = document.getElementById("sidebar-toggle");
  const sidebar = document.getElementById("sidebar");
  if (toggle && sidebar) {
    toggle.addEventListener("click", () => {
      document.body.classList.toggle("sidebar-hidden");
    });
  }

  // Menu utilisateur
  const avatarBtn = document.getElementById("user-avatar-btn");
  const userMenu = document.getElementById("user-menu-dropdown");
  if (avatarBtn && userMenu) {
    avatarBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      userMenu.style.display = userMenu.style.display === "none" ? "" : "none";
    });
    document.addEventListener("click", () => {
      userMenu.style.display = "none";
    });
    document.getElementById("sidebar-roles-btn")?.addEventListener("click", () => {
      userMenu.style.display = "none";
      window.dispatchEvent(new CustomEvent("cetas:open-roles-modal"));
    });
    document.getElementById("sidebar-prompts-btn")?.addEventListener("click", () => {
      userMenu.style.display = "none";
      window.dispatchEvent(new CustomEvent("cetas:open-prompts-modal"));
    });
    document.getElementById("dashboard-btn")?.addEventListener("click", () => {
      userMenu.style.display = "none";
      window.dispatchEvent(new CustomEvent("cetas:open-save-modal"));
    });
    document.getElementById("sidebar-faq-btn")?.addEventListener("click", () => {
      userMenu.style.display = "none";
      window.dispatchEvent(new CustomEvent("cetas:open-config-faq"));
    });
    document.getElementById("theme-cycle-btn")?.addEventListener("click", () => {
      userMenu.style.display = "none";
      const cur = document.documentElement.dataset.theme || "ocean";
      const next = cur === "ocean" ? "sombre" : cur === "sombre" ? "clair" : "ocean";
      window.dispatchEvent(new CustomEvent("cetas:theme", { detail: next }));
    });
    document.getElementById("logout-btn")?.addEventListener("click", () => logout());
  }

  // Configuration
  document.getElementById("apikeys-btn")?.addEventListener("click", () => {
    window.dispatchEvent(new CustomEvent("cetas:open-config"));
  });

  // Modules dev
  const devToast = document.getElementById("dev-toast");
  let devTimer = null;
  function toastDev() {
    if (!devToast) return;
    devToast.style.display = "";
    clearTimeout(devTimer);
    devTimer = setTimeout(() => (devToast.style.display = "none"), 1800);
  }
  document.querySelectorAll(".dev-module-btn").forEach((b) => {
    b.addEventListener("click", () => {
      const mod = b.dataset.module;
      if (mod === "agents") window.dispatchEvent(new CustomEvent("cetas:toggle-agents"));
      else if (mod === "terminal") window.dispatchEvent(new CustomEvent("cetas:toggle-terminal"));
      else toastDev();
    });
  });

  // Effacer toutes les conversations
  const clearBtn = document.getElementById("clear-all-btn");
  const clearOverlay = document.getElementById("clear-all-overlay");
  if (clearBtn && clearOverlay) {
    const codeEl = document.getElementById("clear-all-code");
    const codeInput = document.getElementById("clear-all-code-input");
    const confirmBtn = document.getElementById("clear-all-confirm");
    const cancelBtn = document.getElementById("clear-all-cancel");
    let code = "";
    clearBtn.addEventListener("click", () => {
      code = Array.from({ length: 6 }, () => "ABCDEFGHJKMNPQRSTUVWXYZ23456789"[Math.floor(Math.random() * 32)]).join("");
      if (codeEl) codeEl.textContent = code;
      if (codeInput) codeInput.value = "";
      if (confirmBtn) confirmBtn.disabled = true;
      clearOverlay.style.display = "";
    });
    if (codeInput) {
      codeInput.addEventListener("input", () => {
        if (confirmBtn) confirmBtn.disabled = codeInput.value.trim().toUpperCase() !== code;
      });
    }
    if (cancelBtn) cancelBtn.addEventListener("click", () => (clearOverlay.style.display = "none"));
    clearOverlay.addEventListener("click", (e) => {
      if (e.target === clearOverlay) clearOverlay.style.display = "none";
    });
    if (confirmBtn) {
      confirmBtn.addEventListener("click", async () => {
        confirmBtn.disabled = true;
        const items = archives.slice();
        for (const a of items) {
          try {
            await api("/api/conversations/" + encodeURIComponent(a.id), { method: "DELETE" });
          } catch (e) {}
        }
        clearOverlay.style.display = "none";
        await refresh();
      });
    }
  }

  window.addEventListener("cetas:chat-changed", refresh);
  window.addEventListener("cetas:cats-changed", () => {
    renderCatDropdown();
    render();
  });

  refresh();
  return { refresh };
}
