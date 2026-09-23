---
id: vague2-nerdykings/nerdykings/bloome-equipe-agents-ia-test
title: "Bloome : J'ai Testé Une Vraie Équipe D'Agents IA (Le Verdict)"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "Apple", "Google", "xAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "claude", "gemini", "grok", "opus 4", "pricing"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/bloome-equipe-agents-ia-test.md
source_anchor: ""
source_lines: [1, 67]
sha256: ac9020702159f6bcb6204a250df06ceb859872604ba9354452e5c9a7d6519bba
---

# Bloome : J'ai Testé Une Vraie Équipe D'Agents IA (Le Verdict)

## Metadata

- **Source** : https://www.nerdykings.com/blog/bloome-equipe-agents-ia-test.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Most of the time, we open a chatbot, write a big prompt, and hope it directly produces the right result. That sometimes works by luck, but on a real multi-step project this technique quickly shows its limits. The author tested **Bloome**, a platform that makes several AI agents work together in the same conversation — one that builds, one that critiques, one that improves, and the user who keeps the final decision.

The problem with a single prompt: in a real project, generating the final page is almost the last step. Before that, you must understand what you're building, who it's for, what message to convey, what objections people might have, and how to turn a raw idea into something clear and credible. A single AI can generate fast but doesn't always take the step back a real team would: it may accept a bad angle, write something that sounds good but sells nothing, or rush into execution before the problem is even framed. Bloome's idea is precisely to see what happens when several agents work in the same conversation with different roles rather than betting everything on a single prompt.

How Bloome works: after creating an account, Bloome downloads as an app on macOS, Windows, Google Play, and iOS. You then connect the agents you want to work together: **Codex, OpenCode, Hermes Agent, Claude Code, Gemini CLI, Grok**, or Bloome's native agent. For this test, the author connected Codex, then created a small custom team: a **growth specialist**, a **brand strategist**, and a **product marketing** agent, plus Codex for execution. Each agent has its own name, profile photo, and above all collaboration rules defined in the group settings: who intervenes first, what each must challenge, and a ban on directly producing a final version without going through intermediate steps.

The test: a product team for the "Nerdy Pocket." To test in real conditions, the author set up a group around a fake project: a landing page for the Nerdy Pocket, a fictional retro-futuristic handheld console he regularly uses to test AI tools. Game rule fixed in the collaboration rules: product marketing first defines targets, positioning, and models; the brand strategist then sets the brand universe, premium tone, and visual direction; the growth specialist comes last to challenge anything too vague or too marketing, before final validation by the author. Result: instead of immediately throwing a prompt to generate the site, the team first turned the raw idea into a solid brief. The agents called each other, waited their turn, and the growth specialist even sent back a critique explicitly listing what was missing — **"zero conversion strategy"** — before the author validated and moved to execution.

The critique loop: V1 → V2 → V3. Once the brief was set, Codex generated a first prototype based on the team's work. V1 respected the requested structure but remained too simple, almost a placeholder. The author asked for a more visually premium version without touching positioning or copywriting: V2 arrived clearly improved. Then he had the result reviewed by the three other agents, each giving an opinion on the premium direction, the conversion journey, and priority fixes — producing a V3 integrating their feedback. What's striking is that the author isn't the one doing the back-and-forth between each agent: they reread and critique each other, and he intervenes only to decide.

The final twist: replacing the agent with **Claude Opus 4.8**. Even after V3, the author wasn't satisfied. He made a more radical decision: replace the Codex agent with an agent running under Claude Opus 4.8, asking it to take up all the other agents' work, respect the validated structure, and deliver a much more elaborate version with animations and a wow effect. The final result — the "Nerdy Kings Ultimate Landing" — has nothing in common with V1: premium retro style respected, a playable mini-game integrated directly into the page, two editions with prices (€299 and €499), a materials section detailing brushed aluminum and gold-alloy buttons, and a final call-to-action with email capture. Quite a gap from the starting prototype.

The author's view: what struck him most wasn't the quality of the final render — any good code model can produce a nice landing page. It's that the agents call each other, critique each other, and adjust among themselves without him being the intermediary at every step. It changes the nature of the work: you go from "I prompt and correct" to "I supervise a team and decide." Honestly: if your goal is to send a single prompt and wait for the result, Bloome is useless — just open a classic chatbot. But if you have a real workflow with several steps and the need to build progressively on a good base, it's clearly a tool to test. Free credits are added daily, leaving plenty of room to form an opinion before paying.

## Key points

- **Bloome** is a platform where multiple AI agents collaborate in the same conversation with distinct roles.
- Solves the single-prompt limit for multi-step projects by adding framing and critique stages.
- Available as an app on macOS, Windows, Google Play, and iOS.
- Connects agents: **Codex, OpenCode, Hermes Agent, Claude Code, Gemini CLI, Grok**, or native Bloome agent.
- Custom team built for the test: growth specialist, brand strategist, product marketing, plus Codex for execution.
- Collaboration rules define intervention order, challenges, and ban on skipping intermediate steps.
- Test project: a landing page for the fictional "Nerdy Pocket" retro-futuristic handheld console.
- The growth specialist flagged "zero conversion strategy," showing agents actively critique each other.
- Iterative loop V1 → V2 → V3, with agents reviewing each other; author only decides.
- Final twist: replacing Codex with **Claude Opus 4.8** produced an elaborate final landing page with a mini-game and pricing.
- Free daily credits available; best for real workflows, not single-prompt use.

## Technical data / figures

| Item | Value |
|---|---|
| Platform | Bloome |
| Type | Multi-agent collaboration platform |
| Platforms | macOS, Windows, Google Play, iOS |
| Connectable agents | Codex, OpenCode, Hermes Agent, Claude Code, Gemini CLI, Grok, native agent |
| Test agents used | Codex + growth specialist, brand strategist, product marketing |
| Test project | Nerdy Pocket landing page (fictional) |
| Iterations | V1 → V2 → V3 |
| Agent swap | Codex → Claude Opus 4.8 |
| Final output | "Nerdy Kings Ultimate Landing" with mini-game |
| Editions / prices | 2 editions: €299 and €499 |
| Pricing model | Free credits added daily |

- Key concepts: **multi-agent collaboration**, **collaboration rules**, **critique loop**, **role specialization**, **human-in-the-loop decision**
- Notable critique example: "zero conversion strategy"

## Why this source matters for the RAG

This article is a hands-on evaluation of a multi-agent collaboration platform, showing a concrete workflow where agents frame, critique, and iterate under human supervision. It is valuable for a RAG knowledge base on agent orchestration, multi-agent roles, and AI-assisted product/design workflows.

## Source URL

https://www.nerdykings.com/blog/bloome-equipe-agents-ia-test.html
