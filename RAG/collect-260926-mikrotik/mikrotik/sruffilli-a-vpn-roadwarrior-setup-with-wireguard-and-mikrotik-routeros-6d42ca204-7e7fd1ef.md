---
id: collect-260926-mikrotik/mikrotik/sruffilli-a-vpn-roadwarrior-setup-with-wireguard-and-mikrotik-routeros-6d42ca204-7e7fd1ef
title: "Create the wireguard interface, and generate the pub/pri keys"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/sruffilli-a-vpn-roadwarrior-setup-with-wireguard-and-mikrotik-routeros-6d42ca204-7e7fd1ef.md
source_anchor: ""
source_lines: [1, 44]
sha256: 701fef8429da52ea531cd1e4c9d6e283a219bd72532ce0d2657be85cf4a8dac2
---

# Create the wireguard interface, and generate the pub/pri keys

A VPN roadwarrior setup with WireGuard and Mikrotik RouterOS
Disclaimer: I’ve just put my hands over an hAP ac², my first piece of Mikrotik equipment. I’m very new to RouterOS so take this article as my own notes rather than a prescriptive recipe — comments welcome!
RouterOS 7 (currently available as a Release Candidate) introduced support for WireGuard, the VPN tech that aims to be “faster, simpler, leaner” than IPSec, and “considerably more performant than OpenVPN”.
This brief article explains how I have configured my hAP ac² for a roadwarrior scenario — that is, a VPN gateway that accepts peers connecting from non-static IP Addresses.
This is a simplified diagram of my current networking setup:
An ISP-provided router terminates the (PPPoA) DSL connection, and NATs 1:1 its public interface (1.2.3.4) to the WAN interface of the hAP (192.168.0.2), which through the LAN interface (192.168.1.1) masquerades all traffic going towards WAN.
Creating and configuring the WireGuard interface
We’re going to create a network interface for WireGuard, which will be assigned the IP 192.168.98.1, and we’ll dedicate 192.168.98.0/24 for the remote clients.
# Create the wireguard interface, and generate the pub/pri keys
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1# Print the newly created interface - mark the public-key for later
/interface wireguard print
# Returns, e.g.
# 0 R name="wireguard1" mtu=1420 listen-port=13231 private-key="example-mikrotik-private-key" public-key="example-mikrotik-public-key"# Allocate an IP address to the wireguard interface.
# This will also automatically create a route for 192.168.98.0/24
/ip address
add address=192.168.98.1/24 interface=wireguard1 network=192.168.98.0# Allow incoming traffic to the wireguard service
/ip firewall filter
add action=accept chain=input dst-port=13231 protocol=udp
Client configuration
To have your roadwarriors connecting to WireGuard, you’ll have to generate a configuration file (including a pub/pri key pair) for each client. You can read the WireGuard docs, use a tool such as WireGuard Config Generator (which claims to be client-side only) or your client UI (e.g. the official Android client can import or generate the required config). You’re going to need the generated public key (let’s call it example-client1-public-key) for a later setup stage.
Get Simone Ruffilli’s stories in your inbox
Join Medium for free to get updates from this writer.
Let’s take a look at a sample configuration:
[Interface]
Address = 192.168.98.2/32
DNS = 8.8.8.8
PrivateKey = example-client1-private-key[Peer]
AllowedIPs = 0.0.0.0/0
Endpoint = myrouterpublicip.example.com:13231
PublicKey = example-mikrotik-public-key
This configuration routes all traffic to the VPN gateway (including internet traffic), which might or might not be the desired scenario. In case you want to implement “split tunneling” instead and only route private IPs to the VPN, the configuration would change as follows (notice the change in the “AllowedIPs” bit).
[Interface]
Address = 192.168.98.2/32
DNS = 8.8.8.8
PrivateKey = example-client1-private-key[Peer]
AllowedIPs = 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
Endpoint = myrouterpublicip.example.com:13231
PublicKey = example-mikrotik-public-key
One more thing
One last bit of configuration is required on the Mikrotik side — that is, adding and configuring a (or as many as you have created!) WireGuard peer.
/interface wireguard peers
add allowed-address=192.168.98.2/32 interface=wireguard1 public-key="example-client1-public-key"
That should be all! Your roadwarrior should be able to ping (and access) the local network, and potentially (according to the AllowedIPs configuration) egress from your home/office.
