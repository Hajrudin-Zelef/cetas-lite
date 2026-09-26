---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-8
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 8)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "OpenAI", "Sakana", "United States"]
dates: ["2026-03", "2026-06-12", "2026-06-15", "2026-06-30", "2026-07-20"]
keywords: ["alignment", "claude", "copyright", "cyberattack", "export controls", "fable 5", "fugu", "ipo", "lawsuit", "mythos 5", "opus 4", "pretraining"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [218, 235]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 7eafe52f6b83089358210d4a2018fba56fd30835f67688e125d815e95bcc1bc6
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 8)

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

