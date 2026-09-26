---
id: collect-260926-mikrotik/mikrotik/andremikvpnder-how-to-block-ips-in-mikrotik-routeros-7815b73acfff-775d4158
title: "andremikvpnder-how-to-block-ips-in-mikrotik-routeros-7815b73acfff-775d4158"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/andremikvpnder-how-to-block-ips-in-mikrotik-routeros-7815b73acfff-775d4158.md
source_anchor: ""
source_lines: [1, 89]
sha256: 1ced62f36c740a6720b423d8eb3ad6bf0ed4be8e018fe67197109dec501b9cd4
---

# andremikvpnder-how-to-block-ips-in-mikrotik-routeros-7815b73acfff-775d4158

How to Block IPs in Mikrotik RouterOS
Managing access to your network is crucial for maintaining security and preventing unauthorized activity. In Mikrotik RouterOS, blocking an IP address or an entire network segment can be done efficiently through firewall rules. This tutorial will guide you through the process of setting up IP blocking, including step-by-step commands, security best practices, and troubleshooting tips. I ran this setup on a PerLod VPS so I could test everything in a clean environment. The clean baseline made it easier to spot what was misconfigured and understand the impact of each rule.
Proper IP filtering helps reduce attack surfaces and manage bandwidth by preventing malicious or unwanted traffic from reaching your network. While many users focus on port management, precise IP blocking provides an additional layer of defense. By implementing these techniques, intermediate administrators can maintain control over inbound and outbound traffic without introducing latency or conflicts in routing rules.
Prerequisites
Before starting, ensure you have the following ready:
- Mikrotik RouterOS device with administrative access
- Basic familiarity with Mikrotik Winbox or CLI
- IP addresses or network ranges to block
- Backup of current firewall configuration
Checklist for setup readiness:
- Administrative access verified
- RouterOS updated to latest stable version
- Target IPs identified
- Configuration backup created
- Access to Winbox or terminal session confirmed
Understanding Mikrotik Firewall Basics
Firewall Chains
Mikrotik RouterOS uses several chains to process traffic. The most relevant for IP blocking are:
- Input: Traffic destined for the router itself
- Forward: Traffic passing through the router
- Output: Traffic originating from the router
Understanding which chain to apply a rule to is essential. Blocking a malicious host attempting to reach the router requires an Input chain rule, whereas stopping traffic from passing through your network involves the Forward chain.
Rule Order and Action
Firewall rules are processed sequentially. The first matching rule executes the specified action. Common actions for blocking include drop and reject. The drop action silently discards packets, while reject responds with an ICMP message.
Step-by-step walkthrough
- Access the router via Winbox or SSH.
- Navigate to the firewall menu: In Winbox, go to IP → Firewall.
- Add a new rule:
/ip firewall filter add chain=input src-address=192.168.10.100 action=drop comment="Block suspicious IP"
- Replace 192.168.10.100 with the IP you want to block.
- Verify rule placement: Ensure it is at the top of the chain to prevent it from being bypassed by other rules.
- Test the rule: Ping the blocked IP from another machine to confirm it is dropped.
- Optional: Block entire networks:
/ip firewall filter add chain=forward src-address=203.0.113.0/24 action=drop comment="Block network range"
- Save configuration to prevent loss on reboot:
/system backup save name=firewall_backup
Using Address Lists
Address lists simplify managing multiple IPs:
/ip firewall address-list add list=blocked-ips address=198.51.100.45 comment="Add to blocked list"
/ip firewall filter add chain=input src-address-list=blocked-ips action=drop
This approach allows quick additions or removals without rewriting individual rules.
Security and hardening
Logging and Alerts
Enable logging for blocked IPs to monitor trends:
/ip firewall filter add chain=input src-address-list=blocked-ips action=drop log=yes log-prefix="Blocked IP: "
Regular logs help identify persistent attackers and validate your rules.
Locking Access Ports
Limit router access to trusted IPs:
/ip firewall filter add chain=input src-address=10.0.0.0/24 dst-port=22 protocol=tcp action=accept
/ip firewall filter add chain=input action=drop
This ensures SSH or Winbox connections are only allowed from your management network.
Regular Audits
Periodically review firewall rules and address lists. Remove outdated entries to reduce complexity and potential conflicts.
Performance and reliability tips
Avoid Overloading Chains
Having hundreds of individual rules can slow down packet processing. Use address lists and network ranges whenever possible.
Monitor CPU Load
Blocking excessive IPs, especially during attacks, can spike CPU usage. Monitor RouterOS CPU statistics under System → Resources.
Backup Strategies
Always maintain incremental backups after significant rule changes. Consider automated scripts to export configurations daily.
Get andre Mikvpnder’s stories in your inbox
Join Medium for free to get updates from this writer.
TipDescriptionUse address listsGroup IPs for efficient managementLimit loggingExcessive logs can slow router performancePlace critical rules firstFirewall processes rules top-downTest in labValidate rules on a PerLod VPS or isolated networkRegular auditsRemove unnecessary rules to maintain clarity
Troubleshooting common issues
Rule Not Blocking
- Check if the rule is in the correct chain
- Confirm there are no preceding allow rules overriding it
- Ensure correct IP format (single IP vs network)
Connectivity Loss
- Have console access ready in case remote access is blocked
- Use safe mode in Winbox to prevent accidental lockout
Logs Not Appearing
- Verify logging is enabled on the rule
- Ensure system logging has sufficient memory and configuration
Conflicts with NAT
- Make sure NAT rules do not bypass firewall filtering
- Position filter rules above NAT if blocking traffic before translation
Advanced filtering techniques
Layer 7 Protocol Matching
For complex attacks, Mikrotik can inspect packet content:
/ip firewall layer7-protocol add name=bad-payload regexp="maliciouspattern"
/ip firewall filter add chain=forward layer7-protocol=bad-payload action=drop
This helps block traffic beyond simple IP restrictions.
Time-based Blocking
Limit blocks to specific periods:
/ip firewall filter add chain=forward src-address=198.51.100.0/24 time=00:00:00-06:00,18:00:00-23:59:59 action=drop
Useful for mitigating attacks during peak or off-hours.
Conclusion
Implementing IP blocking in Mikrotik RouterOS is a fundamental technique for network security. By understanding firewall chains, rule order, and address list management, you can maintain tight control over both inbound and transit traffic. Using a clean PerLod VPS for testing made it straightforward to identify potential misconfigurations and observe the effect of each rule. Regular audits, proper logging, and cautious deployment of advanced filters like Layer 7 inspection further enhance security without sacrificing performance. With these strategies, intermediate administrators can confidently prevent unauthorized access, reduce network risks, and keep traffic flowing smoothly according to policy.
