---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-3
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: []
keywords: ["agent", "agents", "arr", "claude", "mcp", "memory", "research", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [140, 252]
sha256: 210debb4c7a91ceb6ea09855d3bccdba3a6ea4ff03bf442afcab6bddb9f113a8
---

# ./scripts/validate-readonly-query.sh

Le champ `tools` fonctionne comme une liste blanche, `disallowedTools` comme une liste noire retirée du pool hérité. Si les deux sont présents, `disallowedTools` s’applique d’abord, puis `tools` est résolu contre ce qui reste. Un exemple qui n’autorise que la lecture et la recherche, sans aucun accès en écriture ni aux outils MCP :

```
---
name: safe-researcher
description: Research agent with restricted capabilities
tools: Read, Grep, Glob, Bash
---
```
Et l’inverse, qui hérite de tout sauf de l’écriture et de la modification de fichiers :

```
---
name: no-writes
description: Inherits the available tools except file writes
disallowedTools: Write, Edit
---
```
Ces deux champs acceptent aussi des motifs au niveau des serveurs MCP : `mcp__

Le champ `mcpServers` va plus loin : il donne à un sous-agent l’accès à des serveurs MCP absents de la conversation principale, ce qui évite que leurs descriptions d’outils consomment du contexte ailleurs que là où elles servent. Reprenons l’exemple officiel d’un testeur de navigateur basé sur Playwright, déjà couvert par notre tutoriel sur les serveurs MCP pour Claude Code :

```
---
name: browser-tester
description: Tests features in a real browser using Playwright
mcpServers:
  - playwright:
      type: stdio
      command: npx
      args: ["-y", "@playwright/mcp@latest"]
  - github
---
Use the Playwright tools to navigate, screenshot, and interact with pages.
```
Le serveur Playwright MCP de Microsoft nécessite Node.js 18 ou plus récent, puisqu’il se lance via `npx`. Notez la différence de syntaxe : une définition en ligne comme `playwright` reste isolée à ce sous-agent, tandis qu’une simple chaîne comme `github` réutilise une connexion déjà configurée au niveau de la session.

## Étape 6 : choisir le modèle IA et maîtriser les coûts

Le champ `model` accepte un alias (`sonnet`, `opus`, `haiku`, `fable`), un identifiant complet comme `claude-opus-4-8`, ou `inherit`, valeur par défaut si le champ est omis. C’est l’un des leviers de coût les plus concrets pour une équipe qui multiplie les sous-agents : depuis juillet 2026, Sonnet 5 est devenu le modèle par défaut des sièges Pro avec, lui aussi, une fenêtre de contexte d’un million de tokens selon Anthropic, tandis qu’un agent de recherche de fichiers simple peut continuer à tourner sur Haiku, nettement moins cher par token.

Claude Code résout le modèle effectif d’un sous-agent dans cet ordre de priorité, du plus fort au plus faible :

1. La variable d’environnement CLAUDE_CODE_SUBAGENT_MODEL, quand elle est réglée sur un alias ou un ID de modèle
2. Le paramètre model passé pour cette invocation précise
3. Le champ model défini dans le frontmatter du sous-agent
4. Le modèle de la conversation principale

Depuis la version 2.1.196, régler CLAUDE_CODE_SUBAGENT_MODEL sur `inherit` équivaut à ne pas la définir du tout : la résolution continue avec le paramètre d’invocation puis le frontmatter. Sur les versions antérieures, `inherit` forçait tous les sous-agents sur le modèle de la conversation principale et ignorait les deux autres sources. Si votre organisation restreint les modèles disponibles via une liste blanche, Claude Code ignore silencieusement une valeur qui pointe vers un modèle exclu et fait tourner le sous-agent sur le modèle hérité à la place.

Pour comparer les tarifs exacts entre Sonnet, Opus et les modèles concurrents au moment de dimensionner votre budget de sous-agents, notez que la tarification API de Sonnet 5 était fixée à 2 $ / 10 $ par million de tokens (entrée/sortie) jusqu’au 31 août 2026 selon Anthropic ; notre comparatif Claude Code vs Cursor 2026 détaille les écarts de coût par requête entre les deux outils.

## Étape 7 : ajouter des hooks pour un contrôle conditionnel

La liste `tools` reste binaire : un outil est autorisé ou non. Pour un contrôle plus fin, comme autoriser Bash mais bloquer uniquement les commandes SQL d’écriture, les hooks `PreToolUse` interceptent chaque appel d’outil avant son exécution. Voici l’exemple officiel d’un sous-agent `db-reader`, limité aux requêtes de lecture :

```
---
name: db-reader
description: Execute read-only database queries. Use when analyzing data or generating reports.
tools: Bash
hooks:
  PreToolUse:
    - matcher: "Bash"
      hooks:
        - type: command
          command: "./scripts/validate-readonly-query.sh"
---
You are a database analyst with read-only access. Execute SELECT queries
to answer questions about the data.
```
Claude Code transmet l’entrée du hook en JSON via stdin. Le script lit ce JSON, extrait la commande Bash, et sort avec le code 2 pour bloquer une opération d’écriture :

```
#!/bin/bash
# ./scripts/validate-readonly-query.sh
INPUT=$(cat)
COMMAND=$(echo "$INPUT" | jq -r '.tool_input.command // empty')
if echo "$COMMAND" | grep -iE '\b(INSERT|UPDATE|DELETE|DROP|CREATE|ALTER|TRUNCATE)\b' > /dev/null; then
  echo "Blocked: Only SELECT queries are allowed" >&2
  exit 2
fi
exit 0
```
Rendez le script exécutable avec `chmod +x ./scripts/validate-readonly-query.sh` sur macOS et Linux. Sur Windows, écrivez le script en PowerShell et ajoutez `shell: powershell` à l’entrée du hook. Les événements les plus utiles pour un sous-agent sont `PreToolUse` (avant l’appel), `PostToolUse` (après l’appel) et `Stop`, converti automatiquement en `SubagentStop` quand l’agent tourne comme sous-agent plutôt que comme session principale.

Ces hooks peuvent aussi être déclarés au niveau du projet, dans `settings.json`, pour réagir au démarrage ou à l’arrêt de n’importe quel sous-agent via les événements `SubagentStart` et `SubagentStop`, avec un filtre par nom d’agent.

## Étape 8 : donner une mémoire persistante à vos sous-agents

Le champ `memory` attribue à un sous-agent un dossier qui survit d’une conversation à l’autre. L’agent s’en sert pour accumuler des connaissances : schémas de base de données récurrents, conventions de nommage, décisions d’architecture déjà expliquées une fois. Trois portées sont possibles.

- **user** — stocké dans ~/.claude/agent-memory/nom-de-l-agent/, pour un savoir valable sur tous vos projets
- **project** — stocké dans .claude/agent-memory/nom-de-l-agent/, partageable via le contrôle de version, portée recommandée par défaut
- **local** — stocké dans .claude/agent-memory-local/nom-de-l-agent/, propre à un projet mais volontairement exclu du dépôt

Une fois la mémoire activée, le prompt système du sous-agent inclut automatiquement les instructions de lecture et d’écriture de son dossier de mémoire, ainsi que les 200 premières lignes ou 25 Ko de son fichier MEMORY.md, selon la limite atteinte en premier. Les outils Read, Write et Edit sont alors activés automatiquement pour que le sous-agent gère lui-même ses fichiers de mémoire :

```
---
name: code-reviewer
description: Reviews code for quality and best practices
memory: user
---
You are a code reviewer. As you review code, update your agent memory with
patterns, conventions, and recurring issues you discover.
```
Demandez explicitement au sous-agent de consulter puis de mettre à jour sa mémoire : « Relis ce fichier de mémoire avant de commencer » en début de tâche, puis « Maintenant que tu as terminé, note ce que tu as appris dans ta mémoire » à la fin. Cette mémoire fait partie du système de mémoire automatique de Claude Code : si vous la désactivez globalement, le champ `memory` d’un sous-agent perd tout effet.

## Étape 9 : invoquer vos sous-agents (langage naturel, @-mention, –agent)

Trois méthodes existent, de la simple suggestion ponctuelle à l’usage par défaut de toute la session. En langage naturel, il suffit de nommer le sous-agent dans votre message, sans syntaxe particulière :

