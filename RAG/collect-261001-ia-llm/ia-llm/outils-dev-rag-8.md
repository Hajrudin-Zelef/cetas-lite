---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-8
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1421, 1624]
sha256: 86514a60af7ec8265d3992aeddca58f1ddfc4acd974467626958eaf9e507934c
---

# Outils dev + ingénierie RAG (chunk & corpus)

`requests` + BeautifulSoup suffisent pour du HTML statique. Dès que le site
est une **SPA** (React/Vue/Angular — comme `support.huawei.com`, vérifié
dans ton expérience), le contenu est généré en JavaScript : `requests`
récupère une coquille vide. **Playwright** (Microsoft) pilote un **vrai
navigateur** (Chromium, Firefox, WebKit) : il exécute le JS, attend le
rendu, clique, remplit des formulaires, télécharge des fichiers.

| | requests | Playwright |
|---|---|---|
| HTML statique | ✅ rapide, léger | overkill |
| SPA / JS | ❌ coquille vide | ✅ rendu complet |
| Anti-bot basique | ⚠ headers manuels | ✅ vrai navigateur |
| Coût | ~0 RAM | **~200-600 Mo par instance** |
| Vitesse | ms | secondes (démarrage navigateur) |

**Règle :** `requests` d'abord (rapide, économe en crédits/temps),
Playwright quand `requests` rend une page vide ou un challenge.

## 46. Installation de Playwright (Python, vérifié)

```bash
# Dans ton venv
python -m venv .venv && source .venv/bin/activate
pip install --upgrade pip
pip install playwright          # vérifié : doc officielle Microsoft

# Télécharge les binaires navigateurs (Chromium, Firefox, WebKit)
playwright install
# Ou seulement Chromium (ton usage scraping) :
playwright install chromium

# Dépendances système manquantes (Debian/Ubuntu) :
playwright install --with-deps chromium
# ou : sudo playwright install-deps

# Vérifier
playwright --version
ls ~/.cache/ms-playwright/      # les binaires sont ICI, pas dans ton repo
```

> **Astuce serveur partagé :** `export PLAYWRIGHT_BROWSERS_PATH=/opt/playwright-browsers`
> + `sudo playwright install` une fois → tous les utilisateurs/venv
> partagent les mêmes binaires (plusieurs Go économisés).

## 47. Premier script Playwright commenté

`scripts/pw_premier.py` :

```python
"""Scrape minimal avec Playwright : titre + texte d'une page."""
from playwright.sync_api import sync_playwright

URL = "https://example.com"

def main() -> None:
    # sync_playwright() gère le cycle de vie (démarrage/arrêt du driver)
    with sync_playwright() as p:
        # launch() démarre Chromium. headless=True = sans fenêtre (serveur).
        browser = p.chromium.launch(headless=True)
        # Un "contexte" = une session isolée (cookies, cache) — comme une
        # fenêtre de navigation privée. Plusieurs contextes = plusieurs
        # identités dans le même navigateur.
        context = browser.new_context(
            # Pas de user_agent forcé ici : Playwright envoie celui du
            # navigateur qu'il a installé. Forcer un UA avec une version
            # figée (ex. Chrome/126) est une MAUVAISE pratique : la version
            # devient obsolète et trahit un bot. Si tu dois vraiment le
            # définir, recopie la version du navigateur installé.
            viewport={"width": 1366, "height": 768},
            locale="fr-FR",
        )
        page = context.new_page()

        # goto() ATTEND par défaut le chargement ("load").
        # Pour les SPA, préfère attendre un sélecteur précis (voir §49).
        page.goto(URL, wait_until="domcontentloaded", timeout=30000)

        print("Titre :", page.title())
        # inner_text = texte visible (rendu) ; content() = HTML brut
        print(page.locator("main").inner_text()[:2000])

        browser.close()

if __name__ == "__main__":
    main()
```

```bash
python scripts/pw_premier.py
```

**Version async** (pour paralléliser N pages) : remplace
`sync_playwright` par `async_playwright`, `await` devant chaque appel,
et lance avec `asyncio.run(main())`. Le mode async est indispensable dès
que tu veux 4+ workers comme ton `collect_still.py`.

## 48. Sélecteurs : viser juste

Playwright recommande cet ordre de priorité :

