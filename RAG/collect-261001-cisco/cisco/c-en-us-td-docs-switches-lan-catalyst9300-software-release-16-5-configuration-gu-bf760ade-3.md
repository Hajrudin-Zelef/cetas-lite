---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [124, 198]
sha256: 8f2cd18c665ec0357944b7dabe27705a9af7d149bce1c71e334e2afed3e6a3c1
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

                                    Port Security is not supported on EtherChannel interfaces.
Information About Port Security
Port Security
You can use the port security feature to restrict input to an interface by limiting and identifying MAC addresses of the stations allowed to access the port. When you assign secure MAC addresses to a secure port, the port does not forward packets with source addresses outside the group of defined addresses. If you limit the number of secure MAC addresses to one and assign a single secure MAC address, the workstation attached to that port is assured the full bandwidth of the port.
If a port is configured as a secure port and the maximum number of secure MAC addresses is reached, when the MAC address of a station attempting to access the port is different from any of the identified secure MAC addresses, a security violation occurs. Also, if a station with a secure MAC address configured or learned on one secure port attempts to access another secure port, a violation is flagged.
Types of Secure MAC Addresses
The switch supports these types of secure MAC addresses:
- 
                                    Static secure MAC addresses—These are manually configured by using the switchport port-security mac-address mac-address interface configuration command, stored in the address table, and added to the switch running configuration.
- 
                                    Dynamic secure MAC addresses—These are dynamically configured, stored only in the address table, and removed when the switch restarts.
- 
                                    Sticky secure MAC addresses—These can be dynamically learned or manually configured, stored in the address table, and added to the running configuration. If these addresses are saved in the configuration file, when the switch restarts, the interface does not need to dynamically reconfigure them.
Sticky Secure MAC Addresses
You can configure an interface to convert the dynamic MAC addresses to sticky secure MAC addresses and to add them to the running configuration by enabling sticky learning. The interface converts all the dynamic secure MAC addresses, including those that were dynamically learned before sticky learning was enabled, to sticky secure MAC addresses. All sticky secure MAC addresses are added to the running configuration.
The sticky secure MAC addresses do not automatically become part of the configuration file, which is the startup configuration used each time the switch restarts. If you save the sticky secure MAC addresses in the configuration file, when the switch restarts, the interface does not need to relearn these addresses. If you do not save the sticky secure addresses, they are lost.
If sticky learning is disabled, the sticky secure MAC addresses are converted to dynamic secure addresses and are removed from the running configuration.
Security Violations
It is a security violation when one of these situations occurs:
- 
                                       The maximum number of secure MAC addresses have been added to the address table, and a station whose MAC address is not in the address table attempts to access the interface.
- 
                                       An address learned or configured on one secure interface is seen on another secure interface in the same VLAN.
- 
                                       
                                       
                                       Running diagnostic tests with port security enabled.
You can configure the interface for one of three violation modes, based on the action to be taken if a violation occurs:
- 
                                       protect—when the number of secure MAC addresses reaches the maximum limit allowed on the port, packets with unknown source addresses are dropped until you remove a sufficient number of secure MAC addresses to drop below the maximum value or increase the number of maximum allowable addresses. You are not notified that a security violation has occurred. 
 Note
 We do not recommend configuring the protect violation mode on a trunk port. The protect mode disables learning when any VLAN reaches its maximum limit, even if the port has not reached its maximum limit. 
- 
                                       restrict—when the number of secure MAC addresses reaches the maximum limit allowed on the port, packets with unknown source addresses are dropped until you remove a sufficient number of secure MAC addresses to drop below the maximum value or increase the number of maximum allowable addresses. In this mode, you are notified that a security violation has occurred. An SNMP trap is sent, a syslog message is logged, and the violation counter increments.
- 
                                       shutdown—a port security violation causes the interface to become error-disabled and to shut down immediately, and the port LED turns off. When a secure port is in the error-disabled state, you can bring it out of this state by entering the errdisable recovery cause psecure-violation global configuration command, or you can manually re-enable it by entering the shutdown and no shut down interface configuration commands. This is the default mode.
- 
                                       shutdown vlan—Use to set the security violation mode per-VLAN. In this mode, the VLAN is error disabled instead of the entire port when a violation occurs
This table shows the violation mode and the actions taken when you configure an interface for port security.
| Table 3. Security Violation Mode Actions |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Violation Mode | Traffic is forwarded 1 | Sends SNMP trap | Sends syslog message | Displays error message 2 | Violation counter increments | Shuts down port | 
|---|---|---|---|---|---|---|
| protect | No | No | No | No | No | No | 
| restrict | No | Yes | Yes | No | Yes | No | 
| shutdown | No | No | No | No | Yes | Yes | 
| shutdown vlan | No | No | Yes | No | Yes | No 3 | 
Port Security Aging
You can use port security aging to set the aging time for all secure addresses on a port. Two types of aging are supported per port:
- 
                                    Absolute—The secure addresses on the port are deleted after the specified aging time.
- 
                                    Inactivity—The secure addresses on the port are deleted only if the secure addresses are inactive for the specified aging time.
Port Security and Switch Stacks
When a switch joins a stack, the new switch will get the configured secure addresses. All dynamic secure addresses are downloaded by the new stack member from the other stack members.
When a switch (either the active switch or a stack member) leaves the stack, the remaining stack members are notified, and the secure MAC addresses configured or learned by that switch are deleted from the secure MAC address table.
Default Port Security Configuration
| Table 4. Default Port Security Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Port security | Disabled on a port. | 
| Sticky address learning | Disabled. | 
| Maximum number of secure MAC addresses per port | 1. | 
| Violation mode | Shutdown. The port shuts down when the maximum number of secure MAC addresses is exceeded. | 
| Port security aging | Disabled. Aging time is 0. Static aging is disabled. Type is absolute. | 
Port Security Configuration Guidelines
- 
                                       					
                                       Port security can only be configured on static access ports or trunk ports. A secure port cannot be a dynamic access port.
- 
                                       					
                                       A secure port cannot be a destination port for Switched Port Analyzer (SPAN).
- 
                                       					
