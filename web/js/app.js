import { api, getToken, clearSession, getUser } from "./api.js";
import { initAuth, showLogin } from "./auth.js";
import { initSidebar } from "./sidebar.js";
import { initChat } from "./chat.js";
import { initModels } from "./model-select.js";
import { initRightPanel } from "./right-panel.js";
import { initModals } from "./modals.js";
import { initAgents } from "./agents.js";
import { initTerminal } from "./terminal.js";
import { initMetrics } from "./metrics.js";

const SPLASH_MIN_MS = 1400;

function hideSplash() {
  const sp = document.getElementById("kiro-splash");
  if (!sp) return;
  sp.classList.add("fading-out");
  setTimeout(() => sp.remove(), 600);
}

async function boot() {
  const t0 = Date.now();
  // Version + registration
  try {
    const cfg = await api("/api/config");
    const badge = document.getElementById("version-badge");
    if (badge && cfg && cfg.version) badge.textContent = "v" + cfg.version;
    window.CETAS_CONFIG = { registrationOpen: !!(cfg && cfg.registration_open) };
  } catch (e) {
    window.CETAS_CONFIG = { registrationOpen: false };
  }

  // Auth
  let me = null;
  if (getToken()) {
    try {
      me = await api("/api/me");
    } catch (e) {
      me = null;
    }
  }
  if (!me) {
    const wait = Math.max(0, SPLASH_MIN_MS - (Date.now() - t0));
    setTimeout(() => {
      hideSplash();
      showLogin();
    }, wait);
    initAuth(() => boot());
    return;
  }

  // Session OK — init des modules
  const initials = document.getElementById("user-avatar-initials");
  const uname = (me && me.username) || getUser().username || "?";
  if (initials) initials.textContent = String(uname).slice(0, 2).toUpperCase();

  try {
    initSidebar();
    initModels();
    initChat();
    initRightPanel();
    initModals();
    initAgents();
    initTerminal();
    initMetrics();
  } catch (e) {
    console.error("[app] init modules:", e);
  }

  const wait = Math.max(0, SPLASH_MIN_MS - (Date.now() - t0));
  setTimeout(hideSplash, wait);

  window.addEventListener("cetas:unauthorized", () => {
    clearSession();
    showLogin();
  });
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", boot);
} else {
  boot();
}
