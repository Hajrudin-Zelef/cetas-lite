---
id: collect-261001-ia-llm/ia-llm/towards-safety-cases-for-frontier-ai-training-2
title: "towards-safety-cases-for-frontier-ai-training"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["training", "alignment", "incident"]
source: docs/RAG/collect-261001-ia-llm/towards-safety-cases-for-frontier-ai-training.md
source_anchor: ""
source_lines: [73, 91]
sha256: fa4c5a8e41d8928df223b3c193f23d5ae1769881ac19b2f9553b0bd3813d9dc8
---

# towards-safety-cases-for-frontier-ai-training

Rollback ability: It should be easy to identify all downstream uses of a misaligned model in the training process (for example: data generation or grading) to enable undoing the effect of misaligned outputs if necessary.

Residual risk completeness: Safety cases should enumerate a comprehensive list of residual risks which are not covered by currently implemented mitigations, as best possible, to enable informed risk-acceptance decisions.

These represent our current recommendations and are in the process of being implemented at OpenAI. We expect our practices to continue to evolve over the coming weeks.

3. Investigations of misalignment incidents

We also have been developing some best practices for investigating severe AI misalignment incidents. Labs should seek to learn as much as possible from individual incidents (similar to investigation practices(opens in a new window) in other high-stakes industries) to be able to prevent such instances from happening in the future. Some examples of such actions could include:

Internal transparency:As an investigation may take significant time to complete, incident investigations should have periodic updates presented internally (e.g., daily updates for ongoing investigations). Employees should have defined pathways to get more access, including raw transcripts and sampling from misaligned models, if it is safe and relevant for their work.

Misalignment root-cause: Researchers should root-cause training dynamics (e.g., via targeted ablations or resampling experiments) to understand how misaligned behaviors were introduced to better understand the science of misalignment and better prevent it in the future.

Postmortem: An operational and cultural postmortem should be conducted to understand all contributing causes to the incident, such as why issues were introduced and left undetected or unescalated prior to the incident.

Detection: We should develop alignment testing methods capable of discovering the propensity to cause the incident, without directly hillclimbing on information derived from the incident (e.g., transcripts or incident summaries). Incident-derived evals should be created as “regression tests,” to ensure future models do not exhibit propensity for misalignment on very similar incidents.

Public disclosures: Investigation results, postmortems, and operational changes should be shared with the public following the conclusion of the investigation. Affected third parties should be notified as soon as possible.
