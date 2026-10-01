---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-4
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "embedding", "embeddings", "memory"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [245, 378]
sha256: c1cc927100afa532d64152f51d552087c47b38eb7e67045043a6283453ed30be
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

fast-plaid is LightOn's Rust implementation of PLAID, the index ColBERT was originally built around. There's no server to start, and it reads the tensors `encode_document` hands back without any conversion.

```
# pip install sentence-transformers datasets fast-plaid
from datasets import load_dataset
from fast_plaid import search
from sentence_transformers import MultiVectorEncoder
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
corpus = list(dict.fromkeys(dataset["answer"]))
model = MultiVectorEncoder("lightonai/LateOn")
query = "when did richmond last play in a preliminary final"
document_embeddings = model.encode_document(corpus, batch_size=32)
query_embedding = model.encode_query(query)
fast_plaid = search.FastPlaid(index="natural-questions", device="cuda")
# 4,874 documents (608,414 token vectors) indexed in 5s
fast_plaid.create(documents_embeddings=document_embeddings)
results = fast_plaid.search(queries_embeddings=query_embedding.unsqueeze(0), top_k=3)  # 11ms
for index, score in results[0]:
    print(f"{score:.4f}  {corpus[index][:90]}")
"""
11.8828  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieve
11.7676  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contes
11.6758  Battle of Appomattox Court House The Battle of Appomattox Court House (Virginia, U.S.), fo
"""
```
The `index` argument is a directory, not just a label, so the index is written to disk as it is built. Pointing a new `FastPlaid` at the same path reopens it for searching or for adding more documents, instead of rebuilding from the embeddings each time. On this corpus it occupies 92 MB, against 311.5 MB for the raw float32 vectors.

This is the only one of the four that is approximate, and it is the one place in this section where the scores do not match the exhaustive MaxSim. PLAID prunes with centroids and stores quantized residuals, so the three scores drift by a few hundredths in both directions against the 11.9192 / 11.7591 / 11.6710 computed earlier. The ranking is unaffected here, and that is the trade PLAID is making. It was designed for corpora far larger than this one, where scanning everything is not an option.

## **Qdrant**

Qdrant needs a server: `docker run -p 6333:6333 qdrant/qdrant`. The client also has a local mode (`QdrantClient(":memory:")`) that needs no server, but it's a pure-Python reimplementation, so use it for trying things out rather than for timing them.

```
# pip install sentence-transformers datasets qdrant-client
from datasets import load_dataset
from qdrant_client import QdrantClient, models
from sentence_transformers import MultiVectorEncoder
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
corpus = list(dict.fromkeys(dataset["answer"]))
model = MultiVectorEncoder("lightonai/LateOn")
query = "when did richmond last play in a preliminary final"
document_embeddings = model.encode_document(corpus, batch_size=32)
query_embedding = model.encode_query(query)
client = QdrantClient("http://localhost:6333")
client.create_collection(
    collection_name="natural-questions",
    vectors_config=models.VectorParams(
        size=model.get_embedding_dimension(),
        distance=models.Distance.COSINE,
        multivector_config=models.MultiVectorConfig(
            comparator=models.MultiVectorComparator.MAX_SIM
        ),
        # MaxSim never walks the HNSW graph, so skip building one
        hnsw_config=models.HnswConfigDiff(m=0),
    ),
)
# 4,874 documents (608,414 token vectors) ingested in 26.3s
client.upload_points(
    collection_name="natural-questions",
    points=[
        models.PointStruct(id=idx, vector=embedding, payload={"text": text})
        for idx, (embedding, text) in enumerate(zip(document_embeddings, corpus))
    ],
    batch_size=64,
)
results = client.query_points(
    collection_name="natural-questions",
    query=query_embedding,
    limit=3,
    with_payload=True,
).points  # 18ms
for result in results:
    print(f"{result.score:.4f}  {result.payload['text'][:90]}")
"""
11.9192  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieve
11.7591  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contes
11.6710  Battle of Appomattox Court House The Battle of Appomattox Court House (Virginia, U.S.), fo
"""
```
`MAX_SIM` is the only comparator Qdrant offers, and `hnsw_config=HnswConfigDiff(m=0)` is their recommendation for late-interaction fields, since the vectors are used for rescoring rather than graph traversal. Note that Qdrant themselves suggest reserving late interaction for reranking a few hundred candidates rather than scanning a whole collection, which is the Retrieve and Rerank pattern. At 4,874 documents the full scan costs 18ms and is exact, but that doesn't extrapolate.

## **Weaviate**

Weaviate needs a server too: `docker run -p 8080:8080 -p 50051:50051 cr.weaviate.io/semitechnologies/weaviate:1.34.0`. Multi-vector support needs 1.29 or newer, and the embedded mode isn't available on Windows.

```
# pip install sentence-transformers datasets weaviate-client
import weaviate
from datasets import load_dataset
from sentence_transformers import MultiVectorEncoder
from weaviate.classes.config import Configure, DataType, Property
from weaviate.classes.query import MetadataQuery
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
corpus = list(dict.fromkeys(dataset["answer"]))
model = MultiVectorEncoder("lightonai/LateOn")
query = "when did richmond last play in a preliminary final"
document_embeddings = model.encode_document(corpus, batch_size=32)
query_embedding = model.encode_query(query)
client = weaviate.connect_to_local()
collection = client.collections.create(
    "Documents",
    # self_provided turns on MaxSim late interaction
    vector_config=[Configure.MultiVectors.self_provided(name="colbert")],
    properties=[Property(name="text", data_type=DataType.TEXT)],
)
# 4,874 documents (608,414 token vectors) ingested in 41s
with collection.batch.fixed_size(batch_size=64) as batch:
    for text, embedding in zip(corpus, document_embeddings):
        batch.add_object(properties={"text": text}, vector={"colbert": embedding.tolist()})
results = collection.query.near_vector(
    near_vector=query_embedding.tolist(),
    target_vector="colbert",
    limit=3,
    return_metadata=MetadataQuery(distance=True),
)  # 17ms
for result in results.objects:
    # Weaviate reports the MaxSim score as a negated distance
    print(f"{-result.metadata.distance:.4f}  {result.properties['text'][:90]}")
"""
11.9192  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieve
11.7591  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contes
11.6710  Battle of Appomattox Court House The Battle of Appomattox Court House (Virginia, U.S.), fo
"""
client.close()
```
Defaults are enough here. Weaviate's dynamic `ef` resolves to 100 for a top-3 query, and this ranking is already exact from about 32 upward. That margin is a property of the embeddings rather than of Weaviate, so it's worth confirming on your own model instead of assuming the defaults hold.

Weaviate also supports MUVERA encoding, which made ingestion 3x faster and queries 1.8x faster in our test. It cost far more accuracy than that speed is worth at this size though, since the correct third passage didn't appear even in its top 50.

## **Vespa**

Vespa also runs in a container, but `pyvespa` starts it for you, so there's no separate `docker run`.

