---
id: collect-240926-huggingface/huggingface/inclusionai-ming-image-0-1-design-hugging-face
title: "inclusionai-ming-image-0-1-design-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Ant", "Hugging Face", "vLLM"]
dates: []
keywords: ["gpu", "inference", "license", "mit license", "omni", "text-to-image", "vllm"]
source: docs/RAG/clean_en/huggingface/inclusionai-ming-image-0-1-design-hugging-face.md
source_anchor: ""
source_lines: [1, 45]
sha256: 079da2d04361d4adbcfac203c16d49b5e22cd0ceca38a183daf64e0318a0d296
---

# inclusionai-ming-image-0-1-design-hugging-face

<!-- source: https://huggingface.co/inclusionAI/Ming-Image-0.1-Design -->

🧩 ModelScope · 🤗 Hugging Face · 📄 Blog · 🖥️ Demo

🎨 Design Skill · 📊 PPT Skill

Ming-Image-0.1-Design is a 6B text-to-image model for UI, infographics, posters, and other text-rich visual designs. It generates complete visual compositions and supports RGBA output with transparent backgrounds.


Use the companion Ming-Image repository for installation and inference:

```
git clone https://github.com/inclusionAI/Ming-Image
cd Ming-Image
pip install -r requirements.txt
python infer.py \
  --model inclusionAI/Ming-Image-0.1-Design \
  --task text-to-image \
  --prompt assets/t2i_four_seasons_cabin_prompt.json \
  --resolution 2048 \
  --output-dir outputs/t2i
```
Prompt enhancement (PE) can use `Ling-3.0-flash-VL` or `qwen3.8-27B`; see
text-to-image prompt rewriting.

For transparent-background generation, prepend exactly one of the recommended RGBA phrases. See the transparent-background generation tip.

We recommend the following inference frameworks to serve the model:

- vLLM-Omni: see the recipes and installation guide.

- Resolution: **2048 x 2048** (recommended), or**1024 x 1024** for faster
generation.
- Sampling steps: **12** .
- CFG scale: **1.0** .
- Precision: **BF16** .
- Hardware: **one CUDA GPU with 80 GiB VRAM** (validated configuration).

The public inference code maps text-to-image resolution requests to the supported 1024 or 2048 bucket.



The checkerboard is used only to preview transparency; it is not part of the generated RGBA images.

This model is released under the MIT License.
