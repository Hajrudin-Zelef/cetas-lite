---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-39
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Fireworks AI", "Google", "Groq", "Lambda", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Z.ai", "vLLM"]
dates: []
keywords: ["agents", "awq", "aws", "benchmark", "benchmarks", "claude", "deepseek", "dpo", "dram", "embedding", "embeddings", "fable 5"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2914, 3032]
sha256: ef0b51f7306dd6958652e666f25ce7df92b378862e1c6b98554254120348ea22
---

# IA — Le grand dossier

- **AI Act** : un LLM interne d'entreprise = généralement « risque limité » (obligations de transparence : informer que c'est une IA) ; devient « haut risque » si décision automatisée à impact (recrutement, maintenance prédictive critique) → analyse d'impact.
- **Données personnelles** : un RAG indexant des tickets nominatifs = traitement RGPD → registre, durées de conservation, droit d'effacement (prévoir la réindexation).
- **Secret des affaires** : le local est l'argument le plus simple (« les données ne sortent pas ») face à un DPO.
- **Traçabilité** : épingler modèle + version + config (qui a répondu quoi, quand, avec quels chunks) — les spend logs LiteLLM + les logs d'accès Qdrant couvrent 80 % du besoin.

### 11.3. Durcissement réseau type

```
[Internet] --(443)--> [Reverse proxy TLS] --(4000)--> [LiteLLM] --+--> [vLLM :8000] (VLAN IA)
                                                                  +--> [Ollama :11434] (VLAN IA)
                                                                  +--> [providers cloud] (HTTPS sortant)
[VLAN users] --> [Open WebUI :3000] (auth SSO/LDAP)
[VLAN IA] : aucun accès Internet entrant ; sortant limité (updates, HF)
```

---

## 12. Évaluer un modèle local avant de le mettre en prod

### 12.1. Les 3 niveaux d'évaluation

| Niveau | Outil | Ce que ça mesure | Durée |
|---|---|---|---|
| **1. Perplexité** | `llama-perplexity` sur un corpus FR | Qualité brute du modèle/quantif | Minutes |
| **2. Benchmarks** | `lm-evaluation-harness` (MMLU-fr, ARC, HellaSwag…) | Capacités générales, comparables | Heures |
| **3. Eval métier** | 50–100 questions réelles + juge (LLM ou humain) | **Ce qui compte vraiment** : ton usage | 1–2 jours à construire |

### 12.2. Protocole minimal (à faire une fois, à rejouer à chaque changement)

```bash
# 1. Perplexité : Q4 vs Q8 sur TON corpus (pas WikiText générique)
./build/bin/llama-perplexity -m qwen3-32b-q4_k_m.gguf -f corpus_interne.txt -c 2048
./build/bin/llama-perplexity -m qwen3-32b-q8_0.gguf   -f corpus_interne.txt -c 2048
# Écart < 0,3 pt -> Q4 validé pour ton usage

# 2. Benchmark standard (comparabilité)
pip install lm-eval
lm_eval --model hf --model_args pretrained=Qwen/Qwen3-32B \
  --tasks mmlu,arc_challenge,hellaswag --batch_size 8

# 3. Eval métier : questions/réponses attendues en JSON, scoring par un juge
python eval_metier.py --model qwen3:32b --dataset eval_rag_100q.json
# Métriques : précision@1, taux de citation correcte, refus propres vs hallucinations
```

**Seuils de décision** (indicatifs) : si le 32B local atteint ≥85 % du score du frontier sur ton eval métier → local par défaut + escalade ; entre 70 et 85 % → hybride par criticité ; <70 % → cloud par défaut, local pour le non-critique.

---

## 13. Coûts cachés du cloud (ce que les pages pricing ne disent pas)

| Coût caché | Mécanisme | Ordre de grandeur | Parade |
|---|---|---|---|
| **Egress** | Sortie de données (checkpoints, datasets) | 0,05–0,09 $/Go (AWS) ; 0 $ (RunPod, Lambda, Vast) | Garder les artefacts lourds chez le même provider |
| **Over-provisioning PTU** | Débit réservé payé même inutilisé | 50–300 $/h selon modèle (à vérifier) | PTU seulement si SLA de latence contractuel |
| **Support** | Business/Enterprise obligatoire au-delà d'un volume | 100–15 000 $/mois (AWS) | Négocier dès 10 k$/mois de spend |
| **Version drift** | Le provider change les poids → tes evals/prompt cassent | Coût en temps de re-validation | Épingler les versions datées quand dispo ; rejouer l'eval |
| **Délistage** | Modèle retiré → migration d'urgence | Jours d'ingénierie | Noms logiques (LiteLLM) + 2 providers par cas d'usage |
| **Retry storms** | Boucle de retry sur panne → facture ×N | Jusqu'à ×10 en cas de bug | Backoff + jitter + plafond de tentatives (§Annexe D) |
| **Shadow AI** | Chaque équipe sa clé, aucun contrôle | +30–200 % de spend non piloté | Clés virtuelles centralisées dès le jour 1 |
| **Tokens invisibles** | Templates, tools, images, raisonnement comptés | ×2–5 vs le texte visible | Mesurer le ratio tokens/prompt réel (spend logs) |

