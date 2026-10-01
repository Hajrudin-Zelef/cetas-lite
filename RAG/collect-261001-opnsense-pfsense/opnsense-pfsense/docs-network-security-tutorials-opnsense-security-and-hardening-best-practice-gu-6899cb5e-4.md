---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e-4
title: "docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "cost", "license", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e.md
source_anchor: ""
source_lines: [501, 596]
sha256: 274085df8a855822ff0817282cd54d75eee4d13acd23c847c671b3748362d7a4
---

# docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e

All administrative access should be restricted to a single interface or VLAN interface in the management network. All other interfaces should have administrative access disabled. To do this in OPNsense, you would utilize firewall rules, keeping in mind that the default is to deny all (implicit deny) incoming connections with the exception of the LAN interface.

## Enable IPS/IDS

Intrusion detection and prevention systems (IDPS) should be activated in OPNsense for hardening since they give an extra layer of protection to the firewall by detecting and blocking hostile or unauthorized network activity. IDPS systems identify threats and preventing them from reaching their intended targets using approaches such as signature-based detection, anomaly-based detection, and reputation-based detection. By activating IDPS in OPNsense, enterprises proactively defend their network against threats and reduce the likelihood of data breaches and other security issues. You can find the details about how to enable IPS on the OPNsense manual page: `https://docs.opnsense.org/manual/ips.html`

## Create Notifications for Changes

Change management notifications are essential for firewall hardening because they give a transparent record of any configuration changes. This helps guarantee that no illegal modifications are done and makes it easier to recognize and resolve any problems that may arise as a consequence of changes. Moreover, change management notifications can assist firms in meeting legal obligations and demonstrating due care in the case of a security problem. By giving a clear record of modifications, change management alerts make it simpler to debug and restore to a prior configuration if necessary, mitigating the impact of any potential issues.

For monitoring purposes, OPNsense uses the Monit framework. Given the breadth of Monit's monitoring capabilities, it comes as no surprise that the tool's setup choices are similarly flexible.

Monit is a program that checks your filesystems, drives, processes, and system, among other things. It operates on the OPNsense firewall host and delivers messages or performs actions in response to a variety of events.

First, you must choose what you want to monitor and what failure looks like. It is advantageous to understand how Monit alerts are configured.

### What is M/Monit?

M/Monit is an IT infrastructure monitoring and management solution that operates automatically and proactively. M/Monit is able to keep tabs on and manage a wide variety of computer systems, perform routine maintenance and repairs automatically, and take appropriate causal action in the event of an issue.

All of your hosts and services can be managed and monitored with M/Monit since it employs Monit as its agent. If a service isn't running, M/Monit can initiate its launch; if it has stopped responding, it can be restarted; and if it's using too many system resources, M/Monit can put a hold on it.

Keep an eye out for changes in your hosts' filesystems, directories, files, processes, and other system properties. If a value deviates from a predetermined threshold, an alert may be generated and a predetermined action is taken.

All of the data from the systems under surveillance is compiled and kept in a central repository. The gathered data may be explored using drill-down and filtering features. Charts and tables are automatically updated with the latest status and events from all monitored systems.

### What are the Benefits of M/Monit?

Monit provides the following advantages:

- 
**Easy-to-use:** In contrast to other solutions, M/Monit does not need extensive configuration or the installation of any additional third-party software.
- 
**Improved Uptime:** The uptime of your computer systems will improve since M/Monit can deal with problem scenarios automatically, in many cases without any human interaction.
- 
**Intuitive User Interface:** Monit's interface is intuitive and well-designed, making it a good choice whether you need to monitor two hosts or a thousand.
- 
**Open-Source:** The whole source code and build system are available for review. The M/Monit framework also includes freely available open-source software.
- 
**Cost-effective:** In exchange for a one-time fee, you get perpetual access to an M/Monit license. When compared to comparable commercial systems, the cost is negligible, and the time investment is little in comparison to that of an equivalent open-source solution.
- 
**Scalable Application:** The M/Monit application server is a state-of-the-art option that is both small and extensible. High performance is achieved with the use of thread-pools and an event-driven, non-blocking i/o architecture. M/Monit is compatible with all POSIX systems and requires just around 10 MB of memory. A connection pool that can communicate with MySQL, PostgreSQL, and SQLite manages access to the database.

### How to Enable and Configure Monit?

To enable and configure the Monit for OPNsense monitoring, you may follow the steps below:

1. 
Navigate to **Services** >**Monit** >**Settings** on OPNsense Web UI.
2. 
Enable Monit by clicking on the **Enable Monit** checkbox.
3. 
Set the port of the mail server in the **Mail Server Port** option. Typically you may set`465` for SSL or`25` for TLS and nonsecure connections.
4. 
Enter your email address into the **Mail Server Username** field for authentication.
5. 
Enter your email password into the **Mail Server Password** field for authentication.
6. 
You may enable encryption for mail server communication by clicking on the **Mail Server SSL Connection** option.

**Figure 27.** *Enabling Monit Service*

1. 
Click **Apply** to activate the settings.
2. 
To start the **Monit** service, click on the**Start** button at the upper right corner of the page.

**Figure 28.** *Starting Monit Service*

1. To view the status of the `Monit` service, navigate to the**Services** >**Monit** >**Status** . You should see the*status* field is`OK` .

You may view all general settings by clicking on the `advanced mode` option on the **General Settings** page.

| Setting | Description | 
|---|---|
| Enable Monit | Turns Monit on or off. | 
| Polling interval | How often Monit checks the health of the things it's watching. | 
| Start delay | Time in seconds before Monit begins verifying its components after being started. | 
| Mail Server | A list of mail servers to send notifications to. | 
| Mail Server Port | The mail server port to use. 25 and 465 are common examples. | 
| Username | The username used to log into your SMTP server, if needed. Often, but not always, the same as your e-mail address. | 
| Password | The password used to log into your SMTP server, if needed. | 
| Secure Connection | Use TLS when connecting to the mail server. | 
| SSL Version | The TLS version to use. AUTO will try to negotiate a working version. | 
| Verify SSL Certificates | Checks the TLS certificate for validity. If you use a self-signed certificate, turn this option off. | 
| Log File | The log file of the Monit process. This can be the keyword syslog or a path to a file. | 
| State File | The state file of the Monit process. | 
| Eventqueue Path | The path to the eventqueue directory. | 
| Eventqueue Slots | The number of eventqueue slots. | 
| Enable HTTPD | Turns on the Monit web interface. | 
| Monit HTTPD Port | The listen port of the Monit web interface service. | 
| Monit HTTPD Access List | The username:password or host/network etc. for accessing the Monit web interface service. | 
| M/Monit URL | The M/Monit URL, e.g. `https://user:[email protected]:8443/collector` | 
| M/Monit Timeout | Set the timeout for requests to M/Monit to this value of seconds. | 
| M/Monit Register Credentials | Registration in M/Monit is performed automatically by the transfer of Monit credentials. | 

