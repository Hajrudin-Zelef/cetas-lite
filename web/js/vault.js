// Coffre chiffré — panneau Configuration.
// Le mot de passe maître ne transite que vers /api/vault/* ; la clé dérivée
// reste côté serveur (session mémoire, verrouillage auto après 15 min).
import { getToken } from "./api.js";

// Appel direct (sans clearSession) : un 401 ici signifie « mauvais mot de
// passe du coffre », pas « session utilisateur expirée ».
export async function vaultFetch(path, opts = {}) {
  const headers = Object.assign({}, opts.headers || {});
  if (opts.body !== undefined) headers["Content-Type"] = "application/json";
  const token = getToken();
  if (token) headers["Authorization"] = "Bearer " + token;
  let resp;
  try {
    resp = await fetch(path, {
      method: opts.method || "GET",
      headers,
      body: opts.body !== undefined ? JSON.stringify(opts.body) : undefined,
    });
  } catch (e) {
    throw new Error("Connexion au serveur impossible.");
  }
  const text = await resp.text();
  let data = null;
  if (text) {
    try {
      data = JSON.parse(text);
    } catch (e) {
      data = text;
    }
  }
  if (!resp.ok) {
    const msg = data && data.error ? data.error : "Erreur " + resp.status;
    const err = new Error(msg);
    err.status = resp.status;
    throw err;
  }
  return data;
}

export function escapeHtml(s) {
  return String(s).replace(/[&<>"']/g, (c) => ({
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
    '"': "&quot;",
    "'": "&#39;",
  })[c]);
}

// Valeur JSON quelconque → texte affichable.
export function formatValue(v) {
  if (typeof v === "string") return v;
  try {
    return JSON.stringify(v);
  } catch (e) {
    return String(v);
  }
}

function fieldRow(label, inputHtml) {
  return (
    '<div class="audio-setting-row"><span class="audio-setting-label">' +
    escapeHtml(label) +
    "</span>" +
    inputHtml +
    "</div>"
  );
}

function pwInput(id, placeholder) {
  return (
    '<input type="password" id="' +
    id +
    '" class="vault-input" placeholder="' +
    escapeHtml(placeholder) +
    '" autocomplete="new-password">'
  );
}

function renderCreate(body) {
  body.innerHTML =
    '<div class="vault-box">' +
    "<p>Aucun coffre pour l'instant. Choisissez un mot de passe maître (12 caractères minimum).</p>" +
    fieldRow("Mot de passe", pwInput("vault-new-pw", "Mot de passe maître")) +
    fieldRow("Confirmation", pwInput("vault-new-pw2", "Confirmez le mot de passe")) +
    '<div class="vault-err" id="vault-err"></div>' +
    '<button type="button" class="models-save-btn" id="vault-create-btn">Créer le coffre</button>' +
    "</div>";
  body.querySelector("#vault-create-btn").addEventListener("click", async () => {
    const pw = body.querySelector("#vault-new-pw").value;
    const pw2 = body.querySelector("#vault-new-pw2").value;
    const err = body.querySelector("#vault-err");
    err.textContent = "";
    if (pw.length < 12) {
      err.textContent = "Minimum 12 caractères.";
      return;
    }
    if (pw !== pw2) {
      err.textContent = "Les deux mots de passe ne correspondent pas.";
      return;
    }
    try {
      await vaultFetch("/api/vault/init", { method: "POST", body: { password: pw } });
      await loadVaultPanel();
    } catch (e) {
      err.textContent = e.message;
    }
  });
}

function renderUnlock(body) {
  body.innerHTML =
    '<div class="vault-box">' +
    "<p>Le coffre est verrouillé. Entrez le mot de passe maître pour l'ouvrir.</p>" +
    fieldRow("Mot de passe", pwInput("vault-pw", "Mot de passe maître")) +
    '<div class="vault-err" id="vault-err"></div>' +
    '<button type="button" class="models-save-btn" id="vault-unlock-btn">Déverrouiller</button>' +
    "</div>";
  const doUnlock = async () => {
    const err = body.querySelector("#vault-err");
    err.textContent = "";
    const pw = body.querySelector("#vault-pw").value;
    try {
      await vaultFetch("/api/vault/unlock", { method: "POST", body: { password: pw } });
      await loadVaultPanel();
    } catch (e) {
      err.textContent = e.message;
    }
  };
  body.querySelector("#vault-unlock-btn").addEventListener("click", doUnlock);
  body.querySelector("#vault-pw").addEventListener("keydown", (e) => {
    if (e.key === "Enter") doUnlock();
  });
}

async function copyText(text, btn) {
  try {
    await navigator.clipboard.writeText(text);
  } catch (e) {
    const ta = document.createElement("textarea");
    ta.value = text;
    document.body.appendChild(ta);
    ta.select();
    try {
      document.execCommand("copy");
    } catch (_) {}
    ta.remove();
  }
  if (btn) {
    const old = btn.textContent;
    btn.textContent = "Copié ✓";
    setTimeout(() => {
      btn.textContent = old;
    }, 1200);
  }
}

