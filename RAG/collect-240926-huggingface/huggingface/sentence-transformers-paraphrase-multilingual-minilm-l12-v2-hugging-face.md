---
id: collect-240926-huggingface/huggingface/sentence-transformers-paraphrase-multilingual-minilm-l12-v2-hugging-face
title: "Mean Pooling - Take attention mask into account for correct averaging"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["attention", "compute", "embedding", "embeddings"]
source: docs/RAG/clean_en/huggingface/sentence-transformers-paraphrase-multilingual-minilm-l12-v2-hugging-face.md
source_anchor: ""
source_lines: [1, 66]
sha256: 4b1209f070cb7beb15c6b048e72c59a59b0f8cafabfcf0e64bf39acc7c4e02f3
---

# Mean Pooling - Take attention mask into account for correct averaging

<!-- source: https://huggingface.co/sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2 -->

This is a sentence-transformers model: It maps sentences & paragraphs to a 384 dimensional dense vector space and can be used for tasks like clustering or semantic search.

Using this model becomes easy when you have sentence-transformers installed:

```
pip install -U sentence-transformers
```
Then you can use the model like this:

```
from sentence_transformers import SentenceTransformer
sentences = ["This is an example sentence", "Each sentence is converted"]
model = SentenceTransformer('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')
embeddings = model.encode(sentences)
print(embeddings)
```
Without sentence-transformers, you can use the model like this: First, you pass your input through the transformer model, then you have to apply the right pooling-operation on-top of the contextualized word embeddings.

```
from transformers import AutoTokenizer, AutoModel
import torch
# Mean Pooling - Take attention mask into account for correct averaging
def mean_pooling(model_output, attention_mask):
    token_embeddings = model_output[0] #First element of model_output contains all token embeddings
    input_mask_expanded = attention_mask.unsqueeze(-1).expand(token_embeddings.size()).float()
    return torch.sum(token_embeddings * input_mask_expanded, 1) / torch.clamp(input_mask_expanded.sum(1), min=1e-9)
# Sentences we want sentence embeddings for
sentences = ['This is an example sentence', 'Each sentence is converted']
# Load model from HuggingFace Hub
tokenizer = AutoTokenizer.from_pretrained('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')
model = AutoModel.from_pretrained('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')
# Tokenize sentences
encoded_input = tokenizer(sentences, padding=True, truncation=True, return_tensors='pt')
# Compute token embeddings
with torch.no_grad():
    model_output = model(**encoded_input)
# Perform pooling. In this case, max pooling.
sentence_embeddings = mean_pooling(model_output, encoded_input['attention_mask'])
print("Sentence embeddings:")
print(sentence_embeddings)
```
```
SentenceTransformer(
  (0): Transformer({'max_seq_length': 128, 'do_lower_case': False}) with Transformer model: BertModel 
  (1): Pooling({'word_embedding_dimension': 384, 'pooling_mode_cls_token': False, 'pooling_mode_mean_tokens': True, 'pooling_mode_max_tokens': False, 'pooling_mode_mean_sqrt_len_tokens': False})
)
```
This model was trained by sentence-transformers.

If you find this model helpful, feel free to cite our publication Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks:

```
@inproceedings{reimers-2019-sentence-bert,
    title = "Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks",
    author = "Reimers, Nils and Gurevych, Iryna",
    booktitle = "Proceedings of the 2019 Conference on Empirical Methods in Natural Language Processing",
    month = "11",
    year = "2019",
    publisher = "Association for Computational Linguistics",
    url = "http://arxiv.org/abs/1908.10084",
}
```
- Downloads last month
- 45,991,751