---

## 14. Plan 90 jours : du zéro au hybride opérationnel

| Semaine | Action | Livrable |
|---|---|---|
| 1–2 | Inventaire : volumes, cas d'usage, sensibilité des données ; ouvrir comptes (OpenRouter + 1 provider vitesse + 1 provider prix) | Matrice besoins/providers |
| 3–4 | LiteLLM en Docker + Postgres ; 2 noms logiques (`triage`, `raisonnement`) ; 1 app pilote branchée | Gateway en pré-prod |
| 5–6 | Clés virtuelles + budgets ; dashboards Grafana ; alertes spend | Pilotage financier |
| 7–8 | Fallbacks testés (simuler une panne) ; runbook écrit ; DPA signés | Résilience validée |
| 9–10 | Ollama sur 1 poste + 1 serveur (Qwen3-8B/32B, Nomic embeddings) ; eval métier 50 questions | Preuve de concept locale |
| 11–12 | vLLM (si GPU dispo) ou Mac Studio ; RAG pilote sur doc interne (Qdrant + reranker) | RAG local pilote |
| 13 | Bascule progressive : `local` par défaut sur les flux sensibles ; escalade cloud sur confiance basse | Hybride en prod |

**Budget indicatif du plan** : ~500–2 000 $ de crédits API + 0–8 000 $ de GPU selon l'existant. Le poste le plus cher reste le temps d'ingénierie (compter 0,5 ETP sur 3 mois).

---

## Annexe H — Matrice : quel modèle pour quel usage (sept. 2026)

| Usage | 1er choix local | 1er choix cloud pas cher | 1er choix qualité max |
|---|---|---|---|
| Chat interne / Q&R doc | Qwen3-32B Q4 (vLLM) | GPT-OSS-120B via DeepInfra/Groq | Claude Opus 5.5 / GPT-6 Terra |
| RAG génération (FR) | Qwen3-32B Q4 / Llama 3.3 70B Q4 | DeepSeek V4 Pro (Fireworks/Together) | Claude Sonnet 5 / GPT-5.6 Terra |
| Code (complétion) | Qwen2.5-Coder-14B Q4 | Kimi K2.6 (Moonshot/K2-Think) | Claude Sonnet 5 / Devstral 2 |
| Code (agents) | Devstral 2 (vérifier licence) | GLM-5 (Zhipu) / MiniMax M2.7 | Claude Opus 5.5 / GPT-5.6 Terra |
| Raisonnement (maths, dimensionnement) | DeepSeek-R1-Distill-32B / GPT-OSS-20B | DeepSeek V4 Pro | GPT-6 Terra / Fable 5 |
| Embeddings (FR) | Qwen3-Embedding-8B / Nomic | text-embedding-3-small (0,02 $/1M) | Voyage-4 / Gemini |
| Classification / extraction bulk | Qwen3-8B Q4 (Ollama) | Llama 4 Scout (Groq, rapide) | — (inutile ici) |
| Vision (doc scannés) | Qwen3-VL-32B / Gemma 4 (MLX) | GPT-OSS-120B (multimodal) | GPT-6 / Fable multimodal |
| Temps réel (voix, jeux) | — (trop lent) | Groq / Cerebras | — |
| Air-gap total | Llama 3.3 70B Q4 / Qwen3-32B Q4 | — | — |

---

## Annexe I — 5 fiches matérielles détaillées (prix sept. 2026, à vérifier à l'achat)

### I.1. « Le dev » — tour existante + RTX 4090 24 Go (occasion ~2 150 $)

- Fait : 8B–14B full-GPU à 100–130 tok/s ; 32B Q4 (~22 Go) confortable ; embeddings ; 70B en offload partiel (lent, ~5–10 tok/s).
- Conso : ~450 W en charge. Idéal pour : dev RAG, classification bulk, tests.
- Limite : 24 Go = mur pour le 70B ; crise DRAM : prix occasion +24 % depuis mars.

### I.2. « Le sprinteur » — tour + RTX 5090 32 Go (~2 500–3 800 $ marché)

- Fait : tout ce que fait la 4090, ~1,5× plus vite (1 792 Go/s) ; 32B Q8 possible (~38 Go ? non — 32 Go : Q4 confort, Q8 limite).
- **Ne fait pas** : 70B Q4 full-GPU (43 Go > 32 Go). Le record de vitesse <32B, pas de capacité.
- Conso : 575 W — prévoir alim 1 000 W+ et refroidissement sérieux.

### I.3. « Le serveur d'équipe » — 2× RTX 4090 48 Go (~4 500 $ occasion) + vLLM

- Fait : 70B Q4 servi à 5–10 utilisateurs (15–40 tok/s/utilisateur) ; 32B AWQ très rapide ; 2 modèles simultanés.
- Conso : ~900 W. Bruit/chaleur : local technique obligatoire.
- Alternative pro : 1× RTX 6000 Ada 48 Go (~7 000 $, à vérifier) — une seule carte, garantie pro.

### I.4. « Le silencieux » — Mac Studio M4 Max 128 Go (~3 800 $) ou M5 Ultra 96 Go (dès 5 499 $)

