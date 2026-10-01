---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers-5
title: "Loading in fp32 is preferred for training if your memory can handle it"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["training", "attention", "benchmark", "consumer", "cost", "distillation", "embedding", "embeddings", "gpu", "lora", "multimodal", "parameters"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multi-vector-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [377, 442]
sha256: 1985645e28c9482c989c87e8de5221f9b1774c59c961512e65c392990df7f4c1
---

# Loading in fp32 is preferred for training if your memory can handle it

The finetuned model tops the table, beating the strongest zero-shot model of any architecture by +0.062 NDCG@10. In other words, the strongest zero-shot model returns the right passage as the very first hit for 75.8% of the queries, while the finetuned model does so for 84.9%, cutting the rank-1 error by more than a third.

The architecture pattern is just as clear, with the top of the table exclusively late interaction. On long documents, one vector per token beats one vector per document, even at matched training and matched backbones. DenseOn and LateOn share training data and architecture except for the head, and the late-interaction sibling wins by +0.12, with the multilingual pair (mDenseOn and mLateOn) replicating this at +0.13. Scale doesn't rescue single vectors either. Qwen3-Embedding-4B, the strongest dense model with roughly 33x the active (non-embedding) parameters of mine, still stops 0.13 short, and the 8B version scores lower than the 4B.

BM25 also performs surprisingly well, beating every sparse model, every truncation-capped multi-vector model, and all but three dense models: the multi-billion Qwen3-Embedding-4B and 8B, and voyage-4-nano, which reads its full 32k token context to edge past by just 0.006. Don't expect that to transfer to your own data though. MIRIAD's questions are generated from the passages, so the lexical overlap between a query and its gold passage is far larger than in typical retrieval, and BM25's unlimited context length lets it use every one of those overlapping words while most neural checkpoints truncate. A BM25 baseline is cheap and always worth running, just don't count on this margin.

The full field at a glance, sorted by score and colored by architecture family.

## Click to see the full evaluation table

Models marked `@N` are evaluated with their document length cap lifted to N tokens, since their native caps (180 to 512 tokens) would otherwise truncate the 941-token average passages. For every multi-vector model this lift was worth +0.08 to +0.24 NDCG@10 over the as-served row, and even the dense DenseOn gained +0.03 from the same treatment.

Note that this does not mean that multi-vector-encoder/mLateOn-medical is the strongest model on *all* domains. It's simply the strongest in *my* domain. This is totally fine, as I just need this model to work well on my data.

Don't underestimate the power of finetuning multi-vector models on your domain. Fourteen and a half hours on a single consumer GPU produced a model that no general-purpose retriever comes close to on this data, and the recipe is a single script with no teacher model and no mined negatives!

The fair objection to multi-vector retrieval is index size, and this domain is close to the worst case for it. Storing one vector per token, my model needs about 878 vectors per passage, so the 200,000-passage corpus takes roughly 45 GB at fp16, where a dense model needs well under 1 GB. Document length is what makes that gap so wide. The Natural Questions passages in the companion post average about 125 token vectors each, seven times fewer, so a corpus of short passages starts from a far smaller index than this one does. The `HierarchicalTokenPooling` module compresses exactly this by clustering each document's token embeddings and storing the cluster means, keeping roughly `1 / pool_factor` of the vectors:

```
from sentence_transformers.multi_vector_encoder.modules import HierarchicalTokenPooling
pooling = HierarchicalTokenPooling(pool_factor=4)
document_embeddings = model.encode_document(passages, token_pooling=pooling)
```
I measured it post-hoc on the finished model, with no pooling-aware training, and on long documents it is remarkably cheap.

The solid points are uncompressed embeddings, so that every family is counted the same way and scored with exact search. You would not deploy any of them like that, though. Dense indexes routinely use int8 or binary quantization with rescoring, sparse indexes compress their postings, and multi-vector indexes use PLAID-style residual compression. Don't read those points as the disk you need to buy, but as relative storage cost.

Token pooling is the solid line. Halving the vector count costs 0.0033 NDCG@10 and leaves rank-1 accuracy untouched, and keeping only a quarter of them, at 11.2 GB, still scores 0.8991. The curve keeps going (I measured out to a tenth of the vectors, still at 0.8765) but there is little reason to push pooling that far once quantization is on the table, which is what the dashed line below is about.

The dashed line is what a real deployment might look like. I gave Omar Khattab early access to the model and the benchmark, and he measured these configurations with fast-plaid at 1-bit residual quantization, using compact 17-bit centroid ids and 18-bit document ids instead of its ordinary unpacked 64-bit integers, plus document-side pruning:

| configuration | vectors kept | index | NDCG@10 | 
|---|---|---|---|
| 1-bit PLAID, all vectors | 100% | 3.37 GB | 0.8984 | 
| 1-bit PLAID + pruning | 65% | 2.23 GB | 0.8830 | 
| 1-bit PLAID + pruning | 42% | 1.45 GB | 0.8642 | 

That first row is 13x smaller than the raw embeddings, for 0.0155 NDCG@10. That is a far better trade than anywhere on the pooling curve. Quantization shrinks each vector while pooling and pruning cut how many you keep, so they compose, and quantization is the one to reach for first. Push further and the last row lands at 1.45 GB, *smaller* than the fp16 embeddings of Qwen3-Embedding-8B (1.64 GB), while scoring 0.0895 higher. The objection that multi-vector indexes are too big does not survive a properly configured index.

The pruning here is naive, meant only to establish that token reduction works on top of quantization, so read the bottom two rows as a floor rather than the frontier. If you would rather not hand-tune quantization at all, the Indexing section of the companion post covers fast-plaid, Qdrant, Weaviate, and Vespa.

Multi-vector retrieval is only as expensive as its index. The raw embeddings for this corpus are 45 GB, and a properly configured index is at least 7x smaller at nearly the same accuracy. The index deserves as much of your attention as the checkpoint.

Thanks to Omar Khattab for measuring the quantized and pruned index configurations in Optimizing the index, and for the discussions around late-interaction index costs.

These pages have training examples with explanations as well as links to training scripts. You can use them to get familiar with the multi-vector training loop:

- MIRIAD: domain-specific training on medical retrieval, an earlier and simpler cousin of this blogpost's recipe
- MS MARCO: contrastive and knowledge distillation recipes
- Multimodal: ColPali-style visual document retrieval training
- PEFT Adapters: parameter-efficient finetuning with LoRA

For further learning, you may also want to explore the following resources on Sentence Transformers:

- Installation
- Quickstart
- Usage
- Creating Custom Models
- Pretrained Models
- Training Overview (This blogpost is a distillation of the Training Overview documentation)
- Loss Overview
- API Reference

And here is an advanced page that might interest you:

And the companion blogpost, covering everything about *using* these models:
