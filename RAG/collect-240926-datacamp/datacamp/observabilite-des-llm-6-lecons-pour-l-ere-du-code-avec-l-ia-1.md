---
id: collect-240926-datacamp/datacamp/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia-1
title: "observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "guardrails", "incident", "kill switch"]
source: docs/RAG/clean_en/datacamp/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia.md
source_anchor: ""
source_lines: [1, 116]
sha256: a8803055013b0ae4f702e58dee23cdc23973247463a4e2e35780868e09eb81d3
---

# observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia

<!-- source: https://www.datacamp.com/fr/blog/llm-observability-six-lessons-from-datadog-cto -->

Cursus

Engineering teams today ship more code than they can read. AI assistants now write much of it, faster than any reviewer can follow line by line. This shift serves as the backdrop for Datadog's DASH conference in New York this week, where co-founder and CTO Alexis Lê-Quôc is hosting a session titled "The New Shape of Engineering."

His observation is simple: the way software is operated hasn't changed: you ship a change, you deploy it, then you observe. What's changing is the volume and the cadence — and that changes what keeps the whole thing safe.

In this article, I synthesize his thinking into six key lessons: the evolution of code review, production as the ultimate test, and what you should take away from it.

If you're new to LLM observability, we recommend reading our guides to get started with MLOps and LLM evaluation as a starting point.

## In brief

Lê-Quôc's throughline: observability becomes the control layer for software written, tested, and deployed by AI, serving both operators and the agents themselves.

The six lessons, summarized:

- **Review moves away from the code itself.** There's too much AI-generated code to read line by line; the real control is the tests, specifications, and evidence you design upfront, while also building in guardrails against agents that might try to game those tests.
- **Production is the only test that matters.** An all-green CI pipeline doesn't prove much against real-world usage you couldn't anticipate, and a model's output is never entirely certain; so you need to monitor it live and keep a kill switch.
- **Let agents take on the drudgery.** Entrust them with monitoring dashboards and verifying hypotheses that exhaust humans, and keep people for high-judgment decisions.
- **Split the work into two loops:** a development loop (write, ship, verify, fix) and an operations and security loop (detect, investigate, resolve).
- **Get AI spending under control.** Size which model does which task based on agent trajectories, and leave that decision to the developers and SREs who make it.
- **Learn how to learn.** Models are tireless tutors, but the skill is in questioning them: understanding systems layer by layer and asking why the code they wrote actually worked.

## Level up your team's AI skills

Transform your business by equipping your team with advanced AI skills through DataCamp for Business. Sharpen your knowledge and your efficiency.

## Lesson 1: AI has shattered the old code review

Let's start with the pressure that shapes everything else: there's more code than any human can read.

Lê-Quôc is direct: the historical model, a human reading a pull request line by line, doesn't hold up to AI-assisted development. The concern he hears across the industry is about the impossibility of review: too much is happening to keep up by reading PRs.

His answer isn't to ask people to read faster, but to move the review elsewhere.

Review is no longer at the line of code; there's too much, you can't keep up. It's about the tests we design upfront, and forbidding the agent from circumventing them.

Alexis Lê-Quôc, CTO at Datadog

That last point is easy to miss. Once you orchestrate one agent to plan, another to write, and a third to test, you also need to prevent the writer from tampering with the automated tests instead of solving the problem.

He goes beyond tests. Datadog now adds semi-formal and formal proofs that the specification actually does what it's supposed to do — an approach too costly to generalize before agents took on the bulk of the work. It works particularly well on backend and coordination systems, where behavior is mathematical enough to reason about precisely.

## Lesson 2: Production is the only test that matters

Passing all tests in CI is necessary, and very far from sufficient. The failures that matter come later.

Where it really matters is in production.

Alexis Lê-Quôc, CTO at Datadog

Every delivery rests on assumptions that can't be fully verified upfront: the shape of the data and user behavior. Expose those assumptions to enough real traffic, and rare cases stop being rare: they become the everyday slowdowns and errors of data and model drift.

LLMs make things even more complicated: with classic code, you can at least reason about each branch. No one can mechanically explain why a model returns a given answer; the same input never guarantees the same output. Occasional strange results can't be eradicated through engineering.

So you stop trying to prove a system is correct before deployment. Instead, you:

- Write evaluations of expected behavior
- Monitor it in production
- Keep a kill switch for a deployment that goes wrong.

The question is no longer whether it succeeded, but whether a problem is isolated or the start of a trend.

That live signal isn't just a dashboard for humans. Wired into the deployment system, it lets an agent roll out a change the way a cautious engineer would: 1% of users, then 5%, judging on real data whether the change produces the expected effect.

## Lesson 3: Let agents take on the drudgery

For Lê-Quôc, agents don't replace engineers: they take on the tasks that wear people down.

Diagnosing an incident means testing hypotheses against a symptom, and during long incidents, it's often the unlikely hypothesis that turns out to be right. Datadog's Bits AI agent checks them all in parallel, ahead of the engineer, while the person steers it toward the intuition no dashboard would surface.

The heart of the matter is fatigue. An on-call deployment is alert spikes followed by hours of nothing, repeated until judgment erodes.

You're in maximum alert mode, then you watch paint dry.

Alexis Lê-Quôc, CTO at Datadog

An agent doesn't get fazed by it and doesn't decline after four hours of staring at numbers. Stress and fatigue degrade human performance, which is why people rotate through on-call shifts.

Entrust the tireless watch to a machine, and teams come back rested for the decisions that matter. Same logic in security, where analysts burn out sorting false positives from real threats.

## Lesson 4: split the work into two loops

Lê-Quôc organizes agent work at Datadog around two loops.

### The development loop

The first loop will be familiar to most engineers:

1. Write code
2. Ship it
3. Check whether it works
4. Fix
5. Repeat

Datadog's angle: a problem born in the code is generally fixed in the code. The platform therefore tries to offer you that fix, based on what it knows about the application: owners, recent changes, reported errors.

He cites database query optimization as an example. Any model can rewrite a slow query; the hard part is proving the rewritten version is faster and safer before it reaches production. Datadog therefore tests it on a realistic copy of production data, then provides a pull request complete with the evidence.

### The operations and security loop

The other loop runs in parallel, among the same people or a different team:

1. Detect
2. Investigate
3. Fix
4. Repeat

This is where Datadog's AI Guard prioritizes security events and blocks attacks faster than an analyst could by hand. Agents can also handle routine operational tasks that engineers perform every day without enthusiasm, like resizing that famous Kubernetes pod.

In both loops, Lê-Quôc remains clear on the order of priorities. Datadog doesn't start from "here's AI, what problem can it solve?": you start from a proven customer pain point, usually a variation of "I don't want to do this repetitive task anymore," then assess whether an agent can be trusted to handle it.

## Lesson 5: get AI spending under control

