---
id: collect-261001-ia-llm/ia-llm/kimi-researcher
title: "kimi-researcher"
domain: ia-llm
role: reference
task: reference
actors: ["Moonshot", "OpenAI"]
dates: []
keywords: ["kimi", "research", "agent", "agentic", "benchmark", "benchmarks", "embedding", "reasoning", "tool use", "training"]
source: docs/RAG/collect-261001-ia-llm/kimi-researcher.md
source_anchor: ""
source_lines: [1, 40]
sha256: 139052d62c873be45497db551c742bb028458366c76f81b7e794d1630ae24def
---

# kimi-researcher

Meet **Kimi-Researcher**, an autonomous agent that excels at multi-turn search and reasoning. It performs an average of 23 reasoning steps and explores over 200 URLs per task. Built on an internal version of the Kimi k-series model and trained entirely through end-to-end agentic reinforcement learning (RL), it achieved a Pass@1 score of **26.9%**—a state-of-the-art result—on Humanity's Last Exam, and Pass@4 accuracy of **40.17%**. Starting from an initial HLE score of 8.6%, Kimi-Researcher reached 26.9% almost entirely through end-to-end RL training, providing compelling evidence that end-to-end agentic RL can significantly advance agent intelligence.
        

Kimi-Researcher has also achieved strong performance across several complex and challenging real-world benchmarks. On xbench, a new, dynamic, professionally-aligned suite designed to bridge AI capabilities with real-world productivity, Kimi-Researcher achieved 69% pass@1 (averaged on 4 runs) on xbench-DeepSearch, outperforming models such as o3 with search tools. On benchmark tests for multi-turn search reasoning (FRAMES, Seal-0) and factual information (SimpleQA), Kimi-Researcher also achieved strong performance.

            Kimi-Researcher is an autonomous agentic and thinking model designed to solve complex problems through multi-step planning, reasoning, and tool use. It leverages three main tools: a parallel, real-time internal **search tool**; a text-based **browser tool** for interactive web tasks; and a **coding tool** for automated code execution.
        

            Formally, given the state observation 

Traditional agent development has key limitations:

End-to-end agentic reinforcement learning trains a single model to solve problems holistically: given a query, the agent explores a large number of possible strategies, receives rewards for correct solutions, and learns from the full trajectory. Unlike SFT, it naturally handles long, on-policy reasoning and adapts to changing tools and environments; unlike modular approaches, all skills—planning, perception, and tool use—are learned together without hand-crafted rules or workflow templates. Previous work like OpenAI's Deep Research also highlights the strong performance of this approach, but it introduces new challenges:

Kimi-Researcher is trained via end-to-end reinforcement learning. We observe a consistent improvement in agent performance across different domains. Figure 2-a illustrates the overall training accuracy of Kimi-Researcher throughout the reinforcement learning process. Figure 2-b presents model performance on several internal datasets.

To address the scarcity of high-quality agentic datasets, we engineered our training corpus with two complementary objectives.

First, we developed a suite of challenging, tool-centric tasks designed to promote robust tool-use learning. These prompts are deliberately constructed such that solving the task requires invoking specific tools—making naive approaches either infeasible or substantially less efficient. By embedding tool dependencies into task design, the agent learns not only when to invoke a tool, but also how to orchestrate tool use effectively in complex, real-world settings. (See Figure 3 for tool invocation rates using these training data.)

Second, we curated and synthesized reasoning-intensive tasks to reinforce the agent's core cognitive abilities and its capacity to integrate reasoning with tool usage. This component is further subdivided into:

                To build this diverse prompt set at scale, we developed a **fully automated pipeline** capable of generating and validating many question-answer pairs with **minimal manual intervention**, ensuring both diversity and correctness at unprecedented scale. **Ensuring accurate ground truth (GT) is critical for synthetic tasks**, so we introduced a robust GT extraction method to guarantee that each question is paired with a reliable answer whenever possible. Additionally, a rigorous filtering funnel removes ambiguous, trivial, or incorrect pairs — with **Pass@N checks ensuring only non-trivial questions are retained**. Figure 4 shows the effectiveness of our synthetic tasks based on two experimental results.
            

The model is primarily trained using the REINFORCE algorithm. We have observed that the following factors contribute to more stable training:

Kimi-Researcher uses outcome rewards for training, aiming to provide a constant preference in a dynamic training environment.

                To promote efficiency, a gamma-decay factor is applied to correct trajectories. Concretely, the reward of step 

To address the efficiency and stability challenges of large-scale Agent RL, we have developed a suite of infrastructure with the following key features:

During end-to-end reinforcement learning, we observed several notable emergent abilities in Kimi-Researcher. Here are two highlights:

Kimi-Researcher is beginning its gradual rollout to users today. It empowers you to conduct deep, comprehensive research on any topic directly within Kimi. Join the waitlist here.

It represents the early stage of our broader vision: evolving from a focused search and reasoning agent into a general-purpose agent capable of solving a wide range of complex tasks with an ever-expanding toolkit. To realize this vision, we are expanding the agent's capabilities across both tools and task domains, while also advancing the underlying reinforcement learning infrastructure and algorithms to ensure greater training stability and efficiency.

To facilitate more research efforts in the field, we are planning on open-sourcing the base pretrained model as well as the reinforcement-learned model underlying Kimi-Researcher in the following months.
