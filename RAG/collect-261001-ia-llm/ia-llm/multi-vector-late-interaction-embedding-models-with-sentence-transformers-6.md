---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-6
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["consumer", "embedding", "embeddings", "memory", "omni", "training", "transcription"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [479, 593]
sha256: 4b387dd82a1c18601e8ba5a2d7298c64f2f6c28857b02d3f0e06839d9f19da5a
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

```
# pip install -U "sentence-transformers[audio,video]"
import torch
from datasets import Audio, load_dataset
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder(
    "vidore/colqwen-omni-v0.1",
    model_kwargs={"dtype": torch.bfloat16},
)
print(model.modalities)
# ['text', 'image', 'audio', 'video', 'message']
# 20 recorded conversations, averaging 28 seconds each
dataset = load_dataset("eustlb/dailytalk-conversations-grouped", split="train[:20]")
dataset = dataset.cast_column("audio", Audio(sampling_rate=16_000))
audio = [row["array"] for row in dataset["audio"]]  # raw mono waveforms, float32 at 16 kHz
query_embeddings = model.encode_query(["medicine for car nausea"])
document_embeddings = model.encode_document(audio, batch_size=2)
scores = model.similarity(query_embeddings, document_embeddings)[0]
top_scores, top_indices = scores.topk(3)
for score, index in zip(top_scores.tolist(), top_indices.tolist()):
    print(f"{score:.4f}  {' / '.join(dataset[index]['texts'][:2])}")
"""
50.8902  Excuse me? Do you have anything for a carsickness? / Yes, but you look fine.
46.1028  Excuse me, could you tell me where you have got that music book? / Certainly. Let me see. Oh, it's on that shelf.
46.0514  Jeff, I'm going to the supermarket. Do you want to come with me? / I think the supermarket is closed now.
"""
```
ColQwen-Omni was trained purely on image-text pairs, so its audio retrieval is zero-shot. It never heard a training example, and there is no transcription step anywhere in the pipeline. The query says `nausea` where the recording says `carsickness`, and it still picks the pharmacy conversation out of twenty by a wide margin.

Video works the same way, but sample the frames or it will eat your VRAM. Its release blogpost is blunt about this, that video "is very memory-intensive, so it's best suited for short clips":

```
import torch
from sentence_transformers import MultiVectorEncoder
model = MultiVectorEncoder(
    "vidore/colqwen-omni-v0.1",
    model_kwargs={"dtype": torch.bfloat16},
)
# Sparse, low-resolution frames: 0.5 fps rather than the full frame rate
model[0].processing_kwargs.update(
    {"video": {"max_pixels": 32 * 28 * 28, "do_sample_frames": True, "fps": 0.5}}
)
query_embeddings = model.encode_query(["How to cook Mapo Tofu?"])
document_embeddings = model.encode_document([
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/mapo_tofu.mp4",
    "https://huggingface.co/datasets/sentence-transformers/example-documents/resolve/main/zhajiang_noodle.mp4",
], batch_size=1)
print(model.similarity(query_embeddings, document_embeddings))
# tensor([[53.3100, 51.0561]])
```
At 1 fps and full resolution the same pair of videos produces 8,426 and 5,137 token vectors and peaks at 20.8 GB of VRAM, against 4,240 and 2,446 vectors and 12.5 GB here, for a model that occupies 9.0 GB on its own. The ranking is identical either way. Long audio wants the same treatment, and the release blogpost recommends 30-second chunks, which come to roughly 800 tokens each.

Because MaxSim is a sum of per-query-token maxima, a ranking decomposes exactly, so every point of a document's score belongs to one query token and one document token. That lets you answer "why did this rank here?" precisely, rather than by eye.

For image documents, `sentence_transformers.multi_vector_encoder.interpretability` overlays that decomposition onto the page as the standard ColPali heatmap, either aggregated over the query or one map per query token. Asking "How much was spent on water resources and power?" against the outlays page from above, this is where the `water` token went:

heatmap.py is the runnable version, including the masking step that lines the document embedding up with the patch grid.

Text documents have no patch grid to overlay, but the same decomposition applies. text_similarity_map.py ranks a corpus and then attributes the top hit's score token by token, here on the Natural Questions corpus from earlier with the 32M-parameter mxbai-edge-colbert-v0-32m:

```
Query: when did richmond last play in a preliminary final
Top 3 of 4874 documents by exhaustive MaxSim (191.0ms):
  12.3489  Richmond Football Club Richmond began 2017 with 5 straight wins, a feat it had not achieved since 19
  12.1771  2017 AFL Grand Final The 2017 AFL Grand Final was an Australian rules football game contested betwee
  12.0591  2018 UEFA Champions League Final The 2018 UEFA Champions League Final was the final match of the 201
  query token       best document token      sim   share
  when              since                 0.9154    7.4%
  did               had                   0.9675    7.8%
  rich              rich                  0.9764    7.9%
  mond              mond                  0.9856    8.0%
  last              to                    0.9249    7.5%
  play              game                  0.9384    7.6%
  in                the                   0.9732    7.9%
  a                 a                     0.9587    7.8%
  preliminary       preliminary           0.9394    7.6%
  final             final                 0.9654    7.8%
  --------------------------------------------------------
  3 special tokens                        2.8038   22.7%
  MaxSim score                           12.3489  100.0%
```
`rich`, `mond`, `preliminary`, and `final` matched themselves, while `when` settled on `since` and `play` on `game`. Three special tokens contribute 22.7% of the score while carrying none of the query's content. Below this table the script prints the passage itself, with the winning tokens highlighted in place.

If the index footprint worries you, the most effective knob is to store fewer token vectors. `HierarchicalTokenPooling` implements the token pooling technique from Clavié, Chaffin, and Adams by clustering each document's token vectors with Ward linkage on cosine distance and replacing each cluster with its mean, keeping roughly `1 / pool_factor` of the tokens. Within one document a lot of token vectors end up close to each other, so much of what you drop is redundancy rather than signal:

```
from datasets import load_dataset
from sentence_transformers import MultiVectorEncoder
from sentence_transformers.multi_vector_encoder.modules import HierarchicalTokenPooling
dataset = load_dataset("sentence-transformers/natural-questions", split="train[:5000]")
documents = list(dict.fromkeys(dataset["answer"]))
model = MultiVectorEncoder("lightonai/LateOn")
pooling = HierarchicalTokenPooling(pool_factor=2)
document_embeddings = model.encode_document(documents, token_pooling=pooling)
```
There are three places to apply it, depending on when you want to pay for it:

```
# 1. Per encode call, as above
document_embeddings = model.encode_document(documents, token_pooling=pooling)
# 2. Standalone, on embeddings you already have saved (e.g. list of [num_tokens, num_dims] tensors)
pooled = pooling.pool(document_embeddings)
# 3. Baked into the model, so every consumer of the checkpoint gets pooled documents
model.append(HierarchicalTokenPooling(pool_factor=2))
model.save_pretrained("my-pooled-colbert")
```
By default, pooling applies to documents only, since queries are short and are the side you can't afford to distort. On the Natural Questions corpus from earlier, the reduction tracks `pool_factor` closely, and pooling all 608k token vectors took about 6 seconds:

| `pool_factor` | Token vectors | Reduction | float32 index | 
|---|---|---|---|
| 1 (off) | 608,414 | 1.00x | 311.5 MB | 
| 2 | 305,438 | 1.99x | 156.4 MB | 
| 3 | 204,407 | 2.98x | 104.7 MB | 
| 4 | 153,936 | 3.95x | 78.8 MB | 

