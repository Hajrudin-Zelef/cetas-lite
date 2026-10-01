---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-14
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Mistral", "Nvidia"]
dates: ["2025-10-31", "2025-12-02", "2026-02-04", "2026-03-12", "2026-03-16", "2026-03-23", "2026-04-28", "2026-07-08", "2026-07-31", "2026-09-27"]
keywords: ["agent", "agents", "apache", "benchmarks", "blackwell", "compute", "datacenter", "distillation", "gpu", "jailbreak", "lean", "mai"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1163, 1253]
sha256: 019b91959766d5fe17cf9f2b0a5d2df05d8f9b7a0b5721f308f849f65f84d743
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Robostral Navigate |
| Sortie | **8 juillet 2026** (annonce officielle Mistral) |
| Statut | ⚠️ annoncé — **poids non publiés, pas d'API publique, pas de pricing** (page officielle : « talk with our team ») |
| Architecture | **8B paramètres**, modèle vision-langage-action (VLA) pour la navigation incarnée ; initialisé depuis un VLM Mistral spécialisé grounding (pointing, counting, object localization) ; entraîné entièrement en simulation |
| Contexte | **non vérifié au 27/09/2026** |
| Features | navigation à partir d'**une seule caméra RGB** (pas de LiDAR/profondeur) + instruction en langage naturel ; navigation « via pointing » (coordonnées image de la cible + orientation) avec fallback en déplacements métriques ; généralise roues/pattes/drones et tailles de robots ; RL en ligne (CISPO) : +3,2 % sans plateau |
| Données | **2,4M trajectoires sur 350k scènes** (source officielle). ⚠️ Conflit : sources secondaires disent ~400K trajectoires / 6K scènes — **la version officielle est retenue**. Prefix-caching : 22× moins de tokens d'entraînement (mois → jours) |
| Benchmarks clés | **R2R-CE : 76,6 % success validation-unseen / 79,4 % validation-seen** ; +9,7 pts vs meilleure approche single-camera, +4,5 pts vs meilleurs systèmes multi-capteurs (chiffres Mistral, non répliqués indépendamment) |
| Poids / licence | **non vérifié** — poids non publiés à ce jour |
| Déploiement | démo entreprise uniquement |

## 102. Voxtral TTS — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | Voxtral TTS (`voxtral-tts-2603`) |
| Sortie | **23 mars 2026** |
| Statut | disponible (poids ouverts selon sources secondaires) |
| Architecture | TTS open-weight sur backbone **Ministral 3B** (secondaire) |
| Features | text-to-speech SOTA (vendor), **zero-shot voice cloning**, multilingue (9 langues, secondaire), streaming temps réel |
| Poids / licence | poids ouverts ; licence exacte **non vérifiée** (une source secondaire indique CC BY-NC 4.0 — non confirmé officiellement) |
| Prix API | **non vérifié au 27/09/2026** |

## 103. Voxtral Mini Transcribe 2 + Realtime ✅ (fiches courtes)

- **Voxtral Mini Transcribe 2** (`voxtral-mini-2602`) + **Voxtral Mini Transcribe Realtime** (`voxtral-mini-transcribe-realtime-2602`) — sortis le **4 février 2026**. Transcription audio (+ diarisation). Specs détaillées : **non vérifiées au 27/09/2026**.

## 104. Mistral Moderation / Shieldstral ✅ (fiches courtes)

- **Mistral Moderation 2603** (`mistral-moderation-2603`) et **Mistral Moderation 2** (`v26.03`, 128K ctx, détection jailbreak) — sortis le **12 mars 2026** (docs officielles).
- **Shieldstral 1.0** — modèle de modération multimodal compact (docs officielles). Date exacte : **non vérifiée**.

## 105. Leanstral, OCR, Codestral ✅ (fiches courtes)

- **Leanstral** (mars 2026) — `labs-leanstral-2603` ; premier code agent open-source pour l'ingénierie de preuves formelles **Lean 4** ; **Leanstral 1.5** existe (docs officielles). Licence/poids : **non vérifié**.
- **Codestral 2508** (`codestral-2508`, 256K ctx) — spécialiste code, successeur de codestral-2501. Date exacte : **non vérifiée** (2025 H2 probable).
- **OCR 4.1** — cité dans la presse comme modèle 2026 ; docs officielles : OCR (25.03), OCR 2 (25.05) dépréciés. Version exacte/date : **non vérifié**.

