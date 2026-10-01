---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-3
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [205, 351]
sha256: f41972e1d2ce529ee1ad5695064ca2e75dc20ec291b5ecd71d066801b07e48b4
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

|---|---|---|---|
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| WARNINGCONNECTIONSASSOC | Threshold |  |  | 
| CRITICALCONNECTIONSASSOC | Threshold |  |  | 
| WARNINGCONNECTIONSAUTH | Threshold |  |  | 
| CRITICALCONNECTIONSAUTH | Threshold |  |  | 
| WARNINGCONNECTIONSDHCP | Threshold |  |  | 
| CRITICALCONNECTIONSDHCP | Threshold |  |  | 
| WARNINGCONNECTIONSDNS | Threshold |  |  | 
| CRITICALCONNECTIONSDNS | Threshold |  |  | 
| WARNINGCONNECTIONSSUCCESS | Threshold |  |  | 
| CRITICALCONNECTIONSSUCCESS | Threshold |  |  | 
| WARNINGLINKLATENCY | Threshold in milliseconds |  |  | 
| CRITICALLINKLATENCY | Threshold in milliseconds |  |  | 
| WARNINGLINKLOSS | Threshold in percentage |  |  | 
| CRITICALLINKLOSS | Threshold in percentage |  |  | 
| WARNINGLINKSTATUS | Threshold |  |  | 
| CRITICALLINKSTATUS | Threshold |  |  | 
| WARNINGLOAD | Threshold |  |  | 
| CRITICALLOAD | Threshold |  |  | 
| WARNINGPORTSTATUS | Threshold |  |  | 
| CRITICALPORTSTATUS | Threshold |  |  | 
| WARNINGPORTTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALPORTTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGPORTTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALPORTTRAFFICOUT | Threshold in b/s |  |  | 
| WARNINGSTATUS | Threshold |  |  | 
| CRITICALSTATUS | Threshold |  |  | 
| WARNINGTOTALALERTING | Threshold |  |  | 
| CRITICALTOTALALERTING | Threshold |  |  | 
| WARNINGTOTALOFFLINE | Threshold |  |  | 
| CRITICALTOTALOFFLINE | Threshold |  |  | 
| WARNINGTOTALOFFLINEPRCT | Threshold in percentage |  |  | 
| CRITICALTOTALOFFLINEPRCT | Threshold in percentage |  |  | 
| WARNINGTOTALONLINE | Threshold |  |  | 
| CRITICALTOTALONLINE | Threshold |  |  | 
| WARNINGTOTALONLINEPRCT | Threshold in percentage |  |  | 
| CRITICALTOTALONLINEPRCT | Threshold in percentage |  |  | 
| WARNINGTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALTRAFFICOUT | Threshold in b/s |  |  | 
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| FILTERDEVICENAME | Filter devices by name (can be a regexp) |  |  | 
| FILTERLINKNAME | Filter VPN links by name (can be a regexp) |  |  | 
| FILTERNETWORKID | Filter devices by network ID (can be a regexp) |  |  | 
| FILTERTAGS | Filter devices by tags (can be a regexp) |  |  | 
| FILTERORGANIZATIONNAME | Filter devices by organization name (can be a regexp) |  |  | 
| FILTERORGANIZATIONID | Filter devices by organization ID (can be a regexp) |  |  | 
| WARNINGCONNECTIONSASSOC | Threshold |  |  | 
| CRITICALCONNECTIONSASSOC | Threshold |  |  | 
| WARNINGCONNECTIONSAUTH | Threshold |  |  | 
| CRITICALCONNECTIONSAUTH | Threshold |  |  | 
| WARNINGCONNECTIONSDHCP | Threshold |  |  | 
| CRITICALCONNECTIONSDHCP | Threshold |  |  | 
| WARNINGCONNECTIONSDNS | Threshold |  |  | 
| CRITICALCONNECTIONSDNS | Threshold |  |  | 
| WARNINGCONNECTIONSSUCCESS | Threshold |  |  | 
| CRITICALCONNECTIONSSUCCESS | Threshold |  |  | 
| WARNINGLINKLATENCY | Threshold in milliseconds |  |  | 
| CRITICALLINKLATENCY | Threshold in milliseconds |  |  | 
| WARNINGLINKLOSS | Threshold in percentage |  |  | 
| CRITICALLINKLOSS | Threshold in percentage |  |  | 
| WARNINGLINKSINEFFECTIVE | Threshold |  |  | 
| CRITICALLINKSINEFFECTIVE | Threshold |  |  | 
| CRITICALLINKSTATUS | Threshold | %{link_status} =~ /failed/i |  | 
| WARNINGLINKSTATUS | Threshold |  |  | 
| WARNINGLOAD | Threshold |  |  | 
| CRITICALLOAD | Threshold |  |  | 
| WARNINGPORTSTATUS | Threshold |  |  | 
| CRITICALPORTSTATUS | Threshold |  |  | 
| WARNINGPORTTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALPORTTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGPORTTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALPORTTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALSTATUS | Threshold | %{status} =~ /alerting/i |  | 
| WARNINGSTATUS | Threshold |  |  | 
| WARNINGTOTALALERTING | Threshold |  |  | 
| CRITICALTOTALALERTING | Threshold |  |  | 
| WARNINGTOTALOFFLINE | Threshold |  |  | 
| CRITICALTOTALOFFLINE | Threshold |  |  | 
| WARNINGTOTALOFFLINEPRCT | Threshold in percentage |  |  | 
| CRITICALTOTALOFFLINEPRCT | Threshold in percentage |  |  | 
| WARNINGTOTALONLINE | Threshold |  |  | 
| CRITICALTOTALONLINE | Threshold |  |  | 
| WARNINGTOTALONLINEPRCT | Threshold in percentage |  |  | 
| CRITICALTOTALONLINEPRCT | Threshold in percentage |  |  | 
| WARNINGTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALTRAFFICOUT | Threshold in b/s |  |  | 
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. | --verbose |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| WARNINGCONNECTIONSASSOC | Threshold |  |  | 
| CRITICALCONNECTIONSASSOC | Threshold |  |  | 
| WARNINGCONNECTIONSAUTH | Threshold |  |  | 
| CRITICALCONNECTIONSAUTH | Threshold |  |  | 
| WARNINGCONNECTIONSDHCP | Threshold |  |  | 
| CRITICALCONNECTIONSDHCP | Threshold |  |  | 
| WARNINGCONNECTIONSDNS | Threshold |  |  | 
| CRITICALCONNECTIONSDNS | Threshold |  |  | 
| WARNINGCONNECTIONSSUCCESS | Threshold |  |  | 
| CRITICALCONNECTIONSSUCCESS | Threshold |  |  | 
| WARNINGTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALTRAFFICOUT | Threshold in b/s |  |  | 
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| FILTERNETWORKNAME | Filter network name (can be a regexp) |  |  | 
| FILTERORGANIZATIONNAME | Filter networks by organization name (can be a regexp) |  |  | 
| FILTERORGANIZATIONID | Filter networks by organization ID (can be a regexp) |  |  | 
| WARNINGCONNECTIONSASSOC | Threshold |  |  | 
| CRITICALCONNECTIONSASSOC | Threshold |  |  | 
| WARNINGCONNECTIONSAUTH | Threshold |  |  | 
| CRITICALCONNECTIONSAUTH | Threshold |  |  | 
| WARNINGCONNECTIONSDHCP | Threshold |  |  | 
| CRITICALCONNECTIONSDHCP | Threshold |  |  | 
| WARNINGCONNECTIONSDNS | Threshold |  |  | 
| CRITICALCONNECTIONSDNS | Threshold |  |  | 
| WARNINGCONNECTIONSSUCCESS | Threshold |  |  | 
| CRITICALCONNECTIONSSUCCESS | Threshold |  |  | 
| WARNINGTRAFFICIN | Threshold in b/s |  |  | 
| CRITICALTRAFFICIN | Threshold in b/s |  |  | 
| WARNINGTRAFFICOUT | Threshold in b/s |  |  | 
| CRITICALTRAFFICOUT | Threshold in b/s |  |  | 
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. | --verbose |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| FILTERNETWORKNAME | Filter VPN tunnels by network name (can be a regexp) |  |  | 
| FILTERORGANIZATIONID | Filter VPN tunnels by organization ID (can be a regexp) |  |  | 
| FILTERORGANIZATIONNAME | Filter VPN tunnels by organization name (can be a regexp) |  |  | 
| FILTERDEVICESERIAL | Filter VPN tunnels by device serial (can be a regexp) |  |  | 
| FILTERVPNNAME | Filter VPN tunnels by VPN name (can be a regexp) |  |  | 
| FILTERVPNTYPE | Filter VPN tunnels by VPN type (can be a regexp) |  |  | 
| WARNINGDEVICESTATUS | Threshold |  |  | 
| CRITICALDEVICESTATUS | Threshold |  |  | 
| WARNINGTOTALUNREACHABLE | Threshold |  |  | 
| CRITICALTOTALUNREACHABLE | Threshold |  |  | 
| WARNINGVPNSTATUS | Threshold |  |  | 
| CRITICALVPNSTATUS | Threshold |  |  | 
