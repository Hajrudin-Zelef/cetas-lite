---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique-1
title: "deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "claude", "deepseek", "agents", "bedrock", "benchmark", "benchmarks", "mcp", "open source", "transcription"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [1, 98]
sha256: b0d04cf478ec8ca163316b14aafb049e369c781c086ccbb3cdfd592f59e46fc8
---

# deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique

Cursus

La plupart des comparatifs d’agents de code se résument à une course à la vitesse, au prix, au nombre d’outils ou aux scores de benchmark. Ce cadrage passe à côté de la vraie question qui distingue ces deux approches : jusqu’où un développeur peut-il remplacer les composants du runtime de l’agent ?

DeepSeek Harness et Claude Code répondent à cette question à des niveaux différents. Harness expose les adaptateurs de modèle, le stockage, les bacs à sable et la boucle d’agent sous forme de plugins remplaçables. Claude Code embarque sa propre boucle autour de Claude et prévoit des points d’extension pour le workflow. Ici, cette frontière compte davantage qu’une simple comparaison des modèles DeepSeek et Claude.

J’ai donné aux deux outils le même dépôt cassé et le même modèle Claude pour observer ce que change le runtime. Une étude de cas ne suffit pas à départager les produits, mais elle montre pourquoi le harness n’est pas un détail d’arrière-plan. Je me concentre sur les différences de choix de modèle, de configuration, de vérification, de journaux d’exécution et de coût.

## À retenir

- Si le runtime fait partie du travail : Harness expose l’adaptateur de modèle, le stockage, le bac à sable et la boucle d’agent comme plugins remplaçables, avec prise en charge de plusieurs fournisseurs de modèles.
- Si l’essentiel est le code applicatif : Claude Code embarque davantage de runtime et étend son workflow via des Skills, des hooks, MCP et fonctionnalités associées.
- Test avec le même modèle : les deux ont produit un patch identique à l’octet et passé la suite de tests d’origine. Dans cet essai sur Windows, Claude Code a annoncé 55,0 secondes ; Harness a enregistré 125,4 secondes après trois demandes d’approbation.
- Coût : DeepSeek Harness n’a pas de frais de licence ; l’usage du modèle est facturé par le fournisseur choisi, avec d’éventuels coûts d’infrastructure. Claude Code est inclus dans les offres Claude payantes.
- Règle simple : commencez avec Claude Code pour les travaux applicatifs courants. Optez pour Harness lorsque modifier ou inspecter le runtime fait partie du mandat.

## Introduction aux agents d'intelligence artificielle

## DeepSeek Harness vs Claude Code : comparaison rapide

| Dimension | DeepSeek Harness | Claude Code | 
|---|---|---|
| Ce que c’est | Runtime d’agent à composants remplaçables | Agent de code packagé | 
| Licence et statut | MIT, aperçu développeur, pas de version stable | Propriétaire, produit en production | 
| Modèles | DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, local | Claude, en direct ou via Bedrock et Vertex | 
| Boucle remplaçable | Oui, point d’extension plugin substituable | Non, des extensions se greffent autour | 
| Unité d’extension | Plugin Cordis, toute capacité du runtime | Plugin regroupant skills, hooks, agents, MCP | 
| Historique d’exécution | Journal d’événements typés en ajout seul, relecturable | Transcription JSONL, checkpoints, OpenTelemetry | 
| Surfaces | Interface web locale, CLI headless, SDK Python | Terminal, IDE, desktop, web, Slack, CI | 
| Configuration | Nécessite la configuration du fournisseur et la maîtrise des profils | Configuration par défaut plus légère ; permissions, hooks, MCP et projet en option | 

## Qu’est-ce que DeepSeek Harness ?

DeepSeek Harness est un harness d’agent open source actuellement en aperçu développeur. Il fournit le runtime dans lequel les modèles utilisent des outils et du contexte, et il peut exécuter des modèles d’autres fournisseurs que DeepSeek. Son README avertit que des mises à jour peuvent casser des configurations existantes.

