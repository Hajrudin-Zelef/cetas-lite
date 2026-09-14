import { api } from "./api.js";
import { ThreadView, el } from "./thread-view.js";

// Panneau multi-agents : liste, creation, suivi, stop, suppression.
export function initAgents() {
  const toolbarBtn = document.getElementById("agents-btn");
  if (!toolbarBtn) return;

  const panel = document.createElement("aside");
  panel.id = "agents-panel";
  panel.className = "agents-panel";
  panel.hidden = true;
  panel.setAttribute("aria-label", "Agents parallèles");
  document.body.appendChild(panel);

  let families = [];
  let selectedId = null;
  let detailView = null;
  let detailEls = null; // { statusSlot, stopBtn }
  let lastAgents = [];
  let refreshTimer = null;
  let aliasesLoaded = false;

  async function loadAliases() {
    if (aliasesLoaded) return;
    try {
      const data = await api("/api/aliases");
      families = (data.families || []).filter((f) => (f.modes || []).some((m) => m.agent));
    } catch (e) {
      families = [];
    }
    aliasesLoaded = true;
  }

  function modeLabel(family, mode) {
    const f = families.find((x) => x.id === family);
    if (!f) return mode;
    const m = (f.modes || []).find((x) => x.mode === mode);
    return (f.label || family) + " · " + (m ? m.label : mode);
  }

  function statusBadge(status) {
    const s = el("span", "agent-status agent-status-" + status, {
      running: "En cours",
      done: "Terminé",
      stopped: "Arrêté",
      error: "Erreur",
    }[status] || status);
    return s;
  }

  function closeDetail() {
    if (detailView) {
      detailView.disconnect();
      detailView = null;
    }
    detailEls = null;
    selectedId = null;
    renderPanel(lastAgents);
  }

  function openDetail(id) {
    selectedId = id;
    // La vue detail est (re)creee une seule fois a la selection.
    renderPanel(lastAgents);
  }

  async function stopAgent(id) {
    try {
      await api("/api/agents/" + encodeURIComponent(id) + "/stop", { method: "POST" });
      refreshList();
    } catch (e) {}
  }

  async function deleteAgent(id) {
    if (!confirm("Supprimer cet agent et son worktree ?")) return;
    try {
      await api("/api/agents/" + encodeURIComponent(id), { method: "DELETE" });
      if (selectedId === id) closeDetail();
      else refreshList();
    } catch (e) {}
  }

  function agentCard(a) {
    const card = el("button", "agent-card" + (a.id === selectedId ? " selected" : ""));
    card.type = "button";
    const top = el("div", "agent-card-top");
    top.appendChild(el("span", "agent-card-id", "#" + String(a.id).slice(0, 8)));
    top.appendChild(statusBadge(a.status));
    card.appendChild(top);
    card.appendChild(el("div", "agent-card-model", modeLabel(a.family, a.mode)));
    if (a.worktree) card.appendChild(el("div", "agent-card-wt", "🌿 worktree isolé"));
    const meta = el("div", "agent-card-meta", (a.turns || 0) + " tour(s)");
    card.appendChild(meta);
    const acts = el("div", "agent-card-actions");
    const stopBtn = el("button", "mini-btn", "Stop");
    stopBtn.type = "button";
    stopBtn.disabled = a.status !== "running";
    stopBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      stopAgent(a.id);
    });
    const delBtn = el("button", "mini-btn mini-btn-danger", "Suppr.");
    delBtn.type = "button";
    delBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      deleteAgent(a.id);
    });
    acts.appendChild(stopBtn);
    acts.appendChild(delBtn);
    card.appendChild(acts);
    card.addEventListener("click", () => openDetail(a.id));
    return card;
  }

  function renderDetail(container, a) {
    const head = el("div", "agent-detail-head");
    const back = el("button", "mini-btn", "← Agents");
    back.type = "button";
    back.addEventListener("click", closeDetail);
    head.appendChild(back);
    head.appendChild(el("span", "agent-detail-title", "#" + String(a.id).slice(0, 8) + " · " + modeLabel(a.family, a.mode)));
    const statusSlot = el("span");
    statusSlot.appendChild(statusBadge(a.status));
    head.appendChild(statusSlot);
    const stopB = el("button", "mini-btn", "Stop");
    stopB.type = "button";
    stopB.disabled = a.status !== "running";
    stopB.addEventListener("click", () => stopAgent(a.id));
    head.appendChild(stopB);
    container.appendChild(head);
    detailEls = { statusSlot, stopBtn: stopB };

    if (a.worktree) {
      const wt = el("div", "agent-detail-wt", "🌿 " + a.worktree);
      wt.title = a.worktree;
      container.appendChild(wt);
    }

    const log = el("div", "agent-log");
    log.id = "agent-log-" + a.id;
    const emptyHTML = '<div class="agent-empty" data-empty>En attente des premiers messages de l’agent…</div>';
    log.innerHTML = emptyHTML;
    container.appendChild(log);

    const badges = el("div", "agent-badges");
    const routeBadge = el("span", "route-badge");
    routeBadge.hidden = true;
    const statsBadge = el("span", "stats-badge");
    statsBadge.hidden = true;
    badges.appendChild(routeBadge);
    badges.appendChild(statsBadge);
    container.appendChild(badges);

    const composer = el("form", "agent-composer");
    const input = el("textarea", "agent-input");
    input.placeholder = "Message de suivi pour cet agent…";
    input.rows = 2;
    const sendBtn = el("button", "btn-primary", "Envoyer");
    sendBtn.type = "submit";
    const stopBtn = el("button", "mini-btn", "Stop");
    stopBtn.type = "button";
    stopBtn.hidden = true;
    composer.appendChild(input);
    composer.appendChild(sendBtn);
    composer.appendChild(stopBtn);
    container.appendChild(composer);

    detailView = new ThreadView({
      log,
      stopBtn,
      routeBadge,
      statsBadge,
      emptyHTML,
      streamURL: (from) => "/api/agents/" + encodeURIComponent(a.id) + "/stream?from=" + from,
      sendURL: "/api/agents/" + encodeURIComponent(a.id) + "/message",
      stopURL: "/api/agents/" + encodeURIComponent(a.id) + "/stop",
      approveURL: "/api/agents/" + encodeURIComponent(a.id) + "/approve",
      getPayload: (text) => ({ message: text, approve: !!a.approve, plan: !!a.plan }),
      onDone: () => refreshList(),
    });
    detailView.connect();

    const send = async () => {
      const ok = await detailView.sendText(input.value);
      if (ok) input.value = "";
    };
    composer.addEventListener("submit", (e) => {
      e.preventDefault();
      send();
    });
    input.addEventListener("keydown", (e) => {
      if (e.key === "Enter" && !e.shiftKey) {
        e.preventDefault();
        send();
      }
    });
    stopBtn.addEventListener("click", () => detailView.stop());
  }

  // Rafraichissement doux : ne reconstruit jamais la vue detail ouverte,
  // met seulement a jour le badge de statut et la liste en arriere-plan.
  async function softRefresh() {
    if (panel.hidden) return;
    try {
      const data = await api("/api/agents");
      lastAgents = data.agents || [];
    } catch (e) {
      return;
    }
    if (selectedId) {
      const a = lastAgents.find((x) => x.id === selectedId);
      if (!a) {
        closeDetail();
        return;
      }
      if (detailEls) {
        detailEls.statusSlot.innerHTML = "";
        detailEls.statusSlot.appendChild(statusBadge(a.status));
        detailEls.stopBtn.disabled = a.status !== "running";
      }
      return;
    }
    renderListOnly(lastAgents);
  }

  function renderListOnly(agents) {
    const list = panel.querySelector(".agents-list");
    if (!list) return;
    list.innerHTML = "";
    if (!agents.length) {
      list.appendChild(el("div", "agents-empty", "Aucun agent. Lance le premier pour paralléliser le travail."));
    } else {
      for (const a of agents) list.appendChild(agentCard(a));
    }
  }

  async function refreshList() {
    await softRefresh();
  }

  function renderPanel(agents) {
    panel.innerHTML = "";
    const head = el("div", "agents-head");
    head.appendChild(el("h2", "", "Agents"));
    const newBtn = el("button", "btn-primary", "+ Nouvel agent");
    newBtn.type = "button";
    newBtn.addEventListener("click", openModal);
    head.appendChild(newBtn);
    const closeBtn = el("button", "mini-btn", "Fermer");
    closeBtn.type = "button";
    closeBtn.addEventListener("click", togglePanel);
    head.appendChild(closeBtn);
    panel.appendChild(head);

    if (selectedId) {
      const found = (agents || []).find((a) => a.id === selectedId);
      const body = el("div", "agent-detail");
      if (found) renderDetail(body, found);
      else {
        body.appendChild(el("div", "agent-empty", "Agent introuvable."));
        const back = el("button", "mini-btn", "← Retour");
        back.type = "button";
        back.addEventListener("click", closeDetail);
        body.appendChild(back);
      }
      panel.appendChild(body);
      return;
    }

    const list = el("div", "agents-list");
    if (!agents || !agents.length) {
      list.appendChild(el("div", "agents-empty", "Aucun agent. Lance le premier pour paralléliser le travail."));
    } else {
      for (const a of agents) list.appendChild(agentCard(a));
    }
    panel.appendChild(list);
  }

  function openModal() {
    loadAliases().then(() => {
      const overlay = el("div", "modal-overlay");
      const modal = el("div", "modal agent-modal");
      modal.appendChild(el("h3", "", "Nouvel agent"));
      const form = el("form");

      const promptLabel = el("label", "", "Mission");
      const prompt = el("textarea", "modal-input");
      prompt.rows = 4;
      prompt.required = true;
      prompt.placeholder = "Décris la tâche de cet agent…";
      promptLabel.appendChild(prompt);
      form.appendChild(promptLabel);

      const row = el("div", "modal-row");
      const famLabel = el("label", "", "Famille");
      const famSel = el("select", "modal-input");
      for (const f of families) {
        const opt = document.createElement("option");
        opt.value = f.id;
        opt.textContent = f.label;
        famSel.appendChild(opt);
      }
      famLabel.appendChild(famSel);
      row.appendChild(famLabel);

      const modeLabelEl = el("label", "", "Mode");
      const modeSel = el("select", "modal-input");
      modeLabelEl.appendChild(modeSel);
      row.appendChild(modeLabelEl);
      form.appendChild(row);

      function renderModes() {
        const f = families.find((x) => x.id === famSel.value);
        modeSel.innerHTML = "";
        for (const m of (f && f.modes ? f.modes : []).filter((m) => m.agent)) {
          const opt = document.createElement("option");
          opt.value = m.mode;
          opt.textContent = m.label + (m.rule ? " · " + m.rule : "");
          modeSel.appendChild(opt);
        }
      }
      famSel.addEventListener("change", renderModes);
      renderModes();

      const opts = el("div", "modal-opts");
      const approveLabel = el("label", "check-label");
      const approveChk = document.createElement("input");
      approveChk.type = "checkbox";
      approveChk.checked = true;
      approveLabel.appendChild(approveChk);
      approveLabel.appendChild(document.createTextNode(" Approbations avant écriture/exécution"));
      opts.appendChild(approveLabel);

      const planLabel = el("label", "check-label");
      const planChk = document.createElement("input");
      planChk.type = "checkbox";
      planLabel.appendChild(planChk);
      planLabel.appendChild(document.createTextNode(" Mode plan (valider avant de coder)"));
      opts.appendChild(planLabel);

      const wtLabel = el("label", "check-label");
      const wtChk = document.createElement("input");
      wtChk.type = "checkbox";
      wtChk.checked = true;
      wtLabel.appendChild(wtChk);
      wtLabel.appendChild(document.createTextNode(" Worktree git isolé"));
      opts.appendChild(wtLabel);
      form.appendChild(opts);

      const repoLabel = el("label", "", "Dépôt git (optionnel, chemin sous ton espace)");
      const repoInput = el("input", "modal-input");
      repoInput.type = "text";
      repoInput.placeholder = "~/workspace/mon-projet";
      repoLabel.appendChild(repoInput);
      form.appendChild(repoLabel);

      const errBox = el("div", "modal-error");
      errBox.hidden = true;
      form.appendChild(errBox);

      const actions = el("div", "modal-actions");
      const cancel = el("button", "mini-btn", "Annuler");
      cancel.type = "button";
      cancel.addEventListener("click", () => overlay.remove());
      const submit = el("button", "btn-primary", "Lancer l’agent");
      submit.type = "submit";
      actions.appendChild(cancel);
      actions.appendChild(submit);
      form.appendChild(actions);

      form.addEventListener("submit", async (e) => {
        e.preventDefault();
        errBox.hidden = true;
        submit.disabled = true;
        try {
          const res = await api("/api/agents", {
            method: "POST",
            body: {
              family: famSel.value,
              mode: modeSel.value,
              message: prompt.value.trim(),
              approve: approveChk.checked,
              plan: planChk.checked,
              worktree: wtChk.checked,
              repo: repoInput.value.trim(),
            },
          });
          overlay.remove();
          await softRefresh();
          openDetail(res.id);
        } catch (err) {
          errBox.textContent = err.message;
          errBox.hidden = false;
          submit.disabled = false;
        }
      });

      modal.appendChild(form);
      overlay.appendChild(modal);
      overlay.addEventListener("click", (e) => {
        if (e.target === overlay) overlay.remove();
      });
      document.body.appendChild(overlay);
      prompt.focus();
    });
  }

  function togglePanel() {
    panel.hidden = !panel.hidden;
    toolbarBtn.classList.toggle("active", !panel.hidden);
    if (!panel.hidden) {
      renderPanel(lastAgents);
      softRefresh();
      refreshTimer = setInterval(softRefresh, 3000);
    } else {
      clearInterval(refreshTimer);
      refreshTimer = null;
      if (detailView) {
        detailView.disconnect();
        detailView = null;
      }
      detailEls = null;
      selectedId = null;
    }
  }

  toolbarBtn.addEventListener("click", togglePanel);
}
