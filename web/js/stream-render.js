export function takeCount(len, dtMs, cps = 30) {
  if (len <= 0) return 0;
  const catchup = Math.ceil(len / 8);
  const base = Math.max(1, Math.round((cps * dtMs) / 1000));
  return Math.min(len, Math.max(base, Math.min(catchup, len)));
}

function now() {
  return typeof performance !== "undefined" && performance.now ? performance.now() : Date.now();
}

function reducedMotion() {
  return (
    typeof window !== "undefined" &&
    typeof window.matchMedia === "function" &&
    window.matchMedia("(prefers-reduced-motion: reduce)").matches
  );
}

function raf(fn) {
  if (typeof requestAnimationFrame === "function") return requestAnimationFrame(fn);
  return setTimeout(() => fn(now()), 16);
}

function cancel(id) {
  if (typeof cancelAnimationFrame === "function") cancelAnimationFrame(id);
  else clearTimeout(id);
}

export function createStreamRenderer({ onRender, onFirst, onDone, cps = 30, reducedMotion: reducedOverride } = {}) {
  const reduced = reducedOverride === undefined ? reducedMotion() : reducedOverride;
  let buffer = "";
  let shown = 0;
  let frame = 0;
  let last = 0;
  let first = true;

  function render() {
    if (onRender) onRender(buffer.slice(0, shown));
  }

  function tick(t) {
    frame = 0;
    const dt = Math.max(0, t - last);
    last = t;
    const take = takeCount(buffer.length - shown, dt, cps);
    if (take > 0) shown += take;
    render();
    if (shown < buffer.length) start();
    else if (onDone) onDone();
  }

  function start() {
    if (frame) return;
    last = now();
    frame = raf(tick);
  }

  function add(text) {
    if (!text) return;
    if (first) {
      first = false;
      if (onFirst) onFirst();
    }
    buffer += String(text);
    if (reduced) {
      shown = buffer.length;
      render();
      return;
    }
    start();
  }

  function replace(text) {
    buffer = String(text || "");
    shown = buffer.length;
    if (frame) {
      cancel(frame);
      frame = 0;
    }
    render();
  }

  function flush() {
    if (frame) {
      cancel(frame);
      frame = 0;
    }
    shown = buffer.length;
    render();
  }

  function reset() {
    buffer = "";
    shown = 0;
    first = true;
    if (frame) {
      cancel(frame);
      frame = 0;
    }
  }

  return {
    add,
    replace,
    flush,
    reset,
    text: () => buffer,
    isActive: () => frame !== 0,
  };
}
