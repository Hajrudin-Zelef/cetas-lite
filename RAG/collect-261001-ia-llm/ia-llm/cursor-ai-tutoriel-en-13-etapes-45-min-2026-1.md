---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-1
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: ["Apple", "Intel", "Microsoft"]
dates: []
keywords: ["agent", "agents", "cloud agent", "copilot", "distribution", "intel", "mcp"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [1, 51]
sha256: dc20dcad99a6951343f6bcd294f68747d48dee32f1d8ea058add8ba840125f47
---

# AGENTS.md

Cursor AI a dépassé 2 milliards de dollars de revenu annualisé et franchi le cap du million d’abonnés payants en février 2026, selon les chiffres de TechCrunch et The Information relayés par CloudZero. L’éditeur new-yorkais Anysphere revendique une adoption dans 64 % des entreprises du Fortune 500. Pour les développeurs qui envisagent de basculer vers cet éditeur de code piloté par l’IA, reste une question très concrète : comment le configurer correctement, sans perdre une heure à deviner où cliquer. En France comme dans le reste de l’Europe, cette question s’accompagne souvent d’une autre, plus réglementaire : comment adopter un outil de ce type sans exposer du code sensible ou des données personnelles à des serveurs mal identifiés.

Ce tutoriel couvre l’installation de Cursor AI, la création du compte, le choix de la formule tarifaire, la mise en place des règles de projet (`.cursor/rules` et `AGENTS.md`), le mode Agent, les serveurs MCP et le déploiement d’un Cloud Agent. Treize étapes, 45 minutes environ, avec un projet complet à la fin pour tout mettre en pratique sur du code réel plutôt que sur un simple test isolé. Une section dédiée au déploiement en équipe et un tableau de dépannage avec dix problèmes fréquents complètent l’ensemble.

## Qu’est-ce que Cursor AI et pourquoi son adoption explose en 2026

Cursor AI est un éditeur de code construit à partir de VS Code, développé par la société Anysphere. Contrairement à une extension comme GitHub Copilot qui s’ajoute à un éditeur existant, Cursor AI intègre l’IA directement dans le cœur du produit : autocomplétion multi-lignes, chat contextuel, édition en ligne et surtout un mode Agent capable de modifier plusieurs fichiers de façon autonome.

Cette bascule dépasse le simple effet de mode. Selon le rapport DORA 2025 sur le développement logiciel assisté par IA, 90 % des professionnels de la tech utilisent désormais l’IA au travail, avec un temps médian de deux heures par jour passé sur ces outils. Cursor AI capte une bonne partie de cette demande grâce à son positionnement agentique : au lieu de suggérer une ligne de code isolée, l’outil comprend une tâche entière, écrit les fichiers nécessaires, exécute les commandes utiles et propose une pull request prête à relire.

Cette croissance a aussi un revers. Toujours selon les données relayées par CloudZero, seuls 22 % des directeurs financiers arrivent à relier précisément les dépenses en IA aux résultats business qu’elles produisent. Pour une équipe de développement, cela veut dire une chose simple : configurer Cursor AI correctement dès le départ, avec des règles de projet claires et une formule adaptée à l’usage réel, évite le gaspillage de crédits et les mauvaises surprises en fin de mois.

Cursor AI n’est pas non plus seul sur ce marché. GitHub Copilot reste l’outil le plus utilisé au niveau mondial, avec, selon GitHub, des millions d’utilisateurs individuels et des dizaines de milliers de clients entreprise. Nous détaillons cette comparaison plus loin, mais retenez déjà que le choix entre les deux dépend surtout du niveau d’autonomie que vous voulez confier à l’IA, et du budget que votre équipe peut y consacrer chaque mois.

Pour les équipes françaises, l’intérêt de Cursor AI dépasse le simple gain de vitesse. Les entreprises qui recrutent des développeurs juniors ou en reconversion constatent que l’outil réduit le temps d’appropriation d’un dépôt existant, à condition que les règles de projet soient rédigées correctement dès l’arrivée d’une nouvelle recrue. C’est un usage moins spectaculaire que la génération de code pure, mais souvent plus rentable sur la durée : un fichier `AGENTS.md` bien tenu vaut parfois mieux qu’une journée entière de documentation technique que personne ne relira.

## Prérequis : versions, comptes et configuration matérielle

Avant de lancer l’installation, vérifiez que votre poste de travail répond aux exigences minimales. Cursor AI tourne sur les trois grands systèmes d’exploitation et ne demande pas de machine surpuissante, mais quelques prérequis évitent les mauvaises surprises, surtout si vous comptez faire tourner des serveurs MCP en local.

| Élément | Minimum requis | Recommandé | 
|---|---|---|
| Système d’exploitation | Windows 10, macOS 12, distribution Linux récente | Windows 11, macOS 14 ou plus | 
| Mémoire vive | 8 Go | 16 Go ou plus | 
| Espace disque | 2 Go libres | 5 Go libres (cache et extensions inclus) | 
| Connexion Internet | Requise en permanence | Connexion stable, faible latence | 
| Compte | Adresse email | Compte GitHub (synchronisation des dépôts) | 
| Node.js (optionnel) | Non requis pour l’éditeur seul | Version LTS, utile pour lancer des serveurs MCP locaux | 

Un compte GitHub n’est pas obligatoire, un simple email suffit pour démarrer. Il devient toutefois utile assez vite : il permet de synchroniser vos dépôts, et surtout de déclencher un Cloud Agent directement depuis un commentaire sur une pull request. Pour les équipes soumises au RGPD, prévoyez aussi un point de vigilance : le code envoyé au modèle transite par les serveurs de Cursor AI ou de ses fournisseurs de modèles, donc les projets sous NDA ou contenant des données personnelles méritent une revue avant connexion, surtout si vous ajoutez des serveurs MCP donnant accès à des bases de données internes.

Si votre équipe vient d’un autre éditeur, la migration reste globalement simple puisque Cursor AI, Windsurf et VS Code partagent la même base technique et le même format d’extensions. La bascule depuis un IDE plus éloigné comme IntelliJ ou Eclipse demande en revanche un peu plus de temps d’adaptation, principalement sur les raccourcis clavier et l’organisation des panneaux, même si l’essentiel des concepts (explorateur de fichiers, terminal intégré, débogueur) reste familier.

## Étapes 1 à 3 : télécharger, installer et lancer Cursor AI

### Étape 1 – Télécharger l’installeur adapté à votre système

Rendez-vous sur le site officiel de Cursor et cliquez sur le bouton de téléchargement. Le site détecte automatiquement votre système d’exploitation et propose le bon fichier : binaire universel pour macOS (Apple Silicon et Intel), exécutable pour Windows, ou paquet AppImage/.deb pour Linux. Aucune inscription n’est nécessaire à cette étape, la création de compte intervient au premier lancement.

### Étape 2 – Installer et importer vos réglages VS Code

Lancez l’installeur téléchargé et suivez l’assistant. Sur macOS, glissez simplement l’application dans le dossier Applications. Au premier démarrage, Cursor AI propose d’importer en un clic vos extensions, raccourcis clavier et thème depuis une installation VS Code existante. C’est l’un des vrais avantages d’un éditeur construit sur la même base : la migration prend quelques secondes plutôt qu’une demi-journée de reconfiguration.

Comme la plupart des forks de VS Code, Cursor installe généralement un raccourci en ligne de commande. Une fois l’application ouverte au moins une fois, vous pouvez normalement ouvrir n’importe quel dossier de travail directement depuis un terminal :

```
cd mon-projet
cursor .
```
### Étape 3 – Premier lancement et prise en main de l’interface

