---
id: briefing-ia-2026/06-labs-china-europe/01-deepseek
title: "DeepSeek: V4-Flash, V4-Pro then V4.1 Flash"
domain: labs-china-europe
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "Apple", "China", "Cohere", "DeepSeek", "Hugging Face", "Inflection AI", "Microsoft", "Mistral", "Moonshot", "Nvidia", "Sakana", "Stability AI", "StepFun", "Z.ai"]
dates: ["2026-09", "2026-09-08"]
keywords: ["deepseek", "accelerator", "advisory", "cohere", "compute", "context window", "distillation", "glm", "inference", "kimi", "license", "licenses"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s06"
source_lines: [7048, 7204]
sha256: 23d121d630807ee241b0a70c553f2f2eab2635746b358ac0e9be85ead21f1453
---

# DeepSeek: V4-Flash, V4-Pro then V4.1 Flash

<a id="s06"></a>
## 6. Chinese, European and other labs in depth

This section covers the period from summer 2026 to mid-September 2026,
with a few earlier milestones mentioned only where they shed light on an
actor's trajectory.
It brings together three families of dynamics that must be read together.
First, the Chinese labs — DeepSeek, Alibaba/Qwen, Moonshot/Kimi, Z.ai/GLM,
StepFun — which dominate the pace of open-weight releases, compress API
prices, and now draw regulators' attention to the question of distillation.
Second, the European and transatlantic players — Mistral, Cohere and its
tie-up with Aleph Alpha — who are betting on sovereignty, European compute,
and record fundraising rounds.
Finally, singular trajectories: Sakana and its orchestrators, Apple and
its Siri pivot, Stability AI and licensed music, Inflection AI and Pi's
return, Ant and its diffusion model.

The through-line of the period, on the Chinese side, is twofold.
On one hand, unprecedented aggressiveness on opening up weights — MIT,
Apache 2.0 licenses, massive publications — which is redrawing the map of
the models available to developers.
On the other, a rising controversy over distillation: the U.S. advisory
AA26-251A of September 8, 2026 accuses several of these labs of industrial-
scale distillation, and Anthropic names Zhipu specifically.
This dossier does not adjudicate these accusations; it documents them,
because they have become a structuring factor of the market.
The European through-line is sovereignty: Mistral secures Vera Rubin
compute in Europe with Microsoft, Cohere merges with Aleph Alpha to build
what they present as the first transatlantic sovereign AI stack,
explicitly sidestepping the U.S. CLOUD Act.

Reading rule: all facts below are drawn from the briefing's verified core;
approximate amounts are flagged as such, out-of-period items are marked
"out of period", and unconfirmed integrations are given in the conditional.

<a id="s06-2"></a>
### DeepSeek: V4-Flash, V4-Pro then V4.1 Flash

DeepSeek kept up a brisk publication pace all through summer 2026,
with three milestones arriving in barely two months.
In July, the lab published V4-Flash-0731; in August, V4-Pro-0813 reached
general availability (GA); on September 10, V4.1 Flash took over and
replaced V4 Pro on the API starting September 14.
This cadence — one major model per month — is a signal in itself: the
Hangzhou lab, already known for R1 and V3, is industrializing its R&D
and shortening its time-to-market cycles.
For developers, that means frequent API migrations; for competitors,
constant pressure on the price-performance ratio.

V4-Flash-0731, published in July, is a 284-billion-parameter MoE model,
with only 13 billion active per token.
That sparsity — under 5% of parameters engaged at each step — is
DeepSeek's economic signature: lots of stored capacity, little compute
spent, hence very low inference cost.
The model is released under the MIT license, which allows commercial
use, modification, and redistribution royalty-free.
On the API side, the listed prices are $0.14 per million input tokens
and $0.27 per million output tokens.
At those price levels, DeepSeek is not chasing unit margin: it is chasing
volume, developer adoption, and the anchoring of its ecosystem in
production pipelines.
It is a strategy of deliberate price aggressiveness, which forces the
entire market — Western players included — to justify their own prices.

V4-Pro-0813, which reached GA in August, changes scale: 1,600 billion
parameters total (1.6T), 49 billion active, and a one-million-token
context window.
The jump is considerable: roughly five to six times the total parameters
of V4-Flash, and nearly four times the active parameters.
The million-token context puts the model in the "long context" category,
capable of ingesting entire codebases or long document corpora in a
single pass.
The GA milestone is a commercial one: the model leaves experimental
status and becomes a stable, billable, supported offering — and thus
integrable into products.
DeepSeek demonstrates here that it can operate at very large scale, not
just publish research weights.

V4.1 Flash, announced September 10, is the third beat of the sequence.
It is a 552-billion-parameter MoE, with 8 to 16 billion active parameters
depending on the mode, and still a one-million-token context window.
The weights are published under the MIT license, continuing the lab's
openness strategy.
The architecture is described as "Causal Encoder-Decoder", a formulation
that signals a hybrid design between encoder-decoder paradigms and the
dominant causal decoders.
One major operational fact: starting September 14, V4.1 Flash replaces
V4 Pro on the DeepSeek API.
In other words, the "Pro" model, which had reached GA just a month
earlier, is already superseded by a more economical "Flash" version.
The implicit message is that efficiency beats raw size: 552 billion
well-exploited parameters replace 1.6 trillion parameters that are more
expensive to serve.

The comparison of the three models deserves to be laid out explicitly,
because it sums up DeepSeek's doctrine for summer 2026:

| Model | Total parameters | Active parameters | Context | License | API price (input/output) |
|---|---|---|---|---|---|
| V4-Flash-0731 (Jul.) | 284B | 13B | — | MIT | $0.14 / $0.27 per M |
| V4-Pro-0813 (Aug., GA) | 1.6T | 49B | 1M | — | — |
| V4.1 Flash (09/10) | 552B | 8B/16B | 1M | MIT | replaces V4 Pro on the API (09/14) |

It shows a bell-shaped trajectory on size — scaling up with V4-Pro, then
coming back down toward efficiency with V4.1 Flash — and one constant:
the million-token context window, now the lab's standard, along with
the MIT license on the open versions.
For a technical decision-maker, the lesson is that "Pro" no longer
necessarily means "bigger": the lineup hierarchy is being recomposed
around inference cost.

Another notable fact of the period: "abliterated" forks of DeepSeek V4.1
— versions stripped of their censorship safeguards — are circulating on
Hugging Face.
The phenomenon is not new in the open-weight ecosystem, but it takes on
particular scale here given the base model's popularity.
The MIT license permits these derivations; the Hugging Face platform
hosts them; the community downloads them.
This once again raises the question of the limits of a lab's control
over the uses of its weights once published — a question all the more
sensitive given that DeepSeek is a Chinese actor subject to censorship
obligations on its own services.
The existence of these forks says nothing about the lab's official
position; it says something about the ecosystem: weight openness is
irreversible, and with it the plurality of uses, including those the
original publisher neither wanted nor foresaw.

On September 8, 2026, DeepSeek is cited in the U.S. advisory AA26-251A,
which concerns industrial-scale distillation.
Two things are worth retaining from this citation.
First, the advisory documents industrial-scale distillation practices
involving the lab.
Second, it calls "misleading" the $5.6 million training cost long touted
for R1 and V3.
This second point matters: the $5.6M figure had fed an entire narrative —
that of a "cheap" training run that supposedly reshuffled the economics
of models — and its undermining weakens that narrative.
That does not mean DeepSeek is not cost-efficient; it means that
transparency on real training costs remains an industry blind spot, and
that marketing figures must be read with caution.
For the market, advisory AA26-251A marks a turning point: distillation —
the technique of training a model on another's outputs — moves from the
status of tolerated gray practice to that of an intellectual-property
issue documented by the authorities.
DeepSeek, on the front line of this controversy, sees its image as an
efficient "outsider" complicated by a geopolitical dimension.

In summary, DeepSeek moves through summer 2026 as a market accelerator:
slashed prices, open weights under MIT, short cycles, long contexts.
But the end of the period — the September 8 advisory, the lightning
replacement of V4 Pro by V4.1 Flash — also shows the fragilities of this
position: dependence on the cost narrative, growing regulatory exposure,
and internal competition between its own product lines.
The lab remains unavoidable for anyone following open weights; it is no
longer just a supplier of cheap models, it is an actor whose every
release moves the sector's prices and licensing standards.

