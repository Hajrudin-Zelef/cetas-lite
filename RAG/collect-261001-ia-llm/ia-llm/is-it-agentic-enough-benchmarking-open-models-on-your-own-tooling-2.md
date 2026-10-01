---
id: collect-261001-ia-llm/ia-llm/is-it-agentic-enough-benchmarking-open-models-on-your-own-tooling-2
title: "Task: classify the sentiment of \"I absolutely loved the movie, it was fantastic!\""
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "MiniMax", "Moonshot", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "agents", "cost", "distribution", "glm", "kimi", "lean", "memory", "quantization", "tool use", "training"]
source: docs/RAG/collect-261001-ia-llm/is-it-agentic-enough-benchmarking-open-models-on-your-own-tooling.md
source_anchor: ""
source_lines: [129, 241]
sha256: 58fd008a7100f830dd4b5fd755c7a55a026fcbfd0c601780330ae44c88453209
---

# Task: classify the sentiment of "I absolutely loved the movie, it was fantastic!"

  *A run rendered in the Hub's agent-traces viewer: MiniMax-M2.7 on the answer-question task.*

  **Open this trace on the Hub ↗**

Before the results, a quick recap of the setup. Each run varies four things: the **model** driving the agent,
the **`transformers` revision** it runs against, the **task**, and the **tier** (`bare` / `clone` / `skill`).
As discussed, we look at different metrics for the two different model categories.

Since a large open model will usually get to the correct result, what you're really measuring is the effort it took to do so. Did it take ten turns or one? Did it follow an API path you deprecated because it trusted obsolete documentation? Did it hit an error you hadn't foreseen?

The natural experiment is to fix one strong model and vary the tool's
revisions: the successive git versions of `transformers` we test against, from released tags like
`v5.8.0` and `v5.9.0` to the specific commit that introduces the CLI and Skill. We want to watch whether the load
it puts on the agent goes up or down. We used the harness on `transformers` to check
whether adding a dedicated CLI and Skill actually lightened the agents' work.

For the three large models we used in our tests, the average time spent on all tasks indicates that the Skill commit results in less time spent working on the tasks:

  

  *Median time per revision, by tier: the skill commit (green dot) is the fastest.*

On the other hand, in the experiments in which we cloned the repository, we can see a significant increase in token consumption due to the commit that introduced the CLI and examples, as we'll see in a moment.

  

  *Median new tokens per revision, by tier: the clone variant jumps once the CLI lands in the repo.*

Reading the clone-variant traces explains why. The commit adds a command, but it also ships the
CLI's implementation and a set of `cli/agentic/*.py` usage examples into the repository directly. 

On the `clone` variant the agent has a full transformers checkout in front of it, and roughly a third of the runs go read the new
surface (the `/cli/` tree and the example scripts) to learn the interface before calling it. This raises the
median input from ~4k to ~6.4k tokens. 

The two charts are then two sides of one tradeoff: the commit buys the large models less time (they reach for the CLI instead of debugging Python) at the cost of more tokens (they read the code that taught them the CLI). A tradeoff worth knowing about before merging PRs.

One caveat works in the CLI's favor, though, which isn't benchmarked yet: the cost of reading it is amortized with successive runs. Our setup is built for one-off experiments. Each run is a fresh agent that rediscovers the CLI from scratch, so it pays the discovery cost every time. In real usage an agent learns the interface once and then solves task after task within the same session, amortizing that cost across many requests. The token bump we measure here is closer to a worst case than to what a user would see day to day.

Open models give us fine-grained control over the variables that matter most here: size, configuration, quantization,
provider, training, and anything that would differ from one model to the next. They're also where a good tool surface
matters most: a small model asked to "use `transformers` to do X" on a `bare` environment can
guess an API that changed some releases ago, may do unnecessary tool calls, and can
get the wrong answer.

So here the experiment is the opposite of the above: hold the revision and sweep the model. This helps see which models actually take care of the task, not just by token count and time, but down to which ones can't reliably handle the tool calls. Our intuition is that the smaller the model, the harder both tool use and the task get; we ran the harness across a range of model sizes to test exactly that:

  

  *Match % across models, by tier: the skill tier lifts the larger models but drops the smaller ones.*

which also seems to be correlated with the number of tokens ingested

  

  *Median new tokens across models, by tier.*

A note on fair comparison: naively averaging across tasks is misleading when
coverage is uneven (a model that only finished the quick tasks looks fast). The
report has a **"shared tasks only"** toggle (across models and/or revisions) so
you compare like-for-like, and a **Coverage** heatmap so you can see exactly which
task × revision × model cells actually ran.


Two things come together here: how to look past whether the agent succeeded to what it did and how it did it; as well as the first results we pulled out of the harness.

Match %, tokens, and time tell you the cost of a run but don't tell you much about what happened under the hood.

This is why we've introduced the concept of markers. A marker is a named pattern the profile (the small per-tool plugin that teaches the harness how to build and drive a given library) matches against a run.

It is a one-line label for a behavior you care about, checked against the shell commands the agent ran, the code it wrote, the files it read, or its final answer. A run can fire several markers or none; the report shows how often each one fired, per model and per revision.

For `transformers` we declare a handful but we'll only look at the two most relevant ones:

- **`cli`** : the agent invoked the`transformers` command-line tool (e.g.`transformers classify …` ) instead of writing Python.
- **`pipeline`** : it reached for the high-level`pipeline(...)` Python API.

These are what we watch to see whether a change actually shifted the agent's behavior. Interestingly here, the larger the model, the more it leverages the new context instead of using its memory; therefore leveraging the newly introduced CLI.

  

  *CLI adoption by tier across models: only the skill tier reaches for it, and more so as models grow.*

CLI adoption is new: the CLI lands in a single commit, isn't in any model's training data, and is only lightly documented. The effect is clear: it's the Skill variant, the one that ships the CLI's documentation, that actually reaches for it, at 55.3%.

Comparing the commit across model sizes, the CLI + Skill helps the bigger models: on the `skill` tier, Kimi and the other large agents reach for the CLI and finish in fewer turns. (On `clone` they spend *more* input tokens first, reading the new CLI code, as we saw above, so the win shows up in time and turns, not raw tokens.)

  

  *Kimi-K2.6, GLM-5.1, and MiniMax-M2.7 across revisions*

But in some smaller-model settings, it appears to hurt performance. One plausible explanation is that small
models lean on memorized API patterns, reproducing `pipeline(...)` snippets
they've seen in their training data. The new concepts are then a larger
surface for them to get wrong. You can watch this directly on the harness: lower
match %, more retries, the `cli` marker barely firing. It is particularly striking on the Qwen3-4B model:
the Skill barely changes its match rate yet its cost distribution is significantly affected. 

Almost all of that comes from the `clone` tier. The checkout now contains
the CLI's implementation and `cli/agentic/*.py` examples, and the 4B agent reads them in bulk: its median new
tokens jump from ~2.4k to ~23k, with time and output skyrocketing as well, for no gain in
accuracy.

  

  *Qwen3-4B across revisions. The CLI + Skill commit fans the cost distribution wide open, on the `clone` tier the agent reads the newly-shipped CLI source in bulk (~10× the new tokens), for no gain in match %. (`repeat tokens` stays flat: this setup uses no prompt caching.)*

Sometimes, though, the Skill breaks correctness outright. Reading the traces shows how, for example for Qwen3-14B: 
adding the Skill drops its overall match rate from 67% (bare) to 43%, and on the simplest tasks the collapse is very
visible: `classify-sentiment` goes from 100% on the `clone` variant to **0%** with the Skill. 

  

