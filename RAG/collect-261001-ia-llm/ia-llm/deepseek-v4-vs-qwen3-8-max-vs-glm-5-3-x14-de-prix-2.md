---
id: collect-261001-ia-llm/ia-llm/deepseek-v4-vs-qwen3-8-max-vs-glm-5-3-x14-de-prix-2
title: "DeepSeek V4 Flash 0731 (via API officielle)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Together AI", "Z.ai"]
dates: []
keywords: ["deepseek", "agent", "benchmark", "benchmarks", "claude", "fable 5", "gemini", "glm", "kimi", "moe", "open source", "qwen"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v4-vs-qwen3-8-max-vs-glm-5-3-x14-de-prix.md
source_anchor: ""
source_lines: [37, 96]
sha256: cd062383a2b954f5c4a42456a24393b4b5c13e41c37251717a207301daa49a56
---

# DeepSeek V4 Flash 0731 (via API officielle)

La fenêtre de contexte atteint elle aussi 1 million de tokens, avec une sortie maximale de 128 000 tokens. Sur les benchmarks propres à Zhipu, GLM-5.3 obtient **28,3 sur Terminal-Bench 3.0** (une version plus récente et donc non directement comparable au Terminal-Bench 2.1 utilisé par DeepSeek et Qwen) et **66,9 sur DeepSWE 1.1**, un résultat que Zhipu présente comme le meilleur parmi les modèles open source sur cette suite, devant Kimi K3. En interne, Z.ai revendique un gain de **50 % sur son benchmark de code maison** par rapport à GLM-5.2, un chiffre à prendre avec la prudence habituelle réservée aux benchmarks propriétaires non audités.

Le vrai atout différenciant de GLM-5.3 est sa compatibilité multi-protocole. Le modèle expose à la fois un endpoint OpenAI Chat Completions, un endpoint OpenAI Responses et un **endpoint compatible Anthropic Messages** (accessible via `/api/anthropic`), ce qui permet de le brancher directement sur des outils construits pour Claude, comme Claude Code, sans réécrire la couche d’intégration. C’est une décision stratégique claire : capter les équipes déjà outillées autour de l’écosystème Anthropic sans qu’elles aient à migrer leur code.

Sur le prix, GLM-5.3 conserve la grille tarifaire de GLM-5.2 : **1,40 $** par million de tokens en entrée, **0,26 $** en entrée mise en cache, et **4,40 $** en sortie. Pour un usage typique en boucle d’agent où la majorité du contexte est relu depuis le cache, le coût d’entrée effectif tombe autour de 0,60 $ par million de tokens selon les estimations de plusieurs analystes indépendants. Au lancement, seule l’API hébergée est disponible : Z.ai prévoit de publier les poids ouverts sur Hugging Face autour du 28 août 2026, deux semaines après l’annonce initiale.

## Tableau comparatif : spécifications techniques des trois modèles

Voici l’ensemble des caractéristiques techniques vérifiées pour chaque modèle, à date du 22 août 2026.

| Caractéristique | DeepSeek V4 Flash 0731 | Qwen3.8 Max | GLM-5.3 | 
|---|---|---|---|
| Développeur | DeepSeek | Alibaba | Zhipu AI (Z.ai) | 
| Date de sortie | 31 juillet 2026 | 2 août 2026 | 14 août 2026 | 
| Architecture | MoE, 284 Md paramètres | MoE, 2,4 T paramètres | MoE, ~743 Md paramètres | 
| Paramètres actifs | ~13 Md / token | 95 Md / requête | ~40 Md / token | 
| Fenêtre de contexte | 1 000 000 tokens | 1 000 000 tokens | 1 000 000 tokens | 
| Sortie maximale | 384 000 tokens | 131 072 tokens | 128 000 tokens | 
| Modalités d’entrée | Texte (mode thinking/non-thinking) | Texte, image, vidéo | Texte (agentique, code) | 
| Licence à date | Open weight, MIT | API seule (poids annoncés) | API seule (poids annoncés ~28 août) | 
| Prix entrée (cache miss) | 0,14 $ / 1M tokens | 2,00 $ / 1M tokens | 1,40 $ / 1M tokens | 
| Prix entrée (cache hit) | 0,0028 $ / 1M tokens | Non communiqué | 0,26 $ / 1M tokens | 
| Prix sortie | 0,28 $ / 1M tokens | 6,00 $ / 1M tokens | 4,40 $ / 1M tokens | 
| Endpoints compatibles | API propre, style OpenAI | API OpenAI-compatible | OpenAI + Anthropic Messages | 
| Disponibilité | API DeepSeek, OpenRouter, Together AI | Alibaba Cloud Model Studio, QwenCloud | Z.ai API, GLM Coding Plan | 

Ce tableau met en évidence un point souvent négligé dans la couverture anglophone : les trois modèles convergent tous vers une fenêtre de contexte d’un million de tokens. En 2025, ce chiffre était encore réservé aux modèles les plus chers. En août 2026, il devient un standard même sur les offres les moins onéreuses du marché chinois.

## Benchmarks : que disent réellement les scores publiés

Comparer des benchmarks entre trois laboratoires qui ne publient pas les mêmes suites de tests est un exercice à manier avec prudence. Terminal-Bench 2.1 et Terminal-Bench 3.0, par exemple, ne sont pas la même échelle : un score de 28,3 sur la version 3.0 de GLM-5.3 n’est pas directement comparable à un score de 82,7 sur la version 2.1 de DeepSeek V4 Flash. Voici néanmoins l’ensemble des chiffres vérifiés, avec leur source, pour que chacun puisse se faire son idée.

| Benchmark | DeepSeek V4 Flash 0731 | Qwen3.8 Max | GLM-5.3 | 
|---|---|---|---|
| Terminal-Bench 2.1 | 82,7 % | 86,6 | Non testé sur cette version | 
| Terminal-Bench 3.0 | Non communiqué | Non communiqué | 28,3 | 
| SWE-Bench Pro | Non communiqué | 67,7 | Non communiqué | 
| GPQA Diamond | Non communiqué | 92,6 | Non communiqué | 
| OSWorld-Verified | Non communiqué | 86,1 | Non communiqué | 
| DeepSWE 1.1 | Non communiqué | Non communiqué | 66,9 | 
| Gain vs génération précédente | +20,9 pts sur Terminal-Bench 2.1 vs préversion d’avril | Nouvelle échelle (2,4T vs Qwen3.7 non chiffré publiquement) | +50 % sur benchmark code interne Z.ai vs GLM-5.2 | 

Trois sources indépendantes confirment ces chiffres : le rapport technique de MarkTechPost pour le score Terminal-Bench 2.1 de DeepSeek, l’agrégateur de spécifications qui recense les scores GPQA Diamond et SWE-Bench Pro de Qwen3.8 Max, et le tracker de sorties de modèles BenchLM qui documente le calendrier de publication des trois modèles. Le constat qui ressort : Qwen3.8 Max est de loin le mieux documenté sur les benchmarks de raisonnement scientifique et de résolution de bugs logiciels, tandis que DeepSeek et GLM se concentrent sur des métriques agentiques et de terminal qui valorisent leur spécialisation coût-efficacité et code respectivement.

Un classement français publié par Blog du Modérateur plaçait mi-août Qwen 3.8 Max en 8e position et Gemini 3.7 Flash High en 9e position de son top 20 des modèles les plus performants sur Text Arena, aux côtés de Claude Fable 5 en tête. Ce classement généraliste confirme que Qwen3.8 Max s’installe déjà dans le haut du panier, alors que GLM-5.3 et DeepSeek V4 Flash, plus récents ou plus spécialisés, n’y figuraient pas encore à cette date.

## Tarification détaillée : quel modèle coûte le moins cher en production

Le prix par million de tokens ne raconte qu’une partie de l’histoire. Ce qui compte en production, c’est le coût réel d’une charge de travail type. Prenons un scénario courant pour un agent de support ou de code : 10 millions de tokens d’entrée (essentiellement du contexte répété) et 2 millions de tokens de sortie sur un mois.

| Modèle | Entrée (cache miss) | Entrée (cache hit) | Sortie | Coût estimé 10M in + 2M out | 
|---|---|---|---|---|
| DeepSeek V4 Flash 0731 | 0,14 $ / 1M | 0,0028 $ / 1M | 0,28 $ / 1M | ~1,96 $ (hors cache) / ~0,59 $ (avec cache à 90 %) | 
| GLM-5.3 | 1,40 $ / 1M | 0,26 $ / 1M | 4,40 $ / 1M | ~22,80 $ (hors cache) / ~11,60 $ (avec cache à 90 %) | 
| Qwen3.8 Max | 2,00 $ / 1M | Non communiqué | 6,00 $ / 1M | ~32,00 $ (tarif marché, sans cache documenté) | 

L’écart entre le moins cher et le plus cher atteint **14 fois** sur le prix de sortie brut (0,28 $ contre 4,40 $ pour le cache miss d’entrée le plus onéreux face au moins cher), et grimpe encore plus haut si l’on compare le tarif de cache hit de DeepSeek face au tarif d’entrée standard de Qwen3.8 Max. Pour une startup française qui fait tourner des milliers de requêtes agentiques par jour, la différence se chiffre vite en dizaines de milliers d’euros sur l’année. C’est un argument qui pèse d’autant plus lourd dans un contexte où la réglementation AI Act impose déjà des coûts de conformité supplémentaires aux éditeurs européens.

## Licence et disponibilité : poids ouverts contre API fermée

