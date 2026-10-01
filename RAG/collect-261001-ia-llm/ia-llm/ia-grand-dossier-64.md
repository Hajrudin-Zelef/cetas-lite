---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-64
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "CoreWeave", "Crusoe", "DeepSeek", "Google", "Groq", "Lambda", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "vLLM"]
dates: []
keywords: ["agent", "attention", "awq", "aws", "benchmarks", "blackwell", "claude", "compute", "datacenter", "deepseek", "diffusion", "embeddings"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5393, 5442]
sha256: e92c9393114fc75aae8b8195044d9a8e014f1d8c77a828b3cb9aaeb00cf23e49
---

# IA — Le grand dossier

49. **GAN** — Generative Adversarial Network (Goodfellow et al., 2014) : générateur vs discriminateur. A dominé la génération d'images avant la diffusion.
50. **Gateway (passerelle)** — Proxy unifié (LiteLLM, OpenRouter) devant N providers : routage, fallbacks, budgets, observabilité.
51. **GGUF** — Format de fichier de llama.cpp ; une seule archive = poids quantifiés + métadonnées.
52. **Goodhart (loi de)** — « Quand une mesure devient un objectif, elle cesse d'être une bonne mesure. » S'applique aux benchmarks et aux reward models.
53. **GPAI** — « General-Purpose AI » : modèles à usage général (GPT, Claude, Gemini…). Régime spécifique dans l'AI Act (transparence, évaluation si risque systémique).
54. **GPTQ** — Ancien standard de quantification 4-bit ; encore présent dans les modèles pré-quantifiés.
55. **Grounding (ancrage)** — Relier les sorties du modèle à des sources vérifiables (documents RAG, citations). L'antidote aux hallucinations.
56. **Guardrails** — Garde-fous : règles/filtres encadrant un système IA (refus de sujets, validation des sorties, limites d'outils).
57. **H100 / B200 / H200** — GPU datacenter NVIDIA (Hopper/Blackwell) ; 80–192 Go HBM ; 2,5–12 $/h en cloud.
58. **Hallucination** — Le modèle affirme avec assurance des informations fausses ou inventées. Inhérent au fonctionnement probabiliste ; le RAG (grounding sur documents) est le principal contre-feu.
59. **Hyperscaler** — Cloud planétaire (AWS, Azure, GCP, Oracle) vendant API + GPU + puces maison.
60. **Inference** — Phase d'utilisation d'un modèle entraîné (vs entraînement) ; ce que paie 99 % des usages.
61. **Inference scaling / test-time compute** — Nouveau paradigme (2024-) : améliorer les réponses en dépensant plus de calcul **pendant** l'inférence (tokens de raisonnement) plutôt qu'en entraînant plus gros. Formalisé par Snell et al. (2024) ; produits : o1/o3, R1.
62. **Interprétabilité** — Comprendre ce que représentent les neurones (travaux Olah/Anthropic). Encore partielle, mais progresse (2024-2026).
63. **Jailbreak** — Technique pour contourner les garde-fous d'un modèle (« DAN », injections...). Course sans fin avec les équipes sûreté.
64. **KV cache** — Mémoire des états intermédiaires de la conversation ; croît avec le contexte (peut OOM).
65. **KV-cache** — Mémoire cache des clés/valeurs d'attention des tokens déjà traités, qui évite de tout recalculer à chaque token généré. Son volume (batch × contexte × couches) est souvent le vrai dimensionnant VRAM du serving.
66. **Lab** — Organisation qui entraîne les modèles (OpenAI, Anthropic, DeepSeek…) vs **provider** qui les sert.
67. **Lethal trifecta** — (Simon Willison) Les 3 conditions qui rendent un agent dangereux : données de valeur + contenu externe non fiable + capacité d'exfiltration.
68. **Load balancing** — Répartition des requêtes entre plusieurs endpoints/clés (round-robin, moindre charge…).
69. **LoRA** — Low-Rank Adaptation : méthode de fine-tuning qui n'entraîne que de petites matrices de rang faible injectées dans le modèle — 10 000× moins de paramètres à entraîner qu'un fine-tuning complet. Le standard du fine-tuning « pauvre ».
70. **Loss (fonction de perte)** — Mesure de l'erreur du modèle pendant l'entraînement (ex. : entropie croisée sur le mot suivant). Les lois de scaling prédisent la loss, pas directement l'utilité.
71. **LPU** — « Language Processing Unit » : puce d'inférence de Groq, SRAM au lieu de HBM, 500–800 tok/s.
72. **MLX** — Framework Apple pour l'IA sur puces M ; l'équivalent CUDA du Mac (plus limité).
73. **MoE (Mixture of Experts)** — Architecture où chaque token n'active que quelques sous-réseaux « experts » parmi beaucoup (ex. : DeepSeek-V3 : 671B totaux, 37B actifs). Capacité d'un géant, coût d'un moyen.
74. **Multimodal** — Modèle traitant plusieurs modalités : texte, image, audio, vidéo.
75. **Multimodalité** — Un seul modèle traite plusieurs modalités : texte, image, audio, vidéo (GPT-4o, Gemini). La multimodalité « native » (un seul réseau) remplace les assemblages de modèles spécialisés.
76. **Neocloud** — Cloud GPU nouvelle génération (CoreWeave, Nebius, Lambda, Crusoe…) entre le marketplace et l'hyperscaler.
77. **NPU / TPU / Trainium** — Accélérateurs maison (Google, AWS…) : moins chers que NVIDIA à charge stable.
78. **Offload** — Déporter une partie du modèle sur CPU/RAM quand la VRAM est insuffisante (lent mais fonctionnel).
79. **Open-weight** — Poids publiés, réutilisables ; ≠ « open source » (code/données d'entraînement rarement inclus).
80. **Over-training (sur-entraînement)** — Entraîner volontairement au-delà de l'optimum Chinchilla (ex. : 1 875 tokens/paramètre pour Llama 3 8B) parce que l'inférence — coût récurrent — domine l'économie totale.
81. **Overthinking** — Dégradation du raisonnement quand le budget de thinking est trop élevé : boucles de vérification stériles (études 2025-2026).
82. **PagedAttention** — Pagination du KV cache (vLLM) : +2 à 24× de débit vs moteurs naïfs.
83. **Perplexité** — Exp(loss) : mesure classique de la qualité d'un modèle de langage (plus bas = mieux). De moins en moins utilisée seule face aux benchmarks de tâches.
84. **PPO** — Proximal Policy Optimization (Schulman et al., 2017) : algorithme de RL utilisé dans le RLHF classique.
85. **Prefill** — Phase de traitement du prompt d'un coup ; limitée par le calcul (FLOPS).
86. **Prompt / prompting** — Texte d'entrée donné au modèle ; l'art de le formuler (instructions, exemples, contexte) conditionne fortement la qualité. Le « prompt engineering » est devenu une discipline, partiellement automatisée (DSPy...).
87. **Prompt caching** — Réutilisation tarifée du préfixe déjà traité (-90 % en lecture chez Anthropic/OpenAI) ; arme n°1 du RAG.
88. **Prompt injection** — Instruction malveillante glissée dans une donnée lue par le modèle (page web, e-mail, commentaire). **Indirecte** = la variante dangereuse en prod. Vecteur n°1 (34 % des incidents, OWASP 2026).
89. **Pré-entraînement** — Phase 1 : apprendre sur un corpus massif (le coûteux). Phase 2 : post-training (SFT, RLHF) — le rentable.
90. **PTU** — « Provisioned Throughput Units » (Azure) : débit réservé avec SLA de latence, vs trafic partagé.
91. **Q4_K_M / Q8_0** — Niveaux de quantification GGUF : Q4 ≈ 0,5 o/param (92–98 % qualité), Q8 ≈ 1 o/param (~100 %).
92. **Quantification** — Réduire la précision des poids (FP16 → INT8/INT4) pour diviser la VRAM par 2-8 avec une perte de qualité maîtrisée. Méthodes : GPTQ, AWQ, GGUF, FP8. Indispensable au self-hosting.
93. **RAG (Retrieval-Augmented Generation)** — Architecture : on récupère les passages pertinents d'un corpus (via embeddings) et on les injecte dans le prompt du générateur. Réduit les hallucinations et ancre les réponses dans des sources vérifiables. **Le cœur du projet de Zelef.**
94. **Rate limit** — Quotas du provider (requêtes/min, tokens/min-jour) ; le 429 est le déclencheur n°1 des fallbacks.
95. **Reasoning (modèle de)** — Modèle qui génère des étapes de raisonnement intermédiaires avant de répondre (chaîne de pensée). Plus lent, meilleur en logique/code/maths.
96. **Red-teaming** — Faire attaquer le système par des testeurs (humains ou IA) pour trouver les failles avant les attaquants. Standard industriel.
97. **Rerank** — Reclassement des résultats de recherche par un modèle plus précis (cross-encoder). Étape clé d'un bon RAG.
98. **Reranker** — Second modèle qui re-classe les résultats de recherche ; +5–15 pts de précision RAG typiques.
