---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d-3
title: "document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d.md
source_anchor: ""
source_lines: [261, 299]
sha256: 181298fe893c492d1a99169a1db7a415e48a02861fd9eff6225d25c117f59552
---

# document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d

                set service ALL
                set nat enable
            next
        end
    next
end
- 
                                                    Configure the firewall policies from SalesLocal to the Internet: config vdom
    edit Sales
        config firewall policy
            edit 3
                set name "Sales-local-to-Management"
                set srcintf port3
                set dstintf SalesVlnk0
                set srcaddr all
                set dstaddr all
                set action accept
                set schedule always
                set service ALL
                set nat enable
            next
        end
    next
    edit root
        config firewall policy
            edit 4
                set name "Sales-VDOM-to-Internet"
                set srcintf SalesVlnk1
                set dstintf port1
                set srcaddr all
                set dstaddr all
                set action accept
                set schedule always
                set service ALL
                set nat enable
            next
        end
    next
end
