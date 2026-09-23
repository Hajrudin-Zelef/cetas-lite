---
id: vague2-nerdykings/nerdykings/openai-navier-stokes-88-heures
title: "OpenAI Aurait Résolu Navier-Stokes En 88 Heures Avec 10 000 Agents"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09-01", "2026-09-10", "2026-09-23"]
keywords: ["agent", "agents", "astra", "claude", "compute", "energy", "governance", "gpt-6", "lean", "reasoning", "research", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/openai-navier-stokes-88-heures.md
source_anchor: ""
source_lines: [1, 67]
sha256: 38b4845f588a9bc00900305b5c45d80ea28968d9b216a51b702323ac924cbd50
---

# OpenAI Aurait Résolu Navier-Stokes En 88 Heures Avec 10 000 Agents

## Metadata

- **Source** : https://www.nerdykings.com/blog/openai-navier-stokes-88-heures.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

OpenAI claims to have obtained, in just 88 hours, a proof for a problem that has resisted mathematicians for nearly 90 years. The model behind the result is not even GPT-6 Astra: it is reportedly a new internal model still in training, presented as significantly more powerful, deployed across nearly 10,000 agents working in parallel. Their goal: solve a formulation of the Navier-Stokes problem, one of the seven Millennium Prize problems, with a $1M reward. If confirmed, this would be one of the most important scientific discoveries ever produced by a model — but the story is far more complicated than OpenAI's announcement suggests.

Navier-Stokes equations describe fluid motion (water, air, smoke, blood, turbulence around an aircraft) and are already used in countless simulations. The real problem is that nobody knows whether they remain mathematically stable in all situations. Take a fluid with perfectly regular initial motion: does its evolution always remain regular, or could a tiny region accelerate indefinitely, reaching infinite speed in finite time? This mathematical breaking point is called a **singularity** or **blow-up**. In reality a fluid cannot reach infinite speed, but if a singularity can appear in the equations, it reveals a fundamental limit of the mathematical model. Since 1934, thanks to Jean Leray, "weak" solutions are known to exist, but nobody had determined whether they can always remain perfectly regular. In 2000, the Clay Institute placed this question among the seven Millennium problems, with a $1M prize.

The system reportedly found a precise scenario in which Navier-Stokes equations actually produce a singularity. Initially the fluid is perfectly still. An external force gradually creates a vortex that spins ever faster, tightens around its center, and stretches up and down. As the vortex narrows, its speed increases until theoretically infinite in finite time. For the demonstration to be valid, it isn't enough to apply an infinite force (that would be like claiming a discovery from an infinite-energy explosion). The difficulty was to use a perfectly regular and bounded force while still producing a singularity — which OpenAI's system reportedly achieved. This type of force is explicitly allowed in certain official versions of the problem, so if all mathematical conditions are met, the proof could genuinely resolve one of the Clay formulations.

How it was obtained: on September 1, 2026, OpenAI learned researchers had made important advances on two Millennium problems. It tested a new internal model, still in training and much more powerful than GPT-6 Astra, deploying thousands of parallel agents — some searching for solutions, others verifying calculations, others gathering the most promising findings. First, ~100 agents worked 50 hours on the Euler equations (fluid motion without viscosity). After encouraging results, OpenAI focused resources on Navier-Stokes: ~10,000 agents, and after 88 hours the system built a complete mathematical demonstration. In total, agents produced 2.7 million messages and nearly 130 billion tokens for this single problem. GPT-6 Astra then spent ~17 additional hours translating the demonstration into **Lean**, a mathematical verification tool, to check each reasoning step and reduce the risk of a hidden error or hallucination.

Why it is not (yet) "solved": Lean only verifies what was encoded. Mathematicians must still ensure the hypotheses given to the software exactly match the official Clay problem. Navier-Stokes is therefore still considered unsolved and the $1M has not been awarded. The discovery also didn't come from nowhere: much of the strategy relies on work by **Diego Córdoba** and **Luis Martínez Zoroa**, who developed layers and vortices producing a cascade toward a singularity, though their force didn't meet all required properties. **Tristan Buckmaster** (NYU professor) had tried to cross that final step with Claude and Codex, partnering with an Anthropic-employed researcher working personally, achieving important results on Euler and related systems — rumors of which reportedly reached OpenAI, which then launched its massive operation.

A controversy erupted: Buckmaster asked whether his private Codex conversations could have influenced OpenAI's model, and OpenAI's first response left uncertainty. On September 10, OpenAI updated its statement after an internal investigation, asserting the relevant Codex prompts could not have influenced the system, directly or via training. So one cannot truly claim OpenAI used their data — but the intellectual debt to Córdoba and Martínez Zoroa is very clear. The affair raises a fundamental question: how to protect scientific priority when a company can mobilize 10,000 agents as soon as it hears of a promising lead?

Why math advances faster than other fields with AI: formal mathematics offers automatic verification. A model proposes a step, Lean verifies it, giving an immediate valid/invalid signal. This propose-verify-correct loop runs at speeds impossible for a human, and thousands of agents exploring directions simultaneously turn scientific research into a giant automated process. The result still depends on decades of human work and choosing the right lead, but scale and speed have abruptly changed.

The author concludes: if confirmed, it would be enormous — an AI producing a real advance on a 90+ year-old problem, using a model not yet public and already more powerful than GPT-6 Astra. But two cautions remain: technically, Lean only verifies what it is given, so independent mathematicians must confirm the hypotheses match the Clay problem; ethically, the priority dispute with Buckmaster is unsettling, and the ability to throw 10,000 agents at any heard-about lead changes the rules for researchers lacking such compute power.

## Key points

- OpenAI reportedly obtained a Navier-Stokes singularity proof in 88 hours using ~10,000 parallel agents.
- The model was a new internal model still in training, more powerful than GPT-6 Astra — not yet public.
- Navier-Stokes is a Millennium Prize problem ($1M), open since Jean Leray's 1934 weak solutions; Clay listed it in 2000.
- The constructed scenario uses a regular, bounded force to create a vortex singularity — allowed in some official formulations.
- ~100 agents worked 50 h on Euler equations first, then ~10,000 agents produced 2.7M messages / ~130B tokens.
- GPT-6 Astra spent ~17 h translating the proof into Lean for formal verification.
- Still not "solved": Lean verifies only encoded hypotheses; mathematicians must confirm they match the official Clay problem.
- Intellectual debt to Córdoba & Martínez Zoroa is clear; a priority controversy involves Tristan Buckmaster, Claude, and Codex.
- OpenAI's Sept 10 internal investigation claims Codex prompts could not have influenced the system.
- Formal math advances fast due to an automatic propose-verify-correct loop.

## Technical data / figures

| Item | Value |
|---|---|
| Time to proof | 88 hours |
| Agents mobilized | ~10,000 |
| Initial Euler phase | ~100 agents, 50 hours |
| Messages produced | 2.7 million |
| Tokens produced | ~130 billion |
| Lean translation time | ~17 hours (GPT-6 Astra) |
| Prize | $1 million (Clay Institute) |
| Problem age | ~90 years (Leray 1934; Clay 2000) |
| OpenAI announcement update | September 10, 2026 |
| Trigger date | September 1, 2026 |

- Key concepts: **singularity / blow-up**, **weak solutions** (Leray), **Euler equations** (inviscid), **Lean** formal verification
- Named researchers: Jean Leray, Diego Córdoba, Luis Martínez Zoroa, Tristan Buckmaster

## Why this source matters for the RAG

This article is a detailed case study of large-scale multi-agent AI applied to frontier mathematics, including the technical mechanism (Lean verification), scale metrics, and the emerging scientific-priority controversy. It is essential for a RAG knowledge base on autonomous research agents, AI-for-science, and the ethics/governance of AI-driven discovery.

## Source URL

https://www.nerdykings.com/blog/openai-navier-stokes-88-heures.html
