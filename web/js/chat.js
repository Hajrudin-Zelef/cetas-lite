import { api, getToken, readSSE } from "./api.js";

export function initChat() {
  const log = document.getElementById("chat-log");
  const empty = document.getElementById("empty-chat");
  const input = document.getElementById("prompt-input");
  const form = document.getElementById("composer");
  const stopBtn = document.getElementById("stop-btn");
  const routeBadge = document.getElementById("route-badge");

  let lastSeq = 0;
  let assistant = null;
  let reasoning = null;
  let controller = null;
  let generating = false;
  const toolBoxes = new Map();

  function el(tag, cls, text) {
    const e = document.createElement(tag);
    if (cls) e.className = cls;
    if (text !== undefined) e.textContent = text;
    return e;
  }

  function clearEmpty() {
    if (empty && empty.parentNode) empty.remove();
  }

  function atBottom() {
    return log.scrollHeight - log.scrollTop - log.clientHeight < 80;
  }

  function scroll(force) {
    if (force || atBottom()) log.scrollTop = log.scrollHeight;
  }

  function selection() {
    return {
      family: document.getElementById("family-select").value,
      mode: document.getElementById("mode-select").value,
    };
  }

  function addUser(text) {
    clearEmpty();
    log.appendChild(el("div", "msg user", text));
    scroll(true);
  }

  function ensureAssistant() {
    if (!assistant) {
      clearEmpty();
      assistant = el("div", "msg assistant");
      log.appendChild(assistant);
    }
    return assistant;
  }

  function appendContent(text) {
    ensureAssistant().textContent += text;
    scroll();
  }

  function replaceContent(text) {
    ensureAssistant().textContent = text;
    scroll();
  }

  function appendReasoning(text) {
    if (!reasoning) {
      ensureAssistant();
      reasoning = el("div", "msg-reasoning");
      log.insertBefore(reasoning, assistant);
    }
    reasoning.textContent += text;
    scroll();
  }

  function removeReasoning() {
    if (reasoning) {
      reasoning.remove();
      reasoning = null;
    }
  }

  function addError(text) {
    clearEmpty();
    log.appendChild(el("div", "msg error", text));
    scroll();
  }

  function summarize(args) {
    if (!args) return "";
    return args.command || args.file_path || args.pattern || args.query || "";
  }

  function addTool(ev) {
    const key = ev.name + "|" + JSON.stringify(ev.args || {});
    if (ev.phase === "start") {
      clearEmpty();
      const box = el("details", "msg-tool");
      box.appendChild(el("summary", null, ev.name + (summarize(ev.args) ? " " + summarize(ev.args) : "")));
      const pre = el("pre");
      box.appendChild(pre);
      log.appendChild(box);
      toolBoxes.set(key, pre);
      assistant = null;
      reasoning = null;
      scroll(true);
      return;
    }
    const pre = toolBoxes.get(key);
    if (!pre) return;
    if (typeof ev.result === "string" && ev.result) pre.textContent = ev.result;
    if (Array.isArray(ev.diff)) {
      for (const line of ev.diff) {
        const cls = line.kind === "+" ? "diff-add" : "diff-del";
        pre.appendChild(el("span", cls, (line.kind === "+" ? "+ " : "- ") + line.text + "\n"));
      }
    }
    scroll();
  }

  function finishTurn() {
    generating = false;
    stopBtn.hidden = true;
    assistant = null;
    reasoning = null;
  }

  function resetLocal() {
    log.innerHTML = "";
    lastSeq = 0;
    assistant = null;
    reasoning = null;
    toolBoxes.clear();
    generating = false;
    stopBtn.hidden = true;
    routeBadge.hidden = true;
  }

  function handleEvent(ev) {
    if (typeof ev.seq === "number" && ev.seq > lastSeq) lastSeq = ev.seq;
    if (ev.reset) {
      resetLocal();
      return;
    }
    if (ev.pad !== undefined || ev.caught_up) return;
    if (ev.user !== undefined) {
      addUser(String(ev.user));
      assistant = null;
      reasoning = null;
      generating = true;
      stopBtn.hidden = false;
      return;
    }
    if (ev.reasoning_content !== undefined) {
      appendReasoning(String(ev.reasoning_content));
      return;
    }
    if (ev.content !== undefined) {
      const text = String(ev.content);
      if (ev.replace) replaceContent(text);
      else appendContent(text);
      return;
    }
    if (ev.tool !== undefined) {
      addTool(ev.tool);
      return;
    }
    if (ev.route !== undefined) {
      const r = ev.route || {};
      routeBadge.hidden = false;
      const name = r.label || r.model || "";
      routeBadge.textContent = (r.provider ? r.provider + "/" : "") + name + (r.local ? " (local)" : "") + (r.fallback ? " (repli)" : "");
      return;
    }
    if (ev.drop_reasoning) {
      removeReasoning();
      return;
    }
    if (ev.error !== undefined) {
      addError(String(ev.error));
      return;
    }
    if (ev.turn_done !== undefined) {
      finishTurn();
    }
  }

  async function connect() {
    if (controller) controller.abort();
    controller = new AbortController();
    try {
      const resp = await fetch("/api/chat/stream?from=" + lastSeq, {
        headers: { Authorization: "Bearer " + getToken() },
        signal: controller.signal,
      });
      if (!resp.ok || !resp.body) return;
      await readSSE(resp, handleEvent);
    } catch (e) {
      if (e && e.name === "AbortError") return;
      setTimeout(connect, 1500);
    }
  }

  async function send() {
    const text = input.value.trim();
    if (!text || generating) return;
    const sel = selection();
    if (!sel.family || !sel.mode) {
      addError("Choisis une famille et un mode.");
      return;
    }
    input.value = "";
    autoGrow();
    try {
      await api("/api/chat/send", { method: "POST", body: { family: sel.family, mode: sel.mode, message: text } });
      generating = true;
      stopBtn.hidden = false;
      scroll(true);
    } catch (err) {
      if (/en cours/i.test(err.message)) {
        generating = true;
        stopBtn.hidden = false;
      } else {
        addError(err.message);
      }
    }
  }

  function autoGrow() {
    input.style.height = "auto";
    input.style.height = Math.min(input.scrollHeight, 200) + "px";
  }

  form.addEventListener("submit", (e) => {
    e.preventDefault();
    send();
  });
  input.addEventListener("input", autoGrow);
  input.addEventListener("keydown", (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      send();
    }
  });
  stopBtn.addEventListener("click", () => {
    api("/api/chat/stop", { method: "POST" }).catch(() => {});
  });
  document.getElementById("new-chat-btn").addEventListener("click", () => {
    api("/api/chat/reset", { method: "POST" }).catch(() => {});
  });

  connect();
}
