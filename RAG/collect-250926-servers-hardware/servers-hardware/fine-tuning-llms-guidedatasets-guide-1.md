---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-guidedatasets-guide-1
title: "Datasets Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "Unsloth"]
dates: []
keywords: ["embeddings", "fine-tuning", "pretraining", "research", "rlhf", "training"]
source: docs/RAG/clean4/fine-tuning-llms-guidedatasets-guide.md
source_anchor: ""
source_lines: [1, 70]
sha256: fed85acdd08026898c1fa7b97b155529b18bfdcfcc2e78c0ec9352aae481d780
---

# Datasets Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/fine-tuning-llms-guide/datasets-guide.md).
# Datasets Guide
Learn how to create & prepare a dataset for fine-tuning.
## What is a Dataset?
For LLMs, datasets are collections of data that can be used to train our models. In order to be useful for training, text data needs to be in a format that can be tokenized. You'll also learn how to [use datasets inside of Unsloth](#applying-chat-templates-with-unsloth).
One of the key parts of creating a dataset is your [chat template](/docs/basics/chat-templates.md) and how you are going to design it. Tokenization is also important as it breaks text into tokens, which can be words, sub-words, or characters so LLMs can process it effectively. These tokens are then turned into embeddings and are adjusted to help the model understand the meaning and context.
### Data Format
To enable the process of tokenization, datasets need to be in a format that can be read by a tokenizer.

Format

Description

Training Type

Raw Corpus

Raw text from a source such as a website, book, or article.

Continued Pretraining (CPT)

Instruct

Instructions for the model to follow and an example of the output to aim for.

Supervised fine-tuning (SFT)

Conversation

Multiple-turn conversation between a user and an AI assistant.

Supervised fine-tuning (SFT)

RLHF

Conversation between a user and an AI assistant, with the assistant's responses being ranked by a script, another model or human evaluator.

Reinforcement Learning (RL)

{% hint style="info" %}
It's worth noting that different styles of format exist for each of these types.
{% endhint %}
## Getting Started
Before we format our data, we want to identify the following:
{% stepper %}
{% step %} Purpose of dataset
Knowing the purpose of the dataset will help us determine what data we need and format to use.
The purpose could be, adapting a model to a new task such as summarization or improving a model's ability to role-play a specific character. For example:
* Chat-based dialogues (Q\&A, learn a new language, customer support, conversations).
* Structured tasks ([classification](https://colab.research.google.com/github/timothelaborie/text_classification_scripts/blob/main/unsloth_classification.ipynb), summarization, generation tasks).
* Domain-specific data (medical, finance, technical).
{% endstep %}
{% step %} Style of output
The style of output will let us know what sources of data we will use to reach our desired output.
For example, the type of output you want to achieve could be JSON, HTML, text or code. Or perhaps you want it to be Spanish, English or German etc.
{% endstep %}
{% step %} Data source
When we know the purpose and style of the data we need, we need to analyze the quality and [quantity](#how-big-should-my-dataset-be) of the data. Hugging Face and Wikipedia are great sources of datasets and Wikipedia is especially useful if you are looking to train a model to learn a language.
The Source of data can be a CSV file, PDF or even a website. You can also [synthetically generate](#synthetic-data-generation) data but extra care is required to make sure each example is high quality and relevant.
{% endstep %}
{% endstepper %}
{% hint style="success" %}
One of the best ways to create a better dataset is by combining it with a more generalized dataset from Hugging Face like ShareGPT to make your model smarter and diverse. You could also add [synthetically generated data](#synthetic-data-generation).
{% endhint %}
## 🦥 Unsloth Data Recipes
[Unsloth Data Recipes](/docs/new/studio/data-recipe.md) lets you upload documents like PDFs or CSVs files and transforms them into useable datasets. Create and edit datasets visually via a graph-node workflow.
The recipes page is the main entry point. Recipes are stored locally in the browser, so you come back to saved work later. From here, you can create a blank recipe or open a guided learning recipe.

Using the dataset example I provided, follow the structure and generate conversations based on the examples.

