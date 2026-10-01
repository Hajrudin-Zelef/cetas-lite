---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-65
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Cerebras", "DeepSeek", "Fireworks AI", "Google", "Groq", "OpenAI", "vLLM"]
dates: ["2022-11-30"]
keywords: ["agents", "attention", "chatgpt", "compute", "deepseek", "dpo", "gpu", "inference", "rlhf", "speculative decoding", "tool calling", "training"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5443, 5497]
sha256: 59a7654b9c49cc17e59c73929c4a7d732eb2f7909f4026d3a00cefa910dc3e66
---

# IA — Le grand dossier

99. **RLHF** — Reinforcement Learning from Human Feedback (InstructGPT, OpenAI, 2022) : SFT → modèle de récompense sur préférences humaines → optimisation PPO. A transformé les prédicteurs de texte en assistants (ChatGPT).
100. **RLHF / RLAIF / DPO** — Méthodes d'alignement : RLHF = renforcement avec feedback *humain* ; RLAIF = feedback d'une *IA* (Constitutional AI) ; DPO = optimisation directe des préférences, sans RL.
101. **RLVR** — Reinforcement Learning with Verifiable Rewards (2025-) : le RL du raisonnement avec des récompenses **vérifiables** (tests, démonstrateurs) plutôt que des jugements humains — réponse au mur des données humaines et à l'« overthinking ».
102. **RPM / TPM / RPD** — Requêtes/tokens par minute, requêtes par jour : les trois compteurs des quotas.
103. **Scaling laws (lois d'échelle)** — Relations en loi de puissance entre compute/données/taille et performance (Kaplan 2020, Chinchilla 2022, lois du test-time 2024). Le cadre quantitatif de toute la partie 4.
104. **Self-attention** — Variante de l'attention où chaque position d'une séquence s'auto-pondère par rapport aux autres positions de la même séquence. Permet le parallélisme total du Transformer.
105. **Serverless (inférence)** — Le provider alloue le GPU à la requête ; facturation au token ou à la seconde ; cold starts possibles.
106. **Souveraineté** — Maîtrise du lieu et du contrôle des données/traitements (UE, FR) ; argument d'achat public.
107. **Speculative decoding** — Un petit modèle « brouillon » accélère le grand ; +30–100 % de débit (vLLM, Fireworks).
108. **Spend log** — Journal de chaque appel (tokens, coût, latence) ; base du pilotage budgétaire (LiteLLM/Postgres).
109. **Streaming (SSE)** — Envoi des tokens au fil de l'eau ; divise la latence perçue.
110. **Structured output** — Réponse contrainte à un schéma (JSON) ; indispensable pour l'extraction/classification fiable.
111. **System card** — Document publié par un labo décrivant un modèle : capacités, limites, évaluations de sécurité. La doc « honnête » — à lire avant d'adopter.
112. **Temperature** — Paramètre (0 à ~2) contrôlant le caractère aléatoire de la génération : 0 = déterministe (factuel, RAG), élevé = créatif/divers.
113. **Token** — Unité de texte traitée par le modèle (~0,75 mot anglais, ~4 caractères). Unité de facturation des API et de mesure des contextes.
114. **Tool calling** — Capacité du modèle à appeler des fonctions/outils (agents) ; nécessite un parser adapté (vLLM).
115. **Transformer** — Architecture (Vaswani et al., 2017) fondée sur la self-attention, sans récurrence. Devenue l'architecture universelle (texte, image, audio, vidéo, protéines).
116. **TTFT** — « Time To First Token » : délai avant le premier token ; la métrique de l'interactif (Groq <1 s).
117. **Vector DB** — Base de données optimisée pour la recherche de similarité entre vecteurs (pgvector, Qdrant, Chroma, Milvus, Weaviate). Cœur du retrieval sémantique.
118. **vLLM** — Serveur d'inférence open-source (UC Berkeley) : PagedAttention + continuous batching, API compatible OpenAI. Le standard du serving auto-hébergé 2024-2026.
119. **VRAM** — Mémoire de la carte GPU ; le facteur n°1 du dimensionnement local.
120. **WSE** — « Wafer-Scale Engine » : puce géante de Cerebras ; 1 800–3 000 tok/s record.
121. **Zero-retention** — Engagement contractuel : le provider ne conserve ni n'exploite tes prompts ; exiger l'écrit.
122. **Émergence** — Apparition brutale d'une capacité au-delà d'une échelle (few-shot de GPT-3). Réelle mais débattue (artefact de mesure possible).

---

# Quiz général — 30 questions corrigées

Les trois quiz des parties, réunis. 10 questions par partie.

### Quiz — Partie 1 : Histoire, révolutions, scaling (10 questions)

**Q1.** Quelle est la performance d'AlexNet à ImageNet 2012, et en quoi l'écart avec le second est-il historique ?

> **Corrigé.** **15,3 % d'erreur top-5** contre 26,2 % pour le second (Krizhevsky, Sutskever, Hinton — « ImageNet Classification with Deep Convolutional Neural Networks », NeurIPS 2012). L'écart est historique car le deep learning **divise quasiment l'erreur par deux** d'un coup, là où les progrès se mesuraient en dixièmes de point — c'est ce choc qui déclenche la décennie du deep learning. Ingrédients : ReLU, dropout, 2 GPU en parallèle, augmentation de données. (Section 1.2.)

**Q2.** Citez les 8 auteurs du papier « Attention Is All You Need » (2017) dans l'ordre, et expliquez en une phrase l'innovation centrale du Transformer par rapport aux RNN/LSTM.

> **Corrigé.** **Ashish Vaswani, Noam Shazeer, Niki Parmar, Jakob Uszkoreit, Llion Jones, Aidan N. Gomez, Łukasz Kaiser, Illia Polosukhin** (arXiv:1706.03762, NeurIPS 2017). Innovation : **supprimer entièrement la récurrence** et la remplacer par la **self-attention** — chaque mot traite tous les autres en une opération parallélisable (Attention(Q,K,V) = softmax(QK^T/√d_k)V), d'où parallélisme total sur GPU et dépendances longue distance. (Section 1.5.)

**Q3.** Quelle est la différence opérationnelle entre les lois de scaling de Kaplan (2020) et celles de Chinchilla (2022) ? Donnez la règle pratique issue de Chinchilla.

> **Corrigé.** Kaplan (OpenAI, 2020) : à budget compute fixé, **privilégier les paramètres** (d'où GPT-3 : 175B paramètres / ~300B tokens). Chinchilla (Hoffmann et al., DeepMind, 2022) corrige un biais méthodologique et montre que **paramètres et données doivent scaler à parts égales** (N_opt ∝ C^0,5, D_opt ∝ C^0,5). Règle pratique : **D ≈ 20 × N** (20 tokens par paramètre). Démonstration : Chinchilla (70B / 1 400B tokens) bat Gopher (280B / 300B tokens) à compute égal. (Section 4.1.)

**Q4.** Décrivez les 3 étapes du RLHF (InstructGPT, 2022) et le résultat « choc » du papier concernant la taille des modèles.

> **Corrigé.** (1) **SFT** : affinage supervisé sur des réponses écrites par des humains. (2) **Modèle de récompense** : entraîné sur des classements par paires de réponses par des annotateurs. (3) **PPO** : optimisation par renforcement contre le modèle de récompense, avec pénalité KL anti-reward-hacking. Résultat choc (Ouyang et al., 2022) : un **InstructGPT de 1,3B paramètres est préféré par les humains au GPT-3 brut de 175B** (~100× plus gros) — l'alignement bat la taille. ChatGPT (30/11/2022) en est le produit direct. (Section 1.9.)

**Q5.** Qu'est-ce que le « test-time compute » (inference scaling), quel modèle l'a popularisé en produit en septembre 2024, et quel modèle ouvert l'a répliqué en janvier 2025 ?

> **Corrigé.** Le test-time compute consiste à dépenser du calcul **pendant la réponse** (longue chaîne de pensée interne, auto-vérification) plutôt qu'en entraînant plus gros ; la performance suit des lois d'échelle à l'inférence (Snell et al., 2024). Popularisé en produit par **OpenAI o1** (ex-« Strawberry », septembre 2024 : 83 % à l'AIME 2024 vs 13 % pour GPT-4o). Répliqué en ouvert par **DeepSeek-R1** (20 janvier 2025, poids MIT, niveau o1 ; R1-Zero montre l'émergence du raisonnement par pur RL). (Sections 4.5, 1.11-1.12.)

**Q6.** Citez 3 arguments du camp « le scaling du training ralentit » et leurs auteurs/sources, puis 2 arguments du camp « le scaling continue » et leurs auteurs/sources.

