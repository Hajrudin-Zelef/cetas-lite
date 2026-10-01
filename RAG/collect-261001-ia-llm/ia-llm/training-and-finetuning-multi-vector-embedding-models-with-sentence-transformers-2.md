---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers-2
title: "Loading in fp32 is preferred for training if your memory can handle it"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["memory", "training", "cost", "distillation", "embedding", "gpu", "gpus", "pretraining"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [77, 160]
sha256: 474e0b859a1f20df9417c0be1e0b9c4d348e3e90a1fa9861694bcf50f1ef4157
---

# Loading in fp32 is preferred for training if your memory can handle it

The classic ColBERT tokenization tricks (`[MASK]` query expansion, `[Q]` / `[D]` prefix tokens, a document length cap, a punctuation skiplist) are all off by default and configurable. See Creating Custom Models for the full set. For what it's worth, I tested `[MASK]` query expansion in four configurations for my domain finetune and none of them made a measurable difference, so don't feel obliged to reach for the classic recipe.

I measured this directly while preparing this blogpost, taking six starting points and training each with the identical recipe on 25k medical question-passage pairs from MIRIAD, then evaluating on 1,000 held-out questions against a 50,000 passage corpus:

| Starting point | Zero-shot NDCG@10 | After 25k pairs | Delta | 
|---|---|---|---|
| lightonai/mLateOn-unsupervised | 0.9087 | **0.9398** | **+0.0311** | 
| lightonai/mLateOn | 0.9277 | 0.9319 | +0.0042 | 
| lightonai/LateOn-unsupervised | 0.9026 | **0.9206** | **+0.0180** | 
| lightonai/LateOn | 0.9185 | 0.9105 | -0.0080 | 
| lightonai/GTE-ModernColBERT-v1 | 0.9198 | 0.9007 | -0.0191 | 
| Fresh head on gte-modernbert-base | - | 0.9177 | - | 

The result surprised me, and it replicated across two model families. *The `-unsupervised` checkpoints adapt to a new domain far better than their finished siblings, overtaking them despite starting lower. These checkpoints sit after large-scale contrastive pretraining but before supervised finetuning on general retrieval, so they carry all the late-interaction structure with none of the general-purpose tuning that domain training then has to undo. The finished checkpoints, by contrast, barely moved or even regressed, at every learning rate I tried.

So, if the model family you like publishes a pre-supervised checkpoint, start there. If not, a fresh projection on a strong retrieval-pretrained backbone is a close runner-up. Continuing from a fully finished checkpoint is the weakest option for domain adaptation, despite being the most natural-feeling one.

The `MultiVectorEncoderTrainer` uses `datasets.Dataset` or `datasets.DatasetDict` instances for training and evaluation. You can load data from the Hugging Face Datasets Hub or use local data in whatever format you prefer (e.g. CSV, JSON, Parquet, Arrow, or SQL).

**Note:** Lots of public datasets that work out of the box with Sentence Transformers have been tagged with `sentence-transformers` on the Hugging Face Hub, so you can easily find them on https://huggingface.co/datasets?other=sentence-transformers. Consider browsing through these to find ready-to-go datasets that might be useful for your tasks, domains, or languages.

You can use the `load_dataset` function to load data from datasets on the Hub:

```
from datasets import load_dataset
train_dataset = load_dataset("tomaarsen/miriad-4.4M-split", split="train")
print(train_dataset)
"""
Dataset({
    features: ['question', 'passage_text'],
    num_rows: 4467542
})
"""
```
This is the dataset I'll train on in this blogpost: 4.4 million medical questions from MIRIAD, each paired with the source passage that contains its answer (averaging 941 tokens). Simple (query, relevant passage) pairs like these are the easiest retrieval training data to collect for your own domain, and as you'll see, they're all you need.

You can also use `load_dataset` for loading local data in common file formats:

```
from datasets import load_dataset
dataset = load_dataset("csv", data_files="my_file.csv")
# or
dataset = load_dataset("json", data_files="my_file.json")
```
And if your local data requires pre-processing, you can use `datasets.Dataset.from_dict` to initialize your dataset with a dictionary of lists:

```
from datasets import Dataset
queries = []
documents = []
# Open a file, perform preprocessing, filtering, cleaning, etc.
# and append to the lists
dataset = Dataset.from_dict({
    "query": queries,
    "document": documents,
})
```
It is important that your dataset format matches your loss function (or that you choose a loss function that matches your dataset format). Verifying whether a dataset format works with a loss function involves two steps:

1. If your loss function requires a *Label* according to the Loss Overview table, then your dataset must have a**column named "label" or "score"** . This column is automatically taken as the label.
2. All columns not named "label" or "score" are considered *Inputs* according to the Loss Overview table. The number of remaining columns must match the number of valid inputs for your chosen loss. The names of these columns are**irrelevant** , only the**order matters** .

There are two multi-vector specific conventions on top of this:

- Positional query and document assignment: the first column is embedded as the *query* and all following columns as*documents* , regardless of the column names. This default can be overridden per column via the standard`router_mapping` training argument.
- Knowledge distillation format: one column per candidate document, i.e. `(query, document_1, ..., document_N, scores)` where`scores` is a list of N teacher scores per row. For KD datasets that store query and document*IDs* alongside separate text datasets (e.g. lightonai/ms-marco-en-bge), you can use`resolve_ids` to resolve the IDs to texts on the fly.

Loss functions quantify how well a model performs for a given batch of data, allowing an optimizer to update the model weights to produce more favourable (i.e., lower) loss values. The right loss function for your task depends on the data you have and what you're trying to achieve. You can find a full list of options in the Loss Overview.

For the common case of question-answer or question-passage pairs, the workhorse is in-batch negatives training with `MultiVectorMultipleNegativesRankingLoss`, where every other document in the batch acts as a negative for each query. Bigger batches mean more negatives and stronger training, so in practice you'll want its GradCache variant, `CachedMultiVectorMultipleNegativesRankingLoss`, which decouples the effective batch size from what fits on your GPU:

```
from sentence_transformers import MultiVectorEncoder
from sentence_transformers.multi_vector_encoder.losses import CachedMultiVectorMultipleNegativesRankingLoss
model = MultiVectorEncoder("lightonai/mLateOn-unsupervised", model_kwargs={"torch_dtype": "float32"})
loss = CachedMultiVectorMultipleNegativesRankingLoss(
    model=model,
    mini_batch_size=16,  # how many documents to encode per chunk: bounds memory, not quality
)
```
The `mini_batch_size` parameter bounds the memory by encoding documents in chunks of this size, while the effective contrastive batch size (128 in my run below, and in my ablations bigger batches bought nothing further) stays a free choice. GradCache guarantees identical results regardless of the chunk size, so lower it for smaller GPUs at only a wall-clock cost. When your document lengths vary a lot, consider its sibling `mini_batch_num_tokens`, which packs each chunk to a total token budget instead of a document count, so a chunk of unusually long documents can never spike your memory (my `mini_batch_size=16` at roughly 940 tokens per document corresponds to `mini_batch_num_tokens=15_000`).

One multi-vector specific trap is that the contrastive losses default to `scale=1.0`, unlike the dense embedding equivalent which defaults to `scale=20.0`. That 20.0 exists because a cosine similarity is a single value in [-1, 1], too narrow a range for a sharp softmax. A MaxSim score instead sums one best-match similarity per query token, so it already spans roughly [0, query_length]: a 32-token query can score up to 32. So don't copy `scale=20.0` over from a dense training script, since it would saturate the softmax and kill your gradients.

