---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-903162-important-dns-cli-commands-4844f4ab
title: "diagnose test application dnsproxy 6"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-903162-important-dns-cli-commands-4844f4ab.md
source_anchor: ""
source_lines: [1, 104]
sha256: 20a27f22763fb9dbf16a6b3035fa23fe1e969a485ee7b3f073c97e4cb0f1d368
---

# diagnose test application dnsproxy 6

Important DNS CLI commands
Important DNS CLI commands
DNS settings can be configured with the following CLI command:
config system dns
    set primary <ip_address>
    set secondary <ip_address>
    set protocol {cleartext dot doh}
    set ssl-certificate <string>
    set server-hostname <hostname>
    set domain <domains>
    set ip6-primary <ip6_address>
    set ip6-secondary <ip6_address>
    set timeout <integer>
    set retry <integer>
    set dns-cache-limit <integer>
    set dns-cache-ttl <integer>
    set cache-notfound-responses {enable | disable}
    set interface-select-method {auto | sdwan | specify}
    set interface <interface>
    set source-ip <class_ip>
    set server-select-method {least-rtt | failover}
    set alt-primary <ip_address>
    set alt-secondary <ip_address>
    set log {disable |error | all}
    set fqdn-cache-ttl <integer>
    set fqdn-min-refresh <integer>
    set fqdn-max-refresh <integer>
end
                                            For a FortiGate with multiple logical CPUs, you can set the DNS process number from 1 to the number of logical CPUs. The default DNS process number is 1.
config system global
    set dnsproxy-worker-count <integer>
end
                                            DNS protocols
The following DNS protocols can be enabled:
- cleartext : Enable clear text DNS over port 53 (default).
- dot : Enable DNS over TLS.
- doh : Enable DNS over HTTPS.
For more information, see DNS over TLS and HTTPS.
cache-notfound-responses
When enabled, any DNS requests that are returned with NOT FOUND can be stored in the cache. The DNS server is not asked to resolve the host name for NOT FOUND entries. By default, this option is disabled.
dns-cache-limit
Set the number of DNS entries that are stored in the cache (0 to 4294967295, default = 5000). Entries that remain in the cache provide a quicker response to requests than going out to the Internet to get the same information.
dns-cache-ttl
The duration that the DNS cache retains information, in seconds (60 to 86400 (1 day), default = 1800).
fqdn-cache-ttl
FQDN cache time to live (TTL), in seconds (0 - 86400, default = 0).
This is the amount of time an FQDN's address record can live if not refreshed. This setting applies globally, across all VDOMs, to FQDNs that have unspecified firewall address cache-ttl settings. If the cache-ttl value is configured for an FQDN address, it will supersede the fqdn-cache-ttl setting for that address.
For example, configure the FQDN cache TTL on the global VDOM:
config system dns
    set fqdn-cache-ttl 2000
end
# diagnose test application dnsproxy 6
...
vfid=0 name=test.bb.com ver=IPv4 wait_list=0 timer=985 min_ttl=1000 cache_ttl=2000 slot=-1 num=1 wildcard=0
         1.1.1.1 (ttl=1000:991:1991)
vfid=0 name=*.google.com ver=IPv4 wait_list=0 timer=0 min_ttl=0 cache_ttl=2000 slot=-1 num=0 wildcard=
...
                                            Change the cache TTL in a VDOM for a specific address:
config firewall address
    edit "test.bb.com"
        set cache-ttl 1000
    next
end
# sudo global diagnose test application dnsproxy 6
...
vfid=0 name=test.bb.com ver=IPv4 wait_list=0 timer=864 min_ttl=1000 cache_ttl=1000 slot=-1 num=1 wildcard=0
1.1.1.1 (ttl=1000:870:1870)
vfid=0 name=*.google.com ver=IPv4 wait_list=0 timer=0 min_ttl=0 cache_ttl=2000 slot=-1 num=0 wildcard=1
...
fqdn-min-refresh
FQDN cache minimum refresh time, in seconds (10 - 3600, default = 60).
An FQDN normally requeries for updates according to the lowest TTL interval returned from all the DNS records in a DNS response. The FortiGate has a default minimum refresh interval of 60 seconds; if a TTL interval is shorter than 60 seconds, it still requires a minimum of 60 seconds for the FortiGate to requery for new addresses. The fqdn-min-refresh setting changes the interval. The settings could be shortened if there are FQDNs that require fast resolutions based on a short TTL interval.
For example, if fqdn_min_refresh is unspecified:
# diagnose test application dnsproxy 3
worker idx: 0
...
FQDN: min_refresh=60 max_refresh=3600
...
# diagnose test application dnsproxy 6
worker idx: 0
vfid=0 name=aa.com ver=IPv4 wait_list=0 timer=28 min_ttl=20 cache_ttl=0 slot=-1 num=1 wildcard=0
         23.202.195.114 (ttl=20:0:0)
                                            The min_refresh is the default value of 60 seconds. Although the min_ttl (TTL returned) value is shorter, the FortiGate only requeries for updates based on the min_refresh value. the timer value is the countdown until the next refresh is triggered. The FortiGate triggers a refresh slightly earlier than the larger of the min_refresh or min_ttl value.
If fqdn_min_refresh is configured:
config system dns
    set fqdn-min-refresh 20
end
# diagnose test application dnsproxy 3
worker idx: 0
...
FQDN: min_refresh=20 max_refresh=3600
...
# diagnose test application dnsproxy 6
worker idx: 0
vfid=0 name=aa.com ver=IPv4 wait_list=0 timer=8 min_ttl=20 cache_ttl=0 slot=-1 num=1 wildcard=0
         23.202.195.114 (ttl=20:14:14)
                                            This setting can be used in combination with fqdn-cache-ttl and cache-ttl to send more frequent queries and store more resolved addresses in cache. This is useful in scenarios where the FQDN has many resolutions and changes very frequently.
fqdn-max-refresh NEW
FQDN cache maximum refresh time, in seconds (3600 - 86400, default = 3600).
The fqdn-max-refresh setting is used to control the global upper limit of the FQDN refresh timer. FQDN entries with a TTL interval that is longer than the fqdn-max-refresh value will have their refresh timer reduced to this upper limit. This allows the FortiGate to dictate the upper limit in querying for DNS updates for its FQDN addresses.
VDOM DNS
When the FortiGate is in multi-vdom mode, DNS is handled by the management VDOM. However in some cases, administrators may want to configure custom DNS settings on a non-management VDOM. For example, in a multi-tenant scenario, each VDOM might be occupied by a different tenant, and each tenant might require its own DNS server.
To configure a custom VDOM within a non-management VDOM:
config vdom edit <vdom> config system vdom-dns set vdom-dns enable set primary <primary_DNS> set secondary <secondary_DNS> set protocol {cleartext dot doh} set ip6-primary <primary_IPv6_DNS> set ip6-secondary <secondary_IPv6_DNS> set source-ip <IP_address> set interface-select-method {auto | sdwan | specify} end
