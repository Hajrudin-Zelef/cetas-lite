---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face-2
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "agentic", "attention", "data centre", "distribution", "gpu", "gpus", "inference", "latency", "license", "moe", "reasoning"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face.md
source_anchor: ""
source_lines: [91, 132]
sha256: 967437a7b16fc71ce174dbd1d27333b2bfe770e3844d5b96ffb0df6c168df869
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face

| Field | Response | 
|---|---|
| Intended Task/Domain | Text generation, reasoning, tool use, and agentic workflows accelerated with speculative decoding | 
| Model Type | Hybrid LatentMoE language model used with a DFlash draft checkpoint | 
| Intended Users | Developers deploying Nemotron-3.5-Lightning-30B-A3B for reasoning, chat, RAG, and agentic workflows that benefit from lower-latency speculative decoding on data centre GPUs and high-end local GPU systems. | 
| Output | Text string(s) | 
| Describe how the model works | NVIDIA-Nemotron-3.5-Lightning-30B-A3B uses interleaved Mamba-2, MoE, and Attention layers for the target model, while DFlash proposes candidate token blocks to improve generation speed during verification | 
| Name the adversely impacted groups this has been tested to deliver comparable outcomes regardless of | Not Applicable | 
| Technical Limitations & Mitigation | The speculative decoding model does not affect the output distribution of the underlying reasoning model. It is only used for lossless inference acceleration. | 
| Verified to have met prescribed quality standards? | Yes | 
| Performance Metrics | Accuracy, throughput, latency, and speculative-decoding acceptance rate | 
| Potential Known Risk | The base model was trained on data that contains toxic language and societal biases originally crawled from the internet. Therefore, the model may amplify those biases and return toxic responses especially when prompted with toxic prompts. Therefore, before deploying any applications of this model, developers should perform safety testing and tuning tailored to their specific applications of the model. | 
| Developers should validate task accuracy, latency, and safety in their own deployment environment and add application-specific safeguards before production use. |  | 
| Licensing | **Governing Terms:** Use of this model is governed by the OpenMDW-1.1 model license | 

| Field | Response | 
|---|---|
| Participation considerations from adversely impacted groups protected classes in model design and testing | None | 
| Measures taken to mitigate against unwanted bias | None | 
| Bias Metric | None | 

| Field | Response | 
|---|---|
| Model Application(s) | Chat, instruction following, RAG, reasoning, and agentic AI workflows | 
| Describe life critical application (if present) | Not Applicable | 
| Use Case Restrictions | Use must comply with the OpenMDW-1.1 model license and applicable laws and regulations | 
| Model and Dataset Restrictions | The Principle of least privilege (PoLP) is applied limiting access for dataset generation and model development. Restrictions enforce dataset access during training, and dataset license constraints adhered to. | 

| Field | Response | 
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
- 213,432
