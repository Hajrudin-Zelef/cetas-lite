---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56.md
source_anchor: ""
source_lines: [47, 123]
sha256: 25bf940dab2e2ce47e2bf95d1e24c22e65806dc2ffc82f8c17a412ef3e7c2bd7
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56

| protect | No | No | No | No | No | No | 
| restrict | No | Yes | Yes | No | Yes | No | 
| shutdown | No | No | No | No | Yes | Yes | 
| shutdown vlan | No | No | Yes | No | Yes | No 3 | 
You can use port security aging to set the aging time for all secure addresses on a port. Two types of aging are supported per port:
Absolute—The secure addresses on the port are deleted after the specified aging time.
Inactivity—The secure addresses on the port are deleted only if the secure addresses are inactive for the specified aging time.
When a switch joins a stack, the new switch will get the configured secure addresses. All dynamic secure addresses are downloaded by the new stack member from the other stack members.
When a switch (either the active switch or a stack member) leaves the stack, the remaining stack members are notified, and the secure MAC addresses configured or learned by that switch are deleted from the secure MAC address table.
| Table 3. Default Port Security Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Port security | Disabled on a port. | 
| Sticky address learning | Disabled. | 
| Maximum number of secure MAC addresses per port | One address | 
| Violation mode | Shutdown. The port shuts down when the maximum number of secure MAC addresses is exceeded. | 
| Port security aging | Disabled. Aging time is 0. Static aging is disabled. Type is absolute. | 
The following guidelines are applicable during port security configuration:
Port security can only be configured on static access ports or trunk ports. A secure port cannot be a dynamic access port.
A secure port cannot be a destination port for Switched Port Analyzer (SPAN).
Voice VLAN is only supported on access ports and not on trunk ports, even though the configuration is allowed.
When you enable port security on an interface that is also configured with a voice VLAN, set the maximum allowed secure addresses on the port to two. When the port is connected to a Cisco IP phone, the IP phone requires one MAC address. The Cisco IP phone address is learned on the voice VLAN, but is not learned on the access VLAN. If you connect a single PC to the Cisco IP phone, no additional MAC addresses are required. If you connect more than one PC to the Cisco IP phone, you must configure enough secure addresses to allow one for each PC and one for the phone.
When a trunk port configured with port security and assigned to an access VLAN for data traffic and to a voice VLAN for voice traffic, entering the switchport voice and switchport priority extend interface configuration commands has no effect.
When a connected device uses the same MAC address to request an IP address for the access VLAN and then an IP address for the voice VLAN, only the access VLAN is assigned an IP address.
When you enter a maximum secure address value for an interface, and the new value is greater than the previous value, the new value overwrites the previously configured value. If the new value is less than the previous value and the number of configured secure addresses on the interface exceeds the new value, the command is rejected.
The switch does not support port security aging of sticky secure MAC addresses.
This table summarizes port security compatibility with other port-based features.
| Table 4. Port Security Compatibility with Other Switch Features |  | 
|---|---|
| Type of Port or Feature on Port | Compatible with Port Security | 
|---|---|
| DTP 4 port 5 | No | 
| Trunk port | Yes | 
| Dynamic-access port 6 | No | 
| Routed port | No | 
| SPAN source port | Yes | 
| SPAN destination port | No | 
| EtherChannel | No | 
| Tunneling port | Yes | 
| Protected port | Yes | 
| IEEE 802.1x port | Yes | 
| Voice VLAN port 7 | Yes | 
| IP source guard | Yes | 
| Dynamic Address Resolution Protocol (ARP) inspection | Yes | 
| Flex Links | Yes | 
A device in a network allows traffic like SNMP, HTTP, HTTPS, Telnet, Secure Shell SSH and Netconf through any port with any IP address. Traffic flow from various interfaces to the local devices might decrease the security strength in a network.
Management traffic control feature allows traffic to enter through a user-defined physical interface and restricts traffic to any other interfaces that is not defined by the user. When the feature is enabled a single IP address is assigned on the device to receive traffic. The user can configure the feature by defining an interface under the management traffic control feature. When the network protocol and the IP address is set according to the user’s preference, traffic flow is allowed only through the defined interface.
The feature is supported on:
Layer 2 physical interface.
Layer 3 physical interface.
Layer 2 port channel.
Layer 3 port channel.
App-hosting interface.
For example, in the following figure Switch1 and Switch2 are devices that are in a network. Management traffic control feature is enabled on Switch1 with the interface Tengig 1/0/1 and destination IP address 30.30.30.30. Traffic is allowed through the interface with the enabled protocols SSH, Telnet, SNMP, HTTP, HTTPS, Netconf, SCP to destination IP address 30.30.30.30. Traffic passing through interface Tengig 1/0/1 to IP address 20.20.20.20 is dropped. Management traffic control feature enables only one destination IP address to be configured. If interface TenGig 1/0/2 is not defined by the management traffic control feature for any of the devices, traffic will not be allowed to Switch1 but can still pass through to its configured destination in the network.
| Note |  | 
| Table 5. Supported protocols and the respective port numbers |  |  | 
|---|---|---|
| Protocol | Keyword | Port Number | 
|---|---|---|
| HTTPS | TCP | 443 | 
| TELNET | TCP | 23 | 
| SSH | TCP | 22 | 
| NETCONF-SSH | TCP | 830 | 
| SNMP | UDP | 161 | 
| HTTP | TCP | 80 | 
This task restricts input to an interface by limiting and identifying MAC addresses of the stations allowed to access the port:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport mode {access \| trunk} Example:  Device(config-if)# switchport mode access   | Sets the interface switchport mode as access or trunk; an interface in the default mode (dynamic auto) cannot be configured as a secure port. | 
| Step 5 | switchport voice vlan vlan-id Example:  Device(config-if)# switchport voice vlan 22   | Enables voice VLAN on a port. vlan-id —Specifies the VLAN to be used for voice traffic. | 
| Step 6 | switchport port-security Example:  Device(config-if)# switchport port-security   | Enables port security on the interface. | 
| Step 7 | switchport port-security [maximum value [vlan {vlan-list \| {access \| voice}}]] Example:  Device(config-if)# switchport port-security maximum 20   | (Optional) Sets the maximum number of secure MAC addresses for the interface. The maximum number of secure MAC addresses that you can configure on a switch or switch stack is set by the maximum number of available MAC addresses allowed in the system. This number is the total of available MAC addresses, including those used for other Layer 2 functions and any other secure MAC addresses configured on interfaces. (Optional) vlan —sets a per-VLAN maximum value Enter one of these options after you enter the vlan keyword:  | 
| Step 8 | switchport port-security violation {protect \| restrict \| shutdown \| shutdown vlan} Example:  Device(config-if)# switchport port-security violation restrict   | (Optional) Sets the violation mode, the action to be taken when a security violation is detected, as one of these:  | 
