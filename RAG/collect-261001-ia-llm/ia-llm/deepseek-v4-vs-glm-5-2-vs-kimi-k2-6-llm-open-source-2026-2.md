---
id: collect-261001-ia-llm/ia-llm/deepseek-v4-vs-glm-5-2-vs-kimi-k2-6-llm-open-source-2026-2
title: "Auto-hebergement avec vLLM (exemple GLM-5.2 en FP8)"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Z.ai", "vLLM", "xAI"]
dates: []
keywords: ["fp8", "glm", "vllm", "agents", "attention", "attribution", "benchmark", "benchmarks", "chatgpt", "deepseek", "gemini", "gpu"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v4-vs-glm-5-2-vs-kimi-k2-6-llm-open-source-2026.md
source_anchor: ""
source_lines: [56, 123]
sha256: c2249b8082c243816727c3bc94b78bf6ebd2d4866e454a6ba80474b28e4f62c0
---

# Auto-hebergement avec vLLM (exemple GLM-5.2 en FP8)

Kimi K2.6 de Moonshot AI est un MoE d’un billion de paramètres, dont 32 milliards actifs par token. Sa singularité tient à deux points : il traite nativement le texte, l’image et la vidéo dans une seule architecture, sans module de vision séparé, et il est livré nativement en quantification INT4, ce qui réduit l’empreinte mémoire à l’inférence. C’est aussi le modèle le mieux classé de notre trio sur l’Artificial Analysis Intelligence Index, avec un score de 54. Ses poids sont publics sur Hugging Face.

### GLM-5.2 : le spécialiste du codage à long horizon

GLM-5.2 est un MoE de 753 milliards de paramètres, publié le 13 juin 2026 avec une build FP8 quantifiée pour un déploiement plus léger. Contrairement à ses deux rivaux, Z.ai a concentré son marketing sur une promesse unique : le codage agentique à long horizon. Le modèle vise explicitement les agents de développement autonomes capables d’enchaîner des dizaines d’actions sur une base de code réelle. Ses poids MIT sont disponibles sous zai-org/GLM-5.2. C’est le plus récent des trois, et son arrivée a redéfini le sommet du **LLM open source** agentique.

## Benchmarks : SWE-bench, GPQA et Intelligence Index décryptés

C’est ici que les trajectoires divergent. Le piège classique consiste à comparer des scores qui ne mesurent pas la même chose. Nous distinguons donc trois familles de tests : le codage vérifié (SWE-bench Verified), le codage agentique long (SWE-bench Pro), et le raisonnement scientifique (GPQA Diamond), avant de revenir à l’indice composite d’Artificial Analysis.

| Benchmark | GLM-5.2 | DeepSeek V4-Pro | Kimi K2.6 | Référence propriétaire | 
|---|---|---|---|---|
| AA Intelligence Index v4.0 | 51,1 | 52 | 54 | – | 
| SWE-bench Verified | non publié | 80,6 % | 80,2 % | Gemini 3.1 Pro : 80,6 % | 
| SWE-bench Pro | 62,1 % | non publié | 58,6 % | GPT-5.5 : 58,6 % | 
| GPQA Diamond | 89,5 % | 90,1 % | non publié | – | 
| LiveCodeBench | non publié | 93,5 % | non publié | – | 
| MMLU-Pro | non publié | 87,5 % | non publié | – | 
| Terminal-Bench 2.1 | 81,0 % | non publié | non publié | – | 

**SWE-bench Verified – DeepSeek V4 en tête.** DeepSeek-V4-Pro-Max atteint 80,6 % sur SWE-bench Verified, le meilleur score jamais enregistré par un modèle à poids ouverts, à égalité avec Gemini 3.1 Pro (source : Artificial Analysis). Kimi K2.6 le talonne à 80,2 %. GLM-5.2 n’a pas publié de score sur ce test précis.

**SWE-bench Pro – GLM-5.2 devant GPT-5.5.** Sur le test agentique long, plus difficile et plus représentatif du travail réel d’ingénierie, GLM-5.2 obtient 62,1 %, devançant GPT-5.5 (58,6 %) et Kimi K2.6 (58,6 %). C’est la première fois qu’un modèle sous licence MIT dépasse un modèle phare d’OpenAI sur ce benchmark que les acheteurs surveillent réellement. Kimi K2.6 se place lui-même devant GPT-5.4 (57,7 %) et Gemini 3.1 Pro (54,2 %).

**Raisonnement scientifique – quasi-égalité.** Sur GPQA Diamond, DeepSeek V4-Pro (90,1 %) et GLM-5.2 (89,5 %) se tiennent en moins d’un point. DeepSeek complète le tableau avec 93,5 % sur LiveCodeBench, 87,5 % sur MMLU-Pro et 95,2 % sur le concours de mathématiques HMMT 2026 – de loin le profil de benchmark le plus complet et le mieux documenté du trio.

**Le juge de paix : l’Intelligence Index.** Pour trancher entre des scores qui ne se recouvrent pas, l’indice composite d’Artificial Analysis reste la meilleure boussole. Kimi K2.6 mène à 54, devant DeepSeek V4-Pro (Max) à 52 et GLM-5.2 à 51,1. Trois points d’écart : autant dire une égalité statistique. Le choix ne se fera donc pas sur « qui est le plus intelligent », mais sur le prix, la licence et l’usage cible.

## Prix et coût réel : DeepSeek V4 casse le marché

C’est le nerf de la guerre pour toute équipe qui passe en production. Les trois modèles étant open source, deux modes de coût coexistent : l’API managée (facturée au token) et l’auto-hébergement (coût d’infrastructure fixe, poids gratuits). Commençons par les tarifs officiels des API.

| Modèle | Entrée ($/M) | Sortie ($/M) | Coût mixte 3:1 estimé* | Contexte | 
|---|---|---|---|---|
| **DeepSeek V4-Flash** | 0,14 | 0,28 | ≈ 0,18 $ | 1 M | 
| **DeepSeek V4-Pro** | 0,435 | 0,87 | ≈ 0,54 $ | 1 M | 
| **Kimi K2.6** | 0,60 | 2,50 | ≈ 1,08 $ | 256 K | 
| **GLM-5.2** (Z.ai) | 1,40 | 4,40 | ≈ 2,15 $ | 1 M | 
| GLM-5.2 (OpenRouter) | 0,77 | 2,42 | ≈ 1,18 $ | 1 M | 

**Coût mixte calculé sur un ratio de 3 tokens d’entrée pour 1 de sortie (convention Artificial Analysis). Chiffres arrondis, à titre indicatif.*

Le verdict prix est sans appel : **DeepSeek V4** écrase la concurrence. À 0,435 $ le million de tokens en entrée, la variante Pro est près de cinq fois moins chère que l’API officielle de GLM-5.2, et la variante Flash tombe à 0,14 $. DeepSeek a annoncé le 22 mai 2026 que sa remise de 75 %, initialement promotionnelle, devenait permanente – un tarif désormais confirmé à l’échelle production depuis le passage en disponibilité générale de **DeepSeek-V4-Pro-0813** les 12 et 13 août 2026. Attention toutefois : le prix officiel Z.ai de GLM-5.2 (1,40 $/4,40 $) est trompeur, car des revendeurs comme OpenRouter le proposent à 0,77 $/2,42 $, et l’entrée mise en cache descend à 0,26 $ le million de tokens. Kimi K2.6, à 0,60 $/2,50 $, se situe entre les deux.

Pour l’auto-hébergement, la logique s’inverse : les poids sont gratuits, seul compte le coût GPU. Kimi K2.6 (INT4 natif) et GLM-5.2 (FP8) sont les plus économes en VRAM à qualité égale, tandis que DeepSeek V4-Pro, avec ses 1,6 T de paramètres, exige l’infrastructure la plus lourde. Notre tutoriel Ollama pour exécuter un LLM en local détaille la mise en place sur poste de travail ; en production, vLLM reste la référence.

```
# Auto-hebergement avec vLLM (exemple GLM-5.2 en FP8)
pip install vllm
vllm serve zai-org/GLM-5.2 \
  --quantization fp8 \
  --max-model-len 131072 \
  --tensor-parallel-size 8 \
  --served-model-name glm-5.2
# L'API exposee est compatible OpenAI :
# POST http://localhost:8000/v1/chat/completions
```
## Licences et poids ouverts : MIT, MIT modifiée et implications commerciales

Le terme « open source » recouvre des réalités différentes. Les trois modèles publient bien leurs poids – c’est-à-dire qu’on peut les télécharger, les auto-héberger, les affiner et les utiliser commercialement – mais les conditions varient.

- **GLM-5.2 et DeepSeek V4 : licence MIT stricte.** Usage commercial libre, modification, redistribution, aucune redevance, aucune obligation d’attribution dans l’interface. C’est le régime le plus permissif possible pour un produit commercial.
- **Kimi K2.6 : licence MIT modifiée.** Une clause s’ajoute : si vous déployez Kimi K2.6 (ou un dérivé) dans un produit qui dépasse 100 millions d’utilisateurs actifs mensuels*ou* génère plus de 20 millions de dollars de revenus mensuels, vous devez afficher visiblement la mention « Kimi K2 » dans l’interface. En dessous de ces seuils, la licence fonctionne comme une MIT standard.

Concrètement, pour 99 % des entreprises françaises – PME, ETI, startups, éditeurs SaaS – la nuance est sans effet : les trois licences autorisent un usage commercial complet. La clause Kimi ne mord que sur les très grands acteurs grand public. En revanche, la disponibilité des poids change tout pour la conformité : un modèle auto-hébergé traite les données *sur votre infrastructure*, sans qu’aucun prompt ne quitte l’Union européenne. C’est l’argument massue face aux API propriétaires américaines comme Grok 4, ChatGPT ou Gemini 3.

## Contexte long et agents autonomes : le vrai terrain de jeu 2026

