---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75-2
title: "document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75.md
source_anchor: ""
source_lines: [191, 290]
sha256: 3cf6a449333b47e848c68da0952486b8fef1c62a35fcaa1bff78409e647c3cd7
---

# document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75

|  |  |  |  |  | 
| azure-ad-autoconnect | Enable/disable Azure AD Auto-Connect for FortiClient. | option | - | disable | 
|  |  |  |  |  | 
| backup-gateway <address> | Instruct unity clients about the backup gateway address(es). Address of backup gateway. | string | Maximum length: 79 |  | 
| banner | Message that unity client should display after connecting. | var-string | Maximum length: 1024 |  | 
| cert-id-validation | Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. | option | - | enable | 
|  |  |  |  |  | 
| cert-peer-username-strip | Enable/disable domain stripping on certificate identity. | option | - | disable | 
|  |  |  |  |  | 
| cert-peer-username-validation | Enable/disable cross validation of peer username and the identity in the peer's certificate. | option | - | none | 
|  |  |  |  |  | 
| cert-trust-store | CA certificate trust store. | option | - | local | 
|  |  |  |  |  | 
| certificate <name> | Names of up to 4 signed personal certificates. Certificate name. | string | Maximum length: 79 |  | 
| childless-ike | Enable/disable childless IKEv2 initiation (RFC 6023). | option | - | disable | 
|  |  |  |  |  | 
| client-auto-negotiate | Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic. | option | - | disable | 
|  |  |  |  |  | 
| client-keep-alive | Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic. | option | - | disable | 
|  |  |  |  |  | 
| client-resume | Enable/disable resumption of offline FortiClient sessions. When a FortiClient enabled laptop is closed or enters sleep/hibernate mode, enabling this feature allows FortiClient to keep the tunnel during this period, and allows users to immediately resume using the IPsec tunnel when the device wakes up. | option | - | disable | 
|  |  |  |  |  | 
| client-resume-interval | Maximum time in seconds during which a VPN client may resume using a tunnel after a client PC has entered sleep mode or temporarily lost its network connection. | integer | Minimum value: 120 Maximum value: 172800 | 1800 | 
| comments | Comment. | var-string | Maximum length: 255 |  | 
| dev-id | Device ID carried by the device ID notification. | string | Maximum length: 63 |  | 
| dev-id-notification | Enable/disable device ID notification. | option | - | disable | 
|  |  |  |  |  | 
| dhcp-ra-giaddr | Relay agent gateway IP address to use in the giaddr field of DHCP requests. | ipv4-address | Not Specified | 0.0.0.0 | 
| dhcp6-ra-linkaddr | Relay agent IPv6 link address to use in DHCP6 requests. | ipv6-address | Not Specified | :: | 
| dhgrp | DH group. | option | - | 14 | 
|  |  |  |  |  | 
| digital-signature-auth | Enable/disable IKEv2 Digital Signature Authentication (RFC 7427). | option | - | disable | 
|  |  |  |  |  | 
| distance | Distance for routes added by IKE. | integer | Minimum value: 1 Maximum value: 255 | 15 | 
| dns-mode | DNS server mode. | option | - | manual | 
|  |  |  |  |  | 
| domain | Instruct unity clients about the single default DNS domain. | string | Maximum length: 63 |  | 
| dpd | Dead Peer Detection mode. | option | - | on-demand | 
|  |  |  |  |  | 
| dpd-retrycount | Number of DPD retry attempts. | integer | Minimum value: 0 Maximum value: 10 | 3 | 
| dpd-retryinterval | DPD retry interval. | user | Not Specified |  | 
| eap | Enable/disable IKEv2 EAP authentication. | option | - | disable | 
|  |  |  |  |  | 
| eap-cert-auth | Enable/disable peer certificate authentication in addition to EAP if peer is a FortiClient endpoint. | option | - | disable | 
|  |  |  |  |  | 
| eap-exclude-peergrp | Peer group excluded from EAP authentication. | string | Maximum length: 35 |  | 
| eap-identity | IKEv2 EAP peer identity type. | option | - | use-id-payload | 
|  |  |  |  |  | 
| ems-sn-check | Enable/disable verification of EMS serial number. | option | - | disable | 
|  |  |  |  |  | 
| enforce-unique-id | Enable/disable peer ID uniqueness check. | option | - | disable | 
|  |  |  |  |  | 
| esn * | Extended sequence number (ESN) negotiation. | option | - | disable | 
|  |  |  |  |  | 
| exchange-fgt-device-id | Enable/disable device identifier exchange with peer FortiGate units for use of VPN monitor data by FortiManager. | option | - | disable | 
|  |  |  |  |  | 
| fallback-tcp-threshold | Timeout in seconds before falling back IKE/IPsec traffic to tcp. | integer | Minimum value: 1 Maximum value: 300 | 15 | 
| fec-base | Number of base Forward Error Correction packets. | integer | Minimum value: 1 Maximum value: 20 | 10 | 
| fec-codec | Forward Error Correction encoding/decoding algorithm. | option | - | rs | 
|  |  |  |  |  | 
| fec-egress | Enable/disable Forward Error Correction for egress IPsec traffic. | option | - | disable | 
|  |  |  |  |  | 
| fec-health-check | SD-WAN health check. | string | Maximum length: 35 |  | 
| fec-ingress | Enable/disable Forward Error Correction for ingress IPsec traffic. | option | - | disable | 
|  |  |  |  |  | 
| fec-mapping-profile | Forward Error Correction (FEC) mapping profile. | string | Maximum length: 35 |  | 
| fec-receive-timeout | Timeout in milliseconds before dropping Forward Error Correction packets. | integer | Minimum value: 1 Maximum value: 1000 | 50 | 
| fec-redundant | Number of redundant Forward Error Correction packets. | integer | Minimum value: 1 Maximum value: 5 | 1 | 
| fec-send-timeout | Timeout in milliseconds before sending Forward Error Correction packets. | integer | Minimum value: 1 Maximum value: 1000 | 5 | 
| fgsp-sync | Enable/disable IPsec syncing of tunnels for FGSP IPsec. | option | - | disable | 
|  |  |  |  |  | 
| fortinet-esp | Enable/disable Fortinet ESP encapsulaton. | option | - | disable | 
|  |  |  |  |  | 
| fragmentation | Enable/disable fragment IKE message on re-transmission. | option | - | enable | 
|  |  |  |  |  | 
| fragmentation-mtu | IKE fragmentation MTU. | integer | Minimum value: 500 Maximum value: 16000 | 1200 | 
| group-authentication | Enable/disable IKEv2 IDi group authentication. | option | - | disable | 
|  |  |  |  |  | 
| group-authentication-secret | Password for IKEv2 ID group authentication. ASCII string or hexadecimal indicated by a leading 0x. | password-3 | Not Specified |  | 
| ha-sync-esp-seqno | Enable/disable sequence number jump ahead for IPsec HA. | option | - | enable | 
|  |  |  |  |  | 
| idle-timeout | Enable/disable IPsec tunnel idle timeout. | option | - | disable | 
|  |  |  |  |  | 
| idle-timeoutinterval | IPsec tunnel idle timeout in minutes. | integer | Minimum value: 5 Maximum value: 43200 | 15 | 
| ike-version | IKE protocol version. | option | - | 1 | 
|  |  |  |  |  | 
| inbound-dscp-copy | Enable/disable copy the dscp in the ESP header to the inner IP Header. | option | - | disable | 
|  |  |  |  |  | 
| include-local-lan | Enable/disable allow local LAN access on unity clients. | option | - | disable | 
|  |  |  |  |  | 
| interface | Local physical, aggregate, or VLAN outgoing interface. | string | Maximum length: 35 |  | 
| internal-domain-list <domain-name> | One or more internal domain names in quotes separated by spaces. Domain name. | string | Maximum length: 79 |  | 
| ip-delay-interval | IP address reuse delay interval in seconds. | integer | Minimum value: 0 Maximum value: 28800 | 0 | 
| ipv4-dns-server1 | IPv4 DNS server 1. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-dns-server2 | IPv4 DNS server 2. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-dns-server3 | IPv4 DNS server 3. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-end-ip | End of IPv4 range. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv4-name | IPv4 address name. | string | Maximum length: 79 |  | 
| ipv4-netmask | IPv4 Netmask. | ipv4-netmask | Not Specified | 255.255.255.255 | 
| ipv4-split-exclude | IPv4 subnets that should not be sent over the IPsec tunnel. | string | Maximum length: 79 |  | 
