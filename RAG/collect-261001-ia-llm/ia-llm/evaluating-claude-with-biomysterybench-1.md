---
id: collect-261001-ia-llm/ia-llm/evaluating-claude-with-biomysterybench-1
title: "evaluating-claude-with-biomysterybench"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["claude", "agent", "agentic", "benchmark", "benchmarks", "protein", "reasoning", "research"]
source: docs/RAG/collect-261001-ia-llm/evaluating-claude-with-biomysterybench.md
source_anchor: ""
source_lines: [1, 36]
sha256: 78930b3063f5a717408ab7c7555c68184fada63e0f43f5c84a517ac2488aa7ac
---

# evaluating-claude-with-biomysterybench

## Subscribe to Anthropic Science

Features on AI-assisted discoveries, practical workflows, and field notes across the sciences.

*In this post, Brianna*, *a researcher on the discovery team, shares results from a recent bioinformatics benchmarking effort.*

Almost as soon as large language models could hold a conversation, people started asking how they’d stack up against human experts. Could models pass the bar exam? Could they answer medical licensing questions, or solve Olympiad math problems? Such *benchmarks*—self-contained sets of human-vetted problems designed to evaluate a capability of a model—have now become a source of competition across AI developers, reported in model release system cards and tracked on many online leaderboards. 

Competition aside, benchmarks help us tackle an important question: whether models are capable and reliable enough to support, or even produce, professional-level work. Scientists are using models to write code for analysis pipelines, propose hypotheses, and draw conclusions from data with the long-term aim of accelerating innovation and discovery. But exactly how proficient is AI in science right now, and how quickly are Claude and other models improving? 

To answer this, the research community has built several benchmarks. MMLU-Pro tests expert-level knowledge and reasoning questions. GPQA poses graduate-level, "Google-proof" questions in biology, physics, and chemistry. LAB-Bench tests biology-specific knowledge work—reading the literature, interpreting figures, reasoning about protocols. Although these benchmarks were developed in the “chatbot” era, they’ve persisted into the agent and tool-use era, joined by even more difficult scientific reasoning evals like FrontierScience and Humanity's Last Exam, because knowledge and reasoning remain a vital measure of scientific capability.

Still, many real-world scientific tasks demand more than that. They require reading papers, querying databases, running experiments, coding and analysis. Now that models can do many of these things, benchmarks have evolved to reflect these workflows. BLADE tasks a model with a dataset and an open-ended task, and checks if the model takes similar analysis steps to a human scientist. BixBench uses biological datasets, and grades models on whether their conclusions line up with scientists’. In SciGym, the model is dropped into a simulated biology lab, where it has to design and run its own experiments to uncover a hidden mechanism.

These benchmarks move us closer to measuring scientific capability, but they don't quite test whether a model can devise creative solutions to the messy, open-ended problems that define research. This is why we developed BioMysteryBench, a bioinformatics benchmark that tasks Claude with the analysis of real-world datasets, while tackling some of the challenges inherent in evaluating complex and noisy biological systems. We learned that Claude's scientific capabilities in biology are improving rapidly across generations, that current models perform on par with human experts, and that the latest generations solved many problems that a panel of human experts could not, sometimes using very different strategies.

Doctors have board exams and lawyers have the bar, but there’s no standardized test for becoming a scientist. The same problem shows up with AI. Despite how badly we want to use these models for science, no agentic science benchmark has become quite as canonical as SWE-bench is for software engineering. We think that’s because scientific research, particularly biology, has several properties that make it especially hard to evaluate via a benchmark.

If there were only one right way to answer a research question, PhD students would earn their degrees in a matter of months, corporate R&D departments wouldn’t exist, and no science fair poster would need a “Methods” section. How a scientist tackles a problem depends on their skills and background, the resources available to them, and their research taste.

Consider a seemingly straightforward question that has mystified metabolic researchers for years: why do some type 2 diabetics respond to the oral drug metformin while others do not? In order to answer this question, you could run a genome-wide association (GWAS) study on responders vs. non-responders and look for predictive genetic variants, or sequence the gut microbiomes of both groups, since metformin is partly metabolized by gut bacteria. Both are reasonable directions, and how you proceed will often just depend on expertise and resources.

BixBench handles this well by grading the model on its conclusions rather than the method used to reach them. The tradeoff is that those conclusions were produced by an individual scientist who made a series of subjective choices along the way that may have shaped the answer itself. This, in turn, has its own pitfalls…

Even within a chosen research direction, individual decisions can be highly subjective: one scientist may approve of a decision, while another researcher may have serious objections. Just ask any frustrated author who’s gotten conflicting suggestions from a round of peer review! Making this all the more difficult is the fact that biological datasets are often noisy enough that small differences in research decisions can lead to entirely different conclusions about the data.

In the decade-long search for metformin response predictors, slight differences in study design have led to entirely different conclusions about metformin response. A 2011 paper reported a variant that predicts metformin response that replicated in two cohorts, with a plausible mechanism involving AMPK activation. A year later, the Diabetes Prevention Program tested the same variant in pre-diabetics and found nothing. Finally, rather than spinning up their own study, a 2012 meta-analysis pooled five cohorts and once again decided the 2011 paper's effect was real but more modest than originally reported.

SciGym's clever way of handling such ambiguity is by choosing tasks with a well-defined answer. Because the underlying biological network is a simulator, there is, in fact, a ground-truth, and noise is controlled rather than inherited from a messy living system. However, it's unclear how closely performance in a simulated lab tracks performance on real data.

The research tasks where models could have the greatest impact are those that humans alone have yet to solve. And ultimately, those are precisely the tasks we’d like to be able to evaluate models on. What, for example, is the mechanism of action of metformin? Thirty years after its development, the field still is not certain of the primary target. Discovering it, or finding a homolog of metformin that is cheaper to synthesize and more stable, would be enormously consequential.

Machine learning has long tackled problems humans perform poorly at, like sequence prediction and protein modeling, by leaning on experimental data instead of expert intuition. ProteinGym scores models on mutation fitness effects using Deep Mutational Scanning experiments as ground-truth, and the long-running CASP competition evaluates protein folding against unpublished crystal structures. Both are grounded in experimental measurements no expert would trust themselves to reproduce. However, these benchmarks are built around a narrow set of tasks and don't capture the breadth of bioinformatics work we actually want to measure.

Because no benchmark perfectly handles the three aforementioned challenges, we developed BioMysteryBench. BioMysteryBench uses messy, real-world bioinformatics data, without allowing the complexity and challenges inherent in this data to corrupt the quality of the evaluation.

