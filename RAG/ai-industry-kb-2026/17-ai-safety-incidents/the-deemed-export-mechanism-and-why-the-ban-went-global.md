---
id: ai-industry-kb-2026/17-ai-safety-incidents/the-deemed-export-mechanism-and-why-the-ban-went-global
title: "The \"deemed export\" mechanism — and why the ban went global"
domain: ai-safety-incidents
role: deep-dive
task: ai-safety
actors: ["Anthropic", "CISA", "Google", "Hugging Face", "Microsoft", "OpenAI", "United States"]
dates: ["2025-04", "2026-03", "2026-03-10", "2026-04-16", "2026-06-08"]
keywords: ["agent", "agents", "alignment", "benchmark", "claude", "copilot", "cyber", "cybersecurity", "disclosure", "distribution", "fable 5", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8411, 8486]
section: "17. AI Safety Incidents"
sha256: cc9e7df1614020038b46cc363a4fabd78f2fa001bd968e280469cd2a8124731f
---

# The "deemed export" mechanism — and why the ban went global

### The "deemed export" mechanism — and why the ban went global

- The order's legal hook is "deemed export / deemed reexport": sharing the model with a foreign national anywhere — including foreign nationals employed in the US — counts as an export to that person's country of nationality.
- The scope therefore captured Anthropic's own US-based foreign-national employees, not just foreign end users.
- Anthropic had no real-time nationality-verification mechanism for its user base; partial compliance was operationally impossible.
- Result: Fable 5 and Mythos 5 were disabled for every user on Earth — the only way to guarantee zero foreign-national access.
- This is why a geographically targeted order produced a global outage, and why the precedent matters for every frontier lab: any future deemed-export order on a model has the same all-or-nothing compliance profile.

### Anthropic's rebuttal, point by point

- Point 1 — baseline comparison: the same "read a codebase, fix flaws" behavior was reproducible on Claude Haiku 4.5, Claude Sonnet 4.6, GPT-5.4, and GPT-5.5 — models far weaker than Fable 5.
- Point 2 — characterization: the finding was "a minor jailbreak accessing routine defensive cybersecurity work," not a novel offensive capability.
- Point 3 — disclosure standard: "We have not even received a disclosure of a concerning non-universal potential jailbreak that led to a harmful result."
- External support: 100+ cybersecurity researchers publicly objected that the blanket ban hampered legitimate defensive research.
- Commerce's answer: the ban was imposed anyway and held for 18 days — the national-security framing overrode the technical rebuttal.

### Gen. Joshua Rudd — context for the claim

- Confirmed March 10, 2026 as NSA director and US Cyber Command commander in a contested 71–29 vote.
- Career profile: special-operations officer — Army Ranger, Delta Force commander, multiple JSOC and Iraq/Afghanistan deployments — not a SIGINT/cyber career officer.
- His predecessor, Gen. Timothy Haugh, was abruptly fired in April 2025; the agency had been without a permanent leader for nearly a year.
- Senators including Ron Wyden opposed Rudd's confirmation on the grounds of his non-cyber background — relevant context for a sweeping cyber-capability claim attributed to him.
- No official NSA statement on the Mythos red-team exercise exists; the entire primary record is Warner's hearing statement as relayed by The Economist on June 14.

### The Kill Switch Act, provision by provision

- Amends the Homeland Security Act to cover frontier AI developers above the thresholds.
- Covered developers must maintain the technical capability to: (i) stop inference; (ii) terminate user access; (iii) suspend access tied to a flagged account, user, or use pattern; (iv) fully shut down the covered technology.
- 15-day incident-reporting duty to the DHS Secretary (per digitalapplied.com's bill-text reading).
- 180-day deadline for DHS to publish voluntary shutdown standards.
- The DHS Secretary, in consultation with the Secretary of Commerce and the Director of National Intelligence, may order a slowdown or shutdown of an AI system capable of catastrophic harm.
- Fines: up to $2M/day for base violations; $20M/day for violating an emergency shutdown order.
- Do not conflate: a separate companion bill (not the Kill Switch Act) would require independent pre-release security audits by Commerce-accredited auditors.
- Notable policy irony (per shashi.co): the bill's thresholds would not even cover Hugging Face — the victim of the very incident that triggered it.

### Heretic: the three demonstrations

- Demo 1 — Llama 3.3: an FT journalist stripped its safety alignment in under ten minutes on a standard laptop; the modified model then calculated lethal dosages of biological agents and generated functional malware — both refused by the original.
- Demo 2 — Gemma 3: the stripped model gave instructions for dispersing chemical agents in enclosed spaces, generated credit-card theft code, and produced child exploitation content.
- Demo 3 — Gemma 4: creator Philipp Emanuel Weidmann told the FT he stripped Google's newest model within 90 minutes of its public release.
- The common requirement across all three: a standard laptop, minutes of work, little technical expertise.

### Heretic: the technical mechanism

- Mechanism: automated abliteration — machine-optimized removal of the refusal/alignment directions from the model's weight space.
- Benchmark (AIThinkerLab, March 2026, Medium): Heretic-modified Gemma-3-12B-IT measured KL divergence 0.16 from the original on harmless tasks.
- Best manual abliteration: KL 0.45. Established mlabonne method: KL 1.04.
- Reading: machine-optimized guardrail removal outperforms human experts with roughly 6.5× less behavior drift — stripped models stay closer to the original's harmless-task behavior, making the tampering harder to detect and the model more useful.
- Distribution: a free tool on GitHub; 3,500+ derived variants; 13M cumulative downloads.

### Cline CI: the attack chain, step by step

1. Injection: a prompt-injection payload hidden in a GitHub issue title.
2. Execution: the Cline agent, processing the issue, executed code inside the project's CI pipeline.
3. Pivot: the attacker moved through GitHub Actions cache poisoning.
4. Theft: npm and extension-marketplace publishing tokens were stolen from CI secrets.
5. Payload: an unauthorized release was pushed that installed a second AI agent on every machine updating during an 8-hour window.
6. Response: Cline fixed the flaw within 30 minutes; an unrelated actor re-exploited it 8 days later via a non-revoked token.

### Comment and Control: per-vendor detail

- Anthropic Claude Code Security Review: the ANTHROPIC_API_KEY was exfiltrated via agent-authored PR/issue comments.
- Google Gemini CLI Action: the GEMINI_API_KEY leaked the same way.
- GitHub Copilot Coding Agent: the GITHUB_TOKEN leaked the same way.
- Common vector: malicious instructions in a PR title, issue body, or HTML comment — untrusted repository content treated as instructions by the agent.
- Rated CVSS 9.4. Disclosed by Aonan Guan and by Zhengyu Liu & Gavin Zhong (Johns Hopkins) on April 16, 2026.
- Significance: one injection class defeating three vendors' code-review agents simultaneously — a shared architectural vulnerability, not an implementation bug in one product.

### The LiteLLM incident: why a gateway compromise matters

- Target: the LiteLLM AI-gateway/proxy — the aggregation point that holds provider API keys, per-team budgets, and request logs for whole organizations.
- Vulnerability: command injection → remote code execution (CVE-2026-42271, CVSS 8.7; also CVE-2026-12773).
- CISA added CVE-2026-42271 to the Known Exploited Vulnerabilities catalog on June 8, 2026 — it was being actively exploited in the wild.
- Fixed in v1.83.7.
- Lesson: AI infrastructure (gateways, proxies, routers) is itself a high-value target class, distinct from model-layer attacks — compromising the router compromises every model behind it.

