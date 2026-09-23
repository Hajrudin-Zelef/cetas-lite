---
id: collect-mindstudio/mindstudio/open-vs-closed-weight-model-token-share
title: "Open Weights vs Closed Weights: Why Token Usage Is Flipping"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "United States"]
dates: ["2026-08", "2026-09-23"]
keywords: ["open weights", "agent", "benchmarks", "claude", "cost", "cost per token", "deepseek", "inference", "kimi", "nvidia", "open-weight", "perplexity"]
source: docs/RAG/Collect RAG/02_mindstudio/open-vs-closed-weight-model-token-share.md
source_anchor: ""
source_lines: [1, 56]
sha256: 203e14838ca20f8f0080c2febec5cac773326a92ad5b739547f340002413ff92
---

# Open Weights vs Closed Weights: Why Token Usage Is Flipping

## Metadata

- **Source** : https://www.mindstudio.ai/blog/open-vs-closed-weight-model-token-share
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article analyzes a crossover in AI token usage data from Vercel's hosting platform tracked between June and August 2026: open-weight models (freely downloadable, self-hostable, fine-tunable) now process more total tokens than closed-weight models like OpenAI's GPT line and Anthropic's Claude. The closed-model line is sliding down while open-weight usage climbs. DeepSeek alone has passed Anthropic in token share, 25.2% versus 24.5%. Crucially, token volume and revenue tell two different stories: Anthropic and OpenAI still capture the overwhelming majority of dollars spent on AI inference even as their share of raw token traffic shrinks. Anthropic captures roughly 64.6% of model spend versus DeepSeek's 2.8%, despite near-equal token shares.

The article explains the shift through cost and control. Open-weight models are generally far cheaper to run, and businesses can self-host them, fine-tune on private data, and avoid sending sensitive information to a third party. China has been a major source of increasingly capable open-weight systems. Structurally, once a task does not require the absolute best model, there is little incentive to pay premium prices — if a flash-tier open model handles everyday tasks, most users have no reason to reach for the most expensive option.

Revenue stays concentrated because being marginally smarter matters enormously in high-stakes tasks. Venture investor Gavin Baker's framing (cited in the source discussion) holds that frontier tokens from labs like Anthropic and OpenAI may represent only 10–25% of total tokens used yet account for 60–90% of all economic value generated. For high-frequency trading, legal analysis, or advanced coding, a small accuracy edge justifies dramatically higher cost. Anthropic's annualized revenue run rate is reportedly above $65 billion, with OpenAI around $40 billion — more than every open-model provider combined. The pricing gap is stark: Claude's higher-tier models run around $50 per million output tokens, while a DeepSeek flash-tier variant can cost closer to 18 cents per million output tokens.

The article stresses that cost per token is not the right comparison; cost per completed task is. Investor Bindu Reddy's point is that some models need significantly more tokens to reach the same conclusion. In one comparison, Kimi K3 priced at roughly half the per-token cost of GPT 5.1 still cost almost the same per completed task (84 cents vs 96 cents) because it required more tokens. Investor Martin Casado highlights cost, privacy, and product fit as the real decision drivers; privacy pushes large companies toward controllable open-weight models. Examples: Thomson Reuters' Harvey (built on Kimi), Airbnb (Qwen), Perplexity (DeepSeek), and Cursor (Kimi K2.5). Harvey fine-tuned Kimi K3 into a legal-specific model that ranked near the top on contract analysis and corporate law agent benchmarks — something only possible with open weights.

The article outlines risks of closed-platform dependency: data exposure revealing internal workflows and strategy, provider pricing risk, availability/terms risk, and the possibility a provider competes directly using insight gained from a customer's usage. Open weights remove most of this risk since the business owns the weights and fine-tuned data and can move between providers or self-host. MIT economist Christian Catalini's three-tier framing is used: bottom = cheap generalist open-weight commodities carrying most volume but small spend share; middle = open-weight "state of the art specialist" models fine-tuned for enterprise, carrying meaningful spend; top = closed frontier generalist models from OpenAI/Anthropic, a small volume share but the majority of revenue. The article notes Nvidia benefits either way, since open-weight adoption still drives chip demand.

## Key points

- Open-weight models now process more total tokens than closed models on Vercel's platform (June–August 2026 data).
- DeepSeek passed Anthropic in token share: 25.2% vs 24.5%.
- Revenue remains concentrated: Anthropic ~64.6% of model spend vs DeepSeek 2.8%.
- Frontier intelligence commands a huge premium: frontier tokens may be 10–25% of volume but 60–90% of economic value.
- Anthropic run rate reportedly >$65B, OpenAI ~$40B, more than all open-model providers combined.
- Cost per token ≠ cost per task (Kimi K3 at half the per-token price of GPT 5.1 cost ~same per task: 84¢ vs 96¢).
- Major US companies build on Chinese open weights: Harvey (Kimi), Airbnb (Qwen), Perplexity (DeepSeek), Cursor (Kimi K2.5).
- Market splitting into three tiers: commodity open models, open-weight specialists, closed frontier generalists.

## Technical data / figures

| Item | Value |
|---|---|
| Data source | Vercel hosting platform |
| Tracking window | June–August 2026 |
| DeepSeek token share | 25.2% |
| Anthropic token share | 24.5% |
| Anthropic model spend share | ~64.6% |
| DeepSeek model spend share | ~2.8% |
| Frontier tokens as share of volume | 10–25% |
| Frontier value as share of economic value | 60–90% |
| Anthropic annualized run rate | >$65B |
| OpenAI annualized run rate | ~$40B |
| Claude high-tier output price | ~$50 / M tokens |
| DeepSeek flash-tier output price | ~$0.18 / M tokens |
| Kimi K3 cost per task | $0.84 (vs GPT 5.1 $0.96) |

## Why this source matters for the RAG

It quantifies the open-vs-closed token/revenue divergence and the cost-per-task versus cost-per-token distinction, which are foundational to any analysis of model selection, self-hosting, and AI economics. The concrete company examples (Harvey, Airbnb, Perplexity, Cursor) provide real deployment evidence for open-weight viability.

