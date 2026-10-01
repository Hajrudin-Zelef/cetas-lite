---
id: collect-261001-ia-llm/ia-llm/modeles-de-langage-aux-origines-des-hallucinations-2
title: "modeles-de-langage-aux-origines-des-hallucinations"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-ia-llm/modeles-de-langage-aux-origines-des-hallucinations.md
source_anchor: ""
source_lines: [57, 80]
sha256: 9ad5764e633890fb81f651ca4d8047e3661d794d0559914d77a63883ecf7d35b
---

# modeles-de-langage-aux-origines-des-hallucinations

Il est deux fois plus difficile de faire la distinction entre les affirmations valides et non valides sans exemples étiquetés d’affirmations non valides. Mais même avec les étiquettes, certaines erreurs restent inévitables. Pour bien en comprendre les raisons, basons nous sur une nouvelle analogie. Dans le domaine de la reconnaissance d’images, l’étiquetage de millions de photos de chats et de chiens permet aux algorithmes de les classer de manière fiable. Imaginons qu’au lieu d’étiqueter chaque photo en fonction de son sujet (chien ou chat), nous indiquions la date d’anniversaire de l’animal. Ces dates étant aléatoires, cette tâche générerait toujours des erreurs, quel que soit le degré de sophistication de l’algorithme.

Il en va de même pour le pré-entraînement. L’orthographe et l’organisation des parenthèses suivent une logique. Les erreurs sont donc éliminées à mesure que les volumes de données augmentent. A contrario, les faits aléatoires dont la fréquence est faible, comme la date d’anniversaire d’un animal, ne peuvent pas être prédits par une logique quelconque et génèrent donc des hallucinations. Notre analyse explique les types d’hallucinations qui résultent de la prédiction du mot suivant. Dans l’idéal, de nouvelles étapes suivant le pré-entraînement devraient pouvoir les éliminer, mais ce n’est aujourd’hui pas parfaitement le cas pour les raisons décrites dans la section précédente.

Conclusions

Nous espérons que l’explication statistique de notre étude clarifie la nature des hallucinations et bat en brèche diverses idées fausses, par exemple :

Affirmation : les hallucinations disparaîtront si nous améliorons l’exactitude des modèles, car un modèle obtenant un score d’exactitude de 100 % ne peut pas halluciner.
Constatation : l’exactitude des modèles n’atteindra jamais 100 %, car quelles que soient sa taille et ses capacités de recherche et de raisonnement,un modèle ne pourra jamais répondre à certaines des questions qui lui sont posées dans le monde réel.

Affirmation : les hallucinations sont inévitables.
Constatation : ce n’est pas vrai, car les modèle de langage peuvent choisir de ne pas répondre en cas d’incertitude.

Affirmation : éviter les hallucinations impose un niveau d’intelligence qui n’est atteignable qu’avec les plus grands modèles.
Constatation : il peut au contraire être plus facile pour un petit modèle de déterminer ses limites. Par exemple, en réponse à une question portant sur le maori, un petit modèle qui ne parle pas le maori pourrait simplement répondre « Je ne sais pas », tandis qu’un modèle qui connaît un peu de maori doit d’abord déterminer son niveau de confiance. Comme l’explique l’étude, la « calibration » demande bien moins de ressources de calcul que la fourniture d’une réponse exacte.

Affirmation : les hallucinations constituent un bug mystérieux des modèles de langage modernes.
Constatation : nous comprenons les mécanismes statistiques qui donnent naissance aux hallucinations et les récompensent lors des évaluations.

Affirmation : pour mesurer les hallucinations, nous avons simplement besoin d’une bonne évaluation spécialisée.
Constatation : des évaluations centrées sur les hallucinations ont déjà été publiées. Pour autant, une bonne évaluation n’a que peu d’effet après les centaines d’évaluations classiques basées sur l’exactitude qui pénalisent l’humilité et récompensent les hypothèses. Il convient plutôt de repenser tous les indicateurs des évaluations principales pour récompenser l’expression de l’incertitude.

Nos derniers modèles de langage présentent des taux d’hallucination réduits, et nous ne cessons de les améliorer.
