---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-6
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "mcp"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [427, 465]
sha256: dcb94083c938e51963f608b3cf6fdc316caf0f7c82147453012474013094f3fd
---

# ./scripts/validate-readonly-query.sh

Enfin, pensez à combiner `skills` et sous-agents plutôt que de choisir l’un ou l’autre par défaut. Un champ `skills` précharge le contenu complet d’un skill dans le contexte du sous-agent au démarrage, ce qui lui donne une connaissance métier ciblée (conventions d’API, patterns de gestion d’erreurs) sans qu’il ait à la découvrir lui-même pendant l’exécution.

## Questions fréquentes

**Un sous-agent Claude Code coûte-t-il plus cher qu’une conversation classique ?**

Pas nécessairement. Un sous-agent qui tourne sur Haiku coûte moins cher par token qu’une conversation principale sur Sonnet ou Opus. Le vrai risque de coût vient du nombre de sous-agents lancés en parallèle plutôt que du principe lui-même.

**Faut-il un dépôt Git pour utiliser des sous-agents de projet ?**

Non. Un dossier vide fonctionne. Un dépôt Git devient utile surtout pour versionner `.claude/agents/` et le partager avec une équipe.

**Quelle est la différence entre un sous-agent et un fork de conversation ?**

Un sous-agent démarre avec un contexte neuf et son propre prompt système. Un fork, lancé avec `/subtask`, hérite au contraire de tout l’historique de la conversation en cours, ce qui le rend utile pour une tâche annexe qui demanderait sinon trop de contexte à réexpliquer.

**Peut-on utiliser des sous-agents en dehors du terminal Claude Code ?**

Oui. Les définitions de sous-agents restent valables via le SDK Agent, ce qui permet de lancer Claude Code de façon programmatique dans un pipeline CI/CD, en dehors de toute session interactive.

**Un sous-agent personnalisé peut-il remplacer Explore ou Plan ?**

Oui, en le nommant explicitement `Explore` ou `Plan` dans un sous-agent de projet ou d’utilisateur : cette définition prend le dessus sur le sous-agent intégré du même nom, en conservant tout de même son propre champ `model`.

**Comment partager mes sous-agents avec toute mon équipe sans passer par –add-dir ?**

Deux options : les placer dans `.claude/agents/` et les committer dans le dépôt, ou les distribuer via un plugin, qui les charge automatiquement pour tous les utilisateurs de ce plugin.

**Que se passe-t-il si mon organisation restreint les modèles disponibles ?**

Claude Code vérifie la variable d’environnement, le paramètre par invocation et la valeur du frontmatter contre la liste blanche availableModels de l’organisation. Une valeur qui pointe vers un modèle exclu est ignorée, et le sous-agent tourne sur le modèle hérité à la place, sans erreur bloquante.

**Les sous-agents fonctionnent-ils de la même façon sur Windows ?**

La syntaxe `claude mcp`, `claude –agent` et les frontmatters YAML sont identiques sur macOS, Linux, WSL et PowerShell. Seuls les scripts de hooks changent : écrivez-les en PowerShell et ajoutez `shell: powershell` à l’entrée du hook sur Windows.

**Peut-on reprendre un sous-agent après qu’il a terminé sa tâche ?**

Oui, tant qu’il ne s’agit pas d’Explore ou de Plan. Demandez simplement à Claude de continuer le travail : le sous-agent reprend avec son historique complet, y compris tous les appels d’outils et résultats précédents, au lieu de repartir d’un contexte vide.
