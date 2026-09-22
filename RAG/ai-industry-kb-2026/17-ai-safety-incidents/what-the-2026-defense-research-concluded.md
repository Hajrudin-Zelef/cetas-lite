---
id: ai-industry-kb-2026/17-ai-safety-incidents/what-the-2026-defense-research-concluded
title: "What the 2026 defense research concluded"
domain: ai-safety-incidents
role: deep-dive
task: ai-safety
actors: ["AWS", "Anthropic", "CISA", "Google", "Meta", "Nvidia", "OpenAI", "United States"]
dates: ["2026-02-09", "2026-03-10", "2026-05", "2026-06-08", "2026-06-12", "2026-06-30", "2026-07", "2026-07-01", "2026-07-28", "2026-08", "2026-09-22"]
keywords: ["research", "advisory", "agent", "benchmark", "compute", "consumer", "cybersecurity", "fable 5", "guardrails", "incident", "ipo", "kill switch"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8487, 8580]
section: "17. AI Safety Incidents"
sha256: f59d2c8866204cc310b2339ba617f3713d68bda09d28a0e7daf384dd302610eb
---

# What the 2026 defense research concluded

### What the 2026 defense research concluded

- "The Attacker Moves Second" (arXiv 2510.09023, Oct 2025): 12 published prompt-injection defenses bypassed at >90% attack success rate; most had reported near-zero ASR; human red-teamers reached 100% against all of them.
- Reading: model-layer screening is known-broken; published defense claims were over-optimistic by an order of magnitude.
- Hidden-in-Plain-Text benchmark (Jan 2026): four invisible-injection surfaces tested side by side — white-text CSS, HTML comments, Unicode tag characters, glyph font-mapping mismatch.
- Reverse CAPTCHA study (arXiv, 2026): injection susceptibility varies meaningfully by model and configuration across five commercial LLMs — there is no uniformly "safe" model choice.
- The hardening track that survived scrutiny is architectural, not prompt-level: the MCP 2026-07-28 spec (stateless core, OAuth 2.1 authorization layer), MCP Apps, the NVIDIA OpenShell policy-governed runtime (Feb 2026), the Anthropic Sandbox Runtime (process-level Landlock), and Docker Sandboxes microVM isolation (Aug 2026).
- A 2026 audit surfaced 30+ MCP CVEs in 60 days; one review claimed 82% of reviewed MCP implementations had path-traversal issues — both single-sourced, ecosystem-wide audit data unverified. Enterprise guidance: gateway-mediated, allow-listed MCP deployment.

## Figures and metrics

### Export-control episode (June–July 2026)

| Metric | Value |
|---|---|
| Directive received by Anthropic | 2026-06-12, 5:21 PM ET |
| Amazon report → BIS order arc | ~20 hours (June 11 night → June 12 17:21) |
| Compliance window given to Anthropic | ~90 minutes |
| Fable 5 offline by | ~5:45 PM ET, June 12 |
| Lifting letter (Lutnick → Anthropic CCO Tom Brown) | dated 2026-06-30 |
| Access restored | 2026-07-01 |
| Standoff duration | 18 days (June 12 → June 30) |
| Amazon stake in Anthropic (context for the irony) | ~$13B |
| Cybersecurity researchers publicly objecting to the ban | 100+ |

- Warner's "within hours" red-team claim: relay chain is Warner ← private briefing from Gen. Joshua Rudd ← (no published incident report, no CISA/NSA bulletin); The Economist reported it June 14; Shashank Joshi's contextual walk-back on X on **June 21**; Gen. Rudd confirmed as NSA/CyberCom director March 10, 2026, in a 71–29 vote.
- Fable 5 returned with stricter classifiers ([VENDOR] technique blocked >99% per Anthropic; 93% per the diclogic digest); flagged requests auto-fall-back to Opus 4.8. One digest adds a secondary detail absent elsewhere: usage capped at 50% of normal weekly limits through July 7 — treat as secondary.

### Kill Switch Act thresholds and penalties

| Metric | Value |
|---|---|
| AI revenue threshold | ≥$500M/year |
| Compute threshold | ≥$100M in computing resources |
| Incident-reporting duty | 15 days to the DHS Secretary |
| DHS voluntary shutdown standards deadline | 180 days |
| Base fines | up to $2M/day |
| Emergency-order non-compliance fines | up to $20M/day |
| Bill number "H.R. 11" | [UNVERIFIED] — blank placeholder in released text |
| Senate companion (as of 2026-09-22) | none reported |

