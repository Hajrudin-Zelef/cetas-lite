---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c.md
source_anchor: ""
source_lines: [273, 498]
sha256: 236595025f692006012ccbe5c3c98e21808d952580dd49ab557ceb62b3ec33d2
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c

Specifies a different interface and enters interface configuration mode.
Step 9
ip addressip-address mask[secondary]
Example:
Switch(config-if)# ip address 172.31.232.182 255.255.255.240
Sets a primary IP address for an interface.
Step 10
ip nat outside
Example:
Switch(config-if)# ip nat outside
Connects the interface to the outside network.
Step 11
end
Example:
Switch(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
Configuring Dynamic Translation of Inside Source Addresses
Dynamic translation establishes a mapping between an inside local address and a pool of global addresses dynamically. Dynamic
translation can be enabled by configuring a dynamic NAT rule and the mapping is established based on the result of the evaluation
of the configured rule at run-time. You can employ an ACL to specify the inside local address and the inside global address
can be specified through an address pool or an interface.
Dynamic translation is useful when multiple users on a private network need to access the Internet. The dynamically configured
pool IP address may be used as needed and is released for use by other users when access to the internet is no longer required.
SUMMARY STEPS
enable
configure terminal
ip nat poolname start-ip end-ip netmasknetmask | prefix-length prefix-length
Defines a standard access list permitting those addresses that are to be translated.
Step 5
ip nat inside source listaccess-list-numberpoolname
Example:
Switch(config)# ip nat inside source list 1 pool net-208
Establishes dynamic source translation, specifying the access list defined in Step 4.
Step 6
interface type number
Example:
Switch(config)# interface ethernet 1
Specifies an interface and enters interface configuration mode.
Step 7
ip addressip-address mask
Example:
Switch(config-if)# ip address 10.114.11.39 255.255.255.0
Sets a primary IP address for the interface.
Step 8
ip nat inside
Example:
Switch(config-if)# ip nat inside
Connects the interface to the inside network, which is subject to NAT.
Step 9
exit
Example:
Switch(config-if)#exit
Exits the interface configuration mode and returns to global configuration mode.
Step 10
interface type number
Example:
Switch(config)# interface ethernet 0
Specifies an interface and enters interface configuration mode.
Step 11
ip addressip-address mask
Example:
Switch(config-if)# ip address 172.16.232.182 255.255.255.240
Sets a primary IP address for the interface.
Step 12
ip nat outside
Example:
Switch(config-if)# ip nat outside
Connects the interface to the outside network.
Step 13
end
Example:
Switch(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
Configuring PAT
Perform this task to allow your internal users access to the Internet and conserve addresses in the inside global address
pool using overloading of global addresses.
SUMMARY STEPS
enable
configure terminal
ip nat poolname start-ip end-ip netmasknetmask | prefix-length prefix-length
Defines a standard access list permitting those addresses that are to be translated.
The access list must permit only those addresses that are to be translated. (Remember that there is an implicit “deny all”
at the end of each access list.) Use of an access list that is too permissive can lead to unpredictable results.
Step 5
ip nat inside source listaccess-list-numberpool name overload
Example:
Switch(config)# ip nat inside source list 1 pool net-208 overload
Establishes dynamic source translation with overloading, specifying the access list defined in Step 4.
Step 6
interface type number
Example:
Switch(config)# interface ethernet 1
Specifies an interface and enters interface configuration mode.
Step 7
ip addressip-address mask[secondary]
Example:
Switch(config-if)# ip address 192.168.201.1 255.255.255.240
Sets a primary IP address for an interface.
Step 8
ip nat inside
Example:
Switch(config-if)# ip nat inside
Connects the interface to the inside network, which is subject to NAT.
Step 9
exit
Example:
Switch(config-if)# exit
Exits interface configuration mode and returns to global configuration mode.
Step 10
interface type number
Example:
Switch(config)# interface ethernet 0
Specifies a different interface and enters interface configuration mode.
Step 11
ip addressip-address mask[secondary]
Example:
Switch(config-if)# ip address 192.168.201.29 255.255.255.240
Sets a primary IP address for an interface.
Step 12
ip nat outside
Example:
Switch(config-if)# ip nat outside
Connects the interface to the outside network.
Step 13
end
Example:
Switch(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
Configuring NAT of External IP Addresses Only
By default, NAT translates the addresses embedded in the packet pay-load as explained in Using Application-Level Gateways with NAT section. There might be situations where the translation of the embedded address is not desirable and in such cases, NAT
can be configured to translate the external IP address only.
Device(config)# ip nat outside source static network 10.1.1.1 192.168.251.0/24 no-payload
Disables network packet translation on the outside host device.
Step 9
exit
Example:
Device(config)# exit
Exits global configuration mode and returns to privileged EXEC mode.
Step 10
showipnattranslations [verbose]
Example:
Device# show ip nat translations
Displays active NAT.
Configuring Translation of Overlapping Networks
Configure static translation of overlapping networks if your IP addresses in the stub network are legitimate IP addresses
belonging to another network and you want to communicate with those hosts or routers using static translation.
SUMMARY STEPS
enable
configure terminal
ip nat inside source static local-ip global-ip
ip nat outside source static local-ip global-ip
interface type number
ip addressip-address mask
ip nat inside
exit
interface type number
ip addressip-address mask
ip nat outside
end
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Switch> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configure terminal
Example:
Switch# configure terminal
Step 3
ip nat inside source static local-ip global-ip
Example:
Switch(config)# ip nat inside source static 10.1.1.1 203.0.113.2
Establishes static translation between an inside local address and an inside global address.
Step 4
ip nat outside source static local-ip global-ip
Example:
Switch(config)# ip nat outside source static 172.16.0.3 10.1.1.3
Establishes static translation between an outside local address and an outside global address.
Step 5
interface type number
Example:
Switch(config)# interface ethernet 1
Specifies an interface and enters interface configuration mode.
Step 6
ip addressip-address mask
Example:
Switch(config-if)# ip address 10.114.11.39 255.255.255.0
Sets a primary IP address for an interface.
Step 7
ip nat inside
Example:
Switch(config-if)# ip nat inside
Marks the interface as connected to the inside.
Step 8
exit
Example:
Switch(config-if)# exit
Exits interface configuration mode and returns to global configuration mode.
Step 9
interface type number
Example:
Switch(config)# interface ethernet 0
Specifies a different interface and enters interface configuration mode.
Step 10
ip addressip-address mask
Example:
Switch(config-if)# ip address 172.16.232.182 255.255.255.240
Sets a primary IP address for an interface.
Step 11
ip nat outside
Example:
Switch(config-if)# ip nat outside
Marks the interface as connected to the outside.
Step 12
end
Example:
Switch(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
Configuring Address Translation Timeouts
You can configure address translation timeouts based on your NAT configuration.
By default, dynamically created translation entries time-out after a period of inactivity to enable the efficient use of various
resources. You can change the default values on timeouts, if necessary. The following are the default time-out configurations
