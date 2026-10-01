---
id: collect-261001-general-networking/general-networking/questions-1077033-wireguard-destination-host-unreachable-on-internal-network-1b956e52
title: "Wireguard Destination Host Unreachable on internal network"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1077033-wireguard-destination-host-unreachable-on-internal-network-1b956e52.md
source_anchor: ""
source_lines: [1, 38]
sha256: 8635392dd598976f926f68704919e2815f8fa95bd9c70ff8c9a4a6086854c8db
---

# Wireguard Destination Host Unreachable on internal network

*Score : 1 | Source : https://serverfault.com/questions/1077033/wireguard-destination-host-unreachable-on-internal-network*

I'm setting up Wireguard to tunnel from a cloud VM to our internal network. The local server is using the Wireguard plugin for OPNSense.
OPNSense acts as firewall, dhcp, etc.
The cloud VM is not behind any firewall or anything.
interface: wg0
  public key: redacted
  private key: (hidden)
  listening port: 42001
peer: redacted
  endpoint: CLOUD_VM_PUBLIC_IP:42001
  allowed ips: 10.0.1.42/32
  latest handshake: 48 seconds ago
  transfer: 184.23 KiB received, 186.37 KiB sent
  persistent keepalive: every 21 seconds
  public key: redacted
  private key: (hidden)
  listening port: 42001
peer: redacted
  endpoint: LOCAL_PUBLIC_IP:42001
  allowed ips: 10.0.0.0/16
  latest handshake: 2 minutes, 14 seconds ago
  transfer: 1.30 KiB received, 1.20 KiB sent
  persistent keepalive: every 21 seconds
[Interface]
# set address to next address
Address = 10.0.1.42/16
ListenPort = 42001
PrivateKey = redacted
DNS = 1.1.1.1
[Peer]
PublicKey = redacted
Endpoint = LOCAL_PUBLIC_IP:42001
AllowedIPs = 10.0.0.0/16
PersistentKeepalive = 21
With this configuration, I can connect to the VM using the internal IP address 10.0.1.42 on OPNSense, but anything else shows 'Destination Host Unreachable'.  And trying to ping any internal ip in 10.0.0.0/16 from the cloud VM times out.
