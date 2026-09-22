---
id: ai-industry-kb-2026-wave6/13-anthropic/asl-3-responsible-scaling-policy
title: "ASL-3 / Responsible Scaling Policy"
domain: anthropic
role: deep-dive
task: regulation
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Microsoft", "OpenAI"]
dates: ["2025-05", "2025-08", "2026-01-12", "2026-02", "2026-02-10", "2026-04", "2026-05", "2026-05-28", "2026-06-01", "2026-07", "2026-08-22"]
keywords: ["asl", "agent", "agents", "aws", "bedrock", "chatgpt", "claude", "compute", "consumer", "context window", "cost", "cybersecurity"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6510, 6593]
section: "§13. Anthropic"
sha256: 904f884bb1e97f172173e10680dda49b0a8f71f9a0f4da59f1bbe1c53d710a7a
---

# ASL-3 / Responsible Scaling Policy

### ASL-3 / Responsible Scaling Policy
- Anthropic activated the ASL-3 Deployment and Security Standards with the launch of Claude Opus 4 (2025) as a precautionary, provisional action [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- Explicitly not a determination that Opus 4 had passed the Capabilities Threshold [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- Anthropic had ruled out ASL-4 for Opus 4 and ASL-3 for Sonnet 4 [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- ASL-3 Security Standard: increased internal security against model-weight theft [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- ASL-3 Deployment Standard: narrowly targeted measures against CBRN misuse [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- Designed not to cause refusals except on a very narrow topic set [VENDOR](https://www.anthropic.com/news/activating-asl3-protections)
- Per secondary reporting, Anthropic activated ASL-3 safeguards for relevant models in May 2025 [SECONDARY](https://cio.economictimes.indiatimes.com/news/artificial-intelligence/anthropic-strengthens-ai-safety-regime-with-new-transparency-and-risk-controls/128770430)
- RSP v2.2 (May 2025) refined security exclusions under ASL-3 [SECONDARY](https://cio.economictimes.indiatimes.com/news/artificial-intelligence/anthropic-strengthens-ai-safety-regime-with-new-transparency-and-risk-controls/128770430)
- RSP v3.0 (reported February 2026, per TIME) was the largest structural overhaul — reorganizing catastrophic risk around four capability thresholds [SECONDARY](https://cio.economictimes.indiatimes.com/news/artificial-intelligence/anthropic-strengthens-ai-safety-regime-with-new-transparency-and-risk-controls/128770430)
- Anthropic's Fable 5 jailbreak severity framework (developed with Glasswing partners) was introduced as a draft proposal to standardize how the industry describes jailbreak risk — a bid to shape regulatory language [SECONDARY](https://github.com/vstorm-co/oss-website/blob/HEAD/src/data/blog/en/gemini-3-deep-think-advancing-science-engineering.mdx)

### Revenue / funding scale (2026)
- Annualized revenue run-rate trajectory (company-stated/leak-reported, unaudited): ~$9B end-2025 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- $14B Feb 2026 (Series G) [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- $19B Mar 2026 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- $30B Apr 2026 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- $47B May 2026 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- $65B end-July 2026 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- Series G February 2026: $30B raised at $380B post-money [SECONDARY](https://github.com/uroshp/scout-ci/blob/HEAD/v2/archive/anthropic__vs__openai__general/current.md)
- Series H May 2026: $65B raised at $965B post-money [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/60-anthropic-ipo-confidential-filing.md)
- Confidential draft Form S-1 filed 2026-06-01 — second frontier lab after OpenAI's May 22 filing [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/60-anthropic-ipo-confidential-filing.md)
- Q2 2026 preliminary revenue $11.5B (~14× YoY) per investor update via CNBC [SECONDARY](https://github.com/duh17/oppi/blob/HEAD/clients/apple/OppiTests/Fixtures/2026-08-22-anthropic-s1-report.md)
- Customer base: 300,000+ business customers [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- 1,000+ enterprise customers paying >$1M/year [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- 8 of the Fortune 10 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- Revenue mix ~70–80% API/enterprise contracts, 10–20% consumer subscriptions — the mirror image of OpenAI's consumer-heavy mix [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- Claude Code reached ~$1B annualized revenue within six months of its mid-2025 launch [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- ~$2.5B run rate by Feb 2026 with enterprise over half [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- Business subscriptions quadrupled [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- ACCOUNTING DISPUTE: OpenAI has publicly disputed ~$8B of Anthropic's reported figure on gross-vs-net cloud accounting [SECONDARY](https://github.com/uroshp/scout-ci/blob/HEAD/v2/archive/anthropic__vs__openai__general/current.md)
- Estimated compute cost per revenue dollar fell from $0.71 (Q1 2026) to projected $0.56 (Q2) [SECONDARY](https://www.ainvest.com/news/15-000x-cost-number-anthropic-26-96-safety-headline-graded-1-2608/)
- Gross margin recovered from estimated −94% (2024) to a 44–60% range for 2026 [SECONDARY](https://www.ainvest.com/news/15-000x-cost-number-anthropic-26-96-safety-headline-graded-1-2608/)
- ~$80B committed through 2029 to AWS and Google [SECONDARY](https://www.ainvest.com/news/15-000x-cost-number-anthropic-26-96-safety-headline-graded-1-2608/)

### Model deprecations/retirements (extended)
- Claude Opus 4.1: deprecated ($15/$75 pricing table status) [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Claude Opus 4: retired except on Vertex AI [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Claude Sonnet 4: retired except on Bedrock and Vertex AI [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Claude Haiku 3.5: retired except on Bedrock and Vertex AI [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Current official docs lineup: Claude Fable 5 [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Opus 4.8 [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Sonnet 5 [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Haiku 4.5 [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Mythos 5 limited to approved customers (Project Glasswing) [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Mythos Preview invitation-only for defensive cybersecurity workflows [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Claude Design is an Anthropic Labs research preview powered by Claude Opus 4.7 (vision-focused), not Opus 4.8 [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- The June-15 Claude Code credit-pool split (Interactive vs Agent-SDK) was announced then paused before taking effect [SECONDARY](https://github.com/beannation/wtclaude/blob/HEAD/site/src/content/blog/claude-billing-changes-june-2026.md)

- Claude Cowork: launched January 12, 2026 as a research preview for Claude Max subscribers on macOS. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Windows support followed on February 10, 2026. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: built on the same Claude Agent SDK that powers Claude Code. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: a Chrome plugin for web navigation extends task scope beyond the local file system. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: parallel task queuing lets multiple tasks run simultaneously. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Anthropic's 2026 usage data reports average task completion dropping from 3.1 hours to ~15 minutes — a 92% reduction. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Anthropic internal research from August 2025 found its own engineers using Claude in 60% of work with a 50% productivity boost. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: HUB International's late-2025 deployment found an 85% productivity increase and 2.5 hours saved per employee per week, with >90% satisfaction. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Anthropic's Economic Index (May 2025) found 57% of AI usage is augmentation — AI as partner rather than replacement. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Claude Projects gives each team a shared knowledge base with a 200,000-token context window. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: artifacts can be shared and edited across Project members in real time. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: the 2026 "Cowork & Plugins for the Enterprise" suite embeds Claude in the apps workers already use daily. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: enterprise plugins cover Excel and PowerPoint with cross-application context (a financial model in Excel becomes a slide deck without manual copy-paste). [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Google Workspace access covers Drive, Gmail, and Calendar with inline citations from internal documents. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Slack plugin covers channel summarization, thread digests, and direct task delegation via the Claude Slack app. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: DocuSign and LegalZoom plugins handle document processing automation for legal workflows. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: the Zapier plugin connects to more than 8,000 third-party applications. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: companies can build private "plugin marketplaces" — internal repositories of department-configured Claude agents. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: subscription coverage spans Pro, Max, Team, and Enterprise (versus Chat alone which also covers Free). [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Claude Cowork: Cowork targets knowledge workers with a graphical interface, while Claude Code (launched 2025) targets engineers with a terminal CLI. [SECONDARY](https://aiagentskit.com/blog/claude-cowork/)
- Opus 4.7: launched in April 2026 as the then-latest flagship before Opus 4.8. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: API model name is claude-opus-4-8. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: fast mode is "three times cheaper than previous fast mode models," corroborating the $30/$150→$10/$50 cut math. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: substantially lower rates of misaligned behavior compared to Opus 4.7 (Anthropic vendor claim via Neowin). [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: Dynamic Workflows lets Claude plan large tasks and run hundreds of parallel subagents in a single session. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: Dynamic Workflows available on Claude Code Enterprise, Team, and Max plans. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: defaults to high effort; users can also choose "extra" or "max" effort. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: the Messages API now supports system entries inside the messages array. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: effort controls are similar to features OpenAI's Codex and ChatGPT "have offered for months" (Neowin comparison). [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Anthropic: "will bring Mythos-class models to all customers in the coming weeks," announced May 28, 2026 alongside Opus 4.8. [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Fable 5.1: scored 55.8% on Terminal-Bench 4.0 versus Fable 5's 42.0% (reported via Decrypt). [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Mythos 5.1: "designed experimentally validated protein binders" (reported via VentureBeat). [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Fable 5.1: "trained a neural network that produced a higher-resolution elevation map covering roughly a third of Venus" (reported via VentureBeat). [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Fable 5.1: available via the Claude API, Amazon Bedrock, Google Cloud, and Microsoft Foundry under model ID "claude-fable-5-1." [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)

