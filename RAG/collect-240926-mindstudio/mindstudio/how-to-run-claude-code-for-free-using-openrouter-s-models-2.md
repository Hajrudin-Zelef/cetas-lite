---
id: collect-240926-mindstudio/mindstudio/how-to-run-claude-code-for-free-using-openrouter-s-models-2
title: "how-to-run-claude-code-for-free-using-openrouter-s-models"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenRouter"]
dates: []
keywords: ["claude", "training"]
source: docs/RAG/clean_en/mindstudio/how-to-run-claude-code-for-free-using-openrouter-s-models.md
source_anchor: ""
source_lines: [75, 95]
sha256: 36cc52510f0a9264217b64ff918ada1ac8d8f59c3dc05695982ec9848fe07843
---

# how-to-run-claude-code-for-free-using-openrouter-s-models

It’s a free model available through OpenRouter from a provider called Stealth. It’s anonymous, meaning its architecture, training data, and infrastructure aren’t publicly disclosed, which is worth considering before sending sensitive data through it.

### Can I use free OpenRouter models in the Claude Code desktop app?

No. The desktop app overrides the settings file and forces the use of an Anthropic model. This method requires running Claude Code through an IDE like VS Code or directly in a terminal.

### Are OpenRouter’s free models always available?

Not permanently. OpenRouter periodically changes which models are offered for free, so a specific model may lose free-tier status. OpenRouter’s free models router can automatically select from whatever free models are currently active.

### Why did tasks take so much longer with the free model?

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

Testing showed significantly slower completion times, including a landing page build that took around six hours and a data-processing task that hit repeated timeout errors. This is likely tied to rate limits, server load on free-tier infrastructure, and the model’s own processing speed.

### Is it safe to use anonymous models like Stealth Ox Alpha for business projects?

It depends on the sensitivity of the data. Since the provider and origin aren’t disclosed, it’s worth avoiding proprietary code, confidential business data, or anything you wouldn’t want handled by an unverified third party.
