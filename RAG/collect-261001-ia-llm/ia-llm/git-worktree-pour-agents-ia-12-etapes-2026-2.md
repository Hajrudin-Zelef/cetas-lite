---
id: collect-261001-ia-llm/ia-llm/git-worktree-pour-agents-ia-12-etapes-2026-2
title: "Exemple de sortie : /Users/marie/projets/api-recettes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot"]
source: docs/RAG/collect-261001-ia-llm/git-worktree-pour-agents-ia-12-etapes-2026.md
source_anchor: ""
source_lines: [66, 156]
sha256: baaf67c34f4325c7e536aed0843664abb8bcf66d49de1b058d728c941d230806
---

# Exemple de sortie : /Users/marie/projets/api-recettes

```
Preparing worktree (new branch 'agent/fix-auth')
HEAD is now at 3f9a21c Corrige la validation des tokens
Branch 'agent/fix-auth' set up to track 'origin/main'.
```
Si vous voulez travailler sur une branche existante plutôt qu’en créer une nouvelle, omettez le drapeau `-b` et indiquez simplement le nom de la branche : `git worktree add ../api-recettes-review feature/pagination`. Git refusera si cette branche est déjà extraite ailleurs, ce qui est précisément le garde-fou qui évite les conflits silencieux entre deux agents.

## Étape 3 – Lister et auditer vos worktrees actifs

Une fois que plusieurs worktrees existent, il devient facile de perdre le fil. La commande `git worktree list` affiche tous les worktrees enregistrés, leur chemin, le commit sur lequel ils pointent et leur branche.

```
git worktree list
/Users/marie/projets/api-recettes            3f9a21c [main]
/Users/marie/projets/api-recettes-agent-claude  3f9a21c [agent/fix-auth]
/Users/marie/projets/api-recettes-review        7c1e4f0 [feature/pagination]
```
Pour un traitement automatisé (par exemple dans un script de nettoyage ou un tableau de bord interne), la variante `git worktree list --porcelain` donne une sortie stable, ligne par ligne, pensée pour être analysée par un programme plutôt que lue à l’œil. C’est la commande à retenir : lancez-la avant chaque session de travail multi-agents pour vérifier qu’aucun ancien worktree oublié ne traîne encore.

## Étape 4 – Adopter une arborescence pensée pour le multi-agent

Les noms de dossiers au hasard fonctionnent pour un test rapide, mais deviennent vite ingérables dès que trois ou quatre agents tournent en même temps. Un schéma répandu dans les guides communautaires consiste à regrouper tous les worktrees d’agents sous un sous-dossier dédié, souvent nommé `.trees`, avec un identifiant de tâche par worktree :

```
api-recettes/                  (dépôt principal, branche main)
.trees/
  fix-auth-401/                (agent Claude Code)
  pagination-v2/                (agent Cursor)
  tests-integration/            (agent Codex CLI)
```
Ce découpage a deux avantages concrets. D’abord, chaque nom de dossier documente lui-même la tâche en cours, ce qui évite d’ouvrir un worktree au hasard pour se souvenir de ce qu’il contient. Ensuite, un simple `.trees/` ajouté au `.gitignore` du dépôt principal empêche ces dossiers de dépendances imbriquées de polluer les commits si jamais quelqu’un travaille par erreur depuis la racine. Gardez le dépôt principal propre et réservez-le à la lecture, aux revues de code et aux fusions, jamais à l’édition directe pendant qu’un agent tourne ailleurs.

Le tableau suivant résume la convention à adopter selon l’agent utilisé, à titre de point de départ pour votre propre équipe.

| Agent IA | Mode de lancement | Fichier de contexte par worktree | Commande d’ouverture | 
|---|---|---|---|
| Claude Code | Terminal, dans le dossier du worktree | CLAUDE.md | cd .trees/<tache> && claude | 
| Cursor | Fenêtre IDE dédiée | .cursor/rules | cursor .trees/<tache> | 
| GitHub Copilot (VS Code) | Fenêtre IDE dédiée | copilot-instructions.md | code .trees/<tache> | 
| OpenAI Codex CLI | Terminal, dans le dossier du worktree | AGENTS.md | cd .trees/<tache> && codex | 

Cette table n’a rien de figé : adaptez les noms de dossiers et les fichiers de contexte aux conventions déjà en place dans votre équipe. Ce qui compte, c’est la cohérence, pas le respect exact de ce modèle.

