import { api, getToken, readSSE } from "./api.js";
import { appendLinkified, createMarkdownRenderer } from "./markdown.js";
import {
  appendReasoningPanel,
  finishReasoningPanel,
  resetReasonPanel,
} from "./reasoning-panel.js";
import { setTurnStats } from "./turn-tokens.js";

const BRAILLE = ["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"];

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
    const a = el("a", "citation-card");
    a.href = url;
    a.target = "_blank";
    a.rel = "noopener noreferrer";
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

function renderDiff(lines, filePath) {
  const wrap = el("div", "tool-diff");
  let adds = 0, dels = 0;
  for (const line of lines) {
    if (line.kind === "+") adds++;
    else if (line.kind === "-") dels++;
  }
  const head = el("div", "diff-head");
  head.appendChild(el("span", "diff-file", filePath ? String(filePath) : "modification"));
  head.appendChild(el("span", "diff-stats", "+" + adds + " / -" + dels));
  wrap.appendChild(head);
  const body = el("div", "diff-body");
  let oldN = 0, newN = 0;
  for (const line of lines) {
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
  wrap.appendChild(body);
  return wrap;
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

    this.empty = this.log.querySelector("[data-empty]");
    this.lastSeq = 0;
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
    this.controller = null;
    this.generating = false;
    this.waitEl = null;
    this.waitTimer = null;
    this.waitIdx = 0;
    this.toolBoxes = new Map();
    this.approvalCards = new Map();
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
    if (this.assistantBody) mdRenderer.finalize(this.assistantBody, this.assistantText);
  }

  resetAssistantState() {
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
  }

  setBusy(v) {
    this.log.setAttribute("aria-busy", v ? "true" : "false");
  }

  showWait() {
    if (this.waitEl) return;
    this.waitEl = el("div", "stream-waiting");
    this.waitEl.appendChild(el("span", "stream-spinner", BRAILLE[0]));
    this.log.appendChild(this.waitEl);
    this.waitTimer = setInterval(() => {
      this.waitIdx = (this.waitIdx + 1) % BRAILLE.length;
      if (this.waitEl && this.waitEl.firstChild) this.waitEl.firstChild.textContent = BRAILLE[this.waitIdx];
    }, 80);
    this.toBottom();
  }

  hideWait() {
    if (this.waitTimer) {
      clearInterval(this.waitTimer);
      this.waitTimer = null;
    }
    if (this.waitEl) {
      this.waitEl.remove();
      this.waitEl = null;
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
  }

  ensureAssistant() {
    if (!this.assistant) {
      this.clearEmpty();
      const wrapper = el("div", "message-wrapper message-wrapper-assistant");
      this.assistant = el("div", "message message-assistant streaming");
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
      this.reasoningEl.open = true;
      this.reasoningEl.appendChild(el("summary", null, "Raisonnement"));
      this.reasoningEl.appendChild(el("div", "thinking-content"));
      this.assistant.insertBefore(this.reasoningEl, this.assistant.firstChild);
    }
    return this.reasoningEl;
  }

  appendReasoning(text, isReplace) {
    if (this.reasonHooks) {
      this.reasonHooks.append(text, isReplace);
      return;
    }
    if (this.reasonPanel) {
      appendReasoningPanel(text, isReplace);
      return;
    }
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
    if (this.reasoningEl) {
      this.reasoningEl.remove();
      this.reasoningEl = null;
      this.reasoningText = "";
    }
  }

  addError(text) {
    this.hideWait();
    this.setBusy(false);
    this.clearEmpty();
    const wrapper = el("div", "message-wrapper message-wrapper-assistant");
    const bubble = el("div", "message message-assistant message-error");
    bubble.appendChild(el("div", "message-text", text));
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
      const regen = el("button", "regen-btn", "Regenerer");
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
    if (window.speechSynthesis) {
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

  addTool(ev) {
    const key = ev.name + "|" + JSON.stringify(ev.args || {});
    if (ev.phase === "start") {
      this.clearEmpty();
      this.hideWait();
      const box = el("details", "msg-tool");
      const summary = el("summary");
      summary.appendChild(el("span", "tool-name", ev.name));
      const hint = summarizeArgs(ev.args);
      if (hint) summary.appendChild(el("span", "tool-hint", " " + hint));
      box.appendChild(summary);
      const body = el("div", "tool-body");
      if (ev.name === "TodoWrite" && ev.args && Array.isArray(ev.args.todos)) {
        body.appendChild(renderTodos(ev.args.todos));
      }
      if (ev.name === "web_search" || ev.name === "web_fetch") {
        body.appendChild(renderSearchStatus(searchLabel(ev.name, ev.args)));
      }
      box.appendChild(body);
      this.log.appendChild(box);
      this.toolBoxes.set(key, body);
      this.finalizeAssistant();
      this.resetAssistantState();
      this.requestFollow();
      return;
    }
    const body = this.toolBoxes.get(key);
    if (!body) return;
    // Fin d'une recherche web : l'indicateur laisse place au panneau Sources.
    const isSearch = ev.name === "web_search" || ev.name === "web_fetch";
    const status = body.querySelector(".search-status");
    if (status) status.remove();
    if (isSearch && Array.isArray(ev.sources) && ev.sources.length) {
      body.appendChild(renderSources(ev.sources));
      this.requestFollow();
      return;
    }
    if (Array.isArray(ev.diff) && ev.diff.length) {
      body.appendChild(renderDiff(ev.diff, ev.args && ev.args.file_path));
    }
    if (typeof ev.result === "string" && ev.result) {
      const pre = el("pre", "tool-result");
      appendLinkified(pre, ev.result);
      body.appendChild(pre);
    }
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
    return "";
  }

  async decideApproval(id, approved, always) {
    if (!this.approveURL) return;
    try {
      await api(this.approveURL, { method: "POST", body: { id, approved, always: !!always } });
    } catch (e) {
      this.addError("Approbation : " + e.message);
    }
  }

  setApprovalResolved(id, approved, timeout) {
    const card = this.approvalCards.get(id);
    if (!card) return;
    card.setAttribute("data-resolved", "1");
    const btns = card.querySelectorAll("button");
    btns.forEach((b) => { b.disabled = true; });
    const status = card.querySelector(".approval-status");
    if (status) {
      status.textContent = timeout ? "Expirée (10 min sans réponse)" : approved ? "Approuvé" : "Refusé";
      status.classList.add(approved && !timeout ? "approved" : "denied");
    }
    this.requestFollow();
  }

  addApproval(ev) {
    const id = ev.id;
    if (ev.phase === "resolved") {
      this.setApprovalResolved(id, ev.approved === true, ev.timeout === true);
      return;
    }
    if (ev.phase !== "request" || !id || this.approvalCards.has(id)) return;
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
      for (const k of ["content", "code", "new"]) {
        if (typeof shown[k] === "string" && shown[k].length > 600) shown[k] = shown[k].slice(0, 600) + "…";
      }
      pre.textContent = JSON.stringify(shown, null, 2);
      det.appendChild(pre);
      card.appendChild(det);
    }
    const status = el("div", "approval-status", "En attente de ta décision…");
    card.appendChild(status);
    const row = el("div", "approval-actions");
    const btnApprove = el("button", "btn-approve", isPlan ? "Valider le plan" : "Approuver");
    btnApprove.addEventListener("click", () => this.decideApproval(id, true, false));
    const btnDeny = el("button", "btn-deny", "Refuser");
    btnDeny.addEventListener("click", () => this.decideApproval(id, false, false));
    row.appendChild(btnApprove);
    row.appendChild(btnDeny);
    if (!isPlan) {
      const btnAlways = el("button", "btn-always", "Toujours approuver (ce tour)");
      btnAlways.addEventListener("click", () => this.decideApproval(id, true, true));
      row.appendChild(btnAlways);
    }
    card.appendChild(row);
    this.log.appendChild(card);
    this.approvalCards.set(id, card);
    this.finalizeAssistant();
    this.toBottom();
  }

  addTurnStats(box) {
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
    this.setBusy(false);
    this.generating = false;
    if (this.stopBtn) this.stopBtn.hidden = true;
    if (this.reasonPanel) finishReasoningPanel();
    else if (this.reasonHooks) this.reasonHooks.finish();
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
    this.log.innerHTML = this.emptyHTML;
    this.empty = this.log.querySelector("[data-empty]");
    this.lastSeq = 0;
    this.assistant = null;
    this.assistantBody = null;
    this.assistantText = "";
    this.assistantStarted = false;
    this.reasoningEl = null;
    this.reasoningText = "";
    this.toolBoxes.clear();
    this.approvalCards.clear();
    this.searchStatus = null;
    this.generating = false;
    this.turnStartTs = 0;
    this.turnStats = null;
    this.turnRoute = null;
    this.turnElapsedMs = null;
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
    if (ev.reset) {
      this.reset();
      return;
    }
    if (ev.pad !== undefined) return;
    if (ev.caught_up) {
      this.toBottom();
      return;
    }
    if (ev.user !== undefined) {
      this.addUser(String(ev.user));
      this.resetAssistantState();
      this.turnStartTs = Date.now();
      this.turnStats = null;
      this.turnRoute = null;
      this.turnElapsedMs = null;
      if (this.reasonPanel) resetReasonPanel();
      else if (this.reasonHooks) this.reasonHooks.reset();
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
      this.addTool(ev.tool);
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
    if (ev.worktree !== undefined) {
      const wt = ev.worktree || {};
      this.addSystem("🌿 Worktree isolé : " + (wt.path || ""));
      return;
    }
    if (ev.worktree_error !== undefined) {
      this.addSystem("⚠️ Worktree indisponible (" + String(ev.worktree_error) + ") — repli sur le workspace partagé.");
      return;
    }
    if (ev.stats !== undefined) {
      const s = ev.stats || {};
      this.turnStats = s;
      if (this.trackTokens) setTurnStats(s.prompt_tokens, s.completion_tokens);
      if (this.statsBadge) {
        this.statsBadge.hidden = false;
        this.statsBadge.textContent = "↑" + (s.prompt_tokens || 0) + " ↓" + (s.completion_tokens || 0);
      }
      return;
    }
    if (ev.compact) {
      this.addSystem("Contexte compacté pour rester dans la fenêtre du modèle.");
      return;
    }
    if (ev.route !== undefined) {
      const r = ev.route || {};
      this.turnRoute = r;
      if (this.routeBadge) {
        this.routeBadge.hidden = false;
        const name = r.label || r.model || "";
        this.routeBadge.textContent =
          (r.provider ? r.provider + "/" : "") + name + (r.local ? " (local)" : "") + (r.fallback ? " (repli)" : "");
      }
      return;
    }
    if (ev.drop_reasoning) {
      this.removeReasoning();
      return;
    }
    if (ev.error !== undefined) {
      this.addError(String(ev.error));
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
    const url = this.streamURL(this.lastSeq);
    const run = async () => {
      try {
        const resp = await fetch(url, {
          headers: { Authorization: "Bearer " + getToken() },
          signal: this.controller.signal,
        });
        if (!resp.ok || !resp.body) return;
        await readSSE(resp, (ev) => this.handleEvent(ev));
      } catch (e) {
        if (e && e.name === "AbortError") return;
        setTimeout(() => {
          if (this.controller && !this.controller.signal.aborted) run();
        }, 1500);
      }
    };
    run();
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
    const payload = this.getPayload(text);
    try {
      await api(this.sendURL, { method: "POST", body: payload });
      this.generating = true;
      if (this.stopBtn) this.stopBtn.hidden = false;
      this.setBusy(true);
      this.toBottom();
      return true;
    } catch (err) {
      if (/en cours/i.test(err.message)) {
        this.generating = true;
        if (this.stopBtn) this.stopBtn.hidden = false;
        this.setBusy(true);
        return true;
      }
      this.addError(err.message);
      return false;
    }
  }

  stop() {
    if (this.stopURL) api(this.stopURL, { method: "POST" }).catch(() => {});
  }
}
