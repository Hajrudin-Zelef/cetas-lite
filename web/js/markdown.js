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
    if (pre.querySelector(".code-copy")) continue;
    const btn = document.createElement("button");
    btn.type = "button";
    btn.className = "code-copy";
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
