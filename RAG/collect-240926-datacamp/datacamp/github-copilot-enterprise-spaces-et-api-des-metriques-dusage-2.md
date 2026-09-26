---
id: collect-240926-datacamp/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-2
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-10", "2026-02", "2026-03-10", "2026-04", "2026-05-14"]
keywords: ["copilot", "agent", "license", "licenses", "training"]
source: docs/RAG/clean_en/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [164, 411]
sha256: 2e03dd50e7b601eb264ce4080d7c5998bd5dac79fb877d4f537d337152e0aa35
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

- Missing repositories
- Low-quality documentation
- Incorrect permissions
- Insufficient indexing time

### Sharing and access controls

Spaces supports two main visibility models:

- Individual Spaces
- Organization-wide Spaces

Members of an enterprise can see their individual spaces managed by global enterprise settings. Enterprise administrators can also centrally manage early access policies and feature availability.

Private Spaces are suitable for teams experimenting or for sensitive initiatives. Organizational Spaces are ideal for engineering standards, onboarding, or common frameworks.

A common mistake is over-centralization. A single sprawling, global Space quickly becomes noisy and less useful for Copilot.

### Organizing Spaces by team or domain

There is no single universal structure.

Common patterns include one Space per team, one Space per product, or shared standards Spaces. Each has a different scope and uses the same settings in a specific way.

#### One Space per team

Useful when engineering groups operate relatively independently.

Examples:

- Platform engineering
- Data engineering
- Mobile development

#### One Space per product

Useful for organizations structured by products rather than by departments.

Examples:

- Payments
- Analytics
- Infrastructure
- Customer platform

#### A shared standards Space

Many organizations maintain a separate shared Space for:

- Security guidelines
- Code conventions
- Deployment workflows
- Architecture standards

In practice, hybrid approaches generally produce the best results: each team has its own Space, supplemented by large shared standards Spaces.

## The Copilot usage metrics API

Spaces solves the context problem. The usage metrics API solves the measurement problem. It replaced several older telemetry systems that GitHub retired during the 2026 API consolidation.

Without clear metrics, organizations quickly lose sight of the success of Copilot adoption. Leadership wants evidence that the investment improves developer workflows, not just another subscription line.

The dashboard has been generally available since February 2026 and is accessible via your enterprise account → AI Controls → Copilot → Metrics → Copilot usage metrics in the Insights tab.

### What the API measures

The usage metrics API exposes several categories of operational telemetry.

Common metrics include:

- Active users
- Lines of code suggested vs lines of code accepted
- IDE usage profiles
- Model usage
- Agent interactions
- Language breakdown

This offers a much finer view than simple license counts.

A team with 100 assigned licenses but only 15 active users does not have the same adoption profile as a team with regular daily usage and high acceptance rates.

### The API transition in 2026

GitHub retired several legacy telemetry APIs (User-level Feature Engagement Metrics API, Direct Data Access API, Copilot Metrics API) between 2025 and 2026, with a complete shutdown in April 2026.

They included:

- The legacy Metrics API
- The Feature Engagement APIs
- The Direct Data Access APIs

The new usage metrics endpoints, available since February 2026, unified these reporting systems into a more consistent model, with versioning for breaking changes.

This is important because many older GitHub posts and examples still reference deprecated endpoints. Before industrializing your integrations, always check the most recent documentation.

## Querying the usage metrics API

Now that the purpose of the metrics API is clear, let's see how to use it concretely.

### Authentication and permissions

GitHub Copilot usage metrics endpoints generally require a few permissions on your Personal Access Token (PAT), whether classic or fine-grained.

- 
For classic PATs, your enterprise admin must grant you the `manage_billing:copilot` and `read:org` permissions.
- 
For fine-grained tokens, use a GitHub App user access token or installation access token with the `Enterprise Copilot metrics enterprise permissions (read)` permission.

In general, fine-grained tokens are preferable because they limit unnecessary exposure of permissions.

### Organization-level endpoints

The two most common organization-level reports are:

- 
`organization-1-day`
- 
`organization-28-day`

#### One-day organizational report

The one-day report is ideal for operational monitoring and short-term trend analysis. History goes back to October 10, 2025, and remains accessible for one year from the current date.

The curl command below calls the 1-day report API and returns a JSON response with download links. Fill in `YOUR_TOKEN` for the Bearer and choose a `DAY` in `YYYY-MM-DD` format.

```
curl -L \
 -H "Accept: application/vnd.github+json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
-H “X-GitHub-Api-Version: 2026-03-10” \
"https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/enterprise-1-day?day=DAY"
```
The URLs in `download_links` are signed and time-limited: they expire quickly after being generated. Your workflow must retrieve the URL and then download the file immediately, in the same execution.

The response may contain only `download_links` and `report_day`, but here is the full possible schema:

```
{
  "type": "object",
  "title": "Copilot Metrics 1 Day Report",
  "description": "Links to download the Copilot usage metrics report for an enterprise/organization for a specific day.",
  "properties": {
    "download_links": {
      "type": "array",
      "items": {
        "type": "string",
        "format": "uri"
      },
      "description": "The URLs to download the Copilot usage metrics report for the enterprise/organization for the specified day."
    },
    "report_day": {
      "type": "string",
      "format": "date",
      "description": "The day of the report in YYYY-MM-DD format."
    }
  },
  "required": [
    "download_links",
    "report_day"
  ]
}
```
#### 28-day organizational report

The 28-day report highlights adoption patterns and longer-term trends. The commands are almost identical, pointing to the 28-day API.

Example request:

```
curl -L \
 -H "Accept: application/vnd.github+json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
-H “X-GitHub-Api-Version: 2026-03-10” \
https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/enterprise-28-day/latest
```
You will get a similar response, although with `response_start_day` and `response_end_day`.

#### Structure of organizational reports

The 1-day and 28-day JSON reports at the organization level may look like this:

```
[
  {
    "user_id": 1001,
    "user_login": "octocat",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 42,
    "slug": "frontend"
  },
  {
    "user_id": 1001,
    "user_login": "octocat",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 43,
    "slug": "backend"
  },
  {
    "user_id": 1002,
    "user_login": "hubot",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 42,
    "slug": "frontend"
  }
]
```
This gives you an overview of an organization's users, their teams, and their team tags.

### User-level endpoints

User-level reports provide more granular visibility into adoption. You can understand, at a high level, how each person uses Copilot.

Common endpoints:

- 
`users-1-day`
- 
`users-28-day`
- 
`user-teams-1-day`

These reports help administrators identify:

- Highly active users
- Teams with low adoption
- Training needs
- Usage trends by department

These requests look very similar to the 1-day and 28-day reports at the organization level, simply pointing to a different endpoint.

#### One-day user report

Example `users-1-day` call:

```
curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
  "https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/users-1-day?day=DAY"
```
#### 28-day user report

Example `users-28-day` call:

