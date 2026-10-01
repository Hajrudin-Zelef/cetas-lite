---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-cannot-rdp-into-pc-when-connected-with-ssl-vpn-145015-40cddbac
title: "fortigate-3-technical-tip-cannot-rdp-into-pc-when-connected-with-ssl-vpn-145015-40cddbac"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-cannot-rdp-into-pc-when-connected-with-ssl-vpn-145015-40cddbac.md
source_anchor: ""
source_lines: [1, 25]
sha256: 7aa0ddca6e96ec9f1b98e8bae12bcd2a039acca51b7210b49d757c424eb8a191
---

# fortigate-3-technical-tip-cannot-rdp-into-pc-when-connected-with-ssl-vpn-145015-40cddbac

Technical Tip: Cannot RDP into PC when connected with SSL VPN
Description
This article describes that if the user cannot RDP into the PC when connected with SSL VPN, but RDP when it is on the same network, and provides troubleshooting steps for this issue.
Scope
FortiGate.
Solution
- Check the SSL VPN setting. Make sure the user is in SSL VPN setting -> Authentication & portal mapping:
- If it has a full access portal assigned, check in the portal if split tunneling is enabled.
- Make sure the SSL VPN to LAN policy has a subnet in which the PC resides as the destination with service ALL or at least RDP.
- If all the configurations are as stated, try to run the following command:
diagnose debug disable
diagnose debub flow filter saddr x.x.x.x <----- IP user is getting when connected with SSL VPN.
diagnose debug flow filter daddr x.x.x.x <-----PC IP which user is trying to RDP in.
diagnose debug flow show function-name en
diagnose debug flow trace start 999
diagnose debug en
- If the traffic is being accepted by the SSL VPN to the LAN policy but still not able to RDP, check below, try to run the command:
diagnose sniffer packet any ‘host x.x.x.x and host y.y.y.y’ 4 0 l
Or
diagnose sniffer packet any ‘host x.x.x.x and port 3389’ 4 0 l
- Here, x.x.x.x is the IP that the user gets when connected with VPN, y.y.y.y ,is the IP of the PC that is RDP into.
- Check if there is a reply from the PC.
- Check Logs & Report -> Forward traffic logs and apply a filter with source and destination addresses.
- If there is 0 Byte in received bytes, check if the Windows firewall is enabled on the PC, disable it, and try again. Make sure there is no other firewall other than FortiGate that can block traffic. Disable if there is any other firewall and try again.
Contact TAC if there is still an issue.
