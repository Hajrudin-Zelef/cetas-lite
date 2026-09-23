---
id: collect-korben/korben/cloudflare-security-audit-skill
title: "Cloudflare lance son skill d'audit de sécurité par IA"
domain: korben
role: reference
task: article
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "claude", "exploit", "guardrails", "license", "memory", "mit license", "sandbox"]
source: docs/RAG/Collect RAG/01_korben/cloudflare-security-audit-skill.md
source_anchor: ""
source_lines: [1, 60]
sha256: 624111e24d7e519e4455b57bd887cc44e5a2c924b69ca783de0fec008d61c3b0
---

# Cloudflare lance son skill d'audit de sécurité par IA

## Metadata

- **Source** : https://korben.info/cloudflare-security-audit-skill.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article presents Cloudflare's "security-audit-skill", an AI-powered security audit skill published on GitHub. It is described as a package that turns a code agent into a real auditor rather than a simple code reviewer. It is a folder of Markdown instructions dropped into a coding agent.

The repository contains one file per attack family: memory corruption, prompt injection, HTTP request framing, and tenant isolation. Alongside these, a JSON schema describes what a report should look like, and two validators written in JavaScript with no dependencies at all.

The agent runs through six phases, from reconnaissance to the final report. It starts by mapping the architecture, trust boundaries and entry points, then writes its own coverage grid. It then sends isolated "hunters" case by case, and each lead that surfaces goes to a fresh verifier that did not participate in the hunt. Each lead concludes with a technical report. A `confirmed` report requires a complete source trace and a genuinely observed result. A `needs_validation` report names the precise missing fact and, importantly, has no right to any severity level. The `rejected` report keeps a memory of what was refuted, so the next pass does not resurface the same thing.

The validator is real code that checks the "form" of the evidence. Fingerprints must be unique, and the test path must run from an entry point to an impact point. Thus a confirmed finding whose path does not hold up is rejected. The independence of the verifier is requested in the instructions but not enforced (at Cloudflare it is their in-house orchestrator that handles it, and it is not publicly available).

Cloudflare explains what one of these agents does when unleashed on code without guardrails: it modifies the source so its exploit works, then proudly announces the bug it just created itself. Or it produces a test demonstrating that exec() executes things, therefore concluding it is necessarily a critical flaw. These instructions are designed to forbid exactly that behavior.

To install it, the Vercel Labs skills CLI is used: `npx skills add https://github.com/cloudflare/security-audit-skill --skill security-audit`. The agent is then launched in the repository to audit and asked for a security audit; the skill triggers automatically and the agent's model drives sub-agents in parallel. By default the report lands in a `~/security-audit-skill` folder, outside the audited repository. The skill limits mistakes: it refuses to execute audited code without an OS-enforced sandbox with network cut, read-only target and CPU/memory limits. The repository does not provide this sandbox; users must set it up themselves, for example with an isolation tool like Fence. Without it, the audit lead ends in `needs_validation` rather than being truly followed through and decided.

Cloudflare notes that a repository of about 30,000 lines takes only 3 to 4 hours to audit with their large internal system for a correct result, but locally it will be much longer. A single pass is not enough, finding roughly half the flaws each time, so it must be relaunched several times. The repository is under MIT license and requires Node.js for the two validators. Anthropic also maintains a security reviewer for Claude Code that reads the diff and comments directly on it.

## Key points

- Cloudflare published "security-audit-skill" on GitHub, an AI security-audit skill built from Markdown instruction files.
- One file per attack family: memory corruption, prompt injection, HTTP request framing, tenant isolation.
- Includes a JSON report schema and two dependency-free JavaScript validators.
- Six phases from reconnaissance to final report; isolated hunters per lead; a fresh independent verifier per finding.
- Three verdicts: `confirmed` (full source trace + observed result), `needs_validation` (no severity allowed), `rejected` (kept to avoid repeats).
- Validators check evidence form: unique fingerprints and a test path from entry point to impact.
- Requires an OS-enforced sandbox (network cut, read-only target, CPU/memory limits); otherwise leads stay `needs_validation`.
- MIT license, needs Node.js; at Cloudflare ~30,000 lines audited in 3-4 hours, but one pass finds only ~half the flaws.

## Technical data / figures

| Item | Value |
|---|---|
| Project | Cloudflare security-audit-skill |
| Repository | github.com/cloudflare/security-audit-skill |
| License | MIT |
| Attack-family files | Memory corruption, prompt injection, HTTP request framing, tenant isolation |
| Validators | 2 JavaScript validators, no dependencies |
| Report schema | JSON |
| Phases | 6 |
| Verdicts | confirmed, needs_validation, rejected |
| Install CLI | Vercel Labs skills CLI (`npx skills add ...`) |
| Default report path | ~/security-audit-skill |
| Requirement | Node.js; OS-enforced sandbox (e.g., Fence) |
| Cloudflare audit time (30k lines) | 3-4 hours |
| Flaws found per pass | ~50% |
| Related tool | Anthropic Claude Code security review |

## Why this source matters for the RAG

This article provides a detailed technical walkthrough of an open-source AI security-audit skill, including its architecture, verdict system, sandbox requirements and performance figures. It is highly useful for questions about AI-assisted penetration testing, agent guardrails, and secure code auditing workflows.
