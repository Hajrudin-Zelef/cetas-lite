---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-6
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02-05", "2026-04-16", "2026-05-28", "2026-06-09", "2026-07-24", "2026-09-22", "2026-09-27"]
keywords: ["agent", "agents", "agi", "asl", "bedrock", "benchmark", "benchmarks", "claude", "cyber", "fable 5", "fine-tuning", "foundry"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [447, 513]
sha256: 2e893eca11e861bbda073380bcebf140921f5407e754194f5f72f4b885f0fe41
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

## 43. Claude Opus 4.6 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Opus 4.6 / `claude-opus-4-6` |
| Sortie | **05/02/2026** (3 mois après Opus 4.5) |
| Statut | disponible (génération précédente) ; fast mode déprécié ~06/2026 |
| Contexte | **1 M tokens (bêta)** — premier Opus à 1 M (5× vs 200K d'Opus 4.5) ; sortie max 128K (doublée vs 64K) |
| Knowledge cutoff | août 2025 |
| Features | **adaptive thinking** (4 niveaux : low/medium/high/max — le modèle décide quand réfléchir), **agent teams** dans Claude Code (multi-instances collaboratives, lead agent + peer-to-peer), context compaction (bêta), fast mode (research preview, 2,5×). Acceptait encore `{type:"enabled", budget_tokens}` (déprécié) et les sampling params |
| Fine-tuning | non offert |
| Poids | fermés. API, claude.ai, Bedrock, Vertex AI, Microsoft Foundry |
| KV cache | prompt caching ; seuil min 4 096 tokens. Tarif premium long-contexte : **$10/$37,50** au-delà de 200K tokens |
| Benchmarks clés | **SWE-bench Verified 80,8 %**, ARC-AGI-2 **68,8 %** (SOTA fév. 2026), ARC-AGI-1 94,0 %, GPQA Diamond 91,3 %, Terminal-Bench 2.0 65,4 %, OSWorld-Verified 72,7 %, CyberGym 66,6 %, τ2-bench retail 91,9 % / telecom 99,3 %, MRCR long-contexte **76 % à 1 M** (vs 18,5 % Opus 4.5). Fait marquant : **500+ zero-days découverts** en open-source pendant les tests pré-release |
| Prix API | **$5 / $25** par MTok |
| Sécurité | déploiement ASL-3 (AI Safety Level 3) |

## 44. Claude Opus 4.7 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Opus 4.7 / `claude-opus-4-7` |
| Sortie | **16/04/2026** |
| Statut | disponible (génération précédente) ; remplacé par 4.8 |
| Contexte | **1 M tokens** ; sortie max 128K |
| Knowledge cutoff | janvier 2026 |
| Features | vision **3× résolution** (jusqu'à 2 576 px grand côté, ~3,75 MP), **nouveau niveau d'effort `xhigh`** (entre high et max), auto-vérification (écrit et exécute des tests), task budgets (bêta publique), **nouveau tokenizer** (1,0–1,35× tokens — controverse : +35–47 % de coût réel mesuré par la communauté malgré « prix inchangé »), garde-fous cybersécurité intégrés + **Cyber Verification Program**. Sampling params supprimés |
| Fine-tuning | non offert |
| Poids | fermés. API, Claude Code, Bedrock, Vertex AI, Microsoft Foundry |
| KV cache | prompt caching ; seuil min 2 048 tokens |
| Benchmarks clés | **SWE-bench Verified 87,6 %** (vs 80,8 %), **SWE-bench Pro 64,3 %** (vs 53,4 %), CursorBench 70 % (vs 58 %), MCP-Atlas +14,6 pts, navigation visuelle 57,7→79,5 %. Devant GPT-5.4 et Gemini 3.1 Pro en code agentique |
| Prix API | **$5 / $25**. ⚠️ Conflit : Neowin annonce $15/$75 — contredit par toutes les autres sources ; probable erreur Neowin, écarté |

## 45. Claude Opus 4.8 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Opus 4.8 / `claude-opus-4-8` |
| Sortie | **28/05/2026** (41 jours après Opus 4.7 ; annonce synchronisée avec la clôture de la Series H) |
| Statut | disponible ; n'est plus le flagship depuis Fable 5 (09/06/2026) mais reste le **modèle de fallback** des garde-fous de Fable 5/5.1 et d'Opus 5 (< 5 % des sessions) |
| Contexte | **1 M tokens** ; sortie max **128K** (300K en bêta Batch API) |
| Knowledge cutoff | janvier 2026 |
| Features | **Dynamic Workflows** (research preview, Claude Code : centaines de sous-agents parallèles par session), **effort control** (low/medium/high/xhigh/max ; xhigh hérité de 4.7), **mid-conversation system messages** (sans bêta header, préserve le cache), task budgets (bêta), adaptive thinking (désactivable explicitement), **nouveau tokenizer** (1,0–1,35× tokens vs 4.6). Sampling params supprimés. Modèle « 4× moins susceptible de laisser passer sans le signaler un défaut de son propre code » |
| Fine-tuning | non offert |
| Poids | fermés. API, claude.ai, Bedrock, Vertex AI. Fast mode : 2,5× plus rapide, **3× moins cher** que le fast mode précédent |
| KV cache | prompt caching standard ; seuil min 1 024 tokens |
| Benchmarks clés | **SWE-bench Pro 69,2 %** (vs 64,3 % pour 4.7), **SWE-bench Verified 88,6 %**, **USAMO 2026 96,7 %** (vs 69,3 %), Online-Mind2Web 84 %, OSWorld 83,4 % (computer use), Legal Agent Benchmark (Sierra) : premier > 10 % all-pass ; Artificial Analysis Intelligence Index 61,4 (1er à la sortie) |
| Prix API | **$5 / $25** (inchangé vs 4.7) ; fast mode **$10/$50** |

## 46. Claude Opus 5 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | Claude Opus 5 / `claude-opus-5` |
| Sortie | **24/07/2026** (annonce officielle lue en lecture directe : anthropic.com/news/claude-opus-5) |
| Statut | disponible. Nouveau défaut sur **Claude Max** ; modèle le plus fort sur Claude Pro ; défaut Opus de Claude Code. Passé en « legacy » (remplacé-par-5.5 recommandé) le 22/09/2026 |
| Contexte | **1 M tokens** ; sortie max **128K** |
| Knowledge cutoff | mai 2026 |
| Features | adaptive thinking + réglages `effort` (arbitrage vitesse/tokens/capacité) + **mode rapide à prix supérieur** ; vision, tool use, streaming ; computer use ; MCP natif ; vérification autonome renforcée. Température/top_p/top_k : rejetés. Annonce officielle : « comes close to the frontier intelligence of Claude Fable 5 at half the price » |
| Fine-tuning | non offert |
| Poids | fermés. API, Bedrock, Vertex AI, Databricks, Snowflake, Claude Code (défaut Opus) |
| KV cache | prompt caching standard ; seuil min 512 tokens |
| Benchmarks clés (annonce officielle) | Frontier-Bench v0.1 **SOTA** (2× Opus 4.8 à moindre coût/tâche) ; CursorBench 3.2 à < 0,5 % du pic Fable 5 (max effort, moitié du coût) ; ARC-AGI 3 : **3× le 2e meilleur modèle** ; Zapier AutomationBench 1,5× le taux de réussite ; OSWorld 2.0 : bat Fable 5 à ~1/3 du coût. Reste **derrière Mythos 5** en cybersécurité/biologie (OSS-Fuzz : proche en *découverte* de vulnérabilités, loin derrière en *exploitation*) |
| Prix API | **$5 / $25** (inchangé vs Opus 4.8). ⚠️ Conflit : une vidéo YouTube annonce $10/$50, 500K de contexte et août 2026 — contredit par 2+ sources écrites ; écartée |

## 47. Claude Opus 5.5 — fiche ✅ (flagship « travail sérieux » au 27/09/2026)

