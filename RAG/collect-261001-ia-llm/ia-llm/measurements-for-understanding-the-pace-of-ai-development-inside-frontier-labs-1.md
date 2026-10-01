---
id: collect-261001-ia-llm/ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs-1
title: "measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-08"]
keywords: ["agent", "agents", "claude", "compute", "guardrails", "latency", "research", "rsp", "safeguards"]
source: docs/RAG/collect-261001-ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs.md
source_anchor: ""
source_lines: [1, 44]
sha256: e3de4fa1e8455aeba707502520b727d9210c61c156bf48f386ebd9bd29b4a989
---

# measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs

AI systems are becoming exponentially more powerful and have begun to automate more of the process of building themselves. As the world considers slowing the pace of frontier AI development, the public needs more information.

In this post, we lay out measurement tools that can illuminate three critical aspects of AI development:

1. The extent to which AI is building the next version of itself, as opposed to being built by humans
2. Our ability to oversee and intervene in actions that AI agents take on Anthropic’s systems
3. The resources that power the development of more capable models

We also provide a snapshot of these metrics from inside Anthropic. It’s important to note that we would expect these numbers to shift if there were coordination on pacing the frontier, as called for by Anthropic CEO Dario Amodei. We plan to embed independent third-party evaluators from multiple organizations at Anthropic, and give them access to internal processes, systems, and data comparable to what internal risk assessment teams have. These third parties will verify safety practices, report incidents, and monitor key metrics such as the ones in this piece.

We are reporting these measurements because they give the public, third parties, and governments better visibility into the pace of AI development inside frontier labs. For each measurement, we describe what we measured, what the measurement showed, and what it would take to publish these measurements regularly in a form others can verify. We share methodological details in the Appendix.

## Reasons to track these measurements

The measurements in this piece are focused on *how models are built.* By better understanding the production process of models, we have a better chance of correlating model inputs, like compute, with model outputs, like capabilities. They complement capability evaluations, which measure *what models can do*. We publish those separately through our Responsible Scaling Policy (RSP) risk reports, which include evidence on how much our models are accelerating AI R&D. In our policy proposal on advanced AI, the Advanced AI Framework (AAIF), we propose rules of the road for how any lab releases safe models, including transparency obligations that governments could require, such as risk reports. Together, these proposed measurements and policies are a starting point for monitoring the pace of AI development from outside the labs.

## (1) Measuring AI-led AI R&D

**Why measure AI-led R&D?** Frontier AI labs increasingly use AI to build future AI models. This process allows labs in democratic countries to develop more capable models more quickly and conduct more safety and testing on models before they are released to secure AI’s benefits while staying on the frontier. However, models accelerating their own development could make it more challenging for humans to understand or control these systems. It is therefore important to share these metrics to understand how close the world is to reaching recursive self improvement (a model fully autonomously building its successor).

**What we measured.** We built a prototype index of how much of Anthropic’s AI research and development (R&D) is performed by Claude, called the Anthropic R&D Automation Index. It’s built by cataloguing every kind of AI R&D work done at the company, rating how automated each task currently is, and aggregating those ratings.

**What we found.** To measure the extent to which AI is doing AI R&D at Anthropic, we use an automation rating scale developed by Epoch AI that measures “Automation Level,” or AL. It runs from AL0 (no AI involvement) to AL5 (AI operates fully autonomously, with no human in the loop). In AL3, AI “collaborates”: it can do large chunks of work under close human direction. In AL4, AI “leads”: it can complete most of the task end-to-end from a high-level prompt, while the human supervises.<sup>1</sup>

As of August 2026,

- Claude is not operating fully autonomously for any measured subset of AI R&D work.
- Claude “leads” 26% of Anthropic’s AI R&D work.
- The share of work at or above “AI collaborates” is above 90%.

**What any AI developer could report today.** Any frontier developer could publish these measures regularly, using a public methodology. This would enable the numbers to be compared over time, and potentially across labs.

Two obstacles stand in the way of cross-lab comparison on this type of reporting. First is the lack of a common methodology. Second, we’re using our own models to evaluate our systems, which could mean that the “judge” model could make the same kinds of errors as the model it is checking. To address these challenges, a developer’s measurements could be verified by a third party, or by other developers’ models, with guardrails in place to limit the exchange of competitively sensitive data. These measures could also become the trigger for stronger requirements, like a fixed testing window before a new model is used for further AI R&D.

## (2) Measuring oversight of AI agents

**Why measure oversight of agents?** Like other frontier developers, Anthropic employees increasingly delegate tasks to agents that work semi-autonomously for long stretches, and that delegate work to one another. As work becomes increasingly automated, from “AI collaborates” toward “AI leads,” agents could make more consequential decisions, such as which research direction to pursue next.

These decisions need safeguards to block harmful behavior from an individual agent, as well as an oversight regime to detect problems that might only arise across many agents. In our monitoring data to date, individual agents rarely misbehave. But when there are millions or billions of agents operating in the economy, even rare events can happen regularly.

**What we measured.** We’ve built a system that lets us oversee and intervene in actions AI agents take on Anthropic’s systems. Here, we consider three different metrics: *coverage,* which describes the share of an agent’s actions that pass through a monitor before or after execution; *review* *latency*, which is the time between an action and its review, first by an automated monitor and then by a human; and *escalation rate,* which is the share of agent activities that are either blocked/redirected (in the case of online monitors) or flagged for further review (in the case of offline monitors).

**What we found.** As of August 2026, there were approximately 30,000 agents doing research and engineering work at Anthropic at any one time in our most-used internal platform. These measurements cover this platform only. The actions of these agents are constrained by two kinds of monitors, summarized below:

