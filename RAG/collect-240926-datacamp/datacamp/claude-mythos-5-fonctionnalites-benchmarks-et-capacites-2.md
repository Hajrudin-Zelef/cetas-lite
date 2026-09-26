---
id: collect-240926-datacamp/datacamp/claude-mythos-5-fonctionnalites-benchmarks-et-capacites-2
title: "claude-mythos-5-fonctionnalites-benchmarks-et-capacites"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "OpenAI"]
dates: []
keywords: ["benchmark", "benchmarks", "claude", "agent", "agentic", "cyber", "cybersecurity", "fable 5", "gemini", "jailbreak", "mythos 5", "opus 4"]
source: docs/RAG/clean_en/datacamp/claude-mythos-5-fonctionnalites-benchmarks-et-capacites.md
source_anchor: ""
source_lines: [86, 170]
sha256: fc5cadb88d7a3f6332e31516b797bf42774f45c5327bb0de87e04a382d333c84
---

# claude-mythos-5-fonctionnalites-benchmarks-et-capacites

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

