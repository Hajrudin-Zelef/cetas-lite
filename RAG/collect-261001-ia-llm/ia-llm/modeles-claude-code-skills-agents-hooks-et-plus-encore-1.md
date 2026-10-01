---
id: collect-261001-ia-llm/ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore-1
title: "Database Migration Skill"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "attention", "claude", "distribution", "mcp"]
source: docs/RAG/collect-261001-ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [1, 116]
sha256: d690e47776c0ad9e6676a5b9ebe5fad818159bfb430c058debee9026de861301
---

# Database Migration Skill

Cours

Claude Code est puissant, mais sans configuration réutilisable, vous vous retrouvez à répéter les mêmes instructions encore et encore. Les modèles Claude Code résolvent ce problème : ils transforment des consignes récurrentes, des workflows, des autorisations d’outils et des intégrations en fichiers de projet réutilisables que Claude peut découvrir et appliquer.

Dans cet article, nous verrons ce que sont les modèles Claude Code, les principaux types disponibles, la manière dont chacun se comporte, comment choisir le bon type pour votre flux de travail, et où trouver des modèles prêts à l’emploi.

Cet article suppose que vous avez déjà une configuration Claude Code de base. Si vous débutez, commencez par ce tutoriel Claude Code avant d’aller plus loin avec les modèles. Si vous apprenez encore comment Claude Code s’intègre au développement en ligne de commande, consultez ce guide d’introduction à la Claude Code CLI.

## En bref

- 
Les modèles Claude Code sont des configurations réutilisables basées sur des fichiers (stockées dans `.claude/` ) qui vous évitent de réexpliquer votre stack et vos workflows à chaque session.
- 
Il en existe six types : skills (workflows répétables), agents (rôles et permissions cadrés), commandes (actions manuelles en slash), hooks (garde-fous automatiques), MCP (connexions à des outils et données externes) et plugins (regroupements des cinq précédents).
- 
`CLAUDE.md` contient toujours le brief de votre projet ; les modèles y ajoutent des comportements modulaires et réutilisables.
- 
Choisissez selon le déclencheur : les hooks appliquent des règles automatiquement, les agents apportent une expertise métier, les skills encodent des workflows répétables, les commandes s’exécutent à la demande, MCP accède à des systèmes externes, et les plugins emballent et partagent une configuration complète.
- 
Commencez petit avec la documentation officielle d’Anthropic, une collection communautaire comme aitmpl.com, ou vos propres fichiers, et créez un premier skill pour votre tâche la plus répétée avant d’étendre.

## Introduction aux agents d'intelligence artificielle

## Que sont les modèles Claude Code ?

Les modèles Claude Code sont des fichiers de configuration réutilisables qui personnalisent le comportement de Claude Code dans un projet ou à l’échelle de votre environnement local.

Point essentiel : les modèles sont basés sur des fichiers. Vous ne les installez pas via une interface de réglages classique pour cliquer sur des écrans de configuration. Claude Code détecte plutôt des fichiers et dossiers spécifiques, charge les métadonnées pertinentes dans le contexte et s’appuie dessus pour décider comment se comporter.

En pratique, il s’agit généralement de fichiers Markdown, JSON ou orientés shell, stockés dans des dossiers au niveau du projet comme `.claude/`, ou empaquetés dans des répertoires de type plugin pour le partage.

Une structure typique au niveau projet peut ressembler à ceci :

```
my-app/
├── CLAUDE.md
├── .mcp.json
└── .claude/
    ├── skills/
    │   └── database-migration/
    │       └── SKILL.md
    ├── agents/
    │   └── security-auditor.md
    ├── commands/
    │   └── summarize-pr.md
    └── settings.json
```
### Modèles Claude Code vs CLAUDE.md

`CLAUDE.md` reste important, mais son rôle est différent. Considérez `CLAUDE.md` comme le brief du projet : ce qu’est le projet, quelles commandes comptent, quels standards de code s’appliquent et quelles conventions d’architecture Claude doit garder en mémoire.

