---
id: collect-mindstudio/mindstudio/google-wiki-skill-agent-memory
title: "Google's Wiki Skill: How AI Agents Get Persistent Memory"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "memory", "fine-tuning", "inference", "research"]
source: docs/RAG/Collect RAG/02_mindstudio/google-wiki-skill-agent-memory.md
source_anchor: ""
source_lines: [1, 55]
sha256: f78cc868148bcef5325479676ab5cb108da37f999e4e8b8beca68d4675c4fa3b
---

# Google's Wiki Skill: How AI Agents Get Persistent Memory

## Metadata

- **Source** : https://www.mindstudio.ai/blog/google-wiki-skill-agent-memory
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains Wiki Skill, a research proposal from Google Research describing how an AI agent can convert its own experience into persistent, evolving knowledge rather than relearning the same lesson every session. The system separates three distinct artifacts: what happened (raw logs), what was learned from it (a wiki of observations), and what the agent should do differently (skills). The paper credits Andrej Karpathy's "LLM wiki" idea as a partial inspiration and builds a more formal, agent-specific architecture on top of it.

The architecture has three layers: (1) The raw layer stores immutable execution traces — every agent attempt is recorded and never edited or deleted, functioning as a permanent audit trail (analogous to Karpathy's evidence store). (2) The wiki layer compiles observations: patterns pulled from the raw traces, notes on what worked/failed and why. It compounds and refines over time but never resets — the agent's evolving understanding of its domain built from evidence. (3) The skills layer is what actually changes agent behavior: instructions (typically plain-text markdown files, matching Anthropic's own skill file format) that describe procedures step by step and evolve based on what the wiki layer has learned.

Why keep wiki and skills separate: the wiki is a record of observations (description of reality — data), while skills are instructions to act on (logic/program). Collapsing them means every new observation directly rewrites behavior, risking instability — one bad trace could corrupt a working skill. Separating them lets the system accumulate evidence without immediately touching skill files, promoting only validated patterns. It's the classic data/logic separation: "the wiki is the database, the skill is the program that reads from it."

The four-agent loop: (1) inference agent — the worker that attempts the task using the current skill; (2) wiki maintainer — performs root cause analysis on raw traces, figures out what succeeded/failed/why, and turns it into a wiki observation; plus two more roles that update the wiki with new findings and refine the skill files based on accumulated wiki content. The loop runs continuously/indefinitely, so the knowledge base is never static — each run is a chance to improve without human rewriting of instructions.

Comparison to Karpathy's LLM wiki: Karpathy's framing was compiling knowledge once and keeping it current rather than rederiving from scratch per query — naturally applicable to personal knowledge management (cross-linked notes like a personal Wikipedia). Wiki Skill applies the same principle to agent behavior, adding the formal observation/instruction split plus the automated four-agent loop that turns traces into updated skills.

Why persistent agent memory matters now: it reflects the broader trend of frontier labs building agents that operate over long stretches without human guidance (e.g., reports about OpenAI's internal models built for persistence and multi-agent collaboration). Wiki Skill is a modest, transparent version of that need — agents that forget everything between sessions are limited by context-window size; agents that compile experience into durable knowledge can improve at a task indefinitely. Practical to build today at a conceptual level: plain text logs, markdown skill files, periodic review — the harder part is designing the wiki-maintainer review process to reliably extract correct root causes instead of noise.

## Key points

- Wiki Skill proposes a three-layer memory system: raw immutable execution traces, a compounding wiki of observations, and a skills layer that changes agent behavior.
- Skills are plain-text markdown files — easy to inspect, edit, and version compared to opaque fine-tuning.
- The wiki/skill separation is a data-vs-logic split: observations accumulate without immediately corrupting working skills.
- A four-agent loop (inference agent, wiki maintainer, + wiki updater and skill refiner) runs continuously, automating the trace→observation→skill pipeline.
- Built on Karpathy's LLM wiki concept: memory is only useful if it compounds instead of resetting.
- The main engineering difficulty is the root-cause-analysis (wiki maintainer) step extracting correct lessons from raw logs.

## Technical data / figures

| Component | Role |
|---|---|
| Raw layer | Immutable execution traces; permanent audit trail; never edited/deleted |
| Wiki layer | Compounding observations; patterns from traces; grows but never resets |
| Skills layer | Instructions (markdown) that change agent behavior; evolve from validated wiki findings |
| Inference agent | Attempts the task using current skill |
| Wiki maintainer | Root-cause analysis on raw traces → new wiki observations |
| Additional loop agents | Update wiki with findings; refine skill files |
| Skill format | Plain text / markdown (same format as Anthropic's skill files) |

## Why this source matters for the RAG

Directly relevant to designing persistent-memory and self-improving RAG systems: the raw/wiki/skills separation is a concrete architecture for turning an agent's retrieval-and-answer experience into durable knowledge. The LLM-wiki concept underpins several related patterns (OKF, LLM wiki knowledge bases, semantic memory), making this a foundational reference for agent memory design.

## Related context from the article

- Anthropic's own skill files are plain-English markdown documents agents read before tasks.
- Same underlying need as OpenAI's reported persistent internal models.
- Agents limited by context-window size vs agents that compile experience into durable knowledge.