Si vous souhaitez l’essayer, commencez par notre tutoriel DeepSeek Harness.

### Comment fonctionne DeepSeek Harness

La boucle d’agent est implémentée comme un plugin. Harness utilise Cordis pour composer les adaptateurs de modèle, les outils, les sessions, le stockage, les bacs à sable et la boucle.

Cordis permet aux plugins de découvrir des services et d’échanger des événements. Les changements de dépendances chargent ou déchargent des plugins ; le rechargement à chaud nettoie les écouteurs obsolètes et les tâches en arrière-plan.

Le profil par défaut expose la boucle comme une configuration. `dsh --profile headless --dump-default-config` inclut cette entrée :

```
- id: agent-loop
  name: '@deepseek-ai/dsh-agent-loop'
  config:
    agents: []
```
Cette entrée peut être remplacée via `cordis.patch.yml`, preuve que la boucle elle-même n’est pas figée.

Panneau des plugins affichant les composants du runtime. Vidéo de l’auteur.

### Les quatre modes d’exécution

DeepSeek Harness propose quatre modes d’exécution distincts.

- **Standard** : mode agent de code complet
- **PTC** /**Code** : combine plusieurs appels d’outils en un programme TypeScript
- **Minimal** : ne conserve que bash persistant et str_replace_editor
- **Creator** : pour inspecter et tester le runtime

DeepSeek a utilisé le mode *Minimal* pour ses benchmarks.

## Qu’est-ce que Claude Code ?

Claude Code est l’agent de code propriétaire d’Anthropic et un harness agentique pour un usage local ou managé via le terminal, l’IDE, le desktop, le web, Slack et le CI.

La CLI se lance avec `claude` dans le répertoire du projet. Ce répertoire devient le périmètre de fichiers par défaut, et la session lit les instructions du projet depuis `CLAUDE.md`. Les utilisateurs peuvent élargir l’accès aux fichiers ou ajouter des services externes ensuite.

Anthropic présente ces interfaces comme des moyens d’accéder à Claude Code. Les sessions locales s’exécutent sur votre machine. Les sessions cloud tournent dans des environnements managés ou sur des serveurs opérés par votre organisation. Le Remote Control permet de piloter un travail local depuis un navigateur.

Les meilleurs points de départ sont notre tutoriel Claude Code et le guide des bonnes pratiques Claude Code.

### Boucle intégrée et extensions

Anthropic décrit Claude Code comme le harness agentique autour de Claude, avec une boucle répétée de contexte, action et vérification. Les utilisateurs peuvent l’orienter pendant l’exécution.

Claude Code stocke les sessions en local et compacte les anciens contextes à mesure que la fenêtre se remplit. `CLAUDE.md` et la mémoire automatique conservent des instructions et détails de projet sélectionnés d’une session à l’autre. Les sous-agents utilisent des fenêtres de contexte séparées et renvoient des synthèses à la session parente.

### Comment fonctionne la couche d’extension de Claude Code

Claude Code prend en charge plusieurs mécanismes d’extension. `CLAUDE.md`, skills, hooks, MCP, subagents, plugins, Agent Teams et l’Agent SDK étendent tous son workflow.

Ces extensions opèrent autour de la boucle intégrée de Claude Code. Les hooks peuvent imposer des règles d’appels d’outils, tandis que l’Agent SDK fournit des outils et la gestion du contexte dans le code.

## DeepSeek Harness vs Claude Code : architecture et contrôle

DeepSeek Harness expose des éléments de runtime de plus bas niveau, remplaçables. Claude Code conserve sa boucle intégrée et prend en charge des extensions autour d’elle.

### Runtime remplaçable vs agent packagé

DeepSeek Harness considère le runtime comme une infrastructure configurable : adaptateurs de modèle, stockage, bacs à sable et boucle peuvent être remplacés via des entrées de profil. Claude Code garde sa boucle intégrée fixe et étend le workflow autour.

Harness peut utiliser ce mécanisme pour échanger les fournisseurs de modèles ou d’autres parties du runtime. Le mot « plugin » n’a pas la même portée ici :

