import { persistPrefs } from "./model-select.js";

const VALID = ["ocean", "clair", "sombre"];

export function applyTheme(theme) {
  const value = VALID.includes(theme) ? theme : "ocean";
  document.documentElement.dataset.theme = value;
  document.body.className = value === "sombre" ? "ocean-theme dark" : value === "ocean" ? "ocean-theme" : "";
  try {
    localStorage.setItem("cetas-lite-theme", value);
  } catch (e) {}
}

export function initTheme() {
  const sel = document.getElementById("theme-select");
  if (!sel) return;
  const current = document.documentElement.dataset.theme || "ocean";
  sel.value = current;
  sel.addEventListener("change", () => {
    applyTheme(sel.value);
    persistPrefs().catch(() => {});
  });
}
