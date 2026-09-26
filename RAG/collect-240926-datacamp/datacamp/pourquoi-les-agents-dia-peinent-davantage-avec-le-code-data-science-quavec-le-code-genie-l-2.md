---
id: collect-240926-datacamp/datacamp/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-l-2
title: "pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "cost", "reasoning"]
source: docs/RAG/clean_en/datacamp/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel.md
source_anchor: ""
source_lines: [119, 173]
sha256: 7addefb22d258c84f92a8a2baff171e41b3526c554f81d8cc892df8694f441ee
---

# pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel

Refactoring a notebook into clean, modular code before the analytical question has stabilized is not always an improvement. It can detach the logic from the context that gave it meaning.

### The opportunity: bridging the gap with production

The real opportunity for AI agents in data science is not to copy what works in software engineering. It is to help bridge the production gap: the distance between an insight discovered during exploration and a reproducible, deployable workflow.

This gap exists because notebooks preserve analytical context at the cost of structural cleanliness. An agent capable of understanding both the analytical context and the structural requirements of production systems could bridge this gap without forcing the human analyst to leave exploratory mode too early.

This is a harder problem than navigating a codebase. But it is also the most important one.

## Key takeaways

In summary, what the analysis of the studied repositories suggests:

- Data science code often exhibits higher entropy on the surface, while software engineering code often exhibits higher entropy at the structural level.
- Data science code more often refers to external state (datasets, columns, tables, intermediate outputs), while software engineering code is more self-referential.
- Notebook-style workflows tend to show tighter coupling, often a consequence of exploration rather than simply bad practice.
- Comments, outputs, and evolving state are part of the semantic content of data science work, not mere peripheral details.
- Agents optimized for software engineering contexts will often underperform in data science if they are not designed to reason about data, outputs, and context in addition to code structure.

## Conclusion

When I launched this analysis, I expected to find that data science code was simply less well structured than software engineering code. Instead, I found that it is structured differently for good reasons, with meaning often residing elsewhere.

This has concrete consequences for anyone designing or evaluating AI agents for analytical work. The real question is not whether an agent can write correct Python, but whether it understands what the data is, why the analysis takes the form it does, and what decisions were made along the way that are not visible in the code alone.

Data science code is not immature software engineering. It is an optimization of a different kind: more dependent on external state, less reliant on durable internal structure during exploration, and shaped as much by context as by code.

Once one sees this, notebooks stop looking like failed software projects and appear for what they are: work surfaces for reasoning about changing data.

**I continue this analysis with a broader corpus of at least 1,000 repositories and additional metrics. If you want to take a deeper look at how AI agents are designed to work with data in context (including persistent execution architectures that maintain data state across sessions), DataCamp's AI agent fundamentals track is an excellent starting point.**

## FAQs

### What is code entropy and why does it matter for AI agents?

**Shannon entropy applied to code is a way of measuring variation or unpredictability at a given level of abstraction. Higher entropy at the structural level may indicate that the code expresses a broader range of internal behaviors. For AI agents, these patterns matter because they determine the kind of context the agent needs to understand.**

### What is the difference between indexical and symbolic code?

**This distinction is a useful shorthand. Software engineering code often carries more of its meaning internally through functions, interfaces, and abstractions. Data science code depends more heavily on external context: datasets, columns, outputs, and the evolving state of the analysis.**

### Does this mean data science code is lower quality than software engineering code?

**No. This is not about saying one is better than the other. They are optimized for different goals. Exploration in data science often prioritizes speed of iteration and preservation of context, while software engineering generally prioritizes structure, reuse, and maintainability.**

### Why do AI coding agents often seem less effective in data science notebooks?

**Because many current agents are optimized for navigating code structure. In data science, a large part of the meaning lies outside the code itself: data state, outputs, comments, and analytical decisions. If an agent cannot reason about that context, it misses part of the task.**

### What would a native AI agent for data science look like?

**It would need to maintain awareness of data state, outputs, and variable provenance, not just source code. It would need to treat comments and intermediate results as meaningful context and help move exploratory work into reproducible production workflows.**

Jason Hillary co-founded Zerve after seeing how much friction slowed down work around data. He holds a PhD in engineering from the University of Limerick and has spent years designing AI systems that are actually operational in production, not just in theory. Before creating Zerve, Jason worked across the entire data and AI ecosystem, with the goal of making technical tasks less painful and more productive.
