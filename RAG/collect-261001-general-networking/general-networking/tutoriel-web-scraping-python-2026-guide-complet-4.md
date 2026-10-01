---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-4
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "sandbox"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [366, 490]
sha256: f986a91a387f1a2e8fbc18812043e8d3e6784bab99dd273c4b39b8062e73dd45
---

# Créer le dossier du projet

Une large proportion des sites web modernes – plateformes e-commerce, réseaux sociaux, dashboards SaaS – reposent massivement sur JavaScript pour afficher leur contenu. Dans ces cas, une simple requête HTTP retourne un HTML vide ou minimal, le contenu étant chargé dynamiquement via des appels API JavaScript. **Playwright** est la solution de référence en 2026 pour ce type de scraping. Pour une comparaison complète des outils d'automatisation de navigateur, consultez notre guide Playwright vs Cypress vs Selenium 2026.

**Étape 9 : Configuration et premier scraping avec Playwright**

Playwright, développé par Microsoft, supporte Chromium, Firefox et WebKit. Il offre une API Python moderne avec support natif de l'asynchronisme via `asyncio`. Avec 71 500+ étoiles GitHub en 2026, c'est sans conteste l'outil le plus adopté pour l'automatisation de navigateur ; sa version Python **1.61.0**, la plus récente en juin 2026 selon WebScraping.ai, améliore encore l'auto-waiting et embarque directement les navigateurs. La documentation officielle Playwright Python est excellente et propose des exemples pour tous les cas d'usage.

```
import asyncio
from playwright.async_api import async_playwright, Browser, Page
import json
async def scraper_avec_playwright(url: str) -> list[dict]:
    """
    Scrape un site JavaScript-heavy avec Playwright.
    Utilise le mode asynchrone pour de meilleures performances.
    Technique indispensable pour l'extraction données web Python moderne.
    """
    resultats = []
    async with async_playwright() as p:
        # Lancer Chromium en mode headless (sans interface graphique)
        browser: Browser = await p.chromium.launch(
            headless=True,
            args=[
                "--no-sandbox",
                "--disable-setuid-sandbox",
                "--disable-dev-shm-usage",  # Important pour Docker/CI
            ]
        )
        # Créer un contexte de navigateur avec paramètres réalistes
        context = await browser.new_context(
            viewport={"width": 1920, "height": 1080},
            locale="fr-FR",
            timezone_id="Europe/Paris",
            user_agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
        )
        # Bloquer les ressources inutiles pour accélérer le scraping
        await context.route("**/*.{png,jpg,jpeg,gif,svg,css,font,woff,woff2}",
                           lambda route: route.abort())
        page: Page = await context.new_page()
        try:
            # Naviguer vers la page et attendre que le réseau soit inactif
            await page.goto(url, wait_until="networkidle", timeout=30000)
            # Attendre qu'un sélecteur spécifique soit visible
            await page.wait_for_selector("article", timeout=10000)
            # Exécuter du JavaScript dans le contexte de la page
            # Utile pour extraire des données depuis window.__NEXT_DATA__ (Next.js)
            donnees_js = await page.evaluate("""
                () => {
                    if (window.__NEXT_DATA__) {
                        return window.__NEXT_DATA__.props;
                    }
                    return null;
                }
            """)
            if donnees_js:
                print("Données Next.js trouvées dans window.__NEXT_DATA__")
            # Extraire le contenu HTML rendu par JavaScript
            contenu_html = await page.content()
            # Parser avec BeautifulSoup pour l'extraction
            from bs4 import BeautifulSoup
            soup = BeautifulSoup(contenu_html, "lxml")
            articles = soup.find_all("article", class_="product_pod")
            for article in articles:
                lien = article.select_one("h3 > a")
                prix = article.select_one("p.price_color")
                resultats.append({
                    "titre": lien["title"] if lien else "",
                    "prix": prix.text.strip() if prix else "",
                })
            print(f"{len(resultats)} éléments extraits")
        except Exception as e:
            print(f"Erreur Playwright : {e}")
            await page.screenshot(path="screenshot_erreur.png")
        finally:
            await context.close()
            await browser.close()
    return resultats
# Étape 10 : Gérer le défilement infini (infinite scroll) avec Playwright
async def scraper_infinite_scroll(url: str, max_scrolls: int = 10) -> list[dict]:
    """
    Gère le défilement infini en simulant le scroll de l'utilisateur.
    Pattern courant sur les réseaux sociaux et les marketplaces modernes.
    """
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=True)
        page = await browser.new_page()
        await page.goto(url, wait_until="domcontentloaded")
        elements_vus = set()
        tous_elements = []
        hauteur_precedente = 0
        for i in range(max_scrolls):
            # Scroller jusqu'au bas de la page
            await page.evaluate("window.scrollTo(0, document.body.scrollHeight)")
            # Attendre le chargement du nouveau contenu
            await page.wait_for_timeout(2000)
            # Extraire les éléments présents sur la page
            contenu = await page.content()
            from bs4 import BeautifulSoup
            soup = BeautifulSoup(contenu, "lxml")
            for el in soup.find_all("article", class_="product_pod"):
                texte = el.get_text(strip=True)
                if texte not in elements_vus:
                    elements_vus.add(texte)
                    lien = el.select_one("h3 > a")
                    tous_elements.append({"titre": lien["title"] if lien else texte[:50]})
            # Vérifier si la hauteur de la page a changé
            hauteur_actuelle = await page.evaluate("document.body.scrollHeight")
            print(f"  Scroll {i+1}/{max_scrolls} : {len(tous_elements)} éléments, hauteur {hauteur_actuelle}px")
            if hauteur_actuelle == hauteur_precedente:
                print("  Fin du contenu détectée – arrêt du scroll")
                break
            hauteur_precedente = hauteur_actuelle
        await browser.close()
        return tous_elements
if __name__ == "__main__":
    resultats = asyncio.run(scraper_avec_playwright("https://books.toscrape.com/"))
    print(f"Résultats : {len(resultats)} livres extraits")
```
## Étapes 11-12 : Scraping Asynchrone Haute Performance avec httpx et asyncio

Le scraping séquentiel – une requête à la fois – est souvent insuffisant pour les projets nécessitant de collecter des données depuis des dizaines ou centaines d'URLs. Le **scraping asynchrone avec Python** permet de lancer plusieurs requêtes HTTP en parallèle grâce à `asyncio`, multipliant les performances par un facteur 10 à 50 selon les contraintes réseau.

Python 3.12 a apporté des améliorations significatives à `asyncio` : meilleure gestion des exceptions dans les tâches concurrentes, `asyncio.TaskGroup` pour une gestion structurée de la concurrence, et des performances globales améliorées. La documentation officielle d'asyncio est une ressource incontournable pour approfondir ces concepts.

Si vous souhaitez visualiser vos données scrapées dans un dashboard interactif, notre Tutoriel Streamlit Python 2026 vous montrera comment créer des interfaces web sans écrire une ligne de JavaScript.

