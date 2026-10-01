---
id: collect-261001-huawei/huawei/dmulyalin-ttp-templates-blob-head-docs-ttp-templates-misc-n2g-cli-ip-data-huawei-ef42d805
title: "dmulyalin-ttp-templates-blob-head-docs-ttp-templates-misc-n2g-cli-ip-data-huawei-ef42d805"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/dmulyalin-ttp-templates-blob-head-docs-ttp-templates-misc-n2g-cli-ip-data-huawei-ef42d805.md
source_anchor: ""
source_lines: [1, 7]
sha256: c23d9d1b4e970697d2b3c339516e90eb60892ad4d7b720f911bed82b50925df3
---

# dmulyalin-ttp-templates-blob-head-docs-ttp-templates-misc-n2g-cli-ip-data-huawei-ef42d805

Reference path:
ttp://misc/N2G/cli_ip_data/Huawei.txt
Template to parse Huawei interfaces configuration and ARP cache.
Template Content
``` Template to parse Huawei interfaces configuration and ARP cache. commands = [ "display current-configuration interface", "display arp all", ] kwargs = {"strip_prompt": False} method = "send_command" platform = ["huawei", "huawei_vrpv8"]
local_hostname="gethostname"
interface {{ interface | resuball("IfsNormalize") }} description {{ port_description | re(".+") }} ip binding vpn-instance {{ vrf }} ip address {{ ip | IP }} {{ netmask }} ip address {{ ip | IP }} {{ netmask }} sub ipv6 address {{ ip | IPV6 }}/{{ netmask }} ipv6 address {{ ip | IPV6 }}/{{ netmask }} sub vrrp vrid {{ group | let("type", "VRRP") }} virtual-ip {{ ip | IP }} {{ ip | IP }} {{ mac | MAC | mac_eui }} {{ expire | re("\\d+") }} {{ type | notdigit }} {{ interface | resuball("IfsNormalize") }} {{ vrf }} {{ ip | IP }} {{ mac | MAC | mac_eui }} {{ type | notdigit }} {{ interface | resuball("IfsNormalize") }} {{ vrf }} {{ ip | IP }} {{ mac | MAC | mac_eui }} {{ expire | re("\\d+") }} {{ type | notdigit }} {{ interface | resuball("IfsNormalize") }} {{ ip | IP }} {{ mac | MAC | mac_eui }} {{ type | notdigit }} {{ interface | resuball("IfsNormalize") }} ```
