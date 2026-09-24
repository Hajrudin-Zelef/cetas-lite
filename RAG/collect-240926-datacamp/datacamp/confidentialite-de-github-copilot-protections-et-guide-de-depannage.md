---
id: collect-240926-datacamp/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-05"]
keywords: ["copilot", "agent", "agents", "attribution", "cloud agent", "copyright", "guardrails", "license", "training"]
source: docs/RAG/clean_en/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [1, 253]
sha256: 5ed2a96460a238b7ee3d3b5c3ae999e5bb939688f9e42066ec69e2a4ff5e3d41
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

You can see code references directly in the IDE when a match is flagged. In VS Code, they appear in the Copilot output panel alongside the suggestion.

### What the filter does not detect

The duplication filter does not match:

- Short snippets and patterns too generic to be flagged.
- Code restructured or partially modified from a source.

Its purpose is to detect textual matches (identical or near-identical), not conceptual similarity.

### IP protection and contractual guarantees

GitHub offers intellectual property (IP) protection to Copilot Business and Enterprise customers. If a suggestion leads to a claim, GitHub handles the legal defense.

Two conditions apply:

1. The duplication filter must be enabled
2. You must have an eligible plan

Free and Pro users are not covered. This is a commercial guarantee: it does not prevent matches from appearing, but covers the legal risk if they do.

## Managing Copilot policies on GitHub.com

The Copilot policy page on GitHub.com is where org and Enterprise admins control what Copilot can do for the team.

### Enabling and disabling Copilot features

Admins can independently enable/disable code completions, Chat, code review, GitHub CLI integration, and agent mode. This granularity is useful for a phased rollout or to restrict certain capabilities to specific teams.

License assignment is also done here: you choose who has access, by user or by group.

### Configuring allowed AI models

Copilot supports several underlying models, and administrators can restrict which ones are allowed for the organization. You can lock it to a specific model or allow all available options and let developers choose.

For public and regulated environments, GitHub offers options compliant with the FedRAMP (Federal Risk and Authorization Management Program).

Check the Copilot policy settings in your organization's Copilot tab to see what is available with your plan.

### Audit and compliance

Copilot usage indicators, including completion rates, number of active users, and trends by feature and model, are available under the Insights tab at the Enterprise and organization level (Insights > Copilot usage).

The "Copilot usage metrics" policy must be enabled to access the dashboard. Per-member details are available via NDJSON export.

Seat and license data are separate and accessible under Org Settings > Copilot > Access.

For a detailed overview of the new Usage Metrics API and other advanced features, feel free to read our GitHub Copilot Enterprise guide.

Note that the audit log does not include client session data such as prompts; a custom solution is required for that. To retain history beyond 180 days or set up anomaly alerts, GitHub recommends streaming the audit log to a SIEM platform via the built-in streaming feature.

## Troubleshooting GitHub Copilot

When Copilot stops working, the cause is almost always one of the following. Check these points before opening a ticket.

### Missing or interrupted suggestions

Start with the Copilot status icon in the IDE status bar. A diagonal line across the icon indicates that a content exclusion is active for the current file.

If the icon appears normal but no suggestions appear, check, in order:

- Update your IDE and the Copilot extension.
- Verify that your subscription is active and that your account has an assigned license.
- Check the exclusion rules for the file and the repo, and test your network connection.

Proxy and VPN configurations frequently block silently: the IDE must reach Copilot servers on GitHub infrastructure, and some corporate proxies block them without an explicit message.

### Unexpected content exclusion

After adding or modifying exclusions, a delay of up to 30 minutes may apply in IDEs where settings are already loaded.

To apply changes immediately:

- In VS Code, open the **Command Palette** and run **Developer: Reload Window**.
- In JetBrains and Visual Studio IDEs, close and reopen the application.
- In Vim/Neovim, no action needed: exclusions are fetched automatically each time a file is opened.

After reloading, test explicitly: open the excluded file and ask Chat to explain it. If Chat describes the file's contents, the exclusion is not being applied; recheck the rule's syntax.

Note that three Copilot features do not handle exclusions: Copilot CLI, the Copilot coding agent (cloud agent), and Agent mode in Copilot Chat in the IDE. If you observe unexpected access with these modes, it is not a misconfiguration.

### Authentication and token issues

If Copilot is unavailable in VS Code despite an open session, sign out via the Accounts icon at the bottom left, reload the window (F1 > Developer: Reload Window in VS Code), then sign back in.

In Visual Studio, verify that the connected GitHub account matches the one with a Copilot license, refresh credentials if needed, or remove/re-add the account and restart Visual Studio.

### Rate limiting

Copilot's usage-based billing model means different capabilities per plan, and premium models consume this capacity faster than base models.

If suggestions stop mid-session or Copilot Chat returns errors, switching to automatic model selection (or to a model with a lower multiplier) can resolve the issue until the usage window resets.

Enterprise admins can proactively monitor usage via the Copilot analytics dashboard to anticipate limits.

For service issues, check githubstatus.com before debugging locally.

## In conclusion

GitHub Copilot gives teams real control over their data management: contractual protections tied to the plan, granular content exclusions, and duplication filter.

Understanding these settings (and knowing how to configure them properly) lets you adopt Copilot with peace of mind, whether you are an individual developer or deploying at enterprise scale. If something does not work as expected, the troubleshooting steps above should quickly get you back on track.

If you would like to practice with GitHub Copilot, learn how to customize it, and leverage all its intelligent features, we highly recommend our Software Development with GitHub Copilot course.

## FAQ on GitHub Copilot privacy and troubleshooting

### Does GitHub Copilot send code from my private repositories to GitHub servers?

**Copilot sends immediate code context from your editor to GitHub servers to generate a suggestion. It does not draw from the code in your private repositories stored at rest on GitHub.**

### How do I prevent GitHub from using my Copilot data to train its models?

**Go to GitHub Settings, then Copilot, and disable "Allow GitHub to use my data for AI model training". This opt-out applies immediately to future collections. Copilot Business and Enterprise users are automatically excluded from training and have nothing to change.**

### What is a content exclusion and how do I configure it?

**A content exclusion is a rule that prevents Copilot from reading or generating suggestions from specific files or paths. You configure it at the repository level via Settings > Copilot > Content Exclusion, with glob patterns like `'*.env'` or `'**/secrets/**'`. Organization owners can define exclusions that apply to all repositories.**

### Does GitHub Copilot's duplication filter protect against copyright issues?

**It filters textual matches (identical or near-identical) with known public code and can flag the source license. It does not detect conceptually similar code or partial rewrites. For full IP protection, Copilot Business and Enterprise customers also benefit from legal protection if a suggestion triggers a claim, provided the duplication filter is enabled.**

### Why has GitHub Copilot stopped showing suggestions in a specific file?

**A slash through the Copilot status icon means an exclusion rule is active for that file. Check the exclusion settings at the repo and org levels to see if the file matches a rule. If so and you want to allow suggestions, modify the exclusion pattern.**

### Does agent mode respect content exclusions?

**No. As of May 2026, Copilot's agent mode and Cloud Agents do not respect content exclusion rules.**

### How do you fix GitHub Copilot authentication errors?

**Sign out of GitHub in your IDE and sign back in. Verify that the account you're using has an active Copilot license. For Enterprise accounts, reauthenticate if you recently changed your password or if SSO has changed. In Visual Studio, check for duplicate or conflicting versions of the Copilot extension.**

### What is GitHub Copilot's IP protection and who is eligible for it?

**IP protection means that GitHub covers defense costs if a Copilot suggestion triggers an intellectual property claim. It is available to Copilot Business and Copilot Enterprise customers, under two conditions: the duplication filter must be enabled and the suggestion must be used as-is. Users of the Free and Pro plans are not covered.**
