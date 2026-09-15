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
                    "pool": [
                        {"provider": "fake", "model": "ok", "label": "Fake", "input_per_1m": 0, "output_per_1m": 0}
                    ],
                }
            ],
        }
    ]
}

SETTINGS = {"theme": "ocean", "family": "code", "mode": "standard", "web_default": False, "mcp_default": True}

CATALOG = {
    "providers": [
        {
            "id": "deepseek",
            "label": "DeepSeek",
            "models": [
                {"id": "deepseek-flash", "label": "DeepSeek Flash", "input_per_1m": 0.15, "output_per_1m": 0.6},
            ],
        }
    ]
}

PROVIDERS = {"providers": [{"id": "deepseek", "label": "DeepSeek", "configured": True}]}

MCP_SERVERS = {"servers": [{"name": "demo", "transport": "stdio", "connected": True, "tools": 2, "error": ""}]}

CONVERSATIONS = {
    "archives": [
        {"id": "20260101_120000_1", "title": "Ancienne question", "updated": 1735732800000, "messages": 4}
    ]
}

STATE = {"put_settings": None, "send_body": None, "deleted": None, "regenerated": None, "exported": None}


def sse(events):
    return "".join("data: " + json.dumps(e) + "\n\n" for e in events)


STREAM = sse(
    [
        {"seq": 1, "user": "salut"},
        {"seq": 2, "reasoning_content": "je reflechis"},
        {"seq": 3, "content": "## Titre\n\n**gras** et texte\n"},
        {"seq": 4, "tool": {"name": "Edit", "args": {"file_path": "a.txt"}, "phase": "start"}},
        {
            "seq": 5,
            "tool": {
                "name": "Edit",
                "args": {"file_path": "a.txt"},
                "phase": "end",
                "result": "ok https://go.dev doc",
                "diff": [{"kind": "-", "text": "ligne"}, {"kind": "+", "text": "LIGNE"}],
            },
        },
        {"seq": 6, "content": "\n\n```js\nconsole.log(1)\n```\n"},
        {"seq": 7, "stats": {"prompt_tokens": 10, "completion_tokens": 5}},
        {"seq": 8, "compact": True},
        {"seq": 9, "turn_done": True, "elapsed_ms": 12},
    ]
)


def free_port():
    s = socket.socket()
    s.bind(("127.0.0.1", 0))
    port = s.getsockname()[1]
    s.close()
    return port


def wait_health(base, timeout=10):
    deadline = time.time() + timeout
    while time.time() < deadline:
        try:
            with urllib.request.urlopen(base + "/api/health", timeout=1) as r:
                if r.status == 200:
                    return True
        except Exception:
            time.sleep(0.1)
    return False


def route_mocks(page):
    # Catch-all : tout endpoint /api non mocke explicitement renvoie un objet vide
    # (evite les 401 -> deconnexion pendant l'init). Les routes specifiques,
    # enregistrees ensuite, ont la priorite.
    page.route("**/api/**", lambda r: r.fulfill(json={}))
    page.route("**/api/config", lambda r: r.fulfill(json={"registration_open": False, "version": "smoke"}))
    page.route("**/api/me", lambda r: r.fulfill(json={"username": "test", "role": "user"}))
    page.route("**/api/aliases", lambda r: r.fulfill(json=FAMILIES))
    page.route("**/api/catalog", lambda r: r.fulfill(json=CATALOG))
    page.route("**/api/providers", lambda r: r.fulfill(json=PROVIDERS))

    def settings(route):
        if route.request.method == "PUT":
            STATE["put_settings"] = route.request.post_data_json
            route.fulfill(json={"ok": True})
        else:
            route.fulfill(json=SETTINGS)

    def send(route):
        STATE["send_body"] = route.request.post_data_json
        route.fulfill(json={"ok": True})

    def conversations(route):
        req = route.request
        if req.method == "DELETE":
            STATE["deleted"] = req.url.rsplit("/", 1)[-1]
            route.fulfill(json={"ok": True})
        elif req.method == "POST":
            route.fulfill(json={"ok": True})
        else:
            route.fulfill(json=CONVERSATIONS)

    page.route("**/api/settings", settings)
    page.route("**/api/mcp", lambda r: r.fulfill(json=MCP_SERVERS))
    page.route("**/api/capabilities", lambda r: r.fulfill(json={"caps": {}}))
    page.route("**/api/chat/send", send)
    page.route("**/api/chat/state", lambda r: r.fulfill(json={"turns": 1, "generating": False}))
    page.route("**/api/conversations**", conversations)

    def regenerate(route):
        STATE["regenerated"] = True
        route.fulfill(json={"ok": True})

    page.route("**/api/chat/regenerate", regenerate)

    def export(route):
        STATE["exported"] = True
        route.fulfill(status=200, headers={"Content-Type": "text/markdown"}, body="# Conversation\n")

    page.route("**/api/chat/export**", export)
    page.route(
        "**/api/chat/stream**",
        lambda r: r.fulfill(status=200, headers={"Content-Type": "text/event-stream", "Cache-Control": "no-cache"}, body=STREAM),
    )