```python
# 1. Rôles (le plus robuste — suit l'accessibilité de la page)
page.get_by_role("button", name="Télécharger").click()
page.get_by_role("heading", name="Documentation").text_content()

# 2. Texte visible
page.get_by_text("Command Reference").first.click()

# 3. CSS (précis, mais cassant si le site change de classes)
titre = page.locator("h1.doc-title").inner_text()
lignes = page.locator("table.cmds tr").all_inner_texts()

# 4. XPath (dernier recours)
page.locator("xpath=//div[@id='content']//a").first.get_attribute("href")

# Bonnes pratiques
page.locator(".resultat").first        # .first évite le mode strict
page.locator(".resultat").nth(2)       # 3e élément
page.locator("a").filter(has_text="PDF")  # filtrer dans une liste
```

> **Mode strict :** `locator("h1")` lève une erreur s'il matche 0 ou 2+
> éléments. C'est une **fonctionnalité**, pas un bug : ça t'évite de
> scraper la mauvaise zone en silence.

## 49. Attentes : l'auto-waiting et ses limites

Playwright **attend automatiquement** avant chaque action (élément visible,
activé, stable). Tu n'as presque jamais besoin de `time.sleep()` :

```python
# ❌ INTERDIT (fragile, lent)
import time; time.sleep(5)

# ✅ Attente explicite ciblée
page.get_by_role("heading", name="Résultats").wait_for(timeout=15000)

# ✅ Attendre un état réseau (utile mais parfois trop long sur les SPA
#    qui pollent en continu — préfère un sélecteur quand possible)
page.goto(url, wait_until="networkidle")

# ✅ Attendre qu'une requête précise parte/revienne (top pour les SPA)
with page.expect_response("**/api/docs/**") as resp_info:
    page.get_by_role("button", name="Charger plus").click()
resp = resp_info.value
print(resp.status, resp.url)
data = resp.json()   # ← parfois l'API JSON est directement exploitable !
```

> **Insight collecte :** sur une SPA, l'onglet Réseau (ou
> `expect_response`) révèle souvent une **API JSON** propre. Si elle est
> accessible, tu peux la requêter en `requests` et **éviter Playwright**
> pour les pages suivantes (plus rapide, moins de RAM).

## 50. Contexte réaliste : headers et anti-bots DE BASE

Rester dans le **réaliste et légal** : tu te comportes comme un navigateur
normal (c'est ce que fait déjà Playwright), tu respectes `robots.txt` et
les CGU, tu espaces tes requêtes. Pas de contournement agressif.

```python
context = browser.new_context(
    # user_agent : non forcé par défaut (UA du navigateur installé).
    # À n'utiliser que si le site exige une valeur explicite, et dans ce cas
    # avec la version RÉELLE du navigateur installé — jamais une version figée.
    viewport={"width": 1366, "height": 768},
    locale="fr-FR",
    timezone_id="Europe/Paris",
    extra_http_headers={
        # Un vrai navigateur envoie toujours ces headers
        "Accept-Language": "fr-FR,fr;q=0.9,en;q=0.8",
    },
    # ignore_https_errors=True  # UNIQUEMENT en labo, jamais en prod
)
# Bloquer les ressources inutiles au scraping (images, fonts) = 2-3x plus rapide
page.route("**/*.{png,jpg,jpeg,gif,svg,woff,woff2}", lambda route: route.abort())
```

**Règles de politesse (anti-ban) :**

```python
import random, time
time.sleep(random.uniform(1.0, 3.0))   # délai aléatoire entre les pages
# 1 contexte = 1 "utilisateur" ; ne pas ouvrir 50 contextes en parallèle
# sur un petit site. 2 à 4 workers suffisent pour de la doc.
```

> Ton expérience `support.huawei.com` (WAF au comportement inverse) le
> montre : chaque site a ses particularités. **Teste petit** (5 URL),
> observe, puis industrialise. Et ne bourrine jamais les retries
> (risque de ban IP — tu l'as noté dans ta mémoire).

## 51. Screenshots et PDF

```python
# Screenshot : debug visuel ("qu'est-ce que voit vraiment le navigateur ?")
page.screenshot(path="debug/page.png", full_page=True)

# Screenshot d'un élément seul
page.locator("article.doc").screenshot(path="debug/article.png")

# PDF : pour archiver une doc telle que rendue
page.pdf(path="exports/doc.pdf", format="A4", print_background=True)

