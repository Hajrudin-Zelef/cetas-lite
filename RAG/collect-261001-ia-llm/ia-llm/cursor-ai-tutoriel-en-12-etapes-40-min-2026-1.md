---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026-1
title: "Windows (winget)"
domain: ia-llm
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft", "Nvidia"]
dates: []
keywords: ["agent", "agents", "copilot", "mai", "mcp", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [1, 40]
sha256: 97f32310e46e486235eaf675ed5596043af3b2867e564faa40f04809afc1dde1
---

# Windows (winget)

**Cursor** s’est imposé en moins de trois ans comme l’éditeur de code doté d’intelligence artificielle le plus discuté de la planète. Édité par la start-up américaine **Anysphere**, il a atteint une valorisation de 9,9 milliards de dollars dès juin 2025 à la faveur d’une Série C de 900 millions de dollars (selon TechCrunch et Unite.AI), franchi la barre des 2 milliards de dollars de revenus annualisés début 2026, bouclé une Série D de 2,3 milliards de dollars le 13 novembre 2025 le valorisant à 29,3 milliards, puis engagé, selon CNBC en avril 2026, des discussions pour lever 2 milliards supplémentaires à plus de 50 milliards de dollars. Il revendique aujourd’hui environ un million d’utilisateurs, dont près de 300 000 payants fin 2025, qui produisent près d’un milliard de lignes de code chaque jour ; signe que l’élan dépasse désormais le poste de travail, l’application Cursor pour iOS est passée en bêta publique fin juin 2026 avec la version 3.9, une annonce qui a cumulé 3,9 millions de vues selon Explainx. Ce tutoriel **Cursor AI** vous accompagne, en 12 étapes et environ 40 minutes, de l’installation jusqu’à la livraison d’un projet complet.

Contrairement à une simple extension, **Cursor AI** est un environnement de développement (IDE) à part entière, dérivé de Visual Studio Code, mais repensé autour de l’IA. Depuis la sortie de **Cursor 2.0** le 29 octobre 2025, l’éditeur embarque son propre modèle maison, *Composer*, et une interface multi-agents capable d’orchestrer jusqu’à huit agents en parallèle ; la troisième génération majeure, **Cursor 3.0**, est arrivée le 2 avril 2026, suivie le 24 avril par la version 3.2 qui ajoute la commande `/multitask` et la prise en charge des espaces de travail multi-racines (multi-root workspaces), puis le 29 mai 2026 par la version 3.6, qui introduit un « Auto-review Run Mode » selon Rapid Dev ; le rythme des sorties ne faiblit pas puisqu’en juin 2026, la version 3.7 a introduit les agents cloud et des sous-agents `/in-cloud` selon Anycap.ai, suivie le 18 juin par la 3.8 et son nouveau skill `/automate` (Pondero.ai), puis le 22 juin par la 3.9, qui ajoute une page Customize et une Marketplace de règles et d’extensions (Anycap.ai). Nous allons exploiter toutes ces briques – Tab, édition en ligne, Chat, Agent, règles de projet, MCP et CLI – pour construire une véritable API de gestion de tâches.

Ce guide s’adresse aux développeurs de tous niveaux, du débutant curieux au professionnel qui veut industrialiser son usage de l’IA. Voici ce que vous saurez faire à la fin :

- Installer et configurer **Cursor** sur Windows, macOS ou Linux, et importer votre environnement VS Code.
- Utiliser les quatre modes clés (Tab, édition en ligne, Chat, Agent) et choisir le bon modèle selon la tâche.
- Cadrer l’IA avec des règles de projet, générer un code multi-fichiers et laisser l’agent tester et corriger seul.
- Étendre l’éditeur avec MCP, l’automatiser via Cursor CLI et sécuriser vos données en entreprise.

## Cursor en 2026 : l’éditeur IA à 9,9 milliards de dollars

**Cursor** est un « fork » de VS Code : vous retrouvez la même ergonomie, les mêmes raccourcis et la compatibilité avec la quasi-totalité des extensions, thèmes et paramètres de l’éditeur de Microsoft. La différence tient à la couche d’IA, profondément intégrée plutôt que greffée. Là où GitHub Copilot reste une extension d’autocomplétion, Cursor transforme l’éditeur en assistant capable de lire l’intégralité de votre base de code, de modifier plusieurs fichiers en même temps et d’exécuter des commandes dans le terminal.

Le tournant majeur remonte à **Cursor 2.0** (29 octobre 2025), qui a introduit deux nouveautés structurantes, comme l’a détaillé InfoWorld. D’abord **Composer**, le premier modèle de codage maison d’Anysphere : un modèle « mixture-of-experts » spécialisé par apprentissage par renforcement, annoncé comme quatre fois plus rapide que les modèles agentiques de niveau comparable, qui boucle la plupart de ses tours en moins de 30 secondes. Ensuite une interface pensée pour le travail multi-agents, où l’on peut lancer plusieurs agents sur des tâches distinctes et comparer leurs résultats. Composer a continué d’évoluer depuis : la version Composer 2.5, lancée le 18 mai 2026, a vu son usage doubler dès sa première semaine, selon LearnAgent, signe de l’appétit des développeurs pour ce modèle maison plus rapide. Le SDK de Cursor a lui aussi été enrichi en juin 2026, avec l’ajout de trois nouvelles primitives – magasins personnalisés (custom stores), outils et auto-révision (auto-review) –, selon JoinNextDev, de quoi faciliter l’intégration de Composer dans des workflows sur mesure.

Cette trajectoire explique l’engouement. Selon les chiffres relayés par la presse spécialisée, Anysphere est passée de 100 millions de dollars de revenus annualisés en janvier 2025 à 500 millions en juin, puis a franchi le milliard en novembre 2025 et les 2 milliards début 2026 ; la valorisation aurait grimpé jusqu’à environ 29,3 milliards de dollars fin 2025, à la faveur d’une Série D de 2,3 milliards de dollars bouclée le 13 novembre 2025, menée par Accel et Coatue selon SiliconANGLE, avec l’entrée de nouveaux investisseurs comme Nvidia et Google, selon Business Wire. Reuters a d’ailleurs souligné que cette valorisation avait presque triplé en cinq mois, passant d’environ 10 milliards de dollars à 29,3 milliards. Plus de la moitié des entreprises du classement Fortune 500 utiliseraient l’outil. Pour comprendre le contexte plus large de cet essor, notre analyse du phénomène du « vibe coding » en 2026 remet ces montants en perspective.

## Prérequis : versions, configuration système et comptes

Avant de lancer ce tutoriel **Cursor AI**, vérifiez que votre poste et vos comptes sont prêts. Cursor est disponible sur les trois grandes plateformes de bureau ; le projet fil rouge, lui, s’appuie sur Python. Voici la liste précise du matériel et des logiciels recommandés en 2026.

| Élément | Minimum | Recommandé | 
|---|---|---|
| Système (Windows) | Windows 10 64 bits | Windows 11 | 
| Système (macOS) | macOS 10.15 Catalina | macOS 13+ (Apple Silicon) | 
| Système (Linux) | Ubuntu 20.04 / .deb, .rpm, AppImage | Ubuntu 22.04+ | 
| Mémoire vive (RAM) | 8 Go | 16 Go ou plus | 
| Espace disque | ~1 Go pour l’éditeur | SSD, 5 Go libres | 
| Connexion Internet | Obligatoire (modèles cloud) | Fibre / haut débit stable | 
| Python (projet) | Python 3.10 | Python 3.12 | 
| Git | Git 2.30+ | Dernière version | 

Côté comptes, prévoyez une adresse e-mail (ou un compte GitHub / Google pour la connexion) et, idéalement, un dépôt Git vide pour versionner votre travail. Cursor fonctionne dès l’offre gratuite *Hobby*, mais l’essai des fonctions agentiques les plus lourdes est plus confortable avec la période d’essai Pro. Installez également **Python 3.12** et **Git** : nous laisserons l’agent créer l’environnement virtuel, mais les interpréteurs doivent être présents sur la machine.

## Le projet fil rouge : une API de gestion de tâches en FastAPI

Pour ne pas rester dans la théorie, nous allons construire un projet complet et fonctionnel : une **API REST de gestion de tâches** (« to-do ») écrite en Python avec le framework FastAPI, une persistance SQLite et une suite de tests automatisés. C’est un cas d’usage réaliste qui met en jeu toutes les capacités de **Cursor** : génération multi-fichiers, exécution du terminal, correction d’erreurs et rédaction de tests.

