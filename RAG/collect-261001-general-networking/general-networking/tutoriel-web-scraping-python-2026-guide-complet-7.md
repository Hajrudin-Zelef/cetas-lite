---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-7
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [666, 779]
sha256: 2b73bb42f99cc8a7086a4d2f955eb4e0c99be25975270c3dc7ada2343047ecfc
---

# Créer le dossier du projet

```
import httpx
import time
import random
from functools import wraps
from typing import Callable, Any
def avec_retry(
    max_tentatives: int = 3,
    delai_initial: float = 1.0,
    facteur_backoff: float = 2.0,
    exceptions_retentables: tuple = (httpx.RequestError, httpx.HTTPStatusError)
):
    """
    Décorateur de retry avec backoff exponentiel et jitter.
    Indispensable pour tout scraper Python robuste en production.
    Args:
        max_tentatives: Nombre maximum de tentatives avant abandon
        delai_initial: Délai initial entre les tentatives (secondes)
        facteur_backoff: Multiplicateur du délai à chaque tentative
        exceptions_retentables: Types d'exceptions qui déclenchent un retry
    """
    def decorateur(func: Callable) -> Callable:
        @wraps(func)
        def wrapper(*args, **kwargs) -> Any:
            delai = delai_initial
            derniere_exception = None
            for tentative in range(1, max_tentatives + 1):
                try:
                    return func(*args, **kwargs)
                except exceptions_retentables as e:
                    derniere_exception = e
                    # Vérifier si c'est une erreur 429 avec header Retry-After
                    if isinstance(e, httpx.HTTPStatusError):
                        if e.response.status_code == 429:
                            retry_after = e.response.headers.get("Retry-After", str(delai))
                            try:
                                delai = float(retry_after)
                                print(f"  Serveur demande d'attendre {delai}s (Retry-After header)")
                            except ValueError:
                                pass
                        elif e.response.status_code in (400, 401, 403, 404):
                            # Ces erreurs ne sont pas temporaires : inutile de réessayer
                            print(f"  Erreur permanente {e.response.status_code} – abandon")
                            raise
                    if tentative < max_tentatives:
                        # Ajouter du jitter pour éviter les requêtes synchronisées
                        jitter = random.uniform(0, delai * 0.2)
                        attente = delai + jitter
                        print(f"  Tentative {tentative}/{max_tentatives} échouée : {type(e).__name__}")
                        print(f"  Prochain essai dans {attente:.1f}s...")
                        time.sleep(attente)
                        # Augmenter le délai (backoff exponentiel)
                        delai *= facteur_backoff
                    else:
                        print(f"  Toutes les {max_tentatives} tentatives épuisées – abandon définitif")
                        raise derniere_exception
            raise derniere_exception
        return wrapper
    return decorateur
# Application du décorateur à une fonction de scraping réelle
@avec_retry(max_tentatives=3, delai_initial=2.0, facteur_backoff=2.0)
def requete_robuste(url: str) -> httpx.Response:
    """Effectue une requête HTTP avec retry automatique en cas d'échec."""
    with httpx.Client(timeout=30.0, headers={
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"
    }) as client:
        response = client.get(url)
        response.raise_for_status()
        return response
# Test du mécanisme de retry
if __name__ == "__main__":
    try:
        response = requete_robuste("https://httpbin.org/status/200")
        print(f"Succès : HTTP {response.status_code}")
    except Exception as e:
        print(f"Échec définitif : {e}")
```
## Astuces Avancées : Proxies, Fingerprinting et Scraping IA avec Crawl4AI

Une fois les bases maîtrisées, plusieurs techniques avancées permettent de construire des scrapers plus robustes, plus furtifs et plus intelligents. Ces techniques sont particulièrement pertinentes pour les projets à grande échelle ou ciblant des sites avec des protections anti-bot sophistiquées.

### Proxies Rotatifs et Gestion des Identités

L'utilisation de proxies rotatifs est incontournable pour les scrapers à grande échelle. Plusieurs services commerciaux existent : Bright Data, Oxylabs, Smartproxy, et IPRoyal offrent des pools de millions d'adresses IP résidentielles et datacenter. Les proxies résidentiels sont beaucoup plus difficiles à détecter car ils proviennent de vraies connexions d'utilisateurs domestiques, mais ils sont aussi plus coûteux (environ 2-15$/GB selon le fournisseur).

Avec `httpx`, la configuration d'un proxy se fait via le paramètre `proxies` du client : `httpx.Client(proxies={"https://": "http://user:[email protected]:8080"})`. Pour les proxies rotatifs, chaque nouvelle instance de client ou chaque requête avec rotation active obtient automatiquement une nouvelle IP depuis le pool.

### curl_cffi : Le Bypass de TLS Fingerprinting

`curl_cffi` est devenu la bibliothèque de référence pour contourner la détection basée sur le fingerprinting TLS/JA3. Elle fonctionne en utilisant la bibliothèque C libcurl avec le support de curl-impersonate, qui reproduit fidèlement les handshakes TLS des navigateurs Chrome, Firefox et Safari. L'API est presque identique à `requests` avec un paramètre supplémentaire `impersonate` :

`from curl_cffi import requests as cf_requests` puis `response = cf_requests.get(url, impersonate="chrome131")`. Cette seule modification peut résoudre 80% des problèmes de blocage Cloudflare sur les sites sans CAPTCHA interactif.

### Crawl4AI : Le Scraping Assisté par Intelligence Artificielle

**Crawl4AI** représente la prochaine génération d'outils de scraping. Avec 18 000+ étoiles GitHub en 2026, cette bibliothèque combine Playwright pour le rendu JavaScript avec des LLMs (Large Language Models) pour extraire des données structurées depuis n'importe quel site web, sans nécessiter la définition manuelle de sélecteurs CSS.

L'approche est fondamentalement différente des scrapers traditionnels : au lieu de définir des sélecteurs pour chaque site, vous décrivez en langage naturel ce que vous voulez extraire, et Crawl4AI utilise un LLM pour comprendre la structure sémantique de la page et retourner les données en JSON ou Markdown. Particulièrement utile pour les sites sans structure HTML prévisible, les articles de blog, ou les pages de profils variés.

Pour une intégration avec des modèles de langage locaux (sans coût API), consultez notre Tutoriel Ollama 2026 qui couvre l'installation et l'utilisation de LLMs open source en local – parfait pour un pipeline Crawl4AI + LLM local entièrement autohébergé.

### Couverture Connexe

Pour aller plus loin dans votre stack Python et web, ces tutoriels complètent parfaitement vos compétences en **automatisation web Python** et en construction de pipelines de données :

- Tutoriel FastAPI Python 2026 – Exposez vos données scrapées via une API REST performante avec validation automatique
- Tutoriel Streamlit Python 2026 – Créez des dashboards interactifs pour visualiser vos datasets en quelques lignes
- Tutoriel Ollama 2026 – Intégrez des LLMs locaux pour analyser et structurer vos données scrapées sans coût API
- Playwright vs Cypress vs Selenium 2026 – Comparatif complet pour choisir le meilleur outil d'automatisation navigateur
- Tutoriel Astro JS 2026 – Construisez des sites web ultra-performants pour présenter vos données collectées
- Guide des Outils de Coding IA – Accélérez votre développement de scrapers avec les meilleurs assistants IA de 2026

## Scrapy 2.12 : Le Framework Industriel pour le Scraping Python à Grande Échelle

