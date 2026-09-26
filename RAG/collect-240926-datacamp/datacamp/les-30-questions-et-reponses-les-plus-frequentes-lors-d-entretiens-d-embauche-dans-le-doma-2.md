---
id: collect-240926-datacamp/datacamp/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-doma-2
title: "les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agentic", "agents", "attention", "context window", "llama", "reasoning"]
source: docs/RAG/clean_en/datacamp/les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i.md
source_anchor: ""
source_lines: [94, 161]
sha256: 4b4d2eb487dbe1ea8109fa327a56b58cf13fcd509228b492361072564cab7700
---

# les-30-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-dans-le-domaine-de-l-i

I think this is a fairly rare question, but it is nevertheless relevant to ask yourself, perhaps even in a general way and not only in the context of an interview. You can consider ideas related to the position you are applying for or more general ideas, such as AI applications that make decisions affecting human life. There is certainly no right answer to this question, which is generally used to assess your interest in and thinking about the field.

### What security risks should be considered when deploying autonomous AI agents?

There are several security concerns to consider when deploying autonomous AI agents. One risk lies in the fact that the model may have access to sensitive internal tools or databases. If the model is not properly isolated or authorized, a malicious user could use prompt injection. prompt injection or adversarial inputs to extract private data or trigger unwanted actions.

Another risk concerns the manipulation of the model's behavior through carefully crafted prompts or external inputs. An attacker could induce the model to ignore security constraints, expand its privileges, or adopt behavior that deviates from its intended function.

There is also the possibility of denial-of-service attacks, in which the model is overwhelmed with requests or misled so that it disrupts its own operations. If an agent controls critical infrastructure or automated workflows, this could lead to larger disruptions.

In order to mitigate these risks, it is essential to apply traditional software security principles: least privilege, rigorous input validation, monitoring, rate limiting, and continuous evaluation of the agent's behavior.

## Introduction to Artificial Intelligence Agents

### Which human jobs do you think will soon be replaced by agentic AI applications, and why?

Interviewers might ask this question in order to understand your knowledge of the current capabilities of agentic AI. They will not only look for a list, but also a detailed explanation of your reasoning.

For example, I personally don't think doctors will be replaced anytime soon, especially those whose decisions have a direct impact on human life, and this comes back to the question of ethics. There is a lot to explore here, and you can even discuss to determine whether you consider it positive or negative that some jobs are replaced by AI.

### Could you describe some of the challenges you faced when developing an AI application?

Although this is a question that concerns you personally, I placed it in the intermediate section because it is quite common and recruiters tend to give it a lot of importance. It is essential that you have prepared a concrete example; don't try to find one on the spot.

If you haven't yet encountered major challenges, at least try to discuss a theoretical situation and how you would handle it.

## Advanced Agentic AI Interview Questions

Finally, let's tackle some more advanced and technical questions. I will try to be as general as possible, although, generally, the questions asked in a real interview are more specific. For example, instead of asking general questions about indexing, you might be asked about the different indexing methods supported by Langchain or Llama-Index.

### What is the difference between the system and the user prompt?

System prompts and user prompts are both inputs provided to a language model, but they serve different roles and generally have different levels of influence.

The system instruction "" is a hidden instruction that defines the model's general behavior or personality. It is not directly visible to the user during a conversation, but it plays a fundamental role. For example, the system could ask the model to act as an efficient assistant, a mathematician, or a travel planner. It defines the tone, style, and constraints of the interaction.

The user prompt, on the other hand, is the input that the user directly makes, such as a question or a request. This is what the model responds to in real time.

In many configurations, the system prompt is more important, helping to maintain consistent behavior from session to session, while the user prompt determines the specific content of each response.

### How do you program an agentic AI system to prioritize certain competing goals or tasks?

Agentic AI systems are generally programmed by defining clear goals, assigning the appropriate tools, and structuring the logic that determines how the agent prioritizes tasks when goals are competing. This often involves using a combination of prompts, function calls, and orchestration logic, sometimes across multiple models or subsystems.

One approach is to define a hierarchy of goals and assign weights or rules that guide the agent in choosing which task to accomplish in case of conflict. Some systems also use planning components or intermediate reasoning steps (such as reflection loops or scratchpads) to evaluate trade-offs before acting.

If you are new to this, I recommend starting with Anthropic's article on agent design patterns. It provides concrete examples and common architectures used in real systems. If you have a background in software engineering, many of these concepts will be familiar to you, especially those related to modular design, state management, and asynchronous task execution.

### How comfortable are you with prompting and prompt engineering? What approaches have you heard of or used?

Prompt engineering is a major component of an agentic AI system, but it is also a topic that tends to generate stereotypes. It is therefore important to avoid vague statements about its importance and instead focus on the technical details of its application.

Here is what I would consider an appropriate response:

I am very comfortable with prompting and prompt engineering, and I have used several techniques both in projects and in my daily tasks. For example, I regularly use the "few-shot prompting" method to guide models toward a specific format or tone by providing examples. I also use chain of thought when I need the model to reason step by step, which is particularly useful for tasks such as coding, logic puzzles, or planning.

In more structured applications, I have experimented with prompt tuning. prompt tuning and prompt compression, especially when working with APIs that charge per token or that require strict control of outputs. These techniques involve reducing prompts to their most essential elements while preserving their intent and performance.

Since the field is evolving rapidly, I make a habit of reading recent articles, GitHub repositories, and documentation updates, in order to stay up to date with techniques such as function calling. function calling, retrieval-augmented prompting, and modular chaining modular chaining of prompts.

### What is a context window? Why is its size limited?

A context window refers to the maximum amount of information, measured in tokens, that a language model can process simultaneously. This includes the current prompt, the history of previous conversations, and system-level instructions. Once the context window limit is reached, older tokens may be truncated or ignored.

The reason why the context window is limited is due to computational and architectural constraints. In transformer-based models, attention mechanisms require computing the relationships between all tokens in the context, which increases quadratically with the number of tokens. This makes processing very long contexts expensive and slow, especially on current hardware. Earlier models, such as RNNs, did not have a strict context limit in the same way, but they struggled to effectively retain long-term dependencies.

### What is Retrieval-Augmented Generation (RAG)?

