---
id: collect-261001-general-networking/general-networking/morningreis-eeda36e8bb07dcb750d77e9a744776e8-067af069-2
title: "Key for OPNsense Demo"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/morningreis-eeda36e8bb07dcb750d77e9a744776e8-067af069.md
source_anchor: ""
source_lines: [143, 161]
sha256: 382b515b3e2cd9c8d575595a01c7b68575a711845ad33fdb1dda66371e625806
---

# Key for OPNsense Demo

- Move this rule above all other rules
- Save and Apply Changes
- Navigate to Firewall > Rules > Floating
- Add a new rule
- Make the following changes:
Action:                Block
Quick:                 Checked
Interface:             WAN
Direction:             out
Click Advanced Options Show/Hide
Match local tag:         NO_WAN_EGRESS
- Move this rule above all other rules
- Save and Apply Changes
This block rule only needs to be made once, regardless of if you have multiple Wireguard tunnels. This is what is serving as the killswitch
You can repeat these steps to add additional tunnels, however I had issue getting more than 3 tunnels to work. I believe there may be a limit on Proton's end. On the WireGuard configuration page, delete any stored configurations that you are not using.
If you have more than one tunnel and want to disable one, go to Navigate to VPN > WireGuard > Endpoints and check the tunnel you want to disable. If you disable it from the Local tab, you will also remove the interface.
If for any reason a Wireguard tunnel drops out or fails to connect, the interface and gateway you created will still be in place, and the killswitch will prevent any traffic leaking out. However if your gateway is disabled for any reason, the default behavior will be to use your regular WAN gateway, meaning the VPN will not be used at all. There is no reason for your gateway to go down however even if the tunnel is down.
Lastly, for DNS Leak protection, you should ensure that your DNS resolver (most likely your OPNsense machine) is included under an Alias to be routed through one of your Wireguard connections. That will force DNS requests to go through the VPN, but past that you will need to configure DNS over TLS or DNS over HTTPS using Unbound DNS, which is outside the scope of this guide.
Thanks for this tutorial it helped me!
