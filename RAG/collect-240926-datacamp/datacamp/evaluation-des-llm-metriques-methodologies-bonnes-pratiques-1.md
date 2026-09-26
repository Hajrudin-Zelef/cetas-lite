---
id: collect-240926-datacamp/datacamp/evaluation-des-llm-metriques-methodologies-bonnes-pratiques-1
title: "evaluation-des-llm-metriques-methodologies-bonnes-pratiques"
domain: datacamp
role: reference
task: reference
actors: ["Cohere", "Perplexity"]
dates: []
keywords: ["alignment", "benchmarks", "distribution", "perplexity"]
source: docs/RAG/clean_en/datacamp/evaluation-des-llm-metriques-methodologies-bonnes-pratiques.md
source_anchor: ""
source_lines: [1, 112]
sha256: dc62bcc3992ed90c82279c96a4418685060848fdda510a499610c1f1b658c3d8
---

# evaluation-des-llm-metriques-methodologies-bonnes-pratiques

<!-- source: https://www.datacamp.com/fr/blog/llm-evaluation -->

Course

Large language models (LLMs) are rapidly establishing themselves in many applications, from chatbots to content creation.

However, evaluating these powerful models remains complex. How can their performance and reliability be precisely measured, given the diversity of their capabilities and implementations?

This guide offers a comprehensive overview of LLM evaluation, covering essential metrics, methodologies, and best practices to help you choose the models best suited to your needs.

## Improving AI for Beginners

## Key Metrics for Evaluating LLMs

Evaluating LLMs requires a holistic approach, mobilizing several measures to assess different aspects of their performance. Here we review the major criteria for evaluating LLMs, including accuracy and performance, bias and fairness, as well as other important metrics.

### Accuracy and Performance Metrics

Measuring performance correctly is a key step in understanding the capabilities of an LLM. This section details the main metrics used to evaluate accuracy and performance.

#### Perplexity

Perplexity is a fundamental metric for evaluating and measuring an LLM's ability to predict the next word in a sequence. Here is how it can be calculated:

1. Probability: first, the model calculates the probability of each word likely to appear next in the sentence.
2. Inverse probability: we take the inverse of this probability. For example, if a word has a high probability (the model judges it likely), its inverse probability will be lower.
3. Normalization: we then calculate the average of this inverse probability over all the words in the test set (the text on which the model is evaluated).

#### Illustration of an LLM predicting the probability of the next word according to context. Source

Lower perplexity scores indicate that the model predicts the next word more accurately, which reflects better performance. In short, perplexity quantifies the ability of a probabilistic model to predict a sample.

For LLMs, low perplexity means that the model is more confident in its predictions, and therefore generates more coherent texts that are better adapted to the context.

#### Accuracy

Accuracy is a metric commonly used for classification tasks, representing the share of correct predictions made by the model. Although intuitive, it can be misleading for open-ended generation tasks.

For example, when it comes to generating creative or nuanced text, the notion of "correctness" is less clear-cut than for tasks such as sentiment analysis or topic classification. Useful for specific cases, accuracy must therefore be complemented by other metrics to evaluate LLMs.

#### BLEU/ROUGE Scores

BLEU (Bilingual Evaluation Understudy) and ROUGE (Recall-Oriented Understudy for Gisting Evaluation) make it possible to evaluate the quality of a generated text by comparing it to reference texts.

BLEU emphasizes precision: if a machine translation uses the same words as a human translation, the BLEU score is high. For example, if the human reference is "The cat is on the mat" and the machine output is "The cat sits on the mat," the BLEU score will be high due to the significant word overlap.

ROUGE favors recall: it checks whether the generated text adequately covers the essential ideas of the reference text. If a human summary states "The study found that people who exercise regularly tend to have lower blood pressure." and the AI summary is "Exercise linked to lower blood pressure," ROUGE will assign a high score because the main idea is well captured, even with different wording.

These metrics are useful for tasks such as machine translation, automatic summarization, and text generation, providing a quantitative measure of alignment with human references.

### Bias and Fairness Metrics

Ensuring fairness and reducing bias in LLMs is essential for equitable uses. Here are the main metrics for evaluating bias and fairness.

#### Demographic Parity

Demographic parity examines whether the model's performance is consistent across different demographic groups. It evaluates the proportion of positive outcomes according to attributes such as origin, gender, or age.

Achieving demographic parity means that the model's predictions neither favor nor disadvantage any group, ensuring fairness and justice in its applications.

#### Equal Opportunity

Equal opportunity focuses on the distribution of model errors across demographic groups. It notably evaluates false negative rates to verify that the model does not fail disproportionately for certain groups.

This metric is crucial for applications where fairness and equal access are fundamental, such as recruitment algorithms or loan granting.

#### Counterfactual Fairness

Counterfactual fairness evaluates whether the model's predictions would change if certain sensitive attributes were different. It consists of generating counterfactual examples where the sensitive attribute (e.g., gender or origin) is modified, with the other characteristics remaining constant.

If the prediction varies following this modification, this reveals a bias related to the sensitive attribute. Counterfactual fairness is essential for detecting and mitigating biases that do not always emerge with other metrics.

### Other Metrics

Beyond performance and fairness, other criteria contribute to a comprehensive evaluation of LLMs. This section highlights these aspects.

#### Fluency

Fluency measures the naturalness and grammatical correctness of the generated text. A fluent LLM produces outputs that are easy to read and understand, reproducing the rhythm of human language.

It can be evaluated using automated tools or human judgment, focusing on grammar, syntax, and overall readability.

#### Coherence

Coherence analyzes the logical flow and consistency of the generated text. A coherent text maintains a clear structure and a logical progression of ideas, making it easier to read. It is particularly important for long texts, such as essays or articles, where the continuity of the narrative thread is key.

#### Factuality

Factuality evaluates the accuracy of the information provided by the LLM, particularly for information retrieval tasks. It verifies that the generated text is not only plausible but also factually correct.

Essential for uses such as news generation, educational content, or customer support, where accuracy is paramount.

## Evaluation Methodologies

A robust evaluation of LLMs combines quantitative and qualitative approaches. This section details several methods, such as reference datasets, human evaluation, and automated evaluations, to rigorously assess performance.

### Reference Datasets

Reference datasets provide standardized tasks for comparing different models. They establish a common baseline and facilitate benchmarking.

#### Existing Benchmarks

Among the most popular reference datasets for various natural language processing (NLP) tasks:

- GLUE (General Language Understanding Evaluation): a set of diverse tasks to evaluate the general linguistic capabilities of LLMs, including sentiment analysis, textual entailment, and question answering.
- SuperGLUE: a more advanced version of GLUE, with more difficult tasks to test the robustness and fine-grained understanding of LLMs.
- SQuAD (Stanford Question Answering Dataset): a reading comprehension dataset where models are scored on their ability to answer questions from Wikipedia articles.

#### Custom Datasets

While valuable, existing benchmarks often need to be supplemented by custom datasets for sector-specific evaluation. They allow evaluation to be adapted to the requirements and constraints specific to an application or sector.

