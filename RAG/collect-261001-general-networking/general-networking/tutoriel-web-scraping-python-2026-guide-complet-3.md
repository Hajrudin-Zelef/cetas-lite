---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-3
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [203, 365]
sha256: 7b3f08e13a76de98bd29582334840a619931824da27602fbcb6bca90badd227a
---

# Créer le dossier du projet

```
import csv
import json
from pathlib import Path
from datetime import datetime
def sauvegarder_csv(donnees: list[dict], nom_fichier: str = "livres.csv") -> None:
    """
    Sauvegarde une liste de dictionnaires dans un fichier CSV.
    Gère automatiquement l'encodage UTF-8 et les caractères spéciaux français.
    """
    if not donnees:
        print("Aucune donnée à sauvegarder.")
        return
    chemin = Path(nom_fichier)
    # Extraire les noms de colonnes depuis le premier enregistrement
    colonnes = list(donnees[0].keys())
    with open(chemin, "w", newline="", encoding="utf-8-sig") as f:
        # utf-8-sig ajoute le BOM pour la compatibilité Excel
        writer = csv.DictWriter(f, fieldnames=colonnes)
        writer.writeheader()
        writer.writerows(donnees)
    print(f"{len(donnees)} enregistrements sauvegardés dans {chemin}")
def sauvegarder_json(donnees: list[dict], nom_fichier: str = "livres.json") -> None:
    """
    Sauvegarde les données en JSON avec métadonnées de scraping.
    """
    export = {
        "metadata": {
            "date_scraping": datetime.now().isoformat(),
            "nombre_enregistrements": len(donnees),
            "version_scraper": "1.0.0"
        },
        "donnees": donnees
    }
    chemin = Path(nom_fichier)
    with open(chemin, "w", encoding="utf-8") as f:
        json.dump(export, f, ensure_ascii=False, indent=2)
    print(f"Données sauvegardées dans {chemin} ({chemin.stat().st_size / 1024:.1f} KB)")
def charger_json(nom_fichier: str) -> list[dict]:
    """Charge les données depuis un fichier JSON précédemment exporté."""
    chemin = Path(nom_fichier)
    if not chemin.exists():
        raise FileNotFoundError(f"Fichier non trouvé : {chemin}")
    with open(chemin, "r", encoding="utf-8") as f:
        export = json.load(f)
    return export.get("donnees", [])
# Utilisation
if __name__ == "__main__":
    donnees_exemple = [
        {"titre": "A Light in the Attic", "prix": "£51.77", "note": "Three", "disponibilite": "In stock"},
        {"titre": "Tipping the Velvet", "prix": "£53.74", "note": "One", "disponibilite": "In stock"},
        {"titre": "Soumission", "prix": "£50.10", "note": "One", "disponibilite": "In stock"},
    ]
    sauvegarder_csv(donnees_exemple, "livres_export.csv")
    sauvegarder_json(donnees_exemple, "livres_export.json")
    donnees_rechargees = charger_json("livres_export.json")
    print(f"{len(donnees_rechargees)} enregistrements rechargés depuis JSON")
```
## Étapes 7-8 : Gérer la Pagination et Scraper Plusieurs Pages

Dans la réalité, les données que vous souhaitez collecter sont rarement concentrées sur une seule page. La gestion de la pagination est une compétence fondamentale du **web scraping Python** : elle vous permet de traverser automatiquement des dizaines, voire des centaines de pages de résultats.

**Étape 7 : Implémenter la gestion de la pagination**

Il existe plusieurs patterns de pagination sur le web : pagination numérique classique (/page/1, /page/2…), bouton “charger plus” (lazy loading), défilement infini (infinite scroll), et pagination par curseur (utilisée par de nombreuses APIs). La gestion robuste de la pagination est une marque des scrapers professionnels.

