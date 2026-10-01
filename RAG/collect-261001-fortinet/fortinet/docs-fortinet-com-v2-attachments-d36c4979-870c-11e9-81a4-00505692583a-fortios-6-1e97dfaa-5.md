---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-5
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [519, 655]
sha256: af4bc968346df3ab83337bfa3236f7fe069bd4e7c39286b8435b0917813ebf80
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

Security best practices 19
The default value of admin-lockout-threshold is 3 and the range of values is between 1 and 10. The admin-
lockout-duration is set to 60 seconds by default and the range of values is between 1 and 2147483647 seconds.
Keep in mind that the higher the lockout threshold, the higher the risk that someone may be able to break into the
FortiGate.
Example
To set the admin-lockout-threshold to one attempt and the admin-lockout-duration to a five minute
duration before the administrator can try to log in again, enter the commands:
config system global
set admin-lockout-threshold 1
set admin-lockout-duration 300
end
If the time span between the first failed login attempt and the admin-lockout-threshold
failed login attempt is less than admin-lockout-duration, the lockout will be triggered.
Rename the admin administrator account
You can improve security by renaming the admin account. To do this, create a new administrator account with the
super_admin admin profile and log in as that administrator. Then go to System > Administrators and edit the admin
administrator and change the User Name. Renaming the admin account makes it more difficult for an attacker to log into
FortiOS.
Add administrator disclaimers
FortiOS can display a disclaimer before or after logging into the GUI or CLI (or both). In either case the administrator
must read and accept the disclaimer before they can proceed.
Use the following command to display a disclaimer before logging in:
config system global
set pre-login-banner enable
end
Use the following command to display a disclaimer after logging in:
config system global
set post-login-banner enable
end
You can customize the replacement messages for these disclaimers by going to System > Replacement Messages.
Select Extended View to view and edit the Administrator replacement messages.
From the CLI:
config system replacemsg admin pre_admin-disclaimer-text
config system replacemsg admin post_admin-disclaimer-text
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 20
Global commands for stronger and more secure encryption
This section describes some best practices for employing stronger and more secure encryption.
Turn on global strong encryption
Enter the following command to configure FortiOS to use only strong encryption and allow only strong ciphers (AES,
3DES) and digest (SHA1) for HTTPS, SSH, TLS, and SSL functions.
config system global
set strong-crypto enable
end
Disable MD5 and CBC for SSH
In some cases, you may not be able to enable strong encryption. For example, your FortiGate may be communicating
with a system that does not support strong encryption. With strong-crypto disabled you can use the following options
to prevent SSH sessions with the FortiGate from using less secure MD5 and CBC algorithms:
config system global
set ssh-hmac-md5 disable
set ssh-cbc-cipher disable
end
Disable static keys for TLS
You can use the following command to prevent all TLS sessions that are terminated by FortiGate from using static keys
(AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256):
config system global
set ssl-static-key-ciphers disable
end
Require larger values for Diffie-Hellman exchanges
Larger Diffie-Hellman values result in stronger encryption. Use the following command to force Diffie-Hellman
exchanges to use 8192 bit values (the highest configurable DH value).
config system global
set dh-params 8192
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 21
Disable auto USB installation
If USB installation is enabled, an attacker with physical access to a FortiGate could load a new configuration or firmware
on the FortiGate using the USB port. You can disable USB installation by entering the following from the CLI:
config system auto-install
set auto-install-config disable
set auto-install-image disable
end
Set system time by synchronizing with an NTP server
For accurate time, use an NTP server to set system time. Synchronized time facilitates auditing and consistency
between expiry dates used in expiration of certificates and security protocols.
From the GUI go to System > Settings > System Time and select Synchronize with NTP Server. By default, this causes
FortiOS to synchronize with Fortinet's FortiGuard secure NTP server.
From the CLI you can use one or more different NTP servers:
config system ntp
set type custom
set ntpsync enable
config ntpserver
edit 1
set server <ntp-server-ip>
next
edit 2
set server <other-ntp-server-ip>
end
Disable the maintainer admin account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI.The maintainer account allows you to log into a FortiGate if you have lost all
administrator passwords.
Once you have logged in with the maintainer account you can:
l Change the password of the admin administrator account (if it exists).
l Reset the FortiGate to the factory default configuration using the execute factoryreset command. This is the
only way to get access to the FortiGate if you have deleted the admin administrator account.
See the Fortinet knowledge base or Resetting a lost Admin password for details about using the maintainer account to
regain access to your FortiGate if you have lost all administrator account passwords.
The methodology for using the maintainer account is publicly available. As long as someone with physical access to the
device has the serial number of the device, which is labeled on the device, they can change the admin administrator
account password and access the FortiGate. This may be an unacceptable risk in some circumstances, especially
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 22
where the hardware is not physically secured. As an added security measure, the maintainer account can be disabled
using the following setting:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be able to
log into your FortiGate.The only way to access your FortiGate will be to start over with a new
firmware installation and default configuration file. All of your settings will be lost.
Enable password policies
Go to System > Settings > Password Policy, to create a password policy that all administrators must follow. Using the
available options you can define the required length of the password, what it must contain (numbers, upper and lower
case, and so on) and an expiry time.
Use the password policy feature to make sure all administrators use secure passwords that meet your organization's
requirements.
Configure auditing and logging
For optimum security go to Log & Report > Log Settings enable Event Logging. For best results send log messages to
FortiAnalyzer or FortiCloud.
From FortiAnalyzer or FortiCloud, you can view reports or system event log messages to look for system events that
may indicate potential problems. You can also view system events by going to FortiView > System Events.
Establish an auditing schedule to routinely inspect logs for signs of intrusion and probing.
Encrypt logs sent to FortiAnalyzer/FortiManager
To keep information in log messages sent to FortiAnalyzer private, go to Log & Report > Log Settings and when you
configure Remote Logging to FortiAnalyzer/FortiManager select Encrypt log transmission.
From the CLI.
config log {fortianalyzer | fortianalyzer2 | fortianalyzer3} setting
set enc-algorithm high
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

