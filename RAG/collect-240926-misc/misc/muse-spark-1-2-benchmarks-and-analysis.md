---
id: collect-240926-misc/misc/muse-spark-1-2-benchmarks-and-analysis
title: "Muse Spark 1.2: Improved Agentic Performance at Higher Cost per Task"
domain: artificialanalysis
role: reference
task: reference
actors: ["Anthropic", "Meta", "Moonshot", "OpenAI", "SpaceX", "United States", "xAI"]
dates: ["2026-09-21", "2026-09-22"]
keywords: ["agent", "agentic", "cost", "muse", "muse spark", "benchmarks", "claude", "context window", "fable 5", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/clean_en/misc/muse-spark-1-2-benchmarks-and-analysis.md
source_anchor: ""
source_lines: [1, 67]
sha256: 23c2e2bc36261a6eae72ee7c4b23d8ac312d306e83ed259b26707ea394759a85
---

# Muse Spark 1.2: Improved Agentic Performance at Higher Cost per Task

<!-- source: https://artificialanalysis.ai/articles/muse-spark-1-2 -->

# Muse Spark 1.2: Improved Agentic Performance at Higher Cost per Task

See model page
**Meta’s Muse Spark 1.2 scores 54 on the Artificial Analysis Intelligence Index. Its Meta's third release in four months, significantly improving agentic knowledge work capabilities over prior releases and putting Meta next to SpaceXAI in a tie for third place amongst US labs**

Muse Spark 1.2 (xhigh) lands at 54, up 3 points from Muse Spark 1.1 (51) and 11 points from Muse Spark 1.0 (43, April). It enters effectively tied with GPT-5.5 (xhigh, 55) and Grok 4.5 (high, 54), narrowly behind the current frontier models Claude Opus 5 (max, 61), Claude Fable 5 (max w/ fallback, 60), GPT-5.6 Sol (max, 59), and Kimi K3 (max, 57). Meta shared access with us ahead of public release for benchmarking.

**Key takeaways:**

➤ Muse Spark 1.2 gets closer to the frontier on agentic knowledge work: At Muse Spark 1.1’s launch, we noted agentic knowledge work as its clearest gap; Muse Spark 1.2’s gains help to close this. Its GDPval-AA v2 Elo rose 260 points to 1631, #5 among all models we have benchmarked and ahead of Claude Opus 4.8 (max, 1588). GDPval-AA v2 is our leading metric for general agentic performance, measuring models on realistic knowledge work tasks such as preparing presentations and analysis, with shell access and web browsing capabilities in an agentic loop via our reference harness Stirrup. Terminal-Bench v2.1, which measures agentic coding and terminal use, gained 2 points (78% to 80%), and τ³-Banking, which measures agentic tool use, rose 2 points (25% to 27%).

➤ Among the most cost-efficient models at its intelligence level: Muse Spark 1.2 costs $0.40 per Intelligence Index task at Meta’s unchanged $1.25/$4.25 per 1M token pricing, with only Grok 4.5 (high, $0.37) and GPT-5.6 Sol (medium, $0.39) cheaper in its intelligence cluster - GPT-5.6 Terra (max, $0.51), Kimi K3 (max, $0.86), and GPT-5.5 (xhigh, $1.18) all cost more per task. The cost increase over Muse Spark 1.1 ($0.29 per task) is driven by increased token usage per Intelligence Index task, with input tokens up ~53% and output tokens up ~36%, concentrated in GDPval-AA v2.

➤ AA-Omniscience abstention rate increases: The score rose from 18 to 22 as the hallucination rate fell 10 points (38% to 28%) and the attempt rate dropped from 82% to 67%. AA-Omniscience measures knowledge reliability and hallucination: it rewards correct answers, penalizes hallucinations, and applies no penalty for declining to answer. Muse Spark 1.2’s heavy abstention (not answering questions when unsure) now drives both the low hallucination rate and a lower accuracy (41% to 38%).

➤ Scientific Reasoning results remain largely unchanged: CritPt, which tests research-level physics reasoning, notably gained 3 points (15% to 18%), while SciCode fell 2 points (58% to 56%) and Humanity’s Last Exam fell 1 point (45% to 44%).

➤ Model details: Muse Spark 1.2 retains Muse Spark 1.1’s 1M token context window and pricing ($1.25/$4.25 per 1M input/output tokens, with cache hits discounted to $0.15 per 1M), and is available on Meta’s first-party API at launch.


**Agentic and intelligence evaluations**

The 3 point gain over Muse Spark 1.1 on the Artificial Analysis Intelligence Index is concentrated in agentic evaluations: GDPval-AA v2 +260 Elo (1371 to 1631), Terminal-Bench v2.1 +2 points (78% to 80%), and τ³-Banking +2 points (25% to 27%). The minor regressions are SciCode (-2 points) and Humanity’s Last Exam (-1 point).


**GDPval-AA v2 Leaderboard**

Muse Spark 1.2 ranks #5 on GDPval-AA v2 at 1631 Elo, behind Claude Opus 5 (max, 1852), GPT-5.6 Sol (max, 1730), and Kimi K3 (max, 1685), and ahead of Claude Opus 4.8 (max, 1588). Muse Spark 1.1 scored 1371 at its launch last month.


**Intelligence Index vs Cost per Intelligence Index Task**

Muse Spark 1.2 sits near the Pareto frontier of Intelligence vs Cost per Intelligence Index Task, at $0.40 per task.


**AA-Omniscience Index**

Muse Spark 1.2 continues Meta’s AA-Omniscience pattern: the score rose from 18 to 22, driven by abstention for the second consecutive release. The hallucination rate fell 10 points (38% to 28%) as the attempt rate dropped to 67%, while accuracy slipped from 41% to 38%.


**Intelligence Evaluations**

Full breakdown of the individual evaluations in the Artificial Analysis Intelligence Index:

See Artificial Analysis for further details and benchmarks of Muse Spark 1.2: https://artificialanalysis.ai/models/muse-spark-1-2

#### Read the latest

### GPT-6 Sol and Luna push the cost efficiency frontier

GPT-6 Sol and Luna push the cost efficiency frontier by halving cost relative to GPT-5.6 Sol and Luna. Intelligence Index and Coding Agent Index scores remain level with GPT-5.6, with progress in some evaluations and regressions in others

September 22, 2026

### Claude Opus 5.5 takes the top spot on the Artificial Analysis Intelligence Index

Anthropic's new Opus scores 58 and arrives with a 20% price cut and a larger cache hit discount

September 22, 2026

### Benchmarking Grok 4.7

Grok 4.7 scores 46 on the Artificial Analysis Intelligence Index to bring SpaceXAI into the top 4 AI labs. Coding Agent Index performance has also improved, overtaking GPT-5.6 Sol

September 21, 2026
