---
id: collect-261001-ia-llm/ia-llm/quest-ce-quun-agent-harness-guide-pour-debutants-1
title: "quest-ce-quun-agent-harness-guide-pour-debutants"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "mcp", "model context protocol", "valuation"]
source: docs/RAG/collect-261001-ia-llm/quest-ce-quun-agent-harness-guide-pour-debutants.md
source_anchor: ""
source_lines: [1, 80]
sha256: f2285795b6b3b7e8562418a14b04ab33f5803a83606c9f4e4c296e0badf85ce6
---

# quest-ce-quun-agent-harness-guide-pour-debutants

Cours

L’idée n’est pas nouvelle. Les développeurs créent depuis des années des wrappers, des échafaudages et des environnements d’exécution autour des modèles. L’expression s’est imposée après que Mitchell Hashimoto, cofondateur de HashiCorp, a parlé de « harness engineering » dans un billet de blog de février 2026 sur son workflow IA. Son point était simple : lorsqu’un agent commet une erreur, modifiez l’environnement pour que l’erreur ne puisse plus se reproduire. OpenAI a repris le terme la même semaine pour ses travaux sur Codex, et LangChain a suivi avec le même cadrage.

Dans cet article, j’explique ce qu’est un agent harness, pourquoi les agents d’IA en ont besoin, en quoi il diffère des frameworks et des runtimes, et quels outils les développeurs utilisent pour bâtir des systèmes de type harness.

## Qu’est-ce qu’un agent harness ?

Une définition vient de LangChain : « Si vous n’êtes pas le modèle, vous êtes le harness. » En pratique, un agent harness est le logiciel qui entoure un modèle de langage : outils, mémoire, état, exécution, garde-fous et observabilité.

Agent = Modèle + Harness

Le modèle raisonne. Le harness lui fournit un environnement pour agir, se souvenir, vérifier les résultats et suivre des règles.

*Modèle au sein de son agent harness opérationnel. Image par l’auteur.*

La formule est utile, mais reste un modèle mental, pas une norme industrielle. Certains éditeurs utilisent encore « harness », « framework » et « scaffold » pour désigner à peu près la même chose.

## Pourquoi les agents d’IA ont besoin d’un harness

Un modèle de langage brut atteint vite ses limites quand vous lui demandez de travailler sur de nombreuses étapes. Il ne maintient pas un état durable tout seul, n’exécute pas de lui-même des outils, ne gère pas un contexte qui s’allonge, et ne se remet pas d’un échec d’appel d’outil sans aide.

Imaginez un agent chargé de corriger un test défaillant dans un projet Python. Sans harness, le modèle peut proposer une correction en apparence, mais il ne peut ni lire le vrai fichier de test, ni lancer pytest, ni voir l’erreur réelle, ni éditer la fonction fautive, ni confirmer que la correction passe. Avec un harness, toute cette boucle devient quelques minutes de travail que l’agent réalise seul, chaque étape étant consignée pour inspection humaine.

La recommandation d’Anthropic reste valable : commencez par l’approche la plus simple possible et n’ajoutez des pièces mobiles que lorsque la tâche l’exige.

## De quoi se compose un agent harness

Les composants varient, mais la plupart partagent quelques briques communes. Voyez-les comme une check-list, pas comme un cahier des charges strict. Un petit agent n’aura besoin que d’une partie de ces éléments, tandis qu’un agent en production en nécessitera davantage.

### Prompts système et règles de comportement

Le harness contrôle généralement les instructions de base du modèle. Cela inclut le prompt système, mais aussi les règles projet, les standards de code, les contraintes de rôle et les politiques de sécurité. Dans Deep Agents de LangChain, par exemple, un fichier `AGENTS.md` peut poser le cadre avant le démarrage d’une tâche.

Certains harnesses en 2026 utilisent aussi une divulgation progressive des instructions. Plutôt que de charger au démarrage la description complète de chaque outil, le harness n’ajoute qu’un résumé de ce qui est disponible. La notice détaillée d’un outil n’est chargée que lorsque le modèle en a besoin.