## Étape 5 – Isoler une session Claude Code dans son propre worktree

Une fois le worktree créé, déplacez-vous dedans et lancez votre agent normalement, comme s’il s’agissait d’un dépôt classique.

```
cd .trees/fix-auth-401
claude
```
Claude Code démarre sa session dans ce dossier précis et n’a aucune visibilité sur ce qui se passe dans les autres worktrees. Si votre projet utilise un fichier `CLAUDE.md` pour donner des instructions persistantes à l’agent, placez-en une copie adaptée dans chaque worktree : un agent qui corrige un bug d’authentification n’a pas besoin des mêmes consignes qu’un agent qui écrit des tests d’intégration. Cette isolation des instructions, en plus de l’isolation des fichiers, rend le multi-agent réellement fiable plutôt qu’une source de confusion. Pour aller plus loin sur la configuration fine de Claude Code, notre tutoriel sur les sous-agents Claude Code détaille comment déléguer des rôles spécifiques à chaque session.

## Étape 6 – Faire tourner Cursor et GitHub Copilot en parallèle sans conflit

Le principe reste identique pour les agents intégrés à un IDE. Ouvrez chaque worktree dans sa propre fenêtre plutôt que dans un onglet du même projet : Cursor et VS Code (avec l’extension Copilot) traitent chaque fenêtre comme un espace de travail totalement séparé, ce qui correspond exactement à l’isolation offerte par le worktree.

```
cursor .trees/pagination-v2
code .trees/tests-integration
```
Un piège classique ici : ouvrir les deux worktrees comme deux dossiers d’un même espace de travail multi-racine. Techniquement possible, cette approche mélange les vues de fichiers à l’écran et augmente le risque qu’un agent modifie le mauvais dossier par inadvertance. Une fenêtre par worktree, sans exception, reste la règle la plus sûre pour un usage sérieux avec plusieurs agents actifs.

## Étape 7 – Ajouter OpenAI Codex CLI à la rotation d’agents

Le même schéma s’applique à un agent en ligne de commande comme Codex CLI. Depuis son propre worktree, lancez simplement l’outil :

```
cd .trees/tests-integration
codex
```
Codex CLI, comme Claude Code, fonctionne en boucle lecture-édition-vérification directement sur les fichiers du dossier courant. Comme ce dossier est un worktree indépendant, l’agent peut exécuter ses propres commandes de test, installer ses propres dépendances et même faire planter son serveur de développement local sans jamais toucher aux deux autres sessions en cours. Notre présentation complète de Codex CLI couvre l’installation et l’authentification si ce n’est pas encore fait.

À ce stade, vous avez trois agents actifs sur trois branches différentes du même dépôt, dans trois dossiers physiquement distincts. C’est le cœur du workflow. Les étapes suivantes traitent des détails pratiques qui déterminent si cette configuration tient sur la durée ou s’effondre au bout de deux jours.

## Étape 8 – Gérer les fichiers .env, les ports et les bases de données par worktree

Un fichier `.env` n’est jamais suivi par Git, donc il n’existe pas automatiquement dans un nouveau worktree. Sans lui, la plupart des projets refusent de démarrer. Copiez-le manuellement, ou via un petit script, à chaque création de worktree :

`cp ../api-recettes/.env .trees/fix-auth-401/.env`
Le vrai piège arrive ensuite : si les trois agents lancent chacun un serveur de développement sur le port 3000, deux d’entre eux échoueront au démarrage. Attribuez un port distinct par worktree, soit directement dans chaque `.env`, soit via une variable d’environnement passée au lancement :

```
PORT=3001 npm run dev    # worktree fix-auth-401
PORT=3002 npm run dev    # worktree pagination-v2
PORT=3003 npm run dev    # worktree tests-integration
```
Pour les projets avec base de données, le même raisonnement s’applique : soit une base SQLite distincte par worktree (le plus simple), soit un schéma ou une base PostgreSQL séparée par tâche si votre projet dépend d’un serveur partagé. Ignorer ce point est la cause numéro un des échecs silencieux quand un agent « ne voit pas » les changements d’un autre : les deux ne travaillent tout simplement pas sur les mêmes données.

## Étape 9 – node_modules, environnements virtuels et l’explosion du disque

