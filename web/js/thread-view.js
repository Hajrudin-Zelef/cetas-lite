import { api, getToken, readSSE } from "./api.js";
import { appendLinkified, createMarkdownRenderer } from "./markdown.js";
import {
  appendReasoningPanel,
  beginReasonTurn,
  dropCurrentReasonTurn,
  finalizeReasonTurn,
  finishReasoning,
  reopenReasonPanel,
  resetReasonPanel,
  restoreReasonSnapshot,
  sealReasonTurn,
  setReasonModel,
  updateReasonUsage,
} from "./reasoning-panel.js";
import { setTurnStats } from "./turn-tokens.js";
import { getFeaturePref } from "./model-select.js";
import { getAgenticStyle, AGENTIC_STYLE_OPENCODE, AGENTIC_STYLE_CODEX } from "./agentic-style.js";

// Loader rond "Marex" pendant la generation — repris trait pour trait du
// CETAS complet (marexcode) : point central lumineux + anneau (arc visible)
// + 4 satellites en orbite. La couleur vient de currentColor, heritee de
// .marex-loader (var(--accent)) : suit le theme et la palette selectionnee.
function buildMarexLoader() {
  const root = el("div", "marex-loader");
  const ring = el("div", "loader");
  ring.appendChild(el("div", "loader__inner"));
  const orbit = el("div", "loader__orbit");
  for (let i = 0; i < 4; i++) orbit.appendChild(el("div", "loader__dot"));
  ring.appendChild(orbit);
  root.appendChild(ring);
  return root;
}

// Défilement calqué sur la vue de référence (ChatView.tsx) :
// seuil d'épinglage, échantillonnage des gestes lecteur, borne DOM.
const FOLLOW_THRESHOLD = 24; // px : en-deçà du bas = "épinglé"
const SCROLL_SAMPLE_MS = 500; // debounce d'échantillonnage du scroll lecteur
const MAX_LOG_NODES = 500; // au-delà, les anciens nœuds sont élagués
const PRUNE_KEEP_TAIL = 120; // la queue active n'est jamais élaguée

// Moteur de rendu streaming accelere (technique Marexcode) : coalesce les
// deltas sur une seule frame + ne re-rend que les blocs markdown modifies.
const mdRenderer = createMarkdownRenderer();

function rafTick(fn) {
  if (typeof requestAnimationFrame === "function") return requestAnimationFrame(fn);
  return setTimeout(fn, 16);
}

function cancelRafTick(id) {
  if (typeof cancelAnimationFrame === "function") cancelAnimationFrame(id);
  else clearTimeout(id);
}

// Vue Agents sur mobile/tablette (≤1024px, même seuil que le CSS) : le
// raisonnement s'affiche EN LIGNE, au-dessus de la réponse, dans un bloc
// repliable. Sur desktop le panneau latéral reste le seul affichage.
function isMobileViewport() {
  return (
    typeof window !== "undefined" &&
    typeof window.matchMedia === "function" &&
    window.matchMedia("(max-width: 1024px)").matches
  );
}

export function el(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

function summarizeArgs(args) {
  if (!args) return "";
  return args.command || args.file_path || args.pattern || args.query || "";
}

// Icône sobre par outil pour les lignes d'outils (style Harness, phase 1
// refonte) : fini l'étincelle ✨, chaque famille d'outils a son glyphe.
// Façon opencode TUI (→ ✱ $) et Codex (•).
function toolGlyph(name) {
  switch (name) {
    case "Bash":
    case "RunScript":
    case "Curl":
      return "$";
    case "Grep":
    case "Glob":
    case "Tree":
      return "✱";
    case "TodoWrite":
      return "☑";
    case "web_search":
    case "web_fetch":
      return "◈";
    default:
      return "→";
  }
}

// Phase 2 refonte — cartes d'approbation premium : question d'action par
// outil (le serveur n'envoie pas de "raison du modèle", la headline décrit
// donc l'action demandée de façon explicite).
function approvalHeadline(kind, tool) {
  if (kind === "plan") return "Valider ce plan ?";
  switch (tool) {
    case "Bash":
    case "RunScript":
      return "Exécuter cette commande ?";
    case "Write":
      return "Écrire ce fichier ?";
    case "Edit":
      return "Modifier ce fichier ?";
    case "Mkdir":
      return "Créer ce dossier ?";
    case "Mv":
      return "Déplacer ce fichier ?";
    case "Curl":
      return "Appeler cette URL ?";
    case "Sed":
    case "Awk":
      return "Appliquer cette transformation ?";
    case "GitHubRepoCreate":
      return "Créer ce dépôt GitHub ?";
    case "GitHubIssueCreate":
      return "Créer cette issue GitHub ?";
    case "GitHubIssueComment":
      return "Publier ce commentaire GitHub ?";
    case "GitHubPRCreate":
      return "Créer cette pull request ?";
    case "GitHubPRMerge":
      return "Fusionner cette pull request ?";
    default:
      return "Autoriser cet outil ?";
  }
}

// Phase 2 refonte : détail d'approbation structuré par outil — jamais de
// JSON brut. Lignes clé/valeur et blocs mono bornés.
function approvalDetail(tool, args) {
  const box = el("div", "ap2-detail");
  const s = (v) => (v === undefined || v === null ? "" : String(v));
  const row = (label, value) => {
    if (!value) return;
    const r = el("div", "ap2-kv");
    r.appendChild(el("span", "ap2-k", label));
    r.appendChild(el("span", "ap2-v", value));
    box.appendChild(r);
  };
  const code = (label, text, maxLines, tone) => {
    if (!text) return;
    const lines = text.split("\n");
    const shown = maxLines && lines.length > maxLines
      ? lines.slice(0, maxLines).join("\n") + "\n… (" + lines.length + " lignes au total)"
      : text;
    const w = el("div", "ap2-codeblock");
    if (label) w.appendChild(el("div", "ap2-k", label));
    // tone "old"/"new" : teinte fade rouge/vert (carte Edit).
    const pre = el("pre", "ap2-code" + (tone === "old" ? " is-old" : tone === "new" ? " is-new" : ""));
    pre.textContent = shown;
    w.appendChild(pre);
    box.appendChild(w);
  };
  const repo = [s(args.owner), s(args.repo)].filter(Boolean).join("/");
  switch (tool) {
    case "Bash":
      code("Commande", s(args.command));
      if (args.timeout) row("Timeout", s(args.timeout) + " s");
      break;
    case "RunScript":
      row("Langage", s(args.language));
      code("Script", s(args.code), 30);
      break;
    case "Write":
      row("Fichier", s(args.file_path));
      code("Contenu", s(args.content), 30);
      break;
    case "Edit":
      row("Fichier", s(args.file_path));
      code("Remplacer", s(args.old), 12, "old");
      code("Par", s(args.new), 12, "new");
      break;
    case "Mkdir":
      row("Dossier", s(args.path));
      break;
    case "Mv":
      row("Source", s(args.src));
      row("Destination", s(args.dst));
      if (args.overwrite === true || s(args.overwrite) === "true") row("Écrasement", "autorisé");
      break;
    case "Sed":
      row("Expression", s(args.expression));
      if (args.file) row("Fichier", s(args.file));
      if (args.in_place === true || s(args.in_place) === "true") row("Mode", "modification sur place");
      break;
    case "Awk":
      code("Programme", s(args.program), 20);
      if (args.file) row("Fichier", s(args.file));
      break;
    case "Curl":
      row("Méthode", s(args.method) || "GET");
      row("URL", s(args.url));
      if (args.body) code("Corps", s(args.body), 15);
      break;
    case "GitHubRepoCreate":
      row("Dépôt", s(args.name));
      if (args.description) row("Description", s(args.description));
      row("Visibilité", s(args.private) === "false" ? "public" : "privé");
      break;
    case "GitHubIssueCreate":
      if (repo) row("Dépôt", repo);
      row("Titre", s(args.title));
      if (args.body) code("Contenu", s(args.body), 20);
      break;
    case "GitHubIssueComment":
      if (repo) row("Dépôt", repo);
      if (args.number) row("Numéro", s(args.number));
      code("Commentaire", s(args.body), 20);
      break;
    case "GitHubPRCreate":
      if (repo) row("Dépôt", repo);
      row("Titre", s(args.title));
      if (args.head) row("Branche", s(args.head) + (args.base ? " → " + s(args.base) : ""));
      if (args.body) code("Description", s(args.body), 20);
      break;
    case "GitHubPRMerge":
      if (repo) row("Dépôt", repo);
      if (args.number) row("Numéro", s(args.number));
      row("Méthode", s(args.merge_method) || "merge");
      break;
    default: {
      // Repli générique : lignes clé/valeur (jamais de JSON brut).
      const keys = Object.keys(args || {}).filter(
        (k) => args[k] !== undefined && args[k] !== null && typeof args[k] !== "object"
      );
      for (const k of keys.slice(0, 12)) {
        const v = s(args[k]);
        row(k, v.length > 300 ? v.slice(0, 300) + "…" : v);
      }
      break;
    }
  }
  return box;
}

// Libellé de statut façon OpenCode selon l'outil : "Writing command"
// quand l'agent écrit une commande, "Preparing edit" quand il prépare une
// modification. null = aucun statut particulier.
function agentStatusForTool(name, args) {
  if (name === "Bash" || name === "RunScript") return "Writing command";
  if (name === "Edit" || name === "Write") return "Preparing edit";
  // Phase 4 : progression en direct et en français pendant les lectures —
  // "Lecture de X…" au lieu d'attendre le bloc final.
  if (name === "Read" && args && args.file_path) return "Lecture de " + args.file_path + "…";
  if (name === "Ls") return "Liste des fichiers…";
  if (name === "Tree") return "Exploration de l'arborescence…";
  if (name === "Grep" && args && args.pattern) return "Recherche de « " + args.pattern + " »…";
  if (name === "Glob" && args && args.pattern) return "Recherche de " + args.pattern + "…";
  return null;
}

// Format compact façon Harness : 950 -> "950", 13700 -> "13,7k".
function fmtK(n) {
  n = Math.round(n || 0);
  if (n < 1000) return String(n);
  const v = n / 1000;
  const txt = v >= 100 ? String(Math.round(v)) : v.toFixed(1).replace(".", ",").replace(/,0$/, "");
  return txt + "k";
}

// Indicateur visuel "recherche web en cours" : spinner + libellé simple.
function renderSearchStatus(label) {
  const row = el("div", "search-status");
  row.appendChild(el("span", "search-spinner"));
  const txt = el("span", "search-status-text");
  txt.textContent = "Recherche web en cours" + (label ? " : " + label : "") + "…";
  row.appendChild(txt);
  return row;
}

function searchLabel(name, args) {
  args = args || {};
  if (name === "web_fetch" && args.url) return String(args.url);
  if (args.query) return String(args.query);
  return "";
}

// Panneau "Sources" : cartes des pages consultees (style de la reference).
function renderSources(sources) {
  const VISIBLE = 4;
  const block = el("div", "citations-block");
  block.appendChild(el("div", "citations-title", "Sources"));
  const list = el("ul", "citations-list");
  sources.forEach((s, i) => {
    const url = typeof s === "string" ? s : s.url || "";
    const title = typeof s === "string" ? "" : s.title || "";
    let host = url;
    try {
      host = new URL(url).hostname.replace(/^www\./, "");
    } catch (e) { /* url non parsable : on affiche telle quelle */ }
    const li = el("li");
    if (i >= VISIBLE) li.classList.add("citation-hidden");
    // F3 : on ne rend cliquable que les URL http(s) — jamais
    // javascript:, data:, file:… (XSS via href).
    const safeLink = /^https?:\/\//i.test(url);
    const a = el(safeLink ? "a" : "span", "citation-card" + (safeLink ? "" : " citation-nolink"));
    if (safeLink) {
      a.href = url;
      a.target = "_blank";
      a.rel = "noopener noreferrer";
    }
    a.title = url;
    const head = el("div", "citation-card-head");
    const fav = el("img", "citation-favicon");
    fav.src = "https://www.google.com/s2/favicons?domain=" + encodeURIComponent(host) + "&sz=32";
    fav.alt = "";
    fav.loading = "lazy";
    fav.onerror = () => fav.remove();
    head.appendChild(fav);
    head.appendChild(el("span", "citation-domain", host));
    head.appendChild(el("span", "citation-num", "[" + (i + 1) + "]"));
    a.appendChild(head);
    a.appendChild(el("div", "citation-card-title", title || host));
    li.appendChild(a);
    list.appendChild(li);
  });
  if (sources.length > VISIBLE) {
    const more = el("li");
    const btn = el("div", "citation-more", "+" + (sources.length - VISIBLE) + " sources");
    btn.addEventListener("click", () => {
      list.querySelectorAll(".citation-hidden").forEach((x) => x.classList.remove("citation-hidden"));
      more.remove();
    });
    more.appendChild(btn);
    list.appendChild(more);
  }
  block.appendChild(list);
  return block;
}


function renderTodos(todos) {
  const list = el("ul", "chat-todo-list");
  for (const t of todos) {
    const status = t.status || "pending";
    list.appendChild(el("li", "chat-todo-item todo-" + status, t.content || ""));
  }
  return list;
}

// "+ Thought: 227ms" / "+ Thought: 2.9s" : format compact façon OpenCode.
function fmtThoughtMs(ms) {
  if (ms < 1000) return Math.round(ms) + "ms";
  return (ms / 1000).toFixed(1) + "s";
}

// Lecture de fichier : 10 premières lignes + bouton Expand/Collapse.
function renderReadMore(text, totalLines) {
  const wrap = el("div", "tool-result-wrap");
  const lines = String(text).split("\n");
  const preShort = el("pre", "tool-result");
  appendLinkified(preShort, lines.slice(0, 10).join("\n"));
  const preFull = el("pre", "tool-result");
  appendLinkified(preFull, String(text));
  preFull.hidden = true;
  const btn = el("button", "tool-expand-btn");
  const setLabel = (expanded) => {
    btn.textContent = expanded ? "▲ Collapse" : "▼ Expand · " + totalLines + " lignes";
  };
  setLabel(false);
  btn.addEventListener("click", () => {
    const expanded = preFull.hidden;
    preFull.hidden = !expanded;
    preShort.hidden = expanded;
    setLabel(expanded);
  });
  wrap.appendChild(preShort);
  wrap.appendChild(preFull);
  wrap.appendChild(btn);
  return wrap;
}

// Sortie bornée façon Codex (phase 1 refonte) : 5 premières + 5 dernières
// lignes, ellipse "… +N lignes" au milieu, déplié complet au clic.
// Mêmes classes que renderReadMore (tool-result-wrap, tool-expand-btn)
// pour ne pas casser les harnais de test.
function renderOutputMore(text) {
  const lines = String(text).split("\n");
  const total = lines.length;
  const wrap = el("div", "tool-result-wrap");
  const preShort = el("pre", "tool-result");
  appendLinkified(preShort, lines.slice(0, 5).join("\n"));
  const dots = el("div", "tool-result-ellipsis", "… +" + (total - 10) + " lignes");
  const preTail = el("pre", "tool-result tool-result-tail");
  appendLinkified(preTail, lines.slice(-5).join("\n"));
  const preFull = el("pre", "tool-result");
  appendLinkified(preFull, String(text));
  preFull.hidden = true;
  const btn = el("button", "tool-expand-btn");
  const setLabel = (expanded) => {
    btn.textContent = expanded ? "▲ Réduire" : "▼ Afficher tout · " + total + " lignes";
  };
  setLabel(false);
  btn.addEventListener("click", () => {
    const expanded = preFull.hidden;
    preFull.hidden = !expanded;
    preShort.hidden = expanded;
    dots.hidden = expanded;
    preTail.hidden = expanded;
    setLabel(expanded);
  });
  wrap.appendChild(preShort);
  wrap.appendChild(dots);
  wrap.appendChild(preTail);
  wrap.appendChild(preFull);
  wrap.appendChild(btn);
  return wrap;
}

// --- RAG : rendu dédié rag_search / rag_read ---
// Sorties compactes et structurées (pas de gros bloc monospacé brut) :
// - rag_search : <details> "Base locale · N passages", un item par passage
//   (titre, chemin, extrait tronqué) ;
// - rag_read : <details> "Lecture · chemin · Lx–Ly", aperçu borné +
//   déplié complet (mêmes classes que renderReadMore pour les harnais).
function parseRagSearch(text) {
  const lines = String(text).split("\n");
  const head = lines[0] || "";
  const m = head.match(/Base locale \((\d+) passage/);
  const count = m ? parseInt(m[1], 10) : 0;
  const hits = [];
  let cur = null;
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i];
    const hm = line.match(/^\[(\d+)\]\s*(.+?)\s*[—–-]\s*(.+)$/);
    if (hm) {
      if (cur) hits.push(cur);
      cur = { ref: hm[1], title: hm[2].trim(), path: hm[3].trim(), snippet: [] };
    } else if (cur && line.trim() !== "") {
      cur.snippet.push(line);
    } else if (cur && line.trim() === "" && cur.snippet.length) {
      hits.push(cur);
      cur = null;
    }
  }
  if (cur) hits.push(cur);
  return { count: count || hits.length, hits };
}

