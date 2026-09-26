---
id: collect-240926-huggingface/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face-2
title: "Load the model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Cohere", "Google"]
dates: []
keywords: ["attention", "cohere", "compute", "embedding", "embeddings", "gemini", "memory", "multimodal", "qwen", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face.md
source_anchor: ""
source_lines: [60, 206]
sha256: b820d1eacca5978c5bd2fa6bfaef6552f76df7aa63148a1fbeef944e285a78f2
---

# Load the model

| Model | Size | Mean (Task) | Mean (Type) | Bitxt Mining | Class. | Clust. | Inst. Retri. | Multi. Class. | Pair. Class. | Rerank | Retri. | STS | 
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| NV-Embed-v2 | 7B | 56.29 | 49.58 | 57.84 | 57.29 | 40.80 | 1.04 | 18.63 | 78.94 | 63.82 | 56.72 | 71.10 | 
| GritLM-7B | 7B | 60.92 | 53.74 | 70.53 | 61.83 | 49.75 | 3.45 | 22.77 | 79.94 | 63.78 | 58.31 | 73.33 | 
| BGE-M3 | 0.6B | 59.56 | 52.18 | 79.11 | 60.35 | 40.88 | -3.11 | 20.1 | 80.76 | 62.79 | 54.60 | 74.12 | 
| multilingual-e5-large-instruct | 0.6B | 63.22 | 55.08 | 80.13 | 64.94 | 50.75 | -0.40 | 22.91 | 80.86 | 62.61 | 57.12 | 76.81 | 
| gte-Qwen2-1.5B-instruct | 1.5B | 59.45 | 52.69 | 62.51 | 58.32 | 52.05 | 0.74 | 24.02 | 81.58 | 62.58 | 60.78 | 71.61 | 
| gte-Qwen2-7b-Instruct | 7B | 62.51 | 55.93 | 73.92 | 61.55 | 52.77 | 4.94 | 25.48 | 85.13 | 65.55 | 60.08 | 73.98 | 
| text-embedding-3-large | - | 58.93 | 51.41 | 62.17 | 60.27 | 46.89 | -2.68 | 22.03 | 79.17 | 63.89 | 59.27 | 71.68 | 
| Cohere-embed-multilingual-v3.0 | - | 61.12 | 53.23 | 70.50 | 62.95 | 46.89 | -1.89 | 22.74 | 79.88 | 64.07 | 59.16 | 74.80 | 
| Gemini Embedding | - | 68.37 | 59.59 | 79.28 | 71.82 | 54.59 | 5.18 | **29.16** | 83.63 | 65.58 | 67.71 | 79.40 | 
| Qwen3-Embedding-0.6B | 0.6B | 64.33 | 56.00 | 72.22 | 66.83 | 52.33 | 5.09 | 24.59 | 80.83 | 61.41 | 64.64 | 76.17 | 
| Qwen3-Embedding-4B | 4B | 69.45 | 60.86 | 79.36 | 72.33 | 57.15 | **11.56** | 26.77 | 85.05 | 65.08 | 69.60 | 80.86 | 
| Qwen3-Embedding-8B | 8B | **70.58** | **61.69** | **80.89** | **74.00** | **57.65** | 10.06 | 28.66 | **86.40** | **65.63** | **70.88** | **81.08** | 
| Qwen3-VL-Embedding-2B | 2B | 63.87 | 55.84 | 69.51 | 65.86 | 52.50 | 3.87 | 26.08 | 78.50 | 64.80 | 67.12 | 74.29 | 
| Qwen3-VL-Embedding-8B | 8B | 67.88 | 58.88 | 77.48 | 71.95 | 55.82 | 4.46 | 28.59 | 81.08 | 65.72 | 69.41 | 75.41 | 

Install Sentence Transformers with `pip install sentence-transformers`, then use the model like this:

```
from sentence_transformers import SentenceTransformer
# Load the model
model = SentenceTransformer("Qwen/Qwen3-VL-Embedding-8B")
# Text queries
queries = [
    "A woman playing with her dog on a beach at sunset.",
    "Pet owner training dog outdoors near water.",
    "Woman surfing on waves during a sunny day.",
    "City skyline view from a high-rise building at night.",
]
# Documents: text, image, and text+image
documents = [
    "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust.",
    "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg",
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust.", "image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"},
]
# Encode queries and documents
query_embeddings = model.encode(queries)
doc_embeddings = model.encode(documents)
print(query_embeddings.shape, doc_embeddings.shape)
# (4, 4096) (3, 4096)
# Compute similarities
similarities = model.similarity(query_embeddings, doc_embeddings)
print(similarities)
# tensor([[0.7438, 0.6556, 0.6244],
#         [0.4430, 0.3323, 0.3929],
#         [0.3685, 0.2310, 0.2874],
#         [0.0602, -0.0162, 0.0167]])
```
By default, all inputs are wrapped with the `"Represent the user's input."` instruction via a system prompt. You can customize this by passing a different prompt:

```
# With a custom prompt
model.encode(queries, prompt="Retrieve relevant documents for the query.")
```
- **requirements**

```
transformers>=4.57.0
qwen-vl-utils>=0.0.14
torch==2.8.0
```
```
from scripts.qwen3_vl_embedding import Qwen3VLEmbedder
# Define a list of query texts
queries = [
    {"text": "A woman playing with her dog on a beach at sunset."},
    {"text": "Pet owner training dog outdoors near water."},
    {"text": "Woman surfing on waves during a sunny day."},
    {"text": "City skyline view from a high-rise building at night."}
]
# Define a list of document texts and images
documents = [
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust."},
    {"image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"},
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust.", "image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"}
]
# Specify the model path
model_name_or_path = "Qwen/Qwen3-VL-Embedding-8B"
# Initialize the Qwen3VLEmbedder model
model = Qwen3VLEmbedder(model_name_or_path=model_name_or_path)
# We recommend enabling flash_attention_2 for better acceleration and memory saving,
# model = Qwen3VLEmbedder(model_name_or_path=model_name_or_path, torch_dtype=torch.float16, attn_implementation="flash_attention_2")
# Combine queries and documents into a single input list
inputs = queries + documents
# Process the inputs to get embeddings
embeddings = model.process(inputs)
# Compute similarity scores between query embeddings and document embeddings
similarity_scores = (embeddings[:4] @ embeddings[4:].T)
# Print out the similarity scores in a list format
print(similarity_scores.tolist())
# [[0.74267578125, 0.6630859375, 0.6328125], [0.443603515625, 0.33349609375, 0.396484375], [0.3671875, 0.2354736328125, 0.289306640625], [0.060821533203125, -0.01557159423828125, 0.0165863037109375]]
```
```
import argparse
import numpy as np
import os
from typing import List, Dict, Any
from vllm import LLM, EngineArgs
from vllm.multimodal.utils import fetch_image
# Define a list of query texts
queries = [
    {"text": "A woman playing with her dog on a beach at sunset."},
    {"text": "Pet owner training dog outdoors near water."},
    {"text": "Woman surfing on waves during a sunny day."},
    {"text": "City skyline view from a high-rise building at night."}
]
# Define a list of document texts and images
documents = [
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust."},
    {"image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"},
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust.", "image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"}
]
def format_input_to_conversation(input_dict: Dict[str, Any], instruction: str = "Represent the user's input.") -> List[Dict]:
    content = []
    
    text = input_dict.get('text')
    image = input_dict.get('image')
    
    if image:
        image_content = None
        if isinstance(image, str):
            if image.startswith(('http', 'https', 'oss')):
                image_content = image
            else:
                abs_image_path = os.path.abspath(image)
                image_content = 'file://' + abs_image_path
        else:
            image_content = image
        
        if image_content:
            content.append({
                'type': 'image', 
                'image': image_content,
            })
    
    if text:
        content.append({'type': 'text', 'text': text})
    
    if not content:
        content.append({'type': 'text', 'text': ""})
    
    conversation = [
        {"role": "system", "content": [{"type": "text", "text": instruction}]},
        {"role": "user", "content": content}
    ]
    
