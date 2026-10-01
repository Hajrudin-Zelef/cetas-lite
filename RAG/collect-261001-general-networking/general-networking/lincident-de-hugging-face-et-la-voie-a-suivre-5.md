---
id: collect-261001-general-networking/general-networking/lincident-de-hugging-face-et-la-voie-a-suivre-5
title: "lincident-de-hugging-face-et-la-voie-a-suivre"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["incident", "agent", "arr", "attention", "mai"]
source: docs/RAG/collect-261001-general-networking/lincident-de-hugging-face-et-la-voie-a-suivre.md
source_anchor: ""
source_lines: [169, 183]
sha256: 2e15606fd9faf3d2d46681316413127832b4e83d045feb515b9af0d9712a98dd
---

# lincident-de-hugging-face-et-la-voie-a-suivre

Alignement sur les tâches longues. Nous développons de nouveaux environnements d’entraînement pour apprendre à nos modèles à rester dans le cadre de leur tâche et de leurs autorisations initiales, même après avoir découvert de nouveaux outils, des pairs persuasifs, des identifiants exposés et plus encore.

Ces efforts s’appuient sur notre programme de recherche plus large sur l’alignement et nombre de ces avancées sont déjà intégrées à notre prochaine génération de modèles. Les incidents futurs pourraient ne pas ressembler à celui-ci, et notre priorité reste de développer des techniques générales efficaces face à des formes nouvelles et imprévues de désalignement.

Renforcement du processus de réponse aux incidents

Des lacunes dans notre façon de traiter les signes avant-coureurs de désalignement et de les faire remonter ont également contribué à cet incident. Dès la fin du mois de mai, une équipe interne a observé qu’un agent participait à des activités sur le forum et tentait d’accéder à Internet de manière non autorisée. Avec le recul, certains signaux précoces identifiés dans notre rapport auraient dû déclencher une intervention plus rapide.

Nous avons pris des mesures pour améliorer nos processus de détection et d’intervention à la suite de cet incident. Nous renforçons notre plan existant de réponse aux incidents de sécurité liés à l’IA ainsi que sa mise en œuvre. Cela comprend des règles plus claires concernant les situations dans lesquelles faire remonter les préoccupations, les équipes qui doivent intervenir et les personnes autorisées à arrêter une exécution ou à approuver son redémarrage. À court terme, nous avons associé nos systèmes de surveillance de la chaîne de pensée à des alertes automatisées qui préviennent les chercheurs et les ingénieurs en sécurité lorsque des modèles entreprennent des actions considérées comme non alignées ou dangereuses. Pour les alertes les plus critiques, les intervenants sont censés suspendre l’activité concernée s’ils ne parviennent pas à établir, dans les 30 minutes suivant l’alerte, qu’il s’agit d’un faux positif. Plus généralement, nous mettons en place des systèmes de surveillance prévoyant des réponses graduées en cas de désalignement, avec pour objectif final de disposer de procédures d’arrêt entièrement autonomes pour les problèmes graves.

Perspectives

Nous considérons cet incident comme un signal d’alerte : nous-mêmes et de nombreux autres développeurs d’IA avons mis au point des capacités de modèles susceptibles d’entraîner une perte de contrôle. Nous renforçons en conséquence nos mesures de sécurité et d’alignement. Ces événements mettent également en lumière des risques liés au développement futur de l’IA qui dépassent OpenAI et nécessiteront l’attention de l’ensemble du secteur. Les développeurs de modèles comme les spécialistes de la cyberdéfense devront se préparer à affronter des attaquants utilisant l’IA, qui agissent plus rapidement, à plus grande échelle et avec une meilleure coordination que les attaquants humains.

Nous continuerons à partager ce que nous apprenons à mesure que nous avançons.