function renderRagResult(toolName, text, args) {
  const wrap = el("div", "tool-result-wrap rag-result-wrap");
  if (toolName === "rag_search") {
    const { count, hits } = parseRagSearch(text);
    if (!hits.length) {
      const pre = el("pre", "tool-result");
      appendLinkified(pre, text);
      wrap.appendChild(pre);
      return wrap;
    }
    const det = el("details", "rag-result");
    det.open = true;
    const sum = el("summary", "rag-summary");
    sum.appendChild(el("span", "rag-summary-title", "Base locale"));
    sum.appendChild(el("span", "rag-summary-count", " · " + count + " passage" + (count > 1 ? "s" : "")));
    det.appendChild(sum);
    const list = el("div", "rag-hit-list");
    for (const h of hits) {
      const item = el("div", "rag-hit");
      const head = el("div", "rag-hit-head");
      head.appendChild(el("span", "rag-hit-ref", "[" + h.ref + "] "));
      head.appendChild(el("span", "rag-hit-title", h.title));
      item.appendChild(head);
      item.appendChild(el("div", "rag-hit-path", h.path));
      const snip = h.snippet.join(" ").trim();
      if (snip) item.appendChild(el("div", "rag-hit-snippet", snip.length > 400 ? snip.slice(0, 400) + "…" : snip));
      list.appendChild(item);
    }
    det.appendChild(list);
    wrap.appendChild(det);
    return wrap;
  }
  // rag_read : lignes numérotées "123\t…".
  const lines = String(text).split("\n").filter((l) => l.trim() !== "");
  const nums = [];
  for (const l of lines) {
    const lm = l.match(/^(\d+)\t/);
    if (lm) nums.push(parseInt(lm[1], 10));
  }
  const range = nums.length ? "L" + nums[0] + "–L" + nums[nums.length - 1] : "";
  const path = (args && args.path) || "";
  const det = el("details", "rag-result");
  const sum = el("summary", "rag-summary");
  sum.appendChild(el("span", "rag-summary-title", "Lecture"));
  if (path) sum.appendChild(el("span", "rag-summary-path", " · " + path));
  if (range) sum.appendChild(el("span", "rag-summary-count", " · " + range));
  det.appendChild(sum);
  const PREVIEW = 20;
  const preShort = el("pre", "tool-result");
  appendLinkified(preShort, lines.slice(0, PREVIEW).join("\n"));
  const preFull = el("pre", "tool-result");
  appendLinkified(preFull, String(text));
  preFull.hidden = true;
  const btn = el("button", "tool-expand-btn");
  const setLabel = (expanded) => {
    btn.textContent = expanded ? "▲ Réduire" : "▼ Afficher tout · " + lines.length + " lignes";
  };
  setLabel(false);
  btn.addEventListener("click", () => {
    const expanded = preFull.hidden;
    preFull.hidden = !expanded;
    preShort.hidden = expanded;
    setLabel(expanded);
  });
  const body = el("div", "rag-read-body");
  body.appendChild(preShort);
  body.appendChild(preFull);
  if (lines.length > PREVIEW) body.appendChild(btn);
  det.appendChild(body);
  wrap.appendChild(det);
  return wrap;
}

// Phase 3 refonte : verbe d'action affiché dans l'en-tête du diff
// (vue Agents uniquement).
export function diffVerb(toolName) {
  switch (toolName) {
    case "Write":
      return "Écrit";
    case "Edit":
    case "Sed":
      return "Modifié";
    default:
      return "Diff";
  }
}

// Lignes du diff affichées avant "… N lignes de plus" (aperçu borné,
// phase 3). Le serveur borne déjà les diffs à 300 lignes ; l'aperçu
// n'affiche que le début.
const DIFF_PREVIEW_LINES = 24;

export function renderDiff(lines, filePath, verb) {  const wrap = el("div", "tool-diff");
  let adds = 0, dels = 0;
  for (const line of lines) {
    if (line.kind === "+") adds++;
    else if (line.kind === "-") dels++;
  }
  const head = el("div", "diff-head");
  if (verb) {
    // Vue Agents, phase 3 : verbe + chemin + "+n -m" colorés.
    head.appendChild(el("span", "diff-verb", verb));
    head.appendChild(el("span", "diff-file", filePath ? String(filePath) : "modification"));
    const stats = el("span", "diff-stats");
    stats.appendChild(el("span", "diff-add-n", "+" + adds));
    stats.appendChild(document.createTextNode(" "));
    stats.appendChild(el("span", "diff-del-n", "-" + dels));
    head.appendChild(stats);
  } else {
    // Chat général : en-tête historique strictement inchangé.
    head.appendChild(el("span", "diff-file", filePath ? String(filePath) : "modification"));
    head.appendChild(el("span", "diff-stats", "+" + adds + " / -" + dels));
  }
  wrap.appendChild(head);
  const body = el("div", "diff-body");
  let oldN = 0, newN = 0;
  const shown = lines.slice(0, DIFF_PREVIEW_LINES);
  for (const line of shown) {
    const kind = line.kind === "+" ? "diff-add" : line.kind === "-" ? "diff-del" : "diff-ctx";
    const row = el("div", "diff-row " + kind);
    let gutter = "";
    if (line.kind === "-") { oldN++; gutter = String(oldN); }
    else if (line.kind === "+") { newN++; gutter = String(newN); }
    row.appendChild(el("span", "diff-gutter", gutter));
    row.appendChild(el("span", "diff-sign", line.kind === "…" ? "…" : line.kind || " "));
    row.appendChild(el("span", "diff-code", line.text != null ? String(line.text) : ""));
    body.appendChild(row);
  }
  if (lines.length > shown.length) {
    body.appendChild(el("div", "diff-row diff-more",
      "… " + (lines.length - shown.length) + " lignes de plus"));
  }
  wrap.appendChild(body);
  return wrap;
}

// Reproduction web du TUI opencode (internal/tui/components/chat/message.go)
// : blocs à bordure gauche épaisse, en-têtes "Nom: paramètres", résultats
// bornés à 10 lignes, statuts exacts, erreurs en rouge, pied "modèle (durée)".

// Libellé d'action affiché dans l'en-tête de l'outil PENDANT son exécution
// (équivalent du "Working..." / "Building command..." du TUI).
function ocToolAction(name) {
  switch (name) {
    case "Bash":
    case "RunScript":
      return "Building command...";
    case "Edit":
      return "Preparing edit...";
    case "Write":
      return "Preparing write...";
    case "Read":
    case "Cat":
      return "Reading file...";
    case "Grep":
      return "Searching content...";
    case "Glob":
      return "Finding files...";
    case "Ls":
      return "Listing directory...";
    case "TodoWrite":
      return "Updating todos...";
    case "web_search":
      return "Searching web...";
    case "web_fetch":
      return "Fetching page...";
    default:
      return "Working...";
  }
}

// Paramètre principal façon TUI (renderParams) : la valeur la plus
// parlante de l'outil, tronquée à ~120 caractères.
function ocToolParams(name, args) {
  args = args || {};
  const val =
    args.file_path || args.command || args.path || args.pattern ||
    args.query || args.url || args.text || args.diff || args.content;
  if (val === undefined || val === null) return "";
  let s = String(val);
  const firstLine = s.split("\n")[0];
  s = firstLine.length < s.length ? firstLine + "…" : firstLine;
  return s.length > 120 ? s.slice(0, 119) + "…" : s;
}

// Un résultat d'outil est une erreur quand le backend le préfixe "[erreur]".
function ocIsError(result) {
  return /^\s*\[erreur\]/i.test(String(result || ""));
}

