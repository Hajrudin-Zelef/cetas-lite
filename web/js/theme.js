import { putPrefs } from "./api.js";

export function initTheme() {
  const sel = document.getElementById("theme-select");
  const current = document.documentElement.dataset.theme || "ocean";
  sel.value = current;
  sel.addEventListener("change", () => {
    const value = sel.value;
    document.documentElement.dataset.theme = value;
    try {
      localStorage.setItem("cetas-lite-theme", value);
    } catch (e) {}
    putPrefs({
      theme: value,
      family: (document.getElementById("family-select") || {}).value || "",
      mode: (document.getElementById("mode-select") || {}).value || "",
      agent_default: false,
    }).catch(() => {});
  });
}
