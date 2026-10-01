---
id: collect-261001-ia-llm/ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs-2
title: "measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-07", "2026-08"]
keywords: ["agent", "agents", "alignment", "claude", "compute", "incident", "latency", "pretraining", "research", "safeguards", "sandbox", "training"]
source: docs/RAG/collect-261001-ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs.md
source_anchor: ""
source_lines: [45, 86]
sha256: e2cb2945e6da64066b3f62f8c352318dc6e7b1caa557cd42e6186293c1587c82
---

# measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs

|  | **Online monitors** | **Offline monitors** | 
| **Description** | Online monitors check the actions of agents *in real-time* and either block dangerous actions or steer agent misbehavior automatically. They’re built to address harms that would happen quickly and could not be undone, such as an agent copying model weights out of our systems. | Offline monitors read agent activities *after the fact* . They are built to help address dangerous behaviors that develop slowly and can be reversed, like signs of model misalignment. We use flagged activities to improve our mitigations, like classifiers and deterministic controls. If the situation caused a large imminent danger, we would shut down the dangerous parts of the agent’s deployment. | 
| **Coverage** | 100% of these agents’ actions pass through the online monitor before they are executed, usually within seconds. | 100% of these agents’ actions are ingested by the monitor after they are taken. | 
| **Review latency** | Automated review happens before an agent takes action; humans review any blocked actions within one week. | The monitor flags roughly 100,000 transcripts per week that are then reviewed by multiple classifiers for false-positives and initial oversight review. Most of the flags are resolved by this process, and the highest priority flags (approximately ~50 per week) are escalated to human review. | 
| **Blocking / escalation rate** | We analyzed over a billion decisions from our research & engineering agents over August 2026 and found that 0.002% of them (about 1 in 47,000) were blocked by this monitor. | The monitor flags roughly one to two transcripts in every thousand for further review. | 

**What an AI developer could report today.** Any developer running agents on its own research and engineering workloads could publish the same measures: *coverage* (the share of agent activity that is monitored), *review latency* (how soon flagged activities are reviewed), and *escalation rate* (what share of agent activities the monitors block or flag). Together, these measures would allow society to see whether oversight is keeping pace with the growing role of AI in AI R&D. We published all of these measurements in our recent risk report.

Developers should share how often known agent misbehavior is caught by monitors, and should have third parties do this kind of testing independently. In the past, METR has independently red-teamed our offline monitoring platform. As described above, we are now setting up external third party evaluators at Anthropic.

## (3) Measuring compute allocation

**Why measure compute allocation?** Broadly speaking, AI developers use compute for building more powerful models, serving customers, and safety-focused work like auditing a model’s “thoughts”, training model organisms to study misalignment, and evaluating whether a model can be safely deployed. Understanding how AI developers allocate their compute can tell you where a developer is focusing its resources and how that focus changes over time.

Additionally, compute is among the most verifiable inputs to the AI R&D process, meaning that it could be a critical lever in a future pacing effort. A coordinated pacing effort could encourage companies to increase the compute allocated to safety across the industry and devote more resources to alignment, interpretability, safety testing, and evaluation.

**What we measured.** We examined a snapshot of how Anthropic used all of its compute from July 13 to July 20.<sup>2</sup> To do that, we sorted every workload into a small number of categories, then asked how much of the compute going to AI R&D was safety work.

Safety research tends to use less compute than frontier training runs by its nature, so compute is an imperfect proxy for how much a company focuses on safety. This is because safety research consists of individual researchers designing experiments, which is time-consuming even though running the experiments is not particularly compute-intensive. The value of this metric, therefore, is less the absolute numbers and more that it provides a straightforward mechanism to compare like with like, across developers and over time.

**What we found.** Over the examined week, about 6% of compute that went to AI R&D was allocated toward safety, and about 12% of compute that went to AI-driven AI R&D was allocated toward safety.

These are deliberately conservative estimates. For example, if a token was used to advance capabilities as much as it was to advance safety, it was not counted in these metrics. Additionally, these metrics do not account for safeguards classifiers, which are a separate, comparable amount of compute that make our models much safer for the world.

**What an AI developer could report today.** Any frontier developer could publish what share of its AI R&D compute goes to safety work, with the category definitions published alongside and the classification checked by an independent third party.

Safety research is hard to distinguish from capabilities research, and each developer will be tempted to draw the line generously. The burden of proof should sit with the developer to show that work is safety-related. Developers, governments, and the wider research community would benefit from converging on a shared definition ahead of time. A measurement like this could inform future actions, such as a lab’s commitments about the share of compute going to safety research, or limits on the share of compute going towards AI research agents.

## Conclusion

As the world considers pacing the frontier, we should do everything possible to minimize the gap between what frontier labs know and what the public knows. This means better measuring the development of AI, reporting on it publicly, and giving society an opportunity to decide how to use this information. We hope to model that transparency by releasing these measurements, and we’ll continue to do so.

## Appendix

Here are methodological details on all of the measurements we’ve prototyped.

### Measuring AI-led R&D

**How we did it.** The Automation Index requires three things: a complete map of all the AI R&D tasks being done at Anthropic, a way to rate the level of automation, and a way to weight the tasks, so that important areas of work count for more than less important ones. No one person can list every AI R&D task at a frontier AI company by hand, at least not at the granularity we want. Instead we constructed this list of tasks in a bottom-up manner from work records including Slack and various sources of internal documentation.

For each week in July 2026, we randomly sampled 20% of staff from each department that make up the model R&D loop. A Claude research agent reviewed each sampled person’s week using Slack and internal documentation, and listed the tasks they worked on. Repeating this for each week in July 2026 gives us a flat list of ~15,000 granular model R&D tasks. We then used Claude to organize these tasks into a hierarchical tree, starting from all model R&D at the root and branching into areas such as training and product, then pretraining and reinforcement learning, and so on down to increasingly specific kinds of work. The resulting tree has 542 nodes at different depths, of which 378 are leaves like “eval platform defect diagnosis and fixes,” “RL sandbox egress and network policy,” and “serving incident postmortems.” We freeze this tree so that every measurement we make happens against the same basket of work.

