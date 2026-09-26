---
id: collect-250926-servers-hardware/servers-hardware/run-models-on-amd-gpus-with-unsloth-2
title: "Train & run models on AMD GPUs with Unsloth"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Unsloth", "vLLM"]
dates: []
keywords: ["amd", "gpu", "gpus", "benchmarks", "consumer", "datacenter", "fine-tuning", "gguf", "inference", "llama", "llama.cpp", "lora"]
source: docs/RAG/clean4/run models on AMD GPUs with Unsloth.md
source_anchor: ""
source_lines: [94, 251]
sha256: 26b3a09cf212df16fc740d2bf5fa275de529780027980b88a47a9a823988304d
---

# Train & run models on AMD GPUs with Unsloth

### :bar\_chart: AMD Analysis
We brought many of our optimizations we made for non AMD GPUs over to make AMD GPUs shine! Getting AMD training to work on Windows devices and ensuring compatbility across nearly every AMD device was a lot of work but we're happy in the current state. We are still actively working on making AMD GPUs even better! Below are some analysis/benchmarks we conducted on a AMD MI300X.
#### Reinforcement learning: more accurate LoRA and QLoRA
During RL on AMD, Unsloth supports **weight sharing** with vLLM as introduced in [Memory Efficient RL](/docs/get-started/reinforcement-learning-rl-guide/memory-efficient-rl.md). We **halve memory usage** by directly accessing vLLM's allocation of the model.
During LoRA and QLoRA, we also do not do merging and demerging of the LoRA weights like below:
$$
W = W + sAB \\
\text{inference} = XW \\
W = W - sAB
$$
The above causes issues as **IEEE floats are not associative due to rounding**. When W is in BF16, this causes subtle differences when subtracting as done in other RL engines. This means in IEEE land:
$$
W \ne(W + sAB) - (sAB)
$$
And when W is in bfloat16 with few mantissa bits, it **rounds small numbers to zero**. Since A and B are float32 during LoRA, this means W can drift over time after merging & demerging. Unsloth uses vLLM's LoRA path directly, so we avoid this. The base weights are never edited and will never drift.
#### Faster Reinforcement Learning
{% columns %}
{% column %}
We benchmarked other setups with FA2 and vLLM co-located. Both use vLLM for generation, Unsloth pushes more of the step into generation while making training and roughly 2x faster.
All three benchmarks are reproducible: same seeds, same token budgets, and the
loss curves match between the two stacks.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### Faster & more memory efficient training
On Llama-3.1-8B LoRA SFT (batch 2 x grad-accum 4 x 2048 = 16,384 tokens/step, packed), Unsloth trains at 2.07 s/step vs 2.87 s/step for TRL + FA2, and peaks at 18.3 GB versus 24.3 GB. That is 1.39x faster and 1.33x less memory usage, with no change in accuracy. We plan to make this even better!

The memory savings are not just a lower peak, they hold for the whole run.
Sampling allocated VRAM every 0.1s, Unsloth stays flat and low the entire time,
while other setups + FA2 spikes toward 22.8 GB on nearly every step and takes longer to
finish (75s vs 56s for the same 25 steps).
### :sparkles:Supported AMD Hardware
Unsloth’s AMD support focuses on popular consumer, workstation, and datacenter GPUs. Training, Unsloth Studio, and GGUF/llama.cpp inference support can vary by architecture, so the table below separates optimized support from compatibility paths.

Architecture

Series

gfx

Support

Platform

RDNA 4

Radeon™ RX 9000 Series

gfx1200, gfx1201

Full

Windows + WSL + Linux

RDNA 3.5

Ryzen AI 300 / Ryzen AI MAX (Strix Halo)

gfx1150, gfx1151, gfx1152

Full

Windows + WSL + Linux

RDNA 3

Radeon™ RX 7000 Series

gfx1100, gfx1101, gfx1102

Full

Windows + WSL + Linux

RDNA 2

Radeon™ RX 6000 Series

gfx1030, gfx1031, gfx1032, gfx1034

Full

Windows + WSL + Linux

RDNA 1

Radeon™ RX 5000 Series

gfx1010, gfx1012

Vulkan inference

Windows + WSL + Linux

CDNA 4

Instinct™ MI350 Series GPUs

gfx950

Full

Linux only

CDNA 3

Instinct™ MI300 Series GPUs

gfx940, gfx941, gfx942

Full

Linux only

CDNA 2

Instinct™ MI200 Series GPUs

gfx90a

Full

Linux only

CDNA 1

Instinct™ MI100 Series GPUs

gfx908

Full

Linux only

We are actively working on supporting more AMD GPUs, however older ones unfortunately do not have the required hardware support for us to work on.
### :desktop: Guide to using Unsloth on AMD
After installation, you can open Unsloth Studio in your browser. From there, you can run and fine-tune local LLMs on AMD hardware automatically. See below for detailed install instructions:
{% columns %}
{% column width="50%" %} AMD Install GuideInstall Unsloth Studio
Select your model, dataset, and training settings to begin. For example, you can run QLoRA 4-bit fine-tuning on Gemma 4 12B with a context length of \~16k, LoRA rank 16, learning rate 0.0002, and 30 max steps.
{% endcolumn %}
{% column width="50%" %}
{% endcolumn %}
{% endcolumns %}
{% columns %}
{% column %}
Unsloth Studio’s fine-tuning screen is fully customizable, letting you configure your model, dataset, and training parameters. During training, Unsloth tracks progress in real time, including loss, RAM and VRAM usage, as well as support for multi-gpu. Once training starts you will see the progress plotted in real time.
After training, you can visit your history to see past and active training sessions. Select from any of your fine-tuned models to view past training logs:
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}

To set it up, first in terminal after installing Unsloth, type `unsloth studio --secure` . Setup a password so no one unauthorized can access the HTTPS link!
Then in the logs - use the **green Cloudflare link** and enter this on our Phone or any remote device!
And you're in! You can always monitor / cancel the live HTTPS served link by CTRL+C ing the terminal where Unsloth Studio was launched!

