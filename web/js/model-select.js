import { api } from "./api.js";

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
    agentPill.hidden = !(m && m.agent);
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

  familySel.addEventListener("change", renderModes);
  modeSel.addEventListener("change", updateAgent);

  async function load() {
    const data = await api("/api/aliases");
    families = data.families || [];
    const prevF = familySel.value;
    const prevM = modeSel.value;
    familySel.innerHTML = "";
    for (const f of families) {
      const opt = document.createElement("option");
      opt.value = f.id;
      opt.textContent = f.label;
      familySel.appendChild(opt);
    }
    if (prevF && families.some((f) => f.id === prevF)) familySel.value = prevF;
    renderModes();
    if (prevM) {
      modeSel.value = prevM;
      updateAgent();
    }
  }

  await load();
  return load;
}
