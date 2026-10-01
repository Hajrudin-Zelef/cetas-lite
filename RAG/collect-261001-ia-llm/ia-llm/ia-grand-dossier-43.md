---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-43
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Groq", "Hugging Face", "OpenAI", "OpenRouter", "United States", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agents", "awq", "benchmarks", "claude", "cyber", "deepseek", "embedding", "embeddings", "fp8", "gpu", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3402, 3510]
sha256: 5996a0c762c7b59cab04825377de117604b43f7f18a3e14c2adb104b6f6bfb1e
---

# IA — Le grand dossier

*Document clos le 27/09/2026 — 2000+ lignes. Toute décision d'achat ou d'engagement budgétaire : re-vérifier prix et disponibilités sur les pages officielles.*

---

## 15. Pour aller plus loin : ressources officielles et communautés

### 15.1. Documentation officielle (à bookmarker)

| Ressource | URL de départ | Pourquoi |
|---|---|---|
| LiteLLM (docs) | docs.litellm.ai | La référence gateway : config, fallbacks, budgets — lire la doc de TA version |
| OpenRouter (docs) | openrouter.ai/docs | Routage, `:free`, BYO-key, paramètres provider |
| Ollama (docs) | docs.ollama.com | Modelfile, API, variables d'environnement |
| llama.cpp (GitHub) | github.com/ggml-org/llama.cpp | `llama-server`, quantifications, builds |
| vLLM (docs) | docs.vllm.ai | Serveur, AWQ/FP8, LoRA, metrics |
| LM Studio | lmstudio.ai/docs | Serveur local GUI |
| Open WebUI | docs.openwebui.com | Chat interne, RAG, pipelines |
| Hugging Face Hub | huggingface.co/models | Le catalogue des poids ouverts (+ licences) |
| Artificial Analysis | artificialanalysis.ai | Benchmarks indépendants (prix, vitesse, qualité) |
| LLM Price Tracker | llm-price-tracker.vercel.app | Comparateur de prix API communautaire |

### 15.2. Veille (rythme recommandé)

- **Hebdomadaire** : Artificial Analysis (nouveaux modèles/benchmarks), releases LiteLLM/Ollama/vLLM sur GitHub.
- **Mensuelle** : pages pricing de tes 3–5 providers principaux (ça bouge), licences des modèles ouverts que tu utilises.
- **Trimestrielle** : recalcul du seuil local/cloud (§8.2), revue des budgets, exercice de bascule des fallbacks.
- **Alertes utiles** : GitHub Watch sur `BerriAI/litellm`, `ollama/ollama`, `vllm-project/vllm` (releases uniquement).

### 15.3. Règle de fin

Ce dossier est un **point de départ daté du 27/09/2026**, pas une vérité permanente. Le marché de l'IA infra change tous les trimestres : ce qui ne change pas, c'est la méthode — **noms logiques, evals métier, budgets plafonnés, fallbacks testés, licences notées**. Avec ça, tu peux changer de provider, de modèle ou de GPU sans réécrire ton SI.

---

*Fin de la Partie 2 — Providers, passerelles & IA locale (2000+ lignes).*

### 15.4. Prochaines étapes concrètes (pour ton RAG)

1. **Semaine 1** : ouvre un compte OpenRouter (+20 $ de crédits) et un compte Groq (free tier) ; compare `qwen3:8b` local (Ollama) vs GPT-OSS-120B (Groq) sur 20 questions de ta doc.
2. **Semaine 2** : déploie LiteLLM en Docker avec 2 noms logiques (`triage`, `raisonnement`) ; branche ton script RAG dessus au lieu d'appeler OpenAI en direct.
3. **Semaine 3** : ajoute `embeddings` en logique double (text-embedding-3-small actuel + `nomic-embed-text` local) ; mesure l'écart de précision@k sur ton corpus.
4. **Mois 2** : si le volume dépasse ~3M tokens/jour, chiffre le seuil (§8.2) et envisage une RTX 4090 d'occasion ou un Mac Studio selon ton besoin (vitesse vs silence/capacité).
5. **En continu** : tiens le tableau de suivi mensuel (§P.1) ; c'est lui qui transformera ce dossier en pilotage réel.

*Le meilleur provider est celui que tu peux quitter en une ligne de YAML.*

---

## Annexe S — Référence des variables d'environnement (template `.env.example`)

```bash
# ── Gateway ──
LITELLM_MASTER_KEY=                    # 32+ caractères, générée (openssl rand -hex 32)
LITELLM_TEAM_KEY=                      # clé virtuelle d'équipe (créée dans l'UI :4000/ui)
POSTGRES_PASSWORD=
DATABASE_URL=postgresql://litellm:${POSTGRES_PASSWORD}@postgres:5432/litellm
# ── Providers cloud (une seule variable par provider utilisé) ──
OPENAI_API_KEY= ANTHROPIC_API_KEY= GOOGLE_API_KEY=
GROQ_API_KEY= CEREBRAS_API_KEY= OPENROUTER_API_KEY=
MISTRAL_API_KEY= DEEPSEEK_API_KEY= TOGETHER_API_KEY= FIREWORKS_API_KEY=
# ── Local ──
VLLM_API_KEY=                          # Bearer pour :8000
OLLAMA_HOST=127.0.0.1:11434            # jamais 0.0.0.0 sans reverse proxy
OLLAMA_KEEP_ALIVE=30m OLLAMA_NUM_PARALLEL=4
# Règle : ce fichier (.env.example) va en git SANS valeurs ; le vrai .env ne va JAMAIS en git.
```

---

# PARTIE 3 — Débats, dangers, Claude, stacks, communautés

*Grand dossier IA — Partie 3. Rédigé le 27 septembre 2026.*
*Public : sysadmin / chef de service systèmes & énergies. Ton : neutre, factuel, opinions toujours attribuées.*

> **Avertissement de lecture.** Cette partie expose des positions contradictoires *sans prendre parti*.
> Chaque opinion est attribuée à son auteur. Les chiffres datés sont issus de sources identifiées ;
> tout fait incertain est marqué « à vérifier ». Les volumes 1 et 2 (modèles, outillage) sont supposés lus.

---

## Sommaire de la partie 3

| Section | Contenu |
|---|---|
| **1. Le grand débat** | Productivité (études chiffrées, attribuées) · emploi (les deux camps) · créativité et propriété intellectuelle · coût environnemental · éducation |
| **2. Dangers et alertes** | Sécurité des modèles (jailbreaks, prompt injection, exfiltration) · mésusages (deepfakes, désinformation, cyber) · voix critiques (qui dit quoi) · régulation au 27/09/2026 (UE, US, Chine) |
| **3. Focus Claude / Anthropic** | Histoire factuelle (fondation, levées) · axe sécurité (Constitutional AI expliquée) · ce qui distingue Claude (factuel) · gamme de modèles (synthèse, renvoi aux volumes IA) |
| **4. Les stacks techniques** | Stack entraînement · stack inférence · stack RAG (exemple chaîné complet) · stack agents · schémas textuels + code réel |
| **5. Les communautés** | Slack · Discord · forums (Reddit, Hacker News) · Hugging Face · GitHub · newsletters et comptes à suivre (noms réels) |
| **6. 20 conseils concrets** | Rester à jour quand on est sysadmin |
| **7. Glossaire** | 30+ termes |
| **8. Quiz** | 10 questions + corrigé |

---

## 1. Le grand débat pour/contre l'IA

### 1.1. Position du dossier

L'IA générative est la première technologie depuis Internet à susciter simultanément des promesses
de productivité mesurées en dizaines de pourcents et des alertes d'extinction formulées par ses
propres inventeurs. Cette section juxtapose les chiffres et les positions **sans les départager** :
à toi de te faire ton opinion. Règle d'hygiène intellectuelle appliquée ici : un chiffre n'est
crédible que s'il est attribué (auteur, méthode, date).

### 1.2. Productivité : les études qui comptent

#### Ce que disent les expériences de terrain

