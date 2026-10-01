---
id: collect-261001-cisco/cisco/hc-en-us-articles-17198673056151-unifi-identity-enterprise-troubleshoot-vpn-issu-7eadb0c5
title: "hc-en-us-articles-17198673056151-unifi-identity-enterprise-troubleshoot-vpn-issu-7eadb0c5"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "ethernet"]
source: docs/RAG/collect-261001-cisco/hc-en-us-articles-17198673056151-unifi-identity-enterprise-troubleshoot-vpn-issu-7eadb0c5.md
source_anchor: ""
source_lines: [1, 73]
sha256: e3516e0e65eedd07587ed3a93f5f0396ad8771221db77706f115a66a693e7ec8
---

# hc-en-us-articles-17198673056151-unifi-identity-enterprise-troubleshoot-vpn-issu-7eadb0c5

Troubleshooting UID Enterprise One-Click VPN
Important
A VPN should always use a private LAN address to avoid routing conflicts, NAT issues, and security risks. Private IP range includes:
- 10.0.0.0/8
- 172.16.0.0/12
- 192.168.0.0/16
Common One-Click VPN Issues
Issues
- One-Click VPN is frequently disconnected.
- The status shows VPN Connected but you still cannot connect to the internet.
Resolutions
Check the following and try connecting to the VPN again:
- Ensure One-Click VPN is set up correctly.
- Ensure port 10118 is enabled.
- Ensure the console's primary WAN IP address is not within the console's private IP range.
- Ensure the console's public IP address is not within the CGNAT IP range (100.64.0.0 to 100.127.255.255).
- If a public IP address is assigned to your UniFi Console, port forwarding does not need to be configured manually.
- If your UniFi Console does not have a public IP address, but the router or gateway connected to the ISP does, configure port forwarding to port 10118 (OpenVPN) or 51820 (WireGuard) on that router.
- Ensure the One-Click VPN status is "Enabled" in UID Enterprise Manager > Services > One-Click VPN.
- Make sure a public IP address is configured. Please refer to VPN Connection Issues or Frequent VPN Disconnections or Timeouts/Ensure a Public IP Is Configured below for details.
- Ensure your UID Enterprise Agent is online.
- If your issue persists, please refer to VPN Connection Issues or Frequent VPN Disconnections/Timeouts below for more information.
VPN Connection Issues or Frequent VPN Disconnections or Timeouts
- For users: If you are unable to connect to One-Click VPN, please contact your UID Enterprise administrator and then submit your feedback on your UID Enterprise mobile app.
- For administrators: Go to your UID Enterprise Manager > Services > One-Click VPN > VPN to modify the VPN settings.
If you've followed the steps above but are still experiencing connection issues, refer to the resolutions below in order.
A. Ensure a Public IP Is Configured
If your UniFi Console does not have a public IP address, you will need to configure port forwarding. Multi-level port forwarding is required for consoles with a public IP address that has multi-level routes. You can use the following methods to check your console's public IP settings:
- 
Method 1: Check in your OS Settings.
  - Go to UniFi OS > Settings > General.
  - Check if the WAN IP is a public IP.
- 
Method 2: Check via SSH.
  - 
Enter the following traceroute command:
  - Check if the first router address is a public IP.
  - 
Enter this command to check the VPN operating environment.
- 
B. Ensure the One-Click VPN Settings Are Correct
- Go to Services and click One-Click VPN.
- Do either of the following:
  - If the workspace has one site: Go to VPN and click the One-Click VPN.
  - If the workspace has multiple sites: Go to Sites, click a site, go to VPN, and click a One-Click VPN.
- Make sure the settings are correct, especially the VPN port and IP address.
C. Enter the Telnet Command to Check the VPN Port Connectivity
Note: This method is only available when the VPN uses TCP protocol. The IP and port are automatically filled in the VPN settings.
- If the telnet IP port can be connected, a "Connected to" status will be shown.
- If the telnet IP port cannot be connected, an "Unable to connect" status will be shown.
D. Sign In to UniFi OS via SSH
Note: Do not close the SSH terminal because you will need to return to it later.
- Enter the following command:
ssh root@UDM_IP
tcpdump -i eth8 dst port 10118
-  
  - eth8 in the command line represents the console WAN port with an Ethernet connection and is calculated as eth(n-1), where n indicates the port number.
  - Example: eth8 on a UDM Pro indicates that Port Number 9 on the console is connected via Ethernet cable (eth(9-1) = eth8 ). For a UDM, it would be eth(5-1) or eth(4).
Issue Related to Single WAN Support with Multiple Public IP Addresses
One-Click VPN only supports multiple public IP addresses on a single WAN when using the TCP network protocol. When this protocol is in use, all console IPs are available.
If you use the UDP protocol to set up a One-Click VPN, you will only be able to use the primary IP for VPN connections. This is because UDP is a connectionless protocol. When client packets sent to the server reach the network layer, the server is likely to forget the client’s requested destination IP. As a result, the server chooses the primary IP as the source address when sending packets back to the client. Therefore, you cannot use multiple public IPs while using the UDP protocol.
Issue Related to Dual WAN Support
One-Click VPN doesn't support multiple WAN configurations over UDP because the protocol cannot remember the source IP address when packets return to the server.
To solve this issue, do either of the following:
- Change Load Balancing to "Failover Only".
  - Go to your UniFi Network application.
  - Go to Settings > Internet > Load Balancing.
  - Select "Failover Only".
- Switch the protocol from UDP to TCP. Note: This method is applicable to OpenVPN only. WireGuard VPN does not support TCP.
  - Sign in to your UID Enterprise Manager.
  - Go to Services > One-Click VPN > VPN.
  - Select a VPN and scroll down to Advanced > Protocol.
  - Change "UDP" to "TCP".
