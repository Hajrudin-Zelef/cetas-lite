---
id: collect-240926-huggingface/huggingface/minimaxai-minimax-m2-5-hugging-face-3
title: "minimaxai-minimax-m2-5-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "MiniMax"]
dates: []
keywords: ["agent", "agentic", "agents", "benchmark", "claude", "cost", "leaderboard", "memory", "pricing", "reasoning", "research", "sandbox"]
source: docs/RAG/clean_en/huggingface/minimaxai-minimax-m2-5-hugging-face.md
source_anchor: ""
source_lines: [112, 136]
sha256: 8ac06d340e7a24476ee0a27ece236978fc994a2fc46cdda445bc603989a3b42f
---

# minimaxai-minimax-m2-5-hugging-face

- SWE benchmark: SWE-bench Verified, SWE-bench Multilingual, SWE-bench-pro, and Multi-SWE-bench were tested on internal infrastructure using Claude Code as the scaffolding, with the default system prompt overridden, and results averaged over 4 runs. Additionally, SWE-bench Verified was also evaluated on the Droid and Opencode scaffoldings using the default prompt.
- Terminal Bench 2: We tested Terminal Bench 2 using Claude Code 2.0.64 as the evaluation scaffolding. We modified the Dockerfiles of some problems to ensure the correctness of the problems themselves, uniformly expanded sandbox specifications to 8-core CPU and 16 GB memory, set the timeout uniformly to 7,200 seconds, and equipped each problem with a basic toolset (ps, curl, git, etc.). While not retrying on timeouts, we added a detection mechanism for empty scaffolding responses, retrying tasks whose final response was empty to handle various abnormal interruption scenarios. Final results are averaged over 4 runs.
- VIBE-Pro: Internal benchmark. Uses Claude Code as the scaffolding to automatically verify the interaction logic and visual effects of programs. All scores are computed through a unified pipeline that includes a requirements set, containerized deployment, and a dynamic interaction environment. Final results are averaged over 3 runs.
- BrowseComp: Uses the same agent framework as WebExplorer (Liu et al., 2025). When token usage exceeds 30% of the maximum context, all history is discarded.
- Wide Search: Uses the same agent framework as WebExplorer (Liu et al., 2025).
- RISE: Internal benchmark. Contains real questions from human experts, evaluating the model's multi-step information retrieval and reasoning capabilities when combined with complex web interactions. A Playwright-based browser tool suite is added on top of the WebExplorer (Liu et al., 2025) agent framework.
- GDPval-MM: Internal benchmark. Based on the open-source GDPval test set, using a custom agentic evaluation framework where an LLM-as-a-judge performs pairwise win/tie/loss judgments on complete trajectories. Average token cost per task is calculated based on each vendor's official API pricing (without caching).
- MEWC: Internal benchmark. Built on MEWC (Microsoft Excel World Championship), comprising 179 problems from the main and other regional divisions of Excel esports competitions from 2021–2026. It evaluates the model's ability to understand competition Excel spreadsheets and use Excel tools to complete problems. Scores are calculated by comparing output and answer cell values one by one.
- Finance Modeling: Internal benchmark. Primarily contains financial modeling problems constructed by industry experts, involving end-to-end research and analysis tasks performed via Excel tools. Each problem is scored using expert-designed rubrics. Final results are averaged over 3 runs.
- AIME25 ~ AA-LCR: Obtained through internal testing based on the public evaluation sets and evaluation methods covered by the Artificial Analysis Intelligence Index leaderboard.

- Downloads last month
- 284,797

## Spaces using MiniMaxAI/MiniMax-M2.5 100

## Collection including MiniMaxAI/MiniMax-M2.5

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  85.2
- mercor/apex-agents · Apex Agents View evaluation results  source leaderboard6.2<sup>*</sup>
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results  source leaderboard75.8<sup>*</sup>
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results    leaderboard  55.4
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results  source  leaderboard  80.1
- internlm/WildClawBench leaderboard
- Overall View evaluation resultssource27.1
