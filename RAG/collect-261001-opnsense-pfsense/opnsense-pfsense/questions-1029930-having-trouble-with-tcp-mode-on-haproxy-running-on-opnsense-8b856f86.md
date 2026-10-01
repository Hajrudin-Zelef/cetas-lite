---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1029930-having-trouble-with-tcp-mode-on-haproxy-running-on-opnsense-8b856f86
title: "Having trouble with tcp mode on HAProxy running on OPNsense"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1029930-having-trouble-with-tcp-mode-on-haproxy-running-on-opnsense-8b856f86.md
source_anchor: ""
source_lines: [1, 67]
sha256: 0a0676cf8f95d3307c4657bb070c64ed879f182fbf92b6bb4542c70fd841cad3
---

# Having trouble with tcp mode on HAProxy running on OPNsense

*Score : 0 | Source : https://serverfault.com/questions/1029930/having-trouble-with-tcp-mode-on-haproxy-running-on-opnsense*

I'm trying to setup HAProxy to do simple tcp passthrough to different backends based on SNI. However, I have serious issues on my clients if I do so:
If I contact the backends directly, bypassing HAProxy, I have no such issues. So something in HAProxy is not configured right but I fail to see what it is.
Technical Details:
My config looks like this:
global
    # NOTE: Could be a security issue, but required for some feature.
    uid                         80
    gid                         80
    chroot                      /var/haproxy
    daemon
    stats                       socket /var/run/haproxy.socket group proxy mode 775 level admin
    nbproc                      1
    nbthread                    1
    tune.ssl.default-dh-param   1024
    spread-checks               0
    tune.chksize                16384
    tune.bufsize                16384
    tune.lua.maxmem             0
    log /var/run/log local0 info
defaults
    log     global
    option redispatch -1
    timeout client 30s
    timeout connect 30s
    timeout server 30s
    retries 3
# Frontend: https_passthrough (https tcp mode main ip)
frontend https_passthrough
    bind x.x.x.x:443 name x.x.x.x:443
    mode tcp
    # tuning options
    timeout client 30s
    # logging options
    # ACL: service_sni
    acl acl_5f34555b7fae76.76760871 req.ssl_sni -i service.example.com
    # ACTION: service_https_passthrough
    use_backend service_https_passthrough if acl_5f34555b7fae76.76760871
    # WARNING: pass through options below this line
    tcp-request inspect-delay 10s
# Backend: service_https_passthrough ()
backend service_https_passthrough
    # health checking is DISABLED
    mode tcp
    balance source
    # stickiness
    stick-table type ip size 50k expire 30m
    stick on src
    # tuning options
    timeout connect 30s
    timeout server 30s
    server service_https_passthrough 10.0.1.105:443
If the connection fails I find entries like
https_passthrough https_passthrough/<NOSRV> -1/-1/0 0 SC 1/1/0/0/0 0/0
in the logs. But as I said, the backend servers are up and running 24/7 and work fine all day from internal LAN (also routing/firewall/VLANs involved here) so I doubt it is a firewall issue.
Any help appreciated.

---

### Reponse — score 2

Looks like adding these two lines to the frontend config fixed the issue for me:
tcp-request inspect-delay 5s tcp-request content accept if { req_ssl_hello_type 1 }
Derrived from a HAProxy blog post.
