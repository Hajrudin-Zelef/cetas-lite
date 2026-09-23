---
id: vague2-vision-ia/vision-ia/opus-robot
title: "Claude Opus 4.7 programme un robot 20x plus vite qu'une équipe humaine"
domain: vision-ia
role: reference
task: article
actors: ["AWS", "Alibaba", "Anthropic", "Meta", "OpenAI", "Perplexity", "Samsung", "United States"]
dates: ["2026-06-22", "2026-09-23"]
keywords: ["claude", "opus 4", "agent", "agentic", "agents", "aws", "benchmark", "chatgpt", "fine-tuning", "governance", "gpu", "inference"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/opus-robot.md
source_anchor: ""
source_lines: [1, 50]
sha256: db5e9edb97a0e65da8a69eb3e9acf228de8d598e667810cecf2eea7f5829fe1e
---

# Claude Opus 4.7 programme un robot 20x plus vite qu'une équipe humaine

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/opus-robot
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This June 22, 2026 edition of the Vision-IA newsletter is led by Anthropic publishing the second phase of Project Fetch. Facing a quadruped robot alone, Claude Opus 4.7 completed 5 tasks in 12 minutes: connecting hardware, reading camera and lidar streams, writing movement code, and locating a ball. A human team—although assisted by Claude—had needed 264 minutes for the same course, roughly 20 times longer. The model produced 1,045 lines of code versus 10,309 for the human team, nearly 10x more compact. Claude wired sensors directly to the robot and generated executable code without long trial-and-error cycles; the main gain came from quickly choosing the right interfaces and scripts that worked on the first try. The key limitation: the model failed to retrieve the ball—closed-loop control (tracking an object drifting in real time) remains out of reach. The newsletter frames this as a concrete step toward robots configured by AI themselves, while marking the frontier: reactive perception and continuous real-time control remain obstacles.

The edition also reports that Anthropic now requires identity verification to access Claude, sparking backlash (696 points on Hacker News, heavy Reddit discussion) over privacy, though Anthropic cites security and regulatory compliance. It covers the Trump administration taking a series of measures against Anthropic (per TechCrunch's Equity podcast), and a petition accusing Meta of training its AI on employee communications and code (69 points on Hacker News).

The "Research" section covers OpenAI's LifeSciBench (750 tasks designed by 173 expert practitioners, evaluating real scientific work: interpreting contradictory data, designing protocols, debugging experiments); a UC Berkeley study of 500,000+ grades showing writing/coding course results rose after ChatGPT, concentrated on homework rather than in-class exams (outsourcing rather than learning); DeepAdapt's ACI (Adaptive Continual Intelligence), an in-production active learning layer claiming up to 82% lower costs and 33x faster inference by processing known queries on CPU and routing novel ones to GPU; Perplexity Brain (self-improving agent memory building a nightly "LLM wiki"); a tutorial on fine-tuning a local Qwen 3:0.6B model for question classification (105 Hacker News points); and an opinion piece questioning the usefulness of geometric algebra (125 points).

Additional briefs include Samsung deploying ChatGPT Enterprise and Codex to all employees worldwide; Recall (local project memory for Claude Code, 101 Hacker News points); AWS launching Continuum (auto-detects/fixes code vulnerabilities) and Context (enterprise knowledge graph for agents); Sam Altman defending scaling at Stanford and criticizing skeptics; phone farms faking social media engagement; tech workers mobilizing against Silicon Valley's AI push (31 points); Zuckerberg struggling to revive Meta's hacker culture after 8,000 layoffs; CleverCrow (tokens to fund open source, 35 points); and a prompt to reproduce a writer's style via a reusable "Voice Card."

## Key points

- In Project Fetch phase 2, Claude Opus 4.7 completed 5 robot tasks in 12 minutes versus 264 minutes for a Claude-assisted human team (~20x faster).
- Claude wrote 1,045 lines of code vs 10,309 for humans (~10x more compact) and wired sensors directly.
- Main limitation: the model failed closed-loop control (retrieving a moving ball), showing real-time reactive control remains unsolved.
- Anthropic now mandates identity verification for Claude, triggering major privacy backlash (696 HN points).
- The Trump administration took measures against Anthropic; TechCrunch analyzed who benefits.
- A petition accuses Meta of training AI on employee communications and code.
- OpenAI's LifeSciBench (750 tasks, 173 experts) evaluates real scientific reasoning, not fact recall.
- DeepAdapt claims up to 82% lower costs and 33x faster inference via CPU/GPU request routing.

## Technical data / figures

| Item | Figure |
|---|---|
| Claude Opus 4.7 robot tasks | 5 tasks in 12 minutes |
| Human team baseline | 264 minutes (~20x longer) |
| Claude code output | 1,045 lines vs human 10,309 lines (~10x compact) |
| Identity verification backlash | 696 points on Hacker News |
| LifeSciBench | 750 tasks, 173 experts |
| UC Berkeley study | 500,000+ grades |
| DeepAdapt ACI | −82% costs, 33x faster inference |
| Qwen 3:0.6B fine-tuning | 600M parameters, 105 HN points |
| CleverCrow | 35 HN points |
| AWS Continuum/Context | announced at NYC Summit |

## Why this source matters for the RAG

It provides a concrete, quantified case study of an AI agent programming physical hardware dramatically faster than a human team (Project Fetch phase 2), while clearly identifying the remaining frontier (closed-loop real-time control). It captures the tension between AI capability expansion and governance (Anthropic identity verification, Trump administration pressure, Meta employee-data petition), plus benchmark/agent-infrastructure developments relevant to evaluating agentic progress.
