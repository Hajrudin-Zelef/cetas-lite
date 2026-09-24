---
id: collect-240926-datacamp/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-01", "2025-10-10", "2025-11-01", "2026-02", "2026-03-10", "2026-04", "2026-05-14"]
keywords: ["copilot", "agent", "benchmarks", "cost", "governance", "license", "licenses", "training"]
source: docs/RAG/clean_en/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [1, 594]
sha256: 7ec5aac7099677f9a9d039f870f369ed093cc82735936e534f12950e1bd11125
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

<!-- source: https://www.datacamp.com/fr/blog/github-copilot-enterprise -->

Course

You've deployed GitHub Copilot Enterprise across the organization, assigned licenses, configured policies, and your developers are already using it in their IDEs. Now you need to answer the hard questions:

- How do you optimize Copilot so it better learns your company's specific engineering context?
- How do you measure Copilot's value? Which departments are adopting it successfully and which are ignoring it entirely?

This is where GitHub Copilot Spaces and the usage metrics API come in. Spaces allows Copilot to ingest your organization's technical knowledge. The usage metrics API helps administrators measure adoption, retention, and productivity trends across the enterprise.

In this article, we'll cover:

- What GitHub Copilot Enterprise includes
- How Copilot Spaces work
- How to configure Spaces at scale
- The GitHub Copilot usage metrics API endpoints
- Authentication and reporting workflows
- Concrete strategies for measuring ROI

If you're not comfortable with GitHub organizations, pull requests, and permission models, the Intermediate GitHub Concepts course covers these fundamentals. If you're also new to Copilot, our How to Use GitHub Copilot tutorial introduces the basic features that this guide builds on.

## Strengthen your data privacy and governance

Ensure compliance and protect your business with DataCamp for Business. Specialized courses and centralized tracking to protect your data.

## What is GitHub Copilot Enterprise?

GitHub Copilot Enterprise sits at the top of GitHub's Copilot offerings.

Compared to GitHub Copilot Business or Pro+, Enterprise emphasizes governance, organizational context, and measurement capabilities. It's designed for enterprises managing large engineering environments, rather than individual developers or small teams.

Two capabilities matter most in practice:

1. Custom organizational context via **Spaces**
2. Organization-wide telemetry via the **usage metrics API**

These two features transform Copilot from simple "smart autocomplete" into something resembling a true internal AI-powered engineering platform.

Enterprises that get the most value from GitHub Copilot Enterprise integrate it as a key element of their internal infrastructure. They curate organizational context, continuously measure adoption, and adjust policies based on usage data, not assumptions.

For a broader overview of the GitHub ecosystem, we recommend our Introduction to GitHub Products guide.

### How Enterprise differs from Business and Pro+

GitHub Copilot Enterprise extends the Business offering with:

- Organization-level usage metrics
- Enhanced governance controls
- Enterprise-wide policy inheritance
- Higher quotas for premium requests (1,000 vs. 300 on Business)
- Additional model access and management

Enterprise requires GitHub Enterprise Cloud in addition to the Copilot Enterprise subscription. This adds an additional per-user cost, so make sure your organization actually needs enterprise-level governance, telemetry, and administration.

| **Feature** | **Pro+** | **Business** | **Enterprise** | 
| Individual use | Yes | No | No | 
| Centralized license management | No | Yes | Yes | 
| Audit logs | No | Yes | Yes | 
| File exclusions | No | Yes | Yes | 
| Spaces support | Yes, with Copilot | Yes, limited admin visibility | Yes, full enterprise-level management | 
| Usage metrics API | No | Organization level | Enterprise + organization level | 
| Enterprise policy inheritance | No | No | Yes | 

**Note:** Business subscribers access the usage metrics API at the organization level (`/orgs/{org}/…`). Enterprise subscribers additionally have access to aggregated reports at the enterprise level (`/enterprises/{enterprise}/…`) covering all organizations in a single view.

### Who GitHub Copilot Enterprise is for

GitHub Copilot Enterprise is aimed at organizations with mature GitHub environments.

Typical Enterprise customers:

- Large engineering organizations
- Regulated industries
- Multi-team platform teams
- Enterprises with internal development standards
- Organizations requiring centralized governance

