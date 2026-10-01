---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-4
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2023-06-01", "2025-04-14", "2025-10-02"]
keywords: ["claude", "agents", "chatgpt", "distribution", "mai", "mcp"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [213, 294]
sha256: 37cd7f18f3f782919826cdd84bdda62434e47f9fb4c303306262f2f4aa95c7ef
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

```
curl https://api.anthropic.com/v1/messages \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: skills-2025-10-02,files-api-2025-04-14" \
  -H "content-type: application/json" \
  -d '{
    "model": "claude-opus-4-8",
    "max_tokens": 4096,
    "container": {
      "skills": [{"type": "pptx", "skill_id": "pptx"}]
    },
    "messages": [
      {"role": "user", "content": "Crée une présentation de 5 slides résumant ce rapport trimestriel."}
    ]
  }'
```
Sur ce canal, l’environnement d’exécution tourne dans un conteneur isolé, sans accès réseau et sans possibilité d’installer de nouveaux paquets à la volée : seules les dépendances préconfigurées par Anthropic sont disponibles. C’est une contrainte à garder en tête si vous comptez faire tourner une Skill personnalisée gourmande en dépendances externes via l’API plutôt que via Claude Code, où l’accès réseau est complet.

Pour créer et téléverser vos propres Skills sur ce canal, l’API expose des points de terminaison dédiés sous `/v1/skills`. Une fois téléversées, les Skills personnalisées de l’API sont partagées à l’échelle de tout l’espace de travail (workspace) : tous les membres y ont accès, contrairement à claude.ai où chaque Skill personnalisée reste individuelle à l’utilisateur qui l’a importée.

## Étape 12 : empaqueter et partager via un plugin ou Git

Deux méthodes de partage principales existent pour une équipe qui travaille dans Claude Code. La plus simple consiste à committer le dossier `.claude/skills/` à la racine du dépôt : chaque membre de l’équipe qui clone le projet et lance Claude Code récupère automatiquement les Skills, sans étape d’installation supplémentaire. C’est l’approche à privilégier pour tout ce qui est spécifique à un seul projet.

Pour une Skill destinée à plusieurs dépôts, la distribution via un plugin Claude Code est plus adaptée : elle permet de regrouper Skills, agents, hooks et serveurs MCP dans un seul paquet installable. Les Skills livrées par un plugin sont automatiquement préfixées par son nom (`mon-plugin:ma-skill`), ce qui évite tout conflit avec les Skills personnelles ou de projet des utilisateurs qui l’installent. À noter : ajouter un fichier `.claude-plugin/plugin.json` directement dans un dossier de Skill suffit à le faire charger comme un plugin à part entière, capable d’embarquer aussi des agents, des hooks et des serveurs MCP.

Pour une distribution en dehors de l’écosystème Claude Code (téléversement sur claude.ai ou via l’API), utilisez le script `package_skill.py` fourni dans le dépôt officiel anthropics/skills. Il valide que votre frontmatter respecte bien les six champs autorisés par la spécification portable avant de générer l’archive à téléverser.

## Étape 13 : publier sur claude.ai et gérer le cross-surface

Sur claude.ai, direction Paramètres puis la section Fonctionnalités (*Settings > Features*), où vous pouvez téléverser une Skill personnalisée sous forme d’archive ZIP, à condition d’être sur un plan Pro, Max, Team ou Enterprise avec l’exécution de code activée. L’enjeu dépasse le simple confort individuel : selon le Ramp AI Index publié en mai 2026, Claude devançait déjà ChatGPT en usage professionnel en avril 2026, avec 34,4 % de parts d’adoption en entreprise contre 32,3 % pour son concurrent, ce qui rend la question du partage propre des Skills d’autant plus stratégique pour les équipes. Gardez cependant un point capital à l’esprit : les trois surfaces (claude.ai, API, Claude Code) ne synchronisent jamais leurs Skills personnalisées entre elles.

- Une Skill téléversée sur claude.ai n’est pas automatiquement disponible via l’API.
- Une Skill téléversée via l’API n’apparaît pas sur claude.ai.
- Les Skills de Claude Code restent basées sur le système de fichiers local, complètement séparées des deux autres surfaces.

Si vous utilisez Cowork ou les sessions cloud de Claude Code (y compris les routines planifiées), sachez qu’elles ne lisent jamais le dossier `~/.claude/skills/` de votre machine locale : ces sessions chargent les Skills activées pour votre compte claude.ai au démarrage, plus les Skills de projet committées dans le dépôt cloné. Pour qu’une Skill personnelle fonctionne aussi dans une routine planifiée, il faut donc soit l’activer explicitement sur votre compte claude.ai, soit la committer dans `.claude/skills/` du dépôt, soit la livrer via un plugin déclaré dans le `.claude/settings.json` du dépôt.

## Projet complet : un Skill de description de PR prêt à l’emploi

Voici la version finale, complète et fonctionnelle de la Skill construite tout au long de ce tutoriel. Copiez-la telle quelle pour démarrer, puis adaptez les sections aux conventions de votre équipe.

### Structure des fichiers

```
~/.claude/skills/pr-description/
├── SKILL.md
├── FORMAT.md
└── scripts/
    └── check_ticket.sh
```
### Code complet

```
---
name: pr-description
description: Génère une description de pull request structurée (résumé, changements, tests, risques) à partir du diff Git en cours. Utiliser quand l'utilisateur demande une description de PR, un message de pull request, ou veut résumer ses changements avant de les pousser.
argument-hint: "[numero-ticket]"
arguments: [ticket]
allowed-tools: Bash(git diff:*) Bash(git branch:*)
---
## Instructions
À partir du diff ci-dessous, rédige une description de pull request
avec ces sections : Résumé (une phrase), Changements (liste à
puces), Tests (comment le changement a été vérifié), Risques (effets
de bord potentiels, "Aucun" si non applicable).
Si un ticket est fourni ($ticket), exécute
scripts/check_ticket.sh $ticket puis ajoute une ligne
"Closes #$ticket" à la fin si le script confirme que le ticket
existe.
Pour le gabarit détaillé attendu par l'équipe, voir [FORMAT.md](FORMAT.md).
## Contexte
Branche actuelle : !`git branch --show-current`
Fichiers modifiés : !`git diff --name-only HEAD`
Diff complet :
!`git diff HEAD`
```
Le champ `allowed-tools: Bash(git diff:*) Bash(git branch:*)` pré-autorise ces deux commandes précises pendant l’invocation de la Skill, pour éviter une confirmation manuelle à chaque test. L’autorisation s’efface automatiquement dès votre prochain message, elle ne persiste pas au-delà du tour en cours. Une fois ce fichier en place et testé selon la méthode de l’étape 5, la Skill est prête à être committée dans `.claude/skills/` à la racine d’un dépôt d’équipe.

## 5 pièges courants à éviter

