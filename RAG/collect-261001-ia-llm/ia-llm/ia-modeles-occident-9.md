---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-9
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: ["2026-02-12", "2026-03-18", "2026-07-23", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "agi", "benchmarks", "cyber", "gemini", "gemini 3.8", "gemini 4", "mai", "mcp", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [692, 827]
sha256: b723807a24888b39ba066eab3eb68d3ac46e83790bb9989a11a04d707bd1f78b
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

## 63. La série Gemini 3.x : comprendre la logique de nommage

La série 3.x de 2026 est la plus fournie du marché… et la plus confuse. Logique officielle :
- **Le premier chiffre (3)** = génération.
- **Le deuxième chiffre (1, 5, 6, 7, 8)** = itération du tier Flash (le seul tier qui bouge en 2026).
- **Le suffixe** = le tier : **Pro** (frontier, quasi figé depuis fév. 2026), **Flash** (équilibré), **Flash-Lite** (économique), **Cyber** (variante cybersécurité restreinte).
- Anomalie : **il n'existe ni Gemini 3.5 Pro ni Gemini 4** au 27/09/2026 (voir § 77). La ligne Pro est restée à **3.1 Pro Preview** depuis le 19 février 2026 pendant que Google itérait frénétiquement sur Flash (3.5 → 3.8 en 4 mois).

## 64. Gemini 3 Pro Preview ✅ (hors période, pour la trajectoire)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3-pro-preview` |
| Sortie | **18 novembre 2025** (premier modèle Gemini 3) |
| Statut | **retiré le 9 mars 2026** ; le slug pointe désormais vers `gemini-3.1-pro-preview` |
| Contexte | **non vérifié au 27/09/2026** |
| Features | multimodal ; raisonnement ; mode « Deep Think » pour les abonnés Google AI Ultra |
| Poids | fermés, licence propriétaire |
| Benchmarks | cité comme top LMArena à l'époque (préférence humaine, plusieurs mois) — détails non vérifiés |
| Prix API | **non vérifié au 27/09/2026** |

## 65. Gemini 3 Flash (Preview déc. 2025 → GA) ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **17 décembre 2025** (Preview) |
| Statut | GA sous l'ID `gemini-3-flash` ; migration officielle recommandée vers `gemini-3.8-flash` |
| Contexte | 1M tokens (reporté par des guides tiers, mars 2026) |
| Features | multimodal ; raisonnement « thinking » limité (niveaux) |
| Poids | fermés |
| Prix API | **$0,30 / $1,50** par MTok (guides tiers, mars 2026) |
| Déploiement | Gemini API, AI Studio, Vertex AI ; a été le modèle par défaut de l'app Gemini |

## 66. Gemini 3 Deep Think ✅ (mode, pas un modèle API)

| Champ | Valeur vérifiée |
|---|---|
| Sortie | mode lancé fin 2025 pour AI Ultra ; **grosse mise à jour : février 2026** (« upgrade majeur », 9to5google 12/02/2026) |
| Statut | disponible (app AI Ultra + early access entreprise API) ; **pas d'ID API public** |
| Features | raisonnement profond parallèle, ciblé maths/sciences/ingénierie |
| Poids | fermés |
| Prix API | **non vérifié au 27/09/2026** |

## 67. Gemini 3.1 Pro Preview ✅ (le Pro le plus récent — depuis février 2026)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.1-pro-preview` |
| Sortie | **19 février 2026** (Preview) |
| Statut | **disponible** ; toujours le modèle Pro le plus récent de l'API en sept. 2026 |
| Contexte | 1M tokens entrée / 64K max sortie |
| Features | vision, function calling, URL context, grounding Google Maps (depuis le 18/03/2026 pour les modèles Gemini 3), thinking levels (low/medium/high, défaut high) |
| Poids | fermés |
| Benchmarks (vendor) | ARC-AGI-2 : 77,1 % (plus du double de Gemini 3 Pro) ; SWE-Bench Verified : 80,6 % ; GPQA Diamond : 94,3 % |
| Prix API | **$2 / $12** (≤200K tokens), **$4 / $18** (>200K) par MTok |
| Déploiement | Gemini API, AI Studio, Vertex AI (Vertex : toujours « Preview », pas de label GA stable au 23/07/2026) |
| Déclinaison | `gemini-3.1-pro-preview-customtools` — variante pour agents « bash-heavy » et outils custom |

## 68. Gemini 3.1 Flash-Lite Preview ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.1-flash-lite` |
| Sortie | **3 mars 2026** (Preview) |
| Statut | disponible (Preview) ; modèle texte le moins cher de la série 3 |
| Architecture | dérivé architectural de Gemini 3 Pro selon la model card officielle DeepMind, optimisé latence/débit — **non vérifié indépendamment au 27/09/2026** |
| Contexte | 1M tokens entrée / 64K max sortie |
| Features | thinking levels : minimal / low / medium / high ; ~287 tok/s (mesure tiers) |
| Poids | fermés |
| Benchmarks | GPQA Diamond 86,9 %, FACTS 40,6 %, LiveCodeBench 72,0 % (sources tiers, mars 2026) |
| Prix API | **$0,25 / $1,50** (entrée audio $0,50) par MTok |

## 69. Gemini 3.5 Flash ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.5-flash` |
| Sortie | **19 mai 2026** (Google I/O 2026, GA immédiate) — devenu le même jour le modèle par défaut de l'app Gemini et d'AI Mode |
| Statut | disponible ; considéré « legacy pipelines only » face aux 3.6/3.7/3.8 (migration officielle) |
| Contexte | 1M tokens entrée / 64K max sortie |
| Features | multimodal (texte, image, vidéo, audio, PDF) ; thinking à 4 niveaux (minimal/low/medium/high) ; structured outputs (JSON schema) ; implicit caching ; ~4× plus rapide que les frontier en tok/s (vendor) ; agentic video understanding ajouté à la ligne Flash le 1er sept. 2026 |
| Poids | fermés |
| Benchmarks (vendor) | Terminal-Bench 2.1 : 76,2 % ; MCP Atlas : 83,6 % ; CharXiv (multimodal) : 84,2 % ; ARC-AGI-2 : 77,1 % ; GDPval-AA : 1656 Elo ; Finance Agent v2 : 57,9 % |
| Prix API | **$1,50 / $9,00** par MTok ; cached input $0,15 (−90 %) |
| Déploiement | Gemini API, AI Studio, Gemini app (défaut), AI Mode Search, plateforme Antigravity |

## 70. Gemini 3.5 Flash-Lite ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.5-flash-lite` |
| Sortie | **21 juillet 2026** (lancé avec 3.6 Flash, moins cher que 3.5 Flash) |
| Statut | disponible ; chemin de migration « cheap » recommandé pour le travail mécanique |
| Contexte | **non vérifié au 27/09/2026** — probablement 1M, à confirmer dans les docs API |
| Features | thinking levels : minimal / low / medium / high |
| Poids | fermés |
| Prix API | **$0,30 / $2,50** par MTok |

## 71. Gemini 3.6 Flash ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.6-flash` |
| Sortie | **21 juillet 2026** (GA sur la Gemini Enterprise Agent Platform) |
| Statut | disponible ; 2ᵉ Flash de l'été 2026 |
| Contexte | **non vérifié précisément au 27/09/2026** — la doc d'annonce ne publie pas de contexte |
| Features | thinking levels : medium (défaut), minimal, low, high ; multimodal |
| Poids | fermés |
| Prix API | lancement : $1,50 / $7,50 ; tarif courant aligné sur la promo Flash **$0,75 / $3,75** (sept. 2026) |

## 72. Gemini 3.5 Flash Cyber ✅ (accès restreint — Fairwind Program)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.5-flash-cyber` (nom exact d'ID à confirmer) |
| Sortie | **21 juillet 2026** (annoncé avec 3.6 Flash) |
| Statut | disponible sous accès restreint — **Fairwind Program** (autorités gouvernementales, opérateurs d'infrastructures critiques, mainteneurs) |
| Features | variante cybersécurité : mitigations cyber-offense assouplies pour découverte de vulnérabilités |
| Poids | fermés |
| Prix API | **non vérifié au 27/09/2026** |

## 73. Gemini 3.7 Flash ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `gemini-3.7-flash` |
| Sortie | **13 août 2026** (« visé coding et agents ») |
| Statut | disponible ; migration officielle recommandée vers 3.8 Flash (mais pas de date de shutdown annoncée) |
| Contexte | 1M tokens entrée / 64K max sortie (docs développeur Google, repris par plusieurs guides) |
| Features | multimodal (texte, image, vidéo, audio, PDF) ; thinking levels low / medium / high (défaut medium) ; pas de `minimal` contrairement à 3.6 |
| Poids | fermés |
| Benchmarks (vendor, table de lancement Google) | DeepSWE v1.1 : 65,3 % ; Terminal-bench 4.0 : 11,2 % ; OSWorld-2.0 : 50,6 % ; Vals Finance Agent v2 : 59,0 % |
| Prix API | **promo $0,75 / $3,75 jusqu'au 31 déc. 2026** (moitié du prix initial de 3.6 Flash) |

## 74. Gemini 3.8 Flash ✅ (modèle texte le plus récent au 27/09/2026)

