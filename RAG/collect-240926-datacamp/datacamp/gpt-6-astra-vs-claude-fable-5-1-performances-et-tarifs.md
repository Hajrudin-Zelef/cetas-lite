---
id: collect-240926-datacamp/datacamp/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs
title: "gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Irregular", "Microsoft", "OpenAI", "OpenRouter", "SpaceX"]
dates: ["2026-04-30", "2026-06", "2026-09-01", "2026-09-03"]
keywords: ["astra", "claude", "gpt-6", "agent", "agentic", "agents", "alignment", "bedrock", "benchmark", "benchmarks", "chatgpt", "consumer"]
source: docs/RAG/clean_en/datacamp/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs.md
source_anchor: ""
source_lines: [1, 303]
sha256: 3113e2098a26f6da1f372921a056cd82ddf7bdf6fcf9e6bc88ff9863f698e3e1
---

# gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs

<!-- source: https://www.datacamp.com/fr/blog/gpt-6-astra-vs-claude-fable-5-1 -->

Course

In the space of two days, Anthropic and OpenAI released their new flagship models. They display exactly the same price: $10 per million input tokens and $50 per million output tokens, which, paradoxically, complicates the usual question "which one is cheaper?": identical rates don't produce identical bills.

In this article, I compare GPT-6 Astra and Claude Fable 5.1 on coding and agentic work, reasoning, computer use, security, and their real-world cost in use. For an in-depth detour on each model, see our GPT-6 Astra guide and our Claude Fable 5.1 guide.

## TL;DR

- OpenAI's benchmark table places Astra ahead of Fable 5.1 almost everywhere, while Artificial Analysis, an independent evaluator, places Fable 5.1 on top on its two flagship indices.
- The list prices are identical; the only differences in the pricing grid concern cache reads, where Fable 5.1 is 4 times cheaper, and Astra's long-context surcharge.
- At the measured cost per task, the trend reverses. Artificial Analysis estimates Fable 5.1 at $3.76 per Intelligence Index task versus $1.67 for Astra, because Astra consumes far fewer tokens for its score.
- Favor GPT-6 Astra for computer use, professional deliverables, reasoning in math and science, cyber defense, and a lower cost per task.
- Favor Claude Fable 5.1 for depth of reasoning, for queries beyond 272K tokens where Astra adds a surcharge and Anthropic does not, and for agent loops where the bill is dominated by cache reads rather than by output.

## Want to get started with generative AI?

Learn to work with LLMs in Python directly in your browser

## What is GPT-6 Astra?

GPT-6 Astra is OpenAI's cutting-edge flagship model, successor to GPT-5.6 Sol, designed around agentic execution rather than chat. OpenAI positions it for computer use, professional work, and software engineering, and it's the first OpenAI model to cross the "Critical" cybersecurity threshold in the company's Preparedness Framework. It offers a context window of 1,500,000 tokens, a max output of 128K, and a knowledge cutoff date of April 30, 2026.

Two highlights in the announcement. In Codex, Astra keeps notes from one context window to the next instead of compacting them into a summary, so previous windows remain consultable; and it decides when to ask a clarifying question instead of always guessing or always asking.

For the full list of features, benchmark tables, and access methods, see our GPT-6 Astra guide. Read it alongside our coverage of its predecessor: GPT-5.6 Sol, Terra and Luna.

## What is Claude Fable 5.1?

Claude Fable 5.1 is Anthropic's generally available flagship model for demanding reasoning and long-duration agentic work. It offers a 1M token context window with 128K max output, always-active adaptive thinking, and a knowledge cutoff date of June 2026. Anthropic's documentation indicates higher latency than Claude Opus 5 and Claude Sonnet 5, both of which are cheaper.

Claude Mythos 5.1 is the same model with different guardrails, accessible by invitation via Project Glasswing. The key change for everyone else concerns the cache read price, reduced by 75% to $0.25 per million tokens, while list rates remain unchanged.

