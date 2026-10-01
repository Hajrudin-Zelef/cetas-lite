---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-24
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["distribution", "governance", "incident", "licenses"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [829, 863]
sha256: 94af3f21e3638a8954e51bb99746deeb1c00b2834216df292b1bf117bfc9c54f
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

The CAM is also a critical part of the authentication, authorization, and posture assessment phases of NAC. Although the CAM does not pass client traffic, the impact of its availability needs to be considered in the network design. Like the CAS, the CAM has a high availability solution that provides for a primary server and a hot standby secondary server. In addition, each CAS may be configured with a "fallback" option (as shown in Figure 60) that defines how it will manage client traffic in situations where the CAM is unavailable.
In both high availability CAM and high availability CAS, high availability licenses are used that address the high availability role of the server.
The use of the high availability features is dependent upon the company's unique requirements, but CAS fallback should always configured to ensure that critical network services are available in the event of a network outage.
Figure 60 CAS Fallback
Basic Clean Access Switch Configuration
For OOB-based Clean Access, some simple configuration must be performed on the switches implementing NAC. This configuration is primarily to enable SNMP communication between the switches and the CAM. Table 3 shows a simple SNMP v1 configuration (SNMPv2c and SNMPv3 are supported).
In addition to the switch SNMP configuration, the required trusted and untrusted VLANs must exist and be operational on the switch, as illustrated in the configuration in Table 3. If a switch has more than one IP address, the snmp-server source interface must be specified, as the CAM must be configured with the source IP address that OOB SNMP messages will originate from; alternatively, all IP addresses of interfaces on the switch can be added to the CAM. If SNMP access filtering is applied on the switch (recommended as a best practice), the CAM must be added as a trusted address.
Basic Clean Access Out of Band Switch Configuration
Table 3 SNMPv1 Configuration
NAC CAS Connection
The NAC CAS connects to the core/distribution switch using two switch ports (which are not configured in EtherChannel). One is the untrusted port that serves the VLANs used by the clients prior to their authentication and authorization and the other is the trusted port that connects to the VLANs used once clients have successfully completed the NAC process. Both trusted and untrusted ports are required, even if OOB NAC is used, as the CAS requires access to the trusted VLANs during the NAC process. The following is an example of NAC CAS port configuration:
interface GigabitEthernet1/0/4description NAC Trusted Eth0switchport trunk encapsulation dot1qswitchport trunk allowed vlan 48,57,62switchport mode trunkspanning-tree portfast trunk!interface GigabitEthernet1/0/8description NAC Untrusted Eth1switchport trunk encapsulation dot1qswitchport trunk allowed vlan 61,248,257switchport mode trunkspanning-tree portfast trunk
Basic 802.1X Switch Configuration
The basic 802.1X configuration controls access to an access VLAN depending upon the success or failure of the 802.1X authentication. If the 802.1X authentication is successful, there are three basic options:
•Access to the VLAN configured on the switch port
•Access to the VLAN configured on the switch port and controlled by a access list downloaded from the AAA server
•Access to a VLAN passed to the switch by the AAA server
Table 4 shows an example of an 802.1X configuration.
Table 4 802.1X Switch Configuration
For more information on Cisco 3750 802.1X configuration, refer to:
•Catalyst 3750-E and 3560-E Switch Software Configuration Guide, 12.2(50)SE ->Configuring IEEE 802.1x Port-Based Authentication: http://www.cisco.com/en/US/docs/switches/lan/catalyst3750e_3560e/software/release/12.2_50_se/configuration/guide/sw8021x.html
•Catalyst 2960 Switch Software Configuration Guide, Rel. 12.2(50)SE Configuring IEEE 802.1x Port-Based Authentication: http://www.cisco.com/en/US/docs/switches/lan/catalyst2960/software/release/12.2_50_se/configuration/guide/sw8021x.html
Cisco Security Services
The Cisco SAFE Security Architecture is complimented by Cisco's rich portfolio of security services designed to support the entire solution lifecycle. Security is integrated everywhere and with the help of a lifecycle services approach, enterprises can deploy, operate, and optimize network platforms that defend critical business processes against attack and disruption, protect privacy, and support policy and regulatory compliance controls. Figure 61 shows how the Cisco Lifecycle Security Services support the entire lifecycle.
Figure 61 Cisco Lifecycle Security Services
Strategy and Assessments
Cisco offers a comprehensive set of assessment services based on a structured IT governance, risk management, and compliance approach to information security. These services help the customer understand the needs and gaps, recommend remediation based on industry and international best practices, and help the customer to strategically plan the evolution of an information security program, including updates to security policy, processes, and technology.
Deployment and Migration
Cisco offers deployment services to support the customer in planning, designing, and implementing Cisco security products and solutions. In addition, Cisco has services to support the customer in evolving its security policy and process-based controls to make people and the security architecture more effective.
Remote Management
Cisco Remote Management services engineers become an extension of the customer's IT staff, proactively monitoring the security technology infrastructure and providing incident, problem, change, configuration, and release management, as well as management reporting, 24 hours a day, 365 days a year.
Security Intelligence
The Cisco Security Intelligence services provide early warning intelligence, analysis, and proven mitigation techniques to help security professionals respond to the latest threats. The customer's IT staff can use the latest threat alerts, vulnerability analysis, and applied mitigation techniques developed by Cisco experts who use in-depth knowledge and sophisticated tools to verify anomalies and develop techniques that help ensure timely, accurate, and quick resolution to potential vulnerabilities and attacks.
Security Optimization
The Cisco security Optimization service is an integrated service offering designed to assess, develop, and optimize the customer's security infrastructure on an ongoing basis. Through quarterly site visits and continual analysis and tuning, the Cisco security team becomes an extension of the customer's security staff, supporting them in long-term business security and risk management, as well as near-term tactical solutions to evolving security threats.
