---
id: collect-240926-datacamp/datacamp/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-doma-3
title: "les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i"
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "attention", "benchmark", "compute", "context window", "distillation", "fine-tuning", "inference", "latency", "memory", "reasoning"]
source: docs/RAG/clean_en/datacamp/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i.md
source_anchor: ""
source_lines: [162, 221]
sha256: b8501c87831163e1f9bcbd223f8d2a18c8b2c7418e2f431847e62dfb0a0fd333
---

# les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i

Retrieval-Augmented Generation (RAG) is a technique that enhances language models by enabling them to retrieve relevant information from external sources before generating a response. Instead of relying solely on what the model learned during training, RAG systems can access up-to-date or domain-specific data at inference time.

A typical RAG setup includes two main components: a retriever, which searches a database or document collection for relevant context based on the input query, and a generator, which uses the retrieved information to produce a more accurate and better-informed response. This approach is particularly useful for tasks that require factual accuracy, long-term memory, or domain-specific knowledge.

### What other LLM architectures do you know of besides the transformer?

Although the transformer is the dominant architecture in AI today, there are several other types of models that are important to know. For example, xLSTM builds on the LSTM architecture with improvements that optimize performance on long sequences while maintaining efficiency.

Mamba is another promising architecture: it uses selective state space models to process long contexts more efficiently than transformers, especially for tasks that do not require attention to every token.

Google's Titans architecture also deserves examination. It is designed to address some of the main limitations of transformers, such as the lack of persistent memory and high computational costs.

These alternative architectures aim to make models more efficient, more scalable, and capable of processing longer or more complex inputs without requiring significant hardware resources.

### What is meant by tool use and function calling in LLMs?

Tool and function calling allows large language models to interact with external systems, such as APIs, databases, or custom functions. Instead of relying solely on pre-acquired knowledge, the model is able to recognize when a task requires up-to-date or specialized information and respond by calling an appropriate tool.

For example, if you ask a model with access to a weather API, "What's the weather like in London?", it can decide to call that API in the background and return real-time data instead of generating a generic or outdated response. This approach makes models more useful and more reliable, especially for tasks involving real-time data, calculations, or actions beyond the model's internal capabilities.

### What is chain of thought (CoT) and why is it important in agentic AI applications?

Chain of thought (CoT) is a prompting technique that helps language models break down complex problems into step-by-step reasoning before producing a final answer. is a prompting technique that helps language models break down complex problems into step-by-step reasoning before producing a final answer. This allows the model to generate intermediate reasoning steps, which improves accuracy and transparency, especially for tasks involving logic, mathematics, or multi-step decision-making.

CoT is widely used in agentic AI systems. For example, when a model acts as a judge in an evaluation, you can ask it to explain its answer step by step in order to better understand its decision-making process. CoT is also a fundamental technique in reasoning-focused models such as OpenAI o1, where the model first generates reasoning tokens before using them to produce the final output. This structured thinking process makes agent behavior more understandable and more reliable.

### What is tracing? What are spans?

Tracing is the process of recording and visualizing the sequence of events that occur during a single execution or a single call of an application. In the context of LLM applications, a trace captures the complete history of interactions, such as multiple model calls, tool use, or decision points, within an execution flow.

A span corresponds to a single event or operation within that trace. For example, a model call, a function invocation, or a retrieval step would each be recorded as individual spans. Together, spans help you understand the structure and behavior of your application.

Tracing and spans are essential for debugging and optimizing agentic systems. They make it easier to detect failures, latency bottlenecks, or undesirable behaviors. Tools such as Arize Phoenix and others provide visual interfaces to examine traces and spans in detail.

### What are evals? How do you evaluate the performance and robustness of an agentic AI system?

Evals are essentially the unit tests of agentic AI engineering. They allow developers to assess system performance across different scenarios and edge cases. There are several types of evals commonly used today. One approach is to use a manually curated benchmark dataset to compare model outputs against known correct answers.

Another approach is to use an LLM as a judge to evaluate the quality, accuracy, or reasoning behind model responses. Some evals measure overall task success, while others focus on individual elements such as tool use, planning, or coherence. Running them regularly helps identify regressions, measure improvements, and ensure system reliability as it evolves. To go deeper into the subject, I recommend checking out this LLM evaluation guide.

### Could you tell us about the transformer architecture and its importance for agentic AI?

The transformer architecture was introduced in the influential 2017 paper "Attention Is All You Need." If you haven't read it yet, it's worth doing so, as it laid the foundation for almost all modern large language models.

Since its release, many variants and improvements have been developed, but most models used in agentic AI systems are still based on some form of the transformer.

One of the main advantages of the transformer lies in its attention mechanism, which allows the model to compute the relevance of each token in the input sequence relative to all other tokens, provided everything fits within the context window. This enables excellent performance on tasks that require understanding long-range dependencies or reasoning from multiple inputs.

With regard specifically to agentic AI, the flexibility and parallelism of the transformer make it particularly well suited to handling complex tasks such as tool use, planning, and multi-turn dialogue, which are the fundamental behaviors of most current agentic systems.

### What is LLM observability and why is it important?

LLM observability refers to the ability to monitor, analyze, and understand the behavior of large-scale language model systems in real time. It is an umbrella term that encompasses tools such as traces, spans, and evals, which allow developers to better understand the internal workings of the system.

Because large language models (LLMs) are often regarded as "black boxes," observability is essential for debugging, improving performance, and ensuring reliability. It allows you to track how models interact with each other and with external tools, identify failure points, and quickly detect unexpected behaviors. In agentic AI systems, where multiple steps and decisions are chained together, observability is especially important for maintaining trust and control.

### Could you explain what fine-tuning and model distillation are?

Fine-tuning a model is the process of taking a pre-trained model and refining it with a new dataset, usually with the goal of specializing it for a specific domain or task. This allows the model to adapt its behavior and responses based on more targeted or updated knowledge.

