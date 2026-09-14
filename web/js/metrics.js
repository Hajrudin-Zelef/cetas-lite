// Section "Métriques" de la sidebar : CPU / RAM / Disque / Réseau.
// Interroge GET /api/metrics toutes les 10 s.
import { api } from "./api.js";
import { currentSelection, getFamilies } from "./model-select.js";
import { getLastTokens } from "./turn-tokens.js";

function updateModelRow() {
  const el = document.getElementById("sb-m-model-val");
  if (!el) return;
  let label = "—";
  try {
    const sel = currentSelection();
    const f = getFamilies().find((x) => x.id === sel.family);
    const m = f && f.modes.find((x) => x.mode === sel.mode);
    const member = m && m.pool && m.pool[0];
    label = (member && (member.label || member.model)) || (f ? f.label : "—");
  } catch (e) {
    /* familles pas encore chargées */
  }
  const t = getLastTokens();
  const total = (t.in || 0) + (t.out || 0);
  el.textContent = label + " · " + total.toLocaleString("fr") + " tok";
}

function fmtBytes(n) {
  n = Number(n) || 0;
  if (n >= 1e12) return (n / 1e12).toFixed(1).replace(".", ",") + " To";
  if (n >= 1e9) return (n / 1e9).toFixed(1).replace(".", ",") + " Go";
  if (n >= 1e6) return (n / 1e6).toFixed(1).replace(".", ",") + " Mo";
  if (n >= 1e3) return (n / 1e3).toFixed(0) + " Ko";
  return n.toFixed(0) + " o";
}

function fmtRate(bps) {
  if (bps >= 1e6) return (bps / 1e6).toFixed(1).replace(".", ",") + " Mo/s";
  if (bps >= 1e3) return (bps / 1e3).toFixed(0) + " Ko/s";
  return bps.toFixed(0) + " o/s";
}

function setBar(fillId, valId, pct, label) {
  const fill = document.getElementById(fillId);
  const val = document.getElementById(valId);
  if (!fill || !val) return;
  const p = Math.min(100, Math.max(0, Number(pct) || 0));
  fill.style.width = p.toFixed(0) + "%";
  fill.style.background =
    p >= 90 ? "var(--danger,#ef4444)" : p >= 70 ? "var(--warning,#eab308)" : "var(--success,#22c55e)";
  val.textContent = label;
}

let prevNet = null;
let prevTs = 0;
let timer = null;

export function initMetrics() {
  if (!document.getElementById("sb-metrics") || timer) return;

  async function tick() {
    updateModelRow();
    let m;
    try {
      m = await api("/api/metrics");
    } catch (e) {
      return; // silencieux : le prochain cycle retentera
    }
    const now = Date.now();
    if (typeof m.cpu === "number") {
      setBar("sb-m-cpu-fill", "sb-m-cpu-val", m.cpu, m.cpu.toFixed(0) + " %");
    }
    if (m.ram) {
      const pct = m.ram.total ? (m.ram.used * 100) / m.ram.total : 0;
      setBar("sb-m-ram-fill", "sb-m-ram-val", pct, fmtBytes(m.ram.used) + " / " + fmtBytes(m.ram.total));
    }
    if (m.disk) {
      const pct = m.disk.total ? (m.disk.used * 100) / m.disk.total : 0;
      setBar("sb-m-disk-fill", "sb-m-disk-val", pct, fmtBytes(m.disk.used) + " / " + fmtBytes(m.disk.total));
    }
    if (m.net) {
      const nv = document.getElementById("sb-m-net-val");
      if (nv) {
        if (prevNet && now > prevTs) {
          const dt = (now - prevTs) / 1000;
          const rx = Math.max(0, (m.net.rx_bytes - prevNet.rx_bytes) / dt);
          const tx = Math.max(0, (m.net.tx_bytes - prevNet.tx_bytes) / dt);
          nv.textContent = "↓ " + fmtRate(rx) + " · ↑ " + fmtRate(tx);
        } else {
          nv.textContent = "↓ " + fmtBytes(m.net.rx_bytes) + " · ↑ " + fmtBytes(m.net.tx_bytes);
        }
      }
      prevNet = { rx_bytes: m.net.rx_bytes || 0, tx_bytes: m.net.tx_bytes || 0 };
      prevTs = now;
    }
  }

  tick();
  timer = setInterval(tick, 10000);
  document.addEventListener("visibilitychange", () => {
    if (document.hidden) {
      clearInterval(timer);
      timer = null;
      prevNet = null;
      prevTs = 0;
    } else if (!timer) {
      tick();
      timer = setInterval(tick, 10000);
    }
  });
}
