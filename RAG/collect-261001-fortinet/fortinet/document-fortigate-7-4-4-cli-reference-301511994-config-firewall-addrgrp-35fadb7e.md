---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-cli-reference-301511994-config-firewall-addrgrp-35fadb7e
title: "config firewall addrgrp"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-cli-reference-301511994-config-firewall-addrgrp-35fadb7e.md
source_anchor: ""
source_lines: [1, 84]
sha256: a0ed882f2907ba067abd70f7209b0de4ae26f5a1d7372e8761f857d989b6c6ad
---

# config firewall addrgrp

# config firewall addrgrp

Configure IPv4 address groups.

```
config firewall addrgrp
    Description: Configure IPv4 address groups.
    edit <name>
        set allow-routing [enable|disable]
        set category [default|ztna-ems-tag|...]
        set color {integer}
        set comment {var-string}
        set exclude [enable|disable]
        set exclude-member <name1>, <name2>, ...
        set fabric-object [enable|disable]
        set member <name1>, <name2>, ...
        config tagging
            Description: Config object tagging.
            edit <name>
                set category {string}
                set tags <name1>, <name2>, ...
            next
        end
        set type [default|folder]
        set uuid {uuid}
    next
end
```
                                            ## config firewall addrgrp

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| allow-routing | Enable/disable use of this group in the static route configuration. | option | - | disable | 
|  |  |  |  |  | 
| category | Address group category. | option | - | default | 
|  |  |  |  |  | 
| color | Color of icon on the GUI. | integer | Minimum value: 0 Maximum value: 32 | 0 | 
| comment | Comment. | var-string | Maximum length: 255 |  | 
| exclude | Enable/disable address exclusion. | option | - | disable | 
|  |  |  |  |  | 
| exclude-member `<name>` | Address exclusion member. Address name. | string | Maximum length: 79 |  | 
| fabric-object | Security Fabric global object setting. | option | - | disable | 
|  |  |  |  |  | 
| member `<name>` | Address objects contained within the group. Address name. | string | Maximum length: 79 |  | 
| name | Address group name. | string | Maximum length: 79 |  | 
| type | Address group type. | option | - | default | 
|  |  |  |  |  | 
| uuid | Universally Unique Identifier (UUID; automatically assigned but can be manually reset). | uuid | Not Specified | 00000000-0000-0000-0000-000000000000 | 

| Option | Description | 
|---|---|
| *enable* | Enable use of this group in the static route configuration. | 
| *disable* | Disable use of this group in the static route configuration. | 

| Option | Description | 
|---|---|
| *default* | Default address group category (cannot be used as ztna-ems-tag/ztna-geo-tag in policy). | 
| *ztna-ems-tag* | Members must be ztna-ems-tag group or ems-tag address, can be used as ztna-ems-tag in policy. | 
| *ztna-geo-tag* | Members must be ztna-geo-tag group or geographic address, can be used as ztna-geo-tag in policy. | 

| Option | Description | 
|---|---|
| *enable* | Enable address exclusion. | 
| *disable* | Disable address exclusion. | 

| Option | Description | 
|---|---|
| *enable* | Object is set as a security fabric-wide global object. | 
| *disable* | Object is local to this security fabric member. | 

| Option | Description | 
|---|---|
| *default* | Default address group type (address may belong to multiple groups). | 
| *folder* | Address folder group (members may not belong to any other group). | 

### config tagging

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| name | Tagging entry name. | string | Maximum length: 63 |  | 
| category | Tag category. | string | Maximum length: 63 |  | 
| tags `<name>` | Tags. Tag name. | string | Maximum length: 79 |  |
