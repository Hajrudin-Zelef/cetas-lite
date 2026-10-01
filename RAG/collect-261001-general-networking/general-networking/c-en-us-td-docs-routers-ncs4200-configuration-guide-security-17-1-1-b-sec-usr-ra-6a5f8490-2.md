---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490-2
title: "c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490.md
source_anchor: ""
source_lines: [132, 189]
sha256: d93ba2828e8a2a48119fe9e009fb8b21ea728c5080b4f2e898b1fc54c67d76af
---

# c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490

aaa group server radius group1
server 10.1.1.1 auth-port 1645 acct-port 1646
server 10.2.2.2 auth-port 2000 acct-port 2001
deadtime 1
! The following commands define the group2 RADIUS server group and associate servers
! with it and configures a deadtime of two minutes.
aaa group server radius group2
server 10.2.2.2 auth-port 2000 acct-port 2001
server 10.3.3.3 auth-port 1645 acct-port 1646
deadtime 2
! The following set of commands configures the RADIUS attributes for each host entry
! associated with one of the defined server groups.
radius-server host 10.1.1.1 auth-port 1645 acct-port 1646
radius-server host 10.2.2.2 auth-port 2000 acct-port 2001
radius-server host 10.3.3.3 auth-port 1645 acct-port 1646
The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use
these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products
and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password.
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for AAA
Server Groups
Feature Name
Releases
Feature
Information
AAA Server
Group
Configuring
the device to use AAA server groups provides a way to group existing server
hosts. This allows you to select a subset of the configured server hosts and
use them for a particular service. A server group is used with a global
server-host list. The server group lists the IP addresses of the selected
server hosts.
The following
commands were introduced or modified:
aaa group server
radius,
aaa group server
tacacs+, and
server
(RADIUS).
AAA Server
Group Enhancements
AAA Server
Group Enhancements enables the full configuration of a server in a server
group.
AAA Server
Group Deadtimer
Configuring
deadtime within a server group allows you to direct AAA traffic to separate
groups of servers that have different operational characteristics.
The following
commands were introduced or modified:
deadtime.