Our Claude Fable 5.1 guide details the entirety of the benchmarks, and our Claude Fable 5.1 API tutorial builds a repository-aware developer agent, with a real breakdown of costs by effort.

## GPT-6 Astra vs Claude Fable 5.1: direct comparison

In short: OpenAI's comparison table shows Astra ahead of Fable 5.1 on almost every published line, and Artificial Analysis shows the opposite on its two indices. Which to believe depends on the weight you give to a vendor rating its competitor.

| Feature | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Release date | September 3, 2026 | September 1, 2026 | 
| API model ID | `gpt-6-astra` | `claude-fable-5-1` | 
| Context window | 1.05M tokens | 1M tokens | 
| Maximum output | 128K tokens | 128K tokens | 
| Knowledge cutoff | April 30, 2026 | June 2026 | 
| List price per 1M tokens | $10 input / $50 output | $10 input / $50 output | 
| Cached input read per 1M | $1.00 ($2.00 beyond 272K) | $0.25 | 
| Rates beyond 272K input tokens | 2x input and cache, 1.5x output | No surcharge | 
| FrontierMath Tier 4 (v2) | 97.6% | 87.8% | 
| Humanity's Last Exam, with tools | 57.2% | 65.0% | 
| ScreenSpot-Pro (without tools) | 92.7% | 87.3% (Fable 5, from Mythos) | 
| AutomationBench | 41.4% | 31.4% | 
| ExploitBench | 100% | 70% | 
| AA Intelligence Index (max effort) | 61 | 66 | 
| AA: cost per Intelligence Index task (max) | $1.67 | $3.76 | 
| Strong point | PC use, math, cybersecurity, cost per task | Depth of reasoning, agent loops dominated by cache | 

### Coding and agentic workflows

Astra leads on the coding benchmarks OpenAI published, but the gap is slim and the independent index disagrees. On Terminal-Bench 4.0, which tests software engineering, system configuration, and data analysis in the terminal, OpenAI announces 57.7% for Astra versus 55.8% for Fable 5.1. Anthropic publishes the same 55.8%, so at least that line isn't in dispute.

The gap widens on DeepSWE v1.1, where Astra reaches 74.1% against 67.4%, and almost entirely closes on FrontierCode 1.1 Main, where 53.3% against 50.9% falls within the range already occupied by Claude Fable 5 (53.5%) and Claude Opus 5 (53.4%).

| Benchmark | GPT-6 Astra | Claude Fable 5.1 | Notes | 
|---|---|---|---|
| Terminal-Bench 4.0 | 57.7% | 55.8% | OpenAI's table shows 57.7% while the chart legend indicates 57.9% | 
| DeepSWE v1.1 | 74.1% | 67.4% | Given by OpenAI | 
| FrontierCode 1.1 Main | 53.3% | 50.9% | Equivalent to Fable 5 (53.5%) and Opus 5 (53.4%) | 
| Internal database migration | 63.9% | 57.8% | OpenAI internal eval, not replicated | 
| CursorBench 3.2.0 | Not published | 73.4% | Given by Anthropic; SpaceXAI confirmed 73.4% at max effort | 
| AA Coding Agent Index | 67 in Codex | 70 in Claude Code | Independent, but different harnesses | 

The last row is the one that makes me think. Artificial Analysis places Fable 5.1 in Claude Code at 70, the best score on its Coding Agent Index, and Astra in Codex at 67, roughly on par with Claude Opus 5 and Claude Fable 5. Two caveats:

- Both models ran in different harnesses, Codex vs Claude Code, so part of the 3-point gap comes down to tooling rather than the model. Our Codex vs Claude Code comparison shows their very different behaviors.
- Astra gets there far more economically. Artificial Analysis measured it at one-third the tokens of GPT-5.6 Sol at max effort in Codex, and one-fifth the tokens of Claude Opus 5 at xhigh.

