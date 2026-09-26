---
id: collect-240926-huggingface/huggingface/baai-bge-m3-hugging-face-1
title: "[[0.6265, 0.3477], [0.3499, 0.678 ]]"
domain: huggingface
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["benchmark", "compute", "cost", "embedding", "embeddings", "fine-tuning", "reranker"]
source: docs/RAG/clean_en/huggingface/baai-bge-m3-hugging-face.md
source_anchor: ""
source_lines: [1, 116]
sha256: fa4a2cc1a2799767a4d1d8d87778558133b91d95724806a38244cc98fdefd836
---

# [[0.6265, 0.3477], [0.3499, 0.678 ]]

<!-- source: https://huggingface.co/BAAI/bge-m3 -->

For more details please refer to our github repo: https://github.com/FlagOpen/FlagEmbedding

In this project, we introduce BGE-M3, which is distinguished for its versatility in Multi-Functionality, Multi-Linguality, and Multi-Granularity.

- Multi-Functionality: It can simultaneously perform the three common retrieval functionalities of embedding model: dense retrieval, multi-vector retrieval, and sparse retrieval.
- Multi-Linguality: It can support more than 100 working languages.
- Multi-Granularity: It is able to process inputs of different granularities, spanning from short sentences to long documents of up to 8192 tokens.

**Some suggestions for retrieval pipeline in RAG**

We recommend to use the following pipeline: hybrid retrieval + re-ranking.

- Hybrid retrieval leverages the strengths of various methods, offering higher accuracy and stronger generalization capabilities. A classic example: using both embedding retrieval and the BM25 algorithm. Now, you can try to use BGE-M3, which supports both embedding and sparse retrieval. This allows you to obtain token weights (similar to the BM25) without any additional cost when generate dense embeddings. To use hybrid retrieval, you can refer to Vespa and Milvus.
- As cross-encoder models, re-ranker demonstrates higher accuracy than bi-encoder embedding model. Utilizing the re-ranking model (e.g., bge-reranker, bge-reranker-v2) after retrieval can further filter the selected text.

- 2024/7/1: **We update the MIRACL evaluation results of BGE-M3** . To reproduce the new results, you can refer to: bge-m3_miracl_2cr. We have also updated our paper on arXiv.## DetailsThe previous test results were lower because we mistakenly removed the passages that have the same id as the query from the search results. After correcting this mistake, the overall performance of BGE-M3 on MIRACL is higher than the previous results, but the experimental conclusion remains unchanged. The other results are not affected by this mistake. To reproduce the previous lower results, you need to add the `--remove-query` parameter when using`pyserini.search.faiss` or`pyserini.search.lucene` to search the passages.
- 2024/3/20: **Thanks Milvus team!** Now you can use hybrid retrieval of bge-m3 in Milvus: pymilvus/examples
/hello_hybrid_sparse_dense.py.
- 2024/3/8: **Thanks for the experimental results from @Yannael. In this benchmark, BGE-M3 achieves top performance in both English and other languages, surpassing models such as OpenAI.**
- 2024/3/2: Release unified fine-tuning example and data
- 2024/2/6: We release the MLDR (a long document retrieval dataset covering 13 languages) and evaluation pipeline.
- 2024/2/1: **Thanks for the excellent tool from Vespa.** You can easily use multiple modes of BGE-M3 following this notebook

- Model

| Model Name | Dimension | Sequence Length | Introduction | 
|---|---|---|---|
| BAAI/bge-m3 | 1024 | 8192 | multilingual; unified fine-tuning (dense, sparse, and colbert) from bge-m3-unsupervised | 
| BAAI/bge-m3-unsupervised | 1024 | 8192 | multilingual; contrastive learning from bge-m3-retromae | 
| BAAI/bge-m3-retromae | -- | 8192 | multilingual; extend the max_length of xlm-roberta to 8192 and further pretrained via retromae | 
| BAAI/bge-large-en-v1.5 | 1024 | 512 | English model | 
| BAAI/bge-base-en-v1.5 | 768 | 512 | English model | 
| BAAI/bge-small-en-v1.5 | 384 | 512 | English model | 

- Data

| Dataset | Introduction | 
|---|---|
| MLDR | Docuemtn Retrieval Dataset, covering 13 languages | 
| bge-m3-data | Fine-tuning data used by bge-m3 | 

**1. Introduction for different retrieval methods**

- Dense retrieval: map the text into a single embedding, e.g., DPR, BGE-v1.5
- Sparse retrieval (lexical matching): a vector of size equal to the vocabulary, with the majority of positions set to zero, calculating a weight only for tokens present in the text. e.g., BM25, unicoil, and splade
- Multi-vector retrieval: use multiple vectors to represent a text, e.g., ColBERT.

**2. How to use BGE-M3 in other projects?**

For embedding retrieval, you can employ the BGE-M3 model using the same approach as BGE. The only difference is that the BGE-M3 model no longer requires adding instructions to the queries.

For hybrid retrieval, you can use Vespa and Milvus.

**3. How to fine-tune bge-M3 model?**

You can follow the common in this example to fine-tune the dense embedding.

If you want to fine-tune all embedding function of m3 (dense, sparse and colbert), you can refer to the unified_fine-tuning example

Install:

```
git clone https://github.com/FlagOpen/FlagEmbedding.git
cd FlagEmbedding
pip install -e .
```
or:

```
pip install -U FlagEmbedding
```
- Dense Embedding

```
from FlagEmbedding import BGEM3FlagModel
model = BGEM3FlagModel('BAAI/bge-m3',  
                       use_fp16=True) # Setting use_fp16 to True speeds up computation with a slight performance degradation
sentences_1 = ["What is BGE M3?", "Defination of BM25"]
sentences_2 = ["BGE M3 is an embedding model supporting dense retrieval, lexical matching and multi-vector interaction.", 
               "BM25 is a bag-of-words retrieval function that ranks a set of documents based on the query terms appearing in each document"]
embeddings_1 = model.encode(sentences_1, 
                            batch_size=12, 
                            max_length=8192, # If you don't need such a long length, you can set a smaller value to speed up the encoding process.
                            )['dense_vecs']
embeddings_2 = model.encode(sentences_2)['dense_vecs']
similarity = embeddings_1 @ embeddings_2.T
print(similarity)
# [[0.6265, 0.3477], [0.3499, 0.678 ]]
```
You also can use sentence-transformers and huggingface transformers to generate dense embeddings. Refer to baai_general_embedding for details.

- Sparse Embedding (Lexical Weight)

```
from FlagEmbedding import BGEM3FlagModel
model = BGEM3FlagModel('BAAI/bge-m3',  use_fp16=True) # Setting use_fp16 to True speeds up computation with a slight performance degradation
sentences_1 = ["What is BGE M3?", "Defination of BM25"]
sentences_2 = ["BGE M3 is an embedding model supporting dense retrieval, lexical matching and multi-vector interaction.", 
               "BM25 is a bag-of-words retrieval function that ranks a set of documents based on the query terms appearing in each document"]
output_1 = model.encode(sentences_1, return_dense=True, return_sparse=True, return_colbert_vecs=False)
output_2 = model.encode(sentences_2, return_dense=True, return_sparse=True, return_colbert_vecs=False)
# you can see the weight for each token:
print(model.convert_id_to_token(output_1['lexical_weights']))
# [{'What': 0.08356, 'is': 0.0814, 'B': 0.1296, 'GE': 0.252, 'M': 0.1702, '3': 0.2695, '?': 0.04092}, 
#  {'De': 0.05005, 'fin': 0.1368, 'ation': 0.04498, 'of': 0.0633, 'BM': 0.2515, '25': 0.3335}]
# compute the scores via lexical mathcing
lexical_scores = model.compute_lexical_matching_score(output_1['lexical_weights'][0], output_2['lexical_weights'][0])
print(lexical_scores)
# 0.19554901123046875
print(model.compute_lexical_matching_score(output_1['lexical_weights'][0], output_1['lexical_weights'][1]))
# 0.0
```
- Multi-Vector (ColBERT)

