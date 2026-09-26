---
id: collect-260926-rattrapage/rattrapage/fine-tuning-llms-on-intel-gpus-with-unsloth-1
title: "Fine-tuning LLMs on Intel GPUs with Unsloth"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Intel", "Unsloth", "vLLM"]
dates: []
keywords: ["fine-tuning", "gpu", "gpus", "intel", "consumer", "inference", "llama", "lora", "memory", "mistral", "qlora", "reasoning"]
source: docs/RAG/lot-rattrapage/fine-tuning/Fine-tuning LLMs on Intel GPUs with Unsloth.md
source_anchor: ""
source_lines: [1, 195]
sha256: b043e43216e6969d4bd20187d1ba8f922adc2499710fd3bde45236330e6a8957
---

# Fine-tuning LLMs on Intel GPUs with Unsloth

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/intel.md).
# Fine-tuning LLMs on Intel GPUs with Unsloth
Learn how to train and fine-tune large language models on Intel GPUs.
You can now fine-tune LLMs on your local Intel device with Unsloth! Read our guide on exactly how to get started with training your own custom model.
Before you begin, make sure you have:
* **Intel GPU:** Data Center GPU Max Series, Arc Series, or Intel Ultra AIPC
* **OS:** Linux (Ubuntu 22.04+ recommended) or Windows 11 (recommended)
* **Windows only:** Install Intel oneAPI Base Toolkit 2025.2.1 (select version 2025.2.1)
* **Intel Graphics driver:** Latest recommended driver for Windows/Linux
* **Python:** 3.10+
### Build Unsloth with Intel Support
{% stepper %}
{% step %}
#### Create a new conda environment (Optional)
```bash
conda create -n unsloth-xpu python==3.10
conda activate unsloth-xpu
```
{% endstep %}
{% step %}
#### Install Unsloth
```bash
git clone https://github.com/unslothai/unsloth.git
cd unsloth
pip install .[intel-gpu-torch290]
```
{% hint style="info" %}
Linux Only: Install [vLLM](/docs/basics/inference-and-deployment/vllm-guide.md) (Optional)\
You can also install vLLM for [inference](/docs/basics/inference-and-deployment.md) and [RL](/docs/get-started/reinforcement-learning-rl-guide.md). Please follow [vLLM's guide](https://docs.vllm.ai/en/latest/getting_started/installation/gpu/#intel-xpu).
{% endhint %}
{% endstep %}
{% step %}
#### Verify your environments
```python
import torch
print(f"PyTorch version: {torch.__version__}")
print(f"XPU available: {torch.xpu.is_available()}")
print(f"XPU device count: {torch.xpu.device_count()}")
print(f"XPU device name: {torch.xpu.get_device_name(0)}")
```
{% endstep %}
{% step %}
#### Start fine-tuning.
You can directly use our Unsloth [notebooks](/docs/get-started/unsloth-notebooks.md) or view our dedicated [fine-tuning](/docs/get-started/fine-tuning-llms-guide.md) or [reinforcement learning](/docs/get-started/reinforcement-learning-rl-guide.md) guides.
{% endstep %}
{% endstepper %}
### Windows Only - Runtime Configurations
In Command Prompt with Administrator privilege, enable long path support in the Windows registry:
```bash
powershell -Command "Set-ItemProperty -Path "HKLM:\\SYSTEM\\CurrentControlSet\\Control\\FileSystem" -Name "LongPathsEnabled" -Value 1
```
This command only needs to be set once on a single machine. It does not need to be configured before each run. Then:
1. Download level-zero-win-sdk-1.20.2.zip from [GitHub](https://github.com/oneapi-src/level-zero/releases/tag/v1.20.2)
2. Unzip the level-zero-win-sdk-1.20.2.zip
3. In Command Prompt, under conda environment unsloth-xpu:
```bash
call "C:\Program Files (x86)\Intel\oneAPI\setvars.bat" -
set ZE_PATH=path\to\the\unzipped\level-zero-win-sdk-1.20.2
```
### Example 1: QLoRA Fine-tuning with SFT
This example demonstrates how to fine-tune a Qwen3-32B model using 4-bit QLoRA on an Intel GPU. QLoRA significantly reduces memory requirements, making it possible to fine-tune large models on consumer-grade hardware.
{% code expandable="true" %}
```python
from unsloth import FastLanguageModel, FastModel
from trl import SFTTrainer, SFTConfig
from datasets import load_dataset
max_seq_length = 2048 # Supports RoPE Scaling internally, so choose any!
# Get LAION dataset
url = "https://huggingface.co/datasets/laion/OIG/resolve/main/unified_chip2.jsonl"
dataset = load_dataset("json", data_files = {"train" :
url}, split = "train")
# 4bit pre quantized models we support for fast downloading + no OOMs.
fourbit_models = [
"unsloth/Qwen3-32B-bnb-4bit",
"unsloth/Qwen3-14B-bnb-4bit",
"unsloth/Qwen3-8B-bnb-4bit",
"unsloth/Qwen3-4B-bnb-4bit",
"unsloth/Qwen3-1.7B-bnb-4bit",
"unsloth/Qwen3-0.6B-bnb-4bit",
# "unsloth/Qwen2.5-32B-bnb-4bit",
# "unsloth/Qwen2.5-14B-bnb-4bit",
# "unsloth/Qwen2.5-7B-bnb-4bit",
# "unsloth/Qwen2.5-3B-bnb-4bit",
# "unsloth/Qwen2.5-1.5B-bnb-4bit",
# "unsloth/Qwen2.5-0.5B-bnb-4bit",
# "unsloth/Llama-3.2-3B-bnb-4bit",
# "unsloth/Llama-3.2-1B-bnb-4bit",
# "unsloth/Llama-3.1-8B-bnb-4bit",
# "unsloth/Llama-3.1-70B-bnb-4bit",
# "unsloth/mistral-7b-bnb-4bit",
# "unsloth/Phi-4",
# "unsloth/Phi-3.5-mini-instruct",
# "unsloth/Phi-3-medium-4k-instruct",
# "unsloth/Phi-3-mini-4k-instruct",
# "unsloth/gemma-2-9b-bnb-4bit",
# "unsloth/gemma-2-27b-bnb-4bit",
] # More models at https://huggingface.co/unsloth
model, tokenizer = FastLanguageModel.from_pretrained(
model_name = "unsloth/Qwen3-32B-bnb-4bit",
max_seq_length = max_seq_length,
load_in_4bit = True,
# token = "hf_...", # use one if using gated models like meta-llama/Llama-2-7b-hf
)
model = FastLanguageModel.get_peft_model(
model,
r = 16, # Choose any number > 0 ! Suggested 8, 16, 32, 64, 128
target_modules = ["q_proj", "k_proj", "v_proj", "o_proj",
"gate_proj", "up_proj", "down_proj",
],
lora_alpha = 16,
lora_dropout = 0, # Supports any, but = 0 is optimized
bias = "none", # Supports any, but = "none" is optimized
use_gradient_checkpointing = "unsloth", # True or "unsloth" for very long context
random_state = 3407,
use_rslora = False, # We support rank stabilized LoRA
loftq_config = None, # And LoftQ
)
trainer = SFTTrainer(
model = model,
tokenizer = tokenizer,
train_dataset = dataset,
dataset_text_field = "text",
max_seq_length = max_seq_length,
dataset_num_proc = 1, # Recommended on Windows
packing = False, # Can make training 5x faster for short sequences.
args = SFTConfig(
per_device_train_batch_size = 2,
gradient_accumulation_steps = 4,
warmup_steps = 5,
max_steps = 60,
learning_rate = 2e-4,
logging_steps = 1,
optim = "adamw_8bit",
weight_decay = 0.01,
lr_scheduler_type = "linear",
seed = 3407,
dataset_num_proc=1, # Recommended on Windows
),
)
trainer.train()
```
{% endcode %}
### Example 2: Reinforcement Learning GRPO
GRPO is a [reinforcement learning](/docs/get-started/reinforcement-learning-rl-guide.md) technique for aligning language models with human preferences. This example shows how to train a model to follow a specific XML output format using multiple reward functions.
#### What is GRPO?
GRPO improves upon traditional RLHF by:
* Using group-based normalization for more stable training
* Supporting multiple reward functions for multi-objective optimization
* Being more memory efficient than PPO
{% code expandable="true" %}
```python
from unsloth import FastLanguageModel
import re
from trl import GRPOConfig, GRPOTrainer
from datasets import load_dataset, Dataset
max_seq_length = 1024  # Can increase for longer reasoning traces
lora_rank = 32  # Larger rank = smarter, but slower
max_prompt_length = 256
# Load and prep dataset
SYSTEM_PROMPT = """
Respond in the following format:
...
...
"""
XML_COT_FORMAT = """\
{reasoning}
{answer}
"""
def extract_xml_answer(text: str) -> str:
    answer = text.split("")[-1]
    answer = answer.split("")[0]
    return answer.strip()
def extract_hash_answer(text: str) -> str | None:
    if "####" not in text:
        return None
    return text.split("####")[1].strip()
# uncomment middle messages for 1-shot prompting
def get_gsm8k_questions(split: str = "train") -> Dataset:
    data = load_dataset("openai/gsm8k", "main")[split]  # type: ignore
    data = data.map(
        lambda x: {  # type: ignore
            "prompt": [
                {"role": "system", "content": SYSTEM_PROMPT},
                {"role": "user", "content": x["question"]},
            ],
            "answer": extract_hash_answer(x["answer"]),
        }
    )  # type: ignore
    return data  # type: ignore
# Reward functions
def correctness_reward_func(prompts, completions, answer, **kwargs) -> list[float]:
    responses = [completion[0]["content"] for completion in completions]
    q = prompts[0][-1]["content"]
    extracted_responses = [extract_xml_answer(r) for r in responses]
    print(
