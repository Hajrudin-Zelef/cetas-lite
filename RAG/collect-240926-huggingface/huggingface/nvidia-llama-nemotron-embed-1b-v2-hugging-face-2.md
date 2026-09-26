---
id: collect-240926-huggingface/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face-2
title: "Compute similarity scores"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "OpenAI", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["benchmark", "benchmarks", "embedding", "embeddings", "fine-tuning", "gpu", "gpus", "inference", "license", "llama", "mistral", "nvidia"]
source: docs/RAG/clean_en/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face.md
source_anchor: ""
source_lines: [118, 248]
sha256: 1599045736432e3c1b6d1a55447bb6a737ff88131f7beb06c881914f958e010a
---

# Compute similarity scores

```
vllm serve \
    nvidia/llama-nemotron-embed-1b-v2 \
    --trust-remote-code
```
If you already have a local copy of the model, you can also pass the local path instead of the HF repo ID.

Optional flags:

- `--dtype <float32|bfloat16|float16>` to force precision (the default is`auto` , which resolves from model config; this model defaults to BF16).
- `--data-parallel-size <num_gpus_to_use>` for multi-GPU serving.
- `--port 8000` to set the server port.

Online serving example (OpenAI SDK):

```
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:8000/v1",
    api_key="EMPTY",  # required by OpenAI SDK; ignored by default in local vLLM
)
response = client.embeddings.create(
    input=['query: summit define'],
    model="nvidia/llama-nemotron-embed-1b-v2",
)
response.data[0].embedding
```
Offline inference example (Python API, no server required):

```
from vllm import LLM
llm = LLM(
    model="nvidia/llama-nemotron-embed-1b-v2",
    runner="pooling",
    trust_remote_code=True,
)
outputs = llm.embed(["query: summit define", "passage: a summit is a meeting"])
for output in outputs:
    print(len(output.outputs.embedding))  # 2048
```
**Runtime Engine:** Llama Nemotron embedding NIM
**Supported Hardware Microarchitecture Compatibility**: NVIDIA Ampere, NVIDIA Hopper, NVIDIA Lovelace
**Supported Operating System(s):** Linux

Llama Nemotron Embedding 1B v2 Short Name: llama-nemotron-embed-1b-v2

The development of large-scale public open-QA datasets has enabled tremendous progress in powerful embedding models. However, one popular dataset named MS MARCO restricts commercial licensing, limiting the use of these models in commercial settings. To address this, NVIDIA created its own training dataset blend based on public QA datasets, which each have a license for commercial applications.

**Data Collection Method by dataset**: Automated, Unknown

**Labeling Method by dataset**: Automated, Unknown

**Properties:** Semi-supervised pre-training on 12M samples from public datasets and fine-tuning on 1M samples from public datasets.

Properties: We evaluated the NeMo Rtriever embdding model in comparison to literature open & commercial retriever models on academic benchmarks for question-answering - NQ, HotpotQA and FiQA (Finance Q&A) from BeIR benchmark and TechQA dataset. Note that the model was evaluated offline on A100 GPUs using the model's PyTorch checkpoint. In this benchmark, the metric used was Recall@5.

| Open & Commercial Retrieval Models | Average Recall@5 on NQ, HotpotQA, FiQA, TechQA dataset | 
|---|---|
| llama-nemotron-embed-1b-v2 (embedding dim 2048) | 68.60% | 
| llama-nemotron-embed-1b-v2 (embedding dim 384) | 64.48% | 
| llama-3.2-nv-embedqa-1b-v1 (embedding dim 2048) | 68.97% | 
| nv-embedqa-mistral-7b-v2 | 72.97% | 
| nv-embedqa-mistral-7B-v1 | 64.93% | 
| nv-embedqa-e5-v5 | 62.07% | 
| nv-embedqa-e5-v4 | 57.65% | 
| e5-large-unsupervised | 48.03% | 
| BM25 | 44.67% | 

