---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-69
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Cerebras", "DeepSeek", "Fireworks AI", "Groq", "OpenRouter", "vLLM"]
dates: []
keywords: ["agent", "aws", "cost", "deepseek", "embeddings", "gpu", "latency", "llama", "nvidia", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5771, 5887]
sha256: ec3827fb8ddf866789c76cf67a08da440a9a75af832532262549c767a2c5ddad
---

# Exemples d'appels :
# ask("local", "Résume cette procédure de consignation.", system="Tu es technicien.")
# ask("triage", "Classe ce ticket.", ...)
# ask_avec_escalade("Dimensionne un onduleur 60 kVA pour 45 min d'autonomie.")
```

### B.3. Supervision : ce qu'on monitor (tes outils existants)

| Sonde | Où | Seuil d'alerte (exemple) |
|---|---|---|
| `GET /health/readiness` LiteLLM | Zabbix HTTP / Prometheus | Down > 2 min → critique |
| `GET /health` vLLM + Ollama (`/api/tags`) | idem | Down → critique |
| `litellm_request_latency` (Prometheus) | Prometheus → Grafana | p95 > 30 s sur `triage` → warning |
| Spend jour vs budget (SQL §4.6) | cron + mail | > 80 % budget mensuel → warning ; 100 % → la clé coupe |
| Taux d'erreur par provider | `/metrics` | > 5 % sur 10 min → warning (préparer le reroutage) |
| VRAM GPU (`nvidia-smi`) | Zabbix agent / node_exporter | > 95 % → warning (risque OOM) |
| Espace disque (poids, logs) | Zabbix | > 85 % → warning |

### B.4. Sauvegarde et reprise

- **Postgres LiteLLM** : `pg_dump` quotidien (clés virtuelles, spend logs, budgets) → rétention 30 j.
- **Poids des modèles** : registre interne (NFS/SMB versionné) ; pas besoin de sauvegarder ce qui se re-télécharge, mais **épingler les révisions** (hash du commit HF ou tag Ollama) pour la reproductibilité.
- **Qdrant** : snapshots natifs (`POST /snapshots`) avant chaque réindexation massive.
- **Config** : `litellm_config.yaml` + `.env.example` (sans secrets) dans git ; secrets dans Vault/Bitwarden.

### B.5. Durcissement minimal avant exposition

1. LiteLLM **jamais** exposé sans TLS : Caddy/Nginx devant, `master_key` forte (32+ caractères).
2. Ollama/vLLM : bindés en `127.0.0.1` ou réseau d'admin ; seul LiteLLM les appelle.
3. Clés virtuelles : une par application, `max_budget` mensuel, allowlist de modèles (pas de `raisonnement` pour le script de tickets).
4. Logs : ne jamais logger le contenu des prompts en prod (ou chiffrer) ; les spend logs ne contiennent que métadonnées.
5. Mises à jour : épingler les tags Docker ; tester les montées de version (tokenizer, templates) en pré-prod.

---

## Annexe C — Prix comparés sur 3 modèles de référence (sept. 2026)

**Lecture** : même modèle, prix très différents selon l'hébergeur. Choisir le provider, c'est arbitrer prix/vitesse/garanties. Tarifs à re-vérifier (ils bougent chaque mois).

| Provider | GPT-OSS-120B (in / out $/1M) | DeepSeek V4 Pro (in / out $/1M) | Llama 3.3 70B (in / out $/1M) | Notes |
|---|---|---|---|---|
| DeepInfra | 0,09 / 0,45 | 1,74 / 3,48 (via partenaires) | ~0,23 / 0,40 (à vérifier) | Prix plancher |
| Groq | 0,15 / 0,60 | — | 0,59 / 0,79 | Vitesse ~500–800 tok/s |
| Fireworks (Standard) | ~0,15 / 0,60 (à vérifier) | 1,74 / 3,48 | ~0,70 / 0,70 (à vérifier) | Tier Priority ×1,5 |
| Together | ~0,15 / 0,60 (à vérifier) | 1,74 / 3,48 | ~0,70 / 0,70 (à vérifier) | Batch -50 % |
| OpenRouter | selon provider routé | selon provider routé | selon provider routé | +5,5 % sur crédits |
| Cerebras | 0,35 / 0,75 | — | — | Vitesse 1 800–3 000 tok/s |
| Datacrunch (serverless) | ~0,50 / 0,50 (à vérifier) | — | — | GPU RTX Pro 6000 |
| Azure / AWS / GCP | selon région/PTU | — | selon région | SLA + conformité max |

**Enseignement** : l'écart va de 1 à 4 sur le même modèle. Un routeur coût (LiteLLM `cost-based-routing`) capte cet écart automatiquement.

---

## Annexe D — Les erreurs API et la stratégie de retry

| Erreur | Sens | Action automatique (gateway) | Action humaine |
|---|---|---|---|
| **429** rate_limit | Quota RPM/TPM dépassé | **Retry avec backoff exponentiel** + fallback vers un autre provider du groupe ; cooldown du deployment fautif | Augmenter le tier, répartir sur 2 clés, activer le cache |
| **429** insufficient_quota | Crédit épuisé | Alerte immédiate (pas de retry utile) | Recharger / vérifier le budget |
| **400** context_length_exceeded | Prompt > fenêtre du modèle | Fallback vers un modèle à grand contexte (cf. §4.4 `triage-long`) | Tronquer/résumer en amont (chunking) |
| **400** invalid_request | Paramètre non supporté (ex. `temperature` sur un modèle de raisonnement) | Normaliser les paramètres par modèle dans la gateway | Corriger l'appel |
| **401/403** | Clé invalide / révoquée | Alerte ; bascule sur clé de secours si configurée | Rotation de clé (Vault) |
| **500/502/503** | Panne provider | Retry 1–2× puis fallback ; cooldown 60 s | Rerouter durablement si ça dure |
| **529** (Anthropic) | Surcharge temporaire | Retry backoff ; fallback | Idem 503 |
| **Timeout** (>120 s) | Réponse trop lente | Couper (timeout client) ; fallback | Choisir un provider plus rapide pour ce cas |

**Règle d'or** : **retry ≠ boucle infinie**. Backoff exponentiel (1 s → 2 s → 4 s…), jitter, max 3 tentatives, puis fallback, puis erreur propre à l'utilisateur. Un retry agressif transforme une panne provider en auto-DDoS facturé.

---

## Annexe E — Checklist de mise en production d'une gateway (à cocher)

- [ ] LiteLLM épinglé (tag Docker), config versionnée en git, secrets en Vault.
- [ ] Postgres dédié, sauvegardé quotidiennement, testé en restauration.
- [ ] Clés virtuelles : 1 par application, budgets mensuels, allowlist de modèles.
- [ ] Fallbacks configurés et **testés** (couper un provider, vérifier la bascule).
- [ ] Alertes : spend >80 % budget, erreur >5 %, latence p95, down des backends.
- [ ] TLS + auth sur tous les endpoints exposés ; Ollama/vLLM non exposés directement.
- [ ] Logs sans contenu de prompts (ou chiffrés) ; DPA signés avec les providers cloud.
- [ ] Runbook : « le provider X est down » (qui reroute, qui décide, qui paie).
- [ ] Eval de non-régression : rejouer 50 questions après chaque changement de modèle/provider.
- [ ] Revue trimestrielle des prix (ils bougent) + des licences des modèles ouverts utilisés.

---

## Annexe F — 3 scénarios budgétaires annuels (ordres de grandeur)

**Hypothèses communes** : 20 utilisateurs, 250 j/an, mix 80 % triage / 15 % raisonnement / 5 % local.

### Scénario 1 — « Tout cloud direct » (sans gateway)

- 10M tokens/j/mois… (10M tokens/jour : 8M triage à 0,30 $/1M + 1,5M raisonnement à 3 $/1M + 0,5M embeddings à 0,02 $/1M)
- ≈ (2,40 + 4,50 + 0,01) $/jour ≈ **7 $/jour → ~1 750 $/an**
- Risques : aucun plafond, aucune bascule, 3 clés à gérer par app.

### Scénario 2 — « Hybride avec gateway » (recommandé)

- Triage via Groq/Cerebras pas cher + prompt caching : ~1,50 $/jour
- Raisonnement plafonné (budget LiteLLM 100 $/mois) : ~3 $/jour
- Local (RTX 4090 amortie, embeddings + bulk) : ~0,90 $/jour tout compris
- Total ≈ **5,40 $/jour → ~1 350 $/an** + 2 000 $ d'amortissement GPU la 1ʳᵉ année
- Gains : plafonds, fallbacks, observabilité, données sensibles en local.

### Scénario 3 — « Local d'abord » (souveraineté)

- Serveur 2×4090 (8 000 $, 36 mois) + élec : ~330 $/mois
- Cloud : pics + frontier, budget 50 $/mois
- Total ≈ **~4 560 $/an** (dont ~3 960 $ d'infra)
- Justifié si : données non externalisables, volume > 50M tokens/jour, ou exigence d'air-gap.

**Moralité** : à volume modéré, l'hybride coûte à peine plus cher que le tout-cloud mais apporte le contrôle ; le local pur ne se justifie économiquement qu'au-delà d'un seuil de volume (cf. §8.2) ou pour des raisons non financières.

---

## Annexe G — Veille : 10 signaux à surveiller (fin 2026 → 2027)

