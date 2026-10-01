---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-4
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "arr", "attention", "claude", "cost"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [253, 379]
sha256: 21598ee451762a92b6fce1ef767b2d7021c06db9e6ddb798247216431a97b2be
---

# ./scripts/validate-readonly-query.sh

```
Utilise le sous-agent test-runner pour corriger les tests en échec
Demande au sous-agent code-reviewer de regarder mes changements récents
```
L’@-mention garantit qu’un sous-agent précis s’exécute, plutôt que de laisser Claude choisir : tapez `@` puis sélectionnez le sous-agent dans la liste, exactement comme pour mentionner un fichier. Votre message complet part quand même vers Claude, qui rédige le prompt de délégation à partir de ce que vous avez demandé : l’@-mention contrôle uniquement quel sous-agent est invoqué, pas le contenu du prompt qu’il reçoit.

Pour faire tourner toute une session comme un sous-agent donné, avec son prompt système, ses restrictions d’outils et son modèle, utilisez l’option `–agent` :

`claude --agent code-reviewer`
Pour en faire le comportement par défaut de tout un projet, ajoutez `agent` dans `.claude/settings.json` :

```
{
  "agent": "code-reviewer"
}
```
Le drapeau passé en ligne de commande prend le dessus si les deux sont présents en même temps. Ce mode fonctionne aussi bien avec les sous-agents intégrés qu’avec vos sous-agents personnalisés, et le choix persiste quand vous reprenez une session interrompue.

## Étape 10 : exécuter en arrière-plan et enchaîner plusieurs sous-agents

Depuis la version 2.1.198, les sous-agents tournent en arrière-plan par défaut : côté SDK, ce changement se traduit par un paramètre `run_in_background` qui, laissé vide, bascule désormais l’exécution du sous-agent en arrière-plan sans intervention explicite. Claude ne les fait passer au premier plan que lorsqu’il a besoin du résultat immédiatement pour continuer, sauf si le frontmatter impose ce comportement via le champ booléen `background: true`, qui force l’arrière-plan quelle que soit la décision de Claude. Un sous-agent en arrière-plan garde un jeu d’outils plus restreint qu’un sous-agent au premier plan, et toute invite de permission qu’il rencontre remonte dans votre session principale, avec le nom du sous-agent concerné. Approuvez pour le laisser continuer, ou appuyez sur Échap pour refuser cet appel précis sans arrêter le sous-agent.

Pour des recherches indépendantes, plusieurs sous-agents peuvent tourner simultanément :

```
Recherche les modules d'authentification, de base de données et d'API
en parallèle avec des sous-agents séparés
```
Chaque sous-agent explore sa zone indépendamment, puis Claude synthétise les résultats. Attention toutefois : quand plusieurs sous-agents renvoient chacun un résultat détaillé, la conversation principale peut se remplir plus vite que prévu. Pour un travail multi-étapes, vous pouvez aussi enchaîner les sous-agents en séquence, chacun transmettant le contexte pertinent au suivant :

```
Utilise le sous-agent code-reviewer pour trouver les problèmes de
performance, puis utilise le sous-agent optimizer pour les corriger
```
Par défaut, un sous-agent ne peut pas lui-même en faire naître d’autres : Claude Code lui retire l’outil Agent tant que l’imbrication n’est pas activée. Pour l’autoriser, réglez la variable CLAUDE_CODE_MAX_SUBAGENT_SPAWN_DEPTH sur le nombre de niveaux souhaités, par exemple 2, dans `settings.json`. Deux limites protègent aussi la session : 200 sous-agents maximum au total par défaut (ajustable via CLAUDE_CODE_MAX_SUBAGENTS_PER_SESSION), et 20 sous-agents simultanés maximum par défaut (ajustable via CLAUDE_CODE_MAX_CONCURRENT_SUBAGENTS) — un plafond fixé par la version 2.1.217 de juillet 2026 selon Digital Applied, la même mise à jour qui a mis fin, depuis le 21 juillet 2026, à la possibilité par défaut pour un sous-agent d’en faire naître d’autres sans configuration explicite de la profondeur d’imbrication.

