---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-administration-guide-614237-snmp-examples-744c8f74-2
title: "SNMP examples"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-administration-guide-614237-snmp-examples-744c8f74.md
source_anchor: ""
source_lines: [180, 259]
sha256: ae6f9801d14a37939192d459062f35f49bae81b825925ea82cf24eecc592d86b
---

# SNMP examples

1. 
                                                    Start the packet capture on interface port1 with the filter set to port 162. See Using the packet capture tool for more information.
2. 
                                                    Overload the DHCP server IP pool.
3. 
                                                    Save the packet capture. The SNMP v3 trap is transmitted from port1 to the SNMP manager. Note that both `msgAuthenticationParameters` and`msgPrivacyParameters` are set up, indicating that authentication and encryption are active. This is further confirmed by`encryptedPDU` in`msgData` .
4. 
                                                    Verify that the SNMP manager has received the trap. See Important SNMP traps for an example of a trap.
5. 
                                                    Verify that the SNMP manager can successfully query and receive DHCP lease usage information for FortiGate: # snmpwalk -v3 -u DHCP_Status -l authPriv -a SHA384 -A xxxxxxxx -x AES256 -X xxxxxxxx 172.16.200.1 1.3.6.1.4.1.12356.101.23 iso.3.6.1.4.1.12356.101.23.1.1.0 = INTEGER: 6 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.1 = INTEGER: 0 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.2 = INTEGER: 0 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.3 = INTEGER: 0 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.4 = INTEGER: 0 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.5 = INTEGER: 0 iso.3.6.1.4.1.12356.101.23.2.1.1.2.1.6 = INTEGER: 100

In the following example, the same IP address will be set on different ports in two VDOMs. The ipAddrTable SNMP Tree output will then be reviewed before and after enabling the *append-index* command.

|  | When the `append-index` command is enabled:   When the `append-index` command is disabled:  | 

###### To enable the INDEX extension:

1. 
                                                    In two different VDOMs, set the same address on two different ports. ```
config system interface
    edit "port3"
        set vdom "vdom1"
        set ip 10.1.1.1 255.255.255.0
        set type physical
        set snmp-index 5
    next
end
config system interface
    edit "port4"
        set vdom "root"
        set ip 10.1.1.1 255.255.255.0
        set type physical
        set snmp-index 6
    next
end
```
2. 
                                                    Configure the SNMP information but do not enable the INDEX extension. ```
config system snmp sysinfo
    set status enable
    set description "REGR-SYS"
end
```
3. 
                                                    On your PC, review the ipAddrTable SNMP Tree (OID 1.3.6.1.2.1.4.20). The IP address 10.1.1.1 is only displayed once. snmpwalk -v2c -c REGR-SYS 172.16.200.1 1.3.6.1.2.1.4.20 **IP-MIB::ipAdEntAddr.10.1.1.1 = IpAddress: 10.1.1.1** IP-MIB::ipAdEntAddr.10.255.1.1 = IpAddress: 10.255.1.1
