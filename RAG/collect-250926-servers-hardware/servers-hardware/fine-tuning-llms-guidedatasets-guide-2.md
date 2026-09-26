---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-guidedatasets-guide-2
title: "Datasets Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["DeepSeek", "Google", "Hugging Face", "OpenAI", "Unsloth"]
dates: []
keywords: ["chatgpt", "deepseek", "fine-tuning", "llama", "multimodal", "reasoning", "research", "training"]
source: docs/RAG/clean4/fine-tuning-llms-guidedatasets-guide.md
source_anchor: ""
source_lines: [71, 134]
sha256: 6c903b89fbc1953399a188f35a4b251f2844a824da2b3c6b9cfceafa7602218f
---

# Datasets Guide

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
