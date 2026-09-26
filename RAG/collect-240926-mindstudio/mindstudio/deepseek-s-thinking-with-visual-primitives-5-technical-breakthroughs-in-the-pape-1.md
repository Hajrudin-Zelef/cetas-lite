---
id: collect-240926-mindstudio/mindstudio/deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape-1
title: "deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI"]
dates: ["2024-10", "2024-12", "2025-01"]
keywords: ["deepseek", "agent", "agents", "attention", "benchmark", "claude", "consumer", "distillation", "fine-tuning", "gemini", "gpu", "multimodal"]
source: docs/RAG/clean_en/mindstudio/deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape.md
source_anchor: ""
source_lines: [1, 77]
sha256: 0ac3abe4db16f1ecd91b60e13e4803dedb814a2db723b0b3be7fda11a545f594
---

# deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape

<!-- source: https://www.mindstudio.ai/blog/deepseek-thinking-visual-primitives-5-technical-breakthroughs-paper -->

## DeepSeek a publié un article de vision, puis l'a retiré. Voici 5 détails techniques qui méritent d'être compris.

L'article s'intitulait « Thinking with Visual Primitives ». Il est apparu, a été cité, puis est devenu difficile à trouver. Cela seul le rend digne d'attention.

Voici 5 idées techniques spécifiques enfouies dans cet article — en commençant par celle qui lui donne son nom.

## 1. Les primitives visuelles sont des tokens en ligne dans la chaîne de raisonnement, pas un appel d'outil

C'est l'idée centrale, et elle est plus propre qu'elle n'en a l'air.

Quand on demande à un modèle multimodal actuel de compter des personnes dans une photo bondée, le modèle essaie de le faire en langage. Il décrit ce qu'il voit, tient un décompte courant, et perd fréquemment le fil. Le problème n'est pas que le modèle ne peut pas voir — c'est que le langage est un mauvais médium pour la référence spatiale. Les mots dérivent. Les pronoms deviennent ambigus. « La troisième personne en partant de la gauche » ne signifie plus rien de précis après quatre phrases de raisonnement supplémentaires.

La réponse de DeepSeek : donner au modèle un moyen de pointer.

Le format est `<ref>label</ref><box>x1,y1,x2,y2</box>`. Ce sont des tokens spéciaux dans le vocabulaire du modèle. Pas un appel de fonction. Pas un module de grounding séparé invoqué via API. Pas un outil que le modèle décide d'utiliser. Ils apparaissent en ligne dans la chaîne de raisonnement, de la même manière qu'un humain pourrait entourer quelque chose sur un tableau blanc au milieu d'une phrase.

## Remy est nouveau. La plateforme ne l'est pas.

Remy est l'expression la plus récente d'années de travail sur la plateforme. Pas un LLM hâtivement emballé.

Ainsi, quand le modèle travaille sur « compter les hommes dans cette photo d'équipe », il ne se contente pas de narrer — il émet une boîte englobante pour chaque personne qu'il identifie, directement dans la trace de raisonnement. Chaque entité reçoit une coordonnée. Le modèle peut y revenir. La référence ne dérive pas.

L'article appelle cela le « fossé de référence » — distinct du « fossé de perception » que la plupart des travaux multimodaux de l'ère 2024 tentaient de résoudre. Le recadrage haute résolution, le patching dynamique, le raisonnement avec des images — tout cela vise à mieux voir. Les primitives visuelles visent à mieux *pointer*. La distinction compte.

## 2. Le pipeline d'entraînement comporte cinq étapes et trois têtes de récompense

L'architecture est un modèle mixture-of-experts de 284 milliards de paramètres avec 13 milliards de paramètres actifs à l'inférence — le backbone DeepSeek V4 Flash. Un raisonnement de niveau frontière pour une fraction du coût d'inférence. Cette partie est cohérente avec le schéma de DeepSeek.

Le pipeline d'entraînement est là où la conception devient intéressante.

Étape un : pré-entraînement multimodal sur des milliers de milliards de tokens. Standard.

Étape deux : fine-tuning supervisé, mais scindé. Ils entraînent *deux modèles séparés* — un pour le grounding (boîtes englobantes) et un pour le pointage (points de coordonnées). Pas un seul modèle essayant de faire les deux. Deux spécialistes.

Étape trois : apprentissage par renforcement avec GRPO sur chaque spécialiste, utilisant trois têtes de récompense : format, qualité et précision. Les récompenses de format garantissent que le modèle émet réellement des tokens `<ref><box>` valides. Les récompenses de qualité poussent vers un raisonnement intermédiaire utile. Les récompenses de précision mesurent si la réponse finale est correcte. Trois signaux séparés, pas une seule perte mélangée.

Étape quatre : RFT unifié qui fusionne les deux spécialistes.

Étape cinq : distillation on-policy dans un seul modèle étudiant.

L'élégance ici réside dans la séparation des préoccupations. Entraîner des spécialistes, puis consolider. C'est la même logique que celle derrière le mixture-of-experts au niveau de l'architecture, appliquée au processus d'entraînement lui-même. On ne demande pas à un seul modèle d'apprendre deux comportements de grounding différents simultanément — on laisse chacun devenir bon dans sa tâche, puis on fusionne.

C'est significativement différent de la façon dont la plupart des laboratoires abordent le fine-tuning multimodal, qui tend à être une seule passe de SFT sur un jeu de données combiné suivie de RLHF. Le pipeline en cinq étapes avec des têtes de récompense RL séparées pour chaque modalité de grounding est une conception plus délibérée.

## 3. Les chiffres de benchmark sont spécifiques — et honnêtement délimités

Le résultat phare est la navigation dans un labyrinthe : DeepSeek obtient 67 %, contre 49 % pour Gemini Flash 3, 50 % pour GPT-5.4 et 49 % pour Claude Sonnet 4.6. C'est environ un écart de 17 points sur GPT-5.4 sur une tâche qui nécessite de suivre un chemin à travers une structure spatiale.

Le traçage de chemin raconte une histoire similaire. Le comptage et le raisonnement spatial sont plus mitigés — Gemini Flash 3 est encore en tête sur le QA de comptage brut.

La raison pour laquelle les primitives visuelles aident spécifiquement sur la navigation dans un labyrinthe est que le langage est particulièrement mauvais pour les descriptions de trajectoire. « Va à droite, puis en haut, puis encore à droite à la jonction » s'effondre vite dans un labyrinthe complexe. Un modèle qui peut émettre des références de coordonnées pendant qu'il raisonne — marquant des points de passage dans sa chaîne de raisonnement — a un avantage structurel sur exactement ces tâches.

- ✕un agent de codage
- ✕no-code
- ✕vibe coding
- ✕un Cursor plus rapide

Celui qui dit aux agents de codage quoi construire.

L'article est honnête sur la portée. Il y a une note de bas de page : « les scores rapportés ne couvrent qu'un sous-ensemble des dimensions d'évaluation directement pertinentes pour le focus de recherche de cet article et ne sont donc pas indicatifs des capacités globales du modèle. » Ils ne prétendent pas battre GPT-5.4 sur tous les tableaux. Ils prétendent le battre sur les tâches de raisonnement spatial visuellement ancrées. C'est une affirmation plus étroite et plus défendable.

Ce type de délimitation est rare. La plupart des annonces de benchmark enterrent les mises en garde. DeepSeek l'a mise dans une note de bas de page, mais ils l'ont mise.

## 4. La lignée vision remonte à deux ans et raconte une histoire cohérente

L'article sur les primitives visuelles ne vient pas de nulle part. DeepSeek a livré environ sept modèles liés à la vision depuis mars 2024, et ils répondent tous à la même question : quelle est la représentation la moins coûteuse qui fonctionne encore ?

DeepSeek VL (mars 2024) : modèles de 1,3 Md et 7 Md, encodeur hybride SigLIP et SAM. Pose des fondations, rien de tapageur.

Janus (October 2024): decoupled visual encoders for understanding versus generation. Most unified multimodal models at the time had a single encoder bottleneck — one encoder trying to serve both comprehension and generation. Janus ran two encoders sharing a transformer. The insight was that the two tasks have different representational needs, and forcing a compromise hurts both.

VL2 (December 2024): ported mixture-of-experts and multi-head latent attention from V2/V3 into vision. A 1B activated-parameter version scored 80.9 on OCR Bench and 88.9 on DocVQA. Small activations, large numbers.

Janus Pro 7B (January 2025): went viral during the R1 moment. 80% on GenEval, runnable on a single consumer GPU.

