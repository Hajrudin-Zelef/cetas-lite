---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers-4
title: "Loading in fp32 is preferred for training if your memory can handle it"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face"]
dates: []
keywords: ["memory", "training", "apache", "embedding", "gpu", "inference", "license", "qwen"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [248, 376]
sha256: 5985a81b7460d59e3bf3a986870400e785363e1a89097412c2418b11cd34fb72
---

# Loading in fp32 is preferred for training if your memory can handle it

```
import logging
import string
import traceback
from datasets import load_dataset
from sentence_transformers import (
    MultiVectorEncoder,
    MultiVectorEncoderModelCardData,
    MultiVectorEncoderTrainer,
    MultiVectorEncoderTrainingArguments,
)
from sentence_transformers.base.sampler import BatchSamplers
from sentence_transformers.multi_vector_encoder.evaluation import MultiVectorInformationRetrievalEvaluator
from sentence_transformers.multi_vector_encoder.losses import CachedMultiVectorMultipleNegativesRankingLoss
logging.basicConfig(format="%(asctime)s - %(message)s", datefmt="%Y-%m-%d %H:%M:%S", level=logging.INFO)
def main():
    # 1. Load the starting checkpoint: contrastively pretrained, not yet supervised
    # Loading in fp32 is preferred for training if your memory can handle it
    model = MultiVectorEncoder(
        "lightonai/mLateOn-unsupervised",
        model_kwargs={"torch_dtype": "float32"},
        processor_kwargs={"model_max_length": 8192},
        model_card_data=MultiVectorEncoderModelCardData(
            language="en",
            license="apache-2.0",
            model_name="mLateOn finetuned on MIRIAD medical retrieval",
        ),
    )
    # 2. Lift the per-task length caps so training and inference see full medical passages
    model[0].query_length = None
    model[0].document_length = None
    # 3. Skip punctuation tokens during scoring: a small quality win and a 9.6% smaller index
    model[2].skiplist_words = list(string.punctuation)
    model[2].resolve_with_tokenizer(model.tokenizer)
    # 4. Load 1 million medical question-passage pairs
    train_dataset = load_dataset("tomaarsen/miriad-4.4M-split", split="train").select(range(1_000_000))
    # 5. In-batch negatives with GradCache: large effective batch, memory-bounded chunks
    loss = CachedMultiVectorMultipleNegativesRankingLoss(model=model, mini_batch_size=16)
    # 6. A light dev evaluator to watch progress during training: 500 held-out questions
    # against the eval split's ~10k unique passages. The full 200k protocol runs afterwards.
    eval_split = load_dataset("tomaarsen/miriad-4.4M-split", split="eval")
    corpus, queries, relevant_docs, passage_to_id = {}, {}, {}, {}
    for idx, row in enumerate(eval_split):
        if row["passage_text"] not in passage_to_id:
            passage_to_id[row["passage_text"]] = f"p{len(passage_to_id)}"
            corpus[passage_to_id[row["passage_text"]]] = row["passage_text"]
        if idx < 500:
            queries[f"q{idx}"] = row["question"]
            relevant_docs[f"q{idx}"] = {passage_to_id[row["passage_text"]]}
    dev_evaluator = MultiVectorInformationRetrievalEvaluator(
        queries=queries, corpus=corpus, relevant_docs=relevant_docs, name="miriad-dev", batch_size=16
    )
    # 7. Training arguments, as discussed above
    run_name = "mLateOn-medical"
    args = MultiVectorEncoderTrainingArguments(
        output_dir=f"models/{run_name}",
        num_train_epochs=1,
        per_device_train_batch_size=128,
        per_device_eval_batch_size=16,
        learning_rate=1e-4,
        warmup_steps=0.05,
        prompts={"question": "[Q] ", "passage_text": "[D] "},
        fp16=False,  # Set to True if you have a GPU that supports FP16
        bf16=True,  # Set to True if you have a GPU that supports BF16
        batch_sampler=BatchSamplers.NO_DUPLICATES,
        eval_strategy="steps",
        eval_steps=0.1,
        save_strategy="steps",
        save_steps=0.05,
        logging_steps=0.01,
        run_name=run_name,
    )
    # 8. Create a trainer & train
    trainer = MultiVectorEncoderTrainer(
        model=model,
        args=args,
        train_dataset=train_dataset,
        loss=loss,
        evaluator=dev_evaluator,
    )
    trainer.train()
    # 9. Save the trained model
    model.save_pretrained(f"models/{run_name}/final")
    # 10. (Optional) Push it to the Hugging Face Hub
    try:
        model.push_to_hub(run_name)
    except Exception:
        logging.error(f"Error uploading model to the Hugging Face Hub:\n{traceback.format_exc()}")
if __name__ == "__main__":
    main()
```
That's the whole recipe: a pre-supervised checkpoint, a million domain pairs, in-batch negatives, full document length, and a higher-than-usual learning rate. The run took 14.5 hours on my single RTX 3090 at a peak of 17.5 GB VRAM, and every one of those choices was the winner of a measured comparison rather than a guess.

For readers on smaller budgets, my scaling experiments put 100k pairs (75 minutes of training) within 0.012 NDCG@10 of the full million-pair run. Most of the gain comes in the first hour.

The MultiVectorEncoder trainer supports various `transformers.TrainerCallback` subclasses, including:

- `WandbCallback` for logging training metrics to W&B if`wandb` is installed
- `TensorBoardCallback` for logging training metrics to TensorBoard if`tensorboard` is accessible
- `CodeCarbonCallback` for tracking carbon emissions during training if`codecarbon` is installed

Enable these via the `report_to` training argument, e.g. `report_to=["wandb", "codecarbon"]`, with the required dependencies installed. It defaults to `"none"`, and `report_to="all"` activates every integration whose dependency is installed.

Refer to the Transformers Callbacks documentation for more information on these callbacks and how to create your own.

Typically, top-performing general-purpose models are trained on multiple datasets simultaneously. However, this approach can be challenging due to the varying formats of each dataset. Fortunately, the `MultiVectorEncoderTrainer` allows you to train on multiple datasets without requiring a uniform format. Additionally, it provides the flexibility to apply different loss functions to each dataset. Here are the steps to train with multiple datasets at once:

- Use a dictionary of `datasets.Dataset` instances (or a`datasets.DatasetDict` ) as the`train_dataset` (and optionally also`eval_dataset` ).
- (Optional) Use a dictionary of loss functions mapping dataset names to losses. Only required if you wish to use different loss functions for different datasets.

Each training/evaluation batch will only contain samples from one of the datasets. The order in which batches are sampled from the multiple datasets is defined by the `MultiDatasetBatchSamplers` enum, which can be passed to the `MultiVectorEncoderTrainingArguments` via `multi_dataset_batch_sampler`. Valid options are:

- `MultiDatasetBatchSamplers.ROUND_ROBIN` : Round-robin sampling from each dataset until one is exhausted. With this strategy, it's likely that not all samples from each dataset are used, but each dataset is sampled from equally.
- `MultiDatasetBatchSamplers.PROPORTIONAL` (default): Sample from each dataset in proportion to its size. With this strategy, all samples from each dataset are used and larger datasets are sampled from more frequently.

To find out where the finetuned model stands, I evaluated it against over 50 retrieval model configurations across four architecture families on the MIRIAD evaluation set, built exactly as in the Evaluator section above, with 1,000 held-out medical questions searching 200,000 unique passages (the 10k gold passages hidden among 190k deduplicated distractors from the training split). This corpus is four times the size of the 50,000-passage one from Which starting point should you pick?, so scores are not comparable between the two tables.

The headline results, with the full table in the collapsible below:

| Model | Family | NDCG@10 | 
|---|---|---|
| **multi-vector-encoder/mLateOn-medical (mine)** | **Multi-vector, finetuned** | **0.9139** | 
| lightonai/mLateOn | Multi-vector, zero-shot | 0.8520 | 
| lightonai/GTE-ModernColBERT-v1 (cap lifted) | Multi-vector, zero-shot | 0.8502 | 
| Qwen/Qwen3-Embedding-4B | Dense, zero-shot | 0.7817 | 
| voyageai/voyage-4-nano | Dense, zero-shot | 0.7563 | 
| BM25 | Lexical | 0.7501 | 
| naver/splade-v3 | Sparse, zero-shot | 0.6853 | 

