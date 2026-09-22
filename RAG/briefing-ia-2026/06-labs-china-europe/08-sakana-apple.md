---
id: briefing-ia-2026/06-labs-china-europe/08-sakana-apple
title: "Sakana Fugu and Apple iOS 27 / Siri AI"
domain: labs-china-europe
role: deep-dive
task: model-release
actors: ["Anthropic", "Apple", "China", "Cohere", "DeepSeek", "Google", "Mistral", "Moonshot", "OpenAI", "Sakana"]
dates: ["2026-09-11", "2026-09-14"]
keywords: ["fugu", "sakana", "benchmarks", "chatgpt", "claude", "cohere", "compute", "deepseek", "distribution", "gemini", "kimi", "memory"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s06-9"
source_lines: [7754, 7842]
sha256: 172a878d9c9088e6a35929e11bdbf19e68e7fa3fee5c9c59fd3b0dbdb6746331
---

# Sakana Fugu and Apple iOS 27 / Siri AI

<a id="s06-9"></a>
### Sakana: Fugu, the orchestrator rather than the foundation model

Sakana AI, the Japanese lab founded notably by researchers from Google,
occupies a position apart in the 2026 landscape: rather than competing
head-on on foundation models, it bets on orchestration.
On September 11, 2026, Sakana published Fugu Max, priced at $2 per
million input tokens and $6 for output, then Fugu Ultra v2.0, at $5 and
$30 respectively.
The originality lies in the product's nature: Fugu is not a foundation
model, but a learned orchestrator that routes requests to a pool of
models.
In other words, instead of training a single giant, Sakana trains a
system that chooses, for each task, the most suitable model among those
available — including, potentially, third-party models.
This is a strong thesis on the future of AI: value would shift from the
"raw material" (the base model) to routing intelligence (routing, task
decomposition, response synthesis).
The prices — $2/$6 for Max, $5/$30 for Ultra v2.0 — read in this logic:
you pay for the orchestration service, not just the raw token, and the
price positioning is intermediate, between aggressive Chinese models and
American flagships.

This approach deserves comparison with the period's dominant strategies.
Where DeepSeek compresses token cost via MoE sparsity, where Moonshot
stacks parameters in open weights, where Mistral and Cohere play
sovereignty, Sakana bets on heterogeneity: no single model is optimal
for everything, so the winning system is the one that knows how to
choose.
It is also a pragmatic answer to the compute constraint: without the
means to train a 2.8T model like Kimi K3, a lab can nevertheless create
differentiating value on top of existing models.
The mirror risk is dependence: an orchestrator is only as good as the
quality and availability of the models in its pool, and the relevance
of its routing decisions — two variables the briefing does not document
with verified benchmarks.
For enterprise buyers, the promise is seductive: a single entry point
that optimizes the cost/quality trade-off task by task, without locking
the client into a single model vendor.
It is, in short, the anti-silo: where each lab sells its lineup, Sakana
sells the arbitrage between lineups.
If the thesis verifies at scale, 2026 could be remembered as the year
orchestration became a product category in its own right — on par with
the models themselves.

<a id="s06-10"></a>
### Apple: iOS 27 and the Siri AI beta

Apple enters the briefing on September 14, 2026 with iOS 27 and the
beta launch of Siri AI — the assistant rethought around Apple
Intelligence.
Access terms are restrictive: waitlist beta, reserved for iPhone 15 Pro
and later models.
This is the classic Apple method: progressive deployment, experience-
quality control, and hardware-based segmentation — only devices with
sufficient chips and memory get the advanced features.
The briefing documents more than 20 new Apple Intelligence features in
this version, without detailing them: the size of the batch suggests a
major update, not a simple polish.
The differentiating fact highlighted is cross-app personal context:
Siri AI can draw on the user's information across their apps — messages,
calendar, files — for personalized responses.
This is Apple's "on-device and private" bet applied to the assistant:
deep personalization without massive sending of personal data to the
cloud, or at least that is the narrative.
For competitors, it is a direct challenge: neither web chatbots nor
cloud assistants have this depth of integration into the operating
system and local data.

On third-party model integrations, the briefing imposes caution: the
Claude, GPT, and Gemini integrations are in the conditional — only the
ChatGPT integration is documented.
In other words, as of the verified facts' date, Apple has publicly
committed only to ChatGPT as an external model partner; the rest is
speculation or unconfirmed discussions.
This contractual caution is consistent with Apple's strategy: keep
control of the experience, integrate third parties only as optional
complements, and avoid any visible dependence on a competitor.
For the labs concerned, being "the model behind Siri" would be a
colossal distribution trophy — hundreds of millions of iPhones — which
explains the intensity of the speculation.
But as things stand, a single name is confirmed: ChatGPT.
The waitlist-beta deployment also means that real usage feedback —
response quality, latency, cross-app context relevance — will only be
measurable after wider rollout.
Apple is here playing at its own pace: where labs publish models every
month, the Cupertino firm moves by annual system versions, with an
integration standard that slows the tempo but locks in the experience.

