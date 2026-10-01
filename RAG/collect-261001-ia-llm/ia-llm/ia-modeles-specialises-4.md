---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-4
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "Meta", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Xiaomi", "vLLM"]
dates: []
keywords: ["agents", "agi", "apache", "attention", "benchmark", "benchmarks", "blackwell", "embedding", "gemini", "gguf", "kimi", "license"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [347, 472]
sha256: 14a10815140754c516ea8689e94d039c1fc7d43179cbad37657a333aeb8278e7
---

# Encyclopédie des modèles IA — Volume 3

Pour ton RAG : la vision sert à indexer les **documents non textuels** — schémas unifilaires d'onduleurs, captures de CLI, topologies eNSP, pages de datasheets scannées. Un VLM décrit l'image en texte, et ce texte rejoint ton index.

## 24. NVIDIA Nemotron 3 Nano Omni

Sorti le 28 avril 2026. Le pari « perception pour agents » de NVIDIA.

| Champ | Valeur |
|---|---|
| Architecture | **30B totaux, ~3B actifs par token** (MoE hybride : couches Mamba + transformer) |
| Modalités | Texte, image, audio, vidéo — unifiés dans une seule boucle de raisonnement |
| Contexte | 256K tokens (selon presse) |
| Positionnement | Couche de perception pour agents : computer use, screen recordings 1920×1080 en temps réel |
| Benchmarks (fournisseur) | Tête de 6 classements (document intelligence, vidéo, audio) dont OCRBenchV2 et VoiceBench ; 9,2× débit système sur vidéo vs autres modèles omni ouverts (comparateur non nommé — non reproduit) |
| Benchmark (indépendant, partiel) | anton-abyzov, v1 indicative, 29 tâches : 3–4× plus rapide qu'Opus 4.7 sur perception, latence médiane image 0,6 s ; audio 74 %, vidéo 75 % |
| Licence | Poids ouverts, **NVIDIA Open Model License** (+ datasets + recettes d'entraînement) |
| Déploiement | Hugging Face, NVIDIA NIM, Ollama, vLLM, SGLang, Jetson (edge), 25+ plateformes |

Le ratio 30B/3B est l'illustration parfaite du MoE (section 71) : tu charges 30B en VRAM mais tu ne calcules que 3B par token. Le mélange Mamba + transformer est notable : Mamba (SSM) traite les longues séquences vidéo en temps linéaire là où l'attention est quadratique.

## 25. Moonshot Kimi K3

Annoncé le 16 juillet 2026, poids publiés le 27 juillet 2026. Le monstre chinois du moment.

| Champ | Valeur |
|---|---|
| Architecture | MoE **~2,8T paramètres totaux, ~104B activés par token** (896 experts : 16 routés + 2 partagés) |
| Checkpoint | ~1,56 To |
| Contexte | **1 048 576 tokens** (1M) |
| Vision | Native |
| Prix API (secondaire) | $3/M entrée, $15/M sortie, cache $0,30/M — NON VÉRIFIÉ contre grille officielle |
| Licence | Poids ouverts, **licence Kimi K3 personnalisée** (propriétaire, pas OSI) |
| Déploiement | API Moonshot ; self-host lourd (64+ accélérateurs conseillés) |

2,8T paramètres : c'est l'ordre de grandeur supposé des plus gros modèles fermés (GPT-4 était estimé ~1,8T). La licence personnalisée n'est pas open-source au sens OSI — à lire avant tout usage au-delà du test. Benchmarks NON VÉRIFIÉS dans cette recherche.

## 26. Google Gemma 4 31B

Poids sortis le 31 mars 2026, API le 2 avril 2026. La vitrine open-weight de Google.

| Champ | Valeur |
|---|---|
| Paramètres | 30,7B **dense** (pas de MoE) |
| Contexte | 256K tokens |
| Entrées | Texte + image ; vidéo par trames (les petits E2B/E4B acceptent aussi l'audio) |
| Langues | 140+ |
| Licence | Poids ouverts, **Apache 2.0** (rupture : Gemma 1–3 utilisaient des conditions custom) |
| Déploiement | Hugging Face, Vertex AI, Ollama, LiteRT |

Deux points notables : le passage à Apache 2.0 (vrai open-source, usage commercial libre) et la stratégie « dense assumé » à 30B là où tout le monde fait du MoE. Benchmarks NON VÉRIFIÉS dans cette recherche (la fiche officielle Google n'a pas été retrouvée). Les checkpoints existent aussi en **NVFP4** (Blackwell) et **GGUF Q4_0 QAT officiels Google** — voir section 85.

## 27. Google Gemini 3.1 Pro

Public preview le 19 février 2026 (update custom-tools le 23 février). Le flagship multimodal de Google au snapshot.

| Champ | Valeur |
|---|---|
| Architecture | MoE (taille non divulguée — Google ne publie jamais) |
| Contexte | 1M tokens |
| Entrées | Texte + image + audio + vidéo ; thinking natif |
| Benchmarks (fournisseur, via secondaires) | SWE-Bench Verified 80,6 % ; GPQA Diamond 94,3 % ; HLE 44,4 % ; Terminal-Bench 2.0 68,5 % ; ARC-AGI-2 77,1 % ; BrowseComp 85,9 % ; tête revendiquée sur 13/16 benchmarks |
| Benchmark (indépendant) | Artificial Analysis, août 2026 : Intelligence Index 48, 113 tok/s, coût pondéré ~$1,74/M |
| Prix (secondaire) | $2/M entrée, $12/M sortie — NON VÉRIFIÉ contre grille officielle |
| Licence | Fermé, API uniquement |

Le croisement est intéressant : le fournisseur revendique la tête sur 13/16 benchmarks, et l'indépendant Artificial Analysis le place à 113 tok/s pour ~$1,74/M pondéré. C'est exactement la lecture croisée recommandée section 119 : ni l'un ni l'autre ne suffit seul.

## 28. Alibaba Qwen4-Exp (preview)

Preview repérée le 26 août 2026, poids ouverts sur Hugging Face. La fiche technique la plus détaillée du lot (config.json vérifiée par un tiers — confiance haute).

| Champ | Valeur |
|---|---|
| Architecture | **125B totaux, 6B actifs par token** (MoE : 512 experts routés, top-10 + 1 expert partagé) |
| Hybride | 48 couches : **36 Gated DeltaNet + 12 attention** (ratio 3:1), hidden 2560 |
| Vision | Oui — `pipeline_tag: image-text-to-text` |
| Détails | Table d'embedding n-gram 51,2B paramètres ; tête MTP 4B |
| Benchmarks | NON VÉRIFIÉS — la doc ne cite pas de scores |
| Licence | Poids ouverts (HF) ; licence NON VÉRIFIÉE |

Le ratio 125B/6B est extrême : 5 % des paramètres actifs par token. Et l'hybride Gated DeltaNet (une variante d'attention linéaire) + attention classique 3:1 montre la direction 2026 : payer l'attention quadratique seulement sur 1 couche sur 4. La tête MTP (Multi-Token Prediction, 4B) sert au décodage spéculatif natif.

## 29. Meta Llama 5 (variantes 5V vision)

Le point le plus fragile du dossier : date du 5 avril 2026 issue d'**une seule source secondaire**, fiche Meta officielle non ouverte. À traiter avec prudence.

| Champ | Valeur |
|---|---|
| Architecture | Famille dense + MoE ; flagship 400B+ évoqué |
| Vision | Via variantes **Llama 5V** (texte+image) |
| Contexte | Une fiche secondaire cite 10M — NON VÉRIFIÉ |
| Licence | Poids ouverts annoncés ; **Llama 5 Community License** (source-available, pas OSI — clause 700M MAU selon analyses secondaires) |
| Contradiction non tranchée | Un digest cite une variante « 405B Apache 2.0 », incompatible avec la licence communautaire Meta |

La clause 700M MAU (monthly active users) est la marque des licences Meta : en dessous, usage libre ; au-dessus, il faut une licence commerciale. Pour un RAG personnel ou PME, c'est transparent. Pour un déploiement à grande échelle, c'est un point juridique.

## 30. Alibaba Qwen3-VL-235B-A22B (Thinking / Instruct)

2025 — hors période, mais le flagship VL open-weight encore leader au snapshot.

| Champ | Valeur |
|---|---|
| Paramètres | 235B totaux (22B actifs selon la nomenclature) |
| Contexte | 262 144 tokens |
| Entrées | Image / vidéo / texte |
| Licence | Poids ouverts, Apache 2.0 (famille Qwen3) |
| Déploiement | vLLM, Transformers, self-host |

235B totaux en Apache 2.0 : c'est le plus gros VLM ouvert permissif du comparatif. Le 22B « actifs » dans le nom est la convention Qwen (A22B = active). Benchmarks à prendre sur la carte officielle — NON VÉRIFIÉS dans cette recherche.

## 31. Xiaomi MiMo-V2-Omni

Signalé le 18 mars 2026 par BenchLM. Le dossier le plus mince : **une seule source** (BenchLM via GitHub), tout le reste NON VÉRIFIÉ.

| Champ | Valeur |
|---|---|
| Contexte | 262K tokens (signalé) |
| Modalités | Multimodal, vision incluse |
| Licence | Propriétaire (pas de poids ouverts repérés) |
| Déploiement | API (selon source secondaire) |

Xiaomi pousse ses modèles MiMo via MiMo Open Platform. À ne citer qu'avec la réserve « source unique » tant qu'une fiche primaire n'est pas ouverte.

## 32. Alibaba Qwen3.5-Omni (0.8B / 2B)

Repérés dans les index de juillet 2026 (date exacte NON VÉRIFIÉE). Les petits VLM de la famille Qwen3.5.