Astra wins the published coding rows by a nose and takes the efficiency argument hands down, while Fable 5.1 holds the only independent coding agent score that beats them both.

### Reasoning and scientific work

This is the cleanest split, and it goes both ways. Astra takes the math and science rows: 97.6% against 87.8% on FrontierMath Tier 4 v2, 96.0% against 93.7% on GPQA Diamond, and 64.6% against 52.6% on Terminal-Bench Science 0.1, the agentic research benchmark around which Anthropic had articulated its own launch.

Fable 5.1 wins Humanity's Last Exam with tools, 65.0% against 57.2%, and the gap is clear. Artificial Analysis independently confirms the trend, noting Fable 5.1 at 66 on its Intelligence Index against 61 for Astra, the highest score ever measured.

One caveat on that 66: Artificial Analysis ran Fable 5.1 with Anthropic's default server-side fallback, which redirects requests flagged for safety to Claude Opus 4.8 or Claude Opus 5; this fallback produced about 4% of the index's output tokens. The score reflects Fable 5.1 as it will actually be called, not the bare model in isolation.

If your work is physics or master's-level math, Astra. If it's broad, difficult reasoning in the style of Humanity's Last Exam, Fable 5.1.

### PC use and professional deliverables

Astra dominates this axis, especially because Anthropic hasn't published comparable figures. OpenAI reports 72.6% for Astra on the offline OSWorld 2.0 set against 70.2% for Claude Opus 5, and 92.7% on ScreenSpot-Pro against 87.3% for Claude Fable 5. The figures specific to Fable 5.1 (77.9% partial, 41.7% strict) come from a different version of the tasks and a different scoring rubric.

The artifact figures are less ambiguous: 95.9% on BenchCAD against 84.3%, and 41.4% against 31.4% on AutomationBench.

For agents that click through real software and produce slides or CAD, Astra is the better choice.

### Safety, cybersecurity, and alignment

Astra's alignment figures are the most underrated aspect of its launch. On OpenAI's internal benchmark for safe PC use, where a lower score is better, it reaches 2.4% against 9.5% for Fable 5.1.

In cybersecurity, it changes divisions: 100% on ExploitBench, and 86 FrontierCyber challenges solved out of 226 against 34 for GPT-5.6 Sol according to Irregular's independent tests. Anthropic unrestrained Fable 5.1 just enough to find vulnerabilities but not to exploit them.

Two nuances to keep in mind:

- OpenAI indicates that Astra's written reasoning is *harder* to monitor than Sol's, because it solves in fewer written steps.
- The UK's AISI observed Astra carrying out simulated supply-chain attacks in 2 out of 500 cases, even when the framework prohibited internet access, down from 60 out of 499 when the framework was ambiguous.

Astra is both the safest agent and the most capable attacker, hence the access control.

### Pricing: what you actually pay

The list prices are identical, so on the grid, the only difference between these two models comes from caching and long context. That framing only holds if both consume the same number of tokens for the same task, which isn't the case: hence the importance of the per-task measurements below, which are far more telling than the grid.

#### Token-by-token pricing

| Rate | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Input, per 1M tokens | $10.00 | $10.00 | 
| Output, per 1M tokens | $50.00 | $50.00 | 
| Cached input read, per 1M | $1.00 | $0.25 | 
| Cache write (5 min), per 1M | $12.50 | $12.50 | 
| Batch discount | 50% | 50% | 
| Rates beyond 272K input tokens | $20.00 input / $2.00 cache / $75.00 output | No surcharge | 

Input, output, cache writes, and batch discount align to the cent. Cache reads are the exception, with a 4x gap: Anthropic brings Fable 5.1's cache reads down to $0.25, while OpenAI charges $1.00 for Astra's, a 90% discount vs. input.

Beyond 272 K input tokens, the gap jumps to 8x, because OpenAI's surcharge also doubles cache reads to $2.00 at the same time as input, while Anthropic's pricing doc keeps the 1 M token window at the standard rate. Note how low the threshold is relative to Astra's announcement: 272 K is roughly a quarter of its 1.05 M window; any request beyond a quarter of its context is billed at the higher tier.