function renderEntries(body, keys) {
  const rows = keys
    .map(
      (k) =>
        '<div class="vault-row" data-key="' +
        escapeHtml(k) +
        '">' +
        '<span class="vault-key">' +
        escapeHtml(k) +
        '</span>' +
        '<span class="vault-val" data-val>••••••</span>' +
        '<span class="vault-actions">' +
        '<button type="button" class="vault-mini-btn" data-act="show">Voir</button>' +
        '<button type="button" class="vault-mini-btn" data-act="copy">Copier</button>' +
        '<button type="button" class="vault-mini-btn vault-danger" data-act="del">Supprimer</button>' +
        "</span></div>"
    )
    .join("");
  body.innerHTML =
    '<div class="vault-box">' +
    '<div class="vault-head"><span>' +
    keys.length +
    " secret" +
    (keys.length > 1 ? "s" : "") +
    "</span>" +
    '<button type="button" class="vault-mini-btn" id="vault-lock-btn">Verrouiller</button></div>' +
    '<div id="vault-list">' +
    (rows || '<p class="vault-empty">Aucun secret. Ajoutez-en un ci-dessous.</p>') +
    "</div>" +
    '<div class="vault-add">' +
    fieldRow(
      "Clé",
      '<input type="text" id="vault-add-key" class="vault-input" placeholder="ex. github_token" maxlength="128">'
    ) +
    fieldRow(
      "Valeur",
      '<input type="text" id="vault-add-val" class="vault-input" placeholder="Valeur du secret">'
    ) +
    '<div class="vault-err" id="vault-err"></div>' +
    '<button type="button" class="models-save-btn" id="vault-add-btn">Ajouter</button>' +
    "</div>" +
    '<details class="vault-change"><summary>Changer le mot de passe maître</summary>' +
    fieldRow("Ancien", pwInput("vault-old-pw", "Ancien mot de passe")) +
    fieldRow("Nouveau", pwInput("vault-change-pw", "Nouveau mot de passe (12 min.)")) +
    '<div class="vault-err" id="vault-change-err"></div>' +
    '<button type="button" class="models-save-btn" id="vault-change-btn">Changer</button>' +
    "</details>" +
    "</div>";

  body.querySelector("#vault-lock-btn").addEventListener("click", async () => {
    await vaultFetch("/api/vault/lock", { method: "POST" }).catch(() => {});
    await loadVaultPanel();
  });

  body.querySelector("#vault-add-btn").addEventListener("click", async () => {
    const err = body.querySelector("#vault-err");
    err.textContent = "";
    const k = body.querySelector("#vault-add-key").value.trim();
    const v = body.querySelector("#vault-add-val").value;
    if (!k) {
      err.textContent = "Clé requise.";
      return;
    }
    try {
      await vaultFetch("/api/vault/entries/" + encodeURIComponent(k), {
        method: "PUT",
        body: { value: v },
      });
      await loadVaultPanel();
    } catch (e) {
      err.textContent = e.message;
    }
  });

  body.querySelector("#vault-change-btn").addEventListener("click", async () => {
    const err = body.querySelector("#vault-change-err");
    err.textContent = "";
    const oldPw = body.querySelector("#vault-old-pw").value;
    const newPw = body.querySelector("#vault-change-pw").value;
    if (newPw.length < 12) {
      err.textContent = "Minimum 12 caractères.";
      return;
    }
    try {
      await vaultFetch("/api/vault/change-password", {
        method: "POST",
        body: { old_password: oldPw, new_password: newPw },
      });
      err.textContent = "";
      body.querySelector("#vault-old-pw").value = "";
      body.querySelector("#vault-change-pw").value = "";
      const ok = document.createElement("div");
      ok.className = "vault-ok";
      ok.textContent = "Mot de passe modifié ✓";
      body.querySelector(".vault-change").appendChild(ok);
      setTimeout(() => ok.remove(), 3000);
    } catch (e) {
      err.textContent = e.message;
    }
  });

  body.querySelector("#vault-list").addEventListener("click", async (e) => {
    const btn = e.target.closest("[data-act]");
    if (!btn) return;
    const row = btn.closest(".vault-row");
    const k = row.dataset.key;
    const valEl = row.querySelector("[data-val]");
    const act = btn.dataset.act;
    if (act === "show") {
      if (btn.dataset.shown) {
        valEl.textContent = "••••••";
        btn.textContent = "Voir";
        delete btn.dataset.shown;
      } else {
        try {
          const data = await vaultFetch("/api/vault/entries/" + encodeURIComponent(k));
          valEl.textContent = formatValue(data.value);
          btn.textContent = "Masquer";
          btn.dataset.shown = "1";
        } catch (err) {
          valEl.textContent = "erreur";
        }
      }
    } else if (act === "copy") {
      try {
        const data = await vaultFetch("/api/vault/entries/" + encodeURIComponent(k));
        await copyText(formatValue(data.value), btn);
      } catch (err) {}
    } else if (act === "del") {
      if (!window.confirm("Supprimer le secret « " + k + " » ?")) return;
      try {
        await vaultFetch("/api/vault/entries/" + encodeURIComponent(k), { method: "DELETE" });
        await loadVaultPanel();
      } catch (err) {}
    }
  });
}

export async function loadVaultPanel() {
  const body = document.getElementById("vault-body");
  if (!body) return;
  body.innerHTML = '<p class="vault-empty">Chargement…</p>';
  let st;
  try {
    st = await vaultFetch("/api/vault/status");
  } catch (e) {
    body.innerHTML = '<div class="vault-err">Impossible de joindre le coffre : ' + escapeHtml(e.message) + "</div>";
    return;
  }
  if (!st.exists) renderCreate(body);
  else if (!st.unlocked) renderUnlock(body);
  else {
    try {
      const data = await vaultFetch("/api/vault/entries");
      renderEntries(body, data.keys || []);
    } catch (e) {
      body.innerHTML = '<div class="vault-err">' + escapeHtml(e.message) + "</div>";
    }
  }
}
