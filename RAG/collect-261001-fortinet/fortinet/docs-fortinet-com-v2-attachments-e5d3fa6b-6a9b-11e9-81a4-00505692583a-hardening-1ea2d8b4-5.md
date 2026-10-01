---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-5
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [533, 676]
sha256: 02087dbe27a84ef0f1bcf9114ac7e8ea8d82c606656abcdb92f278609ed0f3c0
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

Security best practices 21
Global commands for stronger and more secure encryption
This section describes some best practices for employing stronger and more secure encryption.
Turn on global strong encryption
Enter the following command to configure FortiOS to use only strong encryption and allow only strong ciphers (AES,
3DES) and digest (SHA1) for HTTPS, SSH, TLS, and SSL functions.
config sys global
set strong-crypto enable
end
Disable MD5 and CBC for SSH
In some cases, you may not be able to enable strong encryption. For example, your FortiGate may be communicating
with a system that does not support strong encryption. Withstrong-crypto disabled you can use the following
options to prevent SSH sessions with the FortiGate from using less secure MD5 and CBC algorithms:
config sys global
set ssh-hmac-md5 disable
set ssh-cbc-cipher disable
end
Disable static keys for TLS
You can use the following command to prevent TLS sessions from using static keys (AES128-SHA, AES256-SHA,
AES128-SHA256, AES256-SHA256):
config sys global
set ssl-static-key-ciphers disable
end
Require larger values for Diffie- Hellman exchanges
Larger Diffie-Hellman values result in stronger encryption. Use the following command to force Diffie-Hellman
exchanges to use 8192 bit values (the highest configurable DH value).
config sys global
set dh-params 8192
end
Disable sending malware statistics to FortiGuard
By default FortiOS periodically sends encrypted malware statistics to FortiGuard. The malware statistics record
Antivirus, IPS, or Application Control events. This data is used to improved FortiGuard services. The malware statistics
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 22
that FortiOS sends do not include any personal or sensitive customer data. The information is not shared with any
external parties and is used in accordance with Fortinet's Privacy Policy.
To disable sending malware statistics to FortiGuard, enter the following command:
config system global
set fds-statistics disable
end
Disable sending Security Rating statistics to FortiGuard
Security Rating is a Fortinet Security Fabric feature that allows customers to audit their Security Fabric and find and fix
security problems. As part of the feature, FortiOS sends your security rating to FortiGuard every time a security rating
test runs.
You can opt out of submitting Security Rating scores to FortiGuard. If you opt out you won't be able to see how your
organization's scores compare with the scores of other organizations. Instead, an absolute score is shown. Use the
following command to disable FortiGuard Security Rating result submission:
config system global
set fortiguard-audit-result-submission disable
end
Disable auto USB installation
If USB installation is enabled, an attacker with physical access to a FortiGate could load a new configuration or firmware
on the FortiGate using the USB port. You can disable USB installation by entering the following from the CLI:
config system auto-install
set auto-install-config disable
set auto-install-image disable
end
Set system time by synchronizing with an NTP server
For accurate time, use an NTP serverto set system time. Synchronizedtime facilitates auditing and consistency
between expiry dates used in expiration of certificates and security protocols.
From the GUI go to System > Settings > System Time and select Synchronize with NTP Server. By default, this
causes FortiOS to synchronize with Fortinet's FortiGuard secure NTP server.
From the CLI you can use one or more different NTP servers:
config system ntp
set type custom
set ntpsync enable
config ntpserver
edit 1
set server <ntp-server-ip>
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 23
next
edit 2
set server <other-ntp-server-ip>
end
Disable the maintainer admin account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI.The maintainer account allows you to log into a FortiGate if you have lost
all administrator passwords.
Once you have logged in with the maintainer account you can:
l Change the password of the admin administrator account (if it exists).
l Reset the FortiGate to the factory default configuration using theexecute factoryreset command. This is
the only way to get access to the FortiGate if you have deleted the admin administrator account.
See the Fortinet knowledgebase or Resetting a lost Admin password for details about using the maintainer account to
regain access to your FortiGate if you have lost all administrator account passwords.
The methodology for using the maintainer account is publicly available. As long as someone with physical access to the
device has the serial number of the device, which is labeled on the device, they can change the admin administrator
account password and access the FortiGate. This may be an unacceptable risk in some circumstances, especially where
the hardware is not physically secured. As an added security measure, the maintainer account can be disabled using the
following setting:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be able to
log into your FortiGate.The only way to access your FortiGate will be to start over with a new
firmware installation and default configuration file. All of your settings will be lost.
Enable password policies
Go to System > Settings > Password Policy, to create a password policy that all administrators must follow. Using
the available options you can define the required length of the password, what it must contain (numbers, upper and
lower case, and so on) and an expiry time.
Use the password policy feature to make sure all administrators use secure passwords that meet your organization's
requirements.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 24
Configure auditing and logging
For optimum security go to Log & Report > Log Settings enable Event Logging. For best results send log
messages to FortiAnalyzer or FortiCloud.
From FortiAnalyzer or FortiCloud, you can view reports or system event log messages to look for system events that
may indicate potential problems. You can also view system events by going to FortiView > System Events.
Establish an auditing schedule to routinely inspect logs for signs of intrusion and probing.
Encrypt logs sent to FortiAnalyzer/FortiManager
To keep information in log messages sent to FortiAnalyzer private, go to Log & Report > Log Settings and when you
configure Remote Logging to FortiAnalyzer/FortiManager select Encrypt log transmission.
From the CLI.
config log {fortianalyzer | fortianalyzer2 | fortianalyzer3} setting
set enc-algorithm high
end
Disable unused interfaces
To disable an interface from the GUI, go to Network > Interfaces. Edit the interface to be disabled and set Interface
State to Disabled.
From the CLI, to disable the port21 interface:
config system interface
edit port21
set status down
end
Disable unused protocols on interfaces
You can use theconfig system interface command to disable unused protocols that attackers may attempt to
use to gather information about a FortiGate unit. Many of these protocols are disabled by default. Using theconfig
system interface command you can see the current configuration of each of these options for the selected
interface and then choose to disable them if required.
config system interface
edit <interface-name>
set dhcp-relay-service disable
set pptp-client disable
set arpforward disable
set broadcast-forward disable
set l2forward disable
set icmp-redirect disable
set vlanforward disable
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

