---
id: collect-240926-datacamp/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage-1
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "agents", "attribution", "cloud agent", "guardrails", "license", "training"]
source: docs/RAG/clean_en/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [1, 114]
sha256: 8bdb492d1d457d4b33396a3e9d2021c96cfdb42b5c57d57e72d15a3120f9400c
---

# confidentialite-de-github-copilot-protections-et-guide-de-depannage

<!-- source: https://www.datacamp.com/fr/blog/github-copilot-privacy-and-troubleshooting -->

Course

As soon as you enable GitHub Copilot for a team, configuration questions come one after another: what data leaves the IDE? Which repositories should be excluded from suggestions? What happens when a suggestion matches public code?

Whether you're a developer or an administrator deploying Copilot across the organization, mastering privacy and security aspects is part of good use of the tool.

In this article, we review how Copilot manages your data and how to configure privacy settings, content exclusions, and guardrails like the duplication filter, as well as the most effective troubleshooting steps when something stops working.

## How GitHub Copilot handles your data

As we explain in our GitHub Copilot tutorial: best practices, Copilot sends a snapshot of the surrounding code context to GitHub's servers as soon as you type in your IDE. The model processes this context and returns a suggestion: it's this interaction data that powers the experience.

For users of individual plans (Free, Pro, and Pro+), GitHub's privacy statement authorizes the use of this interaction data for model training. Users can opt out at any time in their personal privacy settings. Business and Enterprise plans are governed by separate contractual terms that completely exclude the use of interaction data for training — no action required from users.

### What constitutes interaction data

Code from private repositories stored at rest is not used. However, interaction data generated when you actively use Copilot in a private repository may be used for training, unless you opt out. But what exactly do "interaction data" cover?

When you use Copilot, the system collects several types of signals to improve its assistance:

- **Inputs and prompts:** All commands or questions sent to Copilot Chat or the CLI.
- **Outputs:** The model's code suggestions or text responses, including your acceptance or rejection.
- **Code context:** The snippets surrounding the cursor position and the content of open files used to provide relevant suggestions.
- **Metadata and structure:** File names, repository structure, and your navigation in the IDE.
- **User feedback:** Your ratings (thumbs up/down) and your comments.

### Processing differences by plan

GitHub offers several account tiers: Free, Pro, Business, and Enterprise. Data processing varies depending on the account type, as shown in the table below.

| **Dimension** | **Free / Pro / Pro+** | **Business** | **Enterprise** | 
| Use for model training | Requires opt-out | No. Contractually excluded | No. Contractually excluded | 
| Private repository code at rest | Not used | Not used | Not used | 
| Prompt/output retention | IDE: not retained. Outside IDE: 28 days | IDE: not retained. Outside IDE: 28 days | IDE: not retained. Outside IDE: 28 days | 
| Admin controls | Individual only | Organization-level policies and license management | All Business controls, plus enterprise-wide policy inheritance and audit logs | 
| Content exclusions | Not available | Available at repo and organization level | Available enterprise-wide | 
| IP protection | Not included | Yes, with duplication filter enabled | Yes, with duplication filter enabled | 

In addition to the opt-out for individual plans, it's important to note that Business and Enterprise users benefit from content exclusions, IP protection, and organization-level admin controls, which are absent from free and individual plans. For a detailed comparison beyond data management, see our GitHub Copilot guide: plans.

**A few useful clarifications:** Content exclusions do not yet apply to Edit mode, Agent mode in IDE chats, GitHub Copilot CLI, or the cloud agent. IP protection coverage requires the duplication detection filter to be enabled and the suggestion to be used without modification.

## Configuring Copilot privacy settings

Data privacy has become a crucial issue for any business. Let's look at the settings that determine whether your interactions are used to train models when you use GitHub Copilot.

### Opt-out for Free, Pro, and Pro+

For your individual account, you can decline the use of your data for training purposes by going to GitHub Settings and setting "Allow GitHub to use my data for AI model training" to Disabled, as shown below.

This opt-out stops future collection and does not reduce Copilot's features. However, GitHub cannot guarantee the removal of data already used for previous training; your previously collected data may therefore persist in existing training sets.

### Organization and enterprise policies

Business and Enterprise users are already excluded from model training, but administrators should still review data sharing policies to control the Copilot features enabled in the organization:

- **Organization settings** > **Copilot** > **Policies** lets you manage features, license assignment, and model selection for all members.
- Org-level policies take precedence over individual preferences: any setting defined here applies to everyone.
- Enterprise owners can define policies inherited by multiple organizations and audit the state from a single dashboard.

## Using Copilot content exclusions

You can prevent Copilot from accessing certain content. From repository settings, define the items Copilot should ignore.

### How content exclusions work

For excluded files:

- Inline suggestions will not be offered.
- Their content will not be used to generate suggestions in other files.
- Their content will not be used by GitHub Copilot Chat for its responses.
- Copilot code review will not apply to these files.

Exclusions can be configured by repository administrators, organization owners, and Enterprise owners.

### Configuring exclusions at the repository and organization level

At the repo level, open Settings > Copilot > Content Exclusion, and specify paths using glob patterns. Common examples: `**"**/secrets/**"**` to exclude any path containing a secrets directory, and `**"*.env"**` to exclude all environment files. 

The REST API offers a programmatic option if you manage exclusions at scale and want to version your configuration.

At the org level, the path is: Org Settings > Copilot > Content Exclusion. Rules defined here apply to all repositories in the organization.

Org-level and repo-level rules are additive: they apply simultaneously. Enterprise-level rules take precedence over org-level and repo-level rules.

To verify that exclusions are working: open an excluded file and ask Copilot Chat to "explain this file." If Chat provides a relevant answer about its content, the exclusion is not being applied. This is the signal to reload the extension and check the rule syntax.

### Limitations to be aware of

When using GitHub Copilot content exclusions, keep in mind:

- Copilot CLI, Agent mode, and Cloud Agents do not respect exclusions (mid-2026).
- Semantic leakage: type information and hover definitions from excluded files may still indirectly influence suggestions.
- Symlinks and remote file systems are not covered.

## Code referencing and duplication filter

GitHub Copilot helps you understand the provenance of suggested code by referencing and linking the source. When you accept such a suggestion, Copilot records the source URL and its license.

This lets you decide whether to use the snippet and what type of attribution to provide.

### How the duplication filter works

When a suggestion is generated, Copilot compares it against known public code. If it matches a public repository beyond a similarity threshold, it is blocked or flagged with attribution.

If accepted, Copilot records:

- The date and time of acceptance
- The file where the suggestion was added
- A snippet of the added code
- The license and source URL