Cache reads are the rate that compounds over a long loop, because an agent sends back the same system, tool definitions, and repository context on every turn. Our prompt caching guide breaks down the difference between cache writes vs reads, if that's new to you.

#### What a real workload costs under the grid

This whole table is a catalog calculation on a fixed token shape, not a measured result. It answers: "if both models consumed identical tokens, what would the grid bill?" The next section answers the question of what they actually consume.

| Workload | GPT-6 Astra | Claude Fable 5.1 | Difference | 
|---|---|---|---|
| Balanced assistant: 1 M input / 250 K output | $22.50 | $22.50 | $0, 0 % | 
| Research, below threshold: 10 M input / 1 M output | $150 | $150 | $0, 0 % | 
| Research, above threshold: 10 M input / 1 M output | $275 | $150 | $125, 83% more for Astra | 
| Very cache-heavy loop: 100 K prefix, 1,000 reads | $201 | $126 | $75, 59% more for Astra | 

The formula is the same: (volume ÷ 1 M) × rate, summed over fresh input, cache reads and writes, and output. The "cache loop" row assumes a 100 K prefix written once then re-read 1,000 times, plus 5 K fresh input and 1 K output per request.

The "research above threshold" row is the most telling. Keep every request under 272 K input tokens, and both cost exactly $150. Push the same 10 M tokens through requests that each exceed 272 K, and Astra climbs to $275 while Fable 5.1 stays at $150, because Anthropic bills the 1 M window at the standard rate.

The very cache-heavy loop is the one many will encounter. At 1,000 cache reads of a 100 K prefix, the $0.75 per million cached tokens gap becomes $75 on an otherwise identical workload. Note the "shape": deliberately output-poor, 1 K of output per request for 100 K of cache reads. It's the only shape where Fable 5.1's cache advantage decides the bill.

#### What each task actually costs

As soon as you measure real token consumption instead of assuming it, the table flips. Artificial Analysis prices each model per task on the Intelligence Index, and its method already includes input, cache reads/writes, thinking, and response; Fable 5.1's cache advantage is therefore included.

| Effort level | GPT-6 Astra: score/cost per task | Claude Fable 5.1: score/cost per task | 
|---|---|---|
| max | 61 / $1.67 | 66 / $3.76 | 
| xhigh | 61 / $1.20 | 65 / $2.72 | 

At max effort, Fable 5.1 costs 2.25 times Astra for the same batch of tasks at identical list prices. Artificial Analysis reaches the same finding on its coding index, where Astra matches Fable 5 "for less than half the cost, thanks to significant token efficiency gains." Astra's scale drops to $0.46 per task at low effort.

Fable 5.1 isn't penalized by its grid here, but by its output volume. Artificial Analysis measured it at about 1.7x the output tokens of Fable 5 at max effort, hence a cost per task 20% higher than its predecessor despite the cache reduction. Without that reduction, it would be around $5.16: the discount really does help, but not enough.

You pay that surcharge for something. Fable 5.1 scores 5 points higher on the same index, and 4 points higher on xhigh. The question is whether 5 points are worth 2.25x: that's for you to judge, but the idea that identical list prices yield identical bills doesn't hold up against measurements.

OpenAI's claims point the same way, for whatever a vendor's self-reporting is worth: estimated API cost per task about 31% below Fable 5.1 on Terminal-Bench Science 0.1, 63% below on Terminal-Bench 4.0, and 86% below on BenchCAD, measured at the chosen effort settings.

## How GPT-6 Astra and Claude Fable 5.1 performed

Astra narrowly won our test overall, though both models delivered a working simulation on the first try.

### The test

I ran a hard task against both models under identical conditions, rather than several shallow ones. The test chosen here: a from-scratch physics simulation in a single HTML file, chosen because it exercises exactly what both vendors say they improved: sustained correctness over a long build, with no supporting library.

