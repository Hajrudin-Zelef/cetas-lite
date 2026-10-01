---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a.md
source_anchor: ""
source_lines: [75, 120]
sha256: ea669e448b8bef3fef61e45e2331a0e9152296fa664490697e8b91cb2138fbe0
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a

  - For information on Hub-Spoke topology please refer to Configuring Hub-and-spoke VPN Connections on the MX Security Appliance.
- Layer 3 firewall rules on this page are stateful. Layer 7 firewall rules are stateful starting in MX 26.1.
- In NAT/Routed mode, inbound connections are denied by default except ICMP traffic to the appliance. Use port forwarding or NAT policies to allow inbound traffic.
- To determine the priority of Layer 3 vs Layer 7 rules, please refer to Layer 3 and 7 Firewall Processing Order.
Layer 7 firewall rules are stateful starting in MX 26.1.
Template Firewall Rules
Additional options are available when configuring firewall rules on a configuration template. For details, see the Firewall rules for templates section of the Configuration Templates page.
Cellular Failover Rules
These cellular firewall rules are evaluated separately from, and after, the existing L3 outbound firewall rule table when the appliance has failed over to a cellular modem uplink. For example, an Allow Any Any rule in the L3 outbound firewall rules allows the packet past that ruleset, but the packet is still subsequently evaluated against the cellular firewall rules. If a packet is denied by the L3 outbound firewall rule table, the cellular firewall rule table is not checked. This can be useful for restricting cellular traffic to business-critical uses and helping prevent unnecessary cellular overages.
Note: As the cellular failover rules are appended, if a deny any any is applied to the general Layer 3 firewall rules, cellular failover rules will never trigger.
Note: Prior to MX version 26.x, the default Inbound Cellular Failover Rule was to allow all. In MX version 26.x and above, the default Inbound Cellular Failover Rule is to deny all.
FQDN Support
Fully qualified domain names can be configured in the Destination field.
Note: At this time, cellular firewall rules do not support the use of FQDNs.
FQDN is not supported as a source or on inbound firewall rules.
FQDN-based L3 firewall rules are implemented based on snooping DNS traffic. When a client device attempts to access a web resource, the MX will track the DNS requests and response to learn the IP of the web resource returned to the client device.
There are several important considerations for utilizing and testing this configuration:
- The MX must see a client DNS request and the server's response to learn the proper IP mapping. The DNS request does not need to be from a specific client, the MX will map the FQDN based on the DNS response, regardless of which client made the DNS request. The communication between the client and DNS server can be an inter-VLAN, VPN, or NATed traffic flow, but cannot be intra-VLAN (this DNS traffic is not snooped).
- In some cases, a client device may already have IP information about the web resource it is attempting to access. This could be due to the client having cached a previous DNS response, or a local statically configured DNS entry on the device. The MX may not be able to properly block or allow communications to the web resource in these cases if the client devices do not generate a DNS request for the MX to inspect. Clients caching DNS query responses for longer than the TTL reported in the response must send new queries in order to maintain the IP mapping for that domain.
- FQDN rules imply a wildcard when no subdomain is used by prepending "*" to the domain.tld. This wildcard is not shown on the Dashboard but is visible in syslog messages if syslog is configured for a network. For example, a rule to permit "yahoo.com" would permit any subdomain under yahoo.com such as mail.yahoo.com. Permitting "mail.yahoo.com" in the rule would only permit mail.yahoo.com and not the TLD or other subdomain of yahoo.com.
MX will not be able to snoop TCP-based or encrypted DNS traffic.
An example configuration is included below:
To ensure successful operation, DNS traffic must be allowed by the MX's Layer 3 firewalls. Blocking DNS will result in the MX being unable to learn hostname and IP address mappings and, subsequently, prevent it from blocking or allowing traffic as expected.
Behavior when a Canonical Name (CNAME) record point to a different Top Level Domain (TLD)
If your lookup uses a CNAME that resolves to a different Top Level Domain, FQDN rules will not be applied to that traffic. (Example: Lookup for "update.microsoft.com" resolves to "update.microsoft.com.akadns.net")
WAN Appliance Services
- ICMP Ping: Use this setting to allow the WAN appliance to reply to inbound ICMP ping requests coming from the specified address(es). Supported values for the remote IP address field include None, Any, or a specific IP range (using CIDR notation). You can also enter multiple IP ranges separated by commas. To add specific IP addresses rather than ranges, use the format X.X.X.X/32. This setting has never applied to cellular. To allow ICMP to cellular uplink on MX version 26.x and above, create an Inbound Cellular Failover Rule allowing ICMP.
- Web (local status & configuration): Use this setting to allow or disable access to the local management page (wired.meraki.com) via the WAN IP of the WAN appliance. Supported values for the remote IPs field are the same as for ICMP Ping.
- SNMP: Use this setting to allow SNMP polling of the appliance from the WAN. Supported values for the remote IPs field are the same as for ICMP Ping.
The WAN appliance services configuration section is removed for WAN appliances operating in Passthrough or VPN Concentrator mode as these are configured in the inbound firewall.
Layer 7 Firewall Rules
Using Meraki's unique Layer 7 traffic analysis technology, it is possible to create firewall rules to block specific web-based services, websites, or types of websites without specifying IP addresses or port ranges. This can be particularly useful when applications or websites use more than one IP address, or when their IP addresses or port ranges are subject to change.
It is possible to block applications by category (e.g. 'All video & music sites') or for a specific type of application within a category (e.g. only iTunes within the 'Video & music' category). The figure below illustrates a set of Layer 7 firewall rules that includes both blocking entire categories and blocking specific applications within a category:
It is also possible to block traffic based on HTTP hostname, destination port, remote IP range, and destination IP/port combinations.
The remote IPs cannot be blocked inbound for L2TP VPN or AnyConnect VPN.
Geo-IP Based Firewall Rules
The Layer 7 firewall can block traffic based on the destination country of outbound traffic and the source country of return traffic. Geo-IP classification is provided by MaxMind.
To configure a Geo-IP firewall rule, create a Layer 7 firewall rule and select Countries... from the Application drop-down. You can block traffic to selected countries, or block traffic to all countries except the countries you specify.
Considerations and Limitations
- Geo-IP rules supports IPv4 traffic only. IPv6 traffic is not evaluated against Geo-IP firewall rules.
- Geo-IP rules are supported only on physical MX appliances with an Advanced Security license. They are not supported on Z3 devices or vMX appliances.
- Geo-IP rules do not support exceptions for specific IP ranges within a blocked country. For more granular policy control, use an organization-wide Group Policy, which supports both Layer 3 and Layer 7 firewall rules in a single ruleset.
- Geo-IP rules also apply to internally routed traffic. If a subnet configured on Security & SD-WAN > Configure > Addressing & VLANs geolocates to a blocked country, the MX will drop traffic sourced from that subnet.
Forwarding rules
Use this area to configure port forwarding rules and 1:1 NAT mappings as desired.
Port forwarding
