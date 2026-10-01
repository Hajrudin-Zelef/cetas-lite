---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-5
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Broadcom", "Glasswing", "Google", "Meta", "Microsoft", "Nvidia", "OpenRouter", "United States"]
dates: ["2026-03-05", "2026-03-26", "2026-04-08", "2026-06-09", "2026-07-01", "2026-07-09", "2026-07-20", "2026-07-24", "2026-08-05", "2026-08-10", "2026-09-01", "2026-09-02", "2026-09-08", "2026-09-27"]
keywords: ["agent", "agents", "aws", "bedrock", "benchmarks", "claude", "cyber", "fable 5", "fine-tuning", "kv cache", "muse", "muse spark"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [367, 446]
sha256: 4422ec118b91c002b40d417386841d54aad230cd8b07c7cfe2cb3bb2816086d2
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Modèle / produit | Sortie | Statut | Contexte | Prix | Notes |
|---|---|---|---|---|---|
| Muse Spark 1.0 | 08/04/2026 | dispo | non vérifié | — | Contemplating 58 % HLE |
| Muse Spark 1.1 | 09/07/2026 | dispo | 1M | $1,25/$4,25 (Contributor $0,10/$0,20) | Meta Model API |
| Muse Spark 1.2 | 05/08/2026 | dispo | 1M | $1,25/$4,25 | + Muse Code bêta |
| **Muse Spark 1.3** | **02/09/2026** | **dispo — actuel** | 1M | $1,25/$4,25 | propulse **Muse** |
| Muse Glimmer 30B | 10/08/2026 | **source unique** | — | — | open-weights supposé |
| « Watermelon » | — | **non vérifié** | — | — | nom de code interne |
| Muse (agent) | 08/09/2026 | dispo (US only) | — | gratuit ≤100M tok/sem ; $20/$100/mois | Secure VM + Sentinel |

---

# PARTIE III — Anthropic : la famille Claude

## 37. La politique d'Anthropic : aucune architecture publiée

**Anthropic ne publie aucun détail d'architecture pour aucun modèle Claude** (paramètres, type de réseau, données d'entraînement) — confirmé par l'ensemble des sources (« param_count: Undisclosed », « no weights, no parameter count, no architecture paper »). Seules exceptions partielles : les **System Cards** publiées pour Opus 4.8, Sonnet 5, Fable 5.1 et Opus 5.5 (évaluations comportementales et de capacités, pas d'architecture). Toutes les fiches ci-dessous portent donc « architecture : non publiée » sans le répéter.

## 38. Claude Mythos (nom de code Capybara) — restreint, jamais public ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Claude Mythos (nom de code interne : **Capybara** ; une source cite aussi « Mythopoeia » — conflit non tranché) |
| Sortie | **pas de sortie publique**. Fuite le **26/03/2026** (mauvaise config du CMS d'Anthropic : ~3 000 assets non publiés exposés, dont un brouillon d'annonce ; découvert par Roy Paz, LayerX Security ; histoire révélée par Fortune). Une source secondaire parle d'une fuite dès le 05/03/2026 (bucket S3) — incohérence de dates entre sources. |
| Statut | **modèle interne restreint, jamais commercialisé publiquement** — « by far the most powerful AI model we've ever developed » (brouillon fuité) / « a step change » (porte-parole, confirmé à Fortune) |
| Pourquoi restreint | premier modèle de l'histoire volontairement non publié pour potentiel cyber destructeur. Capacités revendiquées : découverte autonome de milliers de zero-days (faille TCP SACK OpenBSD de 27 ans, bug FFmpeg H.264 de 16 ans, RCE NFS FreeBSD CVE-2026-4747) ; taux de succès d'exploitation autonome de 72,4 % (vs ~0 % pour Opus 4.6) selon une analyse tierce du preview |
| Features | non vérifié dans le détail ; positionné cybersécurité + agents autonomes longue-horizon |
| Poids | fermés ; accès API restreint via **Project Glasswing** (~40 organisations vérifiées : Amazon, Microsoft, Google, Apple, CrowdStrike, Cisco, JPMorgan Chase, Linux Foundation, Nvidia, Broadcom, Palo Alto Networks), sous NDA (« Responsible Deployment Protocol v3 » selon une source). Dotation : **100 M$ de crédits d'usage + 4 M$ de dons** aux orgs de sécurité open-source |
| Prix API | une seule source secondaire (FAQ pilot-shell) indique **$25 / $125** par MTok pour Mythos Preview — **à traiter avec prudence (source unique)** |
| Benchmarks | revendications non chiffrées publiquement (« step change ») ; le brouillon fuité le place « dramatically higher than Claude Opus 4.6 » en code, raisonnement académique et cybersécurité. L'annonce officielle d'Opus 5 (24/07/2026) confirme que Mythos 5 reste **devant** Opus 5 en biologie et cybersécurité offensive |

## 39. Claude Mythos 5 — même modèle que Fable 5, sans les garde-fous ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Claude Mythos 5 — **même modèle sous-jacent que Claude Fable 5**, sans les garde-fous cyber/bio |
| Sortie | 09/06/2026 (annonce conjointe avec Fable 5) |
| Statut | disponible uniquement pour partenaires Project Glasswing + chercheurs biomédicaux sélectionnés |
| Prix API | **non vérifié au 27/09/2026** |

## 40. Claude Mythos 5.1 — restreint ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Claude Mythos 5.1 — même modèle sous-jacent que Fable 5.1, garde-fous assouplis |
| Sortie | 01/09/2026 |
| Statut | restreint à des organisations US sélectionnées (d'abord sciences de la vie ; accès cybersécurité élargi prévu) |

## 41. Claude Fable 5 — fiche ✅ (tier au-dessus d'Opus)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Fable 5 / `claude-fable-5` — « premier modèle public de classe Mythos », tier **au-dessus d'Opus** |
| Sortie | **09/06/2026** (veille de « Code with Claude Tokyo »). Suspension temporaire du 12/06 au 01/07/2026 suite à une directive US de contrôle des exportations, puis redéploiement global |
| Statut | disponible (GA) |
| Contexte | **1 M tokens** (tarif standard, sans surcharge long-contexte) ; sortie max **128K** |
| Features | vision (texte+image), tool use, reasoning (toujours actif ; seul `{type:"adaptive"}` accepté — `disabled` et `budget_tokens` rejetés en 400), adaptive thinking, effort `xhigh` supporté, prompt caching (seuil min 512 tokens). Garde-fous : les requêtes cyber/bio/chimie sensibles **basculent sur Opus 4.8** (< 5 % des sessions) au lieu d'être refusées |
| Fine-tuning | non offert (aucun programme de fine-tuning public Anthropic) |
| Poids | fermés, propriétaires. API Anthropic, claude.ai, AWS Bedrock, Google Vertex AI, OpenRouter. Claude Code : opt-in (`/model fable`), jamais défaut |
| Benchmarks clés | **SWE-bench Verified 95,0 %**, **SWE-bench Pro 80,3 %**, **GPQA Diamond 92,6 %** (sources : best-of-ai, ai-whatchelin) ; LMArena 1440, MMLU 93,0 (best-of-ai, à prendre avec prudence) |
| Prix API | **$10 / $50** par MTok — le double d'Opus 4.8 |
| Déploiement | plans Max (inclus depuis 20/07/2026, plafonné à 50 % des limites hebdo), Team/Enterprise (sièges premium), Pro (crédits d'usage) ; jamais sur le tier gratuit claude.ai |

## 42. Claude Fable 5.1 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Fable 5.1 / `claude-fable-5-1` |
| Sortie | **01/09/2026** |
| Statut | disponible (GA). Décrit comme « le modèle Claude le plus intelligent disponible publiquement » |
| Contexte | **1 M tokens** ; sortie max **128K** ; ~69 tokens/s (deeplearning.ai) |
| Knowledge cutoff | juin 2026 |
| Features | reasoning toujours actif, 5 niveaux d'effort (low/medium/high/xhigh/max, défaut high), changement de niveau **en cours de conversation sans invalider le cache** (bêta), tool use avec mises à jour de progression lisibles (bêta), system messages par tour (bêta), prompt caching, **watermarking statistique du texte généré**, fallback optionnel vers Opus 4.8/Opus 5 sur déclenchement des classifieurs bio/cyber. Interventions des garde-fous cyber réduites de 60 % par session Claude Code et de 85 % sur requêtes bio/médicales bénignes vs Fable 5 |
| Fine-tuning | non offert |
| Poids | fermés. API, claude.ai, Bedrock, Vertex AI, OpenRouter, Azure |
| KV cache | **lectures de cache à 0,025× (0,25 $/M, −75 %)** vs 0,1× standard — principale baisse de coût (jusqu'à −45 % sur charges agentiques, ~−25 % en usage typique). Écritures cache : $12,50/M (5 min) / $20/M (1 h) ; batch $5/$25. Seuil min 512 tokens |
| Benchmarks clés | Terminal-Bench-Science **52,6 %** (vs 24,7 % Fable 5), Terminal-Bench 4.0 **55,8 %** (vs 42,0 %), CursorBench **73,4 %** (vs 70,5 %) ; ex æquo premier de l'Artificial Analysis Intelligence Index v4.3 (53) ; Vals Index 68,83 % |
| Prix API | **$10 / $50** (inchangé vs Fable 5) + cache read $0,25/M. ⚠️ **Infox débunkée :** un article annonçait $3/$15 — contredit par 3+ sources concordantes, écarté |
| Déploiement | mêmes canaux que Fable 5. Cognition (Devin) annonce migrer une partie de son trafic Opus 5 → Fable 5.1 |

