---
id: vague2-datacamp/datacamp/github-copilot-enterprise
title: "GitHub Copilot Enterprise : Spaces et API des métriques d'usage"
domain: datacamp
role: reference
task: article
actors: ["Microsoft"]
dates: ["2025-11-01", "2026-02", "2026-03-10", "2026-04", "2026-09-23"]
keywords: ["copilot", "agent", "benchmark", "benchmarks", "cost", "governance", "license", "mcp"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/github-copilot-enterprise.md
source_anchor: ""
source_lines: [1, 75]
sha256: 82d11d3754bd530afe7ba43b1d4b747320927f83c51f50019852608a38e2bc82
---

# GitHub Copilot Enterprise : Spaces et API des métriques d'usage

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/github-copilot-enterprise
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Tim Lu focuses on two key GitHub Copilot Enterprise capabilities: Copilot Spaces (organizational context) and the usage metrics API (adoption measurement). It addresses the challenges enterprises face after deployment: how to make Copilot learn company-specific engineering context, and how to measure adoption and ROI.

GitHub Copilot Enterprise sits at the top of the Copilot lineup, emphasizing governance, organizational context, and measurement. It extends Business with org-level usage metrics, enhanced governance controls, enterprise-wide policy inheritance, higher premium request quotas (1,000 vs 300 for Business), and additional model access/management. Enterprise requires GitHub Enterprise Cloud in addition to the Copilot Enterprise subscription, adding cost. The article notes Enterprise does not intrinsically improve Copilot's performance — it adds governance and measurement tools.

A comparison table contrasts Pro+, Business, and Enterprise across individual use, centralized license management, audit logs, file exclusions, Spaces support, usage metrics API, and enterprise policy inheritance. Business subscribers get org-level usage metrics (`/orgs/{org}/...`); Enterprise adds aggregated enterprise reports (`/enterprises/{enterprise}/...`) across all organizations.

Copilot Spaces provides curated organizational context that Copilot can use in conversation and coding assistance — internal APIs, architecture decisions, code conventions, deployment workflows, onboarding docs. Spaces supports code files, Markdown, JSON, uploaded files, images, GitHub Issues, and pull requests, with enterprise-grade sharing controls and governance. Spaces replaced the deprecated Copilot Knowledge Bases (retired November 1, 2025).

The guide explains creating/configuring Spaces (access admin page, create, select repos/content sources including MCP servers, add sources, configure sharing, test with prompts), sharing models (individual vs organization-wide), and organization patterns (per team, per product, shared standards, hybrid).

The usage metrics API replaced several older telemetry APIs (User-level Feature Engagement Metrics API, Direct Data Access API, Copilot Metrics API) retired by April 2026. New endpoints launched February 2026. It measures active users, lines proposed vs accepted, IDE usage profiles, model usage, agent interactions, and language breakdown. The dashboard is GA since February 2026.

Authentication uses classic PATs (`manage_billing:copilot`, `read:org`) or fine-grained tokens (`Enterprise Copilot metrics enterprise permissions (read)`). Endpoints include org-level `organization-1-day` and `organization-28-day`, and user-level `users-1-day`, `users-28-day`, `user-teams-1-day`. Reports return signed, time-limited download links. The article provides curl examples and JSON report schemas, and outlines a reporting workflow (scheduled API call → storage → reporting tables → BI visualization) with ROI indicators (active user growth, acceptance rates, code proposed vs retained, PR cycle time reduction, IDE usage frequency) and GitHub benchmarks (55% faster task completion, 88% code retained).

## Key points

- Copilot Enterprise's two key capabilities: Spaces (organizational context) and the usage metrics API (measurement).
- Enterprise extends Business with org-level metrics, policy inheritance, higher quotas (1,000 vs 300 premium requests), and more model access.
- Enterprise requires GitHub Enterprise Cloud in addition to the Copilot subscription.
- Enterprise does not intrinsically improve Copilot performance — it adds governance/measurement.
- Copilot Spaces supports code, Markdown, JSON, uploads, images, Issues, and PRs with enterprise sharing controls.
- Spaces replaced Knowledge Bases, retired November 1, 2025.
- Usage metrics dashboard is GA since February 2026; new API endpoints replaced older telemetry APIs retired by April 2026.
- Metrics include active users, LOC proposed/accepted, IDE profiles, model usage, agent interactions, language breakdown.
- Endpoints: `organization-1-day`, `organization-28-day`, `users-1-day`, `users-28-day`, `user-teams-1-day`.
- Authentication: classic PAT scopes `manage_billing:copilot` and `read:org`, or fine-grained `Enterprise Copilot metrics enterprise permissions (read)`.
- Download links in reports are signed and time-limited.
- GitHub benchmarks: 55% faster task completion, 88% code retained.

## Technical data / figures

| Feature | Pro+ | Business | Enterprise |
|---------|------|----------|-----------|
| Individual use | Yes | No | No |
| Centralized license management | No | Yes | Yes |
| Audit logs | No | Yes | Yes |
| File exclusions | No | Yes | Yes |
| Spaces support | Yes (with Copilot) | Yes (limited admin visibility) | Yes (full enterprise management) |
| Usage metrics API | No | Organization level | Enterprise + organization level |
| Enterprise policy inheritance | No | No | Yes |

| API endpoint | Level |
|--------------|-------|
| `organization-1-day` | Org (single-day) |
| `organization-28-day` | Org (28-day) |
| `users-1-day` | User (single-day) |
| `users-28-day` | User (28-day) |
| `user-teams-1-day` | User-team mapping |

| Detail | Value |
|--------|-------|
| Premium requests (Enterprise vs Business) | 1,000 vs 300 |
| Knowledge Bases retirement | November 1, 2025 |
| Usage metrics dashboard GA | February 2026 |
| Old telemetry APIs retired | April 2026 |
| API version header | `X-GitHub-Api-Version: 2026-03-10` |
| Task completion benchmark | 55% faster |
| Code retained benchmark | 88% |

## Why this source matters for the RAG

This guide provides in-depth, current (2026) technical and governance detail on GitHub Copilot Enterprise's Spaces and usage metrics API, including endpoints, authentication, report schemas, and ROI measurement, making it valuable for enterprise AI-adoption and observability queries. Its concrete API examples and metrics definitions supply precise, retrievable facts for the knowledge base.
