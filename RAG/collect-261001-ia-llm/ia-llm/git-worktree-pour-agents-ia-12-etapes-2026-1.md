---
id: collect-261001-ia-llm/ia-llm/git-worktree-pour-agents-ia-12-etapes-2026-1
title: "Exemple de sortie : /Users/marie/projets/api-recettes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot"]
source: docs/RAG/collect-261001-ia-llm/git-worktree-pour-agents-ia-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 65]
sha256: 03833a5388e7a0d00419632f0c47e7eb89e5b19d0de617ca180a7f072092a261
---

# Exemple de sortie : /Users/marie/projets/api-recettes

Il y a encore deux ans, un développeur lançait un seul assistant IA à la fois et attendait sa réponse avant de passer à autre chose. En 2026, ce modèle a explosé. Claude Code corrige un bug pendant que Cursor écrit une fonctionnalité et que GitHub Copilot génère des tests, le tout sur le même dépôt, en même temps. Le problème, c’est que Git n’a jamais été conçu pour qu’un seul dossier de travail accueille trois sessions concurrentes sans qu’elles se marchent dessus.

La réponse ne vient pas d’un nouvel outil IA, mais d’une commande Git vieille de plus de dix ans : **git worktree**. Elle permet de créer plusieurs dossiers de travail reliés au même dépôt, chacun sur sa propre branche, sans dupliquer l’historique et sans les conflits classiques du changement de branche. Ce tutoriel montre comment l’installer, la structurer et l’utiliser concrètement pour faire tourner plusieurs agents de codage IA en parallèle, avec un projet complet à la fin.

Douze étapes composent ce guide, de la vérification de votre version de Git jusqu’au nettoyage final des dossiers devenus inutiles, en passant par un script d’automatisation et un exemple de bout en bout sur un vrai projet Node.js. Comptez environ 45 minutes pour tout suivre, y compris les tests. Vous repartirez avec une structure de dossiers réutilisable sur n’importe quel projet, quel que soit l’agent IA que vous utilisez au quotidien.

## Pourquoi git worktree s’impose dans les workflows d’agents IA en 2026

Le scénario est devenu banal dans les équipes qui utilisent Claude Code, Cursor ou GitHub Copilot au quotidien : vous lancez un agent sur une tâche, puis vous voulez en démarrer un second sur un autre ticket sans attendre que le premier termine. Avec un dépôt classique, il faut choisir entre attendre, ou changer de branche et perdre le contexte du premier agent (fichiers modifiés, processus lancés, serveur de développement actif). Aucune des deux options ne convient à un usage sérieux.

Git worktree règle ce problème directement à la racine, sans passer par des clones séparés ni des scripts maison. Chaque worktree est un dossier indépendant avec son propre état de fichiers, sa propre zone d’index et sa propre branche extraite, tout en partageant la même base d’objets Git (commits, arbres, blobs). Résultat : trois agents peuvent travailler sur trois branches différentes du même projet, dans trois dossiers distincts, sans jamais se gêner.

Ce n’est pas une fonctionnalité expérimentale. Elle existe depuis Git 2.5, sorti en 2015, bien avant l’essor des agents de codage. Ce qui change en 2026, c’est l’usage : la documentation de Claude Code recommande d’isoler chaque session dans son propre worktree pour éviter les interférences entre tâches, et la pratique s’est généralisée à Cursor, Copilot et OpenAI Codex CLI par la force des choses. Ce guide part de zéro et construit, étape par étape, un vrai poste de travail multi-agents.

## Prérequis : versions, outils et comptes nécessaires

Avant de commencer, réunissez les éléments suivants. Rien d’exotique : l’essentiel tient dans un terminal et une installation Git à jour.

- **Git 2.40 ou supérieur** (git worktree fonctionne dès Git 2.5, mais les sous-commandes move et repair sont plus fiables sur les versions récentes)
- Un terminal (macOS Terminal, Windows Terminal avec WSL2, ou tout terminal Linux)
- Un compte GitHub ou GitLab avec un dépôt existant, ou un projet local à initialiser
- Au moins un agent de codage IA installé : Claude Code CLI, Cursor, GitHub Copilot ou OpenAI Codex CLI
- Node.js 20 ou supérieur, ou Python 3.11 ou supérieur, selon le langage de votre projet (utilisé pour l’exemple de ce tutoriel)
- Environ 2 Go d’espace disque libre par worktree actif si votre projet a des dépendances lourdes
- Des notions de base sur les branches Git (checkout, merge, pull) : ce tutoriel ne les réexplique pas

