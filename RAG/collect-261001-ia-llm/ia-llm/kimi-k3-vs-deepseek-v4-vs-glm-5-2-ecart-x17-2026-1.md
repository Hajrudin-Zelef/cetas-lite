---
id: collect-261001-ia-llm/ia-llm/kimi-k3-vs-deepseek-v4-vs-glm-5-2-ecart-x17-2026-1
title: "kimi-k3-vs-deepseek-v4-vs-glm-5-2-ecart-x17-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "glm", "kimi", "attention", "attribution", "benchmark", "benchmarks", "claude", "fable 5", "fine-tuning", "gemini", "gpt-5.6"]
source: docs/RAG/collect-261001-ia-llm/kimi-k3-vs-deepseek-v4-vs-glm-5-2-ecart-x17-2026.md
source_anchor: ""
source_lines: [1, 34]
sha256: 7adaa39d40cd366f504418e6405f71557f3988eb9fe34c3961f9bcb8eb94210e
---

# kimi-k3-vs-deepseek-v4-vs-glm-5-2-ecart-x17-2026

Trois laboratoires chinois viennent de rebattre les cartes de l’IA open-weight en l’espace de trois mois. **Kimi K3** de Moonshot AI, **DeepSeek V4 Pro** et **GLM-5.2** de Zhipu AI affichent tous une fenêtre de contexte à un million de tokens et des scores qui rivalisent avec les modèles fermés facturés cinq à trente fois plus cher. Mais entre le modèle le plus capable, le moins cher et le plus rapide, l’écart de prix API grimpe jusqu’à **17 fois** selon la méthode de calcul retenue. Selon une analyse publiée par MarkTechPost, ces trois modèles occupent désormais le sommet du classement open-weight mondial, une position qu’aucun trio de modèles ouverts n’avait atteinte simultanément auparavant. Ce comparatif détaille les specs, les benchmarks croisés de trois évaluateurs indépendants, la tarification réelle et les cas d’usage pour aider les équipes techniques européennes à choisir sans se tromper.

## Le contexte : la course aux modèles IA open-weight s’accélère en 2026

Depuis avril 2026, le marché des grands modèles de langage ouverts a changé de dimension. DeepSeek a ouvert le bal le 24 avril avec **DeepSeek V4 Pro**, suivi par Zhipu AI (Z.ai) qui a publié **GLM-5.2** le 13 juin, puis Moonshot AI qui a livré **Kimi K3** le 16 juillet. Les trois modèles partagent un point commun frappant : ce sont des architectures Mixture-of-Experts (MoE) à l’échelle du trillion de paramètres, avec une licence MIT ou dérivée et un contexte d’un million de tokens. Selon Moonshot AI, K3 est le premier modèle open-weight de la « classe 3T ».

Ce basculement compte pour l’Europe. Pendant que Claude Opus 5 et GPT-5.6 dominent le haut du classement fermé, et que la guerre des prix entre GPT-5.6 et Opus 5 fait chuter les tarifs des API propriétaires, les modèles open-weight chinois offrent désormais une alternative crédible pour les équipes qui veulent maîtriser leurs coûts d’inférence ou héberger leurs propres poids pour des raisons de conformité RGPD. Le benchmark LARA sur la conformité à l’AI Act a d’ailleurs rappelé que la traçabilité et le contrôle des données restent un point sensible pour les entreprises européennes, ce qui pousse certaines à regarder du côté de l’auto-hébergement.

Sur l’indice composite Artificial Analysis Intelligence Index, qui évalue tous les modèles sur la même suite de tests, Kimi K3 obtient un score d’environ **57**, DeepSeek V4 Pro (mode Max reasoning) **44**, et GLM-5.2 **51**. K3 se classe parmi les meilleurs modèles toutes catégories confondues, devant Claude Opus 4.8 (56) et juste derrière Claude Fable 5 et GPT-5.6 Sol. C’est la première fois qu’un modèle open-weight s’approche autant du sommet du classement général.

## Kimi K3 de Moonshot AI : le poids lourd à 2,8 billions de paramètres

Kimi K3 est un modèle Stable LatentMoE de **2,8 billions de paramètres** au total, avec environ 50 milliards de paramètres actifs par requête (16 experts activés sur 896 disponibles). Moonshot n’a pas publié le nombre exact de paramètres actifs dans sa documentation officielle, mais Artificial Analysis et DeepInfra convergent sur ce chiffre. L’architecture repose sur deux innovations maison : Kimi Delta Attention (KDA) et les Attention Residuals, qui apportent selon Moonshot un gain d’efficacité d’environ 2,5x par rapport à la génération précédente, K2.6.

