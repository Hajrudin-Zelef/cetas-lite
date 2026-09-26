---
id: collect-240926-huggingface/huggingface/qwen-qwen3-tts-12hz-1-7b-base-hugging-face-3
title: "Download through ModelScope (recommended for users in Mainland China)"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "MiniMax", "vLLM"]
dates: []
keywords: ["attention", "benchmarks", "decode", "inference", "omni", "parameters", "qwen", "training", "vllm", "voice"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-tts-12hz-1-7b-base-hugging-face.md
source_anchor: ""
source_lines: [201, 359]
sha256: 84fb223cc2f80e4cf6683962f35933f910ff80d3f5f9c47dee1285ef6fec8364
---

# Download through ModelScope (recommended for users in Mainland China)

```
import torch
import soundfile as sf
from qwen_tts import Qwen3TTSModel
# create a reference audio in the target style using the VoiceDesign model
design_model = Qwen3TTSModel.from_pretrained(
    "Qwen/Qwen3-TTS-12Hz-1.7B-VoiceDesign",
    device_map="cuda:0",
    dtype=torch.bfloat16,
    attn_implementation="flash_attention_2",
)
ref_text = "H-hey! You dropped your... uh... calculus notebook? I mean, I think it's yours? Maybe?"
ref_instruct = "Male, 17 years old, tenor range, gaining confidence - deeper breath support now, though vowels still tighten when nervous"
ref_wavs, sr = design_model.generate_voice_design(
    text=ref_text,
    language="English",
    instruct=ref_instruct
)
sf.write("voice_design_reference.wav", ref_wavs[0], sr)
# build a reusable clone prompt from the voice design reference
clone_model = Qwen3TTSModel.from_pretrained(
    "Qwen/Qwen3-TTS-12Hz-1.7B-Base",
    device_map="cuda:0",
    dtype=torch.bfloat16,
    attn_implementation="flash_attention_2",
)
voice_clone_prompt = clone_model.create_voice_clone_prompt(
    ref_audio=(ref_wavs[0], sr),   # or "voice_design_reference.wav"
    ref_text=ref_text,
)
sentences = [
    "No problem! I actually... kinda finished those already? If you want to compare answers or something...",
    "What? No! I mean yes but not like... I just think you're... your titration technique is really precise!",
]
# reuse it for multiple single calls
wavs, sr = clone_model.generate_voice_clone(
    text=sentences[0],
    language="English",
    voice_clone_prompt=voice_clone_prompt,
)
sf.write("clone_single_1.wav", wavs[0], sr)
wavs, sr = clone_model.generate_voice_clone(
    text=sentences[1],
    language="English",
    voice_clone_prompt=voice_clone_prompt,
)
sf.write("clone_single_2.wav", wavs[0], sr)
# or batch generate in one call
wavs, sr = clone_model.generate_voice_clone(
    text=sentences,
    language=["English", "English"],
    voice_clone_prompt=voice_clone_prompt,
)
for i, w in enumerate(wavs):
    sf.write(f"clone_batch_{i}.wav", w, sr)
```
If you only want to encode and decode audio for transport or training and so on, `Qwen3TTSTokenizer` supports encode/decode with paths, URLs, numpy waveforms, and dict/list payloads, for example:

```
import soundfile as sf
from qwen_tts import Qwen3TTSTokenizer
tokenizer = Qwen3TTSTokenizer.from_pretrained(
    "Qwen/Qwen3-TTS-Tokenizer-12Hz",
    device_map="cuda:0",
)
enc = tokenizer.encode("https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen3-TTS-Repo/tokenizer_demo_1.wav")
wavs, sr = tokenizer.decode(enc)
sf.write("decode_output.wav", wavs[0], sr)
```
For more tokenizer examples (including different input formats and batch usage), please refer to the example codes. With those examples and the description for `Qwen3TTSTokenizer`, you can explore more advanced usage patterns.

To launch the Qwen3-TTS web ui demo, simply install the `qwen-tts` package and run `qwen-tts-demo`. Use the command below for help:

```
qwen-tts-demo --help
```
To launch the demo, you can use the following commands:

```
# CustomVoice model
qwen-tts-demo Qwen/Qwen3-TTS-12Hz-1.7B-CustomVoice --ip 0.0.0.0 --port 8000
# VoiceDesign model
qwen-tts-demo Qwen/Qwen3-TTS-12Hz-1.7B-VoiceDesign --ip 0.0.0.0 --port 8000
# Base model
qwen-tts-demo Qwen/Qwen3-TTS-12Hz-1.7B-Base --ip 0.0.0.0 --port 8000
```
And then open `http://<your-ip>:8000`, or access it via port forwarding in tools like VS Code.

To avoid browser microphone permission issues after deploying the server, for Base model deployments, it is recommended/required to run the gradio service over **HTTPS** (especially when accessed remotely or behind modern browsers/gateways). Use `--ssl-certfile` and `--ssl-keyfile` to enable HTTPS. First we need to generate a private key and a self-signed cert (valid for 365 days):

```
openssl req -x509 -newkey rsa:2048 \
  -keyout key.pem -out cert.pem \
  -days 365 -nodes \
  -subj "/CN=localhost"
```
Then run the demo with HTTPS:

```
qwen-tts-demo Qwen/Qwen3-TTS-12Hz-1.7B-Base \
  --ip 0.0.0.0 --port 8000 \
  --ssl-certfile cert.pem \
  --ssl-keyfile key.pem \
  --no-ssl-verify
```
And open `https://<your-ip>:8000` to experience it. If your browser shows a warning, it’s expected for self-signed certificates. For production, use a real certificate.

To further explore Qwen3-TTS, we encourage you to try our DashScope API for a faster and more efficient experience. For detailed API information and documentation, please refer to the following:

| API Description | API Documentation (Mainland China) | API Documentation (International) | 
|---|---|---|
| Real-time API for Qwen3-TTS of custom voice model. | https://help.aliyun.com/zh/model-studio/qwen-tts-realtime | https://www.alibabacloud.com/help/en/model-studio/qwen-tts-realtime | 
| Real-time API for Qwen3-TTS of voice clone model. | https://help.aliyun.com/zh/model-studio/qwen-tts-voice-cloning | https://www.alibabacloud.com/help/en/model-studio/qwen-tts-voice-cloning | 
| Real-time API for Qwen3-TTS of voice design model. | https://help.aliyun.com/zh/model-studio/qwen-tts-voice-design | https://www.alibabacloud.com/help/en/model-studio/qwen-tts-voice-design | 

vLLM officially provides day-0 support for Qwen3-TTS! Welcome to use vLLM-Omni for Qwen3-TTS deployment and inference. For installation and more details, please check vLLM-Omni official documentation. Now only offline inference is supported. Online serving will be supported later, and vLLM-Omni will continue to offer support and optimization for Qwen3-TTS in areas such as inference speed and streaming capabilities.

You can use vLLM-Omni to inference Qwen3-TTS locally, we provide examples in vLLM-Omni repo which can generate audio output:

```
# git clone https://github.com/vllm-project/vllm-omni.git
# cd vllm-omni/examples/offline_inference/qwen3_tts
# Run a single sample with CustomVoice task
python end2end.py --query-type CustomVoice
# Batch sample (multiple prompts in one run) with CustomVoice task:
python end2end.py --query-type CustomVoice --use-batch-sample
# Run a single sample with VoiceDesign task
python end2end.py --query-type VoiceDesign
# Batch sample (multiple prompts in one run) with VoiceDesign task:
python end2end.py --query-type VoiceDesign --use-batch-sample
# Run a single sample with Base task in icl mode-tag
python end2end.py --query-type Base --mode-tag icl
```
During evaluation, we ran inference for all models with `dtype=torch.bfloat16` and set `max_new_tokens=2048`. All other sampling parameters used the defaults from the checkpoint’s `generate_config.json`. For the Seed-Test and InstructTTS-Eval test sets, we set `language="auto"`, while for all other test sets we explicitly passed the corresponding `language`. The detailed results are shown below.

## Speech Generation Benchmarks

*Zero-shot speech generation on the Seed-TTS test set. Performance is measured by Word Error Rate (WER, ↓), where lower is better.*

| Datasets | Model | Performance |  | 
|---|---|---|---|
| *Content Consistency* |  |  |  | 
| SEED *test-zh* \|*test-en* | Seed-TTS (Anastassiou et al., 2024) | 1.12 | 2.25 | 
|  | MaskGCT (Wang et al., 2024) | 2.27 | 2.62 | 
|  | E2 TTS (Eskimez et al., 2024) | 1.97 | 2.19 | 
|  | F5-TTS (Chen et al., 2024) | 1.56 | 1.83 | 
|  | Spark TTS (Wang et al., 2025) | 1.20 | 1.98 | 
|  | Llasa-8B (Ye et al., 2025b) | 1.59 | 2.97 | 
|  | KALL-E (Xia et al., 2024) | 0.96 | 1.94 | 
|  | FireRedTTS 2 (Xie et al., 2025) | 1.14 | 1.95 | 
|  | CosyVoice 3 (Du et al., 2025) | **0.71** | 1.45 | 
|  | MiniMax-Speech (Zhang et al., 2025a) | 0.83 | 1.65 | 
|  | Qwen3-TTS-25Hz-0.6B-Base | 1.18 | 1.64 | 
|  | Qwen3-TTS-25Hz-1.7B-Base | 1.10 | 1.49 | 
|  | Qwen3-TTS-12Hz-0.6B-Base | 0.92 | 1.32 | 
|  | Qwen3-TTS-12Hz-1.7B-Base | 0.77 | **1.24** | 

*Multilingual speech generation on the TTS multilingual test set. Performance is measured by Word Error Rate (WER, ↓) for content consistency and Cosine Similarity (SIM, ↑) for speaker similarity.*

