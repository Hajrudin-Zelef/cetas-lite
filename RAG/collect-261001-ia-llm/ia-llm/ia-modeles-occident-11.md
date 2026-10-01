---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-11
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Microsoft", "OpenAI", "SpaceX", "xAI"]
dates: ["2025-10-15", "2025-11-18", "2025-12-17", "2026-01-24", "2026-02-19", "2026-03-03", "2026-03-09", "2026-05-19", "2026-07-18", "2026-07-21", "2026-08-13", "2026-09-02", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "benchmarks", "copilot", "cyber", "gemini", "gemini 3.8", "gemini 4", "grok", "grok 4", "mai"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [907, 1002]
sha256: ce8775ac44677e901ae6ba08b8ee13c3a1a5008454cf1cd813227e44bb6dc53d
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Modèle | Sortie | Statut | Contexte | Prix in/out par MTok | Notes |
|---|---|---|---|---|---|
| Gemini 3 Pro Preview | 18/11/2025 | retiré 09/03/2026 | non vérifié | non vérifié | Deep Think |
| Gemini 3 Flash | 17/12/2025 | GA → migrer 3.8 | 1M | $0,30/$1,50 | ex-défaut app Gemini |
| Gemini 3.1 Pro Preview | 19/02/2026 | **dispo — Pro actuel** | 1M/64K | $2/$12 (≤200K) | customtools variant |
| Gemini 3.1 Flash-Lite | 03/03/2026 | preview | 1M/64K | $0,25/$1,50 | le moins cher |
| Gemini 3.5 Flash | 19/05/2026 | dispo (legacy pipelines) | 1M/64K | $1,50/$9,00 | défaut app Gemini |
| Gemini 3.5 Flash-Lite | 21/07/2026 | dispo | non vérifié | $0,30/$2,50 | cheap path |
| Gemini 3.6 Flash | 21/07/2026 | dispo | non vérifié | $0,75/$3,75 (promo) | 2ᵉ Flash été 2026 |
| Gemini 3.5 Flash Cyber | 21/07/2026 | **Fairwind restreint** | non vérifié | non vérifié | cyber-offense |
| Gemini 3.7 Flash | 13/08/2026 | dispo → migrer 3.8 | 1M/64K | $0,75/$3,75 (promo) | coding/agents |
| **Gemini 3.8 Flash** | **02/09/2026** | **dispo — recommandé** | 1M/64K | **$0,75/$3,75 (promo)** | Terminal-bench 2.1 : 89,4 % |
| Gemini 3.8 Flash Cyber | 02/09/2026 | **Fairwind restreint** | non vérifié | non vérifié | découverte vulns |
| « Gemini 3.5 Pro » | — | **N'EXISTE PAS** | — | — | — |
| « Gemini 4 » | — | **N'EXISTE PAS** | — | — | — |
| Nano Banana 2 / 2 Lite / Pro | 02/2026–06/2026 | dispo | 32K | $0,0336–$0,24/image | SynthID |
| Veo 3.1 | 15/10/2025 | **dispo — actuel** | 1024 tok prompt | non vérifié | audio natif |
| « Veo 4 » | — | **N'EXISTE PAS** | — | — | remplacé par Omni |
| Gemini Omni Flash | 19/05/2026 | dispo (app/Shorts/Flow) | non vérifié | non vérifié | max 10 s/clip |

---

# PARTIE V — xAI : la famille Grok

## 83. Contexte entreprise xAI (2026) — SpaceX, Cursor, Grok Bot

Éléments d'entreprise cités avec leurs sources (secondaires à la recherche modèles) : xAI a été **absorbée par SpaceX** ; SpaceX s'est introduit en bourse en **juin 2026** (ticker SPCX, ~$1 800 Mds de valorisation, $75 Mds levés — ainvest.com, sept. 2026). Musk a **acquis Cursor** (éditeur de code IA) pour **$60 Mds** en juin 2026 (selon mui-api). **Grok Build** — agent de code terminal-native, open-sourcé le 15 juillet 2026 (xai-org/grok-build ; source-available, contributions externes refusées). **Grok Bot** — équipe d'agents always-on (mi-août 2026), ~418K utilisateurs hebdo au 14 sept. 2026 (Bloomberg via PYMNTS).

## 84. Grok 3 / Grok 3 Mini ✅ (hors période, retirés)

| Champ | Valeur vérifiée |
|---|---|
| Noms / IDs | `grok-3`, `grok-3-mini` |
| Sortie | février 2025 (avant la période, inclus pour la trajectoire) |
| Statut | **retirés le 15 mai 2026 à 12h00 PT** (vague de retraits API xAI) ; slugs redirigés vers `grok-4.3` |
| Contexte | 131K tokens (grok-3) |
| Features | raisonnement (Think / DeepSearch) ; support de `reasoning_effort` sur grok-3 |
| Poids | fermés |

## 85. Grok Code Fast 1 ✅ (retiré)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `grok-code-fast-1` (nom de code stealth : « sonic ») |
| Sortie | **26 août 2025** |
| Statut | **retiré le 15 mai 2026** (même vague) ; remplacement recommandé : `grok-4.6` |
| Architecture | 314B params MoE **selon sources tiers** — **non vérifié officiellement au 27/09/2026** ; conçu pour le code de A à Z (pas un distillat) |
| Contexte | 256K tokens entrée ; 10K max sortie (source tiers) |
| Features | code/agentic coding, tool use (grep, terminal, file edits), structured outputs, traces de raisonnement visibles ; pas d'input image |
| Poids | fermés |
| Benchmarks (tiers) | SWE-Bench Verified : 70,8 % ; LiveCodeBench : 80,0 % |
| Prix API | $0,20 / $1,50 (entrée/sortie) ; cache $0,02/M (sources tiers) |
| Déploiement | API xAI (historique) ; offert gratuitement des semaines via GitHub Copilot, Cursor, Cline, Windsurf, OpenCode Zen (tier gratuit terminé le 24/01/2026) |

## 86. Grok 4.1 / Grok 4.1 Fast ✅ (retirés)

| Champ | Valeur vérifiée |
|---|---|
| Noms / IDs | `grok-4-1-fast-reasoning`, `grok-4-1-fast-non-reasoning` |
| Sortie | **17 nov. 2025** (Grok 4.1) ; **19 nov. 2025** (Grok 4.1 Fast API) |
| Statut | **variantes Fast retirées le 15 mai 2026**, slugs redirigés vers `grok-4.3` |
| Contexte | 2M tokens (variantes Fast) |
| Features | 4.1 : écriture créative, nuances émotionnelles, instruction following ; Fast : tool calling, web+X search, file retrieval, code execution, MCP |
| Poids | fermés |

## 87. Grok 4.3 ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `grok-4.3` |
| Sortie | **nuit du 30 avril → 1er mai 2026** (annonce) ; API ouverte à tous le **5–6 mai 2026** |
| Statut | **disponible** — flagship de l'époque ; les anciens slugs (`grok-3`, `grok-4-1-fast-*`, `grok-4-fast-*`, `grok-4-0709`, `grok-code-fast-1`, `grok-imagine-image-pro`) y sont redirigés depuis le 15 mai 2026 |
| Architecture | nouvelle base « Grok 4.20 » améliorée, échelle similaire (The Tech Outlook, mai 2026) — détails **non vérifiés officiellement** |
| Contexte | **1M tokens** |
| Features | raisonnement configurable (none / low / medium / high) ; function calling ; structured outputs ; tool calling agentique ; text+image en entrée, texte en sortie ; input vidéo natif ; génération de documents ; suite de clonage vocal (Custom Voices : 1 min d'enregistrement → clone en <2 min, $4,20/M caractères TTS, $3/h realtime) ; catalogue de 80+ voix, 28 langues |
| Poids | fermés |
| Benchmarks (vendor + tiers, mai 2026) | leaderboards Artificial Analysis (tool calling agentique, instruction following) ; #1 ValsAI entreprise (case law, corporate finance) ; taux d'hallucination le plus bas du marché (claim xAI) |
| Prix API | **$1,25 / $2,50** par MTok (−83 % sur l'output vs génération précédente) |

## 88. Grok 4.5 ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `grok-4.5` |
| Sortie | **8 juillet 2026** (annonce le 10 juillet par Dataconomy) |
| Statut | **disponible** (supersedé comme flagship par 4.6 le 12 août 2026, toujours servi) |
| Architecture | fondation **V9** de xAI ; entraîné aux côtés de Cursor ; taille **non vérifiée officiellement** |
| Contexte | 500K ? — **non vérifié précisément au 27/09/2026** ; cache-read $0,30/M cité pour 4.5 (vs $0,50 pour 4.6) |
| Features | ciblé coding, finance, juridique ; tool use ; `reasoning_effort` supporté (style OpenAI) |
| Poids | fermés |
| Benchmarks (vendor) | DeepSWE 1.1 : 54 % (vs 65,9 % pour 4.6) ; Artificial Analysis Intelligence Index : 61 |
| Prix API | **$2 / $6** par MTok (aligné 4.6/4.7) |
| Déploiement | API xAI, Cloudflare AI Gateway, Grok Build (modèle par défaut depuis le 18/07/2026), Cursor |

## 89. Grok 4.6 ✅

