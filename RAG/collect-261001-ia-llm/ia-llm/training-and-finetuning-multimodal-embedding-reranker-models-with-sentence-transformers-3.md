---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers-3
title: "['text', 'image', 'video', 'message']"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face"]
dates: []
keywords: ["apache", "attention", "embedding", "gpu", "gpus", "license", "memory", "parameters", "qwen", "training"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [168, 309]
sha256: 71cd8e53a0acfa8c5eb649c6fd57a09799808c4f0d2e0b2e90589e5496ff3b67
---

# ['text', 'image', 'video', 'message']

```
from sentence_transformers.sentence_transformer.evaluation import InformationRetrievalEvaluator
# Build the evaluation data from the eval dataset.
# Queries and corpus use integer IDs: query 0's relevant document is corpus 0.
eval_queries = {qid: sample["query"] for qid, sample in enumerate(eval_dataset)}
eval_corpus = {did: sample["image"] for did, sample in enumerate(eval_dataset)}
num_eval = len(eval_dataset)
# Add hard negatives to the corpus with offset IDs (num_eval, 2*num_eval, ...)
# so they don't collide with the positive document IDs (0..num_eval-1).
negative_columns = ["negative_0", "negative_1", "negative_2", "negative_3"]
for neg_idx, neg_col in enumerate(negative_columns):
    for did, sample in enumerate(eval_dataset):
        eval_corpus[num_eval * (neg_idx + 1) + did] = sample[neg_col]
# Each query's relevant document is the positive at the same index
eval_relevant_docs = {idx: [idx] for idx in range(len(eval_dataset))}
eval_evaluator = InformationRetrievalEvaluator(
    queries=eval_queries,
    corpus=eval_corpus,
    relevant_docs=eval_relevant_docs,
    batch_size=1,
    show_progress_bar=True,
    name="vdr-eval-hard",
)
```
The evaluator takes text queries, a corpus of images (including hard negatives), and a mapping of which documents are relevant to which queries. Note that the corpus contains a mix of positive and hard negative document screenshots, making this a challenging evaluation. Using `batch_size=1` prevents out-of-memory issues during evaluation of the large VLM.

The `SentenceTransformerTrainer` brings everything together. Here's the complete training script:

```
from datasets import load_dataset
from sentence_transformers import SentenceTransformer
from sentence_transformers.sentence_transformer.evaluation import InformationRetrievalEvaluator
from sentence_transformers.sentence_transformer.losses import CachedMultipleNegativesRankingLoss, MatryoshkaLoss
from sentence_transformers.sentence_transformer.model_card import SentenceTransformerModelCardData
from sentence_transformers.sentence_transformer.trainer import SentenceTransformerTrainer
from sentence_transformers.sentence_transformer.training_args import (
    BatchSamplers,
    SentenceTransformerTrainingArguments,
)
# 1. Load a model to finetune with (optional) model card data
model = SentenceTransformer(
    "Qwen/Qwen3-VL-Embedding-2B",
    model_card_data=SentenceTransformerModelCardData(
        language="en",
        license="apache-2.0",
        model_name="Qwen3-VL-Embedding-2B model trained on Visual Document Retrieval query-document screenshot pairs",
    ),
    model_kwargs={"attn_implementation": "flash_attention_2", "torch_dtype": "bfloat16"},
    # Control image resolution: lower values save memory, higher values preserve detail
    processor_kwargs={"min_pixels": 28 * 28, "max_pixels": 600 * 600},
)
# 2. Load a dataset to finetune on: (query, positive, negative_0) triplets for training,
# all 4 hard negatives retained for evaluation
train_dataset = load_dataset("tomaarsen/llamaindex-vdr-en-train-preprocessed", "train", split="train")
train_dataset = train_dataset.select_columns(["query", "image", "negative_0"])
eval_dataset = load_dataset("tomaarsen/llamaindex-vdr-en-train-preprocessed", "eval", split="train")
# 3. Define a loss function
loss = CachedMultipleNegativesRankingLoss(model, mini_batch_size=1)
loss = MatryoshkaLoss(model, loss, matryoshka_dims=[2048, 1536, 1024, 512, 256, 128, 64])
# 4. (Optional) Specify training arguments
run_name = "Qwen3-VL-Embedding-2B-vdr"
args = SentenceTransformerTrainingArguments(
    # Required parameter:
    output_dir=f"models/{run_name}",
    # Optional training parameters:
    num_train_epochs=1,
    per_device_train_batch_size=64,
    per_device_eval_batch_size=64,
    learning_rate=2e-5,
    warmup_ratio=0.1,
    fp16=False,  # BF16 is preferred over FP16 for VLMs due to better numerical stability
    bf16=True,  # Set to True if your GPU supports BF16 (most modern GPUs do)
    batch_sampler=BatchSamplers.NO_DUPLICATES,  # MultipleNegativesRankingLoss benefits from no duplicates
    # Optional tracking/debugging parameters:
    eval_strategy="steps",
    eval_steps=0.1,
    save_strategy="steps",
    save_steps=0.1,
    save_total_limit=2,
    logging_steps=0.05,
    run_name=run_name,  # Used in e.g. Trackio if installed
    # report_to=["codecarbon", "trackio"],  # Uncomment to enable logging (pip install codecarbon trackio)
)
# 5. (Optional) Create an evaluator & evaluate the base model
eval_queries = {qid: sample["query"] for qid, sample in enumerate(eval_dataset)}
eval_corpus = {did: sample["image"] for did, sample in enumerate(eval_dataset)}
num_eval = len(eval_dataset)
negative_columns = ["negative_0", "negative_1", "negative_2", "negative_3"]
for neg_idx, neg_col in enumerate(negative_columns):
    for did, sample in enumerate(eval_dataset):
        eval_corpus[num_eval * (neg_idx + 1) + did] = sample[neg_col]
eval_relevant_docs = {idx: [idx] for idx in range(len(eval_dataset))}
eval_evaluator = InformationRetrievalEvaluator(
    queries=eval_queries,
    corpus=eval_corpus,
    relevant_docs=eval_relevant_docs,
    batch_size=1,
    show_progress_bar=True,
    name="vdr-eval-hard",
)
eval_evaluator(model)
# 6. Create a trainer & train
trainer = SentenceTransformerTrainer(
    model=model,
    args=args,
    train_dataset=train_dataset,
    eval_dataset=eval_dataset,
    loss=loss,
    evaluator=eval_evaluator,
)
trainer.train()
# 7. (Optional) Evaluate at each Matryoshka dimension
eval_evaluator(model)
for dim in [2048, 1536, 1024, 512, 256, 128, 64]:
    dim_evaluator = InformationRetrievalEvaluator(
        queries=eval_queries,
        corpus=eval_corpus,
        relevant_docs=eval_relevant_docs,
        truncate_dim=dim,
        batch_size=1,
        show_progress_bar=True,
        name=f"vdr-eval-hard-{dim}d",
    )
    dim_evaluator(model)
# 8. Save the trained model
model.save_pretrained(f"models/{run_name}/final")
# 9. (Optional) Push it to the Hugging Face Hub
# This pushes to your personal namespace, e.g. {your_username}/Qwen3-VL-Embedding-2B-vdr
model.push_to_hub("Qwen3-VL-Embedding-2B-vdr")
```
The training script is nearly identical to a text-only training script. The only differences are:

1. Model loading: We pass `model_kwargs` for precision and attention implementation, and`processor_kwargs` for image resolution bounds.
2. Loss function: We use `CachedMultipleNegativesRankingLoss` with`mini_batch_size=1` to handle the large VLM without running out of memory.
3. Evaluator: The evaluator uses images in the corpus and text as queries, enabling cross-modal retrieval evaluation.

Everything else (the trainer, training arguments, dataset loading) works exactly the same as text-only training.

After training for just 1 epoch, the finetuned tomaarsen/Qwen3-VL-Embedding-2B-vdr model achieves an NDCG@10 of **0.947** on the evaluation set (300 queries, 1500 corpus documents, cosine similarity). This is a significant improvement over the base Qwen/Qwen3-VL-Embedding-2B model's 0.888, and outperforms all existing VDR models:

## Full NDCG@10 numbers by model (20 models)

