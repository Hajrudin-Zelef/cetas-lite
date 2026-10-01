---
id: collect-261001-ia-llm/ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore-1
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Meta", "Microsoft", "OpenAI"]
dates: []
keywords: ["llama", "benchmarks", "deepseek", "distillation", "fine-tuning", "gemini", "gpu", "moe", "multimodal", "muse", "muse spark", "open-weight"]
source: docs/RAG/collect-261001-ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [1, 57]
sha256: 54fbc87abe9393709bffae373bf4a4d2982013ce1ed7d7fe1c4b25a9c83d517d
---

# llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore

Cours

Meta vient d’annoncer la suite de modèles Llama 4, qui comprend deux modèles déjà disponibles — Llama 4 Scout et Llama 4 Maverick — et un troisième encore en entraînement : Llama 4 Behemoth.

Les variantes Scout et Maverick sont disponibles dès maintenant, publiées ouvertement sous la licence open-weight habituelle de Meta — avec une réserve notable : si vos services dépassent 700 millions d’utilisateurs actifs mensuels, vous devez obtenir une licence distincte auprès de Meta, qui peut ou non vous être accordée à sa discrétion.

Llama Scout prend en charge une fenêtre de contexte de 10 millions de tokens, la plus grande de tous les modèles publiés publiquement. Llama Maverick est un modèle généraliste qui vise GPT-4o, Gemini 2.0 Flash et DeepSeek-V3. Llama Behemoth, encore en entraînement, sert de modèle enseignant à haute capacité.

Dans cet article d’introduction, je vous propose un tour d’horizon de la suite Llama 4. Notre équipe a déjà testé le modèle, et je vous recommande ces tutoriels si vous souhaitez aller plus loin :

**Mise à jour :** après environ un an de silence, Meta a publié son tout dernier LLM. Pour une vue d’ensemble, consultez notre guide sur Muse Spark.

Nous tenons nos lecteurs informés des dernières avancées en IA via *The Median*, notre newsletter gratuite du vendredi qui décode les actus clés de la semaine. Abonnez-vous et restez au fait en quelques minutes par semaine :


## Qu’est-ce que Llama 4 ?

Llama 4 est la nouvelle famille de grands modèles de langage de Meta. La publication inclut deux modèles déjà disponibles : Llama 4 Scout et Llama 4 Maverick ; et un troisième, Llama 4 Behemoth, encore en entraînement.

Source : Meta AI

Llama 4 introduit des améliorations substantielles. En particulier, il intègre une architecture mixture-of-experts (MoE), qui vise à améliorer l’efficacité et les performances en activant uniquement les composantes nécessaires pour une tâche donnée (nous y revenons juste après). Ce design marque une évolution vers des modèles d’IA plus extensibles et spécialisés.

Llama 4 prolonge la stratégie de Meta consistant à publier des modèles open-weight — mais avec une réserve. Si votre entreprise exploite des services totalisant plus de 700 millions d’utilisateurs actifs mensuels, vous devrez obtenir une licence distincte de Meta, qui peut ne pas être accordée. Malgré cette contrainte, la sortie reste un événement majeur dans l’écosystème open-weight, même si celui-ci a évolué très rapidement ces derniers mois.

Si Llama 2 et 3 ont un temps défini la catégorie, Llama 4 arrive désormais dans un champ bien plus concurrentiel. DeepSeek s’est imposé avec de fortes capacités de raisonnement. La série Qwen d’Alibaba s’illustre sur des benchmarks multilingues et de code. Les modèles Gemma de Google avancent sur le même terrain avec des architectures plus petites et efficaces. Et tout récemment, OpenAI a annoncé son intention de publier un modèle open-weight, une inflexion impensable il y a un an.

Voyons maintenant les détails de chaque modèle.

## Llama Scout

Llama 4 Scout est le modèle le plus léger de la nouvelle suite, mais c’est sans doute le plus intrigant. Il fonctionne sur un seul GPU H100 et prend en charge une fenêtre de contexte de 10 millions de tokens. Cela en fait le modèle open-weight le plus gourmand en contexte à ce jour et potentiellement le plus utile pour des tâches comme la synthèse multi-documents, le raisonnement sur de longs codes et l’analyse d’activité.

