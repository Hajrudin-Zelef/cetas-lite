---
id: collect-240926-korben/korben/cloudflare-lance-son-skill-d-audit-de-sacurita-par-ia-korben
title: "Cloudflare launches its AI security audit skill"
domain: korben
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "cybersecurity", "exploit", "guardrails", "license", "memory", "mit license", "sandbox"]
source: docs/RAG/clean_en/korben/cloudflare-lance-son-skill-d-audit-de-sacurita-par-ia-korben.md
source_anchor: ""
source_lines: [1, 58]
sha256: d661b940c5cb85bbc316481f2880747c845b0dc04f24ecdf99fd9d8661a9d5cf
---

# Cloudflare launches its AI security audit skill

<!-- source: https://korben.info/cloudflare-security-audit-skill.html -->

# Cloudflare launches its AI security audit skill

## What to remember AI-generated summary

1. Cloudflare publishes on GitHub an AI security audit skill, based on Markdown instruction files per attack family, plus a JSON schema and two JavaScript validators with no dependencies.
2. The agent runs through six phases, from architecture mapping to the final report, with isolated hunters per lead and an independent verifier who did not participate in the hunt.
3. At Cloudflare, auditing a repository of about 30,000 lines takes 3 to 4 hours via their internal system, but a single pass only finds about half of the vulnerabilities.

Cloudflare has put online on its GitHub a super cool security audit skill that served as their starting point for their own security vulnerability hunting system. This package for doing AI-agent-boosted security audits is a folder of Markdown instructions that you drop into your code agent, and that makes it work like a real auditor and not just as a code "reviewer."

In this repo, you'll find one file per attack family: memory corruption, prompt injection, HTTP request framing, tenant isolation. And with that, a JSON schema that describes what a report should look like, and two validators written in JavaScript without the slightest dependency.

*The skill folder on GitHub: one instruction file per attack family, and the two validators at the bottom of the list (
Source
)*

The agent then runs through six phases, from reconnaissance to the final report. It starts by mapping the architecture, trust boundaries, and entry points, then writes its own coverage grid. Then it sends isolated hunters case by case and each lead that comes up goes to a fresh new verifier who did not participate in the hunt.

Each lead then ends with a technical report. A `confirmed` report that requires a complete source trace and a truly observed result. A `needs_validation` report that names the precise fact that is missing, and above all that is not entitled to any severity level. And the `rejected` report that keeps the memory of what has been refuted, so that the next pass doesn't bring you out the same thing again.

The validator, for its part, is real code that verifies the "form" of the evidence. The fingerprints must be unique, and the tracking of the test path that must start from an entry point to end at an impact point. Thus a confirmed finding whose path doesn't hold up is rejected.

However, be aware, the independence of the verifier is required in the instructions, but not controlled (at Cloudflare, it's their in-house orchestrator that handles it, and it's not publicly available).

Cloudflare explains on its site what one of these agents unleashed on code without guardrails gives. It modifies the source so that its exploit works, then proudly announces the bug it just created itself. Or else it churns out a test that demonstrates that exec() executes things, so that it's necessarily a critical vulnerability... And that, you'll have understood, is crap and is absolutely what we don't want.

So here is exactly what these instructions seek to forbid it from doing.

To install it, you need the Vercel Labs skills CLI:

```
npx skills add https://github.com/cloudflare/security-audit-skill --skill security-audit
```
Then, you launch your agent in the repo to be audited and you ask it for a security audit. The skill triggers on its own and it's your agent's model that will drive the sub-agents in parallel (so use a model that knows how to do that, it's better lool). By default, the report then lands in a ~/security-audit-skill folder, outside the audited repo.

What's cool is that this skill is designed to limit mistakes. For example it refuses to execute the code it audits if it doesn't have a sandbox imposed by the operating system, with the network cut, the target in read-only, and CPU and memory limits. This sandbox, the repo doesn't provide it, and it's up to you to set it up with for example
an isolation tool like Fence
. Without it, the audit lead will quite simply end up in `needs_validation` instead of being truly followed up and decided.

After that, it's time that will make the difference. At Cloudflare, they have a big system that allows them, for example with a repo of about 30,000 lines of code, to spend only 3 to 4 hours on the audit to get a correct result. But at your place, locally, it will be much longer. Cloudflare also warns that a single pass is not enough since it finds roughly half of the vulnerabilities each time. You'll have to run it several times.

Come on, if you're also looking for a safety net for everything concerning your pull requests, know that Anthropic maintains on its side a security reviewer for Claude Code that reads the diff and comes to comment directly on it.

Cloudflare's Security Audit Skill repo is under MIT license, and it will need Node.js to run its two validators.

In any case, what I advise you is to go read the SKILL.md before launching the first pass, to make sure everything will be OK within your harness.

Source: security-audit-skill on GitHub

Entirely dedicated to cybersecurity, the Guardia school is accessible either directly after the baccalaureate (post-bac), or after a bac+2 or bac+3. By joining the Guardia school, you will become a computer developer with a cybersecurity option (Bac+3) or a cybersecurity expert (Bac+5).

Guardia CS also trains professionals in cybersecurity via several online courses

## Comments

starfix!in Surfshark doesn't make you invMorganein Discord guesses your age sansts3rv1in The Ray-Ban Display arrive eponpondin Openpilot - The NHTSA passes lesfabiendin Claude Code makes you choose
