---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-3
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["copilot", "agent", "agents", "arr", "claude", "gemini", "incident", "mcp", "model context protocol", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [103, 177]
sha256: 9a1b2ab4bef9d38a98da55f5da6c8c68dce0ebdb1875612866ca745d4397d8ae
---

# .github/copilot-instructions.md

Cette boucle autonome est la vraie différence avec le mode Edit, qui s’arrête après une seule modification sans vérifier le résultat. Elle explique aussi pourquoi la qualité du prompt initial compte autant : plus le contexte fourni est précis, moins l’agent a besoin d’itérations pour converger vers un résultat correct, et moins vous consommez de crédits au passage.

## Étape 6 – Approuver les modifications et les commandes terminal

Par défaut, l’Agent Mode ne vous laisse jamais complètement les mains liées. Chaque modification de fichier s’affiche dans un éditeur de différences multi-fichiers avant d’être définitivement appliquée. Deux boutons apparaissent : **Keep** pour accepter les changements, et **Undo** pour les rejeter. Vous pouvez examiner chaque fichier modifié individuellement avant de valider l’ensemble.

Le même principe s’applique aux commandes terminal. L’agent propose une commande, l’affiche en clair, et attend votre confirmation avant de l’exécuter, sauf si vous avez explicitement autorisé l’exécution automatique pour certains types de commandes dans les paramètres. Prenez l’habitude de lire ce qui s’affiche avant de cliquer, en particulier pour tout ce qui touche à la suppression de fichiers, aux migrations de base de données ou aux commandes Git qui réécrivent l’historique. Ce réflexe simple évite la majorité des incidents liés à l’usage agentique.

Pour les tâches répétitives et sans risque, comme le formatage de code ou l’exécution de tests en lecture seule, il est possible d’élargir la liste des commandes approuvées automatiquement dans les réglages de l’extension. Faites-le progressivement, commande par commande, plutôt que d’activer une confiance totale dès le premier jour.

Chaque session conserve par ailleurs un historique consultable des actions menées : fichiers touchés, commandes lancées, réponses obtenues. Sur un projet d’équipe, ce journal aide à comprendre a posteriori pourquoi tel fichier a changé, un peu comme un message de commit détaillé mais généré automatiquement. Sur les paliers Business et Enterprise, une partie de cet historique remonte aussi dans les journaux d’audit accessibles aux administrateurs, ce qui facilite une revue de sécurité si un incident survient.

## Étape 7 – Personnaliser l’agent avec copilot-instructions.md et AGENTS.md

Un agent qui ignore les conventions de votre équipe perd vite en utilité. GitHub Copilot lit automatiquement un fichier d’instructions placé à la racine du dépôt, dans `.github/copilot-instructions.md`, à chaque nouvelle session. Vous pouvez y décrire l’architecture du projet, les bibliothèques à privilégier, le style de code attendu, ou les zones du code à ne jamais toucher sans validation explicite. Voici un exemple pour un projet Python orienté API :

```
# .github/copilot-instructions.md
## Contexte du projet
API interne construite avec FastAPI et SQLAlchemy.
Base de données PostgreSQL, migrations gérées par Alembic.
## Conventions à respecter
- Toutes les fonctions asynchrones utilisent async/await.
- Les réponses d'erreur suivent le format {"detail": "message"}.
- Chaque nouvelle route doit être accompagnée d'un test dans tests/.
- Ne jamais modifier le dossier migrations/ sans le signaler explicitement.
## À éviter
- Pas de dépendance externe sans la mentionner dans le plan proposé.
- Pas de commit ni de push automatique.
```
Depuis 2026, GitHub Copilot prend également en charge `AGENTS.md`, une convention de fichier partagée par plusieurs éditeurs d’outils IA pour développeurs, qui permet de décrire l’agent une seule fois pour plusieurs assistants compatibles. La structure ressemble à ceci :

```
# AGENTS.md
## Build
pip install -r requirements.txt
## Tests
pytest -v
## Style
Suivre PEP 8. Formatage automatique avec black avant chaque commit.
## Notes pour l'agent
Le dossier legacy/ contient du code hérité non testé.
Éviter d'y toucher sauf demande explicite.
```
Si les deux fichiers coexistent dans un même dépôt, considérez `copilot-instructions.md` comme la source prioritaire pour Copilot spécifiquement, et `AGENTS.md` comme un socle commun utile si votre équipe teste aussi d’autres assistants agentiques. Pour des règles plus ciblées, GitHub Copilot accepte aussi des fichiers d’instructions par dossier avec un motif `applyTo`, utile sur un monorepo où le frontend et le backend suivent des règles différentes.

## Étape 8 – Connecter un serveur MCP pour étendre l’agent

Le Model Context Protocol (MCP) est un standard ouvert qui permet à un agent IA de se connecter à des outils externes : accès à une base de données, appel à une API interne, lecture d’un système de fichiers spécifique, interaction avec un gestionnaire de tickets. Le site officiel du protocole recense les serveurs disponibles et la spécification technique complète. GitHub Copilot Agent Mode supporte nativement les serveurs MCP depuis 2025, et GitHub a confirmé en juin 2026 le déploiement de l’Agent Mode avec support MCP à l’ensemble des utilisateurs de VS Code, ce qui démultiplie ce que l’agent peut accomplir sans quitter l’éditeur.

Pour en ajouter un, ouvrez les paramètres de VS Code (`Ctrl+,`), recherchez « copilot mcp », ou créez directement un fichier `.vscode/mcp.json` à la racine du projet. Voici un exemple de configuration pour le serveur MCP officiel de GitHub, qui donne à l’agent un accès structuré à vos dépôts, issues et pull requests :

```
{
  "servers": {
    "github": {
      "type": "stdio",
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "ghcr.io/github/github-mcp-server"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
      }
    }
  }
}
```
Une fois le fichier enregistré, VS Code demande de confirmer le démarrage du serveur puis affiche un bouton « outils » dans la vue Chat, où vous pouvez activer ou désactiver chaque capacité exposée par le serveur MCP. Redémarrez VS Code après une modification de la configuration pour que les changements soient pris en compte. La syntaxe exacte évolue régulièrement : vérifiez la documentation officielle du serveur MCP que vous ajoutez avant de copier une configuration trouvée en ligne.

## Étape 9 – Choisir le bon modèle pour chaque tâche

Le sélecteur de modèle, situé juste à côté du sélecteur de mode dans la vue Chat, permet de choisir quel modèle traite votre requête. GitHub a déployé GPT-5.4 comme nouveau modèle de référence pour le codage agentique en mars 2026 (accessible sur les paliers Pro, Pro+, Business et Enterprise), avant de le compléter par GPT-5.5 dès avril 2026 ; l’entreprise a par ailleurs annoncé en juillet 2026 la mise en retraite de plusieurs modèles au 1er septembre 2026, dont Claude Opus 4.5 et Gemini 3.1 Pro, ce qui invite à vérifier régulièrement la liste des modèles disponibles dans le sélecteur. Tous les modèles ne consomment pas la même quantité de crédits : un modèle léger convient très bien à une tâche simple comme renommer des variables ou écrire un test unitaire basique, tandis qu’une tâche qui demande un raisonnement plus profond, comme refactoriser une architecture ou déboguer un problème de concurrence, tire davantage parti d’un modèle premium comme ceux accessibles via les paliers Pro+ et Max.

