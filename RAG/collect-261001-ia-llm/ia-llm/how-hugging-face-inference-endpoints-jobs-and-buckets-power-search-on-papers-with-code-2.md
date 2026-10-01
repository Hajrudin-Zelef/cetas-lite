---
id: collect-261001-ia-llm/ia-llm/how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code-2
title: "how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "SGLang", "vLLM"]
dates: []
keywords: ["inference", "agent", "compute", "cost", "embedding", "embeddings", "gpu", "latency", "memory", "qwen", "reranker", "sglang"]
source: docs/RAG/collect-261001-ia-llm/how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code.md
source_anchor: ""
source_lines: [100, 185]
sha256: 2a0dd7ed1a162be1a5c360a5b1cd28842dc2d6436aa530d97a054a08fc9a14e6
---

# how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code

- **Reproducibility:** we can trace a database generation back to an exact corpus snapshot, model revision, and set of artifacts.
- **Safe retries:** Jobs can resume from completed shards in the same run prefix.
- **Cheap experiments:** several models or dimensions can reuse one verified input snapshot.
- **Controlled rollout:** importing a generation does not activate it. We first validate coverage and build its index.
- **Simple rollback:** the previous generation and its artifacts remain available until the new one is proven stable.

Only after the importer rechecks schemas, checksums, dimensions, normalization, unique paper IDs, and current content hashes do we load the vectors into PostgreSQL. We then build a separate HNSW index for the new generation and atomically mark it active only when every eligible current paper is covered (HNSW is the graph-based algorithm that enables fast vector search).

Batch embeddings solve the document side of retrieval. A user query still needs to be embedded at request time using the same model contract.

We deploy the pinned model as an authenticated Inference Endpoint backed by Text Embeddings Inference (TEI). The endpoint accepts the query text and returns a normalized 256-dimensional vector using the model's `query` prompt. Note that one could also leverage vLLM or SGLang here.

The API then performs a cosine-distance search over the active pgvector generation:

```
SELECT paper_id,
       embedding <=> CAST(:query_vector AS halfvec(256)) AS distance
FROM paper_embeddings
WHERE generation_id = :active_generation
ORDER BY embedding <=> CAST(:query_vector AS halfvec(256))
LIMIT 50;
```
The HNSW index keeps this lookup fast. On our 5,000-paper pilot, the 256-dimensional Qwen index achieved 0.9955 Recall@20 against exact search, with 1.31 ms p50 and 2.21 ms p95 HNSW lookup latency. Its table and index used about 27% of the storage of the 1024-dimensional version while retaining essentially the same ANN recall in that test.

The Endpoint is configured with a maximum of one replica and can scale to zero when idle. That is a useful cost lever, as this means you're not paying when there's no usage. However, this also means cold starts must be part of the application design rather than treated as an exceptional event, as it takes some time for the endpoint to spin up and serve traffic.

Our query client therefore has deliberately strict behavior:

- a one-second production timeout;
- a non-blocking concurrency limit;
- response dimension, finiteness, and norm validation;
- a short cache keyed by the query and embedding generation;
- a circuit breaker after repeated failures; and
- no raw query text in logs, only a normalized fingerprint.

If the endpoint is scaling up, times out, returns a malformed vector, or has no concurrency available, we skip the semantic branch immediately. Users still receive lexical results instead of waiting for an unreliable dependency.

Inference Endpoints works really reliably, and includes a nice dashboard so you can quickly see key analytics.

For every query, the lexical branch retrieves up to 50 candidates using weighted PostgreSQL full-text search. The semantic branch retrieves up to 50 candidates from pgvector.

We combine their ranks using weighted reciprocal rank fusion (RRF):


RRF is simple and robust, because it combines ranks rather than scores from two systems with different scales. Basically, if a paper is ranked high both by the lexical branch and the semantic branch, it has a higher chance of being ranked high by the hybrid search. We currently use equal branch weights and (k=60) (k is the "rank constant", a hyperparameter of the RRF algorithm).

Dense retrieval improves recall for conceptual queries. Full-text retrieval remains excellent for exact terminology, identifiers, and rare names. We also preserve deterministic identity behavior on top of the fused ranking:

- exact titles and arXiv IDs stay at the top;
- the method taxonomy recognizes navigational searches such as “the original BERT paper”;
- incomplete titles and bounded spelling mistakes use conservative trigram candidates; and
- ambiguous fuzzy matches abstain rather than forcing a bad result.

Note: hybrid search isn't always the best option, it is recommended to start with keyword search as a cheap and fast baseline, and only adding semantic and/or hybrid search when it turns out those give a reasonable boost in retrieval quality. One could further improve the search by adding a reranker after keyword/semantic/hybrid retrieval, using a model like Qwen3-Reranker.

The large initial corpus is embedded with Jobs, but Papers with Code changes continuously. New papers arrive, abstracts are corrected, and new arXiv versions become current.

Launching a GPU Job for a handful of changed rows would add unnecessary startup and orchestration overhead. Instead, an hourly incremental process selects missing or content-changed papers and sends a bounded delta to the same TEI Endpoint, this time with the `document` prompt.

Each run processes at most 500 papers in batches of 16. Before an embedding is written, the source row is locked and its content hash is checked again. If a paper changed during inference, that vector is discarded and picked up by the next run.

This gives us a useful division of labor:

- **Jobs** handle full rebuilds, new model generations, and large backfills.
- **Inference Endpoints** handle interactive query embeddings and small incremental document updates.
- **Buckets** preserve the large-build artifacts and make those builds resumable and auditable.

The hourly path keeps the active index close to the live catalog without turning an online endpoint into an unbounded batch processor.

The same document embeddings also power related-paper recommendations on each paper page.

Because the source paper already has a stored vector, related-paper retrieval requires no model call at request time. It is a single nearest-neighbor query over the active generation. If a vector is temporarily missing, the application can use a previous arXiv version or fill results from the existing task- and citation-based fallback. We fetch citation data through the Semantic Scholar API and also built `s2-cli`, a command-line interface for querying its citation graph. The latter is used by the agent at https://paperswithcode.co/chat.

Corpus embedding and query embedding use the same model, but they are different infrastructure problems. Jobs optimize for throughput and bounded cost; Inference Endpoints optimize for availability and request latency.

Buckets provide an explicit handoff between compute and production. Checksummed artifacts create a reviewable boundary before data enters the production index.

The revision, dimension, prompt, normalization, and input formatter all affect retrieval. Store them together and validate them everywhere.

Scale-to-zero is valuable when traffic is intermittent, but only if the product has a fast fallback. Hybrid search gave us that fallback naturally: lexical search is always useful on its own.

Matryoshka embeddings let us evaluate quality, memory, index size, and latency as one trade-off. In our pilot, 256 dimensions preserved ANN recall while materially shrinking storage compared with 1024 dimensions.

New generations are imported beside the current one, indexed independently, checked for complete and current coverage, and then activated atomically. Rollback is a configuration change, not an emergency recomputation.

Feel free to try out the search at https://paperswithcode.co or the chat interface at https://paperswithcode.co/chat, and let us know any feedback!