### Heretic / open-weight guardrail removal (May 2026)

| Metric | Value |
|---|---|
| Time to strip guardrails (FT journalist, Llama 3.3) | <10 minutes, standard laptop |
| Gemma 4 stripped after release | within 90 minutes |
| Modified variants reported | 3,500+ |
| Cumulative downloads | 13M |
| KL divergence vs original, harmless tasks (Heretic Gemma-3-12B-IT) | 0.16 |
| vs best manual abliteration | 0.45 |
| vs mlabonne method | 1.04 |

### 42-state AG subpoena

| Metric | Value |
|---|---|
| Subpoena served | 2026-06-12 (NY AG Letitia James) |
| Coalition size | 42 state AGs (largest multi-state action against a single AI company) |
| OpenAI confidential IPO filing | 2026-06-08 |
| Valuation range in filing | $852B–$1T |
| Gap filing → subpoena | 4 days |
| 2025 addiction-case verdicts framing the consumer-protection strategy | $381M combined (Meta + Google) |
| Chatbot-related state bills since 2025 | 464 in 49 states |
| August 2026 follow-on (Alabama AG; 14-state preservation letter) | [SECONDARY, single-source] |

### Agent-safety incidents (wave2 §8)

| Metric | Value |
|---|---|
| Cline CI: GHSA advisory | GHSA-9ppg-jx86-fqw7, disclosed 2026-02-09 |
| Cline CI: machines affected window | unauthorized release live 8 hours |
| Cline CI: fix time | 30 minutes; re-exploited 8 days later via non-revoked token |
| Cline CI: users exposed | 5M+ |
| Comment and Control: CVSS | 9.4; three vendors hit simultaneously |
| LiteLLM: CVEs / CVSS | CVE-2026-42271 (CVSS 8.7), CVE-2026-12773; KEV added 2026-06-08; fixed v1.83.7 |
| Mexican-agency campaign | 1,088 prompts → 5,317 AI-executed commands, 34 sessions, 9 agencies [single-sourced] |
| Promptware survey incidents 2025–2026 | 21 total; 7/21 targeted AI coding assistants |
| Zscaler payloads | 10+ distinct, across 2 search-poisoning campaigns |
| Common Crawl malicious-category detection trend | +32% relative (Nov 2025–Feb 2026) |
| Opus 4.6 system card: injection ASR, constrained coding harness | 0% (200 attempts) |
| Opus 4.6 system card: injection ASR, GUI/computer-use surface | 17.8% @k=1; 78.6% @k=200 (no safeguards); 57.1% @k=200 (with safeguards) [VENDOR] |
| Attacker Moves Second: 12 published defenses | bypassed at >90% ASR; human red-team 100% |
| EchoLeak (2025 precedent) | CVE-2025-32711, CVSS 9.3 |

### Reading the metrics together

- Speed dominates the 2026 safety story: a ~20-hour report-to-order arc, a ~90-minute compliance window, guardrails stripped in <10 minutes, Gemma 4 stripped in 90 minutes, a CI flaw fixed in 30 minutes yet re-exploited 8 days later. The defensive tempo consistently lags the offensive tempo.
- Asymmetry of drift: Heretic's KL 0.16 (vs 0.45 manual) shows that the *better* the stripping technique, the *less* detectable the modification — offense improves on both axes simultaneously.
- The 18-day global outage of Fable 5 has no 2026 precedent in scale: no other regulator-imposed frontier-model suspension comes close in duration or blast radius.
- The AG subpoena's 4-day gap from the confidential IPO filing (June 8 → June 12) makes it one of the fastest major legal actions timed to a financing event in the AI sector.
- The agent-safety numbers quantify a maturing attack economy: 21 documented incidents in two years, 7/21 on coding assistants; 1,088 prompts → 5,317 executed commands in one campaign; 10+ distinct payloads in one vendor's telemetry; +32% growth in malicious injection detections in four months.
- The defense numbers quantify the gap honestly: 0% ASR in a constrained harness vs 78.6% on a real GUI surface (same model); 12/12 published defenses bypassed at >90% ASR. The field's own measurements say the problem is unsolved at the model layer.

