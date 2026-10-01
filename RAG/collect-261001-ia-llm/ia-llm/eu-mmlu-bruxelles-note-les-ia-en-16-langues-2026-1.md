---
id: collect-261001-ia-llm/ia-llm/eu-mmlu-bruxelles-note-les-ia-en-16-langues-2026-1
title: "eu-mmlu-bruxelles-note-les-ia-en-16-langues-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "EU", "Google", "Hugging Face", "Mistral", "OpenAI", "Z.ai"]
dates: ["2026-08"]
keywords: ["benchmark", "benchmarks", "claude", "diffusion", "gemini", "glm", "gpt-5.6", "mistral", "opus 5", "reasoning", "valuation"]
source: docs/RAG/collect-261001-ia-llm/eu-mmlu-bruxelles-note-les-ia-en-16-langues-2026.md
source_anchor: ""
source_lines: [1, 34]
sha256: 4a12f365d4da016d919be3a13c8c6624a098ccb15472809ef56e973c219b3452
---

# eu-mmlu-bruxelles-note-les-ia-en-16-langues-2026

Le 22 juillet 2026, la Commission européenne a mis en ligne un outil que peu de développeurs attendaient sous cette forme précise : un benchmark multilingue baptisé **EU MMLU**, conçu pour évaluer les grands modèles de langage dans 16 langues officielles de l’Union, dont le français, l’allemand, l’italien et le néerlandais. Un mois plus tard, ce classement commence déjà à redistribuer les cartes. Sur la première grille de résultats publiée par BenchLM le 22 août 2026, c’est **Ministral 3 14B (Reasoning)**, un modèle français de Mistral AI, qui occupe la première place chez les modèles européens avec un score composite de 49,7. Ce n’est pas un hasard de calendrier : l’EU MMLU arrive exactement au moment où l’AI Act impose ses premières obligations de transparence documentaire aux fournisseurs de modèles à usage général, et où neuf nouveaux modèles ont été lancés par six laboratoires différents rien qu’au mois d’août 2026. Pour les développeurs et les entreprises françaises qui choisissent une API de LLM, la question n’est plus seulement « quel modèle est le plus puissant », mais « quel modèle comprend vraiment le français, l’allemand ou le polonais, et documente ses données d’entraînement comme la loi l’exige désormais ».

## Qu’est-ce que l’EU MMLU, le nouveau benchmark linguistique de Bruxelles ?

L’EU MMLU est une adaptation du célèbre benchmark américain MMLU (Massive Multitask Language Understanding), mais reconstruite et traduite pour coller aux réalités linguistiques européennes plutôt que de se contenter d’une traduction automatique de l’anglais. Selon la page officielle publiée par la direction générale de la traduction, le jeu de données couvre actuellement 16 langues officielles, structurées en 57 sujets et plus de 1 000 questions dès sa première publication en juillet 2026, un volume confirmé par le média spécialisé Slator : le croate, le tchèque, le néerlandais, le français, l’allemand, le grec, le hongrois, l’irlandais, l’italien, le lituanien, le polonais, le portugais, le roumain, le slovaque, le slovène et l’anglais comme référence. L’ensemble a été publié sous licence ouverte CC BY 4.0 sur Hugging Face par l’équipe EC-DGT-AI, qui y a également déposé un sous-ensemble élargi couvrant 24 langues officielles de l’Union, un choix de diffusion qui facilite l’audit indépendant de la performance linguistique des modèles par des chercheurs extérieurs à la Commission.

L’objectif affiché n’est pas de créer un simple gadget statistique de plus. C’est de répondre à un problème concret que tout développeur ayant testé un LLM en français a déjà remarqué : un modèle qui obtient 88 % sur MMLU en anglais peut chuter de dix à quinze points quand on lui pose exactement les mêmes questions en français ou en roumain. Pour éviter les approximations d’une traduction automatique, la Commission a mobilisé 250 étudiants issus de 21 universités européennes chargés de retravailler et de valider chaque question langue par langue avant la publication de juillet 2026. Jusqu’ici, cette perte de qualité linguistique restait largement invisible dans les benchmarks marketing des éditeurs, qui communiquent presque toujours sur leurs scores en anglais. L’EU MMLU force une transparence que le marché ne s’imposait pas de lui-même, et il tombe au moment précis où l’AI Act rend la documentation des performances par langue beaucoup plus qu’une option de bonne volonté.

