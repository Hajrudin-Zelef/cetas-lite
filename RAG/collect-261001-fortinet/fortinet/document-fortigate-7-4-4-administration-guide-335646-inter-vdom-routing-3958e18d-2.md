---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d-2
title: "document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d.md
source_anchor: ""
source_lines: [85, 260]
sha256: 56dbe78262db448225c599055ea435865ec7abe86df136865db16c15c671ae29
---

# document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d

- 
                                                    IP address: 0.0.0.0/0.0.0.0 (default)
To configure the default static route to the Internet in the Accounting VDOM:
- 
                                                    In the Accounting VDOM, go to Network > Static Routes.
- 
                                                    Click on Create New and select the version you need.
- 
                                                    Enter the following information: Destination Subnet IP address 0.0.0.0/0.0.0.0 Gateway 11.11.11.1 Interface AccountVlink0 Administrative Distance 10
- 
                                                    Click OK.
To configure the default static route to the Internet in the Sales VDOM:
- 
                                                    In the Sales VDOM, go to Network > Static Routes.
- 
                                                    Click on Create New and select the version you need.
- 
                                                    Enter the following information: Destination Subnet IP address 0.0.0.0/0.0.0.0 Gateway 12.12.12.1 Interface SalesVlink0 Administrative Distance 10
- 
                                                    Click OK.
Configure the firewall policies
With the VDOMs, physical interfaces, VDOM links, and static routes configured, the firewall must now be configured to allow the proper traffic. Firewalls are configured per-VDOM, and firewall objects and routes must be created for each VDOM separately.
To configure the firewall policies from AccountingLocal to Internet in the GUI:
- 
                                                    In the Accounting VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Account-Local-to-Management Incoming Interface port2 Outgoing Interface AccountVlnk0 Source All Destination All Schedule always Service ALL Action ACCEPT NAT enabled
- 
                                                    Click OK.
- 
                                                    In the root VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Account-VDOM-to-Internet Incoming Interface AccountVlnk1 Outgoing Interface port1 Source All Destination All Schedule always Service ALL Action ACCEPT NAT enabled
- 
                                                    Click OK.
To configure the firewall policies from SalesLocal to Internet in the GUI:
- 
                                                    In the Sales VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Sales-Local-to-Management Incoming Interface port3 Outgoing Interface SalesVlnk0 Source All Destination All Schedule always Service ALL Action ACCEPT NAT enabled
- 
                                                    Click OK.
- 
                                                    In the root VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Sales-VDOM-to-Internet Incoming Interface SalesVlnk1 Outgoing Interface port1 Source All Destination All Schedule always Service ALL Action ACCEPT NAT enabled
- 
                                                    Click OK.
Test the configuration
When the inter-VDOM routing has been configured, test the configuration to confirm proper operation. Testing connectivity ensures that physical networking connections, FortiGate unit interface configurations, and firewall policies are properly configured.
The easiest way to test connectivity is to use the ping and traceroute commands on hosts in the Accounting and Sales networks, respectively, to confirm the connectivity of different routes on the network. Test connectivity with hosts connected to port2 (AccountingLocal) in the Accounting VDOM to the internet and hosts connected to port3 (SalesLocal) in the Sales VDOM to the internet.
Configuration with the CLI
The example can also be configured in the CLI.
To configure inter-VDOM routing in the CLI:
- 
                                                    Enable multi-VDOM mode: config system global
    set vdom-mode multi-vdom
endYou will be logged out of the device when VDOM mode is enabled.
- 
                                                    Create the Sales and Accounting VDOMs: config vdom
    edit Accounting
    next
    edit Sales
    next
end
- 
                                                    Assign interfaces to the VDOMs: config global
    config system interface
        edit port2
            set vdom Accounting
        next
        edit port3
            set vdom Sales
        next
        edit port1
            set vdom root
        next
    end
end
- 
                                                    Configure the Accounting and management VDOM link: config global
    config system vdom-link
        edit AccountVlnk
        next
    end
    config system interface
        edit AccountVlnk0
            set vdom Accounting
            set ip 11.11.11.2 255.255.255.252
            set allowaccess https ping ssh
            set description "Accounting side of the VDOM link"
        next
        edit AccountVlnk1
            set vdom root
            set ip 11.11.11.1 255.255.255.252
            set allowaccess https ping ssh
            set description "Management side of the VDOM link"
        next
    end
end
- 
                                                    Configure the Sales and management VDOM link: config global
    config system vdom-link
        edit SalesVlnk
        next
    end
    config system interface
        edit SalesVlnk0
            set vdom Sales
            set ip 12.12.12.2 255.255.255.252
            set allowaccess https ping ssh
            set description "Sales side of the VDOM link"
        next
        edit SalesVlnk1
            set vdom root
            set ip 12.12.12.1 255.255.255.252
            set allowaccess https ping ssh
            set description "Management side of the VDOM link"
        next
    end
end
- 
                                                    Configure the default static route to the Internet in the Accounting VDOM: config vdom
    edit Accounting
    config router static
        edit 1
            set gateway 11.11.11.1
            set device "AccountVlnk0"
        next
    end
end
- 
                                                    Configure the default statis route to the Internet in the Sales VDOM: config vdom
    edit Sales
    config router static
        edit 1
            set gateway 12.12.12.1
            set device "SalesVlnk0"
        next
    end
end
- 
                                                    Configure the firewall policies from AccountingLocal to the Internet: config vdom
    edit Accounting
        config firewall policy
            edit 1
                set name "Accounting-Local-to-Management"
                set srcintf port2
                set dstintf AccountVlnk0
                set srcaddr all
                set dstaddr all
                set action accept
                set schedule always
                set service ALL
                set nat enable
            next
        end
    next
    edit root
        config firewall policy
            edit 2
                set name "Accounting-VDOM-to-Internet"
                set srcintf AccountVlnk1
                set dstintf port1
                set srcaddr all
                set dstaddr all
                set action accept
                set schedule always
