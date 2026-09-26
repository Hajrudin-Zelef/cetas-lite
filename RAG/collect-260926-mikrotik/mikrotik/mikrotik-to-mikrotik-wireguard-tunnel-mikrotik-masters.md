---
id: collect-260926-mikrotik/mikrotik/mikrotik-to-mikrotik-wireguard-tunnel-mikrotik-masters
title: "mikrotik-to-mikrotik-wireguard-tunnel-mikrotik-masters"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/mikrotik-to-mikrotik-wireguard-tunnel-mikrotik-masters.md
source_anchor: ""
source_lines: [1, 129]
sha256: 006849f752832c9fe9455164dc8ca542b2feb2b12fc1479ddb3d12de9a790f46
---

# mikrotik-to-mikrotik-wireguard-tunnel-mikrotik-masters

## WireGuard Interface

On both sides (in our case CHR-AU & CHR-UK) we create a Wireguard interface

CHR-AU & CHR-UK

```
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1
```
Make a copy of the Public of each end (not the Private Key – this is only for the device itself and does not need to be shared).

*NB: Keys should not be publicly shared or advertised, and only configured on the required devices. They act like passwords or secrets and should be treated with the same privacy*

| CHR-AU Public Key: qI+Fwh/ZrSGjdkRjU1O74PXRkztEDLS/xRQ4q2fVESM= CHR-UK Public Key: DwtVNA1pXvvbWAZG6j4AhjZoF4dhXw732iErV2wL/X0= | 

## IP Address

Now give each of the WireGuard interfaces an IP (CHR-AU: 10.0.0.1/30 & CHR-UK: 10.0.0.2/30)

CHR-AU

CHR-UK

## WireGuard Peers

Now create a peer on each router (Wireguard > Peer) and use the following details

|  |  | CHR-AU | CHR-UK | 
| Name | Unique name *(name of opposite router)* | CHR-UK | CHR-AU | 
| Interface | Name of Wireguard interface created previously | wireguard1 | wireguard1 | 
| Public Key | Public key of opposite router | DwtVNA1pXvvbWAZG6j4AhjZoF4dhXw732iErV2wL/X0= | qI+Fwh/ZrSGjdkRjU1O74PXRkztEDLS/xRQ4q2fVESM= | 
| Endpoint | Public IP of opposite router | 18.132.39.216 | 3.106.63.236 | 
| Endpoint Port | Port specified on opposite router Wireguard interface  *(default: 13231)* | 13231 | 13231 | 
| Allowed Address | Address to allow over tunnel *(IP range of Wireguard interface and any other LANs intended to route over tunnel*)* | 10.0.0.0/30 10.2.254.0/24* | 10.0.0.0/30 10.1.254.0/24* | 

**  The subnets of the LANs can be ignored if using the GRE/Dynamic routing option outlined later in this tutorial*

CHR-AU

CHR-UK

CHR-AU

```
/interface wireguard peers
add allowed-address=10.0.0.0/30,10.2.254.0/24 endpoint-address=18.132.39.216 endpoint-port=13231 interface=wireguard1 \
    name=CHR-UK public-key="DwtVNA1pXvvbWAZG6j4AhjZoF4dhXw732iErV2wL/X0=
```
CHR-UK

```
/interface wireguard peers
add allowed-address=10.0.0.0/30,10.1.254.0/24 endpoint-address=3.106.63.236 endpoint-port=13231 interface=wireguard1 \
    name=CHR-AU public-key="qI+Fwh/ZrSGjdkRjU1O74PXRkztEDLS/xRQ4q2fVESM="
```
## Firewall Filter Rules (input chain)

To allow the connection to establish, we need to allow this through your deny input chain rule (if you don’t have this set up you should! Click this link to learn how)

### Option 1: Allow all from remote IP

CHR-AU

```
/ip firewall filter
add action=accept chain=input in-interface=ether1 src-address=18.132.39.216
```
CHR-UK

```
/ip firewall filter
add action=accept chain=input in-interface=ether1 src-address=3.106.63.236
```
### Option 2: Allow UDP/13231 (if using the default port)

CHR-AU

```
/ip firewall filter
add action=accept chain=input dst-port=13231 in-interface=ether1 protocol=udp
```
CHR-UK

```
/ip firewall filter
add action=accept chain=input dst-port=13231 in-interface=ether1 protocol=udp
```
## Testing

## Routing Option: via GRE

This option allows you to use dynamic routing to route between routers. This can be useful when there are multiple other links/routers/subnets on the other side of the links.

To do this, start by removing the LAN subnets for the Peer configurations

Now we will create a GRE tunnel to run over the WireGuard tunnel. This will allow the dynamic routing advertisements and a gateway IP to point static routes to. Use the WireGuard IP address as the GRE tunnel IP address.

Go to Interfaces and add a GRE Tunnel

CHR-AU

CHR-UK

CHR-AU

```
/interface gre
add local-address=10.0.0.1 name=gre-tunnel-uk remote-address=10.0.0.2
```
CHR-UK

```
/interface gre
add local-address=10.0.0.2 name=gre-tunnel-au remote-address=10.0.0.1
```
### IP Address

CHR-AU

CHR-UK

### Static Routes

Now we can use those GRE IP addresses to point static routes to (or use as Neighbour IPs for BGP)

CHR-AU

CHR-UK
