---
id: vague2-datacamp/datacamp/github-copilot-plans
title: "Formules GitHub Copilot : guide des fonctionnalités et de l'administration"
domain: datacamp
role: reference
task: article
actors: ["Microsoft"]
dates: ["2026-04", "2026-06", "2026-09-23"]
keywords: ["copilot", "agent", "agentic", "fine-tuning", "governance", "pricing", "training"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/github-copilot-plans.md
source_anchor: ""
source_lines: [1, 67]
sha256: 0bdbac08a3d37d4569e726e8df192a48fe9080cb077470015e5fc8878f1345b0
---

# Formules GitHub Copilot : guide des fonctionnalités et de l'administration

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/github-copilot-plans
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Tim Lu explains GitHub Copilot's tiered plans in 2026, focusing on privacy boundaries, administrative controls, auditability, and governance. It notes the configuration surface is vast because needs vary: a solo developer on personal projects faces different privacy/compliance concerns than an enterprise admin managing thousands of seats on regulated repositories.

GitHub offers four individual tiers (Free, Student, Pro, Pro+) and two organizational tiers (Business, Enterprise). A key 2026 change: GitHub is transitioning to usage-based billing, replacing the "Premium Request Unit" (PRU) framework with GitHub AI Credits in June 2026. Basic code completions and "Next Edit" suggestions remain unlimited and consume no credits; advanced operations (multi-file chat, agentic workflows, long coding sessions, deep code reviews) consume AI Credits based on tokens relative to the model's published API rates.

Plan details: Free (free, limited AI Credits, basic completions/chat); Student (free for verified students/teachers, extended AI Credits); Pro ($10/month, 1,000 Base + 500 Flex = 1,500 total, extended IDE integrations and multi-model support); Pro+ ($39/month, 3,900 Base + 3,100 Flex = 7,000 total, includes GitHub Spark); Business ($19/user/month, 1,900 credits/user, centralized seat management, audit logs, file exclusions, IP indemnity); Enterprise ($39/user/month, 3,900 credits/user, repository indexing, custom fine-tuning, global governance).

Business and Enterprise provide contractual guarantees that interaction data is never used for training. Individual plans moved to opt-out by default in April 2026 (previously opt-in). The article distinguishes "code at rest" (private repo code not read or used in public training sets) from "interaction data" (prompts, chat requests, cursor context, surrounding code blocks, acceptance metrics, feedback logs).

Business adds centralized seat assignment, org-wide reference policies, structured audit logs, content/file exclusions, and commercial IP indemnity. Enterprise adds Copilot Spaces, enhanced GitHub.com Chat, and hierarchical policy inheritance. Enterprise requires an active GitHub Enterprise Cloud subscription ($21/user/month), making the real minimum $60/user/month; Business can be bought natively by orgs on GitHub Free or Team.

The guide covers file exclusion configuration (at repo or org level, using path patterns), organization-wide policy management (feature enablement, public code filter, model restrictions, custom org instructions), Enterprise policy inheritance (Enterprise > Organization > user preferences), audit logs (Copilot events, search/filter/export, SIEM streaming), and managing seats via the REST API (endpoints, authentication with PAT scopes `admin:org`, automation patterns like identity onboarding, inactive seat reclamation, financial dashboards). A Python example shows assigning a seat via `POST /orgs/{org}/copilot/billing/selected_users`.

## Key points

- Six tiers: Free, Student, Pro, Pro+ (individual); Business, Enterprise (organizational).
- June 2026: transition from Premium Request Units to usage-based GitHub AI Credits.
- Basic completions and Next Edit suggestions stay unlimited; advanced operations consume AI Credits.
- Pricing: Free (free), Pro ($10/mo), Pro+ ($39/mo), Business ($19/user/mo), Enterprise ($39/user/mo).
- Business/Enterprise contractually guarantee interaction data is never used for training; individual plans are opt-out by default since April 2026.
- Business adds centralized seats, audit logs, file exclusions, org policies, and IP indemnity; Enterprise adds Spaces, hierarchical policy inheritance, and org-level usage metrics.
- Enterprise requires GitHub Enterprise Cloud, making the real minimum $60/user/month.
- File/content exclusions prevent Copilot from processing specified files/dirs (Business/Enterprise only); not supported by Copilot CLI or agent mode.
- Policy inheritance order: Enterprise > Organization > user preferences.
- Audit logs track seat assignments, policy changes, exclusions, and feature enablement; streamable to SIEMs on Enterprise.
- Seats can be managed as code via the REST API with `admin:org` scoped PATs.

## Technical data / figures

| Tier | Audience | Base price | Monthly volume | Key differentiators |
|------|----------|-----------|----------------|---------------------|
| Free | Casual individuals | Free | Limited AI Credits | Basic completions/chat |
| Student | Verified students/teachers | Free | Extended AI Credits | Broader model access |
| Pro | Individual devs | $10/mo | 1,000 Base + 500 Flex (1,500) | Extended IDE integrations, multi-model |
| Pro+ | Intensive individuals | $39/mo | 3,900 Base + 3,100 Flex (7,000) | Large token quotas; GitHub Spark |
| Business | Teams/orgs | $19/user/mo | 1,900 credits/user | Centralized seats, audit logs, exclusions, IP indemnity |
| Enterprise | Large enterprises | $39/user/mo | 3,900 credits/user | Repo indexing, custom fine-tuning, global governance |

| Privacy/SKU | Training on interaction data? | IP indemnity? | Content/file exclusions? | Audit logs? |
|-------------|-------------------------------|---------------|--------------------------|-------------|
| Free/Student/Pro/Pro+ | Yes (opt-out possible) | No | No | No |
| Business/Enterprise | No | Yes | Yes | Yes |

| Additional detail | Value |
|-------------------|-------|
| Enterprise Cloud subscription | $21/user/month |
| Enterprise real minimum | $60/user/month |
| Training policy change | April 2026 (opt-out for individual plans) |
| AI Credits transition | June 2026 (replaces PRU) |
| PAT scopes for admin | `admin:org` |
| Policy inheritance | Enterprise > Organization > user preferences |

## Why this source matters for the RAG

This guide provides detailed, current (2026) information on GitHub Copilot's plan tiers, privacy boundaries, billing transition, and enterprise administration capabilities, making it highly valuable for procurement, governance, and compliance queries. Its comparison tables, pricing figures, and API examples supply precise, retrievable facts for the knowledge base.
