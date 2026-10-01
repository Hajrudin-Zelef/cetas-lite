---
id: collect-261001-ia-llm/ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-3
title: "Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "vLLM"]
dates: []
keywords: ["embedding", "embeddings", "accelerator", "benchmarks", "gguf", "inference", "llama", "llama.cpp", "memory", "pruning", "training", "vllm"]
source: docs/RAG/collect-261001-ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-best-sub-.md
source_anchor: ""
source_lines: [92, 237]
sha256: 4921157620c644c4bcc78ccdc18a0b9bd7a5e916d86e2250f28680e64bc24ea6
---

# Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("ibm-granite/granite-embedding-311m-multilingual-r2")
# Full 768-dimensional embeddings
full = model.encode(["example text"])
print(full.shape)  # (1, 768)
# Truncated to 384 dimensions
small = model.encode(["example text"], truncate_dim=384)
print(small.shape)  # (1, 384)
```
The 97M model does not support Matryoshka — 384 dimensions is already compact.

Average performance on cross-lingual tasks within MTEB Retrieval. Belebele measures cross-lingual passage matching across 122 languages; MLQA measures extractive cross-lingual question answering retrieval across 7 languages.

| Model | Belebele Retrieval | MLQA Retrieval | 
|---|---|---|
| granite-embedding-107m-multilingual (R1) | 55.1 | 60.5 | 
| granite-embedding-278m-multilingual (R1) | 62.2 | 63.0 | 
| granite-embedding-97m-multilingual-r2 | 52.9 | 60.5 | 
| **granite-embedding-311m-multilingual-r2** | **66.5** | **67.1** | 

The 311M R2 model gains +4.3 on Belebele and +4.1 on MLQA over its R1 predecessor, showing improved cross-lingual transfer at the larger scale across both benchmarks.

The 97M R2 model scores lower on Belebele (52.9 vs 55.1, −2.2) while matching its R1 predecessor on MLQA (60.5). The Belebele gap is a tradeoff inherent in the pruning and vocabulary reduction process — the R2 model's training prioritized the broader 18-language MTEB Multilingual Retrieval set (where it gains +12.2 over R1) and long-document retrieval (+31.3), while the smaller vocabulary (180K vs. 250K tokens) and reduced layer count (12 vs. 22) affect narrow cross-lingual transfer tasks. If cross-lingual transfer across many language pairs is your primary use case, the full-size 311M model is the better choice.

Both models ship with multiple deployment paths for production use. Install the core library with:

```
pip install sentence-transformers
```
**Sentence Transformers** (recommended for most users):

```
from sentence_transformers import SentenceTransformer, util
model = SentenceTransformer("ibm-granite/granite-embedding-97m-multilingual-r2")
queries = [
    "What is the tallest mountain in Japan?",          # English
    "Wer hat das Lied Achy Breaky Heart geschrieben?", # German
    "ドイツの首都はどこですか？",                            # Japanese
]
passages = [
    "富士山は、静岡県と山梨県にまたがる活火山で、標高3776.12 mで日本最高峰の独立峰である。",  # Japanese
    "Achy Breaky Heart is a country song written by Don Von Tress.",                        # English
    "Berlin ist die Hauptstadt und ein Land der Bundesrepublik Deutschland.",                # German
]
q_emb = model.encode(queries)
p_emb = model.encode(passages)
print(util.cos_sim(q_emb, p_emb))
# Each query scores highest against its matching passage — across languages
```
**LangChain** (`pip install langchain-huggingface`):

```
from langchain_huggingface import HuggingFaceEmbeddings
embeddings = HuggingFaceEmbeddings(
    model_name="ibm-granite/granite-embedding-97m-multilingual-r2"
)
docs = embeddings.embed_documents([
    "富士山は日本最高峰の独立峰です。",
    "Mount Fuji is Japan's highest peak.",
])
query = embeddings.embed_query("What is Japan's tallest mountain?")
# Drop-in replacement anywhere LangChain accepts an Embeddings object
```
**LlamaIndex** (`pip install llama-index-embeddings-huggingface`):

```
from llama_index.embeddings.huggingface import HuggingFaceEmbedding
from llama_index.core import Settings
embed_model = HuggingFaceEmbedding(
    model_name="ibm-granite/granite-embedding-97m-multilingual-r2"
)
Settings.embed_model = embed_model  # applies globally to any index or pipeline
```
## **Haystack** (`pip install sentence-transformers haystack-ai`)

```
from haystack.components.embedders import (
    SentenceTransformersDocumentEmbedder,
    SentenceTransformersTextEmbedder,
)
from haystack.components.retrievers.in_memory import InMemoryEmbeddingRetriever
from haystack.dataclasses import Document
from haystack.document_stores.in_memory import InMemoryDocumentStore
doc_embedder = SentenceTransformersDocumentEmbedder(
    model="ibm-granite/granite-embedding-97m-multilingual-r2"
)
query_embedder = SentenceTransformersTextEmbedder(
    model="ibm-granite/granite-embedding-97m-multilingual-r2"
)
doc_embedder.warm_up()
query_embedder.warm_up()
# Embed and index documents
document_store = InMemoryDocumentStore()
result_docs = doc_embedder.run(documents=[
    Document(content="富士山は日本最高峰の独立峰です。"),
    Document(content="Mount Fuji is Japan's highest peak."),
    Document(content="Achy Breaky Heart is a country song written by Don Von Tress."),
    Document(content="Berlin ist die Hauptstadt und ein Land der Bundesrepublik Deutschland."),
])
document_store.write_documents(result_docs["documents"])
# Embed query and retrieve
result_query = query_embedder.run(text="What is Japan's tallest mountain?")
retriever = InMemoryEmbeddingRetriever(document_store=document_store)
results = retriever.run(query_embedding=result_query["embedding"], top_k=2)
for doc in results["documents"]:
    print(f"{doc.score:.3f}  {doc.content}")
