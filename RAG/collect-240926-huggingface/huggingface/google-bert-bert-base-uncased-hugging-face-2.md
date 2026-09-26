---
id: collect-240926-huggingface/huggingface/google-bert-bert-base-uncased-hugging-face-2
title: "google-bert-bert-base-uncased-hugging-face"
domain: huggingface
role: reference
task: reference
actors: []
dates: []
keywords: ["tpu", "training"]
source: docs/RAG/clean_en/huggingface/google-bert-bert-base-uncased-hugging-face.md
source_anchor: ""
source_lines: [152, 182]
sha256: f2a1f1de65074bb10f1ba0d9781b4472d6aeebd81d480abf7f5cc7d646a3b367
---

# google-bert-bert-base-uncased-hugging-face

The model was trained on 4 cloud TPUs in Pod configuration (16 TPU chips total) for one million steps with a batch size of 256. The sequence length was limited to 128 tokens for 90% of the steps and 512 for the remaining 10%. The optimizer used is Adam with a learning rate of 1e-4, and , a weight decay of 0.01, learning rate warmup for 10,000 steps and linear decay of the learning rate after.

When fine-tuned on downstream tasks, this model achieves the following results:

Glue test results:

| Task | MNLI-(m/mm) | QQP | QNLI | SST-2 | CoLA | STS-B | MRPC | RTE | Average | 
|---|---|---|---|---|---|---|---|---|---|
|  | 84.6/83.4 | 71.2 | 90.5 | 93.5 | 52.1 | 85.8 | 88.9 | 66.4 | 79.6 | 

```
@article{DBLP:journals/corr/abs-1810-04805,
  author    = {Jacob Devlin and
               Ming{-}Wei Chang and
               Kenton Lee and
               Kristina Toutanova},
  title     = {{BERT:} Pre-training of Deep Bidirectional Transformers for Language
               Understanding},
  journal   = {CoRR},
  volume    = {abs/1810.04805},
  year      = {2018},
  url       = {http://arxiv.org/abs/1810.04805},
  archivePrefix = {arXiv},
  eprint    = {1810.04805},
  timestamp = {Tue, 30 Oct 2018 20:39:56 +0100},
  biburl    = {https://dblp.org/rec/journals/corr/abs-1810-04805.bib},
  bibsource = {dblp computer science bibliography, https://dblp.org}
}
```
- Downloads last month
- 44,984,761
