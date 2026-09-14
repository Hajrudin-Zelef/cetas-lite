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

SETTINGS = {"theme": "ocean", "family": "code", "mode": "standard"}

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
    page.route("**/api/config", lambda r: r.fulfill(json={"registration_open": False, "version": "smoke"}))
    page.route("**/api/me", lambda r: r.fulfill(json={"username": "test", "role": "user"}))
    page.route("**/api/aliases", lambda r: r.fulfill(json=FAMILIES))

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
    page.wait_for_selector("#app:not([hidden])", state="visible", timeout=8000)
    page.wait_for_selector("#chat-log[aria-busy='false']", timeout=8000)
    page.wait_for_selector(".tool-diff .diff-add", state="attached", timeout=8000)
    page.wait_for_selector(".msg.assistant pre code", timeout=8000)

    title = page.locator(".msg.assistant h2").first.inner_text()
    assert title.strip() == "Titre", f"markdown titre = {title!r}"

    diff = page.locator(".tool-diff .diff-add").first.text_content()
    assert diff and "LIGNE" in diff, f"diff = {diff!r}"

    text = page.locator("#chat-log").inner_text()
    for marker in ["Titre", "gras", "console.log(1)"]:
        assert marker in text, f"contenu tronque, manque {marker!r}"

    busy = page.get_attribute("#chat-log", "aria-busy")
    assert busy == "false", f"aria-busy = {busy!r}"

    toggle = page.locator("#web-toggle")
    assert toggle.count() == 1, "web toggle absent"
    assert toggle.get_attribute("aria-pressed") == "false", "web toggle doit demarrer desactive"
    toggle.click()
    assert toggle.get_attribute("aria-pressed") == "true", "web toggle doit s'activer"
    page.wait_for_timeout(100)
    put = STATE["put_settings"] or {}
    assert put.get("web_default") is True, f"web_default non persiste: {put!r}"

    mcp = page.locator("#mcp-toggle")
    assert mcp.count() == 1, "mcp toggle absent"
    assert mcp.is_visible(), "mcp toggle doit etre visible si des serveurs sont configures"
    assert mcp.get_attribute("aria-pressed") == "true", "mcp doit demarrer actif (auto)"
    mcp.click()
    assert mcp.get_attribute("aria-pressed") == "false", "mcp doit se desactiver"
    page.wait_for_timeout(100)
    put = STATE["put_settings"] or {}
    assert put.get("mcp_default") is False, f"mcp_default non persiste: {put!r}"
    mcp.click()
    assert mcp.get_attribute("aria-pressed") == "true", "mcp doit se reactiver"
    page.wait_for_timeout(100)

    think = page.locator("#thinking-toggle")
    assert think.count() == 1, "thinking toggle absent"
    assert think.is_disabled(), "thinking doit etre force en mode agent"
    assert think.get_attribute("aria-pressed") == "true", "thinking doit etre actif en agent"
    assert page.locator("#effort-select").input_value() == "default", "effort defaut attendu"

    stats = page.locator("#stats-badge")
    assert stats.is_visible(), "stats badge visible attendu"
    stext = stats.inner_text()
    assert "10" in stext and "5" in stext, f"stats = {stext!r}"

    assert page.locator(".msg-system").count() >= 1, "notice de compaction attendue"

    assert page.locator(".conv-item").count() == 1, "une conversation archivee attendue"
    ctitle = page.locator(".conv-title").first.inner_text()
    assert "Ancienne question" in ctitle, f"titre archive = {ctitle!r}"

    page.locator(".conv-del").first.click()
    page.wait_for_timeout(100)
    assert STATE.get("deleted") == "20260101_120000_1", f"DELETE archive attendu: {STATE.get('deleted')!r}"

    assert page.locator('.tool-result a[href="https://go.dev"]').count() >= 1, "citation cliquable attendue"

    assert page.locator(".msg-actions").count() >= 1, "actions message attendues"
    assert page.locator(".msg-action", has_text="Lire").count() >= 1, "bouton TTS (navigateur) attendu"
    assert page.locator("#mic-btn").count() == 1, "bouton micro present"
    page.locator(".msg-action", has_text="Regenerer").last.click()
    page.wait_for_timeout(100)
    assert STATE.get("regenerated"), "regeneration attendue"

    page.locator("#export-btn").click()
    page.wait_for_timeout(100)
    assert STATE.get("exported"), "export actif attendu"

    page.locator("#settings-btn").click()
    page.wait_for_selector("#settings-overlay:not([hidden])", timeout=4000)
    page.wait_for_timeout(200)
    assert page.locator("#mcp-panel .mcp-name").count() == 1, "serveur MCP liste attendu"
    mcpname = page.locator("#mcp-panel .mcp-name").first.inner_text()
    assert mcpname == "demo", f"nom MCP = {mcpname!r}"
    caps = page.locator("#caps-panel .caps-name")
    assert caps.count() >= 1, "capacites modeles attendues"
    assert "fake/ok" == page.locator("#caps-panel .caps-name").first.inner_text(), "modele d'alias attendu"
    page.locator("#settings-close").click()

    page.fill("#prompt-input", "question web")
    page.press("#prompt-input", "Enter")
    page.wait_for_timeout(100)
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
