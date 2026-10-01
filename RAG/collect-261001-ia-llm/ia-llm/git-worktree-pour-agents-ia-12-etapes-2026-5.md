---
id: collect-261001-ia-llm/ia-llm/git-worktree-pour-agents-ia-12-etapes-2026-5
title: "Exemple de sortie : /Users/marie/projets/api-recettes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot"]
source: docs/RAG/collect-261001-ia-llm/git-worktree-pour-agents-ia-12-etapes-2026.md
source_anchor: ""
source_lines: [359, 399]
sha256: 85a633954a8fa85908f42dc863aa2602e030dfa23487b494c9cdf43eaaf10882
---

# Exemple de sortie : /Users/marie/projets/api-recettes

Pour ceux qui travaillent en SSH sur un serveur distant plutôt qu’en local, combiner git worktree avec un multiplexeur de terminal comme tmux ou screen apporte un vrai confort. Chaque worktree correspond alors à une fenêtre tmux nommée d’après la tâche en cours, ce qui permet de détacher la session, de fermer son ordinateur portable, et de la retrouver intacte le lendemain avec l’agent toujours actif au même endroit. C’est particulièrement utile pour les tâches longues, comme une migration de base de données ou une refactorisation étendue, que vous ne voulez pas interrompre en fermant votre terminal par erreur.

## Foire aux questions

### Git worktree ralentit-il mon dépôt ?

Non. Les opérations Git courantes (commit, log, diff) restent aussi rapides que d’habitude puisque la base d’objets reste unique. Seul l’espace disque augmente, à cause des dépendances réinstallées dans chaque worktree, pas à cause de l’historique Git lui-même.

### Combien de worktrees puis-je créer en même temps ?

Il n’existe pas de limite technique fixée par Git. En pratique, la limite vient de votre espace disque et du nombre de sessions d’agents que vous pouvez réellement suivre. Trois à cinq worktrees actifs représentent déjà un maximum raisonnable pour un développeur seul.

### Puis-je utiliser git worktree avec GitHub Desktop ou un autre client graphique ?

La plupart des clients graphiques grand public ne gèrent pas nativement la création de worktrees, même s’ils affichent correctement les dossiers une fois créés en ligne de commande. Le plus simple reste de créer et gérer les worktrees via le terminal, puis d’ouvrir chaque dossier normalement dans votre client ou IDE préféré.

### Un worktree peut-il pointer vers un dépôt distant différent ?

Non. Tous les worktrees d’un même groupe partagent obligatoirement la même configuration de dépôt distant (remote), puisqu’ils partagent la même base d’objets Git. Pour travailler avec un remote totalement différent, un clone séparé reste la bonne solution.

### Que se passe-t-il si je supprime le dépôt principal par erreur ?

Les worktrees secondaires deviennent inutilisables, car ils dépendent tous du dossier .git central du dépôt principal. Ne supprimez jamais le dépôt d’origine tant que des worktrees actifs en dépendent, et gardez toujours une copie poussée sur votre remote (GitHub, GitLab) comme filet de sécurité.

### Git worktree fonctionne-t-il aussi bien sous Windows que sous macOS et Linux ?

Oui, la commande est identique sur les trois systèmes. Sous Windows, l’utilisation via WSL2 reste recommandée pour les workflows avec agents de codage IA, car la plupart des outils (Claude Code, Codex CLI) sont d’abord pensés pour un environnement Unix.

### Faut-il fusionner les branches d’agents une par une, ou toutes en même temps ?

Une par une, dans l’ordre où elles sont validées. Fusionner plusieurs branches d’agents simultanément multiplie le risque de conflits difficiles à attribuer à une tâche précise, alors qu’une fusion séquentielle garde un historique clair et un point de retour simple en cas de problème.

### Peut-on utiliser git worktree sur un très gros monorepo ?

Oui, et c’est même l’un des cas d’usage où le gain est le plus net. Sur un monorepo volumineux, cloner une copie complète pour chaque tâche coûte cher en temps et en espace disque. Le worktree évite ce coût puisque l’historique n’est jamais dupliqué. Le seul point de vigilance reste la réinstallation des dépendances par worktree, qui peut prendre plus de temps que sur un petit projet si le monorepo contient de nombreux paquets internes.

### Related Coverage

## Sources et documentation officielle

Pour approfondir chaque commande présentée dans ce tutoriel, la documentation officielle de Git reste la référence à jour : git-scm.com/docs/git-worktree. Les intégrations avec les agents IA mentionnées s’appuient sur la documentation de Claude Code, sur celle de GitHub Copilot, sur le dépôt officiel d’OpenAI Codex CLI et sur la documentation du contrôle de source de Visual Studio Code.
