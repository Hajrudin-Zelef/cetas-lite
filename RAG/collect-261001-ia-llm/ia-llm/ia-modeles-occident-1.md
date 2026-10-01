---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-1
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Poolside", "United States", "xAI"]
dates: ["2026-09-22", "2026-09-27", "2026-11-21"]
keywords: ["agent", "agents", "agi", "astra", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "fable 5", "fine-tuning"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1, 95]
sha256: 882c8bc20f3b5c20ffa4ec101f6398998986e381bae66e0f4b86d3647fdf88f1
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident
**Période couverte : février 2026 → 27 septembre 2026. Langue : français.**
**Date de référence : 27/09/2026.**

Ce volume couvre les familles occidentales : OpenAI GPT, Meta Muse / Muse Spark, Anthropic Claude, Google Gemini, xAI Grok, Meta Llama 4, Mistral AI, Google Gemma 4, IBM Granite, Poolside, NVIDIA Nemotron et NousResearch Hermes.

> **Règle d'or du volume** : aucune spec n'est inventée. Tout ce qui n'est pas vérifiable est marqué **« non vérifié au 27/09/2026 »**. Les chiffres de benchmarks cités sont des revendications fournisseurs sauf mention contraire. Les modèles recherchés et introuvables sont documentés comme **NON TROUVÉ** avec les requêtes tentées. Les fiches reposent sur une recherche vérifiée menée par huit chercheurs humains/IA le 27/09/2026 ; quand une donnée vient d'une source unique, c'est signalé.

**Conventions utilisées dans ce volume :**
- ✅ RÉEL = modèle vérifié comme existant et disponible (ou sorti puis retiré, si indiqué).
- ❌ NON TROUVÉ = recherché (2–3 angles de recherche), aucune trace.
- « non vérifié au 27/09/2026 » = donné plausible mais non confirmable par les sources consultées.
- Les prix API sont en $US par million de tokens (entrée / sortie), sauf mention contraire.
- « MTok » = million de tokens.

---

# PARTIE I — OpenAI : la famille GPT (5.3 → 6)

## 1. Le système de tiers GPT-5.6 : Sol, Terra, Luna

Avant les fiches, l'essentiel : **Sol, Terra et Luna ne sont ni des noms de code internes ni des erreurs de sources** — ce sont les noms officiels publics des trois tiers de la famille **GPT-5.6** (limited preview 26 juin 2026, GA 9 juillet 2026), reconduits pour la série **GPT-6** (sauf Terra, voir § 8). Logique :
- **Sol** = le tier flagship, capacité maximale, raisonnement `max`/`ultra` réservés, effort de raisonnement le plus élevé.
- **Terra** = le tier équilibré (« balanced/everyday »), prix intermédiaire.
- **Luna** = le tier économique, rapide, pour le volume (remplace la gamme mini/nano).

Cette architecture en tiers a remplacé la vieille nomenclature mini/nano. Les trois tiers partagent ~1 M de contexte ; ce qui les sépare, c'est la capacité, la latence et le prix.

## 2. GPT-5.6 Sol — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-5.6 Sol / `gpt-5.6-sol` |
| Sortie | limited preview **26 juin 2026** ; GA **9 juillet 2026** |
| Statut | Disponible ; remplacé comme flagship par GPT-6 Sol le 22/09/2026, toujours accessible |
| Architecture | **non vérifié au 27/09/2026** — OpenAI ne publie pas (dense/MoE, paramètres total/actifs inconnus) |
| Contexte | **1,05 M tokens**, 128K max output (sources tierces : docs bofai, epic-skills) |
| Knowledge cutoff | février 2026 |
| Features | vision (text + image in), tool use, reasoning réglable (none/low/medium/high/xhigh/**max** — `max` = Sol uniquement), agents (Responses API, computer use, Programmatic Tool Calling, subagents via mode `ultra`), structured outputs, streaming, Batch |
| Fine-tuning | non supporté sur Sol selon les pages de référence OpenAI (endpoints assistants/realtime/live/fine-tuning non supportés pour la famille GPT-6 Sol/Luna ; pour 5.6 Sol non vérifié précisément) |
| Poids | fermés — API uniquement (OpenAI, Azure, AWS Bedrock) |
| KV cache | prompt caching (90 % de remise sur cached input reads) ; cache-write pricing introduit avec la 5.6 ; >272K tokens prompt = tarification long-contexte ×2 in / ×1.5 out |
| Prix API | lancement $5 / $30 → **promo cut le 21 août 2026 → $4 / $20** (cache read $0,40 ; promo tenue jusqu'au moins 21/11/2026) |
| Déploiement | ChatGPT (Plus/Pro/Business/Enterprise), Codex, API OpenAI, Azure, AWS Bedrock |

## 3. GPT-5.6 Terra — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-5.6 Terra / `gpt-5.6-terra` |
| Sortie | GA **9 juillet 2026** (preview 26–27 juin 2026) |
| Statut | Disponible — tier « balanced/everyday » |
| Architecture | **non vérifié au 27/09/2026** |
| Contexte | **1,05 M tokens**, 128K max output ; **knowledge cutoff : 16 février 2026** |
| Features | raisonnement (effort none/low/medium/high/xhigh ; `max`/`ultra` = Sol uniquement ; Terra supporte `reasoning.mode: "pro"`), Responses API, functions, web search, file search, computer use, Programmatic Tool Calling, persisted reasoning, multi-agent bêta |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | fermés |
| Prix API | lancement $2,50 / $15 → **cut −20 % le 30 juillet 2026 → $2 / $12** |
| Déploiement | ChatGPT (Plus/Pro), Codex, API |

## 4. GPT-5.6 Luna — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-5.6 Luna / `gpt-5.6-luna` |
| Sortie | GA **9 juillet 2026** (famille GPT-5.6) |
| Statut | Disponible — tier le moins cher ; remplace la gamme mini/nano |
| Architecture | **non vérifié au 27/09/2026** |
| Contexte | **non vérifié précisément** (famille 5.6 ≈ 1 M ; pas de chiffre publié par OpenAI) |
| Features | tier rapide/économique — résumé, extraction, classification, drafting, automatisation de routine |
| Poids | fermés |
| Prix API | lancement $1 / $6 → **cut −80 % le 30 juillet 2026 → $0,20 / $1,20** (cache read = 1/10e) |
| Déploiement | ChatGPT (Free, Go, Plus), Codex, API |

## 5. GPT-6 Astra — fiche ✅ (flagship au 27/09/2026)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-6 Astra / `gpt-6-astra` |
| Sortie | **3 septembre 2026** (rollout limité organisations) → GA **4 septembre 2026** |
| Statut | Disponible — **flagship actuel d'OpenAI** (remplace GPT-5.6 Sol) |
| Architecture | **non vérifié au 27/09/2026** — OpenAI ne publie pas ; une source tierce (Medium) dit « single dense reasoning model » → **non vérifié** |
| Contexte | **1 050 000 tokens** (plafond input 922 000), **128K max output** |
| Knowledge cutoff | **30 avril 2026** (déclaré par OpenAI) |
| Features | raisonnement (effort low/medium/high/xhigh/max), **computer use natif** (opère le bureau/logiciels comme un humain : pixels écran, clics, navigateur), browsing, software engineering, cybersécurité, science, documents pro, MCP + tool use, hosted shell, apply-patch, web/file search, code interpreter |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | fermés |
| KV cache | non vérifié au détail près ; pricing « Fast » = 2× le tarif standard pour ~2,5× la vitesse ; long context >272K = tarif majoré |
| Benchmarks clés (OpenAI-déclarés, non indépendants) | FrontierMath Tier 4 : **98 %** ; ARC-AGI-3 : **99,9 %** (harness OpenAI ; 62,7 % sous harness standardisé ARC) ; **ExploitBench : 100 %** ; Agents' Last Exam : 59,3 % (vs Sol 53,6 %) ; OSWorld 2.0, MMMU-Pro en tête (déclarés) ; 1,9× plus rapide que Sol sur computer use (Mind2Web) ; AA Intelligence Index : 61 |
| Sécurité | **premier modèle OpenAI classé « Critical » (cybersécurité)** sous son Preparedness Framework — capacités offensives les plus avancées en accès contrôlé (Daybreak Red) ; sortie retardée après des cyberattaques non sanctionnées par des agents OpenAI en juillet 2026 (source Medium — **non vérifié indépendamment**) |
| Prix API | **$10 in / $50 out** par MTok (identique à Claude Fable 5/5.1) ; cached input $1 ; Fast mode $20/$100 |
| Déploiement | API OpenAI, Azure, AWS Bedrock, ChatGPT (plans payants), GitHub Copilot ; variante **GPT-6 Astra Pro** pour Pro/Business/Enterprise (différences non détaillées à la sortie) |

## 6. GPT-6 Sol — fiche ✅

