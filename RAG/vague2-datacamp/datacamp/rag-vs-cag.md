---
id: vague2-datacamp/datacamp/rag-vs-cag
title: "RAG par rapport à CAG : Principales différences, avantages et cas d'utilisation"
domain: datacamp
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: ["agents", "attention", "context window", "latency", "memory", "research", "training"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/rag-vs-cag.md
source_anchor: ""
source_lines: [1, 57]
sha256: 0958878d99d459746afd3e99aad42d1e0e9ac9d775f0d0f33c7796bea8678110
---

# RAG par rapport à CAG : Principales différences, avantages et cas d'utilisation

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/rag-vs-cag
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article compares **Retrieval-Augmented Generation (RAG)** and **Cache-Augmented Generation (CAG)** as approaches to injecting knowledge into LLMs. The author (Benito Martin) argues they are less competitors than different tools for different tasks, sometimes best combined.

**RAG** lets models go beyond fixed training data by connecting to external databases and search mechanisms, retrieving relevant documents at query time. Workflow: the user query is encoded into a vector, used to search a vector database; effective **chunking strategies** split documents into semantic units (typically 100–1,000 tokens); approximate nearest-neighbor search retrieves quickly at scale; retrieved documents are passed to the generation step. External sources can include proprietary databases, scientific papers, legal archives, or real-time APIs — like giving your AI a library card rather than hoping it memorized every book. Strengths: **real-time updates** (a legal policy updated at 3pm is known at 3:01pm), **reduced hallucinations** (answers grounded in retrieved documents), and **flexible data integration** (structured DBs, semi-structured APIs, unstructured text). Limits: **system complexity** (retriever + vector DB + generator = more failure points), **latency** (retrieval adds overhead per query), and **dependence on retrieval quality** (poor retrieval harms answers).

**CAG** instead preloads and retains relevant information in the model's extended context or cache. It relies on two caching mechanisms: **knowledge caching** (documents preloaded into the extended context window and reused across queries) and **key-value (KV) caching** (storing attention states — key and value matrices — generated when processing tokens, so similar/repeated queries reuse them instead of recomputing). This reduces latency and preserves longer context. Strengths: **speed and efficiency** (reuse of cached computations), **cross-session consistency** (good for conversational agents, workflow automation, customer support), and **reduced system complexity**. Limits: **stale information** (cached data ages), **high memory requirements**, and **complex cache management** (coordination across distributed deployments).

**Key differences table:** RAG is just-in-time (retrieves during the query), slower, real-time knowledge, best for dynamic/large datasets, scales horizontally, high complexity, justifies answers with citations. CAG is preloaded (loads before the query), fastest, snapshot knowledge, best for stable/repetitive datasets, memory-bound scalability, moderate complexity, justifies answers within a consistent pre-established context.

**Decision framework:** assess information volatility, latency requirements, consistency needs, and resource availability. High volatility favors RAG; stable domains benefit from CAG. Latency-sensitive apps favor CAG; apps needing the freshest info should use RAG.

**Use RAG for:** dynamic, frequently updated information (news, evolving products, research); large/varied knowledge bases (legal, medical, competitive intelligence); protection against outdated information. **Use CAG for:** stable knowledge requirements (customer service FAQs, established curricula, workflow automation); high query volumes with repetitive patterns; latency-critical applications (real-time recommendation, interactive gaming). CAG is ideal when you can predict 90% of queries and knowledge is stable.

**Industry cases:** healthcare (RAG for latest research/protocols, CAG for standard protocols/patient summaries — hybrid works best); finance (RAG for market analysis/compliance/investment research, CAG for standard calculations/product definitions — regulatory risk often drives RAG); education (RAG for varied/updated content, CAG for repetitive quizzes/standard explanations); software engineering (RAG for docs/API specs/troubleshooting, CAG for code autocomplete/debugging/recurring queries); legal (RAG for case law/contract review, CAG for fixed compliance rules like GDPR Article 15); retail/e-commerce (RAG for product search/inventory/recommendations, CAG for shipping/return policies/order status).

**Hybrid approaches:** combine both — CAG for stable, frequently accessed information and RAG for real-time/specialized queries, with intelligent routing. Benefits: optimized response times, maintained accuracy for dynamic content, reduced system load. Challenges: increased architectural complexity, coordination between cache and retrieval, sync maintenance, and routing logic. Common in customer service, search, and e-commerce.

## Key points

- RAG is just-in-time retrieval; CAG preloads knowledge and KV caches.
- RAG: real-time knowledge, higher latency, horizontal scaling, high complexity, citation grounding.
- CAG: fastest, snapshot knowledge, memory-bound, moderate complexity, consistent context.
- Chunking typically 100–1,000 tokens; approximate nearest-neighbor search for scale.
- CAG uses knowledge caching plus key-value (attention state) caching.
- Choose RAG for dynamic/updated info; CAG for stable/repetitive queries and low latency.
- Hybrid systems route between cached and retrieved knowledge for the best of both.
- No universal answer; choice depends on requirements, constraints, and goals.

## Technical data / figures

| Feature | RAG | CAG |
|---|---|---|
| Core mechanism | Just-in-time retrieval from external DB | Preloaded context/cache before query |
| Latency | Slower (retrieval overhead) | Fastest (memory access) |
| Knowledge freshness | Real-time | Snapshot at last cache update |
| Best use case | Dynamic, large datasets | Stable, repetitive datasets |
| Scalability | Horizontal | Memory-limited |
| Complexity | High (vector DB, pipelines, retrieval) | Moderate (cache lifecycle, context) |
| Hallucination handling | Grounds answers in retrieved documents | Grounds answers in consistent cached context |

Chunk size: 100–1,000 tokens. Key mechanisms: vector databases, approximate nearest neighbor, knowledge caching, key-value (KV) caching of attention states.

## Why this source matters for the RAG

It provides a clear, practical comparison of two knowledge-injection architectures with decision criteria, industry examples, and hybrid guidance. It is valuable for RAG queries about retrieval strategies, caching, latency vs freshness trade-offs, and system design.
