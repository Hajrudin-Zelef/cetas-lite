---
id: briefing-ia-2026/06-labs-china-europe/10-cross-prices-sizes
title: "Cross-cutting readings: API prices and open-weight sizes"
domain: labs-china-europe
role: deep-dive
task: actor-profile
actors: ["China", "Cohere", "DeepSeek", "Moonshot", "Sakana", "StepFun"]
dates: ["2026-09"]
keywords: ["open-weight", "agent", "cohere", "context window", "deepseek", "distillation", "fugu", "kimi", "licenses", "parameters", "price war", "pricing"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s06-14"
source_lines: [7947, 8010]
sha256: 80577df51f5b1f76c3b41d70cac72d90172840b8c1fd1f5c64abe14db77b1c94
---

# Cross-cutting readings: API prices and open-weight sizes

<a id="s06-14"></a>
### Cross-cutting readings: prices, sizes, distillation, licenses, capital

Beyond individual trajectories, the period lends itself to several
quantified comparisons that use only the briefing's verified facts and
illuminate the market's structure at mid-September 2026.

**The API price war, in one table.**
The documented prices sketch a 1-to-100 spread between Chinese entry
level and the top of the range:

| Model / service | Input ($/M tokens) | Output ($/M tokens) |
|---|---|---|
| DeepSeek V4-Flash-0731 | 0.14 | 0.27 |
| Step 5 Preview (StepFun) | 1 | 2.70 |
| Fugu Max (Sakana, orchestrator) | 2 | 6 |
| Kimi K3 (Moonshot) | 3 | 15 |
| Fugu Ultra v2.0 (Sakana, orchestrator) | 5 | 30 |

The gap between the two extremes is a factor of ~20 on input ($0.14
versus $3) and over 100 on output ($0.27 versus $30).
But the raw comparison is misleading: Fugu is not a foundation model
but an orchestrator that routes to a pool of models — its price pays
for the routing service, not the raw token.
Likewise, Kimi K3 is priced for its size (2.8T parameters) while V4-Flash
is priced for its sparsity (13B active).
The lesson is therefore not "the cheapest wins": it is that the market
has segmented into three pricing logics — the commodity token
(DeepSeek), the performant mid-range (StepFun), the capacity premium
(Moonshot) — plus an emerging fourth logic, the price of the
orchestration service (Sakana).
For a buyer, the question is no longer "which is the best model", but
"which pricing logic matches my workload".

**The open-weights size race.**
In September 2026, the hierarchy of documented open-weight publications
is established as follows, by total parameters: Kimi K3 (2.8T, open
weights) > Step 5 Preview (600B, weights announced 10/15) > DeepSeek
V4.1 Flash (552B, MIT) > DeepSeek V4-Flash-0731 (284B, MIT) > Command A+
(218B, Apache 2.0).
The jump between first and second is a factor of ~4.7: Moonshot is
literally playing in another category.
Note that the three largest are Chinese, and that the only Western
entry in the ranking — Command A+ — is also the smallest, but the one
betting most on operability (near-lossless W4A4 on 1× B200 or 2× H100,
native citations, 48 languages).
Two philosophies oppose each other: the show of force through size
(Moonshot) versus adoption through deployment ease (Cohere).
Both can win different segments: research and prestige on one side,
regulated enterprise on the other.

**The million tokens as the Chinese standard.**
Four models of the period advertise a one-million-token context window:
DeepSeek V4-Pro-0813, DeepSeek V4.1 Flash, Kimi K3, and Step 5 Preview —
all Chinese.
Against them, Command A+ makes do with 128K.
The contrast is a positioning choice: Chinese labs have made "long
context" a commodity, an entry standard into the big leagues; Cohere
assumes a shorter context in service of targeted enterprise use cases.
For agent developers — who consume context at high speed — the Chinese
million tokens are a massive commercial argument, even if the quality
of long-context exploitation remains a variable the briefing does not
measure.

