---
id: vague2-datacamp/datacamp/why-ai-agents-struggle-more-with-data-science-code
title: "Pourquoi les agents d’IA peinent davantage avec le code data science qu’avec le code génie logiciel"
domain: datacamp
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: ["agent", "agents"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/why-ai-agents-struggle-more-with-data-science-code.md
source_anchor: ""
source_lines: [1, 54]
sha256: e4c29c8c3c0248a9be25709daf1f274d4277bdef5517c941817d65e52d221b3e
---

# Pourquoi les agents d’IA peinent davantage avec le code data science qu’avec le code génie logiciel

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/why-ai-agents-struggle-more-with-data-science-code
- **Site** : DataCamp
- **Type** : Article (analysis)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The author analyzed hundreds of GitHub repositories to understand why AI coding agents often seem much more competent at software engineering than at data science. The finding is not just a tooling gap but a deeper difference: where the meaning lives in each type of code.

The article introduces the concept of an **entropy inversion**. Measuring Shannon entropy at three abstraction levels (character, token, and AST) across repositories, the author found: at the character level, data science code tends to have higher entropy (varied column names, dataset labels, ad hoc variables, domain-specific identifiers); at the token level, both domains are similar; but at the AST level, the picture inverts—software engineering code encodes far more structural variation (abstractions, interfaces, modules, internal logic), while data science code reuses a smaller set of operations (load, filter, group, aggregate, visualize, inspect, adjust) in changing contexts. So data science code looks more complex on the surface, while software engineering code carries more complexity in its structure.

This maps to an **indexical vs symbolic** distinction. In software engineering, meaning is compressed into structure—functions, interfaces, modules, types, class boundaries stabilize behavior. In data science, meaning is tightly bound to external context: the dataset, columns, intermediate outputs, assumptions behind transformations, and the evolving question. Measuring external vs internal reference densities per 100 lines of code confirmed the difference: data science code refers more often outward (datasets, tables, columns, temporary objects, external state), while software engineering code is more self-referential (functions, classes, modules defined elsewhere in the codebase). "Data" code points outward; software code points inward.

On **coupling**: notebook-style workflows show significantly tighter coupling per 100 lines than software engineering codebases. This is not necessarily bad design—it reflects exploration where the question itself is still evolving, and structuring too early can reduce options. The article also notes that in data science, comments, outputs, and state carry part of the meaning (e.g., a note about outlier removal at the 99th percentile captures an analytical decision not recoverable from code alone; a table or chart mid-notebook explains why the analysis continues). Software engineering code largely explains itself via structure, so comments are peripheral; in data science they often are not.

**Implications for AI agents**: agents work best when their tools match where meaning lives. Software engineering agents navigate architecture, call graphs, and interfaces. Data science agents must understand what the data actually contains, track state across steps, trace variable provenance, and reason about why a transformation was applied. A software-engineering-oriented agent does not automatically transfer. This also explains why notebook-style workflows persist: notebooks keep meaning close to the data while the analytical question evolves—they are a different artifact for a different phase, not just badly structured Python. The real opportunity for AI agents in data science is not copying software engineering but bridging the productionization gap between an exploratory insight and a reproducible, deployable workflow, by understanding both analytical context and structural production requirements.

## Key points

- Analysis of hundreds of GitHub repos reveals an "entropy inversion": data science code has higher surface (character) entropy but lower structural (AST) entropy than software engineering code.
- Software engineering meaning lives in structure (symbolic); data science meaning lives in external context (indexical)—datasets, columns, outputs, evolving questions.
- Data science code refers more often outward; software engineering code is more self-referential.
- Notebook workflows show tighter coupling, a consequence of exploration rather than merely poor practice.
- In data science, comments, outputs, and state carry semantic content, not just peripheral details.
- Agents optimized for software engineering underperform in data science unless designed to reason about data, outputs, and context.
- The key opportunity is bridging the exploration-to-production gap without forcing analysts out of exploratory mode.

## Technical data / figures

| Level of abstraction | Data science code | Software engineering code |
| --- | --- | --- |
| Character entropy | Higher (varied names, ad hoc identifiers) | Lower |
| Token entropy | Similar | Similar |
| AST (structural) entropy | Lower (small reused operation set) | Higher (abstractions, modules, interfaces) |

| Dimension | Data science | Software engineering |
| --- | --- | --- |
| Meaning location | External context (indexical) | Internal structure (symbolic) |
| Reference direction | Outward (datasets, columns, state) | Inward (functions, classes, modules) |
| Coupling (notebooks) | Tighter per 100 lines | Looser |
| Role of comments/outputs/state | Part of semantic content | Peripheral |

- Method: Shannon entropy at character, token, and AST levels; external vs internal reference densities per 100 lines of code; coupling measures.
- Corpus: hundreds of GitHub repositories (expanding to 1,000+).

## Why this source matters for the RAG

It offers a novel, well-reasoned explanation of why AI coding agents underperform on data science code, valuable for agent-design, notebook-vs-production, and AI-assisted data-work questions. It reframes notebooks as purposeful analytical artifacts rather than poorly structured software.
