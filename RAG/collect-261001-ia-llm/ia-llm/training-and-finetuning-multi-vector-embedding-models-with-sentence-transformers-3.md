---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers-3
title: "Loading in fp32 is preferred for training if your memory can handle it"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["training", "distillation", "gpu", "inference", "parameters"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [161, 247]
sha256: 8dd105c38d8dce0c03c1b260785a5df6b1d8dcc35825f8c603bf3a5e30cd72c7
---

# Loading in fp32 is preferred for training if your memory can handle it

For distillation from a stronger teacher, which is how the strongest general-purpose late-interaction models are trained, see `MultiVectorDistillKLDivLoss` and the Knowledge Distillation tab in the Training Overview documentation.

You can customize the training process using the `MultiVectorEncoderTrainingArguments` class. This class lets you adjust parameters that can impact training speed and help you understand what's happening during training.

For more information on the most useful training arguments, check out the Multi-Vector Encoder > Training Overview > Training Arguments. It's worth reading to get the most out of your training.

Here's an example, using the values from my actual training run:

```
from sentence_transformers import MultiVectorEncoderTrainingArguments
from sentence_transformers.base.sampler import BatchSamplers
args = MultiVectorEncoderTrainingArguments(
    # Required parameter:
    output_dir="models/mLateOn-medical",
    # Optional training parameters:
    num_train_epochs=1,
    per_device_train_batch_size=128,  # the effective contrastive batch, thanks to GradCache
    per_device_eval_batch_size=16,
    learning_rate=1e-4,
    warmup_steps=0.05,
    prompts={"question": "[Q] ", "passage_text": "[D] "},  # the checkpoint's markers, keyed by training column
    fp16=False,  # Set to True if you have a GPU that supports FP16
    bf16=True,  # Set to True if you have a GPU that supports BF16
    batch_sampler=BatchSamplers.NO_DUPLICATES,  # in-batch negatives benefit from no duplicates
    # Optional tracking/debugging parameters:
    eval_strategy="steps",
    eval_steps=0.1,
    save_strategy="steps",
    save_steps=0.05,
    logging_steps=0.01,
    run_name="mLateOn-medical",  # Will be used in e.g. Trackio, W&B, etc.
)
```
A few of these deserve a comment:

- `prompts` : training does not automatically apply the prompts stored in the model, so map them onto your training columns explicitly. Here that is the checkpoint's`[Q]` marker for the question column and`[D]` for the passage column, keeping training consistent with inference.
- `max_length` (deliberately not set): this argument caps tokenization during*training only* , for when you want cheaper training than the model's full serving length. I measured what that shortcut costs on this data. Training at 512 tokens lost about 0.015 NDCG@10 for about 2x the speed, and the deficit did not shrink with more data, because the model simply never sees what got cut off. Leave it unset so training matches inference, unless you need the speedup more than the quality.
- `learning_rate=1e-4` : after a sweep from 5e-6 to 2e-4, I had the best luck with this higher-than-usual learning rate.

To track your model's performance during training, you can pass an `eval_dataset` to the trainer for evaluation loss, but concrete retrieval metrics are much more informative. Sentence Transformers includes the following built-in evaluators for multi-vector models:

| Evaluator | Required Data | 
|---|---|
| `MultiVectorInformationRetrievalEvaluator` | Queries, corpus, and relevant document mappings | 
| `MultiVectorNanoBEIREvaluator` | No data required | 
| `MultiVectorTripletEvaluator` | (anchor, positive, negative) triplets | 
| `MultiVectorRerankingEvaluator` | List of `{'query': '...', 'positive': [...], 'negative': [...]}` dictionaries | 
| `MultiVectorDistillationEvaluator` | Queries with candidate documents and teacher scores | 

For domain finetuning, the `MultiVectorInformationRetrievalEvaluator` built from your own held-out data is the one that matters. One tip on constructing it is that the corpus should be hard enough that models can be told apart. In my case the MIRIAD questions are generated from their own source passages, which makes retrieval unusually easy. Against just the 10k gold passages, nearly every model scored above 0.97 NDCG@10. If your evaluation saturates like that, add *distractor* passages (I use deduplicated passages from the training split) until the scores spread out:

```
from datasets import load_dataset
from sentence_transformers.multi_vector_encoder.evaluation import MultiVectorInformationRetrievalEvaluator
dataset = load_dataset("tomaarsen/miriad-4.4M-split")
# Gold: 1,000 evaluation questions, each mapping to its own passage, with the
# eval split's full ~10k unique passages as the initial corpus
corpus = {}
queries = {}
relevant_docs = {}
passage_to_id = {}
for idx, row in enumerate(dataset["eval"]):
    if row["passage_text"] not in passage_to_id:
        passage_to_id[row["passage_text"]] = f"p{len(passage_to_id)}"
        corpus[passage_to_id[row["passage_text"]]] = row["passage_text"]
    if idx < 1_000:
        queries[f"q{idx}"] = row["question"]
        relevant_docs[f"q{idx}"] = {passage_to_id[row["passage_text"]]}
# Distractors: unique train passages that make the haystack realistic
seen = set(passage_to_id)
for row in dataset["train"]:
    if len(corpus) >= 200_000:
        break
    if row["passage_text"] not in seen:
        seen.add(row["passage_text"])
        corpus[f"d{len(corpus)}"] = row["passage_text"]
evaluator = MultiVectorInformationRetrievalEvaluator(
    queries=queries,
    corpus=corpus,
    relevant_docs=relevant_docs,
    name="miriad-dev",
    batch_size=16,
)
# results = evaluator(model)
```
The `MultiVectorEncoderTrainer` is where all previous components come together. Here is the complete script that trained multi-vector-encoder/mLateOn-medical, the model from the introduction:

