---
id: briefing-ia-2026/07-infrastructure-compute/09-aws-agentcore
title: "AWS: AgentCore, WorkSpaces and agentic infrastructure"
domain: infrastructure-compute
role: deep-dive
task: infrastructure
actors: ["AWS"]
dates: ["2026-06"]
keywords: ["agent", "agentic", "aws", "agents", "compute", "mcp", "research", "training"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s07-10"
source_lines: [8957, 9065]
sha256: 552b361e2ffa3a8922503c2aaf090881093764b9b6f0097d5bb5fb8ceb1922ea
---

# AWS: AgentCore, WorkSpaces and agentic infrastructure

<a id="s07-10"></a>
### 7.9 AWS: AgentCore, WorkSpaces, and agentic infrastructure

The AWS Summit in New York, on June 16 and 17, 2026, was the stage of a major offensive on agents.
AWS took the AgentCore harness to general availability (GA) there, as of June 17.
The service had been in public preview since April, making for rapid maturation.
The harness is built around two operations, CreateHarness and InvokeHarness.
It is provider-agnostic, that is, independent of the underlying model provider.
It is a strong strategic choice: AWS sells agent infrastructure, not a captive model.

AgentCore's provider-agnostic character deserves attention.
In a market where each cloud pushes its own models, AWS makes the opposite bet.
The company tells customers: bring your models, we provide the runtime.
It is a pure-infrastructure positioning, in AWS's historical tradition.
It seduces enterprises that do not want to bet on a single model provider.
And it prepares the future: when models commoditize, infrastructure captures the value.
AgentCore is therefore a bet on the commoditization of models themselves.

Around the harness, AWS stacked the bricks of a complete agentic platform.
A managed Knowledge Base feeds agents with enterprise knowledge.
Web Search grounding anchors agents' answers in verifiable web sources.
Observability and A/B testing bring software-engineering rigor to agents.
For an agent in production must be supervised, measured, and compared like any service.
These bricks turn the agent from a demo into an operable system.
It is the whole difference between a prototype and a product.

Two previews completed the picture, with different horizons.
AWS Context, a knowledge graph for agents, was presented in preview.
AWS Continuum, dedicated to agentic security, is also in preview, and not in GA.
This distinction matters: preview signals a direction, GA a commitment.
The knowledge graph answers agents' need to navigate linked data.
Agentic security is the most delicate construction site of all.
An agent that acts autonomously must be framed, and Continuum sketches this answer.
AWS is thus advancing on two fronts: agent power and their control.

The strategic reading is clear: AWS is building the enterprise-agent runtime layer.
The cloud no longer just hosts models, it orchestrates autonomous systems.
It is a natural extension of its historical infrastructure business.
But it is also a response to competition from integrated model platforms.
By staying agnostic, AWS positions itself as neutral ground where all models can run.
It is an audacious bet: betting on infrastructure when everyone bets on models.
AWS's history suggests this kind of bet usually works out well for it.

A few days later, around June 30, AWS crossed another milestone.
Amazon WorkSpaces for AI agents went to general availability, according to available sources.
The principle is striking: agents piloting legacy desktops, enterprises' old workstations.
Rather than rewriting historical applications, agents are sent in to use them like a human.
It is a pragmatic answer to an immense problem: decades of software never modernized.
Automation stumbled on these applications; agents piloting desktops unblock it.
It is perhaps the Summit's most concretely transformative announcement for CIOs.

The chosen architecture rests on an AWS-managed MCP endpoint.
MCP, the protocol connecting agents to tools, becomes here an interface to the desktop.
Authentication goes through IAM, AWS's identity system, anchoring the solution in existing governance.
Audit is ensured via CloudTrail and CloudWatch, the cloud's traceability tools.
MCP tool forwarding lets agents invoke tools across the infrastructure.
Real-time control gives operators the ability to supervise, even interrupt.
And domain-joined fleets integrate these desktops into the enterprise directory.
Each brick answers a classic objection from security departments.

For it is governance that will make or break this offering.
Entrusting enterprise desktops to autonomous agents is an idea that makes CISOs shudder.
AWS understood this and stacked the guarantees: IAM, audit, real-time control.
The message is: the agent is powerful, but it remains under watch and under control.
It is the condition for large enterprises to dare take the plunge.
Without this trust layer, agents would remain confined to demos.
With it, the entire enterprise legacy becomes automatable.
The economic stake is colossal, matching the productivity deposits at play.

Taken together, AgentCore and WorkSpaces sketch AWS's vision for the agentic era.
On one side, a modern, agnostic runtime for building cloud-native agents.
On the other, a bridge to the past, to automate the existing without rewriting it.
It is a two-speed strategy, typical of mature infrastructure players.
It tells customers: whatever your starting point, we have the path.
And it positions AWS as the backbone of the agentic enterprise, from legacy to greenfield.
In June 2026, the cloud ceased to be a mere host to become an agent operator.

The names themselves, CreateHarness and InvokeHarness, tell a story.
Create a harness, then invoke it: it is the lifecycle of a runtime taking shape.
The vocabulary is that of software engineering, not AI research.
It says the agent has become an object one manufactures, deploys, and calls.
This banalization of vocabulary is the sign of a technology entering industry.
We no longer talk of "training" or "aligning," we talk of "creating" and "invoking."
It is the grammar of the agentic era, and AWS writes it first.

A/B testing applied to agents is a small methodological revolution.
Testing two versions of an agent in production, measuring, comparing, deciding: it is web rigor applied to AI.
Until now, evaluating an agent was often impressionistic, a few well-chosen demos.
A/B testing industrializes evaluation and makes it enforceable in organizations.
It is also an answer to models' non-determinism problem: we no longer prove, we measure.
For CIOs, it is a decisive argument: agents can be managed like everything else.
AWS thus brings agents the methods that made modern software successful.

The managed Knowledge Base answers enterprises' most pressing need.
An agent without access to the enterprise's documents is an intern without a file: likeable but useless.
Feeding agents with internal, up-to-date, controlled knowledge is the real deployment challenge.
By offering it as a managed service, AWS considerably lowers the entry barrier.
Web Search grounding completes the picture by anchoring answers in verifiable sources.
It is the fight against hallucinations waged with infrastructure means.
Knowledge becomes a service, like compute or storage before it.

AWS Continuum, in preview, sketches the hardest construction site: agent security.
An agent that acts autonomously, that clicks, buys, modifies systems, must be framed.
Traditional security models, designed for humans or deterministic services, are insufficient.
Continuum promises "agentic" security, that is, designed for autonomous actors.
The fact that it is a preview, not a GA, is a welcome admission of humility.
Nobody yet knows exactly what large-scale agent security looks like.
AWS opens the construction site in public, and that is already a strong signal.

