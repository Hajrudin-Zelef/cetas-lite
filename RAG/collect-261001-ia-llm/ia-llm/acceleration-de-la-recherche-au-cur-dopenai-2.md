---
id: collect-261001-ia-llm/ia-llm/acceleration-de-la-recherche-au-cur-dopenai-2
title: "acceleration-de-la-recherche-au-cur-dopenai"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agents", "agi", "arr", "astra", "gpu", "valuation"]
source: docs/RAG/collect-261001-ia-llm/acceleration-de-la-recherche-au-cur-dopenai.md
source_anchor: ""
source_lines: [43, 91]
sha256: 19ab3b52b95f62baf1d5255905816e56ab4957dd76fb4f9ba39580e602de9902
---

# acceleration-de-la-recherche-au-cur-dopenai

Pour éclairer cette tendance, nous avons analysé les usages récents en recherche avec une taxonomie récemment publiée(ouverture dans une nouvelle fenêtre) par Epoch AI des tâches du cycle de R&D en IA. Inspirée du système de classification des métiers O*NET, cette taxonomie adaptée à la R&D en IA de pointe distingue six phases :

Décider : sujets, poursuite des travaux, allocation des ressources

Concevoir : idées de recherche et spécifications techniques

Construire : code et jeux de données

Exécuter : entraînement, évaluations, matériel, service d’inférence

Analyser : expériences, modèles, déploiement, travaux externes

Ci-dessous, nous classons les tokens des agents de programmation selon cette taxonomie.

Toutes les catégories ont progressé de janvier à août 2026. En janvier, le code de recherche et d’infrastructure dominait. Cette catégorie a grandi, mais d’autres aussi, surtout l’assistance technique et la surveillance des exécutions. La planification de haut niveau reste une part minime des tokens de sortie des agents.

Selon nos collègues, les agents excellent dans le dépannage de l’infrastructure de recherche interne, levant un frein important. Plusieurs équipes proposant des permanences de dépannage ont vu leur fréquentation baisser en 2026. L’une les a supprimées pour se consacrer à d’autres améliorations.

Ce graphique montre les messages initiaux quotidiens d’un des principaux canaux internes où les chercheurs sollicitent l’assistance technique d’autres équipes. À notre connaissance, cette baisse n’a pas été compensée par un transfert vers un autre canal d’assistance assuré par des humains. La baisse du trafic concorde avec cette évolution générale.

Nous pouvons aussi étudier si les agents réussissent les tâches confiées par les chercheurs. Un classificateur agentique montre que, de janvier à juillet, la réussite a généralement progressé dans plusieurs niveaux de difficulté, estimés par la durée humaine, pour les tâches au résultat vérifiable. Les agents ont toutefois encore besoin d’un guidage humain important, surtout pour les tâches complexes. Ces 6 derniers mois, plus de la moitié des tâches de 4 à 8 heures réussies ont nécessité au moins une intervention.

4. Réguler le rythme de développement des modèles

Les progrès vers des systèmes plus performants au service d’une AGI sûre et bénéfique dépendront aussi des garde-fous nécessaires à ces travaux. Notre évaluation des garde-fous nécessaires peut évoluer à mesure que nous comprenons mieux les risques.

Comme nous l’avons expliqué, nous avons récemment actualisé nos normes de surveillance, d’alignement et de sécurité. Nous montrons ici comment les restrictions récentes ont affecté un aspect de l’activité de recherche.

Le 20 juillet, après avoir découvert que des agents avaient compromis notre infrastructure de recherche, nous avons temporairement arrêté le service de conteneurs utilisé pour l’entraînement, puis l’avons rétabli avec d’importantes restrictions supplémentaires.

Cela a entraîné une forte baisse du calcul consacré à l’entraînement en RL, le temps que les équipes reconfigurent leurs flux de travail pour opérer dans cet environnement de recherche renforcé. Le graphique ci-dessus inclut la suspension de deux semaines de l’apprentissage par renforcement sur nos derniers modèles destinés au déploiement. Entre le 20 juillet et le 6 août, la majorité des exécutions des expériences de RL de classe Astra, en termes d’allocation de GPU, visait à tester la mise en œuvre d’améliorations de sûreté et de sécurité.

Le 7 août, des indices préliminaires suggérant qu’Astra pourrait posséder des cybercapacités critiques au sens de notre cadre de préparation(ouverture dans une nouvelle fenêtre) ont conduit à des restrictions de sécurité supplémentaires propres à ce modèle, imposant son exécution dans des environnements de recherche plus sécurisés. La semaine suivante, l’allocation de GPU à la classe Astra a encore baissé de 59,2 %, tandis que celle des autres classes de modèles a augmenté de 17,2 %. Cette hausse a compensé environ 85 % de la baisse de la classe Astra, laissant l’allocation totale des charges de travail de RL analysées largement inchangée. Cette tendance concorde avec le transfert d’une partie de l’entraînement et de l’expérimentation vers des modèles non-Astra pendant les restrictions touchant Astra. Elle rejoint les témoignages de chercheurs ayant trouvé d’autres usages aux ressources de calcul qui ne pouvaient plus servir aux charges de travail soumises aux nouvelles contraintes.

Ces données éclairent utilement les débats actuels sur l’entraînement et la sécurité : lorsque de nouveaux contrôles sont instaurés, les ressources de calcul restent précieuses et flexibles, et sont naturellement réorientées vers d’autres usages au sein de la recherche. À plus long terme, les discussions sur le rythme des progrès de l’IA devraient aussi aborder la meilleure façon d’utiliser les ressources de calcul soumises à des contrôles nouveaux ou envisagés.

5. La voie à suivre

Progresser vers une RSI alignée et comprendre ces progrès est essentiel à notre mission. Nous continuerons d’affiner nos méthodes, de rendre compte de l’évolution de notre compréhension et d’œuvrer pour un débat public éclairé et une véritable gouvernance démocratique des systèmes de pointe.

Annexe : méthodes utilisées pour cet article

La recherche en IA appuyée par des agents est encore nouvelle, et nous apprenons toujours à la mesurer. Certains indicateurs, comme la quantité de code générée par nos équipes de recherche, sont relativement faciles à recueillir, mais difficiles à interpréter, car leur lien avec les progrès de la recherche est incertain. Des indicateurs plus directement centrés sur les progrès de la recherche, comme la fréquence à laquelle les agents réussissent les tâches confiées par les chercheurs, pourraient être plus utiles, mais sont complexes à élaborer et à valider. L’évolution rapide des outils et systèmes utilisés par les chercheurs ajoute à cette difficulté. Approfondir notre compréhension de l’accélération de la recherche constitue un axe majeur dans l’ensemble d’OpenAI.

Dans toutes ces analyses, sauf indication contraire :

Le terme « chercheur » désigne au sens large tout membre de notre organisation de recherche, y compris ceux qui développent l’infrastructure, gèrent les projets ou soutiennent l’activité d’une autre manière.

Les indicateurs d’utilisation des agents de programmation couvrent la plupart des usages, mais pas tous, compte tenu de l’évolution rapide des outils et systèmes utilisés par les chercheurs.
