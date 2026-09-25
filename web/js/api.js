import { clearSessionCache } from "./session-cache.js";

const TOKEN_KEY = "cetas-lite-token";
const USER_KEY = "cetas-lite-user";

export function getToken() {
  try {
    return localStorage.getItem(TOKEN_KEY) || "";
  } catch (e) {
    return "";
  }
}

export function setSession(token, user) {
  try {
    localStorage.setItem(TOKEN_KEY, token);
    localStorage.setItem(USER_KEY, JSON.stringify(user || {}));
  } catch (e) {}
}

export function clearSession() {
  try {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
  } catch (e) {}
  // Le cache des sessions est scopé par utilisateur : on le purge pour ne
  // laisser aucune métadonnée d'un compte précédent sur la machine.
  try {
    clearSessionCache();
  } catch (e) {}
}

export function getUser() {
  try {
    return JSON.parse(localStorage.getItem(USER_KEY) || "{}");
  } catch (e) {
    return {};
  }
}

export async function api(path, opts = {}) {
  const headers = Object.assign({}, opts.headers || {});
  if (opts.body !== undefined) headers["Content-Type"] = "application/json";
  const token = getToken();
  if (token) headers["Authorization"] = "Bearer " + token;

  let resp;
  try {
    resp = await fetch(path, {
      method: opts.method || "GET",
      headers,
      body: opts.body !== undefined ? JSON.stringify(opts.body) : undefined,
      signal: opts.signal,
    });
  } catch (e) {
    // Abort (timeout client) : remonte tel quel pour que l'appelant le
    // distingue d'une vraie panne reseau.
    if (e && e.name === "AbortError") throw e;
    throw new Error("Connexion au serveur impossible.");
  }

  const text = await resp.text();
  let data = null;
  if (text) {
    try {
      data = JSON.parse(text);
    } catch (e) {
      data = text;
    }
  }

  if (resp.status === 401 && path !== "/api/auth/login") {
    clearSession();
    window.dispatchEvent(new CustomEvent("cetas:unauthorized"));
  }
  if (!resp.ok) {
    const msg = data && data.error ? data.error : "Erreur " + resp.status;
    throw new Error(msg);
  }
  return data;
}

export function getPrefs() {
  return api("/api/settings");
}

export function putPrefs(prefs) {
  return api("/api/settings", { method: "PUT", body: prefs });
}

export async function readSSE(resp, onEvent) {
  const reader = resp.body.getReader();
  const decoder = new TextDecoder();
  let buf = "";
  for (;;) {
    const { value, done } = await reader.read();
    if (done) break;
    buf += decoder.decode(value, { stream: true });
    let idx;
    while ((idx = buf.indexOf("\n\n")) >= 0) {
      const block = buf.slice(0, idx);
      buf = buf.slice(idx + 2);
      for (const raw of block.split("\n")) {
        const line = raw.trim();
        if (!line.startsWith("data:")) continue;
        const payload = line.slice(5).trim();
        if (!payload) continue;
        let ev;
        try {
          ev = JSON.parse(payload);
        } catch (e) {
          continue;
        }
        onEvent(ev);
      }
    }
  }
}
