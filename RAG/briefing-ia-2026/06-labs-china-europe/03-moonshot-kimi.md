---
id: briefing-ia-2026/06-labs-china-europe/03-moonshot-kimi
title: "Moonshot / Kimi: K3, open weights and the Hong Kong road"
domain: labs-china-europe
role: deep-dive
task: model-release
actors: ["AWS", "Alibaba", "China", "Cohere", "DeepSeek", "Meta", "Moonshot", "StepFun"]
dates: ["2026-07", "2026-09", "2026-09-04", "2026-09-08", "2026-09-18"]
keywords: ["kimi", "open weights", "advisory", "aws", "bedrock", "cohere", "context window", "deepseek", "distillation", "distribution", "ipo", "open source"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s06-4"
source_lines: [7280, 7381]
sha256: cf5111b8da381445fc4c0c93b60c57d8be34c54d61070ee5e94de76e42bdb5a4
---

# Moonshot / Kimi: K3, open weights and the Hong Kong road

<a id="s06-4"></a>
### Moonshot / Kimi: K3, open weights and the road to Hong Kong

Moonshot AI, the Chinese lab behind the Kimi family, signs one of the
densest sequences of the period, at the intersection of technology,
open source, and finance.
In July 2026, it published Kimi K3: a 2,800-billion-parameter MoE (2.8T),
with a one-million-token context window, priced at $3 per million input
tokens and $15 for output.
The scale is massive: 2.8T is nearly five times the 552 billion
parameters of DeepSeek V4.1 Flash, published two months later, and well
above the 600 billion of StepFun's Step 5 Preview.
The pricing, by contrast, sits at the high end of the period's Chinese
models — compare with the $0.14/$0.27 of DeepSeek V4-Flash or the
$1/$2.70 of Step 5 Preview.
Moonshot is therefore not playing the price-slashing card: it positions
K3 as a premium model, billed at the level of its size.
The million-token context window, now quasi-standard among advanced
Chinese labs, confirms that the "long context" battle is now an expected
commodity, not a differentiator.

The truly historic fact follows shortly after: Moonshot then published
Kimi K3's open weights, making it — according to the briefing's verified
facts — the largest model ever published in open weights.
2,800 billion parameters available for download: the order of magnitude
exceeds anything the open-weight ecosystem had ever seen.
The implications are multiple.
First, a show of force: publishing at this scale means asserting a
mastery of training and infrastructure that few labs in the world can
claim.
Second, a community bet: even if few actors can run a 2.8T model locally,
weight availability feeds research, distributed fine-tuning, and
derivative architectures.
Finally, a geopolitical gesture: by opening the largest model ever
published, a Chinese lab claims the symbolic leadership of global
open weights, ground historically held by Meta with Llama.
The comparison with the period's other open weights is telling:
DeepSeek V4.1 Flash (552B, MIT), Step 5 Preview (600B, announced for
October 15), Cohere's Command A+ (218B, Apache 2.0) — all are an order
of magnitude smaller than Kimi K3.
Moonshot thus plays alone in its size category, which makes the
publication all the more remarkable.

On September 18, 2026, Kimi K3 reached general availability (GA) on
AWS Bedrock.
This is a major distribution milestone: Bedrock is Amazon Web Services'
model marketplace, and being listed there means reaching enterprises
already on AWS with integrated billing and compliance.
For a Chinese lab, being distributed on a U.S. giant's cloud
infrastructure is far from trivial: it proves that commercial barriers,
while not disappearing, remain crossable through the cloud-partnership
channel.
It also positions Kimi K3 as a credible off-the-shelf alternative to
U.S. models for CIOs who want to diversify suppliers while staying
within their AWS perimeter.
The GA on Bedrock, combined with the open weights, sketches a two-speed
distribution strategy: the managed API for enterprises, the weights for
the community and sovereign deployments.

On the financial side, Moonshot is targeting a Hong Kong stock market
listing: about $3 billion raised, for a valuation of about $50 billion,
with a filing submitted on September 4, 2026.
The orders of magnitude are considerable: a $50B valuation would place
Moonshot among the world's most highly valued AI startups, and a $3B
raise in Hong Kong would be one of the largest tech deals on that
exchange in recent years.
The choice of Hong Kong, rather than New York, fits the general movement
of Chinese tech companies refocusing on Asian markets, amid Sino-U.S.
regulatory tensions.
The timeline is tight: filing in early September, just as the lab
publishes its weights and reaches GA on Bedrock — all milestones meant
to feed the IPO narrative.
Still, these figures are, at this stage, those of an IPO project: the
filing is submitted, the deal is not done, and market conditions will
decide its final success.

Finally, Moonshot is cited in advisory AA26-251A of September 8, 2026:
the advisory documents more than 23 million exchanges aimed at
distilling Kimi K3.
The volume is far below the 151 million attributed to Alibaba for Qwen
over the May–July 2026 period, but it remains massive in absolute terms.
The timing is troubling: it is just as Moonshot prepares its IPO and
opens its weights that the question of its models' distillation becomes
a subject of U.S. advisories.
For an investor reading the IPO filing, this is a risk factor to price
in: the intellectual property of the models, and the lab's ability to
defend it, become valuation variables.
For the industry, it is confirmation that distillation spares no major
actor — neither the most open nor the most highly valued.

In summary, Moonshot moves through September 2026 on a spectacularly
ascending trajectory — largest open weights ever published, GA on
Bedrock, $50B IPO in preparation — but under growing scrutiny from U.S.
authorities on distillation.
The lab embodies Chinese AI's move upmarket: gone are the days when
Chinese actors were merely low-price followers; Kimi K3 commands a
premium price, flaunts record size, and targets international financial
markets.
The open question is sustainability: turning technical prowess and a
valuation into recurring revenue remains the challenge of every
foundation lab, Chinese and American alike.

