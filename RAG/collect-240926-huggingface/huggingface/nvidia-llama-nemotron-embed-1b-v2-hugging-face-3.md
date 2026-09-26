---
id: collect-240926-huggingface/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face-3
title: "Compute similarity scores"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["embedding", "latency", "license", "llama", "nvidia", "throughput", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-llama-nemotron-embed-1b-v2-hugging-face.md
source_anchor: ""
source_lines: [249, 285]
sha256: ec6168e164fb2b57e6d4a6052c7866ec6cbcca10f1ba049094c98d8dbefcba15
---

# Compute similarity scores

| Field | Response | 
|---|---|
| Intended Application & Domain: | Passage and query embedding for question and answer retrieval | 
| Model Type: | Transformer encoder | 
| Intended User: | Generative AI creators working with conversational AI models - users who want to build a multilingual question and answer application over a large text corpus, leveraging the latest dense retrieval technologies. | 
| Output: | Array of float numbers (Dense Vector Representation for the input text) | 
| Describe how the model works: | Model transforms the tokenized input text into a dense vector representation. | 
| Performance Metrics: | Accuracy, Throughput, and Latency | 
| Potential Known Risks: | This model does not always guarantee to retrieve the correct passage(s) for a given query. | 
| Licensing & Terms of Use: | Use of this model is governed by the NVIDIA Open Model License Agreement. Additional Information: Llama 3.2 Community Model License Agreement. | 
| Technical Limitations | The model’s max sequence length is 8192. Therefore, the longer text inputs should be truncated. | 
| Name the adversely impacted groups this has been tested to deliver comparable outcomes regardless of: | N/A | 
| Verified to have met prescribed NVIDIA quality standards: | Yes | 

| Field | Response | 
|---|---|
| Generatable or reverse engineerable personally-identifiable information (PII)? | None | 
| Was consent obtained for any personal data used? | Not Applicable | 
| PII used to create this model? | None | 
| How often is the dataset reviewed? | Before Every Release | 
| Is a mechanism in place to honor data subject right of access or deletion of personal data? | No | 
| If personal data was collected for the development of the model, was it collected directly by NVIDIA? | Not Applicable | 
| If personal data was collected for the development of the model by NVIDIA, do you maintain or have access to disclosures made to data subjects? | Not Applicable | 
| If personal data was collected for the development of this AI model, was it minimized to only what was required? | Not Applicable | 
| Is there provenance for all datasets used in training? | Yes | 
| Does data labeling (annotation, metadata) comply with privacy laws? | Yes | 
| Is data compliant with data subject requests for data correction or removal, if such a request was made? | No, not possible with externally-sourced data. | 

| Field | Response | 
|---|---|
| Model Application(s): | Text Embedding for Retrieval | 
| Describe the physical safety impact (if present). | Not Applicable | 
| Use Case Restrictions: | Use of this model is governed by the NVIDIA Open Model License Agreement. Additional Information: Llama 3.2 Community Model License Agreement. | 
| Model and dataset restrictions: | The Principle of least privilege (PoLP) is applied limiting access for dataset generation and model development. Restrictions enforce dataset access during training, and dataset license constraints adhered to. | 

- Downloads last month
- 417,704