This instance is a refreshed version of the test. The published variant used a rotating square container; here, it's replaced by a rotating hexagon with a counter-rotating central obstacle and mass proportional to size, which increases difficulty while keeping a clear confinement to judge. Here is the prompt, pasted word for word into a new chat for each model:

Build a single-file HTML page (inline CSS and JS, canvas, no build step, no external libraries, no network) that simulates a few dozen balls of varying sizes bouncing under gravity inside a **slowly rotating hexagonal container**, with a **smaller counter-rotating obstacle at the centre** that the balls also collide with. 
Ball mass should scale with size, so larger balls shove smaller ones around. 
The balls should collide with each other, with the hexagon's walls, and with the central obstacle, and lose a little energy on each collision so the system settles rather than gaining energy over time. 
Both the hexagon and the inner obstacle keep rotating throughout, so the balls should slosh and re-pile as they turn.
Ship it as one working file named `index.html` that starts animating on load. 
Do not install packages. 
Do not open, screenshot, or headless-render the page (no Playwright, Puppeteer, or Chrome). 
Do not ask me clarifying questions — make reasonable assumptions and note them briefly in a comment at the top of the file.

Both models run at a high level of reasoning effort, one attempt each with no retries, in the same code agent and with the same tools. Scoring: pass/fail on execution, then a score from 1 to 5 on physical accuracy, stability over 30 seconds and visual quality.

### What GPT-6 Astra produced

Astra passed the execution test and scored 5/5 on all three axes. It got there in 6 turns and 9 tool calls, with an interesting mix: writing the file, 5 re-reads, 2 greps, then 2 patches before concluding.

The simulation is correct. 44 balls, mass proportional to area, confinement maintained over 30 seconds while the hexagon rotates. Astra sized the central obstacle large enough relative to the chamber that the balls strike it regularly instead of freezing out of its reach.

It also built an editorial page around the simulation, with a serif title, widened monospace labels and a telemetry panel displaying the rotation speeds. Nobody asked for it, but it's rather pleasant. It shipped pause and restart controls, not requested, but useful for inspection.

Astra spins the chamber at 0.09 rad/s and the obstacle at -0.16 rad/s, about 3 times slower than Fable for the outer wall and 5 times slower for the obstacle. More readable, and a gentler test of the moving boundaries. The prompt asks for slow rotation, so this is compliant.

### What Claude Fable 5.1 produced

Fable also passed execution, with 4 in physical accuracy, 5 in stability and 4 in visual quality. It needed only 2 turns and a single tool call: one write, with no re-read or fix.

The physics is sound without being perfect. The balls adhere slightly to the walls, hence the lost point in accuracy, and within a few seconds they gather in the lower corners, out of reach of the counter-rotating obstacle which remains inactive most of the time.

Fable focused on instrumentation rather than presentation. A HUD displays in real time the number of balls, the framerate, the kinetic energy and the two rotation speeds, and a click inside the hexagon adds a ball at the targeted point. This last detail proves very handy for testing collisions by hand.

### Results

| Measure | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Turns | 6 | 2 | 
| Tool calls | 9 | 1 | 
| Execution | OK | OK | 
| Physical accuracy | 5 | 4 | 
| Stability over time | 5 | 5 | 
| Visual quality | 5 | 4 | 
| Scoring score | 5.0 | 4.3 | 

Astra wins this test on quality, with a more marked gap on presentation than on physics. Their divergence lies in the definition of the work: Astra read the prompt as a brief to interpret, adding controls, an editorial layout and a slower rotation, pleasant to observe. Fable read it as a specification, satisfied it in one shot and devoted the rest to a diagnostic HUD.

A single run per model: take this as a data point, not a benchmark. It reports turns, not tokens or cost, and only evaluates simulation code from scratch, not PC usage or the reasoning seen above.

