---
id: collect-240926-huggingface/huggingface/google-gemma-4-26b-a4b-it-qat-q4-0-gguf-hugging-face-3
title: "Load model"
domain: huggingface
role: reference
task: reference
actors: ["Google"]
dates: ["2025-01"]
keywords: ["compute", "cost", "decode", "gemini", "inference", "llama", "llama.cpp", "multimodal", "reasoning", "refusals", "training", "transcription"]
source: docs/RAG/clean_en/huggingface/google-gemma-4-26b-a4b-it-qat-q4-0-gguf-hugging-face.md
source_anchor: ""
source_lines: [266, 373]
sha256: ba2383a76d9f45661a0afd726c26b31ec8cc6c88cf743a6938e16a3feb0e23ca
---

# Load model

```
# Prompt - add video before text
messages = [
    {
        'role': 'user',
        'content': [
            {"type": "video", "video": "https://github.com/bebechien/gemma/raw/refs/heads/main/videos/ForBiggerBlazes.mp4"},
            {'type': 'text', 'text': 'Describe this video.'}
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
For the best performance, use these configurations and best practices:

Use the following standardized sampling configuration across all use cases:

- `temperature=1.0`
- `top_p=0.95`
- `top_k=64`

Compared to Gemma 3, the models use standard `system`, `assistant`, and `user` roles. To properly manage the thinking process, use the following control tokens:

- **Trigger Thinking:** Thinking is enabled by including the`<|think|>` token at the start of the system prompt. To disable thinking, remove the token.
- **Standard Generation:** When thinking is enabled, the model will output its internal reasoning followed by the final answer using this structure:`<|channel>thought\n`**[Internal reasoning]**`<channel|>`
- **Disabled Thinking Behavior:** For all models except for the E2B and E4B variants, if thinking is disabled, the model will still generate the tags but with an empty thought block:`<|channel>thought\n<channel|>`**[Final answer]**

Note that many libraries like Transformers and llama.cpp handle the complexities of the chat template for you.


- **No Thinking Content in History** : In multi-turn conversations, the historical model output should only include the final response. Thoughts from previous model turns must*not be added* before the next user turn begins, with the exception of tool call turns where thinking content should be preserved.

For optimal performance with multimodal inputs, place:

- Image content **before** the text in your prompt.
- Audio content **after** the text in your prompt.

Aside from variable aspect ratios, Gemma 4 supports variable image resolution through a configurable visual token budget, which controls how many tokens are used to represent an image. A higher token budget preserves more visual detail at the cost of additional compute, while a lower budget enables faster inference for tasks that don't require fine-grained understanding.

- The supported token budgets are: **70** ,**140** ,**280** ,**560** , and**1120** .
  - Use *lower budgets* for classification, captioning, or video understanding, where faster inference and processing many frames outweigh fine-grained detail.
  - Use *higher budgets* for tasks like OCR, document parsing, or reading small text.
- Use 

Use the following prompt structures for audio processing:

- **Audio Speech Recognition (ASR)**

```
Transcribe the following speech segment in {LANGUAGE} into {LANGUAGE} text.
Follow these specific instructions for formatting the answer:
* Only output the transcription, with no newlines.
* When transcribing numbers, write the digits, i.e. write 1.7 and not one point seven, and write 3 instead of three.
```
- **Automatic Speech Translation (AST)**

```
Transcribe the following speech segment in {SOURCE_LANGUAGE}, then translate it into {TARGET_LANGUAGE}.
When formatting the answer, first output the transcription in {SOURCE_LANGUAGE}, then one newline, then output the string '{TARGET_LANGUAGE}: ', then the translation in {TARGET_LANGUAGE}.
```
All models support image inputs and can process videos as frames whereas the E2B, E4B, and 12B models also support audio inputs. Audio supports a maximum length of 30 seconds. Video supports a maximum of 60 seconds assuming the images are processed at one frame per second.

Data used for model training and how the data was processed.

Our pre-training dataset is a large-scale, diverse collection of data encompassing a wide range of domains and modalities, which includes web documents, code, images, audio, with a cutoff date of January 2025. Here are the key components:

- **Web Documents** : A diverse collection of web text ensures the model is exposed to a broad range of linguistic styles, topics, and vocabulary. The training dataset includes content in over 140 languages.
- **Code** : Exposing the model to code helps it to learn the syntax and patterns of programming languages, which improves its ability to generate code and understand code-related questions.
- **Mathematics** : Training on mathematical text helps the model learn logical reasoning, symbolic representation, and to address mathematical queries.
- **Images** : A wide range of images enables the model to perform image analysis and visual data extraction tasks.

The combination of these diverse data sources is crucial for training a powerful multimodal model that can handle a wide variety of different tasks and data formats.

Here are the key data cleaning and filtering methods applied to the training data:

- **CSAM Filtering** : Rigorous CSAM (Child Sexual Abuse Material) filtering was applied at multiple stages in the data preparation process to ensure the exclusion of harmful and illegal content.
- **Sensitive Data Filtering** : As part of making Gemma pre-trained models safe and reliable, automated techniques were used to filter out certain personal information and other sensitive data from training sets.
- **Additional methods** : Filtering based on content quality and safety in line with our policies.

As open models become central to enterprise infrastructure, provenance and security are paramount. Developed by Google DeepMind, Gemma 4 undergoes the same rigorous safety evaluations as our proprietary Gemini models.

Gemma 4 models were developed in partnership with internal safety and responsible AI teams. A range of automated as well as human evaluations were conducted to help improve model safety. These evaluations align with Google’s AI principles, as well as safety policies, which aim to prevent our generative AI models from generating harmful content, including:

- Content related to child sexual abuse material and exploitation
- Dangerous content (e.g., promoting suicide, or instructing in activities that could cause real-world harm)
- Sexually explicit content
- Hate speech (e.g., dehumanizing members of protected groups)
- Harassment (e.g., encouraging violence against people)

For all areas of safety testing, we saw major improvements in all categories of content safety relative to previous Gemma models. Overall, Gemma 4 models significantly outperform Gemma 3 and 3n models in improving safety, while keeping unjustified refusals low. All testing was conducted without safety filters to evaluate the model capabilities and behaviors. For both text-to-text and image-to-text, and across all model sizes, the model produced minimal policy violations, and showed significant improvements over previous Gemma models' performance.

These models have certain limitations that users should be aware of.

Multimodal models (capable of processing vision, language, and/or audio) have a wide range of applications across various industries and domains. The following list of potential uses is not comprehensive. The purpose of this list is to provide contextual information about the possible use-cases that the model creators considered as part of model training and development.

