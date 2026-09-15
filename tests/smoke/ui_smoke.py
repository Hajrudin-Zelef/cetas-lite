"""Smoke test UI (version allégée, 2026-09-15).

Lance le vrai binaire `cetas-lite serve`, pilote un vrai Chromium via
Playwright avec des routes API simulées (aucune clé requise), et vérifie
6 points que les tests unitaires (jsdom) ne peuvent pas voir :

  1. la page charge sans erreur JavaScript ;
  2. un tour de chat simulé (SSE) s'affiche avec le markdown rendu ;
  3. l'envoi (Entrée) émet bien POST /api/chat/send avec le message ;
  4. la vue Agents s'ouvre en plein écran via le bouton module ;
  5. les menus du composer (+, modèle, projet) s'ouvrent ET sont
     réellement visibles dans le viewport (non rognés) ;
  6. aucune erreur JS durant tout le scénario.

Usage : make smoke   (ou : python3 tests/smoke/ui_smoke.py)
"""

import json
import os
import shutil
import socket
import subprocess
import tempfile
import time
import urllib.request

from playwright.sync_api import sync_playwright

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
BIN = os.path.join(ROOT, "bin", "cetas-lite")
if os.name == "nt" and not os.path.exists(BIN) and os.path.exists(BIN + ".exe"):
    BIN += ".exe"  # go build ajoute .exe automatiquement sous Windows

FAMILIES = {
    "families": [
        {
            "id": "code",
            "label": "Code",
            "modes": [
                {
                    "family": "code",
                    "mode": "standard",
                    "label": "Standard",
                    "agent": True,
                    "local": False,
                }
            ],
        }
    ]
}

SETTINGS = {"theme": "ocean", "family": "code", "mode": "standard"}

STATE = {"send_body": None}


def sse(events):
    return "".join("data: " + json.dumps(e) + "\n\n" for e in events)


# Un tour complet rejoué par le flux SSE simulé au chargement.
STREAM = sse(
    [
        {"seq": 1, "user": "salut"},
        {"seq": 2, "content": "## Titre\n\n**gras** et texte\n\n```js\nconsole.log(1)\n```\n"},
        {"seq": 3, "stats": {"prompt_tokens": 10, "completion_tokens": 5}},
        {"seq": 4, "turn_done": True, "elapsed_ms": 12},
    ]
)


def free_port():
    s = socket.socket()
    s.bind(("127.0.0.1", 0))
    port = s.getsockname()[1]
    s.close()
    return port


def wait_health(base, timeout=15):
    deadline = time.time() + timeout
    while time.time() < deadline:
        try:
            with urllib.request.urlopen(base + "/api/health", timeout=1) as r:
                if r.status == 200:
                    return True
        except Exception:
            time.sleep(0.2)
    return False


def route_mocks(page):
    # Catch-all : tout /api non redéfini après renvoie JSON vide (pas 401).
    page.route("**/api/**", lambda r: r.fulfill(json={}))
    page.route("**/api/config", lambda r: r.fulfill(json={"registration_open": False, "version": "smoke"}))
    page.route("**/api/me", lambda r: r.fulfill(json={"username": "test", "role": "user"}))
    page.route("**/api/aliases", lambda r: r.fulfill(json=FAMILIES))

    def settings(route):
        if route.request.method == "PUT":
            route.fulfill(json={"ok": True})
        else:
            route.fulfill(json=SETTINGS)

    def send(route):
        STATE["send_body"] = route.request.post_data_json
        route.fulfill(json={"ok": True})

    page.route("**/api/settings", settings)
    page.route("**/api/chat/send", send)
    page.route(
        "**/api/chat/stream**",
        lambda r: r.fulfill(
            status=200,
            headers={"Content-Type": "text/event-stream", "Cache-Control": "no-cache"},
            body=STREAM,
        ),
    )
    page.route("**/api/conversations**", lambda r: r.fulfill(json={"archives": []}))
    page.route("**/api/metrics", lambda r: r.fulfill(json={}))
    page.route("**/api/chat/state", lambda r: r.fulfill(json={"turns": 0, "generating": False}))
    page.route("**/api/model-info", lambda r: r.fulfill(json={}))
    page.route("**/api/plugins", lambda r: r.fulfill(json={"plugins": []}))
    page.route("**/api/mcp", lambda r: r.fulfill(json={"servers": []}))
    page.route("**/api/capabilities", lambda r: r.fulfill(json={"caps": {}}))


def assert_visible_in_viewport(page, selector, label):
    """Le menu doit être ouvert ET entièrement dans le viewport (anti-régression du rognage)."""
    box = page.locator(selector).bounding_box()
    assert box, f"{label} : pas de boîte de rendu"
    vp = page.viewport_size
    assert box["width"] > 40 and box["height"] > 20, f"{label} : boîte dégénérée {box}"
    assert box["x"] >= 0 and box["y"] >= 0, f"{label} : hors écran (x={box['x']}, y={box['y']})"
    assert box["x"] + box["width"] <= vp["width"] + 1, f"{label} : dépasse à droite"
    assert box["y"] + box["height"] <= vp["height"] + 1, f"{label} : dépasse en bas"
    assert page.locator(selector).is_visible(), f"{label} : non visible"


