---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-673634-traffic-shaping-policies-967a67b7
title: "document-fortigate-7-4-0-administration-guide-673634-traffic-shaping-policies-967a67b7"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-673634-traffic-shaping-policies-967a67b7.md
source_anchor: ""
source_lines: [1, 44]
sha256: 80f2cffebf7f0615077180b0ee6d1c30be50e0f6e62c9f7f7e5d5712e17f0bfd
---

# document-fortigate-7-4-0-administration-guide-673634-traffic-shaping-policies-967a67b7

Traffic shaping policies
Traffic shaping policies
As mentioned in Traffic shaping, traffic shaping starts with the traffic shaping policy. Traffic shaping policies are used to map traffic to a traffic shaper or assign them to a class. Traffic is then shaped by the shaper or the shaping profile that is applied on an interface.
Traffic can also be shaped by applying traffic shapers directly on a firewall policy. However, this legacy approach can only be configured from the CLI, and is not a preferred method for applying traffic shaping. As the number of firewall policies increases, managing shaping on each individual policy becomes increasingly difficult. For the same reason, it is also not recommended to mix the legacy approach with traffic shaping policies to avoid the added complexity.
Overview
A traffic shaping policy is a rule that matches traffic based on certain IP header fields and/or upper layer criteria. When traffic hits the firewall, the FortiGate will first look up a firewall policy, and then match a shaping policy. The matching traffic will apply a traffic shaper, class ID, or assign a DSCP DiffServ tag to the outgoing traffic.
The traffic shaping policies must be placed in the correct order in the traffic shaping policy list page to obtain the desired results. Policies are matched from top-down, so the traffic shaping policies should be arranged in a sequence that places the more granular policies above general policies.
The policy can be configured by going to Policy & Objects > Traffic Shaping and selecting the Traffic Shaping Policies tab. If the menu does not display the traffic shaping settings, go to System > Feature Visibility and enable Traffic Shaping.
Configuring traffic shaping policies
A traffic shaping policy can be split into two parts:
- Options used to match the traffic
- Options used to apply actions to the matched traffic
In the GUI, the options are configured in the If Traffic Matches and Then sections. In the CLI, all options are configured under config firewall shaping-policy. Some options can only be configured from the CLI.
The following options can be configured for traffic matching criteria:
| GUI option |  | CLI option | Description | 
|---|---|---|---|
| Source |  |  |  | 
|  | Address | set srcaddr <address_object> | Select the address object to match the source IP. | 
|  | User | set users <user_object> | Select the user object to match the user authenticated for the session. | 
|  | Internet Service | set internet-service-src enable set internet-service-src-name <name> set internet-service-src-group <group> set internet-service-src-custom <custom> set internet-service-src-custom-group <custom_group> | Select the internet service to match the source of the incoming traffic. Internet service currently cannot be used with source address. | 
| Destination |  |  |  | 
|  | Address | set dstaddr <address_object> | Select the address object to match the destination IP. | 
|  | Internet Service | set internet-service enable set internet-service-name <name> set internet-service-group <group> set internet-service-custom <custom> set internet-service-custom-group <custom_group> | Select the internet service to match the destination of the incoming traffic. Internet service currently cannot be used with destination address and service. | 
| Schedule |  | set schedule <schedule> | Enable to select a schedule (one-time, recurring, or group). | 
| Service |  | set service <service> | Select the service or service group for the traffic. | 
| Application |  |  | Application control must be enabled in the related firewall policy to learn the application of the traffic. | 
|  | Application | set application <application> | Select the application to match the application of the traffic. | 
|  | Category | set app-category <category> | Select the application category to match the application of the traffic. | 
|  | Group | set app-group <groups> | Select the application group to match the application of the traffic. | 
| URL Category |  | set url-category <category> | Select the URL category to match the URL of the traffic. A web filter profile must be enabled in the related firewall policy to know the URL of the traffic (see Web filter). | 
| n/a |  | set tos-mask <hexadecimal_mask> set tos <value> set tos-negate {enable \| disable} | Specify the type of service (ToS) and mask to match. These options can only be configured in the CLI. | 
The following options can be configured for actions to apply to the matched traffic:
| GUI option |  | CLI option | Description | 
|---|---|---|---|
| Outgoing interface |  | set dstintf <interface> | Select the destination interface that the traffic shaping applies to (required). | 
| Apply shaper |  |  |  | 
|  | Shared shaper | set traffic-shaper <shaper> | Select the shared shaper to be applied to traffic in the ingress-to-egress direction. For example, on traffic that egresses on the wan interface, the shaper is applied to upload or outbound traffic. | 
|  | Reverse shaper | set traffic-shaper-reverse <shaper> | Select the reverse shaper to be applied to traffic in the egress-to-ingress direction. For example, on traffic that egresses on the wan interface, the shaper is applied to download or inbound traffic. | 
|  | Per-IP shaper | set per-ip-shaper <shaper> | Select the per-IP shaper. Per-IP shapers affect downloads and uploads. The allotted bandwidth applies to each individual IP. In a shared shaper, the allotted bandwidth applies to all IPs. | 
| Assign shaping class ID |  |  |  | 
|  | Traffic shaping class ID | set class-id <class> | Set the class ID to apply the matching traffic. Class IDs are further prioritized within a traffic shaping profile and applied to an interface. | 
| n/a |  | set diffserv-forward {enable \| disable} set diffservcode-forward <code> set diffserv-reverse {enable \| disable} set diffservcode-reverse <code> | Specify the settings to apply a DSCP tag to the forward or reverse traffic. The DiffServ code is in 6-bit binary format. These options can only be configured in the CLI. | 
Traffic shapers and class IDs can be applied at the same time when configuring traffic shaping policies. However, to reduce the complexity, it is recommended to use one method over the other.
The following topics include examples with traffic shaping policies:
