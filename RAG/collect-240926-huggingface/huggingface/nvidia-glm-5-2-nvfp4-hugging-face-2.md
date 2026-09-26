---
id: collect-240926-huggingface/huggingface/nvidia-glm-5-2-nvfp4-hugging-face-2
title: "nvidia-glm-5-2-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia"]
dates: []
keywords: ["nvidia", "attention", "license", "mit license", "reasoning", "throughput", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-glm-5-2-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [139, 178]
sha256: 9068801fa84eaea7956682cbdbc7fb7536a05c9aa6ab20a0182bb4d94c53ef64
---

# nvidia-glm-5-2-nvfp4-hugging-face

| Field: | Response: | 
|---|---|
| Intended Task/Domain: | Text generation, reasoning, summarization, and question answering. | 
| Model Type: | Text and Image-to-text transformer | 
| Intended Users: | This model is intended for developers, researchers, and customers building/utilizing LLMs, while balancing accuracy and efficiency. | 
| Output: | Text String(s) | 
| Describe how the model works: | Generates text by predicting the next word or token based on the context provided in the input sequence using multiple self-attention layers | 
| Name the adversely impacted groups this has been tested to deliver comparable outcomes regardless of: | Not Applicable | 
| Technical Limitations & Mitigation: | The model was trained on data that contains toxic language and societal biases originally crawled from the internet. Therefore, the model may amplify those biases and return toxic responses especially when prompted with toxic prompts. Therefore, before deploying any applications of this model, developers should perform safety testing and tuning tailored to their specific applications of the model. | 
| Verified to have met prescribed quality standards? | Yes | 
| Performance Metrics: | Accuracy, Throughput, and user-side throughput | 
| Potential Known Risk | The model may generate answers that may be inaccurate, omit key information, or include irrelevant or redundant text producing socially unacceptable or undesirable text, even if the prompt itself does not include anything explicitly offensive. | 
| Licensing: | Your usage is governed by the following **GOVERNING TERMS:** Use of the model is governed by the MIT License, same as the base model. | 

| Field: | Response: | 
|---|---|
| Participation considerations from adversely impacted groups protected classes in model design and testing: | None | 
| Measures taken to mitigate against unwanted bias: | None | 

| Field: | Response: | 
|---|---|
| Model Application(s): | Chat, Instruction Following, Chatbot Development, Code Generation, Reasoning | 
| Describe life critical application (if present): | Not Applicable | 
| Use Case Restrictions: | Abide by the **GOVERNING TERMS:** Use of the model is governed by the MIT License, same as the base model. | 
| Model and Dataset Restrictions: | The Principle of least privilege (PoLP) is applied limiting access for dataset generation. Restrictions enforce dataset access during training, and dataset license constraints adhered to. Model checkpoints are made available on Hugging Face, and may become available on cloud providers' model catalog. | 

| Field: | Response: | 
|---|---|
| Generatable or Reverse engineerable personal data? | No | 
| Personal data used to create this model? | No | 
| Was consent obtained for any personal data used? | Not Applicable | 
| How often is dataset reviewed? | Before Release | 
| Was data from user interactions with the AI model (e.g. user input and prompts) used to train the model? | No | 
| Is there provenance for all datasets used in training? | Yes | 
| Does data labeling (annotation, metadata) comply with privacy laws? | Yes | 
| Is data compliant with data subject requests for data correction or removal, if such a request was made? | Not Applicable | 
| Applicable NVIDIA Privacy Policy | https://www.nvidia.com/en-us/about-nvidia/privacy-policy/ | 

- Downloads last month
- 705,370
