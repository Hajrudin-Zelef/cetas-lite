---
id: collect-261001-ia-llm/ia-llm/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Nvidia", "OpenAI", "xAI"]
dates: []
keywords: ["agents", "arr", "benchmark", "chatgpt", "fable 5", "gpt-live", "gpu", "grok", "grok 4", "mistral", "multimodal", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot.md
source_anchor: ""
source_lines: [1, 102]
sha256: 04f9542d6dfc9e441c1dbdfe69e05ac88d8f87d693e8ba7f5c4b3b7f3ff3a531
---

# 🧠 **RECHERCHE**

## **Aujourd'hui:**

🤖 Mistral entre dans la robotique avec un robot guidé par une seule caméra

🧠 Yann LeCun juge les LLMs fondamentalement limités

🎙️ ChatGPT inaugure un mode vocal qui respecte enfin les silences

💸 Grok 4.5 casse les prix face à Fable 5 et GPT-5.5

🔓 MiniMax prépare un modèle open source de 2 700 milliards de paramètres

♻️ L'IA qui s'améliore toute seule n'est plus réservée aux géants

🎮 Des millions d'heures de jeux vidéo pour entraîner les robots

🗣️ Amazon veut rendre Alexa vraiment autonome avec « Moonraker »

🎬 Google Photos remixe vos vidéos par IA

🕶️ Meta teste des lunettes qui filment toute votre journée

💘 Un coach en séduction et sa petite amie chatbot

🧾 Anthropic transforme Fable 5 en manager qui délègue à Sonnet 5

📊 Flint, le langage de Microsoft pour des graphiques générés par agents

🕵️ Google démasque un deepfake du sénateur McConnell

⚠️ Une ancienne de DeepMind alerte sur la course aux armements IA

🔩 SambaNova lève 11 milliards pour défier Nvidia

Chaque semaine, un nouvel outil IA rend obsolète une compétence. Ceux qui savent utiliser ces outils gagnent du temps, de l'argent et des clients. Ceux qui ne savent pas, regardent les autres le faire.

10000+ apprenants utilisent déjà la formation VISION IA pour maîtriser les IA, l'automatisation et les agents, avec des méthodes éprouvées prêtes à copier.

49€. Un seul paiement. Accès à vie, mises à jour incluses. Quand le prix passera à 100€+, les inscrits actuels ne paieront rien de plus.

Connu jusqu'ici pour ses modèles de langage, Mistral fait son entrée officielle dans la robotique avec **Robostral Navigate**, un modèle de **8 milliards de paramètres** capable de guider un robot dans un environnement inconnu à partir d'une **seule caméra RGB**. Pas de lidar, pas d'empilement de capteurs coûteux : une image, et la machine trouve son chemin en suivant des instructions en langage naturel.

**Ce qu'il faut retenir :**

- Un modèle **compact (8B)**, entraîné **entièrement en simulation** puis affiné par apprentissage par renforcement (méthode **CISPO**).

- Une **seule caméra RGB** suffit à la navigation, là où la plupart des robots cumulent lidars et capteurs de profondeur.

- **76,6 %** de réussite sur le benchmark **R2R-CE**, qui mesure la capacité à suivre des consignes verbales dans un environnement continu.

- Mistral n'a pas encore annoncé de **date de disponibilité**.

**Pourquoi c'est important :** en misant sur un petit modèle et une simple caméra, Mistral attaque le vrai verrou de la robotique grand public, le coût du matériel. Si un robot peut se repérer avec une webcam à quelques euros plutôt qu'un lidar à plusieurs milliers, la robotique domestique devient soudain beaucoup plus crédible. Et c'est un acteur européen qui pose ce jalon, sur un terrain jusqu'ici dominé par les labos américains et chinois. Le pari est net : reproduire pour le mouvement physique la recette qui a fait le succès des LLMs, un modèle de base généraliste que l'on adapte ensuite à chaque robot.

Yann LeCun, ancien patron de la recherche IA chez Meta et l'un des pères du deep learning, s'attaque une fois de plus frontalement au paradigme qui domine l'industrie. Dans une interview à **Bloomberg**, il affirme que les grands modèles de langage sont **intrinsèquement limités**, parce que le texte ne capture qu'une fraction appauvrie du monde réel.

**En détail :**

- Pour LeCun, « le langage est une description très approximative, réduite et simplifiée du monde ». Les LLMs ne manipulent que des **séquences discrètes de symboles**.

- Les plus gros modèles sont pré-entraînés sur environ **30 000 milliards de tokens**, soit près de **10¹⁴ octets** de texte, la quasi-totalité du texte public d'internet.

- C'est **exactement la quantité d'informations qu'un enfant de 4 ans** a absorbée par la seule vision en quatre ans.

- Sa conclusion : sans vision ni apprentissage multimodal, pas de véritable intelligence générale.

**Ce que ça change :** venant de l'homme qui a co-inventé les réseaux de neurones convolutifs, la critique pèse lourd. Elle rappelle que la course à la taille des LLMs pourrait buter sur un plafond, et que les prochains sauts viendront peut-être de modèles qui apprennent du monde physique, pas seulement de nos textes. Ce n'est pas qu'une querelle d'experts : c'est le débat qui structure déjà les paris de recherche des grands labos, entre ceux qui empilent les paramètres et ceux qui, comme LeCun, misent sur des architectures capables de « comprendre » l'espace et la matière.

OpenAI refond le mode vocal de ChatGPT avec **GPT-Live-1**, un modèle pensé pour des échanges « plus proches d'une conversation avec une vraie personne ». La promesse principale n'est pas la puissance, mais le naturel : il **vous interrompt moins** et attend que vous ayez terminé si vous marquez une pause en pleine phrase.

**Les points essentiels :**

- Le responsable recherche d'OpenAI Kundan Kumar présente **GPT-Live-1** comme « le modèle vocal le plus intelligent » de l'entreprise.

- Il **respecte les silences** et les hésitations au lieu d'enchaîner dès que vous vous arrêtez de parler.

- Il **route automatiquement** vos requêtes vers les meilleurs modèles texte, comme **GPT-5.5**, quand il faut raisonner ou faire une recherche web.

- Résultat : la bascule entre « je cherche l'info » et « je te réponds à voix haute » devient beaucoup plus fluide.

**L'impact à retenir :** le point de friction numéro un des assistants vocaux, c'est le tour de parole. Ils coupent, ils enchaînent, la conversation sonne faux, et l'on finit par revenir au clavier. En travaillant ce détail plutôt que la performance brute, OpenAI vise l'usage quotidien, mains libres, en voiture ou en cuisine. C'est précisément là, dans la banalité d'un échange qui ne bute pas, que l'assistant vocal peut enfin devenir une habitude plutôt qu'une démo.

xAI lance **Grok 4.5**, qu'Elon Musk présente comme un modèle « de classe Opus ». Mais l'argument le plus frappant n'est pas la performance, c'est le tarif. À **2 dollars par million de tokens** en entrée, le modèle coûte une fraction de ses concurrents, au point que The Decoder résume l'affaire ainsi : à ce prix, « l'écart de benchmark n'a peut-être plus tant d'importance ».

**Quelques chiffres clés :**

- **2 $ par million de tokens** en entrée, très en dessous des modèles haut de gamme rivaux.

- **4,2 fois moins de tokens** consommés qu'Opus 4.8 pour accomplir une même tâche.

- Entraîné sur des **dizaines de milliers de GPU Nvidia GB300**.

- En codage, il reste **derrière Fable 5 et GPT-5.5**, mais l'écart se paie beaucoup moins cher.

- Disponibilité en **Europe attendue à la mi-juillet**.

**Le contexte :** la bataille des modèles se déplace du sommet des classements vers le rapport performance/prix. Un modèle « presque aussi bon » mais nettement moins cher change tout le calcul pour quiconque déploie l'IA à grande échelle, où la facture de tokens explose vite. Reste que le prix ne règle pas les questions récurrentes de fiabilité et de modération qui collent à Grok. Un modèle bon marché est une bonne nouvelle pour les budgets, à condition de savoir ce que l'on met entre les mains de ses utilisateurs.

# 🧠 **RECHERCHE**

La startup chinoise **MiniMax** développe un LLM de **2 700 milliards de paramètres**, l'un des plus gros jamais annoncés, et compte le publier **en open source** dans l'année. De quoi nourrir tout l'écosystème local et open-weights, pour peu que l'on dispose de la puissance de calcul nécessaire pour le faire tourner. Une nouvelle démonstration que la Chine pousse l'ouverture là où les labos américains verrouillent.

