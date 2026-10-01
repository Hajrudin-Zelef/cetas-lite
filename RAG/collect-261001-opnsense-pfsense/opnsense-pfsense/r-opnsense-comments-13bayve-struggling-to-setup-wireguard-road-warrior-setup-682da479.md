---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-13bayve-struggling-to-setup-wireguard-road-warrior-setup-682da479
title: "r-opnsense-comments-13bayve-struggling-to-setup-wireguard-road-warrior-setup-682da479"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-13bayve-struggling-to-setup-wireguard-road-warrior-setup-682da479.md
source_anchor: ""
source_lines: [1, 75]
sha256: 0cec0878062fe0e5417337e23c14348cefd511bb2cbce5e2dbcaf6178ec45053
---

# r-opnsense-comments-13bayve-struggling-to-setup-wireguard-road-warrior-setup-682da479

Anyone here who's setup remote access via WireGuard on OPNsense?

It's my first time setting it up and I'm struggling to access the network from a outside connection..

I've followed the official guide

And this tutorial as a visual reference. https://youtu.be/b58PpuIsQ3A

Here's the steps I've followed, in great detail.

I've installed the WireGuard Plugin, refreshed the page, went to the VPN 》WireGuard section.

I enabled the service, setup a local configuration.

I ticked, enabled.

Name: wg1

I left public and private key empty, since it's automatically generated.

Listen Port: 51820

Tunnel Address: 10.0.0.1/24 (different compared to the lan and Wan subnets.)

DNS: Local Servers one and two.

--

I then rebooted the OPNsense instance, and assigned my wg1 interface to an OPNsense instance.

In the wg1 interface settings I enabled packet logging and the interface, leaving the rest at the default settings.

I then passed a firewall rule for the Wan interface, allowing all WireGuard traffic inbound.

Action: Pass

Interface: WAN

Direction: in

TCP/IP Version: IPv4

Protocol: UDP

Source: Any

Destination: WAN Address

Destination Port Range Other: 51820

Added packet logging and saved, then applied.

I then moved on to setup Client Configuration.

[Interface] PrivateKey = "Clients Private key" Address = 10.0.0.2/32 DNS = the local DNS Servers I used earlier, separated with a comma.

[Peer] PublicKey = "The public key from the local configuration earlier." AllowedIPs = 0.0.0.0/0 Endpoint = "My Virtual IP WAN"

I then setup a endpoint device in the OPNsense WireGuard Panel.

Name: Client1

Public Key: "Public Key from Client."

AllowedIPs: 10.0.0.2/32

Kept everything else empty or default.

I then added the endpoint as a peer in the local configuration, and restarted the WireGuard service.

Now it's time to test it, I pushed my wg configuration up on my client, and checked for handshakes on the OPNsense WireGuard Pannel, I can see that a tunnel has successfully been connected over the local network.

I then switched my client to a separate cellular network, and attempted to connect over the tunnel.. I can't ping anything in the local network..

Any help is greatly appreciated, Thanks.
