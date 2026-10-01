---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-4
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [405, 518]
sha256: 20c27931afb07e52c8188d8d757cfde58565370a5248854912a676212537feda
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

Security best practices 16
System administrator best practices
This section describes a collection of changes you can implement to make administrative access to the GUI and CLI
more secure.
Disable administrative access to the external (Internet-facing) interface
When possible, don’t allow administration access on the external (Internet-facing) interface.
To disable administrative access, go to Network > Interfaces, edit the external interface and disable HTTPS, PING,
HTTP, SSH, and TELNET under Administrative Access.
From the CLI:
config system interface
edit <external-interface-name>
unset allowaccess
end
Allow only HTTPS access to the GUI and SSH access to the CLI
For greater security never allow HTTP or Telnet administrative access to a FortiGate interface, only allow HTTPS and
SSH access. You can change these settings for individual interfaces by going to Network > Interfaces and adjusting the
administrative access to each interface.
From the CLI:
config system interface
edit <interface-name>
set allowaccess https ssh
end
Require TLS 1.2 for HTTPS administrator access
Use the following command to require TLS 1.2 for HTTPS administrator access to the GUI:
config system global
set admin-https-ssl-versions tlsv1-2
end
TLS 1.2 is currently the most secure SSL/TLS supported version for SSL-encrypted administrator access.
Re-direct HTTP GUI logins to HTTPS
Go to System > Settings > Administrator Settings and enable Redirect to HTTPS to make sure that all attempted HTTP
login connections are redirected to HTTPS.
From the CLI:
config system global
set admin-https-redirect enable
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 17
Change the HTTPS and SSH admin access ports to non-standard ports
Go to System > Settings > Administrator Settings and change the HTTPS and SSH ports.
You can change the default port configurations for HTTPS and SSH administrative access for added security. To
connect to a non-standard port, the new port number must be included in the collection request. For example:
l If you change the HTTPS port to 7734, you would browse to https://<ip-address>:7734.
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
To set the administrator idle timeout, go to System > Settings and enter the amount of time for the Idle timeout. A best
practice is to keep the default time of 5 minutes.
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
Setting up trusted hosts for an administrator limits the addresses from where they can log into FortiOS. The trusted hosts
configuration applies to most forms of administrative access including HTTPS, SSH, and SNMP. When you identify a
trusted host for an administrator account, FortiOS accepts that administrator’s login only from one of the trusted hosts. A
login, even with proper credentials, from a non-trusted host is dropped.
Even if you have configured trusted hosts, if you have enabled ping administrative access on a
FortiGate interface, it will respond to ping requests from any IP address.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 18
To identify trusted hosts, go to System > Administrators, edit the administrator account, enable Restrict login to trusted
hosts, and add up to ten trusted host IP addresses.
To add two trusted hosts from the CLI:
config system admin
edit <administrator-name>
set trustedhost1 172.25.176.23 255.255.255.255
set trustedhost2 172.25.177.0 255.255.255.0
end
Trusted host IP addresses can identify individual hosts or subnets. Just like firewall policies, FortiOS searches through
the list of trusted hosts in order and acts on the first match it finds. When you configure trusted hosts, start by adding
specific addresses at the top of the list. Follow with more general IP addresses. You don't have to add addresses to all of
the trusted hosts as long as all specific addresses are above all of the 0.0.0.0 0.0.0.0 addresses.
Set up two-factor authentication for administrators
FortiOS supports FortiToken and FortiToken Mobile 2-factor authentication. FortiToken Mobile is available for iOS and
Android devices from their respective application stores.
Every registered FortiGate unit includes two trial tokens for free. You can purchase additional tokens from your reseller
or from Fortinet.
To assign a token to an administrator, go to System > Administrators and select Enable Two-factor Authentication for
each administrator.
Create multiple administrator accounts
Rather than allowing all administrators to access ForiOS with the same administrator account, you can create accounts
for each person or each role that requires administrative access. This configuration allows you to track the activities of
each administrator or administrative role.
If you want administrators to have different functions you can add different administrator profiles. Go to System > Admin
Profiles and select Create New.
Modify administrator account lockout duration and threshold values
By default, the FortiGate sets the number of password retries at three, allowing the administrator a maximum of three
attempts to log into their account before locking the account for a set amount of time.
Both the number of attempts (admin-lockout-threshold) and the wait time before the administrator can try to enter
a password again (admin-lockout-duration) can be configured within the CLI.
To configure the lockout options:
config system global
set admin-lockout-threshold <failed_attempts>
set admin-lockout-duration <seconds>
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