def check(page, url):
    errors = []
    console_errors = []
    STATE["send_body"] = None
    page.on("pageerror", lambda e: errors.append(f"pageerror: {e}"))
    page.on("console", lambda m: console_errors.append(m.text) if m.type == "error" else None)
    page.on("dialog", lambda d: d.accept())

    # 1. chargement sans erreur JS
    page.goto(url)
    page.wait_for_selector("#kiro-splash", state="detached", timeout=15000)
    page.wait_for_selector("#prompt-input", state="visible", timeout=8000)

    # 2. tour simulé : markdown rendu, tour terminé
    page.wait_for_selector("#chat-container .message-assistant .message-text h2", timeout=8000)
    title = page.locator("#chat-container .message-assistant .message-text h2").first.inner_text()
    assert title.strip() == "Titre", f"markdown titre = {title!r}"
    code = page.locator("#chat-container .message-assistant .message-text pre code").first.text_content()
    assert code and "console.log(1)" in code, f"bloc de code manquant : {code!r}"
    page.wait_for_selector("#chat-container[aria-busy='false']", timeout=8000)

    # 3. envoi : Entrée -> POST /api/chat/send avec le message
    page.wait_for_function(
        "document.getElementById('family-select') && document.getElementById('family-select').value === 'code'",
        timeout=8000,
    )
    page.fill("#prompt-input", "question smoke")
    page.press("#prompt-input", "Enter")
    deadline = time.time() + 8
    while STATE["send_body"] is None and time.time() < deadline:
        time.sleep(0.1)
    body = STATE["send_body"] or {}
    assert body.get("message") == "question smoke", f"POST /api/chat/send non émis : {body!r}"
    assert body.get("family") == "code" and body.get("mode") == "standard", f"sélection modèle : {body!r}"

    # 4. vue Agents en plein écran via le vrai bouton module
    page.locator('.dev-module-btn[data-module="agents"]').click()
    page.wait_for_selector("#marex-view.open", timeout=8000)
    vbox = page.locator("#marex-view").bounding_box()
    vp = page.viewport_size
    assert vbox["x"] == 0 and vbox["y"] == 0, f"vue Agents non calée : {vbox}"
    assert vbox["width"] >= vp["width"] and vbox["height"] >= vp["height"], "vue Agents non plein écran"

    # 5. menus du composer : ouverts ET visibles (anti-régression rognage)
    for btn_sel, menu_sel, label in [
        ("#mx-plus", "#mx-menu-plus", "menu +"),
        ("#mx-btn-model", "#mx-menu-model", "menu modèle"),
        ("#mx-btn-project", "#mx-menu-project", "menu projet"),
    ]:
        page.locator(btn_sel).click()
        page.wait_for_selector(f"{menu_sel}.open", timeout=4000)
        assert_visible_in_viewport(page, menu_sel, label)
        page.keyboard.press("Escape")
        page.wait_for_selector(f"{menu_sel}.open", state="hidden", timeout=4000)

    # 6. aucune erreur JS sur tout le scénario (les console.error sont
    # affichées à titre informatif mais ne font pas échouer : le serveur
    # réel peut répondre 404 sur des routes non simulées).
    for ce in console_errors[:5]:
        print(f"[info] console.error: {ce[:160]}")
    assert not errors, f"erreurs JS : {errors[:5]}"
    print("[ok] smoke : 6/6 vérifications")


def main():
    if not os.path.exists(BIN):
        subprocess.run(["make", "build"], cwd=ROOT, check=True)

    port = free_port()
    base = f"http://127.0.0.1:{port}"
    home = tempfile.mkdtemp(prefix="cetas-smoke-")
    env = dict(os.environ, CETAS_LITE_HOME=home, CETAS_LITE_ADDR=f"127.0.0.1:{port}")
    proc = subprocess.Popen([BIN, "serve"], cwd=ROOT, env=env, stdout=subprocess.DEVNULL, stderr=subprocess.STDOUT)
    try:
        if not wait_health(base):
            raise SystemExit("serveur non demarre")
        with sync_playwright() as p:
            browser = p.chromium.launch()
            page = browser.new_page(viewport={"width": 1280, "height": 800})
            page.add_init_script("try{localStorage.setItem('cetas-lite-token','smoke');}catch(e){}")
            route_mocks(page)
            check(page, base + "/")
            browser.close()
        print("SMOKE OK")
    finally:
        proc.terminate()
        try:
            proc.wait(timeout=5)
        except Exception:
            proc.kill()
        shutil.rmtree(home, ignore_errors=True)


if __name__ == "__main__":
    main()
