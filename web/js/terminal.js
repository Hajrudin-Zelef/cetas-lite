import { api, getToken, readSSE } from "./api.js";
import { el } from "./thread-view.js";

// Terminal integre : tiroir bas avec onglets, xterm.js, PTY cote serveur.
export function initTerminal() {
  const toolbarBtn = document.getElementById("terminal-btn");
  if (!toolbarBtn) return;
  if (typeof Terminal === "undefined") {
    toolbarBtn.disabled = true;
    toolbarBtn.title = "Terminal indisponible (xterm.js introuvable)";
    return;
  }

  const drawer = document.createElement("div");
  drawer.id = "terminal-drawer";
  drawer.className = "terminal-drawer";
  drawer.hidden = true;
  document.body.appendChild(drawer);

  const bar = el("div", "terminal-bar");
  const tabs = el("div", "terminal-tabs");
  bar.appendChild(tabs);
  const newBtn = el("button", "mini-btn", "+ Nouveau");
  newBtn.type = "button";
  newBtn.addEventListener("click", () => createSession());
  bar.appendChild(newBtn);
  const closeBtn = el("button", "mini-btn", "Fermer");
  closeBtn.type = "button";
  closeBtn.addEventListener("click", toggleDrawer);
  bar.appendChild(closeBtn);
  drawer.appendChild(bar);

  const viewport = el("div", "terminal-viewport");
  drawer.appendChild(viewport);

  // sessions: id -> { id, cwd, term, fit, container, controller, alive }
  const sessions = new Map();
  let activeId = null;
  let fitObserver = null;

  function decodeB64(b64) {
    const bin = atob(b64);
    const bytes = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; i++) bytes[i] = bin.charCodeAt(i);
    return bytes;
  }

  function shortCwd(cwd) {
    if (!cwd) return "shell";
    const parts = String(cwd).split("/").filter(Boolean);
    return parts.length ? parts[parts.length - 1] : cwd;
  }

  function makeTerm(sess) {
    const container = el("div", "terminal-instance");
    container.hidden = true;
    viewport.appendChild(container);
    const term = new Terminal({
      cursorBlink: true,
      fontSize: 13,
      fontFamily: "ui-monospace, 'Cascadia Code', Menlo, Consolas, monospace",
      theme: { background: "#0d1117", foreground: "#e6edf3", cursor: "#58a6ff" },
      scrollback: 5000,
    });
    const fit = new (window.FitAddon ? window.FitAddon.FitAddon : function () {
      this.fit = function () {};
      this.activate = function () {};
    })();
    try {
      term.loadAddon(fit);
    } catch (e) {}
    term.open(container);
    term.onData((data) => {
      api("/api/terminal/" + encodeURIComponent(sess.id) + "/input", {
        method: "POST",
        body: { data },
      }).catch(() => {});
    });
    sess.term = term;
    sess.fit = fit;
    sess.container = container;
  }

  function fitActive() {
    const sess = sessions.get(activeId);
    if (!sess || !sess.term || !sess.fit) return;
    try {
      sess.fit.fit();
    } catch (e) {
      return;
    }
    const cols = sess.term.cols;
    const rows = sess.term.rows;
    if (cols && rows) {
      api("/api/terminal/" + encodeURIComponent(sess.id) + "/resize", {
        method: "POST",
        body: { cols, rows },
      }).catch(() => {});
    }
  }

  function subscribe(sess) {
    if (sess.controller) return;
    sess.controller = new AbortController();
    const url = "/api/terminal/" + encodeURIComponent(sess.id) + "/stream";
    fetch(url, {
      headers: { Authorization: "Bearer " + getToken() },
      signal: sess.controller.signal,
    })
      .then((resp) => {
        if (!resp.ok || !resp.body) throw new Error("stream");
        return readSSE(resp, (ev) => {
          if (ev.output && sess.term) {
            try {
              sess.term.write(decodeB64(ev.output));
            } catch (e) {}
          }
          if (ev.exited) {
            sess.alive = false;
            refreshTabs();
          }
        });
      })
      .catch((e) => {
        sess.controller = null;
        if (e && e.name === "AbortError") return;
        // Reconnexion douce si le tiroir est ouvert.
        if (!drawer.hidden && sessions.has(sess.id)) {
          setTimeout(() => subscribe(sess), 2000);
        }
      });
  }

  function unsubscribe(sess) {
    if (sess.controller) {
      sess.controller.abort();
      sess.controller = null;
    }
  }

  function activate(id) {
    activeId = id;
    for (const [sid, sess] of sessions) {
      const on = sid === id;
      if (on && !sess.term) makeTerm(sess);
      if (sess.container) sess.container.hidden = !on;
      if (on) subscribe(sess);
    }
    refreshTabs();
    requestAnimationFrame(() => {
      fitActive();
      const sess = sessions.get(id);
      if (sess && sess.term) sess.term.focus();
    });
  }

  function refreshTabs() {
    tabs.innerHTML = "";
    for (const [sid, sess] of sessions) {
      const tab = el("button", "terminal-tab" + (sid === activeId ? " active" : ""));
      tab.type = "button";
      const dot = el("span", "terminal-dot" + (sess.alive === false ? " dead" : ""));
      tab.appendChild(dot);
      tab.appendChild(el("span", "", shortCwd(sess.cwd)));
      const x = el("span", "terminal-tab-close", "×");
      x.title = "Fermer ce terminal";
      x.addEventListener("click", (e) => {
        e.stopPropagation();
        killSession(sid);
      });
      tab.appendChild(x);
      tab.addEventListener("click", () => activate(sid));
      tabs.appendChild(tab);
    }
  }

  async function createSession() {
    try {
      const res = await api("/api/terminal", { method: "POST", body: {} });
      const sess = { id: res.id, cwd: res.cwd, term: null, fit: null, container: null, controller: null, alive: true };
      sessions.set(res.id, sess);
      activate(res.id);
    } catch (e) {
      // Erreur visible dans le tiroir.
      const err = el("div", "terminal-error", e.message);
      viewport.appendChild(err);
      setTimeout(() => err.remove(), 4000);
    }
  }

  async function killSession(id) {
    const sess = sessions.get(id);
    if (sess) {
      unsubscribe(sess);
      if (sess.term) sess.term.dispose();
      if (sess.container) sess.container.remove();
      sessions.delete(id);
    }
    try {
      await api("/api/terminal/" + encodeURIComponent(id), { method: "DELETE" });
    } catch (e) {}
    if (activeId === id) {
      activeId = null;
      const next = sessions.keys().next();
      if (!next.done) activate(next.value);
      else refreshTabs();
    } else {
      refreshTabs();
    }
  }

  async function syncSessions() {
    try {
      const data = await api("/api/terminal");
      const list = data.sessions || [];
      const seen = new Set();
      for (const s of list) {
        seen.add(s.id);
        if (!sessions.has(s.id)) {
          sessions.set(s.id, { id: s.id, cwd: s.cwd, term: null, fit: null, container: null, controller: null, alive: true });
        }
      }
      for (const sid of [...sessions.keys()]) {
        if (!seen.has(sid)) {
          const sess = sessions.get(sid);
          unsubscribe(sess);
          if (sess.term) sess.term.dispose();
          if (sess.container) sess.container.remove();
          sessions.delete(sid);
          if (activeId === sid) activeId = null;
        }
      }
      if (!activeId) {
        const next = sessions.keys().next();
        if (!next.done) activate(next.value);
        else if (list.length === 0) await createSession();
      } else {
        refreshTabs();
      }
    } catch (e) {}
  }

  function toggleDrawer() {
    drawer.hidden = !drawer.hidden;
    toolbarBtn.classList.toggle("active", !drawer.hidden);
    if (!drawer.hidden) {
      syncSessions().then(() => fitActive());
      if (!fitObserver && typeof ResizeObserver !== "undefined") {
        fitObserver = new ResizeObserver(() => fitActive());
        fitObserver.observe(viewport);
      }
      window.addEventListener("resize", fitActive);
    } else {
      window.removeEventListener("resize", fitActive);
      for (const sess of sessions.values()) unsubscribe(sess);
      const sess = sessions.get(activeId);
      if (sess && sess.term) sess.term.blur();
    }
  }

  toolbarBtn.addEventListener("click", toggleDrawer);
}
