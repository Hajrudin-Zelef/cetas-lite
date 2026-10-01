---
id: collect-261001-ia-llm/ia-llm/how-claude-is-uplifting-biomolecular-modeling-2
title: "how-claude-is-uplifting-biomolecular-modeling"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Nvidia", "United States", "Z.ai"]
dates: []
keywords: ["claude", "agents", "compute", "cost", "cyber", "glm", "gpu", "gpus", "mythos 5", "nvidia", "opus 5", "protein"]
source: docs/RAG/collect-261001-ia-llm/how-claude-is-uplifting-biomolecular-modeling.md
source_anchor: ""
source_lines: [33, 55]
sha256: 235292116c3efce412eae1a58fc23416d723c9d9f085aec796c92caba1c5e170
---

# how-claude-is-uplifting-biomolecular-modeling

To test the limits of Claude’s optimizations, we asked Claude to predict structures of a greater size than anything that had previously been achieved. Using a single 8-GPU B300 node, Claude generated predictions of entire viral capsids and protein compartments ranging in size from more than 31,000 to more than 70,000 tokens. These systems are nearly two orders of magnitude larger than the training context of these structure prediction models, and, perhaps unsurprisingly, are not predicted correctly. However, the barrier to inferencing at this scale has been significantly lowered now that it takes just one NVIDIA B300 node, suggesting that with improved tools researchers will soon be able to computationally model an increasingly complex set of biological systems.

In our earlier work on protein design, we provided Claude with an approximately 16,000-word prompt that encouraged it to utilize sub-agents and spend up to $10,000 per target on Modal (roughly 2,500 NVIDIA H100 GPU hours) in a 24-hour span. Here, we gave a single Claude model access to one NVIDIA H200 and 24 hours of wall time, a prompt of about 1,100 words, and a reference sheet for the pre-installed tools, with no sub-agents and no human steering the designs.

We ran three Claude models (Mythos 5.1, Mythos 5, and Opus 5) against 16 targets with the accelerated biomolecular models described in this post. We scored designs by ipSAE, an *in silico* score that has been shown to be predictive of binding in the wet lab. Averaged over 16 targets, the median-scoring and highest-scoring designs from all three Claude models evaluated achieve approximately the same ipSAE values as our earlier Mythos 5.1 campaigns despite using about two orders of magnitude fewer GPU hours. We also considered Claude token costs and found that with a combined spend of approximately $150 on GPUs and tokens, we can achieve *in silico* performance matching the levels of our previous campaigns.

The optimizations described above help us predict and design molecules more efficiently, while unlocking capabilities that would have otherwise been resource-prohibitive. To demonstrate the uplift they provide and the impact of Claude on molecule design more broadly, we’re partnering with Adaptyv Bio to launch a protein design competition. We’ve selected five problems at the frontier of today’s protein design capabilities, including challenges such as species cross-reactivity, pH-sensitivity, and peptide-MHC specificity, as well as difficult targets such as GPCRs.

With the Adaptyv team, we’ll be experimentally validating over 5,000 designs submitted by the community against these problems. We will be providing up to $1 million in Claude credits and additional funds for experimental validation at Adaptyv for participating researchers, Modal will provide up to $250,000 in compute credits, and Twist Bioscience will provide DNA for the competition. You can find more information, including eligibility criteria (here) and (apply here).

We have also begun to provide frontier AI capabilities to life scientists for biology-related work via our Life Sciences Verification Program. We recently enrolled our first group of organizations, and opened up the program in public beta today. You can find more information (here).

The following resources provide further technical depth and more detailed information about the results described above:

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Like Claude Mythos Preview, GLM-5.3 has strong capabilities for autonomously building end-to-end cyber exploits. But GLM-5.3 is unlike other frontier models in that it has been released without meaningful safeguards to limit misuse.

Read more
