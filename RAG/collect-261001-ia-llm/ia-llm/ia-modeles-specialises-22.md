---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-22
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cohere", "Google", "SGLang", "vLLM"]
dates: ["2026-09-24"]
keywords: ["apache", "attention", "awq", "benchmark", "benchmarks", "cohere", "datacenter", "embedding", "embeddings", "fine-tuning", "fp8", "gemini"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [2207, 2308]
sha256: 37044b045b67d16a4c61b59e8f304c4fb4dfb8f595ca7ea8edb62b78afdc45bd
---

# Encyclopédie des modèles IA — Volume 3

**Niveau 1 — légendage à l'indexation (recommandé).**
À l'ingestion, chaque image passe dans un VLM léger local (Qwen3.5-2B, 2B paramètres, OCRBench 84,5+, Apache 2.0) avec le prompt : « Décris précisément cette image technique en 150 mots : équipements, étiquettes, valeurs, connexions. » La description devient un chunk texte avec métadonnées `type: image_description`, `image_path`, `page`. Coût : quelques secondes/image en local, une fois. Aucun changement au runtime.

**Niveau 2 — chunks mixtes au runtime.**
Le VLM indexé côté corpus ne suffit pas quand la question porte sur un détail visuel (« où est le port console sur cette face avant ? »). Option : stocker aussi un embedding **multimodal** de l'image (Gemini Embedding 2, Cohere Embed v4 — texte+image dans le même espace) et chercher en texte→image. Plus cher (API), à réserver aux corpus très visuels.

**Niveau 3 — question avec image jointe.**
L'utilisateur envoie une photo (« c'est quoi ce voyant rouge ? ») + question. Le générateur doit être un VLM (Gemini 3.1 Pro, Qwen3-VL, Gemma 4). C'est le niveau « support terrain » : puissant, mais chaque requête coûte une inférence VLM.

Pour toi : le niveau 1 couvre 90 % du besoin (tes schémas deviennent cherchables en texte). Le niveau 3 est le « wow » pour le terrain — à envisager quand le RAG texte sera solide.

---

*Fin du volume 3 — 139 sections, rédigé le 27 septembre 2026.*
*Bon RAG, Zelef.*

## 140. Fiches réflexes : les chiffres à retenir par cœur

Les ordres de grandeur qui doivent devenir des réflexes :

**Tailles.**
- 1B paramètres ≈ 2 Go en FP16, ≈ 1 Go en INT8, ≈ 0,5–0,6 Go en INT4/GGUF Q4.
- 7B ≈ 14 / 7 / 4 Go · 27B ≈ 54 / 27 / 16–18 Go · 70B ≈ 140 / 70 / 40 Go.

**KV cache (GQA-8, head_dim 128, FP16).**
- 7B (32 couches) : 128 Kio/token → 4 Go à 32K, 16 Go à 128K.
- 70B (80 couches) : 320 Kio/token → 10 Go à 32K, 40 Go à 128K.
- Formule : 2 × couches × têtes_KV × dim_tête × octets × tokens.

**Prix API (ordres de grandeur 2026).**
- Embedding : $0,02–0,20/M tokens — négligeable à ton échelle.
- Générateur éco : ~$0,15–0,45/M in — ~$0,005/question RAG.
- Générateur flagship : $4–5/M in — ~$0,10/question RAG.
- Prompt caching : −90 % sur le préfixe. Batch : −50 %.

**Benchmarks.**
- Embedding FR : regarder MMTEB Retrieval / MIRACL, pas MTEB English.
- Rerank : nDCG@10 sur BEIR ; 63+ = excellent en 2026 (0,6B !).
- Écart < 3 pts sur le même benchmark = bruit.

**RAG.**
- Rappel@5 ≥ 90 % sur tes questions = retrieval sain.
- 10–20 chunks de ~1200 tokens = 12–24K tokens de contexte par requête.
- Métadonnées + filtre famille : +10 à +20 pts de précision, gratuit.

**Hardware.**
- 24 Go : 27B Q4 + cache pour ~20K tokens, batch=1.
- Vise 70–80 % de VRAM max, jamais 100 %.
- Le batch multiplie le cache, pas les poids.

## 141. Comparer deux modèles en 30 minutes : la méthode

Quand tu hésites entre deux modèles (deux embeddings, deux générateurs), ne lis pas 10 articles — fais ça :

**Minutes 0–10 : les fiches.**
- Licence (Apache 2.0 ? NC ? fermé ?) — éliminatoire selon ton usage.
- Contexte ≥ tes besoins (chunks + prompt) ?
- Prix ou VRAM : dans ton budget ?
- Si un critère est éliminatoire, stop. Sinon continue.

**Minutes 10–25 : le test réel.**
- 10 questions représentatives de TON usage (pas des benchmarks).
- Même protocole pour les deux (même chunks, même prompt, même température).
- Note en aveugle si possible (mélange les réponses, juge sans savoir lequel est lequel).

**Minutes 25–30 : la décision.**
- Écart net et reproductible (≥ 5 pts ou préférence claire sur 8/10) → le gagnant.
- Écart flou → prends le moins cher / le plus simple / la meilleure licence.
- Documente : 5 lignes dans un fichier (date, modèles, protocole, résultat). Dans 6 mois tu ne te souviendras plus pourquoi tu as choisi.

Ce qui ne marche pas : décider sur un leaderboard seul, sur une démo marketing, ou sur « tout le monde dit que ». Ton usage (français technique, 25 000 lignes de guides) n'est le benchmark de personne.

## 142. Le stack 100 % open-source et local : la recette complète

Si tu veux un RAG sans aucun appel API, sans données qui sortent, en licences propres (Apache 2.0 / MIT) :

| Brique | Modèle | Licence | VRAM |
|---|---|---|---|
| Embedding | Qwen3-Embedding-8B (ou BGE-M3 si petit GPU) | Apache 2.0 | ~16 Go FP16 / ~4 Go INT8 |
| Reranker | Qwen3-Reranker-0.6B (ou bge-reranker-v2-m3) | Apache 2.0 | ~1,5 Go |
| Générateur | Qwen3.8-27B Q4_K_M | Apache 2.0 | ~18 Go |
| Index | Qdrant (ou pgvector) | Apache 2.0 | RAM |
| Serveur | vLLM (ou SGLang) | Apache 2.0 | — |
| VLM légendage | Qwen3.5-2B | Apache 2.0 | ~4–5 Go |
| **Total** | | | **~40 Go (2× 24 Go ou 1× 48 Go)** |

Version « une seule 24 Go » : embedding BGE-M3 (1,2 Go) + reranker 0,6B (1,5 Go) + générateur 27B Q4 (18 Go) ≈ 21 Go + cache. Le VLM de légendage tourne **à l'indexation uniquement** (pas en même temps que le serveur) — pas besoin de le compter dans le serving.

Coût total : la carte (amortie) + l'électricité (~0,50 €/mois pour 1000 questions). Données : 100 % chez toi. C'est le stack « souverain » — et en 2026, il est compétitif avec les API sur la qualité du français technique.

## 143. Antisèche finale : une page, tout le volume

**Modèles (Partie A).** Embeddings : Qwen3-Embedding-8B (Apache 2.0, self-host), Jina v5 (léger, NC), Gemini Embedding 2 (multimodal, API). Rerankers : Jina v3.5 (listwise, 131K, NC), Qwen3-Reranker (Apache 2.0), Cohere 4.0 (API). Vision : Qwen3-VL-235B (Apache 2.0), Nemotron 3 Nano Omni (30B/3B), Qwen3.5-2B (edge). Vidéo : Seedance 2.0 (Elo 1269), Kling 3.0, Runway Aleph (édition) ; Sora 2 = API fermée le 24/09/2026. Image : Nano Banana 2 ($0,067), Pro ($0,134, typo parfaite), Lite ($0,0336).

**Technique (Partie B).** KV cache = 2×couches×têtes_KV×dim×octets×tokens ; le poste VRAM n°1 en long contexte. GQA standard ; MLA = compression ; hybrides 2026 = attention globale sur une minorité de couches. MoE : total en VRAM, actifs en calcul ; bon à gros batch, mauvais en solo. Quantization : GPTQ/AWQ (PTQ), GGUF Q4_K_M (local), FP8/NVFP4 (datacenter, natif 2026). LoRA = A×B sur poids figés ; QLoRA = sur base 4-bit. RAG > fine-tuning pour les faits ; LoRA pour le comportement ; fine-tuner l'embedding = meilleur ROI. Chunking par sections, tableaux atomiques, métadonnées + filtre famille. HyDE pour les questions vagues. Prompt caching (−90 %) + batch (−50 %) = ÷5 à ÷10 sur la facture. Fenêtre 1M = outil d'analyse, pas architecture RAG. Mesure (rappel@5 sur TES questions FR) avant d'optimiser — toujours.

*Fin du volume 3.*

## 144. Les licences : le guide de survie juridique

En 2026, « open weights » ne veut pas dire « open source ». Le tableau qui évite les ennuis :

