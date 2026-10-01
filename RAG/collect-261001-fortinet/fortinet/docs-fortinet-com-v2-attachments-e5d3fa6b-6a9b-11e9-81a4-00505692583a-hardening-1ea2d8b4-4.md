---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-4
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [418, 532]
sha256: 0ac1b7cb3023ab4cae8e83b07d1fa53369ec091a437b6b76a9b6db6f520d7f0a
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

Security best practices 18
Change the HTTPS and SSH admin access ports to non-standard ports
Go to System > Settings > Administrator Settings and change the HTTPS and SSH ports.
You can change the default port configurations for HTTPS and SSH administrative access for added security. To
connect to a non-standard port, the new port number must be included in the collection request. For example:
l If you change the HTTPS port to 7734, you would browse tohttps://<ip-address>:7734.
l If you change the SSH port to 2345, you would connect to ssh admin@<ip-address>:2345
To change the HTTPS and SSH login ports from the CLI:
config system global
set admin-sport 7734
set admin-ssh-port 2345
end
If you change to the HTTPS or SSH port numbers, make sure your changes do not conflict with ports used for other
services.
Maintain short login timeouts
Set the idle timeout to a short time to avoid the possibility of an administrator walking away from their management
computer and leaving it exposed to unauthorized personnel.
To set the administrator idle timeout, go to System > Settings and enter the amount of time for the Idle timeout. A
best practice is to keep the default time of 5 minutes.
To set the administrator idle timeout from the CLI:
config system global
set admintimeout 5
end
You can use the following command to adjust the grace time permitted between making an SSH connection and
authenticating. The range can be between 10 and 3600 seconds, the default is 120 seconds (minutes). By shortening
this time, you can decrease the chances of someone attempting a brute force attack a from being successful. For
example, you could set the time to 30 seconds.
config system global
set admin-ssh-grace-time 30
end
Restrict logins from trusted hosts
Setting up trusted hosts for an administrator limits the addresses from where they can log into FortiOS. The trusted
hosts configuration applies to most forms of administrative access including HTTPS, SSH, and SNMP. When you
identify a trusted host for an administrator account, FortiOS accepts that administrator’s login only from one of the
trusted hosts. A login, even with proper credentials, from a non-trusted host is dropped.
Even if you have configured trusted hosts, if you have enabled ping administrative access on a
FortiGate interface, it will respond to ping requests from any IP address.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 19
To identify trusted hosts, go to System > Administrators, edit the administrator account, enable Restrict login to
trusted hosts, and add up to ten trusted host IP addresses.
To add two trusted hosts from the CLI:
config system admin
edit <administrator-name>
set trustedhost1 172.25.176.23 255.255.255.255
set trustedhost2 172.25.177.0 255.255.255.0
end
Trusted host IP addresses can identify individual hosts or subnets. Just like firewall policies, FortiOS searches through
the list of trusted hosts in order and acts on the first match it finds. When you configure trusted hosts, start by adding
specific addresses at the top of the list. Follow with more general IP addresses. You don't have to add addresses to all
of the trusted hosts as long as all specific addresses are above all of the 0.0.0.0 0.0.0.0 addresses.
Set up two-factor authentication for administrators
FortiOS supports FortiToken and FortiToken Mobile 2-factor authentication. FortiToken Mobile is available for iOS and
Android devices from their respective application stores.
Every registered FortiGate unit includes two trial tokens for free. You can purchase additional tokens from your reseller
or from Fortinet.
To assign a token to an administrator, go to System > Administratorsand select Enable Two-factor
Authenticationfor each administrator.
Create multiple administrator accounts
Rather than allowing all administrators to access ForiOS with the same administrator account, you can create accounts
for each person or each role that requires administrative access. This configuration allows you to track the activities of
each administrator or administrative role.
If you want administrators to have different functions you can add different administrator profiles. Go to System
> Admin Profiles and select Create New.
Modify administrator account lockout duration and threshold values
By default, the FortiGate sets the number of password retries at three, allowing the administrator a maximum of three
attempts to log into their account before locking the account for a set amount of time.
Both the number of attempts (admin-lockout-threshold) and the wait time before the administrator can try to
enter a password again (admin-lockout-duration) can be configured within the CLI.
To configure the lockout options:
config system global
set admin-lockout-threshold <failed_attempts>
set admin-lockout-duration <seconds>
end
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 20
The default value ofadmin-lockout-threshold is 3 and the range of values is between 1 and 10. Theadmin-
lockout-duration is set to 60 seconds by default and the range of values is between 1 and 4294967295 seconds.
Keep in mind that the higher the lockout threshold, the higher the risk that someone may be able to break into the
FortiGate.
Example:
To set theadmin-lockout-threshold to one attempt and theadmin-lockout-duration to a five minute
duration before the administrator can try to log in again, enter the commands:
config system global
set admin-lockout-threshold 1
set admin-lockout-duration 300
end
If the time span between the first failed login attempt and theadmin-lockout-
threshold failed login attempt is less thanadmin-lockout-duration, the lockout will
be triggered.
Rename the admin administrator account
You can improve security by renaming the admin account. To do this, create a new administrator account with the
super_admin admin profile and log in as that administrator. Then go to System > Administratorsand edit the admin
administrator and change the User Name. Renaming the admin account makes it more difficult for an attacker to log
into FortiOS.
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
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

