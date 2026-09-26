---
id: collect-240926-mindstudio/mindstudio/how-to-run-claude-code-for-free-using-openrouter-s-models-1
title: "how-to-run-claude-code-for-free-using-openrouter-s-models"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["claude", "agent", "agentic", "agents", "cost", "deepseek", "gemini", "glm", "memory", "research"]
source: docs/RAG/clean_en/mindstudio/how-to-run-claude-code-for-free-using-openrouter-s-models.md
source_anchor: ""
source_lines: [1, 74]
sha256: d7be49df91461dcb2a98a9dce7590f76863d08c0b31c56187df8aa893255ce34
---

# how-to-run-claude-code-for-free-using-openrouter-s-models

<!-- source: https://www.mindstudio.ai/blog/free-stealth-model-claude-code-openrouter -->

## Can you actually use Claude Code without paying for Claude?

Yes. Claude Code doesn’t check which model is generating responses, it just needs an API endpoint that speaks the right protocol. By editing the tool’s settings file and pointing it at OpenRouter instead of Anthropic’s own API, you can route requests to any model OpenRouter hosts, including a growing list of free ones. The catch is that free models come with rate limits, timeouts, and slower, less reliable output than Claude itself.

## TL;DR

- **OpenRouter** acts as a single API gateway to hundreds of AI models, several of which are listed at zero cost per million tokens.
- You can redirect Claude Code to OpenRouter by editing the `env` section of its**settings.json** file and swapping in an OpenRouter API key instead of an Anthropic one.
- This trick only works with **Claude Code run through an IDE or terminal** (like VS Code), not the desktop app, which forces you to use an Anthropic model.
- A newly released free model called **Stealth Ox Alpha** was tested inside the Claude Code harness and handled full agentic workflows, tool calls, and file edits.
- Free models are **anonymous providers** in some cases, meaning you don’t know the origin of the model or where your data goes, which matters for sensitive projects.
- In testing, tasks that normally take Claude Code a few minutes took **hours** on the free model, and some runs hit repeated**timeout errors** .
- Free models seemed usable for **knowledge work and research-style tasks** , but noticeably weaker for**complex software builds** compared to Claude or GPT-class models.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## What is OpenRouter and why does it matter here?

OpenRouter is a platform that gives developers one API key with access to a large catalog of AI models from different providers. Instead of signing up separately for OpenAI, Anthropic, Google, or smaller labs, you route all your requests through OpenRouter and pick whichever model you want per request.

The catalog includes paid models (cheap options like DeepSeek, as well as full-price models like Claude Opus, GPT, and Gemini variants) and a separate tier of free models. Searching “free” on OpenRouter’s model list surfaces a number of options with $0 listed for both input and output tokens. These free listings typically come with restrictions: rate limits, daily caps, or the possibility that the provider pulls the free tier entirely without much notice.

One model highlighted in recent testing, Stealth Ox Alpha, comes from a provider called Stealth. As the name implies, it’s an anonymous model, meaning there’s no public confirmation of who built it, what it was trained on, or where it’s hosted. That anonymity is worth weighing before sending proprietary code or sensitive business data through it.

## How do you set up Claude Code to use free OpenRouter models?

The setup is a configuration change, not a code change, and takes a few minutes.

1. **Create an OpenRouter account** at openrouter.ai.
2. **Generate an API key.** Go to your account, open the credits or API keys section, and create a new key.
3. **Find a free model.** Browse the models page, search “free,” and copy the exact model identifier (for example,`stealth/ox-alpha` or a free GLM variant).
4. **Open Claude Code’s settings file.** This only works when running Claude Code inside an editor like VS Code or directly in a terminal, not the desktop app, since the desktop app overrides these settings and forces an Anthropic model.
5. **Edit the `env` section.** Inside the settings file, there’s an environment variables block. Replace the Anthropic API token field with your OpenRouter key, and set the model fields to the model identifier you copied.
6. **Save and start a new session.** Claude Code should display the active model name at the top of the session (and in a status line, if configured), confirming the swap worked.

Because OpenRouter periodically rotates which models are free, a specific model like Stealth Ox Alpha may disappear from the free tier over time. OpenRouter also offers a “free models router” option, which automatically selects from whatever free models are currently available, removing the need to manually track which one is active.

## Does it actually work well inside Claude Code?

Functionally, yes. The free model was able to operate inside the full Claude Code agent loop: it acknowledged goals, searched files, wrote code, called tools, and even self-corrected. In one test, it built a multi-page landing site with product listings, a shopping cart, brand-consistent styling, and nutrition-fact style detail pages, all without being handed a project brief beyond a single prompt.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

It also respected the existing project context. In one case it dug through unrelated project folders to find previously written scripts and API keys, used them to pull real YouTube analytics data, and produced a structured quarterly and annual report as a Google Sheet. It picked up on stored “skills” and memory files the same way a standard Claude Code session would, since none of that agent behavior is model-specific.

The problems showed up in speed and reliability. Tasks that normally take a few minutes with a full-price Claude model stretched to 40 minutes or several hours. One run repeatedly failed with an “upstream idle timeout exceeded” error until the task was broken into smaller chunks. Another run got stuck in a retry loop and eventually gave up, reporting it was blocked by an infrastructure issue, after nearly 45 minutes of attempts.

## Where does the free model fall short?

The clearest gap is on complex, technical, multi-step engineering work. Building a full front end took roughly six hours instead of the expected fraction of that time with a paid frontier model. The output quality itself was reasonable (correct branding, working cart functionality, accurate product data) but the process was slow and occasionally errored out entirely before finishing.

For deep technical orchestration, like coordinating multiple sub-agents, planning a large codebase, or handling long dependency chains, the free model does not match what Opus-class or GPT-5-class models deliver. The rate limits and timeouts tied to free-tier access compound this, since a model that has to retry or restart loses even more time.

## Is it worth using free models in Claude Code?

For research, reporting, and general knowledge work, it holds up reasonably well. Pulling analytics, summarizing data, writing structured reports, and doing directory-level file management are tasks the free model completed correctly, if slowly.

For anything resembling serious software development, the tradeoff tips the other way. The time cost (hours instead of minutes) and the risk of a run failing partway through outweigh the savings, especially since many day-to-day coding tasks aren’t complex enough to need a frontier model in the first place, but also aren’t tolerant of multi-hour turnaround times.

The more practical takeaway is that most everyday requests don’t need a top-tier model at all. The real dividing line is task complexity: simple, well-scoped requests can run on cheaper or free infrastructure, while anything requiring an orchestrator-level model to manage many moving parts still benefits from paying for Claude or GPT directly.

## Frequently Asked Questions

### What is Stealth Ox Alpha?

