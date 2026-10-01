---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-1-cli-reference-106491833-config-system-vne-interface-67bdf28a
title: "document-fortigate-7-6-1-cli-reference-106491833-config-system-vne-interface-67bdf28a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-1-cli-reference-106491833-config-system-vne-interface-67bdf28a.md
source_anchor: ""
source_lines: [1, 44]
sha256: ce43fdb416b9c783ff79e32d6839905d94c11dcdf9e0151bd5920f819a47898a
---

# document-fortigate-7-6-1-cli-reference-106491833-config-system-vne-interface-67bdf28a

config system vne-interface
config system vne-interface
Configure virtual network enabler tunnels.
config system vne-interface
    Description: Configure virtual network enabler tunnels.
    edit <name>
        set auto-asic-offload [enable|disable]
        set bmr-hostname {password}
        set br {string}
        set http-password {password}
        set http-username {string}
        set interface {string}
        set ipv4-address {ipv4-classnet-host}
        set mode [map-e|fixed-ip|...]
        set ssl-certificate {string}
        set update-url {string}
    next
end
                                            config system vne-interface
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| auto-asic-offload * | Enable/disable tunnel ASIC offloading. | option | - | enable | 
|  |  |  |  |  | 
| bmr-hostname | BMR hostname. | password | Not Specified |  | 
| br | IPv6 address or FQDN of the border relay. | string | Maximum length: 255 |  | 
| http-password | HTTP authentication password. | password | Not Specified |  | 
| http-username | HTTP authentication user name. | string | Maximum length: 64 |  | 
| interface | Interface name. | string | Maximum length: 15 |  | 
| ipv4-address | Tunnel IPv4 address and netmask. | ipv4-classnet-host | Not Specified | 0.0.0.0 0.0.0.0 | 
| mode | VNE tunnel mode. | option | - | map-e | 
|  |  |  |  |  | 
| name | VNE tunnel name. | string | Maximum length: 15 |  | 
| ssl-certificate | Name of local certificate for SSL connections. | string | Maximum length: 35 | Fortinet_Factory | 
| update-url | URL of provisioning server. | string | Maximum length: 511 |  | 
| Option | Description | 
|---|---|
| enable | Enable auto ASIC offloading. | 
| disable | Disable ASIC offloading. | 
| Option | Description | 
|---|---|
| map-e | Map-e mode. | 
| fixed-ip | Fixed-ip mode. | 
| ds-lite | DS-Lite mode. | 
* This parameter may not exist in some models.
