---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-guidedatasets-guide
title: "Datasets Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["DeepSeek", "Google", "Hugging Face", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agents", "chatgpt", "deepseek", "embeddings", "fine-tuning", "inference", "llama", "multimodal", "pretraining", "reasoning", "research"]
source: docs/RAG/clean4/fine-tuning-llms-guidedatasets-guide.md
source_anchor: ""
source_lines: [1, 240]
sha256: b07592d05703938652df91fc656952a0012e9d77cc5346f56ea023d4d4ea6975
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

* **Prompt if you no have dataset**:
{% code overflow="wrap" %}
```
Create 10 examples of product reviews for Coca-Coca classified as either positive, negative, or neutral.
```
{% endcode %}
* **Prompt for a dataset without formatting**:
{% code overflow="wrap" %}
```
Structure my dataset so it is in a QA ChatML format for fine-tuning. Then generate 5 synthetic data examples with the same topic and format.
```
{% endcode %}
It is recommended to check the quality of generated data to remove or improve on irrelevant or poor-quality responses. Depending on your dataset it may also have to be balanced in many areas so your model does not overfit. You can then feed this cleaned dataset back into your LLM to regenerate data, now with even more guidance.
## Dataset FAQ + Tips
### How big should my dataset be?
We generally recommend using a bare minimum of at least 100 rows of data for fine-tuning to achieve reasonable results. For optimal performance, a dataset with over 1,000 rows is preferable, and in this case, more data usually leads to better outcomes. If your dataset is too small you can also add synthetic data or add a dataset from Hugging Face to diversify it. However, the effectiveness of your fine-tuned model depends heavily on the quality of the dataset, so be sure to thoroughly clean and prepare your data.
### How should I structure my dataset if I want to fine-tune a reasoning model?
If you want to fine-tune a model that already has reasoning capabilities like the distilled versions of DeepSeek-R1 (e.g. DeepSeek-R1-Distill-Llama-8B), you will need to still follow question/task and answer pairs however, for your answer you will need to change the answer so it includes reasoning/chain-of-thought process and the steps it took to derive the answer.\
\
For a model that does not have reasoning and you want to train it so that it later encompasses reasoning capabilities, you will need to utilize a standard dataset but this time without reasoning in its answers. This is training process is known as [Reinforcement Learning and GRPO](/docs/get-started/reinforcement-learning-rl-guide.md).
### Multiple datasets
If you have multiple datasets for fine-tuning, you can either:
* Standardize the format of all datasets, combine them into a single dataset, and fine-tune on this unified dataset.
* Use the [Multiple Datasets](https://colab.research.google.com/drive/1njCCbE1YVal9xC83hjdo2hiGItpY_D6t?usp=sharing) notebook to fine-tune on multiple datasets directly.
### Can I fine-tune the same model multiple times?
You can fine-tune an already fine-tuned model multiple times, but it's best to combine all the datasets and perform the fine-tuning in a single process instead. Training an already fine-tuned model can potentially alter the quality and knowledge acquired during the previous fine-tuning process.
## Using Datasets in Unsloth
### Alpaca Dataset
See an example of using the Alpaca dataset inside of Unsloth on Google Colab:
We will now use the Alpaca Dataset created by calling GPT-4 itself. It is a list of 52,000 instructions and outputs which was very popular when Llama-1 was released, since it made finetuning a base LLM be competitive with ChatGPT itself.
You can access the GPT4 version of the Alpaca dataset [here](https://huggingface.co/datasets/vicgalle/alpaca-gpt4.). Below shows some examples of the dataset:
You can see there are 3 columns in each row - an instruction, and input and an output. We essentially combine each row into 1 large prompt like below. We then use this to finetune the language model, and this made it very similar to ChatGPT. We call this process **supervised instruction finetuning**.
### Multiple columns for finetuning
But a big issue is for ChatGPT style assistants, we only allow 1 instruction / 1 prompt, and not multiple columns / inputs. For example in ChatGPT, you can see we must submit 1 prompt, and not multiple prompts.
This essentially means we have to "merge" multiple columns into 1 large prompt for finetuning to actually function!
For example the very famous Titanic dataset has many many columns. Your job was to predict whether a passenger has survived or died based on their age, passenger class, fare price etc. We can't simply pass this into ChatGPT, but rather, we have to "merge" this information into 1 large prompt.
For example, if we ask ChatGPT with our "merged" single prompt which includes all the information for that passenger, we can then ask it to guess or predict whether the passenger has died or survived.
Other finetuning libraries require you to manually prepare your dataset for finetuning, by merging all your columns into 1 prompt. In Unsloth, we simply provide the function called `to_sharegpt` which does this in 1 go!
Now this is a bit more complicated, since we allow a lot of customization, but there are a few points:
* You must enclose all columns in curly braces `{}`. These are the column names in the actual CSV / Excel file.
* Optional text components must be enclosed in `[[]]`. For example if the column "input" is empty, the merging function will not show the text and skip this. This is useful for datasets with missing values.
* Select the output or target / prediction column in `output_column_name`. For the Alpaca dataset, this will be `output`.
For example in the Titanic dataset, we can create a large merged prompt format like below, where each column / piece of text becomes optional.
For example, pretend the dataset looks like this with a lot of missing data:
| Embarked | Age | Fare |
| -------- | --- | ---- |
| S | 23 | |
| | 18 | 7.25 |
Then, we do not want the result to be:
1. The passenger embarked from S. Their age is 23. Their fare is **EMPTY**.
2. The passenger embarked from **EMPTY**. Their age is 18. Their fare is $7.25.
Instead by optionally enclosing columns using `[[]]`, we can exclude this information entirely.
1. \[\[The passenger embarked from S.]] \[\[Their age is 23.]] \[\[Their fare is **EMPTY**.]]
2. \[\[The passenger embarked from **EMPTY**.]] \[\[Their age is 18.]] \[\[Their fare is $7.25.]]
becomes:
1. The passenger embarked from S. Their age is 23.
2. Their age is 18. Their fare is $7.25.
### Multi turn conversations
An issue if you didn't notice is the Alpaca dataset is single turn, whilst remember using ChatGPT was interactive and you can talk to it in multiple turns. For example, the left is what we want, but the right which is the Alpaca dataset only provides singular conversations. We want the finetuned language model to somehow learn how to do multi turn conversations just like ChatGPT.
So we introduced the `conversation_extension` parameter, which essentially selects some random rows in your single turn dataset, and merges them into 1 conversation! For example, if you set it to 3, we randomly select 3 rows and merge them into 1! Setting them too long can make training slower, but could make your chatbot and final finetune much better!
Then set `output_column_name` to the prediction / output column. For the Alpaca dataset dataset, it would be the output column.
We then use the `standardize_sharegpt` function to just make the dataset in a correct format for finetuning! Always call this!
## Vision Fine-tuning
The dataset for fine-tuning a vision or multimodal model also includes image inputs. For example, the [Llama 3.2 Vision Notebook](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/Llama3.2_$11B$-Vision.ipynb#scrollTo=vITh0KVJ10qX) uses a radiography case to show how AI can help medical professionals analyze X-rays, CT scans, and ultrasounds more efficiently.
We'll be using a sampled version of the ROCO radiography dataset. You can access the dataset [here](https://www.google.com/url?q=https%3A%2F%2Fhuggingface.co%2Fdatasets%2Funsloth%2FRadiology_mini). The dataset includes X-rays, CT scans and ultrasounds showcasing medical conditions and diseases. Each image has a caption written by experts describing it. The goal is to finetune a VLM to make it a useful analysis tool for medical professionals.
Let's take a look at the dataset, and check what the 1st example shows:
```
Dataset({
features: ['image', 'image_id', 'caption', 'cui'],
num_rows: 1978
})
```
| Image | Caption |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| | Panoramic radiography shows an osteolytic lesion in the right posterior maxilla with resorption of the floor of the maxillary sinus (arrows). |
To format the dataset, all vision finetuning tasks should be formatted as follows:
```python
[
{ "role": "user",
"content": [{"type": "text", "text": instruction}, {"type": "image", "image": image} ]
},
{ "role": "assistant",
"content": [{"type": "text", "text": answer} ]
},
]
```
We will craft an custom instruction asking the VLM to be an expert radiographer. Notice also instead of just 1 instruction, you can add multiple turns to make it a dynamic conversation.
```notebook-python
instruction = "You are an expert radiographer. Describe accurately what you see in this image."
def convert_to_conversation(sample):
conversation = [
{ "role": "user",
"content" : [
{"type" : "text", "text" : instruction},
{"type" : "image", "image" : sample["image"]} ]
},
{ "role" : "assistant",
"content" : [
{"type" : "text", "text" : sample["caption"]} ]
},
]
return { "messages" : conversation }
pass
```
Let's convert the dataset into the "correct" format for finetuning:
```notebook-python
converted_dataset = [convert_to_conversation(sample) for sample in dataset]
```
The first example is now structured like below:
```notebook-python
converted_dataset[0]
```
{% code overflow="wrap" %}
```
{'messages': [{'role': 'user',
'content': [{'type': 'text',
'text': 'You are an expert radiographer. Describe accurately what you see in this image.'},
{'type': 'image',
'image': }]},
{'role': 'assistant',
'content': [{'type': 'text',
'text': 'Panoramic radiography shows an osteolytic lesion in the right posterior maxilla with resorption of the floor of the maxillary sinus (arrows).'}]}]}
```
{% endcode %}
Before we do any finetuning, maybe the vision model already knows how to analyse the images? Let's check if this is the case!
```notebook-python
FastVisionModel.for_inference(model) # Enable for inference!
image = dataset[0]["image"]
instruction = "You are an expert radiographer. Describe accurately what you see in this image."
messages = [
{"role": "user", "content": [
{"type": "image"},
{"type": "text", "text": instruction}
]}
]
input_text = tokenizer.apply_chat_template(messages, add_generation_prompt = True)
inputs = tokenizer(
image,
input_text,
add_special_tokens = False,
return_tensors = "pt",
).to("cuda")
from transformers import TextStreamer
text_streamer = TextStreamer(tokenizer, skip_prompt = True)
_ = model.generate(**inputs, streamer = text_streamer, max_new_tokens = 128,
use_cache = True, temperature = 1.5, min_p = 0.1)
```
And the result:
```
This radiograph appears to be a panoramic view of the upper and lower dentition, specifically an Orthopantomogram (OPG).
* The panoramic radiograph demonstrates normal dental structures.
* There is an abnormal area on the upper right, represented by an area of radiolucent bone, corresponding to the antrum.
**Key Observations**
* The bone between the left upper teeth is relatively radiopaque.
* There are two large arrows above the image, suggesting the need for a closer examination of this area. One of the arrows is in a left-sided position, and the other is in the right-sided position. However, only
```
For more details, view our dataset section in the [notebook here](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/Llama3.2_$11B$-Vision.ipynb#scrollTo=vITh0KVJ10qX).
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/fine-tuning-llms-guide/datasets-guide.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
