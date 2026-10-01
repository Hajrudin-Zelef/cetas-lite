---
id: collect-261001-general-networking/general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016-3
title: "docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016.md
source_anchor: ""
source_lines: [359, 450]
sha256: ac535024dc575253fe8dfb040808b141602a5f247b9ec127f952179e1c29b575
---

# docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016

| Protocol | ESP | Choose ESP for encryption | 
| Encryption algorithms | AES256GCM16 | For the sample we use AES256GCM16 | 
| Hash algortihms | SHA512 | Choose a strong hash like SHA512 | 
| PFS Key group | 14 (2048 bit) | Not required but enhanced security | 
| Lifetime | 3600 sec |  | 
Table 21. Phase 2 proposal (SA/Key Exchange) Phase-2 options on Site-B
You may easily configure IPSec Phase-2 on Site-B by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-B OPNsense web UI.
- 
Click add phase 2 entry button with + in the Commands column of the recently added phase 1 entry.
- 
Add a Description, such as Local LAN Site A .
- 
Set Address option for Remote Network, such as 10.10.10.0/24 .Figure 17. General information Phase-2 on Site-B
- 
Select Encryption algorithms, such as AES256GCM16 .
- 
Select Hash algorithms, such as SHA512 .
- 
Select PFS key group, such as 14 (2048) bits .
- 
Set Lifetime, such as 3600 .
- 
You may leave other options as default.
- 
Click Save.
- 
Click the checkbox at the beginning of the Phase 1 pane to view the Phase 2 settings.
7. Enabling IPsec on Site-B
You may quickly enable IPsec service on SIte-B by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-B OPNsense web UI.
- 
Check Enable IPsec option at the bottom of the page.
- 
Click Apply Changes button at the top right corner of the page to activate the IPsec tunnel settings. Figure 18. Enabling IPsec on Site-B
8. Adding Firewall Rule for LAN Access on Both Site
You may easily add firewall rules on OPNsense firewalls located in Site A and Site B by following the next steps to allow IPsec tunnels to access LAN:
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | IPsec | 
| Protocol | any | 
| Source | any | 
| Source Port | any | 
| Destination | LAN net | 
| Destination Port | any | 
| Category | IPsec Tunnel | 
| Description | Allow IPsec Tunnel traffic to LAN | 
Table 22. Firewall rule settings for LAN access
- 
Navigate to the IPsec interface on the Firewall Rules.
- 
Select Pass for the allow rule.
- 
Select any as the Protocol.
- 
Select any as the Source.
- 
Select any as the Source port.
- 
Select any as Type.
- 
Select LAN net as the destination.
- 
Set Category to IPsec Tunnel .
- 
Set Description to Allow IPsec Tunnel traffic to LAN .
- 
Enable Log packets that are handled by this rule option.
- 
Click Save.
- 
Click Apply Changes. Figure 19. Firewall rule for LAN access on IPsec interface
9. Viewing IPsec Tunnel Status
Both networks should now be routed through the tunnel. To view the current IPsec VPN tunnel status, you may follow navigate to VPN → IPsec → Status Overview on OPNsense web UI.
Figure 20. Viewing IPsec Tunnel Status
Attempt a service restart on both endpoints if the tunnel fails to appear.
How to Troubleshoot IPsec S2S Tunnel Problems on OPNsense
You can navigate through the configured tunnels using the VPN → IPsec → Status Overview menu in order to monitor the connected tunnels.
Additionally, it is possible to gain insight into the registered policies by navigating to the VPN → IPsec → Security Association Database; when NAT is in place, the additional SPD entries should also be visible here.
When attempting to diagnose issues with your OPNsense firewall, you will almost certainly be required to examine the records that are accessible on your system. The user interface of OPNsense organizes log files in accordance with the configurations of the component to which they pertain. The location of the log files is specified in the "Log file" menu.
The most common IPsec site-to-site tunnel issues and their solutions explained below:
- 
Phase 1 does not come up: That issue is quite challenging. Before proceeding, verify that the WAN interface is permitted on the appropriate ports and protocols (ESP, UDP 500, and UDP 4500) via the firewall. Examine your IPSec log to determine whether this is a potential cause. Inequality in settings is a prevalent concern. Both endpoints must employ the identical PSK and encryption protocol.
- 
Phase 1 is operational, but phase 2 tunnels are not connected: Have the proper local and remote networks been configured? It is a frequent error to enter the IP address of the remote host rather than the network's x.x.x.0 suffix.Inequality in settings is a prevalent concern. Both endpoints must employ the identical encryption protocol.
How to Tune IPsec Tunnel on OPNsense?
Enabling multithreaded crypto mode on IPsec is advantageous, depending on the burden (single flow or multiple IPsec flows). This mode distributes cryptographic packets across multiple processors, which is particularly beneficial when only one tunnel is in use.
To enable multithreaded crypto mode on IPsec, you may add or modify the following tunable by navigating to System → Settings → Tunables on OPNsense UI.
net.inet.ipsec.async_crypto = 1
