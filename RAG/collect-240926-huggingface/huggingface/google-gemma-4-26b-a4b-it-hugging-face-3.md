---
id: collect-240926-huggingface/huggingface/google-gemma-4-26b-a4b-it-hugging-face-3
title: "Load model"
domain: huggingface
role: reference
task: reference
actors: ["Google"]
dates: ["2025-01"]
keywords: ["compute", "cost", "gemini", "inference", "llama", "llama.cpp", "multimodal", "reasoning", "refusals", "research", "training", "transcription"]
source: docs/RAG/clean_en/huggingface/google-gemma-4-26b-a4b-it-hugging-face.md
source_anchor: ""
source_lines: [297, 373]
sha256: 9648c6664c1bcb8ec9bb0187bca3eeaed53c6cd981db27519dd38efb8cdd9c10
---

# Load model

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

- **Content Creation and Communication**  - **Text Generation** : These models can be used to generate creative text formats such as poems, scripts, code, marketing copy, and email drafts.
  - **Chatbots and Conversational AI** : Power conversational interfaces for customer service, virtual assistants, or interactive applications.
  - **Text Summarization** : Generate concise summaries of a text corpus, research papers, or reports.
  - **Image Data Extraction** : These models can be used to extract, interpret, and summarize visual data for text communications.
  - **Audio Processing and Interaction** : The E2B, E4B, and 12B models can analyze and interpret audio inputs, enabling voice-driven interactions and transcriptions.
- **Research and Education**  - **Natural Language Processing (NLP) and VLM Research** : These models can serve as a foundation for researchers to experiment with VLM and NLP techniques, develop algorithms, and contribute to the advancement of the field.
  - **Language Learning Tools** : Support interactive language learning experiences, aiding in grammar correction or providing writing practice.
  - **Knowledge Exploration** : Assist researchers in exploring large bodies of text by generating summaries or answering questions about specific topics.

