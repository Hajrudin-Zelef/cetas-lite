---
id: collect-261001-ia-llm/ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers-2
title: "For image support"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["compute", "deepseek", "embedding", "embeddings", "llama", "memory", "multimodal", "nvidia", "qwen", "reranker", "revenue"]
source: docs/RAG/collect-261001-ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [114, 260]
sha256: ded1f92d6bcda3ffa9d9bbaaa54b9e2678a73fb4f613948b75a1c692e01f95dd
---

# For image support

```
from sentence_transformers import CrossEncoder
model = CrossEncoder("Qwen/Qwen3-VL-Reranker-2B")
query = "A green car parked in front of a yellow building"
documents = [
    # Image documents (URL or local file path)
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg",
    # Text document
    "A vintage Volkswagen Beetle painted in bright green sits in a driveway.",
    # Combined text + image document
    {
        "text": "A car in a European city",
        "image": "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    },
]
rankings = model.rank(query, documents)
for rank in rankings:
    print(f"{rank['score']:.4f}\t(document {rank['corpus_id']})")
"""
0.9375  (document 0)
0.5000  (document 3)
-1.2500 (document 2)
-2.4375 (document 1)
"""
```
The reranker correctly identifies the car image (document 0) as the most relevant result, followed by the combined text+image document about a car in a European city (document 3). The bee image (document 1) scores lowest. Keep in mind that the modality gap can influence absolute scores: text-image pair scores may occupy a different range than text-text or image-image pair scores.

You can also check which modalities a reranker supports using `modalities` and `supports()`, just like with embedding models:

```
print(model.modalities)
# ['text', 'image', 'video', 'message']
print(model.supports("image"))
# True
# Check if the model supports a specific pair of modalities
print(model.supports(("image", "text")))
# True
```
You can also use `predict()` to get raw relevance scores for specific pairs of inputs:

```
from sentence_transformers import CrossEncoder
model = CrossEncoder("jinaai/jina-reranker-m0", trust_remote_code=True)
scores = model.predict([
    ("A green car", "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg"),
    ("A bee on a flower", "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg"),
    ("A green car", "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg"),
])
print(scores)
# [0.9389156  0.96922314 0.46063158]
```
A common pattern is to use an embedding model for fast initial retrieval, then refine the top results with a reranker:

```
from sentence_transformers import SentenceTransformer, CrossEncoder
# Step 1: Retrieve with an embedding model
embedder = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
query = "revenue growth chart"
query_embedding = embedder.encode_query(query)
# Pre-compute corpus embeddings (do this once, then store them)
document_screenshots = [
    "path/to/doc1.png",
    "path/to/doc2.png",
    # ... potentially millions of document screenshots
]
corpus_embeddings = embedder.encode_document(document_screenshots, show_progress_bar=True)
# Simple cosine similarity retrieval, viable as long as embeddings fit in memory
similarities = embedder.similarity(query_embedding, corpus_embeddings)
top_k_indices = similarities.argsort(descending=True)[0][:10]
# Step 2: Rerank the top-k results with a reranker model
reranker = CrossEncoder("nvidia/llama-nemotron-rerank-vl-1b-v2", trust_remote_code=True)
top_k_documents = [document_screenshots[i] for i in top_k_indices]
rankings = reranker.rank(query, top_k_documents)
for rank in rankings:
    print(f"{rank['score']:.4f}\t{top_k_documents[rank['corpus_id']]}")
```
Since the corpus embeddings are pre-computed, the initial retrieval is fast even over millions of documents. The reranker then provides more accurate scoring over the smaller candidate set.

Multimodal models accept a variety of input formats. Here's a summary of what you can pass to `model.encode()`:

| Modality | Accepted Formats | 
|---|---|
| **Text** | - Strings | 
| **Image** | - `PIL.Image.Image` objects- File paths (e.g. `"./photo.jpg"` )- URLs (e.g. `"https://.../image.jpg"` )- Numpy arrays, torch tensors | 
| **Audio** | - File paths (e.g. `"./audio.wav"` )- URLs (e.g. `"https://.../audio.wav"` )- Numpy/torch arrays - Dicts with `"array"` and`"sampling_rate"` keys- `torchcodec.AudioDecoder` instances | 
| **Video** | - File paths (e.g. `"./video.mp4"` )- URLs (e.g. `"https://.../video.mp4"` )- Numpy/torch arrays - Dicts with `"array"` and`"video_metadata"` keys- `torchcodec.VideoDecoder` instances | 
| **Multimodal** | - Dicts mapping modality names to values, e.g. `{"text": "a caption", "image": "https://.../image.jpg"}`Valid keys: `"text"` ,`"image"` ,`"audio"` ,`"video"` | 
| **Message** | - List of message dicts with `"role"` and`"content"` keys,e.g. `[{"role": "user", "content": [...]}]` | 

You can check which modalities a model supports using the `modalities` property and `supports()` method:

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
# List all supported modalities
print(model.modalities)
# ['text', 'image', 'video', 'message']
# Check for a specific modality
print(model.supports("image"))
# True
print(model.supports("audio"))
# False
```
The `"message"` modality indicates that the model accepts chat-style message inputs with interleaved content. In practice, you rarely need to use this directly. When you pass strings, URLs, or multimodal dicts, the model converts them to the appropriate message format internally. Sentence Transformers supports two message formats:

- **Structured** (most VLMs, e.g. Qwen3-VL): Content is a list of typed dicts, e.g.`[{"type": "text", "text": "..."}, {"type": "image", "image": ...}]`
- **Flat** (e.g. Deepseek-V3): Content is a direct value, e.g.`"some text"`

The format is auto-detected from the model's chat template.

Since all inputs get converted into the same message format internally, you can mix input types in a single `encode()` call:

```
embeddings = model.encode([
    # A text input
    "A green car parked in front of a yellow building",
    # An image input (URL)
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    # A combined text + image input
    {
        "text": "A car in a European city",
        "image": "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    },
])
```
## Click here if you need to pass raw message inputs

If a model doesn't follow either format and you need full control, you can pass raw message dicts with `role` and `content` keys directly:

```
embeddings = model.encode([
    [
        {
            "role": "user",
            "content": [
                {"type": "image", "image": "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg"},
                {"type": "text", "text": "Describe this vehicle."},
            ],
        }
    ],
])
```
This bypasses the automatic format conversion and passes the messages directly to the processor's `apply_chat_template()`.

You may want to control image resolution bounds or model precision. Use `processor_kwargs` and `model_kwargs` when loading the model:

