---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-7
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["attention", "benchmark", "cost", "embedding", "embeddings", "exploit", "flash attention", "gpu", "inference", "omni", "parameters", "quantization"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [594, 679]
sha256: 7b480d5343b6aa2369069a41f5e72034f0b48cd8b6d25a6c36890951a5634807
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

A cluster mean is a worse match for a query token than the best of its members was, and the coarser the clusters, the more that shows. The original experiments measured that cost on BEIR and found very little of it, at 100.6% of the unpooled retrieval performance on average at `pool_factor=2`, and 99.0% at `pool_factor=3`. Halving your index for free is a good deal, so 2 is a reasonable place to start. How much it costs on your data is corpus-specific though, so measure it with an evaluator before you settle on a factor. The runnable comparison is token_pooling.py.

How far you can push `pool_factor` is also partly a property of the model. LightOn's hierarchical pooling regularization trains for exactly that, shaping the embedding space so pooling costs less and reporting 99.4% retention at 5x compression. Training with that regularizer isn't in Sentence Transformers yet, but the resulting checkpoints are ordinary PyLate models, so `lightonai/LateOn-hpool-regularized` loads and pools like any other.

Multi-vector models run through the same backend machinery as the rest of Sentence Transformers, so you get `torch` (default), `onnx`, and `openvino`, alongside half precision, Flash Attention, and `torch.compile`.

On GPU, fp16 with Flash Attention is the best configuration we measured, at 2.44x the throughput of fp32 with no measurable retrieval quality loss. Flash Attention helps multi-vector models more than most, because documents are only truncated and never padded to a shared length, so your batches have widely varying sequence lengths that unpadding can exploit:

```
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder(
    "lightonai/GTE-ModernColBERT-v1",
    model_kwargs={"attn_implementation": "flash_attention_2", "dtype": "float16"},
)
```
Models with non-attend query expansion (`attend=False`, which covers the Stanford-NLP checkpoints like `colbert-ir/colbertv2.0` and `answerdotai/answerai-colbert-small-v1`) reject Flash Attention at load time. Flash Attention strips `attention_mask=0` positions, so the `[MASK]` expansion tokens that MaxSim scores would never receive an attention update. Use `"sdpa"` for those models.


On CPU, OpenVINO is your better bet where the architecture is supported, and int8 quantization buys a further speedup at a cost of about 0.4% accuracy. See Speeding up Inference for the full benchmark details, the export and quantization helpers, and a flowchart for picking a backend.

`MultiVectorNanoBEIREvaluator` runs the NanoBEIR suite of 13 small BEIR subsets with MaxSim scoring, and needs no data preparation on your side:

```
from sentence_transformers import MultiVectorEncoder
from sentence_transformers.multi_vector_encoder.evaluation import MultiVectorNanoBEIREvaluator
model = MultiVectorEncoder("lightonai/GTE-ModernColBERT-v1")
evaluator = MultiVectorNanoBEIREvaluator(batch_size=16)
results = evaluator(model)
print(f"{evaluator.primary_metric}: {results[evaluator.primary_metric]:.4f}")
```
This also makes it easy to check the claim from the top of this post. `lightonai/LateOn` and `lightonai/DenseOn` were trained by LightOn on the same data with the same ModernBERT backbone and the same 149M parameters, differing only in whether they keep one vector per token or pool down to one per document. Running both over all 13 NanoBEIR datasets isolates what that choice buys:

| NanoBEIR dataset | LateOn (multi-vector, 128d) | DenseOn (dense, 768d) | 
|---|---|---|
| MSMARCO | **0.7194** | 0.6517 | 
| NQ | **0.7810** | 0.7511 | 
| HotpotQA | **0.9295** | 0.8802 | 
| FEVER | **0.9702** | 0.9612 | 
| ClimateFEVER | **0.4887** | 0.4846 | 
| DBPedia | **0.6836** | 0.6748 | 
| QuoraRetrieval | **0.9795** | 0.9687 | 
| Touche2020 | **0.5938** | 0.5673 | 
| ArguAna | 0.5562 | **0.5660** | 
| NFCorpus | **0.3949** | 0.3851 | 
| SciFact | 0.7978 | **0.8057** | 
| SCIDOCS | 0.4469 | **0.4484** | 
| FiQA2018 | 0.5871 | **0.6491** | 
| **Mean** | **0.6868** | 0.6764 | 

Late interaction wins on 9 of the 13 datasets and on the mean, by roughly one NDCG point. The four it loses (ArguAna, FiQA2018, SCIDOCS, and SciFact) are the shape of the tradeoff you should expect, a real gain in retrieval quality at the same model size, paid for in index footprint, rather than a universal win on every dataset. The same pair scores 57.22 against 56.20 on the full 15-dataset BEIR, a comparable gap, so the margin is not an artifact of the small benchmark.

Alongside NanoBEIR, `MultiVectorInformationRetrievalEvaluator`, `MultiVectorRerankingEvaluator`, `MultiVectorTripletEvaluator`, and `MultiVectorDistillationEvaluator` cover the usual evaluation setups on your own data. They're documented in the Evaluation API Reference.

`MultiVectorEncoder` absorbs the modeling, inference, training, and evaluation of both libraries. Every PyLate checkpoint loads directly, and Supported Models lists the colpali-engine checkpoints along with the `revision` to pass where one is still needed. If you're migrating, these are the calls that change:

| PyLate | Sentence Transformers | 
|---|---|
| `pylate.models.ColBERT(model_name_or_path=...)` | `MultiVectorEncoder(...)` | 
| `model.encode(..., is_query=True)` | `model.encode_query(...)` | 
| `model.encode(..., is_query=False)` | `model.encode_document(...)` | 
| `pylate.scores.colbert_scores` | `model.similarity` | 
| `pylate.indexes.PLAID` /`pylate.retrieve.ColBERT` | no equivalent, keep PyLate's PLAID or see Indexing | 

| colpali-engine | Sentence Transformers | 
|---|---|
| `ColQwen2.from_pretrained(...)` +`ColQwen2Processor` | `MultiVectorEncoder(...)` | 
| `processor.process_queries(...)` +`model(**batch)` | `model.encode_query(queries)` | 
| `processor.process_images(...)` +`model(**batch)` | `model.encode_document(images)` | 
| `processor.score_multi_vector(qs, ds)` | `model.similarity(query_embeddings, document_embeddings)` | 
| `mask_non_image_embeddings=True` | `MultiVectorMask(keep_only_token_ids=[...])` | 
| `HierarchicalTokenPooler` | `HierarchicalTokenPooling` | 
| `colpali_engine.interpretability` | `sentence_transformers.multi_vector_encoder.interpretability` | 

The difference worth knowing about is on bare (non-ColBERT) checkpoints, where PyLate's `ColBERT("bert-base-uncased")` applies the classic recipe by default while `MultiVectorEncoder("bert-base-uncased")` builds a plain stack, leaving the prefixes, query expansion, and skiplist as explicit choices. The training loss and evaluator equivalents, and the data-handling differences, are in the Migration Guide.

Save compatibility is one-way in every case. PyLate, Stanford-NLP ColBERT, and colpali-engine checkpoints all load into `MultiVectorEncoder`, but `MultiVectorEncoder.save_pretrained` output isn't loadable by any of them.

Models carrying the `multi-vector` and `sentence-transformers` tags on the Hub are the list that stays current, and we're working to get those tags onto every model that works. The tables below are what we test against directly, so treat them as a starting point rather than the full set. For text retrieval in particular, any PyLate or Stanford-NLP ColBERT checkpoint loads whether or not it carries the tag yet.

Some entries need a small Sentence Transformers configuration added to their repository first, and several of those are still open pull requests at the time of writing. Where a `revision` is listed below, pass it until that pull request is merged, after which the plain model name is enough:

```
model = MultiVectorEncoder("vidore/colqwen-omni-v0.1", revision="refs/pr/N")
```
These load with their trained prefix tokens, query expansion, and punctuation skiplist recovered from the saved configuration.

