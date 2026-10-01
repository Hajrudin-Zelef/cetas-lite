---
id: collect-261001-general-networking/general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes-1
title: "Stockez le mot de passe dans un fichier .env hors du repo Git"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["acquisition", "aws", "embeddings", "pricing"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes.md
source_anchor: ""
source_lines: [1, 59]
sha256: b58e152962adf72a2dd761b59bee892896fe3977253c7f8fd637186e380fbfc8
---

# Stockez le mot de passe dans un fichier .env hors du repo Git

**MongoDB Atlas** a passé la barre des **65 200 clients** au quatrième trimestre de l’exercice 2026 (clos le 31 janvier 2026), avec un chiffre d’affaires trimestriel de **695,1 millions de dollars**, en hausse de 27 % sur un an, selon les résultats publiés par MongoDB Inc. La part Atlas pèse désormais 72 % du revenu, soit un rythme annualisé supérieur à **2 milliards de dollars**. Pour les développeurs français qui démarrent un projet en avril 2026, ce tutoriel MongoDB Atlas vous guide pas à pas : création du cluster gratuit M0, connexion Node.js et Python avec PyMongo 4.17, indexation, agrégation, Atlas Search et Vector Search alimenté par Voyage AI. À la fin, vous aurez une API Express opérationnelle et un cluster prêt pour la production.

Publié le 29 avril 2026.

## Pourquoi un tutoriel MongoDB Atlas en avril 2026

Le marché des bases de données NoSQL n’a jamais été aussi tendu. **MongoDB 8.2**, sortie en septembre 2025 sur la branche Rapid Release, et **MongoDB 8.0** (octobre 2024, support de sécurité jusqu’au 31 octobre 2029) constituent les versions actives du moteur en avril 2026. La version 6.0, longtemps cluster par défaut sur Atlas, est sortie de support le 31 juillet 2025 – toute migration tardive expose à des CVE non corrigés. Ce tutoriel cible MongoDB 8.0 LTS, le driver **PyMongo 4.17.0** et le driver Node.js **mongodb 7.2.0**, qui constituent le combo le plus stable pour démarrer un projet en production en France ou en Europe.

L’autre nouveauté majeure de ces douze derniers mois : l’acquisition de **Voyage AI** par MongoDB, finalisée début 2025, a fait basculer Atlas dans la cour des plateformes RAG natives. Atlas Vector Search ne se contente plus de stocker des embeddings – la plateforme génère désormais les vecteurs côté serveur via les modèles voyage-3.5 et voyage-3.5-lite, supprimant un appel réseau pour chaque ingestion. Couplé aux *lexical prefilters* arrivés en février 2026, cela accélère considérablement les requêtes hybrides. Pour un développeur qui veut bâtir un chatbot RAG, c’est un raccourci décisif.

Ce tutoriel se découpe en treize étapes : prérequis, création du compte Atlas, déploiement du cluster M0 gratuit, configuration du réseau et des utilisateurs, connexion avec MongoDB Compass, initialisation d’un projet Node.js / Express, opérations CRUD, conception du schéma, index, agrégation, Atlas Search, Vector Search, et déploiement production. Chaque étape comporte un bloc de code reproductible, les pièges courants, et les sorties attendues. Vous trouverez en fin d’article un projet complet (API REST de catalogue produits) et une section dépannage avec huit erreurs classiques.

## Prérequis avec versions exactes

Avant de toucher la console Atlas, votre poste doit disposer de l’environnement suivant. Ces versions ont été validées en avril 2026 sur Ubuntu 24.04 LTS, macOS Sonoma 14.6 et Windows 11 23H2. Toute version inférieure à celles indiquées peut entraîner des erreurs de TLS 1.3 ou des incompatibilités avec le driver mongodb 7.x.

| Outil | Version minimale | Recommandée (avril 2026) | Utilité | 
|---|---|---|---|
| Node.js | 20.x LTS | 22.11 LTS | Runtime serveur de l’application | 
| npm | 10.x | 10.9 | Gestion des dépendances | 
| Python | 3.10 | 3.12.7 | Scripts d’ingestion et PyMongo | 
| PyMongo | 4.13 | 4.17.0 | Driver Python officiel | 
| mongodb (driver Node) | 7.0 | 7.2.0 | Driver Node.js officiel | 
| MongoDB Compass | 1.42 | 1.46 | GUI d’administration | 
| mongosh | 2.3 | 2.5 | Shell interactif | 
| MongoDB Atlas CLI | 1.32 | 1.40 | Provisioning en ligne de commande | 
| Git | 2.40 | 2.47 | Versioning du projet | 

Vous aurez également besoin d’une **adresse e-mail valide** pour créer le compte Atlas (le tier M0 reste gratuit à vie sans carte bancaire), d’un **terminal** avec accès sortant sur le port 27017 – vérifiez avec votre administrateur réseau si vous travaillez derrière un proxy d’entreprise – et de **10 Go d’espace disque** libre pour Compass et les dumps de test. Comptez environ 2 heures pour parcourir intégralement les treize étapes en exécutant chaque commande.

## Étape 1 – Créer un compte MongoDB Atlas

Rendez-vous sur la page Atlas Database et cliquez sur « Try Free ». L’inscription accepte une connexion fédérée Google ou GitHub, ce qui simplifie la gestion ultérieure des accès en équipe. Sélectionnez la région **Europe** dans le menu de localisation des données par défaut, conformément aux exigences RGPD pour les charges de travail françaises. Atlas vous demandera de créer une *organisation* et un *projet* : nommez-les respectivement « tech-insider-org » et « catalog-prod » pour suivre les exemples du tutoriel.

Atlas vous propose ensuite un questionnaire d’orientation (langage privilégié, expérience MongoDB). Sélectionnez *JavaScript* et *Python*, puis *Beginner* – cela personnalise les snippets de connexion affichés sur la console. Une fois la page d’accueil chargée, cliquez sur « Build a Database » au centre du tableau de bord. C’est ici que vous choisirez le tier de votre premier cluster.

**Piège fréquent :** ne validez pas l’écran de carte bancaire si vous comptez rester sur le tier gratuit. Atlas affiche par défaut une étape de saisie de moyen de paiement pour préparer la migration vers un tier dédié, mais elle est facultative. Skippez-la – vous pourrez toujours l’ajouter via Billing > Payment Methods plus tard.

## Étape 2 – Déployer un cluster M0 gratuit

Sur l’écran « Deploy your database », sélectionnez le plan **M0 (Free)**. Ce tier offre **512 Mo de stockage**, RAM partagée et vCPU partagé, suffisant pour développer et faire tourner un POC. Choisissez ensuite le fournisseur cloud : **AWS** reste le plus performant en latence depuis la France métropolitaine si vous prenez la région *eu-west-3 (Paris)*. Vous pouvez aussi choisir Azure (*francecentral*) ou Google Cloud (*europe-west9*, Paris). Atlas couvre désormais plus de 125 régions sur les trois hyperscalers, selon la documentation officielle Atlas.

Nommez le cluster `catalog-cluster`. Atlas verrouille automatiquement la version MongoDB à **8.0** pour les nouveaux clusters M0 depuis la mise à jour de mars 2026. Cliquez sur « Create Deployment ». Le provisioning prend entre 3 et 7 minutes – vous pouvez profiter de ce délai pour configurer l’accès réseau et l’utilisateur initial dans les étapes suivantes.

| Tier | Stockage | RAM | vCPU | Coût indicatif | Usage cible | 
|---|---|---|---|---|---|
| M0 Free | 512 Mo | partagée | partagé | 0 $ | POC, apprentissage | 
| M2 Flex | 2 Go | partagée | partagé | ~9 $/mois | Petit prod léger | 
| M5 Flex | 5 Go | partagée | partagé | ~25 $/mois | Pré-production | 
| M10 | 10–128 Go | 2 Go | 2 | ~0,08 $/heure (≈58 $/mois) | Production légère | 
| M20 | 20–256 Go | 4 Go | 2 | ~0,20 $/heure (≈146 $/mois) | Production standard | 
| M30 | 40–512 Go | 8 Go | 2 | variable, voir page pricing | Production exigeante | 

**Limitation M0 :** 100 opérations CRUD par seconde maximum, 500 connexions concurrentes, pas de sauvegardes Cloud Backup ni de sharding. Ces plafonds ne posent pas problème en développement, mais dimensionnent votre passage en M10 dès que la charge réelle approche les 50 ops/s sur des heures de pointe.

## Étape 3 – Configurer l’accès réseau et l’utilisateur

Pendant que le cluster se provisionne, allez dans **Security > Network Access**. Cliquez sur « Add IP Address ». Pour le développement local, sélectionnez « Add Current IP Address ». **N’utilisez jamais 0.0.0.0/0 en production** – cela ouvre votre cluster à tout l’Internet. En entreprise, demandez à votre équipe réseau la plage CIDR de votre VPN sortant et déclarez-la nominativement.

