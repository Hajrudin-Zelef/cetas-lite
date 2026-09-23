---
id: vague2-datacamp/datacamp/github-copilot-privacy-and-troubleshooting
title: "Confidentialité de GitHub Copilot : protections et guide de dépannage"
domain: datacamp
role: reference
task: article
actors: ["Microsoft"]
dates: ["2026-04", "2026-09-23"]
keywords: ["copilot", "agent", "agents", "attribution", "cloud agent", "license", "licenses", "training"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/github-copilot-privacy-and-troubleshooting.md
source_anchor: ""
source_lines: [1, 66]
sha256: e818ec8dfe29d09f7d60de9d02304714bb1b7a3120f1d9e969064aa1ad342aa7
---

# Confidentialité de GitHub Copilot : protections et guide de dépannage

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/github-copilot-privacy-and-troubleshooting
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Derrick Mwiti covers GitHub Copilot's privacy protections and troubleshooting. It explains that Copilot sends a snapshot of surrounding code context to GitHub servers when you type; the model processes it and returns a suggestion. For individual plans (Free, Pro, Pro+), GitHub's privacy statement allows interaction data to be used for model training unless users opt out in privacy settings. Business and Enterprise have separate contractual terms fully excluding interaction data from training, with no user action required.

Interaction data includes: inputs/prompts (chat/CLI commands), outputs (code suggestions/text, acceptances/rejections), code context (surrounding snippets, open files), metadata/structure (file names, repo structure, IDE navigation), and user feedback (thumbs up/down, comments). Private repository code "at rest" is not used; only interaction data generated during active use may be used for training unless opted out.

A table compares Free/Pro/Pro+ vs Business vs Enterprise across training use, private repo code at rest, prompt/output retention (IDE: not retained; outside IDE: 28 days), admin controls, content exclusions, and IP protection. Content exclusions do not yet apply to Edit mode, Agent mode in IDE chats, GitHub Copilot CLI, or the cloud agent. IP protection requires the duplication detection filter to be enabled and the suggestion used unmodified.

Configuring privacy: individual opt-out via GitHub Settings → "Allow GitHub to use my data for AI model training" → Disabled (stops future collection, doesn't remove previously collected data). Organization/Enterprise policies: Settings → Copilot → Policies for features, licenses, model selection; org policies override individual preferences; Enterprise owners can set inherited policies.

Content exclusions: exclude files so inline suggestions aren't offered, content isn't used for suggestions in other files or Chat, and code review doesn't apply. Configured at repo level (Settings → Copilot → Content Exclusion, glob patterns like `**/secrets/**` or `*.env`) or org level; rules are additive and Enterprise rules take precedence. Limitations: Copilot CLI, Agent mode, and Cloud Agents don't respect exclusions (mid-2026); semantic leakage via typing/definitions; symlinks and remote filesystems not covered.

Code referencing and the duplication filter: Copilot compares suggestions to known public code; matches above a similarity threshold are blocked or flagged with attribution, recording date/time, file, code snippet, license, and source URL. The filter doesn't detect short/generic snippets or restructured/partially modified code. IP protection (Business/Enterprise only) covers legal defense if a suggestion triggers a claim, requiring the duplication filter enabled and eligible plan.

Policy management on GitHub.com: enable/disable code completions, Chat, code review, GitHub CLI integration, and agent mode; assign licenses; restrict allowed AI models (FedRAMP-compliant options available). Audit/compliance: usage metrics under Insights > Copilot usage, NDJSON export, seat/license data under Org Settings > Copilot > Access. Audit logs don't include client session data (prompts); stream to SIEM for longer retention.

Troubleshooting: missing suggestions (check Copilot status icon — a diagonal indicates an active exclusion; update IDE/extension, verify license, check exclusions, proxy/VPN); unexpected content exclusion (up to 30 min delay; reload window/restart IDE; test by asking Chat to explain a file); authentication issues (sign out/in, verify license, SSO); rate limiting (switch to auto model selection or lower-multiplier model; monitor via analytics dashboard); check githubstatus.com for service issues.

## Key points

- Individual plans (Free/Pro/Pro+) may use interaction data for training unless opted out; Business/Enterprise are contractually excluded.
- Private repository code at rest is never used; only active interaction data may be.
- Interaction data: prompts, outputs, code context, metadata/structure, user feedback.
- Prompt/output retention: not retained in IDE; 28 days outside IDE.
- Content exclusions block Copilot from specific files/paths (Business/Enterprise only); not respected by CLI, Agent mode, or Cloud Agents.
- Code referencing and the duplication filter flag matches to public code; IP protection (Business/Enterprise) requires the filter enabled.
- Org policies override individual preferences; Enterprise policies are inherited.
- Opt-out stops future collection but doesn't remove previously collected data.
- Usage metrics are under Insights > Copilot usage; audit logs exclude client session data (prompts) — stream to SIEM.
- Troubleshooting steps: check status icon, update IDE, verify license, check exclusions, reload window, check proxy/VPN, monitor rate limits, check githubstatus.com.

## Technical data / figures

| Dimension | Free/Pro/Pro+ | Business | Enterprise |
|-----------|---------------|----------|------------|
| Used for model training | Requires opt-out | No (contractually excluded) | No (contractually excluded) |
| Private repo code at rest | Not used | Not used | Not used |
| Prompt/output retention | IDE: not retained; outside IDE: 28 days | Same | Same |
| Admin controls | Individual only | Org-level policies, license management | All Business + enterprise policy inheritance, audit logs |
| Content exclusions | Not available | Repo and org level | Enterprise-wide |
| IP protection | Not included | Yes (with duplication filter) | Yes (with duplication filter) |

| Detail | Value |
|--------|-------|
| Training policy change | April 2026 |
| Outside-IDE retention | 28 days |
| Exclusion propagation delay | Up to 30 minutes |
| Code referencing records | Date/time, file, code snippet, license, source URL |
| Audit log retention (default) | 180 days |
| Features not respecting exclusions | Copilot CLI, Agent mode, Cloud Agents |
| Example exclusion patterns | `**/secrets/**`, `*.env` |

## Why this source matters for the RAG

This guide provides detailed, current (2026) information on GitHub Copilot's data handling, privacy controls, content exclusions, IP protection, and troubleshooting, making it highly valuable for security, compliance, and support queries. Its comparison tables, configuration examples, and troubleshooting checklists supply precise, actionable facts for the knowledge base.
