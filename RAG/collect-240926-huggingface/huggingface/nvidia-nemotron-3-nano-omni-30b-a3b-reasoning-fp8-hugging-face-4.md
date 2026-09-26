---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-4
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmarks", "multimodal", "omni", "reasoning", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [617, 664]
sha256: bc96b7964737606337c5e56d7435c09229d38f34395786a3228c1226e0b1d4f8
---

# Log in once; the token is cached at ~/.cache/huggingface/token

**Time period for testing data collection:** N/A (standard public benchmarks) 

**Time period for validation data collection:** N/A (standard public benchmarks) 

For more information about the datasets used to train this model, please see the Public Summary of Training Content

Nemotron-Omni extends our commitment from text to multimodal, delivering the same level of openness across text, audio, image, and video.

**Adapter and encoder training scale:** ~127B tokens across mixed modalities spanning text+image, text+video, text+audio, and text+video+audio—reflecting real-world, contextualized interactions versus single-modality data.

**Post-training for real-world tasks:** ~124M curated examples across multimodal combinations (text+audio, text+image, text+video, and text+video+audio), structured to support document reasoning, computer use, and long-horizon workflows.

**RL environments for agent training:** 20 RL datasets across 25 environments covering 5 new multimodal tasks—visual grounding, chart and document understanding, vision-critical STEM problems, video understanding, and automatic speech recognition—extending Nemotron's RL pipeline beyond text into vision and audio.

**Modality Breakdown:**

| Modality | Dataset Entries | Samples | Est. Tokens (M) | 
|---|---|---|---|
| text+audio | 220 | 259,178,821 | 143,533.1 | 
| text+image | 750 | 70,143,901 | 180,347.1 | 
| text+video | 241 | 15,837,673 | 239,631.5 | 
| text+video+audio | 155 | 8,720,044 | 152,499.2 | 
| text | 12 | 707,187 | 958.4 | 
| **Total** | **1395** | **354,587,705** | **716,969.2** | 

Training data for Nemotron-Omni was assembled from a diverse collection of audio, image, video, and text datasets. Raw datasets were first converted into a standardized JSONL format with unified conversation-turn structure. Audio data was resampled to 16 kHz where needed. Image and video datasets were paired with question-answer annotations, often regenerated or refined using large vision-language models to improve quality and consistency. Quality filtering was applied using model-based judges to remove low-quality, unsafe, or off-topic samples. Deduplication and CSAM scanning were performed across all image datasets. Data was then packed into fixed-length sequences (32k, 128k, or 256k tokens) for efficient training. 

Multiple safety measures were implemented throughout the data pipeline. All image/text datasets underwent CSAM (Child Sexual Abuse Material) scanning, with results tracked per dataset. Content safety filtering was applied using two independent safety judge models to flag and remove samples containing harmful content including weapons references, criminal planning, sexual content involving minors, harassment, hate speech, profanity, threats, violence, or suicide-related content. Synthetic data generation pipelines included explicit quality and safety filtering stages. Identity-fix processing was applied to correct potential biases in generated responses. The multi-stage pipeline (original → cleaned → clean+safe → clean+safe+holdout) ensured progressive refinement, with each stage removing additional problematic content. 

We built on the base model, applying additional training, enhancements, and optimizations on top of it.

| Dataset | Samples | % of Public | Tokens (M) | Modality | 
|---|---|---|---|---|
| MiraData | 28,252,307 | 55.53% | 14,181.3 | text+audio+video | 
| laion-disco-12M | 7,507,574 | 14.7% | 22,691.0 | text+audio | 
| YouTube Video | 2,057,000 | 4.0% | 15,390 | text+video | 
| YouTube Video and Audio | 1,164,000 | 2.2% | 18,730 | text+video+audio | 

| Dataset | Samples | % of Private | Tokens (M) | Modality | 
|---|---|---|---|---|
| Granary | 23,370,274 | 8.0% | 1,471.7 | text+audio | 
| SIFT-50M | 22,837,500 | 7.8% | 5,241.7 | text+audio | 

- Overall Size: 41,502,625 samples across modalities: text+audio, text+image, text+video
- Description of synthetic data generation methods:

Synthetic data generation (SDG) was used to improve data quality, generate reasoning traces, re-label annotations, and augment existing datasets. Methods include: re-captioning images and audio using vision-language models, generating question-answer pairs from existing media, producing thinking/reasoning chains for complex tasks, paraphrasing prompts for diversity, and applying model-based quality filtering. 

