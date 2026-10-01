---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-31
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Baseten", "ByteDance", "Cerebras", "Cohere", "DeepSeek", "Fireworks AI", "Google", "Groq", "Hugging Face", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "Perplexity", "Together AI", "United States", "Z.ai", "vLLM", "xAI"]
dates: ["2026-08-21"]
keywords: ["agents", "apache", "astra", "aws", "bedrock", "claude", "cohere", "deepseek", "embedding", "embeddings", "fable 5", "fine-tuning"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2142, 2179]
sha256: d02d84b8fb9526ddf6ead6c8efbb4dd60ac4753e5c8e573cb1a269eb0ed8432f
---

# IA — Le grand dossier

| # | Fournisseur | Spécialité | Modèle économique | Positionnement |
|---|---|---|---|---|
| 1 | **OpenRouter** | Agrégateur : 400+ modèles (ouverts + GPT, Claude, Gemini) derrière **une seule clé API** ; routage par provider, modèles `:free` | Prix public du provider **sans marge/token** ; commission 5,5 % sur achats de crédits (min 0,80 $) ; BYO-key +5 % | Le couteau suisse : un seul SDK, bascule de modèle en une chaîne de caractères, fallbacks natifs |
| 2 | **LiteLLM** | **Gateway open source** auto-hébergée : 100+ providers derrière une API compatible OpenAI (`:4000`) ; v1.97.0 (21/08/2026) | Gratuit (MIT) en self-hosted ; offre managée/entreprise payante | Le standard de facto du routage en entreprise : fallbacks, load balancing, budgets, spend tracking |
| 3 | **Groq** | Inférence ultra-basse latence sur **LPU** (puce propriétaire) : 500–800 tok/s ; catalogue réduit (~6 modèles self-serve) | $/1M tokens ; ex. GPT-OSS-120B 0,15/0,60 $ ; free tier ~30 RPM | Temps réel (réécriture de requêtes, re-rankers, agents) ; a levé 350 M$ en août 2026 (valo 3,5 Md$) ; NVIDIA a licencié sa tech LPU |
| 4 | **Cerebras** | Inférence la plus rapide du marché sur **Wafer-Scale Engine** : 1 800–3 000 tok/s | $/1M tokens ; GPT-OSS-120B ~0,35/0,75 $ ; crédit d'essai 5 $ | Vitesse brute ; 8 modèles self-serve ; contextes parfois réduits vs natif |
| 5 | **Together AI** | Cloud complet des modèles ouverts : serverless (200+ modèles) + fine-tuning (LoRA/full) + endpoints dédiés + clusters GPU | $/1M tokens (0,03–4,50 $) ; dédié H100 dès ~3,99 $/h ; prépayé min 5 $ | « Le AWS des poids ouverts » : servir, tuner et scaler au même endroit |
| 6 | **Fireworks AI** | Inférence production sur moteur **FireAttention** : TTFT bas, function calling, JSON garanti | $/1M tokens (0,20–3,00 $ typique) ; Standard/Priority/Fast ; batch -50 % ; 1 $ de crédit offert | Agents en prod (Cursor, Notion, DoorDash, Quora clients) ; fiabilité > prix plancher |
| 7 | **DeepInfra** | **Prix plancher** sur ~77 modèles open-weight hébergés aux US (B200) | $/1M tokens (0,02–1,50 $) ; batch -50 % ; crédits d'inscription | Bulk : résumés de masse, labeling, evals ; SOC 2 / ISO 27001 / GDPR / HIPAA |
| 8 | **Novita AI** | Serverless haut débit à prix cassé | ~0,20 $/1M tokens sur les petits modèles ; crédits d'inscription | Startups, volumes variables, minimiser la facture |
| 9 | **SiliconFlow** | Inférence budget (Chine) avec endpoints **FP8** quantifiés | $/1M tokens très bas ; ex. DeepSeek V3.2 ~0,27/0,42 $ (FP8) | Le moins cher si la conformité CN est acceptable ; doc en anglais limitée (à vérifier) |
| 10 | **Parasail** | Serverless open-weight, déploiement rapide | $/1M tokens aligné marché (ex. DeepSeek V4 Pro 1,74/3,48 $ comme Fireworks/Together) | Alternative crédible à Fireworks/Together sur les gros modèles |
| 11 | **Hyperbolic** | Inférence décentralisée sur GPU distribués + marketplace | $/1M tokens ; prix agressifs | Pari décentralisé ; vérifier SLA et latence avant prod |
| 12 | **Baseten** | **Déploiements dédiés** single-tenant (Truss), autoscaling, scale-to-zero, Model APIs | GPU dédié à la minute (H100 ~6,50 $/h, B200 ~9,98 $/h) ; Model APIs dès ~0,10 $/1M | Vos modèles fine-tunés en prod avec monitoring ; éligible HIPAA |
| 13 | **Modal** | Serverless GPU **Python-first** (décorateurs), facturation à la seconde | H100 ~3,95 $/h ; A100-80 ~2,50 $/h ; **30 $/mois de crédits récurrents** | Batch, jobs ML, prototypes : le préféré des équipes Python |
| 14 | **Replicate** | Marketplace de modèles (image, vidéo, audio, LLM) + déploiements privés ; **Cog** (open source) pour packager | À la seconde GPU (H100 ~5,49 $/h) ou **au résultat** (ex. image à partir de 0,003 $) | Tester un modèle en 2 minutes sans infra ; créatifs (FLUX, Whisper, LTX) |
| 15 | **Hugging Face** | Hub (1M+ modèles) + **Inference Providers** (routage multi-hébergeurs) + Inference Endpoints dédiés | Free/PRO (9 $/mois) ; endpoints à l'heure d'instance | L'écosystème open-weight : tout y transite |
| 16 | **OpenAI** | Modèles fermés frontière (GPT-6 Astra, 5.6 Sol/Terra/Luna) + GPT-OSS ouverts | $/1M tokens (0,15→10,00 $ in) ; Batch -50 % ; cache -90 % lecture | Référence qualité ; écosystème le plus mature |
| 17 | **Anthropic** | Claude (Fable 5, Opus 5.5/4.8, Sonnet 5, Haiku 4.5) ; roi du **prompt caching** | $/1M tokens (1→10 $ in, 5→50 $ out) ; cache lecture -90 % | Code/agents longs ; RAG à gros préfixe réutilisé |
| 18 | **Google Cloud** | Gemini 3.1 (Pro/Flash) via Vertex AI + AI Studio ; TPU ; Gemma ouverts | $/1M tokens ; free tier sans CB ; 3.1 Pro 2,00/12,00 $ | Gros contextes 1M natifs ; prototypage gratuit |
| 19 | **Microsoft Azure** | GPT via Azure OpenAI + catalogue **AI Foundry** (GPT, Claude, Llama, Phi…) ; PTU | $/1M tokens ou débit provisionné (PTU) | Entreprises Microsoft ; SLA de latence via PTU |
| 20 | **AWS** | **Bedrock** (Claude, Llama, Mistral, Nova…), SageMaker, EC2 GPU, Trainium/Inferentia | $/1M tokens (Bedrock) ; $/h GPU ; Nova Micro 0,04/0,14 $ | Déjà sur AWS : Bedrock = un contrat, tous les modèles |
| 21 | **Oracle OCI** | Bare metal GPU (H100 ~10 $/h/GPU, 8×GPU RDMA) + Generative AI | $/h/GPU ; $/1M tokens (API) | Serveurs d'inférence dédiés (vLLM) à prix agressif |
| 22 | **xAI** | API Grok (4.x, Fast) ; 2M contexte sur Fast | $/1M tokens ; Grok 4.1 Fast 0,20/0,50 $ | Haut volume chez xAI ; temps réel (X) |
| 23 | **Mistral AI** | **La Plateforme** (API UE) + poids ouverts Apache 2.0/MIT | $/1M tokens ; Large 3 : 0,50/1,50 $ ; Nemo 0,02/0,04 $ | Souveraineté UE ; appels d'offres publics |
| 24 | **DeepSeek** | API directe V4-Pro/V4-Flash ; poids **MIT** | $/1M tokens ; V4-Pro 0,435/0,87 $ ; **tarifs peak/off-peak selon l'heure** | Le prix le plus bas du frontier-adjacent |
| 25 | **Alibaba Cloud** | Model Studio : Qwen3.x (ouverts Apache 2.0) + API | $/1M tokens ; Qwen3.7 Max 1,25/3,75 $ | Écosystème Qwen complet ; tarifs Asie agressifs |
| 26 | **ByteDance Volcengine** | API Doubao-Seed (Chine/Asie) | $/1M tokens agressifs (ex. DeepSeek V3.2 0,28/0,42 $ via Volcengine) | Marché asiatique ; vérifier la résidence des données |
| 27 | **Moonshot AI** | API Kimi (K2.x, K3) | $/1M tokens ; K3 ~3,00/15,00 $ (à vérifier) | Agents/code ; 1M contexte |
| 28 | **Zhipu AI** | API BigModel : GLM-4.7/5.x | $/1M tokens ; GLM-5.2 1,40/4,40 $ | Open-weight chinois de référence (SWE-bench) |
| 29 | **Cohere** | Command (RAG/entreprise) + **embeddings** multilingues réputés | $/1M tokens ; Command A 2,50/10,00 $ ; embed ~0,10 $/1M | RAG entreprise ; embeddings non-OpenAI |
| 30 | **AI21 Labs** | Jamba (hybride Mamba-Transformer, 256K) | $/1M tokens ; Jamba 1.7 Large 2,00/8,00 $ | Long contexte à coût contenu |
| 31 | **Perplexity** | API **Sonar** (LLM + recherche web intégrée) | $/1M tokens + frais/requête ; Sonar ~1,00/1,00 $ | Réponses avec citations web sans bricoler un search |
| 32 | **Voyage AI** | Embeddings/rerank SOTA (voyage-4-lite 0,02 $/1M) | $/1M tokens (embeddings) | RAG : alternative sérieuse à text-embedding-3 |
| 33 | **Mixedbread AI** | Embeddings (binaire quantifié) + rerank | $/1M tokens | RAG : précision/coût sur la recherche |
| 34 | **Jina AI** | Embeddings + **reranker** + Reader API | $/1M tokens ; free tier | RAG : rerank multilingue, extraction web |
| 35 | **Cloudflare** | **Workers AI** : inférence serverless sur le réseau edge mondial | 10 000 neurons/jour gratuits ; puis $/requête | Edge, sans serveur ; petits modèles open-weight |
| 36 | **NVIDIA** | **NIM** (conteneurs d'inférence optimisés) + DGX Cloud + API Build | Licence NVIDIA AI Enterprise en prod (dev gratuit) ; $/1M tokens (Build) | Standardiser le déploiement GPU (A100/H100+) |
