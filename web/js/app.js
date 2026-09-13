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
  document.getElementById("sidebar-toggle").addEventListener("click", () => {
    document.getElementById("sidebar").classList.toggle("open");
  });
});
