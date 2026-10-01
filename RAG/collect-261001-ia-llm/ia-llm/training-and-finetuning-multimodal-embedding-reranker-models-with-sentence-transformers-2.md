---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers-2
title: "['text', 'image', 'video', 'message']"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face"]
dates: []
keywords: ["cost", "embedding", "embeddings", "gpu", "memory", "multimodal", "parameters", "training"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [83, 167]
sha256: 9cbddb7a17220801503af0e22c165f9c76cdcacd45ce1fb8ab7b32e7846a2777
---

# ['text', 'image', 'video', 'message']

```
from datasets import load_dataset
train_dataset = load_dataset("tomaarsen/llamaindex-vdr-en-train-preprocessed", "train", split="train")
train_dataset = train_dataset.select_columns(["query", "image", "negative_0"])
eval_dataset = load_dataset("tomaarsen/llamaindex-vdr-en-train-preprocessed", "eval", split="train")
```
The `train` config contains the first 10,000 samples, and the `eval` config contains the next 300 samples (a `full` config with all 53,512 samples is also available). For training, I select `query`, `image`, and `negative_0` to form (anchor, positive, hard negative) triplets. Including additional hard negatives would likely improve the training signal, but each extra negative also increases memory usage and training time, so I stick with one. For evaluation, I keep all four hard negatives per query to build a more challenging retrieval corpus (more on that in the Evaluator section).

Just like text-only training, the dataset format must match your chosen loss function. The rules are the same:

1. If your loss function requires a *Label* , your dataset must have a column named**"label"** or**"score"** .
2. All columns other than **"label"** or**"score"** are considered*Inputs* . The number of these columns must match the number of valid inputs for your chosen loss function. Beyond the label column, the column names don't matter, only the order does.

For multimodal datasets, the inputs can contain:

- **Text** : strings.
- **Image** : PIL images, file paths, URLs, or numpy/torch arrays.
- **Audio** : file paths, numpy/torch arrays, dicts with`"array"` and`"sampling_rate"` keys, or (if`torchcodec` is installed)`torchcodec.AudioDecoder` instances.
- **Video** : file paths, numpy/torch arrays, dicts with`"array"` and`"video_metadata"` keys, or (if`torchcodec` is installed)`torchcodec.VideoDecoder` instances.
- **Multimodal dicts** : a dict mapping modality names to values, e.g.`{"text": ..., "image": ...}` . The keys must be`"text"` ,`"image"` ,`"audio"` , or`"video"` .

The data collator automatically calls `model.preprocess()`, which detects the modality of each input and applies the appropriate preprocessing. No manual tokenization or image processing is needed.

Many Hugging Face datasets that work out of the box with Sentence Transformers have been tagged with `sentence-transformers`, allowing you to easily find them at https://huggingface.co/datasets?other=sentence-transformers.


For this training, I use `CachedMultipleNegativesRankingLoss`, a common choice for retrieval tasks. It accepts (query, positive) pairs with any number of additional hard negative columns, from 0 up to n, as long as each sample has the same number of negatives.
During training, the loss pushes each query's similarity to its positive *up* and its similarity to every negative *down*. The negatives come from two sources:

1. **Hard negatives** : the negative column(s) explicitly supplied in the dataset (just`negative_0` in our triplet setup).
2. **In-batch negatives** : the positives and hard negatives from every*other* sample in the same batch, reused as additional negatives for this query at no extra cost.

More negatives per query means a stronger training signal, so a larger batch size directly improves training quality. Beyond that, the "cached" variant of the loss uses gradient caching to make large effective batch sizes feasible even when GPU memory is limited.

The `mini_batch_size` parameter controls how many samples are processed at once during the cached forward passes. For large multimodal models, setting this to a small value (e.g., 1) is important to avoid out-of-memory errors without sacrificing the benefits of large effective batch sizes:

```
from sentence_transformers.sentence_transformer.losses import CachedMultipleNegativesRankingLoss
loss = CachedMultipleNegativesRankingLoss(model, mini_batch_size=1)
```
To produce embeddings that work well at multiple dimensionalities, I wrap the base loss with `MatryoshkaLoss`. This trains the model so that truncating the embedding to a smaller number of dimensions still yields good performance:

```
from sentence_transformers.sentence_transformer.losses import CachedMultipleNegativesRankingLoss, MatryoshkaLoss
loss = CachedMultipleNegativesRankingLoss(model, mini_batch_size=1)
loss = MatryoshkaLoss(model, loss, matryoshka_dims=[2048, 1536, 1024, 512, 256, 128, 64])
```
This is especially useful for multimodal models, where embeddings can be large (2048 dimensions for Qwen3-VL). With Matryoshka training, you can use truncated embeddings (e.g., 256 or 128 dimensions) at deployment time for faster search with minimal quality loss. As I'll show in the Results section, the finetuned model achieves near-peak performance even at 512 dimensions.

The `SentenceTransformerTrainingArguments` class lets you control training hyperparameters. Here's the configuration used for the VDR finetuning:

```
from sentence_transformers.sentence_transformer.training_args import SentenceTransformerTrainingArguments, BatchSamplers
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
    fp16=False,
    bf16=True,
    batch_sampler=BatchSamplers.NO_DUPLICATES,
    # Optional tracking/debugging parameters:
    eval_strategy="steps",
    eval_steps=0.1,
    save_strategy="steps",
    save_steps=0.1,
    save_total_limit=2,
    logging_steps=0.05,
    run_name=run_name,
)
```
A few things to note for (multimodal) training:

- `bf16=True` : bfloat16 is generally preferred over float16 due to better numerical stability.
- `batch_sampler=BatchSamplers.NO_DUPLICATES` : When using`MultipleNegativesRankingLoss` or its cached variant, having no duplicate samples in a batch ensures that every in-batch negative is a truly different sample.
- `per_device_train_batch_size=64` : This may seem large for a 2B parameter VLM, but`CachedMultipleNegativesRankingLoss` with`mini_batch_size=1` handles the memory constraints through gradient caching.
- `eval_steps` ,`save_steps` , and`logging_steps` : Setting these to a fraction (e.g., 0.1) means evaluation, saving, and logging will happen every 10% of an epoch, which is useful for monitoring training progress.

To track retrieval performance before, during, and after training, I use the `InformationRetrievalEvaluator`. It computes standard retrieval metrics like NDCG@10, MAP, and Recall@k:

