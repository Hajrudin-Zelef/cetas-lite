---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-16
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [549, 583]
sha256: 67bfbd620f0823d285ec5cfef1bc3921ff736b0af0c5798f6713eb34f88af19d
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

•Use syslog or SNMP to keep track of system status, traffic statistics, and device access information.
•Authenticate routing neighbors and log neighbor changes.
•Implement firewall access policies (explained in Firewall Access Policies).
The Cisco ASA 5510 and higher appliance models come with a dedicated management interface that should be used whenever possible. Using a dedicated management interface keeps the management plane of the firewall isolated from threats originating from the data plane. The management interface should connect to the OOB management network, if one is available.
The following is an example of the configuration of a dedicated management interface.
interface Management0/0nameif managementsecurity-level 100ip address 172.26.160.225 255.255.252.0management-only!
Note Any physical interface or logical sub-interface can be configured as a management-only interface using the management-only command.
It is recommended that a legal notification banner is presented on all interactive sessions to ensure that users are notified of the security policy being enforced and to which they are subject. The notification banner should be written in consultation with your legal advisors.
The following example displays the banner after the user logs in:
banner motd UNAUTHORIZED ACCESS TO THIS DEVICE IS PROHIBITED.banner motd You must have explicit, authorized permission to access or configure this device.banner motd Unauthorized attempts and actions to access or use this system may result in civil and/or criminal penalties.banner motd All activities performed on this device are logged and monitored.
Management access to the firewall should be restricted to SSH and HTTPS. SSH is needed for CLI access and HTTPS is needed for the firewall GUI-based management tools such as CSM and ADSM. Additionally, this access should only be permitted for users authorized to access the firewalls for management purposes.
The following ASA configuration fragment illustrates the configuration needed to generate a 768 RSA key pair and enabling SSH and HTTPS access for devices located in the management subnet.
! Generate RSA key pair with a key modulus of 768 bitscrypto key generate rsa modulus 768! Save the RSA keys to persistent flash memorywrite memory! enable HTTPShttp server enable! restrict HTTPS access to the firewall to permitted management stationshttp <CSM/ADSM-IP-address> 255.255.255.255 management! restrict SSH access to the firewall to well-known administrative systemsssh <admin-host-IP-address> 255.255.255.255 management! Configure a timeout value for SSH access to 5 minutesssh timeout 5
Administrative users accessing the firewalls for management must be authenticated, authorized, and access should be logged using AAA. The following ASA configuration fragment illustrates the AAA configurations needed to authenticate, authorize, and log user access to the firewall:
aaa-server tacacs-servers protocol tacacs+reactivation-mode timedaaa-server tacacs-servers host <ACS-Server>key <secure-key>aaa authentication ssh console tacacs-servers LOCALaaa authentication serial console tacacs-servers LOCALaaa authentication enable console tacacs-servers LOCALaaa authentication http console tacacs-servers LOCALaaa authorization command tacacs-servers LOCALaaa accounting ssh console tacacs-serversaaa accounting serial console tacacs-serversaaa accounting command tacacs-serversaaa accounting enable console tacacs-serversaaa authorization exec authentication-server! define local username and password for local authentication fallbackusername admin password <secure-password> encrypted privilege 15
As with the other infrastructure devices in the network, it is important to synchronize the time on the firewall protecting the management module using NTP.
The following configuration fragment illustrates the NTP configuration needed on an ASA to enable NTP to an NTP server located in the management network:
ntp authentication-key 10 md5 *ntp authenticatentp trusted-key 10ntp server <NTP-Server-address> source management
Syslog and SNMP can be used to keep track of system status, device access, and session activity. NetFlow Security Event Logging (NSEL), now supported on all Cisco ASA models, may also be used for the monitoring and reporting of session activity. The following configuration fragment illustrates the configuration of Syslog.
logging trap informationallogging host management <Syslog-Server-address>logging enable
The routing protocol running between the Internet firewall and the distribution/core should be secured. The following ASA configuration fragment illustrates the use of EIGRP MD5 authentication to authenticate the peering session between the inside firewall interface and the core/distribution switch:
interface Redundant1nameif insidesecurity-level 100ip address 10.125.33.10 255.255.255.0authentication key eigrp 100 <removed> key-id 1authentication mode eigrp 100 md5
Network Address Translation (NAT)
NAT is required because the enterprise typically gets a limited number of public IP addresses. In addition, NAT helps shield the company's internal address space from reconnaissance and another malicious activity.
The following illustrates the NAT configuration:
! Static translation for servers residing at DMZstatic (dmz,outside) 198.133.219.10 10.25.34.10 netmask 255.255.255.255static (dmz,outside) 198.133.219.11 10.25.34.11 netmask 255.255.255.255static (dmz,outside) 198.133.219.12 10.25.34.12 netmask 255.255.255.255static (dmz,outside) 198.133.219.13 10.25.34.13 netmask 255.255.255.255!! Dynamic Port Address Translation (PAT) for inside hosts going to the Internetglobal (outside) 10 interfacenat (inside) 10 10.0.0.0 255.0.0.0!! Static translation for inside hosts going to the DMZ and vice-versa. The inside IP addresses are visible to the DMZ.static (inside,dmz) 10.0.0.0 10.0.0.0 netmask 255.0.0.0
Firewall Access Policies
As previously explained, the Internet firewall should be configured to:
•Protect the company's internal resources and data from external threats by preventing incoming access from the Internet.
•Protect public resources served by the DMZ by restricting incoming access to the public services and by limiting outbound access from DMZ resources out to the Internet.
•Control users' Internet-bound traffic.
Enforcing such policies requires the configuration of the appropriate interface security levels and the deployment of ACLs governing what traffic is allowed or prevented from transiting between interfaces.
By default, the Cisco ASA appliance allows traffic from higher to lower security level interfaces (i.e., from inside to outside). However, due to the sensitivity of enterprise environments, the security administrators are recommended to override the default rules with more stringent rules indicating exactly what ports and protocols are permitted.
In our configuration example the inside, DMZ, and outside interfaces were configured with the security levels of 100, 50, and 0 respectively. With this, by default any traffic originating from the inside to the DMZ, from inside to outside, and from the DMZ to the outside interface will be allowed freely. At the same time, any traffic originating from the outside to the DMZ, from the outside to the inside, and from the DMZ to the inside interface will be blocked. While this may satisfy the basic access control requirements of the organization, it is always a good idea to reinforce the policies by enforcing granular ACLs.
Before designing the ACLs, it should also be noted that as the Cisco ASA inspects traffic it is able to recognize packets belonging to already established sessions. The stateful inspection engine of the firewall dynamically allows the returning traffic. Therefore, the firewall ACLs should be constructed to match traffic in the direction in which it is being initiated. In our sample configurations ACLs are applied in the ingress direction.
