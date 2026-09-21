// ============================================================================
// Cache IndexedDB ciblé (phase 3) : arborescence /tree + petits aperçus.
//
// - Arbre : réponses ProjectsAPI.tree en cache 60 s, en stale-while-revalidate
//   (réponse immédiate, rafraîchissement discret en tâche de fond).
// - Aperçus : ProjectsAPI.file / fileBlob ≤ 2 Mo, 10 min.
// - Invalidation : upload, suppression de fichier, suppression de projet
//   (invalidateProjectCache, aussi appelé par le hook onDone des tours
//   agent dans agents.js — l'arbre re-rendu après un tour est toujours
//   frais) ; changement de projet inutile (clés namespacées par projet).
// - Pas de localStorage, pas de gros fichiers côté navigateur : le serveur
//   assure la vitesse de l'agent, le client la vitesse perçue.
// - Dégradation : si IndexedDB est indisponible ou en échec, l'overlay
//   mémoire prend le relais puis tout repasse par le réseau — le cache ne
//   casse jamais l'UI.
// ============================================================================
import { ProjectsAPI } from "./projects.js";

export const IDB_TREE_TTL_MS = 60_000;
export const IDB_FILE_TTL_MS = 10 * 60_000;
export const IDB_FILE_MAX_BYTES = 2 << 20; // 2 Mo

const DB_NAME = "cetas-cache-v1";
const STORE = "kv";

// --- overlay mémoire --------------------------------------------------------
// Devant IndexedDB : absorbe les appels rapprochés (le put IDB est async) et
// évite le disque pour les données très récentes. Borné, FIFO.
const MEM_MAX = 300;
const mem = new Map(); // key -> { t, v }

function memGet(key, ttl) {
  const e = mem.get(key);
  if (!e) return undefined;
  if (Date.now() - e.t >= ttl) {
    mem.delete(key);
    return undefined;
  }
  return e.v;
}

function memSet(key, v) {
  if (mem.size >= MEM_MAX) {
    const first = mem.keys().next();
    if (!first.done) mem.delete(first.value);
  }
  mem.set(key, { t: Date.now(), v });
}

function memDelPrefix(prefix) {
  for (const k of mem.keys()) if (k.startsWith(prefix)) mem.delete(k);
}

// --- couche IndexedDB (tout échec => bypass silencieux) --------------------

let dbPromise = null;

function openDb() {
  if (dbPromise) return dbPromise;
  dbPromise = new Promise((resolve, reject) => {
    const idb = globalThis.indexedDB;
    if (!idb || typeof idb.open !== "function") {
      reject(new Error("indexedDB indisponible"));
      return;
    }
    let req;
    try {
      req = idb.open(DB_NAME, 1);
    } catch (e) {
      reject(e);
      return;
    }
    req.onupgradeneeded = () => {
      try {
        req.result.createObjectStore(STORE);
      } catch {
        /* déjà existant */
      }
    };
    req.onsuccess = () => resolve(req.result);
    req.onerror = () => reject(req.error || new Error("open idb"));
    req.onblocked = () => reject(new Error("open idb bloqué"));
  });
  // Échec d'ouverture : on retentera au prochain appel.
  dbPromise.catch(() => {
    dbPromise = null;
  });
  return dbPromise;
}

function reqAsPromise(req) {
  return new Promise((resolve) => {
    req.onsuccess = () => resolve({ ok: true, value: req.result });
    req.onerror = () => resolve({ ok: false, value: undefined });
  });
}

async function kvGet(key) {
  try {
    const db = await openDb();
    const store = db.transaction(STORE, "readonly").objectStore(STORE);
    const r = await reqAsPromise(store.get(key));
    return r.ok ? r.value : undefined;
  } catch {
    return undefined;
  }
}

async function kvSet(key, value) {
  try {
    const db = await openDb();
    const store = db.transaction(STORE, "readwrite").objectStore(STORE);
    await reqAsPromise(store.put(value, key));
  } catch {
    /* bypass */
  }
}

async function kvDelPrefix(prefix) {
  try {
    const db = await openDb();
    const store = db.transaction(STORE, "readonly").objectStore(STORE);
    const r = await reqAsPromise(store.getAllKeys());
    if (!r.ok || !Array.isArray(r.value)) return;
    const targets = r.value.filter((k) => typeof k === "string" && k.startsWith(prefix));
    for (const k of targets) {
      try {
        const wstore = db.transaction(STORE, "readwrite").objectStore(STORE);
        await reqAsPromise(wstore.delete(k));
      } catch {
        /* une clé peut échouer : on continue */
      }
    }
  } catch {
    /* bypass */
  }
}

// --- clés -------------------------------------------------------------------

const treeKey = (id, path, depth) => `tree:${id}:${path || ""}:${depth}`;
const fileKey = (id, path) => `file:${id}:${path}`;
const blobKey = (id, path) => `blob:${id}:${path}`;

// --- arbre : stale-while-revalidate -----------------------------------------

const origTree = ProjectsAPI.tree.bind(ProjectsAPI);

function refreshTreeInBackground(key, id, path, depth) {
  origTree(id, path, depth).then(
    (v) => {
      memSet(key, v);
      kvSet(key, { t: Date.now(), v });
    },
    () => {}
  );
}

ProjectsAPI.tree = async function (id, path = "", depth = 1) {
  const key = treeKey(id, path, depth);
  const now = Date.now();
  const m = memGet(key, IDB_TREE_TTL_MS);
  if (m !== undefined) {
    refreshTreeInBackground(key, id, path, depth);
    return m;
  }
  const hit = await kvGet(key);
  if (hit && typeof hit.t === "number" && now - hit.t < IDB_TREE_TTL_MS) {
    memSet(key, hit.v);
    refreshTreeInBackground(key, id, path, depth);
    return hit.v;
  }
  const v = await origTree(id, path, depth);
  memSet(key, v);
  // Attendu : sur un miss on a déjà payé le réseau, le put IDB (~1 ms) rend
  // l'invariant « après retour, l'entrée est durablement cachée » vrai et
  // évite qu'un put en vol atterrisse après une invalidation.
  await kvSet(key, { t: now, v });
  return v;
};

// --- aperçus ----------------------------------------------------------------

const origFile = ProjectsAPI.file.bind(ProjectsAPI);

ProjectsAPI.file = async function (id, path) {
  const key = fileKey(id, path);
  const m = memGet(key, IDB_FILE_TTL_MS);
  if (m !== undefined) return m;
  const hit = await kvGet(key);
  if (hit && typeof hit.t === "number" && Date.now() - hit.t < IDB_FILE_TTL_MS) {
    memSet(key, hit.v);
    return hit.v;
  }
  const v = await origFile(id, path);
  try {
    if (JSON.stringify(v).length <= IDB_FILE_MAX_BYTES) {
      memSet(key, v);
      await kvSet(key, { t: Date.now(), v });
    }
  } catch {
    /* non sérialisable : pas de cache */
  }
  return v;
};

const origFileBlob = ProjectsAPI.fileBlob.bind(ProjectsAPI);

ProjectsAPI.fileBlob = async function (id, path) {
  const key = blobKey(id, path);
  const m = memGet(key, IDB_FILE_TTL_MS);
  if (m instanceof Blob) return m;
  const hit = await kvGet(key);
  if (hit && hit.b instanceof Blob && Date.now() - hit.t < IDB_FILE_TTL_MS) {
    memSet(key, hit.b);
    return hit.b;
  }
  const b = await origFileBlob(id, path);
  if (b instanceof Blob && b.size <= IDB_FILE_MAX_BYTES) {
    memSet(key, b);
    await kvSet(key, { t: Date.now(), b });
  }
  return b;
};

// --- invalidation sur mutations locales -------------------------------------

// invalidateProjectCache purge l'arbre et les aperçus d'un projet (mémoire +
// IndexedDB). Ne rejette jamais : le cache reste un accélérateur optionnel.
// Exporté pour le hook onDone des tours agent (agents.js).
export function invalidateProjectCache(id) {
  memDelPrefix(`tree:${id}:`);
  memDelPrefix(`file:${id}:`);
  memDelPrefix(`blob:${id}:`);
  return Promise.all([
    kvDelPrefix(`tree:${id}:`),
    kvDelPrefix(`file:${id}:`),
    kvDelPrefix(`blob:${id}:`),
  ]).then(
    () => undefined,
    () => undefined
  );
}

const origUpload = ProjectsAPI.upload.bind(ProjectsAPI);
ProjectsAPI.upload = async function (id, formData) {
  const r = await origUpload(id, formData);
  await invalidateProjectCache(id);
  return r;
};

const origRemoveFile = ProjectsAPI.removeFile.bind(ProjectsAPI);
ProjectsAPI.removeFile = async function (id, path) {
  const r = await origRemoveFile(id, path);
  await invalidateProjectCache(id);
  return r;
};

const origRemove = ProjectsAPI.remove.bind(ProjectsAPI);
ProjectsAPI.remove = async function (id) {
  const r = await origRemove(id);
  await invalidateProjectCache(id);
  return r;
};
