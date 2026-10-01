---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-10
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmarks"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [1002, 1129]
sha256: bce7da0c8741ca4f445330cedd12be50224bec8631e02af32a21f36fea237ed5
---

# Créer le dossier du projet

                    cle, val = th.text.strip(), td.text.strip()
                    if cle == "UPC":
                        livre.upc = val
                    elif cle == "Number of reviews":
                        try:
                            livre.nb_avis = int(val)
                        except ValueError:
                            pass
        except Exception as e:
            log.debug(f"Enrichissement échoué pour {livre.titre[:40]}... : {e}")
    return livre
async def enrichir_batch(livres: list[Livre], max_concurrent: int = 8) -> list[Livre]:
    """Lance l'enrichissement de tous les livres en parallèle."""
    sem = asyncio.Semaphore(max_concurrent)
    async with httpx.AsyncClient(headers=HEADERS, timeout=20.0) as client:
        tasks = [enrichir_livre(client, l, sem) for l in livres]
        return list(await asyncio.gather(*tasks))
# ============================================================
# COUCHE EXPORT : CSV ET JSON
# ============================================================
def export_csv(livres: list[Livre], path: str) -> None:
    """Exporte les livres en CSV (compatible Excel avec BOM UTF-8)."""
    if not livres:
        return
    cols = list(livres[0].to_dict().keys())
    with open(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=cols)
        w.writeheader()
        w.writerows(l.to_dict() for l in livres)
    log.info(f"CSV exporté : {path} ({len(livres)} lignes, {Path(path).stat().st_size//1024} KB)")
def export_json(livres: list[Livre], path: str, stats: dict) -> None:
    """Exporte les livres en JSON structuré avec métadonnées."""
    payload = {
        "metadata": {**stats, "outil": "scraper-python-2026", "site": BASE_URL},
        "donnees": [l.to_dict() for l in livres]
    }
    with open(path, "w", encoding="utf-8") as f:
        json.dump(payload, f, ensure_ascii=False, indent=2)
    log.info(f"JSON exporté : {path} ({len(livres)} entrées, {Path(path).stat().st_size//1024} KB)")
# ============================================================
# POINT D'ENTRÉE PRINCIPAL
# ============================================================
def main(max_pages: int = 3, max_concurrent: int = 8, format_out: str = "both") -> None:
    """
    Orchestre le scraping complet :
    1. Scraping séquentiel des pages de liste
    2. Enrichissement asynchrone des détails
    3. Export des résultats
    """
    debut = time.time()
    stats = {
        "debut": datetime.now().isoformat(),
        "max_pages": max_pages,
        "pages_scrapees": 0,
        "livres_collectes": 0,
        "erreurs": 0,
    }
    log.info(f"Démarrage – max {max_pages} pages, concurrence {max_concurrent}")
    # Phase 1 : Scraping des pages de liste
    tous_livres: list[Livre] = []
    with httpx.Client(headers=HEADERS, timeout=30.0, follow_redirects=True) as client:
        for batch, num_page in iter_pages(client, f"{BASE_URL}/catalogue/page-1.html", max_pages):
            tous_livres.extend(batch)
            stats["pages_scrapees"] = num_page
            log.info(f"Page {num_page} : {len(batch)} livres (total: {len(tous_livres)})")
    stats["livres_collectes"] = len(tous_livres)
    log.info(f"Phase 1 terminée : {len(tous_livres)} livres en {time.time()-debut:.1f}s")
    # Phase 2 : Enrichissement asynchrone
    log.info("Phase 2 : Enrichissement des détails en mode asynchrone...")
    tous_livres = asyncio.run(enrichir_batch(tous_livres, max_concurrent))
    # Phase 3 : Export
    stats["duree_totale"] = round(time.time() - debut, 2)
    stats["fin"] = datetime.now().isoformat()
    if format_out in ("csv", "both"):
        export_csv(tous_livres, "livres.csv")
    if format_out in ("json", "both"):
        export_json(tous_livres, "livres.json", stats)
    # Résumé final
    prix_valides = [l.prix for l in tous_livres if l.prix > 0]
    print("\n" + "="*55)
    print("  RÉSUMÉ DU SCRAPING – TechInsider.fr")
    print("="*55)
    print(f"  Pages scrapées    : {stats['pages_scrapees']}")
    print(f"  Livres collectés  : {len(tous_livres)}")
    print(f"  Durée totale      : {stats['duree_totale']}s")
    if prix_valides:
        print(f"  Prix moyen        : £{sum(prix_valides)/len(prix_valides):.2f}")
        print(f"  Prix min / max    : £{min(prix_valides):.2f} / £{max(prix_valides):.2f}")
    categories = {l.categorie for l in tous_livres if l.categorie}
    print(f"  Catégories        : {len(categories)} catégories uniques")
    print("="*55)
    print(f"  Fichiers créés : livres.csv et livres.json")
    print("="*55 + "\n")
