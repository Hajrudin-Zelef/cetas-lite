---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-administration-guide-614237-snmp-examples-744c8f74-1
title: "SNMP examples"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-administration-guide-614237-snmp-examples-744c8f74.md
source_anchor: ""
source_lines: [1, 179]
sha256: d5857588b364c51f46c344410da6a93905f075aa18f85e1d9c5b980a7c473a3a
---

# SNMP examples

# SNMP examples

This topic includes examples that incorporate several SNMP settings:

- 
                                                    Example 1: SNMP traps for monitoring interface status using SNMP v3 user
- 
                                                    Example 2: SNMP traps and query for monitoring DHCP pool using SNMP v3 user
- 
                                                    Example 3: Enabling the INDEX extension

This configuration enables the SNMP manager (172.16.200.55) to receive notifications when a FortiGate port either goes down or is brought up. The SNMP manager can also query the current status of the FortiGate port.

###### To configure SNMP for monitoring interface status in the GUI:

1. 
                                                    Configure interface access: 
  1. 
                                                            Go to *Network > Interfaces* and edit*port1* .
  2. 
                                                            In the *Administrative Access* options, enable*SNMP* .
  3. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Configure the SNMP agent: 
  1. 
                                                            Go to *System > SNMP* .
  2. 
                                                            Enable *SNMP Agent* and enter the following:Description Branch Location Burnaby Contact Info Jane Doe
  3. 
                                                            Click *Apply* .
4. 
                                                            
5. 
                                                    Configure an SNMP v3 user: 
  1. 
                                                            Go to *System > SNMP* .
  2. 
                                                            In the *SNMP v3* table, click*Create New* .
  3. 
                                                            Configure the following: User Name Interface_Status Security Level Authentication Authentication Algorithm SHA1 Password ******* Hosts IP Address 172.16.200.55
  4. 
                                                            Click *OK* .
6. 
                                                            

###### To configure SNMP for monitoring interface status in the CLI:

1. 
                                                    Configure interface access: ```
config system interface
    edit port1
        append allowaccess snmp
    next
end
```
2. 
                                                    Configure the SNMP agent: ```
config system snmp sysinfo
    set status enable
    set description Branch
    set contact-info Jane Doe
    set location Burnaby
end
```
3. 
                                                    Configure an SNMP v3 user: ```
config system snmp user
    edit "Interface_Status"
        set notify-hosts 172.16.200.55
        set security-level auth-no-priv
        set auth-proto sha
        set auth-pwd ********************
    next
end
```

###### Verification

1. 
                                                    Start the packet capture on interface port1 with the filter set to port 162. See Using the packet capture tool for more information.
2. 
                                                    Turn off one of the FortiGate interface statuses to down, in this case, port2.
3. 
                                                    Save the packet capture. The SNMP v3 trap is transmitted from port1 to the SNMP manager. Note that `msgAuthenticationParameters` is configured, indicating that authentication is active. The absence of`msgPrivacyParameters` suggests that encryption is not configured. This is further confirmed by`plaintext` in`msgData` .
4. 
                                                    Verify that the SNMP manager has received the trap. See Important SNMP traps for an example of a trap.
5. 
                                                    Verify that the SNMP manager can successfully query and receive receive a response on the current status of the FortiGate ports: # snmpwalk -v3 -u Interface_Status -l authNoPriv -a SHA -A xxxxxxxx 172.16.200.1 1.3.6.1.2.1.2.2.1.8 iso.3.6.1.2.1.2.2.1.8.1 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.2 = INTEGER: 2 iso.3.6.1.2.1.2.2.1.8.3 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.4 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.5 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.6 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.7 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.8 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.9 = INTEGER: 1 iso.3.6.1.2.1.2.2.1.8.10 = INTEGER: 1

This configuration enables the SNMP manager (172.16.200.55) to receive DHCP-related notifications from FortiGate.

The SNMP DHCP event contains three traps and one query.

Traps are sent when:

- 
                                                    DHCP server IP pool usage reaches 90%
- 
                                                    DHCP server detect an IP address that is already in use
- 
                                                    DHCP client receives DHCP NAK

SNMP queries are accepted for DHCP lease usage information (OID = 1.3.6.1.4.1.12356.101.23). The query result is based on the leased out percentage.

###### To configure SNMP for monitoring DHCP pool in the GUI:

1. 
                                                    Configure interface access: 
  1. 
                                                            Go to *Network > Interfaces* and edit*port1* .
  2. 
                                                            In the *Administrative Access* options, enable*SNMP* .
  3. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Configure the SNMP agent: 
  1. 
                                                            Go to *System > SNMP* .
  2. 
                                                            Enable *SNMP Agent* and enter the following:Description Branch Location Burnaby Contact Info Jane Doe
  3. 
                                                            Click *Apply* .
4. 
                                                            
5. 
                                                    Configure an SNMP v3 user: 
  1. 
                                                            Go to *System > SNMP* .
  2. 
                                                            In the *SNMP v3* table, click*Create New* .
  3. 
                                                            Configure the following: User Name DHCP_Status Security Level Authentication Authentication Algorithm SHA384 Password ******* Private Enabled Encryption Algorithm AES256 Password ******* Hosts IP Address 172.16.200.55
  4. 
                                                            Click *OK* .
6. 
                                                            

###### To configure SNMP for monitoring the DHCP pool in the CLI:

1. 
                                                    Configure interface access: ```
config system interface
    edit port1
        append allowaccess snmp
    next
end
```
2. 
                                                    Configure the SNMP agent: ```
config system snmp sysinfo
    set status enable
    set description Branch
    set contact-info Jane Doe
    set location Burnaby
end
```
3. 
                                                    Configure an SNMP v3 user: ```
config system snmp user
    edit "DHCP_Status"
        set notify-hosts 172.16.200.55
        set security-level auth-priv
        set auth-proto sha384
        set auth-pwd ********************
        set priv-proto aes256
        set priv-pwd *********************
    next
end
```

###### Verification

