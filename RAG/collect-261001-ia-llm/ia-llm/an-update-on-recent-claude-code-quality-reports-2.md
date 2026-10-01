---
id: collect-261001-ia-llm/ia-llm/an-update-on-recent-claude-code-quality-reports-2
title: "an-update-on-recent-claude-code-quality-reports"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/an-update-on-recent-claude-code-quality-reports.md
source_anchor: ""
source_lines: [63, 71]
sha256: b1004ede5c697ce1c2963ec5e5d5dbd5cfd809fe5fa38ae07ffe23dbb79ecb81
---

# an-update-on-recent-claude-code-quality-reports

We are going to do several things differently to avoid these issues: we’ll ensure that a larger share of internal staff use the exact public build of Claude Code (as opposed to the version we use to test new features); and we'll make improvements to our Code Review tool that we use internally, and ship this improved version to customers.

We’re also adding tighter controls on system prompt changes. We will run a broad suite of per-model evals for every system prompt change to Claude Code, continuing ablations to understand the impact of each line, and we have built new tooling to make prompt changes easier to review and audit. We've additionally added guidance to our CLAUDE.md to ensure model-specific changes are gated to the specific model they're targeting. For any change that could trade off against intelligence, we'll add soak periods, a broader eval suite, and gradual rollouts so we catch issues earlier.

We recently created @ClaudeDevs on X to give us the room to explain product decisions and the reasoning behind them in depth. We'll share the same updates in centralized threads on GitHub.

Finally, we’d like to thank our users: the people who used the `/feedback` command to share their issues with us (or who posted specific, reproducible examples online) are the ones who ultimately allowed us to identify and fix these problems. Today we are resetting usage limits for all subscribers.

We’re immensely grateful for your feedback and for your patience.