if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser(
        description="Scraper Python production-ready – TechInsider.fr 2026"
    )
    parser.add_argument("--max-pages", type=int, default=3,
                        help="Nombre max de pages à scraper (défaut: 3)")
    parser.add_argument("--concurrent", type=int, default=8,
                        help="Requêtes asynchrones simultanées (défaut: 8)")
    parser.add_argument("--format", choices=["csv", "json", "both"], default="both",
                        help="Format d'export (défaut: both)")
    args = parser.parse_args()
    main(
        max_pages=args.max_pages,
        max_concurrent=args.concurrent,
        format_out=args.format
    ) ## Performance et Benchmarks : Choisir la Bonne Architecture

Les performances de votre scraper dépendent de l'architecture choisie. Voici des benchmarks réalisés en conditions réelles sur `books.toscrape.com` pour vous guider dans vos choix d'architecture d'**extraction de données web Python**.

| Architecture | Pages/seconde | RAM (100 pages) | CPU Usage | Complexité code | Recommandé pour | 
|---|---|---|---|---|---|
| requests synchrone | 1-2 | ~50 MB | Très faible | Débutant | Tests, prototypes, moins de 50 pages | 
| httpx synchrone | 2-4 | ~45 MB | Faible | Débutant | Projets simples, support HTTP/2 | 
| httpx + asyncio (10 concurrent) | 15-40 | ~80 MB | Moyen | Intermédiaire | Projets moyens, performance requise | 
| httpx + asyncio (50 concurrent) | 50-150 | ~200 MB | Élevé | Avancé | Grands projets, multi-domaines | 
| Scrapy 2.12 (défaut) | 20-80 | ~120 MB | Moyen | Framework | Production, équipes, scalabilité | 
| Playwright headless | 1-3 | ~500 MB | Très élevé | Intermédiaire | Sites JavaScript uniquement | 

Les principales optimisations de performance pour l'**automatisation web Python** sont les suivantes. Premièrement, utilisez `selectolax` (version **0.4.11**, la plus récente en juillet 2026 selon WebScraping.ai) à la place de BeautifulSoup pour le parsing HTML pur – jusqu'à 30x plus rapide pour les pages volumineuses grâce à son moteur de parsing CSS optimisé pour le scraping Python à haut volume, avec une syntaxe similaire. Deuxièmement, bloquez les ressources inutiles dans Playwright (images, CSS, polices) via `page.route()` pour réduire la consommation mémoire de 60 à 70%. Troisièmement, utilisez des sessions persistantes `httpx.Client` plutôt que de créer une nouvelle connexion pour chaque requête – les connexions HTTP keep-alive réduisent la latence de 20 à 40%. Quatrièmement, implémentez un cache local avec `diskcache` ou Redis pour éviter de re-scraper des URLs déjà visitées lors des ré-exécutions ou des tests. Cinquièmement, utilisez le streaming pour les très grandes pages via `client.stream()` afin de ne pas charger l'intégralité du HTML en mémoire avant de commencer le parsing.

## FAQ : Questions Fréquentes sur le Web Scraping Python en 2026

### Le web scraping Python est-il légal en France ?

