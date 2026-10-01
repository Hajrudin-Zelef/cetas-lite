---
id: collect-261001-ia-llm/ia-llm/kimi-k2-thinking-1
title: "Introducing Kimi K2 Thinking"
domain: ia-llm
role: reference
task: reference
actors: ["Moonshot"]
dates: []
keywords: ["kimi", "agent", "agentic", "agents", "benchmark", "benchmarks", "gpu", "inference", "int4", "latency", "memory", "moe"]
source: docs/RAG/collect-261001-ia-llm/kimi-k2-thinking.md
source_anchor: ""
source_lines: [1, 58]
sha256: 9d4154bbe2e188f0f0099bf2400b2b8cc53cc8b94ea66b1af2d307b559adb02e
---

# Introducing Kimi K2 Thinking

Today, we are introducing **Kimi** **K2** **Thinking**, our best open-source thinking model.

Built as a **thinking agent**, it reasons step by step *while* using tools, achieving state-of-the-art performance on Humanity's Last Exam (HLE), BrowseComp, and other benchmarks, with major gains in reasoning, agentic search, coding, writing, and general capabilities.

Kimi K2 Thinking can execute up to **200 – 300 sequential tool calls** without human interference, reasoning coherently across hundreds of steps to solve complex problems.

It marks our latest efforts in **test-time scaling**, by scaling both thinking tokens and tool calling steps.

K2 Thinking is now live on kimi.ai under the chat mode [1], with its full agentic mode available soon. It is also accessible through the Kimi K2 Thinking API.

### Evaluations

Kimi K2 Thinking sets new records across benchmarks that assess reasoning, coding, and agent capabilities. K2 Thinking achieves 44.9% on HLE with tools, 60.2% on BrowseComp, and 71.3% on SWE-Bench Verified, demonstrating strong generalization as a state-of-the-art thinking agent model.

### Agentic Reasoning

K2 Thinking demonstrates outstanding reasoning and problem-solving abilities. On Humanity’s Last Exam (HLE)—a rigorously crafted, closed‑ended benchmark—spanning thousands of expert‑level questions across more than 100 subjects, K2 Thinking achieved **a** **state-of-the-art** **score of** **44.9%**,  with search, python, and web-browsing tools, establishing new records in multi‑domain expert‑level reasoning performance.

By reasoning while actively using a diverse set of tools, K2 Thinking is capable of **planning, reasoning, executing, and adapting across hundreds of steps** to tackle some of the most challenging academic and analytical problems. In one instance, it successfully solved a **PhD-level mathematics problem** through **23 interleaved reasoning and tool calls**, exemplifying its capacity for deep, structured reasoning and long-horizon problem solving:

### Agentic Coding

K2 Thinking exhibits substantial gains in coding and software development tasks. It achieves scores of 61.1% on SWE-Multilingual, 71.3% on SWE-Bench Verified, and 47.1% on Terminal-Bench, showcasing strong generalization across programming languages and agent scaffolds.

The model delivers notable improvements on HTML, React, and component-intensive front-end tasks—translating ideas into fully functional, responsive products. In agentic coding settings, it reasons while invoking tools, integrating fluidly into software agents to execute complex, multi-step development workflows with precision and adaptability.

Here are some examples that Kimi K2 Thinking has built from a single prompt:

### Agentic Search and Browsing

K2 Thinking demonstrates strong performance in agentic search and browsing scenarios. On BrowseComp—a challenging benchmark designed to evaluate models' ability to **continuously browse, search, and reason over hard-to-find real-world web information**—K2 Thinking achieved a score of **60.2%**, significantly outperforming the human baseline of 29.2%. This result highlights K2 Thinking's superior capability for goal-directed, web-based reasoning and its robustness in dynamic, information-rich environments.

K2 Thinking can execute **200–300 sequential tool calls**, driven by **long-horizon planning** and **adaptive reasoning**. It performs dynamic cycles of *think → search → browser use → think → code*, continually generating and refining hypotheses, verifying evidence, reasoning, and constructing coherent answers. This interleaved reasoning allows it to decompose ambiguous, open-ended problems into clear, actionable subtasks.

Here's one example of what **Kimi ****K2**** Thinking** has researched from just a single prompt:

### General Capabilities

**Creative Writing: **K2 Thinking delivers improvements in completeness and richness. It shows stronger command of style and instruction, handling diverse tones and formats with natural fluency. Its writing becomes more vivid and imaginative—poetic imagery carries deeper associations, while stories and scripts feel more human, emotional, and purposeful. The ideas it expresses often reach greater thematic depth and resonance.

**Practical Writing: **K2 Thinking demonstrates marked gains in reasoning depth, perspective breadth, and instruction adherence. It follows prompts with higher precision, addressing each requirement clearly and systematically—often expanding on every mentioned point to ensure thorough coverage. In academic, research, and long-form analytical writing, it excels at producing rigorous, logically coherent, and substantively rich content, making it particularly effective in scholarly and professional contexts.

**Personal & Emotional: **When addressing personal or emotional questions, K2 Thinking responds with more empathy and balance. Its reflections are thoughtful and specific, offering nuanced perspectives and actionable next steps. It helps users navigate complex decisions with clarity and care—grounded, practical, and genuinely human in tone.

### Inference Efficiency

Low-bit quantization is an effective way to reduce inference latency and GPU memory usage on large-scale inference servers. However, thinking models use excessive decoding lengths, and thus quantization often results in substantial performance drops.

To overcome this challenge, we adopt Quantization-Aware Training (QAT) during the post-training phase, applying INT4 weight-only quantization to the MoE components. It allows K2 Thinking to support native INT4 inference with a roughly 2x generation speed improvement while achieving state-of-the-art performance. All benchmark results are reported under INT4 precision.

### Full Evaluations [2]

The table below shows that Kimi K2 Thinking matches or surpasses the latest open-source and frontier models across a wide range of tasks, excelling on benchmarks for reasoning, agentic search, and coding.

**Footnotes**