```
import httpx
from bs4 import BeautifulSoup
import time
import random
from typing import Generator
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
}
def extraire_livres_page(soup: BeautifulSoup) -> list[dict]:
    """Extrait tous les livres d'une page BeautifulSoup."""
    livres = []
    for article in soup.find_all("article", class_="product_pod"):
        lien = article.select_one("h3 > a")
        prix = article.select_one("p.price_color")
        livres.append({
            "titre": lien["title"] if lien else "",
            "prix": prix.text.strip() if prix else "",
        })
    return livres
def trouver_page_suivante(soup: BeautifulSoup, url_base: str) -> str | None:
    """
    Cherche le lien vers la page suivante dans la navigation.
    Retourne l'URL complète ou None si on est sur la dernière page.
    """
    bouton_suivant = soup.select_one("li.next > a")
    if bouton_suivant:
        href = bouton_suivant.get("href", "")
        if href.startswith("http"):
            return href
        elif href.startswith("catalogue/"):
            return f"https://books.toscrape.com/{href}"
        else:
            return f"{url_base.rstrip('/')}/{href}"
    return None
def scraper_toutes_pages(
    url_depart: str,
    max_pages: int = 50,
    delai_min: float = 0.5,
    delai_max: float = 2.0
) -> Generator[dict, None, None]:
    """
    Scraper générique avec gestion automatique de la pagination.
    Utilise un générateur pour traiter les données au fur et à mesure
    sans tout charger en mémoire – idéal pour les grands datasets
    dans le cadre de l'automatisation web Python.
    Args:
        url_depart: URL de la première page
        max_pages: Nombre maximum de pages à scraper (sécurité)
        delai_min: Délai minimum entre les requêtes (secondes)
        delai_max: Délai maximum entre les requêtes (secondes)
    """
    url_courante = url_depart
    pages_scrapees = 0
    total_items = 0
    with httpx.Client(headers=HEADERS, timeout=30.0, follow_redirects=True) as client:
        while url_courante and pages_scrapees < max_pages:
            print(f"Scraping page {pages_scrapees + 1} : {url_courante}")
            try:
                response = client.get(url_courante)
                response.raise_for_status()
                soup = BeautifulSoup(response.text, "lxml")
                livres = extraire_livres_page(soup)
                total_items += len(livres)
                # Retourner chaque livre via le générateur
                yield from livres
                # Trouver la page suivante
                url_courante = trouver_page_suivante(soup, url_courante)
                pages_scrapees += 1
                # Délai aléatoire pour paraître plus humain
                if url_courante:
                    delai = random.uniform(delai_min, delai_max)
                    print(f"  Pause de {delai:.1f}s... ({len(livres)} items extraits)")
                    time.sleep(delai)
            except httpx.HTTPStatusError as e:
                print(f"  Erreur HTTP {e.response.status_code} - Arrêt du scraping")
                break
    print(f"\nScraping terminé : {pages_scrapees} pages, {total_items} items au total")
# Utilisation du générateur pour scraper toutes les pages
if __name__ == "__main__":
    tous_les_livres = []
    for livre in scraper_toutes_pages("https://books.toscrape.com/", max_pages=5):
        tous_les_livres.append(livre)
    print(f"Total livres collectés : {len(tous_les_livres)}")
    # Sauvegarder en CSV
    import csv
    with open("tous_livres.csv", "w", newline="", encoding="utf-8-sig") as f:
        writer = csv.DictWriter(f, fieldnames=["titre", "prix"])
        writer.writeheader()
        writer.writerows(tous_les_livres)
```
**Étape 8 : Respecter le fichier robots.txt et les limites de taux**

Un scraper responsable vérifie toujours le fichier `robots.txt` avant de commencer. Le module `urllib.robotparser` de la bibliothèque standard Python permet de vérifier programmatiquement si une URL est autorisée à être crawlée. Ne jamais ignorer le `Crawl-delay` indiqué dans ce fichier – c'est à la fois une obligation éthique et une protection contre le bannissement IP.

## Étapes 9-10 : Scraper les Sites JavaScript avec Playwright

