---
id: vague2-nerdykings/nerdykings/recursive-mass-agents-ia-pensees-latentes
title: "Recursive Mass : Les Agents IA Qui Pensent Sans Mots (-75% Tokens)"
domain: nerdykings
role: reference
task: article
actors: ["Alibaba", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "benchmarks", "cost", "llama", "nvidia", "qwen", "reasoning", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/recursive-mass-agents-ia-pensees-latentes.md
source_anchor: ""
source_lines: [1, 48]
sha256: 9c11046466b937240c8ed29212b2a3f3610ec761eb34d62a10e1b35008bdcd16
---

# Recursive Mass : Les Agents IA Qui Pensent Sans Mots (-75% Tokens)

## Metadata

- **Source** : https://www.nerdykings.com/blog/recursive-mass-agents-ia-pensees-latentes.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article covers Recursive Multi-agent System (Recursive Mass), a paper from a team linked to UIUC, Stanford, Nvidia, and MIT, trained for about $4. The core problem: multi-agent AI teams communicate in natural language, burning tokens and losing information at each exchange. When an AI reasons internally, it manipulates rich internal states (large vectors full of nuance), but to communicate with another agent it must compress this into text — lossy compression. Each generated phrase, summary, or reformulation costs tokens, and adding agents slows the system and increases the bill.

Recursive Mass proposes that agents keep their internal numeric state and transmit it directly to the next agent, with no phrases, summaries, or reformulation — just the raw signal passing between models. Text returns only at the very end when the last agent produces the user-facing answer. Since different models (Llama, Qwen, Gemma) think in different internal formats, the researchers built a small module called the Recursive Link that translates one agent's latent state into a format the next can understand — latent-thought-to-latent-thought translation. Crucially, they do not retrain the models themselves; only this tiny inter-agent module is trained, costing about $4.

Tested on 9 benchmarks (math, science, medicine, information retrieval, code), the most striking result is on competition math: with the exact same small models, the score rises from 73% to nearly 87%. This is not just accuracy — it also gains up to 75% fewer tokens and up to 2.4x speedup, with a sweet spot around 80 latent tokens per transmission. The "recursive" aspect: the last agent can send its state back to the first, looping the whole team for another pass, and each loop refines the answer. Classic text-based agents degrade with more discussion (repetition, reinforcing wrong tracks), while latent-state agents improve per loop.

The paper preempts the objection that gains come from a strong teacher model: competing text-based methods had access to the same teacher, terrain, and supervision, and latent-state communication still won. The disturbing limit is interpretability: removing words removes the only human-readable part, creating machine "telepathy" via vectors and matrices. For high-stakes decisions (medical, legal, credit), this raises the question of whether performance gains justify losing the ability to audit inter-agent reasoning.

## Key points

- Multi-agent teams communicating in natural language suffer lossy compression and token costs.
- Recursive Mass transmits raw internal states between agents; text only appears at the end.
- Recursive Link module translates latent states between heterogeneous models.
- Only the tiny link module is trained — cost ~$4; base models untouched.
- Competition math: 73% → ~87% with identical models.
- Efficiency: up to -75% tokens, up to 2.4x faster; sweet spot ~80 latent tokens.
- Recursive looping improves answers per iteration, unlike text-based agents that degrade.
- Major limit: loss of interpretability/auditability (latent "telepathy").

## Technical data / figures

| Item | Value |
|---|---|
| Organizations | UIUC, Stanford, Nvidia, MIT |
| Training cost | ~$4 |
| Benchmarks tested | 9 (math, science, medicine, IR, code) |
| Competition math score | 73% → ~87% |
| Token reduction | up to 75% |
| Speedup | up to 2.4x |
| Optimal latent tokens/transmission | ~80 |
| Module name | Recursive Link |

## Why this source matters for the RAG

This source documents a novel multi-agent communication paradigm with striking efficiency and accuracy gains at minimal cost, plus a significant interpretability trade-off. It is valuable for RAG corpora on agent architectures, latent reasoning, and AI safety/interpretability.
