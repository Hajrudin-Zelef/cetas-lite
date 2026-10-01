---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75-5
title: "document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75.md
source_anchor: ""
source_lines: [625, 785]
sha256: f8f303824c722bb767850c14ac7e2794f855f477a98363975e6bde37b26768a6
---

# document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75

| require | Require use of IKEv2 Postquantum Preshared Key (PPK). | 
| Option | Description | 
|---|---|
| des-md5 | des-md5 | 
| des-sha1 | des-sha1 | 
| des-sha256 | des-sha256 | 
| des-sha384 | des-sha384 | 
| des-sha512 | des-sha512 | 
| 3des-md5 | 3des-md5 | 
| 3des-sha1 | 3des-sha1 | 
| 3des-sha256 | 3des-sha256 | 
| 3des-sha384 | 3des-sha384 | 
| 3des-sha512 | 3des-sha512 | 
| aes128-md5 | aes128-md5 | 
| aes128-sha1 | aes128-sha1 | 
| aes128-sha256 | aes128-sha256 | 
| aes128-sha384 | aes128-sha384 | 
| aes128-sha512 | aes128-sha512 | 
| aes128gcm-prfsha1 | aes128gcm-prfsha1 | 
| aes128gcm-prfsha256 | aes128gcm-prfsha256 | 
| aes128gcm-prfsha384 | aes128gcm-prfsha384 | 
| aes128gcm-prfsha512 | aes128gcm-prfsha512 | 
| aes192-md5 | aes192-md5 | 
| aes192-sha1 | aes192-sha1 | 
| aes192-sha256 | aes192-sha256 | 
| aes192-sha384 | aes192-sha384 | 
| aes192-sha512 | aes192-sha512 | 
| aes256-md5 | aes256-md5 | 
| aes256-sha1 | aes256-sha1 | 
| aes256-sha256 | aes256-sha256 | 
| aes256-sha384 | aes256-sha384 | 
| aes256-sha512 | aes256-sha512 | 
| aes256gcm-prfsha1 | aes256gcm-prfsha1 | 
| aes256gcm-prfsha256 | aes256gcm-prfsha256 | 
| aes256gcm-prfsha384 | aes256gcm-prfsha384 | 
| aes256gcm-prfsha512 | aes256gcm-prfsha512 | 
| chacha20poly1305-prfsha1 | chacha20poly1305-prfsha1 | 
| chacha20poly1305-prfsha256 | chacha20poly1305-prfsha256 | 
| chacha20poly1305-prfsha384 | chacha20poly1305-prfsha384 | 
| chacha20poly1305-prfsha512 | chacha20poly1305-prfsha512 | 
| aria128-md5 | aria128-md5 | 
| aria128-sha1 | aria128-sha1 | 
| aria128-sha256 | aria128-sha256 | 
| aria128-sha384 | aria128-sha384 | 
| aria128-sha512 | aria128-sha512 | 
| aria192-md5 | aria192-md5 | 
| aria192-sha1 | aria192-sha1 | 
| aria192-sha256 | aria192-sha256 | 
| aria192-sha384 | aria192-sha384 | 
| aria192-sha512 | aria192-sha512 | 
| aria256-md5 | aria256-md5 | 
| aria256-sha1 | aria256-sha1 | 
| aria256-sha256 | aria256-sha256 | 
| aria256-sha384 | aria256-sha384 | 
| aria256-sha512 | aria256-sha512 | 
| seed-md5 | seed-md5 | 
| seed-sha1 | seed-sha1 | 
| seed-sha256 | seed-sha256 | 
| seed-sha384 | seed-sha384 | 
| seed-sha512 | seed-sha512 | 
| Option | Description | 
|---|---|
| disable | Disable use of a Quantum Key Distribution (QKD) server. | 
| allow | Allow, but do not require, use of a Quantum Key Distribution (QKD) server. | 
| require | Require use of a Quantum Key Distribution (QKD) server. | 
| Option | Description | 
|---|---|
| disable | Disable IKE SA re-authentication. | 
| enable | Enable IKE SA re-authentication. | 
| Option | Description | 
|---|---|
| enable | Enable phase1 rekey. | 
| disable | Disable phase1 rekey. | 
| Option | Description | 
|---|---|
| any | Match any IPv4 gateway address. | 
| ipmask | Match IPv4 gateway address and mask. | 
| iprange | Match IPv4 gateway address range. | 
| geography | Match IPv4 gateway address from a specified country. | 
| Option | Description | 
|---|---|
| any | Match any IPv6 gateway address. | 
| ipprefix | Match IPv6 gateway address and prefix. | 
| iprange | Match IPv6 gateway address range. | 
| geography | Match IPv6 gateway address from a specified country. | 
| Option | Description | 
|---|---|
| pkcs1 | RSASSA PKCS#1 v1.5. | 
| pss | RSASSA Probabilistic Signature Scheme (PSS). | 
| Option | Description | 
|---|---|
| enable | Enable IKEv2 RSA signature hash algorithm override. | 
| disable | Disable IKEv2 RSA signature hash algorithm override. | 
| Option | Description | 
|---|---|
| disable | Disable saving XAuth username and password on VPN clients. | 
| enable | Enable saving XAuth username and password on VPN clients. | 
| Option | Description | 
|---|---|
| enable | Enable sending certificate chain. | 
| disable | Disable sending certificate chain. | 
| Option | Description | 
|---|---|
| sha1 | SHA1. | 
| sha2-256 | SHA2-256. | 
| sha2-384 | SHA2-384. | 
| sha2-512 | SHA2-512. | 
| Option | Description | 
|---|---|
| disable | Do not use UI suite. | 
| suite-b-gcm-128 | Use Suite-B-GCM-128. | 
| suite-b-gcm-256 | Use Suite-B-GCM-256. | 
| Option | Description | 
|---|---|
| udp | Use UDP transport for IKE. | 
| udp-fallback-tcp | Use UDP transport for IKE, with fallback to TCP transport. | 
| tcp | Use TCP transport for IKE. | 
| Option | Description | 
|---|---|
| static | Remote VPN gateway has fixed IP address. | 
| dynamic | Remote VPN gateway has dynamic IP address. | 
| ddns | Remote VPN gateway has dynamic IP address and is a dynamic DNS client. | 
| Option | Description | 
|---|---|
| disable | Disable Cisco Unity Configuration Method Extensions. | 
| enable | Enable Cisco Unity Configuration Method Extensions. | 
| Option | Description | 
|---|---|
| custom | Custom VPN configuration. | 
| dialup-forticlient | Dial Up - FortiClient Windows, Mac and Android. | 
| dialup-ios | Dial Up - iPhone / iPad Native IPsec Client. | 
| dialup-android | Dial Up - Android Native IPsec Client. | 
| dialup-windows | Dial Up - Windows Native IPsec Client. | 
| dialup-cisco | Dial Up - Cisco IPsec Client. | 
| static-fortigate | Site to Site - FortiGate. | 
| dialup-fortigate | Dial Up - FortiGate. | 
| static-cisco | Site to Site - Cisco. | 
| dialup-cisco-fw | Dialup Up - Cisco Firewall. | 
| simplified-static-fortigate | Site to Site - FortiGate (SD-WAN). | 
| hub-fortigate-auto-discovery | Hub role in a Hub-and-Spoke auto-discovery VPN. | 
| spoke-fortigate-auto-discovery | Spoke role in a Hub-and-Spoke auto-discovery VPN. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| client | Enable as client. | 
| pap | Enable as server PAP. | 
| chap | Enable as server CHAP. | 
| auto | Enable as server auto. | 
* This parameter may not exist in some models.
config ipv4-exclude-range
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| start-ip | Start of IPv4 exclusive range. | ipv4-address | Not Specified | 0.0.0.0 | 
| end-ip | End of IPv4 exclusive range. | ipv4-address | Not Specified | 0.0.0.0 | 
config ipv6-exclude-range
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| start-ip | Start of IPv6 exclusive range. | ipv6-address | Not Specified | :: | 
| end-ip | End of IPv6 exclusive range. | ipv6-address | Not Specified | :: |
