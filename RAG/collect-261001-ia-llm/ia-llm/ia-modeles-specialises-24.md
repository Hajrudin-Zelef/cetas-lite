---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-24
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Hugging Face", "Nvidia", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-27"]
keywords: ["apache", "attention", "awq", "benchmarks", "cohere", "deepseek", "embedding", "embeddings", "fine-tuning", "fp8", "gptq", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [2408, 2501]
sha256: 1afbf65ed6863ea1ec6070a6dc5cc9d9c3d32a6436177699beac41992ee5e9de
---

# Encyclopédie des modèles IA — Volume 3

**TP 4 — Activer/désactiver le rerank (30 min).**
Avec ton pipeline : retrieval top-50 → mesure le rappel@5 SANS rerank sur tes 50 questions, puis AVEC un reranker 0,6B local (jina-reranker-v3 ou Qwen3-Reranker-0.6B). Note le gain en points ET la latence ajoutée. Objectif : décider sur chiffres, pas sur principe, si le rerank reste dans ton pipeline (sections 22, 97).

**TP 5 — Chiffrer trois architectures (20 min).**
Sur papier : (a) tout API flagship, (b) tout API éco, (c) tout local sur ta carte. Pour 500 questions/mois, calcule avec les grilles de la section 121 + le prompt caching (section 122). Objectif : réaliser que l'écart entre (a) et (c) est d'un facteur 50 à 200 — et que le bon choix dépend de ton volume, pas d'une préférence abstraite.

*Fin du volume 3 — 147 sections, 27 septembre 2026.*

## 148. FAQ express : les questions qu'on pose toujours

**Faut-il un GPU pour un RAG ?** Non pour démarrer : embeddings + rerank + génération en API = 0 € de hardware, ~5–15 €/mois pour 1000 questions (section 102). Oui si tu veux du 100 % local / confidentiel : une 24 Go suffit pour un 27B Q4 (section 107).

**Embedding payant ou gratuit ?** À ton échelle, la question ne se pose pas en coût (vectoriser 2M tokens = $0,04–0,40, une fois). Elle se pose en qualité FR (mesure sur tes questions) et en licence (Apache 2.0 si usage pro futur).

**Un seul modèle peut-il tout faire (embedding + rerank + génération) ?** Non : ce sont trois architectures différentes (bi-encoder, cross-encoder, décodeur). Les modèles « unifiés » existent en recherche mais ne sont pas au niveau des spécialistes en 2026.

**Quelle taille de chunks ?** 800–1500 tokens par défaut, par sections (pas taille fixe aveugle), tableaux jamais coupés, 1 commande = 1 chunk pour les CLI (sections 109–111).

**Combien de chunks dans le prompt ?** 5 à 10 après rerank, triés par pertinence. Au-delà, le « lost in the middle » et le coût du cache punissent (sections 70, 114).

**Le fine-tuning remplace-t-il le RAG ?** Non. RAG = faits (changeants, citables). Fine-tuning = comportement (format, ton). Les deux se combinent (section 91).

**vLLM ou Ollama ?** Ollama pour prototyper en 5 minutes ; vLLM dès que tu sers (batching continu, PagedAttention, partage de préfixes) — 2 à 10× le débit à hardware égal (section 106).

**INT4 dégrade-t-il trop ?** Pour un RAG (synthèse factuelle) : non, la perte est de 2–6 % et un 27B INT4 reste excellent. Pour du code critique ou du raisonnement fin : préfère FP8/INT8 (section 86).

**MoE ou dense en local ?** Dense. Le MoE n'est rentable qu'à gros batch (serveur/API) — en solo il gaspille la VRAM (sections 74, 76).

**1M de contexte : je mets tout mon corpus dedans ?** Non : 10–100× plus cher par question que le retrieval, « lost in the middle », latence du prefill. Le 1M sert à l'analyse ponctuelle de gros dossiers (section 115).

**Comment savoir si mon RAG est bon ?** Rappel@5 ≥ 90 % sur 50–100 questions FR réelles et versionnées, rejouées à chaque changement (sections 11, 130). Le reste est du ressenti.

**Par quoi commencer demain matin ?** Checklist section 128 : corpus propre → chunking par sections → métadonnées → embedding actuel → mesure. N'optimise que ce que la mesure désigne.

*Fin du volume 3 — 148 sections, 27 septembre 2026.*

## 149. Pour aller plus loin : lectures recommandées

Par thème, en allant du plus accessible au plus technique. Toutes ces références sont des classiques établis ou des sources primaires citées dans les recherches — pas des découvertes du jour.

**Fondations (à lire en premier).**
- « Attention Is All You Need » (Vaswani et al., 2017) — le papier originel du transformer. Dense mais incontournable pour comprendre d'où vient le KV cache.
- Le blog de Hugging Face (articles NVIDIA Nemotron-3-Embed, Jina) — les annonces fournisseurs les mieux documentées du lot.
- Les model cards Hugging Face (Qwen3-Embedding-8B, nvidia/Nemotron-3-Embed-8B-BF16, jinaai/jina-reranker-v3.5) — toujours lire la carte avant d'adopter un modèle.

**KV cache et attention.**
- « GQA » (Ainslie et al., 2023) — le papier du Grouped-Query Attention, 10 pages, très lisible.
- Les papiers DeepSeek-V2/V3 (MLA) et les notes DeepSeek V4 (attention hybride) — pour comprendre où va l'industrie.
- « StreamingLLM » et « H2O » (2023) — les deux politiques d'éviction de référence.
- La doc vLLM (PagedAttention) et SGLang (RadixAttention) — le système autant que le modèle.

**Quantization et fine-tuning.**
- « GPTQ » (Frantar et al., 2022), « AWQ » (Lin et al., 2023), « QLoRA » (Dettmers et al., 2023), « LoRA » (Hu et al., 2021) — quatre papiers, toute la pratique du local.
- Les docs Unsloth et axolotl — les recettes de fine-tuning qui marchent en 2026.

**RAG.**
- « HyDE » (Gao et al., 2022) — 8 pages, l'idée en une soirée.
- La documentation Qdrant (filtrage, quantization d'index) et pgvector — le côté index.
- Les leaderboards MTEB/MMTEB (Hugging Face Spaces) — à consulter, jamais à suivre aveuglément (section 119).

**Benchmarks et prix.**
- Artificial Analysis (artificialanalysis.ai) — l'indépendant le plus complet : qualité, vitesse, prix par modèle.
- Les grilles officielles des fournisseurs (OpenAI, Google, Anthropic, DeepSeek, Voyage, Cohere) — bookmarkées et revérifiées avant tout budget.

**Veille.**
- Les 8 fichiers de `~/workspace/research_ia/` — ton snapshot du 27/09/2026, à reconduire tous les 6 mois : le paysage 2026 bouge vite (Sora en est la preuve).
- Les changelogs des moteurs (vLLM, SGLang, Ollama, llama.cpp) — c'est là qu'on voit les techniques devenir des produits.

*Fin du volume 3 — 149 sections, 27 septembre 2026. Bon RAG, Zelef.*

## 150. Note de l'auteur et versioning

Ce volume a été rédigé le 27 septembre 2026 comme le tome 3 d'une encyclopédie des modèles IA, à partir de 8 dossiers de recherche web (février → septembre 2026) et de la littérature technique établie. Il est volontairement daté : en IA, un document non daté est un document mensonger.

Ce qui vieillira le plus vite : les prix API (promos, cuts), les benchmarks (saturation), les statuts (Sora a fermé en 3 jours — tout peut fermer en 3 jours), les « leaders du moment ». Ce qui vieillira le moins vite : la formule du KV cache, le principe du MoE, l'arbitrage RAG vs fine-tuning, la méthode d'évaluation. La Partie B est un investissement ; la Partie A est un snapshot.

Si tu reconduis la veille dans 6 mois : refais le snapshot de la Partie A (les 6+6+10+10+3 modèles auront changé), garde la Partie B en l'amendant (les ordres de grandeur bougeront, les principes resteront). Et quand un nouveau modèle « révolutionnaire » sortira, applique-lui la checklist de la section 120 avant de t'enthousiasmer.

Les fichiers sources de la recherche vivent dans `~/workspace/research_ia/` ; ce guide dans `~/workspace/user/files/ia_modeles_specialises.md`. Bon courage pour le RAG — avec 25 000 lignes de guides propres, un chunking par sections et un reranker 0,6B, tu pars déjà avec un meilleur corpus que 90 % des RAG en production.

*Fin.*

---

## Colophon

- **Titre :** Encyclopédie des modèles IA — Volume 3 : Modèles spécialisés + transversal technique
- **Rédaction :** 27 septembre 2026 (Rédacteur 3/3)
- **Sections :** 150 · **Glossaire :** 64 termes · **Quiz :** 15 questions + réponses
- **Sources :** `~/workspace/research_ia/` (8 dossiers, snapshot fév. → sept. 2026)
- **Langue :** français · **Ton :** direct et dense
- **Fichier :** `~/workspace/user/files/ia_modeles_specialises.md`

*Document versionné : toute mise à jour ultérieure devra porter une nouvelle date et un nouveau snapshot.*
