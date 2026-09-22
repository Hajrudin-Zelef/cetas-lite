---
id: briefing-ia-2026/11-consumer-agents-research/08-alma-memory
title: "ALMA: meta-learning of agent memories (arXiv)"
domain: consumer-agents-research
role: deep-dive
task: research
actors: ["Apple", "Inflection AI", "Microsoft"]
dates: ["2026-02"]
keywords: ["agent", "agentic", "agents", "memory", "personal agents", "research", "scout"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11-10"
source_lines: [12400, 12462]
sha256: 47a3c10eca994a5f450c5975dda7f17bb20122ed61622c7a5bad24e3aa18f2bd
---

# ALMA: meta-learning of agent memories (arXiv)

<a id="s11-10"></a>
### 11.9 ALMA: meta-learning of agent memories (arXiv 2602.07755, February 2026)

**The subject.** Published in February 2026 on arXiv
(reference 2602.07755) under the title "Automated
meta-Learning of Memory designs for Agentic systems",
ALMA tackles a central problem of agentic systems: memory.
Rather than hand-designing an agent's memory architecture
(what to store, how to index it, when to forget, how to
update it), ALMA proposes learning it automatically —
meta-learning applied to memory designs. The components —
agents, database schemas, retrieval and update mechanisms
— are expressed in code, which lets the system explore
the space of possible designs programmatically.

**The stake: continual learning without manual
engineering.** The stated goal is continual learning — an
agent that keeps learning through its interactions —
without manual memory engineering. This is the missing
link between 2026's product promises (Pi Journeys and
its "new memory approaches", Scout and its persistent
identity, Siri AI and its cross-app context) and their
technical realization: all these agents need reliable
memory over long horizons, and today that memory is
largely hand-cobbled, with the drift and forgetting
risks that entails. ALMA proposes making memory design
itself a learning problem — a step toward agents whose
cognitive infrastructure improves automatically.

**Scope and positioning.** By its date (February 2026),
ALMA is one of the year's founding papers on this theme,
predating the summer's product launches. It illustrates
the classic research → product pipeline: early-year work
on memory feeds the summer and fall personal agents. The
fact that designs are "expressed in code" is also a
significant methodological choice: it makes the search
space composable and verifiable, as opposed to purely
opaque neural approaches. For the dossier, ALMA is the
reminder that behind every marketing promise of
"persistent memory" lies an open research question — and
that 2026 is the year this question moved from the
laboratory to the product requirements document.

**ALMA as the missing link of product promises.**
Rereading the summer 2026 launches in ALMA's light is
illuminating: Pi Journeys promises "new memory
approaches" to accompany multi-year life transitions;
Scout promises a nameable persistent identity; Siri AI
promises cross-app personal context. Three promises, one
shared technical prerequisite: reliable agent memory,
that accumulates without drifting, forgets without
amputating, and stays continuously usable. ALMA, by
proposing to automatically learn memory designs instead
of hand-cobbling them, is research's answer to this
prerequisite. The gap between the marketing promise
("your agent remembers everything") and the state of the
art (February 2026: we are only just beginning to
meta-learn memory architectures) is the most
underestimated risk zone of 2026's personal agents: it is
there that disappointments will be born — the agent that
forgets a critical detail after six months — or
breakthroughs.

