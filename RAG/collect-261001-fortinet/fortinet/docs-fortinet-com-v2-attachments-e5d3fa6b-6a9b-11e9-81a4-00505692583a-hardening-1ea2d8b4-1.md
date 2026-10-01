---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-1
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2018-06-18", "2018-07-26", "2018-10-04", "2019-04", "2019-04-29"]
keywords: ["license", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [1, 134]
sha256: 1e611ceb185090d5640e8425cb29531a2c7c563b2e8544b536664df069262279
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

FortiOS - Handbook - Hardening your FortiGate
Version 6.0.4

FORTINET DOCUMENT LIBRARY
https://docs.fortinet.com
FORTINET VIDEO GUIDE
https://video.fortinet.com
FORTINET BLOG
https://blog.fortinet.com
CUSTOMER SERVICE & SUPPORT
https://support.fortinet.com
FORTINET COOKBOOK
https://cookbook.fortinet.com
FORTINET TRAINING & CERTIFICATION PROGRAM
https://www.fortinet.com/support-and-training/training.html
NSE INSTITUTE
https://training.fortinet.com
FORTIGUARD CENTER
https://fortiguard.com/
END USER LICENSE AGREEMENT
https://www.fortinet.com/doc/legal/EULA.pdf
FEEDBACK
Email: techdoc@fortinet.com
29 April 2019
FortiOS 6.0.4 Handbook - Hardening your FortiGate
01-604-467596-20190429

TABLE OF CONTENTS
Change log 5
Hardening your FortiGate 6
Building security into FortiOS 7
Boot PROM and BIOS security 7
FortiOS kernel and user processes 7
Administrationaccesssecurity 7
Admin administratoraccount 7
Secure password storage 7
Maintainer account 8
Administrativeaccesssecurity 8
Network security 9
Network interfaces 9
TCP sequencechecking 10
Reverse path forwarding 10
FIPS and Common Criteria 10
PSIRT advisories 11
FortiOS ports and protocols 12
FortiOS open ports 12
Closing open ports 14
Security best practices 16
Install the FortiGate unit in a physicallysecure location 16
Register your product with Fortinet Support 16
Keep your FortiOS firmware up to date 16
System administratorbest practices 17
Disable administrativeaccessto the external (Internet-facing) interface 17
Allow only HTTPS accessto the GUI and SSH accessto the CLI 17
Require TLS 1.2 for HTTPS administratoraccess 17
Re-direct HTTP GUI logins to HTTPS 17
Change the HTTPS and SSH admin accessports to non-standard ports 18
Maintain short login timeouts 18
Restrict logins from trusted hosts 18
Set up two-factor authenticationfor administrators 19
Create multiple administratoraccounts 19
Modify administratoraccount lockout duration and threshold values 19
Rename the admin administratoraccount 20
Add administratordisclaimers 20
Global commandsfor stronger and more secure encryption 21
Turn on global strong encryption 21
Disable MD5 and CBC for SSH 21
Disable static keys for TLS 21
Require larger values for Diffie-Hellman exchanges 21
Disable sending malware statisticsto FortiGuard 21
Disable sending Security Rating statisticsto FortiGuard 22
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

4
Disable auto USB installation 22
Set system time by synchronizingwith an NTP server 22
Disable the maintainer admin account 23
Enable password policies 23
Configure auditing and logging 24
Encrypt logs sent to FortiAnalyzer/FortiManager 24
Disable unused interfaces 24
Disable unused protocols on interfaces 24
Use local-in policiesto close open ports or restrict access 25
Close ICMP ports 25
Close the BGP port 26
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Change log 5
Change log
Date Change description
April 29, 2019 FortiOS 6.0.4 document release.
October 4, 2018 FortiOS 6.0.3 document release. See Hardening your FortiGate on page 6.
July 26, 2018 FortiOS 6.0.2 document release. Correction to Restrict logins from trusted hosts on
page 18.
June 18, 2018 FortiOS 6.0.1 document release.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Hardening your FortiGate 6
Hardening your FortiGate
This guide describes some of the techniques used to harden (improve the security of) FortiGate devices and FortiOS.
This document contains the following sections:
l Building security into FortiOS
l FortiOS ports and protocols
l Security best practices
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 7
Building security into FortiOS
The FortiOS operating system, FortiGate hardware devices, and FortiOS virtual machines (VMs) are built with security
in mind, so many security features are built into the hardware and software. Fortinet maintains an ISO:9001 certified
software and hardware development processesto ensure that FortiOS and FortiGate products are developed in a
secure manner
Boot PROM and BIOS security
The boot PROM and BIOS in FortiGate hardware devices use Fortinet's own FortiBootLoader that is designed and
controlled by Fortinet. FortiBootLoader is a secure, proprietary BIOS for all FortiGate appliances. FortiGate physical
devices always boot from FortiBootLoader.
FortiOS kernel and user processes
FortiOS is a multi-process operating system with kernel and user processes. The FortiOS kernel runs in a privileged
hardware mode while higher-level applications run in user mode. FortiOS is a closed system that does not allow the
loading or execution of third-party code in the FortiOS user space. All non-essential services, packages, and
applications are removed.
FortiGate appliances with SD drives are encrypted to prevent unauthorized access to data.
Administration access security
This section describes FortiOS and FortiGate administration access security features.
Admin administrator account
All FortiGate firewalls ship with a default administrator account called admin. By default, this account does not have a
password. FortiOS allows administrators to add a password for this account or to remove the account and create new
custom super_admin administrator accounts.
For more information, see Rename the admin administrator account on page 20.
Secure password storage
User and administrator passwords are stored securely on the system in an encrypted format. The encryption hash used
for admin account passwords is SHA256/SHA1. The value that is seen in the configuration file is the Base64 encoded
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

