---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-31
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Mistral", "Nvidia", "OpenAI", "SGLang", "Unsloth"]
dates: ["2026-09-27"]
keywords: ["benchmark", "deepseek", "embedding", "embeddings", "gemini", "gemini 4", "gpt-6", "kv cache", "llama", "mistral", "nvidia", "qwen"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3594, 3650]
sha256: d1bb285923300af7d9f690cb41cb9b7b32314c0f1c39b7f27c4478121122a38e
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

1. **Un compteur par brique** : `tiktoken` pour les API OpenAI-compatibles, `AutoTokenizer` du modèle pour le local. Jamais d'estimation « au mot » dans du code de prod.
2. **Garde-fou avant chaque appel** : `if tokens > MAX: tronquer/alerter` — une boucle qui accumule des chunks sans compter, c'est une facture surprise ou un 400 Bad Request.
3. **Mesure ton corpus une fois** : script qui tokenize tout ton corpus avec le tokenizer de ton modèle d'embedding → tu connais ton coût d'indexation AVANT de payer, et tu repères les documents pathologiques (un PDF de 500 pages scannées en OCR bruité peut exploser le compteur).
4. **Français = budget ×1,5** : quand tu lis un benchmark de coût en anglais (« $X pour 1M tokens »), multiplie par ~1,5–1,8 pour ton usage FR réel.
5. **Tokens spéciaux** : les templates de chat (`<|im_start|>`, etc.) ajoutent des tokens invisibles — compte le prompt **après** application du template, pas avant.
6. **Ne pas optimiser prématurément** : entraîner ton propre tokenizer (167.3) ne sert que si tu entraînes aussi ton modèle. Pour un RAG sur modèles existants, le tokenizer est imposé — ton levier, c'est la taille et la densité des chunks.

### 167.8. Quel tokenizer pour quel modèle : le tableau de correspondance

| Modèle / API | Tokenizer à utiliser pour compter | Vocabulaire | Remarque |
|---|---|---|---|
| GPT-4o, o-series, embeddings OpenAI | `tiktoken` → `o200k_base` | 200 019 | `tiktoken.encoding_for_model("gpt-4o")` |
| GPT-3.5 / GPT-4 | `tiktoken` → `cl100k_base` | 100 277 | L'ancienne génération |
| Qwen3, Qwen2.5 | `AutoTokenizer.from_pretrained("Qwen/Qwen3-8B")` | ~152K | BPE byte-level, bon sur le FR |
| Llama 3.x | `AutoTokenizer.from_pretrained("meta-llama/Llama-3.1-8B")` | 128K | Logique proche de tiktoken |
| Mistral / Mixtral | Tokenizer Mistral (`mistral-common`) | 32K (v1/v2) | Vocabulaire petit → FR plus fragmenté |
| Gemma 3 | `AutoTokenizer.from_pretrained("google/gemma-3-4b-it")` | 256K | Gros vocabulaire, SentencePiece |
| DeepSeek V3/V4 | Tokenizer DeepSeek (`deepseek-ai/DeepSeek-V3`) | ~128K | Proche de l'esprit Llama |
| Doubao Seed | Tokenizer Volcano/Doubao (via doc Ark) | à vérifier | À vérifier sur la doc — ne pas supposer cl100k |

Erreur classique : compter avec `cl100k_base` puis appeler Qwen — l'écart de ±10–20 % suffit à faire dépasser une fenêtre ou à fausser un devis sur gros volumes. **Le compteur doit matcher le modèle servi.** Dans ton routeur, associe chaque modèle à son compteur (un dict `model_id → fonction de comptage`).

### 167.9. Exercice chiffré : dimensionner le budget tokens de ton RAG

Prenons ton cas réel — RAG FR technique, 10 000 documents d'~5 000 caractères chacun (PDF de manuels), chunks de 1 000 caractères avec 200 de recouvrement :

```text
1. Indexation (une fois) :
   10 000 docs × 5 000 car. = 50 M caractères
   FR technique : ~100 tokens / 1 000 car. (mesuré §167.4, texte répétitif)
   → ~5 M tokens d'embedding
   text-embedding-3-small ~$0,02/M (prix public, à vérifier) → ~$0,10 d'indexation
   → L'indexation est quasi gratuite. La RÉINDEXATION fréquente, non.

2. Requête type (Doubao Seed 2.1 Pro, §167.5) :
   entrée ~1 030 tokens → ¥0,0062 | sortie ~600 tokens → ¥0,018
   → ~¥0,024 / question (~0,0035 €)

3. À 500 questions/jour ouvrés (équipe support) :
   500 × 22 jours × ¥0,024 ≈ ¥264/mois ≈ 38 €/mois
   → Le LLM de génération : 38 €/mois. Le VRAI coût : ton temps de curation.

4. Piège : chunks ×2 (2 000 car.) sans rerank
   → entrée ~2 000 tokens/question → ~76 €/mois pour une qualité pas forcément meilleure.
   → D'où la section 140 : des chunks DENSES et un top-k serré valent mieux que du volume.
```

Moralité chiffrée : sur un RAG interne, le coût tokens est un bruit de fond (dizaines d'euros/mois) — **l'optimisation rentable n'est pas le prix au token, c'est la précision** (moins de questions reformulées, moins de chunks inutiles envoyés). Le jour où tu exposes le RAG à 500 utilisateurs, refais ce calcul : c'est là que le cache de prompt (Doubao ¥1,20/M) et le choix du modèle (Turbo vs Pro) changent la donne.

### 167.10. Tokens et sécurité : deux angles morts

1. **Saturation par document piégé** : un attaquant qui contrôle un document ingéré peut y glisser des milliers de tokens invisibles (texte blanc sur blanc, div cachées) pour saturer ta fenêtre ou gonfler ta facture. Contre-mesures : plafond de tokens PAR SOURCE à l'ingestion, détection des écarts (un chunk « normal » fait 100–400 tokens ; un chunk à 8 000 tokens mérite un examen).
2. **Usurpation de rôle via les délimiteurs** : les templates de chat ajoutent des marqueurs (`<|im_start|>`…) — si ton RAG concatène naïvement des contenus non fiables SANS passer par le template du modèle, un document peut forger un faux marqueur système. Toujours appliquer le chat template du tokenizer (`tok.apply_chat_template`) sur des messages STRUCTURÉS, jamais sur du texte brut concaténé.

> **Rappel de doctrine :** le token est l'unité de compte de toute ta stack IA — facturation API, fenêtre de contexte, VRAM du KV cache, coût d'indexation. Celui qui ne compte pas ses tokens pilote à l'aveugle. Un compteur de 10 lignes (167.4) dans ton routeur vaut tous les tableaux de prix du monde.

*Fin du guide — 167 sections. Extension du 27/09/2026 : sections 160 (Doubao), 161 (LM Studio), 162 (Unsloth), 163 (SGLang), 164 (PyTorch), 165 (NVIDIA NIM), 166 (NotebookLM), 167 (tokenisation). Prix « vérifiés le 27/09/2026 » par recherche web ; chiffres de tokenisation mesurés localement avec tiktoken (cl100k_base) le 27/09/2026 ; le reste « à vérifier » sur les pages officielles avant toute décision budgétaire. Rappel : « gemini 4 », « GPT-6 Terra » et assimilés n'existent pas au 27/09/2026 — ne pas les coder en dur (cf. section 159).*
