---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/home2021-src-posts-opnsense-multiwan-index-mdx-at-ff92b52695bf840040f2954adbc6e02d04ba5aad-2
title: "home2021-src-posts-opnsense-multiwan-index-mdx-at-ff92b52695bf840040f2954adbc6e02d04ba5aad-ndom91-ho"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["kill switch"]
source: docs/RAG/collect-261001-opnsense-pfsense/home2021-src-posts-opnsense-multiwan-index-mdx-at-ff92b52695bf840040f2954adbc6e02d04ba5aad-ndom91-ho.md
source_anchor: ""
source_lines: [129, 178]
sha256: c80a62b7ffd4a4ac0dadd70038480fccbc0389698eca29023087c937e61cabdd
---

# home2021-src-posts-opnsense-multiwan-index-mdx-at-ff92b52695bf840040f2954adbc6e02d04ba5aad-ndom91-ho

| Field | Description | 
|---|---|
| Action | Pass | 
| Quick | Unchecked | 
| Interface | Do not select any | 
| Direction | out | 
| TCP/IP Version | IPv4 | 
| Protocol | any | 
| Source / Invert | Unchecked | 
| Source | Select the interface **address** for your Tailscale interface (eg`TSCL address` ) | 
| Destination / Invert | Checked | 
| Destination | Select the interface **network** for your Tailscale interface (eg`TSCL network` ) | 
| Destination port range | any | 
| Description | Add one if you wish to | 
| Gateway | Select the gateway you created above (eg `tailscale_gw` ) | 
| allow options | Checked | 

1. Click **Save** and**Apply**

Next, we'll need to setup a NAT rule to map the internal LAN host addresses (for our chosen target LAN hosts) to the Tailscale interface address and vice-versa for traffic coming and going.

1. Go to **Firewall** ->**NAT** ->**Outbound**
2. If not yet enabled, select **Hybrid outbound NAT rule generation** and click**Save** and**Apply** to apply the hybrid rule setting.
3. Click **Add** to add a new NAT rule
4. Configure as follows

| Field | Description | 
|---|---|
| Interface | Select your Tailscale interface (i.e. `TSCL` ) | 
| TCP/IP Version | IPv4 | 
| Protocol | any | 
| Source invert | Unchecked | 
| Source address | Select the Alias we created for the hosts intended to use the tunnel (eg `tailscale_target_hosts` ) | 
| Source port | any | 
| Destination invert | Unchecked | 
| Destination address | any | 
| Destination port | any | 
| Translation / target | Interface address | 
| Description | Add one if you wish to | 

1. Click **Save** and**Apply**

After applying that last NAT rule, we should successfully have internet connectivity again from the selected hosts, only this time their source IP, from the point of internet hosts, is your Tailscale exit-node! A simple way to test this is to use the `wtfismyip.com` service. You can simply `curl` their `/json` endpoint to get a quick summary of your IP info as seen by that host.

`curl wtfismyip.com/json`
You can toggle this functionality on/off by enabling/disabling the firewall alias `tailscale_target_hosts`, this allows you to easily turn the selective routing on/off whenever you need it as well as control which hosts these special rules should apply to.

This post is based off of the guide in the OPNsense documentation for "Selective Routing to External VPN Endpoints" with modifications for Tailscale and multi-wan. Check out that guide for additional options like "adding a Kill Switch", to disable network connectivity for the targeted hosts if the Tailscale gateway is offline instead of falling back to your default gateway. Or "adding IPv6 support", or the aforementioned "DNS Leak prevention".

If you find any errors, please don't hesitate to open a PR at ndom91/home2021!