Note that this doesn't inherently improve Copilot's performance. This distinction is important: many teams initially oversize their purchase thinking that Enterprise = "better Copilot," when Enterprise mainly adds governance and measurement tools.

## Copilot Spaces: custom context for your organization

Copilot Spaces solves one of the main limitations of general-purpose code assistants.

By default, Copilot has a good grasp of public programming knowledge. It doesn't automatically understand your internal APIs, architectural decisions, code conventions, deployment workflows, or onboarding documentation.

Spaces provides curated organizational context that Copilot can leverage in conversation and for coding assistance.

In practice, Spaces helps Copilot answer questions such as:

- "How do we structure our API handlers internally?"
- "Which authentication library does our platform team recommend?"
- "Which deployment workflow should this microservice use?"
- "What naming conventions does our backend team follow?"

### What Spaces support

Spaces covers a broader range of organizational content than the old Knowledge Bases system.

Supported content types:

- Code files
- Markdown documentation
- JSON files
- Uploaded files
- Images
- GitHub Issues
- Pull requests

Each content type provides a different kind of value.

Code files help Copilot understand implementation patterns. Markdown files detail architecture and onboarding. Pull requests expose review discussions and past engineering decisions. Together, they improve understanding of your organization's development practices.

A subtle but important point: Spaces is not just a vector store backed by GitHub. It includes sharing controls and governance workflows designed for the enterprise.

### The end of Knowledge Bases

GitHub retired the old Copilot Knowledge Bases feature on November 1, 2025.

Spaces replaces Knowledge Bases with:

- Broader content support
- Better sharing controls
- Improved administration
- More flexible organization-level management

You will still find outdated documentation and blog posts mentioning Knowledge Bases. Be careful with old tutorials: many endpoints and workflows changed between 2025 and 2026.

## Creating and configuring Copilot Spaces

On the administration side, creating a Copilot Space is fairly straightforward. The challenge is managing dozens, or even hundreds, of them across teams.

The initial structure tends to persist. I've seen organizations inadvertently create a "documentation jungle" in Spaces due to a lack of clear ownership rules from the start.

Anyone can create a Copilot Space, so let's test it in a personal repository. The steps are similar at the Enterprise level, with a few different pages.

### Configuring a Space

Creation generally follows this workflow:

1. Go to the Copilot Spaces page in the Enterprise admin area
2. Create a new Space

1. Select repositories and content sources, including MCPs and other useful tools

1. Add sources via the "+ Add sources" button on the right

1. Choose to share the Space or set sharing settings at this stage

1. Verify that Copilot can reference the content during conversations

Note for Enterprise users: your administrator may disable sharing of personal Spaces. If you are using your own account, this may limit sharing of a Copilot Space that does not use the company's repositories.

After configuration, administrators should test the Space with concrete prompts.

For example:

`How does our authentication middleware handle token refresh logic?`
Or:

`Show me an example of how our backend services structure database migrations.`
If Copilot cannot answer accurately, the cause is usually:

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

```
curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
   https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/users-28-day/latest
```
#### user-teams report for a single day

A `user-teams-1-day` endpoint also exists and associates each user with their teams. It does not contain usage metrics: it serves as a join key to aggregate data by team.

#### Structure of user-level reports

The level of detail is much higher, since it concerns the usage of a given user:

```
[{
  "code_acceptance_activity_count": 1,
  "code_generation_activity_count": 1,
  "day": "2025-10-01",
  "enterprise_id": "1",
  "loc_added_sum": 8,
  "loc_deleted_sum": 0,
  "loc_suggested_to_add_sum": 10,
  "loc_suggested_to_delete_sum": 0,
  "totals_by_cli": {
    "last_known_cli_version": {
      "cli_version": "1.0.8",
      "sampled_at": "2025-10-01T00:01:43.000Z"
    },
    "prompt_count": 2,
    "request_count": 2,
    "session_count": 2,
    "token_usage": {
      "avg_tokens_per_request": 4400.0,
      "output_tokens_sum": 5000,
      "prompt_tokens_sum": 3800
    }
  },
  "totals_by_feature": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "feature": "code_completion",
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0,
    "user_initiated_interaction_count": 0
  }],
  "totals_by_ide": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "ide": "vscode",
    "last_known_ide_version": {
      "ide_version": "1.85.0",
      "sampled_at": "2025-10-01T00:00:02.000Z"
    },
    "last_known_plugin_version": {
      "plugin": "",
      "plugin_version": "",
      "sampled_at": "2025-10-01T00:00:02.000Z"
    },
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0,
    "user_initiated_interaction_count": 0
  }],
  "totals_by_language_feature": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "feature": "code_completion",
    "language": "unknown",
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0
  }],
  "totals_by_language_model": [],
  "totals_by_model_feature": [],
  "used_agent": false,
  "used_chat": false,
  "used_cli": true,
  "user_id": 1,
  "user_login": "login1",
  "user_initiated_interaction_count": 0,
  "etl_id": "green",
  "day_partition": "2025-10-01",
  "entity_id_partition": 1
}]
```
These metrics are mostly useful as adoption signals at the team level. Acceptance rates and usage volumes are operational indicators, not measures of developer quality.

To see the full potential set of exposed metrics, refer to the most up-to-date GitHub usage metrics data documentation.

User-level reports include CLI interactions. If your teams use Copilot on the command line, our GitHub Copilot CLI Tutorial covers setup and common workflows.

## Setting up a Copilot reporting workflow

Calling the API by hand is useful for experimenting and understanding the schema. To take action, it is better to automate.

Teams that get the most value from Copilot Enterprise generally build lightweight reporting pipelines that cross-reference usage telemetry with their internal engineering metrics.

### The key indicators to demonstrate ROI

Not all Copilot metrics are equal. The most useful include:

- Growth in the number of active users
- Trend in acceptance rates
- Code suggested vs code kept
- Reduction in PR cycle times
- Frequency of IDE usage

GitHub has published benchmarks such as:

- 55% faster task completion
- 88% of code kept

These figures indicate significant productivity gains. Your results will vary by team and workflow, which is why the usage metrics API is valuable. A backend infrastructure team does not use Copilot the same way as a frontend prototyping team.

### From raw data to a team dashboard

A lightweight reporting workflow often looks like this:

1. Scheduled API call
2. Storage of responses in a database or spreadsheet
3. Transformation into reporting tables
4. Visualization in your BI platform

The tech stack matters less than consistency.

Even a simple chain of scheduled Python scripts and CSV exports can provide useful operational visibility.

Typical architecture:

GitHub API

↓

Scheduled Python script

↓

PostgreSQL / CSV / Spreadsheet

↓

Power BI / Tableau / Looker

## To conclude

GitHub Copilot Enterprise is about preparing your infrastructure for "AI-ready" code. Spaces provides the organizational context that makes Copilot more relevant in real engineering environments. The usage metrics API provides the telemetry needed to evaluate the success of adoption.

The organizations that get the best results with Copilot Enterprise have common traits:

- They take care of internal context
- They track adoption continuously
- They take Copilot governance seriously
- They measure outcomes rather than assuming productivity gains

This mindset matters far more than assigning licenses.

To deepen your Copilot and AI skills, we recommend the Software Development with GitHub Copilot course or the AI for Software Engineering skills path.

## GitHub Copilot FAQ

### What is GitHub Copilot Spaces?

**GitHub Copilot Spaces are curated collections of repositories, documentation, issues, and other organizational content that ground Copilot's responses in company-specific knowledge.**

### What did GitHub Copilot Knowledge Bases replace?

**GitHub ended Knowledge Bases on November 1, 2025. Spaces is the replacement system, with broader content support and better sharing controls.**

### What does the GitHub Copilot usage metrics API track?

**The API tracks active users, code suggestions, acceptance rates, usage by language, IDE telemetry, and other organizational adoption metrics.**

### What permissions are required for the usage metrics API?

**Most usage metrics API endpoints require permissions such as `manage_billing:copilot` or `read:org`, depending on the authentication model and the endpoint used.**

I am a data scientist with experience in spatial analysis, machine learning, and data pipelines. I have worked with GCP, Hadoop, Hive, Snowflake, Airflow, and other data engineering and data science processes.
