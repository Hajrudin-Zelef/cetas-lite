---
id: collect-261001-fortinet/fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5-5
title: "poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5.md
source_anchor: ""
source_lines: [601, 613]
sha256: 355e8f547d9d03f7619f618e41bab428b462fe60b08791ac422521a1eeb2a116
---

# poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5

- vrdst_priority - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
- vrgrp - VRRP group ID (1 - 65535).
- vrid - Virtual router identifier (1 - 255).
- vrip - IP address of the virtual router.
- proxy_arp - VRRP Proxy ARP configuration. The structure ofproxy_arp block is documented below.
The proxy_arp block contains:
- id - ID.
- ip - Set IP addresses of proxy ARP.
In addition to all the above arguments, the following attributes are exported:
- id - an identifier for the resource with format {{mkey}}.
Check out allow_append to auto import upon resource creation.
fortios_system_interface can be imported using:
terraform import fortios_system_interface.labelname {{mkey}}
