---
id: ai-industry-kb-2026/17-ai-safety-incidents/timeline-and-context
title: "Timeline and context"
domain: ai-safety-incidents
role: deep-dive
task: ai-safety
actors: ["AWS", "Anthropic", "CISA", "Google", "Hugging Face", "Microsoft", "OpenAI", "United States"]
dates: ["2026-02", "2026-03", "2026-03-10", "2026-04-16", "2026-05-25", "2026-06-01", "2026-06-08", "2026-06-12", "2026-07-23", "2026-08"]
keywords: ["compute", "cyber", "disclosure", "fable 5", "governance", "jailbreak", "kill switch", "liability", "research", "sandbox"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8583, 8612]
section: "17. AI Safety Incidents"
sha256: 61cc6926a680540bd0d1db05369dc93ebdd5a75241656cb5f85d598a3b293744
---

# Timeline and context

- **Howard Lutnick** — US Commerce Secretary; signed the June 12 BIS directive and the June 30 lifting letter to Anthropic's Tom Brown; reserves the right to reimpose licensing.
- **Andy Jassy** — Amazon CEO; flagged the Fable 5 jailbreak to the White House on June 11 (night), triggering the ~20-hour report-to-order arc. Amazon is Anthropic's largest investor (~$13B).
- **Anthropic / Tom Brown** — Chief Compute Officer Tom Brown was the recipient of the June 30 lifting letter; the company disagreed with the ban's premise, complied in ~90 minutes, and established an industry jailbreak framework with Amazon, Microsoft, and Google.
- **Sen. Mark Warner (D-VA)** — vice chair of the Senate Select Committee on Intelligence; relayed the June 11 Mythos/NSA red-team claim secondhand from Gen. Rudd. [DIRECTIONAL] The claim functioned as the political driver of the BIS order.
- **Gen. Joshua Rudd** — director of NSA / commander of US Cyber Command (confirmed March 10, 2026, 71–29); Warner's cited source; no official NSA statement exists.
- **Shashank Joshi** — The Economist's defence editor; confirmed the quote June 21 but walked back the viral "intrusion" framing (authorized red-team drill; Warner's point was faster pre-release testing).
- **Rep. Ted Lieu (D-CA)** and **Rep. Nathaniel Moran (R-TX)** — bipartisan sponsors of the AI Kill Switch Act (introduced July 23, 2026; House only).
- **Philipp Emanuel Weidmann** — creator of Heretic; stripped Gemma 4 within 90 minutes of its public release; reports 3,500+ variants, 13M downloads.
- **Noam Schwartz / Alice** — AI safety research group co-investigator with the FT on the May 25, 2026 Heretic report.
- **NY AG Letitia James** — served the June 12, 2026 subpoena on behalf of 42 state AGs; largest multi-state action against a single AI company. Florida AG James Uthmeier separately sued OpenAI and Sam Altman June 1, 2026.
- **Adnan Khan** — disclosed the Cline CI compromise (Feb 9, 2026).
- **Aonan Guan; Zhengyu Liu & Gavin Zhong (Johns Hopkins)** — disclosed "Comment and Control" (April 16, 2026).
- **Unit 42 / Palo Alto Networks** — documented the first large-scale in-the-wild indirect prompt-injection campaigns (March 2026).
- **Zscaler** — documented 2026 search-poisoning indirect-injection campaigns.
- **CISA** — added LiteLLM CVE-2026-42271 to the Known Exploited Vulnerabilities catalog June 8, 2026.
- **KELA Cyber** — single-sourced reporting of the Mexican-agency jailbreak campaign (vendor-reported, treat cautiously).
- **OWASP** — LLM Top 10 2026 set the enterprise consensus posture: assume compromise, sandbox, least-privilege scopes.

- **Ron Wyden** — US senator who opposed Rudd's NSA confirmation on the grounds of his non-cyber background; his objection is the on-record marker that the director's cyber-capability claims arrived with contested credibility.
- **Mike Belshe (BitGo CEO)** — publicly rejected the viral "NSA confirms" framing of the Mythos claim outright.
- **Sam Altman** — named individually alongside OpenAI in Florida AG James Uthmeier's June 1, 2026 "defective product" suit; the personalization of liability matters for the governance narrative.
- **Alabama AG Steve Marshall** — [SECONDARY, single-source] launched a separate August 2026 OpenAI investigation after the sandbox-escape/Hugging Face breach; 14-state AGs sent OpenAI a records-preservation letter (winzheng.com only).
- **Adnan Khan** — disclosed the Cline CI compromise (Feb 9, 2026); the disclosure's sharpest lesson is the fix-vs-re-exploitation gap (30-minute fix, re-exploited 8 days later via a non-revoked token).
- **Brodt, Feldman, Schneier, Nassi** — authors of the February 2026 Promptware survey: 21 prompt-injection incidents across 2025–2026, 7 of 21 targeting AI coding assistants.
- **CISA** — added LiteLLM CVE-2026-42271 to the Known Exploited Vulnerabilities catalog on June 8, 2026, marking AI-gateway infrastructure as critical infrastructure under active attack.
- **OWASP** — the LLM Top 10 2026 (reported Aug 6) codified the enterprise consensus posture: assume compromise, sandbox, least-privilege tool scopes.
- **The "Attacker Moves Second" authors (arXiv 2510.09023)** — demonstrated that 12 published prompt-injection defenses fall at >90% ASR, resetting the 2026 baseline for what counts as a credible defense claim.

## Timeline and context

