---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-4
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agent", "alignment", "fine-tuning", "nvidia", "open source", "reasoning", "sglang", "tensorrt", "tensorrt-llm", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [452, 575]
sha256: e649791f004c56568eaa5ba855ba163bf05f234b78388f89ab0ee87a0f0d555a
---

# Set the IP for the head node in RAY_HEAD_IP

```
{
    "$schema": "https://opencode.ai/config.json",
    "model": "local/nvidia-nemotron-3-ultra",
    "provider": {
        "local": {
            "npm": "@ai-sdk/openai-compatible",
            "name": "local_backend",
            "options": {
                "baseURL": "http://localhost:8000/v1",
                "apiKey": "EMPTY"
            },
            "models": {
                "nvidia-nemotron-3-ultra": {
                    "name": "nvidia/nemotron-3-ultra",
                    "limit": {
                        "context": 1000000,
                        "output": 32768
                    }
                }
            }
        }
    },
    "agent": {
        "build": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        },
        "plan": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        }
    }
}
```
All backends above default to port `8000`, so the `baseURL` works as-is for vLLM, SGLang, and TensorRT-LLM.

To learn more about other supported agent scaffolds - check out this resource

##  **Advanced: Budget-Controlled Reasoning**

Set a hard token ceiling on the reasoning trace using `reasoning_budget`. The model will attempt to close the trace at the next newline before the budget is hit; if none is found within 500 tokens it closes abruptly at `reasoning_budget + 500`.

```
from typing import Any, Dict, List
import openai
from transformers import AutoTokenizer
class ThinkingBudgetClient:
    def __init__(self, base_url: str, api_key: str, tokenizer_name_or_path: str):
        self.tokenizer = AutoTokenizer.from_pretrained(tokenizer_name_or_path)
        self.client = openai.OpenAI(base_url=base_url, api_key=api_key)
    def chat_completion(        self,
        model: str,
        messages: List[Dict[str, Any]],
        reasoning_budget: int = 512,
        max_tokens: int = 1024,
        **kwargs,
    ) -> Dict[str, Any]:
        assert max_tokens > reasoning_budget, (
            f"reasoning_budget must be less than max_tokens. "
            f"Got {max_tokens=} and {reasoning_budget=}"
        )
        # Step 1: generate the reasoning trace up to the budget
        response = self.client.chat.completions.create(
            model=model, messages=messages, max_tokens=reasoning_budget, **kwargs
        )
        reasoning_content = response.choices[0].message.content
        if "</think>" not in reasoning_content:
            reasoning_content = f"{reasoning_content}.\n\n</think>\n"
        reasoning_tokens_len = len(
            self.tokenizer.encode(reasoning_content, add_special_tokens=False)
        )
        remaining_tokens = max_tokens - reasoning_tokens_len
        assert remaining_tokens > 0, (
            f"No tokens remaining for response ({remaining_tokens=}). "
            "Increase max_tokens or lower reasoning_budget."
        )
        # Step 2: continue from the closed reasoning trace
        messages.append({"role": "assistant", "content": reasoning_content})
        prompt = self.tokenizer.apply_chat_template(
            messages, tokenize=False, continue_final_message=True
        )
        response = self.client.completions.create(
            model=model, prompt=prompt, max_tokens=remaining_tokens, **kwargs
        )
        return {
            "reasoning_content": reasoning_content.strip().strip("</think>").strip(),
            "content": response.choices[0].text,
            "finish_reason": response.choices[0].finish_reason,
        }
```
**Example usage** (32-token reasoning budget):

```
client = ThinkingBudgetClient(
    base_url="http://localhost:8000/v1",
    api_key="EMPTY",
    tokenizer_name_or_path="nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16",
)
result = client.chat_completion(
    model="nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16",
    messages=[
        {"role": "system", "content": "You are a helpful assistant. /think"},
        {"role": "user", "content": "What is 2+2?"},
    ],
    reasoning_budget=32,
    max_tokens=512,
    temperature=1.0,
    top_p=0.95,
)
print(result)
```
**Data Modality:** Text**The total size:** 53.8 TiB (14.8 trillion tokens)**Total number of datasets:** 226**Dataset partition:** *Training [100%], testing [0%], validation [0%]***Time period for training data collection:** 2013 to 2026**Time period for testing data collection:** 2013 to 2026**Time period for validation data collection:** 2013 to 2026**Data Collection Method by dataset:** Hybrid: Automated, Human, Synthetic**Labeling Method by dataset:** Hybrid: Automated, Human, Synthetic  

NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 is pre-trained on a large corpus of high-quality curated and synthetically-generated data. It is trained in the English language, as well as 11 other languages and 43 programming languages. Our sources cover a variety of document types such as: webpages, dialogue, articles, and other written materials. The corpus spans domains including legal, math, science, finance, and more. We also include a small portion of question-answering, and alignment style data to improve model accuracy. The model was pre-trained for approximately 20 trillion tokens.

The post-training corpus for NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 consists of high-quality curated and synthetically-generated data. Primary languages used for post-training include English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, and Chinese.

These datasets, such as FinePDFs, EssentialWeb, HotpotQA, SQuAD, and HelpSteer3, do not collectively or exhaustively represent all demographic groups (and proportionally therein). For instance, these datasets do not contain explicit mentions of demographic classes such as age, gender, or ethnicity in 64-99% of samples, depending on the source. In the subset where such terms are present, document-based datasets (FinePDFs and EssentialWeb) contain representational skews, such as references to "male" outnumbering those to "female", and mentions of "White" as the most frequent among ethnic identifiers (comprising 43-44% of ethnicity mentions). To mitigate these imbalances, we recommend considering evaluation techniques such as bias audits, fine-tuning with demographically balanced datasets, and mitigation strategies like counterfactual data augmentation to align with the desired model behavior. This evaluation used a 3,000-sample subset per dataset, identified as the optimal threshold for maximizing embedder accuracy.

During post-training, we generate synthetic data by distilling trajectories, solutions, and translations from strong teacher models and agent systems, often grounded in real tasks or documents and aggressively filtered for quality. For math, code, and science, we start from curated problem sets and use open source permissive models such as GPT-OSS-120B to produce step-by-step reasoning traces, candidate solutions, best-of-n selection traces, and verified CUDA kernels. For long-context and science, we build synthetic QA and reasoning data by retrieving passages from long documents, generating MCQ/OpenQA questions and answers, and paraphrasing them into multiple prompt/response formats to ensure diversity. Across all pipelines we stack automated verification—compilers, numerical checks, language identification—to ensure our data is high quality.

