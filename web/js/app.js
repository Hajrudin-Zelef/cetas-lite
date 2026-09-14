import { initAuth } from "./auth.js";
import { initChat } from "./chat.js";
import { initConversations } from "./conversations.js";
import { initModels } from "./model-select.js";
import { initSettings } from "./settings.js";
import { initTheme } from "./theme.js";
import { initAgents } from "./agents.js";
import { initTerminal } from "./terminal.js";

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
  initConversations();
  initAgents();
  initTerminal();

  const sidebar = document.getElementById("sidebar");
  const toggle = document.getElementById("sidebar-toggle");
  const mobile = window.matchMedia("(max-width: 768px)");

  function applySidebar(open) {
    sidebar.classList.toggle("collapsed", !open);
    toggle.classList.toggle("collapsed", !open);
    document.body.classList.toggle("sidebar-open", open && mobile.matches);
    toggle.setAttribute("aria-expanded", open ? "true" : "false");
  }

  applySidebar(!mobile.matches);
  toggle.addEventListener("click", () => applySidebar(sidebar.classList.contains("collapsed")));
  sidebar.addEventListener("click", (e) => {
    if (mobile.matches && e.target.closest(".conv-item")) applySidebar(false);
  });
  window.addEventListener("keydown", (e) => {
    if (e.key === "Escape") applySidebar(false);
  });
});
