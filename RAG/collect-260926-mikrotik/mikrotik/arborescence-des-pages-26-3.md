---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-3
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [203, 295]
sha256: 5db250d9cd8a7553b995a361c76399e95949a855b2e81401f6be43210bcf8962
---

# Overview

There are 2 types of interfaces on CAPsMAN - "master" and "slave". The master interface holds the configuration for an actual wireless interface (radio), while a slave interface links to the master interface and is intended to hold the configuration for a Virtual-AP (multiple SSID support). There are settings that are meaningful only for master interface, i.e. mainly hardware setup related settings such as radio channel settings. Note that in order for a radio to accept clients, it's master interface needs to be enabled. Slave interfaces will become operational only if enabled and the master interface is enabled.

Interfaces on CAPsMAN can be static or dynamic. Static interfaces are stored in RouterOS configuration and will persist across reboots. Dynamic interfaces exist only while a particular CAP is connected to CAPsMAN.

## **CAPsMAN Global Configuration**

Settings to enable CAPsMAN functionality are found in **/caps-man manager** menu:

| Property | Description | 
|---|---|
| **enabled** (*yes \| no* ; Default:**no** ) | Disable or enable CAPsMAN functionality | 
| **certificate** (*auto \| certificate name \| none* ; Default:**none** ) | Device certificate | 
| **ca-certificate** (*auto \| certificate name \| none* ; Default:**none** ) | Device CA certificate | 
| **require-peer-certificate** (*yes \| no* ; Default:**no** ) | Require all connecting CAPs to have a valid certificate | 
| **package-path** (*string \|* ; Default: ) | Folder location for the RouterOS packages. For example, use "/upgrade" to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded. | 
| **upgrade-policy** (*none \| require-same-version \| suggest-same-upgrade* ; Default:**none** ) | Upgrade policy options  | 

## **Radio Provisioning**

CAPsMAN distinguishes between CAPs based on an identifier. The identifier is generated based on the following rules:

- if CAP provided a certificate, identifier is set to the Common Name field in the certificate
- otherwise identifier is based on Base-MAC provided by CAP in the form: '[XX:XX:XX:XX:XX:XX]'.

When the DTLS connection with CAP is successfully established (which means that CAP identifier is known and valid), CAPsMAN makes sure there is no stale connection with CAP using the same identifier. Currently connected CAPs are listed in **/caps-man remote-cap** menu:

[admin@CM] /caps-man> remote-cap print
 # ADDRESS                                    IDENT           STATE               RADIOS
 0 00:0C:42:00:C0:32/27044                    MT-000C4200C032 Run                      1

CAPsMAN distinguishes between actual wireless interfaces (radios) based on their builtin MAC address (radio-mac). This implies that it is impossible to manage two radios with the same MAC address on one CAPsMAN. Radios currently managed by CAPsMAN (provided by connected CAPs) are listed in **/caps-man radio** menu:

[admin@CM] /caps-man> radio print
Flags: L - local, P - provisioned 
 #    RADIO-MAC         INTERFACE                               REMOTE-AP-IDENT
 0  P 00:03:7F:48:CC:07 cap1                                    MT-000C4200C032

When CAP connects, CAPsMAN at first tries to bind each CAP radio to CAPsMAN master interface based on radio-mac. If an appropriate interface is found, radio gets set up using master interface configuration and configuration of slave interfaces that refer to particular master interface. At this moment interfaces (both master and slaves) are considered bound to radio and radio is considered provisioned.

If no matching master interface for radio is found, CAPsMAN executes 'provisioning rules'. Provisioning rules is an ordered list of rules that contain settings that specify which radio to match and settings that specify what action to take if a radio matches.


Provisioning rules for matching radios are configured in **/caps-man provisioning** menu:

| Property | Description | 
|---|---|
| **action** (*create-disabled \| create-enabled \| create-dynamic-enabled \| none* ; Default:**none** ) | Action to take if rule matches are specified by the following settings:  | 
| **comment** (*string* ; Default: ) | Short description of the Provisioning rule | 
| **common-name-regexp** (*string* ; Default: ) | Regular expression to match radios by common name | 
| **hw-supported-modes** (*a\|a-turbo\|ac\|an\|b\|g\|g-turbo\|gn* ; Default: ) | Match radios by supported wireless modes | 
| **identity-regexp** (*string* ; Default: ) | Regular expression to match radios by router identity | 
| **ip-address-ranges** (*IpAddressRange[,IpAddressRanges] max 100x* ; Default:**""** ) | Match CAPs with IPs within configured address range. | 
| **master-configuration** (*string* ; Default: ) | If **action** specifies to create interfaces, then a new master interface with its configuration set to this configuration profile will be created | 
| **name-format** (*cap \| identity \| prefix \| prefix-identity* ; Default:**cap** ) | specify the syntax of the CAP interface name creation  | 
| **name-prefix** (*string* ; Default: ) | name prefix which can be used in the name-format for creating the CAP interface names | 
| **radio-mac** (*MAC address* ; Default:**00:00:00:00:00:00** ) | MAC address of radio to be matched, empty MAC (00:00:00:00:00:00) means match all MAC addresses | 
| **slave-configurations** (*string* ; Default: ) | If **action** specifies to create interfaces, then a new slave interface for each configuration profile in this list is created. | 

 If no rule matches radio, then implicit default rule with action **create-enabled** and no configurations set is executed.

To get the active provisioning matchers:

```
[admin@CM] /caps-man provisioning> print
Flags: X - disabled 
 0   radio-mac=00:00:00:00:00:00 action=create-enabled master-configuration=main-cfg 
     slave-configurations=virtual-ap-cfg name-prefix=""
```
For user's convenience there are commands that allow the re-execution of the provisioning process for some radio or all radios provided by some AP:

[admin@CM] > caps-man radio provision 0

and

[admin@CM] > caps-man remote-cap provision 0

## **Interface Configuration**

CAPsMAN interfaces are managed in **/caps-man interface** menu:

[admin@CM] > /caps-man interface print          
Flags: M - master, D - dynamic, B - bound, X - disabled, I - inactive, R - running 
 #      NAME                                 RADIO-MAC         MASTER-INTERFACE                               
 0 M BR cap2                                 00:0C:42:1B:4E:F5 none                                           
 1   B  cap3                                 00:00:00:00:00:00 cap2                   

## **Master Configuration Profiles**

Configuration profiles permit pre-defined 'top level' master settings to be applied to CAP radios being provisioned.


Configuration Profiles are configured in **/caps-man configuration** menu:

