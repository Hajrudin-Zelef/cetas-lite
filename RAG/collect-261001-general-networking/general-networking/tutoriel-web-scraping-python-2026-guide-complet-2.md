---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-2
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [73, 202]
sha256: c4b384ed978cb73ed913843b9b5e4072388b60c4375c79c240a16de98b918059
---

# Créer le dossier du projet

Installez maintenant l’ensemble des bibliothèques dont vous aurez besoin pour ce tutoriel. Le guide du scraper Python de Tech Insider, mis à jour en août 2026, recommande d’installer **7** paquets essentiels – requests (toujours en version 2.34.2, la dernière documentée en mai 2026 selon WebScraping.ai), bs4, lxml, playwright, httpx, scrapy et pandas – pour couvrir la quasi-totalité des cas d’usage. Notez que l’installation de Playwright nécessite une étape supplémentaire pour télécharger les binaires des navigateurs.

```
# Installation des bibliothèques essentielles de web scraping python
pip install requests==2.32.3
pip install httpx[http2]==0.28.0
pip install beautifulsoup4==4.12.3
pip install lxml==5.3.0
pip install selectolax==0.3.17
pip install scrapy==2.12.0
pip install playwright==1.49.0
pip install curl_cffi==0.7.3
pip install pandas==2.2.0
pip install aiohttp==3.10.0
# Installation des navigateurs Playwright (Chromium, Firefox, WebKit)
playwright install
# Optionnel : Crawl4AI pour le scraping assisté par IA
pip install crawl4ai==0.4.0
# Créer le fichier requirements.txt
pip freeze > requirements.txt
echo "Installation terminée !"
```
**Étape 3 : Première requête HTTP avec httpx**

Nous utiliserons `httpx` plutôt que `requests` dès le début car c’est la bibliothèque recommandée en 2026 : elle supporte nativement HTTP/2, les requêtes asynchrones, et présente une API quasi-identique à `requests`. Voici notre premier scraper basique qui constitue la fondation de tout projet de **scraping python beautifulsoup**.

```
import httpx
from bs4 import BeautifulSoup
import time
# Configuration des headers pour simuler un navigateur réel
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
    "Accept-Language": "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7",
    "Accept-Encoding": "gzip, deflate, br",
    "Connection": "keep-alive",
    "Upgrade-Insecure-Requests": "1",
}
def scraper_page(url: str) -> BeautifulSoup | None:
    """
    Effectue une requête HTTP GET sur l'URL donnée
    et retourne un objet BeautifulSoup pour le parsing.
    Args:
        url: L'URL de la page à scraper
    Returns:
        Objet BeautifulSoup ou None en cas d'erreur
    """
    try:
        # Utiliser un client httpx avec timeout de 30 secondes
        with httpx.Client(headers=HEADERS, timeout=30.0, follow_redirects=True) as client:
            response = client.get(url)
            # Vérifier que la requête a réussi
            response.raise_for_status()
            print(f"Statut HTTP : {response.status_code}")
            print(f"Type de contenu : {response.headers.get('content-type', 'inconnu')}")
            # Parser le HTML avec BeautifulSoup et le backend lxml (plus rapide)
            soup = BeautifulSoup(response.text, "lxml")
            return soup
    except httpx.HTTPStatusError as e:
        print(f"Erreur HTTP {e.response.status_code} pour {url}")
        return None
    except httpx.RequestError as e:
        print(f"Erreur de requête : {e}")
        return None
# Test avec un site de démonstration
if __name__ == "__main__":
    url = "https://books.toscrape.com/"
    soup = scraper_page(url)
    if soup:
        # Extraire le titre de la page
        titre = soup.find("title")
        print(f"Titre de la page : {titre.text if titre else 'Non trouvé'}")
        # Compter le nombre de livres sur la page
        livres = soup.find_all("article", class_="product_pod")
        print(f"Nombre de livres trouvés : {len(livres)}")
    # Pause courtoise entre les requêtes
    time.sleep(1)
```
## Étapes 4-6 : Parser le HTML avec BeautifulSoup et Extraire les Données

Maintenant que nous savons effectuer des requêtes HTTP, passons à l’étape cruciale : l’extraction des données depuis le HTML. **BeautifulSoup** avec le backend `lxml` est la solution idéale pour débuter grâce à son API intuitive et sa robustesse face aux HTMLs mal formés. Selon les estimations de DataResearchTools publiées en mars 2026, BeautifulSoup comptait **2,5 millions** d’utilisateurs actifs mensuels, contre 1,8 million pour Scrapy – les deux bibliothèques Python les plus utilisées pour ce type d’extraction. La bibliothèque continue d’évoluer activement : selon WebScraping.ai, **BeautifulSoup 4.15.0** était la version la plus récente en juin 2026, avec un parsing HTML encore plus tolérant face aux pages mal formées. La documentation complète de BeautifulSoup 4 couvre tous les cas d’usage avancés.

**Étape 4 : Comprendre les sélecteurs CSS avec BeautifulSoup**

BeautifulSoup offre deux approches complémentaires pour cibler des éléments HTML : les méthodes `find()` / `find_all()` qui travaillent avec les attributs HTML directement, et la méthode `select()` qui accepte des sélecteurs CSS – plus expressifs et familiers pour les développeurs web. Maîtriser ces deux approches est essentiel pour tout **scraping python beautifulsoup** sérieux.

**Étape 5 : Extraire et structurer les données d’une page produit**

Voici un exemple complet montrant comment extraire des données structurées depuis une page de liste de produits, en utilisant à la fois `find_all()` et `select_one()` selon les besoins.

```
import httpx
from bs4 import BeautifulSoup
import json
HEADERS = {
    "User-Agent": "Mozilla/5.0 (compatible; MonScraper/1.0; +https://monsite.fr/scraper)",
}
def extraire_livres(url: str) -> list[dict]:
    """
    Extrait les données de livres depuis books.toscrape.com.
    Démontre l'utilisation des sélecteurs CSS et des méthodes BeautifulSoup
    dans le cadre d'un projet d'extraction données web Python.
    """
    livres_extraits = []
    with httpx.Client(headers=HEADERS, timeout=30.0) as client:
        response = client.get(url)
        response.raise_for_status()
        # Parser avec lxml pour de meilleures performances
        soup = BeautifulSoup(response.text, "lxml")
        # Méthode 1 : find_all() avec classe CSS
        articles = soup.find_all("article", class_="product_pod")
        for article in articles:
            # Extraire le titre via l'attribut 'title' de la balise 
            # select_one() retourne le premier élément correspondant au sélecteur CSS
            lien_titre = article.select_one("h3 > a")
            titre = lien_titre["title"] if lien_titre else "Titre inconnu"
            # Extraire le prix avec sélecteur CSS
            prix_element = article.select_one("p.price_color")
            prix = prix_element.text.strip() if prix_element else "Prix inconnu"
            # Extraire la disponibilité
            dispo_element = article.select_one("p.availability")
            disponibilite = dispo_element.text.strip() if dispo_element else "Inconnu"
            # Extraire la note (rating) via l'attribut class
            # Ex: 
```
**Étape 6 : Sauvegarde des données en CSV et JSON**

Une fois les données extraites, il faut les persister dans un format exploitable. Le module `csv` de la bibliothèque standard Python est parfait pour les exports tabulaires, tandis que JSON est idéal pour les données hiérarchiques ou les échanges avec des APIs.

