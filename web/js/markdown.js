const marked = window.marked;
const purify = window.DOMPurify;

if (marked && marked.setOptions) {
  marked.setOptions({ gfm: true, breaks: true });
}

const ESCAPES = { "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;" };

function escapeHtml(text) {
  return String(text).replace(/[&<>"]/g, (c) => ESCAPES[c]);
}

export function renderMarkdown(text) {
  if (!marked || !marked.parse) {
    return "<p>" + escapeHtml(text) + "</p>";
  }
  const html = marked.parse(String(text || ""));
  return purify && purify.sanitize ? purify.sanitize(html) : html;
}

function wrapTables(root) {
  for (const table of root.querySelectorAll("table")) {
    if (table.parentElement && table.parentElement.classList.contains("table-wrap")) continue;
    const wrap = document.createElement("div");
    wrap.className = "table-wrap";
    table.parentElement.insertBefore(wrap, table);
    wrap.appendChild(table);
  }
}

function addCopyButtons(root) {
  for (const pre of root.querySelectorAll("pre")) {
    if (pre.querySelector(".code-copy-btn")) continue;
    const btn = document.createElement("button");
    btn.type = "button";
    btn.className = "code-copy-btn";
    btn.textContent = "Copier";
    btn.addEventListener("click", () => {
      const code = pre.querySelector("code");
      const text = code ? code.innerText : pre.innerText;
      if (!navigator.clipboard) return;
      navigator.clipboard.writeText(text).then(() => {
        btn.textContent = "Copie";
        setTimeout(() => {
          btn.textContent = "Copier";
        }, 1200);
      }).catch(() => {});
    });
    pre.appendChild(btn);
  }
}

export function renderInto(el, text) {
  el.innerHTML = renderMarkdown(text);
  wrapTables(el);
  addCopyButtons(el);
}

const URL_RE = /(https?:\/\/[^\s<>()"']+)/g;

export function appendLinkified(el, text) {
  const s = String(text || "");
  let last = 0;
  URL_RE.lastIndex = 0;
  let m;
  while ((m = URL_RE.exec(s)) !== null) {
    if (m.index > last) el.appendChild(document.createTextNode(s.slice(last, m.index)));
    const a = document.createElement("a");
    a.href = m[1];
    a.textContent = m[1];
    a.target = "_blank";
    a.rel = "noopener noreferrer";
    el.appendChild(a);
    last = m.index + m[1].length;
  }
  if (last < s.length) el.appendChild(document.createTextNode(s.slice(last)));
}

// ---------------------------------------------------------------------------
// Rendu streaming accelere (technique Marexcode) :
// - update() coalesce les deltas sur une seule frame (rAF) par conteneur :
//   jamais plus d'un rendu markdown par frame, meme si 50 deltas arrivent.
// - rendu incremental par blocs stables : le texte est decoupe en blocs
//   separes par une ligne vide hors fence, plus un "tail" live. Seul le tail
//   et les blocs modifies sont re-rendus (memoization via _html) -> pas de
//   re-parse O(n) de tout le message a chaque frame.
// ---------------------------------------------------------------------------

function streamRaf(fn) {
  if (typeof requestAnimationFrame === "function") return requestAnimationFrame(fn);
  return setTimeout(() => fn(typeof performance !== "undefined" && performance.now ? performance.now() : Date.now()), 16);
}

function streamCancelRaf(id) {
  if (typeof cancelAnimationFrame === "function") cancelAnimationFrame(id);
  else clearTimeout(id);
}

// Decoupe le markdown en blocs stables separes par une ligne vide hors
// fence cloture, plus un tail "live" (dernier bloc, susceptible de changer).
export function splitBlocks(text) {
  const lines = String(text == null ? "" : text).split("\n");
  const blocks = [];
  let buf = [];
  let fence = null;
  for (const line of lines) {
    const m = line.match(/^ {0,3}(`{3,}|~{3,})/);
    if (m) {
      if (!fence) fence = m[1][0];
      else if (m[1][0] === fence) fence = null;
    }
    buf.push(line);
    if (!fence && line.trim() === "") {
      const block = buf.join("\n");
      if (block.trim()) blocks.push(block);
      buf = [];
    }
  }
  return { blocks, tail: buf.join("\n") };
}

export function createMarkdownRenderer() {
  const pending = new WeakMap(); // container -> texte integral connu
  const frames = new WeakMap(); // container -> id rAF en cours

  function childAt(container, index) {
    let child = container.children[index];
    if (!child) {
      child = document.createElement("div");
      child.className = "md-block";
      child._html = null;
      container.appendChild(child);
    }
    return child;
  }

  function setBlock(container, index, html) {
    const child = childAt(container, index);
    if (child._html !== html) {
      child.innerHTML = html;
      child._html = html;
      wrapTables(child);
      addCopyButtons(child);
    }
  }

  function renderSync(container, text, streaming) {
    const value = String(text == null ? "" : text);
    const { blocks, tail } = splitBlocks(value);
    let total = 0;
    for (let i = 0; i < blocks.length; i++, total++) {
      setBlock(container, i, renderMarkdown(blocks[i]));
    }
    if (streaming || tail.trim()) {
      setBlock(container, total, renderMarkdown(tail));
      total++;
    }
    while (container.children.length > total) container.removeChild(container.lastElementChild);
  }

  function cancel(container) {
    const f = frames.get(container);
    if (f) {
      streamCancelRaf(f);
      frames.set(container, null);
    }
  }

  return {
    // Rendu synchrone complet (remplace le contenu).
    render(container, text) {
      cancel(container);
      pending.set(container, String(text == null ? "" : text));
      renderSync(container, text, false);
    },
    // Streaming : bufferise et rend au plus une fois par frame.
    update(container, text) {
      pending.set(container, String(text == null ? "" : text));
      if (frames.get(container)) return;
      frames.set(
        container,
        streamRaf(() => {
          frames.set(container, null);
          renderSync(container, pending.get(container), true);
        })
      );
    },
    // Fin de stream : annule la frame en attente et rend tout, synchrone.
    finalize(container, text) {
      cancel(container);
      const value = text === undefined ? pending.get(container) || "" : text;
      pending.set(container, String(value));
      renderSync(container, value, false);
    },
    // Texte integral connu pour ce conteneur.
    text(container) {
      return pending.get(container) || "";
    },
  };
}
