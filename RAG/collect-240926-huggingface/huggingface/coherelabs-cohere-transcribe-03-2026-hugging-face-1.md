---
id: collect-240926-huggingface/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face-1
title: "coherelabs-cohere-transcribe-03-2026-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Cohere", "vLLM"]
dates: []
keywords: ["cohere", "apache", "decode", "inference", "leaderboard", "license", "open source", "training", "transcription", "vllm"]
source: docs/RAG/clean_en/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face.md
source_anchor: ""
source_lines: [1, 163]
sha256: 58f6cf355f879efe7d7c06388180e6e3fa19f8770f4a7bfe2e80dbf59e4f75f8
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