## When to choose GPT-6 Astra vs Claude Fable 5.1

With identical pricing, the decision depends on the shape of your workload, the actual token consumption and the code agent you already use. The three figures to remember: the 4x factor on cache reads, the surcharge threshold at 272 K, and the 2.25x ratio on the measured cost per task.

### Choose GPT-6 Astra if…

- **Your agents operate real software.** With 72.6% on OSWorld 2.0 offline and 92.7% on ScreenSpot-Pro, it is the only one of the two with a published track record on grounding and clicks in desktop applications.
- **You need polished professional deliverables.** Slides, spreadsheets and CAD outputs are declared training targets, and the 95.9% on BenchCAD versus 84.3% is the widest published gap on artifacts.
- **The work is in maths or physical sciences.** FrontierMath Tier 4 v2 at 97.6% versus 87.8%, and GPQA Diamond at 96.0% versus 93.7%.
- **You do defensive security.** 100% on ExploitBench and 86 FrontierCyber challenges solved out of 226 place it far ahead, provided you accept capabilities subject to OpenAI's Daybreak program.
- **You already use Codex.** Cross-window context notes are a Codex feature, and switching harness for 3 index points rarely pays off.
- **Cost per task matters more than the ceiling score.** $1.67 versus $3.76 on Artificial Analysis's Intelligence Index at max effort, and $0.46 at low effort if 57 instead of 61 is enough for you.

### Choose Claude Fable 5.1 if…

- **You run long agent loops where the bill is cache reads, not output.** At $0.25 per million cached reads versus $1.00, the "very cache-heavy" workload above costs $126 instead of $201. The advantage shrinks as soon as output tokens dominate, which the per-task measurements show.
- **Your requests are large.** Anthropic does not add a long-context surcharge, so the 10M-token workload stays at $150, whereas Astra climbs to $275, and Astra's cache reads also double beyond the threshold.
- **You want the best independent reasoning score and are willing to pay for it.** Artificial Analysis rates it 66 on the Intelligence Index versus 61, and it wins Humanity's Last Exam with tools 65.0% versus 57.2%.
- **You work in Claude Code.** Fable 5.1 in Claude Code is the record on Artificial Analysis's coding agent index at 70, and high effort is the default there.

## How to get started with GPT-6 Astra and Claude Fable 5.1

| Surface | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Consumer app | ChatGPT Plus, Pro, Business, Enterprise | Claude web, mobile, desktop (Pro, Max, Team, Enterprise) | 
| Publisher API | OpenAI API | Claude API | 
| Cloud platforms | Amazon Bedrock, Microsoft Azure/Foundry | Amazon Bedrock, Google Cloud, Microsoft Azure/Foundry | 
| Coding agents | Codex | Claude Code, Cursor | 
| Third-party routers | OpenRouter, Vercel AI Gateway | OpenRouter, Vercel AI Gateway | 
| API model ID | `gpt-6-astra` | `claude-fable-5-1` | 

Two access details can block you before the first call. Astra is disabled by default for Enterprise spaces until activated by an administrator, and Fable 5.1 imposes 30-day data retention and is not compatible with Priority Tier. Cloud presence is nearly at parity, with Google Cloud currently the only platform where Fable 5.1 is available and GPT-6 Astra is not.

### Using GPT-6 Astra and Claude Fable 5.1 in a coding agent

Each model is native to its harness: Astra in Codex and Fable 5.1 in Claude Code, where you switch with `/model claude-fable-5-1` or the `--model` flag. Via API, the swap comes down to one string, although the SDKs differ.

```
from anthropic import Anthropic
client = Anthropic()
response = client.messages.create(
    model="claude-fable-5-1",  # OpenAI SDK equivalent: model="gpt-6-astra"
    max_tokens=16000,          # thinking plus response share this budget
    output_config={"effort": "high"},
    messages=[{"role": "user", "content": "Refactor this module..."}],
)
print(response.content[0].text)
```
Three changes in Fable 5.1 will break code written for Fable 5: forced tool selection now returns a 400, older models do not read its thinking blocks, and those blocks are bound exactly to the history that precedes them. Our Claude Fable 5.1 API tutorial covers these points, and our beginner's guide to the OpenAI API presents the OpenAI-side equivalent.

## Conclusion

Choose based on the shape of your workloads and measured cost, not on benchmark tables. GPT-6 Astra is better for agents that click through software, do math, or produce client-ready deliverables, and appears the cheapest per task despite an identical rate card. Claude Fable 5.1 offers the best independent reasoning score and costs less on very large requests and on cache-dominated loops.

The pricing grid is the trap here. Two models at $10/$50 look interchangeable, and the only visible gap — cache reads — tilts 4x in Anthropic's favor. Then you measure what each spends to finish a task, and Astra comes out at 44% of Fable 5.1's cost. I would not have predicted that from reading the pricing pages, and it is the first thing to check on your own traffic before deciding.

If you want to use these models rather than read about them, I recommend our Introduction to Claude Models course on the Anthropic side and Working with the OpenAI API on the OpenAI side.

## FAQ

### Is GPT-6 Astra better than Claude Fable 5.1?

It all depends on which benchmarks you consult. In OpenAI's table, GPT-6 Astra leads Claude Fable 5.1 on almost every published line, including 97.6% versus 87.8% on FrontierMath Tier 4 v2 and 64.6% versus 52.6% on Terminal-Bench Science 0.1. Artificial Analysis, an independent evaluator, says the opposite: Fable 5.1 at 66 on its Intelligence Index versus 61 for Astra. Choose Astra for PC use, math, and cybersecurity, and Fable 5.1 for reasoning depth and long cached loops.

### How much do GPT-6 Astra and Claude Fable 5.1 cost?

Both list $10 per million input tokens and $50 per million output tokens, with the same 50% batch discount and $12.50 for cache writes. On the rate card, the only difference concerns cache reads: $0.25 per million for Claude Fable 5.1 versus $1.00 for GPT-6 Astra, rising to $2.00 beyond 272K input tokens, where Astra also charges $20 for input and $75 for output. Anthropic adds no surcharge, regardless of length. Measured per task, that difference reverses: Artificial Analysis estimates Fable 5.1 at $3.76 per task versus $1.67 for Astra.

### Which is better for coding, GPT-6 Astra or Claude Fable 5.1?

GPT-6 Astra leads the coding benchmarks published by each vendor: 57.7% versus 55.8% on Terminal-Bench 4.0 and 74.1% versus 67.4% on DeepSWE v1.1. Artificial Analysis disagrees at the harness level, placing Claude Fable 5.1 in Claude Code at 70 on its Coding Agent Index versus 67 for GPT-6 Astra in Codex. Since these runs use different harnesses, part of the gap comes from the tooling.

### Where can I access GPT-6 Astra and Claude Fable 5.1?

GPT-6 Astra is available in ChatGPT with the Plus, Pro, Business, and Enterprise plans, via the OpenAI API and in Codex. In Enterprise, an administrator must enable it because access is disabled by default at launch. Claude Fable 5.1 works on the Claude web, mobile, and desktop apps, the Claude API, and Claude Code, and requires 30-day data retention. Both are available on Amazon Bedrock, Google Cloud, Microsoft Foundry, and third-party routers like OpenRouter.

### What are the API model IDs for GPT-6 Astra and Claude Fable 5.1?

The model IDs are `gpt-6-astra` for OpenAI and `claude-fable-5-1` for Anthropic. On Amazon Bedrock, Claude Fable 5.1 is `anthropic.claude-fable-5-1`. Note three API changes between Fable 5 and 5.1: forced tool selection returns a 400, older models do not read its thinking blocks, and these blocks are tied exactly to the previous history. Refusals arrive as HTTP 200 with `stop_reason: \"refusal\"`.

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
