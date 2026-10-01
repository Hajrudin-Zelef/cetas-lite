---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-19
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [4224, 4465]
sha256: a485b857f87d9e15746b3e5bbed910cb7b2553cf284a64aaef6e8895d7252f09
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

If you are starting this section with a non-configured device, begin here and then
move to step 3. For those continuing from previous labs, begin at step 2.
Establish the addressing for the lab and temporarily shut down the interfaces Gigabit
Ethernet 0/0/2 of R1 and GigabitGigabitEthernet 0/0/1 of R3.
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R1
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]ip address 10.0.12.1 24
[R1-GigabitEthernet0/0/1]quit


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R3
[R3]interface GigabitEthernet 0/0/1
[R3-GigabitEthernet0/0/1]ip address 10.0.12.3 24
[R3-GigabitEthernet0/0/1]shutdown
[R3-GigabitEthernet0/0/1]quit
[R3]interface GigabitEthernet 0/0/2
[R3-GigabitEthernet0/0/2]ip address 10.0.23.3 24



<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S1



<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S2




                                                   HUAWEI TECHNOLOGIES    Page93
Step 2 Cleaning up the previous configuration

Re-enable to Gigabit Ethernet 0/0/2 interface on R3.


[R3]interface GigabitEthernet 0/0/2
[R3-GigabitEthernet0/0/2]undo shutdown



Step 3 Additional configuration

Disable the port interfaces between S1 and S2 as well as other interfaces to prevent

interference from other devices.



[S1]interface GigabitEthernet 0/0/9
[S1-GigabitEthernet0/0/9]shutdown
[S1-GigabitEthernet0/0/9]quit
[S1]interface GigabitEthernet 0/0/10
[S1-GigabitEthernet0/0/10]shutdown
[S1-GigabitEthernet0/0/10]quit
[S1]interface GigabitEthernet 0/0/13
[S1-GigabitEthernet0/0/13]shutdown
[S1-GigabitEthernet0/0/13]quit
[S1]interface GigabitEthernet 0/0/14
[S1-GigabitEthernet0/0/14]shutdown


[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]shutdown
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]shutdown
[S2-GigabitEthernet0/0/10]quit
[S2]interface GigabitEthernet 0/0/7
[S2-GigabitEthernet0/0/7]shutdown
[S2-GigabitEthernet0/0/23]quit
[S2]interface GigabitEthernet 0/0/6
[S2-GigabitEthernet0/0/6]shutdown


[R1]interface GigabitEthernet 0/0/2



                                         HUAWEI TECHNOLOGIES              Page94
[R1-GigabitEthernet0/0/2]ip address 10.0.23.1 24
[R1-GigabitEthernet0/0/2]shutdown




Verify that Gigabit Ethernet interfaces 0/0/9, 0/0/10, 0/0/13 and 0/0/14, have been

shut down on S1 and that Gigabit Ethernet interfaces 0/09, 0/0/10, 0/0/6 and 0/0/7

have been shut down on S2.


<S1>display interface brief
…output omitted…
Interface               PHY     Protocol InUti OutUti    inErrors   outErrors
GigabitEthernet0/0/1    up         up       0.01%       0.01%         0         0
GigabitEthernet0/0/2    up         up       0.01%       0.01%         0         0
GigabitEthernet0/0/3    down       down        0%         0%          0         0
GigabitEthernet0/0/4    up         up          0%       0.01%         0         0
GigabitEthernet0/0/5    up         up          0%       0.01%         0         0
GigabitEthernet0/0/6    down       down        0%         0%          0         0
GigabitEthernet0/0/7    down       down        0%         0%          0         0
GigabitEthernet0/0/8    down       down        0%         0%          0         0
GigabitEthernet0/0/9    *down      down        0%         0%          0         0
GigabitEthernet0/0/10 *down        down        0%         0%          0         0
GigabitEthernet0/0/11 down         down        0%         0%          0         0
GigabitEthernet0/0/12 down         down        0%         0%          0         0
GigabitEthernet0/0/13 *down        down        0%         0%          0         0
GigabitEthernet0/0/14 *down        down        0%         0%          0         0
…output omitted…


<S2>display interface brief
…output omit…
GigabitEthernet0/0/1          up    up              0% 4.06%              0         0
GigabitEthernet0/0/2          up    up              0% 4.06%              0         0
GigabitEthernet0/0/3          up    up              0% 4.06%              0         0
GigabitEthernet0/0/4          up    up              0% 20.40%             0         0
GigabitEthernet0/0/5          up    up              0% 20.40%             0         0
GigabitEthernet0/0/6          *down down            0% 2.04%              0         0
GigabitEthernet0/0/7          *down down        2.03% 2.03%               0         0



                                                   HUAWEI TECHNOLOGIES                  Page95
GigabitEthernet0/0/8           down down                0%       0%            0       0
GigabitEthernet0/0/9           *down down             1.91% 1.91%          0       0
GigabitEthernet0/0/10          *down down             3.95% 0.12%          0       0
GigabitEthernet0/0/11          up      up              0% 4.06%            0       0
GigabitEthernet0/0/12          up      up              0% 4.06%            0       0
…output omit…




Verify that only interface Gigabit Ethernet 0/0/2 is disabled on R1 and that only

interface GigabitGigabitEthernet 0/0/1 is disabled on R3.


<R1>display ip interface brief
…output omitted…
GigabitEthernet0/0/1                10.0.12.1/24         up           up
GigabitEthernet0/0/2                10.0.23.1/24         *down        down
…output omitted…


<R3>display ip interface brief
…output omitted…
GigabitEthernet0/0/1                10.0.12.3/24         *down        down
GigabitEthernet0/0/2                10.0.23.3/24         up           up
…output omitted…



Step 4 Enable the DHCP function.

The DHCP service is not enabled by default, enable the DHCP service on the
router(s).
[R1]dhcp enable


[R3]dhcp enable

Step 5 Create a global IP address pool

Create an address pool named pool1 for R1 and pool2 for R3. Configure attributes
for pool1 and pool2, including address range, egress gateway, and IP address lease
period.
[R1]ip pool pool1
Info: It's successful to create an IP address pool.
[R1-ip-pool-pool1]network 10.0.12.0 mask 24


                                                      HUAWEI TECHNOLOGIES                  Page96
[R1-ip-pool-pool1]gateway-list 10.0.12.1
[R1-ip-pool-pool1]lease day 1 hour 12
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]dhcp select global


[R3]ip pool pool2
Info: It's successful to create an IP address pool.
[R3-ip-pool-pool2]network 10.0.23.0 mask 24
[R3-ip-pool-pool2]gateway-list 10.0.23.3
[R3-ip-pool-pool2]lease day 1 hour 12
[R3]interface GigabitEthernet 0/0/2
[R3-GigabitEthernet0/0/2]dhcp select global


Run the display ip pool name <name> command on the router to view the
assigned IP address pool configuration parameters.
<R1>display ip pool name pool1
  Pool-name          : pool1
  Pool-No            :0
  Lease              : 1 Days 12 Hours 0 Minutes
  Domain-name        :-
  DNS-server0        :-
  NBNS-server0       :-
  Netbios-type       :-
  Position           : Local             Status             : Unlocked
  Gateway-0          : 10.0.12.1
  Network            : 10.0.12.0
  Mask               : 255.255.255.0
  VPN instance       : --


----------------------------------------------------------------------------       Start   End       Total
Used Idle(Expired) Conflict Disable
----------------------------------------------------------------------------
   10.0.12.1         10.0.12.254     253       0          253(0)               0     0
----------------------------------------------------------------------------




Configure the default management interface for S1 to request an IP address from

the DHCP server (R1). Perform the same steps on S2 for R3.


[S1]dhcp enable


                                                      HUAWEI TECHNOLOGIES                        Page97
[S1]interface Vlanif 1
[S1-Vlanif1]ip address dhcp-alloc


