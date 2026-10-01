---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-15
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Huawei", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["amd", "awq", "embedding", "embeddings", "fine-tuning", "fp8", "gguf", "gptq", "gpu", "int4", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1508, 1595]
sha256: 471408846553aeae5669d1f224501150da6e776f900ac80e0c75314fab2adf27
---

# Encyclopédie des modèles IA — Volume 3

| Critère | vLLM | Ollama | llama.cpp |
|---|---|---|---|
| Cas d'usage | Serveur production, API OpenAI-compatible | Local simple, « ça marche » | Local expert, CPU, embarqué |
| Batching | Continu + PagedAttention (le meilleur débit) | Basique | Basique (server) |
| Débit (req/s) | 2–10× les autres à hardware égal | Moyen | Faible à moyen |
| Formats | FP16, FP8, GPTQ, AWQ, GGUF (partiel) | GGUF (+ safetensors récent) | GGUF (+ tous les quants exotiques) |
| GPU supportés | NVIDIA (CUDA), ROCm partiel | NVIDIA, AMD, Apple Silicon, CPU | Tout (CPU excellent, tous GPU) |
| Partage de préfixes | Oui (prompt système en cache une fois) | Non documenté | Non |
| Complexité | Docker, Python, config | Une commande | Compilation, flags |
| Speculative decoding | Oui (MTP, EAGLE) | Non | Oui (partiel) |

Recommandation :

- **Serveur RAG (même mono-utilisateur intensif)** : vLLM. Le partage de préfixes (ton prompt système RAG identique à chaque appel) et le batching continu justifient à eux seuls le choix.
- **Test rapide / laptop** : Ollama. `ollama run qwen3.8:27b` et c'est parti — imbattable pour prototyper.
- **CPU-only / hardware exotique / contrôle fin** : llama.cpp. Le seul qui exploite vraiment bien un CPU moderne ou un Mac.

Note 2026 : SGLang est l'alternative montante à vLLM (support day-0 des nouveautés : Qwen3.5, FP8, speculative decoding — la recherche Granite cite un support day-0 SGLang). À évaluer si tu montes un serveur sérieux.

## 107. Quelle carte pour quoi : tableau de correspondance

| Carte (VRAM) | Ordre de prix (2026, neuf/occasion, indicatif) | Ce que tu sers en local | RAG associé |
|---|---|---|---|
| RTX 4060 Ti 16 Go | ~500 € | 7B–14B Q4 | Générateur correct, embedding+rerank sur CPU |
| RTX 3090 24 Go (occasion) | ~700–900 € | 27B Q4 (18 Go) + cache | **Sweet spot** : Qwen3.8-27B + reranker 0,6B |
| RTX 4090 24 Go | ~1800–2000 € | 27B Q4/Q5 + cache large | Idem, plus rapide ; fine-tuning LoRA 7B |
| RTX 6000 Ada 48 Go | ~7000 €+ | 70B Q4 (40 Go) + cache | Serveur mono-GPU sérieux |
| 2× RTX 4090 48 Go | ~4000 € | 70B Q4 en tensor-parallel | vLLM multi-GPU à la maison |
| H100 80 Go (cloud/occasion pro) | Location ~2–3 $/h | 70B FP8 + gros batch | Production |
| Mac Studio M3/M4 Ultra (128–512 Go unifiés) | ~4000–8000 € | 70B Q4 en RAM unifiée (lent mais ça passe) | Option « silencieuse » : gros modèle, faible débit |
| DGX Spark GB10 (128 Go unifiés) | ~3000–4000 $ | 70B Q4 / INT4 (ex : Ling 72 Go en INT4) | Le « mini-serveur » NVIDIA 2026 |

