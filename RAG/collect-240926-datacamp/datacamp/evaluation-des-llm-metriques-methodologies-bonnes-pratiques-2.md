---
id: collect-240926-datacamp/datacamp/evaluation-des-llm-metriques-methodologies-bonnes-pratiques-2
title: "evaluation-des-llm-metriques-methodologies-bonnes-pratiques"
domain: datacamp
role: reference
task: reference
actors: ["Perplexity"]
dates: []
keywords: ["benchmarks", "exploit", "fine-tuning", "perplexity"]
source: docs/RAG/clean_en/datacamp/evaluation-des-llm-metriques-methodologies-bonnes-pratiques.md
source_anchor: ""
source_lines: [113, 166]
sha256: 0b8969c566fd7fdf62e5c2764cb5be2d4ea05b98817640ca2009ec902eea9192
---

# evaluation-des-llm-metriques-methodologies-bonnes-pratiques

For example, a healthcare organization can build a corpus of medical records and clinical notes to evaluate an LLM's mastery of medical terminology and context. These datasets ensure evaluation aligned with real use cases and more actionable insights.

### Human Evaluation

Human evaluation methods are essential for assessing nuances that automated metrics may miss. They rely on direct feedback from evaluators, offering qualitative insights into performance.

#### Direct Evaluation

Human evaluation remains the gold standard for judging the quality of an LLM's outputs. Direct approaches collect opinions via surveys and rating scales.

They capture fine dimensions such as fluency, coherence, and relevance, often overlooked by automated metrics. Evaluators can also point out concrete strengths and weaknesses, to target areas for improvement.

#### Comparative Judgment

Comparative judgment, such as pairwise comparison, involves directly pitting the outputs of different models against each other. This method is sometimes more reliable than absolute scores, by reducing individual subjectivity.

Evaluators choose the best text among pairs, which provides a relative ranking of models. Particularly useful for fine-tuning and selecting higher-performing variants.

### Automated Evaluation

Automated methods offer a fast and objective way to evaluate LLM performance. They mobilize various metrics to quantify multiple dimensions of outputs, ensuring comprehensive evaluation.

#### Metric-Based

Automated metrics provide fast and objective evaluation. Indicators such as perplexity and BLEU are widely used to assess different facets of text generation.

As seen previously, perplexity measures the model's prediction capability, with lower scores indicating better performance. BLEU, meanwhile, evaluates the quality of generated text by comparing it to references, focusing on n-gram precision.

### Adversarial Evaluation

Adversarial evaluation involves subjecting LLMs to adversarial attacks to test their robustness. These attacks exploit model flaws and biases, revealing vulnerabilities that escape classical evaluations.

An attack can, for example, introduce slightly modified or misleading inputs to analyze the model's reaction. This approach is useful when reliability and security are paramount, as it helps identify and mitigate potential risks.

## Best Practices for Evaluating LLMs

To effectively evaluate LLM capabilities, a structured approach should be adopted. Best practices ensure a thorough, transparent evaluation tailored to your needs. Here are the ones to prioritize.

| Good practice | Description | Use case | Relevant metric(s) | 
| Define clear objectives | Identify the tasks and expected outcomes of the LLM before starting the evaluation. | Improve the machine translation performance of an LLM | BLEU/ROUGE scores | 
| Take your audience into account | Adapt the evaluation to the intended users of the LLM, according to their expectations and needs. | Text generation LLM | Perplexity, fluency, coherence | 
| Transparency and reproducibility | Document the evaluation process so that it can be reproduced and audited by third parties. | Publication of the evaluation dataset and the code used to measure the LLM's capabilities | Any relevant metric, depending on the task and objectives | 

## Conclusion

This guide has presented a comprehensive overview of the essential metrics and methodologies for evaluating LLMs, from perplexity and accuracy to measures of bias and fairness.

By combining quantitative and qualitative approaches, and by following best practices, you ensure a reliable and thorough evaluation of these models.

With these benchmarks, you will be better equipped to select and deploy the most suitable LLMs, in order to ensure performance and reliability in your applications.

## Obtain a high-level certification in AI

Recent master's graduate in science, specializing in artificial intelligence
