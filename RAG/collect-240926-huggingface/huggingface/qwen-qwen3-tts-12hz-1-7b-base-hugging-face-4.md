---
id: collect-240926-huggingface/huggingface/qwen-qwen3-tts-12hz-1-7b-base-hugging-face-4
title: "Download through ModelScope (recommended for users in Mainland China)"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Google", "MiniMax", "OpenAI"]
dates: []
keywords: ["benchmark", "benchmarks", "gemini", "qwen", "research", "voice"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-tts-12hz-1-7b-base-hugging-face.md
source_anchor: ""
source_lines: [360, 485]
sha256: 860241293bb68e71f5ef0d28d97b497567b11a0a3ba7775a1c8dd712b10b69e7
---

# Download through ModelScope (recommended for users in Mainland China)

| Language | Qwen3-TTS-25Hz |  | Qwen3-TTS-12Hz |  | MiniMax | ElevenLabs | 
|---|---|---|---|---|---|---|
|  | 0.6B-Base | 1.7B-Base | 0.6B-Base | 1.7B-Base |  |  | 
| *Content Consistency* |  |  |  |  |  |  | 
| Chinese | 1.108 | **0.777** | 1.145 | 0.928 | 2.252 | 16.026 | 
| English | 1.048 | 1.014 | **0.836** | 0.934 | 2.164 | 2.339 | 
| German | 1.501 | 0.960 | 1.089 | 1.235 | 1.906 | **0.572** | 
| Italian | 1.169 | 1.105 | 1.534 | **0.948** | 1.543 | 1.743 | 
| Portuguese | 2.046 | 1.778 | 2.254 | 1.526 | 1.877 | **1.331** | 
| Spanish | 2.031 | 1.491 | 1.491 | 1.126 | **1.029** | 1.084 | 
| Japanese | 4.189 | 5.121 | 6.404 | 3.823 | **3.519** | 10.646 | 
| Korean | 2.852 | 2.631 | **1.741** | 1.755 | 1.747 | 1.865 | 
| French | 2.852 | **2.631** | 2.931 | 2.858 | 4.099 | 5.216 | 
| Russian | 5.957 | 4.535 | 4.458 | **3.212** | 4.281 | 3.878 | 
| *Speaker Similarity* |  |  |  |  |  |  | 
| Chinese | 0.797 | 0.796 | **0.811** | 0.799 | 0.780 | 0.677 | 
| English | 0.811 | 0.815 | **0.829** | 0.775 | 0.756 | 0.613 | 
| German | 0.749 | 0.737 | 0.769 | **0.775** | 0.733 | 0.614 | 
| Italian | 0.722 | 0.718 | 0.792 | **0.817** | 0.699 | 0.579 | 
| Portuguese | 0.790 | 0.783 | 0.794 | **0.817** | 0.805 | 0.711 | 
| Spanish | 0.732 | 0.731 | 0.812 | **0.814** | 0.762 | 0.615 | 
| Japanese | **0.810** | 0.807 | 0.798 | 0.788 | 0.776 | 0.738 | 
| Korean | **0.824** | 0.814 | 0.812 | 0.799 | 0.779 | 0.700 | 
| French | 0.698 | 0.703 | 0.700 | **0.714** | 0.628 | 0.535 | 
| Russian | 0.734 | 0.744 | 0.781 | **0.792** | 0.761 | 0.676 | 

*Cross-lingual speech generation on the Cross-Lingual benchmark. Performance is measured by Mixed Error Rate (WER for English, CER for others, ↓).*

| Task | Qwen3-TTS-25Hz-1.7B-Base | Qwen3-TTS-12Hz-1.7B-Base | CosyVoice3 | CosyVoice2 | 
|---|---|---|---|---|
| en-to-zh | 5.66 | **4.77** | 5.09 | 13.5 | 
| ja-to-zh | 3.92 | 3.43 | **3.05** | 48.1 | 
| ko-to-zh | 1.14 | 1.08 | **1.06** | 7.70 | 
| zh-to-en | 2.91 | **2.77** | 2.98 | 6.47 | 
| ja-to-en | 3.95 | **3.04** | 4.20 | 17.1 | 
| ko-to-en | 3.48 | **3.09** | 4.19 | 11.2 | 
| zh-to-ja | 9.29 | 8.40 | **7.08** | 13.1 | 
| en-to-ja | 7.74 | 7.21 | **6.80** | 14.9 | 
| ko-to-ja | 4.17 | **3.67** | 3.93 | 5.86 | 
| zh-to-ko | 8.12 | **4.82** | 14.4 | 24.8 | 
| en-to-ko | 6.83 | **5.14** | 5.87 | 21.9 | 
| ja-to-ko | 6.86 | **5.59** | 7.92 | 21.5 | 

*Controllable speech generation on InstructTTSEval. Performance is measured by Attribute Perception and Synthesis accuracy (APS), Description-Speech Consistency (DSD), and Response Precision (RP).*

| Type | Model | InstructTTSEval-ZH |  |  | InstructTTSEval-EN |  |  | 
|---|---|---|---|---|---|---|---|
|  |  | APS (↑) | DSD (↑) | RP (↑) | APS (↑) | DSD (↑) | RP (↑) | 
| *Target Speaker* | Gemini-flash | 88.2 | **90.9** | **77.3** | **92.3** | **93.8** | **80.1** | 
|  | Gemini-pro | **89.0** | 90.1 | 75.5 | 87.6 | 86.0 | 67.2 | 
|  | Qwen3TTS-25Hz-1.7B-CustomVoice | 83.1 | 75.0 | 63.0 | 79.0 | 82.8 | 69.3 | 
|  | Qwen3TTS-12Hz-1.7B-CustomVoice | 83.0 | 77.8 | 61.2 | 77.3 | 77.1 | 63.7 | 
|  | GPT-4o-mini-tts | 54.9 | 52.3 | 46.0 | 76.4 | 74.3 | 54.8 | 
| *Voice Design* | Qwen3TTS-12Hz-1.7B-VD | **85.2** | **81.1** | **65.1** | 82.9 | **82.4** | **68.4** | 
|  | Mimo-Audio-7B-Instruct (Zhang et al., 2025b) | 75.7 | 74.3 | 61.5 | 80.6 | 77.6 | 59.5 | 
|  | VoiceSculptor (Hu et al., 2026) | 75.7 | 64.7 | 61.5 | - | - | - | 
|  | Hume | - | - | - | **83.0** | 75.3 | 54.3 | 
|  | VoxInstruct (Zhou et al., 2024) | 47.5 | 52.3 | 42.6 | 54.9 | 57.0 | 39.3 | 
|  | Parler-tts-mini (Lyth & King, 2024) | - | - | - | 63.4 | 48.7 | 28.6 | 
|  | Parler-tts-large (Lyth & King, 2024) | - | - | - | 60.0 | 45.9 | 31.2 | 
|  | PromptTTS (Guo et al., 2023) | - | - | - | 64.3 | 47.2 | 31.4 | 
|  | PromptStyle (Liu et al., 2023) | - | - | - | 57.4 | 46.4 | 30.9 | 

*Target-Speaker Multilingual Speech Generation on the TTS multilingual test set. Performance is measured by Word Error Rate (WER, ↓).*

| Language | Qwen3-TTS-25Hz |  | Qwen3-TTS-12Hz |  | GPT-4o-Audio Preview | 
|---|---|---|---|---|---|
|  | 0.6B-CustomVoice | 1.7B-CustomVoice | 0.6B-CustomVoice | 1.7B-CustomVoice |  | 
| Chinese | 0.874 | **0.708** | 0.944 | 0.903 | 3.519 | 
| English | 1.332 | 0.936 | 1.188 | **0.899** | 2.197 | 
| German | 0.990 | **0.634** | 2.722 | 1.057 | 1.161 | 
| Italian | 1.861 | 1.271 | 2.545 | 1.362 | **1.194** | 
| Portuguese | 1.728 | 1.854 | 3.219 | 2.681 | **1.504** | 
| Spanish | 1.309 | 1.284 | **1.154** | 1.330 | 4.000 | 
| Japanese | **3.875** | 4.518 | 6.877 | 4.924 | 5.001 | 
| Korean | 2.202 | 2.274 | 3.053 | **1.741** | 2.763 | 
| French | 3.865 | **3.080** | 3.841 | 3.781 | 3.605 | 
| Russian | 6.529 | **4.444** | 5.809 | 4.734 | 5.250 | 

*Long speech generation results. Performance is measured by Word Error Rate (WER, ↓).*

| Datasets | Model | Performance |  | 
|---|---|---|---|
| *Content Consistency* |  |  |  | 
| *long-zh* \|*long-en* | Higgs-Audio-v2 (chunk) (Boson AI, 2025) | 5.505 | 6.917 | 
|  | VibeVoice (Peng et al., 2025) | 22.619 | 1.780 | 
|  | VoxCPM (Zhou et al., 2025) | 4.835 | 7.474 | 
|  | Qwen3-TTS-25Hz-1.7B-CustomVoice | **1.517** | **1.225** | 
|  | Qwen3-TTS-12Hz-1.7B-CustomVoice | 2.356 | 2.812 | 

## Speech Tokenizer Benchmarks

*Comparison between different supervised semantic speech tokenizers on ASR Task.*

| Model | Codebook Size | FPS | C.V. EN | C.V. CN | Fluers EN | Fluers CN | 
|---|---|---|---|---|---|---|
| S3 Tokenizer(VQ) (Du et al., 2024a) | 4096 | 50 | 12.06 | 15.38 | - | - | 
| S3 Tokenizer(VQ) (Du et al., 2024a) | 4096 | 25 | 11.56 | 18.26 | 7.65 | 5.03 | 
| S3 Tokenizer(FSQ) (Du et al., 2024a) | 6561 | 25 | 10.67 | **7.29** | 6.58 | 4.43 | 
| Qwen-TTS-Tokenizer-25Hz (Stage 1) | 32768 | 25 | **7.51** | 10.73 | **3.07** | **4.23** | 
| Qwen-TTS-Tokenizer-25Hz (Stage 2) | 32768 | 25 | 10.40 | 14.99 | 4.14 | 4.67 | 

*Comparison between different semantic-related speech tokenizers.*

| Model | NQ | Codebook Size | FPS | PESQ_WB | PESQ_NB | STOI | UTMOS | SIM | 
|---|---|---|---|---|---|---|---|---|
| SpeechTokenizer (Zhang et al., 2023a) | 8 | 1024 | 50 | 2.60 | 3.05 | 0.92 | 3.90 | 0.85 | 
| X-codec (Ye et al., 2025a) | 2 | 1024 | 50 | 2.68 | 3.27 | 0.86 | 4.11 | 0.84 | 
| X-codec 2 (Ye et al., 2025b) | 1 | 65536 | 50 | 2.43 | 3.04 | 0.92 | 4.13 | 0.82 | 
| XY-Tokenizer (Gong et al., 2025) | 8 | 1024 | 12.5 | 2.41 | 3.00 | 0.91 | 3.98 | 0.83 | 
| Mimi (Défossez et al., 2024) | 16 | 2048 | 12.5 | 2.88 | 3.42 | 0.94 | 3.87 | 0.87 | 
| FireredTTS 2 Tokenizer (Xie et al., 2025) | 16 | 2048 | 12.5 | 2.73 | 3.28 | 0.94 | 3.88 | 0.87 | 
| Qwen-TTS-Tokenizer-12Hz | 16 | 2048 | 12.5 | **3.21** | **3.68** | **0.96** | **4.16** | **0.95** | 

If you find our paper and code useful in your research, please consider giving a star :star: and citation :pencil: :)

```
@article{Qwen3-TTS,
  title={Qwen3-TTS Technical Report},
  author={Hangrui Hu and Xinfa Zhu and Ting He and Dake Guo and Bin Zhang and Xiong Wang and Zhifang Guo and Ziyue Jiang and Hongkun Hao and Zishan Guo and Xinyu Zhang and Pei Zhang and Baosong Yang and Jin Xu and Jingren Zhou and Junyang Lin},
  journal={arXiv preprint arXiv:2601.15621},
  year={2026}
}
```
- Downloads last month
- 3,911,478
