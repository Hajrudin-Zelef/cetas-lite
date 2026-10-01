---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-1
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2019-06-04", "2020-10-08", "2020-12-10", "2022-04-28", "2022-06-13"]
keywords: ["license", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [1, 134]
sha256: 6f9e4254541634d4dca6da5bdc1179054cfe4bcf1df29a8eed76e661a4b8a32d
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

FortiOS - Hardening your FortiGate
Version 6.2.0

FORTINET DOCUMENT LIBRARY
https://docs.fortinet.com
FORTINET VIDEO GUIDE
https://video.fortinet.com
FORTINET BLOG
https://blog.fortinet.com
CUSTOMER SERVICE & SUPPORT
https://support.fortinet.com
FORTINET TRAINING & CERTIFICATION PROGRAM
https://www.fortinet.com/support-and-training/training.html
NSE INSTITUTE
https://training.fortinet.com
FORTIGUARD CENTER
https://fortiguard.com/
END USER LICENSE AGREEMENT
https://www.fortinet.com/doc/legal/EULA.pdf
FEEDBACK
Email: techdoc@fortinet.com
April 28, 2022
FortiOS 6.2.0 Hardening your FortiGate
01-620-554155-20220428

TABLE OF CONTENTS
Change Log 5
Hardening your FortiGate 6
Building security into FortiOS 7
Boot PROM and BIOS security 7
FortiOS kernel and user processes 7
Administration access security 7
Admin administrator account 7
Secure password storage 8
Maintainer account 8
Administrative access security 8
Non-factory SSL certificates 9
Network security 10
Network interfaces 10
TCP sequence checking 10
Reverse path forwarding 10
FIPS and Common Criteria 11
PSIRT advisories 11
FortiOS ports and protocols 12
FortiOS open ports 12
Closing open ports 14
Security best practices 15
Install the FortiGate unit in a physically secure location 15
Register your product with Fortinet Support 15
Keep your FortiOS firmware up to date 15
System administrator best practices 16
Disable administrative access to the external (Internet-facing) interface 16
Allow only HTTPS access to the GUI and SSH access to the CLI 16
Require TLS 1.2 for HTTPS administrator access 16
Re-direct HTTP GUI logins to HTTPS 16
Change the HTTPS and SSH admin access ports to non-standard ports 17
Maintain short login timeouts 17
Restrict logins from trusted hosts 17
Set up two-factor authentication for administrators 18
Create multiple administrator accounts 18
Modify administrator account lockout duration and threshold values 18
Rename the admin administrator account 19
Add administrator disclaimers 19
Global commands for stronger and more secure encryption 20
Turn on global strong encryption 20
Disable MD5 and CBC for SSH 20
Disable static keys for TLS 20
Require larger values for Diffie-Hellman exchanges 20
Disable auto USB installation 21
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

4
Set system time by synchronizing with an NTP server 21
Disable the maintainer admin account 21
Enable password policies 22
Configure auditing and logging 22
Encrypt logs sent to FortiAnalyzer/FortiManager 22
Disable unused interfaces 23
Disable unused protocols on interfaces 23
Use local-in policies to close open ports or restrict access 24
Close ICMP ports 24
Close the BGP port 24
Optional settings 25
Send malware statistics to FortiGuard 25
Send Security Rating statistics to FortiGuard 25
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Change Log
Date Change Description
2019-06-04 Initial release.
2020-10-08 Updated Building security into FortiOS on page 7.
2020-12-10 Updated Building security into FortiOS on page 7.
2022-04-28 Updated System administrator best practices on page 16.
2022-06-13 Updated Building security into FortiOS on page 7.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Hardening your FortiGate
This guide describes some of the techniques used to harden (improve the security of) FortiGate devices and FortiOS.
This guide contains the following sections:
l Building security into FortiOS on page 7
l FortiOS ports and protocols on page 12
l Security best practices on page 15
l Optional settings on page 25
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS
The FortiOS operating system, FortiGate hardware devices, and FortiGate virtual machines (VMs) are built with security
in mind, so many security features are built into the hardware and software. Fortinet maintains an ISO:9001 certified
software and hardware development processes to ensure that FortiOS and FortiGate products are developed in a
secure manner.
Boot PROM and BIOS security
The boot PROM and BIOS in FortiGate hardware devices use Fortinet's own FortiBootLoader that is designed and
controlled by Fortinet. FortiBootLoader is a secure, proprietary BIOS for all FortiGate appliances. FortiGate physical
devices always boot from FortiBootLoader.
FortiOS kernel and user processes
FortiOS is a multi-process operating system with kernel and user processes. The FortiOS kernel runs in a privileged
hardware mode while higher-level applications run in user mode. FortiOS is a closed system that does not allow the
loading or execution of third-party code in the FortiOS user space. All non-essential services, packages, and
applications are removed.
Administration access security
This section describes FortiOS and FortiGate administration access security features.
As the first step on a new deployment, review default settings such as administrator passwords, certificates for GUI and
SSL VPN access, SSH keys, open administrative ports on interfaces, and default firewall policies. As soon as the
FortiGate is connected to the internet it is exposed to external risks, such as unauthorized access, man-in-the-middle
attacks, spoofing, DoS attacks, and other malicious activities from malicious actors. Either use the start up wizard or
manually reconfigure the default settings to tighten your security from the beginning, thereby securing your network to its
full potential.
Admin administrator account
All FortiGate firewalls ship with a default administrator account called admin. By default, this account does not have a
password, except for FortiGate VMs on public clouds. FortiOS allows administrators to add a password for this account
or to remove the account and create new custom super_admin administrator accounts.
For more information, see Rename the admin administrator account on page 19.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

