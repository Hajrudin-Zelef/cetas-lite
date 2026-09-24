---
id: collect-240926-huggingface/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face
title: "coherelabs-cohere-transcribe-03-2026-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Apple", "Cohere", "EU", "Hugging Face", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["cohere", "apache", "benchmarks", "decode", "distribution", "inference", "leaderboard", "license", "nvidia", "open source", "qwen", "research"]
source: docs/RAG/clean_en/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face.md
source_anchor: ""
source_lines: [1, 238]
sha256: 515e1807d4eb9381b31f04249fb779b3445bd0aebabe057f8a442d867923e111
---

# coherelabs-cohere-transcribe-03-2026-hugging-face

<!-- source: https://huggingface.co/CohereLabs/cohere-transcribe-03-2026 -->

Cohere Transcribe is an open source release of a 2B parameter dedicated audio-in, text-out automatic speech recognition (ASR) model. The model supports 14 languages.

Developed by: Cohere and Cohere Labs. Point of Contact: Cohere Labs.

| Name | **cohere-transcribe-03-2026** | 
|---|---|
| Architecture | conformer-based encoder-decoder | 
| Input | audio waveform → log-Mel spectrogram. Audio is automatically resampled to 16kHz if necessary during preprocessing. Similarly, multi-channel (stereo) inputs are averaged to produce a single channel signal. | 
| Output | transcribed text | 
| Model size | 2B | 
| Model | a large Conformer encoder extracts acoustic representations, followed by a lightweight Transformer decoder for token generation | 
| Training objective | supervised cross-entropy on output tokens; trained from scratch | 
| Languages | Trained on 14 languages:  | 
| License | Apache 2.0 | 

✨**Try the Cohere Transcribe** **demo**✨

Cohere Transcribe is supported natively in `transformers`. This is the recommended way to use the model for 
offline inference. For online inference, see the vLLM integration example below.

```
pip install transformers>=5.4.0 torch huggingface_hub soundfile librosa sentencepiece protobuf accelerate
pip install datasets  # only needed for long-form and non-English examples
```
Testing was carried out with `torch==2.10.0` but it is expected to work with other versions.

Transcribe any audio file in a few lines:

```
from transformers import AutoProcessor, CohereAsrForConditionalGeneration
from transformers.audio_utils import load_audio
from huggingface_hub import hf_hub_download
processor = AutoProcessor.from_pretrained("CohereLabs/cohere-transcribe-03-2026")
model = CohereAsrForConditionalGeneration.from_pretrained("CohereLabs/cohere-transcribe-03-2026", device_map="auto")
audio_file = hf_hub_download(
    repo_id="CohereLabs/cohere-transcribe-03-2026",
    filename="demo/voxpopuli_test_en_demo.wav",
)
audio = load_audio(audio_file, sampling_rate=16000)
inputs = processor(audio, sampling_rate=16000, return_tensors="pt", language="en")
inputs.to(model.device, dtype=model.dtype)
outputs = model.generate(**inputs, max_new_tokens=256)
text = processor.decode(outputs, skip_special_tokens=True)
print(text)
```
## **Long-form transcription**

For audio longer than the feature extractor's `max_audio_clip_s`, the feature extractor automatically splits the waveform into chunks.
The processor reassembles the per-chunk transcriptions using the returned `audio_chunk_index`.

This example transcribes a 55 minute earnings call:

```
from transformers import AutoProcessor, CohereAsrForConditionalGeneration
from datasets import load_dataset
import time
processor = AutoProcessor.from_pretrained("CohereLabs/cohere-transcribe-03-2026")
model = CohereAsrForConditionalGeneration.from_pretrained("CohereLabs/cohere-transcribe-03-2026", device_map="auto")
ds = load_dataset("distil-whisper/earnings22", "full", split="test", streaming=True)
sample = next(iter(ds))
audio_array = sample["audio"]["array"]
sr = sample["audio"]["sampling_rate"]
duration_s = len(audio_array) / sr
print(f"Audio duration: {duration_s / 60:.1f} minutes")
inputs = processor(audio=audio_array, sampling_rate=sr, return_tensors="pt", language="en")
audio_chunk_index = inputs.get("audio_chunk_index")
inputs.to(model.device, dtype=model.dtype)
start = time.time()
outputs = model.generate(**inputs, max_new_tokens=256)
text = processor.decode(outputs, skip_special_tokens=True, audio_chunk_index=audio_chunk_index, language="en")[0]
elapsed = time.time() - start
rtfx = duration_s / elapsed
print(f"Transcribed in {elapsed:.1f}s — RTFx: {rtfx:.1f}")
print(f"Transcription ({len(text.split())} words):")
print(text[:500] + "...")
```
## **Punctuation control**

