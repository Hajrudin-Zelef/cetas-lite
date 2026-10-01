---
id: collect-261001-ia-llm/ia-llm/qwen3-8-flash-next-alibaba-previsualise-qwen4-2026-1
title: "qwen3-8-flash-next-alibaba-previsualise-qwen4-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "OpenAI", "Z.ai"]
dates: []
keywords: ["qwen", "apache", "astra", "attribution", "benchmark", "benchmarks", "claude", "deepseek", "embeddings", "fable 5", "fine-tuning", "gemini"]
source: docs/RAG/collect-261001-ia-llm/qwen3-8-flash-next-alibaba-previsualise-qwen4-2026.md
source_anchor: ""
source_lines: [1, 34]
sha256: 438296dd7d7bde43d1188eb08ef37ce2f0ac65e2636c98ea8acd737f4d6d5de4
---

# qwen3-8-flash-next-alibaba-previsualise-qwen4-2026

Le 26 août 2026, l’équipe Qwen d’Alibaba a mis en ligne les poids de **Qwen3.8-Flash-Next** sur Hugging Face et ModelScope, sans conférence de presse ni grand lancement marketing. Trois semaines plus tard, ce modèle s’impose comme l’un des sujets IA les plus commentés de la rentrée technique 2026 : il ne s’agit pas d’un simple modèle “Flash” économique de plus, mais d’un aperçu assumé de l’architecture qui équipera **Qwen4**, la prochaine génération de la famille de modèles chinoise. Avec 180 milliards de paramètres au total dont seulement 6 milliards actifs par requête, une fenêtre de contexte native de 262 144 tokens et une licence commerciale conditionnelle, Qwen3.8-Flash-Next relance en France et en Europe le débat sur la place des modèles ouverts face à GPT-6 Astra, Claude Fable 5.1 et Gemini 3.8 Flash.

## Qu’est-ce que Qwen3.8-Flash-Next exactement ?

Qwen3.8-Flash-Next est un modèle de langage multimodal à mélange d’experts (Mixture-of-Experts, ou MoE), doté d’un encodeur visuel qui lui permet de traiter du texte, des images et de la vidéo en entrée, pour une sortie textuelle. Le dépôt officiel sur GitHub indique une date de publication au 26 août 2026, confirmée par plusieurs suivis de sorties de modèles indépendants. Le modèle est décrit par Alibaba comme le jumeau open-weight de l’API hébergée Qwen3.8-Flash, disponible sur QwenCloud, mais avec une différence de taille : n’importe quel développeur peut télécharger les poids depuis Hugging Face ou ModelScope et les faire tourner sur sa propre infrastructure, y compris en Europe.

Ce qui distingue vraiment Qwen3.8-Flash-Next des publications précédentes de la marque, c’est sa fonction annoncée : il s’agit explicitement d’un **galop d’essai pour Qwen4**, l’architecture de nouvelle génération qu’Alibaba prépare pour la suite de sa gamme. Plusieurs analyses spécialisées, dont celle de Datacamp et du suivi de sorties AI Release Tracker, présentent le modèle comme la première démonstration publique du design MoE multimodal qui structurera Qwen4 dans ses différentes tailles.

## Une architecture MoE à 180 milliards de paramètres, 6 milliards actifs

Sur le plan technique, Qwen3.8-Flash-Next combine trois blocs : un modèle principal de 125 milliards de paramètres, une table d’embeddings n-gramme de 51 milliards de paramètres, et un module de prédiction multi-tokens (MTP) de 4 milliards de paramètres, pour un total d’environ 180 milliards de paramètres sur disque. Le point clé, et la raison pour laquelle ce modèle intéresse autant les ingénieurs, c’est que seulement **6 milliards de paramètres sont activés par token** traité, grâce à un routage vers 512 experts au total, dont 10 experts spécialisés plus 1 expert partagé sollicités à chaque étape. Le réseau utilise une dimension cachée de 2 560 et compte 48 couches de transformeur, selon les spécifications relevées sur le modèle publié sur Hugging Face.

Cette conception permet à Alibaba de revendiquer un entraînement environ neuf fois moins coûteux que celui de Qwen3.7-Plus, tout en visant des performances de code et de raisonnement comparables à des modèles denses beaucoup plus lourds. C’est cette efficacité, plus que la taille brute, qui explique pourquoi plusieurs commentateurs qualifient Qwen3.8-Flash-Next de “modèle à 6 milliards de paramètres actifs qui rivalise avec des modèles à près de 400 milliards”.

## Contexte natif de 262 144 tokens, extensible à 1 million

Le modèle open-weight est entraîné nativement sur une fenêtre de contexte de 262 144 tokens, soit environ 262 000 mots-équivalents. Ce chiffre correspond à la configuration de base des poids téléchargeables. L’API hébergée sur QwenCloud, elle, étend cette capacité jusqu’à 1 million de tokens via une technique de mise à l’échelle de type YaRN, avec une limite pratique d’environ 991 000 tokens en entrée et 131 072 tokens en sortie. Il faut noter une nuance importante souvent gommée dans les comparatifs rapides : l’extension à 1 million de tokens est une capacité de service, pas une caractéristique native des poids ouverts, qui restent calibrés à 262 144 tokens tant qu’ils ne sont pas explicitement étendus par l’utilisateur.

## La licence Qwen Community License 1.0 : ni Apache 2.0, ni gratuit pour tous

Contrairement à une partie de la gamme Qwen publiée sous licence Apache 2.0, Qwen3.8-Flash-Next est distribué sous la **Qwen Community License 1.0**, une licence plus restrictive. Le texte de la licence autorise l’usage commercial et le fine-tuning dans la plupart des cas, mais impose deux garde-fous concrets : une obligation d’attribution pour toute entreprise dépassant 100 millions d’utilisateurs mensuels actifs ou 20 millions de dollars de revenus mensuels, et une licence séparée obligatoire pour quiconque propose un service de type “modèle en tant que service” ou un assistant de codage IA commercial basé sur Qwen3.8-Flash-Next. Pour une PME française qui veut simplement l’utiliser en interne, ces seuils ne changent rien au quotidien. Pour un éditeur de logiciels qui viserait à revendre l’accès au modèle, la situation est nettement plus contraignante qu’avec du Apache 2.0 pur.

## Un prix agressif : 0,16 $ en entrée, 0,47 $ en sortie par million de tokens

Sur QwenCloud, l’API hébergée sous le nom “qwen3.8-flash” est facturée 0,16 dollar par million de tokens en entrée et 0,47 dollar par million de tokens en sortie, avec un tarif réduit à environ 0,016 dollar par million de tokens pour les lectures mises en cache. Le site de comparaison de modèles Artificial Analysis qualifie ce positionnement de “prix modéré”, en le situant sous la médiane du marché qu’il évalue à 0,30 dollar en entrée et 1,15 dollar en sortie pour l’ensemble des modèles suivis. Pour les entreprises qui préfèrent l’auto-hébergement plutôt que l’API propriétaire d’Alibaba, l’argument économique change de nature : il n’y a plus de coût par token, mais un coût d’infrastructure GPU, largement compensé par la faible proportion de paramètres actifs par requête.

## Les benchmarks disponibles : DeepSWE, SWE-bench Pro et les inconnues

Sur le terrain des chiffres, il faut être précis pour ne pas céder à la surenchère marketing qui entoure souvent ces annonces. Alibaba revendique, en auto-évaluation, un score de 58,7 sur le benchmark DeepSWE et de 62,5 sur SWE-bench Pro, deux tests centrés sur les capacités d’ingénierie logicielle. Ce sont des scores solides pour un modèle à seulement 6 milliards de paramètres actifs, mais à ce stade, **aucune source consultée ne publie de tableau chiffré comparant directement Qwen3.8-Flash-Next à GPT-6 Astra, Claude Fable 5.1, Gemini 3.8 Flash ou DeepSeek V4** sur des benchmarks communs et vérifiés. Les analyses disponibles évoquent une performance “compétitive avec des modèles denses bien plus lourds” en code et en raisonnement, une formulation qualitative qu’il convient de traiter comme telle plutôt que comme un classement définitif.

Sur le podium des modèles généralistes de septembre 2026, c’est toujours le trio GPT-6 Astra, Claude Opus 5 et Gemini 3.8 Flash qui domine les classements toutes catégories, avec des tarifs d’API très supérieurs à ceux de Qwen3.8-Flash-Next. La comparaison pertinente pour ce dernier se situe plutôt face à d’autres modèles ouverts orientés production à bas coût, comme GLM-5.3-Flash ou les déclinaisons Flash de DeepSeek V4.

## Qwen3.8-Flash-Next face à Qwen3.8-Max : deux stratégies opposées