K3 est le seul des trois modèles à embarquer une modalité vision nativement dès son lancement, ce qui en fait un candidat naturel pour des cas d’usage multimodaux (analyse de documents scannés, capture d’écran, génération de contenu vidéo léger). Sa fenêtre de contexte atteint 1 million de tokens, à égalité avec ses deux rivaux.

Côté benchmarks, K3 prend la tête sur l’indice composite Artificial Analysis (57 points) et sur le classement humain Arena.ai Frontend Code Arena, où il dépasse Claude Fable 5 et GPT-5.6 Sol. Des tests indépendants de BenchLM placent son agrégat BenchAlign à 80,96, avec un point fort marqué sur les tâches agentiques (89,5), loin devant DeepSeek V4 Pro (59,1) sur cette même mesure. En revanche, son score SWE-bench Verified de **76,8 %** reste inférieur à celui de DeepSeek V4 Pro (80,6 %), ce qui illustre qu’un score composite élevé ne garantit pas la première place sur chaque benchmark pris isolément.

Le vrai bémol de K3 en août 2026 reste la disponibilité de ses poids. Moonshot s’est engagé à les publier le **27 juillet 2026** sous une licence Modified MIT, avec une clause d’attribution qui ne s’active qu’au-delà de 100 millions d’utilisateurs actifs mensuels (donc sans incidence pour la quasi-totalité des entreprises). En attendant, K3 n’est accessible que via API ou l’application Kimi, ce qui exclut d’emblée les déploiements air-gapped ou les contraintes strictes de souveraineté des données.

## DeepSeek V4 Pro : le champion incontesté du rapport qualité-prix

DeepSeek V4 Pro, lancé le 24 avril 2026, est un modèle MoE de **1,6 billion de paramètres** avec 49 milliards de paramètres actifs, répartis sur 384 experts routés plus un expert partagé. Son architecture Compressed Sparse Attention (CSA) combinée à la Heavily Compressed Attention (HCA) réduit les besoins en cache KV à environ 10 % de ce que demandait l’architecture précédente V3.2 sur un contexte d’un million de tokens. Trois modes de raisonnement configurables (Non-Think, Think High, Think Max) permettent d’arbitrer entre latence et profondeur de raisonnement directement au niveau de l’appel API.

Sur le terrain de la programmation compétitive, V4 Pro écrase la concurrence : **93,5 % sur LiveCodeBench**, la meilleure note toutes catégories mondiales, et un classement Codeforces de 3206, un niveau qui correspond à une résolution de problèmes de compétition de très haut niveau. Son score SWE-bench Verified de **80,6 %** est le plus élevé des trois modèles comparés ici, à égalité avec Gemini 3.1 Pro selon les données publiées au moment de son lancement. Sur GPQA Diamond, il obtient 90,1 %, un cran sous Kimi K3 (93,5 %) mais dans la fourchette de Claude Opus 4.8.

Le point faible documenté de V4 Pro mérite d’être mentionné clairement : sur le benchmark AA-Omniscience d’Artificial Analysis, qui mesure la calibration d’un modèle (sait-il reconnaître qu’il ne sait pas ?), V4 Pro affiche un **taux d’hallucination de 94 %**, c’est-à-dire qu’il produit une réponse quasiment à chaque fois, indépendamment de sa certitude réelle. Pour des cas d’usage où la fiabilité des réponses compte plus que le débit de tokens bon marché, ce chiffre doit peser dans la décision.

Les poids de V4 Pro sont disponibles dès aujourd’hui sur Hugging Face sous licence MIT, ce qui autorise l’usage commercial, le fine-tuning et l’auto-hébergement sans restriction. Une variante plus légère, **DeepSeek V4 Flash** (284 milliards de paramètres au total, 13 milliards actifs), cible les workloads à très gros volume avec un tarif encore plus agressif.

## GLM-5.2 de Zhipu AI : le sprinteur le plus rapide du trio

GLM-5.2, sorti le 13 juin 2026, est le plus compact des trois avec environ **744 milliards de paramètres** au total (753 milliards selon Artificial Analysis) et environ 40 milliards de paramètres actifs. Son mécanisme de routage IndexShare, associé à une couche de prédiction multi-token (MTP) et au décodage spéculatif KVShare, porte selon Zhipu la longueur d’acceptation des tokens de brouillon jusqu’à +20 %. Résultat concret : GLM-5.2 tourne à environ **168 tokens par seconde** sur les mesures Artificial Analysis, soit près de trois fois plus vite que Kimi K3 et DeepSeek V4 Pro, tous deux autour de 62 tokens/seconde.

