---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-8
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [780, 791]
sha256: 9152562f9c495cd34b70888d9a89c50c083a7c2add600eba418ac5ba192bfa60
---

# Créer le dossier du projet

**Scrapy**, avec ses 54 800+ étoiles GitHub, est le framework de référence pour les projets de **scraping python** à grande échelle. Contrairement aux approches bibliothèque que nous avons vues jusqu'ici, Scrapy est un framework complet avec son propre moteur asynchrone (basé sur Twisted), un système de middleware, des pipelines de traitement des données, et une architecture de spiders réutilisables. Le projet reste très actif : la version **Scrapy 2.18.0**, publiée le 20 août 2026, a ajouté de nouvelles fonctionnalités prêtes pour la production, selon l'équipe du projet Scrapy elle-même. La documentation officielle de Scrapy est particulièrement exhaustive.

Les versions récentes de Scrapy ont considérablement fait évoluer le framework. **Scrapy 2.15.0**, sorti le 9 avril 2026, a introduit un mode sans reactor Twisted ainsi qu'un nouveau download handler basé sur httpx. **Scrapy 2.16.0**, publié le 19 mai 2026, a ajouté le support officiel de Python 3.14 et la compatibilité avec Twisted 26.4.0 et supérieur. Puis **Scrapy 2.17.0**, livré le 7 juillet 2026, a apporté le support HTTP/2 et des proxies SOCKS pour le scraping à grande échelle, selon l'équipe Scrapy – en plus de la meilleure intégration existante avec Playwright via `scrapy-playwright` pour les sites JavaScript. La commande `scrapy bench` intégrée permet toujours de mesurer précisément les performances de votre setup, désormais enrichi de ces nouvelles capacités.

L'architecture de Scrapy repose sur le concept de **Spider** : une classe Python qui définit comment naviguer sur un site et quoi extraire. Créez votre spider avec `scrapy genspider mon_spider example.com`, puis implémentez la méthode `parse()` pour extraire les données et les `Item`s pour les structurer. Les données transitent par des **Pipelines** (nettoyage, validation, persistance) et des **Middlewares** gèrent les aspects transversaux : rotation de proxies, gestion des cookies, respect du robots.txt, et retries automatiques.

Scrapy est particulièrement adapté aux cas suivants : projets nécessitant de scraper des milliers à millions de pages, équipes où plusieurs développeurs collaborent sur le même scraper, pipelines de nettoyage et transformation des données complexes, projets nécessitant monitoring et scheduling via Scrapy Cloud ou Scrapyd, et tous les cas où la scalabilité et la maintenabilité à long terme sont prioritaires.

## Projet Complet : Scraper de Données Production-Ready

Mettons en pratique tout ce que nous avons appris en construisant un scraper complet, robuste et prêt pour la production. Ce projet illustre les meilleures pratiques de l'**extraction données web Python** : architecture modulaire, gestion d'erreurs, logging structuré, persistance des données et configuration externalisée. Il servira de base solide pour vos propres projets.

