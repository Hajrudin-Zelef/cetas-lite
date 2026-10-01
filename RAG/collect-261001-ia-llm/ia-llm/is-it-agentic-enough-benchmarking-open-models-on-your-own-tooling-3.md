---
id: collect-261001-ia-llm/ia-llm/is-it-agentic-enough-benchmarking-open-models-on-your-own-tooling-3
title: "Task: classify the sentiment of \"I absolutely loved the movie, it was fantastic!\""
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face"]
dates: []
keywords: ["agent", "agentic", "inference"]
source: docs/RAG/collect-261001-ia-llm/is-it-agentic-enough-benchmarking-open-models-on-your-own-tooling.md
source_anchor: ""
source_lines: [242, 294]
sha256: c98adaa560c5578f7b634af8073da810e3b56cd84e44641a0d4363e369de914d
---

# Task: classify the sentiment of "I absolutely loved the movie, it was fantastic!"

  *Qwen3-14B on `classify-sentiment`, by tier: `clone` (blue) holds at 100% across revisions, but the Skill variant (green) collapses to 0% at the CLI + Skill revision.*

Looking at the traces, the model mistakes the CLI for a *tool it can call directly* (as in an agentic-harness
tool, like web-search). The Skill is **not** an executable tool: it's documentation loaded
into the agent's context, and the `transformers` CLI is only ever meant to be run from the shell (via `bash`); so this
will not work.

Qwen3-14B reads the Skill and, in 39 of its 56 Skill runs, either emits a `transformers(command="classify", ...)`
tool call (a tool that was never registered) or, finding nothing like it among its `read`/`bash`/`edit`/`write`
tools, concludes it *can't* run a model and gives up. Either way, rather than fall back to the one-line
`pipeline(...)` that scored 100% on the `clone` checkout, it declares the task impossible.

  

  *Qwen3-14B on classify-sentiment (Skill variant): it reasons that read/bash/edit/write can't run a model, and gives up.*

This is exactly what we built the harness to catch: the same change that speeds the large models
ends up breaking the small ones, which seemed a bit counterintuitive to us at first and something we'd likely have
shipped as-is. The takeaway for maintainers: **agent-facing APIs should be evaluated across model sizes, because a
new affordance can reduce work for strong models while adding ambiguity for smaller ones.** It also hints at a fix:
rather than hand-write a Skill and check it after the fact, you could generate and validate one against the weaker
models up front. 

This is exactly what Upskill does: it turns a strong model's solution into a Skill only when it measurably helps the smaller ones.

The harness is one CLI, `agent-eval`. Install it, run a suite, fan it out across models × revisions on HF Jobs, and publish the
report as a Hugging Face Space. 

**Trusted local use only.** The harness runs a coding agent with bypassed permissions and executes code from
whatever revision you point it at, and traces can contain prompts, output, and local paths. See
SECURITY.md before pointing it at code you didn't write or sharing results.


The full, kept-current setup and usage instructions live in the README.

Checking the final answer tells you whether an agent *can* use your library. It
doesn't tell you what it costs: the turns, tokens, errors, and the path it took to
get there. This harness measures that, across the revisions and models you pick.

On `transformers`, it caught something we'd have shipped on faith: the CLI + Skill
helps the largest open models and hurts the smallest ones. Worth knowing before merging!

It's profile-based, and designed to be adaptable: point it at your own library, define a few tasks and their expected answers, get the same report. Code and tasks are in the repo, traces are on the Hub. Let us know if you use it for your project!

This harness stands entirely on
pi, Mario
Zechner's coding-agent CLI: it drives every open-model run, and only needs an
`HF_TOKEN` to serve a model, which is what made the open-model sweep practical at
all.

Thanks to the model builders and inference providers behind the models we
swept. Across the board they performed well above what the `bare` baseline would
suggest.
