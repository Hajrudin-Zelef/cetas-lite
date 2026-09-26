---
id: collect-240926-datacamp/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-3
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-01", "2025-11-01", "2026-03-10"]
keywords: ["copilot", "agent", "benchmarks", "governance", "licenses"]
source: docs/RAG/clean_en/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [412, 594]
sha256: 2a39846ddff32100160f019d0b741922e3270b388057ec5a975a03ffb6d590a7
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

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
