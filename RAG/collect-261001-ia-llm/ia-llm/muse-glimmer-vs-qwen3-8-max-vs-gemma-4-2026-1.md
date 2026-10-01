---
id: collect-261001-ia-llm/ia-llm/muse-glimmer-vs-qwen3-8-max-vs-gemma-4-2026-1
title: "Exemple : servir Muse Glimmer quantifié en local avec llama.cpp"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Mistral", "Moonshot", "Z.ai"]
dates: []
keywords: ["muse", "agent", "apache", "attention", "benchmarks", "claude", "compute", "deepseek", "fp8", "gemini", "glm", "gpu"]
source: docs/RAG/collect-261001-ia-llm/muse-glimmer-vs-qwen3-8-max-vs-gemma-4-2026.md
source_anchor: ""
source_lines: [1, 30]
sha256: 1dd8f5505a45716433d5630d98d6ef27f05fa4a3ceec8e91da050b94f044a48e
---

# Exemple : servir Muse Glimmer quantifié en local avec llama.cpp

Trois nouveaux modèles à poids ouverts ont débarqué en l’espace de dix semaines cet été 2026, et ils ne visent pas du tout le même public. Meta a publié **Muse Glimmer**, un modèle dense de 30 milliards de paramètres pensé pour tourner sur un seul GPU grand public. Alibaba a mis en ligne les poids de **Qwen3.8-Max**, un mastodonte MoE de 2,4 billions de paramètres réservé aux clusters de data center. Et Google DeepMind a sorti **Gemma 4**, un modèle multimodal de 12 milliards de paramètres à l’architecture inédite, sans encodeur séparé pour l’image ou l’audio. Trois tailles, trois licences, trois philosophies de déploiement, pour un seul objectif commun : proposer une alternative crédible aux API propriétaires pour les équipes européennes qui veulent héberger leur IA localement ou dans l’UE.

Ce comparatif détaille les spécifications techniques, les benchmarks disponibles, les licences, les coûts de déploiement et les cas d’usage réels de ces trois modèles, avec un focus particulier sur ce que cela change pour les développeurs, startups et entités publiques en France et en Europe qui cherchent une alternative aux modèles fermés de Claude, GPT ou Gemini hébergés hors UE.

## Pourquoi comparer Muse Glimmer, Qwen3.8-Max et Gemma 4 maintenant

La fenêtre du 3 juin au 13 août 2026 a vu se succéder trois publications qui redessinent la carte des LLM à poids ouverts. Google DeepMind a dégainé le premier avec Gemma 4 12B, présenté début juin comme “un modèle multimodal unifié, sans encodeur”, où les entrées image et audio passent directement dans le tronc du transformeur sans passer par une brique de vision séparée. Alibaba a suivi début août avec Qwen3.8-Max, annoncé le 3 août puis publié en poids ouverts sous la référence Qwen3.8-2.4T-A95B le 12 et 13 août. Meta a clos la séquence le 10 août avec Muse Glimmer, un modèle dense distillé de la famille Muse Spark, taillé pour tourner sur un seul GPU de 24 Go.

Cette accélération n’est pas un hasard de calendrier. Elle traduit une bascule stratégique chez les trois éditeurs : proposer des poids téléchargeables devient un argument commercial à part entière face à la pression des modèles chinois (DeepSeek V4, GLM-5.3, Kimi K3) qui dominaient jusque-là le classement des modèles ouverts sur des plateformes comme BenchLM ou le comparatif de Thunder Compute publié en août 2026. Pour les équipes techniques françaises et européennes, cette vague change concrètement les arbitrages : faut-il continuer à payer un abonnement API vers un fournisseur américain, ou héberger soi-même un modèle sous licence Apache 2.0 dans un data center européen ? Les trois modèles de ce comparatif répondent chacun différemment à cette question, avec des contraintes matérielles radicalement différentes.

Il faut aussi noter que ni Muse Glimmer, ni Qwen3.8-Max, ni Gemma 4 ne sont des modèles européens. Sur le classement spécifique aux modèles européens de BenchLM, mis à jour au 28 août 2026, c’est **Ministral 3 14B (Reasoning)** de Mistral AI qui occupe la première place avec un score de 50, loin devant les alternatives non-UE sur ce classement dédié. Ce comparatif garde donc une dimension pratique : ces trois modèles sont des options d’auto-hébergement pertinentes pour la conformité RGPD (les données ne quittent pas votre infrastructure), mais ils ne remplacent pas un modèle développé en Europe pour les cas d’usage où la provenance du modèle lui-même compte, notamment dans le cadre de l’AI Act.

