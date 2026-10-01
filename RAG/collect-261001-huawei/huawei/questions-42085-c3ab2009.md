---
id: collect-261001-huawei/huawei/questions-42085-c3ab2009
title: "questions-42085-c3ab2009"
domain: huawei
role: reference
task: reference
actors: ["Broadcom", "Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/questions-42085-c3ab2009.md
source_anchor: ""
source_lines: [1, 31]
sha256: b065a9358fa0b79027122986c0dd384d898a819f4d306063cd2b825892e933b8
---

# questions-42085-c3ab2009

I'm trying to do up a wiring diagram for a VMWare cluster in a remote data center.  I typically will check the MAC addresses in VMWare VSphere client (vmnic[x]) and/or use esxcli to do so from SSH on the ESXi hosts.
I'll then log into the switch and look at the MAC table to see what port a MAC address was learned from and go from there.
Having an odd issue with Huawei CloudEngine 6800 series switches where the MAC address that's reported in the MAC table (display mac-addresses | include xxxx-xxxx-xxxx) is off on the last bit.  
For example, vmnic0 is one port on a Broadcom QLogic 57800 2x10GbE + 2x1GbE quad adapter, which reports as xx:xx:xx:xx:xx:b5.  Filtering on xxxx-xxxx-xxb5 produced zero results in the filtering options.  However, there's an entry for xxxx-xxxx-xxb6.
I checked a few more NICs and it appears that it's just the Broadcom QLogic 57800 (and 57810) ones are the only ones with inconsistencies, both 10GbE and 1GbE: it's as if the MACs are all off by a bit (c0 shows up as c1 on the switch, c2 as c3 and so on). 
Now I'm a bit rusty, but I know that converged interfaces can be "sliced up" into sub-interfaces, which would have different MAC addresses, but that would also be represented higher up in the "stack" in ESXi, so presumably the source MAC is being parsed out of the frame correctly.
Just not sure if this is a "UX" problem (the display mac-addresses command/switch has bugs) or potentially a driver/incompatibility issue or I'm missing something here.
Command Outputs:
vmnic02 is the only MAC address that's reporting "correctly".  I guess you can assume that b5 is what the Huawei thinks is b6 and b7 is what the Huawei thinks is b8, but I have no explanation as to why.
Card is this from what I can tell, going by the PCI slots (quad card, not dual) http://accessories.euro.dell.com/sna/productdetail.aspx?c=ie&l=en&s=dhs&cs=iedhs1&sku=540-bbfh#Overview
Name    PCI Device    Driver  Admin Status  Link Status  Speed  Duplex  MAC Address         MTU  Description
------  ------------  ------  ------------  -----------  -----  ------  -----------------  ----  -------------------------------------------------------------
vmnic0  0000:01:00.0  bnx2x   Up            Up           10000  Full    11:22:33:44:55:b5  1500  Broadcom Corporation QLogic 57800 10 Gigabit Ethernet Adapter
vmnic1  0000:01:00.1  bnx2x   Up            Up           10000  Full    11:22:33:44:55:b7  1500  Broadcom Corporation QLogic 57800 10 Gigabit Ethernet Adapter
vmnic2  0000:01:00.2  bnx2x   Up            Up            1000  Full    11:22:33:44:55:b9 1500  Broadcom Corporation QLogic 57800 1 Gigabit Ethernet Adapter
vmnic3  0000:01:00.3  bnx2x   Up            Up            1000  Full    11:22:33:44:55:bb  1500  Broadcom Corporation QLogic 57800 1 Gigabit Ethernet Adapter
vmnic4  0000:83:00.0  bnx2x   Up            Up           10000  Full    aa:bb:cc:dd:ee:20  1500  Broadcom Corporation QLogic 57810 10 Gigabit Ethernet Adapter
vmnic5  0000:83:00.1  bnx2x   Up            Up           10000  Full    aa:bb:cc:dd:ee:22  1500  Broadcom Corporation QLogic 57810 10 Gigabit Ethernet Adapter
display mac-address | include 1122-3344-55
---- Flags: * - Backup
-------------------------------------------------------------------------------
MAC Address    VLAN/VSI                          Learned-From        Type
-------------------------------------------------------------------------------
1122-3344-55b6 15/-                               10GE1/0/1           dynamic
1122-3344-55b8 16/-                               10GE2/0/1           dynamic
1122-3344-55b9 19/-                               10GE2/0/9           dynamic
-------------------------------------------------------------------------------
EDIT 2017/06/28:
Checked iDRAC on host, two 10GbE ports are configured in Standard Mode and iSCSI is not enabled on the adapters nor in ESXi (portgroup has iSCSI Bindings disabled) so I don't think is a case of a "leaking" iSCSI frame: if it was, you'd likely see an extra MAC address, not an incorrect one as the total MAC count, despite being off incrementally, is still correct.
One thing I noticed is that ESXi vSphere Win32 client sees it as two dual adapters under host > Configuration > Network Adapters.
Checked driver, seems ok, so still thinking this is a Huawei issue. Think next step is to do a packet capture and/or call Huawei support.