function speakable(md) {
  return String(md || "")
    .replace(/```[\s\S]*?```/g, " (bloc de code) ")
    .replace(/`([^`]+)`/g, "$1")
    .replace(/^#{1,6}\s*/gm, "")
    .replace(/[*_>#|]/g, " ")
    .replace(/\s+/g, " ")
    .trim();
}

// ThreadView : une vue de conversation complete (rendu + SSE + envoi),
// reutilisable pour le chat principal et pour chaque agent parallele.
//
// opts :
//   log, stopBtn, routeBadge, statsBadge : elements DOM
//   emptyHTML : placeholder quand le fil est vide
//   streamURL(from) : url du flux SSE
//   sendURL, stopURL, approveURL : endpoints (approveURL peut etre null)
//   getPayload(text) : corps JSON pour POST sendURL
//   actions : affiche les boutons Copier/Lire (+ Regenerer si regenerateURL)
//   regenerateURL : endpoint de regeneration (optionnel)
//   onFirstUser, onDone : callbacks optionnels
//   reasonPanel : affiche le raisonnement dans le panneau lateral (vue principale)
//   reasonHooks : { append(text, replace), finish(), reset() } — panneau custom
//     (ex. vue Agents façon Marexcode) ; prioritaire sur reasonPanel.
//   trackTokens : met a jour la ligne de tokens du composer (vue principale)
//   onEvent : callback(ev) appele pour chaque evenement SSE (ex. panneau Todos)
export class ThreadView {
  constructor(opts) {
    this.log = opts.log;
    this.stopBtn = opts.stopBtn || null;
    this.routeBadge = opts.routeBadge || null;
    this.statsBadge = opts.statsBadge || null;
    this.emptyHTML = opts.emptyHTML || "";
    this.streamURL = opts.streamURL;
    this.sendURL = opts.sendURL;
    this.stopURL = opts.stopURL;
    this.approveURL = opts.approveURL || null;
    this.getPayload = opts.getPayload || ((text) => ({ message: text }));
    this.actions = opts.actions !== false;
    this.regenerateURL = opts.regenerateURL || null;
    this.onDone = opts.onDone || null;
    this.reasonPanel = opts.reasonPanel === true;
    this.reasonHooks = opts.reasonHooks || null;
    this.trackTokens = opts.trackTokens === true;
    this.onEvent = typeof opts.onEvent === "function" ? opts.onEvent : null;
    // Module Agentic (vue Agents uniquement) : Harness = rendu actuel,
    // OpenCode = reproduction fidele du TUI OpenCode.
    this.agentic = opts.agentic === true;
    // Écho optimiste à l'envoi (vue Agents uniquement) : le message et
    // l'indicateur d'attente s'affichent sans attendre le POST. Le delta
    // "user" du serveur porte le même client_msg_id et est dédupliqué.
    // Désactivé par défaut : le chat général garde son comportement.
    this.optimisticEcho = opts.optimisticEcho === true;

    this.empty = this.log.querySelector("[data-empty]");
    this.lastSeq = 0;
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
    this.reasoningActive = false;
    this.reasonBtn = null;
    this.reasonAssistantWrapper = null;
    this.streamSpinnerEl = null;
    this.liveStatsEl = null;
    this.liveStatsTimer = 0;
    this.controller = null;
    this.generating = false;
    this.waitEl = null;
    this.waitLabel = null;
    this.waitTimer = 0;
    this.waitStart = 0;
    this.toolBoxes = new Map();
    // Compteur incrémental unique par appel d'outil (F8) : name+JSON(args)
    // seul collisionne quand le modèle appelle deux fois le même outil avec
    // les mêmes arguments. toolPending associe chaque clé de base à la file
    // des clés uniques en cours, dans l'ordre des "start".
    this.toolSeq = 0;
    this.toolPending = new Map();
    // Phase 1 refonte : checklist TodoWrite "live" (un seul bloc mis à jour
    // en place). threadGen invalide le bloc quand le fil est réinitialisé.
    this.todoChecklistEl = null;
    this.todoChecklistGen = 0;
    this.threadGen = 0;
    // Reconnexions SSE (F9) : backoff exponentiel, plafond de tentatives.
    this.sseRetries = 0;
    this.approvalCards = new Map();
    // Dernier événement reçu (pour "+ Thought: Xs" façon OpenCode).
    this.lastEventTs = 0;
    // Statut façon OpenCode ("Writing command", "Preparing edit") : pile
    // des outils en cours ayant un libellé, ligne transitoire en fin de fil.
    this.agentStatusEl = null;
    this.agentStatusStack = [];
    // Indicateur "thinking" + spinner dans le fil pendant la réflexion.
    this.thinkingEl = null;
    // Compteurs façon Harness pour la barre de statut : tours, outils,
    // tokens cumulés sur la conversation.
    this.convTurns = 0;
    this.convTools = 0;
    this.convIn = 0;
    this.convOut = 0;
    this.turnInTok = 0;
    this.turnOutTok = 0;
    // Indicateur de recherche web native (plugin provider) en cours.
    this.searchStatus = null;
    // Suivi du tour en cours (pied de message "modèle · temps · tokens").
    this.turnStartTs = 0;
    this.turnStats = null;
    this.turnRoute = null;
    this.turnElapsedMs = null;
    // État du suivi de défilement (cf. initScrollFollow) : même comportement
    // que la vue de référence (pinned-follow 24px + ledger lecteur/programmatique).
    this.pinned = true;
    this.observedTop = 0;
    this.followFrame = 0;
    this.sampleTimer = 0;
    this.newWhileUnpinned = false;
    this.toBottomBtn = null;
    this.prunedCount = 0;
    this.initScrollFollow();
  }

  clearEmpty() {
    if (this.empty && this.empty.parentNode) this.empty.remove();
    this.empty = null;
  }

  reAddEmpty() {
    this.log.innerHTML = this.emptyHTML;
    this.empty = this.log.querySelector("[data-empty]");
  }

  // ---- Défilement : même comportement que la vue de référence ----
  //
  // Principe (ChatView.tsx du fichier fourni) :
  // - "pinned" = le lecteur est à <= 24px du bas -> le nouveau contenu fait
  //   suivre la vue ; on ne re-épingle jamais sur un simple re-render
  //   (sinon les scrolls inertiels seraient "snappés" jusqu'en bas).
  // - observedTop enregistre chaque écriture programmatique de scrollTop ;
  //   un événement scroll qui n'en dévie pas n'est PAS un geste lecteur
  //   (pas de changement de propriété pinned).
  // - les gestes lecteur sont échantillonnés (500ms + scrollend).

  initScrollFollow() {
    const log = this.log;
    const onScroll = () => {
      if (this.pinned) {
        const floor = this.floorTop();
        // Livraison non-lecteur (écriture programmatique différée, clamp
        // navigateur) : échantillonner aussitôt sans toucher à pinned.
        if (Math.abs(log.scrollTop - Math.min(this.observedTop, floor)) <= 0.5) {
          this.sampleScroll();
          return;
        }
      }
      if (!this.sampleTimer) {
        this.sampleTimer = setTimeout(() => {
          this.sampleTimer = 0;
          this.sampleScroll();
        }, SCROLL_SAMPLE_MS);
      }
    };
    log.addEventListener("scroll", onScroll, { passive: true });
    log.addEventListener("scrollend", () => this.sampleScroll(), { passive: true });
    this.buildToBottomBtn();
  }

  floorTop() {
    return Math.max(0, this.log.scrollHeight - this.log.clientHeight);
  }

  // Ne change la propriété "pinned" que sur un geste lecteur réel.
  sampleScroll() {
    const top = this.log.scrollTop;
    const floor = this.floorTop();
    const readerMoved = Math.abs(top - Math.min(this.observedTop, floor)) > 0.5;
    if (!readerMoved) return; // livraison programmatique : on garde l'état
    const near = floor - top <= FOLLOW_THRESHOLD + 1;
    this.pinned = near;
    this.observedTop = top;
    if (near) this.newWhileUnpinned = false;
    this.updateToBottomBtn();
  }

  writeBottom() {
    this.log.scrollTop = this.log.scrollHeight;
    this.observedTop = this.log.scrollTop;
  }

  // Force le bas : nouveau message utilisateur, carte d'approbation,
  // rattrapage d'historique, clic sur le bouton flottant.
  toBottom() {
    if (this.followFrame) {
      cancelRafTick(this.followFrame);
      this.followFrame = 0;
    }
    this.newWhileUnpinned = false;
    this.writeBottom();
    this.pinned = true;
    this.updateToBottomBtn();
    this.maybePrune();
  }

  // Suivi coalescé sur une frame : le contenu qui grandit (streaming,
  // tool boxes, raisonnement) ne fait défiler que si épinglé.
  requestFollow() {
    if (this.followFrame) return;
    this.followFrame = rafTick(() => {
      this.followFrame = 0;
      this.maybePrune(); // borne DOM même quand le lecteur lit plus haut
      if (!this.pinned) {
        this.newWhileUnpinned = true;
        this.updateToBottomBtn();
        return;
      }
      this.writeBottom();
    });
  }

  buildToBottomBtn() {
    const parent = this.log.parentElement;
    if (!parent || parent.querySelector(":scope > .thread-to-bottom")) return;
    if (getComputedStyle(parent).position === "static") parent.style.position = "relative";
    const btn = el("button", "thread-to-bottom");
    btn.type = "button";
    btn.hidden = true;
    btn.setAttribute("aria-label", "Retour en bas");
    btn.innerHTML = "<svg viewBox=\"0 0 16 16\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M8 3v9M4.5 8.5 8 12l3.5-3.5\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"1.8\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/></svg>";
    btn.addEventListener("click", () => this.toBottom());
    parent.appendChild(btn);
    this.toBottomBtn = btn;
  }

  updateToBottomBtn() {
    const btn = this.toBottomBtn;
    if (!btn) return;
    const show = !this.pinned && this.log.scrollHeight > this.log.clientHeight + 8;
    btn.hidden = !show;
    btn.classList.toggle("has-new", this.newWhileUnpinned && show);
    btn.setAttribute(
      "aria-label",
      this.newWhileUnpinned && show ? "Retour en bas — nouveaux messages" : "Retour en bas",
    );
  }

  // ---- Borne DOM : les longues sessions restent fluides ----
  //
  // Au-delà de MAX_LOG_NODES, les nœuds les plus anciens sont retirés
  // (l'historique serveur reste intact). Protégés : cartes d'approbation
  // en attente, tool boxes en cours de stream, queue active, placeholder.
  maybePrune() {
    if (this.log.children.length <= MAX_LOG_NODES) return;
    let cut = this.log.querySelector(":scope > .history-cut");
    const tailStart = this.log.children.length - PRUNE_KEEP_TAIL;
    let i = 0;
    while (i < tailStart && this.log.children.length > MAX_LOG_NODES) {
      const n = this.log.children[i];
      if (
        n === cut ||
        n === this.empty ||
        n.classList.contains("streaming") ||
        (n.classList.contains("msg-approval") && !n.hasAttribute("data-resolved"))
      ) {
        i++;
        continue;
      }
      n.remove();
      this.prunedCount++;
      // pas d'incrément : le nœud suivant glisse à l'index i
    }
    if (this.prunedCount > 0) {
      if (!cut) {
        cut = el("div", "history-cut");
        this.log.prepend(cut);
      }
      cut.textContent =
        "… " + this.prunedCount + " message(s) précédent(s) masqué(s) pour garder l'interface fluide …";
    }
  }

  // Fige le rendu du message assistant en cours (fin de tour, outil, approbation).
  finalizeAssistant() {
    this.stopStreamSpinner();
    if (this.assistantBody) mdRenderer.finalize(this.assistantBody, this.assistantText);
  }

  resetAssistantState() {
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
    this.reasoningActive = false;
    this.reasonBtn = null;
    this.reasonAssistantWrapper = null;
    this.stopStreamSpinner();
  }

  setBusy(v) {
    this.log.setAttribute("aria-busy", v ? "true" : "false");
  }

  // Attente du premier token : le loader rond s'affiche immediatement
  // (animation 100 % CSS). Libellé façon Harness avec secondes écoulées
  // ("En cours… 12s") — seul le compteur utilise un timer JS.
  showWait() {
    if (this.waitEl) return;
    this.waitEl = el("div", "stream-waiting");
    this.waitEl.appendChild(buildMarexLoader());
    this.waitStart = Date.now();
    this.waitLabel = el("div", "wait-label", "✨ En cours… 0s");
    this.waitEl.appendChild(this.waitLabel);
    this.waitTimer = setInterval(() => {
      if (this.waitLabel)
        this.waitLabel.textContent =
          "✨ En cours… " + Math.floor((Date.now() - this.waitStart) / 1000) + "s";
    }, 1000);
    this.log.appendChild(this.waitEl);
    this.toBottom();
  }

  hideWait() {
    if (this.waitTimer) {
      clearInterval(this.waitTimer);
      this.waitTimer = 0;
    }
    this.waitLabel = null;
    if (this.waitEl) {
      this.waitEl.remove();
      this.waitEl = null;
    }
  }

  // Loader rond + stats live pendant le streaming (repris du CETAS complet) :
  // le loader est insere en tete de la bulle assistant, hors du corps
  // markdown — il est donc insensible aux re-rendus (plus de re-ancrage).
  // Les stats ("5s · 35 tok · 7.8 tok/s", monospace) suivent le texte,
  // actualisees toutes les 500 ms, et sont retirees en fin de tour.
  startStreamSpinner() {
    if (this.streamSpinnerEl) return;
    this.ensureAssistant();
    if (!this.assistant) return;
    this.streamSpinnerEl = buildMarexLoader();
    this.assistant.insertBefore(this.streamSpinnerEl, this.assistantBody);
    this.startLiveStats();
  }

  startLiveStats() {
    this.stopLiveStats();
    if (!this.assistant) return;
    const span = el("div", "gen-stats-live", "0s");
    this.assistant.appendChild(span);
    this.liveStatsEl = span;
    const start = Date.now();
    const tick = () => {
      const secs = (Date.now() - start) / 1000;
      const tok = Math.ceil((this.assistantText || "").length / 4);
      let txt = secs.toFixed(0) + "s";
      if (tok > 0) {
        txt += " · " + tok + " tok";
        if (secs >= 1) txt += " · " + (tok / secs).toFixed(1) + " tok/s";
      }
      if (this.liveStatsEl) this.liveStatsEl.textContent = txt;
    };
    tick();
    this.liveStatsTimer = setInterval(tick, 500);
    // Node (tests) : ne pas retenir la boucle d'evenements.
    if (this.liveStatsTimer && typeof this.liveStatsTimer.unref === "function") {
      this.liveStatsTimer.unref();
    }
  }

  stopLiveStats() {
    if (this.liveStatsTimer) {
      clearInterval(this.liveStatsTimer);
      this.liveStatsTimer = 0;
    }
    if (this.liveStatsEl) {
      this.liveStatsEl.remove();
      this.liveStatsEl = null;
    }
  }

  stopStreamSpinner() {
    this.stopLiveStats();
    if (this.streamSpinnerEl) {
      this.streamSpinnerEl.remove();
      this.streamSpinnerEl = null;
    }
  }

  addUser(text) {
    this.clearEmpty();
    const wrapper = el("div", "message-wrapper message-wrapper-user");
    const bubble = el("div", "message message-user");
    bubble.appendChild(el("div", "message-text", text));
    wrapper.appendChild(bubble);
    this.log.appendChild(wrapper);
    this.toBottom();
    return wrapper;
  }

  // Écho optimiste : affiche le message immédiatement à l'envoi, sans
  // attendre l'aller-retour POST. Le nœud est marqué data-optimistic
  // jusqu'à ce que le delta "user" du serveur (même client_msg_id)
  // le confirme — ou retiré si le POST échoue.
  addUserOptimistic(text) {
    const wrapper = this.addUser(text);
    wrapper.setAttribute("data-optimistic", "1");
    return wrapper;
  }

  removeOptimistic(wrapper) {
    if (wrapper && wrapper.isConnected) wrapper.remove();
    this.pendingUserId = null;
  }

  // Écho optimiste quand le POST est déjà parti (création d'agent) :
  // affiche le message tout de suite, le delta "user" rejoué via SSE
  // (même client_msg_id) le confirmera sans doublon.
  primeOptimistic(text, clientMsgId) {
    this.pendingUserId = clientMsgId;
    this.addUserOptimistic(text);
    this.showWait();
    this.setBusy(true);
  }

  newClientMsgId() {
    try {
      if (typeof crypto !== "undefined" && crypto.randomUUID) return crypto.randomUUID();
    } catch (e) {}
    return "m" + Date.now().toString(36) + Math.random().toString(36).slice(2);
  }

  ensureAssistant() {
    if (!this.assistant) {
      this.clearEmpty();
      const wrapper = el("div", "message-wrapper message-wrapper-assistant");
      this.assistant = el("div", "message message-assistant streaming");
      // Bouton "Raisonnement" en tete de la bulle assistant, comme dans le
      // CETAS complet : reconsulte le panneau une fois le raisonnement
      // termine (visible seulement en mode panneau, i.e. vue principale,
      // et quand le tour a produit du raisonnement).
      if (this.reasonPanel) {
        const btn = el("button", "reason-btn", "▸ Raisonnement");
        btn.type = "button";
        btn.hidden = true;
        btn.addEventListener("click", () => {
          const snap = wrapper._reasonSnap;
          if (snap && snap.text) restoreReasonSnapshot(snap);
          else reopenReasonPanel();
        });
        this.assistant.appendChild(btn);
        this.reasonBtn = btn;
        this.reasonAssistantWrapper = wrapper;
        // Le raisonnement a pu demarrer avant le premier contenu : l'en-tete
        // est alors visible des l'apparition de la bulle.
        if (this.reasoningActive) btn.hidden = false;
      }
      const body = el("div", "message-text");
      this.assistant.appendChild(body);
      wrapper.appendChild(this.assistant);
      this.log.appendChild(wrapper);
      this.assistantBody = body;
      this.assistantText = "";
      this.assistantStarted = false;
    }
    return this.assistant;
  }

  appendContent(text, isReplace) {
    this.ensureAssistant();
    if (!this.assistantStarted) {
      this.assistantStarted = true;
      this.hideWait();
    }
    // Fin de la phase de raisonnement : le panneau se masque tout seul.
    if (this.reasoningActive) {
      this.reasoningActive = false;
      if (this.reasonPanel) finishReasoning();
      else if (this.reasonHooks) this.reasonHooks.finish();
    }
    this.hideThinking();
    this.startStreamSpinner();
    // Streaming accelere (technique Marexcode) : update() bufferise et rend
    // au plus une fois par frame, en ne re-rendant que les blocs modifies.
    if (isReplace) {
      this.assistantText = String(text);
      mdRenderer.render(this.assistantBody, text);
    } else {
      this.assistantText += String(text);
      mdRenderer.update(this.assistantBody, this.assistantText);
    }
    this.requestFollow();
  }

  ensureReasoning() {
    if (!this.reasoningEl) {
      this.ensureAssistant();
      this.reasoningEl = el("details", "thinking-block");
      // Phase 3 refonte (vue Agents uniquement) : sur mobile le bloc
      // "Raisonnement" est replié par défaut — même seuil que le routage
      // inline du raisonnement (isMobileViewport). Desktop et chat général :
      // inchangés (déplié).
      this.reasoningEl.open = !(this.agentic && isMobileViewport());
      this.reasoningEl.appendChild(el("summary", null, "Raisonnement"));
      this.reasoningEl.appendChild(el("div", "thinking-content"));
      this.assistant.insertBefore(this.reasoningEl, this.assistant.firstChild);
    }
    return this.reasoningEl;
  }

  appendReasoning(text, isReplace) {
    // L'agent réfléchit : "thinking" + spinner visibles dans le fil.
    this.showThinking();
    // Vue Agents sur mobile : bloc "Raisonnement" repliable en tête de la
    // bulle, au-dessus de la réponse — le panneau latéral ne reçoit rien
    // (pas de doublon) et reste réservé au desktop.
    if (this.reasonHooks && isMobileViewport()) {
      this.appendReasoningInline(text, isReplace);
      return;
    }
    if (this.reasonHooks) {
      this.reasonHooks.append(text, isReplace);
      return;
    }
    if (this.reasonPanel) {
      appendReasoningPanel(text, isReplace);
      this.reasoningActive = true;
      if (this.reasonBtn) this.reasonBtn.hidden = false;
      return;
    }
    this.appendReasoningInline(text, isReplace);
  }

  // Bloc "Raisonnement" repliable (<details> natif : chevron + toggle
  // collapse/expand) inséré en tête de la bulle assistant, au-dessus du
  // contenu de la réponse.
  appendReasoningInline(text, isReplace) {
    const box = this.ensureReasoning().querySelector(".thinking-content");
    // Rendu incremental : on n'ajoute QUE le nouveau morceau au noeud texte
    // (appendData). Jamais de textContent sur tout le texte -> pas de O(n)
    // par delta quand le raisonnement est long (technique Marexcode).
    if (isReplace) {
      box.textContent = String(text);
      this.reasoningText = String(text);
    } else {
      this.reasoningText += String(text);
      let node = box.firstChild;
      if (!node || node.nodeType !== 3) {
        box.textContent = "";
        node = document.createTextNode("");
        box.appendChild(node);
      }
      node.appendData(String(text));
    }
    this.requestFollow();
  }

  removeReasoning() {
    if (this.reasonPanel) {
      dropCurrentReasonTurn();
      this.reasoningActive = false;
      if (this.reasonBtn) this.reasonBtn.hidden = true;
      return;
    }
    if (this.reasoningEl) {
      this.reasoningEl.remove();
      this.reasoningEl = null;
      this.reasoningText = "";
    }
  }

  addError(text) {
    this.hideWait();
    this.hideThinking();
    this.setBusy(false);
    this.clearEmpty();
    const wrapper = el("div", "message-wrapper message-wrapper-assistant");
    const bubble = el("div", "message message-assistant message-error");
    // Module Agentic, style OpenCode : erreurs préfixées "Error:" (TUI).
    bubble.appendChild(el("div", "message-text", (this.agenticIsOpenCode() || this.agenticIsCodex()) ? "Error: " + text : text));
    wrapper.appendChild(bubble);
    this.log.appendChild(wrapper);
    this.requestFollow();
  }

  addSystem(text) {
    this.clearEmpty();
    this.log.appendChild(el("div", "msg-system", text));
    this.requestFollow();
  }

  addActions(box, raw) {
    if (!this.actions) return;
    const wrapper = box.closest(".message-wrapper") || box;
    if (!wrapper || wrapper.querySelector(".message-btn-row")) return;
    const bar = el("div", "message-btn-row");
    const copy = el("button", "message-copy-btn", "Copier");
    copy.type = "button";
    copy.addEventListener("click", () => {
      if (!navigator.clipboard) return;
      navigator.clipboard.writeText(raw).then(() => {
        copy.textContent = "Copie";
        setTimeout(() => { copy.textContent = "Copier"; }, 1200);
      }).catch(() => {});
    });
    bar.appendChild(copy);
    if (this.regenerateURL) {
      const regen = el("button", "regen-btn", "Régénérer");
      regen.type = "button";
      regen.addEventListener("click", async () => {
        if (this.generating) return;
        try {
          await api(this.regenerateURL, { method: "POST" });
        } catch (e) {
          this.addError(e.message);
        }
      });
      bar.appendChild(regen);
    }
    if (window.speechSynthesis && getFeaturePref("tts", "system") !== "none") {
      const speak = el("button", "message-tts-btn", "Lire");
      speak.type = "button";
      speak.addEventListener("click", () => {
        window.speechSynthesis.cancel();
        const u = new SpeechSynthesisUtterance(speakable(raw));
        u.lang = document.documentElement.lang || "fr";
        window.speechSynthesis.speak(u);
      });
      bar.appendChild(speak);
    }
    wrapper.appendChild(bar);
  }

  // Module Agentic : vrai uniquement pour la vue Agents en style OpenCode.
  // Le chat general (Cetas) et le style Harness gardent le rendu actuel.
  agenticIsOpenCode() {
    return this.agentic && getAgenticStyle() === AGENTIC_STYLE_OPENCODE;
  }

  // Module Agentic : vrai uniquement pour la vue Agents en style Codex.
  agenticIsCodex() {
    return this.agentic && getAgenticStyle() === AGENTIC_STYLE_CODEX;
  }

  // Clé unique par appel d'outil : deux "start" identiques (même nom,
  // mêmes arguments) ne doivent pas partager la même carte, sinon le
  // second écrase le premier et un bloc "running" fantôme subsiste.
  toolKeyStart(ev) {
    const base = ev.name + "|" + JSON.stringify(ev.args || {});
    const key = base + "#" + (++this.toolSeq);
    let q = this.toolPending.get(base);
    if (!q) { q = []; this.toolPending.set(base, q); }
    q.push(key);
    return key;
  }
  // Chaque "end" reprend la plus ancienne clé en attente pour sa base
  // (les fins arrivent dans l'ordre des débuts pour des appels identiques).
  toolKeyEnd(ev) {
    const base = ev.name + "|" + JSON.stringify(ev.args || {});
    const q = this.toolPending.get(base);
    if (q && q.length) {
      const key = q.shift();
      if (!q.length) this.toolPending.delete(base);
      return key;
    }
    // Repli : "end" sans "start" vu (replay partiel, etc.).
    for (const k of this.toolBoxes.keys()) {
      if (k === base || k.startsWith(base + "#")) return k;
    }
    return base;
  }

  // Checklist TodoWrite "live" (phase 1 refonte) : un seul bloc dans le
  // fil, mis à jour en place à chaque appel TodoWrite au lieu d'empiler des
  // instantanés. Tâche en cours = spinner, terminée = coche animée,
  // en-tête avec compteur "n/m" et barre de progression.
  upsertTodoChecklist(todos) {
    if (!Array.isArray(todos) || !todos.length) return;
    let wrap = this.todoChecklistEl;
    if (!wrap || this.todoChecklistGen !== this.threadGen) {
      wrap = el("div", "todo-checklist");
      const head = el("div", "todo-checklist-head");
      head.appendChild(el("span", "todo-checklist-title", "Tâches"));
      const count = el("span", "todo-checklist-count", "");
      head.appendChild(count);
      const bar = el("div", "todo-checklist-bar");
      const fill = el("div", "todo-checklist-fill");
      bar.appendChild(fill);
      head.appendChild(bar);
      wrap.appendChild(head);
      const list = el("ul", "todo-checklist-items");
      wrap.appendChild(list);
      // Références directes (pas de querySelector : les harnais de test
      // utilisent un faux DOM minimal).
      wrap._todoList = list;
      wrap._todoCount = count;
      wrap._todoFill = fill;
      this.clearEmpty();
      this.log.appendChild(wrap);
      this.todoChecklistEl = wrap;
      this.todoChecklistGen = this.threadGen;
    }
    const list = wrap._todoList;
    const done = todos.filter((t) => (t.status || "pending") === "completed").length;
    wrap._todoCount.textContent = done + "/" + todos.length;
    wrap._todoFill.style.width = Math.round((done / todos.length) * 100) + "%";
    todos.forEach((t, i) => {
      const status = t.status || "pending";
      let li = list.children[i] || null;
      if (!li || li.tagName !== "LI") {
        li = el("li", "todo-checklist-item");
        li.appendChild(el("span", "todo-checklist-box"));
        li.appendChild(el("span", "todo-checklist-text"));
        list.insertBefore(li, list.children[i] || null);
      }
      li.className = "todo-checklist-item todo-" + status;
      li.children[1].textContent = t.content || "";
    });
    while (list.children.length > todos.length) {
      list.children[list.children.length - 1].remove();
    }
    this.requestFollow();
  }

  addTool(ev, gapMs) {
    // Vue Agents + style OpenCode : rendu fidele au TUI OpenCode.
    if (this.agenticIsOpenCode()) return this.addToolOpenCode(ev, gapMs);
    // Vue Agents + style Codex : rendu fidele au TUI Codex.
    if (this.agenticIsCodex()) return this.addToolCodex(ev, gapMs);
    if (ev.phase === "start") {
      const key = this.toolKeyStart(ev);
      this.clearEmpty();
      this.hideWait();
      this.hideThinking();
      // Temps de réflexion écoulé depuis l'événement précédent, façon
      // OpenCode ("+ Thought: 2.9s"). gapMs = 0 quand inconnu.
      this.maybeAddThought(gapMs || 0);
      this.convTools++;
      // Ligne compacte : "<glyphe> Read · chemin" (sans le verbiage
      // "Tool call"). Phase 1 refonte (vue Agents uniquement) : icône
      // sobre par outil (toolGlyph), fini l'étincelle ✨. Le chat général
      // (hors Agents) garde son rendu d'origine.
      const box = el("details", "msg-tool harness-tool running");
      const summary = el("summary");
      if (this.agentic) {
        summary.appendChild(el("span", "tool-glyph", toolGlyph(ev.name)));
      } else {
        summary.appendChild(el("span", "tool-spark", "✨"));
      }
      summary.appendChild(el("span", "tool-name", ev.name));
      const hint = summarizeArgs(ev.args);
      if (hint) {
        summary.appendChild(el("span", "tool-sep", "·"));
        summary.appendChild(el("span", "tool-hint", hint));
      }
      box.appendChild(summary);
      const body = el("div", "tool-body");
      if (ev.name === "TodoWrite" && ev.args && Array.isArray(ev.args.todos)) {
        if (this.agentic) {
          // Phase 1 refonte : la checklist vit dans le fil et se coche au
          // fur et à mesure, pas dans la carte d'outil.
          this.upsertTodoChecklist(ev.args.todos);
        } else {
          body.appendChild(renderTodos(ev.args.todos));
        }
      }
      if (ev.name === "web_search" || ev.name === "web_fetch") {
        body.appendChild(renderSearchStatus(searchLabel(ev.name, ev.args)));
      }
      box.appendChild(body);
      this.log.appendChild(box);
      this.toolBoxes.set(key, body);
      // Statut façon OpenCode pendant l'exécution ("Writing command",
      // "Preparing edit", "Lecture de X…").
      const stLabel = agentStatusForTool(ev.name, ev.args);
      if (stLabel) this.pushAgentStatus(key, stLabel);
      this.finalizeAssistant();
      this.resetAssistantState();
      this.requestFollow();
      return;
    }
    const key = this.toolKeyEnd(ev);
    const body = this.toolBoxes.get(key);
    if (!body) return;
    const det = body.closest("details");
    if (det) det.classList.remove("running");
    // Fin d'exécution : le statut façon OpenCode disparaît.
    this.popAgentStatus(key);
    // Fin d'une recherche web : l'indicateur disparaît.
    const isSearch = ev.name === "web_search" || ev.name === "web_fetch";
    const status = body.querySelector(".search-status");
    if (status) status.remove();
    if (this.agentic && isSearch) {
      // Vue Agents : pas de panneau Sources ni de sortie brute — seul le
      // résumé du modèle reste visible dans le fil. L'état d'erreur reste
      // signalé visuellement (bordure rouge via .is-error).
      if (ocIsError(ev.result) && det) det.classList.add("is-error");
      this.requestFollow();
      return;
    }
    if (isSearch && Array.isArray(ev.sources) && ev.sources.length) {
      body.appendChild(renderSources(ev.sources));
      this.requestFollow();
      return;
    }
    if (Array.isArray(ev.diff) && ev.diff.length) {
      body.appendChild(renderDiff(ev.diff, ev.args && ev.args.file_path, this.agentic ? diffVerb(ev.name) : null));
      // Les modifications ne sont pas masquées : le diff s'affiche déplié.
      if (det) det.open = true;
    }
    // RAG : rendu dédié compact (rag_search / rag_read), pas de gros <pre>.
    if ((ev.name === "rag_search" || ev.name === "rag_read") && typeof ev.result === "string" && ev.result) {
      body.appendChild(renderRagResult(ev.name, ev.result, ev.args));
      this.requestFollow();
      return;
    }
    if (typeof ev.result === "string" && ev.result) {
      if (this.agentic) {
        // Vue Agents, phase 1 refonte : sorties longues en 5 premières +
        // 5 dernières lignes, déplié complet au clic ; erreurs en rouge.
        const rlines = ev.result.split("\n");
        if (ocIsError(ev.result)) {
          if (det) det.classList.add("is-error");
        }
        if (rlines.length > 12) {
          body.appendChild(renderOutputMore(ev.result));
        } else {
          const pre = el("pre", "tool-result");
          appendLinkified(pre, ev.result);
          body.appendChild(pre);
        }
      } else {
        // Chat général : rendu d'origine inchangé (10 premières lignes +
        // Expand/Collapse pour les lectures).
        const rlines = ev.result.split("\n");
        if ((ev.name === "Read") && rlines.length > 10) {
          body.appendChild(renderReadMore(ev.result, rlines.length));
        } else {
          const pre = el("pre", "tool-result");
          appendLinkified(pre, ev.result);
          body.appendChild(pre);
        }
      }
    }
    this.requestFollow();
  }

  // Module Agentic, style Codex : reproduction du TUI Codex. Ligne
  // "• Running <nom> <params>" pendant l'exécution, "• Ran <nom> <params>"
  // après ; pastille verte/rouge selon le succès ; sorties en 5 premières
  // + 5 dernières lignes avec compteur explicite.
  addToolCodex(ev, gapMs) {
    if (ev.phase === "start") {
      const key = this.toolKeyStart(ev);
      this.clearEmpty();
      this.hideWait();
      this.hideThinking();
      // Temps de réflexion écoulé depuis l'événement précédent
      // ("+ Thought: 2.9s"). gapMs = 0 quand inconnu.
      this.maybeAddThought(gapMs || 0);
      this.convTools++;
      const box = el("div", "msg-tool cx-tool running");
      const head = el("div", "cx-tool-head");
      head.appendChild(el("span", "cx-bullet", "•"));
      head.appendChild(el("span", "cx-verb", "Running"));
      head.appendChild(el("span", "cx-name", ev.name));
      const params = ocToolParams(ev.name, ev.args);
      if (params) head.appendChild(el("span", "cx-params", params));
      box.appendChild(head);
      const body = el("div", "cx-tool-body");
      box.appendChild(body);
      this.log.appendChild(box);
      this.toolBoxes.set(key, { root: box, head, body });
      if (ev.name === "TodoWrite" && ev.args && Array.isArray(ev.args.todos)) {
        this.upsertTodoChecklist(ev.args.todos);
      }
      if (ev.name === "web_search" || ev.name === "web_fetch") {
        body.appendChild(renderSearchStatus(searchLabel(ev.name, ev.args)));
      }
      this.finalizeAssistant();
      this.resetAssistantState();
      this.requestFollow();
      return;
    }
    const key = this.toolKeyEnd(ev);
    const slot = this.toolBoxes.get(key);
    if (!slot || !slot.body) return;
    slot.root.classList.remove("running");
    const isErr = ocIsError(ev.result);
    // Fin d'exécution : "• Ran <nom> <params>", pastille verte/rouge.
    if (slot.head) {
      slot.head.innerHTML = "";
      slot.head.appendChild(el("span", "cx-bullet", "•"));
      slot.head.appendChild(el("span", "cx-verb", "Ran"));
      slot.head.appendChild(el("span", "cx-name", ev.name));
      const params = ocToolParams(ev.name, ev.args);
      if (params) slot.head.appendChild(el("span", "cx-params", params));
      slot.root.classList.add(isErr ? "is-error" : "is-ok");
    }
    const body = slot.body;
    const isSearch = ev.name === "web_search" || ev.name === "web_fetch";
    const status = body.querySelector(".search-status");
    if (status) status.remove();
    if (isSearch) {
      // Vue Agents (style Codex) : pas de panneau Sources ni de sortie
      // brute — seul le résumé du modèle reste visible dans le fil.
      this.requestFollow();
      return;
    }
    if (Array.isArray(ev.diff) && ev.diff.length) {
      body.appendChild(renderDiff(ev.diff, ev.args && ev.args.file_path, this.agentic ? diffVerb(ev.name) : null));
    }
    // RAG : rendu dédié compact (rag_search / rag_read), pas de gros <pre>.
    if ((ev.name === "rag_search" || ev.name === "rag_read") && typeof ev.result === "string" && ev.result) {
      body.appendChild(renderRagResult(ev.name, ev.result, ev.args));
      this.requestFollow();
      return;
    }
    if (typeof ev.result === "string" && ev.result) {
      if (isErr) {
        const pre = el("pre", "tool-result cx-tool-error");
        appendLinkified(pre, ev.result);
        body.appendChild(pre);
      } else {
        const rlines = ev.result.split("\n");
        if (rlines.length > 12) body.appendChild(renderOutputMore(ev.result));
        else {
          const pre = el("pre", "tool-result");
          appendLinkified(pre, ev.result);
          body.appendChild(pre);
        }
      }
    }
    this.requestFollow();
  }

  // Module Agentic, style OpenCode : reproduction fidele du TUI OpenCode
  // (internal/tui/components/chat/message.go). Blocs a bordure gauche
  // epaisse, en-tete "Nom: parametres" ("Nom: action..." pendant
  // l'execution), resultats bornes a 10 lignes, erreurs en rouge.
  addToolOpenCode(ev, gapMs) {
    if (ev.phase === "start") {
      const key = this.toolKeyStart(ev);
      this.clearEmpty();
      this.hideWait();
      this.hideThinking();
      // Temps de reflexion ecoule depuis l'evenement precedent
      // ("+ Thought: 2.9s"). gapMs = 0 quand inconnu.
      this.maybeAddThought(gapMs || 0);
      this.convTools++;
      const box = el("div", "msg-tool oc-tool running");
      const head = el("div", "oc-tool-head");
      head.appendChild(el("span", "oc-tool-name", ev.name + ": "));
      head.appendChild(el("span", "oc-tool-action", ocToolAction(ev.name)));
      box.appendChild(head);
      const body = el("div", "oc-tool-body");
      if (ev.name === "TodoWrite" && ev.args && Array.isArray(ev.args.todos)) {
        this.upsertTodoChecklist(ev.args.todos);
      }
      if (ev.name === "web_search" || ev.name === "web_fetch") {
        body.appendChild(renderSearchStatus(searchLabel(ev.name, ev.args)));
      }
      box.appendChild(body);
      this.log.appendChild(box);
      this.toolBoxes.set(key, { root: box, head, body });
      this.finalizeAssistant();
      this.resetAssistantState();
      this.requestFollow();
      return;
    }
    const key = this.toolKeyEnd(ev);
    const slot = this.toolBoxes.get(key);
    if (!slot || !slot.body) return;
    slot.root.classList.remove("running");
    // Fin d'execution : l'en-tete affiche "Nom: parametres".
    if (slot.head) {
      slot.head.innerHTML = "";
      slot.head.appendChild(el("span", "oc-tool-name", ev.name + ": "));
      slot.head.appendChild(
        el("span", "oc-tool-params", ocToolParams(ev.name, ev.args))
      );
    }
    const body = slot.body;
    const isSearch = ev.name === "web_search" || ev.name === "web_fetch";
    const status = body.querySelector(".search-status");
    if (status) status.remove();
    if (isSearch) {
      // Vue Agents (style OpenCode) : pas de panneau Sources ni de sortie
      // brute — seul le résumé du modèle reste visible dans le fil.
      // L'état d'erreur reste signalé visuellement (bordure rouge).
      if (ocIsError(ev.result)) slot.root.classList.add("is-error");
      this.requestFollow();
      return;
    }
    if (Array.isArray(ev.diff) && ev.diff.length) {
      body.appendChild(renderDiff(ev.diff, ev.args && ev.args.file_path, this.agentic ? diffVerb(ev.name) : null));
    }
    // RAG : rendu dédié compact (rag_search / rag_read), pas de gros <pre>.
    if ((ev.name === "rag_search" || ev.name === "rag_read") && typeof ev.result === "string" && ev.result) {
      body.appendChild(renderRagResult(ev.name, ev.result, ev.args));
      this.requestFollow();
      return;
    }
    if (typeof ev.result === "string" && ev.result) {
      if (ocIsError(ev.result)) {
        body.appendChild(el("pre", "oc-tool-error", ev.result));
      } else {
        const rlines = ev.result.split("\n");
        if (rlines.length > 10) body.appendChild(renderReadMore(ev.result, rlines.length));
        else {
          const pre = el("pre", "tool-result oc-result");
          appendLinkified(pre, ev.result);
          body.appendChild(pre);
        }
      }
    }
    this.requestFollow();
  }

  // "+ Thought: 2.9s" façon OpenCode : matérialise le temps de réflexion
  // du modèle avant son action. Ignoré si trop bref (< 300 ms, rejouement
  // d'historique ou enchaînement immédiat).
  maybeAddThought(gapMs) {
    if (!gapMs || gapMs < 300) return;
    const line = el("div", "thought-line");
    line.appendChild(el("span", "thought-prefix", "+ Thought: "));
    line.appendChild(el("span", "thought-time", fmtThoughtMs(gapMs)));
    this.log.appendChild(line);
    this.requestFollow();
  }

  // Indicateur "thinking" + spinner dans le fil pendant la réflexion.
  showThinking() {
    if (this.thinkingEl) return;
    this.hideWait();
    this.thinkingEl = el("div", "thinking-indicator");
    this.thinkingEl.appendChild(el("span", "thinking-spinner"));
    // Module Agentic, style OpenCode : libellé exact du TUI ("Thinking...").
    this.thinkingEl.appendChild(el("span", "thinking-label", this.agenticIsOpenCode() ? "Thinking..." : "thinking"));
    this.log.appendChild(this.thinkingEl);
    this.requestFollow();
  }

  hideThinking() {
    if (this.thinkingEl) {
      this.thinkingEl.remove();
      this.thinkingEl = null;
    }
  }

  // Statut façon OpenCode pendant l'exécution d'un outil : "Writing
  // command" (Bash, RunScript), "Preparing edit" (Edit, Write). Pile pour
  // les blocs parallèles : le dernier outil démarré donne le libellé.
  pushAgentStatus(key, text) {
    this.agentStatusStack = this.agentStatusStack.filter((s) => s.key !== key);
    this.agentStatusStack.push({ key, text });
    this.renderAgentStatus();
  }

  popAgentStatus(key) {
    const n = this.agentStatusStack.length;
    this.agentStatusStack = this.agentStatusStack.filter((s) => s.key !== key);
    if (this.agentStatusStack.length !== n) this.renderAgentStatus();
  }

  clearAgentStatus() {
    if (this.agentStatusStack.length || this.agentStatusEl) {
      this.agentStatusStack = [];
      this.renderAgentStatus();
    }
  }

  renderAgentStatus() {
    const top = this.agentStatusStack[this.agentStatusStack.length - 1];
    if (!top) {
      if (this.agentStatusEl) {
        this.agentStatusEl.remove();
        this.agentStatusEl = null;
      }
      return;
    }
    if (!this.agentStatusEl) {
      this.agentStatusEl = el("div", "agent-status");
      this.log.appendChild(this.agentStatusEl);
    }
    if (this.agentStatusEl.textContent !== top.text) this.agentStatusEl.textContent = top.text;
    this.requestFollow();
  }

  // Recherche web native du provider (hors outils) : indicateur pendant la
  // generation, puis panneau "Sources" a partir des annotations recues.
  handleSearch(s) {
    if (!s || typeof s !== "object") return;
    if (s.phase === "start") {
      this.clearEmpty();
      this.hideWait();
      if (!this.searchStatus) {
        this.searchStatus = renderSearchStatus("");
        this.log.appendChild(this.searchStatus);
      }
      this.requestFollow();
      return;
    }
    if (this.searchStatus) {
      this.searchStatus.remove();
      this.searchStatus = null;
    }
    if (this.agentic) {
      // Vue Agents : pas de panneau Sources — seul le résumé du modèle
      // reste visible dans le fil.
      this.requestFollow();
      return;
    }
    const srcs = Array.isArray(s.sources) ? s.sources : [];
    if (srcs.length) {
      this.clearEmpty();
      this.log.appendChild(renderSources(srcs));
      this.requestFollow();
    }
  }

  summarizeApprovalArgs(tool, args) {
    if (!args) return "";
    if (args.file_path) return String(args.file_path);
    if (args.command) return String(args.command);
    if (args.pattern) return String(args.pattern);
    if (args.query) return String(args.query);
    // Sed / Awk : afficher l'expression ou le programme (première ligne)
    // au lieu de laisser la carte sans résumé.
    if (args.expression) return String(args.expression).split("\n")[0].slice(0, 200);
    if (args.script) return String(args.script).split("\n")[0].slice(0, 200);
    if (args.program) return String(args.program).split("\n")[0].slice(0, 200);
    return "";
  }

  async decideApproval(id, approved, always, comment) {
    if (!this.approveURL) return;
    // F5 : feedback immédiat — on désactive les boutons dès le clic pour
    // éviter les doubles décisions, avec un état "Envoi…" visible.
    const card = this.approvalCards.get(id);
    if (card && card.dataset.sending === "1") return; // déjà en cours d'envoi
    const premium = !!card && card.classList.contains("ap2");
    let status = null;
    let bannerText = null;
    if (card) {
      card.dataset.sending = "1";
      card.querySelectorAll("button").forEach((b) => { b.disabled = true; });
      status = card.querySelector(premium ? ".ap2-status" : ".approval-status");
      if (status) status.textContent = "Envoi…";
      if (premium) {
        bannerText = card.querySelector(".ap2-banner-text");
        if (bannerText) bannerText.textContent = "Envoi de la décision…";
        // Mémorise la décision pour la trace persistante (construite à la
        // réception de l'événement "resolved").
        card._ap2 = card._ap2 || {};
        card._ap2.decision = { approved: !!approved, always: !!always, comment: comment || "" };
      }
    }
    try {
      await api(this.approveURL, {
        method: "POST",
        body: { id, approved, always: !!always, comment: comment || "" },
      });
    } catch (e) {
      // Échec : on réarme les boutons pour permettre un nouvel essai.
      if (card) {
        delete card.dataset.sending;
        card.querySelectorAll("button").forEach((b) => { b.disabled = false; });
        if (status) status.textContent = premium ? "" : "En attente de ta décision…";
        if (bannerText) bannerText.textContent = "En attente de décision";
      }
      this.addError("Approbation : " + e.message);
    }
  }

  setApprovalResolved(id, approved, timeout) {
    const card = this.approvalCards.get(id);
    if (!card) return;
    card.setAttribute("data-resolved", "1");
    const btns = card.querySelectorAll("button");
    btns.forEach((b) => { b.disabled = true; });
    // Chat général : comportement historique strictement inchangé.
    if (!card.classList.contains("ap2")) {
      const status = card.querySelector(".approval-status");
      if (status) {
        status.textContent = timeout ? "Expirée (10 min sans réponse)" : approved ? "Approuvé" : "Refusé";
        status.classList.add(approved && !timeout ? "approved" : "denied");
      }
      this.requestFollow();
      return;
    }
    const denyBox = card.querySelector(".ap2-denybox");
    if (denyBox) denyBox.hidden = true;
    const info = card._ap2 || {};
    const tool = info.tool || "";
    const arg = info.arg || "";
    const dec = info.decision || {};
    // Trace persistante dans le fil : "✔ Approuvé · Edit · src/app.js".
    let trace, traceCls, bannerLabel;
    if (timeout) {
      trace = "⌛ Approbation expirée (10 min sans réponse)";
      traceCls = "is-timeout";
      bannerLabel = "Expirée";
    } else if (approved) {
      trace = dec.always ? "✔ Toujours approuver (ce tour)" : "✔ Approuvé";
      traceCls = "is-ok";
      bannerLabel = "Décision enregistrée";
    } else {
      trace = "✖ Refusé";
      traceCls = "is-denied";
      bannerLabel = "Décision enregistrée";
      if (dec.comment) trace += " — " + dec.comment;
    }
    if (tool) trace += " · " + tool;
    if (arg) trace += " · " + arg;
    const banner = card.querySelector(".ap2-banner");
    if (banner) banner.classList.add(traceCls);
    const dot = card.querySelector(".ap2-dot");
    if (dot) dot.classList.add(traceCls);
    const bannerText = card.querySelector(".ap2-banner-text");
    if (bannerText) bannerText.textContent = bannerLabel;
    const status = card.querySelector(".ap2-status");
    if (status) {
      status.textContent = timeout
        ? "Expirée (10 min sans réponse)"
        : approved ? (dec.always ? "Toujours approuver (ce tour)" : "Approuvé") : "Refusé";
      status.classList.add(traceCls);
    }
    // La trace suit la carte dans le fil et survit au reset (cartes
    // résolues retirées, trace conservée comme les autres messages).
    const line = el("div", "ap2-trace " + traceCls, trace);
    const parent = card.parentNode;
    if (parent) {
      const kids = parent.children;
      let idx = -1;
      for (let i = 0; i < kids.length; i++) {
        if (kids[i] === card) { idx = i; break; }
      }
      parent.insertBefore(line, idx >= 0 && idx + 1 < kids.length ? kids[idx + 1] : null);
    }
    this.requestFollow();
  }

  // Phase 2 refonte : carte d'approbation premium (vue Agents uniquement).
  // Identique pour les 3 styles (Harness / OpenCode / Codex) : bandeau
  // d'attente ambré, headline = question d'action, ligne d'identité outil
  // (glyphe + nom + argument clé), détail structuré par outil (jamais de
  // JSON brut), actions à droite, champ de commentaire au refus, trace
  // persistante après décision.
  addApproval(ev) {
    const id = ev.id;
    if (ev.phase === "resolved") {
      this.setApprovalResolved(id, ev.approved === true, ev.timeout === true);
      return;
    }
    if (ev.phase !== "request" || !id || this.approvalCards.has(id)) return;
    // Chat général : carte historique strictement inchangée
    // (hors périmètre de la refonte, préservation du comportement).
    if (!this.agentic) {
      this.addApprovalLegacy(ev);
      return;
    }
    this.clearEmpty();
    this.hideWait();
    const isPlan = ev.kind === "plan";
    const tool = ev.tool || "";
    const args = ev.args && typeof ev.args === "object" ? ev.args : {};
    const keyArg = isPlan ? "" : this.summarizeApprovalArgs(tool, args);

    const card = el("div", "msg-approval ap2");
    card._ap2 = { tool: isPlan ? "Plan" : tool, arg: keyArg };

    // Bandeau d'attente : pastille pulsante + libellé.
    const banner = el("div", "ap2-banner");
    banner.appendChild(el("span", "ap2-dot"));
    banner.appendChild(el("span", "ap2-banner-text", "En attente de décision"));
    card.appendChild(banner);

    // Headline : la question d'action.
    card.appendChild(el("div", "ap2-headline", approvalHeadline(ev.kind, tool)));

    // Ligne d'identité de l'outil.
    if (!isPlan && tool) {
      const ident = el("div", "ap2-tool");
      ident.appendChild(el("span", "tool-glyph", toolGlyph(tool)));
      ident.appendChild(el("span", "tool-name", tool));
      if (keyArg) ident.appendChild(el("span", "ap2-tool-arg", keyArg));
      card.appendChild(ident);
    }

    // Détail structuré (plan : texte borné).
    if (isPlan && ev.plan) {
      const pre = el("pre", "ap2-plan");
      pre.textContent = String(ev.plan).slice(0, 4000);
      card.appendChild(pre);
    } else if (!isPlan) {
      card.appendChild(approvalDetail(tool, args));
    }

    // Zone de refus : révélée au clic sur "Refuser".
    const denyBox = el("div", "ap2-denybox");
    denyBox.hidden = true;
    denyBox.appendChild(el("label", "ap2-deny-label", "Que faire différemment ?"));
    const denyInput = el("textarea", "ap2-deny-input");
    denyInput.placeholder = "Ex. : utilise npm plutôt que yarn… (optionnel)";
    denyInput.rows = 2;
    denyBox.appendChild(denyInput);
    card.appendChild(denyBox);

    const status = el("div", "ap2-status", "");
    card.appendChild(status);

    // Actions à droite : Refuser (outline) / Toujours ce tour / Approuver.
    const row = el("div", "ap2-actions");
    const btnDeny = el("button", "ap2-btn ap2-btn-deny", "Refuser");
    const btnAlways = el("button", "ap2-btn ap2-btn-always", "Toujours ce tour");
    const btnApprove = el("button", "ap2-btn ap2-btn-approve", isPlan ? "Valider le plan" : "Approuver");
    const btnCancelDeny = el("button", "ap2-btn ap2-btn-ghost", "Annuler");
    btnCancelDeny.hidden = true;
    const btnConfirmDeny = el("button", "ap2-btn ap2-btn-deny-solid", "Confirmer le refus");
    btnConfirmDeny.hidden = true;

    btnDeny.addEventListener("click", () => {
      denyBox.hidden = false;
      btnDeny.hidden = true;
      btnApprove.hidden = true;
      btnAlways.hidden = true;
      btnCancelDeny.hidden = false;
      btnConfirmDeny.hidden = false;
      if (typeof denyInput.focus === "function") denyInput.focus();
    });
    btnCancelDeny.addEventListener("click", () => {
      denyBox.hidden = true;
      btnDeny.hidden = false;
      btnApprove.hidden = false;
      btnAlways.hidden = false;
      btnCancelDeny.hidden = true;
      btnConfirmDeny.hidden = true;
    });
    btnApprove.addEventListener("click", () => this.decideApproval(id, true, false, ""));
    if (!isPlan) btnAlways.addEventListener("click", () => this.decideApproval(id, true, true, ""));
    btnConfirmDeny.addEventListener("click", () =>
      this.decideApproval(id, false, false, denyInput.value.trim()));
    row.appendChild(btnDeny);
    row.appendChild(btnCancelDeny);
    if (!isPlan) row.appendChild(btnAlways);
    row.appendChild(btnApprove);
    row.appendChild(btnConfirmDeny);
    card.appendChild(row);

    this.log.appendChild(card);
    this.approvalCards.set(id, card);
    this.finalizeAssistant();
    this.toBottom();
  }

  // Carte d'approbation historique du chat général — strictement inchangée
  // (phase 2 refonte : périmètre = module Agents uniquement).
  addApprovalLegacy(ev) {
    const id = ev.id;
    this.clearEmpty();
    this.hideWait();
    const card = el("div", "msg-approval");
    const head = el("div", "approval-head");
    const isPlan = ev.kind === "plan";
    head.appendChild(el("span", "approval-title", isPlan ? "Plan à valider" : "Approbation requise"));
    if (!isPlan && ev.tool) head.appendChild(el("span", "tool-name", " " + ev.tool));
    card.appendChild(head);
    const hint = isPlan ? "" : this.summarizeApprovalArgs(ev.tool, ev.args);
    if (hint) card.appendChild(el("div", "approval-hint", hint));
    if (isPlan && ev.plan) {
      const pre = el("pre", "approval-plan");
      pre.textContent = String(ev.plan).slice(0, 4000);
      card.appendChild(pre);
    } else if (ev.args && typeof ev.args === "object") {
      const det = el("details", "approval-args");
      det.appendChild(el("summary", "", "Détails"));
      const pre = el("pre");
      const shown = Object.assign({}, ev.args);
      const truncated = [];
      for (const k of ["content", "code", "new"]) {
        if (typeof shown[k] === "string" && shown[k].length > 600) {
          shown[k] = shown[k].slice(0, 600) + "…";
          truncated.push(k);
        }
      }
      if (truncated.length) {
        det.appendChild(el("div", "approval-truncated",
          "Contenu tronqué à 600 caractères (" + truncated.join(", ") +
          ") : relis la valeur complète avant d'approuver."));
      }
      pre.textContent = JSON.stringify(shown, null, 2);
      det.appendChild(pre);
      card.appendChild(det);
    }
    const status = el("div", "approval-status", "En attente de ta décision…");
    card.appendChild(status);
    const row = el("div", "approval-actions");
    const btnApprove = el("button", "btn-approve", isPlan ? "Valider le plan" : "Approuver");
    btnApprove.addEventListener("click", () => this.decideApproval(id, true, false, ""));
    const btnDeny = el("button", "btn-deny", "Refuser");
    btnDeny.addEventListener("click", () => this.decideApproval(id, false, false, ""));
    row.appendChild(btnApprove);
    row.appendChild(btnDeny);
    if (!isPlan) {
      const btnAlways = el("button", "btn-always", "Toujours approuver (ce tour)");
      btnAlways.addEventListener("click", () => this.decideApproval(id, true, true, ""));
      row.appendChild(btnAlways);
    }
    card.appendChild(row);
    this.log.appendChild(card);
    this.approvalCards.set(id, card);
    this.finalizeAssistant();
    this.toBottom();
  }

  // Barre de statut façon Harness : "2 tours · 6 outils · ↑13,7k ↓1,2k".
  // Alimente statsBadge, recopié vers le compteur du composer par onDone.
  updateConvBadge() {
    if (!this.statsBadge) return;
    this.statsBadge.hidden = false;
    const inT = this.convIn + this.turnInTok;
    const outT = this.convOut + this.turnOutTok;
    const t = this.convTurns;
    const parts = [
      t + (t > 1 ? " tours" : " tour"),
      this.convTools + (this.convTools > 1 ? " outils" : " outil"),
      "↑" + fmtK(inT) + " ↓" + fmtK(outT),
    ];
    this.statsBadge.textContent = parts.join(" · ");
    this.statsBadge.title =
      "Entrée : " + inT.toLocaleString("fr") + " tokens · Sortie : " +
      outT.toLocaleString("fr") + " tokens (conversation)";
  }

  addTurnStats(box) {
    // Module Agentic, style Codex (phase 3 refonte) : ligne discrète
    // "modèle · durée" façon TUI Codex, tokens en infobulle.
    if (this.agenticIsCodex()) {
      const r = this.turnRoute || {};
      const model = r.label || r.model || "";
      let secs = 0;
      if (this.turnElapsedMs != null) secs = this.turnElapsedMs / 1000;
      else if (this.turnStartTs) secs = (Date.now() - this.turnStartTs) / 1000;
      const txt = (model ? model + " · " : "") + secs.toFixed(secs < 10 ? 1 : 0) + "s";
      const wrapper = box.closest(".message-wrapper") || box;
      const s = this.turnStats || {};
      const div = el("div", "cx-turn-stats", txt);
      div.title =
        "Entrée : " + (s.prompt_tokens || 0).toLocaleString("fr") + " tokens · Sortie : " +
        (s.completion_tokens || 0).toLocaleString("fr") + " tokens";
      wrapper.appendChild(div);
      return;
    }
    // Module Agentic, style OpenCode : pied façon TUI " modèle (durée)".
    if (this.agenticIsOpenCode()) {
      const r = this.turnRoute || {};
      const model = r.label || r.model || "";
      let secs = 0;
      if (this.turnElapsedMs != null) secs = this.turnElapsedMs / 1000;
      else if (this.turnStartTs) secs = (Date.now() - this.turnStartTs) / 1000;
      const txt =
        " " + (model ? model + " " : "") + "(" + secs.toFixed(secs < 10 ? 1 : 0) + "s)";
      const wrapper = box.closest(".message-wrapper") || box;
      wrapper.appendChild(el("div", "oc-turn-stats", txt));
      return;
    }
    const s = this.turnStats || {};
    const inTok = s.prompt_tokens || 0;
    const outTok = s.completion_tokens || 0;
    const r = this.turnRoute || {};
    const model = r.label || r.model || "";
    let secs = 0;
    if (this.turnElapsedMs != null) secs = this.turnElapsedMs / 1000;
    else if (this.turnStartTs) secs = (Date.now() - this.turnStartTs) / 1000;
    const parts = [];
    if (model) parts.push(model);
    parts.push(secs.toFixed(secs < 10 ? 1 : 0) + "s");
    let tokTxt = outTok.toLocaleString("fr") + " tokens";
    if (secs > 0 && outTok > 0) tokTxt += " (" + Math.round(outTok / secs) + "/s)";
    parts.push(tokTxt);
    const wrapper = box.closest(".message-wrapper") || box;
    const div = el("div", "turn-stats", parts.join(" · "));
    div.title =
      "Entrée : " + inTok.toLocaleString("fr") + " tokens · Sortie : " + outTok.toLocaleString("fr") + " tokens";
    wrapper.appendChild(div);
  }

  finishTurn() {
    this.finalizeAssistant();
    const box = this.assistant;
    const raw = this.assistantText;
    if (box) box.classList.remove("streaming");
    this.hideWait();
    this.hideThinking();
    this.clearAgentStatus();
    this.setBusy(false);
    this.generating = false;
    // Valide les tokens du tour vers le cumul de la conversation.
    this.convIn += this.turnInTok;
    this.convOut += this.turnOutTok;
    this.turnInTok = 0;
    this.turnOutTok = 0;
    this.updateConvBadge();
    if (this.stopBtn) this.stopBtn.hidden = true;
    if (this.reasonPanel) {
      // Fige les infos du panneau et garde un instantane pour le bouton
      // "Raisonnement" de la reponse.
      const snap = finalizeReasonTurn(this.turnElapsedMs);
      if (snap.text && snap.text.trim()) {
        // Tour sans contenu : on cree quand meme la bulle pour accueillir
        // l'en-tete, le raisonnement reste reconsultable.
        if (!this.reasonBtn) this.ensureAssistant();
        if (this.reasonAssistantWrapper) {
          this.reasonAssistantWrapper._reasonSnap = snap;
          if (this.reasonBtn) this.reasonBtn.hidden = false;
        }
      }
      // Sécurité : si le raisonnement n'a jamais basculé sur du contenu
      // (tour sans réponse), on masque quand même le panneau.
      if (this.reasoningActive) {
        this.reasoningActive = false;
        finishReasoning();
      }
    } else if (this.reasonHooks) this.reasonHooks.finish();
    if (box && raw.trim()) {
      this.addTurnStats(box);
      this.addActions(box, raw);
    }
    this.resetAssistantState();
    this.turnStartTs = 0;
    this.turnStats = null;
    this.turnRoute = null;
    this.turnElapsedMs = null;
    if (this.onDone) this.onDone();
  }

  reset() {
    this.hideWait();
    this.setBusy(false);
    if (this.reasonPanel) resetReasonPanel();
      else if (this.reasonHooks) this.reasonHooks.reset();
    // F10 : on ne jette pas les cartes d'approbation en attente lors d'un
    // reset — seules les cartes résolues sont retirées. Les cartes en
    // attente sont mises de côté avant le vidage du fil puis réaffichées.
    const pendingCards = [];
    for (const [id, card] of this.approvalCards) {
      if (card && !card.hasAttribute("data-resolved")) pendingCards.push(card);
      else this.approvalCards.delete(id);
    }
    this.log.innerHTML = this.emptyHTML;
    for (const card of pendingCards) this.log.appendChild(card);
    this.empty = this.log.querySelector("[data-empty]");
    this.lastSeq = 0;
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
    this.reasoningActive = false;
    this.reasonBtn = null;
    this.reasonAssistantWrapper = null;
    this.stopStreamSpinner();
    this.toolBoxes.clear();
    this.toolPending.clear();
    // Phase 1 refonte : invalide la checklist TodoWrite live (le fil est
    // vidé, le bloc sera recréé au prochain appel TodoWrite).
    this.threadGen++;
    this.todoChecklistEl = null;
    this.searchStatus = null;
    this.lastEventTs = 0;
    this.hideThinking();
    this.clearAgentStatus();
    this.generating = false;
    this.turnStartTs = 0;
    this.turnStats = null;
    this.turnRoute = null;
    this.turnElapsedMs = null;
    this.convTurns = 0;
    this.convTools = 0;
    this.convIn = 0;
    this.convOut = 0;
    this.turnInTok = 0;
    this.turnOutTok = 0;
    if (this.followFrame) {
      cancelRafTick(this.followFrame);
      this.followFrame = 0;
    }
    if (this.sampleTimer) {
      clearTimeout(this.sampleTimer);
      this.sampleTimer = 0;
    }
    this.pinned = true;
    this.observedTop = 0;
    this.newWhileUnpinned = false;
    this.prunedCount = 0;
    this.updateToBottomBtn();
    if (this.stopBtn) this.stopBtn.hidden = true;
    if (this.routeBadge) this.routeBadge.hidden = true;
    if (this.statsBadge) {
      this.statsBadge.hidden = true;
      this.statsBadge.textContent = "";
    }
  }

  handleEvent(ev) {
    if (typeof ev.seq === "number" && ev.seq > this.lastSeq) this.lastSeq = ev.seq;
    if (this.onEvent) {
      try { this.onEvent(ev); } catch (e) { /* jamais bloquant pour le fil */ }
    }
    if (ev.reset) {
      this.reset();
      return;
    }
    if (ev.pad !== undefined) return;
    // Horodatage pour "+ Thought: Xs" : le délai entre deux événements
    // mesure le temps de réflexion du modèle (les keepalives "pad"
    // ne comptent pas).
    const nowTs = Date.now();
    const gapMs = this.lastEventTs ? nowTs - this.lastEventTs : 0;
    this.lastEventTs = nowTs;
    if (ev.caught_up) {
      this.toBottom();
      return;
    }
    if (ev.user !== undefined) {
      this.resetAssistantState();
      // Déduplication de l'écho optimiste : le message a déjà été affiché
      // à l'envoi avec ce client_msg_id — on le confirme au lieu de le
      // dupliquer. Le reste du traitement du tour reste identique.
      if (ev.client_msg_id && ev.client_msg_id === this.pendingUserId) {
        this.pendingUserId = null;
        const opt = this.log.querySelector(
          ':scope > .message-wrapper-user[data-optimistic="1"]'
        );
        if (opt) opt.removeAttribute("data-optimistic");
      } else {
        const text = String(ev.user);
        // F16 : cas limite — le delta "user" du serveur arrive sans
        // client_msg_id correspondant, mais un écho optimiste IDENTIQUE est
        // déjà affiché : on le confirme au lieu d'ajouter un doublon.
        const opt = this.log.querySelector(
          ':scope > .message-wrapper-user[data-optimistic="1"]'
        );
        const optText = opt && opt.querySelector(".message-text");
        if (opt && optText && optText.textContent === text) {
          this.pendingUserId = ev.client_msg_id || null;
          opt.removeAttribute("data-optimistic");
        } else {
          this.addUser(text);
        }
      }
      this.turnStartTs = Date.now();
      this.turnStats = null;
      this.turnRoute = null;
      this.turnElapsedMs = null;
      this.convTurns++;
      this.updateConvBadge();
      if (this.reasonPanel) {
        sealReasonTurn();
        beginReasonTurn();
      } else if (this.reasonHooks) this.reasonHooks.reset();
      this.generating = true;
      if (this.stopBtn) this.stopBtn.hidden = false;
      this.showWait();
      this.setBusy(true);
      return;
    }
    if (ev.reasoning_content !== undefined) {
      this.appendReasoning(String(ev.reasoning_content), ev.replace === true);
      return;
    }
    if (ev.content !== undefined) {
      this.appendContent(String(ev.content), ev.replace === true);
      return;
    }
    if (ev.tool !== undefined) {
      this.addTool(ev.tool, gapMs);
      return;
    }
    if (ev.search !== undefined) {
      this.handleSearch(ev.search);
      return;
    }
    if (ev.approval !== undefined) {
      this.addApproval(ev.approval);
      return;
    }
    if (ev.workspace !== undefined) {
      // Espace de travail lié au run, annoncé en tête de fil par le
      // backend : l'utilisateur voit toujours OÙ l'agent travaille
      // (projet lié ou espace partagé).
      const w = ev.workspace || {};
      const wname = String(w.name || "?");
      const wmode = w.mode ? " (" + String(w.mode) + ")" : "";
      this.addSystem("📁 Espace de travail : " + wname + wmode);
      return;
    }
    if (ev.worktree !== undefined) {
      const wt = ev.worktree || {};
      this.addSystem("🌿 Worktree isolé : " + (wt.path || ""));
      return;
    }
    if (ev.worktree_error !== undefined) {
      this.addSystem("⚠️ Worktree indisponible : " + String(ev.worktree_error) + ".");
      return;
    }
    if (ev.stats !== undefined) {
      const s = ev.stats || {};
      this.turnStats = s;
      // Les stats partent à chaque appel provider du tour : on remplace
      // (pas de cumul ici), la validation vers le cumul se fait au turn_done.
      this.turnInTok = s.prompt_tokens || 0;
      this.turnOutTok = s.completion_tokens || 0;
      this.updateConvBadge();
      if (this.reasonPanel) updateReasonUsage(s);
      if (this.trackTokens) setTurnStats(s.prompt_tokens, s.completion_tokens);
      return;
    }
    if (ev.compact) {
      this.addSystem("Contexte compacté pour rester dans la fenêtre du modèle.");
      return;
    }
    if (ev.system !== undefined) {
      this.addSystem(String(ev.system));
      return;
    }
    if (ev.route !== undefined) {
      const r = ev.route || {};
      this.turnRoute = r;
      if (this.reasonPanel) setReasonModel(r.label || r.model || "", r.provider || "", r.model || "");
      if (this.routeBadge) {
        this.routeBadge.hidden = false;
        const name = r.label || r.model || "";
        // L'effort affiché est celui RÉSOLU par le moteur ("" si thinking
        // désactivé) : rend visible la résolution du mode "Défaut"
        // (texte long > 400 caractères → Moyen).
        const effortFr = { low: "Faible", medium: "Moyen", high: "Max" }[r.effort];
        this.routeBadge.textContent =
          (r.provider ? r.provider + "/" : "") + name + (r.local ? " (local)" : "") + (r.fallback ? " (repli)" : "") +
          (effortFr ? " · effort " + effortFr : "");
      }
      return;
    }
    if (ev.drop_reasoning) {
      this.removeReasoning();
      return;
    }
    if (ev.error !== undefined) {
      this.addError(String(ev.error));
      this.clearAgentStatus();
      return;
    }
    if (ev.turn_done !== undefined) {
      const td = ev.turn_done || {};
      if (td.elapsed_ms != null) this.turnElapsedMs = td.elapsed_ms;
      this.finishTurn();
    }
  }

  connect() {
    if (this.controller) this.controller.abort();
    this.controller = new AbortController();
    this.sseRetries = 0;
    const url = this.streamURL(this.lastSeq);
    const run = async () => {
      try {
        const resp = await fetch(url, {
          headers: { Authorization: "Bearer " + getToken() },
          signal: this.controller.signal,
        });
        if (resp.status === 401) {
          // Session expirée : même signal que le reste du code (api.js).
          window.dispatchEvent(new CustomEvent("cetas:unauthorized"));
          return;
        }
        if (!resp.ok || !resp.body) {
          this.scheduleReconnect(run);
          return;
        }
        this.sseRetries = 0; // connexion réussie : on repart de zéro
        await readSSE(resp, (ev) => this.handleEvent(ev));
      } catch (e) {
        if (e && e.name === "AbortError") return;
        this.scheduleReconnect(run);
      }
    };
    run();
  }

  // Reconnexion SSE avec backoff exponentiel et plafond de tentatives :
  // plus de boucle infinie silencieuse, et un message visible quand le
  // plafond est atteint.
  scheduleReconnect(run) {
    if (!this.controller || this.controller.signal.aborted) return;
    const maxRetries = 8;
    if (this.sseRetries >= maxRetries) {
      this.addError("Connexion temps réel perdue après plusieurs tentatives. Recharge la page pour reprendre le fil.");
      return;
    }
    const delay = Math.min(1500 * Math.pow(2, this.sseRetries), 30000);
    this.sseRetries++;
    setTimeout(() => {
      if (this.controller && !this.controller.signal.aborted) run();
    }, delay);
  }

  disconnect() {
    if (this.controller) {
      this.controller.abort();
      this.controller = null;
    }
  }

  async sendText(text) {
    text = String(text || "").trim();
    if (!text || this.generating) return false;
    // Écho optimiste (vue Agents uniquement) : le message et l'indicateur
    // d'attente s'affichent immédiatement, sans attendre l'aller-retour
    // POST. Le delta "user" du serveur porte le même client_msg_id et
    // sera dédupliqué dans handleEvent.
    let clientMsgId = null;
    let optimistic = null;
    if (this.optimisticEcho) {
      clientMsgId = this.newClientMsgId();
      optimistic = this.addUserOptimistic(text);
      this.pendingUserId = clientMsgId;
      this.showWait();
      this.setBusy(true);
      if (this.stopBtn) this.stopBtn.hidden = false;
    }
    const payload = this.getPayload(text);
    if (clientMsgId) payload.client_msg_id = clientMsgId;
    try {
      await api(this.sendURL, { method: "POST", body: payload });
      this.generating = true;
      this.lastSendError = null;
      if (this.stopBtn) this.stopBtn.hidden = false;
      this.setBusy(true);
      this.toBottom();
      return true;
    } catch (err) {
      if (this.optimisticEcho) {
        // Échec du POST : on retire l'écho optimiste, le serveur n'émettra
        // jamais le delta correspondant.
        this.removeOptimistic(optimistic);
        this.hideWait();
        this.setBusy(false);
      }
      if (/en cours/i.test(err.message)) {
        this.generating = true;
        this.lastSendError = null;
        if (this.stopBtn) this.stopBtn.hidden = false;
        this.setBusy(true);
        return true;
      }
      this.lastSendError = err.message;
      this.addError(err.message);
      return false;
    }
  }

  stop() {
    if (!this.stopURL) return;
    // F12 : ne plus avaler silencieusement l'échec — on affiche l'erreur
    // pour que l'utilisateur sache que l'arrêt n'est peut-être pas parti
    // (le bouton reste visible pour réessayer ; le succès est confirmé par
    // l'événement serveur qui masque le bouton).
    api(this.stopURL, { method: "POST" }).catch((e) => {
      this.addError("Stop : " + (e && e.message ? e.message : "requête non envoyée") + " — le tour continue peut-être encore.");
    });
  }
}
