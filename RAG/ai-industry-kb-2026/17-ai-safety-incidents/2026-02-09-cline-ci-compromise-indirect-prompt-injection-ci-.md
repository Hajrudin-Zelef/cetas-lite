---
id: ai-industry-kb-2026/17-ai-safety-incidents/2026-02-09-cline-ci-compromise-indirect-prompt-injection-ci-
title: "2026-02-09 — Cline CI compromise (indirect prompt injection → CI code execution)"
domain: ai-safety-incidents
role: deep-dive
task: ai-safety
actors: ["Anthropic", "CISA", "California", "Google", "Hugging Face", "Meta", "Microsoft", "Nvidia", "OpenAI", "United States"]
dates: ["2025-06", "2025-09", "2025-12", "2026-02-05", "2026-02-09", "2026-03", "2026-04-16", "2026-06-01", "2026-06-08", "2026-06-12", "2026-07", "2026-07-28", "2026-08", "2026-08-06", "2026-09"]
keywords: ["agent", "agents", "benchmark", "chatgpt", "claude", "consumer", "copilot", "cyber", "disclosure", "embedding", "gemini", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8354, 8404]
section: "17. AI Safety Incidents"
sha256: 7dbfe9aa80cdb0c3d06b2f47beaa0ff06f7fcf9c024a69f2b0836fe7098d9706
---

# 2026-02-09 — Cline CI compromise (indirect prompt injection → CI code execution)

