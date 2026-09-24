---
id: collect-240926-huggingface/huggingface/google-t5-t5-small-hugging-face
title: "forward pass"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face"]
dates: []
keywords: ["apache", "compute", "inference", "license", "parameters", "research", "tpu", "training"]
source: docs/RAG/clean_en/huggingface/google-t5-t5-small-hugging-face.md
source_anchor: ""
source_lines: [1, 112]
sha256: b710b9191f3750960a25f2f2891e8294d131847d8430f6a49798c0057c80d5a7
---

# forward pass

<!-- source: https://huggingface.co/google-t5/t5-small -->

The developers of the Text-To-Text Transfer Transformer (T5) write:

With T5, we propose reframing all NLP tasks into a unified text-to-text-format where the input and output are always text strings, in contrast to BERT-style models that can only output either a class label or a span of the input. Our text-to-text framework allows us to use the same model, loss function, and hyperparameters on any NLP task.


T5-Small is the checkpoint with 60 million parameters.

- **Developed by:** Colin Raffel, Noam Shazeer, Adam Roberts, Katherine Lee, Sharan Narang, Michael Matena, Yanqi Zhou, Wei Li, Peter J. Liu. See associated paper and GitHub repo
- **Model type:** Language model
- **Language(s) (NLP):** English, French, Romanian, German
- **License:** Apache 2.0
- **Related Models:** All T5 Checkpoints
- **Resources for more information:**

The developers write in a blog post that the model:

Our text-to-text framework allows us to use the same model, loss function, and hyperparameters on any NLP task, including machine translation, document summarization, question answering, and classification tasks (e.g., sentiment analysis). We can even apply T5 to regression tasks by training it to predict the string representation of a number instead of the number itself.


See the blog post and research paper for further details.

More information needed.

More information needed.

More information needed.

The model is pre-trained on the Colossal Clean Crawled Corpus (C4), which was developed and released in the context of the same research paper as T5.

The model was pre-trained on a on a **multi-task mixture of unsupervised (1.) and supervised tasks (2.)**.
Thereby, the following datasets were being used for (1.) and (2.):

1. **Datasets used for Unsupervised denoising objective** :

1. **Datasets used for Supervised text-to-text language modeling objective**

- Sentence acceptability judgment
- Sentiment analysis 
  - SST-2 Socher et al., 2013
- Paraphrasing/sentence similarity
  - MRPC Dolan and Brockett, 2005
  - STS-B Ceret al., 2017
  - QQP Iyer et al., 2017
- Natural language inference
- Sentence completion
- Word sense disambiguation
- Question answering
  - MultiRC Khashabi et al., 2018
  - ReCoRD Zhang et al., 2018
  - BoolQ Clark et al., 2019

In their abstract, the model developers write:

In this paper, we explore the landscape of transfer learning techniques for NLP by introducing a unified framework that converts every language problem into a text-to-text format. Our systematic study compares pre-training objectives, architectures, unlabeled datasets, transfer approaches, and other factors on dozens of language understanding tasks.


The framework introduced, the T5 framework, involves a training procedure that brings together the approaches studied in the paper. See the research paper for further details.

The developers evaluated the model on 24 tasks, see the research paper for full details.

For full results for T5-small, see the research paper, Table 14.

Carbon emissions can be estimated using the Machine Learning Impact calculator presented in Lacoste et al. (2019).

- **Hardware Type:** Google Cloud TPU Pods
- **Hours used:** More information needed
- **Cloud Provider:** GCP
- **Compute Region:** More information needed
- **Carbon Emitted:** More information needed

**BibTeX:**

```
@article{2020t5,
  author  = {Colin Raffel and Noam Shazeer and Adam Roberts and Katherine Lee and Sharan Narang and Michael Matena and Yanqi Zhou and Wei Li and Peter J. Liu},
  title   = {Exploring the Limits of Transfer Learning with a Unified Text-to-Text Transformer},
  journal = {Journal of Machine Learning Research},
  year    = {2020},
  volume  = {21},
  number  = {140},
  pages   = {1-67},
  url     = {http://jmlr.org/papers/v21/20-074.html}
}
```
**APA:**

- Raffel, C., Shazeer, N., Roberts, A., Lee, K., Narang, S., Matena, M., ... & Liu, P. J. (2020). Exploring the limits of transfer learning with a unified text-to-text transformer. J. Mach. Learn. Res., 21(140), 1-67.

This model card was written by the team at Hugging Face.

Use the code below to get started with the model.

## Click to expand

```
from transformers import T5Tokenizer, T5Model
tokenizer = T5Tokenizer.from_pretrained("t5-small")
model = T5Model.from_pretrained("t5-small")
input_ids = tokenizer(
    "Studies have been shown that owning a dog is good for you", return_tensors="pt"
).input_ids  # Batch size 1
decoder_input_ids = tokenizer("Studies show that", return_tensors="pt").input_ids  # Batch size 1
# forward pass
outputs = model(input_ids=input_ids, decoder_input_ids=decoder_input_ids)
last_hidden_states = outputs.last_hidden_state
```
See the Hugging Face T5 docs and a Colab Notebook created by the model developers for more examples.

- Downloads last month
- 24,567,008
