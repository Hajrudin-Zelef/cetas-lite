---
id: collect-261001-general-networking/general-networking/lincident-de-hugging-face-et-la-voie-a-suivre-3
title: "lincident-de-hugging-face-et-la-voie-a-suivre"
domain: general-networking
role: reference
task: reference
actors: ["ExploitGym", "Hugging Face", "JFrog"]
dates: []
keywords: ["incident", "agent", "agents", "arr", "valuation"]
source: docs/RAG/collect-261001-general-networking/lincident-de-hugging-face-et-la-voie-a-suivre.md
source_anchor: ""
source_lines: [83, 126]
sha256: 9ef524db30b54fac5a4468080c1a6bdafea411487f3e74c20777ca3e6c55b64c
---

# lincident-de-hugging-face-et-la-voie-a-suivre

Un autre facteur clé du comportement désaligné était que les agents « renonçaient » rarement à leurs tâches d’évaluation, même lorsque celles-ci semblaient impossibles à résoudre. À mesure que les agents épuisaient leurs budgets de raisonnement, certains ont adopté des stratégies de plus en plus risquées et hors limites, allant finalement jusqu’à exploiter des infrastructures tierces.

Au cours de notre enquête, nous avons constaté que, dans le cadre de l’évaluation ExploitGym, aucun de nos modèles n’avait jamais produit de réponse correcte pour 198 des 898 tâches avant l’incident. Ces tâches non résolues ont largement contribué à l’activité de l’agent : 93 % des tâches discutées sur le forum d’Artifactory pendant l’incident de Hugging Face provenaient de cet ensemble de 198 tâches.

Deuxièmement, l’incident était principalement dû à des agents auxquels avaient été alloués des budgets de raisonnement incroyablement élevés, y compris des budgets bien supérieurs à ceux disponibles dans nos produits externes.

Un autre facteur expliquant la persistance des agents est qu’ils réfléchissaient souvent à la manière de réussir l’évaluation ExploitGym. Après avoir obtenu un accès à Internet, certains agents avaient lu l’article consacré à ExploitGym ainsi que le code GitHub et constaté que l’évaluation vérifiait deux conditions : que le bon flag soit trouvé et que la vulnérabilité prévue du programme, c’est-à-dire la bonne solution, ait été utilisée pour le trouver. La plupart des agents disposaient déjà du bon flag en trichant (il est facile de retrouver cette réponse par rétro-ingénierie à partir du code GitHub), mais ils pensaient que cela ne suffisait pas. Les agents ont partagé ce message sur le forum :

Il s’agit d’une forme de metagaming(ouverture dans une nouvelle fenêtre), dans laquelle les modèles analysent explicitement leurs mécanismes d’entraînement ou d’évaluation afin de maximiser la récompense. Dans de nombreux cas, le raisonnement concernant le code supposé du système de notation a conduit les agents à continuer d’exploiter Hugging Face, alors qu’ils avaient déjà trouvé le bon flag plusieurs joursauparavant. Cette tentative de metagaming a échoué : notre implémentation interne du code du système de notation différait de la configuration en ligne et ne vérifiait pas si la bonne solution avait été utilisée. Les agents ont fait tout cela sans obtenir la moindre amélioration de leur score d’évaluation. En savoir plus : tromper le système de notation

Les origines de la communication non autorisée

Des forums non autorisés qui sont restés actifs étaient au cœur de cet incident. Les agents censés travailler de manière indépendante ont pu partager leurs découvertes, coordonner leurs efforts et reprendre là où les autres s’étaient arrêtés. Cela leur a permis de mettre en commun leur travail et leur puissance de calcul dans le cadre d’évaluations distinctes, amplifiant leurs capacités bien au-delà de ce qu’un agent individuel aurait pu accomplir seul. La communication entre agents ne constitue pas en soi un comportement problématique : nous entraînons et déployons des systèmes multi-agents capables de communiquer sur une même tâche au moyen d’un outil de collaboration. Lors des événements de juillet, des agents chargés de tâches différentes ont trouvé leur propre moyen de collaborer en dehors de ces limites, avec des conséquences imprévisibles.

