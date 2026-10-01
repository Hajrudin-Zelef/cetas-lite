---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5-97faca3f-1
title: "docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2018-02-15", "2018-03-21", "2019-06-06", "2022-30-03"]
keywords: ["license", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f.md
source_anchor: ""
source_lines: [1, 176]
sha256: 115b58d2de1e14af8637799a70eb9e07dc63f51543c703eff827ac43ec75aea5
---

# docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f

FortiOS™ Handbook - Hardening your FortiGate
VERSION 5.6

FORTINET DOCUMENT LIBRARY
https://docs.fortinet.com
FORTINET VIDEO GUIDE
https://video.fortinet.com
FORTINET KNOWLEDGE BASE
http://kb.fortinet.com
FORTINET BLOG
https://blog.fortinet.com
CUSTOMER SERVICE & SUPPORT
https://support.fortinet.com 
FORTINET NSE INSTITUTE (TRAINING)
https://training.fortinet.com/
FORTIGUARD CENTER
https://fortiguard.com
FORTICAST
http://forticast.fortinet.com
END USER LICENSE AGREEMENT AND PRIVACY POLICY
https://www.fortinet.com/doc/legal/EULA.pdf
https://www.fortinet.com/corporate/about-us/privacy.html
FEEDBACK
Email: techdoc@fortinet.com
3/30/2022
FortiOS™ Handbook - Hardening your FortiGate
01-560-467596-20180209

TABLE OF CONTENTS
Change log 5
Hardening 6
Building security into FortiOS 7
Boot PROM and BIOS security 7
FortiOS kernel and user processes 7
Administration access security 7
Admin administrator account 7
Secure password storage 7
Maintainer account 8
Administrative access security 8
Network security 9
Network interfaces 9
TCP sequence checking 10
Reverse path forwarding 10
FIPS and Common Criteria 10
PSIRT advisories 10
FortiOS ports and protocols 12
FortiOS open ports 12
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
Global commands for stronger and more secure encryption 20
Turn on global strong encryption 21
Disable MD5 and CBC for SSH 21
Disable static keys for TLS 21
Require larger values for Diffie-Hellman exchanges 21
Disable sending Malware statistics to FortiGuard 21
Disable auto USB installation 22
Set system time by synchronizing with an NTP server 22
Enable password policies 22
Configure auditing and logging 23
Encrypt logs sent to FortiAnalyzer/FortiManager 23
Disable interfaces that not used 23
Disable unused protocols on interfaces 23
Use local-in policies to close open ports or restrict access 24
Close ICMP ports 24
Close the BGP port 25

Change log
Change log
Date Change description
June 6, 2019 Minor updates.
March 21, 2018 Updated with new information throughout.
February 15, 2018 Updated for FortiOS 5.6.3.
Hardening
Fortinet Technologies Inc.
5

Hardening
This guide describes some of the techniques used to harden (improve the security of) FortiGate devices and
FortiOS.
This document contains the following sections:
l Building security into FortiOS
l FortiOS ports and protocols
l Security best practices
6 Hardening
Fortinet Technologies Inc.

Building security into FortiOS Boot PROM and BIOS security
Building security into FortiOS
The FortiOS operating system, FortiGate hardware devices, and FortiOS virtual machines (VMs) are built with
security in mind, so many security features are built into the hardware and software. Fortinet maintains an
ISO:9001 certified software and hardware development processes to ensure that FortiOS and FortiGate products
are developed in a secure manner
Boot PROM and BIOS security
The boot PROM and BIOS in FortiGate hardware devices use Fortinet's own FortiBootLoader that is designed and
controlled by Fortinet. FortiBootLoader is a secure, proprietary BIOS for all FortiGate appliances. FortiGate
physical devices always boot from FortiBootLoader.
FortiOS kernel and user processes
FortiOS is a multi-process operating system with kernel and user processes. The FortiOS kernel runs in a
privileged hardware mode while higher-level applications run in user mode. FortiOS is a closed system that does
not allow the loading or execution of third-party code in the FortiOS user space. All non-essential services,
packages, and applications are removed.
Administration access security
This section describes FortiOS and FortiGate administration access security features.
Admin administrator account
All FortiGate firewalls ship with a default administrator account called admin. By default, this account does not
have a password. FortiOS allows administrators to add a password for this account or to remove the account and
create new custom super_admin administrator accounts.
For more information, see Rename the admin administrator account on page 20.
Secure password storage
Passwords are encrypted when stored on the FortiGate, and encoded when displayed in the CLI and
configuration file.
To enhance your password security, you can specify your own private key for the encryption process. This
ensures that your key is unique. The key is also required to restore the system from a configuration file. In HA
clusters, the same key should be used on all of the units.
Hardening
Fortinet Technologies Inc.
7

Administration access security Building security into FortiOS
To enable and enter your own private encryption key:
config system global
set private-data-encryption enable
end
Please type your private data encryption key (32 hexadecimal numbers):
0123456789abcdef0123456789abcdef
Please re-enter your private data encryption key (32 hexadecimal numbers) again:
0123456789abcdef0123456789abcdef
Your private data encryption key is accepted.
This is an example. Using 0123456789abcdef0123456789abcdef as
your private key is not recommended.
Maintainer account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI. When enabled, the maintainer account can be used to log in from the
console after a hard reboot. The password for the maintainer account is bcpb followed by the FortiGate serial
number. An administrator has 60-seconds to complete this login. See Resetting a lost Admin password for details.
The only action the maintainer account has permissions to perform is to reset the passwords of super_admin
accounts. Logging in with the maintainer account requires rebooting the FortiGate. FortiOS generates event log
messages when you login with the maintainer account and for each password reset.
The maintainer account is enabled by default; however, there is an option to disable this feature. The maintainer
account can be disabled using the following command:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be
able to log into your FortiGate.
Administrative access security
Secure administrative access features:
l SSH, Telnet, and SNMP are disabled by default. If required, these admin services must be explicitly enabled on
each interface from the GUI or CLI.
l SSHv1 is disabled by default. SSHv2 is the default version.
l SSLv3 and TLS1.0 are disabled by default. TLSv1.1 and TLSv1.2 are the SSL versions enabled by default for
HTTPS admin access.
l HTTP is disabled by default, except on dedicated MGMT, DMZ, and predefined LAN interfaces. HTTP redirect to
HTTPS is enabled by default.
l The strong-crypto global setting is enabled by default and configures FortiOS to use strong ciphers (AES,
3DES) and digest (SHA1) for HTTPS/SSH/TLS/SSL functions.
8 Hardening
Fortinet Technologies Inc.

