---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes-1
title: "Verifier la version de Node.js (20 ou superieur requis)"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Mistral", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "agents", "license", "mistral", "nvidia", "open source"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes.md
source_anchor: ""
source_lines: [1, 56]
sha256: 417d469b71709b28a2f40883e97402e155391a1c439fbb9ea25329f7b64b16fb
---

# Verifier la version de Node.js (20 ou superieur requis)

Avec plus de **74 000 recherches mensuelles** en France et une communauté qui dépasse les **100 000 étoiles sur GitHub**, **n8n** s’est imposé en 2026 comme la référence de l’automatisation de workflows pour les développeurs et les équipes techniques européennes. Contrairement à Zapier ou Make, n8n est *auto-hébergeable* : vos données ne quittent jamais votre infrastructure, un argument décisif à l’heure du RGPD et de la souveraineté numérique. Ce **tutoriel n8n** complet vous guide, en 13 étapes concrètes, de l’installation avec Docker jusqu’au déploiement en production d’un véritable agent d’automatisation piloté par IA.

À la fin de ce guide, vous disposerez d’un **n8n auto-hébergé** fonctionnel, connecté à une base PostgreSQL, exécutant un workflow complet qui récupère des données via API, les transforme en JavaScript, les fait analyser par un grand modèle de langage (OpenAI, Mistral ou Ollama local) et envoie un résumé par e-mail ou sur Slack. Comptez environ **45 à 60 minutes** pour suivre l’intégralité du tutoriel. Tous les fichiers de configuration et extraits de code sont fournis et testés sur la version stable **n8n 2.38.5** (avec la bêta **2.39.1** déjà accessible aux testeurs), en vigueur depuis septembre 2026 selon la documentation officielle du projet.

## n8n en 2026 : pourquoi l’automatisation auto-hébergée explose

**n8n** (prononcé « n-eight-n », pour *nodemation*) est un outil d’automatisation de workflows open source développé par **n8n GmbH**, une société basée à Berlin et fondée par **Jan Oberhauser**. Lancé fin 2019, le projet a connu une croissance fulgurante : en mars 2025, l’entreprise a d’abord bouclé un tour de série B de **55 millions d’euros** (environ 60 millions de dollars selon TechCrunch) mené par **Highland Europe**, valorisant n8n autour de **250 millions d’euros** (~270 millions de dollars). Sept mois plus tard, le **9 octobre 2025**, n8n a annoncé une série C de **180 millions de dollars** menée par **Accel** et **NVentures** (la branche investissement de NVIDIA), portant son financement total cumulé à **240 millions de dollars** à l’époque et sa valorisation à **2,5 milliards de dollars**, comme l’ont rapporté Bloomberg et TechFundingNews. Selon le décompte plus récent de Startup Intros, le financement cumulé de n8n atteint désormais **258,2 millions de dollars** sur quatre levées de fonds au total, à la date de juillet 2026. Cette dynamique illustre l’appétit du marché pour des alternatives auto-hébergées aux plateformes SaaS américaines.

Le principe de n8n est simple : vous construisez visuellement des **workflows** en reliant des « nœuds » (nodes) dans un éditeur en glisser-déposer. Chaque nœud effectue une action – déclencher sur un événement, appeler une API, transformer des données, interroger un LLM, envoyer un message. n8n propose **plus de 400 intégrations natives** (et plus de 1 800 dans son répertoire élargi incluant les nœuds communautaires), couvrant Google, Slack, Notion, GitHub, les bases de données SQL, et désormais une large gamme de fournisseurs d’IA.

Pour les entreprises françaises et européennes, l’argument central reste la **souveraineté des données**. En auto-hébergeant n8n sur un serveur situé dans l’UE (ou sur votre propre machine), vous conservez la maîtrise totale du traitement des données personnelles, ce qui simplifie considérablement la conformité au **RGPD**. C’est exactement la logique défendue par les acteurs de la souveraineté numérique européenne, et c’est ce qui distingue radicalement n8n des solutions cloud propriétaires.

Côté licence, n8n adopte un modèle **fair-code** sous *Sustainable Use License* : le code source est public et utilisable gratuitement pour un usage interne, y compris commercial, mais certaines utilisations (revente du produit en tant que service concurrent) sont restreintes. Pour la quasi-totalité des cas d’usage – automatiser sa veille, ses notifications, ses synchronisations de données ou ses agents IA – l’auto-hébergement de n8n est **100 % gratuit**.

## Prérequis : versions, matériel et comptes nécessaires

Avant de lancer ce **tutoriel n8n**, assurez-vous de disposer des éléments suivants. La méthode recommandée – et celle que nous suivrons – repose sur Docker, qui évite tout conflit de dépendances et garantit un environnement reproductible. Côté version auto-hébergée, la documentation officielle de n8n référence désormais la **2.38.5** comme version stable, accompagnée d’une bêta **2.39.1** déjà ouverte aux testeurs en septembre 2026 : privilégiez la 2.38.5 (ou une révision ultérieure) pour une installation stable en production. Si vous débutez avec la conteneurisation, notre tutoriel Docker complet couvre les bases nécessaires.

| Prérequis | Version minimale | Rôle | 
|---|---|---|
| Docker Engine | 24.0 ou supérieur | Exécuter le conteneur n8n | 
| Docker Compose | v2.20+ | Orchestrer n8n + PostgreSQL | 
| Node.js (si install npm) | 20 LTS ou supérieur | Installation alternative sans Docker | 
| RAM disponible | 2 Go minimum (4 Go recommandés) | Exécutions et éditeur | 
| Espace disque | 5 Go minimum | Image Docker + base + données | 
| Système d’exploitation | Linux, macOS, Windows 10+ | Hôte du conteneur | 
| Clé API LLM | OpenAI / Mistral / Ollama | Étape agent IA | 

Pour la partie IA, vous aurez besoin d’une clé API. Trois options sont possibles : un compte **OpenAI** (clé commençant par `sk-`), un compte **Mistral AI** – solution souveraine française idéale pour rester dans l’UE – ou une instance **Ollama** locale pour une confidentialité totale et un coût nul. Si vous optez pour la voie 100 % locale, consultez notre tutoriel Ollama pour exécuter un LLM local avant de commencer.

Vérifiez vos versions de Docker avant de continuer :

```
$ docker --version
Docker version 24.0.7, build afdd53b
$ docker compose version
Docker Compose version v2.23.3
```
## n8n vs Zapier vs Make : le comparatif 2026

Avant de plonger dans l’installation, situons n8n face à ses concurrents directs. Zapier et Make (ex-Integromat) sont des plateformes SaaS propriétaires facturées à l’exécution ou à la tâche, sans option d’auto-hébergement. n8n se distingue par son modèle ouvert, son nœud **Code** qui autorise du JavaScript arbitraire, et sa capacité à tourner sur votre propre infrastructure.

| Critère | n8n | Zapier | Make | 
|---|---|---|---|
| Auto-hébergement | Oui (gratuit) | Non | Non | 
| Code personnalisé | JavaScript & Python (nœud Code) | Limité | Limité | 
| Modèle de licence | Fair-code (open source) | Propriétaire | Propriétaire | 
| Intégrations natives | 400+ (1 800+ avec communauté) | 7 000+ | 1 700+ | 
| Conformité RGPD (données dans l’UE) | Totale (auto-hébergé) | Dépend du plan | Dépend du plan | 
| Nœuds IA / agents | Oui (LangChain intégré) | Oui | Oui | 
| Coût pour gros volumes | Fixe (coût serveur) | Croît avec les tâches | Croît avec les opérations | 

Zapier conserve l’avantage du nombre brut d’intégrations (plus de 7 000) et d’une prise en main grand public. Mais dès que les volumes augmentent ou que la confidentialité devient critique, l’économie penche nettement vers n8n : son coût est celui d’un serveur, indépendant du nombre d’exécutions. Pour une équipe technique qui automatise des milliers de tâches par jour, l’écart de facture devient considérable. n8n rejoint ainsi la logique des agents IA construits avec l’API Mistral : maîtrise, ouverture et souveraineté.

## Installation alternative : n8n via npm sans Docker

