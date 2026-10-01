---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-4
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-16", "2026-02-24", "2026-03-02", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-09-27"]
keywords: ["agent", "agents", "apache", "attention", "benchmarks", "claude", "deepseek", "exploit", "fable 5", "gguf", "glm", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [302, 397]
sha256: da681f1a45ee8843e63f80196cf7559bdcfafb7475c4bf1d74b45a4d18fb230c
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | 16/02/2026 | MoE, GDN + MoE | 397B / 17B | ~262K (non vérifié) | non vérifié (−60 % vs préd.) | Apache 2.0 |
| Qwen3.5-122B-A10B | 24/02/2026 | MoE, GDN + MoE | 122B / 10B | ~262K (non vérifié) | non vérifié | Apache 2.0 |
| Qwen3.5-35B-A3B | 24/02/2026 | MoE, GDN + MoE | 35B / 3B | ~262K (non vérifié) | non vérifié | Apache 2.0 |
| Qwen3.5-27B | 24/02/2026 | Dense, hybride GDN | 27B | non vérifié | non vérifié | Apache 2.0 |
| Qwen3.5 (0.8B–9B) | 02/03/2026 | Denses | 0.8–9B | non vérifié | non vérifié | Apache 2.0 |
| Qwen3.6-35B-A3B | 16/04/2026 | MoE, GDN + Gated Attn | 35B / 3B | 256K (>1M YaRN) | ~0,14 / 1,00 (tiers) | Apache 2.0 |
| Qwen3.6-27B | 22/04/2026 | Dense, hybride GDN | 27B | 256K (tiers) | non vérifié | Apache 2.0 |
| Qwen3.6-Max-Preview | 20/04/2026 | MoE (presse) | 35B / 3B (presse) | 256K | non vérifié | Fermé |
| Qwen3.6-Plus | 02/04/2026 | non vérifié | non vérifié | 1M | ~0,29 / ~1,74 (Bailian) | Fermé |

## 25. Qwen3.7-Max (fermé, l'agent natif)

- **Nom exact** : Qwen3.7-Max (previews : Qwen3.7-Max-Preview). **Preview : 19 mai 2026**, annonce
  officielle **20 mai 2026** (Alibaba Cloud Summit, Hangzhou) — les deux previews (Max et Plus) étaient
  apparues silencieusement sur le leaderboard LMArena avant l'annonce.
- **Statut** : disponible, **fermé, API uniquement**. Architecture : non vérifié (paramètres total/actifs
  non publiés par Alibaba). **Contexte : 1M tokens**.
- **Fonctionnalités** : modèle **agent natif** — conçu pour l'exécution autonome longue, pas un chat avec
  tool calling greffé : tâches agentiques jusqu'à **35 heures** en continu ; reasoning à chaîne de pensée
  étendue, toujours active ; **texte uniquement** (pas de vision — voir 3.7-Plus) ; généralisation
  cross-scaffold (Claude Code, OpenClaw, Qwen Code) ; multilingue (55+ langues, WMT24++ 85,8).
- **Benchmarks** : SWE-bench Verified **80,4** (≈ Claude Opus-4.6 Max 80,8 ; DeepSeek-v4-Pro Max 80,6),
  Terminal-Bench 2.0-Terminus **69,7** (devant DeepSeek-v4-pro-Max 67,9), GPQA Diamond **92,4** (devant
  Opus-4.6 : 91,3), HLE **41,4** (devant Opus-4.6 : 40,0), MCP-Mark **60,8** (devant GLM-5.1 : 57,5),
  MCP-Atlas **76,4** (devant Opus-4.6 : 75,8), SpreadSheetBench-v1 **87**. Artificial Analysis Intelligence
  Index ≈ **56,6–57** (5e mondial). LMArena Text #13 mondial.
- **Prix API** : à partir de $1,25 / 1M input, $0,25 / 1M cached input, $3,75 / 1M output (via Novita, prix
  le plus bas suivi) — prix officiel Alibaba non vérifié.
- **Déploiement** : API (Model Studio/QwenCloud) ; tiers : Novita, DeepInfra, Together.
- **Note hardware** : lancé avec la puce **Zhenwu M890** et le serveur **Panjiu AL128** (TechTimes) — Qwen3.7-Max
  « a écrit son propre logiciel de puce » pendant un run de 35 heures selon la presse tech, anecdote qui a
  fait le tour des médias.

## 26. Qwen3.7-Plus (fermé, le multimodal)

- **Nom exact** : Qwen3.7-Plus (preview : Qwen3.7-Plus-Preview). **Sortie : 20 mai 2026** (avec Max, au
  Summit) — API grand public quelques semaines après (juin 2026 d'après VentureBeat).
- **Statut** : disponible, **fermé, API uniquement**. Architecture : non vérifié. Contexte : non vérifié
  (probablement 1M comme le reste de la gamme — non vérifié).
- **Fonctionnalités** : **multimodal** (texte, vidéo, images) — le pendant vision du 3.7-Max texte-only ;
  version « équilibrée » axée raisonnement.
- **Poids** : fermés. **Benchmarks** : LMArena Vision #16 mondial (preview) — autres chiffres non vérifiés.
- **Prix API** : **$0,40 / 1M input, $1,60 / 1M output** (VentureBeat) — ~60 % moins cher que Qwen3.7-Max.

## 27. Qwen3.8-Max (2,4T paramètres, retour à l'open-weight)

- **Nom exact** : Qwen3.8-Max (snapshot mis à jour : **Qwen3.8-Max-0902**, 2 septembre 2026).
- **Preview Arena.AI : 19 juillet 2026** ; lancement officiel **3 août 2026** ; snapshot 0902 le 2 septembre 2026.
- **Statut** : disponible — **poids OUVERTS** (retour à l'open-weight au flagship : annonce de publication
  des poids la semaine du 10 août 2026 ; décrit comme « open-weight multimodal system » par plusieurs médias).
  Licence exacte des poids : non vérifiée au 27/09/2026 (texte final non lu directement).
- **Architecture** : MoE — **2,4 billions (2.4T) de paramètres total / ~95B actifs** par token (chiffre
  actif cité dans les supports presse d'Alibaba). Nombre de couches / type d'attention : non publiés.
- **Contexte** : **1M tokens**.
- **Fonctionnalités** : multimodal (texte, images, vidéos, gros fichiers), coding autonome longue durée
  (16 jours de dev autonome en test interne), analyse visuelle, `reasoning_effort` (low/medium/**xhigh**
  par défaut), `preserve_thinking` activé par défaut, `/think` et `/no_think` sur les poids ouverts.
- **Benchmarks** : Artificial Analysis Intelligence Index v4.3 : **40** (3 août) → **45** (snapshot 0902) ;
  Code Arena WebDev : 4e (1669) → **1er** (1691) pour 0902 (devant Claude Opus 5 Max) ; Terminal-Bench 2.1 :
  86,6 ; PaperBench 93,0 ; QwenSWEBench 80,7 (chiffres carte constructeur) ; LMArena : 2e Vision Arena,
  5e Text Arena (derrière Anthropic Fable 5).
- **Prix API** : **$2,00 / 1M input, $6,00 / 1M output** (Model Studio, identique pour les deux snapshots) —
  coût/tâche mesuré par AA : $2,67 → $5,41 (le snapshot 0902 parle beaucoup plus : 108k tokens/tâche).
- **Déploiement** : QwenCloud, Alibaba Cloud Model Studio, QwenWork ; plans API : Standard (pay-as-you-go)
  ou Token Plan (pas le Coding Plan).

## 28. Qwen3.8-Max-0902 : le snapshot qui parle trop

Le snapshot du 2 septembre 2026 mérite sa fiche : AA Intelligence Index passé de 40 à 45, 1er de Code
Arena WebDev (1691, devant Claude Opus 5 Max), mais coût par tâche doublé ($2,67 → $5,41) parce que le
modèle génère ~108k tokens par tâche. C'est l'illustration parfaite du dilemme « xhigh reasoning » :
plus de raisonnement = meilleurs scores, mais facture API qui explose. Alibaba laisse `reasoning_effort`
réglable (low/medium/xhigh) pour que le client arbitre lui-même.

## 29. Qwen3.8-27B (dense ouvert, tueuse de géants)

- **Nom exact** : Qwen3.8-27B (27,78B params ; 28B avec l'encodeur vision). **Sortie : 14 août 2026**.
  Statut : disponible, poids ouverts.
- **Architecture** : **dense 27,78B** ; hybride : **48 couches Gated DeltaNet + 16 couches full-attention**.
- **Contexte** : 262 144 natifs, extensible à 1M via YaRN.
- **Fonctionnalités** : multimodal (texte, image, vidéo), thinking activé par défaut (désactivable),
  coding/reasoning/agents long-horizon.
- **Poids** : **ouverts, Apache 2.0**. Ollama : `qwen3.8:27b` (18 Go en Q4_K_M, support day-0, tête MTP
  incluse → décodage spéculatif par défaut) ; GGUF Unsloth (UD-Q4_K_XL 17,6 Go) ; tourne sur 1× GPU 24 Go
  en 4-bit.
- **Inférence** : GDN → état compact ; tête MTP exploitée par Ollama (`draft-mtp`).
- **Benchmarks** (agrégateur regolo.ai, à manier avec précaution) : SWE-bench Pro **61,7** (vs Claude Opus
  4.6 Max 53,4), LiveCodeBench v6 **90,3** (vs 88,8), IFBench **79,5** (vs 62,5), QwenSWEBench 79,0 (vs 63,8),
  CoWorkBench 70,7 (vs 68,2), DeepSWE 1.1 42,2, Arena WebDev 18e/128. Un dense 27B qui bat un flagship
  fermé occidental sur le code : c'est le fait technique le plus disruptif de l'été 2026 côté ouvert.
- **Déploiement** : Ollama (day-0), llama.cpp, vLLM, LM Studio, DGX Spark.

## 30. Qwen3.8-Flash (la « preview de Qwen4 »)

