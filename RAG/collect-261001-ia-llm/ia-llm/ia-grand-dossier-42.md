---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-42
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Baseten", "Cerebras", "Google", "Groq", "Lambda", "Mistral", "OpenAI", "vLLM"]
dates: []
keywords: ["dpo", "embeddings", "fine-tuning", "gguf", "gpu", "llama", "llama.cpp", "memory", "mistral", "nvidia", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3274, 3401]
sha256: 521ca0d59ee4dd1bc748698de1c034578ddc4a1e89fb99b5cfca578ce252b40b
---

# IA — Le grand dossier

```
┌─ PROVIDERS ─────────────────────────────────────────────────┐
│ Vitesse : Groq (500-800 tok/s) · Cerebras (1800-3000 tok/s)  │
│ Prix    : DeepInfra · Novita · SiliconFlow (dès 0,02 $/1M)   │
│ Qualité : Anthropic · OpenAI · Google (frontier)             │
│ Europe  : Mistral · Scaleway · OVHcloud (à vérifier)        │
│ GPU/h   : Vast.ai (~0,14 $ 4090 spot) · RunPod · Lambda      │
├─ GATEWAY : LiteLLM (self-hosted, MIT) ───────────────────────┤
│ 1 YAML → noms logiques : triage / raisonnement / local /     │
│ embeddings · fallbacks · budgets · spend logs · /metrics     │
├─ LEVIERS PRIX ───────────────────────────────────────────────┤
│ batch -50 % · prompt caching -90 % (lecture) · provider le   │
│ moins cher du même modèle (écart ×1-4) · routage coût/qualité│
├─ LOCAL ──────────────────────────────────────────────────────┤
│ Ollama (simple) · vLLM (équipe) · llama.cpp (perf) · GGUF Q4  │
│ par défaut · Qwen3-32B = sweet spot (~22 Go VRAM)            │
│ 70B Q4 ≈ 48 Go → 2×24 Go, 48 Go pro, ou Mac ≥64 Go unifiés  │
├─ RÈGLES D'OR ────────────────────────────────────────────────┤
│ • Le code ne connaît que des noms logiques                   │
│ • 2 providers min par cas critique + fallbacks testés        │
│ • Budgets + alertes dès le jour 1                            │
│ • Eval métier (50-100 Q) avant chaque changement            │
│ • Licence notée à chaque nouveau modèle                      │
│ • Température 0-0,2 en exploitation ; jamais d'action         │
│   destructive auto sans validation humaine                   │
└──────────────────────────────────────────────────────────────┘
```

---

*Fin de la Partie 2 — Providers, passerelles & IA locale.*
*Partie 3 (à définir avec Zelef) : [sujet à venir].*

---

## Annexe P — Templates opérationnels : suivi des coûts et charte d'usage

### P.1. Tableau de suivi mensuel des coûts IA (à remplir)

| Mois | Provider / gateway | Tokens in (M) | Tokens out (M) | Coût API ($) | Coût infra locale ($) | Coût total ($) | Coût / utilisateur ($) | Écart vs budget |
|---|---|---|---|---|---|---|---|---|
| 2026-10 | LiteLLM (détail par clé virtuelle) | | | | | | | |
| 2026-10 | GPU local (amort. + élec) | — | — | — | | | | |
| 2026-10 | **Total** | | | | | | | |

**Formules** : coût/utilisateur = total / utilisateurs actifs ; alerte si écart vs budget > 10 % deux mois de suite → revue des routages.

### P.2. Fiche de choix d'un nouveau modèle (à remplir avant ajout au YAML)

```
Modèle : ______________________  Version/révision : ______________________
Provider : _____________________  Prix in/out ($/1M) : ______ / ______
Licence (si poids ouverts) : ________________  URL dépôt : ________________
Fenêtre de contexte : ________  Supporte : [ ] tools  [ ] JSON  [ ] vision
Eval métier : score _____ (baseline _____ )  Date : __________
Cas d'usage : [ ] triage  [ ] raisonnement  [ ] local  [ ] embeddings
Données sensibles ? [ ] oui → local/gateway UE uniquement  [ ] non
Fallback prévu : ______________________
Validé par : ___________________  Date de revue (3 mois) : _______________
```

### P.3. Charte d'usage interne de l'IA (modèle, 10 règles)

1. **Données classifiées** (contrats, plans, données personnelles) : uniquement via les modèles logiques `local` ou les providers validés par le DPO — jamais via un compte personnel.
2. **Aucune clé API personnelle** dans le code, les scripts ou les dépôts git de l'entreprise.
3. Toute nouvelle application IA passe par la gateway (LiteLLM) : pas d'appel direct aux providers.
4. Les budgets par équipe sont des plafonds durs : à 100 %, ça coupe — prévoir, pas subir.
5. **Vérification humaine obligatoire** avant toute action à impact (commande sur équipement, envoi client, décision RH/maintenance).
6. Les sorties de l'IA sont des propositions : la responsabilité reste à l'opérateur qui valide.
7. Signaler tout comportement anormal (fuite de données, réponse suspecte, coût anormal) au responsable IA sous 24 h.
8. Ne pas entraîner/fine-tuner sur des données personnelles sans base légale documentée.
9. Citer les sources quand l'IA s'appuie sur la documentation interne (traçabilité).
10. Cette charte est revue tous les 6 mois ; les manquements relèvent du règlement intérieur.

### P.4. Checklist DPA (accord de traitement des données) avec un provider cloud

- [ ] Localisation du traitement (UE/FR si exigé) écrite au contrat.
- [ ] **Zero-retention** : engagement écrit de non-conservation et non-réutilisation des prompts/completions pour l'entraînement.
- [ ] Durées de conservation des logs (30 jours max recommandé) et droit d'audit.
- [ ] Sous-traitants listés ; notification en cas de changement.
- [ ] Suppression des données à la résiliation (délai + attestation).
- [ ] Mesures de sécurité : chiffrement en transit (TLS 1.2+) et au repos, contrôle d'accès, SOC 2/ISO 27001.
- [ ] Procédure de notification de violation (< 72 h).
- [ ] Responsabilité et assurance en cas de fuite.

---

## Annexe Q — Comparatif express : serverless vs dédié vs spot (inférence)

| Critère | **Serverless** (Together, DeepInfra, RunPod serverless) | **Dédié** (Baseten, Together Dedicated, PTU) | **Spot** (Vast.ai, RunPod spot) |
|---|---|---|---|
| Facturation | Au token ou à la seconde d'exécution | À l'heure/instance, 24/7 | À la seconde, -50 à -90 % |
| Cold start | Oui (sauf provisionné) | Non | N/A (VM persistante) |
| Disponibilité | Haute (multi-tenant) | Garantie (single-tenant) | **Préemptible** (coupure possible) |
| Idéal pour | Trafic variable, dev, API produit | SLA, latence garantie, conformité | Batch, fine-tuning, expérimentation |
| Piège | Coût au token × volume = surprise | Sur-provisionnement = gaspillage | Perte du travail sans checkpoints |
| Exemple prix | 0,03–4,50 $/1M (Together) | H100 ~3,99 $/h (Together) | 4090 ~0,14 $/h (Vast) |

**Règle** : serverless par défaut → dédié quand le SLA l'exige ou que le volume rend l'heure moins chère que le token → spot uniquement pour le batch checkpointé.

---

## Annexe R — Les 10 commandes à connaître par cœur

```bash
# Santé de la stack
curl -s http://litellm:4000/health/readiness | head -c 200          # gateway OK ?
curl -s http://vllm:8000/health && echo " vLLM OK"                  # vLLM OK ?
ollama list                                                          # modèles Ollama présents
nvidia-smi --query-gpu=name,memory.used,memory.total,utilization.gpu --format=csv

# Coût du jour (LiteLLM + Postgres)
docker exec postgres psql -U litellm -d litellm -c \
 "SELECT model, SUM(total_tokens)/1e6 AS mtok, SUM(spend) AS usd
  FROM litellm_spendlogs WHERE starttime > now() - interval '1 day'
  GROUP BY 1 ORDER BY 3 DESC LIMIT 10;"

# Test d'un nom logique (doit répondre même si un provider est down)
curl http://litellm:4000/v1/chat/completions -H "Authorization: Bearer $KEY" \
  -H "Content-Type: application/json" \
  -d '{"model":"triage","messages":[{"role":"user","content":"ping"}]}' --max-time 60

# Bench local avant achat
./build/bin/llama-bench -m modele.gguf -p 512 -n 128                # tok/s réels
```

---

