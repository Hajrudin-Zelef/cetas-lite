---
id: collect-240926-misc/misc/mimo-v2-6-flash-a-pricing-benchmarks-speed-command-code
title: "MiMo V2.6 Flash"
domain: commandcode
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Meta"]
dates: []
keywords: ["agent", "agentic", "claude", "cost", "deepseek", "multimodal", "muse", "muse spark"]
source: docs/RAG/clean_en/misc/mimo-v2-6-flash-a-pricing-benchmarks-speed-command-code.md
source_anchor: ""
source_lines: [1, 60]
sha256: 2b9d6088576539ca05e0d855eee81c7512bf169c4c26805378687e4557e6693d
---

# MiMo V2.6 Flash

<!-- source: https://commandcode.ai/models/mimo-v2-6-flash -->

# MiMo V2.6 Flash

efficient multimodal agentic coding with 1M context.

## vs. the lineup

MiMo V2.6 Flash beside its stablemates and nearest rivals. The â marks the best value in each column across every row shown.

| Model | Intelligence | Coding | Speed | Input $/M | Output $/M | Blended $/M | Context | 
|---|---|---|---|---|---|---|---|
| Muse Spark 1.3 Contributor | 48.1â | 75.8â | 247.5â | $0.10â | $0.20â | $0.13â | 1.05Mâ | 
| MiMo V2.6 Pro | 46.3â | ââ | 54.5â | $0.43â | $0.87â | $0.54â | 1.05Mâ | 
| MiMo V2.5 Pro | 26â | 60.2â | 50.4â | $0.43â | $0.87â | $0.54â | 1Mâ | 
| MiMo V2.5 | 25.2â | 56.8â | 39.5â | $0.14â | $0.28â | $0.18â | 1Mâ | 
| MiMo V2.6 Flash â | ââ | ââ | ââ | $0.14â | $0.28â | $0.18â | 1.05Mâ | 
| MiMo V2.6 Pro UltraSpeed | ââ | ââ | ââ | $4.35â | $8.70â | $5.44â | 1.05Mâ | 
| DeepSeek V4 Pro (latest)pinned | 36â | 68.8â | 78.6â | $0.66â | $1.98â | $0.99â | 1Mâ | 
| DeepSeek V4 Flash (latest)pinned | 34â | 69.1â | ââ | $0.15â | $0.60â | $0.26â | 1Mâ | 

## usage calculator

How far a month of credits goes on MiMo V2.6 Flash.

**$0.0003**Â· in $0.14 Â· out $0.28 Â· cache $0.0028 per M

**~6.6K**quick fixes

**~1.3K**bug fixes

**~220**feature PRs

## what real work costs

Real coding tasks priced end to end on MiMo V2.6 Flash, from a quick lookup to a full-repo agent run.

| Task | Tokens in Â· out | MiMo V2.6 Flash | Muse Spark 1.3 Contributor | Claude Haiku 4.5 | 
|---|---|---|---|---|
| Quick lookup / one-liner | 8K Â· 1K | $0.0004 | $0.0003 | $0.0065 | 
| Review a 500-line PR | 60K Â· 4K | $0.0038 | $0.0027 | $0.04 | 
| Fix a bug (agent loop) | 180K Â· 12K | $0.01 | $0.0072 | $0.12 | 
| Refactor a module | 320K Â· 20K | $0.02 | $0.01 | $0.21 | 
| Full-repo agent run | 900K Â· 45K | $0.04 | $0.03 | $0.49 | 

## frequently asked

## How much does MiMo V2.6 Flash cost?

`$0.14`/M input and `$0.28`/M output, cache reads `$0.0028`/M. In an agent loop most input is cache-read, so the effective input rate is about `$0.04`/M.
## Which plan do I need?

**Go**and above.

## How do I switch to it?

`cmd --model xiaomi/mimo-v2.6-flash`, or type `/model` in a session and pick it. You can switch mid-session without losing context.
## Ship code that matches your taste

Command Code is the AI coding agent that continuously learns your taste. Start for $1.
