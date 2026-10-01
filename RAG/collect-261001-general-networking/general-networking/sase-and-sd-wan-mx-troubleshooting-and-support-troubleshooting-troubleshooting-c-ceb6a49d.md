---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-troubleshooting-and-support-troubleshooting-troubleshooting-c-ceb6a49d
title: "sase-and-sd-wan-mx-troubleshooting-and-support-troubleshooting-troubleshooting-c-ceb6a49d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-troubleshooting-and-support-troubleshooting-troubleshooting-c-ceb6a49d.md
source_anchor: ""
source_lines: [1, 43]
sha256: dc881c9a4514d0b699f6333213777c3e0c67ba64da701eae8a2a09a96331288f
---

# sase-and-sd-wan-mx-troubleshooting-and-support-troubleshooting-troubleshooting-c-ceb6a49d

Troubleshooting Client VPN When all Devices Cannot Connect
Overview
This article provides troubleshooting steps for issues where no client VPN users can connect. If some users can connect, refer to the Unable to Connect to Client VPN from Some Devices.
Troubleshooting MX availability issues
Troubleshooting steps
- Verify that your MX is online and accessible over the internet.
- In the Meraki dashboard, navigate to Security & SD-WAN > Monitor > Appliance status.
- Select the Tools tab at the top of the Appliance status page.
- Select the Ping appliance button.
- Confirm that the MX successfully returns a ping.
Troubleshooting incorrect MX IP address
The MX is offline or unreachable over the internet, preventing all VPN clients from connecting.
Possible causes
When using two uplink connections, the MX IP address may change when the uplink fails over from primary to secondary. VPN connections configured to use the primary MX IP address would no longer work.
Troubleshooting steps
- Verify that the client VPN is configured to connect to the MX using the correct IP address.
- In the Meraki dashboard, navigate to Security & SD-WAN > Monitor > Appliance status to confirm the current MX IP address.
Troubleshooting Dynamic DNS (DDNS)
Troubleshooting upstream NAT/firewall issue on the MX
Troubleshooting steps
- If the MX is behind a NAT device (for example, an upstream router or ISP modem), the MX uplink IP may have a private IP in the 172.16.X.X, 192.168.X.X, or 10.X.X.X subnet range. Verify that UDP traffic on ports 500 and 4500 is being forwarded to the private uplink IP address of the MX.
- Verify that no firewall is blocking UDP traffic on ports 500 or 4500.
- Take a packet capture on the WAN interface of the MX and confirm that traffic from the public IP of the VPN client on UDP ports 500 and 4500 is reaching the MX.
Refer to the Troubleshooting Client VPN with Packet Captures article for more information.
Troubleshooting authentication issue
VPN clients receive an authentication error when attempting to connect.
Troubleshooting steps
If receiving authentication errors:
- 
    Verify that the VPN client is configured with the correct username, password, and shared secret.
- 
    If the issue persists, try a different authentication method, such as Meraki Cloud Authentication, RADIUS, or Active Directory.
Troubleshooting shared secret mismatch
The VPN tunnel fails to establish because the shared secret configured on the VPN client does not match the shared secret on the MX.
Troubleshooting steps
VPNs require the shared secret to match on the VPN server and client before tunnels can be established.
To view the shared secret:
- In the Meraki dashboard, navigate to Security & SD-WAN > Configure > Client VPN.
- Select the IPSec Settings tab and scroll down to Shared secret.
- Select Show secret and confirm the shared secret matches the pre-shared key used in the VPN client configuration.
- If the issue persists, try changing the shared secret. As a best practice, the shared secret should not contain any special characters at the beginning or end.
Encryption method
Client VPN uses the L2TP/IP protocol, with 3DES encryption and SHA1 hashing.