# 0.961  Mount Fuji is Japan's highest peak.
# 0.913  富士山は日本最高峰の独立峰です。
```
## **Milvus** (`pip install pymilvus sentence-transformers`)

```
from pymilvus import MilvusClient
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("ibm-granite/granite-embedding-97m-multilingual-r2")
# Use "./milvus.db" for local persistence or a server URI for production
client = MilvusClient(":memory:")
client.create_collection(collection_name="multilingual_docs", dimension=384)
docs = [
    "富士山は日本最高峰の独立峰です。",
    "Mount Fuji is Japan's highest peak.",
    "Achy Breaky Heart is a country song written by Don Von Tress.",
    "Berlin ist die Hauptstadt und ein Land der Bundesrepublik Deutschland.",
]
embeddings = model.encode(docs).tolist()
client.insert(
    collection_name="multilingual_docs",
    data=[{"id": i, "vector": emb, "text": doc} for i, (emb, doc) in enumerate(zip(embeddings, docs))],
)
query_emb = model.encode(["What is Japan's tallest mountain?"]).tolist()
results = client.search(
    collection_name="multilingual_docs",
    data=query_emb,
    limit=2,
    output_fields=["text"],
)
for hit in results[0]:
    print(f"{hit['distance']:.3f}  {hit['entity']['text']}")
# 0.961  Mount Fuji is Japan's highest peak.
# 0.913  富士山は日本最高峰の独立峰です。
```
Both models also ship with pre-converted **ONNX** and **OpenVINO** weights for optimized CPU/accelerator inference, work as embedding endpoints via **vLLM** (`vllm serve ... --task embed`), and can be converted to GGUF for **Ollama** using llama.cpp. See the model cards for full deployment examples.

If you maintain an embedding framework, vector store, or RAG pipeline library and are evaluating these models as a default, here's what you need to know:

