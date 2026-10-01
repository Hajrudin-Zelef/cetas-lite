---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-11
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agent", "agents", "agi", "awq", "aws", "benchmarks", "blackwell", "capex", "claude", "compute", "deepseek", "embedding"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [672, 769]
sha256: 5c80f3e4ba0f4664a304cd32d41a6427f255ecf0192a77e39dc20885ea44dfec
---

# IA — Le grand dossier

### 2026 (janvier → 27 septembre)
- **E** — Février : Mistral lève/runway — Série D **3 Md€** (septembre, Samsung + EQT, valorisation > 21 Md€) ; Le Chat → « Vibe » (presse, « à vérifier »).
- **M** — **DeepSeek V4** (mai : -75 % définitifs sur les prix API) ; **GPT-5.5**, **Claude Opus 4.7** (« à vérifier » les numéros exacts au jour près) ; **Mistral Large 3 / Magistral** ; **Qwen3**, **Kimi K2** (« à vérifier »).
- **E** — L'inférence devient « too cheap to meter » sur les modèles courants ; la valeur migre vers les **agents** et les volumes.
- **P** — Lovelace et al. (2026) : pénalité de surparamétrisation sous contrainte de données ; Weng (2026) : synthèse des « fissures » des lois de scaling (Techtimes, juin 2026).
- **I** — GB300/Rubin (NVIDIA), Ironwood v7 (Google TPU), Trainium (AWS) : la bataille du **silicium d'inférence** ; pénurie HBM ; datacenters en attente de raccordement électrique.
- **État du débat** : consensus mou — le **training naïf** ne scale plus comme avant (camp « mur » validé sur ce point), mais le **système global** (inférence, RL, efficacité, agents) continue de progresser vite (camp « continuation » validé sur ce point). Le désaccord porte désormais sur **la vitesse vers l'AGI**, pas sur l'existence du progrès.

---

## 6. Ce que ça change concrètement pour un sysadmin

> Zelef est chef de service systèmes & énergies et construit son RAG. Cette section traduit les sections 1-5 en décisions d'infrastructure, de budget et de compétences.

### 6.1. Le basculement fondamental : inference > training

**Avant (2020-2023) :** la valeur était dans **l'entraînement** — seuls quelques labos pouvaient se payer des runs à 100 M$. Un sysadmin d'entreprise n'avait rien à entraîner.

**Maintenant (2024-2026) :** la valeur est dans **l'inférence** — servir des modèles, à grande échelle, au moindre coût par token. Et ça, c'est **le métier du sysadmin** : dimensionnement, haute dispo, supervision, optimisation des coûts, sécurité.

| Ère training (2020-2023) | Ère inference (2024-2026+) |
|---|---|
| CAPEX : clusters d'entraînement géants | OPEX : flotte de serving, facturation au token |
| Quelques runs énormes, batch | Des milliards de requêtes, temps réel |
| Métrique : loss, benchmarks | Métriques : **latence (TTFT, tokens/s), coût/1M tokens, dispo** |
| Compétence rare : distributed training | Compétence demandée : **serving (vLLM, TGI, TensorRT-LLM), KV-cache, batching** |
| Données : corpus d'entraînement | Données : **prompts, RAG, logs, feedback** |

**TTFT** = Time To First Token (temps avant le premier token — critique pour l'UX chat). Le débit en tokens/s détermine le dimensionnement.

### 6.2. Hardware : quoi acheter, quoi ne pas acheter

**Règle n° 1 : ne pas entraîner, servir.** Sauf cas très particulier (fine-tune LoRA sur données métier), une entreprise n'entraîne plus de modèle from scratch. Le budget GPU va au **serving**.

**Ordres de grandeur VRAM (poids seuls, sans KV-cache) :**

| Modèle | Quantification | VRAM poids | Matériel typique |
|---|---|---|---|
| 7-8B (Llama 3.1 8B, Mistral 7B, Qwen 7B) | Q4 (GGUF) | ~5 Go | Laptop, RTX 3060 12 Go |
| 7-8B | FP16 | ~16 Go | RTX 4070 Ti / 4080, L4 |
| 32B (Qwen 32B) | Q4 | ~20 Go | RTX 4090 24 Go |
| 70B (Llama 3.3 70B) | Q4 | ~40 Go | 2× RTX 4090, ou 1× A100 80 Go / H100 |
| 70B | FP16 | ~140 Go | 2× H100, ou 4× A100 40 Go |
| 671B MoE (DeepSeek-V3/R1) | FP8 | ~700 Go+ | 8× H100/H200 (nœud complet) |

Formule mémo : **VRAM ≈ 2 octets × paramètres en FP16** (÷4 en Q4). Ajouter **20-40 % pour le KV-cache** (proportionnel à : batch × longueur de contexte × couches). Le contexte long (128k-1M tokens) est le **vrai mangeur de VRAM** — un 8B en contexte 1M peut exiger plus de VRAM qu'un 70B en contexte 8k.

**Recommandations pratiques :**

1. **RAG d'entreprise (cas de Zelef) :** un **8B-32B quantifié** en local (ou via API) suffit pour 95 % des cas : la qualité du RAG vient du **retrieval** (chunking, embeddings, reranking), pas de la taille du générateur. Investir dans la **qualité du corpus** plutôt que dans un 70B.
2. **Embeddings :** `text-embedding-3-small` via API, ou **BGE/E5/Qwen-embeddings** en local (quelques Go de VRAM) — trivial à servir.
3. **Raisonnement local :** les **distillés R1** (7B/14B/32B) donnent du raisonnement « o1-class » sur une RTX 4090 — réserver le gros raisonnement API (o3, R1-671B) aux cas vraiment durs.
4. **Ne pas acheter de cluster d'entraînement** : le « 80/20 reversal » dit que les gros clusters généralistes risquent le sous-emploi. Préférer du **serving élastique** (scale-to-zero, spot/preemptible pour le batch).
5. **Énergie (lien direct avec le métier de Zelef) :** une baie de 8× H100 tire **~10 kW** ; un rack Blackwell complet dépasse les **100 kW** — on ne refroidit plus ça en air seul (eau/immersion). Le PUE, la contractualisation électrique et la thermographie deviennent des compétences IA à part entière. **L'IA est un problème d'énergéticien autant d'informaticien.**

### 6.3. Coûts : le modèle économique de l'inférence

**API (ordres de grandeur 2026, variables selon fournisseur — « à vérifier » au jour près) :**

| Classe de modèle | Entrée / 1M tokens | Sortie / 1M tokens |
|---|---|---|
| Léger (type GPT-4o-mini, Haiku) | ~0,15 $ | ~0,60 $ |
| Moyen (type GPT-4o, Sonnet) | ~2-3 $ | ~8-15 $ |
| Frontière raisonnant (type o3, Opus) | ~10-15 $ | ~40-60 $ |
| DeepSeek (casse les prix) | ~0,27-0,44 $ | ~0,87-2,19 $ |

**Règle de calcul :** 1M tokens ≈ 750 000 mots anglais (~500 000 mots français, ordre de grandeur). Un agent qui « réfléchit » (10k tokens de thinking par requête) coûte **10-50×** une requête simple : **le test-time compute est le nouveau poste de coût à piloter** (paramètre `reasoning_effort`, plafonds par cas d'usage, routage simple/complexe).

**Self-hosting vs API — grille de décision :**

| Critère | API | Self-hosting |
|---|---|---|
| Données sensibles / souveraineté | ⚠️ (DPA, zéro-rétention à négocier) | ✅ (on-premise, air-gap possible) |
| Coût à faible volume | ✅ | ❌ (GPU immobilisé) |
| Coût à fort volume (>100M tokens/mois) | ❌ (facture linéaire) | ✅ (coût fixe amorti) |
| Latence / SLA maîtrisé | ⚠️ (dépend du fournisseur) | ✅ |
| Maintenance / MLOps | ✅ (zéro) | ❌ (vLLM, drivers, pannes GPU) |
| Modèles frontières (raisonnement extrême) | ✅ (seul accès) | ⚠️ (distillés seulement) |

**Point de bascule empirique :** au-delà de ~50-100M tokens/mois sur des modèles moyens, le self-hosting d'un 8B-70B devient rentable — **à vérifier** avec vos tarifs locaux d'électricité et d'amortissement GPU.

### 6.4. Stack technique du serving moderne (à connaître)

```
Requête → [Gateway/LB] → [vLLM / TGI / TensorRT-LLM]
                              ├─ Continuous batching (remplit le GPU en continu)
                              ├─ PagedAttention (KV-cache en pages, ~0 gaspillage)
                              ├─ Quantification (AWQ/GPTQ/FP8)
                              └─ [RAG] → vector DB (Qdrant/pgvector) → reranker → prompt
```

- **vLLM** : le serveur d'inférence open-source de référence (UC Berkeley) — PagedAttention, continuous batching, OpenAI-compatible API.
- **Ollama / llama.cpp** : le « Docker du LLM local » — parfait pour dev/test et petits déploiements.
- **Vector DB** : pgvector (si déjà PostgreSQL — cas de Zelef), Qdrant, Milvus. Le choix dépend du volume : < 5M chunks → pgvector suffit.
- **Observabilité** : tracer les **tokens in/out par requête**, la latence P50/P99, le taux de cache-hit du préfixe (prompt caching) — ce sont les KPI du FinOps IA.

### 6.5. Compétences à développer (plan 6 mois, réaliste pour un chef de service)

