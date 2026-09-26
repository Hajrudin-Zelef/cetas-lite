---
id: collect-240926-huggingface/huggingface/unsloth-gemma-4-12b-it-gguf-hugging-face-4
title: "Read our How to Run Gemma 4 12B Guide!"
domain: huggingface
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["decode", "fine-tuning", "gguf", "gpu", "llama", "llama.cpp", "multimodal", "omni", "reasoning", "safeguards", "training"]
source: docs/RAG/clean_en/huggingface/unsloth-gemma-4-12b-it-gguf-hugging-face.md
source_anchor: ""
source_lines: [386, 454]
sha256: 23fbf16dd945dde6f0a10e0f6a5dec8ba21cd2fe046fdb42bf9db53a1cc82a42
---

# Read our How to Run Gemma 4 12B Guide!

- **Training Data**  - The quality and diversity of the training data significantly influence the model's capabilities. Biases or gaps in the training data can lead to limitations in the model's responses.
  - The scope of the training dataset determines the subject areas the model can handle effectively.
- **Context and Task Complexity**  - Models perform well on tasks that can be framed with clear prompts and instructions. Open-ended or highly complex tasks might be challenging.
  - A model's performance can be influenced by the amount of context provided (longer context generally leads to better outputs, up to a certain point).
- **Language Ambiguity and Nuance**  - Natural language is inherently complex. Models might struggle to grasp subtle nuances, sarcasm, or figurative language.
- **Factual Accuracy**  - Models generate responses based on information they learned from their training datasets, but they are not knowledge bases. They may generate incorrect or outdated factual statements.
- **Common Sense**  - Models rely on statistical patterns in language. They might lack the ability to apply common sense reasoning in certain situations.

The development of vision-language models (VLMs) raises several ethical concerns. In creating an open model, we have carefully considered the following:

- **Bias and Fairness**  - VLMs trained on large-scale, real-world text and image data can reflect socio-cultural biases embedded in the training material. Gemma 4 models underwent careful scrutiny, input data pre-processing, and post-training evaluations as reported in this card to help mitigate the risk of these biases.
- **Misinformation and Misuse**  - VLMs can be misused to generate text that is false, misleading, or harmful.
  - Guidelines are provided for responsible use with the model, see the Responsible Generative AI Toolkit.
- **Transparency and Accountability**  - This model card summarizes details on the models' architecture, capabilities, limitations, and evaluation processes.
  - A responsibly developed open model offers the opportunity to share innovation by making VLM technology accessible to developers and researchers across the AI ecosystem.

**Risks identified and mitigations**:

- **Generation of harmful content** : Mechanisms and guidelines for content safety are essential. Developers are encouraged to exercise caution and implement appropriate content safety safeguards based on their specific product policies and application use cases.
- **Misuse for malicious purposes** : Technical limitations and developer and end-user education can help mitigate against malicious applications of VLMs. Educational resources and reporting mechanisms for users to flag misuse are provided.
- **Privacy violations** : Models were trained on data filtered for removal of certain personal information and other sensitive data. Developers are encouraged to adhere to privacy regulations with privacy-preserving techniques.
- **Perpetuation of biases** : It's encouraged to perform continuous monitoring (using evaluation metrics, human review) and the exploration of de-biasing techniques during model training, fine-tuning, and other use cases.

At the time of release, this family of models provides high-performance open vision-language model implementations designed from the ground up for responsible AI development compared to similarly sized models.

This is an omni GGUF, so the same files handle text, images, and audio. Grab any recent stock llama.cpp build and start the server. The multimodal projector (mmproj) is downloaded automatically when you use `-hf`, so you do not need to pass it yourself:

```
llama-server -hf unsloth/gemma-4-12b-it-GGUF:UD-Q4_K_XL --jinja -c 8192
# add -ngl 999 if you have a GPU build
```
Then query it through the OpenAI compatible API:

```
import json, base64, urllib.request
def ask(content, max_tokens=256):
    body = {
        "messages": [{"role": "user", "content": content}],
        "max_tokens": max_tokens,
        # Gemma 4 is a thinking model. Set this to False (or raise max_tokens),
        # otherwise the reply lands in reasoning_content and "content" is empty.
        "chat_template_kwargs": {"enable_thinking": False},
    }
    req = urllib.request.Request("http://127.0.0.1:8080/v1/chat/completions",
                                 json.dumps(body).encode(),
                                 {"Content-Type": "application/json"})
    return json.loads(urllib.request.urlopen(req).read())["choices"][0]["message"]["content"]
b64 = lambda p: base64.b64encode(open(p, "rb").read()).decode()
# Text
print(ask("What is 1+1?"))
# Vision (any image file)
print(ask([
    {"type": "text", "text": "What is in this image?"},
    {"type": "image_url", "image_url": {"url": "data:image/jpeg;base64," + b64("image.jpg")}},
]))
# Audio (16 kHz mono WAV works best)
print(ask([
    {"type": "text", "text": "Transcribe this audio."},
    {"type": "input_audio", "input_audio": {"data": b64("audio.wav"), "format": "wav"}},
]))
```
Tips:

- Pass `--jinja` so the Gemma 4 chat template is applied.
- For audio, feed a 16 kHz mono WAV (convert with `ffmpeg -i in.mp3 -ar 16000 -ac 1 out.wav` ). Clean speech transcribes best.
- To force a specific projector precision add `--mmproj-url .../mmproj-F16.gguf` , or pass`--no-mmproj` to disable multimodal.

- Downloads last month
- 1,167,254
