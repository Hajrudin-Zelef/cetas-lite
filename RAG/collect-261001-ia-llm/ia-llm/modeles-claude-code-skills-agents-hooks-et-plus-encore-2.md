---
id: collect-261001-ia-llm/ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore-2
title: "Database Migration Skill"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "mcp", "model context protocol", "safeguards"]
source: docs/RAG/collect-261001-ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [117, 277]
sha256: 9461f7bb44cb568f71a799e29f2a4a7d9a0b18aefc2318780164d6e04597f20a
---

# Database Migration Skill

Ils peuvent résider dans `.claude/agents/` pour une portée projet, ou `~/.claude/agents/` pour une portée personnelle. Les agents sont créés en les demandant à Claude ou en éditant directement les fichiers Markdown du dossier `.claude/agents/`.

Il existe une vraie différence entre skills et agents. Un skill définit comment exécuter une tâche. Un agent définit qui Claude doit être pendant le travail : son rôle, son focus, ses permissions et ses limites.

Voyons un exemple d’agent :

```
---
name: security-auditor
description: Reviews code for security vulnerabilities and produces a findings report without modifying files.
tools: Read, Glob, Grep, Bash
model: sonnet
---
You are a security auditor.
Your task is to inspect the codebase for vulnerabilities, risky patterns, and missing safeguards.
Rules:
- Do not edit files.
- Do not suggest broad rewrites unless directly tied to a security issue.
- Focus on authentication, authorization, input validation, secrets, dependency risk, and unsafe shell or SQL usage.
- Produce a findings report with severity, affected files, evidence, and recommended next steps.
```
Cet agent est utile car il fixe des limites claires. En session générale, Claude pourrait se mettre à corriger les problèmes dès qu’il en trouve. Un agent auditeur sécurité est instrué d’inspecter et de rapporter uniquement, sans modifications.

Les agents sont idéaux pour des domaines spécialisés comme l’audit sécurité, la revue documentaire, la revue d’architecture, l’ingénierie des données ou les contrôles de qualité du code, où l’isolement du contexte et des permissions est crucial.

Ils sont également efficaces en tandem avec des skills spécialisés. Par exemple, un agent auditeur sécurité peut appeler un skill de rapport de constats, tandis qu’un agent relecteur frontend peut utiliser un skill de test de composants.

### 3. Commandes

Les commandes sont des raccourcis invoqués en slash, comme `/generate-tests`, `/check-deps` ou `/summarize-pr`. Historiquement, les commandes personnalisées étaient stockées comme fichiers Markdown sous `.claude/commands/`, le nom du fichier servant de nom de commande.

Claude Code prend toujours en charge cet ancien format, mais nous recommandons d’utiliser des skills pour les nouveaux workflows de type commande : ils gèrent le même appel `/name` et peuvent être déclenchés automatiquement lorsque pertinent.

Les commandes sont préférables lorsque vous souhaitez un déclenchement explicite. Un skill peut se lancer automatiquement si Claude détecte une tâche correspondante, mais une commande ne doit s’exécuter que lorsque vous la lancez. Elles sont donc idéales comme points de contrôle : « générer les tests maintenant », « résumer cette PR maintenant », « vérifier les dépendances maintenant », « préparer un message de commit maintenant ».

### 4. Hooks

Les hooks sont des règles d’automatisation exécutées en réponse à des événements du cycle de vie de Claude Code. Ce sont des commandes shell définies par l’utilisateur, exécutées à des moments précis, qui offrent un contrôle déterministe du comportement.

La différence avec les autres modèles précédemment abordés est qu’ils ne sont pas déclenchés par ce que vous demandez, mais par ce que fait Claude.

Concrètement, vous n’avez pas à espérer que Claude pense à formater un fichier après l’avoir édité ; un hook peut le faire automatiquement.

Les événements de hook actuels incluent notamment `PreToolUse`, `PostToolUse`, `Notification` et `Stop`.

Exemple : lancer un formateur après que Claude a édité ou écrit un fichier :

```
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Edit|Write",
        "hooks": [
          {
            "type": "command",
            "command": "jq -r '.tool_input.file_path' | xargs npx prettier --write"
          }
        ]
      }
    ]
  }
}
Example: block risky shell commands before Claude runs them:
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Bash",
        "hooks": [
          {
            "type": "command",
            "command": "python3 .claude/hooks/block-dangerous-bash.py"
          }
        ]
      }
    ]
  }
}
```
Les hooks sont idéaux pour les règles que Claude ne doit pas pouvoir contourner : exécuter un linter, formater les fichiers modifiés, bloquer l’édition de fichiers protégés, vérifier le code généré ou envoyer des notifications lorsque Claude a besoin d’entrées.

Pour un tutoriel approfondi, lisez notre guide des hooks Claude Code.

### 5. Intégrations MCP

Les intégrations MCP connectent Claude Code à des outils, sources de données et API externes via le Model Context Protocol. MCP joue le rôle de couche connecteur entre systèmes d’IA et outils externes. Dans Claude Code, cela permet à Claude d’aller au-delà des fichiers locaux et des commandes shell.

Claude peut ainsi interagir avec des services externes comme GitHub, des bases de données, des systèmes de documentation, des plateformes cloud ou des API internes, selon les serveurs MCP que vous configurez. Pour une explication complète et un projet de démo, consultez notre tutoriel Model Context Protocol.

Un serveur MCP peut exposer trois grands types de capacités :

- **Tools** : fonctions exécutables appelables par Claude, comme créer un ticket GitHub ou exécuter une requête de base de données.
- **Resources** : sources de contexte en lecture seule, comme un fichier, une ligne de base ou un document.
- **Prompts** : modèles de tâches réutilisables exposés par le serveur.

Un `.mcp.json` au niveau projet peut configurer plusieurs serveurs côte à côte :

```
{
  "mcpServers": {
    "github": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${GITHUB_TOKEN}"
      }
    },
    "sqlite": {
      "type": "stdio",
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-sqlite",
        "./data/app.db"
      ]
    }
  }
}
```
C’est important car Claude ne peut raisonner qu’à partir du contexte et des outils auxquels il a accès. Sans MCP, il peut inspecter des fichiers locaux, mais pas votre gestionnaire de tickets, votre base de données, votre environnement cloud ou vos API internes.

MCP est idéal lorsque Claude doit travailler avec votre stack réelle plutôt qu’un instantané de code statique, et lorsqu’il a besoin d’accès à des données externes.

### 6. Plugins

Les plugins sont des bundles empaquetés. Ils peuvent inclure des skills, des agents, des hooks, des configurations MCP, des commandes et d’autres composants au sein d’une seule structure installable.

Dans Claude Code, un plugin comprend généralement un manifest `.claude-plugin/plugin.json` et des dossiers de composants tels que `skills/`, `agents/`, `hooks/` et `.mcp.json` à la racine du plugin.

Exemple de structure de plugin :

```
frontend-workflow-plugin/
├── .claude-plugin/
│   └── plugin.json
├── skills/
│   └── component-test/
│       └── SKILL.md
├── agents/
│   └── frontend-reviewer.md
├── hooks/
│   └── hooks.json
└── .mcp.json
```
Exemple de plugin json :

```
{
  "name": "frontend-workflow",
  "displayName": "Frontend Workflow",
  "version": "1.0.0",
  "description": "Frontend development workflow with review agents, test skills, and formatting hooks",
  "author": {
    "name": "Your Team"
  }
}
```
Les plugins n’ajoutent pas un nouveau type de comportement ; ils rendent les autres types portables. Utilisez-les pour partager une configuration complète au sein d’une équipe, réutiliser le même workflow sur plusieurs projets, ou installer un bundle communautaire plutôt que de créer chaque fichier manuellement.

