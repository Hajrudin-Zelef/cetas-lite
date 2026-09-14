import { api, getPrefs, putPrefs } from "./api.js";

export function applyWebToggle(on) {
  const el = document.getElementById("web-toggle");
  if (!el) return;
  el.setAttribute("aria-pressed", on ? "true" : "false");
  el.classList.toggle("active", !!on);
}

export function applyMCPToggle(on) {
  const el = document.getElementById("mcp-toggle");
  if (!el) return;
  el.setAttribute("aria-pressed", on ? "true" : "false");
  el.classList.toggle("active", !!on);
}

let thinkingPref = false;

export function applyThinkingToggle(on) {
  const el = document.getElementById("thinking-toggle");
  if (!el || el.disabled) return;
  el.setAttribute("aria-pressed", on ? "true" : "false");
  el.classList.toggle("active", !!on);
}

export function setThinking(on) {
  thinkingPref = !!on;
  applyThinkingToggle(thinkingPref);
}

export function persistPrefs() {
  const body = {};
  const theme = document.documentElement.dataset.theme;
  if (theme) body.theme = theme;
  const familySel = document.getElementById("family-select");
  if (familySel && familySel.value) body.family = familySel.value;
  const modeSel = document.getElementById("mode-select");
  if (modeSel && modeSel.value) body.mode = modeSel.value;
  const webToggle = document.getElementById("web-toggle");
  if (webToggle) body.web_default = webToggle.getAttribute("aria-pressed") === "true";
  const mcpToggle = document.getElementById("mcp-toggle");
  if (mcpToggle && !mcpToggle.hidden) {
    body.mcp_default = mcpToggle.getAttribute("aria-pressed") === "true";
  }
  body.thinking_default = thinkingPref;
  const effort = document.getElementById("effort-select");
  if (effort && effort.value) body.thinking_effort = effort.value;
  return putPrefs(body);
}

export async function initModels() {
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  const agentPill = document.getElementById("agent-toggle");
  let families = [];

  function modesFor(id) {
    const f = families.find((x) => x.id === id);
    return f ? f.modes : [];
  }

  function updateAgent() {
    const m = modesFor(familySel.value).find((x) => x.mode === modeSel.value);
    const agent = !!(m && m.agent);
    agentPill.hidden = !agent;
    const th = document.getElementById("thinking-toggle");
    if (th) {
      if (agent) {
        th.disabled = true;
        th.setAttribute("aria-pressed", "true");
        th.classList.add("active");
        th.title = "Raisonnement (actif en mode Agent)";
      } else {
        th.disabled = false;
        th.title = "Raisonnement";
        applyThinkingToggle(thinkingPref);
      }
    }
  }

  function persist() {
    persistPrefs().catch(() => {});
  }

  function renderModes() {
    modeSel.innerHTML = "";
    for (const m of modesFor(familySel.value)) {
      const opt = document.createElement("option");
      opt.value = m.mode;
      opt.textContent = m.label + (m.rule ? " · " + m.rule : "");
      modeSel.appendChild(opt);
    }
    updateAgent();
  }

  familySel.addEventListener("change", () => {
    renderModes();
    persist();
  });
  modeSel.addEventListener("change", () => {
    updateAgent();
    persist();
  });

  function applyTheme(theme) {
    if (!["ocean", "sombre", "clair"].includes(theme)) return;
    document.documentElement.dataset.theme = theme;
    const sel = document.getElementById("theme-select");
    if (sel) sel.value = theme;
    try {
      localStorage.setItem("cetas-lite-theme", theme);
    } catch (e) {}
  }

  async function load() {
    const data = await api("/api/aliases");
    families = data.families || [];
    const prefs = await getPrefs().catch(() => null);
    if (prefs && prefs.theme) applyTheme(prefs.theme);
    applyWebToggle(!!(prefs && prefs.web_default));
    thinkingPref = !!(prefs && prefs.thinking_default);
    applyThinkingToggle(thinkingPref);
    const effortSel = document.getElementById("effort-select");
    if (effortSel) effortSel.value = (prefs && prefs.thinking_effort) || "default";

    const mcp = await api("/api/mcp").catch(() => null);
    const mcpToggle = document.getElementById("mcp-toggle");
    if (mcpToggle) {
      const servers = (mcp && mcp.servers) || [];
      if (servers.length > 0) {
        mcpToggle.hidden = false;
        const on = prefs && typeof prefs.mcp_default === "boolean" ? prefs.mcp_default : true;
        applyMCPToggle(on);
      } else {
        mcpToggle.hidden = true;
        applyMCPToggle(false);
      }
    }
    const prevF = familySel.value || (prefs && prefs.family) || "";
    const prevM = modeSel.value || (prefs && prefs.mode) || "";
    familySel.innerHTML = "";
    for (const f of families) {
      const opt = document.createElement("option");
      opt.value = f.id;
      opt.textContent = f.label;
      familySel.appendChild(opt);
    }
    if (prevF && families.some((f) => f.id === prevF)) familySel.value = prevF;
    renderModes();
    if (prevM && modesFor(familySel.value).some((m) => m.mode === prevM)) {
      modeSel.value = prevM;
      updateAgent();
    }
  }

  await load();
  return load;
}
