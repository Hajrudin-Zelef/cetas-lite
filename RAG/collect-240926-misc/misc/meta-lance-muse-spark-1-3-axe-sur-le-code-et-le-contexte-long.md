---
id: collect-240926-misc/misc/meta-lance-muse-spark-1-3-axe-sur-le-code-et-le-contexte-long
title: "meta-lance-muse-spark-1-3-axe-sur-le-code-et-le-contexte-long"
domain: blogdumoderateur
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "OpenAI"]
dates: []
keywords: ["muse", "agent", "agentic", "benchmarks", "claude", "gemini", "gpt-5.6", "muse spark", "opus 5", "reasoning", "sol"]
source: docs/RAG/clean_en/misc/meta-lance-muse-spark-1-3-axe-sur-le-code-et-le-contexte-long.md
source_anchor: ""
source_lines: [1, 20]
sha256: 05b1baf4b4deda2950693301f6e3a05f0c2d0d0bbf1923cd42211e500e528489
---

# meta-lance-muse-spark-1-3-axe-sur-le-code-et-le-contexte-long

<!-- source: https://www.blogdumoderateur.com/meta-muse-spark-1-3/ -->

As with Google and Gemini, Meta is continuing with its in-house models. A month after Muse Spark 1.2, released at the same time as its Muse Code programming agent, Meta Superintelligence Labs is rolling out Muse Spark 1.3. The model is accessible in Muse Code and via the Meta Model API.

## A model designed for long projects

Meta positions its update around sustained performance over time rather than raw power. The model is presented as capable of carrying out several tasks in parallel within a single conversation thread, building its own context from messy or contradictory sources, and correcting gaps in its plan along the way. The company also highlights more cooperative behavior. The model asks questions when the request is ambiguous, reaches out to the user when it gets stuck, and asks for confirmation before an action with serious consequences.

Muse Spark 1.3 is designed to optimize long-term project management by facilitating collaboration with users and managing multiple workflows simultaneously, according to the press release.


The most notable argument concerns awareness of its own limitations. Meta says it trained the model to better grasp "*what it can and cannot do, what it knows and does not know*," rather than inventing a result when it hits an obstacle. On code, the group claims a less chatty style and fewer unnecessary loop iterations, with roughly 20% fewer tool calls and 25% fewer tokens according to comparisons run by its own engineers.

## Comparisons backed by a mode that is still locked

As with every release, Meta publishes its own score table, against OpenAI's GPT-5.6 Sol and Anthropic's Claude Opus 5. The model leads its two rivals on code and on long-context management. However, it takes no first place on the six agentic evaluations selected, even though these tasks are one of the selling points of the announcement.

The table above all requires a reading caveat. The scores attributed to Muse Spark 1.3 are those of its "max reasoning" mode, which is not accessible at launch: Meta is delaying it while it conducts additional safety tests. The 1.2, meanwhile, is measured in a lower reasoning mode. Part of the gap shown between the two versions therefore stems from the setting, not from the model generation. Further proof of the limits of benchmarks published by vendors.

Meta also announces larger models and the release of the weights of a version of Muse Spark, without a timeline. The group is thereby renewing a commitment made in August, when Mark Zuckerberg's superintelligence manifesto was published, accompanied by the open weight model Muse Glimmer.