Pass `punctuation=False` to obtain lower-cased output without punctuation marks.

```
inputs_pnc = processor(audio, sampling_rate=16000, return_tensors="pt", language="en", punctuation=True)
inputs_nopnc = processor(audio, sampling_rate=16000, return_tensors="pt", language="en", punctuation=False)
```
By default, punctuation is enabled.

## **Batched inference**

Multiple audio files can be processed in a single call. When the batch mixes short-form and long-form audio, the processor handles chunking and reassembly.

```
from transformers import AutoProcessor, CohereAsrForConditionalGeneration
from transformers.audio_utils import load_audio
processor = AutoProcessor.from_pretrained("CohereLabs/cohere-transcribe-03-2026")
model = CohereAsrForConditionalGeneration.from_pretrained("CohereLabs/cohere-transcribe-03-2026", device_map="auto")
audio_short = load_audio(
    "https://huggingface.co/datasets/hf-internal-testing/dummy-audio-samples/resolve/main/bcn_weather.mp3",
    sampling_rate=16000,
)
audio_long = load_audio(
    "https://huggingface.co/datasets/hf-internal-testing/dummy-audio-samples/resolve/main/obama_first_45_secs.mp3",
    sampling_rate=16000,
)
inputs = processor([audio_short, audio_long], sampling_rate=16000, return_tensors="pt", language="en")
audio_chunk_index = inputs.get("audio_chunk_index")
inputs.to(model.device, dtype=model.dtype)
outputs = model.generate(**inputs, max_new_tokens=256)
text = processor.decode(
    outputs, skip_special_tokens=True, audio_chunk_index=audio_chunk_index, language="en"
)
print(text)
```
## **Non-English transcription**

Specify the language code to transcribe in any of the 14 supported languages. This example transcribes Japanese audio from the FLEURS dataset:

```
from transformers import AutoProcessor, CohereAsrForConditionalGeneration
from datasets import load_dataset
processor = AutoProcessor.from_pretrained("CohereLabs/cohere-transcribe-03-2026")
model = CohereAsrForConditionalGeneration.from_pretrained("CohereLabs/cohere-transcribe-03-2026", device_map="auto")
ds = load_dataset("google/fleurs", "ja_jp", split="test", streaming=True)
ds_iter = iter(ds)
samples = [next(ds_iter) for _ in range(3)]
for sample in samples:
    audio = sample["audio"]["array"]
    sr = sample["audio"]["sampling_rate"]
    inputs = processor(audio, sampling_rate=sr, return_tensors="pt", language="ja")
    inputs.to(model.device, dtype=model.dtype)
    outputs = model.generate(**inputs, max_new_tokens=256)
    text = processor.decode(outputs, skip_special_tokens=True)
    print(f"REF: {sample['transcription']}\nHYP: {text}\n")
```
For production serving we recommend running via vLLM following the instructions below.

## **Run cohere-transcribe-03-2026 via vLLM**

First install vLLM (refer to vLLM installation instructions):

```
uv venv --python 3.12 --seed
source .venv/bin/activate
uv pip install -U vllm==0.19.0 --torch-backend=auto
uv pip install vllm[audio]
uv pip install librosa
```
Start vLLM server

```
vllm serve CohereLabs/cohere-transcribe-03-2026 --trust-remote-code
```
Send request

```
curl -v -X POST http://localhost:8000/v1/audio/transcriptions \
 -H "Authorization: Bearer $VLLM_API_KEY" \
-F "file=@$(realpath ${AUDIO_PATH})" \
-F "model=CohereLabs/cohere-transcribe-03-2026"
```
## **English ASR Leaderboard (as of 03.26.2026)**

| Model | Average WER | AMI | Earnings 22 | Gigaspeech | LS clean | LS other | SPGISpeech | Tedlium | Voxpopuli | 
|---|---|---|---|---|---|---|---|---|---|
| **Cohere Transcribe** | **5.42** | **8.15** | 10.84 | 9.33 | **1.25** | **2.37** | 3.08 | 2.49 | 5.87 | 
| Zoom Scribe v1 | 5.47 | 10.03 | 9.53 | 9.61 | 1.63 | 2.81 | **1.59** | 3.22 | **5.37** | 
| IBM Granite 4.0 1B Speech | 5.52 | 8.44 | **8.48** | 10.14 | 1.42 | 2.85 | 3.89 | 3.10 | 5.84 | 
| NVIDIA Canary Qwen 2.5B | 5.63 | 10.19 | 10.45 | 9.43 | 1.61 | 3.10 | 1.90 | 2.71 | 5.66 | 
| Qwen3-ASR-1.7B | 5.76 | 10.56 | 10.25 | **8.74** | 1.63 | 3.40 | 2.84 | **2.28** | 6.35 | 
| ElevenLabs Scribe v2 | 5.83 | 11.86 | 9.43 | 9.11 | 1.54 | 2.83 | 2.68 | 2.37 | 6.80 | 
| Kyutai STT 2.6B | 6.40 | 12.17 | 10.99 | 9.81 | 1.70 | 4.32 | 2.03 | 3.35 | 6.79 | 
| OpenAI Whisper Large v3 | 7.44 | 15.95 | 11.29 | 10.02 | 2.01 | 3.91 | 2.94 | 3.86 | 9.54 | 
| Voxtral Mini 4B Realtime 2602 | 7.68 | 17.07 | 11.84 | 10.38 | 2.08 | 5.52 | 2.42 | 3.79 | 8.34 | 

Link to the live leaderboard: Open ASR Leaderboard.

We observe similarly strong performance in human evaluations, where trained annotators assess transcription quality across real-world audio for accuracy, coherence and usability. The consistency between automated metrics and human judgments suggests that the model’s improvements translate beyond controlled benchmarks to practical transcription settings.

*Figure: Human preference evaluation of model transcripts. In a head-to-head comparison, 
annotators were asked to express preferences for generations which primarily preserved meaning - 
but also avoided hallucination, correctly identified named entities, 
and provided verbatim transcripts with appropriate formatting. 
A score of 50% or higher indicates that Cohere Transcribe was preferred on average in the comparison.*

## **per-language WERs**

*Figure: per-language error rate averaged over FLEURS, Common Voice 17.0, MLS and Wenet tests sets (where relevant for a given language). CER for zh, ja, ko — WER otherwise*

For more details and results:

- Technical blog post contains WERs and other quality metrics.
- Announcement blog post for more information about the model.
- English, EU and long-form transcription WERs/RTFx are on the Open ASR Leaderboard.

Cohere Transcribe is a performant, dedicated ASR model intended for efficient speech transcription.

Cohere Transcribe demonstrates best-in-class transcription accuracy in 14 languages. As a dedicated speech recognition model, it is also efficient, benefitting from a real-time factor up to three times faster than that of other, dedicated ASR models in the same size range. The model was trained from scratch, and from the outset, we deliberately focused on maximizing transcription accuracy while keeping production readiness top-of-mind.

- **Single language.** The model performs best when remaining in-distribution of a single, pre-specified language amongst the 14 in the range it supports. It does not feature explicit, automatic language detection and exhibits inconsistent performance on code-switched audio.
- **Timestamps/Speaker diarization.** The model does not feature either of these.
- **Silence.** Like most AED speech models, Cohere Transcribe is eager to transcribe, even non-speech sounds. The model thus benefits from prepending a noise gate or VAD (voice activity detection) model in order to prevent low-volume, floor noise from turning into hallucinations.

Cohere Transcribe is supported on the following libraries/platforms:

- `transformers` (see Quick Start above).
- `vLLM` (see vLLM integration above).
- `mlx-audio` for Apple Silicon.
- Rust implementation: `cohere_transcribe_rs`
- In the browser ✨**demo** ✨ (via`transformers.js` and WebGPU)
- Chrome extension: `cohere_transcribe_extension`
- `nano-cohere-transcribe` - blazing fast, particularly for long-form audio. e.g. 18 mins audio transcribed in 2.36s
- Whisper Memos (iOS App).
- Whisperian (Android App).
- hyprwhspr (Linux app).

If you have added support for the model somewhere not included above please raise an issue/PR!

If you find issues with any of these please raise an issue with the respective library.

For errors or additional questions about details in this model card, contact labs@cohere.com or raise an issue.

Terms of Use: We hope that the release of this model will make community-based research efforts more accessible, by releasing the weights of a highly performant 2 billion parameter model to researchers all over the world. This model is governed by an Apache 2.0 license.

To cite this model please use the following bibtex:

```
@misc{julian_mack_2026,
    author       = { Julian Mack and Ekagra Ranjan and Walter Beller-Morales and Bharat Venkitesh and Pierre Richemond },
    title        = { cohere-transcribe-03-2026 (Revision d96e814) },
    year         = 2026,
    url          = { https://huggingface.co/CohereLabs/cohere-transcribe-03-2026 },
    doi          = { 10.57967/hf/8653 },
    publisher    = { Hugging Face }
}
```
- Downloads last month
- 211,468
