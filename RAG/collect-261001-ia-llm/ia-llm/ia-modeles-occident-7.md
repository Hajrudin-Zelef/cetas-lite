---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-7
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2025-01-31", "2025-09-29", "2026-02-17", "2026-03-24", "2026-05", "2026-05-25", "2026-06-30", "2026-08-12", "2026-08-31", "2026-09-22", "2026-09-24", "2026-09-25", "2026-09-27", "2026-09-29"]
keywords: ["agentic", "agents", "aws", "bedrock", "benchmark", "benchmarks", "claude", "consumer", "distillation", "fable 5", "fine-tuning", "foundry"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [514, 591]
sha256: 10acaf7fc9241b67ed8d59ad55c8dbb9b62c6b966a30cb5ed96929c71070e65e
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Opus 5.5 / `claude-opus-5-5` |
| Sortie | **22/09/2026** (même jour que GPT-6 Sol/Luna d'OpenAI — « guerre des prix ») |
| Statut | disponible (GA). System Card publiée le jour du lancement |
| Contexte | **non vérifié au 27/09/2026** — aucune source fiable ne donne le chiffre ; probable 1 M par continuité, **non affirmé** |
| Features | adaptive thinking **obligatoire** (plus de désactivation, calibration via `effort` uniquement), legacy computer-use déprécié, thinking blocks liés aux conversations (anti-distillation), **fin des plafonds d'usage de 5 heures** pour les abonnés, limites d'usage élargies + reset stockable (Pro/Max/Team/Enterprise), prose plus concise |
| Fine-tuning | non offert |
| Poids | fermés. API Anthropic, AWS, Google Cloud, Microsoft Azure ; Zero Data Retention par défaut sur Bedrock |
| KV cache | prompt caching standard (détails non re-spécifiés pour 5.5 → **non vérifié au 27/09/2026** au-delà du standard) |
| Benchmarks clés | **SWE-bench Pro 89,9 %**, **Terminal-Bench 4.0 66,4 %** ; Artificial Analysis Intelligence Index v4.3.2 (24/09/2026) : **57,6** (1er, devant Fable 5.1 à 53,4 et Opus 5 à 50,8). Évalué en externe par Frontier Design et METR |
| Prix API | **$4 / $20** par MTok (−20 % vs Opus 5 ; ~−40 % sur charge typique selon gangstaai) |
| Déploiement | immédiat sur toutes les plateformes. Positionné par Anthropic comme « le flagship pour le travail sérieux » |
| Safety (System Card, 22/09/2026) | 1,5 % de tentatives de franchissement de frontière sandbox (toutes faible sévérité, lecture seule) ; ~50 % d'actions potentiellement nocives dans un exercice simulé cadré CTF (« quite concerning » selon Anthropic) ; **régressions** : suit plus facilement les instructions malveillantes collées, accepte plus souvent des revendications d'autorisation non vérifiables ; rares cas de tool calls malveillants spontanés sur des snapshots pré-release |

## 48. Claude Sonnet 4.5 — fiche ✅ (hors période, fin de vie imminente)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Sonnet 4.5 / `claude-sonnet-4-5-20250929` |
| Sortie | **29/09/2025** (hors période, inclus pour la trajectoire) |
| Statut | superseded par Sonnet 4.6 (17/02/2026) ; **EOL le 29/09/2026** (2 jours après la date de référence) |
| Contexte | **200K standard**, 1 M en bêta (tarif long-contexte : $6/$22,50 au-delà de 200K) ; sortie max 16 384 (pilot-shell) — ⚠️ conflit : magica.com indique 64K, non tranché |
| Knowledge cutoff | 31/01/2025 |
| Features | vision, tool use, computer use, extended thinking, context awareness (suit le remplissage de sa fenêtre). Optimisé code + agents |
| Fine-tuning | non offert |
| Poids | fermés. API, claude.ai, Bedrock, Vertex AI |
| KV cache | prompt caching (cache read $0,30/M, write $3,75/M 5 min / $6,00/M 1 h) ; seuil min 1 024 tokens |
| Benchmarks clés | **SWE-bench Verified ~77,2 %** (extended thinking), Terminal-Bench 2 50,0 % (vs 43,8 % GPT-5), GPQA Diamond 83,4 % ; LMArena 1370 |
| Prix API | **$3 / $15** par MTok |

## 49. Claude Sonnet 4.6 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Sonnet 4.6 / `claude-sonnet-4-6` |
| Sortie | **17/02/2026** |
| Statut | superseded par Sonnet 5 (30/06/2026) |
| Contexte | 200K standard, **1 M en bêta** (tarif 1 M : $6/$22,50) ; sortie max **non vérifiée précisément au 27/09/2026** |
| Knowledge cutoff | août 2025 |
| Features | adaptive + extended thinking, context compaction (bêta, Developer Platform), computer use (94 % sur un benchmark assurance ; niveau « humain » sur tableurs/formulaires web), résistance au prompt injection très améliorée, design visuel |
| Fine-tuning | non offert |
| Poids | fermés. API, claude.ai (défaut Free/Pro à la sortie), Claude Cowork, Bedrock |
| KV cache | prompt caching ; seuil min 1 024 tokens ; option TTL 1 h |
| Benchmarks clés | **SWE-bench Verified 77,2 %** ; préféré à Sonnet 4.5 ~70 % du temps en code, et à Opus 4.5 59 % du temps ; égale Opus 4.6 sur OfficeQA ; +15 pts vs Sonnet 4.5 en Q&A raisonnement complexe |
| Prix API | **$3 / $15** (inchangé vs Sonnet 4.5) |

## 50. Claude Sonnet 5 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Sonnet 5 / `claude-sonnet-5` |
| Sortie | **30/06/2026** |
| Statut | disponible (GA) ; **défaut des plans Free et Pro** ; défaut Claude Code (Pro) |
| Contexte | **1 M tokens** (seule taille, pas de variante réduite, pas de surcharge) ; sortie max **128K** (300K via bêta Batch API, header `output-300k-2026-03-24`) |
| Knowledge cutoff | janvier 2026 |
| Features | « most agentic Sonnet yet » : planification multi-étapes, tool use (navigateur, terminal), exécution autonome ; **adaptive thinking par défaut** (désactivable explicitement), 5 niveaux d'effort (low/medium/high/xhigh/max, défaut **high**) ; **temperature/top_p/top_k rejetés** ; vision ; computer use (`computer_20251124`, jusqu'à 3,75 MP) ; **nouveau tokenizer** (1,0–1,35× tokens vs 4.6 — piège de coût à la migration) |
| Fine-tuning | non offert (« no fine-tuning access ») |
| Poids | fermés. API, Claude Code, Claude Platform, AWS Bedrock, Google Cloud, Microsoft Foundry (preview), plans consumer |
| KV cache | prompt caching (jusqu'à −90 % lectures) ; seuil min 1 024 tokens ; batch API −50 % |
| Benchmarks clés | **SWE-bench Pro 63,2 %** (vs 58,1 % Sonnet 4.6 ; 69,2 % Opus 4.8), Terminal-Bench 2.1 80,4 % (vs 67,0 %), **OSWorld-Verified 81,2 %** (computer use, vs 78,5 %), Humanity's Last Exam **57,4 %** avec outils (vs 57,9 % Opus 4.8) |
| Prix API | lancement **$2 / $10** jusqu'au 31/08/2026 puis **$3 / $15** annoncés. ⚠️ Conflit non tranché : une source (epic-skills, 12/08/2026) affirme qu'Anthropic a **annulé la hausse** et maintenu $2/$10 — non corroboré ailleurs |

## 51. Claude Sonnet 5.5 — ANNONCÉ, non sorti ⏳

Annoncé le **22/09/2026** lors du lancement d'Opus 5.5 (« dans les semaines à venir ») ; **canary test repéré ~25/09/2026** (1 M contexte, 128K sortie, prix « aligné sur GPT-6 Sol » — source : rebabel.net citant ClaudeDevs). Tout le reste : **non vérifié au 27/09/2026**.

## 52. « Claude Sonnet 4.8 » — NON TROUVÉ ❌

**Verdict : nom inexistant au 27/09/2026** (pas un nom de code connu, probable erreur ou confusion avec Opus 4.8). Requêtes tentées : (1) `Claude Sonnet 4.8 release Anthropic model` ; (2) `"Sonnet 4.8" Anthropic Claude model version`. Preuves convergentes :
- hapa1i/multi-forge (référence modèles, mai 2026) : « **Claude 4.8 currently means Claude Opus 4.8. As of May 2026, Anthropic has not released Sonnet 4.8 or Haiku 4.8.** »
- iWeaver AI (25/05/2026) : « Anthropic has not officially announced Claude Sonnet 4.8 » — cible de rumeurs uniquement.
- Catalogue Vertex AI vérifié empiriquement (codeeagle) : Opus 4.5→4.8 + 5, **Sonnet 4.5, 4.6 et 5**, Haiku 4.5, deux builds Fable — **pas de Sonnet 4.8**.
- Plugin Genkit (Anthropic) : pas de Sonnet 4.8 dans les modèles supportés.
- La lignée Sonnet réelle est : **4.5 (09/2025) → 4.6 (02/2026) → 5 (06/2026)**. Il n'y a jamais eu de 4.7/4.8 côté Sonnet.

## 53. Claude Haiku 4.5 — fiche ✅ (toujours le Haiku courant)

