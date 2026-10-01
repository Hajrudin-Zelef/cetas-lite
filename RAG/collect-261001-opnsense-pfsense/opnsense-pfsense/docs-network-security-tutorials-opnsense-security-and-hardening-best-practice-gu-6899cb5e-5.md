---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e-5
title: "docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e.md
source_anchor: ""
source_lines: [597, 620]
sha256: 576c5ab99f2a8ddbe59f20e04f7ecc5f4d235fd59e1ca9583ff82f2ced926e2b
---

# docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e

More than one server may be entered into the "Mail Server" field. When sending an email, Monit will attempt to connect to each mail server in sequence, beginning with the first and moving on to the second if the first fails. No more attempts will be made to send the email by Monit if no server is found to be functional.

To aggregate information from several instances of Monit, M/Monit provides a paid service. Fill up the required data and add the necessary firewall rules in order to utilize it via OPNsense.

1. 
Type `From: [email protected]` in the "Mail format" box if your mail server needs the "From" field to be correctly configured.
2. 
Put the new settings into effect and **Save** the alert.

The Monit daemon is configured using the `os-monit` plugin. Install the `os-monit` plugin before using Monit. As a dependency, it installs the `monit` package.

Reload the GUI when the installation is complete and browse to Services->Monit->Settings.

The first step is to ensure that the plugin installer successfully imported your System->Notification settings. Then check out the other tabs. The installer has included some common entries to help you get started.

To establish a monitoring system, first construct Service Tests, then Services to check, and finally Alerts.

Begin with the Service Test Settings. A condition and an action are both part of a test. It is possible to assign it to one or more services. The Monit documentation includes examples of potential testing. Simply remove the IF and THEN statements to use it.

The following step is to set up service checks. Depending on the service type, we must specify a route, start/stop scripts, and assign already established tests. The same tests might be assigned to various service checks.

On the Alert Settings page, you can choose who receives notifications for particular events and who does not.

You can also format the email text. E.g. Subject: $SERVICE on $HOST failed on $DATE