IP-MIB::ipAdEntAddr.172.16.200.1 = IpAddress: 172.16.200.1
IP-MIB::ipAdEntAddr.192.168.1.99 = IpAddress: 192.168.1.99**IP-MIB::ipAdEntIfIndex.10.1.1.1 = INTEGER: 5** IP-MIB::ipAdEntIfIndex.10.255.1.1 = INTEGER: 39
IP-MIB::ipAdEntIfIndex.172.16.200.1 = INTEGER: 3
IP-MIB::ipAdEntIfIndex.192.168.1.99 = INTEGER: 2**IP-MIB::ipAdEntNetMask.10.1.1.1 = IpAddress: 255.255.255.0** IP-MIB::ipAdEntNetMask.10.255.1.1 = IpAddress: 255.255.255.0
IP-MIB::ipAdEntNetMask.172.16.200.1 = IpAddress: 255.255.255.0
IP-MIB::ipAdEntNetMask.192.168.1.99 = IpAddress: 255.255.255.0**IP-MIB::ipAdEntBcastAddr.10.1.1.1 = INTEGER: 1** IP-MIB::ipAdEntBcastAddr.10.255.1.1 = INTEGER: 1
IP-MIB::ipAdEntBcastAddr.172.16.200.1 = INTEGER: 1
IP-MIB::ipAdEntBcastAddr.192.168.1.99 = INTEGER: 1**IP-MIB::ipAdEntReasmMaxSize.10.1.1.1 = INTEGER: 65535** IP-MIB::ipAdEntReasmMaxSize.10.255.1.1 = INTEGER: 65535
IP-MIB::ipAdEntReasmMaxSize.172.16.200.1 = INTEGER: 65535
IP-MIB::ipAdEntReasmMaxSize.192.168.1.99 = INTEGER: 65535
4. 
                                                    Enable the INDEX extension. ```
config system snmp sysinfo
    set status enable
    set description "REGR-SYS"
    
```
**set append-index enable** end
5. 
                                                    Review the ipAddrTable SNMP Tree (OID 1.3.6.1.2.1.4.20) again. The IP address 10.1.1.1 is now displayed twice. snmpwalk -v2c -c REGR-SYS 172.16.200.1 1.3.6.1.2.1.4.20 **IP-MIB::ipAdEntAddr.10.1.1.1.1 = IpAddress: 10.1.1.1
IP-MIB::ipAdEntAddr.10.1.1.1.2 = IpAddress: 10.1.1.1** IP-MIB::ipAdEntAddr.10.255.1.1.1 = IpAddress: 10.255.1.1
IP-MIB::ipAdEntAddr.172.16.200.1.2 = IpAddress: 172.16.200.1
IP-MIB::ipAdEntAddr.192.168.1.99.1 = IpAddress: 192.168.1.99**IP-MIB::ipAdEntIfIndex.10.1.1.1.1 = INTEGER: 6
IP-MIB::ipAdEntIfIndex.10.1.1.1.2 = INTEGER: 5** IP-MIB::ipAdEntIfIndex.10.255.1.1.1 = INTEGER: 39
IP-MIB::ipAdEntIfIndex.172.16.200.1.2 = INTEGER: 3
IP-MIB::ipAdEntIfIndex.192.168.1.99.1 = INTEGER: 2**IP-MIB::ipAdEntNetMask.10.1.1.1.1 = IpAddress: 255.255.255.0
IP-MIB::ipAdEntNetMask.10.1.1.1.2 = IpAddress: 255.255.255.0** IP-MIB::ipAdEntNetMask.10.255.1.1.1 = IpAddress: 255.255.255.0
IP-MIB::ipAdEntNetMask.172.16.200.1.2 = IpAddress: 255.255.255.0
IP-MIB::ipAdEntNetMask.192.168.1.99.1 = IpAddress: 255.255.255.0**IP-MIB::ipAdEntBcastAddr.10.1.1.1.1 = INTEGER: 1
IP-MIB::ipAdEntBcastAddr.10.1.1.1.2 = INTEGER: 1** IP-MIB::ipAdEntBcastAddr.10.255.1.1.1 = INTEGER: 1
IP-MIB::ipAdEntBcastAddr.172.16.200.1.2 = INTEGER: 1
IP-MIB::ipAdEntBcastAddr.192.168.1.99.1 = INTEGER: 1**IP-MIB::ipAdEntReasmMaxSize.10.1.1.1.1 = INTEGER: 65535
IP-MIB::ipAdEntReasmMaxSize.10.1.1.1.2 = INTEGER: 65535** IP-MIB::ipAdEntReasmMaxSize.10.255.1.1.1 = INTEGER: 65535
IP-MIB::ipAdEntReasmMaxSize.172.16.200.1.2 = INTEGER: 65535
IP-MIB::ipAdEntReasmMaxSize.192.168.1.99.1 = INTEGER: 65535
