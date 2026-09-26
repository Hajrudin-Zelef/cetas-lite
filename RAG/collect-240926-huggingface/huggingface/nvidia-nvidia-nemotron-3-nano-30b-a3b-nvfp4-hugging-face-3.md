---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-3
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "OpenAI"]
dates: ["2025-05-01"]
keywords: ["agent", "alignment", "fine-tuning", "license", "nvfp4", "nvidia", "open source", "reasoning", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [271, 368]
sha256: c277c426c442ddcd4623ee2641b89a0fff92d2209c29b4092b1ad03dc0efa21d
---

# Load tokenizer and model

```
from typing import Any, Dict, List
import openai
from transformers import AutoTokenizer
class ThinkingBudgetClient:
   def __init__(self, base_url: str, api_key: str, tokenizer_name_or_path: str):
       self.base_url = base_url
       self.api_key = api_key
       self.tokenizer = AutoTokenizer.from_pretrained(tokenizer_name_or_path)
       self.client = openai.OpenAI(base_url=self.base_url, api_key=self.api_key)
   def chat_completion(       self,
       model: str,
       messages: List[Dict[str, Any]],
       reasoning_budget: int = 512,
       max_tokens: int = 1024,
       **kwargs,
   ) -> Dict[str, Any]:
       assert (
           max_tokens > reasoning_budget
       ), f"thinking budget must be smaller than maximum new tokens. Given {max_tokens=} and {reasoning_budget=}"
       # 1. first call chat completion to get reasoning content
       response = self.client.chat.completions.create(
           model=model, messages=messages, max_tokens=reasoning_budget, **kwargs
       )
       content = response.choices[0].message.content
       reasoning_content = content
       if not "</think>" in reasoning_content:
           # reasoning content is too long, closed with a period (.)
           reasoning_content = f"{reasoning_content}.\n</think>\n\n"
       reasoning_tokens_len = len(
           self.tokenizer.encode(reasoning_content, add_special_tokens=False)
       )
       remaining_tokens = max_tokens - reasoning_tokens_len
       assert (
           remaining_tokens > 0
       ), f"remaining tokens must be positive. Given {remaining_tokens=}. Increase the max_tokens or lower the reasoning_budget."
       # 2. append reasoning content to messages and call completion
       messages.append({"role": "assistant", "content": reasoning_content})
       prompt = self.tokenizer.apply_chat_template(
           messages,
           tokenize=False,
           continue_final_message=True,
       )
       response = self.client.completions.create(
           model=model, prompt=prompt, max_tokens=remaining_tokens, **kwargs
       )
       response_data = {
           "reasoning_content": reasoning_content.strip().strip("</think>").strip(),
           "content": response.choices[0].text,
           "finish_reason": response.choices[0].finish_reason,
       }
       return response_data
```
Calling the server with a budget (Restricted to 32 tokens here as an example)

```
tokenizer_name_or_path = "nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4"
client = ThinkingBudgetClient(
   base_url="http://localhost:8000/v1",  # Nemotron 3 Nano deployed in thinking mode
   api_key="EMPTY",
   tokenizer_name_or_path=tokenizer_name_or_path,
)
result = client.chat_completion(
   model="nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4",
   messages=[
       {"role": "system", "content": "You are a helpful assistant. /think"},
       {"role": "user", "content": "What is 2+2?"},
   ],
   reasoning_budget=32,
   max_tokens=512,
   temperature=1.0,
   top_p=1.0,
)
print(result)
```
You should see output similar to the following:

```
{'reasoning_content': "Okay, the user asked, What is 2+2? Let me think. Well, 2 plus 2 equals 4. That's a basic.", 'content': '2 + 2 equals **4**.\n', 'finish_reason': 'stop'}
```
- v1.0 NVFP4

**Data Modality:**  Text**The total size:**  10,648,823,153,919 Tokens**Total number of datasets:** 141**Dataset partition:** *Training [100%], testing [0%], validation [0%]***Time period for training data collection:** 2013 to May 1, 2025**Time period for testing data collection:** 2013 to May 1, 2025**Time period for validation data collection:** 2013 to May 1, 2025**Data Collection Method by dataset:** Hybrid: Automated, Human, Synthetic**Labeling Method by dataset: Hybrid:** Automated, Human, Synthetic

NVIDIA-Nemotron-3-Nano-30B-A3B-BF16 is pre-trained on a large corpus of high-quality curated and synthetically-generated data. It is trained in the English language, as well as 19 other languages and 43 programming languages. Our sources cover a variety of document types such as: webpages, dialogue, articles, and other written materials. The corpus spans domains including legal, math, science, finance, and more. We also include a small portion of question-answering, and alignment style data to improve model accuracy. The model was trained for approximately 25 trillion tokens.

The post-training corpus for NVIDIA-Nemotron-3-Nano-30B-A3B-BF16 of high-quality curated and synthetically-generated data. Primary languages used for post-training include English, German, Spanish, French, Italian, and Japanese.

These datasets, such as FinePDFs, EssentialWeb, HotpotQA, SQuAD, and HelpSteer3, do not collectively or exhaustively represent all demographic groups (and proportionally therein). For instance, these datasets do not contain explicit mentions of demographic classes such as age, gender, or ethnicity in 64-99% of samples, depending on the source. In the subset where such terms are present, document-based datasets (FinePDFs and EssentialWeb) contain representational skews, such as references to "male" outnumbering those to "female", and mentions of "White" as the most frequent among ethnic identifiers (comprising 43-44% of ethnicity mentions). To mitigate these imbalances, we recommend considering evaluation techniques such as bias audits, fine-tuning with demographically balanced datasets, and mitigation strategies like counterfactual data augmentation to align with the desired model behavior. This evaluation used a 3,000-sample subset per dataset, identified as the optimal threshold for maximizing embedder accuracy.

During post-training, we generate synthetic data by distilling trajectories, solutions, and translations from strong teacher models and agent systems, often grounded in real tasks or documents and aggressively filtered for quality. For math, code, and science, we start from curated problem sets and use open source permissive models such as GPT-OSS-120B to produce step-by-step reasoning traces, candidate solutions, best-of-n selection traces, and verified CUDA kernels. For long-context and science, we build synthetic QA and reasoning data by retrieving passages from long documents, generating MCQ/OpenQA questions and answers, and paraphrasing them via into multiple prompt/response formats to ensure diversity. Across all pipelines we stack automated verification—compilers, numerical checks, language identification—to ensure our data is high quality.

For all domains, we apply a unified data filtering pipeline to ensure that only high-quality, license-compliant, and verifiable samples are used for post-training. We first discard malformed examples using structural checks (e.g., missing tool definitions when tool calls are present). We then aggressively filter reasoning traces exhibiting pathological repetition, such as repeated n-grams within a sliding window or across the entire trajectory, which we found to be a strong indicator of malformed or low-quality reasoning. Finally, based on internal audits of synthetically generated datasets, we observed that some teacher models occasionally produce reasoning traces and final responses that implicitly align with specific political entities or promote nationalistic narratives. To mitigate this, we apply targeted keyword- and regex-based filters and remove all trajectories matching such behavior.

Alongside the model, we release our final pre-training and post-training data, as outlined in this section. For ease of analysis, there is a sample set that is ungated. For all remaining code, math and multilingual data, gating and approval is required, and the dataset is permissively licensed for model training purposes.

More details on the datasets and synthetic data generation methods can be found in the technical report NVIDIA Nemotron 3 Nano. Data Designer is one of the libraries used to prepare the pre and post training datasets.

