---
id: collect-240926-datacamp/datacamp/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-l
title: "pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "cost", "distribution", "reasoning"]
source: docs/RAG/clean_en/datacamp/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel.md
source_anchor: ""
source_lines: [1, 173]
sha256: b65b6a2f5b141439b9a7ed622130ef154a6fcd28a6e8b13d501efd6c3849e127
---

# pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel

<!-- source: https://www.datacamp.com/fr/blog/why-ai-agents-struggle-more-with-data-science-code -->

Cursus

I analyzed hundreds of GitHub repositories to understand why AI coding agents often appear much more competent in software engineering than in data science. What I found shows that it is not just a tooling gap. It points to a deeper difference: where meaning resides in each type of code.

## What this article is about

If you have used an AI coding agent on a software engineering codebase, you have probably seen its effectiveness. The agent navigates the architecture, follows abstractions, and makes changes that fit surprisingly well with the rest of the system.

Then you open a data science notebook, and the experience often changes.

The agent still knows how to write valid code. It still knows how to follow instructions. But it does not always grasp the essential point: why this dataset, why this filter, why this time window, why this output changed the direction of the analysis.

It treats the notebook as a software project, when it is only part of one.

I wanted to understand why. So I analyzed hundreds of GitHub repositories in data science and software engineering, measuring entropy, reference patterns, and coupling behaviors.

My findings surprised me, and I think they have real implications for anyone building or using AI agents in analytical work.

## The entropy inversion: data science code is not what it seems

### What I actually measured

I measured Shannon entropy at three levels of abstraction for each repository: at the character level, at the token level, and at the AST level. Each captures a different dimension of code variation.


Code entropy distribution: data science vs software engineering (violins)

The pattern that emerges is what I call an entropy inversion.

### The surface seems complex, the structure often is not

At the character level, data science code tends to show higher entropy than software engineering code. This seems intuitive. Data science is full of varied column names, dataset labels, ad hoc variables, and domain-specific identifiers that make the code noisy and irregular.

At the token level, the two domains are much closer. They rely largely on the same syntactic building blocks.

But at the AST level, where structural diversity is observed, the picture reverses.

Software engineering code generally encodes far more structural variation. It creates distinct behaviors through abstractions, interfaces, modules, and internal logic.

Data science code, by contrast, often reuses a smaller set of operations in changing contexts: load, filter, group, aggregate, visualize, inspect, adjust.

In short, data science code often appears more complex on the surface, while software engineering code carries more complexity in its structure.

This is not just a difference in style. It reveals a deeper divergence in how each type of work stores meaning.

## Indexical vs symbolic: where meaning resides

### Two different natures of code

The entropy inversion is best explained by considering what each type of code actually does.

In software engineering, meaning is often compressed into structure. Functions, interfaces, modules, types, and class boundaries do most of the work. Once these abstractions are established, they stabilize behaviors and reduce future uncertainty.

A large part of the meaning is inside the code itself.

In data science, meaning remains far more closely tied to external context. It depends on the dataset, the columns, the intermediate outputs, the assumptions behind a transformation, and the question that evolves throughout the analysis. The code does not only express logic; it refers to a specific analytical situation.

This difference helps explain why agents often behave so differently in the two domains.

### You can see it in where the code "points"

I also measured the densities of external and internal references per 100 lines of code in both groups.


External and internal reference densities with statistical significance

The difference is clear. Data science code more often points outward: datasets, tables, columns, temporary objects, and states that exist outside the code itself.

Software engineering code is more self-referential. It more often builds meaning by pointing to functions, classes, modules, and abstractions defined elsewhere in the codebase.

"Data" code points outward. Software code points inward.

And this matters for agents. In one case, much of the relevant context is in the codebase. In the other, it resides in the surrounding analytical state.

## Why this makes abstraction behave differently

### The coupling problem in notebooks

One of the interesting observations concerns coupling. Notebook-style workflows show significantly tighter coupling per 100 lines of code than software engineering codebases.

This is not necessarily bad design. It reflects something fundamental about exploration: the question itself is still evolving. You test hypotheses, follow unexpected results, check edge cases, and change direction as you learn.

In this context, abstraction does not necessarily pay off as it does in software engineering. Structuring too early can reduce options before you even understand the essentials. Tighter coupling is often a consequence of exploration, not simply bad engineering.

### The role of comments and state

There is also an important difference in how the two domains explain themselves.

In software engineering, code often explains itself through its structure. Types, interfaces, and abstractions carry most of the meaning, and comments are generally secondary.

In data science, comments, outputs, and state often carry part of the meaning. A note like "removed outliers beyond the 99th percentile, confirmed this does not affect the main cohort" is not mere documentation.

It captures an analytical decision that may not be recoverable from the code alone. A table or chart in the middle of the notebook can explain why the rest of the analysis exists at all.

This matters for agents. An agent that ignores comments, inline results, and the evolving state of a notebook misses part of the analysis. In a software codebase, this information is often peripheral. In data science, it often is not.

## What this implies for AI agents

### By default, agents are better suited to one of the two domains

Practical consequence: agents work best when their tools match where the meaning is.

In software engineering, an effective agent navigates the architecture. It follows the call graph, respects interfaces, and maintains consistency of changes across the codebase. This works because much of the meaning is encoded in the structure.

In data science, an effective agent must do more than navigate the code. It must understand what the data actually contains, track state across steps, trace the provenance of variables, and reason about why a transformation or filter was applied several steps earlier.

An agent designed for software engineering does not automatically transfer to data science. The problem is not merely syntactic. The underlying informational structure is different.

### Why notebook-style workflows persist

This also helps explain what often confuses engineers: why notebooks remain central despite their obvious limitations.

The reason is that notebooks keep meaning as close as possible to the data while the analytical question evolves. A notebook is not simply a poorly structured Python module. It is a different artifact, designed for a different phase of work. It keeps code, outputs, and decisions in the same place while the analysis is being built.

This is why a notebook can seem intuitive to an analyst immersed in the context, and awkward for an agent that sees only the code.

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
