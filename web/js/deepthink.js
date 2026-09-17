// Module "DeepThink Global" : reglages de la traduction du raisonnement
// (langue cible + modele traducteur). Onglet des parametres, persistance
// via PUT /api/deepthink (sauvegarde immediate a chaque changement).

import { api } from "./api.js";

function el(tag, cls, text) {
  const e = document.createElement(tag);
  if (cls) e.className = cls;
  if (text !== undefined) e.textContent = text;
  return e;
}

const LANG_NAMES = { es: "Espagnol", fr: "Français", it: "Italien", de: "Allemand", en: "Anglais" };

let settings = null;
let saveTimer = 0;

function setSaveState(text) {
  const s = document.getElementById("deepthink-save-state");
  if (s) s.textContent = text || "";
}

function persist() {
  setSaveState("Enregistrement…");
  clearTimeout(saveTimer);
  saveTimer = setTimeout(async () => {
    try {
      const data = await api("/api/deepthink", { method: "PUT", body: settings });
      if (data && data.settings) settings = data.settings;
      setSaveState("Enregistré ✓");
    } catch (e) {
      setSaveState("Erreur : " + ((e && e.message) || e));
    }
  }, 350);
}

function segButtons(values, isActive, onPick, labelOf) {
  const seg = el("div", "ms-seg");
  seg.style.marginLeft = "0";
  for (const v of values) {
    const b = el("button", "ms-seg-btn" + (isActive(v) ? " active" : ""), labelOf(v));
    b.type = "button";
    b.addEventListener("click", () => {
      if (isActive(v)) return;
      onPick(v);
      seg.querySelectorAll(".ms-seg-btn").forEach((x) => x.classList.remove("active"));
      b.classList.add("active");
      persist();
    });
    seg.appendChild(b);
  }
  return seg;
}

function render(bodyEl, data) {
  bodyEl.innerHTML = "";

  const langTitle = el("div", "dt-row-title", "Langue de traduction");
  bodyEl.appendChild(langTitle);
  bodyEl.appendChild(
    segButtons(
      data.langs && data.langs.length ? data.langs : ["es", "fr", "it", "de", "en"],
      (l) => settings.lang === l,
      (l) => {
        settings.lang = l;
      },
      (l) => LANG_NAMES[l] || l
    )
  );

  const modelTitle = el("div", "dt-row-title", "Modèle traducteur");
  bodyEl.appendChild(modelTitle);
  const models = data.models && data.models.length ? data.models : [];
  bodyEl.appendChild(
    segButtons(
      models,
      (m) => settings.provider === m.provider && settings.model === m.model,
      (m) => {
        settings.provider = m.provider;
        settings.model = m.model;
      },
      (m) => m.label
    )
  );

  const note = el(
    "p",
    "apikey-intro",
    "Le bouton 🌐 de l'en-tête REASONING (panneau latéral) traduit le raisonnement affiché et l'affiche sous l'original — sans jamais le remplacer. Le premier clic active aussi le mode auto : chaque raisonnement suivant est traduit automatiquement 5 secondes après sa fin. Un second clic désactive l'auto."
  );
  note.style.marginTop = "14px";
  bodyEl.appendChild(note);
}

export async function loadDeepThinkPanel() {
  const bodyEl = document.getElementById("deepthink-body");
  if (!bodyEl || bodyEl.dataset.loaded) return;
  bodyEl.dataset.loaded = "1";
  bodyEl.textContent = "Chargement…";
  try {
    const data = await api("/api/deepthink");
    settings = (data && data.settings) || { lang: "fr", provider: "deepseek", model: "deepseek-chat" };
    render(bodyEl, data || {});
  } catch (e) {
    bodyEl.textContent = "Erreur de chargement : " + ((e && e.message) || e);
    delete bodyEl.dataset.loaded;
  }
}
