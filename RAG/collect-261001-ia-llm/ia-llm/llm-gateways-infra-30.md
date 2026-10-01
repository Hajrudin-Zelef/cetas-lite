---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-30
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Huawei", "Hugging Face", "Meta", "Mistral", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["embedding", "embeddings", "kv cache", "llama", "memory", "mistral", "open source", "qwen", "sglang", "vllm"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3469, 3593]
sha256: c258521d1ac3d4f696fda33bd2e440749b7136c5cd0eb4457f0b4590ec189ef4
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

### 167.2. Les algorithmes : BPE, WordPiece, Unigram

**BPE (Byte-Pair Encoding)** — le standard (GPT-2/3/4, RoBERTa, LLaMA, Qwen…) :
1. Vocabulaire de départ = caractères (ou octets) individuels.
2. Trouve la paire de symboles adjacents la plus fréquente dans le corpus d'entraînement.
3. Fusionne-la en un nouveau token, ajoute au vocabulaire.
4. Répète jusqu'à la taille de vocabulaire cible (ex. 100K, 200K).

Résultat : les séquences fréquentes (« maintenance », « https:// », « \n\n ») deviennent des tokens uniques ; les mots rares se décomposent en sous-mots connus. Jamais de mot « inconnu » : au pire, ça retombe sur les octets.

| Algorithme | Principe | Utilisé par |
|---|---|---|
| **BPE** | Fusions gloutonnes par fréquence | GPT, LLaMA, Qwen, Mistral, RoBERTa |
| **WordPiece** | Fusions par vraisemblance (pas fréquence brute) | BERT, DistilBERT |
| **Unigram / SentencePiece** | Part d'un gros vocabulaire, élague par probabilité ; traite le texte brut sans pré-découpe | T5, ALBERT, LLaMA (SentencePiece), Gemma |
| **tiktoken** | Implémentation BPE open source d'OpenAI (Rust) | GPT-3.5/4/4o, o-series, embeddings OpenAI |

Point crucial : **chaque modèle a SON tokenizer**. Compter avec `cl100k_base` pour dimensionner un appel à Qwen ou Llama donne une approximation (±10–20 %), pas le chiffre exact. Pour le chiffre exact, utilise le tokenizer du modèle (`AutoTokenizer.from_pretrained` — voir 167.4).

### 167.3. La librairie `tokenizers` (Hugging Face)

L'implémentation de référence : cœur **Rust**, bindings Python/Node — ~10 à 100x plus rapide que du Python pur, < 20 s pour tokeniser 1 Go de texte sur CPU. C'est ce que `transformers` utilise en interne quand tu charges un `AutoTokenizer`.

```bash
pip install tokenizers          # seul
pip install tokenizers transformers   # avec l'écosystème HF
```

```python
from tokenizers import Tokenizer
from tokenizers.models import BPE
from tokenizers.trainers import BpeTrainer
from tokenizers.pre_tokenizers import ByteLevel

# Entraîner un mini-tokenizer BPE sur TON corpus (utile pour mesurer,
# ou pour un tokenizer métier — voir l'encadré RAG plus bas)
tok = Tokenizer(BPE(unk_token="[UNK]"))
tok.pre_tokenizer = ByteLevel()
trainer = BpeTrainer(vocab_size=30000,
                     special_tokens=["[UNK]", "[CLS]", "[SEP]", "[PAD]", "[MASK]"])
tok.train(["corpus_maintenance.txt"], trainer)
tok.save("tokenizer-metier.json")

out = tok.encode("Remplacer le kit de fusion FK-475")
print(out.tokens)   # ex. ['Remplacer', ' le', ' kit', ' de', ' fusion', ' FK', '-', '475']
print(out.ids)      # les IDs numériques envoyés au modèle
print(len(out.ids)) # le compteur qui fait ta facture
```

Pipeline interne d'un encode : `texte brut → normalisation → pré-tokenisation → modèle (BPE/…) → post-processing (tokens spéciaux) → IDs`. Quand un découpage te surprend, c'est presque toujours la pré-tokenisation (espaces, ponctuation, accents) qui en est la cause.

### 167.4. Compter ses tokens AVANT d'appeler : les 3 patterns

**Pattern 1 — tiktoken (API OpenAI / compatible) :**

```python
import tiktoken

def count_openai(text: str, model: str = "gpt-4o") -> int:
    enc = tiktoken.encoding_for_model(model)  # choisit o200k_base pour gpt-4o
    return len(enc.encode(text))

# Garde-fou budgétaire : tronquer AVANT l'appel, pas après la facture
MAX_TOKENS = 8000
prompt = system_prompt + "\n\n" + "\n\n".join(chunks)
if (n := count_openai(prompt)) > MAX_TOKENS:
    raise ValueError(f"Prompt trop long : {n} tokens > {MAX_TOKENS}")
```

**Pattern 2 — le tokenizer du modèle local (le chiffre exact) :**

```python
from transformers import AutoTokenizer

tok = AutoTokenizer.from_pretrained("Qwen/Qwen3-8B")  # le VRAI tokenizer du modèle servi
n = len(tok.encode("Le rapport de maintenance doit être archivé avant vendredi."))
print(n, "/", tok.model_max_length)   # tokens / fenêtre de contexte du modèle
```

**Pattern 3 — estimation rapide sans lib (règle du pouce) :**

```text
Anglais : 1 token ≈ 4 caractères ≈ 0,75 mot
Français (cl100k_base, texte varié) : 1 token ≈ 4,5 caractères  → ×1,5 à ×1,8 vs anglais
Français technique RÉPÉTITIF : ~10 caractères/token (mesuré : 292 tokens pour ~3000 car.
de procédure copieur répétée 3× — la répétition se re-tokenize efficacement)
```

### 167.5. Impact direct n°1 : le coût

Toutes les API du guide facturent **au million de tokens**, entrée et sortie séparément. Exemple chiffré avec ton cas RAG (prix « vérifiés le 27/09/2026 », section 160) :

```text
Question RAG typique sur Doubao Seed 2.1 Pro (¥6 entrée / ¥30 sortie par M) :
- prompt système :            ~500 tokens
- 5 chunks FR de ~1000 car. : ~5 × 100 tokens = ~500 tokens   (mesuré §167.4)
- question utilisateur :      ~30 tokens
- total entrée :              ~1030 tokens → ¥0,00618  (~0,0009 $)
- réponse générée (300 mots FR) : ~600 tokens → ¥0,018  (~0,0027 $)
- COÛT TOTAL par question :   ~¥0,024  (moins d'un demi-centime d'euro)
```

À 1 000 questions/jour : ~¥24/jour (~3,5 $/jour). Le poste qui explose n'est jamais UNE question, c'est le volume × la taille des chunks × l'absence de cache. D'où :
- **compte tes chunks** : un chunk de 2000 caractères FR ≈ 200–450 tokens selon répétitivité — diviser la taille des chunks par 2 divise presque par 2 ton coût d'entrée ;
- **active le cache de prompt** quand l'API le propose (Doubao : ¥1,20/M en cache vs ¥6 — ton prompt système RAG est le candidat idéal) ;
- **l'embedding aussi se paie au token** : text-embedding-3-small facture ses tokens d'entrée (prix public OpenAI ~$0,02/M, à vérifier) — indexer 1 Go de docs FR ≈ ~100–220 M tokens ≈ $2–4,5 d'indexation. Négligeable une fois, à surveiller si tu réindexe souvent.

### 167.6. Impact direct n°2 : la fenêtre de contexte et la VRAM

La fenêtre de contexte se paie deux fois : en **prix** (tokens d'entrée) et en **mémoire** (KV cache).

```text
256K tokens, ça représente quoi en français ?
  256 000 tokens ÷ ~700 tokens/page A4  ≈  366 pages A4
128K tokens ≈ 183 pages A4. Ton corpus Huawei complet ne tient PAS dans un prompt —
  d'où le RAG (section 140) : on n'envoie que les bons chunks, pas tout.
```

KV cache (ordre de grandeur, à mesurer sur ton setup avec `torch.cuda.memory_summary()`) : pour un 8B en bfloat16, compte ~1–2 Mo de VRAM **par token de contexte**. 32K tokens de contexte ≈ 32–64 Go de KV cache — plus que les 16 Go des poids eux-mêmes. C'est pour ça que :
- les serveurs (vLLM/SGLang, section 163) gèrent le KV cache au cordeau (PagedAttention, RadixAttention) ;
- un prompt système de 10K tokens « parce qu'on a la place » est une mauvaise idée même quand la fenêtre le permet : tu paies le cache à chaque requête ;
- la règle d'or du RAG reste : **le plus petit contexte qui répond bien** (top-k serré, chunks denses — section 140).

### 167.7. Check-list token pour ton pipeline

