---
id: collect-261001-rattrapage/rattrapage/githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d-2
title: "githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d.md
source_anchor: ""
source_lines: [16, 125]
sha256: 21235386e70b7f634bd2c673065c8b682c35745846e7dbb64fc5d77847927e58
---

# githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d

  |  | #===================================================================== | 
|  | #Bard | 
|  | #configure an email alert on each successful VPN SSL connection on FortiGate 7.2: | 
|  | Go to Security Fabric > Automation. | 
|  | Click Create New. | 
|  | In the Name field, enter a name for the automation stitch. | 
|  | In the Trigger field, select FortiOS Event Log. | 
|  | In the Event field, select SSL VPN Tunnel Up. | 
|  | In the Action field, select Email. | 
|  | Configure the following settings in the Email action: | 
|  | Recipient: Enter the email address that you want to receive the alerts. | 
|  | Subject: Enter the subject line for the alerts. | 
|  | Body: Enter the body of the alerts. | 
|  | Click OK | 
|  | #===================================================================== | 
|  | #disable temporarily SSL VPN | 
|  | VPN > SSL-VPN Settings > Enable SSL-VPN | 
|  | or | 
|  | #assign it a Loopback interface which is in a Down state | 
|  | loopbck1 | 
|  | #===================================================================== | 
|  | #SSL-VPN tunnel interface(ssl.root) policies | 
|  | allow access only to specific destinations and services, not "all" | 
|  | #===================================================================== | 
|  | #disable simultaneous logins for users, per ssl-vpn portal | 
|  | # by default, same VPN user can connect from different locations simultaneously | 
|  | # the connected user will be disconnected when someone else logs in with his/her credentials | 
|  | VPN > SSL-VPN Portals > Limit Users to One SSL-VPN Connection at a Time | 
|  | #===================================================================== | 
|  | #Firewall Hardening | 
|  | #https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/555436/hardening | 
|  | #===================================================================== | 
|  | #Firewall Hardening | 
|  | Administrator access - Management network | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/317406/management-network | 
|  | A single interface or VLAN interface in the management network should be dedicated for all administrative access. | 
|  | Administrative access should be disabled on all other interfaces | 
|  | #Verify that Administrative access should be disabled on all other interfaces | 
|  | FortiGate web interface > Navigate to System > Network > Interfaces > Edit | 
|  | Under "Administrative Access," ensure that the "HTTPS" and "SSH" options are set to "Disable." | 
|  | #CLI | 
|  | show system interface | 
|  | #===================================================================== | 
|  | #Firewall Hardening | 
|  | Administrator access - User authentication for management network access | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/127480/user-authentication-for-management-network-access | 
|  | Who can access the FortiGate | 
|  | Users can log in to the FortiGate by authenticating locally with the FortiGate, | 
|  | or with a remote access server that is integrated with the FortiGate, such as LDAP or RADIUS servers. | 
|  | Remote authentication servers enforce their own password policies. | 
|  | They also provide more configuration options. | 
|  | For example, you can use pre-defined security groups to enable access to a group of users. | 
|  | If an administrator's access needs to be removed, when their account is disabled in the remote access server, | 
|  | they are no longer able to log in to the FortiGate. | 
|  | Do not use shared accounts to access the FortiGate. Shared accounts are more likely to be compromised, | 
|  | are more difficult to maintain as password updates must be disseminated to all users, | 
|  | and make it impossible to audit access to the FortiGate | 
|  | #Ansible Automation | 
|  | If the FortiGate is managed by running scripts over SSH, authenticate users using certificates to avoid storing and maintaining passwords in the application that is making the SSH connection | 
|  | What can administrators access | 
|  | The list of users with access should be audited regularly to ensure that it is current. | 
|  | How can users access the FortiGate | 
|  | Trusted hosts can also be used to specify the IP addresses or subnets that can log in to the FortiGate | 
|  | When authenticating to the FortiGate, implement multi-factor authentication (MFA). This makes it significantly more difficult for an attacker to gain access to the FortiGate. | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | Administrator access - Administrative settings | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/103945/administrative-settings | 
|  | Set the idle timeout time for administrators to a low value, preferably less that ten minutes | 
|  | Use non-standard HTTPS and SSH ports for administrative access. | 
|  | #7.2.3 | 
|  | System > Settings > HTTPS port | 
|  | System > Settings > SSH port | 
|  | System > Settings > Idle timeout | 
|  | The maintainer account has been removed in FortiOS 7.2.4 and later. | 
|  | Replace the certificate that is offered for HTTPS access with a trusted certificate that has the FQDN or IP address of the FortiGate. | 
|  | Configure the Fortinet Security Fabric when multiple FortiGates and fabric devices are used. | 
|  | It provides a single-pane-of-glass administration, allowing administrators access to each device in the fabric using SSO. | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | #Telnet is enabled by default. | 
|  | #config system global | 
|  | #show | 
|  | #config system global | 
|  | #set admin-telnet disable | 
|  | #end | 
|  | #config system global #verify | 
|  | #show | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | Access limitation to internal network – Fortigate Web GUI - Local-In-Policy (Trusted Hosts optional) | 
|  | Country-based VPN access limitation | 
|  | Rename “admin” user | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | #https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/432273/basic-configuration#SystemTime | 
|  | System time: | 
|  | Several FortiGate features rely on an accurate system time, such as logging and certificate related functions. It is recommended that you use a Network Time Protocol (NTP) or Precision Time Protocol (PTP) server to set the system time. If necessary, the system time can be set manually. | 
|  | Dashboard > Status > System Time > FortiGuard NTP | 
|  | System > Settings > Time zone | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/691328/logging-and-reporting | 
|  | The records can be stored locally (data at rest) or remotely (data in motion). | 
|  | Due to the sensitivity of the log data, it is important to encrypt data in motion through the logging transmission channel. | 
|  | Communication with FortiAnalyzer and FortiCloud is encrypted by default | 
|  | When logging to third party devices, make sure that the channel is secure. | 
|  | If it is not secure, it is recommended that you form a VPN to the remote logging device before transmitting logs to it. | 
|  | Logging options include FortiAnalyzer, syslog, and a local disk. | 
|  | Logging with syslog only stores the log messages. | 
