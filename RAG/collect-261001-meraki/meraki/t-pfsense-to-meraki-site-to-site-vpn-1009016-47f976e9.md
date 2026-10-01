---
id: collect-261001-meraki/meraki/t-pfsense-to-meraki-site-to-site-vpn-1009016-47f976e9
title: "t-pfsense-to-meraki-site-to-site-vpn-1009016-47f976e9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t-pfsense-to-meraki-site-to-site-vpn-1009016-47f976e9.md
source_anchor: ""
source_lines: [1, 67]
sha256: 7a7b6f04a088e787c0ba28107717511a90d58131af9642fecacfe0bb16c4a4b4
---

# t-pfsense-to-meraki-site-to-site-vpn-1009016-47f976e9

Our HQ in Topeka just received our Meraki MR18 and MX80 appliances! Thanks to Meraki for the free MR18 for attending one of their webinars! We plan to deploy more wireless appliances throughout our HQ, and then install MX80 and MR18’s in our branch locations. We have one location that will be sticking with their pfSense box, currently running on 2.1-RELEASE. Here’s how I got the Meraki security appliance to connect to the pfSense box!
Step 1: The Meraki
Browse to www.dashboard.meraki.com, log in, select the organization and then the network that holds the device you will be setting up.
Step 2: Site-to-Site VPN Settings
On the left-pane toolbar select “Security Appliance” or “Configure”> “Site-to-Site VPN”
Most of these settings will be specific to your organization’s needs. The only settings that we will need to worry about for this KB are under the “Organization-wide settings” and then “Non-Meraki VPN peers”. All settings need to be filled out completely and the Preshared Secret must be the same as what we enter under the pfSense box settings.
Non-Meraki VPN peers:
Name: TestVPN
Public IP: xx.xx.xx.xx (Remote site’s Public IP)
Private Subnets: 192.168.1.0/24 (Remote Site’s Private Subnet)
Preshared secret: Secret1
Click “Save Changes”!
Step 3: The pfSense
Log in to your remote pfSense box and on the left-pane toolbar, navigate to VPN> IPsec.
Under the “Tunnels” tab, make sure that “Enable IPsec” is checked.
Step 4: pfSense Phase 1 Settings
Add a Phase1 entry with these settings, depending on your needs. Items followed by “MUST” are required as per Meraki. Two settings were not specifically required by Meraki but would not open the tunnel with the specified settings. I have mentioned these in the Conclusion, below.
General Information
Internet Protocol: IPv4
Interface: WAN
Remote Gateway: yy.yy.yy.yy (Meraki’s Public IP)
Description: Meraki
Phase 1 Proposal
Authentication Method: Mutual PSK
Negotiation Mode: Main (MUST be Main mode. Refer to note in Conclusion regarding Aggressive mode)
My Identifier: IP Address xx.xx.xx.xx (pfSense Box’s Public IP)
Peer Identifier: IP Address yy.yy.yy.yy (Meraki’s Public IP)
Preshared Key: Secret1 (the same as under Meraki’s "Preshared Secret)
Policy Generation: Default
Proposal Checking: Default (MUST be Default. Refer to note in Conclusion regarding Proposal Checking)
Encryption Algorithm: 3DES (MUST be 3DES)
Hash Algorithm: SHA1 (MUST be SHA1)
DH Key Group: 2 (MUST be 2)
Lifetime: 28800 seconds (MUST be 28800 seconds, or 8 hours)
Advanced Options
NAT Traversal: Enable
DPD: Enabled, 10 second Delay, 5 Retries
SAVE!
Step 5: pfSense Phase 2 Settings
Disabled: UNchecked
Mode: Tunnel IPv4
Local Network: Type: Network  Address: 192.168.1.0/24 (same as the “Private Subnet” on Meraki’s settings)
Description: TestVPN
Phase 2 Proposal
Protocol: ESP
Encryption Algorithm: 3DES (Meraki claims any of 3DES, DES or AES would work)
Hash Algorithm: SHA1 (Meraki claims either MD5 or SHA1 would work)
PFS Key Group: Off (MUST be Off)
Lifetime: 28800 seconds (MUST be 28800 seconds or 8 hours)
Advanced Options
Automatically ping host: 192.168.0.1 (I just like having this set to ping the remote’s local IP)
SAVE!
Now click “Apply these configurations” box at the top-right.
Step 6: Confirm
Test a ping from the local to the remote subnet!
On the pfSense box you can check the status by going to Status> IPsec, or click the “Status of items on this page” icon at the top-right of the IPsec settings page. If the Status is not a green square with a with triangle, try clicking the “Start Tunnel” button to the right of the Status column.
On the Meraki, you cannot see a graphical indicator of the VPN working with a third party VPN device. You should see any VPN errors in the “Network-Wide> Monitor> Event Logs” or "Monitor> Event Log> deselect all filters other than VPN to search through.
I followed the Meraki KB article for “Connecting to a third-party VPN device” but I was unable to get the tunnels up with their required settings.
"Currently, MX series support for third-party VPN interoperability requires the following:
    Preshared keys (no certificates)
    LAN static routes (no routing protocol for the VPN interface)
    Phase 1 (IKE Policy): 3DES, SHA1, DH group 2, lifetime 8 hours
    Phase 2 (IPsec Rule): Any of 3DES, DES, or AES; either MD5 or SHA1; PFS disabled; lifetime 8 hours"
What I needed to configure in the pfSense box was:
Phase 1:
Negotiation Mode: Main
Proposal Checking: Default
