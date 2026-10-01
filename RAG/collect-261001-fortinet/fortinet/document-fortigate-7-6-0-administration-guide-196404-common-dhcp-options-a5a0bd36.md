---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-administration-guide-196404-common-dhcp-options-a5a0bd36
title: "execute dhcp lease-clear <ip_address>"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-administration-guide-196404-common-dhcp-options-a5a0bd36.md
source_anchor: ""
source_lines: [1, 75]
sha256: c717f5af443f358001c122a075de413d42272dc82b509399fea2d4e901743f89
---

# execute dhcp lease-clear <ip_address>

Common DHCP options
Common DHCP options
All FortiGate models come with predefined DHCP options. These DHCP options are widely used and required in most scenarios. The following DHCP options can be set straight from the DHCP server section of the Edit Interface dialog:
| Option Code | Option Name | Purpose | 
|---|---|---|
| *1 | Netmask | Assign subnet mask to the DHCP client. | 
| *3 | Default Gateway | Assign default gateway to the DHCP client. | 
| 6 | DNS server | Assign DNS server to the DHCP client. | 
| 42 | NTP server | Assign NTP server to the DHCP client. | 
| *51 | Lease time | Lease time for the DHCP client. | 
| 138 | Wireless controllers | Assign CAPWAP Access Controller addresses to the DHCP client. | 
| 150 | TFTP server(s) | Assign TFTP server to the DHCP client. | 
The parameter marked with an asterisk (*) are mandatory and must be filled in.
Configuring the lease time
This configuration implements DHCP option code 51. The global lease time (measured in seconds, 300 - 864000) determines the length of time an IP address remains assigned to a client. Once the lease expires, the address is released for allocation to the next client that requests an IP address.
To configure the global lease time:
config system dhcp server
    edit <id>
        set interface <interface>
        set netmask <netmask>
        set lease-time <integer>
    next
end
                                            The default lease time is seven days (604800 seconds). To have an unlimited lease time, set the value to zero.
The lease time can also be configured in the GUI in the Lease time field within the DHCP server section of the Edit Interface dialog.
Configuring the lease time for IP ranges
The lease time can be also be configured for an IP range. Measured in seconds, the range is similar to the global lease time (300 - 864000), but the default value is zero (0). If the default (0) is used for an IP range, it applies the global DHCP server lease time value.
To configure the lease time for an IP range:
config system dhcp server
    edit <id>
        config ip-range
            edit <id>
                set lease-time <integer>
            next
        end
    next
end
                                            This setting can only be configured in the CLI.
Customizing DHCP lease backup during power cycles
FortiOS allows customization of the backup interval of DHCP leases during power cycles using the dhcp-lease-backup-interval command. This provides enhanced control and flexibility, ensuring lease preservation during events like outages or reboots. After a power cycle, expired IP addresses are released from the lease list and unexpired IP addresses are retained.
The backup interval can be set between 10 and 3600 seconds, with the default value being 60.
config system global
    set dhcp-lease-backup-interval <integer>
end
                                            Breaking an address lease
If you need to end an IP address lease, you can break the lease. This is useful if you have limited addresses and longer lease times when some leases are no longer necessary, for example, with corporate visitors.
To break a lease:
# execute dhcp lease-clear <ip_address>
To break a lease for all IP addresses for the DHCP servers in the current VDOM:
# execute dhcp lease-clear all
Configuring NTP servers
This configuration implements DHCP option code 42. NTP server can be used by the client to synchronize their time which is very important as for many features to work, including scheduling, logging, and SSL-dependent features, the FortiOS system time must be accurate. This option specifies a list of the NTP servers available to the client by IP address.
To configure NTP servers:
config system dhcp server 
    edit 2
        set ntp-service {local | default | specify}
        set ntp-server1 <class_ip>
        set ntp-server2 <class_ip>
        set ntp-server3 <class_ip>
    next
end
                                            NTP servers can also be configured in the GUI in the NTP server field within the DHCP server > Advanced section of the Edit Interface dialog.
| ntp-service {local \| default \| specify} | Set the option for assigning NTP servers to DHCP clients:  | 
Configuring TFTP servers
This configuration implements DHCP option code 150. TFTP server are used by VoIP phones to obtain the VoIP Configuration. You can configure multiple TFTP servers for a DHCP server. For example, you may want to configure a main TFTP server and a backup TFTP server.
The tftp-server command allows you to configure the TFTP servers, using either their hostnames or IP addresses. Separate multiple server entries with spaces.
To configure TFTP servers:
config system dhcp server
    edit <id>
        set interface <interface>
        set netmask <netmask>
        set tftp-server <hostname/IP address> <hostname/IP address>
    next
end
                                            TFTP servers can also be configured in the GUI in the TFTP server(s) field within the DHCP server > Advanced section of the Edit Interface dialog.
