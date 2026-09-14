import { api, clearSession, getToken, setSession } from "./api.js";

export function initAuth(onAuthenticated) {
  const overlay = document.getElementById("login-overlay");
  const app = document.getElementById("app");
  const form = document.getElementById("login-form");
  const errEl = document.getElementById("login-error");
  const submit = document.getElementById("login-submit");
  const toggle = document.getElementById("login-toggle");
  let registering = false;

  fetch("/api/config")
    .then((r) => r.json())
    .then((cfg) => {
      if (!cfg.registration_open) toggle.hidden = true;
    })
    .catch(() => {});

  function showLogin() {
    app.hidden = true;
    overlay.hidden = false;
  }

  async function showApp() {
    overlay.hidden = true;
    app.hidden = false;
    const me = await api("/api/me").catch(() => null);
    if (!me) {
      showLogin();
      return;
    }
    document.getElementById("user-name").textContent = me.username;
    document.getElementById("user-initials").textContent = (me.username || "?").slice(0, 1).toUpperCase();
    onAuthenticated(me);
  }

  toggle.addEventListener("click", () => {
    registering = !registering;
    submit.textContent = registering ? "Creer le compte" : "Se connecter";
    toggle.textContent = registering ? "J'ai deja un compte" : "Creer un compte";
    errEl.hidden = true;
  });

  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    errEl.hidden = true;
    const username = document.getElementById("login-username").value.trim();
    const password = document.getElementById("login-password").value;
    try {
      if (registering) {
        await api("/api/auth/register", { method: "POST", body: { username, password } });
      }
      const out = await api("/api/auth/login", { method: "POST", body: { username, password } });
      setSession(out.token, out.user);
      await showApp();
    } catch (err) {
      errEl.textContent = err.message;
      errEl.hidden = false;
    }
  });

  document.getElementById("logout-btn").addEventListener("click", () => {
    clearSession();
    location.reload();
  });
  window.addEventListener("cetas:unauthorized", () => {
    errEl.textContent = "Session expirée. Reconnecte-toi pour continuer.";
    errEl.hidden = false;
    showLogin();
  });

  if (getToken()) showApp();
  else showLogin();
}
