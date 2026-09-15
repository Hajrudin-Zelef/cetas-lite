#!/usr/bin/env python3
"""Génère web/css/features/marex-agents.css à partir du CSS Marexcode de référence.

- Renomme les classes structurelles (collisions avec le design system Cetas) :
  .sidebar -> .mx-sidebar, .main -> .mx-main, .app -> .mx-app, .frame -> .mx-frame,
  .hero -> .mx-hero, .composer -> .mx-composer, .chat-panel -> .mx-chat-panel,
  .chat-log -> .mx-chat-log, .side-panel -> .mx-side-panel, .mobile-bar -> .mx-mobile-bar
- Scope chaque sélecteur sous #marex-view (variables CSS isolées, aucun leak).
- :root -> #marex-view, body[data-accent] -> #marex-view[data-accent],
  body.light -> #marex-view.light, html,body -> #marex-view.
- @media : récursion interne. @keyframes : conservés tels quels.
"""
import re
import sys

RENAMES = [
    (r"\.side-panel(?![\w-])", ".mx-side-panel"),
    (r"\.mobile-bar(?![\w-])", ".mx-mobile-bar"),
    (r"\.chat-panel(?![\w-])", ".mx-chat-panel"),
    (r"\.chat-log(?![\w-])", ".mx-chat-log"),
    (r"\.sidebar(?![\w-])", ".mx-sidebar"),
    (r"\.composer(?![\w-])", ".mx-composer"),
    (r"\.frame(?![\w-])", ".mx-frame"),
    (r"\.hero(?![\w-])", ".mx-hero"),
    (r"\.main(?![\w-])", ".mx-main"),
    (r"\.app(?![\w-])", ".mx-app"),
]

def rename_classes(css):
    for pat, repl in RENAMES:
        css = re.sub(pat, repl, css)
    return css

def scope_selector(sel):
    sel = sel.strip()
    if not sel:
        return sel
    if sel == ":root":
        return "#marex-view"
    if sel in ("html", "body", "html,body", "html, body"):
        return "#marex-view"
    if sel.startswith("body[data-accent"):
        return "#marex-view" + sel[4:]
    if sel.startswith("body.light"):
        return "#marex-view.light" + sel[len("body.light"):]
    if sel == "*":
        return "#marex-view *"
    # ne pas re-scoper un sélecteur déjà scopé
    if sel.startswith("#marex-view"):
        return sel
    return "#marex-view " + sel

def scope_rule_block(css):
    """Scope les règles de premier niveau d'un bloc CSS (sans @media imbriqués)."""
    out = []
    i, n = 0, len(css)
    while i < n:
        # sauter espaces/commentaires
        if css.startswith("/*", i):
            j = css.find("*/", i + 2)
            j = n if j < 0 else j + 2
            out.append(css[i:j])
            i = j
            continue
        if css[i].isspace():
            out.append(css[i])
            i += 1
            continue
        # @-rule
        if css[i] == "@":
            m = re.match(r"@([a-zA-Z-]+)", css[i:])
            name = m.group(1) if m else ""
            # trouver la fin de l'en-tête (jusqu'à { ou ;)
            j = i
            depth = 0
            while j < n:
                if css[j] == "{":
                    break
                if css[j] == ";":
                    break
                j += 1
            header = css[i:j]
            if j < n and css[j] == "{":
                # trouver la fin du bloc
                k = j + 1
                d = 1
                while k < n and d > 0:
                    if css.startswith("/*", k):
                        kk = css.find("*/", k + 2)
                        k = n if kk < 0 else kk + 2
                        continue
                    if css[k] == "{":
                        d += 1
                    elif css[k] == "}":
                        d -= 1
                    k += 1
                inner = css[j + 1:k - 1]
                if name == "media":
                    out.append(header + "{" + scope_rule_block(inner) + "}")
                else:
                    # keyframes, font-face, etc. : inchangés
                    out.append(css[i:k])
                i = k
            else:
                out.append(header + ";")
                i = j + 1
            continue
        # règle normale : sélecteurs { ... }
        j = css.find("{", i)
        if j < 0:
            out.append(css[i:])
            break
        selectors = css[i:j]
        k = j + 1
        d = 1
        while k < n and d > 0:
            if css.startswith("/*", k):
                kk = css.find("*/", k + 2)
                k = n if kk < 0 else kk + 2
                continue
            if css[k] == "{":
                d += 1
            elif css[k] == "}":
                d -= 1
            k += 1
        body = css[j:k]
        parts = [scope_selector(s) for s in selectors.split(",")]
        # règle html,body mappée : si un sélecteur devient vide on garde les autres
        parts = [p for p in parts if p]
        if parts:
            out.append(",".join(parts) + body)
        i = k
    return "".join(out)