Lecture : pour **ton** RAG (générateur 27B + embedding + reranker), une **24 Go** (3090 d'occasion ou 4090) suffit. Le 70B local est un luxe, pas un besoin — l'écart de qualité sur de la synthèse de chunks est faible devant l'écart de coût.

## 108. Le cas particulier : CPU, Mac et DGX Spark

Trois options « sans gros GPU NVIDIA » :

- **CPU moderne (llama.cpp)** : un 7B Q4 tourne à ~10–20 tok/s sur un bon CPU desktop (AVX2/VNNI). Utilisable pour un RAG peu intensif (quelques questions/heure). Au-delà, trop lent.
- **Mac Apple Silicon (MLX / Ollama)** : la mémoire unifiée permet de charger de gros modèles (70B Q4 dans 128 Go unifiés), avec un débit correct (~10–30 tok/s). Les portages MLX existent pour Jina (embeddings v5, reranker v3.5). Bon choix « salon », mauvais choix « serveur ».
- **DGX Spark (GB10, 128 Go unifiés)** : la machine 2026 pour le local sérieux — la recherche Ling la cite faisant tourner un modèle en INT4 (72 Go). C'est le chaînon entre le PC et le serveur : 128 Go unifiés, form factor desktop.

Règle : si ton débit cible est < 1 question/minute, le CPU/Mac suffit. Au-delà, il faut un GPU NVIDIA + vLLM.

## 109. Chunking : taille fixe vs sections

Le chunking découpe tes documents en morceaux indexables. Deux philosophies :

**Taille fixe** (ex : 1000 tokens, overlap 200) :

- ✅ Simple, prévisible, compatible avec tous les modèles (contexte 512+).
- ❌ Coupe au milieu des phrases, des tableaux, des procédures. Un tableau de specs coupé en deux = deux chunks inutilisables.
- Overlap : l'assurance anti-coupe — 10–20 % de recouvrement pour que l'info à cheval sur la frontière survive.

**Par sections** (titres, paragraphes, unités logiques) :

- ✅ Respecte la structure : une procédure reste entière, un tableau reste entier.
- ❌ Tailles irrégulières : une section de 50 tokens (bruit) ou de 5000 tokens (trop gros pour l'embedding).
- Hybride recommandé : découper par sections, puis **re-fusionner les petites** (< 300 tokens) et **subdiviser les grosses** (> 2000 tokens) par taille fixe.

Pour tes corpus (guides en markdown avec titres numérotés), le découpage par sections est naturel : les `## N. Titre` sont des frontières de chunks toutes faites. C'est exactement ce que fait la section suivante.

## 110. L'exemple des guides Huawei (ton corpus réel)

Tes guides Huawei font 2500–3500 lignes en markdown, sections numérotées `## N. Titre`. Stratégie de chunking appliquée :

1. **Frontière primaire : les sections `##`.** Chaque section = un chunk candidat. Une section « 12. Configuration du PoE » contient tout le raisonnement PoE au même endroit — l'embedding la comprend comme une unité.
2. **Garde-fous de taille** : sections < 200 tokens → fusionner avec la suivante (évite les chunks « titre seul ») ; sections > 2500 tokens → subdiviser sur les `###` ou par paragraphes.
3. **Métadonnées par chunk** (section 112) : `guide`, `section_n`, `section_titre`, `famille_equipement` (AP361, S310…), `type_contenu` (procédure, spec, tableau, code).
4. **Cas spécial : les tableaux.** Un tableau de 40 lignes (ex : correspondances IOS↔VRP, 50 entrées) ne doit jamais être coupé : c'est une unité atomique. Si un tableau dépasse la taille max, le dupliquer en entier dans un chunk dédié plutôt que le scinder.
5. **Les quiz et glossaires** : les chunker séparément avec le tag `type: quiz` / `type: glossaire` — ce sont d'excellents chunks pour les questions de définition (« c'est quoi le THDi ? »).

Résultat attendu : ~150–250 chunks par guide de 3000 lignes, taille médiane 800–1200 tokens. Index total pour 11 guides Huawei : ~2000 chunks ≈ 2–2,5M tokens ≈ $0,05 en embedding API.

## 111. Chunking des CLI Huawei (12 620 commandes)

Ton corpus CLI nettoyé (huawei_cli_ref_*_clean.tar) a déjà la bonne granularité : **1 fichier .md par commande** (Function, Format, Parameters en tableau, Views, Default Level, Usage Guidelines, Example). C'est le chunking idéal — ne le casse pas :

- **1 commande = 1 chunk.** La question « comment configurer le PoE sur un port ? » matche le chunk `poe enable` exactement.
- **Métadonnées critiques** : `commande` (nom exact), `famille` (AR/Switch/AP), `view` (system/port/vlan…), `niveau` (défaut). La `view` est un filtre puissant : « en vue port » élimine 90 % des faux positifs.
- **L'Example est de l'or** : les exemples de commandes contiennent la syntaxe réelle — l'embedding les adore (le texte d'exemple ressemble aux questions des utilisateurs).
- **Ne pas chunker les tableaux de paramètres** : le tableau Parameters d'une commande est atomique (section 110, règle 4).
- **Volume** : 12 620 chunks ≈ 35 Mo de markdown. En embedding à 1024 dims : 12 620 × 1024 × 4 octets ≈ 52 Mo d'index — ridicule. Le CLI complet tient dans un index de poche.

Leçon générale : quand la source est déjà structurée (1 commande = 1 page), **le meilleur chunking est de ne pas chunker** — juste indexer l'unité naturelle avec de bonnes métadonnées.

## 112. Métadonnées : le multiplicateur silencieux

