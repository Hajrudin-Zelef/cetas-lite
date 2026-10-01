---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-576158-configuring-firewall-authen-52aa3be9-2
title: "Configuring FSSO firewall authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-576158-configuring-firewall-authen-52aa3be9.md
source_anchor: ""
source_lines: [203, 258]
sha256: 7a309db11208b4a466571d07880a371b84bfd1ea8296c777a7a2bcc4cde0c841
---

# Configuring FSSO firewall authentication

1. 
                                                    Go to *Policy & Objects > Addresses* and select*Address* .
2. 
                                                    Click *Create New* .
3. 
                                                    Configure the following settings: Name Internal_net Type Subnet IP/Netmask 10.11.102.0/24 Interface Port 3
4. 
                                                    Click *OK* .
5. 
                                                    Create another new address by repeating steps 2-4 using the following settings: Name Windows_net Type Subnet IP/Netmask 10.11.101.0/24 Interface Port 2

You must create two security policies: one for the firewall group connecting through port 3, and one for the FSSO group connecting through port 2.

###### To create security policies using the GUI:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* .
2. 
                                                    Click *Create New* .
3. 
                                                    Configure the following settings: Incoming Interface Port2 Source Address Windows_net Source User(s) FSSO_Internet_users Outgoing Interface Port1 Destination Address all Schedule always Service ALL NAT Enabled. Security Profiles You can enable security profiles as desired.
4. 
                                                    Click *OK* .
5. 
                                                    Create another new policy by repeating steps 2-4 using the following settings: Incoming Interface Port3 Source Address Internal_net Source User(s) Internet_users Outgoing Interface Port1 Destination Address all Schedule always Service ALL NAT Enabled. Security Profiles You can enable security profiles as desired.
6. 
                                                    Click *OK* .

###### To create security policies using the CLI:

```
config firewall policy
    edit 0
        set srcintf port2
        set dstintf port1
        set srcaddr Windows_net
        set dstaddr all
        set action accept
        set groups FSSO_Internet_users
        set schedule always
        set service ANY
        set nat enable
    next
    edit 1
        set srcintf port3
        set dstintf port1
        set srcaddr internal_net
        set dstaddr all
        set action accept
        set schedule always
        set groups Internet_users
        set service ANY
        set nat enable
    next
end
```
