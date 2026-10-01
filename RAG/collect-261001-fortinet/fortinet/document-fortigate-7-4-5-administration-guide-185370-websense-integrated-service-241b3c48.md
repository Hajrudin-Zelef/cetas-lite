---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-administration-guide-185370-websense-integrated-service-241b3c48
title: "Websense Integrated Services Protocol"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-administration-guide-185370-websense-integrated-service-241b3c48.md
source_anchor: ""
source_lines: [1, 59]
sha256: c66f5bc63600dfb43165fea01fea0f16b28a7837b8d43557003d32869208fa62
---

# Websense Integrated Services Protocol

# Websense Integrated Services Protocol

Websense Integrated Services Protocol (WISP) is supported on the FortGate, which allows the firewall to send traffic to the third-party web filtering service for rating and approval checking.

When WISP is enabled, the FortiGate maintains a pool of TCP connections to the WISP server. The TCP connections are used to forward HTTP request information and log information to the WISP server and receive policy decisions.

When a WISP server is used in a web filter profile, in flow or proxy mode, the following web filter scanning priority sequence is used:

1. 
                                                    Local URL filter
2. 
                                                    Websense web filtering service
3. 
                                                    FortiGuard web filtering service

The following example uses a WISP server configured in a flow mode web filter profile.

###### To use a WISP server in flow mode:

1. 
                                                    Configure the WISP servers: ```
config web-proxy wisp
    edit "wisp1"
        set server-ip 10.2.3.4
    next
    edit "wisp2"
        set server-ip 10.2.3.5
    next
    edit "wisp3"
        set server-ip 192.168.1.2
    next
    edit "wisp4"
        set server-ip 192.168.3.4
    next
end
```
2. 
                                                    Configure the web filter profile: ```
config webfilter profile
    edit "webfilter_flowbase"
        set feature-set flow
        config ftgd-wf
            unset options
            config filters
                edit 64
                    set category 64
                    set action block
                next
            end
        end
        set wisp enable
        set wisp-servers "wisp1" "wisp2"
        set wisp-algorithm {primary-secondary | round-robin | auto-learning}
        set log-all-url enable
    next
end
```
