---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-5
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "Google", "Moonshot", "Nvidia", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-09-24"]
keywords: ["agents", "apache", "arr", "benchmark", "benchmarks", "claude", "gemini", "glm", "gpt-5.6", "gpu", "grok", "grok 4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [473, 576]
sha256: d3b5dbb78bf811dacd0e196abd65b6d1b26a3d66aad5011cfca356700dc24042
---

# Encyclopédie des modèles IA — Volume 3

| Champ | Valeur |
|---|---|
| Tailles | 0.8B / 2B (+ builds ONNX) |
| Entrées | Texte + image + audio (famille omni) |
| Benchmarks (cartes HF, fournisseur/communauté) | Qwen3.5-2B : MMMU 64,2 ; MMMU-Pro 50,3 ; OCRBench 84,5/85,4 |
| Licence | Poids ouverts, Apache 2.0 |
| Déploiement | ONNX, Transformers, **edge** |

C'est la tendance « VLM de poche » : 2B paramètres avec 84,5+ en OCRBench, ça tourne sur un NPU de PC portable ou un Jetson. Pour ton usage : décrire des schémas ou des captures d'écran **en local**, sans envoyer tes documents à une API. Le build ONNX officiel (`onnx-community/Qwen3.5-0.8B-ONNX`) simplifie le déploiement.

## 33. xAI Grok 4.5

Public launch le 16 juillet 2026 (corroboré par uniaibench, avec liens vers les docs xAI).

| Champ | Valeur |
|---|---|
| Entrées | Texte + image + fichiers (selon secondaires) |
| Raisonnement | Configurable |
| Specs rapportées | ~1,5T paramètres, contexte 500K — **NON VÉRIFIÉS** |
| Licence | Fermé, API xAI |

Date et API corroborées ; tout le reste (taille, contexte) est du rapporté non vérifié. xAI ne publie pas d'architecture — même opacité qu'Anthropic et OpenAI.

## 34. Mentions vision (prudence)

- **Claude Opus 4.8** (28 mai 2026 selon source secondaire ; 1M contexte, $5/$25 par M ; SWE-bench Verified 88,6 %) — **aucune source Anthropic officielle ouverte dans cette recherche**. Le chiffre SWE-bench 88,6 % est remarquable s'il se confirme.
- **GPT-5.6 Sol** (9 juillet 2026 ; 1,05M contexte / 128K sortie ; ~$4/$20 en promo) — non confirmé par OpenAI directement.
- **Qwen3.7-Plus / Qwen3.7-Max** (juin 2026) : agents multimodaux propriétaires, pas open.
- **GLM-4.6V / Flash** (décembre 2025 ; 106B et 9B ; 128K ; poids téléchargeables ; revendications fournisseur MathVista 88,2 / WebVoyager 81,0).
- **Gemini 3.6/3.7 Flash** (juillet–août 2026), **GPT-5.5** : cités dans des comparatifs secondaires, fiches primaires non ouvertes.

## 35. Tableau comparatif vision

| Modèle | Taille (total/actifs) | Contexte | Modalités | Licence | Statut vérification |
|---|---|---|---|---|---|
| Nemotron 3 Nano Omni | 30B / 3B | 256K | T+I+A+V | NVIDIA Open Model | Solide (blog NVIDIA + bench ind.) |
| Kimi K3 | 2,8T / 104B | 1M | T+I (+V/A ?) | Licence custom | Prix/benchs non vérifiés |
| Gemma 4 31B | 30,7B dense | 256K | T+I (+V trames) | Apache 2.0 | Benchs non vérifiés |
| Gemini 3.1 Pro | MoE (non publié) | 1M | T+I+A+V | Fermé | Prix non vérifié |
| Qwen4-Exp | 125B / 6B | — | T+I | Ouvert (licence ?) | Benchs non vérifiés |
| Llama 5V | 400B+ évoqué | 10M (?) | T+I | Community License | Source unique |
| Qwen3-VL-235B | 235B / 22B | 262K | T+I+V | Apache 2.0 | Benchs non vérifiés |
| MiMo-V2-Omni | — | 262K (?) | Multimodal | Propriétaire | Source unique |
| Qwen3.5-Omni | 0.8B / 2B | — | T+I+A | Apache 2.0 | Benchs cartes HF |
| Grok 4.5 | ~1,5T (?) | 500K (?) | T+I | Fermé | Specs non vérifiées |

## 36. La vision dans un RAG documentaire

Concrètement, à quoi sert un VLM dans ton pipeline ? Trois usages, par ordre de maturité :

1. **Légender les images du corpus.** Tes guides contiennent des schémas (unifilaires d'onduleurs, topologies eNSP, faces avant d'équipements). Un Qwen3.5-2B en local décrit chaque image en 200 mots ; cette description devient un chunk comme les autres. Coût : quelques secondes par image sur un GPU modeste.
2. **L'OCR intelligent.** Sur un PDF scanné ou une datasheet en image, un VLM fait mieux qu'un OCR classique : il comprend les tableaux (cellules fusionnées, en-têtes) et les restitue en markdown structuré. OCRBench 84,5+ (Qwen3.5-2B) mesure exactement ça.
3. **La question multimodale.** « Sur ce schéma unifilaire, où est le bypass ? » avec l'image jointe à la question. Ça demande un VLM en inférence à chaque requête — plus cher, à réserver aux usages où l'image est indispensable.

Piège classique : indexer la description générée **sans garder le lien vers l'image source**. Quand le LLM répond en citant la description, l'utilisateur veut voir le schéma d'origine. Toujours stocker `image_path` + `page` en métadonnées (section 112).

## 37. Vidéo : panorama 2026

La génération vidéo a basculé en 2026 : l'audio natif synchronisé (dialogues, bruitages, musique) est devenu standard, les clips s'allongent (15–30 s), et le multi-plan (jusqu'à 6 plans en une génération) arrive. Le fait marquant du snapshot : **Sora 2, le pionnier d'OpenAI, a vu son API fermée le 24 septembre 2026** — trois jours avant l'arrêt de cette recherche. Le marché se recompose autour de ByteDance (Seedance), Kuaishou (Kling), Runway et Google (Veo).

Avertissement transversal : specs et prix ci-dessous viennent majoritairement de **comparatifs secondaires** (GitHub, agrégateurs), pas des grilles officielles. Les prix varient selon le chemin d'accès (API officielle vs fal.ai vs Replicate vs Krea) et la résolution. Ne jamais chiffrer un budget dessus sans revalidation.

## 38. OpenAI Sora 2 / Sora 2 Pro — STATUT : ARRÊTÉ

| Champ | Valeur |
|---|---|
| Lancement | 2025 |
| **Fin de vie API** | **24 septembre 2026** — les modèles `sora-2` et `sora-2-pro` ne sont plus servis |
| Specs historiques | 20–25 s ; 720p (Sora 2) / 1080p (Pro) ; audio natif ; génération, extension, édition |
| Prix historiques | $0,10/s (Sora 2) ; $0,30–0,70/s (Pro) |
| Statut benchmark | Elo Artificial Analysis historiquement élevé, dépassé début 2026 par Seedance 2.0 (1269) |

Le statut « API fermée le 24/09/2026 » repose sur un **miroir GitHub** des docs développeur OpenAI, pas sur openai.com ouvert en direct — même si deux sources secondaires concordent, c'est à confirmer avant d'en faire un fait définitif. Conséquence pratique : **ne pas intégrer Sora à un nouveau projet**. Les générations existantes restent lisibles, mais il n'y a plus de route API.

## 39. ByteDance Seedance 2.0

Lancé en février 2026, disponible mondialement via API partenaires depuis avril 2026. Le n°1 indépendant du moment.

| Champ | Valeur |
|---|---|
| Concept | Modèle **unifié multimodal vidéo+audio** : texte + images + audio + vidéo en un seul appel |
| Clips | 4–15 s, 24 fps ; 480p/720p natif (1080p/4K selon hébergeurs) |
| Audio | Natif synchronisé en une passe : dialogues + lip-sync (`Character says: "..."`), SFX, ambiances, musique |
| Plans | Multi-plans (jusqu'à ~6 plans) dans une génération ; contrôle caméra niveau réalisateur |
| Références | Jusqu'à 9 images + 3 vidéos + 3 audios (12 assets) ; cohérence d'identité |
| Formats | auto / 21:9 / 16:9 / 4:3 / 1:1 / 3:4 / 9:16 ; extension vidéo |
| Benchmark (indépendant) | **Elo 1269 Artificial Analysis** — #1 en février 2026, devant Veo 3, Sora 2, Runway Gen-4.5 |
| Prix (secondaires, selon le chemin) | BytePlus ModelArk ~$0,01–0,03/s (KYC entreprise) ; fal.ai $0,24 (Fast 720p) – $0,68 (Std 1080p) ; Krea $0,0677–0,0849 |
| Accès | fal.ai, Replicate, Runway, Higgsfield, BytePlus ModelArk |

L'écart de prix entre le chemin officiel BytePlus ($0,01–0,03/s) et fal.ai ($0,24–0,68/s) illustre la règle : le même modèle coûte **10 à 20× plus cher** via un agrégateur que via le fournisseur direct. Pour un usage régulier, le KYC entreprise chez BytePlus se rentabilise vite.

## 40. ByteDance Seedance 2.5

Version suivante, repérée en 2026 (date exacte NON VÉRIFIÉE), disponible via agrégateurs (MuAPI cité).

| Champ | Valeur |
|---|---|
| Clips | Jusqu'à **30 s**, **4K** |
| Références | Jusqu'à 50 images de référence |
| Continuité | Multi-scènes ; audio+vidéo en une passe |
| Endpoints dédiés | **Video Edit** et **Spicy** (non censuré) selon un comparatif |
| Prix (secondaire) | ~$0,05–0,25/s selon résolution (MuAPI) — NON VÉRIFIÉ |

