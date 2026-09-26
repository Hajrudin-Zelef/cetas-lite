---
id: collect-260926-mikrotik/mikrotik/blog-mikrotik-ipsec-site-to-site-vpn-a-comprehensive-guide-1764809120-5cc8f45b
title: "MikroTik IPSec Site-to-Site VPN: A Comprehensive Guide"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/ipsec/blog-mikrotik-ipsec-site-to-site-vpn-a-comprehensive-guide-1764809120-5cc8f45b.md
source_anchor: ""
source_lines: [1, 75]
sha256: 6206bf135360a4f54faac62d774f75aa0f1e4743d5778beeff07dd3f2be6fb34
---

# MikroTik IPSec Site-to-Site VPN: A Comprehensive Guide

*Source : https://pocket-or-print-pop.com/blog/mikrotik-ipsec-site-to-site-vpn-a-comprehensive-guide-1764809120*
*(Note : l'extraction de la page source s'interrompt en fin d'article — la dernière section est tronquée.)*

Hey guys! Ever wondered how to create a secure, reliable connection between two networks using MikroTik routers? Well, you're in the right place! This guide is all about setting up a **MikroTik IPSec site-to-site VPN**, and we'll break it down step-by-step. Whether you're a seasoned network admin or just starting out, this article will walk you through everything you need to know to establish a secure tunnel between your networks.

## Understanding the Basics of IPSec and Site-to-Site VPNs

**IPSec** (Internet Protocol Security) is a suite of protocols that provides secure communication over an IP network. It encrypts and authenticates the data packets being sent between two points, protecting data from eavesdropping and tampering as it travels across the internet or any other public network.

A **site-to-site VPN** connects two or more networks together, creating a secure tunnel between two physical locations so devices on those networks communicate as if on the same local network. Combined with IPSec, traffic through the tunnel is encrypted — ideal for sharing sensitive data or accessing internal applications remotely.

## Planning Your MikroTik IPSec VPN Configuration

- **Identify your networks:** e.g., main office `192.168.1.0/24`, branch office `192.168.2.0/24`. They must not overlap.
- **Router IPs:** each router needs a public IP to establish the tunnel; note internal interface addresses too.
- **IPSec parameters:** encryption algorithm (AES, 3DES), hash algorithm (SHA1, MD5, SHA256), Diffie-Hellman group (2, 5, 14), lifetimes for Phase 1 and Phase 2.
- **Topology:** hub-and-spoke vs full mesh — impacts routing. Also consider **MTU** (VPN overhead can cause fragmentation; typically reduce MTU on the tunnel interface).

## Step-by-Step MikroTik IPSec Configuration

### Phase 1: IKE Configuration

**IKE** (Internet Key Exchange) negotiates security parameters, authenticates peers, and exchanges encryption keys. Go to **IP > IPsec > Profiles** and create a profile (e.g., "vpn-profile"):

- `Hash Algorithm`: `sha256` (sha1 is less secure)
- `Encryption Algorithm`: `aes256`
- `DH Group`: `modp1024` or `modp2048`
- `Lifetime`: e.g., `1h`

Then **IP > IPsec > Proposals**, add a proposal (e.g., "vpn-proposal"):

- `Name`: `vpn-proposal`
- `Authentication Algorithms`: `sha256`
- `Encryption Algorithms`: `aes256`
- `DH Group`: `modp1024`

Then **IP > IPsec > Peers**, add a peer (one per router, pointing at the other):

- `Address`: public IP of the *other* router
- `Secret`: strong pre-shared key (PSK), identical on both sides
- `Profile`: `vpn-profile`
- `Exchange Mode`: `main` (`aggressive` if dynamic IPs, but less secure)
- `My ID` / `Peer ID`: the public IPs (e.g., `1.2.3.4` / `5.6.7.8`)

### Phase 2: IPSec Configuration

Go to **IP > IPsec > Policies**, add a policy:

- `Action`: `encrypt`
- `Src. Address`: local network, e.g., `192.168.1.0/24`
- `Dst. Address`: remote network, e.g., `192.168.2.0/24`
- `Protocol`: `all`, `Src. Port`/`Dst. Port`: `any`
- `Tunnel`: checked; `IPsec Protocols`: `esp`
- `Proposal`: `vpn-proposal`; `Peer`: the other router

Repeat on the other router swapping Src./Dst. addresses, and **enable the policies** on both.

### Firewall and Routing Considerations

- **Firewall:** allow UDP port 500 (IKE) and ESP (IP protocol 50); ensure inter-network traffic passes the firewall.
- **Routing:** add static routes — **IP > Routes**: `Dst. Address` = remote network, `Gateway` = the other router's public IP, `Distance` = 1. Mirror on the other side.

## Troubleshooting Common Issues

- **VPN fails to connect:** check RouterOS logs (IKE negotiation, auth failures, mismatched params); verify public IPs; check firewall (UDP 500, ESP); ensure PSKs match *exactly* on both ends.
- **Network connectivity issues:** verify routing tables/gateways; reduce MTU on the IPsec tunnel interface (typical 1400) to fix fragmentation; confirm networks don't overlap.
- **Performance issues:** check CPU usage (IPsec is CPU-intensive; use hardware encryption if available); `aes128` is faster than `aes256` but less secure; tunnel bandwidth is limited by each site's internet connection.

## Advanced MikroTik IPSec Configurations

- **Dynamic IP Addresses:** if your routers have dynamic IP addresses, you can configure IPsec to work with them using the

*[article tronqué dans l'extraction à cet endroit — la page source s'interrompt ici]*
