// Panneau flottant "Requêtes" façon DeepSeek : sommaire des requêtes
// (messages utilisateur) de la CONVERSATION EN COURS, affiché sur le
// côté du chat avec scroll haut/bas. Clic sur une requête = défilement
// vers ce message. La requête la plus proche du haut est surlignée.
// Vue principale uniquement (la vue Agents reste identique à Marexcode).
export function initRequestsPanel() {
  const main = document.querySelector("main.main");
  const log = document.getElementById("chat-container");
  if (!main || !log) return null;
  if (main.querySelector(":scope > .requests-panel")) return null;
  if (getComputedStyle(main).position === "static") main.style.position = "relative";

  const panel = document.createElement("aside");
  panel.className = "requests-panel";
  panel.setAttribute("aria-label", "Requêtes");
  panel.hidden = true;
  main.appendChild(panel);

  /** @type {{btn: HTMLButtonElement, target: Element}[]} */
  const items = [];

  function labelFor(userEl) {
    const txt =
      userEl.querySelector(".message-text")?.textContent || userEl.textContent || "";
    return txt.trim().split("\n")[0] || "(requête)";
  }

  function rebuild() {
    const users = log.querySelectorAll(".message-user");
    panel.innerHTML = "";
    items.length = 0;
    panel.hidden = users.length === 0;
    users.forEach((u) => {
      const label = labelFor(u);
      const b = document.createElement("button");
      b.type = "button";
      b.className = "request-item";
      b.textContent = label;
      b.title = label;
      b.addEventListener("click", () => {
        if (u.scrollIntoView) u.scrollIntoView({ behavior: "smooth", block: "start" });
      });
      panel.appendChild(b);
      items.push({ btn: b, target: u });
    });
    updateActive();
  }

  function updateActive() {
    if (!items.length) return;
    const top = log.getBoundingClientRect().top;
    let best = null;
    let bestDist = Infinity;
    for (const it of items) {
      if (!it.target.isConnected) continue;
      const d = Math.abs(it.target.getBoundingClientRect().top - top);
      if (d < bestDist) {
        bestDist = d;
        best = it;
      }
    }
    for (const it of items) it.btn.classList.toggle("active", it === best);
  }

  let raf = 0;
  function onScroll() {
    if (raf) return;
    const r = window.requestAnimationFrame || ((fn) => setTimeout(fn, 16));
    raf = r(() => {
      raf = 0;
      updateActive();
    });
  }
  log.addEventListener("scroll", onScroll, { passive: true });

  // Nouveaux messages / élagage du DOM : on reconstruit le sommaire.
  const mo = new MutationObserver(() => rebuild());
  mo.observe(log, { childList: true });

  window.addEventListener("cetas:chat-reset", rebuild);
  rebuild();
  return { refresh: rebuild, panel };
}