## Ministral 3 14B et Mistral Small 4 : la France en tête du classement européen

Le classement établi par BenchLM au 22 août 2026 place Ministral 3 14B en version « Reasoning » à la première place des modèles européens, avec un score composite de 49,7. C’est un résultat notable pour un modèle de seulement 14 milliards de paramètres, largement en dessous de la taille des modèles frontières américains ou chinois, ce qui confirme la stratégie de Mistral AI consistant à optimiser l’efficacité par paramètre plutôt que de courir après le nombre brut de paramètres.

Sur le segment entreprise, un second classement de BenchLM daté du 21 août 2026 distingue **Mistral Small 4** comme le modèle généraliste européen le plus complet pour les usages professionnels, avec une couverture de benchmarks jugée particulièrement riche pour un déploiement en production. Ces deux résultats, obtenus à un jour d’intervalle par la même méthodologie, dessinent une image cohérente : sur le terrain spécifique de la compréhension multilingue européenne, les modèles français ne sont plus de simples alternatives de repli face à OpenAI, Google ou Anthropic, ils occupent la tête du classement sur leur propre terrain de jeu réglementaire.

Cela ne veut pas dire que Mistral devance les géants américains sur l’ensemble des tâches. Les benchmarks généralistes en anglais, sur le raisonnement pur ou le codage, continuent de favoriser des modèles bien plus massifs comme GPT-5.6 ou Claude Opus 5. Mais l’EU MMLU mesure autre chose : la capacité d’un modèle à rester fiable quand la question change de langue sans changer de sens, un critère qui compte directement pour un service client français, une administration publique ou une entreprise qui déploie un assistant vocal en Europe.

## Le tableau des scores : qui comprend vraiment le français ?

Voici un aperçu comparatif des positionnements rapportés par les classements disponibles à fin août 2026. Les scores composites reflètent la méthodologie propre à chaque source citée et ne doivent pas être comparés terme à terme avec d’autres benchmarks non EU MMLU. La robustesse de la couverture par sujet commence d’ailleurs à être vérifiée langue par langue : une adaptation bulgare, MMLU-BG, publiée dans les actes de la conférence LREC 2026, couvre déjà 56 des 57 sujets du référentiel européen dès septembre 2026, un niveau de granularité qui donne du poids aux scores composites du tableau ci-dessous.

| Modèle | Éditeur | Origine | Score composite européen | Point fort constaté | 
|---|---|---|---|---|
| Ministral 3 14B (Reasoning) | Mistral AI | France | 49,7 | 1ᵉʳ modèle européen, efficacité par paramètre | 
| Mistral Small 4 | Mistral AI | France | Non chiffré publiquement | Couverture entreprise la plus large | 
| Gemini 3.7 Flash | Google DeepMind | États-Unis | Non classé EU MMLU | Codage et automatisation métier | 
| GLM-5.3 | Zhipu AI | Chine | Non classé EU MMLU | Poids ouverts, coût réduit | 
| Qwen3.8-Max | Alibaba | Chine | Non classé EU MMLU | Multilingue asiatique, contexte long | 

Le fait que la majorité des modèles américains et chinois ne soient pas encore classés officiellement sur l’EU MMLU est en soi une information : la plupart des laboratoires non européens communiquent leurs propres suites de benchmarks (MMLU-Pro, GPQA, SWE-bench) sans se soumettre systématiquement à l’évaluation multilingue européenne, ce qui laisse pour l’instant le haut du tableau largement occupé par les modèles conçus dès l’origine pour le marché européen.

## L’AI Act rattrape le calendrier des benchmarks le 2 août 2026

L’arrivée de l’EU MMLU ne se produit pas dans le vide réglementaire. Selon les lignes directrices publiées par la Commission européenne, « Article 50 of the AI Act applies from 2 August 2026 » (l’article 50 de l’AI Act s’applique à compter du 2 août 2026). Ce même point est répété dans la documentation officielle sur les obligations de transparence, qui précise que « these transparency obligations apply from 2 August 2026 » (ces obligations de transparence s’appliquent à compter du 2 août 2026).

