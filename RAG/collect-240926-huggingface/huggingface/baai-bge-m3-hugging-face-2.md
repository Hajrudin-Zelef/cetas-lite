---
id: collect-240926-huggingface/huggingface/baai-bge-m3-hugging-face-2
title: "[[0.6265, 0.3477], [0.3499, 0.678 ]]"
domain: huggingface
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["benchmark", "compute", "distillation", "embedding", "embeddings", "fine-tuning", "latency", "leaderboard", "training"]
source: docs/RAG/clean_en/huggingface/baai-bge-m3-hugging-face.md
source_anchor: ""
source_lines: [117, 203]
sha256: cb7cb4084fb7ebb5803d84b2fc74139936e0f6bd2a962c1236a88712c31489ab
---

# [[0.6265, 0.3477], [0.3499, 0.678 ]]

```
from FlagEmbedding import BGEM3FlagModel
model = BGEM3FlagModel('BAAI/bge-m3',  use_fp16=True) 
sentences_1 = ["What is BGE M3?", "Defination of BM25"]
sentences_2 = ["BGE M3 is an embedding model supporting dense retrieval, lexical matching and multi-vector interaction.", 
               "BM25 is a bag-of-words retrieval function that ranks a set of documents based on the query terms appearing in each document"]
output_1 = model.encode(sentences_1, return_dense=True, return_sparse=True, return_colbert_vecs=True)
output_2 = model.encode(sentences_2, return_dense=True, return_sparse=True, return_colbert_vecs=True)
print(model.colbert_score(output_1['colbert_vecs'][0], output_2['colbert_vecs'][0]))
print(model.colbert_score(output_1['colbert_vecs'][0], output_2['colbert_vecs'][1]))
# 0.7797
# 0.4620
```
Input a list of text pairs, you can get the scores computed by different methods.

```
from FlagEmbedding import BGEM3FlagModel
model = BGEM3FlagModel('BAAI/bge-m3',  use_fp16=True) 
sentences_1 = ["What is BGE M3?", "Defination of BM25"]
sentences_2 = ["BGE M3 is an embedding model supporting dense retrieval, lexical matching and multi-vector interaction.", 
               "BM25 is a bag-of-words retrieval function that ranks a set of documents based on the query terms appearing in each document"]
sentence_pairs = [[i,j] for i in sentences_1 for j in sentences_2]
print(model.compute_score(sentence_pairs, 
                          max_passage_length=128, # a smaller max length leads to a lower latency
                          weights_for_different_modes=[0.4, 0.2, 0.4])) # weights_for_different_modes(w) is used to do weighted sum: w[0]*dense_score + w[1]*sparse_score + w[2]*colbert_score
# {
#   'colbert': [0.7796499729156494, 0.4621465802192688, 0.4523794651031494, 0.7898575067520142], 
#   'sparse': [0.195556640625, 0.00879669189453125, 0.0, 0.1802978515625], 
#   'dense': [0.6259765625, 0.347412109375, 0.349853515625, 0.67822265625], 
#   'sparse+dense': [0.482503205537796, 0.23454029858112335, 0.2332356721162796, 0.5122477412223816], 
#   'colbert+sparse+dense': [0.6013619303703308, 0.3255828022956848, 0.32089319825172424, 0.6232916116714478]
# }
```
We provide the evaluation script for MKQA and MLDR

The BGE-M3 model emerged as the top performer on this benchmark (OAI is short for OpenAI). For more details, please refer to the article and Github Repo

- Multilingual (Miracl dataset)

- Cross-lingual (MKQA dataset)

- Long Document Retrieval 
  - MLDR: 
 Please note that MLDR is a document retrieval dataset we constructed via LLM, covering 13 languages, including test set, validation set, and training set. We utilized the training set from MLDR to enhance the model's long document retrieval capabilities. Therefore, comparing baselines with`Dense w.o.long` (fine-tuning without long document dataset) is more equitable. 
Additionally, this long document retrieval dataset will be open-sourced to address the current lack of open-source multilingual long text retrieval datasets.
We believe that this data will be helpful for the open-source community in training document retrieval models.
  - NarritiveQA:
- Comparison with BM25

We utilized Pyserini to implement BM25, and the test results can be reproduced by this script. We tested BM25 using two different tokenizers: one using Lucene Analyzer and the other using the same tokenizer as M3 (i.e., the tokenizer of xlm-roberta). The results indicate that BM25 remains a competitive baseline, especially in long document retrieval.

- Self-knowledge Distillation: combining multiple outputs from different retrieval modes as reward signal to enhance the performance of single mode(especially for sparse retrieval and multi-vec(colbert) retrival)
- Efficient Batching: Improve the efficiency when fine-tuning on long text. The small-batch strategy is simple but effective, which also can used to fine-tune large embedding model.
- MCLS: A simple method to improve the performance on long text without fine-tuning. If you have no enough resource to fine-tuning model with long text, the method is useful.

Refer to our report for more details.

Thanks to the authors of open-sourced datasets, including Miracl, MKQA, NarritiveQA, etc. Thanks to the open-sourced libraries like Tevatron, Pyserini.

If you find this repository useful, please consider giving a star :star: and citation

```
@misc{bge-m3,
      title={BGE M3-Embedding: Multi-Lingual, Multi-Functionality, Multi-Granularity Text Embeddings Through Self-Knowledge Distillation}, 
      author={Jianlv Chen and Shitao Xiao and Peitian Zhang and Kun Luo and Defu Lian and Zheng Liu},
      year={2024},
      eprint={2402.03216},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```
- Downloads last month
- 37,173,706

## Spaces using BAAI/bge-m3 100

## Collection including BAAI/bge-m3

## Papers for BAAI/bge-m3

- mteb/BRIGHT leaderboard
- BrightAopsRetrieval Default Standard View evaluation resultssourceObtained using MTEB v2.10.124.56<sup>*</sup>
- BrightAopsRetrieval View evaluation resultssourceObtained using MTEB v2.10.124.56<sup>*</sup>
- BrightBiologyLongRetrieval Default Long View evaluation resultssourceObtained using MTEB v2.10.1214.73<sup>*</sup>
- mteb/arguana leaderboard
- ArguAna Default Test View evaluation resultssourceObtained using MTEB v1.12.7554.04<sup>*</sup>
- ArguAna View evaluation resultssourceObtained using MTEB v1.12.7554.04<sup>*</sup>
