---
id: collect-240926-datacamp/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants-3
title: "Course"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/clean_en/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants.md
source_anchor: ""
source_lines: [190, 198]
sha256: 2fc0e7f3452884d0f52ba3f562220726a07c6085fde1a877cf35a9bf75de2a54
---

# Course

**Not explicitly. Some harnesses inform the model, via the system prompt, of the available tools, but the model has no notion of the harness as a system around it. It only sees the provided context, generates a response, and sometimes produces a tool call. Consequence: when something breaks, the model often cannot explain why, because it is unaware of the harness's existence. Debugging an agent therefore mostly means debugging the harness, not the model.**

### How does the choice of model influence which harness to use?

**More than one might think. State-of-the-art coding models are sometimes post-trained with their own agent harness in the loop, so replacing that harness can degrade performance. Practical heuristic: if your team commits to a model family, the short list of agent harnesses often imposes itself. The hardest case is changing models later: this generally involves rewriting the harness logic, not just changing a configuration value.**

### Is this different from what used to be called "LLM scaffolding"?

**Not really. It is the same idea under a more recent name. "LLM scaffolding," "agent wrapper," and "execution environment" all point in the same direction. The nuance in 2026: "scaffolding" suggests a temporary structure to be dismantled once the model is good enough, whereas "agent harness" suggests something the model keeps around it. This changes how to budget: scaffolding is removed, agent harnesses become an integral part of the system.**
