---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-6
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [461, 490]
sha256: e8ff9d09c31aa7b95d87c70373a33acc859d4ba479aea3fffb07472e49528541
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

In OPNsense, port forwarding enables you to redirect incoming network traffic from one port to another on your internal network, or from one IP address to another. It is frequently employed to facilitate external access to services that are operating on devices within your private network, including web servers, gaming servers, and remote desktop access. Port forwarding in OPNsense operates as follows:

1. 
**Incoming Traffic** : Traffic is directed to your public IP address on a specific port and arrives at your OPNsense firewall/router from the internet.
2. 
**Firewall NAT Rules** : OPNsense evaluates its NAT (Network Address Translation) rules to ascertain whether the incoming traffic corresponds to any configured port forwarding rule.
3. 
**Port Forwarding Rule Matching** : OPNsense forwards traffic to the specified internal IP address and port on your private network if a matching rule is found.
4. 
**Traffic Routing** : The internal device receives the traffic, processes the request, and responds to the firewall/router.
5. 
**Reverse NAT** : OPNsense subsequently performs reverse NAT, rewriting the source address of the reply packets to your public IP address before transmitting them back to the original external sender.

Assume that a web server is operating on your internal network at IP `192.168.1.100` and is listening on port `80`.  You desire that external users have access to this web server via port `8080` on your public IP address. Port forwarding would operate as follows in this scenario.

| Internal Server | OPNSense Firewall | External Request | 
|---|---|---|
| User `>` Public IP:8080 | `>` NAT Rule forwards traffic to`>` | 192.168.1.100:80 | 
| User `<` Response | `<` NAT Rule translates response`<` | 192.168.1.100 | 

While configuring port forwarding on your OPNsense consider that OPNsense generates firewall rules for port forwarding automatically.  Nevertheless, it is imperative to consistently confirm that the rules are being applied accurately by navigating to **Firewall** > **Rules** > **WAN**. Secondly, you should only forward ports that are required; superfluous open ports can pose security hazards.Lastly, utilize online tools or external networks to verify the port forwarding configuration.

## How to Troubleshoot Outbound NAT Problems on OPNsense?

When your OPNsense Outbound NAT is not working as expected, you may consider the following troubleshooting steps:

1. **Check Firewall Rules** : Ensure that the firewall rules for outbound traffic are correctly configured. Verify that the rules allow traffic from the desired source and destination, and that the correct outbound NAT mode is selected.
2. **Review NAT Configuration** : Double-check the outbound NAT configuration, including the interface assignments and the source/destination rules. Ensure that the NAT mode is set to "Automatic" or "Manual" as needed.
3. **Verify Interface Configuration** : Confirm that the interface settings for the WAN and LAN interfaces are accurate, including the IP addresses, subnet masks, and default gateways.
4. **Check DNS Resolution** : Verify that DNS resolution is working correctly on the OPNsense firewall. If DNS is not resolving properly, outbound NAT may not function as expected.