Scout compte 17 milliards de paramètres actifs, organisés via 16 experts, pour un total d’environ 109 milliards de paramètres. Il a été pré-entraîné et post-entraîné avec une fenêtre de contexte de 256 K, mais Meta affirme qu’il généralise bien au-delà (à vérifier). En pratique, cela ouvre la voie à des workflows englobant des bases de code entières, des historiques de session ou des documents juridiques — le tout traité en un seul passage avant.

Sur le plan architectural, Scout s’appuie sur le cadre MoE de Meta, où seul un sous-ensemble de paramètres s’active par token — contrairement aux modèles denses comme GPT-4o, où tous les paramètres sont activés. Résultat : efficacité de calcul et grande scalabilité.

Au-delà de l’architecture, Meta met en avant les capacités multimodales de Scout. Il a été pré-entraîné sur des données texte, image et vidéo avec une fusion précoce, ce qui lui permet de gérer nativement des combinaisons d’invites textuelles et visuelles. Sur des tâches riches en images comme l’ancrage visuel et la VQA (visual question answering), Scout surpasse tous les précédents modèles Llama — et tient tête à des systèmes bien plus grands.

En bref, Scout est conçu pour la polyvalence et l’échelle. Il fonctionne efficacement, accepte plus d’entrée que tout autre modèle open-source publié auparavant et se montre performant sur les tâches texte et image. Nous testerons bientôt cette limite de contexte à 10 M — et nous vous ferons un retour.

## Llama Maverick

Llama 4 Maverick est le généraliste de la gamme : un modèle multimodal de grande ampleur, conçu pour la performance en conversation, raisonnement, compréhension d’images et code. Tandis que Scout repousse les limites de longueur de contexte, Maverick privilégie une qualité d’output équilibrée et élevée sur l’ensemble des tâches. C’est la réponse de Meta à GPT-4o, DeepSeek-V3 et Gemini 2.0 Flash.

Maverick dispose du même nombre de 17 milliards de paramètres actifs que Scout, mais avec une configuration MoE plus vaste : 128 experts et un total d’environ 400 milliards de paramètres. Comme Scout, il utilise une architecture MoE qui n’active qu’une partie du modèle par token — réduisant les coûts d’inférence tout en augmentant la capacité. Le modèle tourne sur un hôte H100 DGX unique, mais peut aussi être déployé en inférence distribuée pour des applications à grande échelle.

Meta a adopté une approche différente pour le post-entraînement, mêlant fine-tuning supervisé léger, apprentissage par renforcement en ligne et optimisation directe des préférences. L’objectif : affûter les performances sur des prompts difficiles sans surcontraindre le modèle. Pour ce faire, Meta a écarté plus de 50 % des exemples d’entraînement jugés « faciles » par d’anciennes versions de Llama et a bâti un programme de formation axé sur des tâches plus ardues de raisonnement, de codage et multimodales.

Maverick a également été co-distillé depuis Llama 4 Behemoth, le modèle interne bien plus grand de Meta, ce qui a permis d’améliorer les performances sans coûts d’entraînement additionnels. D’après Meta, cette chaîne de distillation a entraîné une nette progression du raisonnement et de la qualité en conversation.

## Llama Behemoth

Llama 4 Behemoth est à ce jour le modèle le plus puissant et le plus volumineux de Meta — mais il n’est pas encore disponible. Toujours en entraînement, Behemoth n’est pas un modèle de raisonnement au même sens que DeepSeek-R1 ou o3 d’OpenAI, qui sont conçus et optimisés pour des raisonnements multi-étapes de type chain-of-thought.

D’après les informations disponibles, il ne semble pas non plus conçu comme un produit d’usage direct. Il agit plutôt comme un modèle enseignant, utilisé pour distiller et façonner Scout et Maverick. Une fois publié, il pourrait aussi permettre à d’autres de distiller leurs propres modèles.

