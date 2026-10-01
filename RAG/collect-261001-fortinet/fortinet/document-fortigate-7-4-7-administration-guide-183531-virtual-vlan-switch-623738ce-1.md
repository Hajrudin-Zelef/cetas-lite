---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-183531-virtual-vlan-switch-623738ce-1
title: "document-fortigate-7-4-7-administration-guide-183531-virtual-vlan-switch-623738ce"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-183531-virtual-vlan-switch-623738ce.md
source_anchor: ""
source_lines: [1, 34]
sha256: 235fce16ff36ff5ed0d9ec22ee7d307eedf90f710ae80b7213b0fedfdd118efe
---

# document-fortigate-7-4-7-administration-guide-183531-virtual-vlan-switch-623738ce

Virtual VLAN switch
Virtual VLAN switch
The hardware switch ports on FortiGate models that support virtual VLAN switches can be used as a layer 2 switch. Virtual VLAN switch mode allows 802.1Q VLANs to be assigned to ports, and the configuration of one interface as a trunk port.
The following FortiGate series are supported in FortiOS 7.4: 40F, 60F, 70F, 80F, 90G, 100F, 120G, 140E, 200F, 300E, 400E, 400F, 600F, 900G, 1000F, 1100E, 1800F, 2600F, 3000F, 3200F, 3500F, 3700F, 4200F, 4400F, and 4800F. FortiWiFi 60F models are not supported.
The virtual-switch-vlan option must be enabled in the CLI to configure VLAN switch mode from the GUI or CLI.
To enable VLAN switches:
config system global
    set virtual-switch-vlan enable
end
                                            After this setting is enabled, any previously configured hardware switches will appear in the Network > Interfaces page under VLAN Switch.
To enable VLAN switch mode in the GUI:
- 
                                                    Go to System > Settings.
- 
                                                    In the View Settings section, enable VLAN switch mode.
- 
                                                    Click Apply.
Basic configurations
Hardware switch ports can be configured as either a VLAN switch port or a trunk port. The available interfaces and allowable VLAN IDs that can be used depend on the FortiGate model. It is recommended to remove ports from the default VLAN switch before you begin configurations.
To create a new VLAN and assign ports in the GUI:
- 
                                                    Go to Network > Interfaces and click Create New > Interface.
- 
                                                    Enter a name and configure the following: 
  - 
                                                            Set the Type to VLAN Switch.
  - 
                                                            Enter a VLAN ID.
  - 
                                                            Click the + and add the Interface Members.
  - 
                                                            Configure the Address and Administrative Access settings as needed.
- 
                                                            
