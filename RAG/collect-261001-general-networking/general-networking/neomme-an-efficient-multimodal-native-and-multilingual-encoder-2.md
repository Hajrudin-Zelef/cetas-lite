---
id: collect-261001-general-networking/general-networking/neomme-an-efficient-multimodal-native-and-multilingual-encoder-2
title: "accelerate is an optional dependency needed only when using device_map=\"auto\"."
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Nvidia"]
dates: []
keywords: ["attention", "benchmark", "benchmarks", "compute", "cost", "embedding", "embeddings", "energy", "fine-tuning", "gpu", "inference", "multimodal"]
source: docs/RAG/collect-261001-general-networking/neomme-an-efficient-multimodal-native-and-multilingual-encoder.md
source_anchor: ""
source_lines: [47, 150]
sha256: d619d61b4a381e7d44a3055f46cf0fac389c23a4353f0a30b05ac4cd8e637303
---

# accelerate is an optional dependency needed only when using device_map="auto".

| Visual document retrieval performance on the ViDoRe benchmarks. |  |  |  |  | 
|---|---|---|---|---|
| Model details |  | ViDoRe (nDCG@ *k* ) |  |  | 
|---|---|---|---|---|
| Model | Params. | v3 (@10) | v2 (@5) | v1 (@5) | 
| <300M |  |  |  |  | 
| ColModernVBERT | 250M | 0.261 <sup>†</sup> | 0.407 <sup>‡</sup> | 0.806 <sup>‡</sup> | 
| ColSmol-256M <sup>†</sup> | 256M | 0.207 | 0.348 | 0.797 | 
| *NeoMME* -260M<sup>‡</sup> | 260M | **0.523** | **0.522** | **0.860** | 
| 300M to 1B |  |  |  |  | 
| ColSmol-500M | 500M | 0.340 <sup>‡</sup> | 0.455 <sup>†</sup> | 0.825 <sup>†</sup> | 
| Vultron Flash <sup>†</sup> | 850M | **0.565** | **0.604** | **0.882** | 
| *NeoMME* -800M<sup>‡</sup> | 800M | 0.556 | 0.559 | 0.874 | 
| >1B |  |  |  |  | 
| ColQwen2.5-v0.2 <sup>†</sup> | 3.75B | **0.524** | **0.601** | **0.895** | 
| ColPali v1.3 <sup>†</sup> | 2.92B | 0.430 | 0.547 | 0.848 | 

<sup>†</sup> Scores from MTEB. <sup>‡</sup> Results from our own evaluations.

Late-interaction storage scales linearly with the number of vectors in the output embedding. Higher-resolution images contain more patches, so they produce larger embeddings. For example, a 2048×2048 square page produces embeddings containing 4,200 vectors with *NeoMME*-Retriever, or about 2.1 MB in float32. Across the ViDoRe v3 benchmark, the measured average is about 1.5 MB per document.

To reduce the storage footprint of the late-interaction index, we combine two complementary compression methods:

1. Hierarchical token pooling clusters similar document vectors in a given multi-vector embedding and replaces each cluster with its mean, hence reducing the number of vectors stored for each page.
2. Asymmetric quantization quantizes document embeddings to int8 or binary. Because query embeddings are not stored and only generated on-the-fly, they can be kept at a higher precision.

We tested this setup on ViDoRe v3. With a pooling factor 10 and int8 queries and documents, storage decreased from about 1.5 MB to 39 kB per page, a 39× reduction, while keeping more than 99% of the baseline nDCG@10. A more aggressive configuration uses pooling factor 8, int8 queries, and binary documents. That version uses 6 kB per page (255× smaller) and keeps more than 95% of the original retrieval quality.

Users can pick a compression setting from that frontier based on storage budget and required retrieval quality.

Before you can search a corpus, a retriever model must turn your documents into embeddings, which will be stored in a vector store like Qdrant, Weaviate, or Milvus. Faster encoding makes building and adding new documents to the index faster, thus reducing the GPU uptime and compute cost required.

So we measured image encoding speeds for *NeoMME*-Retriever against other multimodal document retrievers. We used preprocessed image tensors and calibrated the batch size separately for each model and image size. At a matched 2048×2048 input size on one NVIDIA L40S, *NeoMME*-Retriever-260M encodes about 51 pages per second, nearly twice ColModernVBERT's 26 pages per second. Both 260M and 800M *NeoMME*-Retriever models are also faster than the other models we compared on smaller input images.

*NeoMME*-Retriever (260M and 800M) returns dense and multi-vector embeddings together. The example below scores two text queries against two document-page images with MeanMaxSim late interaction and dense cosine similarity.

## Click to see the complete 🤗 `transformers` example snippet

```
# accelerate is an optional dependency needed only when using device_map="auto".
pip install -U accelerate "transformers @ git+https://github.com/huggingface/transformers.git@main" "sentence-transformers>=6.0.0"
```
```
from typing import Any, Literal
import requests
import torch
from PIL import Image
from sentence_transformers.util import cos_sim, mean_maxsim
from transformers import BatchFeature, NeoMMEForRetrieval, NeoMMEProcessor
def encode(    messages: list[list[dict[str, Any]]],
    task: Literal["query", "document"],
return processor.apply_chat_template(
        messages,
        task=task,
        tokenize=True,
        return_dict=True,
        return_tensors="pt",
        processor_kwargs={"padding": "longest"},
    )
model_name = "Hcompany/NeoMME-260M-Retriever"
processor = NeoMMEProcessor.from_pretrained(model_name)
model = NeoMMEForRetrieval.from_pretrained(model_name, device_map="auto")
# Document images (our corpus)
image_urls = [
    "https://github.com/tonywu71/colpali-cookbooks/blob/6ef1332da6bcb48c7ef1f19b25bfa555be7031a8/examples/data/shift_kazakhstan.jpg?raw=true",
    "https://github.com/tonywu71/colpali-cookbooks/blob/6ef1332da6bcb48c7ef1f19b25bfa555be7031a8/examples/data/energy_electricity_generation.jpg?raw=true",
]
documents = [Image.open(requests.get(url, stream=True).raw) for url in image_urls]
# Queries
queries = [
    "Quelle partie de la production pétrolière du Kazakhstan provient de champs en mer ?",
    "Which hour of the day had the highest overall electricity generation in 2019?",
]
document_messages = [
    [{"role": "user", "content": [{"type": "image", "image": document}]}] for document in documents
]
query_messages = [[{"role": "user", "content": query}] for query in queries]
inputs_documents = encode(document_messages, "document").to(model.device)
inputs_text = encode(query_messages, "query").to(model.device)
with torch.inference_mode():
    document_outputs = model(**inputs_documents)
    query_outputs = model(**inputs_text)
late_scores = mean_maxsim(
    query_outputs.embeddings,
    document_outputs.embeddings,
    a_mask=inputs_text["attention_mask"],
    b_mask=inputs_documents["attention_mask"],
)
dense_scores = cos_sim(query_outputs.dense_embeddings, document_outputs.dense_embeddings)
# Expected: late_scores[0, 0] > late_scores[0, 1] and late_scores[1, 1] > late_scores[1, 0].
print(late_scores, dense_scores)
```
We provide separate dense and late-interaction checkpoints for fine-tuning with Sentence Transformers v6. Following the same pattern as text encoders such as ModernBERT, Sentence Transformers loads the backbone through `NeoMMEModel` rather than the dual-head `NeoMMEForRetrieval` class. Sentence Transformers currently supports one retrieval head per model, so each checkpoint lets you fine-tune the dense or late-interaction head independently. To train both heads together, use `NeoMMEForRetrieval` with a custom `Trainer`.

Visual document retrieval can be used as the first stage of a visual retrieval-augmented generation (RAG) system. Unlike text RAG, which retrieves extracted text chunks, visual RAG retrieves the original page images and sends them to a visual language model. The model can then use tables, plots, diagrams, and page layout that text extraction may flatten or omit. Here is how visual RAG works:

1. Indexing: Convert each PDF page to an image, generate an embedding with a retrieval model, and store the embeddings in a vector store.
2. Retrieval: Generate an embedding for the user's query with the same model and retrieve the top-k most relevant pages.
3. Generation: Append the images after the query in the chat message (*e.g.* ,`{query}{img_1}{img_2}...{img_k}` ) and send it to a VLM to generate the answer.

You can test visual RAG directly with *NeoMME*-Retriever in our HF Space: 🤗 tonywu71/neomme-retriever-demo.

*NeoMME* replaces separate pretrained image and text encoders with one long-context bidirectional Transformer. We train it from scratch to process both multilingual text tokens and raw 32×32 image patches.

