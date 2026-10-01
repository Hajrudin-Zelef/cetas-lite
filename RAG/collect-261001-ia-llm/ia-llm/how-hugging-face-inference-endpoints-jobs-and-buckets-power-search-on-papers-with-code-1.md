---
id: collect-261001-ia-llm/ia-llm/how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code-1
title: "how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia"]
dates: []
keywords: ["inference", "agents", "benchmark", "compute", "embedding", "embeddings", "gpu", "latency", "leaderboard", "memory", "nvidia", "qwen"]
source: docs/RAG/collect-261001-ia-llm/how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code.md
source_anchor: ""
source_lines: [1, 99]
sha256: 59f88ce95ba17a7d749c289bd11fdde7773410314260fa300800d539a3ac8f95
---

# how-hugging-face-inference-endpoints-jobs-and-buckets-power-search-on-papers-with-code

Of course, making AI research accessible requires a powerful search engine, so that humans and agents can quickly find relevant and related work, either through the website or the `pwc search` CLI command, which agents can use via the Skill.

It's important to note that searching for research is not quite the same as searching for regular text. A useful paper search engine should find an exact title or arXiv identifier, but it should also understand a query such as “small language models for code generation” even when those words do not appear together in a paper. It needs to recognize that “the original BERT paper” is a navigational request, tolerate an incomplete title or typos, and still respond quickly when a model service is cold or temporarily unavailable.

For Papers with Code, we built this as a **hybrid search** system. This is also based on our prior experience at ML6, where we developed RAG-based systems for clients. It turned out that hybrid search typically outperforms keyword- and vector-based search systems, as it combines the best of both worlds (see also this blog for more info). Keyword search finds exact mentions, whereas vector search finds more fuzzy, semantically similar terms. Note that rerankers (also called cross-encoders) can further improve the results, although they also come with additional overhead and latency.

Papers with Code relies on a PostgreSQL database, hence its full-text search capabilities provide a fast lexical baseline. For dense embeddings, pgvector is used to add semantic recall, and the reciprocal rank fusion (RRF) algorithm combines the two. Three Hugging Face services are used for the dense embeddings:

- Hugging Face Jobs gives us burstable GPU compute for embedding the paper corpus.
- Hugging Face Storage Buckets provides the durable handoff between our database, experiments, and Jobs.
- Hugging Face Inference Endpoints serves low-latency embeddings for live queries and incremental updates.

Today, the system maintains embeddings for more than 110,000 current papers sourced from arXiv and Daily Papers. This post explains the architecture, the design decisions behind it, and the lessons we learned while taking it to production.

We deliberately split search into an offline corpus build and an online search service:

The expensive, throughput-oriented work runs as Jobs. Durable artifacts live in a Bucket. Only the small query-embedding step sits on the request path, behind a protected Inference Endpoint, to power the online search. If that endpoint is cold, busy, or unhealthy, search immediately falls back to full-text retrieval. This separation makes the system both powerful and fast.

Embedding pipelines often fail in subtle ways: a model revision changes, query and document prompts are mixed up, vectors are truncated differently, or an updated abstract no longer matches its stored vector.

We avoid this by treating the embedding format as a versioned API. Every paper is encoded as:

```
normalized title + "\n\n" + normalized abstract
```
For each vector generation, we record:

- the model repository and exact revision;
- the output dimension;
- the input-format version;
- whether the input is a query or a document;
- the normalization method;
- a content hash for the source title and abstract.

Our production generation uses `Qwen/Qwen3-Embedding-0.6B`, pinned to an exact revision, with 256-dimensional L2-normalized vectors. We selected the model with help from the MTEB leaderboard, the go-to benchmark for comparing embedding models. Note that newer embedding models like Qwen3 allow for 2 new features:

- one can specify a **dynamic embedding size** , which allows to trade-off quality with speed/storage costs. Qwen models call this "MRL" which is short for Matryoshka Representation Learning. You can learn all about it here. We chose an embedding size of 256 to make the search fast.
- one can provide an **instruction prompt** . Qwen embedding models support a`document` prompt (which we use to embed the papers) and live searches use their`query` prompt (to embed the user query).

This contract follows an embedding from export, through GPU inference, into PostgreSQL, and finally into online retrieval.

Full-corpus embedding is a classic batch workload. It needs a GPU for a relatively short period, benefits from high throughput, and should not consume resources between runs. Hugging Face Jobs fits that shape well: a Job is defined by a command, a hardware flavor, and optionally a Docker image, and can run uv scripts with their dependencies declared inline.

Our corpus build starts by exporting the latest version of every paper from a repeatable-read PostgreSQL snapshot. The exporter streams rows rather than loading the catalog into memory, writes bounded JSONL shards, and creates a manifest containing row counts and SHA-256 checksums.

We sync that immutable run directory to a private Storage Bucket and mount the Bucket directly (using hf-mount) into an `l4x1` Job (an NVIDIA L4 GPU, which has 24GB of VRAM). From the worker's perspective it is simply a filesystem:

```
hf jobs uv run \
  --flavor l4x1 \
  --timeout 6h \
  --volume hf://buckets/OWNER/pwc-paper-embeddings:/bucket \
  embed_papers_job.py \
  --input /bucket/runs/RUN_ID/input \
  --output /bucket/runs/RUN_ID/output \
  --model Qwen/Qwen3-Embedding-0.6B \
  --revision MODEL_REVISION \
  --dimensions 256 \
  --allow-matryoshka
```
The worker:

1. verifies the input manifest and every shard checksum;
2. loads the pinned model revision;
3. sorts texts by length to reduce padding;
4. calls `encode_document` in batches (as noted in the model card);
5. reduces the batch size automatically if the GPU runs out of memory;
6. truncates the Matryoshka representation to 256 dimensions and normalizes it;
7. writes float16 Parquet shards atomically; and
8. records throughput, package versions, hardware, peak VRAM, row counts, and output checksums.

Each completed shard has its own marker, so a restarted Job can skip verified work. This is useful for a large corpus: retrying should just resume work rather than overwriting existing embeddings.

In our 5,000-paper pilot, the Qwen Job encoded about 75 papers per second at 1024 dimensions on an L4 GPU. The same pass could be deterministically materialized at 512 and 256 dimensions, so we could compare the storage and retrieval trade-offs without paying for more inference.

Storage Buckets are mutable, S3-like object storage on the Hub, optimized for AI workloads. They can be accessed through `hf://buckets/...` paths and mounted read-write in Jobs without building a separate storage integration.

For us, the Bucket is more than a place to put vectors. It is the boundary between three systems with different lifecycles:

- the production database exports source records;
- ephemeral Jobs consume those records and produce vectors;
- the importer validates the results before touching the search index.

We organize artifacts under immutable run prefixes:

```
runs/<run-id>/
├── input/
│   ├── manifest.json
│   └── papers-*.jsonl
└── output/
    ├── manifest.json
    ├── embeddings-*.parquet
    └── embeddings-*.complete.json
```
Buckets themselves are intentionally mutable, so immutability is an application-level rule: a run ID is never overwritten, and every artifact is covered by a manifest and checksum.

This gives us several useful properties:

