---
id: collect-261001-general-networking/general-networking/welcome-inkling-by-thinking-machines-3
title: "model_id = \"thinkingmachines/Inkling-NVFP4\""
domain: general-networking
role: reference
task: reference
actors: ["China", "Mistral", "Nvidia"]
dates: []
keywords: ["nvfp4", "agent", "agentic", "benchmark", "blackwell", "decode", "distillation", "fine-tuning", "gpu", "nvidia", "reasoning", "training"]
source: docs/RAG/collect-261001-general-networking/welcome-inkling-by-thinking-machines.md
source_anchor: ""
source_lines: [287, 367]
sha256: 23076c29b21fc3b930b16472117553e169e9c680d7d7212399c4612846748309
---

# model_id = "thinkingmachines/Inkling-NVFP4"

```
import torch
from transformers import AutoModelForMultimodalLM, AutoProcessor
processor = AutoProcessor.from_pretrained("thinkingmachines/Inkling")
model = AutoModelForMultimodalLM.from_pretrained(
    "thinkingmachines/Inkling",
    dtype=torch.bfloat16,
    device_map="auto",
)
# Preprocess the inputs.
...
generated = model.generate(
    **inputs,
    max_new_tokens=1000,
    do_sample=False,
    use_mtp=True,
)
print(processor.decode(generated[0], skip_special_tokens=True))
```
We have prepared a small suite of reasoning questions from expert-level sources and university entrance exams. We have taken photos of the screen with watermarks in the screenshot to challenge the model. The model has solved all of them on high one, failed one in highest and medium reasoning efforts, so we provide a link to the model answers for you to check out how the model sounds and provide the number of tokens the model has taken to solve each of them. Note that we provide no system prompts in these vibe evals, and these reasoning questions should often be run with a good system prompt. The vibe eval images and results live here.

| Category | Question | Number of Tokens (Reasoning Effort Medium) | Number of Tokens (Reasoning Effort High) | Number of Tokens (Reasoning Effort Max) | 
|---|---|---|---|---|
| Open-ended Drug Interactions | Which components interact here? | 1,893 ✅ | 2,367 ✅ | 3,688 ✅ | 
| Physics Question (MMMU-Pro) | Answer the question in the image. | 1,357 ✅ | 3,323 ✅ | 3,314 ✅ | 
| Multilingual Physics Question | Answer the Turkish question given in the image. | 1,435 ✅ | 2,129 ✅ | 3,162 ✅ | 
| Bar Exam | Answer the question in the image. | 1,117 ✅ | 2,137 ✅ | 1,676 ✅ | 
| Infographics Question Answering (Open-ended) | Based on the information presented, approximately how many times larger is the projected summer warming period in the Arctic than the time over which substantial Arctic warming has already been observed? | 1,378 ❌ | 3,859 ✅ | 6000 (exceeded token budget) | 

**Few notes on vibes:**

- Instead of directly answering the question on infographic, the model first turns text on image to text to ground itself.
- Prompting matters a lot to save tokens in reasoning, for instance, asking vague questions like “which components interact here?” with an image of the back of a pill, the model first needs to see what we mean by interactions here.
- Multi-choice question answers helped the model a lot in structuring its own reasoning, for open-ended questions the model struggled compared to MCQA, however, this is a common issue for many models. The usual chain of thought was OCR → characterize → evaluate each option → answer.
- 0.7 reasoning effort (medium) seems to provide a good trade-off.

We have vibe-evaluated the model on some audio reasoning examples from BigBenchAudio and a few multilingual audio examples of GlobeAudio (Russian and Chinese multi-choice questions asking the last word in transcription). The BigBenchAudio examples we tested consist of logical statements and questions that either ask for formal fallacies (whether an argument can be logically deduced from the context given in audio) or object counting (stating multiple distinctive objects in the audio, asking for the total count of a certain one). Although this benchmark is initially made for speech-to-speech reasoning, we just want to see audio reasoning capabilities of this model. For GlobeAudio, the questions are relatively straightforward, so we ran with reasoning efforts of 0.1. We ran the first example of each language within GlobeAudio. All tests pass on all questions and efforts, except for second formal fallacy example on lowest effort, so we only provide the number of tokens spent in each question against reasoning effort. Vibe eval results and audio files live here.

| GlobeAudio | Question | Number of completion tokens (Reasoning effort lowest) | Number of completion tokens (Reasoning effort medium) | 
|---|---|---|---|
| Russian (asks for last word) | Какое последнее слово в аудиозаписи? 1. Россия 2. Свидетелем 3. Москва 4. Событий Choose the single correct option and answer with its exact text. | 130 | 179 | 
| Russian (asks for profession of the speaker) | Кем, скорее всего, работает говорящая? 1. Репортершей 2. Блоггершей 3. Учительницей истории 4. Ведущей развлекательного шоу Choose the single correct option and answer with its exact text. | 105 | 136 | 
| Chinese (asks for speaking rate) | 播报员的语速有何变化？ 1. 突然变快 2. 突然变慢 3. 保持不变 4. 时快时慢 Choose the single correct option and answer with its exact text. | 111 | 289 | 

| Big Bench Audio | Completion Tokens (lowest) | Completion Tokens (medium) | Number of completion tokens ( highest) | 
|---|---|---|---|
| Formal Fallacy (10) | 285 | 335 | 444 | 
| Formal Fallacy (39) | 275 (fails) | 555 | 778 | 
| Object Counting (680) | 150 | 233 | 161 | 

**Some notes on the vibes:**

- Similar to vision, the model first transcribes the speech before answering the question.
- It resists decoys: in Russian test, the model picked the right answer despite other answers appearing in the audio.
- Similar to vision, usual chain of thought is transcribe → characterize → evaluate each option → answer.
- The effort helps reasoning and not hearing. Audio question answering was much cheaper than images.

If you would like to use Inkling for post-training, Thinking Machines have built `tinker`, a managed tool for post-training open weight models. Their cookbook includes examples for fine-tuning, distillation, and reinforcement learning.

We post trained Inkling with tinker and OpenEnv, an agentic RL environment tool. We used the ECHO algorithm that trains a model to predict the environment without a verifier, applying next-token cross-entropy loss to tokens produced by the environment, alongside the usual policy learning on agent actions. This teaches the policy an implicit world model without requiring a separate model, teacher, or additional rollouts. Check out the example.

## RL Example with Tinker and OpenEnv

```
git clone https://github.com/huggingface/OpenEnv.git
cd OpenEnv
# Add TINKER_API_KEY=... to .env, then run:
uv run --env-file .env \
  examples/echo_world_model/backends/tinker_echo_demo.py
```
If you’re working with Transformers Reinforcement Learning we suggest using Inkling as a teacher model in a knowledge distillation setup. For example, take advantage of Inkling’s document understanding abilities to improve the performance of a smaller (on-device) model. In this example, we use the transformer reinforcement learning library and the GOLD algorithm to distill knowledge. GOLD is handy here because it matches token logits between different tokenizers, so you can distill to any model on the hub.

Below you can find each Inkling checkpoint as well as their VRAM requirements.

| Model Variant | Aggregated VRAM | Recommended GPU Configurations | Deployment Notes | 
|---|---|---|---|
| **Inkling (BF16)** | 2 TB | • 8× NVIDIA B300 / GB200 • 16× NVIDIA H200 | Full precision deployment; requires multi-node interconnect for H200 clusters. | 
| **Inkling (NVFP4)** | 600 GB | • 4× NVIDIA B300 / GB200 (W4A4) • 8× NVIDIA H200 (W4A16) | W4A4 requires Blackwell architecture (SM100+). | 
| **Inkling-Small (BF16)** | 600 GB | • 4× NVIDIA B300 / GB200 (W4A4) • 8× NVIDIA H200 (W4A16) | Does not require Blackwell architecture; easily deployed to 8× H200s. | 
| **Inkling-Small (NVFP4)** | 180 GB | • 1× NVIDIA B300 (W4A4) • 2× NVIDIA H200 (W4A16) | W4A4 supported on single Blackwell GPU; W4A16 supported on 2× H200s. | 

