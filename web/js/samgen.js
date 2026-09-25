/* Panneau "IA locale (SamGen)" : logique separee des providers cloud.
   Trois moteurs (llama.cpp, Ollama, LM Studio) : URL + cle API facultative
   (chiffree au coffre, jamais renvoyee), bouton de test. Namespace :
   /api/local/engines */

import { api } from "./api.js";

function escHtml(s) {
  return String(s == null ? "" : s)
    .replace(/&/g, "&amp;").replace(/</g, "&lt;")
    .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
}

let _loaded = false;

export async function initSamGenPanel() {
  const body = document.getElementById("samgen-body");
  if (!body) return;
  if (_loaded) { await _refresh(); return; }
  _loaded = true;
  body.innerHTML = "<p>Chargement…</p>";
  try {
    const d = await api("/api/local/engines");
    const engines = (d && d.engines) || [];
    if (!engines.length) {
      body.innerHTML = "<p>Aucun moteur local disponible.</p>";
      return;
    }
    body.innerHTML = engines.map(_cardHtml).join("");
    engines.forEach(_bindCard);
  } catch (e) {
    body.innerHTML = "<p>Chargement impossible : " + escHtml(e.message || e) + "</p>";
    _loaded = false;
  }
}

async function _refresh() {
  _loaded = false;
  await initSamGenPanel();
}

function _cardHtml(e) {
  const id = escHtml(e.id);
  const keyRow = e.has_key
    ? '<div class="apikey-masked-row" id="samgen-masked-' + id + '">' +
      '<span class="apikey-masked-key">••••••••</span>' +
      '<button type="button" class="apikey-delete-link" data-engine="' + id + '" title="Supprimer la clé">Supprimer la clé</button>' +
      "</div>"
    : "";
  return '<div class="provider-section samgen-card" data-engine="' + id + '">' +
    '<div class="apikey-label-row"><label class="sp-modal-label">' + escHtml(e.label) + "</label></div>" +
    '<div class="apikey-field"><div class="apikey-input-wrap">' +
    '<input type="text" id="samgen-url-' + id + '" class="sp-modal-input apikey-input" style="padding-right:12px" placeholder="http://…" value="' + escHtml(e.url || "") + '" autocomplete="off" spellcheck="false">' +
    "</div></div>" +
    keyRow +
    '<div class="apikey-field"><div class="apikey-input-wrap">' +
    '<input type="password" id="samgen-key-' + id + '" class="sp-modal-input apikey-input" placeholder="' + (e.has_key ? "Nouvelle clé (laisser vide pour conserver)" : "Clé API (facultatif)…") + '" autocomplete="off" spellcheck="false">' +
    "</div></div>" +
    '<div class="samgen-actions">' +
    '<button type="button" class="apikey-validate-btn" id="samgen-save-' + id + '">Enregistrer</button>' +
    '<button type="button" class="apikey-validate-btn" id="samgen-test-' + id + '">Tester</button>' +
    "</div>" +
    '<p class="apikey-local-status" id="samgen-status-' + id + '"></p>' +
    "</div>";
}

function _status(id, kind, msg) {
  const el = document.getElementById("samgen-status-" + id);
  if (!el) return;
  el.className = "apikey-local-status " + (kind || "");
  el.textContent = msg || "";
}

function _bindCard(e) {
  const id = e.id;
  const card = document.querySelector('.samgen-card[data-engine="' + id + '"]');
  if (!card) return;
  const urlInput = card.querySelector("#samgen-url-" + id);
  const keyInput = card.querySelector("#samgen-key-" + id);
  const saveBtn = card.querySelector("#samgen-save-" + id);
  const testBtn = card.querySelector("#samgen-test-" + id);
  const delBtn = card.querySelector(".apikey-delete-link");

  if (saveBtn) saveBtn.addEventListener("click", async () => {
    const url = (urlInput.value || "").trim();
    const key = (keyInput.value || "").trim();
    if (!url && !key) { _status(id, "error", "Saisissez une URL ou une clé."); return; }
    saveBtn.disabled = true;
    _status(id, "", "Enregistrement…");
    try {
      const r = await api("/api/local/engines/" + encodeURIComponent(id), {
        method: "PUT", body: { url, key },
      });
      keyInput.value = "";
      _status(id, "success", "Enregistré." + (r && r.has_key ? " Clé configurée." : ""));
      if (key) await _refresh(); // affiche la ligne "••••••••" + Supprimer
    } catch (err) {
      _status(id, "error", "Échec : " + (err.message || err));
    } finally { saveBtn.disabled = false; }
  });

  if (testBtn) testBtn.addEventListener("click", async () => {
    testBtn.disabled = true;
    _status(id, "", "Test en cours…");
    try {
      const r = await api("/api/local/engines/" + encodeURIComponent(id) + "/test", { method: "POST" });
      if (r && r.ok) {
        _status(id, "success", "OK — " + (r.models || 0) + " modèle(s), " + (r.latency_ms || 0) + " ms.");
      } else {
        _status(id, "error", "Échec : " + ((r && r.error) || "inconnu"));
      }
    } catch (err) {
      _status(id, "error", "Échec : " + (err.message || err));
    } finally { testBtn.disabled = false; }
  });

  if (delBtn) delBtn.addEventListener("click", async () => {
    if (!confirm("Supprimer la clé API de " + e.label + " ?")) return;
    try {
      await api("/api/local/engines/" + encodeURIComponent(id), { method: "DELETE" });
      await _refresh();
    } catch (err) {
      _status(id, "error", "Échec : " + (err.message || err));
    }
  });
}