## 106. Mistral Large 3 ✅ (hors période — déc. 2025, contexte)

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **2 décembre 2025** (hors période) |
| Architecture | **MoE — 675B total / 41B actifs** |
| Contexte | **256K–262K selon sources (conflit : 262K redhat-et/genkit vs 256K 99sono/dim-s — non tranché)** |
| Poids / licence | ouverts, **Apache 2.0** ; multimodal ; partenariat NVIDIA |
| Prix API | **non vérifié** |

## 107. Ministral 3 ✅ (hors période — déc. 2025, contexte)

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **2 décembre 2025** (hors période) |
| Architecture | denses **3B/8B/14B** (base/instruct/reasoning, 9 variantes) ; Cascade Distillation depuis Mistral Small 3.1 (arXiv 2601.08584, secondaire) |
| Contexte | 128K (conflit mineur : 131K/262K selon genkit — non tranché) |
| Poids / licence | **Apache 2.0** ; vision |
| Benchmarks | le 14B reasoning atteint 85 % sur AIME 2025 (vendor) |

## 108. Dépréciations Mistral 2026

Fenêtre se terminant le **31/07/2026** (docs Mistral officielles + issue braintrust) : `magistral-medium-2509` → remplacé par Mistral Medium 3.5 · `magistral-small-2509` → remplacé par Mistral Small 4 · `devstral-2512` → remplacé par Mistral Medium 3.5 · `open-mistral-nemo` → remplacé par Ministral 3 8B · `magistral-medium/small-2506`, `devstral-medium-latest` (dépréciés depuis le 31/10/2025). Devstral 2 (`devstral-2512`, 256K ctx, coding agentique, déc. 2025) est donc **déprécié** au profit de Medium 3.5.

## 109. Vibe, Forge, Mistral Compute

- **Vibe** (26–28 mai 2026) — Le Chat rebrandé en **agent unifié** (Work / Code / Chat) ; Mistral Medium 3.5 en moteur des agents distants ; Pro $14,99/mois (secondaire).
- **Forge** (mars 2026, GTC) — plateforme d'entraînement custom (pré-training + post-training + RL sur données entreprise, on-prem ou Mistral Compute). Secondaire.
- **Mistral Compute** — cloud IA proprio : 18 000 GPU NVIDIA Blackwell, datacenter 44 MW près de Paris (secondaire).

## 110. Tableau récapitulatif Mistral (sept. 2026)

| Modèle | Sortie | Statut | Contexte | Licence poids | Notes |
|---|---|---|---|---|---|
| Voxtral Mini Transcribe 2 | 04/02/2026 | dispo | non vérifié | — | + Realtime |
| Mistral Moderation 2603/2 | 12/03/2026 | dispo | 128K | — | jailbreak detection |
| **Mistral Small 4** | **16/03/2026** | **dispo** | 256K | **Apache 2.0** | MoE 119B/6–6,5B ; unifié |
| Voxtral TTS | 23/03/2026 | dispo | — | ouverte (exacte non vérifiée) | zero-shot voice cloning |
| Leanstral | 03/2026 | dispo | non vérifié | non vérifié | preuves formelles Lean 4 |
| **Mistral Medium 3.5** | **28/04/2026** | **dispo** | 256K | **Modified MIT** | dense 128B ; SWE-bench 77,6 % |
| Vibe | 05/2026 | dispo | — | — | agent unifié |
| Robostral Navigate | 08/07/2026 | **annoncé (pas d'API)** | non vérifié | non publiée | VLA 8B, single-camera |
| Mistral Large 3 | 02/12/2025 | dispo (hors période) | 256K/262K ? | Apache 2.0 | MoE 675B/41B |
| Ministral 3 | 02/12/2025 | dispo (hors période) | 128K | Apache 2.0 | 3B/8B/14B |

---

# PARTIE VIII — Google Gemma 4 (open-weight, Apache 2.0)

## 111. Gemma 4 : le retour à Apache 2.0 — la rupture majeure

