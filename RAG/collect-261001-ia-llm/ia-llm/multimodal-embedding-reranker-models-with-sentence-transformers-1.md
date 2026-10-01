---
id: collect-261001-ia-llm/ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers-1
title: "For image support"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Google"]
dates: []
keywords: ["compute", "embedding", "embeddings", "gpu", "inference", "multimodal", "parameters", "qwen", "reranker", "training"]
source: docs/RAG/collect-261001-ia-llm/multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [1, 113]
sha256: 336015319d5d5b6562220b9d92d9763dadac660790de05760aee4322668654dc
---

# For image support

**encode and compare texts, images, audio, and videos**using the same familiar API. In this blogpost, I'll show you how to use these new multimodal capabilities for both embedding and reranking.

Multimodal embedding models map inputs from different modalities into a shared embedding space, while multimodal reranker models score the relevance of mixed-modality pairs. This opens up use cases like visual document retrieval, cross-modal search, and multimodal RAG pipelines.

If you want to train your own multimodal models, check out the companion blogpost: Training and Finetuning Multimodal Embedding & Reranker Models with Sentence Transformers.


Traditional embedding models convert text into fixed-size vectors. Multimodal embedding models extend this by mapping inputs from different modalities (text, images, audio, or video) into a shared embedding space. This means you can compare a text query against image documents (or vice versa) using the same similarity functions you're already familiar with.

Similarly, traditional reranker (Cross Encoder) models compute relevance scores between pairs of texts. Multimodal rerankers can score pairs where one or both elements are images, combined text-image documents, or other modalities.

For example, you can compare a text query against image documents, find video clips matching a description, or build RAG pipelines that work across modalities.

Multimodal models require some extra dependencies. Install the extras for the modalities you need (see Installation for more details):

```
# For image support
pip install -U "sentence-transformers[image]"
# For audio support
pip install -U "sentence-transformers[audio]"
# For video support
pip install -U "sentence-transformers[video]"
# Mix and match as needed
pip install -U "sentence-transformers[image,video,train]"
```
VLM-based models like Qwen3-VL-2B require a GPU with at least ~8 GB of VRAM. For the 8B variants, expect ~20 GB. If you don't have a local GPU, consider using a cloud GPU service or Google Colab. On CPU, these models will be extremely slow; text-only or CLIP models are better suited for CPU inference.


Loading a multimodal embedding model works exactly like loading a text-only model:

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
```
Some models might require a `revision` argument for now if the integration pull requests for the model is still pending. Once they're merged, you'll be able to load them without specifying a revision, like above.


The model automatically detects which modalities it supports, so there's nothing extra to configure. See Processor and Model kwargs if you want to control things like image resolution or model precision.

With a multimodal model loaded, `model.encode()` accepts images alongside text. Images can be provided as URLs, local file paths, or PIL Image objects (see Supported Input Types for all accepted formats):

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
# Encode images from URLs
img_embeddings = model.encode([
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg",
])
print(img_embeddings.shape)
# (2, 2048)
```
You can compute similarities between text embeddings and image embeddings, since the model maps both into the same space:

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
# Encode images
img_embeddings = model.encode([
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg",
])
# Encode text queries (one matching + one hard negative per image)
text_embeddings = model.encode([
    "A green car parked in front of a yellow building",
    "A red car driving on a highway",
    "A bee on a pink flower",
    "A wasp on a wooden table",
])
# Compute cross-modal similarities
similarities = model.similarity(text_embeddings, img_embeddings)
print(similarities)
# tensor([[0.5115, 0.1078],
#         [0.1999, 0.1108],
#         [0.1255, 0.6749],
#         [0.1283, 0.2704]])
```
As expected, "A green car parked in front of a yellow building" is most similar to the car image (0.51), and "A bee on a pink flower" is most similar to the bee image (0.67). The hard negatives ("A red car driving on a highway", "A wasp on a wooden table") correctly receive lower scores.

You might notice that even the best matching scores (0.51, 0.67) aren't very close to 1.0. This is due to the modality gap: embeddings from different modalities tend to cluster in separate regions of the space. Cross-modal similarities are typically lower than within-modal ones (e.g., text-to-text), but the relative ordering is preserved, so retrieval still works well.

For retrieval tasks, `encode_query()` and `encode_document()` are the recommended methods. Many retrieval models prepend different instruction prompts depending on whether the input is a query or a document, similar to how chat models might apply different system prompts depending on the goal. Model authors can specify their prompts in the model config, and `encode_query()` / `encode_document()` automatically load and apply the correct one:

- `encode_query()` uses the model's`"query"` prompt (if available) and sets`task="query"` .
- `encode_document()` uses the first available prompt from`"document"` ,`"passage"` , or`"corpus"` , and sets`task="document"` .

Under the hood, both are thin wrappers around `encode()`, they just handle prompt selection for you. Here's what cross-modal retrieval looks like:

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-2B")
# Encode text queries with the query prompt
query_embeddings = model.encode_query([
    "Find me a photo of a vehicle parked near a building",
    "Show me an image of a pollinating insect",
])
# Encode document screenshots with the document prompt
doc_embeddings = model.encode_document([
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/transformers/tasks/car.jpg",
    "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/bee.jpg",
])
# Compute similarities
similarities = model.similarity(query_embeddings, doc_embeddings)
print(similarities)
# tensor([[0.3907, 0.1490],
#         [0.1235, 0.4872]])
```
These methods accept the same input types as `encode()` (images, URLs, multimodal dicts, etc.) and pass through the same parameters. For models without specialized query/document prompts, they behave identically to `encode()`.

Multimodal reranker (CrossEncoder) models score the relevance between pairs of inputs, where each element can be text, an image, audio, video, or a combination. They tend to outperform embedding models in terms of quality, but are slower since they process each pair individually. The currently available pretrained multimodal rerankers focus on text and image inputs, but the architecture supports any modality that the underlying model can handle.

The `rank()` method scores and ranks a list of documents against a query, supporting mixed modalities:

