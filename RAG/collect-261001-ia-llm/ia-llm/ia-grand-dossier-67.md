---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-67
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Fireworks AI", "Groq", "Microsoft", "Nvidia", "OpenAI", "OpenRouter", "Together AI", "vLLM", "xAI"]
dates: ["2026-08-21", "2026-08-26", "2026-09-27"]
keywords: ["agents", "attention", "awq", "claude", "cost", "deepseek", "embedding", "embeddings", "fine-tuning", "gpu", "grok", "grok 4"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5541, 5607]
sha256: 63572be150295439c3c2a33dabb0d1f302c0d48dd851c365c1bac5270efe46ef
---

# IA — Le grand dossier

**Q8.** Votre RAG consomme 3M tokens/jour en entrée sur Claude Sonnet 5 (3,00 $/1M). Une RTX 4090 (2 000 $, amortie 36 mois) vous coûte ~270 $/mois tout compris (amortissement + élec + 4 h d'exploitation). Le local est-il rentable ? Montrez le calcul.
> **Corrigé.** Cloud : 3M × 30 j = 90M tokens/mois × 3,00 $/1M = **270 $/mois**. Local : **~270 $/mois**. Égalité sur le seul prix — le local ne se justifie ici que par la confidentialité/la latence/le coût marginal nul en cas de croissance. Si le volume double, le local gagne ; avec Luna à 0,20 $/1M (18 $/mois), le cloud écrase le local sur le prix.

**Q9.** Qu'est-ce que le KV cache, et pourquoi un modèle qui « tient » en VRAM peut-il quand même planter (OOM) ?
> **Corrigé.** Le KV cache stocke les états d'attention de la conversation en cours ; il croît avec la longueur du contexte et le nombre de requêtes parallèles (ex. +10 Go à 128K sur un 7B). La quantification des poids ne le réduit pas (sauf option dédiée : `--cache-type-k q8_0` sur llama.cpp, cache quantifié sur vLLM). Un dimensionnement se fait donc sur **poids + KV cache (contexte cible) + overhead**, pas sur les poids seuls.

**Q10.** Dessinez (en 5 lignes) l'architecture cible « hybride » recommandée pour une PME/ETI, en nommant les composants et les noms logiques de modèles.
> **Corrigé (exemple).** `[Apps/Agents] → [LiteLLM :4000]` avec 4 noms logiques : `triage` → Groq/Cerebras/OpenRouter (interactif pas cher) ; `raisonnement` → Anthropic/OpenAI (qualité, budget) ; `local` → vLLM :8000 (Qwen3-32B-AWQ, données sensibles) ; `embeddings` → text-embedding-3-small ou Nomic local. Fallbacks : triage→raisonnement, raisonnement→triage. Observabilité : Postgres (spend logs) + Prometheus + alertes budget. Le code ne connaît que les noms logiques : changer de provider = éditer le YAML.

---

## Sources et méthode

- Prix et versions vérifiés par recherche web le **27/09/2026** (pages pricing officielles, trackers communautaires : llm-price-tracker, Artificial Analysis, OpenRouter). Les tarifs API changent vite : **re-vérifier avant tout engagement budgétaire**.
- Mentions « à vérifier » : faits rapportés par une seule source ou en cours d'évolution (ex. tarifs Grok 4.6, disponibilité OVHcloud GPU, bascule Groq Llama en enterprise-only du 26/08/2026).
- Aucune donnée confidentielle, aucune clé réelle dans ce document. Les configs sont des modèles à adapter (variables d'environnement).

*Fin de la Partie 2 — la Partie 3 couvrira : [à définir avec Zelef].*

---

## Annexe A — Fiches détaillées des 12 providers à connaître par cœur

### A.1. OpenRouter — l'agrégateur universel

- **Fondé** : 2023. **Modèle** : marketplace/routeur managé, une clé API pour 400+ modèles.
- **Économie** : prix public du provider sous-jacent **sans marge au token** ; 5,5 % de frais sur achats de crédits (min 0,80 $) ; +5 % en BYO-key (tu apportes ta clé provider).
- **Free tier** : variantes `:free` — 20 RPM ; 50 req/jour (<10 $ de crédits achetés), 1 000 req/jour au-delà. Lentes/en file : pour dev/tests, pas pour la prod.
- **Routage** : `provider.order` (liste ordonnée), `allow_fallbacks`, `require_parameters` (ex. n'exposer que les providers supportant les tools), `data_collection: deny` (refuser les providers qui logguent).
- **Points durs** : un modèle peut être délisté sans préavis ; les providers sous-jacents peuvent logger les prompts (lire chaque fiche modèle) ; latence = routage + provider (variable).
- **LiteLLM** : préfixe `openrouter/<fournisseur>/<modèle>`, env `OPENROUTER_API_KEY`.
- **Quand l'utiliser** : prototypage multi-modèles, comparaisons A/B, fallback universel, petite équipe sans gateway.

### A.2. LiteLLM — la gateway open source

- **Éditeur** : BerriAI. **Licence** : MIT. **Version vérifiée** : v1.97.0 (21/08/2026).
- **Deux formes** : bibliothèque Python (`litellm.completion(model="groq/...")`) et **serveur proxy** (`:4000`, API compatible OpenAI) — c'est le proxy qui sert en entreprise.
- **Capacités** : 100+ providers ; fallbacks (globaux + par type d'erreur : 429, contexte dépassé, refus de modération) ; load balancing (5 stratégies : simple-shuffle, usage-based-v2, latency-based, cost-based, least-busy) ; cooldowns automatiques ; clés virtuelles avec budgets ; spend logs (Postgres) ; UI d'admin (`:4000/ui`) ; `/metrics` Prometheus ; `/health/readiness` ; MCP Gateway (`/v1/mcp`) stable depuis v1.85.
- **Économie** : gratuit en self-hosted (ton serveur + Postgres) ; offre cloud/entreprise payante (à vérifier).
- **Points durs** : à opérer (HA, backups Postgres, rotation des clés) ; la syntaxe des fallbacks fins évolue vite → épingler la version et lire la doc correspondante.
- **Quand l'utiliser** : dès que 2+ apps ou 2+ providers ; dès qu'un budget doit être plafonné ; dès que les données sensibles exigent un point de contrôle unique.

### A.3. Groq — la vitesse (LPU)

- **Fondé** : 2016. **Tech** : LPU (Language Processing Unit), SRAM on-chip au lieu de HBM → 500–800 tok/s, TTFT ~0,7 s.
- **Catalogue** : réduit et volontaire (~6 modèles self-serve) : GPT-OSS-20B/120B, Llama 4 Scout/Maverick, Qwen, DeepSeek-distills, Whisper (STT). **Attention** : le 26/08/2026, des Llama (3.1 8B, 3.3 70B) sont passés en « enterprise-only / contact sales » — à vérifier avant de pinner un modèle.
- **Prix (sept. 2026)** : GPT-OSS-120B 0,15/0,60 $/1M ; Llama 3.3 70B ~0,59/0,79 $/1M. Free tier : ~30 RPM, quotas TPM selon modèle, sans carte.
- **Actu** : levée 350 M$ en août 2026 (valo 3,5 Md$, en baisse vs 6,9 Md$ en 2025) ; NVIDIA a licencié la tech LPU ; devenu NVIDIA Cloud Partner (prévoit d'ajouter du GPU à son offre).
- **Quand l'utiliser** : temps réel (agents, réécriture de requêtes, re-ranking, classification interactive). Pas pour : gros contextes stables, fine-tuning self-serve.

### A.4. Cerebras — la vitesse extrême (WSE)

- **Tech** : Wafer-Scale Engine (WSE-3) : 44 Go de SRAM sur un wafer entier, ~27 kW par système → **1 800–3 000 tok/s** mesurés (le record du marché).
- **Catalogue** : ~8 modèles self-serve (GPT-OSS, Llama, Qwen3, Gemma…).
- **Prix** : GPT-OSS-120B ~0,35/0,75 $/1M (plus cher que Groq, 3–5× plus rapide). Essai : 5 $ de crédit (30 jours, CB requise) ; pas de free tier permanent (à vérifier).
- **Points durs** : contextes parfois rognés vs natif (ex. 65K en essai vs 131K payant sur certains modèles) ; déploiement d'un nouveau modèle = jours/semaines (silicium propriétaire).
- **Quand l'utiliser** : quand chaque milliseconde compte et que le budget suit (trading, jeux, voix temps réel).

### A.5. Together AI — le cloud des poids ouverts

- **Fondé** : 2022. **Position** : « servir, tuner, scaler » les modèles ouverts au même endroit.
- **Offres** : Serverless Inference (200+ modèles, $/1M tokens de 0,03 à 4,50 $) ; fine-tuning (LoRA/full, jusqu'à 405B) ; Dedicated Endpoints (single-tenant, H100 dès ~3,99 $/h) ; GPU Clusters (entraînement) ; Batch API (-50 %).
- **Économie** : prépayé (min 5 $, plus de free trial au 09/2026 — à vérifier) ; ex. DeepSeek V4 Pro 1,74/3,48 $/1M, identique à Fireworks/Parasail (prix de marché du modèle).
- **Quand l'utiliser** : équipe qui veut un seul fournisseur pour inférence + fine-tuning + capacité réservée, sans opérer de GPU.

### A.6. Fireworks AI — l'inférence production

