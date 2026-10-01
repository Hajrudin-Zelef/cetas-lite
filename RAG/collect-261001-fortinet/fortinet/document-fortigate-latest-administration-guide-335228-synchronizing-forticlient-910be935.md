---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-335228-synchronizing-forticlient-910be935
title: "Synchronizing FortiClient ZTNA tags"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-335228-synchronizing-forticlient--910be935.md
source_anchor: ""
source_lines: [1, 88]
sha256: d5a9107b39f7a0059004015f46a2a421cb167691e2361047889212a6b1a4960e
---

# Synchronizing FortiClient ZTNA tags

# Synchronizing FortiClient ZTNA tags

        ZTNA tags (formerly FortiClient EMS tags in FortiOS 6.4 and earlier) are tags synchronized from FortiClient EMS as dynamic address objects on the FortiGate. FortiClient EMS uses zero-trust tagging rules to automatically tag managed endpoints based on various attributes detected by the FortiClient. When the FortiGate establishes a connection with the FortiClient EMS server through the EMS Fabric connector, it pulls zero-trust tags containing device IP and MAC addresses and converts them to read-only dynamic address objects. It also establishes a persistent WebSocket connection to monitor for changes in zero-trust tags, which keeps the device information current.
     These ZTNA tags can then be used in ZTNA rules, firewall rules, and NAC policies to perform security posture checks. ZTNA tags are displayed in the *Device Inventory* widget, *FortiClient* widget, and *Asset Identity Center* page.

By enabling ZTNA EMS tag checking in a firewall policy, you can include EMS tag information in the traffic log.

When primary and secondary ZTNA EMS tag checking is enabled using address groups, the *Primary EMS tag* and *Secondary EMS tag* fields will be included in the GUI traffic logs. Likewise, the `emstag` and `emstag2` fields will be included in the CLI traffic logs.

When using WebSocket, EMS pushes notifications to the corresponding FortiGate when there are updates to tags or other monitored attributes. The FortiGate then fetches the updated information using the REST API over TCP/8013. When WebSocket is not used (due to an override or unsupported EMS version), updates are triggered on demand from the FortiGate side over the REST API.

If the WebSocket capability is detected, the capabilities setting will automatically display the WebSocket option. You can use the `diagnose test application fcnacd 2` command to view the status of the WebSocket connection.

In the following example, the FortiGate connects to and retrieves ZTNA tags from a FortiClient EMS configured with tagging rules. It is assumed that zero-trust tags and rules are already created on the FortiClient EMS. For more information, see the Zero Trust Tags section of the EMS Administration Guide.

###### To verify zero-trust tags in FortiClient EMS:

1. 
                                                    Go to *Zero Trust Tags > Zero Trust Tagging Rules* to view the tags.
2. 
                                                    Go to *Zero Trust Tags > Zero Trust Tag Monitor* to view the registered users who match the defined tag.

###### To configure the EMS Fabric connector to synchronize ZTNA tags in the GUI:

1. 
                                                    Configure the EMS Fabric connector: 
  1. 
                                                            On the root FortiGate, go to *Security Fabric > Fabric Connectors* and double-click the*FortiClient EMS* card.
  2. 
                                                            Set the *Status* to*Enabled* for one of the EMS servers.
  3. 
                                                            Enable *Synchronize firewall addresses* .
  4. 
                                                            Configure the other settings as needed and validate the certificate.
  5. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Enable ZTNA: 
  1. 
                                                            Go to *System > Feature Visibility* and enable*Zero Trust Network Access* .
  2. 
                                                            Click *Apply* .
4. 
                                                            
5. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Tags* tab. You will see the ZTNA IP and ZTNA MAC tags synchronized from the FortiClient EMS.

###### To configure the EMS Fabric connector to synchronize ZTNA tags in the CLI:

1. 
                                                    Configure the EMS Fabric connector on the root FortiGate: ```
config endpoint-control fctems
    edit 1
        set status enable
        set name "WIN10-EMS"
        set server "192.168.20.10"
        set https-port 443
        set pull-sysinfo enable
        set pull-vulnerabilities enable
        set pull-tags enable
        set pull-malware-hash enable
        set capabilities fabric-auth silent-approval websocket
    next
end
```
2. 
                                                    Verify which IPs the dynamic firewall address resolves to: ```
# diagnose firewall dynamic list 
List all dynamic addresses:
FCTEMS0000100000_all_registered_clients: ID(51)
        ADDR(172.17.194.209)
        ADDR(10.10.10.20)
...
FCTEMS0000100000_Low: ID(78)
        ADDR(172.17.194.209)
        ADDR(10.10.10.20)
...
FCTEMS0000100000_Malicious-File-Detected: ID(190)
        ADDR(172.17.194.209)
        ADDR(10.10.10.20)
...
```

When running the FortiGate in multi-VDOM mode, by default, EMS is configured in the global VDOM. All ZTNA tags synchronized with the globally configured EMS are shared by all VDOMs. FortiOS 7.4 and later supports configuring EMS on a per-VDOM basis. See Configuring FortiClient EMS and FortiClient EMS Cloud on a per-VDOM basis for more information.