Pour un guide pas-à-pas, consultez notre guide de rédaction CLAUDE.md.

Les modèles sont plus modulaires :

- Un skill peut encoder un workflow de migration.
- Un agent peut isoler une posture de revue sécurité.
- Un hook peut s’exécuter après des éditions de fichiers.
- Une configuration MCP peut connecter Claude à GitHub, SQLite ou un autre système externe.

C’est aussi là que les modèles s’articulent avec la conception globale de votre workflow Claude Code. Des modèles solides donnent le meilleur résultat lorsqu’ils s’accompagnent de bonnes pratiques de planification, de test et de passation de contexte.

Retrouvez davantage de ces pratiques dans notre guide des bonnes pratiques.

Les commandes personnalisées relèvent également du système de skills, même si l’ancien format `.claude/commands/` fonctionne encore. Le nouveau format recommandé est `.claude/skills/<name>/SKILL.md`, qui prend en charge l’appel par commande slash et l’appel automatique par Claude.

## Quels types de modèles Claude Code puis-je utiliser ?

L’écosystème des modèles Claude Code se structure généralement en six catégories : skills, agents, commandes, hooks, intégrations MCP et plugins.

Les cinq premiers modifient directement le comportement de Claude. Les plugins sont un peu différents : il s’agit d’un format de distribution pouvant regrouper skills, agents, hooks, commandes, serveurs MCP et autres composants dans un package réutilisable.

Nous allons examiner chacune de ces catégories ci-dessous.

### 1. Skills

Les skills sont des ensembles d’instructions pour des tâches multi-étapes répétables. Un skill est généralement un dossier contenant un fichier `SKILL.md` avec un frontmatter YAML et un corps en Markdown.

Le frontmatter décrit ce que fait le skill et comment il doit se comporter ; le corps indique à Claude les étapes à suivre. Pour un décryptage dédié, consultez ce guide Claude Skills.

Claude utilise la description du skill pour décider quand il est pertinent. Par défaut, l’utilisateur comme Claude peuvent appeler un skill : vous pouvez taper `/skill-name`, ou Claude peut le charger automatiquement lorsque la tâche en cours correspond à sa description. Vous pouvez aussi désactiver l’appel automatique du modèle pour les workflows où vous souhaitez garder la main, comme le déploiement.

Voici un court exemple de fichier : `.claude/skills/database-migration/SKILL.md`

```
---
name: database-migration
description: Use when creating, reviewing, or modifying database migrations. Ensures migrations are reversible, tested, and checked before and after execution.
allowed-tools:
  - Read
  - Write
  - Bash
---
# Database Migration Skill
When working on a database migration:
1. Inspect the existing schema and migration history before writing changes.
2. Confirm whether the migration is additive, destructive, or data-transforming.
3. Create a reversible migration whenever the framework supports rollback.
4. Run the project’s migration check command before applying the migration.
5. Run tests that cover the affected models, queries, or API endpoints.
6. After writing the migration, summarize:
   - schema changes
   - rollback behavior
   - affected tables
   - test commands run
```
C’est utile car les instructions sont procédurales. Vous ne dites pas seulement à Claude de « faire attention aux migrations » ; vous lui fournissez une checklist répétable.

Les skills conviennent à tout ce que vous colleriez sinon dans Claude plus de deux fois : génération d’endpoints d’API, rédaction de journaux de modifications, échafaudage de tests, création de notes de version, revue de pull requests ou vérifications de migration.

Pour plus d’inspiration sur les workflows que les développeurs transforment en automatisations IA réutilisables, consultez notre liste Agent Skills.

### 2. Agents

Les agents, plus précisément des sous-agents personnalisés dans Claude Code, sont des assistants IA spécialisés avec leur propre définition Markdown, frontmatter YAML, restrictions d’outils, choix de modèle et prompt système.