### Outils : comment les agents interagissent avec le monde

Les outils permettent à l’agent d’aller au-delà de la simple génération de texte. Exemples courants : recherche web, lecture/écriture de fichiers, requêtes de base de données, appels d’API, actions navigateur, exécution de code et commandes terminal. Le harness contrôle quels outils sont disponibles, quand le modèle est autorisé à les appeler, et comment les résultats sont formatés et réinjectés dans le contexte de l’agent.

Model Context Protocol (MCP) est devenu en 2026 l’interface standard pour cela. De nombreux harnesses, dont Anthropic Agent SDK, LangChain Deep Agents et OpenAI Agents SDK, utilisent MCP pour connecter des serveurs d’outils externes sans écrire d’intégration spécifique pour chacun.

### Mémoire et état

Les agents doivent savoir ce qui s’est passé plus tôt dans une tâche. Un harness peut conserver l’état de court terme dans la conversation active et l’état de long terme dans des fichiers, journaux, résumés ou préférences enregistrées. Certains compressent aussi de longues histoires en résumés pour éviter de surcharger le contexte.

### Environnement d’exécution : où l’agent s’exécute et agit

Beaucoup d’agents utiles ont besoin d’un véritable lieu de travail : un système de fichiers, un conteneur, un terminal isolé, une instance de navigateur ou un runtime cloud. Sans environnement d’exécution géré par le harness, les appels d’outils n’ont nulle part où atterrir.

Nombre de harnesses utilisent désormais des conteneurs bac à sable isolés : des environnements éphémères limités à une session, nettoyés en fin de tâche, afin d’éviter que les écritures de fichiers, installations de paquets et appels réseau d’une tâche ne débordent sur une autre.

### Orchestration et planification

Certaines tâches ne se prêtent pas à une suite d’étapes linéaires. Le harness peut fournir un outil de planification qui décompose un objectif en sous-tâches et en suit l’avancement. Il peut aussi lancer des sous-agents dédiés à une partie du travail et ne renvoyer au principal qu’un résumé.

LangChain Deep Agents, par exemple, suit les étapes du plan dans un fichier du système de fichiers, en les passant de « en attente » à « terminé » au fil de l’exécution.

### Garde-fous et autorisations

Le harness est l’endroit où poser les règles : approbation humaine, blocage d’appels d’outils, permissions basées sur les rôles et contrôles de sortie. OpenAI Agents SDK, LangChain Deep Agents et Microsoft Agent Framework prennent en charge ce type de contrôle. Le schéma le plus sûr consiste à vérifier séparément les entrées, les sorties et les permissions d’outils.

### Observabilité et traçage

Quand une tâche de cinquante étapes échoue à la trente-septième, un trace révèle ce qui s’est passé. Le traçage enregistre appels modèle, appels d’outils, passages de relais, erreurs, latence et coût sur l’ensemble du run. L’OpenAI Agents SDK active le traçage par défaut. LangSmith ajoute des tableaux de bord de débogage et d’évaluation. OpenTelemetry est devenu la norme d’export de traces au format neutre, pour éviter l’enfermement dans un outil d’observabilité.

## Agent harness vs framework vs runtime : quelle différence ?

La question revient souvent, et la réponse est plus subtile que ne le laissent entendre certains articles. La taxonomie est utile, mais elle n’est pas figée.

*Trois couches, abstraction croissante de bas en haut. Image par l’auteur.*

Commençons par le framework, que beaucoup de développeurs ont déjà utilisé.

### Qu’est-ce qu’un framework d’agent ?

Un framework d’agent offre aux développeurs des briques pour créer des agents. Il couvre les appels modèle, la définition d’outils, les schémas de mémoire et la boucle d’agent. Exemples : les premières versions de LangChain, CrewAI et Google ADK. Un framework vous indique comment structurer un agent, mais pas toujours comment l’exécuter de façon fiable en production.

### Qu’est-ce qu’un runtime d’agent ?

