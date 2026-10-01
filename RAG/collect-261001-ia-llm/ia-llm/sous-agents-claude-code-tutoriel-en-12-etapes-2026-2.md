---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-2
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "arr", "claude", "mcp", "memory", "opus 5"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [51, 139]
sha256: 18eae4849876c7ea6cf0e2c69479087ee5a0e39af0e7f39419b8dc80aafd6e1e
---

# ./scripts/validate-readonly-query.sh

```
---
name: code-improver
description: Scans files and suggests improvements for readability, performance, and best practices. Use after writing or modifying code.
tools: Read, Grep, Glob
model: sonnet
---
You are a code improvement specialist. For each issue you find, explain
the problem, show the current code, and provide an improved version.
```
Comme ce fichier vit dans `~/.claude/agents/`, le sous-agent devient disponible dans tous vos projets. Pour le limiter à un seul projet, déplacez-le dans le dossier `.claude/agents/` de ce projet.

Un changement de version mérite d’être signalé si vous suivez un ancien tutoriel : jusqu’à la version 2.1.197, la commande `/agents` ouvrait un assistant interactif avec un onglet Running listant les sous-agents actifs et un onglet Library pour créer, modifier et supprimer des définitions. Depuis la version 2.1.198, `/agents` n’ouvre plus cet assistant : elle affiche simplement un rappel d’éditer `.claude/agents/` directement ou de demander à Claude de le faire. Les emplacements de fichiers et les champs du frontmatter, eux, n’ont pas changé, seul l’assistant en terminal a été retiré.

### Méthode B : écrire le fichier à la main

Le même résultat s’obtient directement, sans passer par Claude :

```
---
name: code-reviewer
description: Reviews code for quality and best practices
tools: Read, Glob, Grep
model: sonnet
---
You are a code reviewer. When invoked, analyze the code and provide
specific, actionable feedback on quality, security, and best practices.
```
Testez ensuite votre sous-agent en lui déléguant explicitement une tâche : « Utilise l’agent code-improver pour suggérer des améliorations sur ce projet ». Si Claude ne trouve pas le nouveau sous-agent, redémarrez Claude Code. Ce cas se produit uniquement quand `~/.claude/agents/` n’existait pas encore au lancement de la session : un processus déjà démarré ne détecte pas un dossier `agents` fraîchement créé. En dehors de ce cas précis, Claude Code surveille `~/.claude/agents/` et `.claude/agents/` et applique une modification de fichier en quelques secondes, sans redémarrage.

## Étape 3 : maîtriser les champs du frontmatter YAML

Seuls deux champs sont obligatoires : `name` et `description`. Tous les autres sont optionnels et prennent une valeur par défaut raisonnable si vous les omettez. Voici la liste complète documentée par Anthropic.

| Champ | Obligatoire | Rôle | 
|---|---|---|
| name | Oui | Identifiant unique en minuscules avec traits d’union | 
| description | Oui | Indique à Claude quand déléguer à ce sous-agent | 
| tools | Non | Liste blanche d’outils autorisés, hérite de tout si omis | 
| disallowedTools | Non | Liste noire d’outils retirés du pool hérité ou spécifié | 
| model | Non | sonnet, opus, haiku, fable, un ID complet, ou inherit (défaut) | 
| permissionMode | Non | default, acceptEdits, auto, dontAsk, bypassPermissions ou plan | 
| maxTurns | Non | Nombre maximal de tours agentiques avant arrêt | 
| skills | Non | Skills à précharger intégralement dans le contexte au démarrage | 
| mcpServers | Non | Serveurs MCP disponibles pour ce sous-agent uniquement | 
| hooks | Non | Hooks de cycle de vie propres à ce sous-agent | 
| memory | Non | Portée de mémoire persistante : user, project ou local | 
| background | Non | Force l’exécution en arrière-plan si réglé sur true | 
| isolation | Non | worktree pour une copie Git isolée du dépôt | 
| color | Non | Couleur d’affichage dans la liste des tâches | 

Le corps du fichier, sous l’en-tête YAML, devient le prompt système du sous-agent. Un point souvent mal compris : un sous-agent ne reçoit que ce prompt, plus quelques détails d’environnement comme le répertoire de travail. Il ne reçoit pas le prompt système complet de Claude Code. Gardez donc vos instructions autonomes et explicites, sans supposer que le sous-agent connaît le contexte de votre conversation en cours.

Ce que le sous-agent charge réellement à son démarrage va plus loin que son seul prompt. Il reçoit aussi le message de délégation que Claude rédige pour résumer la tâche, toute la hiérarchie de vos fichiers CLAUDE.md (utilisateur, projet, CLAUDE.local.md), et un instantané de l’état Git pris au tout début de la session parente. À l’inverse, certains éléments de la conversation principale ne franchissent jamais cette frontière : votre style de sortie personnalisé, la mémoire automatique de la session principale, et la taille de fenêtre de contexte, qui dépend du modèle propre au sous-agent plutôt que de celui de la conversation qui l’a lancé. Depuis juillet 2026, Claude Opus 5 est devenu le modèle Opus par défaut de Claude Code et conserve la fenêtre de contexte d’un million de tokens, selon Anthropic : un sous-agent configuré sur ce modèle peut à lui seul ingérer bien plus de code, ce qui change la donne pour les workflows qui enchaînent plusieurs sous-agents volumineux.

## Étape 4 : choisir la bonne portée (scope) pour chaque sous-agent

L’emplacement du fichier détermine qui peut utiliser le sous-agent. Quand plusieurs sous-agents partagent le même nom, Claude Code applique celui de l’emplacement le plus prioritaire.

| Emplacement | Portée | Priorité | Création | 
|---|---|---|---|
| Paramètres managés | Toute l’organisation | 1 (la plus haute) | Déployé via les paramètres managés | 
| Option CLI –agents | Session en cours | 2 | JSON passé au lancement de Claude Code | 
| .claude/agents/ | Projet en cours | 3 | Demandé à Claude ou créé manuellement | 
| ~/.claude/agents/ | Tous vos projets | 4 | Demandé à Claude ou créé manuellement | 
| Dossier agents/ d’un plugin | Là où le plugin est actif | 5 (la plus basse) | Installé avec un plugin | 

Les sous-agents de projet conviennent à un travail spécifique à un dépôt : versionnez-les pour que toute l’équipe en profite. Claude Code les découvre en remontant depuis le répertoire de travail courant, donc chaque dossier `.claude/agents/` rencontré entre ce répertoire et la racine du dépôt est scanné. Les sous-agents utilisateur, eux, restent personnels mais actifs sur toutes vos machines et tous vos projets.

Gardez les valeurs `name` uniques dans toute l’arborescence d’un même dossier `.claude/agents/`, sous-dossiers compris. Si deux fichiers déclarent le même nom, Claude Code n’en charge qu’un seul, choisi par l’ordre de lecture du système de fichiers plutôt que par une règle documentée. La commande `/doctor` signale ces doublons et propose de renommer ou de supprimer tous les fichiers sauf un.

Pour une définition rapide qui ne survit qu’à la session, la syntaxe `–agents` accepte du JSON directement en ligne de commande, pratique pour tester ou automatiser un script : c’est la version CLI du même mécanisme que le SDK Agent expose via le paramètre `agents` des options de `query()`, qui permet de faire naître plusieurs sous-agents en une seule requête programmatique. Au total, le SDK reconnaît trois modes de création de sous-agents : programmatique via ce paramètre, fichier Markdown dans `.claude/agents/`, ou sous-agent intégré comme Explore ou Plan.

```
claude --agents '{
  "code-reviewer": {
    "description": "Expert code reviewer. Use proactively after code changes.",
    "prompt": "You are a senior code reviewer. Focus on code quality, security, and best practices.",
    "tools": ["Read", "Grep", "Glob", "Bash"],
    "model": "sonnet"
  },
  "debugger": {
    "description": "Debugging specialist for errors and test failures.",
    "prompt": "You are an expert debugger. Analyze errors, identify root causes, and provide fixes."
  }
}'
```
## Étape 5 : restreindre les outils et les permissions

