---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers-4
title: "['text', 'image', 'video', 'message']"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["attention", "embedding", "embeddings", "inference", "llama", "multimodal", "nvidia", "omni", "parameters", "quantization", "qwen", "reranker"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [310, 427]
sha256: 6154f05e31ecbe77de62555e3089631dbe965d6432b765d4cc7db0f08354ecc3
---

# ['text', 'image', 'video', 'message']

| Model | Parameters | NDCG@10 | 
|---|---|---|
| **tomaarsen/Qwen3-VL-Embedding-2B-vdr** | **2.1B** | **0.947** | 
| Qwen/Qwen3-VL-Embedding-8B | 8.1B | 0.923 | 
| nvidia/omni-embed-nemotron-3b | 4.7B | 0.915 | 
| nvidia/llama-nemotron-embed-vl-1b-v2 | 1.7B | 0.912 | 
| nomic-ai/nomic-embed-multimodal-7b | 8.3B | 0.912 | 
| llamaindex/vdr-2b-multi-v1 | 2.2B | 0.912 | 
| llamaindex/vdr-2b-v1 | 2.2B | 0.911 | 
| nomic-ai/nomic-embed-multimodal-3b | 3.8B | 0.899 | 
| Qwen/Qwen3-VL-Embedding-2B | 2.1B | 0.888 | 
| LCO-Embedding/LCO-Embedding-Omni-7B | 8.9B | 0.888 | 
| LCO-Embedding/LCO-Embedding-Omni-3B | 4.7B | 0.860 | 
| BAAI/BGE-VL-v1.5-zs | 7.6B | 0.800 | 
| BAAI/BGE-VL-v1.5-mmeb | 7.6B | 0.797 | 
| BAAI/BGE-VL-MLLM-S2 | 7.6B | 0.792 | 
| BidirLM/BidirLM-Omni-2.5B-Embedding | 2.5B | 0.775 | 
| royokong/e5-v | 8.4B | 0.767 | 
| BAAI/BGE-VL-MLLM-S1 | 7.6B | 0.710 | 
| sentence-transformers/clip-ViT-L-14 | 428M | 0.611 | 
| BAAI/BGE-VL-large | 428M | 0.467 | 
| BAAI/BGE-VL-base | 150M | 0.335 | 

The finetuned 2B model outperforms even the 8B Qwen3-VL-Embedding model, demonstrating the power of task-specific finetuning. Finetuning on your own domain is often worth considering, even when a larger general-purpose model is available!

The comparison above uses full-size 2048-dim embeddings. Thanks to the Matryoshka training, the finetuned model also holds up well when truncated to fewer dimensions, letting you trade off embedding size and retrieval quality at deployment time:

The finetuned model's peak is at the full 2048 dimensions (0.948), but it stays within 0.3% of peak all the way down to 512 (4x smaller), and retains over 92% of peak even at 64 (32x smaller). Matryoshka training concentrates the most important information in the earlier dimensions, so moderate truncation costs very little performance.


## Full NDCG@10 numbers by dimension

| Dimensions | Base NDCG@10 | Finetuned NDCG@10 | 
|---|---|---|
| 2048 (full) | 0.8961 (100%) | 0.9480 (100%) | 
| 1536 | 0.8940 (99.8%) | 0.9439 (99.6%) | 
| 1024 | 0.8941 (99.8%) | 0.9464 (99.8%) | 
| 512 | 0.8760 (97.8%) | 0.9451 (99.7%) | 
| 256 | 0.8347 (93.2%) | 0.9372 (98.9%) | 
| 128 | 0.7888 (88.0%) | 0.9058 (95.5%) | 
| 64 | 0.6852 (76.5%) | 0.8758 (92.4%) | 

The gap between 1024 and 2048 dimensions is small (0.946 vs. 0.948), so I've saved the model with `truncate_dim=1024` set in its configuration. This means that `SentenceTransformer("tomaarsen/Qwen3-VL-Embedding-2B-vdr")` produces 1024-dimensional embeddings by default, halving the storage footprint compared to the full 2048. If you want a different dimensionality, pass `truncate_dim=N` when loading to override it.

You can also finetune multimodal Cross Encoder (reranker) models using the same training infrastructure. The key difference is using `CrossEncoderTrainer` and Cross Encoder-specific loss functions. This section provides a brief overview; see the full training examples for complete, runnable scripts with dataset preparation and evaluation.

Here's a simplified example based on the doodles training script, which trains a reranker to match images with text captions:

```
from sentence_transformers.cross_encoder import CrossEncoder
from sentence_transformers.cross_encoder.losses import BinaryCrossEntropyLoss
from sentence_transformers.cross_encoder.modules import LogitScore, Transformer
from sentence_transformers.cross_encoder.trainer import CrossEncoderTrainer
from sentence_transformers.cross_encoder.training_args import CrossEncoderTrainingArguments
# 1. Build the model from modules
transformer = Transformer(
    "Qwen/Qwen3.5-0.8B",
    transformer_task="any-to-any",
    model_kwargs={"torch_dtype": "bfloat16", "device_map": "auto", "attn_implementation": "flash_attention_2"},
    processing_kwargs={"chat_template": {"add_generation_prompt": True}},
)
# Extend chat template to support "query" and "document" roles
transformer.processor.chat_template = transformer.processor.chat_template.replace(
    'message.role == "user"', 'message.role in ["user", "query", "document"]'
)
# LogitScore: score = log(P("1")) - log(P("0"))
score_head = LogitScore(
    true_token_id=transformer.tokenizer.convert_tokens_to_ids("1"),
    false_token_id=transformer.tokenizer.convert_tokens_to_ids("0"),
)
model = CrossEncoder(
    modules=[transformer, score_head],
    num_labels=1,
    prompts={
        "image_to_text": "Given the image, judge whether the text matches it. Respond with 1 if they match, 0 if they don't.",
        "text_to_image": "Given the text, judge whether the image matches it. Respond with 1 if they match, 0 if they don't.",
    },
)
# 2. Define the loss
loss = BinaryCrossEntropyLoss(model)
# 3. Multi-dataset training with separate directions
trainer = CrossEncoderTrainer(
    model=model,
    args=args,
    train_dataset={"image_to_text": train_image_to_text, "text_to_image": train_text_to_image},
    eval_dataset={"image_to_text": eval_image_to_text, "text_to_image": eval_text_to_image},
    loss=loss,
    evaluator=[image_to_text_evaluator, text_to_image_evaluator],
)
trainer.train()
```
There are multiple valid architectural choices for multimodal rerankers, including:

1. Any-to-Any + LogitScore: Uses the multimodal language model to generate a token, then computes the log-odds of "1" vs "0".
2. Feature Extraction + Pooling + Dense: Uses only the multimodal base model, and extracts the last token's hidden state and projects it to a score via a Dense layer, avoiding the language modeling head computation.

Both approaches are demonstrated in the multimodal cross encoder training examples.

The two scripts linked above split the training data into two datasets, one per direction (image-to-text and text-to-image), with a task-specific prompt for each that tells the model how to score in that direction. Each positive pair is then expanded with randomly sampled negatives so the loss sees a balanced mix of matches and non-matches.

The Sentence Transformers repository includes several multimodal training examples:

- Visual Document Retrieval: The training script used in this blogpost to finetune a VLM-based embedding model for document screenshot retrieval
- Multimodal Reranker (Any-to-Any): Train a multimodal reranker using LogitScore
- Multimodal Reranker (Feature Extraction): Train a multimodal reranker using Pooling + Dense

Additionally, the following pages may be useful to learn more about training with Sentence Transformers:

The direct prerequisite and the prior training guides are listed first, followed by related technique posts that stack with multimodal training:

- Multimodal Embedding & Reranker Models with Sentence Transformers: multimodal inference, the companion to this training post.
- Training and Finetuning Embedding Models with Sentence Transformers: the general training guide for text-only embedding models.
- Training and Finetuning Reranker Models with Sentence Transformers: Cross Encoder (reranker) training, including the ModernBERT-base reranker example.
- Training and Finetuning Sparse Embedding Models with Sentence Transformers: SPLADE and other sparse encoder training.
- 🪆 Introduction to Matryoshka Embedding Models: background on the `MatryoshkaLoss` used in this post to produce embeddings that truncate gracefully.
- Train 400x faster Static Embedding Models with Sentence Transformers: CPU-friendly text embeddings, a natural counterpart to heavy VLM-based models.
- Binary and Scalar Embedding Quantization for Significantly Faster & Cheaper Retrieval: post-training compression that also applies to multimodal embeddings.
- Visual Document Retrieval Goes Multilingual: the LlamaIndex post that introduced the dataset used in this blogpost's Visual Document Retrieval example.
