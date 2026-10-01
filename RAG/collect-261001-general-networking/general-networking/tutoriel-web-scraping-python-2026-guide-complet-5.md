---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-5
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmark"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [491, 625]
sha256: bad88d6629f05b11af3932de4dffcf5efa64bc1a82ad5cf5706d5b34b30b8e70
---

# Créer le dossier du projet

```
import asyncio
import httpx
from bs4 import BeautifulSoup
import time
from typing import Any
# Sémaphore pour limiter les connexions simultanées
# Évite de surcharger le serveur cible et d'être bloqué
MAX_CONNEXIONS_SIMULTANEES = 10
async def scraper_url_async(
    client: httpx.AsyncClient,
    url: str,
    semaphore: asyncio.Semaphore
) -> dict[str, Any]:
    """
    Scrape une seule URL de manière asynchrone.
    Utilise un sémaphore pour respecter les limites de taux.
    Cœur de l'automatisation web Python haute performance.
    """
    async with semaphore:  # Acquérir le sémaphore avant la requête
        try:
            response = await client.get(url)
            response.raise_for_status()
            soup = BeautifulSoup(response.text, "lxml")
            titre = soup.find("title")
            description = soup.find("meta", {"name": "description"})
            return {
                "url": url,
                "statut": response.status_code,
                "titre": titre.text.strip() if titre else "",
                "description": description.get("content", "") if description else "",
                "taille_html": len(response.text),
                "erreur": None
            }
        except httpx.HTTPStatusError as e:
            return {"url": url, "statut": e.response.status_code, "erreur": str(e)}
        except httpx.RequestError as e:
            return {"url": url, "statut": 0, "erreur": str(e)}
async def scraper_multiple_urls(
    urls: list[str],
    max_concurrent: int = MAX_CONNEXIONS_SIMULTANEES
) -> list[dict]:
    """
    Scrape plusieurs URLs en parallèle avec httpx + asyncio.
    Utilise asyncio.gather() pour lancer toutes les tâches simultanément
    tout en respectant la limite de connexions via un sémaphore.
    Args:
        urls: Liste des URLs à scraper
        max_concurrent: Nombre maximum de requêtes simultanées
    Returns:
        Liste des résultats de scraping
    """
    semaphore = asyncio.Semaphore(max_concurrent)
    headers = {
        "User-Agent": "Mozilla/5.0 (compatible; AsyncScraper/1.0)",
        "Accept": "text/html,application/xhtml+xml",
    }
    limits = httpx.Limits(
        max_connections=max_concurrent * 2,
        max_keepalive_connections=max_concurrent
    )
    debut = time.time()
    async with httpx.AsyncClient(
        headers=headers,
        limits=limits,
        timeout=30.0,
        follow_redirects=True
    ) as client:
        # Créer toutes les tâches et les lancer en parallèle avec asyncio.gather()
        taches = [
            scraper_url_async(client, url, semaphore)
            for url in urls
        ]
        resultats = await asyncio.gather(*taches, return_exceptions=False)
    duree = time.time() - debut
    succes = sum(1 for r in resultats if r.get("erreur") is None)
    print(f"\nScraping terminé en {duree:.2f}s")
    print(f"{succes}/{len(urls)} URLs scrapées avec succès")
    print(f"Débit moyen : {len(urls)/duree:.1f} pages/seconde")
    return list(resultats)
# Démonstration du gain de performance async vs sync
async def demo_performance():
    """Compare les performances async vs sync sur un ensemble d'URLs réelles."""
    urls = [f"https://books.toscrape.com/catalogue/page-{i}.html" for i in range(1, 16)]
    print(f"Test de performance sur {len(urls)} URLs en mode asynchrone...")
    debut = time.time()
    resultats = await scraper_multiple_urls(urls, max_concurrent=8)
    duree_async = time.time() - debut
    print(f"\nRésultats du benchmark :")
    print(f"  Mode ASYNC : {duree_async:.2f}s pour {len(urls)} pages")
    print(f"  Mode SYNC estimé : ~{len(urls) * 1.8:.1f}s (1.8s/page moyenne)")
    print(f"  Gain de vitesse : ~{(len(urls) * 1.8) / duree_async:.1f}x plus rapide")
    return resultats
if __name__ == "__main__":
    resultats = asyncio.run(demo_performance())
    for r in resultats[:3]:
        print(f"  {r.get('url', '')[-40:]} -> HTTP {r.get('statut', '?')}")
```
## Comparatif Complet des Outils de Web Scraping Python en 2026

Choisir le bon outil est crucial pour la réussite de votre projet de **scraping python**. Voici une comparaison détaillée des principales solutions disponibles en 2026, basée sur des critères objectifs de performance, facilité d'utilisation et cas d'usage réels.

| Outil | Type | Vitesse (pages/sec) | JS Rendering | Courbe apprentissage | Idéal Pour | 
|---|---|---|---|---|---|
| **requests + BeautifulSoup** | Bibliothèques | 1-5 | Non | Très facile | Débutants, sites statiques simples | 
| **httpx + asyncio** | Bibliothèques | 15-150 | Non | Intermédiaire | Multi-URLs, performance critique | 
| **Scrapy 2.12** | Framework complet | 50-200 | Via scrapy-playwright | Avancée | Projets larges, pipeline de données | 
| **Playwright 1.49** | Automatisation navigateur | 1-3 | Oui (natif) | Intermédiaire | SPAs, sites JavaScript-heavy | 
| **Selenium 4.x** | Automatisation navigateur | 0.5-2 | Oui | Intermédiaire | Legacy, interactions complexes | 
| **Selectolax** | Parser HTML | 30-150 (parsing) | Non | Facile | Parsing haute performance | 
| **curl_cffi** | Client HTTP spécialisé | 5-20 | Non | Facile | Bypass TLS fingerprinting Cloudflare | 
| **Crawl4AI** | Framework IA | Variable | Oui (Playwright) | Facile | Extraction IA, output Markdown/JSON | 

| Critère Technique | BeautifulSoup 4.12 | Scrapy 2.12 | Playwright 1.49 | Selenium 4.x | 
|---|---|---|---|---|
| GitHub Stars (2026) | – | 54 800+ | 71 500+ | 30 000+ | 
| Consommation mémoire | Faible (~50 MB) | Moyenne (~120 MB) | Élevée (~500 MB) | Très élevée (~700 MB) | 
| Support HTTP/2 | Via httpx | Partiel | Oui (natif) | Non | 
| Support async natif | Non | Oui (Twisted) | Oui (asyncio) | Non | 
| Anti-détection | Manuel (headers) | Via middlewares | Via stealth plugins | Via undetected-chromedriver | 
| Facilité de déploiement | Simple | Moyen (Scrapyd) | Simple (Docker) | Complexe (ChromeDriver) | 
| Écosystème/plugins | Riche | Très riche | Croissant | Mature mais vieillissant | 

## Techniques Anti-Bot et Comment les Gérer Éthiquement

Les sites web déploient de plus en plus de mécanismes sophistiqués pour détecter et bloquer les scrapers automatisés. Comprendre ces techniques est essentiel, non pas pour les contourner de manière abusive, mais pour construire des scrapers robustes qui fonctionnent de manière fiable et respectueuse.

### Les Principales Techniques de Détection Anti-Bot

**TLS/JA3 Fingerprinting :** Chaque client HTTP présente une signature TLS unique lors de la négociation SSL. La bibliothèque `requests` a une signature JA3 très distincte des navigateurs réels, ce qui permet aux CDN comme Cloudflare de la détecter instantanément. La solution en 2026 est `curl_cffi`, qui impersonne les signatures TLS de navigateurs réels (Chrome, Firefox, Safari) de manière transparente.

**Headers HTTP Analysis :** Les navigateurs envoient des dizaines de headers spécifiques dans un ordre précis (ordre des headers). L'absence de certains headers (`Sec-Fetch-*`, `Accept-CH`), ou leur présence dans le mauvais ordre, est un signal fort de scraping automatisé. `curl_cffi` gère également cet aspect en reproduisant fidèlement les headers des navigateurs cibles.

**CAPTCHA et Challenges JavaScript :** Cloudflare Turnstile, reCAPTCHA v3, hCAPTCHA sont des défis courants. Pour Cloudflare JS Challenge, `playwright-stealth` ou le plugin `undetected-playwright` peuvent contourner certaines protections de niveau 1. Pour les CAPTCHAs d'images nécessitant une interaction humaine, des services tiers existent mais leur utilisation doit être pesée éthiquement et peut violer les CGU.

