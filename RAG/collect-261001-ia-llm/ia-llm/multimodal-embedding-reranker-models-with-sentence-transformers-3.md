---
id: collect-261001-ia-llm/ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers-3
title: "For image support"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["attention", "compute", "embedding", "embeddings", "llama", "memory", "multimodal", "nvidia", "omni", "parameters", "quantization", "qwen"]
source: docs/RAG/collect-261001-ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [261, 389]
sha256: 9711959a792ad1ee1152d7d12c90ecb3755a4b00d00c5718cffbd0dddd6b38e5
---

# For image support

```
model = SentenceTransformer(
    "Qwen/Qwen3-VL-Embedding-2B",
    model_kwargs={"attn_implementation": "flash_attention_2", "torch_dtype": "bfloat16"},
    processor_kwargs={"min_pixels": 28 * 28, "max_pixels": 600 * 600},
)
```
- `processor_kwargs` controls how inputs are preprocessed (e.g., image resolution bounds). Higher`max_pixels` means higher quality but more memory and compute. These are passed directly to`AutoProcessor.from_pretrained(...)` .
- `model_kwargs` controls how the underlying model is loaded (e.g., precision, attention implementation). These are passed directly to the appropriate`AutoModel.from_pretrained(...)` call (e.g.,`AutoModel` ,`AutoModelForCausalLM` ,`AutoModelForSequenceClassification` , etc., depending on the configuration of the model modules).

See the SentenceTransformer API Reference documentation for more details on these kwargs.

In Sentence Transformers v5.4, `tokenizer_kwargs` has been renamed to `processor_kwargs` to reflect that multimodal models use processors rather than just tokenizers. The old name is still accepted but deprecated.


Here are the multimodal models supported in v5.4, also available in the v5.4 integrations collection:

| Model | Parameters | Modalities | Revision | 
|---|---|---|---|
| Qwen/Qwen3-VL-Embedding-2B | 2B | Text, Image, Video | No `revision` needed | 
| Qwen/Qwen3-VL-Embedding-8B | 8B | Text, Image, Video | No `revision` needed | 
| nvidia/llama-nemotron-embed-vl-1b-v2 | 1.7B | Text, Image | No `revision` needed | 
| nvidia/omni-embed-nemotron-3b | 4.7B | Text, Image | No `revision` needed | 
| LCO-Embedding/LCO-Embedding-Omni-3B | 5B | Text, Image, Audio, Video | No `revision` needed | 
| LCO-Embedding/LCO-Embedding-Omni-7B | 9B | Text, Image, Audio, Video | No `revision` needed | 
| BidirLM/BidirLM-Omni-2.5B-Embedding | 2.5B | Text, Image, Audio | No `revision` needed | 
| BAAI/BGE-VL-base | 0.1B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-large | 0.4B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-MLLM-S1 | 8B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-MLLM-S2 | 8B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-v1.5-zs | 8B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-v1.5-mmeb | 8B | Text, Image | No `revision` needed | 
| BAAI/BGE-VL-Screenshot | 4B | Text, Image | No `revision` needed | 
| royokong/e5-v | 8B | Text, Image | No `revision` needed | 
| eagerworks/eager-embed-v1 | 4B | Text, Image | `revision="refs/pr/2"` | 
| nomic-ai/nomic-embed-multimodal-3b | 5B | Text, Image | `revision="refs/pr/4"` | 
| nomic-ai/nomic-embed-multimodal-7b | 9B | Text, Image | `revision="refs/pr/3"` | 
| Haon-Chen/e5-omni-3B | 5B | Text, Image, Audio, Video | `revision="refs/pr/2"` | 
| Haon-Chen/e5-omni-7B | 9B | Text, Image, Audio, Video | `revision="refs/pr/1"` | 

| Model | Parameters | Modalities | Revision | 
|---|---|---|---|
| Qwen/Qwen3-VL-Reranker-2B | 2B | Text, Image, Video | No `revision` needed | 
| Qwen/Qwen3-VL-Reranker-8B | 8B | Text, Image, Video | No `revision` needed | 
| nvidia/llama-nemotron-rerank-vl-1b-v2 | 2B | Text, Image | No `revision` needed | 
| jinaai/jina-reranker-m0 | 2B | Text, Image | No `revision` needed | 

| Model | Parameters | Revision | 
|---|---|---|
| Qwen/Qwen3-Reranker-0.6B | 0.6B | No `revision` needed | 
| Qwen/Qwen3-Reranker-4B | 4B | No `revision` needed | 
| Qwen/Qwen3-Reranker-8B | 8B | No `revision` needed | 
| mixedbread-ai/mxbai-rerank-base-v2 | 0.5B | No `revision` needed | 
| mixedbread-ai/mxbai-rerank-large-v2 | 2B | No `revision` needed | 
| ContextualAI/ctxl-rerank-v2-instruct-multilingual-1b | 1B | `revision="refs/pr/2"` | 
| ContextualAI/ctxl-rerank-v2-instruct-multilingual-2b | 3B | `revision="refs/pr/1"` | 
| ContextualAI/ctxl-rerank-v2-instruct-multilingual-6b | 7B | `revision="refs/pr/1"` | 

## Click here for a text-only reranker usage example

```
from sentence_transformers import CrossEncoder
model = CrossEncoder("mixedbread-ai/mxbai-rerank-base-v2")
query = "How do I bake sourdough bread?"
documents = [
    "Sourdough bread requires a starter made from flour and water, fermented over several days.",
    "The history of bread dates back to ancient Egypt around 8000 BCE.",
    "To bake sourdough, mix your starter with flour, water, and salt, then let it rise overnight.",
    "Rye bread is a popular alternative to wheat-based breads in Northern Europe.",
]
pairs = [(query, doc) for doc in documents]
scores = model.predict(pairs)
print(scores)
# [ 7.3077507 -2.6217823  8.724761  -2.2488995]
rankings = model.rank(query, documents)
for rank in rankings:
    print(f"{rank['score']:.4f}\t{documents[rank['corpus_id']]}")
# 8.7248  To bake sourdough, mix your starter with flour, water, and salt, then let it rise overnight.
# 7.3078  Sourdough bread requires a starter made from flour and water, fermented over several days.
# -2.2489 Rye bread is a popular alternative to wheat-based breads in Northern Europe.
# -2.6218 The history of bread dates back to ancient Egypt around 8000 BCE.
```
The older CLIP models continue to be supported:

| Model | ImageNet Zero-Shot Top-1 Accuracy | Notes | 
|---|---|---|
| sentence-transformers/clip-ViT-L-14 | 75.4 |  | 
| sentence-transformers/clip-ViT-B-16 | 68.1 |  | 
| sentence-transformers/clip-ViT-B-32 | 63.3 |  | 
| sentence-transformers/clip-ViT-B-32-multilingual-v1 | N/A | Multilingual text encoder, 50+ languages | 

These simple CLIP models still work well on lower-resource hardware.

## Click here for a CLIP usage example

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("sentence-transformers/clip-ViT-L-14")
images = [
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg",
    "https://huggingface.co/datasets/huggingface/cats-image/resolve/main/cats_image.jpeg"
]
texts = ["A green car", "A bee on a flower", "Some cats on a couch", "One cat sitting in the window"]
image_embeddings = model.encode(images)
text_embeddings = model.encode(texts)
print(image_embeddings.shape, text_embeddings.shape)
# (3, 768) (4, 768)
similarities = model.similarity(image_embeddings, text_embeddings)
print(similarities)
# tensor([[0.2208, 0.1042, 0.0617, 0.0907],  First image (car) is most similar to "A green car"
#         [0.1205, 0.2303, 0.0632, 0.0917],  Second image (bee) is most similar to "A bee on a flower"
#         [0.1107, 0.0196, 0.2425, 0.1162]]) Third image (multiple cats) is most similar to "Some cats on a couch"
```
To learn how to finetune these multimodal models on your own data, see the companion blogpost: Training and Finetuning Multimodal Embedding & Reranker Models with Sentence Transformers.

- Sentence Transformers models on the Hub
- Sentence Transformers datasets on the Hub
- v5.4 Integrations Collection

The training companion to this post and adjacent Sentence Transformers guides:

- Training and Finetuning Multimodal Embedding & Reranker Models with Sentence Transformers: the direct training companion to this post, with a Visual Document Retrieval walkthrough.
- Training and Finetuning Embedding Models with Sentence Transformers: the general training guide for text-only bi-encoder embedding models.
- Training and Finetuning Reranker Models with Sentence Transformers: Cross Encoder (reranker) training, applicable to text-only and multimodal rerankers.
- Training and Finetuning Sparse Embedding Models with Sentence Transformers: SPLADE training for sparse retrieval.
- 🪆 Introduction to Matryoshka Embedding Models: variable-size embeddings; also applied to multimodal models in the training companion post.
- Train 400x faster Static Embedding Models with Sentence Transformers: CPU-friendly text embedding models.
- Binary and Scalar Embedding Quantization for Significantly Faster & Cheaper Retrieval: post-training compression of embedding vectors.