We evaluated the multilingual capabilities on the academic benchmark MIRACL across 15 languages and translated the English and Spanish version of MIRACL into additional 11 languages. The reported scores are based on an internal version of MIRACL by selecting hard negatives for each query to reduce the corpus size.

| Open & Commercial Retrieval Models | Average Recall@5 on multilingual | 
|---|---|
| llama-nemotron-embed-1b-v2 (embedding dim 2048) | 60.75% | 
| llama-nemotron-embed-1b-v2 (embedding dim 384) | 58.62% | 
| llama-3.2-nv-embedqa-1b-v1 | 60.07% | 
| nv-embedqa-mistral-7b-v2 | 50.42% | 
| BM25 | 26.51% | 

We evaluated the cross-lingual capabilities on the academic benchmark MLQA based on 7 languages (Arabic, Chinese, English, German, Hindi, Spanish, Vietnamese). We consider only evaluation datasets when the query and documents are in different languages. We calculate the average Recall@5 across the 42 different language pairs.

| Open & Commercial Retrieval Models | Average Recall@5 on MLQA dataset with different languages | 
|---|---|
| llama-nemotron-embed-1b-v2 (embedding dim 2048) | 79.86% | 
| llama-nemotron-embed-1b-v2 (embedding dim 384) | 71.61% | 
| llama-3.2-nv-embedqa-1b-v1 (embedding dim 2048) | 78.77% | 
| nv-embedqa-mistral-7b-v2 | 68.38% | 
| BM25 | 13.01% | 

We evaluated the support of long documents on the academic benchmark Multilingual Long-Document Retrieval (MLDR) built on Wikipedia and mC4, covering 12 typologically diverse languages. The English version has a median length of 2399 tokens and 90th percentile of 7483 tokens using the llama 3.2 tokenizer. The MLDR dataset is based on synthetic generated questions with a LLM, which has the tendency to create questions with similar keywords than the positive document, but might not be representative for real user queries. This characteristic of the dataset benefits sparse embeddings like BM25.

| Open & Commercial Retrieval Models | Average Recall@5 on MLDR | 
|---|---|
| llama-nemotron-embed-1b-v2 (embedding dim 2048) | 59.55% | 
| llama-nemotron-embed-1b-v2 (embedding dim 384) | 54.77% | 
| llama-3.2-nv-embedqa-1b-v1 (embedding dim 2048) | 60.49% | 
| nv-embedqa-mistral-7b-v2 | 43.24% | 
| BM25 | 71.39% | 

**Data Collection Method by dataset**: Unknown

**Labeling Method by dataset:** Unknown

**Properties:** The evaluation datasets are based on MTEB/BEIR, TextQA, TechQA, MIRACL, MLQA, and MLDR. The size ranges between 10,000s up to 5M depending on the dataset.

**Inference**
**Engine:** TensorRT
**Test Hardware:** H100 PCIe/SXM, A100 PCIe/SXM, L40s, L4, and A10G

```
@article{moreira2024nv,
  title={NV-Retriever: Improving text embedding models with effective hard-negative mining},
  author={Moreira, Gabriel de Souza P and Osmulski, Radek and Xu, Mengyao and Ak, Ronay and Schifferer, Benedikt and Oldridge, Even},
  journal={arXiv preprint arXiv:2407.15831},
  year={2024}
}
```
NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their supporting model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

For more detailed information on ethical considerations for this model, please see the Model Card++ tab for the Explainability, Bias, Safety & Security, and Privacy subcards.

Please report security vulnerabilities or NVIDIA AI Concerns here.

Get access to knowledge base articles and support cases or submit a ticket at the NVIDIA AI Enterprise Support Services page..

Visit the NeMo Retriever docs page for release documentation, deployment guides and more.

| Field | Response | 
|---|---|
| Participation considerations from adversely impacted groups protected classes in model design and testing | None | 
| Measures taken to mitigate against unwanted bias | None | 

