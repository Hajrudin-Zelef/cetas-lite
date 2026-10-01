---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-1
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "agents", "incident", "mai", "open source", "valuation"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: ee214d2f191b3fb78d66748f93ac42addc39bfe602c275e1a5ff2c1b94f10e64
---

# Redémarre l'agent et relance les scans SCA

*Mis à jour le 16 août 2026.* Pendant que les licences Splunk et la facturation au gigaoctet de Microsoft Sentinel explosent les budgets, une alternative open source s’est imposée dans les SOC européens : **Wazuh**. Ce SIEM/XDR gratuit sous licence GPLv2 cumule plus de 16 000 étoiles sur GitHub et couvre, sans surcoût de licence, la détection d’intrusion, la surveillance d’intégrité des fichiers, la détection de vulnérabilités et la réponse automatisée aux incidents. Dans ce tutoriel, vous allez déployer **Wazuh 4.14.7** (version stable publiée le 29 juillet 2026, listée comme dernière version dans la documentation officielle) de A à Z, enrôler un agent, détecter une vraie attaque et automatiser sa neutralisation – en 12 étapes, environ 40 minutes. Les plus curieux pourront aussi garder un œil sur la future branche majeure, déjà accessible sous forme de préversion **Wazuh 5.0.0-beta2** publiée en mai 2026.

Au-delà de la technique, l’enjeu est réglementaire. Avec la transposition de la directive NIS 2 qui vise près de 15 000 entités françaises et impose une notification d’incident sous 24 puis 72 heures, disposer d’un système capable de *détecter, journaliser et corréler* les événements de sécurité n’est plus optionnel. Un SIEM open source auto-hébergé comme Wazuh répond à cette obligation tout en gardant les journaux sur votre infrastructure – un atout de souveraineté et de conformité au RGPD que les plateformes cloud américaines peinent à offrir.

## Pourquoi Wazuh, le SIEM open source, séduit l’Europe en 2026

La vague de cyberattaques de 2025-2026 a rappelé une évidence : sans visibilité sur ses journaux, une organisation est aveugle. La fuite de données de l’ANTS ou l’exfiltration de 350 Go subie par la Commission européenne partagent un point commun : la détection tardive. Un SIEM (Security Information and Event Management) centralise et analyse en temps réel les journaux de tous vos serveurs, postes et équipements réseau pour repérer les comportements anormaux avant qu’ils ne deviennent une brèche.

Là où Splunk facture à l’ingestion et Microsoft Sentinel au gigaoctet ingéré dans Azure, Wazuh reste **entièrement gratuit** et auto-hébergé – même la propre offre cloud gérée de son éditeur illustre l’écart de coût : selon le comparateur MSP Compared, Wazuh Cloud facturait en juin 2026 à partir de 571 $/mois pour 100 agents, jusqu’à 923 $/mois pour 250 agents et 1 467 $/mois pour 500 agents, quand la version auto-hébergée de ce tutoriel reste, elle, à coût de licence nul quel que soit le nombre d’agents. Développé par la société américaine Wazuh, Inc. mais distribué sous licence libre GPLv2, le projet affichait déjà plus de 47 900 commits sur son dépôt GitHub principal en juin 2026, preuve d’un rythme de développement soutenu, et il combine dans un seul produit ce que la concurrence vend en modules séparés : SIEM, XDR (détection et réponse étendues), surveillance d’intégrité des fichiers (FIM), évaluation de configuration (SCA), détection de vulnérabilités et réponse active. Pour une PME, une collectivité ou un hébergeur français soumis à NIS 2, l’équation est simple : la souveraineté des données, un coût de licence nul et une conformité facilitée. Cela s’inscrit dans la logique de la stratégie nationale de cybersécurité 2026-2030 qui met l’accent sur l’autonomie technologique.

Ce choix de l’open source rejoint une tendance de fond : comme pour l’auto-hébergement d’un gestionnaire de mots de passe avec Vaultwarden ou le blocage d’attaques avec CrowdSec, reprendre le contrôle de sa sécurité passe par des briques libres, éprouvées et maîtrisées de bout en bout. Wazuh en est la pièce maîtresse côté détection.

## Wazuh, c’est quoi ? L’architecture SIEM/XDR décryptée

Avant d’installer Wazuh, il faut comprendre ses quatre composants. Une installation « tout-en-un » les regroupe sur un seul serveur, mais en production, l’indexeur et le tableau de bord peuvent être répartis sur des machines dédiées pour absorber la montée en charge.

### Les quatre briques de Wazuh

- **Wazuh Manager (serveur)** : le cœur du système. Il reçoit les données des agents, applique les règles de détection, corrèle les événements, déclenche les alertes et pilote la réponse active.
- **Wazuh Indexer** : moteur de recherche et de stockage basé sur OpenSearch. Il indexe les alertes et les journaux pour permettre des recherches quasi instantanées, même sur des téraoctets de données.
- **Wazuh Dashboard** : l’interface web (basée sur OpenSearch Dashboards). Elle offre tableaux de bord, visualisations, exploration des alertes, cartographie MITRE ATT&CK et rapports de conformité.
- **Wazuh Agent** : un collecteur léger (environ 35 Mo de RAM en moyenne) installé sur chaque machine surveillée. Il remonte journaux, événements FIM, inventaire et résultats d’analyses vers le manager.

La communication agent-manager est chiffrée et s’appuie sur des clés d’enregistrement uniques. Un agent consomme très peu de ressources, ce qui permet de couvrir un parc de milliers de machines depuis un seul serveur Wazuh correctement dimensionné. Le tableau ci-dessous positionne Wazuh face aux principaux SIEM du marché en 2026.

| Solution | Type | Licence / coût | Hébergement | Points forts | 
|---|---|---|---|---|
| **Wazuh 4.14.7** | SIEM + XDR open source | GPLv2, gratuit | Auto-hébergé (souverain) | FIM, SCA, réponse active, MITRE, sans coût de licence | 
| Splunk Enterprise Security | SIEM commercial | Licence à l’ingestion / au workload | Cloud ou on-premise | Écosystème d’apps, langage SPL, maturité | 
| Microsoft Sentinel | SIEM cloud natif | Facturation au Go ingéré (Azure) | Cloud Azure | Intégration Microsoft 365 / Entra ID | 
| Elastic Security | SIEM / analytics | Abonnement (offre gratuite limitée) | Cloud ou auto-hébergé | Recherche, machine learning | 
| IBM QRadar | SIEM commercial | Licence (EPS / volume) | Cloud ou on-premise | Corrélation avancée, grands comptes | 

La documentation officielle détaille chaque composant sur la page Getting started. Retenez l’essentiel : Wazuh offre en open source ce que les autres facturent, avec une architecture modulaire qui grandit avec vos besoins.

## Ce que vous allez construire dans ce tutoriel Wazuh

Ce guide ne se limite pas à une installation. À la fin des 12 étapes, vous disposerez d’un **projet complet et fonctionnel** : un serveur Wazuh tout-en-un supervisant un serveur web Linux, avec surveillance d’intégrité des fichiers sensibles, détection des vulnérabilités logicielles, contrôle de configuration CIS, et surtout une chaîne de détection-réponse capable de repérer une attaque SSH par force brute et de **bannir automatiquement l’adresse IP de l’attaquant** en quelques secondes. Nous cartographierons enfin les alertes sur le référentiel MITRE ATT&CK et alignerons le déploiement sur les exigences NIS 2 et RGPD.

- **Machine 1** – le serveur Wazuh (manager + indexeur + tableau de bord), Ubuntu 24.04 LTS, IP d’exemple`10.0.0.10` .
- **Machine 2** – un serveur web à surveiller (agent Wazuh), Ubuntu 24.04 LTS, IP d’exemple`10.0.0.21` , nom`web01` .
- **Résultat** – un SIEM opérationnel, des alertes en temps réel, une réponse active fonctionnelle et un socle de conformité.

Toutes les adresses IP utilisées (plages `203.0.113.0/24` et `198.51.100.0/24`) sont des blocs de documentation réservés : remplacez-les par vos valeurs réelles. Passons aux prérequis.

## Prérequis : versions, matériel et ports réseau

