---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-7
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [243, 349]
sha256: 9a787ccd37de59a2456cc542dc6b10f3345d7102767b59cf4d4ddcf343ee0bfa
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

!--- The source of the traffic should be known and authorized.
!
!--- Permit external BGP to peer 64.104.10.113 
access-list 110 permit tcp host 64.104.10.114 host 64.104.10.113 eq bgp
access-list 110 permit tcp host 64.104.10.114 eq bgp host 64.104.10.113
!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!--- Module 3:  Explicit Deny to Protect Infrastructure
access-list 110 deny ip 64.104.10.0 0.0.0.255 any
!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!--- Module 4:  Explicit Permit for Traffic to School's Public
!--- Subnet.
access-list 110 permit ip any 198.133.219.0 0.0.0.255
!
Note The 64.104.0.0/16 and 198.133.219.0/24 address blocks used in the examples here provided are reserved for the exclusive use of Cisco Systems, Inc.
Internet Firewall
The mission of the Internet firewall is to protect the school's internal resources and data from external threats, secure the public services provided by the DMZ, and to control user's traffic to the Internet. The Schools Service Ready architecture uses a Cisco ASA appliance as illustrated in Figure 4-11.
Figure 4-11 Internet Edge Firewall
The Cisco ASA is implemented with three interface groups, each one representing a distinct security domain:
•Inside—The Inside is the interface connecting to the core/distribution switch that faces the interior of the network where internal users and resources reside.
•Outside—Interface connecting to the Internet border router. The router may be managed either by the school or a service provider.
•Demilitarized Zone (DMZ)—The DMZ hosts school services that are accessible over the Internet. These services may include a web portal and E-mail services.
The Internet firewall acts as the primary gateway to the Internet; Therefore, its deployment should be carefully planned. The following are key aspects to be considered when implementing the firewall:
•Firewall Hardening and Monitoring
•Network Address Translation (NAT)
•Firewall Access Policies
•Firewall Redundancy
•Routing
Firewall Hardening and Monitoring
The Cisco ASA should be hardened in a similar fashion as the infrastructure routers and switches. According to the Cisco SAFE security best practices, the following is a summary of the measures to be taken:
•Implement dedicated management interfaces to the OOB management network.
•Present legal notification for all access attempts.
•Use HTTPS and SSH for device access. Limit access to known IP addresses used for administrative access.
•Configure AAA for role-based access control and logging. Use a local fallback account in case AAA server is unreachable.
•Use NTP to synchronize the time.
•Use syslog or SNMP to keep track of system status, traffic statistics, and device access information.
•Authenticate routing neighbors and log neighbor changes.
•Implement firewall access policies (explained in the Firewall Access Policies).
The Cisco ASA 5510 and higher appliance models come with a dedicated management interface that should be used whenever possible. Using a dedicated management interface keeps the management plane of the firewall isolated from threats originating from the data plane. The management interface should connect to the OBB management network, if one is available.
The following is an example of the configuration of a dedicated management interface.
interface Management0/0
 nameif management
 security-level 100
 ip address 172.26.160.225 255.255.252.0 
 management-only
!
Note Any physical interface or logical sub-interface can be configured as a management-only interface using the management-only command.
It is recommended that a legal notification banner is presented on all interactive sessions to ensure that users are notified of the security policy being enforced and to which they are subject. The notification banner should be written in consultation with your legal advisors.
The following example displays the banner after the user logs in:
banner motd UNAUTHORIZED ACCESS TO THIS DEVICE IS PROHIBITED. 
banner motd You must have explicit, authorized permission to access or configure this 
device. 
banner motd Unauthorized attempts and actions to access or use this system may result in 
civil and/or criminal penalties. 
banner motd All activities performed on this device are logged and monitored.
Management access to the firewall should be restricted to SSH and HTTPS. SSH is needed for CLI access and HTTPS is needed for the firewall GUI-based management tools such as CSM and ADSM. Additionally, this access should only be permitted for users authorized to access the firewalls for management purposes.
The following ASA configuration fragment illustrates the configuration needed to generate a 768 RSA key pair and enabling SSH and HTTPS access for devices located in the management subnet.
! Generate RSA key pair with a key modulus of 768 bits
crypto key generate rsa modulus 768
! Save the RSA keys to persistent flash memory
write memory
! enable HTTPS
http server enable
! restrict HTTPS access to the firewall to permitted management stations
http <CSM/ADSM-IP-address> 255.255.255.255 management
! restrict SSH access to the firewall to well-known administrative systems
ssh <admin-host-IP-address> 255.255.255.255 management
! Configure a timeout value for SSH access to 5 minutes
ssh timeout 5
Administrative users accessing the firewalls for management must be authenticated, authorized, and access should be logged using AAA. The following ASA configuration fragment illustrates the AAA configurations needed to authenticate, authorize, and log user access to the firewall:
aaa-server tacacs-servers protocol tacacs+
 reactivation-mode timed
aaa-server tacacs-servers host <ACS-Server>
 key <secure-key>
aaa authentication ssh console tacacs-servers LOCAL
aaa authentication serial console tacacs-servers LOCAL
aaa authentication enable console tacacs-servers LOCAL
aaa authentication http console tacacs-servers LOCAL
aaa authorization command tacacs-servers LOCAL
aaa accounting ssh console tacacs-servers
aaa accounting serial console tacacs-servers
aaa accounting command tacacs-servers
aaa accounting enable console tacacs-servers
aaa authorization exec authentication-server
! define local username and password for local authentication fallback
username admin password <secure-password> encrypted privilege 15
As with the other infrastructure devices in the network, it is important to synchronize the time on the firewall protecting the management module using NTP.
The following configuration fragment illustrates the NTP configuration needed on an ASA to enable NTP to an NTP server located in the management network:
ntp authentication-key 10 md5 *
ntp authenticate
ntp trusted-key 10
ntp server <NTP-Server-address> source management
Syslog and SNMP can be used to keep track of system status, device access, and session activity. NetFlow Security Event Logging (NSEL), now supported on all Cisco ASA models, may also be used for the monitoring and reporting of session activity. The following configuration fragment illustrates the configuration of Syslog.
logging trap informational
logging host management <Syslog-Server-address>
logging enable
The routing protocol running between the Internet firewall and the core/distribution should be secured. The following ASA configuration fragment illustrates the use of EIGRP MD5 authentication to authenticate the peering session between the inside firewall interface and the core/distribution switch:
interface Redundant1
 nameif inside
 security-level 100
 ip address 10.125.33.10 255.255.255.0 
 authentication key eigrp 100 <removed> key-id 1
 authentication mode eigrp 100 md5
Network Address Translation (NAT)
NAT is required because the school typically gets a limited number of public IP addresses. In addition, NAT helps shield the school's internal address space from reconnaissance and another malicious activity.
The following illustrates the NAT configuration:
