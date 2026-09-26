---
id: collect-240926-datacamp/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage-2
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "cloud agent", "copyright", "license", "training"]
source: docs/RAG/clean_en/datacamp/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [115, 242]
sha256: bed4df1c0a6a12b5aecd5922dcfb801c30391da9edadf735f7620d16baefbcf6
---

# confidentialite-de-github-copilot-protections-et-guide-de-depannage

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

