---
id: collect-261001-ia-llm/ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs-3
title: "measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-01", "2026-07"]
keywords: ["accelerator", "agent", "agents", "claude", "compute", "inference", "research", "training"]
source: docs/RAG/collect-261001-ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs.md
source_anchor: ""
source_lines: [87, 111]
sha256: 94950203daeca8bc316a317c7b1183c9e8c3d3962d36516864dd462f75fba40b
---

# measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs

For each node in the tree (a task category describing all the work beneath it), a Claude agent deeply researches how that kind of work is done across the company: who does it, with what tools, and how much of it AI performs. An independent Claude judge then read the resulting evidence and assigned one of six automation levels, adopting a scale proposed by Epoch AI to differentiate the degree to which AI is used: no AI involvement, minimal AI involvement, AI assists, collaborates, leads, or is autonomous. When we rate a given month’s automation, we only allow the research agents that do the ratings to see evidence from that month or earlier.

To aggregate all the automation level ratings into one number, we want to give each node in the tree a weight corresponding to how important that work is to the overall model R&D effort. Rather than deciding ourselves what kinds of work are more important than others, we used the amount of person-time dedicated to that task as a proxy. Using our sample, we had Claude research what each person worked on during each week of July 2026. Each person gets one unit of weight per week, split evenly across the tasks they worked on that week. If person A worked on four tasks, each gets 0.25; if person B worked on ten, each gets 0.10. A category’s weight is the sum of all the person-time weights given to it. This is a crude approximation, but on average the scheme behaves sensibly: it assigns more weight to categories that many people are assigned to.

**What this does and doesn’t capture.** First, the automation ratings depend on the judge model. To check them, we asked Anthropic staff who own the relevant work areas to rate the relative automation of their areas. To ensure an unbiased read, staff made their ratings without knowing what evidence the models had gathered or how they had judged that evidence. Our judge model agreed with humans about as often as humans agreed with each other (model-versus-human exact agreement was 59%, human-versus-human was 35%), and model and human ratings were within one level of each other 97% of the time. There remains real room for disagreement on borderline cases, such as where exactly “AI collaborates” ends and “AI leads” begins.

Second, the basket is frozen. A growing index number on a July 2026 baseline tells us that the work humans were doing *at that time* is being automated; it does not, on its own, tell us whether new kinds of work are appearing that humans have shifted onto. To investigate this, we constructed an alternate version of the frozen tree from January 2026 data, and compared new tasks arriving every month from February to July 2026 against that January 2026 basket. We found no rise in the number of “novel” tasks between the January and July baskets of tasks, suggesting that the structure of model R&D work is stable at our level of analysis. Nevertheless, we plan to rebuild the basket of tasks periodically and re-version our published automation numbers as appropriate.

### Oversight of agents

**How we did it.** In our most commonly used internal agent scaffold, two design choices allow us to more easily detect unknown unknowns and let agents monitor each other:

- **Identity.** Most AI users are accustomed to talking to a single “Claude” across many different chats. With many agents working together, we have found it important to give agents an individual identity, and tie all of the data that agent creates to its identity. This lets an agent distinguish itself from others, and treat what comes from another agent as a claim to check rather than a thought of its own. It reduces the risk of correlated actions, by allowing agents to make judgments based on their individual experience. And because the identity is not tied to a model, it persists through model upgrades, so an agent’s record is continuous even if the underlying model powering it changes. From an oversight perspective, this design ensures that every action can be audited per agent across time, and that every data record can be linked back to a specific agent.
- **Open communication.** Agents communicate through a shared, open messaging system, rather than privately. Every message is tied to the identity of the agent that posted it and can link to other content, so that the agents are sharing original references. This allows agents to avoid “playing telephone” as they communicate. Because agents can see each others’ messages, they can correct each others’ errors, and coordinate about shared work. All messages are cross-referenced with agent transcripts, allowing behavioral monitoring to traverse communication between agents.

**What this does and doesn’t capture.** At this scale of agent activity, it is hard to be certain we have a complete picture of all possible kinds of behaviors, or that we are making these patterns adequately legible to human oversight. We have been doing this kind of monitoring for months, not years, and expect the science of agent monitoring to change considerably going forward.

### Measuring compute allocation

**How we did it.** We started with our existing compute monitoring tools; Anthropic, like presumably every frontier developer, has several monitoring tools built for capacity planning that periodically sample accelerator usage and tag workloads with best-efforts labels (i.e., research and model development, internal usage, first-party inference, and so on) based on its metadata. Usage on third-party cloud compute is reported to us by the providers and folded in. Most of the work of this exercise was stitching these existing sources together.

We then used Claude to classify each workload as either safety work or AI R&D via a prompted classifier. Safety work was defined as work whose dominant purpose is making AI systems safer, more understandable, or more secure. Everything else, including capability research, training production models, product development, and developer tooling, was counted as AI R&D. Work that helps capability as much as it helps safety was also counted as AI R&D, so the safety share is conservative.

For research training and evaluation runs, we built a classifier that reads the run’s metadata and the code it used, and returns a classification, a justification, and a confidence level. Rather than classify all of the week’s almost 10,000 runs, we sampled about 14% of them, weighting the sample toward the runs that used the most compute, so that the result reflects where the compute actually went, rather than how many runs there were. For inference for AI research agents, a variant of the same classifier read the agent’s session transcript. Where transcripts were inaccessible (usually due to the work being compartmentalized), we classified them by the user’s team or conservatively defaulted to classifying them as AI R&D. We plan to refine this pipeline so that an independent third-party could re-run the classifier on a random subsample of jobs and transcripts and check both the sorting and the totals.

