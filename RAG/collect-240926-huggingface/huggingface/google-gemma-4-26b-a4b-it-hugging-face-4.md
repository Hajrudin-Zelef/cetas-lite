---
id: collect-240926-huggingface/huggingface/google-gemma-4-26b-a4b-it-hugging-face-4
title: "Load model"
domain: huggingface
role: reference
task: reference
actors: ["vLLM"]
dates: []
keywords: ["fine-tuning", "leaderboard", "reasoning", "safeguards", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/google-gemma-4-26b-a4b-it-hugging-face.md
source_anchor: ""
source_lines: [374, 426]
sha256: acdeb8e11d8e1117b156ff6b881da5ef67402fe9530b47c39351770c4ee5f84a
---

# Load model

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

If you find our work helpful, please consider citing it:

```
@misc{gemmateam2026gemma4,
      title={Gemma 4 Technical Report}, 
      author={Gemma Team},
      year={2026},
      eprint={2607.02770},
      archivePrefix={arXiv},
      primaryClass={cs.CL},
      url={https://arxiv.org/abs/2607.02770}, 
}
```
- Downloads last month
- 10,715,414

## Spaces using google/gemma-4-26B-A4B-it 100

## Collection including google/gemma-4-26B-A4B-it

## Paper for google/gemma-4-26B-A4B-it

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  82.3
- llamaindex/ExtractBench leaderboard
- Mean View evaluation resultssourcePipeline name: gemma4_26b_vllm_extract_oneshot_structured_output_file; self-hosted on vLLM, one-shot JSON-Schema structured output70.27<sup>*</sup>
- Short View evaluation resultssourcePipeline name: gemma4_26b_vllm_extract_oneshot_structured_output_file; self-hosted on vLLM, one-shot JSON-Schema structured output77.7<sup>*</sup>
- Medium View evaluation resultssourcePipeline name: gemma4_26b_vllm_extract_oneshot_structured_output_file; self-hosted on vLLM, one-shot JSON-Schema structured output59.61<sup>*</sup>
- llamaindex/ParseBench leaderboard
