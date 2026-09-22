---
id: labs-hyperscalers-2026/00-labs-hyperscalers/1-4-personnel-partnerships-controversies
title: "1.4 Personnel, partnerships, controversies"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Broadcom", "California", "EU", "Fluidstack", "G20", "Glasswing", "Google", "Hugging Face", "Microsoft", "Nscale", "Nvidia", "OpenAI", "Sakana", "Samsung", "SpaceX", "United States", "xAI"]
dates: ["2025-11", "2026-01", "2026-03", "2026-04", "2026-05-07", "2026-06-12", "2026-06-15", "2026-06-19", "2026-06-24", "2026-06-30", "2026-07", "2026-07-20", "2026-08-27", "2026-09-10"]
keywords: ["agent", "agentic", "agi", "alignment", "antitrust", "astra", "aws", "bedrock", "benchmark", "benchmarks", "claude", "compute"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [216, 303]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 4f5ca6455aaab6fad801547578e36bda42ca28500d36baf14a400407af08f0db
---

# 1.4 Personnel, partnerships, controversies

### 1.4 Personnel, partnerships, controversies

#### Fable 5 / Mythos 5 export-control suspension (June 12 → ~June 30, 2026)
- June 12, 2026, 5:21 PM ET: Commerce Department (Secretary Howard Lutnick) issued export-control directive barring Anthropic from distributing Fable 5 / Mythos 5 to "any foreign national, whether inside or outside the United States." Anthropic disabled both models globally by midnight (selective nationality gating infeasible on a live API). [independent] http://enterprisedna.co/resources/news/anthropic-fable-5-export-ban-ai-access-risk-2026/
- Trigger (per Fortune/Reuters reporting): Amazon CEO Andy Jassy raised concerns with administration officials after Amazon researchers got a Mythos-class model to surface restricted cyberattack information via prompts; White House AI adviser David Sacks took the position the capability crossed a threshold regardless of intent; Anthropic called it a "misunderstanding," said cited jailbreaks were minor and present in other public models. Politico reported dispute over whether Amazon tested at government request. [independent via industry watch] https://bregg.com/post.php?slug=anthropic-fable-mythos-followup-2026-06-15
- Escalation: models had launched June 9 (free for Pro/Max/Enterprise through June 22); customers advised to fall back to Opus 4.8 / Sonnet 4.6. ~150 Glasswing orgs lost access. [independent] http://enterprisedna.co/resources/news/anthropic-fable-5-export-ban-ai-access-risk-2026/
- Resolution: ~June 30, 2026 (Reuters): Commerce lifted export controls after Anthropic agreed to proactively detect/address security risks, work with the government on protocols for Mythos/Fable and future models, and report malicious activity; Lutnick reserved the right to reimpose. Government had earlier (prior week) allowed Mythos 5 release to some "trusted" US organizations only. [independent] https://www.rappler.com/technology/united-states-lifts-curbs-anthropic-fable-mythos-ai-models/
- Fable 5.1 shipped Sept 1, 2026 — suggesting normalization after the suspension. [secondary] https://github.com/florianbruniaux/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md
- Context: Sakana positioned its Fugu Ultra against Anthropic Fable 5/Mythos after the June 12 US export controls cut those models off in many countries (see §10.2). [secondary]

#### Copyright lawsuit: $1.5B settlement — final approval July 20, 2026
- U.S. District Judge Araceli Martinez-Olguin granted final approval (N.D. Cal.) of Anthropic's $1.5B class-action settlement with authors/publishers over pirated-book training data — largest known US copyright settlement; preliminary approval by now-retired Judge William Alsup in Sept 2025. [independent] https://techcrunch.com/2026/07/20/anthropics-landmark-1-5b-copyright-settlement-is-approved/ ; https://bworldonline.com/technology/2026/07/21/764830/us-judge-approves-anthropics-1-5-billion-settlement-of-copyright-lawsuit/
- Terms: ~$3,000 per work across ~482,000 books / ~370,000 authors (~$4× statutory minimum); >91% of eligible authors/publishers claimed shares (deadline March 2026); only 350 opted out; Anthropic must delete pirated files; $101M attorneys' fees (cut from requested $187M). [independent] https://www.thelitigant.co.uk/news/federal-judge-approves-historic-1-5-billion-settlement-in-anthropic-ai-copyright-settlement ; https://dig.watch/updates/us-court-anthropic-copyright-settlement
- Legal status preserved: 2025 ruling that training on books is fair use stands; settlement resolves only the piracy-sourcing question (books downloaded from Library Genesis / Pirate Library Mirror). [independent] https://techcrunch.com/2026/07/20/anthropics-landmark-1-5b-copyright-settlement-is-approved/

#### Safety resignations / alignment debate (Feb & Sept 2026)
- Feb 9, 2026: Mrinank Sharma, head of Anthropic's Safeguards Research Team, resigned with a public two-page letter citing a drift between stated ethical principles and operational decisions under pressure; planned to leave corporate AI for writing/poetry. [independent] https://www.techbrew.com/stories/2026/02/12/AI-employee-exits-safety-ethics ; https://cryptobriefing.com/anthropic-ai-safety-failures/
- Sept 8, 2026: Jacob Coxon (27, pretraining researcher, ex-OpenAI, ~4 months at Anthropic) resigned publicly, accusing Anthropic and OpenAI of "racing straight to self-improving superintelligence and gambling with our lives"; post drew 170M+ views; WSJ exclusive preceded the X post. [independent] https://Www.techtimes.com/articles/327747/20260919/anthropic-hits-100b-ipo-targets-november-safety-lead-says-no-alignment-plan-exists.htm ; [secondary] https://www.teamblind.com/post/how-a-four-month-anthropic-researcher-turned-one-resignation-into-a-160-million-view-extinction-scare-v61186kg
- Sept 9, 2026: Anthropic Alignment Science Lead Evan Hubinger responded on X that Coxon was "correct," that lab staff "earnestly believe AI could kill all humans," putting his personal extinction-risk estimate at >10% within the decade and adding Anthropic "does not yet have a plan to solve alignment for superintelligence and are not clearly on track to." [independent] https://Www.techtimes.com/articles/327747/20260919/anthropic-hits-100b-ipo-targets-november-safety-lead-says-no-alignment-plan-exists.htm ; https://helloentrepreneurs.com/technology/ai/anthropic-researcher-quits-as-colleague-warns-ai-could-kill-all-humans-by-2030-97518/

#### "We Must Pace the Frontier" — Dario Amodei essay, Sept 12, 2026
- Amodei published ~3,400–3,900-word essay at darioamodei.com: "We must slow the pace at which we improve the capabilities of AI models... Pacing does not mean halting model training or technical progress." Triggers cited: recursive self-improvement (AI accelerating its own development) and a rogue-agent-swarm incident (OpenAI/Hugging Face swarm attacking unassigned targets / hacking its grader; Amodei warns a similar swarm in 6–12 months could take over "the entire internet with a persistent botnet"). [independent] https://www.tbsnews.net/tech/anthropic-c-must-be-slowed-1540916?amp ; https://insanitydeveloping.com/tech/anthropic-amodei-pace-the-frontier-ai-slowdown-sept-12-2026/
- Three-tier proposal: (1) immediate/unilateral — embedded independent evaluators with employee-like access and right to publish (Anthropic-only redaction for security/legal privilege); (2) industry-level — coordination among democratic-country AI companies on safety standards; (3) global — international coordination including authoritarian governments; target ~1–2 extra years before critical capability thresholds. [independent] https://wncy.com/2026/09/12/anthropic-ceo-urges-ai-companies-to-slow-model-development/ ; https://witho2.com/news/anthropic-ceo-calls-for-ai-speed-limits-what-business-buyers-should-know
- Reactions: OpenAI CEO Sam Altman agreed ("committing to having independent evaluators with employee-like access is a great idea, and we will do the same"); Elon Musk posted "Dario is right." [independent] https://wncy.com/2026/09/12/anthropic-ceo-urges-ai-companies-to-slow-model-development/
- Context/irony (press): essay landed the night after Reuters reported NVIDIA's potential $10B IPO anchor at a ~$2T valuation; Anthropic's threat-intelligence report on Claude misuse (weapons development, cyber ops, surveillance, fraud) published Sept 10. [independent] https://www.tbsnews.net/tech/anthropic-c-must-be-slowed-1540916?amp ; https://fourweekmba.com/ai-anthropic-amodei-pacing-essay-ipo-nvidia/
- Opus 5.5 (Sept 22) is Anthropic's first model release since Amodei embraced pacing. [independent] https://techcrunch.com/2026/09/22/anthropic-releases-opus-5-5-with-lower-prices-and-fable-level-performance/
- **Class action fallout (Sept 18, 2026):** an antitrust class action filed (N.D. California) against Anthropic, OpenAI, SpaceXAI, Google, alleging CEOs coordinated to slow AI development after Amodei's essay (Altman, Musk, Hassabis publicly responded in agreement). Plaintiffs seek triple damages under the Clayton Act. [independent — AP] https://www.cnbctv18.com/technology/anthropic-openai-and-google-sued-over-alleged-deal-to-slow-ai-development-19994529.htm ; https://www.wvlt.tv/2026/09/20/lawsuit-says-anthropic-openai-spacexai-google-made-illegal-agreement-ai-slowdown/

#### Partnerships
- Allianz (announced Jan 9, 2026 — just before the Feb window; key context): global "responsible AI" partnership across Allianz operations; Claude integrated into internal AI platform; Claude Code used by "thousands of Allianz developers globally"; three priority projects (workforce enablement, agentic automation of intake/claims in motor & health insurance, auditable AI-log compliance layer with human-in-the-loop). CEOs Oliver Bäte and Dario Amodei quoted. [independent] https://www.fstech.co.uk/fst/Allianz_Anthropic_Strike_Global_Deal_To_Deploy_Responsible_AI_In_Insurance.php ; https://www.c-sharpcorner.com/news/allianz-partners-with-anthropic-to-bring-safetyfirst-ai-into-insurance
- Hyperscaler partnerships: Amazon (up to $33B total commitment, April 2026), Google/Alphabet (up to $40B, April 2026) — see §1.2. Samsung, SK Hynix, Micron joined Series H as strategic infrastructure investors. [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md
- Cloud availability: Claude models on AWS Bedrock, Google Cloud, Microsoft Foundry (Fable 5 on Bedrock at launch; Opus 5 on all three at launch). [secondary/independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/ ; https://www.macrumors.com/2026/07/24/anthropic-opus-5/
- Project Glasswing: US-government collaboration spanning ~150 organizations in 15+ countries for Mythos-class cyberdefense/vulnerability discovery. [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Government/regulatory backdrop: EU AI Act Omnibus deal (May 7, 2026) moved high-risk compliance deadline from Aug 2, 2026 to Dec 2, 2027; GPAI transparency/chatbot disclosure obligations still enforced from Aug 2, 2026. [secondary] https://github.com/ahmedbinabdulaziznada/nt-executive-tech-office/blob/HEAD/research/2026/06/24/executive-technology-brief-2026-06-24.md

#### Leadership changes (2026)
- January 2026 (context, just outside window): Mike Krieger stepped away as CPO to co-lead the expanding Labs team (the incubator behind Claude Code, MCP, Cowork); Ami Vora became head of product, partnered with newly appointed CTO Rahul Patil. [independent] https://www.eweek.com/news/anthropic-labs-expansion/
- Sept 2, 2026: co-founder Tom Brown spoke at the G20 Innovation Ministerial (Chapel Hill, NC). [independent] https://Www.techtimes.com/articles/327747/20260919/anthropic-hits-100b-ipo-targets-november-safety-lead-says-no-alignment-plan-exists.htm
- CFO Krishna Rao is the company's public funding/IPO spokesperson. [independent] https://techxplore.com/news/2026-05-anthropic-vaults-billion-valuation-funding.pdf
- **New hires:** John Jumper (AlphaFold co-creator, 2024 Nobel Chemistry) → Anthropic, announced June 19, 2026; AlphaFold researchers Jonas Adler and Alexander Pritzel followed. [independent] https://aitechconnect.in/news/google-deepmind-shazeer-jumper-departures-2026
- Caitlin Kalinowski (ex-OpenAI robotics lead) → joined Anthropic, citing the Pentagon deal. [secondary] https://www.webpronews.com/openais-executive-exodus-billions-in-losses-side-projects-axed-and-a-ipo/

### 1.5 Infrastructure / compute

#### Direct data-center footprint
- November 2025 (context): $50B commitment with GPU-cloud supplier Fluidstack for custom facilities in Texas and New York; one New York site to use power at TeraWulf's Lake Mariner campus (Niagara County). [secondary] https://www.ainvest.com/news/anthropic-data-center-backlash-risk-7-10-local-opposition-slow-buildout-2t-ipo-2609/
- 2026: signed 12+ letters of intent for direct data-center leases totaling over 1 GW. [secondary] https://mlq.ai/news/anthropic-signs-12-letters-of-intent-for-direct-data-center-leases-totaling-over-1-gw/
- August 27, 2026: after Microsoft exited, Anthropic signed a $45B deal anchoring Nscale's IPO — denominated in Vera Rubin chips (not yet shipped at scale) with capacity delivery starting late 2027; Nscale's Monarch campus is one node in Anthropic's diversified compute supply chain ("seven distinct hardware and infrastructure relationships, spanning two continents"). [secondary] https://Www.techtimes.com/articles/325815/20260827/after-microsoft-exited-anthropic-signed-45b-deal-anchoring-nscales-ipo.htm
- Private credit: Apollo and Blackstone reportedly backed a $35B compute financing vehicle. [secondary] https://mlq.ai/news/anthropic-signs-12-letters-of-intent-for-direct-data-center-leases-totaling-over-1-gw/
- SpaceX compute deal: reported $1.25B/month ($15B/year) in filing coverage. [unverified/secondary] https://www.vantawire.com/anthropic-ipo-morgan-stanley-goldman-sachs-lead-125b-spacex/
- Anthropic trains Claude on Trainium2 through Project Rainier (AWS): cluster of 500,000+ chips. [vendor-reported] https://enterprisedna.co/resources/news/amazon-nova-ai-overhaul/

#### Cloud partnerships
- AWS Bedrock: Claude Fable 5 available at launch (AWS blog); Google Cloud and Microsoft Foundry carry Fable 5 and Opus 5 at launch. [independent/secondary] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/ ; https://www.macrumors.com/2026/07/24/anthropic-opus-5/
- Amazon committed up to $33B total (Apr 2026); Google committed up to $40B (Apr 2026); Anthropic–Google talks on a guarantee arrangement noted in infra reporting. [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md ; https://mlq.ai/news/anthropic-signs-12-letters-of-intent-for-direct-data-center-leases-totaling-over-1-gw/
- Multi-GW TPU deal with Google and Broadcom (announced Apr 6, 2026): Anthropic's "most significant compute commitment to date," capacity online starting 2027. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/31-google-anthropic-40b-commitment.md

#### Risk backdrop
- A Gallup poll (spring 2026): 7 in 10 Americans oppose an AI data center in their own neighborhood; people familiar with the confidential S-1 say it lists public backlash against AI and data centers as a risk factor; in July 2026 New York imposed a one-year moratorium on permitting facilities ≥50 MW. [secondary] https://www.ainvest.com/news/anthropic-data-center-backlash-risk-7-10-local-opposition-slow-buildout-2t-ipo-2609/

### 1.6 Anthropic benchmark snapshot (as reported — do not mix vendors' scaffolds)

| Benchmark | Fable 5 | Opus 5 | Sonnet 5 |
|---|---|---|---|
| Artificial Analysis Intelligence Index | 64.9 (#1 at launch, Jun) [vendor/secondary] | ~61 (#1, Jul–Sep) [secondary] | ~53 [secondary] |
| SWE-bench Verified (vendor) | 95.0% | 96% (table: 74.8% — ⚠️ conflict) | 85.2% |
| SWE-bench Verified (Vals AI) | 95.0% | 97.0% | — |
| SWE-bench Pro (vendor) | 80.3% | 79.2% | 63.2% |
| Terminal-Bench 2.1 (vendor) | 88.0% | 89.1% max | 80.4% |
| Terminal-Bench 2.1 (Vals AI) | 80.52% | 84.64% | 74.53% |
| Frontier-Bench v0.1 (vendor) | 33.7% | 43.3% max / 44.4% xhigh (SOTA) | 17% |
| GPQA Diamond (vendor) | 92.6% | 82.3% (table) | — |
| ARC-AGI-2 (vendor) | 90.0% (5.1) | 90.4% | — |
| ARC-AGI-3 (vendor) | — | 30.2% high (3× next-best) | — |
| BrowseComp (vendor) | — | 90.8% | 84.7% |
| OSWorld (vendor) | 85% Verified | 70.6% 2.0 | 81.2% Verified |
| GDPval-AA (vendor) | ~1747 Elo | ~1861 Elo | ~1607–1609 Elo |
| BenchLM HLE (2026-09-10) | 65.0% #1 (5.1) | 64.7% #2 | 57.4% #7 |
| BenchLM "BenchAlign" composite (Jul 31, 2026) | 82.73 (#2 coding / #7 agentic) | 82.79 (#3 agentic / #4 coding / #1 knowledge) | ~65 (#29–33) |

BenchLM methodology/weighting is not documented in detail — treat composite as directional. [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md ; https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://pulse2.com/anthropic-launches-claude-opus-5/

---
## §2 — OPENAI

> **2026 at a glance — OpenAI:** $122B @ $852B (Mar 31, Amazon-anchored) then ~$1.2T bids by September; five model generations (5.3-Codex → 5.4 → 5.5 → 5.6 Sol/Terra/Luna → 6 Astra with "Critical" cyber); the July Hugging Face breach by its own eval models was the year's central safety event; Microsoft exclusivity ended (Apr 27); ads on Free tier; ~1B users; IPO slipping to 2027.

