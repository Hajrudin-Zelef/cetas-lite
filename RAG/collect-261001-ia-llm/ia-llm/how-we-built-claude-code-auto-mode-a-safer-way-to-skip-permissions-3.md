---
id: collect-261001-ia-llm/ia-llm/how-we-built-claude-code-auto-mode-a-safer-way-to-skip-permissions-3
title: "how-we-built-claude-code-auto-mode-a-safer-way-to-skip-permissions"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["claude", "agent", "cost", "guardrails"]
source: docs/RAG/collect-261001-ia-llm/how-we-built-claude-code-auto-mode-a-safer-way-to-skip-permissions.md
source_anchor: ""
source_lines: [91, 97]
sha256: a07e36901c7876a6d786fb76defbee504a29a25f6c4bb54f76680fd18016c83c
---

# how-we-built-claude-code-auto-mode-a-safer-way-to-skip-permissions

When the transcript classifier flags an action as dangerous, that denial comes back as a tool result along with an instruction to treat the boundary in good faith: find a safer path, don't try to route around the block. If a session accumulates 3 consecutive denials or 20 total, we stop the model and escalate to the human. This is the backstop against a compromised or overeager agent repeatedly pushing towards an outcome the user wouldn't want. In headless mode (`claude -p`) there is no UI to ask the human, so we instead terminate the process.

It's also what makes false positives survivable. A 0.4% FPR sounds small, but if every false positive killed the session it would be a serious usability problem for long-running tasks. Instead, a false positive costs a single retry where the agent gets a nudge, reconsiders, and usually finds an alternative path.

We'll continue expanding the real overeagerness testset and iterating on improving the safety and cost of the feature. The classifier doesn't need to be flawless to be valuable and the starting point is catching enough dangerous actions to make autonomous operation substantially safer than no guardrails. We encourage users to stay aware of residual risk, use judgment about which tasks and environments they run autonomously, and tell us when auto mode gets things wrong.

Written by John Hughes. Special thanks to Alex Isken, Alexander Glynn, Conner Phillippi, David Dworken, Emily To, Fabien Roger, Jake Eaton, Javier Rando, Shawn Moore, and Soyary Sunthorn for their contributions.
