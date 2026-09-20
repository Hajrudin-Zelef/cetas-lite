// Panneau Paramètres → Crédit API.
// Affiche le crédit restant par provider, saisie manuelle, actualisation
// via les APIs OpenRouter / DeepSeek. Si le crédit est épuisé, le
// Sélecteur affiche un avertissement au choix du modèle.
import { api } from "./api.js";

function fmtCredit(c) {
  if (c === null || c === undefined) return "non renseigné";
  return Number(c).toFixed(2) + " $";
}

function creditRowClass(c) {
  if (c.exhausted) return "credit-row credit-exhausted";
  if (c.credit === null || c.credit === undefined) return "credit-row credit-unknown";
  return "credit-row credit-ok";
}

export async function loadCreditsPanel() {
  const body = document.getElementById("credits-body");
  if (!body) return;
  body.innerHTML = '<p class="apikey-intro">Chargement…</p>';
  let credits;
  try {
    credits = await api("/api/settings/credits");
  } catch (e) {
    body.innerHTML = '<p class="apikey-intro">Impossible de charger les crédits.</p>';
    return;
  }
  body.innerHTML = "";
  for (const c of credits) {
    const row = document.createElement("div");
    row.className = creditRowClass(c);
    row.dataset.provider = c.provider;

    const badge = c.exhausted
      ? '<span class="credit-badge credit-badge-bad">épuisé</span>'
      : (c.credit === null || c.credit === undefined)
        ? '<span class="credit-badge credit-badge-neutral">non renseigné</span>'
        : '<span class="credit-badge credit-badge-ok">ok</span>';
    const src = c.source ? ' <span class="credit-src">(' + c.source + (c.updated_at ? " · " + c.updated_at.slice(0, 10) : "") + ")</span>" : "";
    const manualNote = c.manual_only ? ' <span class="credit-src">saisie manuelle uniquement</span>' : "";

    row.innerHTML =
      '<div class="credit-label">' + c.label + " " + badge + src + manualNote + "</div>" +
      '<div class="credit-value">' + fmtCredit(c.credit) + "</div>" +
      '<div class="credit-edit">' +
        '<input type="number" min="0" step="0.01" placeholder="Crédit restant ($)" ' +
          (c.credit !== null && c.credit !== undefined ? 'value="' + c.credit + '"' : "") + " />" +
        '<button type="button" class="config-save-btn credit-save-btn">Enregistrer</button>' +
        '<button type="button" class="config-save-btn credit-clear-btn">Effacer</button>' +
      "</div>";

    const input = row.querySelector("input");
    const saveBtn = row.querySelector(".credit-save-btn");
    const clearBtn = row.querySelector(".credit-clear-btn");
    const state = document.getElementById("credits-save-state");

    saveBtn.addEventListener("click", async () => {
      const v = input.value.trim();
      if (v === "" || isNaN(Number(v)) || Number(v) < 0) {
        if (state) state.textContent = "Montant invalide.";
        return;
      }
      try {
        await api("/api/settings/credits", { method: "POST", body: { provider: c.provider, credit: Number(v) } });
        if (state) state.textContent = "Enregistré.";
        loadCreditsPanel();
      } catch (e) {
        if (state) state.textContent = "Échec de l'enregistrement.";
      }
    });
    clearBtn.addEventListener("click", async () => {
      try {
        await api("/api/settings/credits", { method: "POST", body: { provider: c.provider, credit: null } });
        if (state) state.textContent = "Effacé.";
        loadCreditsPanel();
      } catch (e) {
        if (state) state.textContent = "Échec.";
      }
    });
    body.appendChild(row);
  }

  const refreshBtn = document.getElementById("credits-refresh-btn");
  const refreshState = document.getElementById("credits-refresh-state");
  if (refreshBtn && !refreshBtn.dataset.bound) {
    refreshBtn.dataset.bound = "1";
    refreshBtn.addEventListener("click", async () => {
      refreshBtn.disabled = true;
      if (refreshState) refreshState.textContent = "Actualisation…";
      try {
        const res = await api("/api/settings/credits/refresh", { method: "POST" });
        const errs = res && res.errors ? res.errors : {};
        const keys = Object.keys(errs);
        if (refreshState) {
          refreshState.textContent = keys.length
            ? "Partiel : " + keys.map((k) => k + " (" + errs[k] + ")").join(", ")
            : "Actualisé.";
        }
        loadCreditsPanel();
      } catch (e) {
        if (refreshState) refreshState.textContent = "Échec de l'actualisation.";
      }
      refreshBtn.disabled = false;
    });
  }
}
