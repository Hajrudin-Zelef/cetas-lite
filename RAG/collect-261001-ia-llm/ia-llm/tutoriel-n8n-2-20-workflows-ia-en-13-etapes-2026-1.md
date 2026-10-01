---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026-1
title: "Nœud Code – Langage Python"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Intel", "OpenAI"]
dates: []
keywords: ["agent", "agents", "intel", "license", "mai", "mcp"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 48]
sha256: f09dda33c590de543ba45444d013beea4088d9a93a5558bb76580d8df59a6d98
---

# Nœud Code – Langage Python

n8n est devenu en 2026 la plateforme d’automatisation de référence en Europe avec plus de **74 000 recherches mensuelles** en France selon DataForSEO, devant Zapier et Make sur le marché de l’orchestration low-code. Selon Taskade, la plateforme revendique déjà plus de **230 000 utilisateurs actifs** et plus de **150 000 étoiles GitHub** dès février 2026. La version **n8n 2.36.8**, publiée en août 2026 selon le suivi GitHub de n8n-assistant, reste à ce jour la dernière stable de la branche 2.x (la bêta 2.37.4 constituant la ligne expérimentale la plus récente du même mois selon n8n Docs), et confirme un rythme de mises à jour mineures hebdomadaire, un AI Builder repensé, des nœuds Code en Python plus rapides et une stabilité accrue sur l’auto-hébergement Docker. Ce tutoriel vous guide pas à pas pour **installer n8n**, créer vos premiers workflows et déployer un projet IA complet en moins de 90 minutes.

Publié le **16 avril 2026** par la rédaction Tech Insider et mis à jour à la lumière du tutoriel officiel « Build your first AI workflow », publié le 5 février 2026 par l’équipe n8n Community en ouverture de sa série 2026 dédiée à l’IA, ce guide couvre l’installation Docker, la connexion aux LLM (OpenAI, Anthropic, Ollama), les webhooks, les nœuds de base de données PostgreSQL et un workflow IA complet de tri d’e-mails avec LangChain. Tous les exemples ont été testés sur Ubuntu 24.04 LTS et macOS 15. À la fin, vous disposerez d’un projet n8n auto-hébergé prêt pour la production, capable d’orchestrer des appels API, des bases vectorielles et des agents IA.

## Pourquoi n8n s’impose en 2026 face à Zapier et Make

La plateforme n8n a connu un tournant majeur fin 2025 avec la sortie de **n8n 2.0**, publié en bêta le 8 décembre 2025 puis en stable le 15 décembre 2025 selon le blog officiel – une dynamique confirmée par la croissance de son dépôt GitHub, passé de **100 000 étoiles en mai 2025** à **150 000 en octobre 2025**, soit une hausse de **+50 %** en cinq mois selon Taskade. Là où Zapier facture chaque tâche exécutée et Make limite les opérations mensuelles, n8n applique un modèle **fair-code** qui autorise l’auto-hébergement gratuit pour un usage interne ou personnel, avec une licence Sustainable Use License gérée par la maison-mère allemande n8n GmbH basée à Berlin.

Le rythme de publication confirme la maturité de la plateforme. La documentation officielle n8n Docs indique qu’une nouvelle version mineure sort **environ une fois par semaine**, avec une bifurcation claire entre canaux stable et bêta : après la branche **2.20**, lancée en mai 2026 avec la vérification du Netlify Trigger, vient la **2.23** du 27 mai 2026, puis la **2.31.0** du 14 juillet 2026 qui ajoute une API de politique de sécurité et des mises à jour de l’assistant IA, avant la **2.34** du 4 août 2026 et son support de la déconnexion OIDC parmi 18 autres fonctionnalités, jusqu’à la **2.36.8**, dernière stable en date. Cette cadence rapide tranche avec les plateformes propriétaires concurrentes, où les utilisateurs attendent souvent des trimestres pour des correctifs.

Le second atout de n8n tient à son positionnement **code + no-code**. Contrairement à Zapier qui interdit l’écriture de code arbitraire sur les plans gratuits, le nœud Code de n8n exécute du JavaScript ou du Python natif depuis l’éditeur visuel. Selon la note de version du Latenode, la mise à jour récente du nœud Code prend en charge une version Python plus moderne et un éventail élargi de bibliothèques, ce qui ouvre la porte à des transformations complexes – parsing de PDF, calcul scientifique, requêtes pandas – sans sortir du workflow.

