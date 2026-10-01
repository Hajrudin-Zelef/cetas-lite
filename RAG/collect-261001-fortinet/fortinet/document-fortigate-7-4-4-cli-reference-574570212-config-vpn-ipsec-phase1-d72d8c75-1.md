---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75-1
title: "document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75.md
source_anchor: ""
source_lines: [1, 190]
sha256: 6e0a705ec2f6c9c2e2ad724ff2eb9a4df4575fa71b64c9b21dc482309006b5aa
---

# document-fortigate-7-4-4-cli-reference-574570212-config-vpn-ipsec-phase1-d72d8c75

config vpn ipsec phase1
config vpn ipsec phase1
Configure VPN remote gateway.
config vpn ipsec phase1
    Description: Configure VPN remote gateway.
    edit <name>
        set acct-verify [enable|disable]
        set add-gw-route [enable|disable]
        set add-route [disable|enable]
        set assign-ip [disable|enable]
        set assign-ip-from [range|usrgrp|...]
        set authmethod [psk|signature]
        set authmethod-remote [psk|signature]
        set authpasswd {password}
        set authusr {string}
        set authusrgrp {string}
        set auto-negotiate [enable|disable]
        set azure-ad-autoconnect [enable|disable]
        set backup-gateway <address1>, <address2>, ...
        set banner {var-string}
        set cert-id-validation [enable|disable]
        set cert-peer-username-strip [disable|enable]
        set cert-peer-username-validation [none|othername|...]
        set cert-trust-store [local|ems]
        set certificate <name1>, <name2>, ...
        set childless-ike [enable|disable]
        set client-auto-negotiate [disable|enable]
        set client-keep-alive [disable|enable]
        set client-resume [enable|disable]
        set client-resume-interval {integer}
        set comments {var-string}
        set dev-id {string}
        set dev-id-notification [disable|enable]
        set dhcp-ra-giaddr {ipv4-address}
        set dhcp6-ra-linkaddr {ipv6-address}
        set dhgrp {option1}, {option2}, ...
        set digital-signature-auth [enable|disable]
        set distance {integer}
        set dns-mode [manual|auto]
        set domain {string}
        set dpd [disable|on-idle|...]
        set dpd-retrycount {integer}
        set dpd-retryinterval {user}
        set eap [enable|disable]
        set eap-cert-auth [enable|disable]
        set eap-exclude-peergrp {string}
        set eap-identity [use-id-payload|send-request]
        set ems-sn-check [enable|disable]
        set enforce-unique-id [disable|keep-new|...]
        set esn [require|allow|...]
        set exchange-fgt-device-id [enable|disable]
        set fallback-tcp-threshold {integer}
        set fec-base {integer}
        set fec-codec [rs|xor]
        set fec-egress [enable|disable]
        set fec-health-check {string}
        set fec-ingress [enable|disable]
        set fec-mapping-profile {string}
        set fec-receive-timeout {integer}
        set fec-redundant {integer}
        set fec-send-timeout {integer}
        set fgsp-sync [enable|disable]
        set fortinet-esp [enable|disable]
        set fragmentation [enable|disable]
        set fragmentation-mtu {integer}
        set group-authentication [enable|disable]
        set group-authentication-secret {password-3}
        set ha-sync-esp-seqno [enable|disable]
        set idle-timeout [enable|disable]
        set idle-timeoutinterval {integer}
        set ike-version [1|2]
        set inbound-dscp-copy [enable|disable]
        set include-local-lan [disable|enable]
        set interface {string}
        set internal-domain-list <domain-name1>, <domain-name2>, ...
        set ip-delay-interval {integer}
        set ipv4-dns-server1 {ipv4-address}
        set ipv4-dns-server2 {ipv4-address}
        set ipv4-dns-server3 {ipv4-address}
        set ipv4-end-ip {ipv4-address}
        config ipv4-exclude-range
            Description: Configuration Method IPv4 exclude ranges.
            edit <id>
                set start-ip {ipv4-address}
                set end-ip {ipv4-address}
            next
        end
        set ipv4-name {string}
        set ipv4-netmask {ipv4-netmask}
        set ipv4-split-exclude {string}
        set ipv4-split-include {string}
        set ipv4-start-ip {ipv4-address}
        set ipv4-wins-server1 {ipv4-address}
        set ipv4-wins-server2 {ipv4-address}
        set ipv6-dns-server1 {ipv6-address}
        set ipv6-dns-server2 {ipv6-address}
        set ipv6-dns-server3 {ipv6-address}
        set ipv6-end-ip {ipv6-address}
        config ipv6-exclude-range
            Description: Configuration method IPv6 exclude ranges.
            edit <id>
                set start-ip {ipv6-address}
                set end-ip {ipv6-address}
            next
        end
        set ipv6-name {string}
        set ipv6-prefix {integer}
        set ipv6-split-exclude {string}
        set ipv6-split-include {string}
        set ipv6-start-ip {ipv6-address}
        set keepalive {integer}
        set keylife {integer}
        set kms {string}
        set link-cost {integer}
        set local-gw {ipv4-address}
        set localid {string}
        set localid-type [auto|fqdn|...]
        set loopback-asymroute [enable|disable]
        set mesh-selector-type [disable|subnet|...]
        set mode [aggressive|main]
        set mode-cfg [disable|enable]
        set mode-cfg-allow-client-selector [disable|enable]
        set nattraversal [enable|disable|...]
        set negotiate-timeout {integer}
        set network-id {integer}
        set network-overlay [disable|enable]
        set npu-offload [enable|disable]
        set peer {string}
        set peergrp {string}
        set peerid {string}
        set peertype [any|one|...]
        set ppk [disable|allow|...]
        set ppk-identity {string}
        set ppk-secret {password-3}
        set priority {integer}
        set proposal {option1}, {option2}, ...
        set psksecret {password-3}
        set psksecret-remote {password-3}
        set qkd [disable|allow|...]
        set qkd-profile {string}
        set reauth [disable|enable]
        set rekey [enable|disable]
        set remote-gw {ipv4-address}
        set remote-gw-country {string}
        set remote-gw-end-ip {ipv4-address-any}
        set remote-gw-match [any|ipmask|...]
        set remote-gw-start-ip {ipv4-address-any}
        set remote-gw-subnet {ipv4-classnet-any}
        set remote-gw6-country {string}
        set remote-gw6-end-ip {ipv6-address}
        set remote-gw6-match [any|ipprefix|...]
        set remote-gw6-start-ip {ipv6-address}
        set remote-gw6-subnet {ipv6-network}
        set remotegw-ddns {string}
        set rsa-signature-format [pkcs1|pss]
        set rsa-signature-hash-override [enable|disable]
        set save-password [disable|enable]
        set send-cert-chain [enable|disable]
        set signature-hash-alg {option1}, {option2}, ...
        set split-include-service {string}
        set suite-b [disable|suite-b-gcm-128|...]
        set transport [udp|udp-fallback-tcp|...]
        set type [static|dynamic|...]
        set unity-support [disable|enable]
        set usrgrp {string}
        set wizard-type [custom|dialup-forticlient|...]
        set xauthtype [disable|client|...]
    next
end
                                            config vpn ipsec phase1
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| acct-verify | Enable/disable verification of RADIUS accounting record. | option | - | disable | 
|  |  |  |  |  | 
| add-gw-route | Enable/disable automatically add a route to the remote gateway. | option | - | disable | 
|  |  |  |  |  | 
| add-route | Enable/disable control addition of a route to peer destination selector. | option | - | disable | 
|  |  |  |  |  | 
| assign-ip | Enable/disable assignment of IP to IPsec interface via configuration method. | option | - | enable | 
|  |  |  |  |  | 
| assign-ip-from | Method by which the IP address will be assigned. | option | - | range | 
|  |  |  |  |  | 
| authmethod | Authentication method. | option | - | psk | 
|  |  |  |  |  | 
| authmethod-remote | Authentication method (remote side). | option | - |  | 
|  |  |  |  |  | 
| authpasswd | XAuth password (max 35 characters). | password | Not Specified |  | 
| authusr | XAuth user name. | string | Maximum length: 64 |  | 
| authusrgrp | Authentication user group. | string | Maximum length: 35 |  | 
| auto-negotiate | Enable/disable automatic initiation of IKE SA negotiation. | option | - | enable | 
