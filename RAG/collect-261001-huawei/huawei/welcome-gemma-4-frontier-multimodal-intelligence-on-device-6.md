---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-6
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: ["China", "Google", "Hugging Face", "Nvidia", "SGLang", "Unsloth"]
dates: []
keywords: ["accelerator", "benchmark", "benchmarks", "decode", "fine-tuning", "llama", "llama.cpp", "mistral", "moe", "multimodal", "nvidia", "reasoning"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [511, 628]
sha256: 74b56e0b29c77cb430e00f35989a90e8f39df8b7253fc4d659d7074ac27edab6
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

```
from transformers import DiffusionGemmaForBlockDiffusion, AutoProcessor
MODEL_ID = "google/diffusiongemma-26B-A4B-it"
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = DiffusionGemmaForBlockDiffusion.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto",
)
message = [{"role": "user", "content": "Why is the sky blue?"}]
input_ids = processor.apply_chat_template(
    message,
    tokenize=True,
    add_generation_prompt=True,
    return_dict=True,
    return_tensors="pt",
).to(model.device)
output = model.generate(**input_ids, max_new_tokens=512)
text = processor.decode(output[0], skip_special_tokens=False)
```
For best results, Google recommends the **Entropy-Bounded (EB) sampler** with adaptive stopping (up to 48 denoising steps, a temperature schedule decaying from 0.8 → 0.4, and an entropy bound of 0.1 for token selection). As with Gemma 4, place image content **before** text in your prompt, and toggle reasoning with the `<|think|>` control token.

Gemma 4 models are ideal for fine-tuning in your favorite tools and platforms and at any budget.

Gemma 4 is fully supported for fine-tuning with TRL. To celebrate, TRL has been upgraded with support for multimodal tool responses when interacting with environments, meaning models can now receive images back from tools during training, not just text.

To showcase this, we've built an example training script where Gemma 4 learns to drive in the CARLA simulator. The model sees the road through a camera, decides what to do and learns from the outcome. After training, it consistently changes lanes to avoid pedestrians. The same approach works for any task where a model needs to see and act: robotics, web browsing, or other interactive environments.

Get started:

```
# pip install git+https://github.com/huggingface/trl.git
python examples/grpo_carla/carla_vlm_gemma.py \
    --env-urls https://sergiopaniego-carla-env.hf.space \
            https://sergiopaniego-carla-env-2.hf.space \
    --model google/gemma-4-E2B-it
```
Find the example here.

Additionally, we have prepared an example on how to fine-tune Gemma 4 with TRL on Vertex AI using SFT, to showcase how to extend the function calling capabilities, whilst freezing both the vision and audio towers. The examples include how to build a custom Docker container with latest Transformers, TRL, etc. with CUDA support on Google Cloud, and how to run it via Vertex AI Serverless Training Jobs.

```
# pip install google-cloud-aiplatform --upgrade --quiet
from google.cloud import aiplatform
aiplatform.init(
    project="<PROJECT_ID>",
    location="<LOCATION>",
    staging_bucket="<BUCKET_URI>",
)
job = aiplatform.CustomContainerTrainingJob(
    display_name="gemma-4-fine-tuning",
    container_uri="<CONTAINER_URI>",
    command=["python", "/gcs/gemma-4-fine-tuning/train.py"],
)
job = job.submit(
    replica_count=1,
    machine_type="a3-highgpu-1g",
    accelerator_type="NVIDIA_H100_80GB",
    accelerator_count=1,
    base_output_dir="<BUCKET_URI>/output-dir",
    environment_variables={
        "MODEL_ID": "google/gemma-4-E2B-it",
        "HF_TOKEN": <HF_TOKEN>,
    },
    boot_disk_size_gb=500,
)
```
You can find the complete example in the "Hugging Face on Google Cloud" docs at https://hf.co/docs/google-cloud/examples/vertex-ai-notebooks-fine-tune-gemma-4.

If you want to fine tune and run a Gemma 4 model in a UI, try out Unsloth Studio. It runs locally or on Google Colab. First, install and start the app:

```
# install unsloth studio on MacOS, Linux, WSL
curl -fsSL https://unsloth.ai/install.sh | sh
# install unsloth studio on Windows
irm https://unsloth.ai/install.ps1 | iex
# launch unsloth studio
unsloth studio -H 0.0.0.0 -p 8888
# Search for for a Gemma 4 model like google/gemma-4-E2B-it
```
Then select any of the Gemma 4 models from the hub.

We have shipped demos for you to try different Gemma 4 models. We include demos based on the transformers implementation for E4B, 12B Unified, 26B/A4B MoE, and 31B dense models. There's also a WebGPU demo with transformers.js 🚀

Gemma 4 models demonstrate exceptional performance across diverse benchmarks, from reasoning and coding to vision and long-context tasks. The graph below shows model performance vs size, with Gemma 4 models forming an impressive Pareto frontier:

Source: Google (blog.google)

Here are detailed benchmark results for the instruction-tuned models:

| Benchmark | Gemma 4 31B | Gemma 4 26B A4B | Gemma 4 12B Unified | Gemma 4 E4B | Gemma 4 E2B | Gemma 3 27B (no think) | 
|---|---|---|---|---|---|---|
| **Reasoning & Knowledge** |  |  |  |  |  |  | 
| MMLU Pro | 85.2% | 82.6% | 77.2% | 69.4% | 60.0% | 67.6% | 
| AIME 2026 no tools | 89.2% | 88.3% | 77.5% | 42.5% | 37.5% | 20.8% | 
| GPQA Diamond | 84.3% | 82.3% | 78.8% | 58.6% | 43.4% | 42.4% | 
| Tau2 (average over 3) | 76.9% | 68.2% | 69.0% | 42.2% | 24.5% | 16.2% | 
| BigBench Extra Hard | 74.4% | 64.8% | 53.0% | 33.1% | 21.9% | 19.3% | 
| MMMLU | 88.4% | 86.3% | 83.4% | 76.6% | 67.4% | 70.7% | 
| **Coding** |  |  |  |  |  |  | 
| LiveCodeBench v6 | 80.0% | 77.1% | 72.0% | 52.0% | 44.0% | 29.1% | 
| Codeforces ELO | 2150 | 1718 | 1659 | 940 | 633 | 110 | 
| HLE no tools | 19.5% | 8.7% | 5.2% | - | - | - | 
| HLE with search | 26.5% | 17.2% | - | - | - | - | 
| **Vision** |  |  |  |  |  |  | 
| MMMU Pro | 76.9% | 73.8% | 69.1% | 52.6% | 44.2% | 49.7% | 
| OmniDocBench 1.5 (edit distance) | 0.131 | 0.149 | 0.164 | 0.181 | 0.290 | 0.365 | 
| MATH-Vision | 85.6% | 82.4% | 79.7% | 59.5% | 52.4% | 46.0% | 
| MedXPertQA MM | 61.3% | 58.1% | 48.7% | 28.7% | 23.5% | - | 
| **Audio** |  |  |  |  |  |  | 
| CoVoST | - | - | 38.5 <sup>*</sup> | 35.54 | 33.47 | - | 
| FLEURS (lower is better) | - | - | 0.069 <sup>*</sup> | 0.08 | 0.09 | - | 
| **Long Context** |  |  |  |  |  |  | 
| MRCR v2 8 needle 128k (average) | 66.4% | 44.1% | 43.4% | 25.4% | 19.1% | 13.5% | 

<sup>*</sup>Excluding Chinese language.

This work wouldn't have been possible without Google's extensive contribution with the model artefact, but also the significant effort contributing the model to transformers in an effort to standardize it. The open-source ecosystem is now more complete, with a very capable, freely-licensed, open-source model. The Gemma 4 transformers integration was handled by Cyril, Raushan, Eustache, Arthur, Lysandre. We thank Joshua for the transformers.js integration and demo, Eric for mistral.rs integration, Son for Llama.cpp, Prince for MLX, Quentin, Albert and Kashif for TRL, Adarsh for SGLang transformers backend, and Toshihiro for building several demos.
