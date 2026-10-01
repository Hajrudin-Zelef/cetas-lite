---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-10
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "EU", "Google"]
dates: ["2025-10-15", "2026-02", "2026-02-26", "2026-06-25", "2026-06-30", "2026-09-27", "2026-12-31"]
keywords: ["agent", "agentic", "agents", "arr", "benchmark", "benchmarks", "cyber", "gemini", "gemini 3.8", "gemini 4", "mai", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [828, 906]
sha256: ff5b07e4767e8edf81cf89c7ea80341d5b5757dfa017cde5c9bfb83906a71e7b
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.8-flash` |
| Sortie | **2 septembre 2026** (GA immédiate, 3ᵉ Flash en 6 semaines) |
| Statut | **disponible** — modèle recommandé par Google pour software engineering, agents autonomes, workflows multi-étapes complexes |
| Contexte | 1M tokens entrée (1 048 576) / 64K (65 536) max sortie |
| Features | multimodal entrée (texte, image, vidéo, audio, PDF), sortie texte ; thinking levels low / medium / high (défaut medium) ; function calling ; agentic video understanding (1er sept. 2026) |
| Poids | fermés |
| Benchmarks (table de lancement Google, claims vendor) | Terminal-bench 2.1 : 89,4 % (vs Opus 5 : 89,1 %) ; DeepSWE v1.1 : **73,7 %** (secondtalent) — **71,0 %** selon Neowin (**écart entre sources, non réconcilié**) ; HLE-Verified : 54,9 % ; long vidéo : 87,8 % ; chart reasoning : 86,2 % ; Terminal-bench 4.0 : 19,1 % (vs Opus 5 : 51,8 %) ; OSWorld-2.0 : 59,0 % (vs 75,4 %) ; Vals Finance Agent v2 : 61,4 % ; Harvey Legal Agent : 10,0 % ; GDPVal-AA Elo : 1545 |
| Prix API | **promo $0,75 / $3,75 jusqu'au 31 déc. 2026** ; $1,50 / $7,50 à partir du 1er jan. 2027 |

## 75. Gemini 3.8 Flash Cyber ✅ (accès restreint)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | ID exact **non vérifié au 27/09/2026** |
| Sortie | **2 septembre 2026** (avec 3.8 Flash) |
| Statut | disponible sous accès restreint — **Fairwind Program** ; mitigations cyber-offense assouplies |
| Features | découverte autonome de vulnérabilités ; framework sécurité Google Frontier Safety Framework pour la variante standard (CBRN / cyber-offense mitigé sur le modèle standard) |
| Poids | fermés |
| Benchmarks (vendor) | CWE-Bench (Collinear) : 47,2 % pass@1 ; CyberGym : frontier en découverte autonome C/C++ ; benchmark interne 20 langages : >70 % de succès |
| Prix API | **non vérifié au 27/09/2026** — prix non publié |

## 76. Variantes TTS de Gemini 3.x ⚠️

- `gemini-3.1-flash-tts-preview` → migration recommandée vers `gemini-3.8-flash-tts` ou `gemini-3.8-flash-lite-tts` (guide de migration Google, sept. 2026).
- Existence des IDs `gemini-3.8-flash-tts` / `gemini-3.8-flash-lite-tts` : **non vérifiée au 27/09/2026** au-delà du guide de migration.

## 77. « Gemini 3.5 Pro » et « Gemini 4 » — NON TROUVÉ ❌

Aucune trace de sortie ni d'annonce vérifiable au 27 sept. 2026. Plusieurs sources tiers affirment explicitement : « Gemini 3.5 Pro remains unreleased with no date » / « Gemini 3.5 Pro and Gemini 4 are still nowhere ». La ligne Pro n'a pas bougé depuis le 19 fév. 2026 (3.1 Pro Preview). Requêtes tentées : « Google Gemini 3.5 3.6 3.7 3.8 models 2026 » ; « Gemini 3.1 Pro Preview release February 2026 » ; « Gemini 3.1 Flash-Lite launch date ».

## 78. Nano Banana : la génération d'images de Google ✅

« Nano Banana » est le nom donné par Google à ses capacités natives de génération d'images dans Gemini (historiquement le nom de code stealth de `gemini-2.5-flash-image`), pas un modèle figé.

| Variante | ID | Sortie | Statut | Prix par image | Notes |
|---|---|---|---|---|---|
| Nano Banana (original) | `gemini-2.5-flash-image` | mi-2025 (date exacte **non vérifiée**) | **legacy** — migrer vers 2 Lite | non vérifié | 32K contexte ; sortie 1K/2K ; pas de search grounding, pas de thinking |
| Nano Banana Pro | `gemini-3-pro-image` | nov. 2025 (Preview dépréciée 25/06/2026) | **dispo — premium actuel** | $0,134 (1K/2K) ; $0,24 (4K) | sortie 1K/2K/4K ; 14 images de référence ; search grounding ; typographie quasi parfaite ; intégré Adobe Firefly/Photoshop ; raisonnement LLM avant génération |
| Nano Banana 2 | `gemini-3.1-flash-image` | 26/02/2026 (Preview ; stable depuis 25/06/2026) | **dispo — workhorse par défaut** | $0,067 (1K) / $0,101 (2K) / $0,151 (4K) | 32K contexte ; sortie 512px/1K/2K/4K ; ratios extrêmes ; 14 références ; search grounding ; entrée PDF ; latence ~1 s |
| Nano Banana 2 Lite | `gemini-3.1-flash-lite-image` | juin 2026 | **dispo — la plus récente** | $0,0336 (1K) | 1K uniquement ; brouillons rapides ; pas de search grounding |

- Toutes les images générées portent un watermark **SynthID** (guide officiel Google).
- Les IDs preview (`gemini-3.1-flash-image-preview`, `gemini-3-pro-image-preview`) ont été dépréciés le 25 juin 2026.

## 79. Veo : la génération vidéo Google ✅ (3.1 = version actuelle)

| Version | ID | Sortie | Statut | Notes |
|---|---|---|---|---|
| Veo 1 | — | I/O 2024 | **déprécié** | 1080p texte→vidéo, research preview |
| Veo 2 | `veo-2.0-generate-001` | déc. 2024 | **arrêté via Gemini API le 30/06/2026** | 4K, physique ; variantes Fast |
| Veo 3 | `veo-3.0-generate-001` | mai 2025 (I/O 2025) | arrêté Gemini API 30/06/2026 ; listé GA sur Vertex AI | **premier modèle vidéo avec audio natif synchronisé** |
| **Veo 3.1** | `veo-3.1-generate-preview` | **15/10/2025** ; update 4K/1080p jan. 2026 ; Veo 3.1 Lite en public preview (printemps 2026) | **dispo — version officielle actuelle** | audio natif ; 4K/1080p ; sortie verticale native 9:16 ; consistance personnages via références ; extension de scène jusqu'à 148 s ; SynthID |
| **Veo 4** | — | — | **NON TROUVÉ** | pas d'annonce, pas de doc ; à l'I/O 2026 Google a lancé **Gemini Omni** au lieu d'un Veo 4 |

Prix API Veo : **non vérifiés au 27/09/2026** (tarification par seconde/clip non confirmée dans les recherches).

## 80. Gemini Omni Flash ✅ (remplace la suite Veo)

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **19 mai 2026** (Google I/O 2026) — disponible le même jour sur l'app Gemini, YouTube Shorts, Google Flow |
| Statut | disponible (app, Shorts, Flow) ; **API entreprise (Vertex AI / Gemini Enterprise Agent Platform) toujours non lancée au sept. 2026** (4 mois après l'annonce) |
| Features | plateforme multimodale unifiée (texte, image, audio, vidéo en une sortie) ; audio natif synchronisé ; édition post-génération conversationnelle sans re-prompting ; watermark SynthID ; **max 10 s par clip** ; Ask YouTube |
| Poids | fermés |
| Prix API | **non vérifié au 27/09/2026** |
| Note | pas de région EU sélectionnable / résidence EU garantie via Gemini API (sept. 2026) — point de vigilance conformité |

## 81. Prix Gemini : la stratégie du Flash agressif

- **Flash** : $0,75/$3,75 en promo jusqu'au 31/12/2026 (puis $1,50/$7,50) — Google vend le frontier-approchant au prix du milieu de gamme.
- **Flash-Lite** : $0,30/$2,50 (3.5) — le moins cher de la série 3.
- **Pro** : $2/$12 (≤200K), $4/$18 (>200K) — inchangé depuis février 2026.
- **Implicit caching** sur la ligne Flash : le cache se fait sans configuration explicite.
- Les variantes Cyber (Fairwind) : prix non publiés.

## 82. Tableau récapitulatif Google (sept. 2026)

