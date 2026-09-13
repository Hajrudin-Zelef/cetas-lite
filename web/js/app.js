import { initAuth } from "./auth.js";
import { initChat } from "./chat.js";
import { initModels } from "./model-select.js";
import { initSettings } from "./settings.js";
import { initTheme } from "./theme.js";

initTheme();

initAuth(async () => {
  let reloadModels = null;
  try {
    reloadModels = await initModels();
  } catch (e) {
    console.error("chargement des alias impossible", e);
  }
  initSettings({ reloadModels });
  initChat();

  const sidebar = document.getElementById("sidebar");
  const scrim = document.getElementById("sidebar-scrim");
  const toggle = document.getElementById("sidebar-toggle");

  function setSidebar(open) {
    sidebar.classList.toggle("open", open);
    scrim.classList.toggle("show", open);
    scrim.hidden = !open;
    toggle.setAttribute("aria-expanded", open ? "true" : "false");
  }

  toggle.addEventListener("click", () => setSidebar(!sidebar.classList.contains("open")));
  scrim.addEventListener("click", () => setSidebar(false));
  window.addEventListener("keydown", (e) => {
    if (e.key === "Escape") setSidebar(false);
  });
});
