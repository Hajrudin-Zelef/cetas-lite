---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-1
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "incident", "mai", "open source"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [1, 32]
sha256: a6aa866950f20108c98e35cb4fe069f508d42873d3298aaf758c99dd91fe1653
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

Un serveur piraté laisse toujours une trace dans les logs, sauf que personne ne les regarde avant qu’il soit trop tard. Le ministère de l’Intérieur a recensé 453 200 cyberattaques en France en 2025, soit une hausse de 87 % en cinq ans selon son rapport annuel sur la cybercriminalité. Face à ce volume, l’ANSSI pousse depuis 2026 les entreprises vers des outils de sécurité open source plutôt que vers des solutions propriétaires coûteuses. Graylog fait partie des rares plateformes capables de centraliser, indexer et analyser des millions de lignes de logs sans les frais de licence d’un Splunk. Ce tutoriel installe Graylog 7.1.8, publiée le 18 août 2026 et toujours la version la plus largement documentée à ce jour, sachant que la page de téléchargement officielle affiche désormais la 7.1.9 comme dernier correctif de la branche depuis sa mise à jour du 2 septembre 2026, sur un serveur Linux avec Docker, puis configure la collecte, le traitement et l’alerte en 11 étapes concrètes.

À la fin de ce guide, vous disposerez d’une pile fonctionnelle capable d’ingérer des dizaines de milliers de messages par seconde, de les structurer via des pipelines, et de déclencher une alerte dès qu’un comportement suspect apparaît. Le tout tient sur un seul serveur avec un investissement de temps d’environ 90 minutes, sans dépendre d’un fournisseur SaaS ni d’un budget de licence annuel.

## Pourquoi centraliser ses logs de sécurité avec Graylog en 2026

Un pare-feu, un serveur web, un annuaire Active Directory et une dizaine de machines Linux génèrent chacun leurs propres journaux, dans des formats différents, stockés sur des disques différents. Sans centralisation, un administrateur doit se connecter à chaque machine pour reconstituer la chronologie d’un incident. C’est précisément le scénario que décrit l’ENISA dans son Threat Landscape 2025 (version 1.2) : l’ingénierie sociale reste le point d’entrée principal des attaquants dans l’écosystème européen, et la détection tardive amplifie les dégâts.

Graylog répond à ce problème en agrégeant tous les flux de logs (syslog, GELF, logs applicatifs, logs Docker, logs Windows via NXLog) dans un moteur de recherche unique basé sur OpenSearch. L’intérêt n’est pas seulement le stockage : c’est la capacité à écrire une seule requête pour retrouver une tentative de connexion SSH échouée sur douze serveurs en même temps, ou à déclencher une alerte automatique dès qu’un compte utilisateur échoue cinq fois de suite. Sur le marché français, l’ANSSI a mis à jour sa politique open source en 2026 autour de quatre axes : publier des logiciels de sécurité sous licence libre, contribuer aux projets externes, renforcer l’écosystème et utiliser des solutions open source en interne. Graylog coche toutes ces cases pour une PME ou une ETI qui ne dispose pas du budget d’un SOC dédié.

Le panorama de la cybermenace publié par le CERT-FR confirme cette tendance : les incidents détectés tardivement, faute de centralisation des journaux, restent parmi les plus coûteux à traiter pour les entités concernées. Les métiers de la cybersécurité recrutent en masse en France en 2026 : ingénieurs cloud security, analystes SOC, responsables conformité NIS2. Tous ont un point commun, ils passent leurs journées dans un outil de gestion de logs. Savoir déployer et exploiter Graylog est donc une compétence directement monétisable, que l’on gère l’infrastructure d’une entreprise ou que l’on prépare une certification en sécurité opérationnelle.

## Graylog 7.1.8 : nouveautés, versions et calendrier de support

