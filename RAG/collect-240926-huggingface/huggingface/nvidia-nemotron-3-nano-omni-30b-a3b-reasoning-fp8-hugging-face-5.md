---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-5
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "fp8", "gemini", "glm", "gqa", "kimi", "nvidia", "omni", "qwen", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [665, 787]
sha256: 42a794cdadf850c37ca440be46072a941fdfbe75065d8adaf115e7567fd18ca1
---

# Log in once; the token is cached at ~/.cache/huggingface/token

| Dataset | Modality | Count | Models Used | 
|---|---|---|---|
| GroundCUA | text+image | 2,797,851 | gpt-oss-120b, Qwen3-VL-30B-A3B-Instruct | 
| OpenImages | text+image | 2,556,412 | Qwen3-VL-30B-A3B-Instruct | 
| MMTrail | text+audio | 1,620,533 | Qwen3-omni-captioner, gpt-oss-120B | 
| Localized Narratives | text+image | 1,511,812 | Qwen3-VL-30B-A3B-Instruct | 
| ALLaVA | text+image | 1,414,130 | Qwen3-VL-30B-A3B-Instruct | 
| VGG-Sound | text+audio | 1,371,167 | Qwen3-omni-captioner, gpt-oss-120B | 
| PIXMO-CAP | text+image | 1,308,838 | Qwen3-VL-30B-A3B-Instruct | 
| TTS-Synthesized Nemotron-Nano-3 SFT Data | text+audio | 1,226,784 | NVIDIA Magpie TTS | 
| MINT-1T | text+image | 904,035 | Qwen3-VL-32B-Instruct, Gemini 3 Pro for filtering, Scene Text models (RTX) translate | 
| ScaleCUA | text+image | 889,010 | Qwen3-VL-30B-A3B-Instruct | 
| AgentNet | text+image | 878,986 | Kimi-K2.5 | 
| Conceptual Captions 3M-30b | text+image | 867,065 | Qwen3-VL-30B-A3B-Thinking-FP8 | 
| MetaMathQA | text+image | 860,656 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Mulberry-SFT COT | text+image | 566,982 | GLM-4.1V-9B-Thinking, Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| CC for OCR | text+image | 522,595 | SwinDocSegmenter, DeepSeek OCR, Qwen3.5-122B-A10B, Qwen3-32B, Gemini 3 Flash Preview for filtering, GPT-4o mini for filtering & quality checks, Qwen3-VL-30B-A3B-Thinking-FP8, gpt-oss-120b | 
| Charxiv-100K | text+image | 272,104 | Qwen3-VL-235B-A22B-Instruct, Qwen3-VL-235B-A22B-Thinking, GPT-4o for filtering, Qwen3.5-122B-A10B | 
| SwinDocSegmenter | text+image | 207,200 | SwinDocSegmenter, DeepSeek OCR | 
| CLEVR | text+image, text+video | 197,027 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| InternVL-Data | text+image | 185,395 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Flickr30k Entities | text+image | 154,760 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Metropolis and Lita | text+video | 150,434 | Qwen3.5-122B-A10B | 
| TextCaps | text+image | 136,911 | Commercial VILA model, Qwen3-VL-30B-A3B-Instruct | 
| Vision R1 Llava CoT | text+image | 126,024 | GLM-4.1V-9B-Thinking | 
| HC-STVG | text+video | 124,902 | NVIDIA relabeled using Qwen model (Qwen2.5-VL-72B-Instruct) | 
| nvPDFtex | text+image | 118,351 | gpt-oss-120b, Qwen3.5-122B-A10B | 
| ChartQA | text+image | 111,602 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering, Qwen2-VL-72B (NV) | 
| ECD-10k-Images | text+image | 110,697 | Qwen3.5-122B-A10B | 
| SAMA-COCO | text+image | 102,965 | gpt-oss-120B | 
| VisualWebInstruct | text+image | 97,746 | Earlier SDG, GLM-4.1V-9B-Thinking | 
| Spatial | text+image | 95,532 | Microsoft Florence-2-large | 
| DoubtNut | text+image | 94,919 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Cosmos Nemotron SFTv13.9 | text+image | 92,128 | Qwen3-VL-30B-A3B-Instruct, Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| CrossTask | text+video | 76,495 | NVIDIA relabeled using Qwen model (Qwen2.5-VL-72B-Instruct) | 
| RefCOCO | text+image | 69,850 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Mantis Instruct | text+image | 66,975 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Visual7W | text+image | 62,589 | Qwen3.5-122B-A10B | 
| ScreenQA | text+image | 62,186 | Qwen3.5-122B-A10B | 
| VQAV2 | text+image | 54,899 | Qwen3.5-122B-A10B | 
| TallyQA | text+image | 50,073 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| KeenSight | text+image | 49,849 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| GQA | text+image | 42,182 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| AskFilo | text+image | 41,807 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Raven | text+image | 41,996 | gpt-oss-120b | 
| DocVQA | text+image | 35,759 | Qwen3.5-122B-A10B | 
| TextVQA | text+image | 34,602 | Commercial VILA model, Qwen3-VL-30B-A3B-Instruct | 
| COCO | text+image | 32,111 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| PlotQA | text+image | 30,665 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Llava | text+video | 30,250 | Qwen3-Omni-30B-A3B-Instruct, Qwen3-VL-32B-Instruct | 
| NVCLIP | text+image | 29,680 | Qwen2.5-72B-Instruct | 
| Tapos | text+video | 29,250 | Qwen2.5-VL-72B-Instruct | 
| Vedantu Chemistry | text+audio | 26,338 | NVIDIA Magpie TTS | 
| NV-CC-Img-Text-Dataset | text+image | 24,998 | Qwen3-VL-30B-A3B-Instruct | 
| DocLayNet | text+image | 22,709 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering, gpt-oss-120b | 
| Taloka Grounding | text+image | 22,218 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Wikipedia OCR | text+image | 21,440 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| InternVL2.5 | text+image | 20,770 | Qwen3-VL-235B-A22B-Instruct, Qwen3-VL-235B-A22B-Thinking, GPT-4o for filtering, Qwen3.5-122B-A10B | 
| PromptPG | text+image | 20,305 | Qwen2-VL-72B | 
| PubTables | text+image | 20,174 | gpt-oss-120b | 
| InfoVQA | text+image | 18,679 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Azure Tables | text+image | 18,188 | gpt-oss-120b, Qwen3.5-122B-A10B | 
| TabRecSet | text+image | 17,437 | GPT-4o mini, Qwen3-VL-30B-A3B-Thinking-FP8, gpt-oss-120b, Qwen3.5-122B-A10B | 
| CD Questions | text+audio, text+image | 16,335 | NVIDIA Magpie TTS, Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Linguistic Data Consortium | text+image | 15,499 | Qwen3.5-122B-A10B, GPT-4o mini, Qwen3-VL-30B-A3B-Thinking-FP8, gpt-oss-120b, Ask Kateryna | 
| MapQA | text+image | 12,480 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| SlideVQA | text+image | 11,199 | Qwen3.5-122B-A10B | 
| OCR Reason Finance | text+image | 9,389 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| GeomVerse | text+image | 9,298 | GLM-4.1V-9B-Thinking | 
| NextQA | text+video | 8,903 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| UniGeo | text+image | 8,822 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| Vedantu | text+audio, text+image | 8,750 | NVIDIA Magpie TTS, Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| GPQA | text+audio | 7,657 | NVIDIA Magpie TTS | 
| SLAKE | text+image | 7,294 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| OpenGVLab | text+image | 7,269 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering, Qwen3-VL-235B-A22B-Instruct, Qwen3-VL-235B-A22B-Thinking, GPT-4o for filtering | 
| PerceptionTest | text+video | 5,192 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| InvoicesQA | text+image | 4,817 | Qwen3.5-122B-A10B | 
| EgoProcel | text+video | 4,660 | Qwen2.5-VL-72B-Instruct | 
| SynthTabNet | text+image | 4,364 | gpt-oss-120b | 
| SerpAPI | text+image | 3,784 | Qwen3.5-122B-A10B, Gemini 3 Flash Preview for filtering | 
| FinTabNet | text+image | 3,852 | gpt-oss-120b | 
| FastMath | text+image | 3,718 | Qwen3-VL-235B-A22B-Instruct-FP8 | 
| ASR Data Derived Speech-to-Text Chat Data | text+audio | 3,608 | GPT-OSS 120B | 
| Geometry3k | text+image | 2,078 | Qwen3-VL-235B-A22B-Thinking-FP8 | 
| VQA-RAD | text+image | 1,270 | Qwen3.5-122B-A10B | 
| RQA | text+audio | 959 | NVIDIA Magpie TTS | 
| HierText OCRQA Qwen | text+image | 514 | Qwen2.5-VL-32B-Instruct | 

**Data Modality** 

- Audio 
- Image 
- Text 
- Video 

**Audio Training Data Size** 

- 10,000 to 1 Million Hours 
(267,898,865 audio-containing samples)

**Image Training Data Size** 

- 1 Million to 1 Billion Images 
(70,143,901 image-containing samples)

**Text Training Data Size** 

- 1 Billion to 10 Trillion Tokens 
(~717.0B tokens total across all modalities)

**Video Training Data Size** 

- 10,000 to 1 Million Hours 
(24,557,717 video-containing samples)

**Data Collection Method by dataset** 

- Hybrid: Human, Automated, Synthetic 

**Labeling Method by dataset** 

- Hybrid: Human, Automated, Synthetic 

