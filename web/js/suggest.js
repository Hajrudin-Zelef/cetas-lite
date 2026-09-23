// Questions suggérées — diversion d'accueil (itération 4).
//
// Au premier « salut » sur un fil vide, le navigateur tire 3 questions du
// pool servi par GET /api/chat/suggestions et les affiche en chips
// cliquables. Le clic pose la question avec focus corpus (le serveur
// sollicite forcément le RAG sur ce corpus). Tirage local : épinglés en
// priorité, mixité des corpus, anti-répétition via localStorage.
import { api } from "./api.js";

const SEEN_KEY = "cetas-lite-suggest-seen";
const SEEN_MAX = 15;

let poolCache = null;

// --- Détection de salutation (miroir de chat.IsGreeting côté serveur) ---
const GREETINGS = new Set([
  "salut",
  "bonjour",
  "bonsoir",
  "hello",
  "hi",
  "hey",
  "coucou",
  "yo",
]);

export function isGreeting(text) {
  const t = String(text || "")
    .toLowerCase()
    .trim()
    .replace(/[ !?.,…]+$/u, "")
    .trim();
  return GREETINGS.has(t);
}

// --- Pool ---
export async function loadPool() {
  if (poolCache) return poolCache;
  const data = await api("/api/chat/suggestions");
  const list = data && Array.isArray(data.suggestions) ? data.suggestions : [];
  poolCache = list.filter((s) => s && s.question && s.corpus);
  return poolCache;
}

export function clearPoolCache() {
  poolCache = null;
}

export function keyOf(s) {
  return String(s.question || "").trim().toLowerCase();
}

function shuffle(arr) {
  const a = arr.slice();
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [a[i], a[j]] = [a[j], a[i]];
  }
  return a;
}

// --- Tirage : épinglés d'abord, puis diversité des corpus ---
export function drawSuggestions(pool, { count = 3, exclude = [] } = {}) {
  const excluded = new Set(exclude);
  const avail = shuffle(
    (pool || []).filter((s) => s && s.question && s.corpus && !excluded.has(keyOf(s)))
  );
  if (!avail.length || count <= 0) return [];
  const ordered = [...avail.filter((s) => s.pinned), ...avail.filter((s) => !s.pinned)];
  const picked = [];
  const usedCorpora = new Set();
  // Passe 1 : un chip par corpus au maximum (mixité des sujets).
  for (const s of ordered) {
    if (picked.length >= count) break;
    if (!usedCorpora.has(s.corpus)) {
      picked.push(s);
      usedCorpora.add(s.corpus);
    }
  }
  // Passe 2 : compléter si le pool est trop homogène.
  for (const s of ordered) {
    if (picked.length >= count) break;
    if (!picked.includes(s)) picked.push(s);
  }
  return picked;
}

// --- Anti-répétition (localStorage ; stockage injectable pour les tests) ---
export function readSeen(storage) {
  const st = storage || (typeof localStorage !== "undefined" ? localStorage : null);
  if (!st) return [];
  try {
    const v = JSON.parse(st.getItem(SEEN_KEY) || "[]");
    return Array.isArray(v) ? v : [];
  } catch {
    return [];
  }
}

export function markSeen(suggs, storage) {
  const st = storage || (typeof localStorage !== "undefined" ? localStorage : null);
  if (!st) return;
  const seen = readSeen(st);
  for (const s of suggs || []) {
    const k = keyOf(s);
    if (k && !seen.includes(k)) seen.push(k);
  }
  try {
    st.setItem(SEEN_KEY, JSON.stringify(seen.slice(-SEEN_MAX)));
  } catch {}
}

// --- Rendu des chips ---
export function renderChips(box, suggestions, { onPick, onMore } = {}) {
  box.innerHTML = "";
  box.classList.add("suggest-chips");
  for (const s of suggestions) {
    const b = document.createElement("button");
    b.type = "button";
    b.className = "suggest-chip";
    b.textContent = s.question;
    b.title = s.corpus;
    b.addEventListener("click", () => onPick && onPick(s));
    box.appendChild(b);
  }
  if (onMore) {
    const more = document.createElement("button");
    more.type = "button";
    more.className = "suggest-chip suggest-chip-more";
    more.textContent = "Autre…";
    more.addEventListener("click", onMore);
    box.appendChild(more);
  }
}

// Remplit une boîte de chips : tirage initial puis « Autre » => nouveau
// tirage hors questions déjà vues. onPick reçoit la suggestion cliquée.
export async function fillChips(box, onPick) {
  let pool = [];
  try {
    pool = await loadPool();
  } catch {
    pool = [];
  }
  if (!pool.length) {
    box.innerHTML = "";
    return;
  }
  const draw = () => {
    const seen = readSeen();
    let picked = drawSuggestions(pool, { count: 3, exclude: seen });
    if (!picked.length) picked = drawSuggestions(pool, { count: 3 }); // tout vu : on recycle
    markSeen(picked);
    renderChips(box, picked, {
      onPick,
      onMore: () => {
        const again = drawSuggestions(pool, {
          count: 3,
          exclude: [...readSeen(), ...picked.map(keyOf)],
        });
        const final = again.length ? again : drawSuggestions(pool, { count: 3 });
        markSeen(final);
        renderChips(box, final, { onPick, onMore: draw });
      },
    });
  };
  draw();
}
