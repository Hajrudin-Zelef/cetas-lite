---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-8
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Microsoft", "OpenRouter"]
dates: ["2025-04-11", "2025-09-29", "2025-10-15", "2025-11-24", "2026-02-05", "2026-02-17", "2026-03-05", "2026-03-26", "2026-04-07", "2026-04-16", "2026-05", "2026-05-28", "2026-06-09", "2026-06-15", "2026-06-30", "2026-07-24", "2026-08-31", "2026-09-01", "2026-09-22", "2026-09-27", "2026-09-29", "2026-10-15", "2026-11-24"]
keywords: ["agent", "agentic", "aws", "bedrock", "benchmarks", "claude", "consumer", "fable 5", "fine-tuning", "foundry", "gemini", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [592, 691]
sha256: 36bdc73edd1a556a34e0860ce99f8b6f8b488cc6c18e5c103f4cc7931301c060
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Haiku 4.5 / `claude-haiku-4-5-20251001` |
| Sortie | **15/10/2025** (hors période, inclus pour la trajectoire) |
| Statut | **toujours le Haiku courant au 27/09/2026** (rôle : routeur rapide / filtre basse latence) ; EOL 15/10/2026 |
| Contexte | **200K tokens** ; sortie max **64K** |
| Knowledge cutoff | février 2025 |
| Features | modèle rapide/économique ; **extended thinking manuel uniquement** (pas d'adaptive thinking, pas de paramètre `effort`) ; vision, tool use |
| Fine-tuning | non offert sur la gamme 2026 (seul le vieux Claude 3.5 Haiku avait un programme de fine-tuning, via Bedrock uniquement) |
| Poids | fermés. API, Bedrock, Vertex AI, plans consumer |
| KV cache | prompt caching ; seuil min **4 096 tokens** |
| Benchmarks clés | 91–104 tokens/s ; benchmarks de capacité **non vérifiés au 27/09/2026** |
| Prix API | **$1 / $5** par MTok |

## 54. « Claude Haiku 5 » — NON TROUVÉ ❌

**Verdict : la génération « 5 » n'existe pas pour Haiku au 27/09/2026** (probable erreur ; la famille saute de 4.5 vers 5.5 annoncé). Requêtes tentées : (1) `Claude Haiku latest version 2026 release Haiku 5` ; (2) `"Claude Haiku 5" model Anthropic`. Preuve empirique : codeeagle a testé `claude-haiku-5` contre le catalogue : « **does not exist — the catalogue rejects it. Haiku 4.5 is the current Haiku** » ; le catalogue Vertex liste Haiku 4.5 uniquement. Note : quelques docs de SDK mentionnent `claude-haiku-5` dans du **code d'exemple** — usage spéculatif/aspirationnel, pas une preuve d'existence.

## 55. Claude Haiku 5.5 — ANNONCÉ, non sorti ⏳

Annoncé le **22/09/2026** avec Opus 5.5 (« dans les semaines à venir ») ; Anthropic a formellement confirmé qu'il **n'est pas annulé** (démenti des rumeurs). Aucune spec officielle, aucune date ferme → tout le reste **non vérifié au 27/09/2026**.

## 56. Adaptive thinking et effort : le récapitulatif vérifié

- Échelle : **low / medium / high (défaut) / xhigh / max** ; `xhigh` introduit par Opus 4.7, disponible sur Opus 4.7+, Sonnet 5, Fable 5.
- Opus 5.5 : thinking **obligatoire** (plus d'opt-out). Fable 5/5.1 : thinking **toujours actif**.
- Sampling params (temperature/top_p/top_k) : **rejetés** sur Sonnet 5, Opus 4.7/4.8, Fable 5 ; encore acceptés sur 4.6.
- Haiku 4.5 : extended thinking **manuel uniquement**, pas d'adaptive thinking.

## 57. Prompt caching Anthropic (vérifié par 4+ sources concordantes)

- Jusqu'à **4 breakpoints `cache_control`** par requête.
- **Écriture cache : 1,25×** le prix input (TTL 5 min, défaut) ; **2×** (TTL 1 h, via header `anthropic-beta: extended-cache-ttl-2025-04-11`).
- **Lecture cache : 0,1×** (−90 %) — **sauf Fable 5.1 / Mythos 5.1 : 0,025× (0,25 $/M)**.
- TTL **glissant** (rafraîchi à chaque lecture). Les lectures de cache ne comptent pas dans l'utilisation rate-limit.
- Seuils minimums par modèle : 512 (Opus 5, Fable 5/5.1, Mythos 5/5.1) · 1 024 (Opus 4.8, Sonnet 5/4.6/4.5…) · 2 048 (Mythos Preview, Opus 4.7) · 4 096 (Opus 4.6/4.5, Haiku 4.5).

## 58. Fine-tuning : absent en 2026 (vérifié comme absent)

**Pas de programme de fine-tuning public Anthropic sur la gamme 2026** (« Anthropic doesn't offer fine-tuning as of May 2026 » ; « Fine-Tuning: Not available » ; « no fine-tuning access » sur Sonnet 5). Seule exception connue : **Claude 3.5 Haiku** finetunable via **AWS Bedrock** (Oregon) — programme hérité, hors gamme 2026. Une source faible évoque du « custom fine-tuning » enterprise — non corroboré, **non vérifié au 27/09/2026**.

## 59. Poids fermés et canaux de déploiement Claude

**Tous les modèles : poids fermés, licence propriétaire, API-only.** Aucun self-hosting. Canaux : API Anthropic, **AWS Bedrock**, **Google Vertex AI / Agent Platform**, **Microsoft Foundry / Azure**, claude.ai, Claude Code, Claude Cowork ; ponctuellement Databricks, Snowflake, OpenRouter (Fable). Alias par défaut (API directe) : `opus` → Opus 4.8, `sonnet` → Sonnet 5 (sept 2026). Zero Data Retention : défaut sur Bedrock (Opus 5.5) ; « Enterprise Frontier Safeguards » en déploiement phasé automne 2026 pour Fable 5/5.1. Retraits : Sonnet 4 + Opus 4 originel retirés le 15/06/2026 ; EOL : Sonnet 4.5 → 29/09/2026, Haiku 4.5 → 15/10/2026, Opus 4.5 → 24/11/2026.

## 60. Conflits entre sources non tranchés (Anthropic)

1. **Prix Sonnet 5 post-31/08/2026** : $3/$15 (annoncé, 4+ sources) vs maintien à $2/$10 (epic-skills seul). → Non tranché.
2. **Sortie max Sonnet 4.5** : 16 384 (pilot-shell) vs 64K (magica). → Non tranché.
3. **Position Opus 5.5 vs Fable 5.1 sur l'Artificial Analysis Index** : v4.3 donne Fable 5.1 ex æquo 1er à 53 ; v4.3.2 donne Opus 5.5 1er à 57,6, Fable 5.1 à 53,4. → Les deux reportées telles quelles (versions d'index différentes).
4. **Fuite Mythos** : 26/03/2026 (CMS, Fortune — majoritaire) vs 05/03/2026 (S3, une source). → Retenu : 26/03 comme date de révélation publique.
5. **Nom de code Mythos** : Capybara (majoritaire) vs Mythopoeia (webpronews/BI). → Retenu : Capybara.

## 61. Tableau récapitulatif Claude (sept. 2026)

| Modèle | Sortie | Statut | Contexte | Prix in/out par MTok | Notes |
|---|---|---|---|---|---|
| Mythos | — (fuite 26/03/2026) | **restreint, jamais public** | non vérifié | $25/$125 (source unique) | Project Glasswing |
| Mythos 5 | 09/06/2026 | restreint | non vérifié | non vérifié | = Fable 5 sans garde-fous |
| Mythos 5.1 | 01/09/2026 | restreint | non vérifié | non vérifié | = Fable 5.1 sans garde-fous |
| Fable 5 | 09/06/2026 | dispo | 1M | $10/$50 | tier au-dessus d'Opus |
| Fable 5.1 | 01/09/2026 | dispo | 1M | $10/$50 (+cache $0,25/M) | le plus intelligent public |
| Opus 4.6 | 05/02/2026 | dispo | 1M | $5/$25 | adaptive thinking |
| Opus 4.7 | 16/04/2026 | dispo | 1M | $5/$25 | effort xhigh |
| Opus 4.8 | 28/05/2026 | dispo | 1M | $5/$25 | fallback garde-fous |
| Opus 5 | 24/07/2026 | dispo (legacy) | 1M | $5/$25 | défaut Claude Max |
| **Opus 5.5** | **22/09/2026** | **dispo — flagship** | **non vérifié** | **$4/$20** | thinking obligatoire |
| Sonnet 4.5 | 29/09/2025 | EOL 29/09/2026 | 200K (1M bêta) | $3/$15 | — |
| Sonnet 4.6 | 17/02/2026 | superseded | 200K (1M bêta) | $3/$15 | — |
| Sonnet 5 | 30/06/2026 | dispo (défaut Free/Pro) | 1M | $2/$10 → $3/$15 ? | most agentic Sonnet |
| Sonnet 5.5 | annoncé 22/09/2026 | **non sorti** | non vérifié | non vérifié | canary ~25/09 |
| « Sonnet 4.8 » | — | **N'EXISTE PAS** | — | — | confusion avec Opus 4.8 |
| Haiku 4.5 | 15/10/2025 | dispo (EOL 15/10/2026) | 200K | $1/$5 | routeur rapide |
| « Haiku 5 » | — | **N'EXISTE PAS** | — | — | catalogue le rejette |
| Haiku 5.5 | annoncé 22/09/2026 | **non sorti** | non vérifié | non vérifié | confirmé non annulé |

## 62. Timeline Claude (sept. 2025 → sept. 2026)

| Date | Modèle | Statut au 27/09/2026 |
|---|---|---|
| 2025-09-29 | Sonnet 4.5 | EOL 2026-09-29 |
| 2025-10-15 | Haiku 4.5 | courant ; EOL 2026-10-15 |
| 2025-11-24 | Opus 4.5 | EOL 2026-11-24 |
| 2026-02-05 | Opus 4.6 | superseded |
| 2026-02-17 | Sonnet 4.6 | superseded (remplacé par Sonnet 5 le 30/06) |
| ~03/2026 | Mythos (fuite) | jamais public ; nom de code Capybara |
| 2026-04-07 | Mythos Preview | accès restreint (Glasswing) |
| 2026-04-16 | Opus 4.7 | superseded |
| 2026-05-28 | Opus 4.8 | dispo ; fallback garde-fous |
| 2026-06-09 | Fable 5 (+ Mythos 5 restreint) | dispo (GA) |
| 2026-06-30 | Sonnet 5 | dispo ; défaut Free/Pro |
| 2026-07-24 | Opus 5 | dispo ; défaut Claude Max |
| 2026-09-01 | Fable 5.1 (+ Mythos 5.1 restreint) | dispo (GA) |
| 2026-09-22 | Opus 5.5 | dispo (GA) ; System Card publiée |
| (à venir) | Sonnet 5.5 / Haiku 5.5 | **annoncés** le 22/09/2026, « dans les semaines à venir » |

---

# PARTIE IV — Google : la série Gemini 3.x (texte), Nano Banana (image), Veo (vidéo)

