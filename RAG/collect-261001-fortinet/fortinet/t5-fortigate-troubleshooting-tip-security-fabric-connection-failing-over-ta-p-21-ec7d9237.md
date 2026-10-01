---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-security-fabric-connection-failing-over-ta-p-21-ec7d9237
title: "t5-fortigate-troubleshooting-tip-security-fabric-connection-failing-over-ta-p-21-ec7d9237"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-security-fabric-connection-failing-over-ta-p-21-ec7d9237.md
source_anchor: ""
source_lines: [1, 1]
sha256: 1d5355e8188b12503d5f843913e7ee5fd5bef0c9946a22e8c8e3b8aa6556752d
---

# t5-fortigate-troubleshooting-tip-security-fabric-connection-failing-over-ta-p-21-ec7d9237

| Follow these troubleshooting steps:    Make sure the tunnel is up and running with traffic on both sides of the tunnels (Head Office and Branch).  Try to connect the upstream FortiGate with a loop-back IP. It should show a 'connecting' state, but never 'connected'.   Make sure the IP addresses are configured on the VPN interface.   Example:   The following topology shows a downstream FortiGate (Branch) connected to the root FortiGate (HQ) over IPsec VPN to join the Security Fabric:     Configure the IPsec VPN interface IP address which will be used to form the Security Fabric.    Go to Network - > Interfaces. Edit the Tunnel in question. Set the Role to LAN. Set the IP/Network Mask to 10.10.10.1/255.255.255.255. Set the Remote IP/Network Mask to 10.10.10.3/255.255.255.0.    Make sure that phase2 selectors allow the source and destination IP addresses.  If the issue still surfaces, test the fabric connectivity using a sniffer: diagnose sniffer packet any “port 8013” 6 0 l   Note:  By default,  Port 8013 (set upstream-port 8013) is the port used for security fabric syncing. There should be no doubt on adding remote IP even in the case of dial-up IPsec VPN. Configuring IP addresses is still possible on VPN interfaces for dial-up VPN, similar to ADVPN, which is a dial-up VPN too. This is also valid for a site-to-site IPsec VPN.   Related article:  Troubleshooting Tip: Troubleshooting Security Fabric Issues |
