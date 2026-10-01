---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6-293b419e-1
title: "docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2020-03-31", "2020-10-06", "2020-12-10", "2022-06-13", "2023-08-25", "2024-05-05", "2024-05-06"]
keywords: ["license", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e.md
source_anchor: ""
source_lines: [1, 136]
sha256: d81cc46b277f5c6a09d812ed94d29fd4b45a75826305118bad23765807a44d5c
---

# docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e

FortiOS - Hardening your FortiGate
Version 6.4.0

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
May 05, 2024
FortiOS 6.4.0 Hardening your FortiGate
01-640-619384-20240506

TABLE OF CONTENTS
Change Log 5
Hardening your FortiGate 6
Building security into FortiOS 7
Boot PROM and BIOS security 7
FortiOS kernel and user processes 7
Administration access security 7
Admin administrator account 7
Secure password storage 8
Configuration backup 8
Maintainer account 8
Administrative access security 9
Non-factory SSL certificates 10
Network security 10
Network interfaces 10
TCP sequence checking 11
Reverse path forwarding 11
FIPS and Common Criteria 11
PSIRT advisories 11
FortiOS ports and protocols 13
FortiOS open ports 13
Closing open ports 15
Security best practices 16
Install the FortiGate unit in a physically secure location 16
Register your product with Fortinet Support 16
Keep your FortiOS firmware up to date 16
System administrator best practices 17
Disable administrative access to the external (Internet-facing) interface 17
Allow only HTTPS access to the GUI and SSH access to the CLI 17
Require TLS 1.2 for HTTPS administrator access 17
Re-direct HTTP GUI logins to HTTPS 17
Change the HTTPS and SSH admin access ports to non-standard ports 18
Maintain short login timeouts 18
Restrict logins from trusted hosts 18
Set up two-factor authentication for administrators 19
Create multiple administrator accounts 19
Modify administrator account lockout duration and threshold values 19
Rename the admin administrator account 20
Add administrator disclaimers 20
Global commands for stronger and more secure encryption 21
Turn on global strong encryption 21
Disable MD5 and CBC for SSH 21
Disable static keys for TLS 21
Require larger values for Diffie-Hellman exchanges 21
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

4
Disable auto USB installation 22
Set system time by synchronizing with an NTP server 22
Disable the maintainer admin account 22
Enable password policies 23
Configure auditing and logging 23
Encrypt logs sent to FortiAnalyzer/FortiManager 23
Disable unused interfaces 24
Disable unused protocols on interfaces 24
Use local-in policies to close open ports or restrict access 25
Close ICMP ports 25
Close the BGP port 25
Optional settings 26
Send malware statistics to FortiGuard 26
Send Security Rating statistics to FortiGuard 26
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Change Log
Date Change Description
2020-03-31 Initial release.
2020-10-06 Updated Building security into FortiOS on page 7.
2020-12-10 Updated Building security into FortiOS on page 7.
2022-06-13 Updated Building security into FortiOS on page 7.
2023-08-25 Added Configuration backup on page 8.
2024-05-06 Updated Building security into FortiOS on page 7.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Hardening your FortiGate
This guide describes some of the techniques used to harden (improve the security of) FortiGate devices and FortiOS.
This guide contains the following sections:
l Building security into FortiOS on page 7
l FortiOS ports and protocols on page 13
l Security best practices on page 16
l Optional settings on page 26
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
For more information, see Rename the admin administrator account on page 20.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