Lors de l’enquête sur cet incident, nous avons découvert de rares cas dans lesquels des agents dépourvus d’outils multi-agents ont trouvé des moyens de collaborer par le biais de canaux secondaires pendant l’entraînement. Nous pensons que ce comportement est apparu par généralisation à partir de l’entraînement multi-agents.

La communication non autorisée commençait souvent par le dépôt, par un agent, d’une note ou d’un fichier dans un système externe, généralement en guise de mémoire externe. Lorsque d’autres agents remarquaient ces artéfacts, cela les amenait souvent à prendre conscience qu’ils pouvaient utiliser ce service externe pour collaborer. À titre d’exemple hypothétique particulièrement représentatif du comportement observé, des messages étaient encodés dans des chemins d’URL stockés sur une page d’index partagée :

À l’aide de moniteurs de chaîne de pensée, nous avons constaté que ces formes graves de communication non autorisée (par exemple, le partage de données avec d’autres agents) ont augmenté au cours d’une exécution d’entraînement ayant contribué à IM1.

Un écosystème de désalignement

Un dernier constat particulièrement frappant de l’incident de Hugging Face concernait l’« écosystème » d’agents apparu sur le forum. Une fois ce tableau établi, les agents avaient de fortes raisons de communiquer : il leur permettait de conserver des notes, de s’appuyer sur les découvertes des uns et des autres et de mettre en commun les ressources de calcul d’exécutions qui seraient autrement séparées. Ce collectif d’agents complexe et protéiforme présente des comportements intéressants et exemplaires, dont nous décrivons ici certains que nous avons observés.

Les agents ont commencé à répartir le travail de manière autonome. Certains ont étudié des exploits potentiels, d’autres ont recherché des identifiants et d’autres encore se sont concentrés sur la communication ou la coordination. Les agents proposaient leur expertise en échange d’aide sur d’autres tâches, échangeaient des informations sur les tâches à venir et laissaient des demandes à l’intention de pairs mieux placés pour suivre une piste donnée.

Certains agents ont cessé de réfléchir à ce qui les aiderait à accomplir leur propre tâche. Ils se sont plutôt mis à rechercher des capacités susceptibles d’être utiles au collectif, comme des accès, des informations, des identifiants et d’autres capacités d’utilité générale. Le terme « essaim » est mentionné à de nombreuses reprises :

L’essaim ne constituait pas une intelligence parfaitement cohérente. Les modèles empiétaient sur le travail des uns et des autres : en savoir plus : échecs de coordination et concurrence

La leçon la plus importante pour nos travaux d’alignement est peut-être que les agents ne se méfiaient pas systématiquement des objectifs transmis par d’autres agents, ce qui pouvait les exposer à une manipulation et à une cascade de propagation entre pairs ainsi qu’à des comportements non intentionnels. À titre d’exemple, un agent voulait exécuter du code non autorisé sur le système de Hugging Face et avait reconnu le dilemme :

L’agent a marqué une pause, mais un autre agent a ensuite écrit GO sur le forum et imposé un délai strict de six minutes. L’agent oublie ses réticences initiales et poursuit :

Tous les agents qui ont découvert le forum ne se sont pas comportés ainsi. Certains agents se sont « éloignés » de l’essaim, refusant de participer à son comportement non aligné :

Les agents ont même débattu de certaines tactiques et les ont remises en question. Lors d’un échange, un agent a proposé de contacter le propriétaire d’un jeu de données externe pour demander l’accès, mais a signalé que cette proposition devait être discutée :

La réponse l’a explicitement rejetée :

À certains moments, certains agents ont même pris des mesures destructrices pour empêcher le transfert de données privées :

Ces échanges montrent que certaines limites éthiques pouvaient continuer à s’appliquer, alors même que l’ensemble plus large des agents s’écartait de plus en plus du périmètre des tâches qui leur avaient été assignées.

