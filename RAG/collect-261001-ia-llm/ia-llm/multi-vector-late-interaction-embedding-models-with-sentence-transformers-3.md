---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-3
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["embeddings", "gpu", "memory", "perplexity", "reranker"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [155, 244]
sha256: ccd46cad888c5b7b9e7beef78e7f54257fd4cefa5c40dc0a8671679222b2045a
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

```
scores = model.similarity_pairwise(query_embeddings, document_embeddings[:1])
print(scores)
# tensor([10.7942])
```
MaxSim sums over query tokens, so its magnitude scales with how many query tokens there are, which means you can't compare scores across models with different query recipes. LateOn encodes the Red Planet query above as 12 tokens. Run that same query and those same documents through ColBERTv2, which pads and truncates every query to exactly 32 tokens, and the scores land in a completely different range:

```
model = MultiVectorEncoder("colbert-ir/colbertv2.0")
# ... same encode_query / encode_document / similarity calls ...
print(scores)
# tensor([[12.7970, 27.1945, 23.8495, 24.5656]])
```
Within one model the ordering is all you need, but if you want scores on a bounded scale, switch the model's similarity function to MeanMaxSim, which divides by the query token count. Back on LateOn:

```
model = MultiVectorEncoder("lightonai/LateOn", similarity_fn_name="meanmaxsim")
# or on an already-loaded model: model.similarity_fn_name = "meanmaxsim"
print(model.similarity(query_embeddings, document_embeddings))
# tensor([[0.8995, 0.9259, 0.9145, 0.9234]])
```
Now every score is an average cosine similarity in `[-1, 1]`, although you'll only see `[0, 1]` in practice.

If your corpus is small, exhaustive MaxSim over all of it is the simplest thing that works. Encode the corpus once, then score each query against everything:

```
import time
from datasets import load_dataset
from sentence_transformers import MultiVectorEncoder
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
# Several questions share an answer passage, so drop repeats but keep the order
corpus = list(dict.fromkeys(dataset["answer"]))  # 5,000 rows -> 4,874 passages
model = MultiVectorEncoder("lightonai/LateOn")
corpus_embeddings = model.encode_document(corpus, show_progress_bar=True)
query = "when did richmond last play in a preliminary final"
start = time.perf_counter()
query_embeddings = model.encode_query([query])
scores = model.similarity(query_embeddings, corpus_embeddings)[0]  # 98ms
top_scores, top_indices = scores.topk(3)
print(f"Search took {(time.perf_counter() - start) * 1000:.1f}ms")
for score, index in zip(top_scores.tolist(), top_indices.tolist()):
    print(f"{score:.4f}  {corpus[index][:100]}")
"""
Search took 122.7ms
11.9192  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieved
11.7591  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contest
11.6710  Battle of Appomattox Court House The Battle of Appomattox Court House (Virginia, U.S.), fou
"""
```
Those 4,874 passages encoded in 20 seconds on an RTX 3090, and each search takes about 120ms end to end, most of that the MaxSim scoring against all 608,414 token vectors. This is exact, but it scales linearly in total corpus tokens and keeps every token vector in memory, so reach for it when you have a few thousand documents rather than a few million. The runnable version of this script is semantic_search.py.

Past that size you want a real late-interaction index, which Sentence Transformers doesn't ship. It doesn't need to, because these indexes store whatever `encode_document` produced, so you encode here and hand the token embeddings to something built for them. Indexing has working snippets for four of the options, and the section directly below covers how to skip the index entirely.

You can also get late-interaction quality without maintaining a late-interaction index, by using a multi-vector model as your *reranker*. A fast bi-encoder narrows a large corpus to a handful of candidates, then the multi-vector model rescores only those:

```
from datasets import load_dataset
from sentence_transformers import MultiVectorEncoder, SentenceTransformer
from sentence_transformers.util import semantic_search
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:50000]")
corpus = list(dict.fromkeys(dataset["answer"]))
retriever = SentenceTransformer("jinaai/jina-embeddings-v5-text-nano-retrieval")
reranker = MultiVectorEncoder("perplexity-ai/pplx-embed-v1-late-0.6b", trust_remote_code=True)
# First stage: index the corpus once with a fast bi-encoder
corpus_embeddings = retriever.encode_document(corpus, convert_to_tensor=True, show_progress_bar=True)
# Retrieve the top 50
query = "when did richmond last play in a preliminary final"
hits = semantic_search(retriever.encode_query([query], convert_to_tensor=True), corpus_embeddings, top_k=50)[0]
candidates = [corpus[hit["corpus_id"]] for hit in hits]
# Second stage: rescore just those candidates with MaxSim
query_embeddings = reranker.encode_query([query])
document_embeddings = reranker.encode_document(candidates)
scores = reranker.similarity(query_embeddings, document_embeddings)[0]
for index in scores.argsort(descending=True)[:3].tolist():
    print(f"{scores[index].item():.4f}  {candidates[index][:100]}")
```
Only the 50 candidates are ever encoded as multi-vectors, so your index stays a normal dense index and the token vectors are transient. This is the same role a cross-encoder plays in a retrieve-and-rerank stack, but a multi-vector model is considerably cheaper per candidate. You encode the documents in one batch and score them with a matrix multiplication, instead of one forward pass per query-document pair. The runnable script is retrieve_rerank.py, which prints the timings of both stages.

Several vector databases index and score multi-vectors natively: Qdrant since v1.10, Weaviate since v1.29, Vespa for years now, LanceDB since v0.15.0, and VectorChord, which adds a MaxSim operator to Postgres that plain pgvector doesn't have. Milvus joined them in v2.6.4, under array-of-structs rather than the unrelated feature it calls multi-vector search. If you would rather not run a server at all, LightOn's fast-plaid is a `pip install` away and implements PLAID directly, and PyLate wraps it in a fuller retrieval stack.

A few others get you partway. OpenSearch and Elasticsearch can rescore candidates with MaxSim but not retrieve on it, and the Elasticsearch field is additionally in technical preview and Enterprise-tier. turbopuffer has late-interaction indexing in private beta.

The snippets below index text, but nothing in them is text-specific. `encode_document` hands back the same list of token-vector matrices whether the document was a passage, a page image, an audio clip, or a video, so the ColPali-style models from Visual Document Retrieval go into any of these unchanged. There are simply more vectors per document, which is what makes Token Pooling worth reaching for sooner there.

fast-plaid, Qdrant, Weaviate, and Vespa all take exactly what `encode_document` returns, so the code is the same up to the client library. Here's a working snippet for each, run against the 4,874 passages and 608,414 token vectors from the Semantic Search example. Each one carries the ingestion and query times it produced on one machine (RTX 3090, i7-13700K), with no tuning beyond what the code shows, to give a sense of the shape of the work. All four answer the query faster than the 98ms `model.similarity` took in that section, and three of them do it on the CPU, since fast-plaid is the only one here using the GPU.

All four returned the same three passages in the same order as the exhaustive PyTorch MaxSim earlier in this post, and the three databases reproduce its scores to four decimals! That is because their snippets score every document, which is affordable at this size and removes approximation as a variable. fast-plaid is approximate by design, so its scores differ slightly. The notes under each one say what changes when you switch to an approximate index, which is where rankings start to drift.

## **fast-plaid**

