---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-9
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [792, 1001]
sha256: 3668941abb46ebe4166ef57ae5e93dad1db4470f164684a44993599152b79c91
---

# Créer le dossier du projet

```
#!/usr/bin/env python3
"""
scraper_complet_production.py
Scraper production-ready avec architecture modulaire - TechInsider.fr
Site cible : books.toscrape.com (site de démonstration, scraping autorisé)
Ce projet illustre les meilleures pratiques du web scraping python :
- Architecture en couches séparées (HTTP / Parsing / Storage)
- Dataclasses typées pour les modèles de données
- Scraping de liste synchrone + enrichissement asynchrone
- Logging structuré avec rotation de fichiers
- Export CSV et JSON avec métadonnées
Usage :
    python scraper_complet_production.py --max-pages 5
    python scraper_complet_production.py --max-pages 20 --concurrent 8 --format json
Auteur : TechInsider.fr – 29 mars 2026
"""
import asyncio
import csv
import json
import logging
import random
import time
from dataclasses import dataclass, asdict, field
from datetime import datetime
from pathlib import Path
from typing import Generator
import httpx
from bs4 import BeautifulSoup
# ============================================================
# CONSTANTES DE CONFIGURATION
# ============================================================
BASE_URL = "https://books.toscrape.com"
HEADERS = {
    "User-Agent": (
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
        "AppleWebKit/537.36 (KHTML, like Gecko) "
        "Chrome/131.0.0.0 Safari/537.36"
    ),
    "Accept-Language": "fr-FR,fr;q=0.9,en;q=0.8",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "Sec-Fetch-Dest": "document",
    "Sec-Fetch-Mode": "navigate",
    "Sec-Fetch-Site": "none",
}
NOTES_MAP = {"One": 1, "Two": 2, "Three": 3, "Four": 4, "Five": 5}
# ============================================================
# MODÈLE DE DONNÉES TYPÉ
# ============================================================
@dataclass
class Livre:
    """
    Modèle de données pour un livre scrapé.
    Utilise les dataclasses Python pour la sérialisation automatique.
    """
    titre: str
    prix: float
    note: int  # 1-5 étoiles
    disponibilite: str
    url: str
    categorie: str = ""
    description: str = ""
    upc: str = ""
    nb_avis: int = 0
    date_scraping: str = field(
        default_factory=lambda: datetime.now().strftime("%Y-%m-%dT%H:%M:%S")
    )
    def to_dict(self) -> dict:
        return asdict(self)
    @classmethod
    def depuis_article_html(cls, article: BeautifulSoup) -> "Livre":
        """
        Fabrique un Livre depuis un élément HTML 
```
.
        Toutes les erreurs de parsing sont gérées silencieusement.
        """
        lien = article.select_one("h3 > a")
        prix_el = article.select_one("p.price_color")
        dispo_el = article.select_one("p.availability")
        note_el = article.select_one("p.star-rating")
        # Extraction sécurisée du prix (supprimer symbole monétaire)
        prix_brut = prix_el.text.strip() if prix_el else "0"
        try:
            prix = float(
                prix_brut.replace("£", "").replace("€", "")
                         .replace(",", ".").replace("Â", "").strip()
            )
        except ValueError:
            prix = 0.0
        # Extraction de la note étoiles depuis la classe CSS
        classes_note = note_el.get("class", []) if note_el else []
        note_texte = classes_note[1] if len(classes_note) > 1 else "Zero"
        # Construction de l'URL absolue
        href = (lien["href"] if lien else "").replace("../", "")
        url_complete = f"{BASE_URL}/catalogue/{href}"
        return cls(
            titre=lien["title"].strip() if lien and lien.get("title") else "Titre inconnu",
            prix=prix,
            note=NOTES_MAP.get(note_texte, 0),
            disponibilite=dispo_el.text.strip() if dispo_el else "",
            url=url_complete,
        )
# ============================================================
# CONFIGURATION DU LOGGING
# ============================================================
def setup_logging() -> logging.Logger:
    """Configure le système de logging avec sortie console et fichier."""
    logger = logging.getLogger("scraper")
    logger.setLevel(logging.INFO)
    # Format avec timestamp, niveau et message
    fmt = logging.Formatter(
        "%(asctime)s [%(levelname)-8s] %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S"
    )
    # Handler console
    ch = logging.StreamHandler()
    ch.setFormatter(fmt)
    logger.addHandler(ch)
    # Handler fichier
    fh = logging.FileHandler("scraper.log", encoding="utf-8")
    fh.setFormatter(fmt)
    logger.addHandler(fh)
    return logger
log = setup_logging()
# ============================================================
# COUCHE HTTP : REQUÊTES ET PAGINATION
# ============================================================
def get_page(client: httpx.Client, url: str, max_retries: int = 3) -> BeautifulSoup | None:
    """
    Récupère et parse une page avec retry automatique.
    Couche HTTP isolée du reste de la logique.
    """
    for tentative in range(1, max_retries + 1):
        try:
            resp = client.get(url)
            resp.raise_for_status()
            return BeautifulSoup(resp.text, "lxml")
        except httpx.HTTPStatusError as e:
            code = e.response.status_code
            if code in (400, 403, 404):
                log.warning(f"Erreur permanente {code} pour {url}")
                return None
            delai = 2 ** tentative + random.uniform(0, 1)
            log.warning(f"HTTP {code} – tentative {tentative}/{max_retries}, attente {delai:.1f}s")
            time.sleep(delai)
        except httpx.RequestError as e:
            delai = 2 ** tentative
            log.warning(f"Erreur réseau ({e}) – tentative {tentative}/{max_retries}")
            time.sleep(delai)
    log.error(f"Impossible d'accéder à {url} après {max_retries} tentatives")
    return None
def iter_pages(
    client: httpx.Client,
    url_start: str,
    max_pages: int = 50
) -> Generator[tuple[list[Livre], int], None, None]:
    """
    Générateur qui itère sur toutes les pages de liste.
    Yields: (liste de Livres, numéro de page)
    """
    url = url_start
    page_num = 0
    while url and page_num < max_pages:
        soup = get_page(client, url)
        if not soup:
            break
        articles = soup.find_all("article", class_="product_pod")
        livres = [Livre.depuis_article_html(a) for a in articles]
        page_num += 1
        yield livres, page_num
        # Trouver la page suivante
        next_link = soup.select_one("li.next > a")
        if next_link:
            href = next_link["href"]
            if "catalogue/" not in url:
                url = f"{BASE_URL}/catalogue/{href}"
            else:
                url = f"{url.rsplit('/', 1)[0]}/{href}"
        else:
            url = None
        if url:
            time.sleep(random.uniform(0.5, 1.5))
# ============================================================
# COUCHE ENRICHISSEMENT : SCRAPING ASYNCHRONE DES DÉTAILS
# ============================================================
async def enrichir_livre(
    client: httpx.AsyncClient,
    livre: Livre,
    sem: asyncio.Semaphore
) -> Livre:
    """Enrichit un livre avec ses données détaillées (async)."""
    async with sem:
        try:
            resp = await client.get(livre.url)
            resp.raise_for_status()
            soup = BeautifulSoup(resp.text, "lxml")
            # Description
            desc = soup.select_one("#product_description + p")
            if desc:
                livre.description = desc.text.strip()[:300]
            # Catégorie depuis le fil d'Ariane
            crumbs = soup.select("ul.breadcrumb li")
            if len(crumbs) >= 3:
                livre.categorie = crumbs[-2].text.strip()
            # UPC et nombre d'avis depuis le tableau produit
            for row in soup.select("table.table-striped tr"):
                th = row.find("th")
                td = row.find("td")
                if th and td:
