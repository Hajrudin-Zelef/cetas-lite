---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3-2
title: "hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3.md
source_anchor: ""
source_lines: [75, 92]
sha256: b71f30bbfd7b3e565b60b77c1589e3bb018776270a752c1bc716a3ca762d2f07
---

# hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3

- Place the Rule: By default, your custom rule takes precedence over built-in rules but follows other custom rules. Use the "Reorder" option to adjust this hierarchy if needed.
Built-in Firewall Policies
Built-in Firewall policies can be identified via the lock icon. Although these cannot be modified or removed, you can add new policies that overrule them by placing them higher in the table.
Default policies are created as follows:
The built-in firewall policies applied to these zone pairings are:
- Allow All Traffic - Allows all traffic.
The built-in firewall policies applied to these zone pairings are:
- Block Invalid Traffic - Blocks traffic with an invalid firewall connection state.
- Allow All Traffic - Allows all traffic.
The built-in firewall policies applied to these zone pairings are:
- Allow Return Traffic - Allows traffic from the internet that are a reply to traffic sent by devices. This is done by matching the established and related firewall connection states.
- Block Invalid Traffic - Blocks traffic with an invalid firewall connection state.
- Block All Traffic - Blocks all traffic.
Next to these policies, there will be others created depending on which options are configured on the UniFi Gateway. For example, there will be additional policies added when using IPTV Streaming, Port Forwarding or setting up a VPN server.
Important Considerations for Zone-Based Firewall Management
- Removing Custom Zones: Deleting a custom zone will also delete all associated firewall policies. Use caution when a policy spans multiple zones.
- Blocking Traffic to the Gateway Zone: Blocking traffic to the gateway zone may disrupt critical network functions like DHCP and DNS. Always double-check configurations when blocking gateway traffic.
- Blocking All Traffic Between Zones: To block all traffic between zones while allowing specific access, create an allow policy for the desired traffic (e.g., to a storage server's IP) before adding a block policy to deny everything else.
