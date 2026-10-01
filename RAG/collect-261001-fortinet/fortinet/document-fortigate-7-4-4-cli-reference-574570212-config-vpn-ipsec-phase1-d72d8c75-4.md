---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75-4
title: "document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75.md
source_anchor: ""
source_lines: [418, 624]
sha256: 1273c5195b3d271e5cee77e4db312bc45d5c42dcbfa7e6a040e29ef0220f618e
---

# document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75

| enable | Enable automatic initiation of IKE SA negotiation. | 
| disable | Disable automatic initiation of IKE SA negotiation. | 
| Option | Description | 
|---|---|
| enable | Enable Azure AD Auto-Connect for FortiClient. | 
| disable | Disable Azure AD Auto-Connect for FortiClient. | 
| Option | Description | 
|---|---|
| enable | Enable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. | 
| disable | Disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. | 
| Option | Description | 
|---|---|
| disable | Disable domain stripping on certificate identity. | 
| enable | Enable domain stripping on certificate identity. | 
| Option | Description | 
|---|---|
| none | Disable cross validation of peer username and the identity in the peer's certificate. | 
| othername | Validate principal name in SAN othername. | 
| rfc822name | Validate RFC822 email address in SAN. | 
| cn | Validate CN in subject. | 
| Option | Description | 
|---|---|
| local | Use local CA certificate. | 
| ems | Use EMS CA certificate. | 
| Option | Description | 
|---|---|
| enable | Enable childless IKEv2 initiation (RFC 6023). | 
| disable | Disable childless IKEv2 initiation (RFC 6023). | 
| Option | Description | 
|---|---|
| disable | Disable allowing the VPN client to bring up the tunnel when there is no traffic. | 
| enable | Enable allowing the VPN client to bring up the tunnel when there is no traffic. | 
| Option | Description | 
|---|---|
| disable | Disable allowing the VPN client to keep the tunnel up when there is no traffic. | 
| enable | Enable allowing the VPN client to keep the tunnel up when there is no traffic. | 
| Option | Description | 
|---|---|
| enable | Enable client session resumption. | 
| disable | Disable client session resumption. | 
| Option | Description | 
|---|---|
| disable | Disable device ID notification. | 
| enable | Enable device ID notification. | 
| Option | Description | 
|---|---|
| 1 | DH Group 1. | 
| 2 | DH Group 2. | 
| 5 | DH Group 5. | 
| 14 | DH Group 14. | 
| 15 | DH Group 15. | 
| 16 | DH Group 16. | 
| 17 | DH Group 17. | 
| 18 | DH Group 18. | 
| 19 | DH Group 19. | 
| 20 | DH Group 20. | 
| 21 | DH Group 21. | 
| 27 | DH Group 27. | 
| 28 | DH Group 28. | 
| 29 | DH Group 29. | 
| 30 | DH Group 30. | 
| 31 | DH Group 31. | 
| 32 | DH Group 32. | 
| Option | Description | 
|---|---|
| enable | Enable IKEv2 Digital Signature Authentication (RFC 7427). | 
| disable | Disable IKEv2 Digital Signature Authentication (RFC 7427). | 
| Option | Description | 
|---|---|
| manual | Manually configure DNS servers. | 
| auto | Use default DNS servers. | 
| Option | Description | 
|---|---|
| disable | Disable Dead Peer Detection. | 
| on-idle | Trigger Dead Peer Detection when IPsec is idle. | 
| on-demand | Trigger Dead Peer Detection when IPsec traffic is sent but no reply is received from the peer. | 
| Option | Description | 
|---|---|
| enable | Enable IKEv2 EAP authentication. | 
| disable | Disable IKEv2 EAP authentication. | 
| Option | Description | 
|---|---|
| enable | Enable peer certificate authentication in addition to EAP if peer is a FortiClient endpoint. | 
| disable | Disable peer certificate authentication in addition to EAP if peer is a FortiClient endpoint. | 
| Option | Description | 
|---|---|
| use-id-payload | Use IKEv2 IDi payload to resolve peer identity. | 
| send-request | Use EAP identity request to resolve peer identity. | 
| Option | Description | 
|---|---|
| enable | Enable EMS serial number verification. | 
| disable | Disable EMS serial number verification. | 
| Option | Description | 
|---|---|
| disable | Disable peer ID uniqueness enforcement. | 
| keep-new | Enforce peer ID uniqueness, keep new connection if collision found. | 
| keep-old | Enforce peer ID uniqueness, keep old connection if collision found. | 
| Option | Description | 
|---|---|
| require | Require extended sequence number. | 
| allow | Allow extended sequence number. | 
| disable | Disable extended sequence number. | 
| Option | Description | 
|---|---|
| enable | Enable exchange of FortiGate device identifier. | 
| disable | Disable exchange of FortiGate device identifier. | 
| Option | Description | 
|---|---|
| rs | Reed-Solomon FEC algorithm. | 
| xor | XOR FEC algorithm. | 
| Option | Description | 
|---|---|
| enable | Enable Forward Error Correction for egress IPsec traffic. | 
| disable | Disable Forward Error Correction for egress IPsec traffic. | 
| Option | Description | 
|---|---|
| enable | Enable Forward Error Correction for ingress IPsec traffic. | 
| disable | Disable Forward Error Correction for ingress IPsec traffic. | 
| Option | Description | 
|---|---|
| enable | Enable IPsec syncing of tunnels to other cluster members. | 
| disable | Disable IPsec syncing of tunnels to other cluster members. | 
| Option | Description | 
|---|---|
| enable | Enable Fortinet ESP encapsulation. | 
| disable | Disable Fortinet ESP encapsulaton. | 
| Option | Description | 
|---|---|
| enable | Enable intra-IKE fragmentation support on re-transmission. | 
| disable | Disable intra-IKE fragmentation support. | 
| Option | Description | 
|---|---|
| enable | Enable IKEv2 IDi group authentication. | 
| disable | Disable IKEv2 IDi group authentication. | 
| Option | Description | 
|---|---|
| enable | Enable HA syncing of ESP sequence numbers. | 
| disable | Disable HA syncing of ESP sequence numbers. | 
| Option | Description | 
|---|---|
| enable | Enable IPsec tunnel idle timeout. | 
| disable | Disable IPsec tunnel idle timeout. | 
| Option | Description | 
|---|---|
| 1 | Use IKEv1 protocol. | 
| 2 | Use IKEv2 protocol. | 
| Option | Description | 
|---|---|
| enable | Enable copy the dscp in the ESP header to the inner IP Header. | 
| disable | Disable copy the dscp in the ESP header to the inner IP Header. | 
| Option | Description | 
|---|---|
| disable | Disable local LAN access on Unity clients. | 
| enable | Enable local LAN access on Unity clients. | 
| Option | Description | 
|---|---|
| auto | Select ID type automatically. | 
| fqdn | Use fully qualified domain name. | 
| user-fqdn | Use user fully qualified domain name. | 
| keyid | Use key-id string. | 
| address | Use local IP address. | 
| asn1dn | Use ASN.1 distinguished name. | 
| Option | Description | 
|---|---|
| enable | Allow ingress/egress IKE traffic to be routed over different interfaces. | 
| disable | Ingress/egress IKE traffic must be routed over the same interface. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| subnet | Enable addition of matching subnet selector. | 
| host | Enable addition of host to host selector. | 
| Option | Description | 
|---|---|
| aggressive | Aggressive mode. | 
| main | Main mode. | 
| Option | Description | 
|---|---|
| disable | Disable Configuration Method. | 
| enable | Enable Configuration Method. | 
| Option | Description | 
|---|---|
| disable | Mode-cfg client to use wildcard selectors. | 
| enable | Mode-cfg client to use custom selectors. | 
| Option | Description | 
|---|---|
| enable | Enable IPsec NAT traversal. | 
| disable | Disable IPsec NAT traversal. | 
| forced | Force IPsec NAT traversal on. | 
| Option | Description | 
|---|---|
| disable | Disable network overlays. | 
| enable | Enable network overlays. | 
| Option | Description | 
|---|---|
| enable | Enable NPU offloading. | 
| disable | Disable NPU offloading. | 
| Option | Description | 
|---|---|
| any | Accept any peer ID. | 
| one | Accept this peer ID. | 
| dialup | Accept peer ID in dialup group. | 
| peer | Accept this peer certificate. | 
| peergrp | Accept this peer certificate group. | 
| Option | Description | 
|---|---|
| disable | Disable use of IKEv2 Postquantum Preshared Key (PPK). | 
| allow | Allow, but do not require, use of IKEv2 Postquantum Preshared Key (PPK). | 
