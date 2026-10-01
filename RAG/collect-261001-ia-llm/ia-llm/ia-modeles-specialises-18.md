---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-18
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Google", "Nvidia", "Oracle", "vLLM"]
dates: []
keywords: ["attention", "awq", "benchmark", "blackwell", "datacenter", "deepseek", "distillation", "embedding", "embeddings", "fine-tuning", "fp4", "fp8"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1834, 1969]
sha256: dd578ae719752a908a61808076f5f75ead6ecf9327823551d6d224f3a97cbfe9
---

# Encyclopédie des modèles IA — Volume 3

**AWQ** — Activation-aware Weight Quantization : méthode PTQ qui protège les poids « saillants » identifiés via les activations (MIT, 2023).

**Batch API** — traitement différé de requêtes en lot, facturé ~−50 % ; pour l'évaluation et le non-temps-réel.

**BEIR** — benchmark de retrieval zero-shot (18 datasets) ; le nDCG@10 y est la métrique standard des rerankers.

**BF16** — Brain Float 16 : format 16-bit privilégié pour l'entraînement et l'inférence (meilleure dynamique que FP16).

**Chunking** — découpage des documents en morceaux indexables ; par taille fixe, par sections, ou hybride.

**Contexte (fenêtre)** — nombre max de tokens que le modèle traite en une fois ; 4K (petit) à 1M+ (2026).

**Cross-encoder** — modèle qui lit la question ET le passage ensemble pour scorer leur pertinence ; principe des rerankers.

**Dense (modèle)** — tous les paramètres sont actifs à chaque token ; opposé de MoE.

**Distillation** — entraîner un petit modèle à imiter un grand ; ex : MiMo Distill-9B.

**Embedding** — vecteur dense représentant le sens d'un texte/image ; brique de la recherche sémantique.

**EP (Expert Parallelism)** — répartition des experts MoE sur plusieurs GPU ; nécessite une interconnexion rapide (NVLink).

**Éviction (cache)** — jeter une partie du KV cache quand il déborde (StreamingLLM, H2O) ; dégrade le rappel du milieu.

**Fine-tuning** — continuer l'entraînement sur tes données ; complet (cher) ou PEFT/LoRA (léger).

**FP8** — format 8-bit (E4M3/E5M2) ; ÷2 vs FP16, quasi sans perte ; standard datacenter 2026, aussi pour le KV cache.

**GGUF** — format de modèle quantizé de llama.cpp ; le standard du local (Q4_K_M = défaut recommandé).

**GQA** — Grouped-Query Attention : les têtes de requêtes partagent des têtes KV par groupes ; le standard 2024–2026.

**GPTQ** — méthode PTQ par compensation d'erreur couche par couche (2022) ; la référence GPU historique.

**H2O** — Heavy-Hitter Oracle : politique d'éviction du KV cache gardant les tokens à forte attention cumulée.

**Hallucination** — le modèle invente des faits avec aplomb ; le RAG la réduit (chunks cités), le fine-tuning l'aggrave parfois.

**HyDE** — Hypothetical Document Embeddings : embedder une réponse hypothétique plutôt que la question, pour mieux matcher les chunks.

**Inférence** — phase d'utilisation du modèle (vs entraînement) ; c'est elle qu'on optimise (quantization, vLLM).

**INT4 / INT8** — quantization entière sur 4/8 bits ; ÷4/÷2 sur la taille, perte croissante sous INT4.

**KV cache** — stockage des vecteurs Key/Value des tokens passés, couche par couche ; croît linéairement avec le contexte.

**Latent (MLA)** — vecteur compressé stockant K/V dans le MLA de DeepSeek ; ~93 % de réduction du cache.

**Listwise (rerank)** — le reranker lit toute la liste de candidats d'un coup et l'ordonne ; le plus moderne (Jina v3/v3.5).

**LoRA** — Low-Rank Adaptation : fine-tuning via deux petites matrices A×B ajoutées aux poids figés ; 256× moins de paramètres.

**Lost in the middle** — les modèles performent moins bien sur l'info au milieu d'un long contexte ; d'où le tri des chunks.

**Matryoshka (MRL)** — embeddings à dimensionnalité emboîtée : tronquables (1024→256) sans réentraînement.

**MHA** — Multi-Head Attention : le mécanisme originel, 1 tête KV par tête de requête ; gourmand en cache.

**MIRACL** — benchmark de retrieval multilingue (18 langues) ; pertinent pour le français.

**MLX** — framework Apple Silicon pour le ML local ; portages Jina (embeddings, reranker).

**MMTEB** — version multilingue de MTEB ; la référence pour comparer des embeddings FR.

**MoE** — Mixture of Experts : N experts + un routeur ; capacité (total) découplée du coût (actifs).

**MQA** — Multi-Query Attention : une seule tête KV partagée ; cache ÷32, légère perte qualité.

**MTEB** — Massive Text Embedding Benchmark (41 tâches EN) ; ne regarder que la sous-partie Retrieval pour un RAG.

**MTP** — Multi-Token Prediction : le modèle prédit plusieurs tokens d'un coup ; sert au décodage spéculatif (Qwen4-Exp).

**nDCG@10** — métrique de classement : qualité de l'ordre des 10 premiers résultats, pondérée par la position.

**Needle-in-a-haystack** — test de rappel d'une info noyée dans un long contexte ; ne mesure pas le classement.

**NVFP4** — format 4-bit de NVIDIA pour Blackwell (micro-scaling par blocs) ; « quasi-identique au 8-bit ».

**Ollama** — runtime local simple (GGUF) ; idéal pour prototyper, pas pour servir en production.

**PagedAttention** — gestion du KV cache en pages (vLLM) ; zéro gaspillage, batching continu, partage de préfixes.

**PEFT** — Parameter-Efficient Fine-Tuning : n'entraîner qu'une fraction des paramètres (LoRA, QLoRA…).

**Prefill** — phase où le modèle traite tout le prompt d'un coup (avant de générer) ; coûteuse en long contexte.

**Prompt caching** — remise (−90 % typique) quand le préfixe du prompt est déjà en cache fournisseur.

**PTQ** — Post-Training Quantization : quantizer après entraînement, avec calibration (GPTQ, AWQ, GGUF).

**QAT** — Quantization-Aware Training : entraîner en simulant la quantization ; meilleure qualité (modèles natifs INT4/FP4 2026).

**QLoRA** — LoRA sur modèle de base quantizé 4-bit (NF4) ; fine-tune un 65B sur un GPU 48 Go.

**RAG** — Retrieval-Augmented Generation :检索 des chunks pertinents puis génération à partir d'eux ; l'architecture de ton projet.

**Rappel@k** — proportion de bonnes réponses présentes dans les k premiers résultats ; LA métrique d'un retrieval.

**Rerank** — second passage qui re-classe les candidats du retrieval avec un cross-encoder ; +5 à +15 pts nDCG typiques.

**RoPE** — Rotary Positional Encoding : l'encodage positionnel standard des LLM modernes.

**Routeur (MoE)** — petit réseau qui choisit les top-k experts par token ; risque de déséquilibre de charge.

**RTEB** — benchmark retrieval 2026, récent ; à suivre avec prudence (alerte split privé, section 6).

**SFT** — Supervised Fine-Tuning : fine-tuning sur paires instruction/réponse ; la base de l'alignement.

**Sliding window** — l'attention ne porte que sur une fenêtre glissante de tokens récents ; borne le KV cache.

**StreamingLLM** — éviction gardant les premiers tokens (attention sink) + fenêtre récente ; génération « infinie » en mémoire constante.

**SWE-bench** — benchmark de résolution de tickets GitHub ; 80 %+ en 2026 = saturation proche.

**SynthID** — watermark invisible de Google sur les images générées (Nano Banana).

**Temperature** — paramètre de créativité/échantillonnage ; basse (0,1–0,3) pour un RAG factuel.

**Token** — unité de texte du modèle (~0,75 mot en anglais, ~1–1,5 mot en français technique) ; unité de facturation.

**Top-k (experts)** — nombre d'experts sélectionnés par token dans un MoE (top-2 à top-16 en 2026).

**vLLM** — moteur d'inférence serveur (PagedAttention, batching continu) ; le choix production.

**ViDoRe** — benchmark de retrieval visuel sur documents ; pour les corpus PDF/images.

## 126. Quiz : 15 questions

**Q1.** Ton embedding actuel est `text-embedding-3-small`. Quel est son principal point faible face aux modèles 2026, et pourquoi n'est-ce pas urgent de migrer ?

**Q2.** Pourquoi Jina Embeddings v5 (licence CC BY-NC 4.0) est-il parfait pour ton RAG personnel mais problématique si ton RAG devient un outil d'entreprise ?

**Q3.** Calcule le KV cache par token (FP16) d'un modèle 32 couches, GQA 8 têtes KV, head_dim 128. Puis le cache total à 32K tokens.

**Q4.** Un Qwen3.6-35B-A3B : combien de paramètres faut-il charger en VRAM, et combien calculent réellement par token ? Quelle en est la conséquence pour le serving ?

