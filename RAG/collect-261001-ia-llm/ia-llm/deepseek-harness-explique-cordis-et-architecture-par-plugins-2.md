---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins-2
title: "deepseek-harness-explique-cordis-et-architecture-par-plugins"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["deepseek", "agent", "agents", "benchmarks", "claude", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [88, 188]
sha256: f34f928edb0b1642a60f7d1615d977ad783a910279d991b2e03d96456ca00eb9
---

# deepseek-harness-explique-cordis-et-architecture-par-plugins

DeepSeek distingue une étape d’un tour. Une étape est une requête modèle plus ses appels d’outils. Un tour regroupe zéro ou plusieurs étapes : il s’ouvre avant que sa première entrée ne soit prise en charge et se ferme quand plus rien n’est dû. La plupart des tours exécutent plusieurs étapes avant que l’agent ne réponde, mais une entrée rejetée ferme un tour sans étape consommée.

Un tour peut contenir plusieurs étapes. Image de l’auteur.

### Les sessions utilisent un journal d’événements en ajout-only

C’est, à mon sens, l’élément le plus important. Une session est un journal append-only d’événements typés, pas un tableau de messages de chat. Harness reconstruit l’historique du modèle à partir de ce journal, et la documentation des sessions exige que tout ce qui est envoyé au modèle soit récupérable depuis celui-ci.

La reprise, le fork, la recherche, la relecture et la vue Trajectory reposent tous sur ce flux d’événements.

Reconstituer l’historique n’est pas une réexécution déterministe. La sortie du modèle et l’état externe peuvent différer, mais le journal fournit tout de même une trace consultable de ce qui s’est passé.

L’historique de session est un journal en ajout-only. Image de l’auteur.

### Comment DeepSeek Harness contrôle les outils et les bacs à sable

Un modèle peut demander un outil par son nom, mais il ne peut pas l’exécuter directement. Deux contrôles distincts s’interposent entre la demande et une modification du système de fichiers.

#### Le pipeline d’exécution des outils

L’appel passe par une vérification de politique, l’exécution, puis le traitement du résultat. Le modèle choisit l’outil ; le runtime décide si et comment il s’exécute.

Le runtime décide comment les outils s’exécutent. Image de l’auteur.

#### Bac à sable versus approbations

- **Approbation** demande si l’utilisateur doit confirmer une action.
- **Bac à sable** limite où et comment l’action s’exécute.

DeepSeek les sépare, même si des presets d’autorisations regroupent les deux contrôles, à l’image d’un runtime de conteneur qui sépare permissions de processus et limites d’exécution.

À signaler dès maintenant, car j’y reviendrai dans les limites : dire à un modèle dans un prompt système de « ne lire que des fichiers » est une consigne qu’il peut choisir de suivre, pas une barrière imposée comme le ferait une restriction de bac à sable au niveau OS.

## Modes de DeepSeek Harness : Standard, PTC, Minimal et Creator

DeepSeek Harness propose quatre modes. Aucun n’est « meilleur » que les autres. Ce sont quatre réponses à « quelle part du runtime doit être exposée à cette session », et le bon choix dépend de la tâche. Comme l’a montré l’architecture, chaque mode modifie l’ensemble d’outils disponible pour l’agent.

Quatre modes, une base runtime commune. Image de l’auteur.

### Mode Standard

Le socle polyvalent :

- Édition de fichiers
- Accès shell
- Recherche de fichiers et sur le web
- Skills
- Planification
- Objectifs
- Sous-agents
- Workflows

Pour un travail habituel sur dépôt, c’est mon point de départ.

### Mode PTC

Le mode PTC conserve presque tout l’outillage de Standard mais change la façon dont le modèle y accède. (Depuis la version 0.1.2, le mode Web PTC n’expose plus par défaut l’outil générique `workflow`.) 

Au lieu d’appeler des outils un par un sur plusieurs étapes, le modèle écrit un programme contre un SDK généré. Ce programme peut invoquer plusieurs outils via `run_code`. Chaque appel passe toujours par les mêmes contrôles de politique : PTC change la façon d’énoncer le plan, pas ce que le modèle est autorisé à faire.

La page produit utilise encore l’étiquette « Code mode », mais une version officielle plus récente l’a renommé « mode PTC » tout en gardant lisibles les anciennes conversations. J’utiliserai « mode PTC » partout ; la FAQ revient sur la signification possible de ces initiales.

### Mode Minimal

Le mode Minimal réduit l’environnement à deux outils : un shell persistant et un éditeur de fichiers par substitution de chaînes. DeepSeek l’utilise pour les benchmarks de modèles, car les résultats dépendent en partie du harness, pas seulement des poids du modèle.

### Mode Creator

Le mode Creator permet aux développeurs d’inspecter le runtime et de tester des plugins Cordis en mémoire. Il sert à construire des presets ; je ne le qualifierais pas « d’auto-améliorant » dans un sens plus profond.

## Ce qui distingue DeepSeek Harness des autres frameworks d’agents

DeepSeek Harness se distingue de nombreux frameworks d’agents en rendant remplaçables les couches basses du runtime. J’aurais pu l’intégrer à l’architecture, mais la nuance est facile à manquer. Cordis gère ces changements via un système unique de plugins.

Vous pouvez modifier le fonctionnement même de l’agent, pas seulement les outils qu’il appelle. Le journal d’événements crée également une exécution inspectable par les développeurs, plutôt qu’un simple relevé de chat. Les modes Minimal et Creator permettent ensuite de tester le runtime depuis deux angles opposés.

## DeepSeek Harness vs Claude Code, Codex et OpenCode

Une checklist de fonctionnalités passerait à côté de l’essentiel. Chaque concurrent prend en charge des extensions ; la vraie question est : quelles parties les développeurs peuvent-ils changer ? La nuance semble minime, elle ne l’est pas. Notre comparatif dédié Harness vs Claude Code utilise le même modèle des deux côtés et couvre configuration, journaux et coût.

### DeepSeek Harness vs Claude Code

Claude Code prend en charge des instructions de projet, des skills, des hooks, le MCP, des sous-agents et un Agent SDK, tout en gardant sa boucle interne fixe. DeepSeek Harness permet de remplacer via configuration la boucle, l’adaptateur de modèle et la couche de stockage.

### DeepSeek Harness vs Codex

Codex nécessite une comparaison plus fine, car son CLI et son App Server sont également open source. Il fournit un agent harness que les développeurs étendent via des points d’entrée documentés. DeepSeek Harness est conçu pour modifier le runtime lui-même. Les niveaux de contrôle offerts diffèrent.

### DeepSeek Harness vs OpenCode

OpenCode est déjà open source, fonctionne avec plusieurs fournisseurs de modèles et utilise une architecture client-serveur. Vous pouvez configurer ses outils, permissions, sessions et fournisseurs. Ses plugins étendent un noyau serveur fixe, tandis que DeepSeek rend aussi la boucle et le magasin de sessions remplaçables.

## Quand utiliser DeepSeek Harness

Remplacer des parties du runtime n’a pas d’intérêt en soi. Ce contrôle supplémentaire ne compte que s’il résout un problème que vous avez déjà.

- **Quand le runtime fait partie du projet.** Si vous modifiez des adaptateurs de modèle, la boucle d’agent, le stockage ou le comportement des sessions — pas seulement si vous construisez au-dessus d’un agent —, c’est un meilleur choix.
- **Quand vous comparez des modèles en environnement contrôlé.** Utiliser le même runtime fixe davantage de paramètres lors des échanges de modèle, même si les modèles diffèrent encore par l’usage des outils et le style de raisonnement.
- **Quand le débogage d’une exécution complexe est crucial.** Le journal d’événements de session et la vue Trajectory facilitent la reconstitution de ce que le modèle a vu et quels outils ont tourné.
- **Quand vous testez l’interne des agents.** Le mode Creator et Cordis s’adressent aux développeurs qui étudient la composition des agents, plus qu’à ceux qui veulent seulement générer du code applicatif.

