---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33-1
title: "document-fortigate-7-4-8-cli-reference-333889629-1717ac33"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33.md
source_anchor: ""
source_lines: [1, 185]
sha256: 756ff2707ae038b6475d8dca65fe87a13156c603201033d7e2590568433cb576
---

# document-fortigate-7-4-8-cli-reference-333889629-1717ac33

config firewall policy
config firewall policy
Configure IPv4/IPv6 policies.
config firewall policy
    Description: Configure IPv4/IPv6 policies.
    edit <policyid>
        set action [accept|deny|...]
        set anti-replay [enable|disable]
        set application-list {string}
        set auth-cert {string}
        set auth-path [enable|disable]
        set auth-redirect-addr {string}
        set auto-asic-offload [enable|disable]
        set av-profile {string}
        set block-notification [enable|disable]
        set captive-portal-exempt [enable|disable]
        set capture-packet [enable|disable]
        set casb-profile {string}
        set cifs-profile {string}
        set comments {var-string}
        set custom-log-fields <field-id1>, <field-id2>, ...
        set decrypted-traffic-mirror {string}
        set delay-tcp-npu-session [enable|disable]
        set diameter-filter-profile {string}
        set diffserv-copy [enable|disable]
        set diffserv-forward [enable|disable]
        set diffserv-reverse [enable|disable]
        set diffservcode-forward {user}
        set diffservcode-rev {user}
        set disclaimer [enable|disable]
        set dlp-profile {string}
        set dnsfilter-profile {string}
        set dsri [enable|disable]
        set dstaddr <name1>, <name2>, ...
        set dstaddr-negate [enable|disable]
        set dstaddr6 <name1>, <name2>, ...
        set dstaddr6-negate [enable|disable]
        set dstintf <name1>, <name2>, ...
        set dynamic-shaping [enable|disable]
        set email-collect [enable|disable]
        set emailfilter-profile {string}
        set fec [enable|disable]
        set file-filter-profile {string}
        set firewall-session-dirty [check-all|check-new]
        set fixedport [enable|disable]
        set fsso-agent-for-ntlm {string}
        set fsso-groups <name1>, <name2>, ...
        set geoip-anycast [enable|disable]
        set geoip-match [physical-location|registered-location]
        set groups <name1>, <name2>, ...
        set http-policy-redirect [enable|disable]
        set icap-profile {string}
        set identity-based-route {string}
        set inbound [enable|disable]
        set inspection-mode [proxy|flow]
        set internet-service [enable|disable]
        set internet-service-custom <name1>, <name2>, ...
        set internet-service-custom-group <name1>, <name2>, ...
        set internet-service-group <name1>, <name2>, ...
        set internet-service-name <name1>, <name2>, ...
        set internet-service-negate [enable|disable]
        set internet-service-src [enable|disable]
        set internet-service-src-custom <name1>, <name2>, ...
        set internet-service-src-custom-group <name1>, <name2>, ...
        set internet-service-src-group <name1>, <name2>, ...
        set internet-service-src-name <name1>, <name2>, ...
        set internet-service-src-negate [enable|disable]
        set internet-service6 [enable|disable]
        set internet-service6-custom <name1>, <name2>, ...
        set internet-service6-custom-group <name1>, <name2>, ...
        set internet-service6-group <name1>, <name2>, ...
        set internet-service6-name <name1>, <name2>, ...
        set internet-service6-negate [enable|disable]
        set internet-service6-src [enable|disable]
        set internet-service6-src-custom <name1>, <name2>, ...
        set internet-service6-src-custom-group <name1>, <name2>, ...
        set internet-service6-src-group <name1>, <name2>, ...
        set internet-service6-src-name <name1>, <name2>, ...
        set internet-service6-src-negate [enable|disable]
        set ippool [enable|disable]
        set ips-sensor {string}
        set ips-voip-filter {string}
        set logtraffic [all|utm|...]
        set logtraffic-start [enable|disable]
        set match-vip [enable|disable]
        set match-vip-only [enable|disable]
        set name {string}
        set nat [enable|disable]
        set nat46 [enable|disable]
        set nat64 [enable|disable]
        set natinbound [enable|disable]
        set natip {ipv4-classnet}
        set natoutbound [enable|disable]
        set network-service-dynamic <name1>, <name2>, ...
        set network-service-src-dynamic <name1>, <name2>, ...
        set np-acceleration [enable|disable]
        set ntlm [enable|disable]
        set ntlm-enabled-browsers <user-agent-string1>, <user-agent-string2>, ...
        set ntlm-guest [enable|disable]
        set outbound [enable|disable]
        set passive-wan-health-measurement [enable|disable]
        set pcp-inbound [enable|disable]
        set pcp-outbound [enable|disable]
        set pcp-poolname <name1>, <name2>, ...
        set per-ip-shaper {string}
        set permit-any-host [enable|disable]
        set permit-stun-host [enable|disable]
        set policy-expiry [enable|disable]
        set policy-expiry-date {datetime}
        set policy-expiry-date-utc {user}
        set poolname <name1>, <name2>, ...
        set poolname6 <name1>, <name2>, ...
        set port-preserve [enable|disable]
        set profile-group {string}
        set profile-protocol-options {string}
        set profile-type [single|group]
        set radius-mac-auth-bypass [enable|disable]
        set redirect-url {var-string}
        set replacemsg-override-group {string}
        set reputation-direction [source|destination]
        set reputation-direction6 [source|destination]
        set reputation-minimum {integer}
        set reputation-minimum6 {integer}
        set rtp-addr <name1>, <name2>, ...
        set rtp-nat [disable|enable]
        set schedule {string}
        set schedule-timeout [enable|disable]
        set sctp-filter-profile {string}
        set send-deny-packet [disable|enable]
        set service <name1>, <name2>, ...
        set service-negate [enable|disable]
        set session-ttl {user}
        set sgt <id1>, <id2>, ...
        set sgt-check [enable|disable]
        set src-vendor-mac <id1>, <id2>, ...
        set srcaddr <name1>, <name2>, ...
        set srcaddr-negate [enable|disable]
        set srcaddr6 <name1>, <name2>, ...
        set srcaddr6-negate [enable|disable]
        set srcintf <name1>, <name2>, ...
        set ssh-filter-profile {string}
        set ssh-policy-redirect [enable|disable]
        set ssl-ssh-profile {string}
        set status [enable|disable]
        set tcp-mss-receiver {integer}
        set tcp-mss-sender {integer}
        set tcp-session-without-syn [all|data-only|...]
        set timeout-send-rst [enable|disable]
        set tos {user}
        set tos-mask {user}
        set tos-negate [enable|disable]
        set traffic-shaper {string}
        set traffic-shaper-reverse {string}
        set users <name1>, <name2>, ...
        set utm-status [enable|disable]
        set uuid {uuid}
        set videofilter-profile {string}
        set virtual-patch-profile {string}
        set vlan-cos-fwd {integer}
        set vlan-cos-rev {integer}
        set vlan-filter {user}
        set voip-profile {string}
        set vpntunnel {string}
        set waf-profile {string}
        set wanopt [enable|disable]
        set wanopt-detection [active|passive|...]
        set wanopt-passive-opt [default|transparent|...]
        set wanopt-peer {string}
        set wanopt-profile {string}
        set wccp [enable|disable]
        set webcache [enable|disable]
        set webcache-https [disable|enable]
        set webfilter-profile {string}
        set webproxy-forward-server {string}
        set webproxy-profile {string}
        set ztna-device-ownership [enable|disable]
        set ztna-ems-tag <name1>, <name2>, ...
        set ztna-ems-tag-secondary <name1>, <name2>, ...
        set ztna-geo-tag <name1>, <name2>, ...
        set ztna-policy-redirect [enable|disable]
        set ztna-status [enable|disable]
        set ztna-tags-match-logic [or|and]
    next
end
                                            config firewall policy
