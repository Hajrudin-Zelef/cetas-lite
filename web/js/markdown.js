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
