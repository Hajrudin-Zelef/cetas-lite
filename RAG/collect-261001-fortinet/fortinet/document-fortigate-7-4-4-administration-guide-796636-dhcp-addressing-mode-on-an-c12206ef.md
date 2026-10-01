---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-796636-dhcp-addressing-mode-on-an-c12206ef
title: "document-fortigate-7-4-4-administration-guide-796636-dhcp-addressing-mode-on-an--c12206ef"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-796636-dhcp-addressing-mode-on-an--c12206ef.md
source_anchor: ""
source_lines: [1, 74]
sha256: 631c633654554f7d49db413c68d87aa4ce7cfc0d5215b5e0975e612b2ca596e3
---

# document-fortigate-7-4-4-administration-guide-796636-dhcp-addressing-mode-on-an--c12206ef

DHCP addressing mode on an interface
DHCP addressing mode on an interface
Any FortiGate interface can be configured to obtain an IP address dynamically using DHCP. If you configure DHCP on an interface on the FortiGate, the FortiGate automatically broadcasts a DHCP request from the interface. The interface is configured with the IP address, any DNS server addresses, and the default gateway address that the DHCP server provides.
Configuring an Interface as a DHCP Client
You can configure interface as a DHCP client.
To configure an interface as a DHCP client in the GUI:
- 
                                                    Go to Network > Interfaces.
- 
                                                    Edit an interface.
- 
                                                    Select the DHCP option in the Addressing mode.
- 
                                                    Configure the rest of the setting as required.
- 
                                                    Click OK.
The following table describes the DHCP status information when DHCP is configured for an interface.
| Field | Description | 
|---|---|
| Status | Displays DHCP status messages as the interface connects to the DHCP server and gets addressing information. Status can be one of the following values:  | 
| Obtained IP/Netmask | The IP address and netmask leased from the DHCP server. This is only displayed if the Status is Connected. | 
| Renew | Select this to renew the DHCP license for this interface. This is only displayed if the Status is Connected. | 
| Expiry Date | The time and date when the leased IP address and netmask is no longer valid for the interface. The IP address is returned to the pool to be allocated to the next user request for an IP address. This is only displayed if the Status is Connected. | 
| Default Gateway | The IP address of the gateway defined by the DHCP server. This is displayed only if the Status is Connected, and if Retrieve default gateway from server is enabled. | 
| Acquired DNS | The DNS server IP defined by the DHCP server. This is displayed only if the Status is Connected. | 
| Retrieve default gateway from server | Enable this to retrieve a default gateway IP address from the DHCP server. The default gateway is added to the static routing table. | 
| Distance | Enter the administrative distance for the default gateway retrieved from the DHCP server. The administrative distance is an integer from 1 to 255 that specifies the relative priority of a route when there are multiple routes to the same destination. A lower administrative distance indicates a more preferred route. | 
| Override internal DNS | Enable this to use the DNS addresses retrieved from the DHCP server instead of the DNS server IP addresses on the DNS page. When VDOMs are enabled, you can override the internal DNS only on the management VDOM. | 
To configure an interface as a DHCP client in the CLI:
config system interface
    edit <name>
        set mode dhcp
        set defaultgw {enable | disable}
        set distance <integer>
        set dns-server-override {enable | disable}
    next
end
                                            Configuring the DHCP renew time
You can set a minimum DHCP renew time for an interface acting as a DHCP client. This option is available only when mode is set to DHCP.
To set the DHCP renew time:
config system interface
    edit <name>
        set vdom <vdom>
        set interface <interface>
        set mode dhcp
        set dhcp-renew-time <integer>
    next
end
                                            The possible values for dhcp-renew-time are 300 to 605800 seconds (five minutes to seven days). To use the renew time that the server provides, set this entry to 0.
DHCP client options
When an interface is in DHCP addressing mode, DHCP client options can be configured in the CLI. For example, a vendor class identifier (usually DCHP client option 60) can be specified so that a request can be matched by a specific DHCP offer.
Multiple options can be configured, but any options not recognized by the DHCP server are discarded.
To configure client option 60 - vendor class identifier:
config system interface
    edit port1
        set vdom vdom1
        set mode dhcp
        config client-options
            edit 1
                set code 60
                set type hex
                set value aabbccdd
            next
        end
        set type physical
        set snmp-index 4
    next
end
                                            | Variable | Description | 
|---|---|
| code <integer> | DHCP client option code (0 - 255, default = 0). See Dynamic Host Configuration Protocol (DHCP) and Bootstrap Protocol (BOOTP) Parameters for a list of possible options. | 
| type {hex \| string \| ip \| fqdn} | DHCP client option type (default = hex). | 
| value <string> | DHCP client option value. | 
| ip <ip> | DHCP client option IP address. This option is only available when type isip . |
