---
id: collect-261001-general-networking/general-networking/comprendre-les-reseaux-neuronaux-grace-aux-circuits-clairsemes-1
title: "comprendre-les-reseaux-neuronaux-grace-aux-circuits-clairsemes"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "valuation"]
source: docs/RAG/collect-261001-general-networking/comprendre-les-reseaux-neuronaux-grace-aux-circuits-clairsemes.md
source_anchor: ""
source_lines: [1, 40]
sha256: 4a57e7e67a323f82764aa8f8d8f9d9ef685c615617cacd3ab70c082560dfb669
---

# comprendre-les-reseaux-neuronaux-grace-aux-circuits-clairsemes

Les réseaux neuronaux alimentent les systèmes d’IA les plus performants d’aujourd’hui, mais ils restent difficiles à comprendre. Nous ne rédigeons pas ces modèles avec des instructions explicites et détaillées, étape par étape. Au lieu de cela, ils apprennent en ajustant des milliards de connexions internes, ou « poids », jusqu’à maîtriser une tâche. Nous concevons les règles de l’entraînement, mais pas les comportements spécifiques qui en découlent, et le résultat est un réseau dense de connexions qu’aucun être humain ne peut facilement déchiffrer.

Notre vision de l’interprétabilité

À mesure que les systèmes d’IA deviennent plus performants et ont un impact réel sur les décisions dans les domaines de la science, de l’éducation et des soins de santé, il est essentiel de comprendre leur fonctionnement. L’interprétabilité désigne les méthodes qui nous aident à comprendre pourquoi un modèle a produit une sortie donnée. Il existe de nombreuses manières d’y parvenir.

Par exemple, les modèles de raisonnement sont encouragés à expliquer leur raisonnement jusqu’à la réponse finale. L’interprétabilité de la chaîne de pensée utilise ces explications pour surveiller le comportement du modèle. Ceci est immédiatement utile : les chaînes de pensée des modèles de raisonnement actuels semblent être instructives en ce qui concerne les comportements préoccupants tels que la tromperie. Toutefois, le fait de s’appuyer entièrement sur cette propriété est une stratégie fragile, qui risque de se dégrader au fil du temps.

D’autre part, l’interprétabilité mécanistique, qui est l’objet de ce travail, cherche à inverser complètement les calculs d’un modèle. Jusqu’à présent, elle s’est révélée moins utile dans l’immédiat, mais pourrait, en principe, offrir une explication plus complète du comportement du modèle. En cherchant à expliquer le comportement du modèle au niveau le plus granulaire, l’interprétabilité mécanistique permet de faire moins d’hypothèses et renforcer notre confiance. Mais passer des détails de bas niveau à l’explication de comportements complexes est un cheminement long et difficile.

L’interprétabilité soutient plusieurs objectifs clés, par exemple en permettant un meilleur suivi et en fournissant des signes avant-coureurs en cas de comportement dangereux ou stratégiquement mal aligné. Elle complète également nos autres efforts en matière de sécurité, tels que la surveillance évolutive, la formation contradictoire et l’équipe d’intervention.

Dans ce travail, nous montrons qu’il est souvent possible d’entraîner des modèles de manière à les rendre plus faciles à interpréter. Nous voyons notre travail comme un complément prometteur à l’analyse post-hoc des réseaux denses.

C’est un pari ambitieux ; il y a encore un long chemin à parcourir avant de bien comprendre les comportements complexes des modèles les plus puissants. Néanmoins, pour les comportements les plus simples, nous constatons que les modèles clairsemés entraînés avec notre méthode contiennent de petits circuits désenchevêtrés qui sont à la fois compréhensibles et suffisants pour exécuter le comportement. Cela suggère qu’il pourrait y avoir une voie praticable pour entraîner des systèmes plus importants dont nous pouvons comprendre les mécanismes.

Une nouvelle approche : l’apprentissage des modèles clairsemés

Les travaux précédents sur l’interprétabilité mécanistique ont commencé à partir de réseaux denses et enchevêtrés, et ont tenté de les démêler. Dans ces réseaux, chaque neurone est connecté à des milliers d’autres neurones. La plupart des neurones semblent remplir de nombreuses fonctions distinctes, ce qui rend leur compréhension apparemment impossibles.

Mais que se passerait-il si nous entraînions des réseaux neuronaux non enchevêtrés comportant beaucoup plus de neurones, mais où chaque neurone n’a que quelques dizaines de connexions ? Le réseau qui en résulterait serait peut-être plus simple et plus facile à comprendre. C’est le principal objectif de recherche de notre travail.

En gardant ce principe à l’esprit, nous avons entraîné des modèles de langage avec une architecture très similaire à celle des modèles de langage existants tels que GPT‑2, avec une petite modification : nous contraignons la grande majorité des poids du modèle à être nuls. Cela contraint le modèle à n’utiliser qu’un très petit nombre des connexions possibles entre les neurones. Il s’agit d’une modification simple qui, selon nous, désenchevêtre de manière substantielle les calculs internes du modèle.

Dans les réseaux neuronaux denses normaux, chaque neurone est connecté aux neurones de la couche suivante. Dans les modèles clairsemés, les neurones ne se connectent qu’à quelques neurones de la couche suivante. Nous espérons que cela facilitera la compréhension des neurones et du réseau dans son ensemble.

Évaluation de l’interprétabilité

Nous souhaitons mesurer le degré de démêlage des calculs de nos modèles clairsemés. Nous avons envisagé divers comportements simples du modèle et vérifié si nous pouvions isoler les parties du modèle responsables de chaque comportement, que nous appelons circuits.

Nous avons soigneusement sélectionné une série de tâches algorithmiques simples. Pour chacun, nous avons réduit le modèle au plus petit circuit capable d’effectuer la tâche, et nous avons examiné la simplicité de ce circuit. (Pour en savoir plus, veuillez consulter cet article(ouverture dans une nouvelle fenêtre).) Nous avons constaté qu’en utilisant un entraînement de modèles plus grands et plus clairsemés, nous pouvions produire des modèles de plus en plus performants avec des circuits de plus en plus simples.

Nous représentons l’interprétabilité en fonction de la capacité des modèles (en bas à gauche étant la meilleure). Pour une taille de modèle clairsemée fixe, l’augmentation de la parcimonie (définition d’un plus grand nombre de pondérations à zéro) réduit la capacité, mais augmente l’interprétabilité. L’augmentation de la taille du modèle repousse cette limite, ce qui suggère que nous pouvons construire des modèles plus grands, à la fois performants et interprétables.

Pour rendre cela plus concret, considérez une tâche dans laquelle un modèle formé au code Python doit compléter une chaîne de caractères avec le bon type de guillemets. En Python, 'hello' doit se terminer par une apostrophe, et "hello" doit se terminer par des guillemets doubles. Le modèle peut résoudre ce problème en se souvenant du type de guillemet qui a ouvert la chaîne et en le reproduisant à la fin.

Nos modèles les plus interprétables semblent contenir des circuits démêlés qui mettent en œuvre cet algorithme précis.

Exemple de circuit dans un transformateur clairsemé qui prédit si une chaîne de caractères doit se terminer par des guillemets simples ou doubles. Ce circuit utilise seulement cinq canaux résiduels (lignes verticales grises), deux neurones MLP dans la couche 0, et un canal de requête-clé d’attention et un canal de valeur dans la couche 10. Le modèle (1) encode les guillemets simples dans un canal résiduel et les guillemets doubles dans un autre ; (2) utilise une couche MLP pour le convertir en un canal qui détecte tout guillemet et un autre qui classe entre guillemets simples et doubles ; (3) utilise une opération d’attention pour ignorer les jetons intermédiaires, trouver le guillemet précédent et copier son type dans le jeton final ; et (4) prédit le guillemet fermant correspondant.

