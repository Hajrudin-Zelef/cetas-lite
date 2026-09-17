// Cache local (localStorage) de la liste des sessions.
//
// Rôle : peinture instantanée de la sidebar (pas d'attente réseau au
// démarrage). Le serveur reste la source de vérité : chaque lecture
// revalide en arrière-plan, chaque mutation écrit en direct.
//
// Sécurité / hygiène :
//  - clé scopée par nom d'utilisateur : aucun mélange entre comptes ;
//  - ne contient que les métadonnées (id, titre, dates, nb messages),
//    jamais le contenu des messages ;
//  - purge complète à la déconnexion (via clearSession dans api.js).

const PREFIX = "cetas.sessions.v1:";
const USER_KEY = "cetas-lite-user";
const MAX_ITEMS = 300;

function username() {
  try {
    const u = JSON.parse(localStorage.getItem(USER_KEY) || "{}");
    return u.username || u.name || "anon";
  } catch (e) {
    return "anon";
  }
}

function cacheKey() {
  return PREFIX + username();
}

function read() {
  try {
    const raw = localStorage.getItem(cacheKey());
    if (!raw) return null;
    const data = JSON.parse(raw);
    if (!data || !Array.isArray(data.sessions)) return null;
    return data.sessions;
  } catch (e) {
    return null;
  }
}

function write(sessions) {
  try {
    localStorage.setItem(
      cacheKey(),
      JSON.stringify({ v: 1, at: Date.now(), sessions: sessions.slice(0, MAX_ITEMS) })
    );
  } catch (e) {}
}

// Lecture du cache (peinture instantanée). Retourne null si vide/invalide.
export function getCachedSessions() {
  return read();
}

// Remplacement complet (après revalidation serveur).
export function setCachedSessions(list) {
  write(Array.isArray(list) ? list : []);
}

// Insertion / mise à jour d'une session (write-through après création).
export function upsertCachedSession(info) {
  if (!info || !info.id) return;
  const cur = read() || [];
  const i = cur.findIndex((s) => s && s.id === info.id);
  if (i >= 0) cur[i] = Object.assign({}, cur[i], info);
  else cur.unshift(info);
  write(cur);
}

// Retrait d'une session (write-through après suppression) : une session
// supprimée ne doit jamais réapparaître, même fugitivement.
export function removeCachedSession(id) {
  const cur = read() || [];
  write(cur.filter((s) => s && s.id !== id));
}

// Purge de tous les caches de sessions (tous utilisateurs) : déconnexion.
export function clearSessionCache() {
  try {
    const gone = [];
    for (let i = 0; i < localStorage.length; i++) {
      const k = localStorage.key(i);
      if (k && k.indexOf(PREFIX) === 0) gone.push(k);
    }
    gone.forEach((k) => localStorage.removeItem(k));
  } catch (e) {}
}
