// --- Dialogues personnalisés (alert/confirm) ---
// Utilise la modale #custom-dialog-overlay de index.html, avec repli sur
// window.confirm si elle est absente (ex. tests jsdom).
export function confirmDialog(message, opts = {}) {
  return new Promise((resolve) => {
    const overlay = document.getElementById("custom-dialog-overlay");
    const icon = document.getElementById("custom-dialog-icon");
    const msg = document.getElementById("custom-dialog-message");
    const okBtn = document.getElementById("custom-dialog-ok");
    const cancelBtn = document.getElementById("custom-dialog-cancel");
    if (!overlay) {
      resolve(window.confirm(message));
      return;
    }
    msg.textContent = message;
    icon.textContent = opts.danger ? "⚠️" : "❓";
    okBtn.textContent = opts.okLabel || "OK";
    cancelBtn.style.display = opts.hideCancel ? "none" : "";
    const done = (v) => {
      overlay.style.display = "none";
      okBtn.removeEventListener("click", onOk);
      cancelBtn.removeEventListener("click", onCancel);
      overlay.removeEventListener("click", onBg);
      resolve(v);
    };
    const onOk = () => done(true);
    const onCancel = () => done(false);
    const onBg = (e) => {
      if (e.target === overlay) done(false);
    };
    okBtn.addEventListener("click", onOk);
    cancelBtn.addEventListener("click", onCancel);
    overlay.addEventListener("click", onBg);
    overlay.style.display = "";
  });
}

export function alertDialog(message) {
  return confirmDialog(message, { okLabel: "Compris", hideCancel: true });
}
