---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-2
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "embedding", "embeddings", "training"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [55, 154]
sha256: 5470846c57b28e9429bb047183164cd617f7ca58932af41b10adbb04ef45f725
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder("lightonai/LateOn")
```
To find models that work, look for the `multi-vector` and `sentence-transformers` tags on the Hub. Any model with those tags loads with the line above, whether it started life as a PyLate checkpoint, a Stanford-NLP ColBERT checkpoint, or a ColPali-family model for visual document retrieval. We're working through the ecosystem to get that tag onto every model that works, so the list keeps growing.

Underneath, `MultiVectorEncoder` reads each of the formats these checkpoints have been published in over the years, so PyLate and Stanford-NLP checkpoints load directly even where the tag hasn't been added yet:

```
from sentence_transformers import MultiVectorEncoder
# Native Sentence Transformers checkpoints. PyLate builds on the same schema,
# so any PyLate checkpoint loads identically
model = MultiVectorEncoder("lightonai/LateOn")
model = MultiVectorEncoder("mixedbread-ai/mxbai-edge-colbert-v0-17m")
model = MultiVectorEncoder("LiquidAI/LFM2.5-ColBERT-350M", trust_remote_code=True)
# Any Stanford-NLP ColBERT checkpoint, detected via the `HF_ColBERT` architecture
# marker. The inline projection weight and the recipe come from `artifact.metadata`
model = MultiVectorEncoder("colbert-ir/colbertv2.0")
model = MultiVectorEncoder("answerdotai/answerai-colbert-small-v1")
# A bare transformer: a fresh random projection is appended, so training is required
model = MultiVectorEncoder("answerdotai/ModernBERT-base")
```
Visual document retrieval models are the exception. ColPali-family checkpoints ship in colpali-engine's own format, which carries no information Sentence Transformers can use, so each one needs a small configuration added to its repository before it loads. Most of that work is done and waiting to be merged. See Supported Models for the current state and how to load them today.

Multi-vector models carry a handful of recipe knobs that differ per checkpoint: marker prefixes for queries and documents, length caps, whether queries are padded out with `[MASK]` tokens, and which tokens are skipped when scoring documents. All of them live in the module configs, so `print(model)` shows you exactly what you loaded. Here's the original ColBERTv2 checkpoint, which pads every query to exactly 32 tokens and truncates documents at 180:

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder("colbert-ir/colbertv2.0")
print(model)
"""
MultiVectorEncoder(
  (0): Transformer({..., 'document_length': 180,
                    'query_expansion': {'strategy': 'fixed', 'attend': False, 'token': None, 'length': 32}})
  (1): Dense({'in_features': 768, 'out_features': 128, 'bias': False, ...})
  (2): MultiVectorMask({'skiplist_words': ['!', '"', '#', ...], 'skiplist_tasks': ['document'], ...})
  (3): Normalize({...})
)
"""
print(model.prompts)
# {'query': '[unused0] ', 'document': '[unused1] '}
```
That's the classic ColBERT pipeline: a `Transformer` producing contextualized token embeddings, a token-level `Dense` projecting each of them to 128 dimensions, a `MultiVectorMask` deciding which tokens count during scoring, and a token-level `Normalize`. Other checkpoints fill in different values. `lightonai/GTE-ModernColBERT-v1` uses the same four modules with `[Q]`  and `[D]`  prompts, no query expansion, and caps of 48 and 300.

You rarely need to touch any of this, since every released checkpoint configures its own. It matters when you build a model from a bare backbone, which is covered in Creating Custom Models.

One value is worth checking against your own data, though. `document_length` truncates, so anything past it never reaches the index. For example, a 662-token passage through LateOn's cap of 300 comes back as 273 vectors, with the rest of the passage simply gone. Most of these checkpoints were trained on short passages, so if your chunks are longer than the cap, you can lift it for a single call with `encode_document(..., processing_kwargs={"text": {"max_length": 512}})`, keeping in mind that you would be running the model past the length it was trained on and that the index grows roughly in proportion. Multi-vector models tend to tolerate that well. On MLDR, a long-document retrieval benchmark, the multilingual siblings of the pair above show the gap clearly: mLateOn scores 77.92 against mDenseOn's 51.59.

Multi-vector models are asymmetric, so queries and documents go through different prefixes, different length caps, and different scoring masks. Unlike many dense models, where the two are interchangeable, `encode_query()` and `encode_document()` are required to get correct embeddings:

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder("lightonai/mLateOn")
queries = ["What is the capital of France?"]
documents = [
    "Paris is the capital of France.",
    "Berlin is the capital and largest city of Germany, by both area and population.",
]
query_embeddings = model.encode_query(queries)
document_embeddings = model.encode_document(documents)
print(query_embeddings[0].shape)
# (10, 128)
print(document_embeddings[0].shape, document_embeddings[1].shape)
# (10, 128) (19, 128)
```
What you get back is a *list* of 2D tensors, one per input, each of shape `(num_tokens, embedding_dim)`. Unlike dense embeddings, you can't stack these into one rectangular tensor, because every input has its own token count. The second document is longer than the first, so it comes back as a taller matrix.

Each call applies the model's own recipe for you. `encode_query` prepends the query marker, expands the query to a fixed length if the checkpoint asks for it, and caps it at the query length. `encode_document` prepends the document marker, caps at the document length, and drops any skiplisted tokens (punctuation, for most checkpoints) from the scoring mask.

The usual `encode()` arguments all still apply, so `batch_size`, `show_progress_bar`, `convert_to_numpy`, `device`, and multi-process pools work the way you'd expect:

```
document_embeddings = model.encode_document(
    documents,
    batch_size=64,
    show_progress_bar=True,
)
```
`model.similarity()` computes the full all-pairs MaxSim matrix:

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder("lightonai/LateOn")
query_embeddings = model.encode_query(["Which planet is known as the Red Planet?"])
document_embeddings = model.encode_document([
    "Venus is often called Earth's twin because of its similar size and proximity.",
    "Mars, known for its reddish appearance, is often referred to as the Red Planet.",
    "Jupiter, the largest planet in our solar system, has a prominent red spot.",
    "Saturn, famous for its rings, is sometimes mistaken for the Red Planet.",
])
scores = model.similarity(query_embeddings, document_embeddings)
print(scores)
# tensor([[10.7942, 11.1104, 10.9743, 11.0811]])
```
Mars wins, as it should, though the runners-up are close behind. Saturn also contains the literal phrase "the Red Planet", and Jupiter is a planet with a red spot, so a token-level operator has plenty to latch onto in all three. The ordering is what matters.

Scores often sit this close together, as GLInt shows by measuring the spread across a full candidate pool. MaxSim takes a *maximum* per query token, so a document will usually give every query token some decent best match, and scores start from a floor. Contextualized token embeddings are also anisotropic, clustering in a narrow cone rather than spreading out, so even arbitrary token pairs tend to score high.

There is also `model.similarity_pairwise()`, for when you already have matched pairs and just want the pair scores instead of the full similarity matrix:

