---
id: collect-240926-datacamp/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-1
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-04", "2026-06"]
keywords: ["agentic", "copilot", "cost", "fine-tuning", "governance", "license", "licenses", "training"]
source: docs/RAG/clean_en/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [1, 92]
sha256: 40328e7bff6144d2c315df7ef281e6963195e884bf151dbf44e8156ead6b3129
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

<!-- source: https://www.datacamp.com/fr/blog/github-copilot-plans -->

Curriculum

Your team has just secured a budget to deploy GitHub Copilot across the engineering organization. To get the most out of it, you need to understand how policy settings, file exclusions, and audit log queries fit together, because that is where the platform's real value lies.

The configuration surface is vast because the needs are just as vast. A solo developer on personal projects does not face the same privacy and compliance challenges as an enterprise administrator managing thousands of seats across regulated repositories. GitHub Copilot's tiered plan structure is designed to cover that spectrum.

This guide covers each Copilot plan tier, the privacy and intellectual property (IP) boundaries that distinguish them, and the administration mechanisms needed to scale within an organization.

Before diving into administration, you should already be comfortable with GitHub organizations, repositories, and the permissions system. If you are new to the ecosystem, start with our guide How to Use GitHub Copilot.

If you are still weighing Copilot against the rest of the market, our selection of the 13 best AI coding assistants in 2026 covers the entire competitive landscape. For a focused comparison with one of the main competitors, see our guide Cursor vs. GitHub Copilot.

## In brief

- GitHub offers four individual tiers (Free, Student, Pro, Pro+) and two organizational tiers (Business and Enterprise) for Copilot, each with distinct privacy, governance, and usage perimeters.
- The Business and Enterprise plans provide contractual guarantees that interaction data is never used for training, while individual plans are opt-out by default as of April 2026.
- Choose your GitHub Copilot plan first based on your compliance and governance requirements; then optimize model selection and usage quotas.
- File exclusion rules and organization-wide policy settings are only available on Business and Enterprise, making them the foundation for teams working with proprietary code.
- GitHub Copilot Enterprise requires an active GitHub Enterprise Cloud subscription, bringing the real minimum cost to $60 per user per month.
- Seat management, audit log queries, and policy enforcement can be automated via the REST API, turning licenses into infrastructure-as-code.

## Overview of GitHub Copilot plans

GitHub offers several tiers for its ecosystem. Notably, the platform is finalizing its transition to usage-based billing, replacing the former "Premium Request Unit" (PRU) framework with GitHub AI Credits in June 2026.

In the new system, basic code completions and "Next Edit" suggestions remain unlimited and do not consume credits.

By contrast, advanced operations such as multi-file chat, agentic workflows, long coding sessions, and in-depth code reviews consume AI Credits based on tokens (input, output, and cache) relative to the published API rates of the model concerned.

Base monthly subscription prices have remained stable, but this change alters how administrators budget for overages and track actual usage.

| **Plan tier** | **Target audience** | **Base price** | **Included monthly volume** | **Key differentiators** | 
| Free | Occasional individual users | Free | Limited AI Credits | Basic access to completions and Chat. | 
| Student | Verified students and teachers | Free | Extended AI Credits | Broader model access for learning. | 
| Pro | Individual developers | $10 / month | 1,000 Base + 500 Flex (1,500 total) | Extended IDE integrations and multi-model support. | 
| Pro+ | Intensive individual users | $39 / month | 3,900 Base + 3,100 Flex (7,000 total) | Large token quotas; includes access to GitHub Spark. | 
| Business | Teams and organizations | $19 / user / month | 1,900 credits / user (3,000 from June 1 to Sept. 1, 2026) | Centralized seat management, audit logs, file exclusions, IP indemnification. | 
| Enterprise | Large enterprises | $39 / user / month | 3,900 credits / user (7,000 from June 1 - Sept. 1, 2026) | Repository indexing, custom fine-tuning, global governance. | 

### Individual plans: Free, Student, Pro, and Pro+

Individual plans differ in model access, usage limits, and experimental capabilities. For example, the Free plan allows basic exploration, while Pro+ provides access to GitHub Spark, an environment designed for building AI-assisted applications.

Currently, new sign-ups for paid individual GitHub accounts, such as Pro, Pro+, and Student, are suspended. Existing accounts can upgrade from Pro to Pro+, but new accounts cannot sign up until GitHub has completed the transition to AI Credits billing.

### Business and Enterprise

With Business and Enterprise, GitHub Copilot plans shift from a simple IDE extension to a true enterprise infrastructure asset, fully auditable.

GitHub Copilot Business introduces essential management features:

- Centralized seat assignment and removal.
- Organization-wide reference policies.
- Structured audit logs and compliance event tracking.
- Repository content and file exclusions.
- Commercial intellectual property indemnification.

GitHub Copilot Enterprise goes further in terms of control and capabilities:

- Copilot Spaces: a knowledge hub that allows you to query Copilot about internal documentation, wikis, and system code standards.
- Enhanced integration of GitHub.com Chat.
- Hierarchical policy inheritance between child organizations.

GitHub Copilot Enterprise requires an active GitHub Enterprise Cloud subscription. Since GitHub Enterprise Cloud costs $21 per user per month and the Copilot Enterprise license is $39 per user per month, the actual minimum cost is $60 per user per month for Enterprise. This does not apply to the GitHub Copilot Business plan, which can be purchased natively by organizations using GitHub Free or GitHub Team.

Organizations do not have access to Enterprise benefits such as policy inheritance, but they still have IP indemnification, auditing, file exclusion, and organization-level policy management; a good alternative for mid-sized engineering teams.

If you are considering an Enterprise subscription, our GitHub Copilot Enterprise guide will show you how to leverage its features, such as Copilot Spaces and the new Usage Metrics API.

## What sets individual plans apart from Business plans

Data management, IP indemnification, and billing are the main areas where individual and Business plans differ significantly. The additional features for users are useful, but understanding these gaps is essential when deciding between managing a stack of personal Pro licenses and subscribing to a Business plan.

### Data processing and default training settings

For teams managing proprietary systems, data confidentiality is generally the deciding factor between personal plans and a Business subscription.

In April 2026, GitHub changed interaction data collection for individual Copilot plans. For Free, Pro, and Pro+ users, interaction data may now be used for model training by default, unless the user explicitly opts out.

Let's clarify the difference between code at rest and interaction data, so you know what is used for AI training:

- **Code at rest:** the raw code present in your private repository is not read or incorporated into public training sets.
- **Interaction data:** includes prompts, chat requests, cursor context, surrounding code blocks transmitted via the IDE API during active sessions, suggestion acceptance metrics, and feedback logs.

Business and Enterprise contracts strictly guarantee that interaction data is never used for training purposes, under any circumstances. No manual user action is required.

To dive deeper into data usage and troubleshooting in Copilot, read our GitHub Copilot: privacy and troubleshooting guide.

### Intellectual property indemnification

