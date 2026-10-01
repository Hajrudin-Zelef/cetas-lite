---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-5
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "Cohere", "DeepSeek", "Fireworks AI", "Google", "Groq", "Hugging Face", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["agent", "cohere", "deepseek", "embeddings", "gemini", "glm", "inference", "kimi", "llama", "mistral", "multimodal", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [475, 567]
sha256: 2211f46fe02b20059a26fe9752c03e1bbfca93a214aef4e014cd862ddb684d65
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Fournisseur | Ce qui est gratuit | Limites constatées (sept 2026) | Modèles notables |
|---|---|---|---|
| **Groq** | Tous les modèles du catalogue | 30 RPM ; 1K–14,4K req/j ; 8K–30K TPM selon modèle | gpt-oss 120B/20B, Qwen3.6-27b, Whisper |
| **Cerebras** | Tous les modèles | **1M tokens/jour** ; ~5 RPM ; 8K ctx cap (temporaire) | gpt-oss-120b, llama-3.1-8b |
| **Google AI Studio (Gemini)** | Famille Flash / Flash-Lite | ~15–30 RPM ; ~1 500 req/j (Flash) ; 1M TPM | Gemini 3.x Flash, Gemma |
| **OpenRouter** | Modèles `:free` | 50 req/j (1 000/j si $10 déjà chargés) ; 20 RPM | Nemotron 3 Ultra 550B, GLM 5.2, gpt-oss-20b |
| **Mistral La Plateforme** | Plan « Experiment » | ~1 Md tokens/mois mais ~2 RPM ; TPM 25K–20M selon modèle | Mistral Large, Medium 3.5, Codestral |
| **Mistral Codestral** | Clé dev Codestral | 30 RPM, 2 000 req/j | Codestral (code) |
| **NVIDIA NIM** | Build / API hosted | ~40 RPM ; crédits à l'inscription (~1 000) | Llama, Nemotron, Qwen, GLM |
| **Cohere** | Clé d'essai | 20 RPM ; ~1 000 req/mois ; usage non commercial | Command A/R, Aya |
| **Cloudflare Workers AI** | Edge inference | 10 000 neurons/jour | Llama 3.x, Mistral 7B, Qwen, **BGE embeddings** |
| **Hugging Face** | Inference Providers | $0.10/mois de crédits (Free) ; $2/mois (PRO) | 200+ modèles via partenaires |
| **GitHub Models** | Playground + API sandbox | Rate-limité ; **pas de prod** | GPT-5.x, Llama, Mistral, Qwen |
| **DeepSeek** | Inscription | **5M tokens offerts** (one-shot) | DeepSeek V4.1 |
| **SambaNova** | Free tier | 20–480 RPM ; $5 de crédit d'essai | Llama 4 Maverick, DeepSeek R1 |
| **OVH AI Endpoints** | Bêta | 2 RPM anonyme / 400 RPM authentifié | Qwen3, 20+ poids ouverts |
| **Together / Fireworks** | Crédits d'essai | ~$1–5 à l'inscription | 200+ poids ouverts |
| **Chutes.ai** | Communautaire | Selon dispo (réseau Bittensor) | DeepSeek-R1, Llama 70B |

**Pas de free tier permanent** : OpenAI et Anthropic (crédits d'essai temporaires uniquement).

## 39. Gemini gratuit : ce qu'il faut savoir (sept 2026)

- Depuis **avril 2026**, seuls les modèles **Flash / Flash-Lite** sont gratuits ; la série **Pro est payante**.
- Quotas constatés : ~15–30 RPM, ~1 500 req/jour sur Flash, jusqu'à 1M TPM ; **fenêtre 1M tokens** sur Flash — le plus grand contexte gratuit du marché.
- Depuis **juin 2026**, les clés API « sans restriction » sont bloquées : passer par AI Studio avec restrictions d'usage.
- **Point de vigilance majeur** : les prompts du free tier **peuvent servir à améliorer les produits Google** (revue humaine possible). Ne jamais y envoyer de données pro sensibles.
- Cas d'usage roi : résumer de très longs documents (1M ctx gratuit), multimodal (images/PDF) gratuit.

## 40. Mistral gratuit : plan Experiment

- Le plan **Experiment** de La Plateforme couvre **tous les modèles** (y compris Large et Codestral), de l'ordre d'**1 Md tokens/mois**, mais bridé à **~2 req/min** — c'est fait pour tester, pas pour du temps réel.
- **Données** : Mistral peut **entraîner sur les données du free tier, avec opt-out** dans les réglages — à désactiver explicitement si tu testes avec du contenu sensible. Rétention configurable (jamais → 1 an).
- **Codestral** (le modèle code) a sa propre clé dev gratuite : 30 RPM / 2 000 req/j — excellent pour de l'autocomplétion perso.
- Vérification téléphone requise sur les deux.

## 41. Autres gratuits notables (détails)

- **GitHub Models** : avec n'importe quel compte GitHub, playground + API sandbox sur 100+ modèles (GPT-5.x, Llama, Mistral, Qwen). **Interdit en production**, quotas journaliers par palier. Parfait pour tester un modèle fermé avant de payer.
- **NVIDIA NIM** : ~40 RPM après vérification téléphone, crédits à l'inscription ; catalogue de gros poids ouverts (Nemotron, Kimi, GLM). Endpoint `https://integrate.api.nvidia.com/v1` compatible OpenAI.
- **Cohere** : 1 000 req/mois, **usage non commercial uniquement** sur la clé d'essai — intéressant pour tester le **reranking** (modèles Rerank) et les embeddings multilingues avant d'acheter.
- **Cloudflare Workers AI** : 10 000 neurons/jour, ~80 modèles dont **BGE pour les embeddings** — un backend d'embeddings gratuit potentiel pour ton RAG (à évaluer qualité/latence).
- **SambaNova** : puces RDU propriétaires, free tier 20–480 RPM + $5 d'essai ; Llama 4 Maverick, DeepSeek R1 hébergés.

## 42. Le coût caché du gratuit : TES DONNÉES (à lire absolument)

| Fournisseur | Entraînement sur tes prompts ? | Rétention |
|---|---|---|
| Groq | Non | Jusqu'à 30 jours (ZDR selon offre) |
| Cerebras | Non communiqué clairement (sept 2026) | À vérifier |
| Google AI Studio (free) | **Oui possible** (amélioration produits, revue humaine) | Selon CGU |
| Mistral Experiment | **Oui, avec opt-out** | Configurable (jamais → 1 an) |
| OpenRouter | Selon le provider final (`data_collection: "deny"` pour filtrer) | Selon provider |
| DeepSeek (API) | **Oui, avec opt-out** | « Aussi longtemps que nécessaire » |
| OpenAI / Anthropic (payant) | **Non par défaut** | 30 jours |

**Règle** : gratuit = tu es le produit (tes prompts). Données perso/projet sensible → local (Ollama) ou offre payante avec ZDR. C'est non négociable en contexte d'entreprise.

## 43. Stratégie d'empilement « $0 » (concrète)

Ordre d'appel pour une appli perso :

1. **Cache local** (SQLite/Redis) : question déjà vue → 0 appel.
2. **Groq free** (`gpt-oss-20b`) : triage, résumés, extractions.
3. **Cerebras free** (`llama-3.1-8b`, 1M tok/j) : volume simple.
4. **Gemini Flash free** : gros documents (1M ctx), multimodal.
5. **OpenRouter `:free`** (`nemotron-3-ultra-550b`, `glm-5.2`) : quand il faut du muscle gratuit.
6. **Mistral Experiment** : modèles FR / Mistral Large en test.
7. **Payant ciblé** : uniquement ce qui échoue aux étapes 1–6.

Avec FreeLLMAPI (section 36), les étapes 2–6 tiennent derrière **un seul endpoint** avec failover auto. Pour ton RAG : embeddings en local (TEI, partie G) + génération sur cette pile = **facture mensuelle proche de 0**.

## 44. Checklist écosystème gratuit

- [ ] Créer les comptes : Groq, Cerebras, Google AI Studio, OpenRouter, Mistral, NVIDIA, GitHub.
- [ ] Noter chaque quota dans un tableau (RPM / req/j / TPM) + date de vérification.
- [ ] Déployer FreeLLMAPI en Docker, y enregistrer les clés, tester le failover (couper une clé → vérifier la bascule).
- [ ] Coder le **retry 429 avec backoff exponentiel** (obligatoire partout en gratuit).
- [ ] Désactiver l'entraînement sur données là où l'opt-out existe (Mistral, DeepSeek).
- [ ] Marquer les données sensibles « local only » dans ton code (garde-fou).

---

# PARTIE F — DEEPSEEK HARNESS + SELF-HOSTING LLM LOCAL

## 45. « Harness » : les deux sens (important)

Le mot **harness** a deux sens qu'il ne faut pas mélanger :
1. **Sens générique** (industrie) : le « harnais » autour d'un modèle — scaffolding, outils, gestion du contexte, retries, évals, permissions, observabilité. « Le modèle se loue, le harnais se construit. » C'est ce que tu construis autour de ton RAG.
2. **DeepSeek Harness** (outil précis, sorti le **13 août 2026**) : l'agent de code open source de DeepSeek, commande `dsh`. Les deux sections suivantes couvrent l'outil ; le reste de la partie F couvre le **self-hosting** (faire tourner DeepSeek / un LLM ouvert en local), qui était l'autre interprétation possible — les deux sont traitées, rien n'est laissé dans le flou.

## 46. DeepSeek Harness : identification et état (sept 2026)

