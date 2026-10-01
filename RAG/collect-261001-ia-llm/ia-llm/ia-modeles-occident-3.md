---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-3
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "Meta", "Nvidia", "OpenAI"]
dates: ["2026-04-24", "2026-06-22", "2026-08-10", "2026-09-24", "2026-09-26", "2026-09-27", "2026-11-21"]
keywords: ["agent", "agentic", "agents", "agi", "astra", "benchmarks", "chatgpt", "claude", "cyber", "fine-tuning", "gemini", "gpt-5.6"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [182, 256]
sha256: 9878c87cb9d9c35a3f348f21958925a6cc8280378bd14f9617d0542cb9186189
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

- **GPT-5.4 Pro** (`gpt-5.4-pro`) : variante haute-précision, « hardest reasoning », $30 / $180 par MTok (Responses API uniquement). Sorti avec le 5.4 standard le 5 mars 2026.
- **GPT-5.4 mini** : $0,75 / $4,50 par MTok. Sorti le 17 mars 2026.
- **GPT-5.4 nano** : $0,20 / $1,25 par MTok. Sorti le 17 mars 2026.
- Les trois restent disponibles en API au 27/09/2026. Fine-tuning : **non vérifié au 27/09/2026**.

## 15. GPT-5.4-Cyber et les variantes Daybreak ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | GPT-5.4-Cyber ; puis GPT-5.5-Cyber (22/06/2026), GPT-5.6-Cyber (10/08/2026) |
| Sortie | **14 avril 2026** (GPT-5.4-Cyber, via **Trusted Access for Cyber**) |
| Statut | accès restreint — programme **Daybreak** / Daybreak Red d'OpenAI (cybersécurité offensive contrôlée) |
| Architecture / contexte / prix | **non vérifiés au 27/09/2026** (accès contrôlé, pas de fiche publique) |
| Note | les variantes -Cyber sont les déclinaisons « capabilities offensives » des flagships, pendant d'OpenAI au programme Glasswing d'Anthropic (Mythos) et au Fairwind Program de Google (Gemini Flash Cyber). |

## 16. GPT-5.5 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-5.5 (`gpt-5.5`) ; variantes GPT-5.5 Pro, GPT-5.5 Thinking, GPT-5.5 Instant ; GPT-5.5-Cyber |
| Sortie | **23 avril 2026** (annonce officielle ; API le 24 avril selon le pricing changelog) ; GPT-5.5 Instant : 5 mai 2026 (source unique bofai — **non vérifié au 27/09/2026**) |
| Statut | disponible — flagship d'avril à juillet 2026 (remplacé par GPT-5.6 Sol) |
| Architecture | « premier base model entièrement réentraîné depuis GPT-4.5 » ; pretraining terminé **mars 2026** sur NVIDIA GB200/GB300 NVL72 ; architecture unifiée **omnimodale** (text/image/audio/vidéo end-to-end) — synthèse tierce (jamoeight) ; paramètres **non vérifiés** |
| Contexte | **1 M tokens en API**, **400K dans Codex** |
| Features | agentic (planification, tool use, vérification, continuation à travers l'ambiguïté) ; coding, recherche, data analysis ; Thinking (raisonnement long) ; Pro (max accuracy) ; Instant (basse latence, usage quotidien) |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | fermés |
| Benchmarks clés | **non vérifié précisément via sources primaires** (revendiqué : plus fort que GPT-5.4 et Claude Opus 4.7 en raisonnement/autonomie — sources secondaires) |
| Prix API | GPT-5.5 : **$5 / $30** ; **Pro : $30 / $180** ; Batch/Flex à 50 %, Priority à 2,5× |
| Déploiement | ChatGPT (Plus/Pro/Business/Enterprise), Codex, API (Responses + Chat Completions) |
| Sécurité | system card + bug bounty bio $25K (the-ledger, 24/04/2026) |

## 17. GPT-5.5 Pro, Thinking, Instant ✅

- **GPT-5.5 Pro** (`gpt-5.5-pro`) : $30 / $180 par MTok — la variante « max accuracy » de la génération 5.5.
- **GPT-5.5 Thinking** : le mode raisonnement long du 5.5, exposé comme variante.
- **GPT-5.5 Instant** : 5 mai 2026 (source unique bofai — **non vérifié au 27/09/2026**), basse latence.
- **GPT-5.5-Cyber** : 22 juin 2026, via Daybreak.

## 18. Computer use : d'OpenAI, l'agent de bureau est devenu natif

Le computer use natif (opérer un OS via pixels d'écran, souris, clavier) est arrivé avec **GPT-5.4** (mars 2026, Playwright sur screenshots) et a culminé avec **GPT-6 Astra** (sept. 2026, 1,9× plus rapide que Sol sur Mind2Web). C'est devenu un champ de bataille explicite : Astra revendique OSWorld 2.0 en tête, Claude Opus 4.6–4.8 réplique (OSWorld 72,7 % → 83,4 %), et Meta annonce « Mac desktop control » pour Muse (sept. 2026). En entreprise, c'est le computer use — pas le chatbot — qui fait vendre les tiers Pro.

## 19. Prix API OpenAI : la guerre des prix de l'été 2026

Chronologie des baisses (par MTok in/out) :
- **30 juillet 2026** : Luna −80 % ($1/$6 → $0,20/$1,20), Terra −20 % ($2,50/$15 → $2/$12).
- **21 août 2026** : Sol $5/$30 → $4/$20 (promo jusqu'au 21/11/2026).
- **22 septembre 2026** : GPT-6 Sol à **$2/$10** (−50 % vs 5.6 Sol), GPT-6 Luna à **$0,10/$0,50** (−50 % vs 5.6 Luna) — le jour même où Anthropic lançait Claude Opus 5.5 à $4/$20 (−20 %).
- Le cache read reste à 10 % du prix input (90 % de remise) ; le cache-write pricing est arrivé avec la 5.6 ; au-delà de 272K tokens de prompt, tarif long-contexte ×2 in / ×1.5 out.
- Leçon : le coût par tâche s'effondre plus vite que les prix affichés (tokens en −47 % sur certaines tâches entre 5.4 et 5.5, revendiqué par OpenAI).

## 20. Prompt caching et cache-write pricing chez OpenAI

OpenAI applique 90 % de remise sur les cached input reads. Avec GPT-5.6, arrivée du **cache-write pricing** (on paie l'écriture du cache). GPT-6 améliore les hit rates par défaut (« higher cache hit rates by default »). Détail des optimisations d'inférence au-delà : **non vérifié au 27/09/2026** (KV cache, quantification serveur non publiés).

## 21. Fine-tuning chez OpenAI : l'état au 27/09/2026

- **GPT-6 Sol et GPT-6 Luna : fine-tuning NON supporté** (endpoints assistants/realtime/live/fine-tuning listés comme non supportés).
- Pour GPT-5.3 → GPT-5.6 et GPT-6 Astra : support exact **non vérifié au 27/09/2026** dans les sources consultées.
- En pratique, OpenAI pousse le RL post-training côté entreprise et les variantes -Pro plutôt que le fine-tuning supervisé classique sur les flagships.

## 22. GPT-6 Astra Pro et variantes annoncées

**GPT-6 Astra Pro** : variante pour les plans Pro/Business/Enterprise, différences non détaillées par OpenAI à la sortie (sept. 2026) — **non vérifié au 27/09/2026** au-delà de l'existence du nom. Rumeur TestingCatalog (24–26/09/2026) : une variante **« aeon »** d'Astra pour un agent « o » — **non vérifié au 27/09/2026**.

## 23. Rumeurs non vérifiées autour d'OpenAI (sept. 2026)

- **« ChatGPT Pro Max » $500/mois** (TestingCatalog via X, 24/09/2026, lié au DevDay) — **non vérifié**.
- **GPT-6 « aeon »** — **non vérifié**.
- **Raison du retard d'Astra** (cyberattaques non sanctionnées par des agents OpenAI, juillet 2026) — source unique Medium, **non vérifié indépendamment**.
- **Benchmarks OpenAI-déclarés** (FrontierMath 98 %, ARC-AGI-3 99,9 %, ExploitBench 100 %) — pas de vérification indépendante dans les sources consultées ; le 99,9 % ARC-AGI-3 tombe à 62,7 % sous harness standardisé (intelligentliving.co).

## 24. Tableau récapitulatif OpenAI (sept. 2026)

