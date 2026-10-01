---
id: collect-261001-ia-llm/ia-llm/evaluating-claude-with-biomysterybench-2
title: "evaluating-claude-with-biomysterybench"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "benchmark", "benchmarks", "opus 4", "pretraining", "research"]
source: docs/RAG/collect-261001-ia-llm/evaluating-claude-with-biomysterybench.md
source_anchor: ""
source_lines: [37, 72]
sha256: b80d482d99a8d73a3c479b6f22fa357bd9d2fe858ed2a7561e5c185c955930da
---

# evaluating-claude-with-biomysterybench

BioMysteryBench consists of 99 questions from various fields of bioinformatics, written by domain experts. Experts were instructed to gather a dataset, and create a question based on controlled, objective properties of the data, rather than unverifiable scientific conclusions. By deriving answers from an experimental or clinical finding, it was possible to develop questions without requiring they be human-solvable.

Although these questions are created from verified ground truth, they still have the same flavor as tasks a research scientist would want to answer. Claude is tasked with each question and put in a container with a minimal set of canonical bioinformatics tools, the ability to install additional tools via pip and conda, and permissions to access canonical bioinformatics databases (such as NCBI and Ensembl) to download additional resources such as reference genomes.

BioMysteryBench has a tetrad of unique properties that make it a particularly powerful benchmark for science, and tackle the challenges above:

In developing this eval, questions were primarily derived from raw or minimally processed DNA or RNA sequencing data since this is where many biological processing pipelines begin (WGS, scRNA-seq, methylation, ChIP-seq, metagenomics, Hi-C), and also included several questions drawn from proteomics and metabolomics.

Questions developers came up with included:

To minimize inherently unsolvable questions while still leaving room for those that might be AI-solvable, we required each question author to submit a validation notebook demonstrating that the signal does, in fact, exist in the data (even if finding it from scratch might be difficult). Think of this as the high-school algebra principle: verifying an answer is much easier than deriving one.

For each question, we tasked up to five domain experts to answer the question from scratch. Once a question was answered correctly by at least one human, we considered it human-solvable. BioMysteryBench contained 76 such tasks.

Sometimes Claude mirrored human strategies. Perhaps humans have landed on a near-optimal approach, or because the method is well-represented in pretraining data.

Other times, Claude took a completely different route, illustrating there is no strictly correct way to solve these problems and that models may have genuine preferences that diverge from ours.

The examples above showcase a particularly interesting strategy: whereas our human experts used algorithms or databases to identify and annotate properties of a dataset, Claude intuitively recognizes certain patterns or sequences. Admittedly, such clever abstraction is not entirely unique to AI—the first eukaryotic promoter, for example, was discovered when a scientist noticed the sequence “TATA” appearing over and over in sequences upstream of genes. *Intuition* like this has been difficult to build into traditional biology machine learning models, but LLMs might be able to turn up patterns like this at unprecedented scale.

That left us with a set of questions that could not be solved by our panel of experts. This could mean (1) the question was malformed or broken, (2) the question is inherently unsolvable (e.g.,the signal isn’t in the data), or (3) the question is theoretically solvable but humans lack the knowledge required to solve it. After QC’ing with benchmarkers and additional experts, we removed 4 questions that were due to (1), leaving 23 human-difficult questions.

Interestingly, Claude Sonnet 4.6 and more capable models were able to solve significant fractions of human-difficult problems, with Claude Mythos Preview topping out at a 30% solve rate. So what exactly is Claude doing that humans aren’t?

Analyzing transcripts from Opus 4.6, we identified two primary strategies used by Claude compared to humans: one is fairly AI-specific: Claude’s vast underlying knowledge base contains information about structural biology, molecular profiles, and meta-analysis from hundreds of thousands of papers. The other strategy is something we human scientists could learn from: when Claude is uncertain about an answer, it layers multiple methods and combines different lines of evidence to arrive at a conclusion.

In some of the human-difficult tasks, Opus’s vast underlying knowledge base helped it solve the problem. Tasks that would require a human expert to run a meta-analysis or stitch together databases, Opus solved directly by combining its internal knowledge of mechanisms and ontologies with live analysis. Often, this allowed Claude to solve human-unsolvable tasks! Here are a few examples:

Even though prior knowledge seemed overwhelmingly helpful to Claude, we saw one interesting case (in the human-solvable set) where this became its downfall:

When Opus 4.6 was not confident about an answer, it often tried multiple different ways of solving the problem and chose the answer that multiple approaches converged on.

Like many of the benchmarks we've discussed, BioMysteryBench has its own limitation: for tasks that neither humans nor models have solved, we can never be fully certain whether they're impossible or just extraordinarily difficult. The validation notebooks help ensure the signal is there and the data is well-formed, but they do not guarantee a model or human can find the answer from scratch. So we ask both our models and our human benchmarkers not to be too frustrated if, a year from now, no one has solved the human-difficult set. That uncertainty is also part of what makes the benchmark exciting: a more scientifically capable model might be the first to crack a problem that no human or model has solved before.

Claude showed solid improvement across generations and did well enough at both the human-solvable and human-difficult tasks that we thought it would be interesting to let Claude Mythos Preview conduct some of its own scientific analysis. Here are a couple of additional insights about its predecessor Claude’s performance on BioMysteryBench:

