---
id: collect-261001-ia-llm/ia-llm/glm-5-3-and-the-spread-of-advanced-cyber-capabilities-2
title: "glm-5-3-and-the-spread-of-advanced-cyber-capabilities"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "United States", "Z.ai"]
dates: []
keywords: ["cyber", "glm", "benchmarks", "claude", "cost", "exploit", "gpu", "mythos 5", "open-weight", "prefill", "research", "safeguards"]
source: docs/RAG/collect-261001-ia-llm/glm-5-3-and-the-spread-of-advanced-cyber-capabilities.md
source_anchor: ""
source_lines: [31, 55]
sha256: 254ed55ee2cdd554d7560b2a1ed114399d32c572c6091ecc6e72e0d769401a7b
---

# glm-5-3-and-the-spread-of-advanced-cyber-capabilities

To research how far abliteration allows attackers to bypass GLM-5.3’s safeguards, we produced an abliterated copy ourselves, and then ran it on three public benchmarks (JailbreakBench, HarmBench, and StrongREJECT) that measure how often a model complies with clearly harmful requests. Abliterating the model took our team—which had never previously attempted this task—about 2,200 GPU hours at a computation cost of roughly $4,400.<sup>3</sup> Abliterating GLM-5.3-Flash took about 600 GPU hours. The edit took GLM-5.3’s refusal rate from above 90% to about 3% and 2% on the first two benchmarks (JailbreakBench and HarmBench) and to 12% on the third (StrongREJECT). Abliteration did not significantly reduce the model’s capabilities: on GPQA-Diamond, an evaluation that measures general scientific capabilities, the standard and abliterated models scored the same results; on a tested subset of the CyberGym evaluations, the abliterated version scored a few percent lower (as shown in the chart below).

In our testing, we observed that GLM-5.3’s safeguards can also be circumvented *without* using an abliterated version of the model. We placed the model in a simulated world<sup>4</sup> in which it was given overtly malicious requests to attack critical systems. Out of the box, GLM-5.3 refused in all trials (as with the other models we tested). But we identified several simple ways to bypass the GLM models’ safeguards, such that it would respond to these requests in most or all cases. These include:

In our testing, none of these techniques got safeguarded Claude models to carry out the harmful tasks we tested. Claude’s safeguards blocked the requests that used deceptive prompts. The Anthropic API provides would-be attackers with no way to prefill Claude’s thinking. And since Claude’s weights are not provided to users, they cannot be abliterated to change Claude’s behavior.

To demonstrate how the abliterated version of GLM-5.3 is willing to engage in harmful tasks, we highlight one quote from the chain of thought that it generated:

GLM-5.3 will likely give malicious actors access to capabilities that will allow them to find and exploit cyber vulnerabilities without meaningful restrictions. This is unlike any other similarly capable AI model, all of which were released with safeguards or through limited access programs. The release of GLM-5.3 is a meaningful step change in the cyber capabilities available to attackers. Anthropic and other US AI labs have published recent reports that disclose how cyber attackers have tried to use AI systems. Given this evidence, we think it’s likely both state and non-state actors will use models like GLM-5.3 to cause real-world harm.

On the other hand, models with this level of capability can also be used by defenders. Our view is that cyber defenders should use the best available tools that meet their needs. We're working to safely expand access to Claude's cyber capabilities to as many defenders as we can. Cyber defenders face attackers who will use every capable tool they can, and we believe defenders should be equipped with frontier models that are at least as good as those their adversaries are using.

Through Project Glasswing (and other efforts, like Patch the Planet), cyber defenders have made meaningful progress towards securing critical systems in advance of this moment—but much work remains to be done. While vetted defenders can now use even more advanced models like Claude Mythos 5.1 through our trusted access programs, a critical threshold in freely accessible capabilities has now been crossed. GLM-5.3 underscores the urgency of expanding access to advanced frontier models to a broader set of entities to empower cyber defenders.

Governments should conduct safety testing on sufficiently capable AI models, including successors to GLM-5.3. Without high-quality evaluations from independent sources, the impact of these capabilities might not become fully clear to model developers until it is too late. As AI developers across the world build increasingly capable open-weight models, we hope they work to appropriately safeguard these capabilities and prevent misuse.

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Guest writer and physicist Matt von Hippel shares what happened when he issued a challenge to AI companies to solve a problem in his former subfield of theoretical physics.

Read more
