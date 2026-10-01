---
id: collect-261001-huawei/huawei/hc-en-us-articles-27699646208279-unifi-gateway-advanced-firewall-rules-de392323
title: "hc-en-us-articles-27699646208279-unifi-gateway-advanced-firewall-rules-de392323"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hc-en-us-articles-27699646208279-unifi-gateway-advanced-firewall-rules-de392323.md
source_anchor: ""
source_lines: [1, 47]
sha256: 1f1da14a5f1cad6e041fe9dfe354dfb5b1997707e05fe87df719f5b5438dc94f
---

# hc-en-us-articles-27699646208279-unifi-gateway-advanced-firewall-rules-de392323

UniFi Gateway - Advanced Firewall Rules
| This is an outdated article on our previous firewall approach.  | 
UniFi Gateways include a powerful Firewall engine to maximum security in your network architecture.
For most users, we recommend creating Simple Rules. They provide an intuitive interface that streamlines rule creation for common use-cases such as VLAN segmentation, application and domain filtering, or even bandwidth limiting. To learn more, see our article on Traffic and Firewall Rules.
To learn how to effectively implement network/VLAN and client isolation, click here.
Preconfigured Rules
UniFi pre-configures certain rules to optimize local network traffic, while preventing certain potentially dangerous internet traffic. Additionally, UniFi will configure similar rules for each additional network you add.
Rule Indexing
Firewall rules are executed in order of the Rule Index. A lower number (top of the list) means that the rule is processed before the other rules. When creating a new rule, you can choose to apply it before or after the predefined rules. It is important to be aware of this index because incorrect placement may create the perception of a rule "not working".
Rule Types
The rules are grouped based on the type of network that they apply to. The following network types are used:
- Internet: Contains IPv4 firewall rules that apply to the Internet network.
- LAN: Contains IPv4 firewall rules that apply to the LAN (Corporate) network.
- Guest: Contains IPv4 firewall rules that apply to the Guest network.
- Internet v6: Contains IPv6 firewall rules that apply to the Internet network.
- LAN v6: Contains IPv6 firewall rules that apply to the LAN (Corporate) network.
- Guest v6: Contains IPv6 firewall rules that apply to the Guest network.
Rule Directionality
Besides the network type, the firewall rules also apply to a direction. The following directions are used:
- Local: Applies to traffic that is destined for the UDM/USG itself.
- In: Applies to traffic that is entering the interface (ingress), destined for other networks.
- Out: Applies to traffic that is exiting the interface (egress), destined for this network.
For example, firewall rules configured under LAN In will apply to traffic from the LAN (Corporate) network, destined for other networks. Firewall rules configured under LAN Local will apply to traffic from the LAN (Corporate) network, destined for the UDM/USG itself.
Rule State
In addition to a direction or network type, the firewall rules can also be matched to a state:
- New: The incoming packets are from a new connection.
- Established: The incoming packets are associated with an already existing connection.
- Related: The incoming packets are new, but associated with an already existing connection.
- Invalid: The incoming packets do not match any of the other states.
For example, the predefined Internet Local and Internet In firewall rules ensure that outside connection attempts from the Internet cannot access the UDM/USG and the LAN network behind it. However, the UDM/USG and the LAN network can reach destinations on the Internet and the return traffic is allowed back. The predefined Internet Local and Internet In firewall rules are:
Rule Index: 3001
Enabled: Yes
Description: allow established/related sessions (see states above)
Action: Accept
Protocol: All
Type: Internet In and Internet LocalRule Index: 3002
Enabled: Yes
Description: drop invalid state (see states above)
Action: Drop
Protocol: All
Type: Internet In and Internet Local
IPsec Rule
Firewall rules can also match on traffic that is encrypted with IPsec. This is useful when filtering traffic that is passed over an IPsec Site-to-Site VPN.
- Do not match - Matches all traffic and not specifically IPsec or non-IPsec traffic (default).
- IPsec - Match traffic that is encrypted by IPsec, e.g. passing over a Site-to-Site VPN.
- Non-IPsec - Match specifically on unencrypted traffic.
An example when IPsec matching firewall rules are used is when configuring a Policy-Based IPsec Site-to-Site VPN. The UniFi Gateway will match encrypted traffic from the remote network destined to the local network.
