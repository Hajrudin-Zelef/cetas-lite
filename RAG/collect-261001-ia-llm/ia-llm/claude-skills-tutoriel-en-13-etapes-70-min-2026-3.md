---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-3
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2025-04-14", "2025-10-02"]
keywords: ["claude", "agent"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [121, 212]
sha256: cb1c45ad22be8efba748763a8797b1a5d6f02f03784433f357873f9af18ea8e9
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

Une Skill ne se limite pas à un seul fichier. C’est le “niveau 3” de la divulgation progressive : fichiers Markdown additionnels, scripts exécutables, gabarits, données de référence — rien de tout cela ne consomme de tokens tant que Claude ne l’ouvre pas explicitement. Pour notre Skill `pr-description`, ajoutons un script qui valide qu’un ticket existe avant de générer la description :

```
pr-description/
├── SKILL.md              (instructions principales)
├── FORMAT.md             (gabarit détaillé, chargé si besoin)
└── scripts/
    └── check_ticket.sh   (script exécuté, jamais chargé en contexte)
```
Deux mécaniques différentes cohabitent ici, et la distinction compte. Un fichier Markdown référencé (comme `FORMAT.md`) charge son contenu texte dans le contexte quand Claude le lit. Un script (comme `check_ticket.sh`) s’exécute via bash, et seul son résultat — “ticket #482 trouvé” ou un message d’erreur — entre dans le contexte, jamais le code du script lui-même. C’est ce qui rend les scripts nettement plus économes qu’une Skill qui demanderait à Claude de régénérer la même logique à chaque fois : Firecrawl citait en août 2026 l’exemple de la Skill Caveman Claude, qui réduit de 65 % en moyenne le nombre de tokens de sortie générés en forçant des réponses délibérément minimalistes. Référencez toujours ces fichiers annexes explicitement depuis le `SKILL.md` principal, sinon Claude n’a aucun moyen de savoir qu’ils existent :

```
## Ressources additionnelles
- Pour le gabarit détaillé, voir [FORMAT.md](FORMAT.md)
- Pour valider un numéro de ticket, exécute scripts/check_ticket.sh
```
## Étape 7 : injecter du contexte dynamique avec Bash

Revenons sur la ligne `` !`git diff HEAD` `` introduite à l’étape 4. La syntaxe `` !`commande` `` déclenche l’injection de contexte dynamique, une extension propre à Claude Code : la commande s’exécute avant que Claude ne lise le contenu de la Skill, et son résultat remplace la ligne dans le texte final. Concrètement, Claude ne voit jamais la commande brute — il voit directement la sortie de `git diff HEAD`, déjà intégrée aux instructions.

Cette technique évite un aller-retour inutile : sans elle, Claude devrait d’abord lire la Skill, comprendre qu’il doit exécuter `git diff`, lancer la commande lui-même, puis relire le résultat. Avec l’injection dynamique, les données arrivent déjà prêtes dans le même chargement. C’est particulièrement utile pour tout ce qui change à chaque exécution : diff Git, date du jour, statut d’un service, nombre de tests en échec.

```
## Contexte
Branche actuelle : !`git branch --show-current`
Fichiers modifiés : !`git diff --name-only HEAD`
Diff complet :
!`git diff HEAD`
```
Par défaut, ces commandes s’exécutent via bash. Si votre équipe travaille sous Windows sans Git Bash, le champ de frontmatter `shell: powershell` bascule l’exécution vers PowerShell pour les blocs `` !`commande` `` de cette Skill précise.

## Étape 8 : utiliser arguments et substitutions de variables

Notre Skill doit accepter un numéro de ticket en argument, par exemple via `/pr-description 482`. Claude Code propose plusieurs formes de substitution pour ce cas de figure, résumées dans le tableau suivant.

| Variable | Résultat | 
|---|---|
| `$ARGUMENTS` | Tous les arguments passés, sous forme de chaîne complète | 
| `$ARGUMENTS[N]` ou`$N` | Argument à l’index N (base 0), ex. `$0` pour le premier | 
| `$nom` | Argument nommé, déclaré via le champ `arguments` du frontmatter | 
| `${CLAUDE_SESSION_ID}` | Identifiant de la session en cours | 
| `${CLAUDE_SKILL_DIR}` | Dossier contenant le SKILL.md, utile pour référencer un script joint | 
| `${CLAUDE_PROJECT_DIR}` | Racine du projet, indépendamment de l’emplacement de la Skill | 

Pour donner un nom explicite à notre argument plutôt que d’utiliser l’indice `$0`, déclarez-le dans le frontmatter :

```
---
name: pr-description
description: Génère une description de pull request structurée à partir du diff Git. Utiliser quand l'utilisateur demande une description de PR ou veut résumer ses changements.
argument-hint: "[numero-ticket]"
arguments: [ticket]
---
Si un ticket est fourni ($ticket), ajoute une ligne "Closes #$ticket"
à la fin de la description.
```
Un piège à connaître : un indice sans argument correspondant (`$1` alors qu’un seul argument a été passé) reste affiché tel quel dans le texte, sans erreur. Un argument nommé sans valeur, lui, se réduit silencieusement à une chaîne vide. Testez toujours votre Skill avec et sans argument avant de la considérer prête.

## Étape 9 : contrôler qui peut invoquer le Skill

Toutes les Skills ne devraient pas se déclencher automatiquement. Un déploiement en production ou l’envoi d’un message Slack sont des actions à effet de bord : vous ne voulez pas que Claude décide seul de les lancer parce que le code “a l’air prêt”. Deux champs de frontmatter permettent de restreindre l’invocation.

| Frontmatter | Vous pouvez invoquer | Claude peut invoquer | 
|---|---|---|
| (par défaut) | Oui | Oui | 
| `disable-model-invocation: true` | Oui | Non | 
| `user-invocable: false` | Non | Oui | 

`disable-model-invocation: true` convient aux workflows à effet de bord comme un déploiement : seule une invocation manuelle via `/deploy` fonctionne, et si Claude tente de reproduire les étapes de son propre chef, Claude Code bloque l’appel et lui indique explicitement de ne pas improviser la même procédure autrement. À l’inverse, `user-invocable: false` convient à une Skill de connaissance de fond, comme la documentation d’un système legacy : ce n’est pas une action que l’utilisateur déclenche lui-même, mais une information que Claude doit connaître quand c’est pertinent.

## Étape 10 : exécuter un Skill dans un sous-agent isolé

Pour des Skills longues ou consommatrices de contexte (un audit de sécurité complet, un run de tests exhaustif), le champ `context: fork` exécute la Skill dans un sous-agent isolé plutôt que dans la conversation principale. Le sous-agent tourne en arrière-plan par défaut ; seul le résultat final revient enrichir le contexte principal, pas chaque étape intermédiaire.

```
---
name: full-pr-review
description: Effectue une revue complète d'une pull request (sécurité, style, tests, performance). Utiliser pour une revue approfondie avant fusion.
context: fork
agent: code-reviewer
background: true
---
Analyse la pull request en cours sous quatre angles : sécurité,
respect des conventions du projet, couverture de tests, et impact
sur les performances. Produis un rapport structuré par section.
```
Le champ `agent` précise quel type de sous-agent utiliser quand `context: fork` est actif. Si vous préférez attendre le résultat dans le tour de conversation qui a déclenché la Skill plutôt que de la laisser tourner en tâche de fond, passez `background: false` — utile en interactif, moins pertinent pour une Skill que vous invoquez puis oubliez le temps qu’elle travaille.

## Étape 11 : utiliser les Skills pré-construites via l’API

Les quatre Skills documentaires d’Anthropic (`pptx`, `xlsx`, `docx`, `pdf`) sont accessibles directement via l’API Claude, sans rien installer. Elles nécessitent l’outil d’exécution de code (*code execution tool*) comme conteneur d’exécution, ainsi qu’un en-tête bêta spécifique : `skills-2025-10-02`. Si vous envoyez ou récupérez des fichiers via la Files API dans le même appel, ajoutez également l’en-tête `files-api-2025-04-14`.