def check(page, url, reduced):
    errors = []
    STATE["put_settings"] = None
    STATE["send_body"] = None
    STATE["deleted"] = None
    STATE["regenerated"] = None
    STATE["exported"] = None
    page.on("pageerror", lambda e: errors.append(str(e)))
    page.on("dialog", lambda d: d.accept())
    page.goto(url)
    # Le splash disparait et le chat devient visible.
    page.wait_for_selector("#kiro-splash", state="detached", timeout=8000)
    page.wait_for_selector("#chat-container", state="visible", timeout=8000)
    page.wait_for_selector(".msg-tool .tool-diff .diff-add", state="attached", timeout=8000)
    page.wait_for_selector(".message-wrapper-assistant .message-text pre code", timeout=8000)

    title = page.locator(".message-wrapper-assistant .message-text h2").first.inner_text()
    assert title.strip() == "Titre", f"markdown titre = {title!r}"

    diff = page.locator(".tool-diff .diff-add").first.text_content()
    assert diff and "LIGNE" in diff, f"diff = {diff!r}"

    text = page.locator("#chat-container").inner_text()
    for marker in ["Titre", "gras", "console.log(1)"]:
        assert marker in text, f"contenu tronque, manque {marker!r}"

    # Toggle web (checkbox dans le menu plus, hors ecran : pilotage via change)
    web = page.locator("#web-toggle")
    assert web.count() == 1, "web toggle absent"
    assert not web.is_checked(), "web toggle doit demarrer desactive"
    page.eval_on_selector("#web-toggle", "el => { el.checked = true; el.dispatchEvent(new Event('change', {bubbles:true})); }")
    assert web.is_checked(), "web toggle doit s'activer"
    page.wait_for_timeout(100)
    put = STATE["put_settings"] or {}
    assert put.get("web_default") is True, f"web_default non persiste: {put!r}"

    # Toggle MCP (checkbox) : actif par defaut (serveurs configures)
    mcp = page.locator("#mcp-toggle")
    assert mcp.count() == 1, "mcp toggle absent"
    assert mcp.is_checked(), "mcp doit demarrer actif (auto)"

    # Bascule top-level en mode Agent : thinking verrouille/actif.
    page.eval_on_selector("#mode-agent-btn", "el => el.click()")
    page.wait_for_timeout(100)
    think = page.locator("#thinking-toggle")
    assert think.count() == 1, "thinking toggle absent"
    assert think.is_disabled(), "thinking doit etre force en mode agent"
    assert think.is_checked(), "thinking doit etre actif en agent"
    assert page.locator("#effort-select").input_value() == "default", "effort defaut attendu"

    stats = page.locator("#cost-info")
    assert stats.is_visible(), "stats badge visible attendu"
    stext = stats.inner_text()
    assert "10" in stext and "5" in stext, f"stats = {stext!r}"

    assert page.locator(".msg-system").count() >= 1, "notice de compaction attendue"

    assert page.locator(".conv-item").count() == 1, "une conversation archivee attendue"
    ctitle = page.locator(".conv-item-title").first.inner_text()
    assert "Ancienne question" in ctitle, f"titre archive = {ctitle!r}"

    page.eval_on_selector(".conv-item-actions .danger", "el => el.click()")
    page.wait_for_selector("#custom-dialog-overlay", state="visible", timeout=4000)
    page.eval_on_selector("#custom-dialog-ok", "el => el.click()")
    page.wait_for_timeout(150)
    assert STATE.get("deleted") == "20260101_120000_1", f"DELETE archive attendu: {STATE.get('deleted')!r}"

    assert page.locator('.tool-result a[href="https://go.dev"]').count() >= 1, "citation cliquable attendue"

    assert page.locator(".message-btn-row").count() >= 1, "actions message attendues"
    assert page.locator(".message-tts-btn").count() >= 1, "bouton TTS (navigateur) attendu"
    assert page.locator("#mic-btn").count() == 1, "bouton micro present"
    page.eval_on_selector(".regen-btn", "el => el.click()")
    page.wait_for_timeout(100)
    assert STATE.get("regenerated"), "regeneration attendue"

    # Export via le menu partage.
    page.eval_on_selector("#share-btn", "el => el.click()")
    page.eval_on_selector("#share-menu-md", "el => el.click()")
    page.wait_for_timeout(150)
    assert STATE.get("exported"), "export actif attendu"

    # Modale de reglages (API et Modeles).
    page.locator("#apikeys-btn").click()
    page.wait_for_selector("#apikeys-modal-overlay", state="visible", timeout=4000)
    page.wait_for_timeout(300)
    assert page.locator(".apikeys-tabs .apikeys-tab").count() >= 4, "onglets de reglages attendus"
    assert page.locator("#providers-list .provider-row").count() >= 1, "providers listes attendus"
    assert page.locator("#families-list .family-row").count() >= 1, "familles listees attendues"
    page.locator("#apikeys-close-btn").click()

    page.fill("#prompt-input", "question web")
    page.press("#prompt-input", "Enter")
    page.wait_for_timeout(150)
    body = STATE["send_body"] or {}
    assert body.get("web") is True, f"le tour doit porter web=true: {body!r}"
    assert body.get("mcp") is True, f"le tour doit porter mcp=true: {body!r}"
    assert body.get("think") is True, f"le mode agent doit porter think=true: {body!r}"
    assert body.get("effort") == "default", f"effort par defaut attendu: {body!r}"

    assert not errors, f"erreurs page: {errors}"
    label = "reduced" if reduced else "normal"
    print(f"[ok] smoke {label}")


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
        url = base + "/"
        with sync_playwright() as p:
            browser = p.chromium.launch()
            for reduced in (False, True):
                ctx = browser.new_context(reduced_motion="reduce" if reduced else "no-preference")
                page = ctx.new_page()
                page.add_init_script("try{localStorage.setItem('cetas-lite-token','smoke');}catch(e){}")
                route_mocks(page)
                check(page, url, reduced)
                ctx.close()
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
