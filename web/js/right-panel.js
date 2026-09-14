const ROLES_KEY = "cetas-lite-roles";

function loadRoles() {
  try {
    return JSON.parse(localStorage.getItem(ROLES_KEY) || "[]");
  } catch (e) {
    return [];
  }
}
export function saveRoles(roles) {
  try {
    localStorage.setItem(ROLES_KEY, JSON.stringify(roles));
  } catch (e) {}
  window.dispatchEvent(new CustomEvent("cetas:roles-changed"));
}

export function initRightPanel() {
  const panel = document.getElementById("right-panel");
  const settingsBtn = document.getElementById("chat-header-settings");

  function toggle(force) {
    if (!panel) return;
    const show = force !== undefined ? force : panel.hidden;
    panel.hidden = !show;
    panel.style.display = show ? "" : "none";
    document.body.classList.toggle("right-panel-open", show);
  }

  if (settingsBtn) {
    settingsBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      toggle();
    });
  }
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape" && panel && !panel.hidden) toggle(false);
  });

  // Rôles (snippets locaux)
  const spSelect = document.getElementById("sp-select");
  const spTextarea = document.getElementById("sp-textarea");
  const spActions = document.getElementById("rp-role-actions");
  const spApply = document.getElementById("sp-apply-btn");

  function renderRoles() {
    if (!spSelect) return;
    const roles = loadRoles();
    spSelect.innerHTML = '<option value="">Sélectionner un rôle enregistré</option>';
    roles.forEach((r, i) => {
      const opt = document.createElement("option");
      opt.value = String(i);
      opt.textContent = r.name;
      spSelect.appendChild(opt);
    });
  }

  if (spSelect && spTextarea) {
    spSelect.addEventListener("change", () => {
      const roles = loadRoles();
      const r = roles[parseInt(spSelect.value, 10)];
      spTextarea.value = r ? r.content : "";
      if (spActions) spActions.style.display = spTextarea.value ? "" : "none";
    });
    spTextarea.addEventListener("input", () => {
      if (spActions) spActions.style.display = spTextarea.value ? "" : "none";
    });
  }
  if (spApply && spTextarea) {
    spApply.addEventListener("click", () => {
      const input = document.getElementById("prompt-input");
      if (input && spTextarea.value) {
        input.value = spTextarea.value + (input.value ? "\n\n" + input.value : "");
        input.dispatchEvent(new Event("input"));
        input.focus();
        toggle(false);
      }
    });
  }

  window.addEventListener("cetas:roles-changed", renderRoles);
  renderRoles();

  return { toggle };
}
