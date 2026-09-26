---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-guidedatasets-guide-3
title: "Datasets Guide"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "fine-tuning", "inference", "research"]
source: docs/RAG/clean4/fine-tuning-llms-guidedatasets-guide.md
source_anchor: ""
source_lines: [135, 240]
sha256: a302bf3d71e769153996acf787cdd8efadff6518b24d0f37196143eaff3c0e47
---

# Datasets Guide

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
