---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-3
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "claude", "cloud agent", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [127, 195]
sha256: fde7b9a46cf74d71870cf4a583c65ef38145024b0dd8d13b7d64b2481da53a9b
---

# AGENTS.md

```
---
description: "Conventions API Python"
globs: api/**/*.py
alwaysApply: false
---
Chaque fonction publique a une signature typée et une docstring courte en une ligne.
Les exceptions métier héritent d'une classe commune AppException, jamais d'Exception nue.
Toute requête vers la base de données passe par le repository dédié, jamais d'ORM direct dans les routes.
```
### Étape 9 – Organiser les règles par dossier et cibler les bons fichiers avec globs

Sur un projet de taille réelle, un seul fichier de règles devient vite ingérable. Cursor AI permet d’organiser les fichiers `.mdc` en sous-dossiers, par exemple `.cursor/rules/frontend/composants.mdc` ou `.cursor/rules/backend/validation.mdc`. Le champ `globs` détermine quels fichiers déclenchent la règle, et `alwaysApply` détermine si elle s’applique en permanence ou seulement quand un fichier correspondant est ouvert. Selon la documentation officielle, ces règles sont **« incluses au tout début du contexte du modèle, ce qui donne à l’IA des instructions cohérentes pour générer du code, interpréter des modifications ou aider sur un workflow »**. Concrètement, mieux vaut trois règles courtes et ciblées qu’un seul fichier de 300 lignes que le modèle finit par ignorer partiellement faute de place dans son contexte.

## Étape 10 : utiliser AGENTS.md pour un contexte simplifié

Pour un projet simple ou un prototype, écrire des fichiers `.mdc` avec en-tête YAML peut sembler excessif. Cursor AI accepte une alternative plus légère : un fichier `AGENTS.md` classique à la racine du projet, sans métadonnées à gérer. Le système prend aussi en charge des fichiers `AGENTS.md` imbriqués dans les sous-dossiers, avec, selon la documentation, **« les instructions les plus spécifiques qui prennent le pas »** sur les plus générales.

```
# AGENTS.md
## Stack technique
- Backend : Python 3.12, FastAPI, PostgreSQL
- Tests : pytest, couverture minimale 80 %
- Style : formatage via ruff, pas d'exceptions silencieuses
## Conventions
- Toute route API retourne un schéma Pydantic explicite
- Les migrations passent par Alembic, jamais de modification manuelle du schéma
- Les secrets viennent des variables d'environnement, jamais du code source
## Commandes utiles
- Lancer les tests : pytest -v
- Démarrer le serveur local : uvicorn app.main:app --reload
```
En pratique, beaucoup d’équipes démarrent avec un `AGENTS.md` unique, puis migrent vers des fichiers `.mdc` ciblés dès que le projet grossit et que certaines règles doivent s’appliquer uniquement à un sous-dossier précis. Les deux mécanismes cohabitent sans problème dans un même dépôt, et rien n’empêche de garder un `AGENTS.md` général pour le contexte métier (ce que fait l’application, qui sont les utilisateurs) tout en déléguant les conventions de code strictes aux fichiers `.mdc`.

## Étape 11 : maîtriser le mode Agent et les workflows agentiques

Cursor AI propose trois niveaux d’intervention. L’édition en ligne (`Ctrl+K`) modifie un bloc de code précis sur votre demande. Le chat (`Ctrl+L`) répond à des questions et peut proposer des modifications que vous appliquez manuellement. Le mode Agent, lui, va plus loin : il lit les fichiers pertinents du projet, écrit ou modifie plusieurs fichiers à la fois, exécute des commandes dans le terminal intégré, corrige les erreurs qu’il rencontre, puis présente un diff complet à valider avant application.

Pour basculer en mode Agent, ouvrez le panneau latéral et sélectionnez Agent dans le menu déroulant au-dessus de la zone de saisie. Un bon prompt agentique décrit un objectif plutôt qu’une action isolée : plutôt que “ajoute une validation”, préférez “ajoute une validation d’email sur le formulaire d’inscription, avec un message d’erreur clair et un test unitaire associé”. L’agent découpe alors la tâche en sous-étapes, les exécute, et signale les points qui nécessitent votre arbitrage, par exemple un choix de bibliothèque ou une ambiguïté dans la spécification.

La différence se voit bien sur un cas concret. Demander en mode Chat “comment ajouter la pagination à cette liste ?” renvoie une explication et un extrait de code à copier-coller soi-même. Demander la même chose en mode Agent déclenche une série d’actions : lecture du composant de liste existant, modification du composant pour accepter des paramètres de page, ajustement de l’appel API correspondant, et proposition d’un test qui vérifie le comportement sur une deuxième page de résultats. Le gain de temps est réel, mais il déplace aussi la charge de travail : moins de temps à écrire, davantage de temps à relire attentivement ce qui a changé dans plusieurs fichiers à la fois.

Ce mode agentique n’est pas propre à Cursor AI. Des outils comme Claude Code ou Google Antigravity suivent la même logique avec des approches différentes de la gestion du contexte et de l’autonomie accordée à l’IA. La bonne pratique, quel que soit l’outil, reste la même : ne jamais valider un diff sans l’avoir lu, même quand le résultat semble correct au premier coup d’œil.

## Étape 12 : connecter des serveurs MCP à Cursor AI

Le Model Context Protocol, ou MCP, est un standard ouvert introduit par Anthropic et désormais repris par la plupart des éditeurs IA du marché, dont Cursor. Sur le site officiel du protocole, la documentation de Cursor résume sa fonction ainsi : **« le Model Context Protocol permet à Cursor de se connecter à des outils et des sources de données externes »**. Concrètement, un serveur MCP donne à l’agent un accès structuré à une base de données, un gestionnaire de tickets, un outil de design ou une API interne, sans que vous ayez à copier-coller manuellement du contexte à chaque conversation.

La configuration se fait dans un fichier `.cursor/mcp.json` à la racine du projet pour une portée locale, ou `~/.cursor/mcp.json` pour une portée globale à tous vos projets. Voici un exemple de connexion à un serveur MCP local lancé via Node.js :

```
{
  "mcpServers": {
    "postgres-projet": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-postgres"],
      "env": {
        "DATABASE_URL": "postgresql://user:password@localhost:5432/mabase"
      }
    }
  }
}
```
Les serveurs MCP ne se limitent pas à Node.js : la documentation officielle précise qu’il est possible d’**« écrire des serveurs MCP dans n’importe quel langage capable d’écrire sur stdout ou de servir un point de terminaison HTTP, comme Python, JavaScript ou Go »**. Pour un serveur distant exposé en HTTP, la syntaxe change légèrement : un champ `url` remplace `command` et `args`. Une fois le serveur déclaré, redémarrez Cursor AI et vérifiez son statut dans Settings > MCP : un point vert confirme la connexion, un point rouge signale généralement un problème de commande ou de variable d’environnement manquante.

Trois cas d’usage reviennent le plus souvent en entreprise. Un serveur MCP relié à un gestionnaire de tickets permet à l’agent de lire directement la description d’un ticket avant de coder, plutôt que de la faire recopier manuellement. Un serveur relié à un outil de design donne accès aux spécifications visuelles sans captures d’écran intermédiaires. Un serveur relié à une base de données de lecture seule permet à l’agent de vérifier un schéma réel plutôt que de le deviner à partir d’un fichier de migration potentiellement obsolète. Dans les trois cas, la règle reste la même : ne jamais donner plus de droits que la tâche ne l’exige.

## Étape 13 : déployer un Cloud Agent et automatiser vos tâches

