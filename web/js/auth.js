import { api, clearSession, getToken, setSession } from "./api.js";

let authedCb = null;

export function showLogin() {
  const overlay = document.getElementById("login-overlay");
  if (overlay) overlay.style.display = "flex";
  const u = document.getElementById("login-username");
  const p = document.getElementById("login-password");
  try {
    const saved = JSON.parse(localStorage.getItem("cetas-lite-user") || "{}");
    if (saved && saved.username && u && !u.value) u.value = saved.username;
  } catch (e) {}
  setTimeout(() => {
    if (u && u.value && p) p.focus();
    else if (u) u.focus();
  }, 60);
}

export function hideLogin() {
  const overlay = document.getElementById("login-overlay");
  if (overlay) overlay.style.display = "none";
}

export function initAuth(onAuthenticated) {
  authedCb = onAuthenticated;
  const form = document.getElementById("login-form");
  const errEl = document.getElementById("login-error");
  const submit = document.getElementById("login-btn");
  const toggle = document.getElementById("login-mode-toggle");
  let registering = false;

  function setMode(reg) {
    registering = reg;
    if (submit) submit.textContent = reg ? "Créer le compte" : "Se connecter";
    if (toggle) {
      toggle.textContent = reg ? "J'ai déjà un compte" : "Créer un compte";
      // Masquer le lien si l'inscription est fermée et aucun compte n'existe pas encore
      toggle.style.display = !reg && window.CETAS_CONFIG && window.CETAS_CONFIG.registrationOpen === false ? "none" : "";
    }
    if (errEl) errEl.style.display = "none";
  }

  if (toggle) {
    toggle.addEventListener("click", (e) => {
      e.preventDefault();
      setMode(!registering);
    });
  }

  // Si l'inscription est fermée, rester en mode login
  if (window.CETAS_CONFIG && window.CETAS_CONFIG.registrationOpen === false) {
    setMode(false);
  }

  if (!form) return;
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("login-username").value.trim();
    const password = document.getElementById("login-password").value;
    if (!username || !password) return;
    if (errEl) errEl.style.display = "none";
    if (submit) submit.disabled = true;
    try {
      const data = await api(registering ? "/api/auth/register" : "/api/auth/login", {
        method: "POST",
        body: { username, password },
      });
      setSession(data.token, { username });
      hideLogin();
      if (authedCb) authedCb();
    } catch (err) {
      if (errEl) {
        errEl.textContent = err.message || "Échec de connexion.";
        errEl.style.display = "block";
      }
    } finally {
      if (submit) submit.disabled = false;
    }
  });
}

export function logout() {
  clearSession();
  location.reload();
}