La branche 7.1 de Graylog est passée en disponibilité générale le 4 mai 2026, sous le nom de code “Spring 2026 Release” que l’éditeur a présenté comme son premier SIEM “AI-powered”. Elle a ensuite reçu une série de correctifs rapprochés : 7.1.1 le 8 mai, 7.1.3 le 3 juin, 7.1.5 le 8 juillet, 7.1.7 le 5 août, 7.1.8 le 18 août 2026 (la version que ce tutoriel installe), puis 7.1.9, devenue le dernier correctif listé sur la page de téléchargement officielle depuis sa mise à jour du 2 septembre 2026. Graylog a même publié une première bêta de la branche suivante, la 7.2.0-beta.1, dès le 31 août 2026, pour les équipes qui veulent tester en avance les prochaines fonctions. Cette cadence de publication mensuelle montre un projet activement maintenu, ce qui compte quand on choisit un outil de sécurité destiné à tourner plusieurs années en production.

La branche 7.1 introduit deux capacités qui rapprochent Graylog d’un véritable SIEM : les investigations automatisées et la détection native d’anomalies comportementales. Concrètement, l’outil ne se contente plus d’indexer des logs, il peut désormais regrouper automatiquement les événements liés à un même incident et signaler un comportement utilisateur qui s’écarte de la normale, sans règle de corrélation écrite à la main. Ces fonctions s’appuient sur les content packs Illuminate, dont la version 7.0.7 est sortie le 8 juin 2026 avec de nouveaux packs de traitement et des “spotlights” orientés détection.

| Branche | Date de sortie | Dernier correctif | Fin de support | Statut | 
|---|---|---|---|---|
| Graylog 7.1 | 4 mai 2026 | 7.1.8 (18 août 2026) | 4 mai 2027 (support entreprise jusqu’au 4 mai 2028) | Supportée, recommandée | 
| Graylog 7.0 | 3 novembre 2025 | 7.0.12 (5 août 2026) | 3 novembre 2026 | Supportée, en fin de vie proche | 
| Graylog 6.3 | 30 juin 2025 | 6.3.15 (5 août 2026) | 30 juin 2026 | Support terminé | 

Ce tableau justifie à lui seul le choix de la version 7.1.8 pour toute nouvelle installation : la branche 6.3 est déjà en fin de vie, et la 7.0 approche de son terme en novembre 2026. Un point technique à connaître avant de migrer : la matrice de compatibilité officielle exige au minimum MongoDB 7.x (maximum 8.2.x) et OpenSearch 2.0.x (maximum 2.19.5) pour les branches 7.0.x et 7.1.x. Graylog a définitivement abandonné le support d’Elasticsearch au profit d’OpenSearch depuis la version 5.0, un détail qui piège encore des administrateurs qui suivent d’anciens tutoriels.

## Graylog vs Wazuh vs ELK Stack : quel outil de logs choisir

Avant de se lancer dans l’installation, il faut savoir si Graylog est le bon outil pour le besoin. Trois familles de solutions open source dominent le marché de la gestion de logs et de la détection en 2026. Wazuh, en version 4.14.7 publiée le 29 juillet 2026, mise sur des agents installés directement sur chaque machine surveillée, ce qui en fait un outil de détection d’intrusion basée sur l’hôte (HIDS) avant d’être un outil de logs. Elastic, avec Elasticsearch 9.5.1 et Elastic Security 9.5 sorties le 3 août 2026, propose l’écosystème le plus riche mais demande le plus de travail d’intégration pour obtenir une expérience SIEM comparable à un produit clé en main.

Graylog se positionne entre les deux : il est plus léger et moins cher à exploiter qu’un Splunk sur le plan des licences, tout en offrant une interface de sécurité plus aboutie qu’un stack ELK construit à la main. Sa force réside dans le traitement structuré des logs via des pipelines, et dans le data tiering introduit en version 6.0, qui déplace automatiquement les données anciennes vers un stockage moins coûteux tout en gardant tout consultable. Ce mécanisme réduit la facture de stockage d’un ordre de grandeur par rapport à un index unique conservé indéfiniment sur du SSD rapide.

