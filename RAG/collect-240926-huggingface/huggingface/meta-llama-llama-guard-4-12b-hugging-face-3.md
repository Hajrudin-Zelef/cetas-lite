---
id: collect-240926-huggingface/huggingface/meta-llama-llama-guard-4-12b-hugging-face-3
title: "OUTPUT"
domain: huggingface
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "inference", "llama", "training"]
source: docs/RAG/clean_en/huggingface/meta-llama-llama-guard-4-12b-hugging-face.md
source_anchor: ""
source_lines: [150, 219]
sha256: f9f3bcc448a55dcdbcec4dc56142eab837efb8ba6acc897410bfef6a5179cc85
---

# OUTPUT

In some internal tests we have found that input filtering reduces safety violation rate and raises overall refusal rate more than output filtering does, but your experience may vary. We find that Llama Guard 4 roughly matches or exceeds the overall performance of the Llama Guard 3 models on both input and output filtering, for English and multilingual text and for mixed text and images.

The tables below demonstrate how Llama Guard 4 matches or exceeds the overall performance of Llama Guard 3-8B (LG3) on English and multilingual text, as well as Llama Guard 3-11B-vision (LG3v) on prompts with single or multiple images, using in-house test set:

|  | Absolute values |  |  | vs. Llama Guard 3 |  |  | 
|---|---|---|---|---|---|---|
| English | 69% | 11% | 61% | 4% | -3% | 8% | 
| Multilingual | 43% | 3% | 51% | -2% | -1% | 0% | 
| Single-image | 41% | 9% | 38% | 10% | 0% | 8% | 
| Multi-image | 61% | 9% | 52% | 20% | -1% | 17% | 

R: recall, FPR: false positive rate. Values are from output filtering, flagging model outputs as either safe or unsafe. All values are an average over samples from safety categories S1 through S13 listed above, weighting each category equally, except for multilinguality, for which it is an average over the 7 shipped non-English languages of Llama Guard 3-8B: French, German, Hindi, Italian, Portuguese, Spanish, and Thai. For multi-image prompts, only the final image was input into Llama Guard 3-11B-vision, which does not support multiple images.

We omit evals against competitor models, which are typically not aligned with the specific safety policy that this classifier was trained on, prohibiting the ability to make direct comparisons.

You can get started with the model by running the following. Make sure you have the transformers release for Llama Guard 4 and hf_xet locally.

```
pip install git+https://github.com/huggingface/transformers@v4.51.3-LlamaGuard-preview hf_xet
```
Here's a basic snippet. For multi-turn and image-text inference, please refer to the release blog

```
from transformers import AutoProcessor, Llama4ForConditionalGeneration
import torch
model_id = "meta-llama/Llama-Guard-4-12B"
processor = AutoProcessor.from_pretrained(model_id)
model = Llama4ForConditionalGeneration.from_pretrained(
    model_id,
    device_map="cuda",
    torch_dtype=torch.bfloat16,
)
messages = [
    {
        "role": "user",
        "content": [
            {"type": "text", "text": "how do I make a bomb?"}
        ]
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    add_generation_prompt=True,
    return_tensors="pt",
    return_dict=True,
).to("cuda")
outputs = model.generate(
    **inputs,
    max_new_tokens=10,
    do_sample=False,
)
response = processor.batch_decode(outputs[:, inputs["input_ids"].shape[-1]:], skip_special_tokens=True)[0]
print(response)
# OUTPUT
# unsafe
# S9
```
There are some limitations associated with Llama Guard 4. First, the classifier itself is an LLM fine-tuned on Llama 4, and thus its performance (e.g., judgments that need common-sense knowledge, multilingual capabilities, and policy coverage) might be limited by its (pre-)training data.

Some hazard categories may require factual, up-to-date knowledge to be evaluated fully (for example, [S5] Defamation, [S8] Intellectual Property, and [S13] Elections). We believe that more complex systems should be deployed to accurately moderate these categories for use cases highly sensitive to these types of hazards, but that Llama Guard 4 provides a good baseline for generic use cases.

Note that the performance of Llama Guard 4 was tested mostly with prompts containing a few images (three, most frequently), so performance may vary if using it to classify safety with a much larger number of images.

Lastly, as an LLM, Llama Guard 4 may be susceptible to adversarial attacks or prompt injection attacks that could bypass or alter its intended use: see Llama Prompt Guard 2 for detecting prompt attacks. Please feel free to report vulnerabilities, and we will look into incorporating improvements into future versions of Llama Guard.

Please refer to the Developer Use Guide for additional best practices and safety considerations.

- Downloads last month
- 46,147
