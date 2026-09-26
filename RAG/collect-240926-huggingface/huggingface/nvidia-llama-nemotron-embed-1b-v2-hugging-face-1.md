---
id: collect-240926-huggingface/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face-1
title: "Compute similarity scores"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "vLLM"]
dates: []
keywords: ["compute", "attention", "embedding", "embeddings", "inference", "latency", "license", "llama", "nvidia", "protein", "research", "throughput"]
source: docs/RAG/clean_en/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face.md
source_anchor: ""
source_lines: [1, 117]
sha256: 17b75ec7c0178ef348c2c2ae2b2777cc78e5a2ccdff4de60d0cc87b303ba0995
---

# Compute similarity scores

<!-- source: https://huggingface.co/nvidia/llama-nemotron-embed-1b-v2 -->

The Llama Nemotron Embedding 1B model is optimized for **multilingual and cross-lingual** text question-answering retrieval with **support for long documents (up to 8192 tokens) and dynamic embedding size (Matryoshka Embeddings)**. This model was evaluated on 26 languages: English, Arabic, Bengali, Chinese, Czech, Danish, Dutch, Finnish, French, German, Hebrew, Hindi, Hungarian, Indonesian, Italian, Japanese, Korean, Norwegian, Persian, Polish, Portuguese, Russian, Spanish, Swedish, Thai, and Turkish.

In addition to enabling multilingual and cross-lingual question-answering retrieval, this model reduces the data storage footprint by 35x through dynamic embedding sizing and support for longer token length, making it feasible to handle large-scale datasets efficiently.

An embedding model is a crucial component of a text retrieval system, as it transforms textual information into dense vector representations. They are typically transformer encoders that process tokens of input text (for example: question, passage) to output an embedding.

This model is ready for commercial use.

The Llama Nemotron Embedding 1B model is a part of the NVIDIA NeMo Retriever collection of NIM, which provide state-of-the-art, commercially-ready models and microservices, optimized for the lowest latency and highest throughput. It features a production-ready information retrieval pipeline with enterprise support. The models that form the core of this solution have been trained using responsibly selected, auditable data sources. With multiple pre-trained models available as starting points, developers can also readily customize them for domain-specific use cases, such as information technology, human resource help assistants, and research & development research assistants.

We are excited to announce the open sourcing of this commercial embedding model. For users interested in deploying this model in production environments, it is also available via the model API in NVIDIA Inference Microservices (NIM) at llama-nemotron-embed-1b-v2.

The Llama Nemotron Embedding 1B model is most suitable for users who want to build a multilingual question-and-answer application over a large text corpus, leveraging the latest dense retrieval technologies.

Use of this model is governed by the NVIDIA Open Model License Agreement. Additional Information: Llama 3.2 Community Model License Agreement.

**Architecture Type:** Transformer
**Network Architecture:** Fine-tuned Llama3.2 1B Retriever

This NeMo embedding model is a transformer encoder - a fine-tuned version of Llama3.2 1b, with 16 layers and an embedding size of 2048, which is trained on public datasets. The AdamW optimizer is employed incorporating 100 warm up steps and 5e-6 learning rate with WarmupDecayLR scheduler. Embedding models for text retrieval are typically trained using a bi-encoder architecture. This involves encoding a pair of sentences (for example, query and chunked passages) independently using the embedding model. Contrastive learning is used to maximize the similarity between the query and the passage that contains the answer, while minimizing the similarity between the query and sampled negative passages not useful to answer the question.

**Input Type:** Text
**Input Format:** List of strings
**Input Parameter:** 1D
**Other Properties Related to Input:** The model's maximum context length is 8192 tokens. Texts longer than maximum length must either be chunked or truncated.

**Output Type:** Floats
**Output Format:** List of float arrays
**Output:** Model outputs embedding vectors of maximum dimension 2048 for each text string (can be configured based on 384, 512, 768, 1024, or 2048).
**Other Properties Related to Output:** N/A

The model supports transformers versions 4.44 through 5.0+.

```
pip install transformers sentence-transformers
```
```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("nvidia/llama-nemotron-embed-1b-v2", trust_remote_code=True)
queries = [
    "how much protein should a female eat",
    "summit define",
]
documents = [
    "As a general guideline, the CDC's average requirement of protein for women ages 19 to 70 is 46 grams per day. But, as you can see from this chart, you'll need to increase that if you're expecting or training for a marathon. Check out the chart below to see how much protein you should be eating each day.",
    "Definition of summit for English Language Learners. : 1  the highest point of a mountain : the top of a mountain. : 2  the highest level. : 3  a meeting or series of meetings between the leaders of two or more governments."
]
query_embeddings = model.encode_query(queries, convert_to_tensor=True)
document_embeddings = model.encode_document(documents, convert_to_tensor=True)
# Compute similarity scores
scores = model.similarity(query_embeddings, document_embeddings)
"""
tensor([[ 0.5968, -0.0454],
        [-0.0336,  0.4613]], device='cuda:0')
"""
```
You can also use transformers directly to run the model. The model supports transformers versions 4.44 through 5.0+.

```
pip install transformers
```
```
import torch
import torch.nn.functional as F
from transformers import AutoTokenizer, AutoModel
def average_pool(last_hidden_states, attention_mask):
    """Average pooling with attention mask."""
    last_hidden_states_masked = last_hidden_states.masked_fill(~attention_mask[..., None].bool(), 0.0)
    embedding = last_hidden_states_masked.sum(dim=1) / attention_mask.sum(dim=1)[..., None]
    embedding = F.normalize(embedding, dim=-1)
    return embedding
tokenizer = AutoTokenizer.from_pretrained("nvidia/llama-nemotron-embed-1b-v2")
model = AutoModel.from_pretrained("nvidia/llama-nemotron-embed-1b-v2", trust_remote_code=True)
model = model.to("cuda:0")
model.eval()
query_prefix = "query:"
document_prefix = "passage:"
queries = [
    "how much protein should a female eat",
    "summit define",
]
documents = [
    "As a general guideline, the CDC's average requirement of protein for women ages 19 to 70 is 46 grams per day. But, as you can see from this chart, you'll need to increase that if you're expecting or training for a marathon. Check out the chart below to see how much protein you should be eating each day.",
    "Definition of summit for English Language Learners. : 1  the highest point of a mountain : the top of a mountain. : 2  the highest level. : 3  a meeting or series of meetings between the leaders of two or more governments."
]
queries = [f"{query_prefix} {query}" for query in queries]
documents = [f"{document_prefix} {document}" for document in documents]
batch_queries = tokenizer(queries, padding=True, truncation=True, return_tensors='pt').to("cuda:0")
batch_documents = tokenizer(documents, padding=True, truncation=True, return_tensors='pt').to("cuda:0")
with torch.no_grad():
    outputs_queries = model(**batch_queries)
    outputs_documents = model(**batch_documents)
# Average Pooling
embeddings_queries = average_pool(outputs_queries.last_hidden_state, batch_queries["attention_mask"])
print("Query embeddings:")
print(embeddings_queries)
print(embeddings_queries.shape)
#torch.Size([2, 2048])
embeddings_documents = average_pool(outputs_documents.last_hidden_state, batch_documents["attention_mask"])
print("\nDocument embeddings:")
print(embeddings_documents)
print(embeddings_documents.shape)
#torch.Size([2, 2048])
# Compute similarity scores
scores = (embeddings_queries @ embeddings_documents.T)
print("\nSimilarity scores:")
print(scores.tolist())
#Similarity scores:
#[[0.5968121290206909, -0.04534469544887543], [-0.03361201286315918, 0.46140915155410767]]
```
1. Ensure you are using `vllm>=0.14.0` .
2. Start the vLLM server with the following command:

Minimal command (required):

