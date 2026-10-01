---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-troubleshooting-ipsec-vpn-routing-issues-ta-p-2-48850c67
title: "t5-fortigate-troubleshooting-tip-troubleshooting-ipsec-vpn-routing-issues-ta-p-2-48850c67"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-troubleshooting-ipsec-vpn-routing-issues-ta-p-2-48850c67.md
source_anchor: ""
source_lines: [1, 1]
sha256: da353694fad4eb9382dc02b0491c7844466d371b460b3418c693015e10f17726
---

# t5-fortigate-troubleshooting-tip-troubleshooting-ipsec-vpn-routing-issues-ta-p-2-48850c67

| Solution |   Understanding the Issue:Users operating IPsec VPNs on FortiGate might notice that while VPNs are active for a specific host, other hosts on the destination network face communication barriers. Closer inspection often reveals that traffic exits from the IPsec VPN on Azure FortiGate without receiving a corresponding response, attributed to an RPF check failure.  Steps for Diagnosis: a. Conduct a Ping Test:  Purpose: Check for live connections between source and destination. Instructions:  Use the command: ping [destination host address] Monitor for timeouts. If present, it indicates a connectivity issue. b. Run a Traceroute:  Purpose: Identify the path packets take and any potential drops. Instructions:  Input the command: traceroute [destination host address] Analyze the results. If the path is incomplete, further diagnosis is needed. c. Check Traffic Flow:  Purpose: Determine if traffic exits the Azure FortiGate via IPsec VPN and reaches the destination. Instructions:  Access VPN logs: diagnose vpn tunnel list Use packet-capturing tools to view the traffic: diagnose sniffer packet Inspect results for traffic egressing but not receiving an acknowledgment. d. Validate RPF:  Purpose: Ensure packets follow the appropriate path. Instructions:  Navigate to routing configurations on the remote end: get router info routing-table all Check for dropped packets due to RPF checks.  Resolution Procedure:  Adjust Route Priority:  Purpose: Ensure traffic passes the RPF check. Instructions:  Access the remote end's routing table: get router info routing-table all Modify the routes: config router static, then edit [route number] Alternatively  'exchange-ip-addr4' setting can be used on site-to-site tunnels, in case traffic is generated from the remote site's private tunnel IP. Adjust the priorities to ensure RPF checks are passed. Confirm traffic flows through the VPN. Notes: Always back up the configuration before making changes. Follow instructions meticulously to avoid misconfiguration.   If BGP is being used, check the routing table for the source address. If a black hole route is found for the address, restarting the routing engine may resolve the issue and make the BGP route active again. Run 'execute router restart'.  |
