---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-13
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cerebras", "DeepSeek", "Google", "Groq", "Hugging Face", "Microsoft", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "vLLM"]
dates: ["2026-08-13", "2026-09-10", "2026-09-27"]
keywords: ["gpu", "attention", "awq", "claude", "context window", "deepseek", "embeddings", "gemini", "gemini 4", "gguf", "gpt-5.6", "gpt-6"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1604, 1700]
sha256: 41fe7a119d8b8b72fddd6600f62b1129557ee5c75f77c730435797aec08874d4
---

# PARTIE K — NOMS NON IDENTIFIÉS (HONNÊTETÉ)

## 111. Noms recherchés mais NON TROUVÉS au 27/09/2026

Recherche web sérieuse effectuée ; ces noms n'ont pas pu être confirmés comme des produits/services réels à cette date. **Ne pas les utiliser** dans des décisions ; re-vérifier avant toute action :

| Nom | Statut | Détail |
|---|---|---|
| **Claude Sonnet 4.8** | **NON TROUVÉ (sept 2026)** | Opus 4.8 existe (mentionné par Anthropic/GitHub), Sonnet 4.8 : aucune annonce trouvée. Ne pas confondre. |
| **Claude Haiku 5 / 5.5** | Partiellement | Haiku 4.5 confirmé ($1/$5). « Haiku 5.5 » évoqué comme « à venir » par la presse — pas de fiche produit trouvée. |
| **Gemini 3.5 Pro / Gemini 4** | **NON TROUVÉ (sept 2026)** | Famille Gemini 3.x (Flash) confirmée ; pas de « 3.5 Pro » ni « 4 » dans les résultats. |
| **Veo 4** | **NON TROUVÉ (sept 2026)** | Aucune annonce trouvée (vidéo Google). |
| **DeepSeek V4.1 Pro (annonce)** | **NON TROUVÉ (sept 2026)** | V4.1 Flash (10/09/2026) et V4 Pro (13/08/2026) confirmés ; pas de « V4.1 Pro ». |
| **GPT-6 Terra** | **NON CONFIRMÉ (sept 2026)** | « gpt-5.6-terra » ($2/$12) apparaît dans un comparatif tiers ; « GPT-6 Terra » : recherche antérieure infructueuse. À re-vérifier sur la page prix OpenAI. |
| **Llama 4 Behemoth** | **NON TROUVÉ (sept 2026)** | Llama 4 Scout/Maverick confirmés ; pas de Behemoth. |
| **Poolside Malibu** | **NON TROUVÉ (sept 2026)** | Poolside Laguna S 2.1 confirmé (OpenRouter :free) ; pas de « Malibu ». |
| **mxbai-rerank-v3.1-listwise** | **NON TROUVÉ (sept 2026)** | Famille mxbai-rerank existe en principe, cette version exacte non confirmée. |
| **ctxl-rerank / ReasonRank** | **NON TROUVÉ (sept 2026)** | Aucune fiche produit trouvée. |
| **Mythos (Anthropic)** | Existe mais **non public** | Évoqué comme modèle interne (~40 orgs) — pas d'accès public, inutile pour ton usage. |

**Noms en revanche IDENTIFIÉS** (contrairement au doute initial) : `freellmapi` = FreeLLMAPI (section 35) ; `harness` (DeepSeek) = DeepSeek Harness `dsh` (section 46).

---

# PARTIE L — GLOSSAIRE

## 112. Glossaire (A–M)

- **AWQ** : quantization 4-bit GPU (Activation-aware Weight Quantization).
- **Batching (dynamique/continu)** : regrouper les requêtes pour remplir le GPU ; le batching continu (vLLM) insère les nouvelles requêtes sans attendre la fin du batch.
- **BF16 / FP16** : flottants 16 bits — précision standard d'inférence.
- **Candle** : moteur Rust de Hugging Face (utilisé par TEI).
- **Context window (ctx)** : nombre max de tokens (entrée + sortie) par requête.
- **Embeddings** : vecteurs numériques représentant le sens d'un texte — la base du RAG.
- **Fallback** : bascule automatique sur un provider/modèle de secours.
- **GGUF** : format de modèle quantifié pour llama.cpp/Ollama.
- **GPTQ** : quantization 4-bit GPU (plus ancienne qu'AWQ).
- **KV cache** : mémoire des clés/valeurs d'attention — grandit avec contexte × batch.
- **LPU** : Language Processing Unit — puce d'inférence de Groq.
- **Matryoshka (MRL)** : embeddings à dimensions emboîtées — on peut tronquer le vecteur (1024→512) sans réentraîner.
- **MoE** : Mixture of Experts — seuls quelques « experts » sont actifs par token (ex. 671B total, ~37B actifs chez DeepSeek).

## 113. Glossaire (N–Z)

- **PagedAttention** : gestion mémoire paginée du KV cache (vLLM) — moins de gaspillage.
- **Prefix caching** : réutiliser le calcul du début de prompt (système + chunks RAG récurrents).
- **Quantization** : réduire la précision des poids (16→8→4 bits) pour diviser la mémoire.
- **Rate limit** : quota (req/min, tokens/min, req/jour) — au-delà : HTTP 429.
- **Reranker** : modèle qui re-classe des documents candidats par pertinence fine (cross-encoder).
- **RPM / RPD / TPM / TPD** : requêtes ou tokens par minute / jour.
- **Serverless (inférence)** : API sans machine à gérer, facturée à l'usage.
- **Spot / interruptible** : instance louée moins chère, préemptible par l'hôte.
- **TEI** : Text Embeddings Inference — serveur d'embeddings de Hugging Face.
- **Tiers** : niveau d'offre (free tier, pay-as-you-go…).
- **Token** : unité de texte (~4 caractères en anglais, ~2–3 en français) — unité de facturation.
- **TPS** : tokens par seconde — mesure de débit.
- **TTFT** : time to first token — latence perçue.
- **vLLM** : serveur d'inférence LLM haute performance (UC Berkeley).
- **WSE** : Wafer-Scale Engine — puce géante de Cerebras.
- **ZDR** : Zero Data Retention — garantie de non-conservation des données.

---

# QUIZ (10 questions + réponses)

## 114. Quiz — questions

1. Quelle est la différence fondamentale entre OpenRouter d'une part, Groq et Cerebras d'autre part ?
2. Cite les ordres de grandeur du free tier Groq (RPM) et Cerebras (tokens/jour).
3. Pourquoi le free tier de Google AI Studio est-il déconseillé pour des données professionnelles sensibles ?
4. Que signifie `provider: {"allow_fallbacks": true}` dans un appel OpenRouter ?
5. Quelle VRAM faut-il prévoir (ordre de grandeur) pour servir un modèle 32B en Q4_K_M avec 16K de contexte sur un seul GPU ?
6. Quelle est la différence entre AWQ et GGUF Q4_K_M en termes de matériel requis ?
7. Pourquoi faut-il activer `--enable-prefix-caching` dans vLLM pour un RAG ?
8. Qu'est-ce que FreeLLMAPI et quelle est sa limite d'usage affichée par le projet ?
9. Quel est l'intérêt d'ajouter un reranker (ex. bge-reranker-v2-m3) entre la récupération et la génération dans un RAG ?
10. Sur Vast.ai, quelle est l'erreur n°1 qui fait exploser la facture, et comment l'éviter ?

## 115. Quiz — réponses

1. OpenRouter est une **passerelle logicielle** (route vers des fournisseurs tiers, 300+ modèles) ; Groq et Cerebras sont des **fondeurs** (puces LPU / WSE) qui vendent leur propre inférence ultra-rapide sur des poids ouverts.
2. Groq : **~30 req/min** (quotas par modèle, 1K–14,4K req/jour) ; Cerebras : **1M tokens/jour** (~5 RPM).
3. Parce que Google peut **utiliser les prompts du free tier pour améliorer ses produits** (revue humaine possible) — banni pour du sensible.
4. Si le fournisseur routé échoue ou est rate-limité, OpenRouter **bascule automatiquement** sur un autre fournisseur du modèle — sans changer ton code.
5. Poids Q4 ≈ 20 Go + KV cache 16K ≈ 3–5 Go + marge système ≈ **24 Go** → une RTX 4090 24 Go (juste) ; confortable à 32 Go.
6. **AWQ = GPU NVIDIA uniquement** (4-bit, via vLLM) ; **GGUF Q4_K_M = CPU et GPU** (llama.cpp/Ollama, offload partiel possible).
7. Parce que le début du prompt (instructions système + chunks récurrents) est identique d'un appel à l'autre : le prefix caching **évite de le ré-encoder**, ce qui divise latence et calcul.
8. Un **proxy open source auto-hébergé** qui agrège les tiers gratuits de ~28 fournisseurs derrière un endpoint OpenAI-compatible avec failover auto ; le projet précise **« personal experimentation only »** (pas de prod).
9. La similarité cosinus sur embeddings est rapide mais grossière ; le cross-encoder **relit vraiment** query+document et re-classe finement → top-5 bien meilleur pour un coût local de ~10 ms/doc.
10. **Oublier de détruire l'instance** (`stopped` ≠ `destroyed`) → script d'auto-destruction en fin de job + solde prépayé limité.

---

# PIÈGES RÉCAPITULATIFS (15+)

## 116. Les 18 pièges à ne jamais oublier

