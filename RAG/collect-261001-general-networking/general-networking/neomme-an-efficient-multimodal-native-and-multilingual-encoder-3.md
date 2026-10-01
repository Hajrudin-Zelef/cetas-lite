---
id: collect-261001-general-networking/general-networking/neomme-an-efficient-multimodal-native-and-multilingual-encoder-3
title: "accelerate is an optional dependency needed only when using device_map=\"auto\"."
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["compute", "embeddings", "fine-tuning", "inference", "multimodal", "parameters", "quantization", "throughput"]
source: docs/RAG/collect-261001-general-networking/neomme-an-efficient-multimodal-native-and-multilingual-encoder.md
source_anchor: ""
source_lines: [151, 167]
sha256: a880e592668a203b3eedc22e99f3a96a9a38c8b23b093920c6241a8cead525e5
---

# accelerate is an optional dependency needed only when using device_map="auto".

*NeoMME*-Retriever is a fine-tuned version of *NeoMME* for visual document retrieval. One forward pass produces both dense and late-interaction representations. The 260M model outperforms all evaluated models strictly below 800M parameters and, at a matched 2048×2048 input size, encodes pages at about 2× ColModernVBERT's throughput. To reduce the large storage footprint of late-interaction embeddings for high-resolution documents, we experimented with hierarchical token pooling and asymmetric quantization and managed to reduce the late-interaction embeddings from roughly 1.5 MB to 6 kB per page, a 255× compression, while retaining more than 95% of the baseline nDCG@10.

We release all *NeoMME* model checkpoints and a day-zero Hugging Face Transformers implementation to allow practitioners to build efficient multimodal and multilingual representation models on top of our work.

*NeoMME* began as a side quest between two good friends. We worked with limited time and compute, and we decided to share the results so the community can build on them. We thank H Company for supporting the work and providing the compute used to train *NeoMME*.

```
@misc{lac2026neommesingletowermultimodalnativemultilingual,
      title={NeoMME: A Single-Tower Multimodal-Native Multilingual Foundation Encoder for Efficient Fine-Tuning and Inference},
      author={Aurélien Lac and Tony Wu},
      year={2026},
      eprint={2609.01657},
      archivePrefix={arXiv},
      primaryClass={cs.IR},
      url={https://arxiv.org/abs/2609.01657},
}
```