Un sous-agent terminé n’est pas forcément fini pour de bon. Demandez à Claude de continuer un travail déjà livré (« Continue cette revue de code et analyse maintenant la logique d’autorisation ») et Claude Code reprend le même sous-agent avec tout son historique de conversation, plutôt que de repartir de zéro. Cette reprise passe par l’outil SendMessage, avec l’identifiant ou le nom du sous-agent comme destinataire. Seuls les sous-agents intégrés Explore et Plan ne renvoient aucun identifiant réutilisable : ce sont des agents à usage unique, donc préférez general-purpose ou un sous-agent personnalisé dès que la tâche demande plusieurs allers-retours.

## Projet complet : une équipe de sous-agents pour la revue de code

Voici un projet fonctionnel qui assemble trois sous-agents complémentaires : un relecteur en lecture seule, un débogueur qui corrige, et un analyste de données pour les requêtes SQL. Créez la structure suivante à la racine de votre projet :

```
mon-projet/
├── .claude/
│   └── agents/
│       ├── code-reviewer.md
│       ├── debugger.md
│       └── data-scientist.md
└── (le reste de votre code)
```
### 1. Le relecteur, en lecture seule

Ce sous-agent ne modifie jamais rien : ses outils s’arrêtent à Read, Grep, Glob et Bash, ce qui suffit pour lancer `git diff` et lire les fichiers changés.

```
---
name: code-reviewer
description: Expert code review specialist. Proactively reviews code for quality, security, and maintainability. Use immediately after writing or modifying code.
tools: Read, Grep, Glob, Bash
model: inherit
---
You are a senior code reviewer ensuring high standards of code quality and security.
When invoked:
1. Run git diff to see recent changes
2. Focus on modified files
3. Begin review immediately
Review checklist:
- Code is clear and readable
- Functions and variables are well-named
- No duplicated code
- Proper error handling
- No exposed secrets or API keys
- Input validation implemented
- Good test coverage
- Performance considerations addressed
Provide feedback organized by priority:
- Critical issues (must fix)
- Warnings (should fix)
- Suggestions (consider improving)
```
### 2. Le débogueur, avec droit de correction

À la différence du relecteur, ce sous-agent inclut Edit, puisque corriger un bug suppose de modifier le code :

```
---
name: debugger
description: Debugging specialist for errors, test failures, and unexpected behavior. Use proactively when encountering any issues.
tools: Read, Edit, Bash, Grep, Glob
---
You are an expert debugger specializing in root cause analysis.
When invoked:
1. Capture error message and stack trace
2. Identify reproduction steps
3. Isolate the failure location
4. Implement minimal fix
5. Verify solution works
For each issue, provide:
- Root cause explanation
- Evidence supporting the diagnosis
- Specific code fix
- Testing approach
- Prevention recommendations
```
### 3. L’analyste de données, pour les requêtes SQL

Ce troisième sous-agent illustre un cas d’usage hors codage pur, avec un modèle fixé explicitement sur Sonnet plutôt que sur `inherit` :

```
---
name: data-scientist
description: Data analysis expert for SQL queries, BigQuery operations, and data insights. Use proactively for data analysis tasks and queries.
tools: Bash, Read, Write
model: sonnet
---
You are a data scientist specializing in SQL and BigQuery analysis.
When invoked:
1. Understand the data analysis requirement
2. Write efficient SQL queries
3. Use BigQuery command line tools (bq) when appropriate
4. Analyze and summarize results
5. Present findings clearly
Always ensure queries are efficient and cost-effective.
```
Une fois les trois fichiers en place, redémarrez Claude Code pour que le dossier `.claude/agents/` fraîchement créé soit détecté. Enchaînez ensuite naturellement : « Utilise le sous-agent code-reviewer sur mes changements récents, puis passe les problèmes trouvés au sous-agent debugger pour qu’il les corrige ». Committez le dossier `.claude/agents/` dans votre dépôt Git : toute l’équipe récupère la même configuration au prochain `git pull`.

