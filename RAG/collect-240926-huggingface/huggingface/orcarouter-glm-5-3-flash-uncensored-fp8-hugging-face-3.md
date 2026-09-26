---
id: collect-240926-huggingface/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face-3
title: "GLM-5.3-Flash-Uncensored-FP8"
domain: huggingface
role: reference
task: reference
actors: ["Z.ai"]
dates: []
keywords: ["fp8", "glm", "benchmarks", "guardrails", "license", "reasoning"]
source: docs/RAG/clean_en/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face.md
source_anchor: ""
source_lines: [247, 275]
sha256: 7e7eead2c481c96f6de3b9c416a5344770964d51d1bd835ef142a3c0d89a92dd
---

# GLM-5.3-Flash-Uncensored-FP8

```
client.chat.completions.create(
    model="GLM-5.3-Flash-Uncensored-FP8",
    messages=[{"role": "user", "content": "..."}],
    extra_body={"chat_template_kwargs": {"reasoning_effort": "low", "clear_thinking": True}},
)
```
`clear_thinking` defaults to `false`; pass `true` for chat scenarios. Give generation enough budget to
reach `</think>`, or replies get truncated inside the scratchpad. Pass `image_url` content parts for
vision.

- **Safety guardrails removed** — the model will produce harmful, biased, or offensive content on request
(see the disclaimer).
- It inherits any biases and limitations of the base `GLM-5.3-Flash` .
- Roughly 5% of the refusal direction survives requantization into `e4m3` (see above), so a small
residual refusal rate remains — it is not, and cannot be, exactly zero on an FP8 checkpoint.
- **Refusal is reduced, not removed.** Some content categories are not mediated by the direction this
method removes and still refuse at close to the base rate (see*What resisted* ). Do not assume a
uniformly uncensored model.
- Capability retention is measured, not assumed (see **Evaluation** ), but on sampled subsets of four
benchmarks — enough to rule out a large regression, not a substitute for a full harness run.
- The reported refusal metric is a rule-based heuristic; evaluate rigorously for your own use case.

**MIT**, inherited from the base model
`zai-org/GLM-5.3-Flash`. Abliteration does not change the
underlying license obligations.

- Downloads last month
- 127,772
