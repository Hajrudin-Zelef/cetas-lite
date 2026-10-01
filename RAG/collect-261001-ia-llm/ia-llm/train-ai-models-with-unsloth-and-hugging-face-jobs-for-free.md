---
id: collect-261001-ia-llm/ia-llm/train-ai-models-with-unsloth-and-hugging-face-jobs-for-free
title: "mac or linux"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "Unsloth"]
dates: []
keywords: ["agent", "agents", "claude", "cost", "fine-tuning", "gpu", "gpus", "lora", "memory", "pricing", "qwen", "training"]
source: docs/RAG/collect-261001-ia-llm/train-ai-models-with-unsloth-and-hugging-face-jobs-for-free.md
source_anchor: ""
source_lines: [1, 145]
sha256: c8f1e97644890309783bbb277eae97d6fbda8d024c6b9613428432a72bb8dff8
---

# mac or linux

`LiquidAI/LFM2.5-1.2B-Instruct` ) through coding agents like Claude Code and Codex. Unsloth provides ~2x faster training and ~60% less VRAM usage compared to standard methods, so training small models can cost just a few dollars.
Why a small model? Small language models like LFM2.5-1.2B-Instruct are ideal candidates for fine-tuning. They are cheap to train, fast to iterate on, and increasingly competitive with much larger models on focused tasks. LFM2.5-1.2B-Instruct runs under 1GB of memory and is optimized for on-device deployment, so what you fine-tune can be served on CPUs, phones, and laptops.

We are giving away free credits to fine-tune models on Hugging Face Jobs. Join the Unsloth Jobs Explorers organization to claim your free credits and one-month Pro subscription.

- A Hugging Face account (required for HF Jobs)
- Billing setup (for verification, you can monitor your usage and manage your billing in your billing page).
- A Hugging Face token with write permissions
- (optional) A coding agent (`Open Code` ,`Claude Code` , or`Codex` )

If you want to train a model using HF Jobs and Unsloth, you can simply use the `hf jobs` CLI to submit a job.

First, you need to install the `hf` CLI. You can do this by running the following command:

```
# mac or linux
curl -LsSf https://hf.co/cli/install.sh | bash
```
Next you can run the following command to submit a job:

```
hf jobs uv run https://huggingface.co/datasets/unsloth/jobs/resolve/main/sft-lfm2.5.py \
    --flavor a10g-small  \
    --secrets HF_TOKEN  \
    --timeout 4h \
    --dataset mlabonne/FineTome-100k \
    --num-epochs 1 \
    --eval-split 0.2 \
    --output-repo your-username/lfm-finetuned
```
Check out the training script and Hugging Face Jobs documentation for more details.

Hugging Face model training skill lowers barrier of entry to train a model by simply prompting. First, install the skill with your coding agent.

Claude Code discovers skills through its plugin system, so we need to install the Hugging Face skills first. To do so:

1. Add the marketplace:

```
/plugin marketplace add huggingface/skills
```
1. Browse available skills in the `Discover` tab:

```
/plugin
```
1. Install the model trainer skill:

```
/plugin install hugging-face-model-trainer@huggingface-skills
```
For more details, see the documentation on using the hub with skills or the Claude Code Skills docs.

Codex discovers skills through `AGENTS.md` files and `.agents/skills/` directories.

Install individual skills with `$skill-installer`:

```
$skill-installer install https://github.com/huggingface/skills/tree/main/skills/huggingface-llm-trainer
```
For more details, see the Codex Skills docs and the AGENTS.md guide.

A generic install method is simply to clone the skills repository and copy the skill to your agent's skills directory.

```
git clone https://github.com/huggingface/skills.git
mkdir -p ~/.agents/skills && cp -R skills/skills/hugging-face-model-trainer ~/.agents/skills/
```
Once the skill is installed, ask your coding agent to train a model:

```
Train LiquidAI/LFM2.5-1.2B-Instruct on mlabonne/FineTome-100k using Unsloth on HF Jobs
```
The agent will generate a training script based on an example in the skill, submit the training to HF Jobs, and provide a monitoring link via Trackio.

Training jobs run on Hugging Face Jobs, fully managed cloud GPUs. The agent:

1. Generates a UV script with inline dependencies
2. Submits it to HF Jobs via the `hf` CLI
3. Reports the job ID and monitoring URL
4. Pushes the trained model to your Hugging Face Hub repository

The skill generates scripts like this based on the example in the skill.

```
# /// script
# dependencies = ["unsloth", "trl>=0.12.0", "datasets", "trackio"]
# ///
from unsloth import FastLanguageModel
from trl import SFTTrainer, SFTConfig
from datasets import load_dataset
model, tokenizer = FastLanguageModel.from_pretrained(
    "LiquidAI/LFM2.5-1.2B-Instruct",
    load_in_4bit=True,
    max_seq_length=2048,
)
model = FastLanguageModel.get_peft_model(
    model,
    r=16,
    lora_alpha=32,
    lora_dropout=0,
    target_modules=[
        "q_proj",
        "k_proj",
        "v_proj",
        "out_proj",
        "in_proj",
        "w1",
        "w2",
        "w3",
    ],
)
dataset = load_dataset("trl-lib/Capybara", split="train")
trainer = SFTTrainer(
    model=model,
    tokenizer=tokenizer,
    train_dataset=dataset,
    args=SFTConfig(
        output_dir="./output",
        push_to_hub=True,
        hub_model_id="username/my-model",
        per_device_train_batch_size=4,
        gradient_accumulation_steps=4,
        num_train_epochs=1,
        learning_rate=2e-4,
        report_to="trackio",
    ),
)
trainer.train()
trainer.push_to_hub()
```
| Model Size | Recommended GPU | Approx Cost/hr | 
|---|---|---|
| <1B params | `t4-small` | ~$0.40 | 
| 1-3B params | `t4-medium` | ~$0.60 | 
| 3-7B params | `a10g-small` | ~$1.00 | 
| 7-13B params | `a10g-large` | ~$3.00 | 

For a full overview of Hugging Face Spaces pricing, check out the guide here.

- Be specific about the model and dataset to use, and include Hub IDs (for example, `Qwen/Qwen2.5-0.5B` and`trl-lib/Capybara` ). Agents will search for and validate those combinations.
- Mention Unsloth explicitly if you want it used. Otherwise, the agent will choose a framework based on the model and budget.
- Ask for cost estimates before launching large jobs.
- Request Trackio monitoring for real-time loss curves.
- Check job status by asking the agent to inspect logs after submission.
