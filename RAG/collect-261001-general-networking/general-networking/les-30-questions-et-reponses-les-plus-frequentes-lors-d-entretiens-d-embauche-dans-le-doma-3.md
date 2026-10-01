---
id: collect-261001-general-networking/general-networking/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-doma-3
title: "les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "attention"]
source: docs/RAG/collect-261001-general-networking/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i.md
source_anchor: ""
source_lines: [136, 183]
sha256: 5326b104ec907378a9847ce29354b1235a462d27a10355216280a41fd21810f3
---

# les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i

Une approche consiste à définir une hiérarchie d'objectifs et à attribuer des pondérations ou des règles qui guident l'agent dans le choix de la tâche à accomplir en cas de conflit. Certains systèmes utilisent également des composants de planification ou des étapes de raisonnement intermédiaires (telles que des boucles de réflexion ou des blocs-notes) pour évaluer les compromis avant d'agir.

Si vous êtes novice en la matière, je vous recommande de commencer par l'article d'Anthropic sur les modèles de conception d'agents. Il fournit des exemples concrets et des architectures courantes utilisées dans des systèmes réels. Si vous avez une formation en génie logiciel, bon nombre de ces concepts vous seront familiers, en particulier ceux liés à la conception modulaire, à la gestion d'état et à l'exécution asynchrone des tâches.

### Dans quelle mesure êtes-vous à l'aise avec l'invite et l'ingénierie des invites ? Quelles approches avez-vous entendues ou utilisées ?

L'ingénierie des invites est un élément majeur d'un système d'IA agentique, mais c'est également un sujet qui tend à susciter des stéréotypes. Il est donc important d'éviter les déclarations vagues sur son importance et de se concentrer plutôt sur les détails techniques de son application.

Voici ce que je considérerais comme une réponse appropriée :

Je suis tout à fait à l'aise avec le prompting et le prompt engineering, et j'ai utilisé plusieurs techniques tant dans le cadre de projets que dans mes tâches quotidiennes. Par exemple, j'utilise régulièrement la méthode « few-shot prompting » pour orienter les modèles vers un format ou un ton spécifique en fournissant des exemples. J'utilise également la chaîne de pensée lorsque j'ai besoin que le modèle raisonne étape par étape, ce qui est particulièrement utile pour des tâches telles que le codage, les casse-têtes logiques ou la planification.

Dans des applications plus structurées, j'ai expérimenté le réglage rapide. ajustement des invites et la compression des invites, en particulier lorsque je travaille avec des API qui facturent au nombre de jetons ou qui nécessitent un contrôle strict des résultats. Ces techniques consistent à réduire les invites à leurs éléments les plus essentiels tout en préservant leur intention et leur performance.

Étant donné que le domaine évolue rapidement, je prends l'habitude de lire les articles récents, les dépôts GitHub et les mises à jour de la documentation, afin de me tenir au courant des techniques telles que l'appel de fonction. appel de fonction, le prompting augmenté par la récupération et le chaînage modulaire de chaînage modulaire des invites.

### Qu'est-ce qu'une fenêtre contextuelle ? Pourquoi sa taille est-elle limitée ?

Une fenêtre contextuelle désigne la quantité maximale d'informations, mesurée en tokens, qu'un modèle linguistique peut traiter simultanément. Cela inclut l'invite actuelle, l'historique des conversations précédentes et les instructions au niveau du système. Une fois la limite de la fenêtre contextuelle atteinte, les jetons plus anciens peuvent être tronqués ou ignorés.

La raison pour laquelle la fenêtre contextuelle est limitée est due à des contraintes informatiques et architecturales. Dans les modèles basés sur des transformateurs, les mécanismes d'attention nécessitent le calcul des relations entre tous les tokens du contexte, ce qui augmentent de manière quadratique avec le nombre de jetons. Cela rend le traitement de contextes très longs coûteux et lent, en particulier sur le matériel actuel. Les modèles antérieurs, tels que les RNN, ne comportaient pas de limite contextuelle stricte de la même manière, mais ils rencontraient des difficultés pour conserver efficacement les dépendances à long terme.

### Qu'est-ce que la génération augmentée par la récupération (RAG) ?

La génération augmentée par la récupération (RAG) est une technique qui améliore les modèles linguistiques en leur permettant de récupérer des informations pertinentes à partir de sources externes avant de générer une réponse. Au lieu de se baser uniquement sur ce que le modèle a appris pendant la formation, les systèmes RAG peuvent accéder à des données actualisées ou spécifiques à un domaine au moment de l'inférence.

Une configuration RAG typique comprend deux composants principaux : un récupérateur, qui recherche dans une base de données ou une collection de documents le contexte pertinent en fonction de la requête saisie, et un générateur, qui utilise les informations récupérées pour produire une réponse plus précise et mieux informée. Cette approche est particulièrement utile pour les tâches qui exigent une précision factuelle, une mémoire à long terme ou des connaissances spécifiques à un domaine.

### Quelles autres architectures LLM connaissez-vous en dehors du transformateur ?

Bien que le transformateur soit l'architecture dominante dans le domaine de l'IA aujourd'hui, il existe plusieurs autres types de modèles qu'il est important de connaître. Par exemple, xLSTM s'appuie sur l'architecture LSTM avec des améliorations qui optimisent les performances sur les séquences longues tout en conservant l'efficacité.

Mamba est une autre architecture prometteuse : elle utilise des modèles d'espace d'état sélectifs pour traiter plus efficacement que les transformateurs les contextes longs, en particulier pour les tâches qui ne nécessitent pas une attention particulière à chaque token.

L'architecture Titans de Google mérite également d'être examinée. Il est conçu pour remédier à certaines des principales limites des transformateurs, telles que le manque de mémoire persistante et les coûts de calcul élevés.

Ces architectures alternatives visent à rendre les modèles plus efficaces, plus évolutifs et capables de traiter des entrées plus longues ou plus complexes sans nécessiter d'importantes ressources matérielles.

### Qu'entend-on par utilisation d'outils et appel de fonctions dans les LLM ?

L'appel d'outils et de fonctions permet aux modèles linguistiques de grande taille d'interagir avec des systèmes externes, tels que des API, des bases de données ou des fonctions personnalisées. Au lieu de s'appuyer uniquement sur des connaissances pré-acquises, le modèle est capable de reconnaître quand une tâche nécessite des informations actualisées ou spécialisées et de réagir en faisant appel à un outil approprié.

Par exemple, si vous demandez à un modèle ayant accès à une API météo « Quel temps fait-il à Londres ? », il peut décider d'appeler cette API en arrière-plan et renvoyer les données en temps réel au lieu de générer une réponse générique ou obsolète. Cette approche rend les modèles plus utiles et plus fiables, en particulier pour les tâches impliquant des données en temps réel, des calculs ou des actions dépassant les capacités internes du modèle.

### Qu'est-ce que la chaîne de pensée (CoT) et pourquoi est-elle importante dans les applications d'IA agentique ?

La chaîne de pensée (CoT) est une technique de guidage qui aide les modèles linguistiques à décomposer les problèmes complexes en raisonnements étape par étape avant de produire une réponse finale. est une technique d'incitation qui aide les modèles linguistiques à décomposer des problèmes complexes en un raisonnement étape par étape avant de produire une réponse finale. Cela permet au modèle de générer des étapes de raisonnement intermédiaires, ce qui améliore la précision et la transparence, en particulier pour les tâches impliquant la logique, les mathématiques ou la prise de décision en plusieurs étapes.