def main():
    base = "/tmp/marex_full/marexcode/css"
    css = ""
    for name in ("marexcode.css", "ocean-fx.css"):
        with open(f"{base}/{name}", encoding="utf-8") as f:
            css += f.read() + "\n"
    css = rename_classes(css)
    scoped = scope_rule_block(css)

    header = """/* ============================================================
   Vue Agents façon Marexcode — généré depuis marexcode.tar (référence).
   Tout est scopé sous #marex-view : aucun impact sur le reste de l'app.
   Classes structurelles renommées en mx-* (collisions évitées).
   NE PAS ÉDITER À LA MAIN — régénérer via scripts/gen-marex-css.py
   ============================================================ */
#marex-view{
  position:fixed;
  inset:0;
  z-index:400;
  display:none;
  background:#000;
  font-family:"Inter",-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,sans-serif;
  -webkit-font-smoothing:antialiased;
  overflow:hidden;
}
#marex-view.open{ display:block; }
#marex-view .mx-app{ font-family:inherit; }
"""

    # Adaptations ThreadView -> look Marexcode (.msg)
    overrides = """
/* ---------- Adaptations ThreadView -> bulles Marexcode ---------- */
#marex-view .mx-chat-log .message-wrapper{ margin-bottom:14px; }
#marex-view .mx-chat-log .message-wrapper-user{ display:flex; justify-content:flex-end; }
#marex-view .mx-chat-log .message{
  max-width:75%;
  padding:10px 14px;
  border-radius:12px;
  font-size:13.5px;
  line-height:1.5;
  border:1px solid var(--border);
  background:transparent;
  color:var(--text-primary);
  box-shadow:none;
  white-space:pre-wrap;
  word-wrap:break-word;
}
#marex-view .mx-chat-log .message-text{ white-space:pre-wrap; }
#marex-view .mx-chat-log .message-error{ background:#3a2222; color:#ffb4b4; border-color:#5a3030; }
#marex-view .mx-chat-log .msg-system{
  text-align:center; font-size:12px; color:var(--text-tertiary);
  margin:10px 0;
}
#marex-view .mx-chat-log .thinking-block{
  border:1px solid var(--border); border-radius:8px;
  margin-bottom:8px; font-size:12.5px; color:var(--text-secondary);
}
#marex-view .mx-chat-log .thinking-block summary{ cursor:pointer; padding:6px 10px; color:var(--accent); }
#marex-view .mx-chat-log .thinking-content{ padding:0 10px 8px; white-space:pre-wrap; }
#marex-view .mx-chat-log .message .md{ white-space:normal; }
/* Étapes d'outils façon Marexcode */
#marex-view .mx-chat-log .tool-block{
  border:1px solid var(--border); border-radius:8px;
  margin:8px 0; overflow:hidden; font-size:12.5px;
}
#marex-view .mx-chat-log .chat-steps{ padding:8px 12px; }
/* Badge statut agent dans le hero */
#marex-view .mx-agent-status{
  display:inline-flex; align-items:center; gap:6px;
  font-size:11.5px; color:var(--text-secondary);
  border:1px solid var(--border); border-radius:999px; padding:3px 10px;
}
#marex-view .mx-agent-status .dot{ width:7px; height:7px; border-radius:50%; background:var(--accent); }
#marex-view .mx-agent-status.running .dot{ animation:sb-metric-pulse 1.2s ease-in-out infinite; }
#marex-view .mx-agent-status.done .dot{ background:#22c55e; animation:none; }
#marex-view .mx-agent-status.stopped .dot, #marex-view .mx-agent-status.error .dot{ background:#ef4444; animation:none; }
/* Étoile favori */
#marex-view .mx-fav-btn.active{ color:#f5c518; }
#marex-view .mx-fav-btn.active svg{ fill:#f5c518; }
/* Bouton flottant réouverture panneau raisonnement */
#marex-view .mx-think-fab{
  position:absolute; right:12px; bottom:12px; z-index:15;
  display:none; align-items:center; gap:6px;
  background:var(--bg-input-btn); border:1px solid var(--border);
  color:var(--accent); font-size:12px; font-family:inherit;
  border-radius:999px; padding:7px 12px; cursor:pointer;
}
#marex-view .mx-think-fab.show{ display:inline-flex; }
/* Menu modèle : groupes familles */
#marex-view .cdrop-menu .cdrop-section-label{ position:sticky; top:0; }
#marex-view .mx-model-group{ padding:2px 0; }
/* Recherche discussions */
#marex-view .mx-search-wrap{ padding:2px 4px 6px; }
#marex-view .mx-search-input{
  width:100%; background:var(--bg-input-btn); border:1px solid var(--border);
  border-radius:8px; color:var(--text-primary); font-size:12.5px; font-family:inherit;
  padding:7px 10px; outline:none;
}
#marex-view .mx-search-input::placeholder{ color:var(--text-tertiary); }
#marex-view .mx-search-input:focus{ border-color:var(--accent); }
/* Le ThreadView pilote le bouton stop via l'attribut hidden */
#marex-view .stop-btn:not([hidden]){ display:flex; }
/* Item discussion avec actions */
#marex-view .sb-hist-item{ position:relative; }
#marex-view .mx-hist-del{
  flex-shrink:0; display:none; color:var(--text-tertiary);
  background:none; border:none; cursor:pointer; padding:2px; border-radius:4px;
  font-size:14px; line-height:1;
}
#marex-view .sb-hist-item:hover .mx-hist-del{ display:block; }
#marex-view .mx-hist-del:hover{ color:var(--danger); background:var(--bg-active); }
/* Scrollbar du panneau latéral */
#marex-view .side-panel-body{ scrollbar-width:thin; }
"""

    out = header + scoped + overrides
    dest = "/home/hatch/workspace/cetas-lite/web/css/features/marex-agents.css"
    with open(dest, "w", encoding="utf-8") as f:
        f.write(out)
    print("écrit", dest, len(out), "octets")

if __name__ == "__main__":
    main()
