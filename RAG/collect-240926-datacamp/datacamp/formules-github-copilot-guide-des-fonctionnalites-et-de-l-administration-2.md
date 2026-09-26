---
id: collect-240926-datacamp/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-2
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: ["2026-04"]
keywords: ["agent", "agentic", "agents", "attention", "claude", "cloud agent", "copilot", "cost", "fine-tuning", "governance", "latency", "liability"]
source: docs/RAG/clean_en/datacamp/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [93, 199]
sha256: 373e2b4c7b3594efda457b81246189b9a1a609b65d667de41b6848fd493cd3c2
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

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

