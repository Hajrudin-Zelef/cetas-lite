---
id: collect-240926-datacamp/datacamp/claude-mythos-5-fonctionnalites-benchmarks-et-capacites
title: "claude-mythos-5-fonctionnalites-benchmarks-et-capacites"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "OpenAI", "Stripe"]
dates: ["2026-04", "2026-06-09", "2026-07-01"]
keywords: ["benchmark", "benchmarks", "claude", "agent", "agentic", "cyber", "cybersecurity", "export control", "fable 5", "gemini", "guardrails", "jailbreak"]
source: docs/RAG/clean_en/datacamp/claude-mythos-5-fonctionnalites-benchmarks-et-capacites.md
source_anchor: ""
source_lines: [1, 175]
sha256: 0de30fc8d3e0279c3cb1c7b70d79b9de107af3f80a5122fa07c95dd209dd3c8f
---

# claude-mythos-5-fonctionnalites-benchmarks-et-capacites

<!-- source: https://www.datacamp.com/fr/blog/claude-mythos-5 -->

Cursus

**Update (July 1, 2026):** access to Claude Fable 5 has been restored and is now available worldwide on Claude Platform, Claude.ai, Claude Code, and Claude Cowork, following the lifting of the export control order. Mythos 5 remains restricted to duly approved Project Glasswing partners.

Anthropic launched two models on June 9, 2026: Claude Fable 5, the public version of the Mythos class with conservative safety guardrails, and Claude Mythos 5, the same underlying model with those guardrails removed for a restricted group of trusted partners. This article focuses on Mythos 5, the version Anthropic describes as having "the most advanced cybersecurity capabilities in the world."

In this article, we'll look at what Claude Mythos 5 is, what it can do in software engineering, life sciences, and scientific research, its performance on benchmarks, and how to access it. You can also check out our analysis of Claude Opus 4.8 to see where Mythos 5 fits within the Anthropic lineup. If you're new to Claude and want to get started quickly, we recommend our guide on how to use Claude.

Stay up to date with the latest AI developments. Subscribe to *The Median*, our free Friday newsletter that breaks down the week's news. Get caught up in just a few minutes.


## What is Claude Mythos 5?

Claude Mythos 5 is Anthropic's most capable model, positioned above the Opus class in what Anthropic calls the Mythos tier. The first model in this tier, Claude Mythos Preview, launched in April 2026 as part of Project Glasswing, a collaboration with the U.S. government focused on cybersecurity. Mythos 5 is the second version in this tier and a direct update to Mythos Preview.

Mythos 5 and Fable 5 share the same architecture. The difference lies in the guardrails: Fable 5 includes classifiers that redirect sensitive requests in cybersecurity and biology to Claude Opus 4.8. Mythos 5 removes these classifiers for specific domains for partners vetted through the trusted access program. Anthropic notes that the naming difference reflects the guardrails, not the capabilities.

The headline number: Mythos 5 scores 80.3% on SWE-bench Pro, compared to 77.8% for Mythos Preview and 69.2% for Opus 4.8. On Humanity's Last Exam with tools, it reaches 64.5%, ahead of Opus 4.8's 57.9% and GPT-5.5's 52.2%. These are not marginal gains over the Opus class.

## Overview of Claude models

## What's new with Claude Mythos 5?

Mythos 5 represents an improvement over Mythos Preview across all major capabilities tested by Anthropic. The gains are particularly visible in long-running autonomous work, especially for scientific reasoning and vision tasks. Here's what that looks like in practice.

### Autonomous, secure software engineering at scale

Mythos 5 can work autonomously on large codebases for longer than any other Claude model. Stripe reported that the model compressed months of engineering work into a few days, carrying out a migration across a 50-million-line Ruby codebase in a single day. On FrontierCode (Diamond), it achieves the best scores among frontier models, even at medium effort.

For security missions, Mythos 5 extends the capabilities that made Mythos Preview valuable to Project Glasswing partners. These partners used Mythos Preview to identify more than 10,000 major and critical vulnerabilities in production systems.

### Drug design and protein engineering

Anthropic's internal protein design team used Mythos 5 to accelerate drug design by a factor of roughly ten. In a controlled comparison, Mythos 5 matched or surpassed experienced human operators on 14 protein targets across the entire pipeline:

- selection of binding sites
- choice of tools
- recovery after failure

Nine produced promising drug candidates currently under investigation.

### Generation of novel scientific hypotheses

Mythos 5 is the first Anthropic model to regularly produce original scientific hypotheses, rather than mere syntheses of the literature. In blind comparisons, Anthropic scientists preferred its hypotheses in molecular biology in about 80% of cases, and several were advanced to experimental evaluation. One hypothesis concerning a novel protein mechanism in *E. coli* was independently corroborated by a lab working on the same topic.

### Autonomous genomics research

Mythos 5 conducted novel genomics research over more than a week of largely autonomous work, assembling single-cell data for millions of cells across 138 animal species and training a custom ML model to identify equivalent cell types in distantly related organisms. The trained model outperformed a model recently published in Science while being 100 times smaller.

### Vision and long-context performance

Mythos 5 achieves 93.2% on CharXiv Reasoning with tools and can extract precise figures from detailed scientific charts or reconstruct a web application from simple screenshots. On long-context tasks, adding file-based memory tripled its performance gain compared to an identical setup with Opus 4.8, and it reached the final act of Slay the Spire three times as often.

## Claude Mythos 5 benchmarks

Mythos 5 leads or ties on nearly every benchmark tested by Anthropic, with gains over Opus 4.8 that are consistent across categories rather than concentrated in a single domain. The comparison table pits it against Claude Mythos Preview, Claude Opus 4.8, GPT 5.5, and Gemini 3.1 Pro.

| Category | Benchmark | Claude Mythos 5 / Fable 5 | Claude Mythos Preview | Claude Opus 4.8 | GPT 5.5 | Gemini 3.1 Pro | 
|---|---|---|---|---|---|---|
| Agentic coding | SWE-Bench Pro | 80.3% | 77.8% | 69.2% | 58.6% | 54.2% | 
| Agentic coding | FrontierCode (Diamond) | 29.3% (xhigh) | — | 13.4% (xhigh) | 5.7% (xhigh) | — | 
| Knowledge work | GDPval-AA | 1932 | — | 1890 | 1769 | 1314 | 
| Knowledge work vision | GDP.pdf | 29.8% (no tools) | — | 22.5% (no tools) | 24.9% (no tools) | 16.7% (no tools) | 
| Spatial reasoning | Blueprint-Bench 2 | 38.6% | — | 14.5% | 36.2% | 26.5% | 
| Tool use | AutomationBench | 17.4% | — | 15.5% | 12.9% | 9.6% | 
| Computer use | OSWorld-Verified | 85.0% | 85.4% | 83.4% | 78.7% | 76.2% | 
| Legal | Legal Agent Benchmark | 13.3% | — | 10.4% | 2.1% | 0.0% | 
| Multidisciplinary reasoning | Humanity's Last Exam (no tools) | 59.0%* | 56.8% | 49.8% | 41.4% | 44.4% | 
| Multidisciplinary reasoning | Humanity's Last Exam (with tools) | 64.5%* | 64.7% | 57.9% | 52.2% | 51.4% | 
| Biology | BioMysteryBench (hard) | 46.1%* | 29.6% | 40.0% | — | — | 
| Biology | BioMysteryBench (solved by humans) | 83.9%* | 82.6% | 80.4% | — | — | 
| Agentic coding | Terminal-Bench 2.1 | 88.0%* | — | 82.7% | 83.4% (Codex CLI) | 70.7% (Gemini CLI) | 
| Cybersecurity | ExploitBench (Cap%) | 78.0%* | 69.0% | 40.0% | 34.0% | — | 
| Health | HealthBench Professional | 66.0%* | 64.7% | 56.9% | 51.8% | — | 

Anthropic reports shared scores for Mythos 5 and Fable 5, noting that they typically vary by 1 to 3 points. Benchmarks marked with an asterisk (*) show a larger gap, because Fable 5's safety classifiers redirect sensitive queries to Opus 4.8; on these benchmarks, Fable 5 comes closer to the performance of the Opus class.

### Agentic coding: SWE-Bench Pro, FrontierCode, and Terminal-Bench 2.1

On SWE-Bench Pro, Mythos 5 reaches 80.3%, versus 77.8% for Mythos Preview, 69.2% for Opus 4.8, 58.6% for GPT 5.5, and 54.2% for Gemini 3.1 Pro. The 11-point gap with Opus 4.8 is significant on a benchmark designed to avoid ground-truth leakage.

For the quality and maintainability of agentic code, measured by FrontierCode (Diamond), the gap is even more pronounced. Mythos 5 scores 29.3% at the xhigh effort level, versus 13.4% for Opus 4.8 and 5.7% for GPT 5.5.

For terminal work, Mythos 5 gives Anthropic the edge back over OpenAI: Mythos 5 scores 88.0%* on Terminal-Bench 2.1, versus 82.7% for Opus 4.8, 83.4% for GPT 5.5 (Codex CLI), and 70.7% for Gemini 3.1 Pro (with Gemini CLI).

### Knowledge work: GDPval-AA and GDPpdf

GDPval-AA evaluates knowledge work on a numerical scale. Mythos 5 scores 1932, versus 1890 for Opus 4.8, 1769 for GPT 5.5, and 1314 for Gemini 3.1 Pro.

The gap widens for knowledge work on PDFs without tool access. Mythos 5 reaches 29.8% on GDPpdf, versus 22.5% for Opus 4.8, 24.9% for GPT 5.5, and 16.7% for Gemini 3.1 Pro.

### Multidisciplinary reasoning: Humanity's Last Exam

Humanity's Last Exam (HLE) tests master's-level reasoning in science, mathematics, and the humanities. Mythos 5 scores 59.0%* without tools and 64.5%* with tools. Mythos Preview reaches 56.8% and 64.7% respectively — practically tied with tools but 2 points behind without tools. The gap with Opus 4.8 is already pronounced (49.8% without, 57.9% with), and even more so with competing flagship models (GPT 5.5: 41.4% and 52.2%, Gemini 3.1 Pro: 44.4% and 51.4%).

Mythos 5's advantage appears most clearly without tools, where it leads Opus 4.8 by more than 9 points. These scores are asterisked, meaning that Fable 5 performs slightly worse because of its safety classifiers.

### Computer use, tools, and spatial reasoning

On OSWorld-Verified, which evaluates the model's ability to perform tasks on a real computer interface, Mythos 5 reaches 85.0%. Mythos Preview edges it slightly at 85.4%, making this the only benchmark where Mythos Preview is ahead. Opus 4.8 follows closely (83.4%), while the competitors fall behind: GPT 5.5 at 78.7% and Gemini 3.1 Pro at 76.2%.

AutomationBench measures the ability to use tools. Mythos 5 reaches 17.4%, versus 15.5% for Opus 4.8, 12.9% for GPT 5.5, and 9.6% for Gemini 3.1 Pro. The low absolute values suggest that tool orchestration remains a challenge for all frontier models.

Spatial reasoning is an area where Mythos 5's lead is the most pronounced. It scores 38.6% on Blueprint-Bench 2, more than double Opus 4.8's 14.5%. GPT 5.5 is closer at 36.2%, and Gemini 3.1 Pro reaches 26.5%.

### Cybersecurity and biology

These are undoubtedly the two areas most highlighted in the release notes, and the results explain why.

ExploitBench measures the fraction of exploits the model can successfully reproduce (Cap%). Mythos 5 reaches 78.0%*, a notable improvement over Mythos Preview (69.0%) and a dramatic leap compared with Opus 4.8 (40.0%) and GPT 5.5 (34.0%).

The 38-point gap with Opus 4.8 is the largest in the comparison table, and it justifies the existence of cyber safeguards for Fable 5. Anthropic's external penetration tests found no universal jailbreak on long agentic tasks, even though the UK AISI made progress toward one during an initial testing window.

BioMysteryBench evaluates biological reasoning at two difficulty levels. On the hard subset, Mythos 5 scores 46.1%*, compared to 29.6% for Mythos Preview and 40.0% for Opus 4.8. On the "human-solved" subset, Mythos 5 reaches 83.9%*, Mythos Preview 82.6%, and Opus 4.8 80.4%. GPT 5.5 and Gemini 3.1 Pro do not show reported scores on these subsets.

As with ExploitBench, Fable 5's scores approach those of Opus 4.8 due to its biology-related safety classifiers.

### Health and Legal

Claude Mythos 5 shows notable strength in two high-stakes professional domains, where precision and quality of reasoning have concrete consequences: medicine and law.

On HealthBench Professional, Mythos 5 scores 66.0%*, slightly above Mythos Preview's 64.7%. Opus 4.8 reaches 56.9% and GPT 5.5 51.8%.

On the Legal Agent Benchmark, Mythos 5 reaches 13.3%, compared to 10.4% for Opus 4.8 and only 2.1% for GPT 5.5. The absolute scores remain low, but the gap with GPT 5.5 and Gemini is marked. Legal reasoning remains a challenge for all models.

## Pricing and Availability of Claude Mythos 5

Claude Mythos 5 is priced at $10 per million input tokens and $50 per million output tokens. That's less than half the price of Claude Mythos Preview ($25/$125), which makes upgrading easier for existing Glasswing partners. Developers can access the model via the Claude API with the identifier `claude-mythos-5`.

Access is currently limited to two groups:

- All users who had access to Claude Mythos Preview via Project Glasswing can upgrade to Mythos 5 with cyber safeguards lifted
- A small group of biomedical researchers who can access Mythos 5 with biology and chemistry safeguards lifted, but cyber safeguards maintained.

Anthropic plans to expand both programs over time.

A broader trusted access program is being considered to allow cybersecurity organizations to apply more systematically, in consultation with the U.S. government. Anthropic has not communicated a timeline for general availability. For most developers, Claude Fable 5 is currently the most practical option, with the same underlying model and access via subscriptions and the standard API.

One operational point to note: Anthropic has instituted a 30-day data retention policy for all traffic on Mythos-class models. The data is not used for training and is deleted after 30 days in almost all cases, but it is retained for security purposes. If you are developing with Mythos 5 on sensitive data, consult Anthropic's support documentation on this policy before deployment.

## Conclusion

Claude Mythos 5 is, to date, the clearest proof that Anthropic intends to deploy cutting-edge AI in high-stakes professional contexts, and the results confirm it.

The gap on SWE-bench Pro (80.3% vs 69.2%), on Terminal-Bench 2.1 (88.0% vs 82.7%), and on ExploitBench (78.0% vs 40.0%) points to a model that handles the hardest tasks more reliably than any available alternative.

The restricted access model is a reasonable approach given the dual-use risks, and the ExploitBench scores argue for keeping the most powerful offensive tools out of reach of the general public. The real question is whether Anthropic will be able to expand the trusted access program fast enough to serve the broader security and biomedical community, before competitors close the gap.

For eligible organizations, upgrading from Mythos Preview is simple and less than half the price.

## FAQ about Claude Mythos 5

### What is the difference between Claude Mythos 5 and Claude Fable 5?

Mythos 5 and Fable 5 share the same architecture but differ in their safety safeguards. Fable 5 redirects, via classifiers, sensitive cybersecurity and biology requests to Claude Opus 4.8, while Mythos 5 lifts these classifiers for approved partners. The difference in name reflects the safeguards, not the capabilities.

### Who can access Claude Mythos 5?

Access is currently limited to two groups: the cybersecurity partners of Project Glasswing, who can use Mythos 5 with cyber safeguards lifted, and a small group of vetted biomedical researchers, who can access it with biology and chemistry safeguards lifted, but cyber safeguards maintained. Anthropic plans to expand both programs, with a broader trusted access program for cybersecurity organizations, in consultation with the U.S. government.

### How does Claude Mythos 5 compare to GPT-5.5 and Gemini 3.1 Pro?

Mythos 5 leads both competitors on all benchmarks tested. The largest gaps are on ExploitBench (78.0% vs 34.0% for GPT-5.5), FrontierCode Diamond (29.3% vs 5.7%), and Humanity's Last Exam with tools (64.5% vs 52.2%). Gemini 3.1 Pro is even further behind on most benchmarks.

### Is Claude Mythos 5 safe to use with sensitive data?

Anthropic has instituted a 30-day data retention policy for all traffic on Mythos-class models. The data is not used for training and is deleted after 30 days in almost all cases, but it is retained for security oversight purposes. Organizations handling sensitive data should consult Anthropic's support documentation on this policy before deploying.

### What does the 30-day data retention policy mean for Mythos 5 users?

Unlike standard API usage, where data is not retained, traffic from the Mythos class is retained for up to 30 days for security monitoring before deletion. This applies to all Mythos 5 API calls and is not used for training. This is an important operational point for any organization deploying the model in production with confidential or regulated data.

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