Vérifiez votre version avec `git --version`. Sous macOS, une mise à jour passe par `brew upgrade git`. Sous Ubuntu ou Debian, `sudo apt update && sudo apt install git` suffit généralement. Sous Windows, réinstallez la dernière version depuis git-scm.com ou passez par WSL2, largement recommandé pour les agents de codage en 2026.

## Comprendre git worktree : les concepts essentiels avant de commencer

Un dépôt Git classique contient un seul dossier `.git` et un seul dossier de travail. Quand vous changez de branche, Git réécrit les fichiers de ce même dossier pour refléter l’état de la nouvelle branche. C’est rapide, mais ça impose un choix : un seul état visible à la fois.

Git worktree change la donne en séparant deux choses qui semblaient liées : la base de données Git (l’historique, les commits, les objets) et le dossier de travail (les fichiers que vous voyez et éditez). Une seule base de données peut alimenter plusieurs dossiers de travail simultanés, chacun pointant vers une branche différente. Trois règles gouvernent ce fonctionnement :

- **Une branche ne peut être extraite que dans un seul worktree à la fois.** Git refuse explicitement de la checkout ailleurs tant qu’elle reste active quelque part.
- **Les objets sont partagés, pas dupliqués.** Un commit créé dans un worktree est immédiatement visible dans les autres, sans opération réseau.
- **Chaque worktree a son propre index et son propre HEAD.** Vous pouvez avoir des fichiers modifiés et non validés dans un worktree sans que cela affecte les autres.

Cette architecture est exactement ce dont un agent de codage IA a besoin. Un agent qui modifie dix fichiers pour corriger un bug ne doit pas voir ses changements se mélanger avec ceux d’un second agent qui construit une fonctionnalité sur une autre branche. Avec des clones séparés, on obtient la même isolation, mais au prix d’un espace disque et d’un temps de clonage démultipliés. Le worktree partage l’historique et ne duplique que ce qui doit vraiment l’être.

Concrètement, chaque worktree ajouté laisse une trace dans un dossier administratif caché, `.git/worktrees/<nom>`, situé dans le dépôt principal. C’est là que Git range les informations propres à ce worktree : sa référence HEAD, son fichier de configuration local et l’emplacement du dossier de travail associé. Vous n’avez jamais besoin d’y toucher directement, mais savoir que cette structure existe explique pourquoi une commande comme `git worktree repair` fonctionne : elle recalcule simplement le contenu de ces dossiers administratifs quand un chemin a changé.

## Étape 1 – Vérifier et préparer votre installation Git

Ouvrez un terminal dans le dossier de votre projet et confirmez que git worktree est disponible.

```
git --version
git worktree --help
```
Vous devriez voir un résumé des sous-commandes disponibles : add, list, lock, move, prune, remove, repair et unlock. Si la commande `git worktree` renvoie une erreur du type « git: ‘worktree’ is not a git command », votre version de Git est trop ancienne et il faut la mettre à jour avant d’aller plus loin.

Profitez-en pour repérer le chemin racine de votre dépôt, car vous en aurez besoin à chaque étape suivante :

```
git rev-parse --show-toplevel
# Exemple de sortie : /Users/marie/projets/api-recettes
```
## Étape 2 – Créer votre premier worktree avec git worktree add

La commande centrale de ce tutoriel est `git worktree add`. Elle prend un chemin de destination et, en option, une branche à extraire ou à créer. Depuis la racine de votre dépôt principal, lancez :

`git worktree add ../api-recettes-agent-claude -b agent/fix-auth`
Cette ligne crée un nouveau dossier `api-recettes-agent-claude` à côté de votre dépôt principal, avec une nouvelle branche `agent/fix-auth` créée à partir de votre HEAD actuel. Git affiche une confirmation :

