---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5-97faca3f-4
title: "docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f.md
source_anchor: ""
source_lines: [475, 593]
sha256: 507a9c09fd51703f89283e30e850389cf6206d7345791dae8fcbaeeb99ae82ca
---

# docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f

Security best practices System administrator best practices
set trustedhost2 172.25.177.0 255.255.255.0
end
Trusted host IP addresses can identify individual hosts or subnets. Just like firewall policies, FortiOS searches
through the list of trusted hosts in order and acts on the first match it finds. When you configure trusted hosts, start
by adding specific addresses at the top of the list. Follow with more general IP addresses. You don't have to add
addresses to all of the trusted hosts as long as all specific addresses are above all of the 0.0.0.0 0.0.0.0
addresses.
Set up two-factor authentication for administrators
FortiOS supports FortiToken and FortiToken Mobile 2-factor authentication. FortiToken Mobile is available for iOS
and Android devices from their respective application stores.
Every registered FortiGate unit includes two trial tokens for free. You can purchase additional tokens from your
reseller or from Fortinet.
To assign a token to an administrator, go to System > Administrators and select Enable Two-factor
Authentication for each administrator.
Create multiple administrator accounts
Rather than allowing all administrators to access ForiOS with the same administrator account, you can create
accounts for each person or each role that requires administrative access. This configuraion allows you to track
the activities of each administrator or administrative role.
If you want administrators to have different functions you can add different administrator profiles. Go to System
> Admin Profiles and select Create New.
Modify administrator account lockout duration and threshold values
By default, the FortiGate sets the number of password retries at three, allowing the administrator a maximum of
three attempts to log into their account before locking the account for a set amount of time.
Both the number of attempts (admin-lockout-threshold) and the wait time before the administrator can try to
enter a password again (admin-lockout-duration) can be configured within the CLI.
To configure the lockout options:
config system global
set admin-lockout-threshold <failed_attempts>
set admin-lockout-duration <seconds>
end
The default value of admin-lockout-threshold is 3 and the range of values is between 1 and 10. The
admin-lockout-duration is set to 60 seconds by default and the range of values is between 1 and
4294967295 seconds.
Keep in mind that the higher the lockout threshold, the higher the risk that someone may be able to break into the
FortiGate unit.
Hardening
Fortinet Technologies Inc.
19

Global commands for stronger and more secure encryption Security best practices
Example:
To set the admin-lockout-threshold to one attempt and the admin-lockout-duration to a five minute
duration before the administrator can try to log in again, enter the commands:
config system global
set admin-lockout-threshold 1
set admin-lockout-duration 300
end
If the time span between the first failed login attempt and the admin-lockout-
threshold failed login attempt is less than admin-lockout-duration, the lockout
will be triggered.
Rename the admin administrator account
You can improve security by renaming the admin account. To do this, create a new administrator account with the
super_admin admin profile and log in as that administrator. Then go to System > Administrators and edit the
admin administrator and change the User Name. Renaming the admin account makes it more difficult for an
attacker to log into FortiOS.
Add administrator disclaimers
FortiOS can display a disclaimer before or after logging into the GUI or CLI (or both). In either case the
administrator must read and accept the disclaimer before they can proceed.
Use the following command to display a disclaimer before logging in:
config system global
set pre-login-banner enable
end
Use the following command to display a disclaimer after logging in:
config system global
set post-login-banner enable
end
You can customize the replacement messages for these disclaimers by going to System > Replacement
Messages. Select Extended View to view and edit the Administrator replacement messages.
From the CLI:
config system replacemsg admin pre_admin-disclaimer-text
config system replacemsg admin post_admin-disclaimer-text
Global commands for stronger and more secure encryption
This section describes some best practices for employing stronger and more secure encryption.
20 Hardening
Fortinet Technologies Inc.

Security best practices Disable sending Malware statistics to FortiGuard
Turn on global strong encryption
Enter the following command to configure FortiOS to use only strong encryption and allow only strong ciphers
(AES, 3DES) and digest (SHA1) for HTTPS, SSH, TLS, and SSL functions.
config sys global
set strong-crypto enable
end
Disable MD5 and CBC for SSH
In some cases, you may not be able to enable strong encryption. For example, your FortiGate may be
communicating with a system that does not support strong encryption. With strong-crypto disabled you can
use the following options to prevent SSH sessions with the FortiGate from using less secure MD5 and CBC
algorithms:
config sys global
set ssh-hmac-md5 disable
set ssh-cbc-cipher disable
end
Disable static keys for TLS
You can use the following command to prevent TLS sessions from using static keys (AES128-SHA, AES256-
SHA, AES128-SHA256, AES256-SHA256):
config sys global
set ssl-static-key-ciphers disable
end
Require larger values for Diffie-Hellman exchanges
Larger Diffie-Helman values result in stronger encryption. Use the following command to force Diffie-Hellman
exchanges to use 8192 bit values (the highest configurable DH value).
config sys global
set dh-params 8192
end
Setting higher DH but values may not be compatible with some systems that the
FortiGate is communicating with. For example, some versions of FortiClient SSL
VPN may not support 8192 DH bit values. Make sure the DH bit value setting that you
choose is compatible with the systems that your FortiGate will be communicating with.
Disable sending Malware statistics to FortiGuard
By default FortiOS periodically sends encrypted Malware statistics to FortiGuard. The Malware statistics record
Antivirus, IPS, or Application Control events. This data is used to improved FortiGuard services. The Malware
statistics that are sent do not include any personal or sensitive customer data. The information is not shared with
any external parties and is used in accordance with Fortinet's Privacy Policy.
To disable sending Malware statistics to FortiGuard, enter the following command:
Hardening
Fortinet Technologies Inc.
21