## Muse Glimmer 30B de Meta : l’agent local sur un seul GPU

Muse Glimmer est un transformeur causal dense, sans architecture MoE, comptant environ 29,6 à 30 milliards de paramètres au total : à peu près 28 milliards pour le décodeur textuel et 1,8 à 2 milliards pour la tour de perception visuelle dédiée. Meta décrit une architecture à 52 couches de décodeur, une dimension cachée de 6 656, un réseau feed-forward SwiGLU et une attention par groupes de requêtes (GQA) avec 32 têtes de requête pour seulement 2 têtes clé-valeur, soit un ratio proche de 16:1, alterné selon un motif d’attention locale, locale, locale, globale.

La fenêtre de contexte annoncée est de 131 072 tokens. Le modèle accepte du texte et des images en entrée, via un encodeur de perception de type ViT capable de traiter jusqu’à 4 096 tokens visuels par image, et certaines analyses techniques évoquent une prise en charge de séquences vidéo allant jusqu’à 96 images à 2 images par seconde. La sortie reste exclusivement textuelle : Muse Glimmer ne génère ni image ni audio.

Le point différenciant, c’est le gabarit mémoire. En précision complète BF16, le modèle pèse environ 55 à 60 Go, ce qui nécessite un GPU professionnel ou plusieurs cartes. Mais quantifié en 4 bits, Muse Glimmer tombe sous les 20 Go, ce qui le fait tenir sur une carte grand public de 24 Go comme une RTX 4090 ou une RTX 5090, ou sur un Mac disposant de suffisamment de mémoire unifiée. C’est exactement le positionnement revendiqué par Meta : un agent capable de tourner en permanence sur une machine personnelle plutôt que dans le cloud. Sur ce point, Eigent AI résume la promesse du modèle comme un moyen de faire tourner “un agent IA disponible en permanence sur votre propre ordinateur portable, pas dans le cloud” (source).

Côté performance brute, une analyse technique de Developers Digest rapporte un débit de 233,4 tokens par seconde sur une RTX 5090, grâce au décodage spéculatif DFlash livré avec le modèle, une technique qui accélère la génération en pré-calculant plusieurs tokens candidats avant validation (source). Sur les benchmarks classiques, les chiffres publics restent parcellaires à la date de rédaction : plusieurs analystes évoquent un score MMLU autour de 80 à 82 %, ce qui le positionnerait dans la fourchette haute des modèles ouverts de 20 à 40 milliards de paramètres, mais aucune fiche technique officielle ne détaille pour l’instant GPQA Diamond, SWE-Bench ou AIME 2025 pour ce modèle précis. Meta a surtout évalué Muse Glimmer sur des tâches agentiques (usage d’outils, lecture de captures d’écran, récupération après échec) plutôt que sur les suites d’examens statiques classiques.

Sur le plan de la licence, Muse Glimmer est publié sous **Apache 2.0**, sans clause de seuil de revenu ni restriction d’usage commercial particulière, ce qui en fait l’option la plus simple juridiquement des trois pour une entreprise européenne souhaitant l’intégrer dans un produit.

## Qwen3.8-Max Open Weights d’Alibaba : la puissance à l’échelle cluster

Qwen3.8-Max change complètement de catégorie. Le modèle a été annoncé le 3 août 2026 comme le nouveau porte-drapeau propriétaire d’Alibaba, avant que la variante à poids ouverts, baptisée Qwen3.8-2.4T-A95B, ne soit publiée sur Hugging Face les 12 et 13 août, en versions BF16 et FP8, chacune répartie sur 213 fichiers de poids. L’architecture est un mélange d’experts (MoE) totalisant 2,4 billions de paramètres, dont seulement 95 milliards sont activés à chaque passage, ce qui explique pourquoi le coût de calcul par requête reste comparable à un modèle dense d’environ 95 milliards de paramètres, même si le stockage complet des experts reste massif.

La fenêtre de contexte native de la variante à poids ouverts est annoncée à 256 000 tokens, contre jusqu’à 1 million de tokens pour la version API propriétaire Qwen3.8-Max accessible via le cloud d’Alibaba. Cette variante open weights est, à ce stade, présentée comme un modèle essentiellement textuel : les sources techniques disponibles ne mentionnent pas de prise en charge image ou audio native pour ce checkpoint précis, contrairement à d’autres membres de la famille Qwen.