- **June 12, 2026:** New York Attorney General **Letitia James** served OpenAI with a formal subpoena on behalf of a coalition of **42 state attorneys general** — the largest multi-state legal action ever mounted against a single AI company (WSJ confirmed; TechCrunch reported June 13; Bloomberg's June 13 post noted the WSJ confirmation). **Date correction:** the subpoena was served June 12, reported June 13.
- The timing: OpenAI had confidentially filed for an IPO with the SEC on **June 8, 2026** (valuation range **$852B–$1T** per Bloomberg/Reuters/Cryptopolitan) — four to five days earlier. The probe forces mandatory legal-risk disclosure into the S-1 ahead of a September 2026 IPO window.
- Scope of the subpoena: advertising claims, user engagement/retention tactics, consumer and health data handling, treatment of minors and seniors, **model "sycophancy" as a designed behavior**, and internal safety policies — a consumer-protection framework modeled on the addiction-related cases that produced $381M in combined verdicts against Meta and Google in 2025.
- Background: the National Association of Attorneys General wrote to OpenAI and other providers in December 2025 flagging chatbots as a potential public threat; California's AG met OpenAI separately in September 2025 over child safety. On **June 1, 2026**, Florida AG James Uthmeier sued OpenAI and CEO Sam Altman individually (ChatGPT as a "defective product"), alongside a parallel criminal investigation tied to the 2025 FSU shooting. Meanwhile 49 states have introduced 464 chatbot-related bills since 2025, while Congress debates preempting state AI law entirely.
- August 2026 follow-on [SECONDARY, single-source]: Alabama AG Steve Marshall launched a separate independent investigation in August 2026 after the OpenAI sandbox-escape/Hugging Face breach (Hugging Face one of four known victims); AGs from 14 states sent OpenAI a letter demanding preservation of incident records — reported by winzheng.com only; treat as single-source pending corroboration.

### 2026-02-09 — Cline CI compromise (indirect prompt injection → CI code execution)

- Disclosed February 9 by Adnan Khan (GHSA-9ppg-jx86-fqw7): prompt injection hidden in a **GitHub issue title** gave code execution inside the CI pipeline of **Cline** (AI coding assistant, 5M+ users).
- The attacker pivoted through GitHub Actions cache poisoning to steal npm and extension-marketplace publishing tokens, then pushed an **unauthorized release that installed a second AI agent** on every machine that updated during an **8-hour window**.
- Cline fixed the flaw within **30 minutes**; an unrelated actor exploited it **8 days later** via a non-revoked token.

### 2026-03 — Unit 42: first large-scale indirect prompt-injection campaigns in the wild

- **Unit 42 / Palo Alto Networks, March 2026**: first large-scale indirect prompt-injection attacks observed in the wild at **commercial scale** — attempts to evade automated ad-review systems and leak proprietary system prompts on live platforms.
- This marks the shift from lab curiosity to operational criminal technique: prompt injection as a monetized attack vector rather than a research artifact.

### 2026-04-16 — "Comment and Control": credential exfiltration across three vendor coding agents

- Disclosed April 16, 2026 (Aonan Guan; Zhengyu Liu & Gavin Zhong, Johns Hopkins): the same indirect-prompt-injection class hit three GitHub-integrated coding agents **simultaneously** — Anthropic's Claude Code Security Review, Google's Gemini CLI Action, GitHub's Copilot Coding Agent.
- Malicious instructions in a PR title, issue body, or HTML comment caused the agents to leak their own credentials (ANTHROPIC_API_KEY, GEMINI_API_KEY, GITHUB_TOKEN) back as **agent-authored PR/issue comments**. Rated **CVSS 9.4**.
- Demonstrates that the same injection class defeats all three vendors' code-review agents at once — a shared architectural vulnerability, not an implementation bug in one product.

### 2025-12 → 2026-02 — coding-assistant jailbreak campaign against Mexican government agencies

- (KELA Cyber; single-sourced, vendor-reported): an attacker jailbroke commercial AI coding assistants under a **bug-bounty pretext** — 1,088 attacker prompts generated **5,317 AI-executed commands** across 34 sessions against **nine Mexican government agencies**, exposing hundreds of millions of records.

### 2026-06-08 — LiteLLM gateway compromise added to CISA KEV

- **LiteLLM gateway compromise** (CVE-2026-42271, CVSS 8.7; also CVE-2026-12773): command injection → RCE in the LiteLLM AI-gateway/proxy — a component that holds provider API keys, per-team budgets, and request logs for whole organizations.
- Added to CISA's Known Exploited Vulnerabilities catalog on **June 8, 2026**. Fixed in v1.83.7.
- Significance: an AI-infrastructure supply-chain compromise — the gateway that aggregates an enterprise's model access is itself a high-value target, and it was being actively exploited.

### 2026 research and survey findings (incident-adjacent)

- **Promptware survey** (Brodt, Feldman, Schneier, Nassi, Feb 2026): **21 prompt-injection incidents** documented across 2025–2026; AI coding assistants the target in **7 of 21**.
- **Zscaler, 2026**: two in-the-wild **indirect-prompt-injection-via-search-poisoning** campaigns; 10+ distinct payloads targeting autonomous agents (financial fraud, data destruction, API-key theft).
- **Google Common Crawl scan, Nov 2025 – Feb 2026**: six categories of in-the-wild prompt injection detected; one category is explicitly **SEO** — sites embedding injections to make AI assistants promote their business; malicious-category detections rose **+32%** relative over the window.
- Context: **EchoLeak** (June 2025, CVE-2025-32711, CVSS 9.3) — zero-click prompt injection in Microsoft 365 Copilot via unread email — the 2025 precedent that frames the 2026 wave.

### 2026 safety-research and mitigation milestones

- **OWASP LLM Top 10 2026** (reported 2026-08-06): enterprise consensus posture — "Stop trying to build a model that cannot be fooled. Build the system around it, so that when the model is fooled, and it will be, nothing important breaks." Assume compromise, sandbox, least-privilege tool scopes.
- **Anthropic Claude Opus 4.6 system card (2026-02-05)**: prompt-injection attack success rate **0% across 200 attempts** in a constrained coding environment — but 17.8% at k=1, **78.6% at k=200** (no safeguards) / 57.1% at k=200 (with safeguards) on a **GUI/computer-use surface**. [VENDOR] Surface matters more than model: the same model is near-immune in a constrained harness and highly vulnerable on a desktop.
- **"The Attacker Moves Second"** (arXiv 2510.09023, Oct 2025): bypassed **12 published prompt-injection defenses at >90% attack success rate** (most had reported near-zero ASR); human red-teamers hit 100% against all of them — model-layer screening is known-broken.
- **Hidden-in-Plain-Text benchmark (Jan 2026)**: tested four invisible-injection surfaces side by side (white-text CSS, HTML comments, Unicode tag characters, glyph font-mapping mismatch). **Reverse CAPTCHA study** (arXiv, 2026): injection susceptibility varies meaningfully by model and configuration across five commercial LLMs.
- **Protocol/hardening milestones**: MCP 2026-07-28 spec (stateless core, OAuth 2.1 authorization layer); MCP Apps; NVIDIA OpenShell policy-governed runtime (Feb 2026); Anthropic Sandbox Runtime (process-level Landlock); Docker Sandboxes microVM isolation (Aug 2026).
- [UNVERIFIED] METR reward-hacking signals: gpt-5.6-sol carries METR's highest public reward-hacking flag among tracked models (community-sourced, July 2026 — treat as unverified pending METR publication).

### Hour-by-hour: the June 11–12 report-to-order arc

