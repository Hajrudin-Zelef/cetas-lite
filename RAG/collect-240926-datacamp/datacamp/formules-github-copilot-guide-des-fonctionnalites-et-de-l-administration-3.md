---
id: collect-240926-datacamp/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-3
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2022-11-28", "2026-05-01", "2026-05-31"]
keywords: ["agentic", "copilot", "cost", "governance", "license", "pricing", "training"]
source: docs/RAG/clean_en/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [200, 328]
sha256: 1cba01e7fa4e25072a22c33b8cb0b960f86fd0fe6ae7344e6af20e026878a24d
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

Enterprise administrators can lock policies globally across all entities, allow selective overrides by organization, or fully delegate control down the hierarchy. For example, the enterprise can enforce global settings to restrict the use of certain models.

At the team level, it can impose strict public code filters on the financial services division, while allowing more experimentation in the internal software R&D division.

## Audit logs

When compliance auditors request verification of your software supply chain, or security teams need to trace a data leak, GitHub Copilot logs platform changes.

### Copilot events in the audit log

The system keeps a complete record of administration operations, including:

- Explicit seat assignments and revocations, billing group changes.
- Changes to the public code duplication filter.
- Changes to file and directory exclusion patterns.
- Feature activation states (e.g., enabling agentic code review modes)

The level of granularity depends on your subscription. Business plans focus on action streams scoped to the organization, while Enterprise accounts provide access to cross-organization forensic telemetry.

### Search, filtering, and export

Audit log streams are accessible natively via the Organization Settings panel. Administrators can query the interface with specific action qualifiers:

```
# Filter to identify who changed Copilot access rights
action:copilot.cfb_seat_assignment_created
# Identify changes to global exclusions within a period
action:copilot.content_exclusion_updated created:2026-05-01..2026-05-31
```
Enterprise accounts allow these audit events to be streamed to external SIEMs (such as Splunk or Datadog) for automated alerting and centralized, immutable retention.

## Managing Copilot seats with the REST API

Manually assigning seats from a dashboard works for small teams, but does not scale with massive onboarding flows. Using the seats endpoints of the GitHub Copilot REST API allows you to handle identity and access entirely as code.

This is one of my favorite parts of Copilot administration, because it turns license management into a process that engineering teams can automate cleanly.

### Key API endpoints

Common API workflows include:

- List seat assignments
- Assign seats
- Remove seats
- Retrieve usage metrics
- Read the organization's Copilot settings

Authentication generally requires:

- Fine-grained personal access tokens
- GitHub App permissions
- Organization administrator privileges

To access these administration functions, your integration scripts must authenticate with a Personal Access Token (PAT) that has the admin:org scopes or run through an authorized GitHub App with explicit Copilot management privileges at the organization level.

To go deeper into programmatic platform integrations, I recommend following our GitHub Foundations skills path.

### Common automation patterns

Practical patterns include:

- 
**Automated identity onboarding:** connect an HRIS (Workday, Okta, etc.) directly to GitHub via webhooks. When an engineer joins a given team, a script triggers a `POST` request to automatically provision Copilot for them.
- 
**Reclaiming inactive seats:** a scheduled Cron script queries active usage via the API. If a user has not used Copilot for more than 30 days, the script executes a `DELETE` to reclaim the license and preserve the credit pool.
- 
**Financial dashboards:** daily extraction of allocation and consumption data to feed internal BI platforms (such as Tableau) and facilitate chargeback by cost center.

#### Example: assigning a Copilot seat with Python

The following script shows how to programmatically assign an organization seat to a specific developer with Python:

```
	import requests
	# Identity Configuration
TOKEN = "YOUR_ORGANIZATION_ADMIN_PAT"
ORG = "your-corporate-org"
USERNAME = "target-developer-user"
url = f"https://api.github.com/orgs/{ORG}/copilot/billing/selected_users"
headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Accept": "application/vnd.github+json",
    "X-GitHub-Api-Version": "2022-11-28"
}
payload = {
    "selected_usernames": [USERNAME]
}
response = requests.post(url, json=payload, headers=headers)
if response.status_code == 201:
    print(f"Successfully allocated Copilot seat to {USERNAME}.")
else:
    print(f"Failed allocation. Status: {response.status_code}")
    print(response.json())
```
## Final thoughts

The structure of GitHub Copilot plans looks simple from the pricing page. Once you're managing teams, the differences become much more substantial.

Privacy boundaries, training policies, auditability, and governance controls often weigh more heavily than raw access to models. That's why discussions of GitHub Copilot Business vs Enterprise usually become conversations about security and operations, rather than simple engineering topics.

If I were advising a team today, I would start from governance requirements:

- Do you need contractual privacy guarantees?
- Do you need audit logs?
- Do you need centralized policy management?

Then I would optimize usage volume and feature access.

To develop your team's technical skills and prepare for official certifications, explore these advanced paths:

## FAQ on GitHub Copilot plans

### What is the difference between GitHub Copilot Business and Enterprise?

**Business includes centralized seat management, audit logs, IP indemnification, and policy controls. Enterprise adds enterprise-wide policy inheritance and expanded governance features.**

### Does GitHub Copilot train its models on code from private repositories?

**No. GitHub states that code from private repositories is not used directly for training. However, interaction data from individual plans may be collected unless the user opts out. Business and Enterprise plans contractually prevent training on interaction data.**

### What are GitHub Copilot audit logs used for?

**Audit logs help administrators track seat assignments, policy changes, feature enablements, and governance activity within the organization.**

### What is file exclusion in GitHub Copilot?

**File exclusion prevents Copilot from accessing specified files or directories for completions, chat, and AI-generated suggestions. This feature is available only with Business and Enterprise plans.**

I am a data scientist with experience in spatial analysis, machine learning, and data pipelines. I have worked with GCP, Hadoop, Hive, Snowflake, Airflow, and other engineering and data science processes.
