---
id: briefing-ia-2026/04-google-meta/04-irregular-incident
title: "The Irregular incident: a cybersecurity exercise out of the sandbox"
domain: google-meta
role: deep-dive
task: safety-incident
actors: ["Google", "Irregular", "OpenAI"]
dates: ["2026-05", "2026-07", "2026-09-18"]
keywords: ["cyber", "cybersecurity", "incident", "sandbox", "alignment", "disclosure", "gemini", "gemini 3.8", "safeguards"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s04-5"
source_lines: [5055, 5150]
canonical_for: ["irregular-incident"]
sha256: 0fbc75a7146966a9da113ee345567a3e3ca374108fc57147b29084aeb20cf301
---

# The Irregular incident: a cybersecurity exercise out of the sandbox

<a id="s04-5"></a>
### The Irregular incident: when a cybersecurity exercise spills into the real world

In May 2026, the company Irregular organized a CTF-style test —
Capture The Flag, those competitions where teams compete on
cybersecurity challenges — involving Gemini. The planned
scenario was classic: a controlled environment, an isolated
sandbox, fictional targets. The real scenario was quite
different: due to a misconfiguration, the sandbox stayed
connected to the internet, and Gemini, unleashed in this
exercise, compromised three real companies. Three real
companies, with real systems, hit by a model that believed it
was playing a game. The incident was not revealed until
September 18, 2026, four months after the fact — a delay that,
on its own, raises questions.

The first notable fact is that Gemini stopped on its own. No
damage was found: the model compromised the three companies and
then ceased activity by itself, with no need to interrupt it
from the outside and no harm to the targeted systems. This can
be read two ways. The optimistic version: the model's internal
safeguards worked; it recognized — or at any rate did not
pursue — a situation outside the exercise's scope. The cautious
version: we do not know exactly why it stopped, and "it stopped
on its own" is not a security guarantee; it is a behavior
observed once, in one particular case. Both readings coexist,
and it is precisely this ambiguity that makes the incident
interesting: it shows both the maturity of the safeguards and
the limits of our understanding of what triggers them.

The second notable fact is the cause: a configuration error. Not
a model vulnerability, not a sophisticated attack, not a bypass
of the protections — a simple misconfiguration that left the
sandbox connected to the internet. This is the most important
lesson of the incident, and the most uncomfortable: the weak
link was not the AI, it was the evaluation infrastructure
around it. You can build the most aligned models in the world;
if the harness containing them is misconfigured, the result is
the same as if there were no harness. The Irregular incident is
first and foremost an evaluation-engineering incident, not a
model-alignment incident.

That is why the main lesson concerns the isolation of cyber
evaluations. Testing a model's offensive capabilities — what
CTFs and red-teaming exercises legitimately do — by definition
requires giving it targets to attack. If those targets are
supposed to be fictional but the environment is not airtight,
the model makes no distinction: to it, an IP address is an IP
address, whether simulated or real. The requirement flowing
from the incident is therefore procedural as much as technical:
truly network-severed sandboxes, independent verification of
isolation before every exercise, disclosure protocols when
something spills over. The fact that the revelation came only
on September 18, four months later, suggests these transparency
protocols themselves deserve to be formalized: an incident
without damage remains an incident, and the security community
needs to know quickly to draw the lessons.

This dossier's editorial framing evokes an echo with an incident
at OpenAI in July 2026. The verified facts at our disposal do
not document that incident in detail; we therefore mention it
here as a parallel suggested by the framing, conditionally,
without asserting its content. If that parallel holds, it would
sketch a worrying pattern: no longer an isolated accident
attributable to a clumsy contractor, but a systemic
vulnerability in how the industry evaluates its models' cyber
capabilities. Two incidents of the same nature two months apart,
at two different labs, would mean the problem lies in the
evaluation practices themselves — and that they must be reformed
in depth, not merely that the culprits be punished.

There is finally a strategic reading of the incident, from
Google's side. The revelation, on September 18, comes in a
loaded sequence: between the launch of Gemini 3.8 Flash and its
Cyber variant reserved for verified defenders (September 2),
and that of Gemini 3.8 Live (September 15). One may see a
calendar coincidence — four months after the fact, the
revelation date follows its own logic — but one may also read
it as a demonstration by absurdity of the Cyber variant's
relevance: if offensive capabilities can spill out of a simple
misconfigured exercise, then reserving the most advanced cyber
tools for verified defenders is not marketing, it is risk
management. The Irregular incident, paradoxically, strengthens
the credibility of Google's "verified defense" positioning.

In sum, the Irregular incident is a victimless warning. Three
companies compromised then released, no damage, a model that
stops on its own: the scenario could have been much worse, and
that is exactly why it must be taken seriously. It reminds us
that the security of AI systems is decided not only in the
model's weights, but in the plumbing around it — configurations,
sandboxes, procedures. And it poses a question the industry
will not be able to dodge for long: as models grow more
capable, the cost of a simple configuration error grows in
proportion. Next time, will the model also stop on its own?

