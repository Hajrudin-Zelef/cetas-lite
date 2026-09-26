---
id: collect-240926-huggingface/huggingface/unsloth-gemma-4-26b-a4b-it-qat-gguf-hugging-face-2
title: "Read our How to Run Gemma 4 QAT Guide!"
domain: huggingface
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["agentic", "decode", "multimodal", "reasoning", "tool use", "transcription"]
source: docs/RAG/clean_en/huggingface/unsloth-gemma-4-26b-a4b-it-qat-gguf-hugging-face.md
source_anchor: ""
source_lines: [88, 283]
sha256: 9a16fdeab52242f665c0f0650980fc0d5819c9c757157c4a6b85326640af00fd
---

# Read our How to Run Gemma 4 QAT Guide!

|  | Gemma 4 31B | Gemma 4 26B A4B | Gemma 4 12B Unified | Gemma 4 E4B | Gemma 4 E2B | Gemma 3 27B (no think) | 
|---|---|---|---|---|---|---|
| MMLU Pro | 85.2% | 82.6% | 77.2% | 69.4% | 60.0% | 67.6% | 
| AIME 2026 no tools | 89.2% | 88.3% | 77.5% | 42.5% | 37.5% | 20.8% | 
| LiveCodeBench v6 | 80.0% | 77.1% | 72.0% | 52.0% | 44.0% | 29.1% | 
| Codeforces ELO | 2150 | 1718 | 1659 | 940 | 633 | 110 | 
| GPQA Diamond | 84.3% | 82.3% | 78.8% | 58.6% | 43.4% | 42.4% | 
| Tau2 (average over 3) | 76.9% | 68.2% | 69.0% | 42.2% | 24.5% | 16.2% | 
| HLE no tools | 19.5% | 8.7% | 5.2% | - | - | - | 
| HLE with search | 26.5% | 17.2% | - | - | - | - | 
| BigBench Extra Hard | 74.4% | 64.8% | 53.0% | 33.1% | 21.9% | 19.3% | 
| MMMLU | 88.4% | 86.3% | 83.4% | 76.6% | 67.4% | 70.7% | 
| **Vision** |  |  |  |  |  |  | 
| MMMU Pro | 76.9% | 73.8% | 69.1% | 52.6% | 44.2% | 49.7% | 
| OmniDocBench 1.5 (average edit distance, lower is better) | 0.131 | 0.149 | 0.164 | 0.181 | 0.290 | 0.365 | 
| MATH-Vision | 85.6% | 82.4% | 79.7% | 59.5% | 52.4% | 46.0% | 
| MedXPertQA MM | 61.3% | 58.1% | 48.7% | 28.7% | 23.5% | - | 
| **Audio** |  |  |  |  |  |  | 
| CoVoST | - | - | 38.5 <sup>*</sup> | 35.54 | 33.47 | - | 
| FLEURS (lower is better) | - | - | 0.069 <sup>*</sup> | 0.08 | 0.09 | - | 
| **Long Context** |  |  |  |  |  |  | 
| MRCR v2 8 needle 128k (average) | 66.4% | 44.1% | 43.4% | 25.4% | 19.1% | 13.5% | 

<sup>*</sup>Excluding Chinese language.

Gemma 4 models handle a broad range of tasks across text, vision, and audio. Key capabilities include:

- **Thinking** – Built-in reasoning mode that lets the model think step-by-step before answering.
- **Long Context** – Context windows of up to 128K tokens (E2B/E4B) and 256K tokens (12B, 26B A4B/31B).
- **Image Understanding** – Object detection, Document/PDF parsing, screen and UI understanding, chart comprehension, OCR (including multilingual), handwriting recognition, and pointing. Images can be processed at variable aspect ratios and resolutions.
- **Video Understanding** – Analyze video by processing sequences of frames.
- **Interleaved Multimodal Input** – Freely mix text and images in any order within a single prompt.
- **Function Calling** – Native support for structured tool use, enabling agentic workflows.
- **Coding** – Code generation, completion, and correction.
- **Multilingual** – Out-of-the-box support for 35+ languages, pre-trained on 140+ languages.
- **Audio** (E2B, E4B, and 12B only) – Automatic speech recognition (ASR) and speech-to-translated-text translation across multiple languages.

You can use all Gemma 4 models with the latest version of Transformers. To get started, install the necessary dependencies in your environment:

`pip install -U transformers torch accelerate`

Once you have everything installed, you can proceed to load the model with the code below:

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "google/gemma-4-12B-it"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto"
)
```
Once the model is loaded, you can start generating output:

```
# Prompt
messages = [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Write a short joke about saving RAM."},
]
# Process input
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    enable_thinking=False
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
# Generate output
outputs = model.generate(**inputs, max_new_tokens=1024)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
# Parse output
processor.parse_response(response)
```
To enable reasoning, set `enable_thinking=True` and the `parse_response` function will take care of parsing the thinking output.

Below, you will also find snippets for processing audio (E2B, E4B, 12B only), images, and video alongside text:

## Code for processing Audio

Make sure to install the following packages:

`pip install -U transformers torch torchvision librosa accelerate`

You can then load the model with the code below:

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "google/gemma-4-12B-it"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID, 
    dtype="auto", 
    device_map="auto"
)
```
Once the model is loaded, you can start generating output by directly referencing the audio URL in the prompt:

```
# Prompt - add audio after text
messages = [
    {
        "role": "user",
        "content": [
            {"type": "text", "text": "Transcribe the following speech segment in its original language. Follow these specific instructions for formatting the answer:\n* Only output the transcription, with no newlines.\n* When transcribing numbers, write the digits, i.e. write 1.7 and not one point seven, and write 3 instead of three."},
            {"type": "audio", "audio": "https://raw.githubusercontent.com/google-gemma/cookbook/refs/heads/main/apps/sample-data/journal1.wav"},
        ]
    }
]
# Process input
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
# Generate output
outputs = model.generate(**inputs, max_new_tokens=512)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
# Parse output
processor.parse_response(response)
```
## Code for processing Images

Make sure to install the following packages:

`pip install -U transformers torch torchvision accelerate`

You can then load the model with the code below:

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "google/gemma-4-12B-it"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID, 
    dtype="auto", 
    device_map="auto"
)
```
Once the model is loaded, you can start generating output by directly referencing the image URL in the prompt:

```
# Prompt - add image before text
messages = [
    {
        "role": "user", "content": [
            {"type": "image", "url": "https://raw.githubusercontent.com/google-gemma/cookbook/refs/heads/main/apps/sample-data/GoldenGate.png"},
            {"type": "text", "text": "What is shown in this image?"}
        ]
    }
]
# Process input
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
# Generate output
outputs = model.generate(**inputs, max_new_tokens=512)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
# Parse output
processor.parse_response(response)
```
## Code for processing Videos

Make sure to install the following packages:

`pip install -U transformers torch torchvision librosa accelerate`

You can then load the model with the code below:

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "google/gemma-4-12B-it"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID, 
    dtype="auto", 
    device_map="auto"
)
```
Once the model is loaded, you can start generating output by directly referencing the video URL in the prompt:

