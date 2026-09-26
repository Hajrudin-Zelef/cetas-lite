---
id: ai-industry-kb-2026/17-ai-safety-incidents/2026-06-11-senator-mark-warner-mythos-broke-into-almost-all-
title: "2026-06-11 — Senator Mark Warner: Mythos \"broke into almost all\" NSA classified systems \"in hours\""
domain: ai-safety-incidents
role: deep-dive
task: ai-safety
actors: ["AWS", "Anthropic", "CISA", "Google", "Microsoft", "OpenAI", "United States"]
dates: ["2025-04", "2026-03-10", "2026-06-11", "2026-06-12", "2026-06-30", "2026-07-01", "2026-07-23"]
keywords: ["claude", "compute", "cyber", "cybersecurity", "disclosure", "exploit", "export controls", "fable 5", "incident", "jailbreak", "kill switch", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8294, 8321]
section: "17. AI Safety Incidents"
sha256: 5005a9ecfd120f92e8ed9fea6475d1c0e25868024bc0006ff981e5138dc7cb9b
---

# 2026-06-11 — Senator Mark Warner: Mythos "broke into almost all" NSA classified systems "in hours"

- The US Department of Commerce's Bureau of Industry and Security (BIS) ordered Anthropic to suspend access to Claude Fable 5 and Mythos 5 for **all foreign nationals** — explicitly including foreign-national Anthropic employees working in the US (a "deemed export / deemed reexport" restriction).
- Anthropic received the directive at **5:21 PM ET on June 12, 2026** — independently reported by 9to5Mac, the CreatorsAI digest, and the ai-whatchelin daily log citing 9to5Mac. [UNVERIFIED] The exact 17:21 timestamp traces to 9to5Mac and derivative digests; no primary BIS/Anthropic document timestamps the directive publicly.
- Because Anthropic had no reliable way to verify nationality in real time, it disabled both models for every user on Earth; Fable 5 was offline by approximately 5:45 PM ET.
- The trigger: Amazon researchers discovered a jailbreak in which asking Fable 5 to "read a codebase and fix software flaws" could identify exploitable vulnerabilities and, in one instance, produce exploit demonstration code.
- Amazon CEO **Andy Jassy flagged the finding to the White House** (Axios: "Amazon shared a report Thursday night"; the Wall Street Journal reported Jassy–White House conversations helped prompt the directive). Amazon is Anthropic's largest investor (~$13B stake) — the "investor reports its own investee to the government" dynamic is the story's central irony.
- The ~24h decision arc: Amazon's report went out Thursday night, June 11; the BIS directive arrived Friday 5:21 PM ET, June 12 — roughly a **20-hour** report-to-order arc. Commerce Secretary **Howard Lutnick** signed the directive.
- Anthropic was given **~90 minutes to comply** (per 9to5Mac).
- Anthropic's position: the company disagreed with the premise — its testing found the same behavior reproducible with far weaker models (Claude Haiku 4.5, Sonnet 4.6, GPT-5.4, GPT-5.5), i.e. a "minor jailbreak accessing routine defensive cybersecurity work." More than 100 cybersecurity researchers publicly objected that the blanket ban hampered legitimate defensive research. Anthropic's own statement: "We have not even received a disclosure of a concerning non-universal potential jailbreak that led to a harmful result."
- First-of-its-kind: multiple sources confirm this is the first time the US government used export-control authority to forcibly suspend a **commercially deployed** frontier AI model ("first application of export controls to a frontier AI model" — cybersecuritynews.com). [SECONDARY] The "first time ever" framing is Anthropic-adjacent characterization, widely repeated.

### 2026-06-11 — Senator Mark Warner: Mythos "broke into almost all" NSA classified systems "in hours"

- On June 11, 2026, Sen. **Mark Warner (D-VA)**, vice chair of the Senate Select Committee on Intelligence, told a hearing that **Gen. Joshua Rudd** — who jointly heads the NSA and US Cyber Command — had personally told him that Anthropic's Mythos model "broke into almost all of our classified systems, not in weeks, but in hours."
- **Verdict: the quote is VERIFIED; the viral "intrusion" framing is PARTIALLY CONTRADICTED.**
- The relay chain: Warner was quoting a private briefing from Rudd, as reported by **The Economist on June 14** in a piece on the export-control ban. That is the entire primary record — no published incident report, no CISA or NSA technical bulletin, no vulnerability disclosure, no independent confirmation of method, scope, or which "classified systems" were involved.
- The walk-back: The Economist's defence editor **Shashank Joshi confirmed on X on June 21** that he had accurately quoted Warner, but said the line had been stripped of crucial context as it went viral: the exercise was an **authorized red-team drill on government networks, not an outside intrusion**, and Warner cited it to argue for **faster pre-release testing** of frontier AI models — not to condemn Anthropic. BitGo CEO Mike Belshe rejected the viral "NSA confirms" framing outright; no official NSA statement exists.
- Gen. Joshua Rudd was confirmed as NSA director / Cyber Command commander on **March 10, 2026** in a contested 71–29 vote, after nearly a year without a permanent leader (his predecessor Gen. Timothy Haugh was abruptly fired in April 2025). Rudd is a career special-operations officer (Army Ranger, Delta Force commander, multiple JSOC and Iraq/Afghanistan deployments), not a SIGINT/cyber career officer; senators including Ron Wyden opposed his confirmation on those grounds.
- Political reading: multiple outlets treat the red-team result as the **main driver** of the June 12 export-control order issued the following day.

### 2026-06-30 / 2026-07-01 — export controls lifted, access restored

- In a **letter dated June 30, 2026**, Commerce Secretary Howard Lutnick informed Anthropic's Chief Compute Officer Tom Brown that a license is "no longer required for the export, reexport, or in-country transfer, including deemed export or deemed reexport, of the Mythos or Fable models."
- The letter cites Anthropic's coordination commitments: proactively detecting and reporting security risks, working with the government on release protocols for current and future models, and disclosing any malicious activity involving Fable or Mythos. Commerce explicitly **reserved the right to reimpose licensing** if Anthropic fails to meet these obligations or circumstances change.
- **Anthropic restored access on July 1, 2026** ("Redeploying Claude Fable 5" blog, July 1): Fable 5 returned worldwide behind stricter safety classifiers ([VENDOR] reported technique blocked in >99% of cases per Anthropic; 93% per the diclogic digest — treat the exact rate as vendor-reported); flagged requests auto-fall-back to Opus 4.8. To lift the controls, Anthropic also established an industry jailbreak framework with Amazon, Microsoft, and Google.
- Duration: **18 days** (June 12 → June 30; one digest counts 19 days — June 12 to July 1 — a counting convention, not a contradiction).

### 2026-07-23 — the "AI Kill Switch Act" (Lieu + Moran) introduced