Enfin, l’écosystème **AI Builder** intégré depuis fin 2025 propulse n8n dans la catégorie des orchestrateurs d’agents IA. Le nœud AI Agent dispose désormais d’une gestion plus efficace des tokens et d’une intégration native avec LangChain, ce qui fait de n8n une alternative crédible à LangGraph ou CrewAI pour les équipes qui préfèrent le visuel au code Python.

## Comparatif n8n vs Zapier vs Make en avril 2026

| Critère | n8n 2.20.9 | Zapier | Make | 
|---|---|---|---|
| Modèle de licence | Fair-code (Sustainable Use) | Propriétaire SaaS | Propriétaire SaaS | 
| Auto-hébergement | Gratuit illimité | Non disponible | Non disponible | 
| Plan Cloud de départ | 20 € / mois (Starter) | 19,99 $ / mois | 9 € / mois | 
| Nœud Code (JS/Python) | Oui, natif | Limité (payant) | Oui (JS uniquement) | 
| Nœud AI Agent natif | Oui, intégré LangChain | Via Zaps IA payants | Limité | 
| Cadence de releases | Hebdomadaire | Mensuelle | Trimestrielle | 
| Volume recherches France | 74 000 / mois | 40 500 / mois | 27 100 / mois | 
| Versions LTS | Stable + Beta | N/A | N/A | 

Les chiffres de recherche issus de DataForSEO d’avril 2026 placent n8n loin devant ses concurrents en France. Cette dynamique s’explique par la popularité croissante des architectures **self-hosted** dans un contexte européen où le RGPD et le futur AI Act poussent les entreprises à reprendre le contrôle de leurs flux de données. n8n permet de garder données et clés API dans un cluster Kubernetes interne ou sur un VPS Hetzner sans dépendre d’un SaaS américain.

## Prérequis et versions exactes pour ce tutoriel n8n

Avant de démarrer, vérifiez que votre poste de travail dispose des composants suivants. La compatibilité a été testée sur Ubuntu 24.04 LTS, Debian 12 et macOS 15 (Apple Silicon et Intel). Sur Windows, utilisez impérativement WSL 2 avec Ubuntu pour éviter les problèmes de permissions Docker. Notez que le suivi de versions WinterFlow.io recense la **2.19.5** du 7 mai 2026, qui corrigeait des problèmes HTTPS liés à Simple-Git, puis la **2.29.0** du 30 juin 2026 parmi plus de cinq versions 2.x publiées au premier semestre 2026 ; le même catalogue documente désormais la version 1.120.0 avec un support élargi de MySQL/MariaDB en plus de PostgreSQL, ainsi qu’une prise en charge native d’OAuth et du protocole MCP, des options utiles si votre infrastructure existante repose déjà sur MySQL.

- **Node.js 22 LTS** ou supérieur (npm 10+) pour l’installation globale via npm
- **Docker 27.5+** et**Docker Compose v2.32+** pour le déploiement conteneurisé recommandé
- **PostgreSQL 17** recommandé pour la base de données de production (SQLite par défaut)
- **4 Go de RAM** minimum, 8 Go conseillés pour les workflows IA
- **10 Go d’espace disque** pour les binaires Docker et l’historique d’exécutions
- Un compte **OpenAI** ou**Anthropic** pour les exemples IA, ou Ollama auto-hébergé
- Un domaine et un certificat **HTTPS** Let’s Encrypt pour les webhooks de production
- **Git 2.46+** pour cloner le dépôt de configuration

Côté connaissances, ce tutoriel suppose une familiarité basique avec la ligne de commande Linux, le format JSON et le concept de variable d’environnement. Aucune expérience préalable de n8n n’est requise. Si vous avez déjà utilisé Ollama dans Docker, le déploiement vous semblera familier.

## Étape 1 : Installer n8n 2.20.9 avec Docker Compose

L’installation via Docker reste la méthode recommandée par la documentation officielle pour 2026. Elle isole l’application, simplifie les mises à jour hebdomadaires et permet de basculer rapidement entre versions stable et bêta. Créez un dossier dédié et un fichier `docker-compose.yml` avec la configuration suivante.

