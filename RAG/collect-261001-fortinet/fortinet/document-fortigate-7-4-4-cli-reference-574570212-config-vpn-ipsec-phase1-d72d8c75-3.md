---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75-3
title: "document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75.md
source_anchor: ""
source_lines: [291, 417]
sha256: 78730cf9477aa06a7d8fad95aca60e05a7bbe9877cfb5192fe209a84dffa88e7
---

# document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75

| ipv4-split-include | IPv4 split-include subnets. | string | Maximum length: 79 |  | 
| ipv4-start-ip | Start of IPv4 range. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-wins-server1 | WINS server 1. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-wins-server2 | WINS server 2. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv6-dns-server1 | IPv6 DNS server 1. | ipv6-address | Not Specified | :: | 
| ipv6-dns-server2 | IPv6 DNS server 2. | ipv6-address | Not Specified | :: | 
| ipv6-dns-server3 | IPv6 DNS server 3. | ipv6-address | Not Specified | :: | 
| ipv6-end-ip | End of IPv6 range. | ipv6-address | Not Specified | :: | 
| ipv6-name | IPv6 address name. | string | Maximum length: 79 |  | 
| ipv6-prefix | IPv6 prefix. | integer | Minimum value: 1 Maximum value: 128 | 128 | 
| ipv6-split-exclude | IPv6 subnets that should not be sent over the IPsec tunnel. | string | Maximum length: 79 |  | 
| ipv6-split-include | IPv6 split-include subnets. | string | Maximum length: 79 |  | 
| ipv6-start-ip | Start of IPv6 range. | ipv6-address | Not Specified | :: | 
| keepalive | NAT-T keep alive interval. | integer | Minimum value: 5 Maximum value: 900 | 10 | 
| keylife | Time to wait in seconds before phase 1 encryption key expires. | integer | Minimum value: 120 Maximum value: 172800 | 86400 | 
| kms | Key Management Services server. | string | Maximum length: 35 |  | 
| link-cost | VPN tunnel underlay link cost. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| local-gw | Local VPN gateway. | ipv4-address | Not Specified | 0.0.0.0 | 
| localid | Local ID. | string | Maximum length: 63 |  | 
| localid-type | Local ID type. | option | - | auto | 
|  |  |  |  |  | 
| loopback-asymroute | Enable/disable asymmetric routing for IKE traffic on loopback interface. | option | - | enable | 
|  |  |  |  |  | 
| mesh-selector-type | Add selectors containing subsets of the configuration depending on traffic. | option | - | disable | 
|  |  |  |  |  | 
| mode | ID protection mode used to establish a secure channel. | option | - | main | 
|  |  |  |  |  | 
| mode-cfg | Enable/disable configuration method. | option | - | disable | 
|  |  |  |  |  | 
| mode-cfg-allow-client-selector | Enable/disable mode-cfg client to use custom phase2 selectors. | option | - | disable | 
|  |  |  |  |  | 
| name | IPsec remote gateway name. | string | Maximum length: 35 |  | 
| nattraversal | Enable/disable NAT traversal. | option | - | enable | 
|  |  |  |  |  | 
| negotiate-timeout | IKE SA negotiation timeout in seconds. | integer | Minimum value: 1 Maximum value: 300 | 30 | 
| network-id | VPN gateway network ID. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| network-overlay | Enable/disable network overlays. | option | - | disable | 
|  |  |  |  |  | 
| npu-offload * | Enable/disable offloading NPU. | option | - | enable | 
|  |  |  |  |  | 
| peer | Accept this peer certificate. | string | Maximum length: 35 |  | 
| peergrp | Accept this peer certificate group. | string | Maximum length: 35 |  | 
| peerid | Accept this peer identity. | string | Maximum length: 255 |  | 
| peertype | Accept this peer type. | option | - | peer | 
|  |  |  |  |  | 
| ppk | Enable/disable IKEv2 Postquantum Preshared Key (PPK). | option | - | disable | 
|  |  |  |  |  | 
| ppk-identity | IKEv2 Postquantum Preshared Key Identity. | string | Maximum length: 35 |  | 
| ppk-secret | IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x). | password-3 | Not Specified |  | 
| priority | Priority for routes added by IKE. | integer | Minimum value: 1 Maximum value: 65535 | 1 | 
| proposal | Phase1 proposal. | option | - |  | 
|  |  |  |  |  | 
| psksecret | Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x). | password-3 | Not Specified |  | 
| psksecret-remote | Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x). | password-3 | Not Specified |  | 
| qkd | Enable/disable use of Quantum Key Distribution (QKD) server. | option | - | disable | 
|  |  |  |  |  | 
| qkd-profile | Quantum Key Distribution (QKD) server profile. | string | Maximum length: 35 |  | 
| reauth | Enable/disable re-authentication upon IKE SA lifetime expiration. | option | - | disable | 
|  |  |  |  |  | 
| rekey | Enable/disable phase1 rekey. | option | - | enable | 
|  |  |  |  |  | 
| remote-gw | Remote VPN gateway. | ipv4-address | Not Specified | 0.0.0.0 | 
| remote-gw-country | IPv4 addresses associated to a specific country. | string | Maximum length: 2 |  | 
| remote-gw-end-ip | Last IPv4 address in the range. | ipv4-address-any | Not Specified | 0.0.0.0 | 
| remote-gw-match | Set type of IPv4 remote gateway address matching. | option | - | any | 
|  |  |  |  |  | 
| remote-gw-start-ip | First IPv4 address in the range. | ipv4-address-any | Not Specified | 0.0.0.0 | 
| remote-gw-subnet | IPv4 address and subnet mask. | ipv4-classnet-any | Not Specified | 0.0.0.0 0.0.0.0 | 
| remote-gw6-country | IPv6 addresses associated to a specific country. | string | Maximum length: 2 |  | 
| remote-gw6-end-ip | Last IPv6 address in the range. | ipv6-address | Not Specified | :: | 
| remote-gw6-match | Set type of IPv6 remote gateway address matching. | option | - | any | 
|  |  |  |  |  | 
| remote-gw6-start-ip | First IPv6 address in the range. | ipv6-address | Not Specified | :: | 
| remote-gw6-subnet | IPv6 address and prefix. | ipv6-network | Not Specified | ::/0 | 
| remotegw-ddns | Domain name of remote gateway. For example, name.ddns.com. | string | Maximum length: 63 |  | 
| rsa-signature-format | Digital Signature Authentication RSA signature format. | option | - | pkcs1 | 
|  |  |  |  |  | 
| rsa-signature-hash-override | Enable/disable IKEv2 RSA signature hash algorithm override. | option | - | disable | 
|  |  |  |  |  | 
| save-password | Enable/disable saving XAuth username and password on VPN clients. | option | - | disable | 
|  |  |  |  |  | 
| send-cert-chain | Enable/disable sending certificate chain. | option | - | enable | 
|  |  |  |  |  | 
| signature-hash-alg | Digital Signature Authentication hash algorithms. | option | - | sha2-512 | 
|  |  |  |  |  | 
| split-include-service | Split-include services. | string | Maximum length: 79 |  | 
| suite-b | Use Suite-B. | option | - | disable | 
|  |  |  |  |  | 
| transport | Set IKE transport protocol. | option | - | udp | 
|  |  |  |  |  | 
| type | Remote gateway type. | option | - | static | 
|  |  |  |  |  | 
| unity-support | Enable/disable support for Cisco UNITY Configuration Method extensions. | option | - | enable | 
|  |  |  |  |  | 
| usrgrp | User group name for dialup peers. | string | Maximum length: 35 |  | 
| wizard-type | GUI VPN Wizard Type. | option | - | custom | 
|  |  |  |  |  | 
| xauthtype | XAuth type. | option | - | disable | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| enable | Enable verification of RADIUS accounting record. | 
| disable | Disable verification of RADIUS accounting record. | 
| Option | Description | 
|---|---|
| enable | Automatically add a route to the remote gateway. | 
| disable | Do not automatically add a route to the remote gateway. | 
| Option | Description | 
|---|---|
| disable | Do not add a route to destination of peer selector. | 
| enable | Add route to destination of peer selector. | 
| Option | Description | 
|---|---|
| disable | Do not assign an IP address to the IPsec interface. | 
| enable | Assign an IP address to the IPsec interface. | 
| Option | Description | 
|---|---|
| range | Assign IP address from locally defined range. | 
| usrgrp | Assign IP address via user group. | 
| dhcp | Assign IP address via DHCP. | 
| name | Assign IP address from firewall address or group. | 
| Option | Description | 
|---|---|
| psk | PSK authentication method. | 
| signature | Signature authentication method. | 
| Option | Description | 
|---|---|
