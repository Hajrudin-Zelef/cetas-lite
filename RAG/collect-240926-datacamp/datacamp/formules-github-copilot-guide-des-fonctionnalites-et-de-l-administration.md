---
id: collect-240926-datacamp/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: ["2022-11-28", "2026-04", "2026-05-01", "2026-05-31", "2026-06"]
keywords: ["agent", "agentic", "agents", "attention", "claude", "cloud agent", "copilot", "cost", "fine-tuning", "governance", "latency", "liability"]
source: docs/RAG/clean_en/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [1, 328]
sha256: b657197421953e902e420d8ae9fd3dc87d14e8debbbf4d4d9e8e9266e850e45d
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

GitHub Copilot Business and Enterprise include intellectual property (IP) indemnification for generated code. Individual plans do not benefit from it.

In practical terms, indemnification means that GitHub contractually commits to providing legal protection under specified circumstances if the generated code leads to IP disputes. This does not eliminate all legal risks, but it changes the liability discussion for teams delivering commercial software.

A freelancer who delivers code to clients should pay attention to this. The difference between a "personal productivity tool" and an "organization-backed development platform" becomes very concrete once contracts and commercial deliverables come into play.

### Billing, seats, and the shift to AI Credits

Individual billing is self-service and tied to personal accounts. Business plans centralize billing with seats assigned by the administrator. Moreover, instead of each user managing an independent pool of credits, the organization pools its monthly AI Credits based on the number of users.

Enterprise plans go further, with granular budget enforcement limits, cost center groupings, and department-level allocations to prevent a single development group, through intensive agentic workflows, from exhausting the entire company's credit stock.

## SKUs and privacy considerations

Understanding privacy protections and SKUs is essential. The architectural boundaries governing data flows, legal protections, and tracking across the different tiers are summarized below:

| **Plan tier** | **Interaction data used for training?** | **Contractual IP indemnification?** | **Content / file exclusions?** | **Access to audit logs?** | 
| Free | Yes (opt-out possible) | No | No | No | 
| Student | Yes (opt-out possible) | No | No | No | 
| Pro | Yes (opt-out possible) | No | No | No | 
| Pro+ | Yes (opt-out possible) | No | No | No | 
| Business | No | Yes | Yes | Yes | 
| Enterprise | No | Yes | Yes | Yes | 

### April 2026 training policy changes

The shift from an opt-in model to an opt-out framework for individual plans constitutes a major compliance leakage vector. The interaction data payload automatically captured during an active IDE session includes:

- Detailed chat histories and prompt context.
- Multi-line code suggestions and local acceptance rates.
- The active editor’s cursor context, which often retrieves context from adjacent files, import statements, and variable declarations from open tabs.

Imagine a developer using a personal Copilot Pro account in a corporate repository. If training remains enabled, interaction data related to that session may enter GitHub’s training ecosystem. This is a common reason why organizations adopt Business plans.

### Choosing the right SKU based on your privacy requirements

Depending on the nature of the work, you will not need the same SKU.

- **Solo developer / personal projects:** Free or Pro plans offer maximum flexibility. Simply disable training in your privacy settings if you are working on proprietary code.
- **Freelancers / subcontractors:** The Business plan provides a protective barrier. Client contracts often prohibit sending data to external LLM providers; a dedicated seat within the organization protects your commitments.
- **Enterprise teams with compliance obligations**: The Business plan is the standard baseline, ensuring isolation of data flows and administrative governance.
- **Regulated sectors (finance, healthcare):** The Enterprise plan is generally essential, enabling integration with specialized security configurations, strict data residency requirements, and localized fine-tuning layers.

## Excluding specific files from Copilot

Setting up file exclusion rules in GitHub Copilot is one of the most effective ways to secure your environment. Content exclusion prevents the local IDE agent from processing certain files, making them completely invisible to inline completions, chat dialogs, and background agentic operations.

Note that GitHub Copilot CLI, the Copilot cloud agent, and Agent mode in Copilot Chat in the IDEs do not support content exclusion.

### Configuring exclusion rules

Administration teams can apply exclusions at the global organization settings level or at the level of targeted repositories. Simply open the repository or organization settings by clicking the Settings button in the top right.

Choose “Code and automation” in the Copilot settings in the sidebar. Then enter your exclusions in the “Paths to exclude in this repository” area as follows:

```
# Ignore the /src/some-dir/kernel.rs file in this repository.
- "/src/some-dir/kernel.rs"
# Ignore files called secrets.json anywhere in this repository.
- "secrets.json"
# Ignore all files whose names begin with secret anywhere in this repository.
- "secret*"
# Ignore files whose names end with .cfg anywhere in this repository.
- "*.cfg"
# Ignore all files in or below the /scripts directory of this repository.
- "/scripts/**"
```
**Organization-level configuration is similar, except that the option is found under “Repositories and Paths to exclude” and uses the following format:**

```
REPOSITORY-REFERENCE:
  - "/PATH/TO/DIRECTORY/OR/FILE"
  - "/PATH/TO/DIRECTORY/OR/FILE"
  - …
```
Keeping the `REPOSITORY-REFERENCE` is an integral part of the configuration. Common configuration bases should prioritize sensitive identifiers, production orchestration profiles, sensitive proprietary algorithmic modules, or highly regulated folders.

### How exclusions apply to Copilot features

When an exclusion match is detected, data isolation is complete across all Copilot subsystems:

- **Inline completions:** inability to generate context within the file or draw from it to feed adjacent files.
- **Copilot chat / agents:** the system returns a notice indicating that the file cannot be examined due to organization policies.

Standard local IDE engines work the same way. Convenience features such as text parsing, syntax highlighting, and localized IntelliSense compile normally, because the exclusion layer applies explicitly to Copilot’s external telemetry flows.

Administrators should carefully test path patterns in staging repositories; malformed wildcards can fail in “open” mode and expose data you intended to isolate.

## Organization-wide policy management

Applying GitHub Copilot policies at the organization level ensures that enterprise security is defined by the administration team, not by individual developer preferences.

### Available policy settings

Organizations can control several settings for developers:

- **Feature activation:** globally enable or disable Copilot Chat in development environments, command-line interfaces (via Copilot CLI), or advanced agentic code review systems.
- **Public code filter:** legal mechanism that prevents Copilot from suggesting code too similar to public open source repositories on GitHub, reducing the risk of license non-compliance.
- **Model choice restrictions:** limit the models (for example, specific variants of GPT or Claude) that developers can select, in order to manage latency, credit consumption, and performance. For an overview of the models available on the GitHub platform, see this practical guide to GitHub Models.
- **Custom organization instructions:** inject standard markdown files that add code conventions, security frameworks, and architecture paradigms to every request issued by your developers.

If your team is less familiar with the GitHub organization and permissions model, the Intermediate GitHub Concepts course provides a good foundation. For engineering teams rolling out command-line tools at scale, see our GitHub Copilot CLI Tutorial.

### Policy inheritance at the Enterprise level

In large enterprise environments, the policy engine follows a strict inheritance cascade: Enterprise policy > organization policy > user preferences

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
