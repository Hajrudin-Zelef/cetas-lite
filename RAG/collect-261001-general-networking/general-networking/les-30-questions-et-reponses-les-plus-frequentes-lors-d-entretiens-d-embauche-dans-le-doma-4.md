---
id: collect-261001-general-networking/general-networking/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-doma-4
title: "les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agents", "attention", "deepseek", "distillation", "rlhf", "valuation"]
source: docs/RAG/collect-261001-general-networking/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i.md
source_anchor: ""
source_lines: [184, 229]
sha256: 11888a96ebb0d78e94dbf1c40c37d44adf472618702e55556f48e1de08f50555
---

# les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i

Le CoT est largement utilisé dans les systèmes d'IA agentique. Par exemple, lorsqu'un modèle agit en tant que juge dans une évaluation, vous pouvez lui demander d'expliquer sa réponse étape par étape afin de mieux comprendre son processus décisionnel. Le CoT est également une technique fondamentale dans les modèles axés sur le raisonnement tels que OpenAI o1, où le modèle génère d'abord des jetons de raisonnement avant de les utiliser pour produire le résultat final. Ce processus de réflexion structuré rend le comportement des agents plus compréhensible et plus fiable.

### Qu'est-ce que le traçage ? Que sont les portées ?

Le traçage est le processus qui consiste à enregistrer et à visualiser la séquence d'événements qui se produisent lors d'une seule exécution ou d'un seul appel d'une application. Dans le contexte des applications LLM, une trace capture l'historique complet des interactions, telles que les multiples appels de modèles, l'utilisation d'outils ou les points de décision, au sein d'un flux d'exécution.

Une période correspond à un événement ou une opération unique au sein de cette trace. Par exemple, un appel de modèle, une invocation de fonction ou une étape de récupération seraient chacun enregistrés comme des intervalles individuels. Ensemble, les spans vous aident à comprendre la structure et le comportement de votre application.

Le traçage et les intervalles sont essentiels pour le débogage et l'optimisation des systèmes agentifs. Ils facilitent la détection des défaillances, des goulots d'étranglement liés à la latence ou des comportements indésirables. Des outils tels qu'Arize Phoenix et d'autres offrent des interfaces visuelles permettant d'examiner en détail les traces et les intervalles.

### Que sont les évaluations ? Comment évaluez-vous les performances et la robustesse d'un système d'IA agentique ?

Les évaluations constituent essentiellement les tests unitaires de l'ingénierie de l'IA agentique. Ils permettent aux développeurs d'évaluer les performances du système dans différents scénarios et cas limites. Il existe plusieurs types d'évaluations couramment utilisés aujourd'hui. Une approche consiste à utiliser un ensemble de données de référence élaboré manuellement afin de comparer les résultats du modèle aux réponses correctes connues.

Une autre approche consiste à utiliser un LLM comme juge pour évaluer la qualité, la précision ou le raisonnement derrière les réponses du modèle. Certaines évaluations mesurent la réussite globale de la tâche, tandis que d'autres se concentrent sur des éléments individuels tels que l'utilisation d'outils, la planification ou la cohérence. Les exécuter régulièrement permet d'identifier les régressions, de mesurer les améliorations et de garantir la fiabilité du système au fur et à mesure de son évolution. Pour approfondir le sujet, je vous recommande de consulter ce guide d'évaluation LLM.

### Pourriez-vous nous parler de l'architecture du transformateur et de son importance pour l'IA agentique ?

L'architecture du transformateur a été présentée dans l'article influent publié en 2017 « Attention Is All You Need ». Si vous ne l'avez pas encore lu, cela vaut la peine de le faire, car il a posé les bases de presque tous les grands modèles linguistiques modernes.

Depuis sa sortie, de nombreuses variantes et améliorations ont été développées, mais la plupart des modèles utilisés dans les systèmes d'IA agentique sont toujours basés sur une forme ou une autre du transformateur.

L'un des principaux avantages du transformateur réside dans son mécanisme d'attention, qui permet au modèle de calculer la pertinence de chaque jeton de la séquence d'entrée par rapport à tous les autres jetons, à condition que tout s'inscrive dans la fenêtre contextuelle. Cela permet d'obtenir d'excellentes performances dans les tâches qui nécessitent de comprendre des dépendances à long terme ou de raisonner à partir de plusieurs entrées.

En ce qui concerne spécifiquement l'IA agentique, la flexibilité et le parallélisme du transformateur le rendent particulièrement adapté à la gestion de tâches complexes telles que l'utilisation d'outils, la planification et les dialogues à plusieurs tours, qui constituent les comportements fondamentaux de la plupart des systèmes agentics actuels.

### Qu'est-ce que l'observabilité LLM et pourquoi est-elle importante ?

L'observabilité LLM désigne la capacité à surveiller, analyser et comprendre le comportement des systèmes de modèles linguistiques à grande échelle en temps réel. Il s'agit d'un terme générique qui englobe des outils tels que les traces, les spans et les evals, qui permettent aux développeurs de mieux comprendre le fonctionnement interne du système.

Les modèles d'apprentissage profond (LLM) étant souvent considérés comme des « boîtes noires », l'observabilité est essentielle pour le débogage, l'amélioration des performances et la garantie de la fiabilité. Il vous permet de suivre la manière dont les modèles interagissent entre eux et avec des outils externes, d'identifier les points de défaillance et de détecter rapidement les comportements inattendus. Dans les systèmes d'IA agentique, où plusieurs étapes et décisions sont enchaînées, l'observabilité est particulièrement importante pour maintenir la confiance et le contrôle.

### Pourriez-vous expliquer ce que sont le réglage fin et la distillation d'un modèle ?

Le réglage fin d'un modèle est le processus qui consiste à prendre un modèle pré-entraîné et à le perfectionner à l'aide d'un nouvel ensemble de données, généralement dans le but de le spécialiser pour un domaine ou une tâche spécifique. Cela permet au modèle d'adapter son comportement et ses réponses en fonction de connaissances plus ciblées ou actualisées.

La distillation de modèles est une technique connexe qui consiste à entraîner un modèle plus petit ou moins performant à partir des résultats d'un modèle plus grand et plus puissant. L'objectif est de transférer les connaissances et les comportements du modèle plus grand vers le plus petit, ce qui permet souvent d'obtenir des modèles plus rapides et plus efficaces avec des performances comparables. Par exemple, depuis la sortie de Deepseek R1, de nombreux modèles plus petits ont été développés à partir de ses réponses et ont atteint une qualité remarquable par rapport à leur taille.

### Quelle est la prochaine tâche de prédiction de jetons et pourquoi est-elle importante ? Que sont les modèles assistants ?

La prédiction du token suivant, également appelée modélisation autorégressive du langage, constitue la tâche d'entraînement principale derrière la plupart des grands modèles linguistiques. Le modèle est formé pour prédire le prochain token dans une séquence à partir de tous les tokens précédents. Cet objectif simple permet au modèle d'apprendre la grammaire, des faits, des schémas de raisonnement et même certaines capacités de planification. Le résultat de cette phase initiale de formation est appelé «modèle de base d' ».

Les modèles assistants sont des modèles de base qui ont été perfectionnés afin d'offrir une assistance plus efficace, plus sûre et plus conversationnelle. Ce réglage fin implique généralement des techniques telles que l'ajustement supervisé de l'instruction et l'apprentissage par renforcement avec rétroaction humaine (RLHF). apprentissage par renforcement avec rétroaction humaine (RLHF), qui guident le modèle pour qu'il réponde davantage comme un assistant plutôt que de se contenter de compléter un texte.

### Qu'est-ce que l'approche « human-in-the-loop » (HITL) ?

