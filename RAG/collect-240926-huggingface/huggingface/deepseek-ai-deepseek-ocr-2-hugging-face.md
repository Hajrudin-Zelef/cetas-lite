---
id: collect-240926-huggingface/huggingface/deepseek-ai-deepseek-ocr-2-hugging-face
title: "prompt = \"<image>\\nFree OCR. \""
domain: huggingface
role: reference
task: reference
actors: ["DeepSeek", "Nvidia"]
dates: []
keywords: ["attention", "benchmark", "deepseek", "gpus", "inference", "leaderboard", "nvidia", "safetensors"]
source: docs/RAG/clean_en/huggingface/deepseek-ai-deepseek-ocr-2-hugging-face.md
source_anchor: ""
source_lines: [1, 76]
sha256: 5fc85d291f90488e87cb1be5df5c57968b29a66fc57bc09097af99361efa35bc
---

# prompt = "<image>\nFree OCR. "

<!-- source: https://huggingface.co/deepseek-ai/DeepSeek-OCR-2 -->

**🌟 Github** |
  **📥 Model Download** |
  **📄 Paper Link** |
  **📄 Arxiv Paper Link** |


Inference using Huggingface transformers on NVIDIA GPUs. Requirements tested on python 3.12.9 + CUDA11.8：

```
torch==2.6.0
transformers==4.46.3
tokenizers==0.20.3
einops
addict 
easydict
pip install flash-attn==2.7.3 --no-build-isolation
```
```
from transformers import AutoModel, AutoTokenizer
import torch
import os
os.environ["CUDA_VISIBLE_DEVICES"] = '0'
model_name = 'deepseek-ai/DeepSeek-OCR-2'
tokenizer = AutoTokenizer.from_pretrained(model_name, trust_remote_code=True)
model = AutoModel.from_pretrained(model_name, _attn_implementation='flash_attention_2', trust_remote_code=True, use_safetensors=True)
model = model.eval().cuda().to(torch.bfloat16)
# prompt = "<image>\nFree OCR. "
prompt = "<image>\n<|grounding|>Convert the document to markdown. "
image_file = 'your_image.jpg'
output_path = 'your/output/dir'
res = model.infer(tokenizer, prompt=prompt, image_file=image_file, output_path = output_path, base_size = 1024, image_size = 768, crop_mode=True, save_results = True)
```
Refer to 🌟GitHub for guidance on model inference acceleration and PDF processing, etc.

- Dynamic resolution
  - Default: (0-6)×768×768 + 1×1024×1024 — (0-6)×144 + 256 visual tokens ✅

```
# document: <image>\n<|grounding|>Convert the document to markdown.
# without layouts: <image>\nFree OCR.
```
We would like to thank DeepSeek-OCR, Vary, GOT-OCR2.0, MinerU, PaddleOCR for their valuable models and ideas.

We also appreciate the benchmark OmniDocBench.

```
@article{wei2025deepseek,
  title={DeepSeek-OCR: Contexts Optical Compression},
  author={Wei, Haoran and Sun, Yaofeng and Li, Yukun},
  journal={arXiv preprint arXiv:2510.18234},
  year={2025}
}
@article{wei2026deepseek,
  title={DeepSeek-OCR 2: Visual Causal Flow},
  author={Wei, Haoran and Sun, Yaofeng and Li, Yukun},
  journal={arXiv preprint arXiv:2601.20552},
  year={2026}
}
```
- Downloads last month
- 857,124

## Spaces using deepseek-ai/DeepSeek-OCR-2 38

## Collection including deepseek-ai/DeepSeek-OCR-2

## Papers for deepseek-ai/DeepSeek-OCR-2

- allenai/olmOCR-bench leaderboard
- Overall View evaluation resultssource76.3
- Arxiv Math View evaluation resultssource82
- Old Scans Math View evaluation resultssource72
- llamaindex/ParseBench leaderboard
- Mean View evaluation resultssourcePipeline name: deepseekocr2_vllm41.2<sup>*</sup>
