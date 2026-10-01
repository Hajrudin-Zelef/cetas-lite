---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-5
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["embedding", "embeddings", "memory", "multimodal", "omni", "parameters", "reranker"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [379, 478]
sha256: 28f7d88b1e066b717dd6ae900258b12b83d973e64f93861c04fdd3c061ab3aa6
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

```
# pip install sentence-transformers datasets pyvespa
from datasets import load_dataset
from sentence_transformers import MultiVectorEncoder
from vespa.deployment import VespaDocker
from vespa.package import (
    ApplicationPackage, Document, Field, FirstPhaseRanking, Function, RankProfile, Schema,
)
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
corpus = list(dict.fromkeys(dataset["answer"]))
model = MultiVectorEncoder("lightonai/LateOn")
query = "when did richmond last play in a preliminary final"
document_embeddings = model.encode_document(corpus, batch_size=32)
query_embedding = model.encode_query(query)
# "dt" is a mapped dimension over the variable token count, "x" the dense 128-dim vector
package = ApplicationPackage(
    name="colbert",
    schema=[
        Schema(
            name="doc",
            document=Document(fields=[
                Field(name="text", type="string", indexing=["summary"]),
                Field(name="colbert", type="tensor<float>(dt{}, x[128])", indexing=["attribute"]),
            ]),
            rank_profiles=[
                RankProfile(
                    name="colbert",
                    inputs=[("query(qt)", "tensor<float>(qt{}, x[128])")],
                    functions=[Function(
                        name="max_sim",  # per query token take the best document token, then sum
                        expression="sum(reduce(sum(query(qt) * attribute(colbert), x), max, dt), qt)",
                    )],
                    first_phase=FirstPhaseRanking(expression="max_sim"),
                )
            ],
        )
    ],
)
app = VespaDocker(port=8080).deploy(application_package=package)  # ~40s to boot
# Vespa reads a mixed tensor as {token index: vector}, for documents and queries alike
def to_tensor(embedding):
    return {str(token): vector for token, vector in enumerate(embedding.tolist())}
# 4,874 documents (608,414 token vectors) ingested in ~80s
app.feed_iterable(
    ({"id": str(idx), "fields": {"text": text, "colbert": to_tensor(embedding)}}
     for idx, (text, embedding) in enumerate(zip(corpus, document_embeddings))),
    schema="doc",
)
response = app.query(body={
    "yql": "select text from doc where true",
    "ranking.profile": "colbert",
    "hits": 3,
    "input.query(qt)": to_tensor(query_embedding),
})  # ~75ms warm, ~115ms on the first call
for hit in response.hits:
    print(f"{hit['relevance']:.4f}  {hit['fields']['text'][:90]}")
"""
11.9192  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieve
11.7591  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contes
11.6710  Battle of Appomattox Court House The Battle of Appomattox Court House (Virginia, U.S.), fo
"""
```
Vespa asks for the most upfront structure of the four, because you're declaring a ranking pipeline rather than just an index. In exchange you get to write MaxSim out as a tensor expression and see exactly what it computes. This version puts MaxSim in `first-phase` over `where true`, which scores all 4,874 documents and is why the output matches exhaustive MaxSim exactly. It's deliberately not what Vespa recommends at scale. Their ColBERT sample app stores int8-binarized vectors and moves MaxSim into `second-phase` to rerank a cheaper first stage.

Moving to that phased setup needs care, because `second-phase` rescores only the best 100 candidates by default, and here that window left two of the three correct passages unscored entirely. Raising `rerank-count` to cover your candidate set fixes that, though at this size the phased version still came out slower than simply scanning everything.

Late interaction is the state of the art for visual document retrieval, matching a text query against page *images*, with charts, tables, and layout intact, and no OCR step. This is what the ColPali family of models does, and those checkpoints load and run through the same API, with the `revision` pinning the open pull request that adds this one's Sentence Transformers configuration (Supported Models has the full list). Image documents are passed as URLs, local paths, or PIL images:

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder("vidore/colqwen2.5-v0.2")
queries = [
    "What is the variable represented on the y-axis of the graph?",
    "Total outlay is maximum in which year?",
]
images = [
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/doc1.jpg",
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/doc2.jpg",
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/doc3.jpg",
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/doc4.jpg",
]
query_embeddings = model.encode_query(queries)
document_embeddings = model.encode_document(images)
print(query_embeddings[0].shape, document_embeddings[0].shape)
# (25, 128) (755, 128)
scores = model.similarity(query_embeddings, document_embeddings)
print(scores)
# tensor([[13.8672, 12.3115, 12.1670, 11.0293],
#         [ 7.2012, 14.7207,  6.9414,  6.9746]])
```
Each query retrieves its own page (the diagonal), and the second query separates much more cleanly than the first, since only one of the four pages is about outlay over time.

The code is unchanged. Underneath, the processor handles the visual prompt and the image patches, and MaxSim scores query text tokens against document image patches. A page holds many separate regions, which is exactly what makes late interaction a natural fit here, since a single vector would have to average a chart, a table, and three paragraphs into one summary. That fidelity costs index space, though. The shapes above are 755 token vectors for one page against 25 for the query, where a Natural Questions passage from earlier averaged about 125, so token pooling is worth reaching for earlier here than it is for text.

These are VLMs, so plan for the memory they need. The table in Supported Models runs from 252M to 8.8B parameters, and the small end of it stays practical on CPU where the multi-billion ones don't.

Page images are the common case, but they're not the only non-text modality. Sentence Transformers accepts text, images, audio, and video, and a checkpoint supports whichever of those its processor does, which `model.modalities` reports. A single document can combine modalities too, by passing a dict like `{"text": ..., "image": ...}` in place of a bare value. Multimodal Embedding & Reranker Models covers multimodal models in Sentence Transformers more broadly, and the Usage documentation lists exactly which input formats each modality accepts.

vidore/colqwen-omni-v0.1 is built on Qwen2.5-Omni and takes all four modalities. Retrieving a recorded conversation with it is the same two calls as retrieving a page:

