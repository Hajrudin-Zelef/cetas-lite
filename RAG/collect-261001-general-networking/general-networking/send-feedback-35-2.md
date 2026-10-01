---
id: collect-261001-general-networking/general-networking/send-feedback-35-2
title: "Inter-VDOM routing configuration example: Internet access"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/send-feedback-35.md
source_anchor: ""
source_lines: [134, 349]
sha256: 303b798c133b75ce907915efbe898c4ecee1e7fb12fdbde5d23d0be5b32c81b0
---

# Inter-VDOM routing configuration example: Internet access

1. 
                                                    In the *Accounting* VDOM, go to*Policy & Objects > Firewall Policy* .
2. 
                                                    Click *Create New* .
3. 
                                                    Enter the following information: **Name**Account-Local-to-Management **Incoming Interface**port2 **Outgoing Interface**AccountVlnk0 **Source**All **Destination**All **Schedule**always **Service**ALL **Action**ACCEPT **NAT**enabled
4. 
                                                    Click *OK* .
5. 
                                                    In the *root* VDOM, go to*Policy & Objects > Firewall Policy* .
6. 
                                                    Click *Create New* .
7. 
                                                    Enter the following information: **Name**Account-VDOM-to-Internet **Incoming Interface**AccountVlnk1 **Outgoing Interface**port1 **Source**All **Destination**All **Schedule**always **Service**ALL **Action**ACCEPT **NAT**enabled
8. 
                                                    Click *OK* .

###### To configure the firewall policies from SalesLocal to Internet in the GUI:

1. 
                                                    In the *Sales* VDOM, go to*Policy & Objects > Firewall Policy* .
2. 
                                                    Click *Create New* .
3. 
                                                    Enter the following information: **Name**Sales-Local-to-Management **Incoming Interface**port3 **Outgoing Interface**SalesVlnk0 **Source**All **Destination**All **Schedule**always **Service**ALL **Action**ACCEPT **NAT**enabled
4. 
                                                    Click *OK* .
5. 
                                                    In the *root* VDOM, go to*Policy & Objects > Firewall Policy* .
6. 
                                                    Click *Create New* .
7. 
                                                    Enter the following information: **Name**Sales-VDOM-to-Internet **Incoming Interface**SalesVlnk1 **Outgoing Interface**port1 **Source**All **Destination**All **Schedule**always **Service**ALL **Action**ACCEPT **NAT**enabled
8. 
                                                    Click *OK* .

When the inter-VDOM routing has been configured, test the configuration to confirm proper operation. Testing connectivity ensures that physical networking connections, FortiGate unit interface configurations, and firewall policies are properly configured.

The easiest way to test connectivity is to use the `ping` and `traceroute` commands on hosts in the Accounting and Sales networks, respectively, to confirm the connectivity of different routes on the network. Test connectivity with hosts connected to port2 (AccountingLocal) in the Accounting VDOM to the internet and hosts connected to port3 (SalesLocal) in the Sales VDOM to the internet.

The example can also be configured in the CLI.

###### To configure inter-VDOM routing in the CLI:

1. 
                                                    Enable multi-VDOM mode: ```
config system global
    set vdom-mode multi-vdom
end
```
You will be logged out of the device when VDOM mode is enabled.
2. 
                                                    Create the Sales and Accounting VDOMs: ```
config vdom
    edit Accounting
    next
    edit Sales
    next
end
```
3. 
                                                    Assign interfaces to the VDOMs: ```
config global
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
```
4. 
                                                    Configure the Accounting and management VDOM link: ```
config global
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
```
5. 
                                                    Configure the Sales and management VDOM link: ```
config global
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
```
6. 
                                                    Configure the default static route to the Internet in the Accounting VDOM: ```
config vdom
    edit Accounting
    config router static
        edit 1
            set gateway 11.11.11.1
            set device "AccountVlnk0"
        next
    end
end
```
7. 
                                                    Configure the default statis route to the Internet in the Sales VDOM: ```
config vdom
    edit Sales
    config router static
        edit 1
            set gateway 12.12.12.1
            set device "SalesVlnk0"
        next
    end
end
```
8. 
                                                    Configure the firewall policies from AccountingLocal to the Internet: ```
config vdom
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
                set service ALL
                set nat enable
            next
        end
    next
end
```
9. 
                                                    Configure the firewall policies from SalesLocal to the Internet: ```
config vdom
    edit Sales
        config firewall policy
            edit 3
                set name "Sales-local-to-Management"
                set srcintf port3
                set dstintf SalesVlnk0
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
            edit 4
                set name "Sales-VDOM-to-Internet"
                set srcintf SalesVlnk1
                set dstintf port1
                set srcaddr all
                set dstaddr all
                set action accept
                set schedule always
                set service ALL
                set nat enable
            next
        end
    next
end
```
