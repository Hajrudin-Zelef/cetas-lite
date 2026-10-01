---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-20
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [4466, 4675]
sha256: db1e094aea1fa0d85eca75824f6789a141e336fdd4e0c1c5d07a492ec4aa1361
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

<S1>display ip interface brief
…output omitted…
Interface                      IP Address/Mask      Physical     Protocol
MEth0/0/1                              unassigned                 down         down
NULL0                                  unassigned                 up           up(s)
Vlanif1                                10.0.12.254/24             up           up




Verify that this address was taken from the DHCP pool named pool1 on R1,

and for S2, from the DHCP pool named pool2 on R3.



<R1>display ip pool name pool1
  Pool-name          : pool1
  Pool-No            :0
  Lease              : 1 Days 12 Hours 0 Minutes
  Domain-name        :-
  DNS-server0        :-
  NBNS-server0       :-
  Netbios-type       :-
  Position         : Local                Status           : Unlocked
  Gateway-0          : 10.0.12.1
  Network            : 10.0.12.0
  Mask               : 255.255.255.0
  VPN instance       : --


----------------------------------------------------------------------------
     Start           End           Total Used Idle(Expired)       Conflict Disable
----------------------------------------------------------------------------
   10.0.12.1     10.0.12.254        253       1         252(0)           0          0
----------------------------------------------------------------------------




 <R3>display ip pool name pool2
  Pool-name          : pool2
  Pool-No            :0
  Lease              : 1 Days 12 Hours 0 Minutes



                                                        HUAWEI TECHNOLOGIES             Page98
  Domain-name        :-
  DNS-server0        :-
  NBNS-server0       :-
  Netbios-type       :-
  Position           : Local             Status             : Unlocked
  Gateway-0          : 10.0.23.3
  Network           : 10.0.23.0
  Mask               : 255.255.255.0
  VPN instance       : --
 ----------------------------------------------------------------------------
          Start             End        Total Used Idle(Expired)        Conflict Disable
 ----------------------------------------------------------------------------
       10.0.23.1     10.0.23.254     253       1         252(0)                 0         0
 ----------------------------------------------------------------------------



Ensure that global pool configuration has been completed for both R1 and R3
before continuing!

Step 6 Create an interface based IP address pool

Disable the interface GigabitGigabitEthernet 0/0/1 R1. For R3 disable interface
Gigabit Ethernet 0/0/2.
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]shutdown


[R3]interface GigabitEthernet 0/0/2
[R3-GigabitEthernet0/0/2]shutdown


Configure an interface address pool to allow the clients connected via Gigabit
Ethernet 0/0/2 of R1 to obtain IP addresses. Perform the same operation for
GigabitGigabitEthernet 0/0/1 of R3. Do not enable these interfaces, as we do not yet
wish to activate the DHCP service on the network.
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]dhcp select interface




[R3]interface GigabitEthernet 0/0/1
[R3-GigabitEthernet0/0/1]dhcp select interface
Isolate addresses from the pool GigabitEthernet0/0/2 for R1, and the pool
GigabitEthernet0/0/1 for R3, for DNS services. Additionally, set the IP address lease


                                                      HUAWEI TECHNOLOGIES                     Page99
period for the interface address pool.
[R1-GigabitEthernet0/0/2]dhcp server dns-list 10.0.23.254
[R1-GigabitEthernet0/0/2]dhcp server excluded-ip-address 10.0.23.254
[R1-GigabitEthernet0/0/2]dhcp server lease day        1 hour 12


[R3-GigabitEthernet0/0/1]dhcp server dns-list 10.0.12.254
[R3-GigabitEthernet0/0/1]dhcp server excluded-ip-address 10.0.12.254
[R3-GigabitEthernet0/0/1]dhcp server lease day 1 hour 12


Run the display ip pool interface command on the router to view the configured
parameters of the interface address pool. For R3 the interface is
GigabitGigabitEthernet 0/0/1.


<R1>display ip pool interface GigabitEthernet0/0/2
  Pool-name          : GigabitEthernet0/0/2
  Pool-No            :1
  Lease              : 1 Days 12 Hours 0 Minutes
  Domain-name        :-
  DNS-server0        : 10.0.23.254
  NBNS-server0       :-
  Netbios-type       :-
  Position           : Interface        Status              : Unlocked
  Gateway-0          : 10.0.23.1
  Network           : 10.0.23.0
  Mask               : 255.255.255.0
  VPN instance       : --
 ----------------------------------------------------------------------------
          Start           End        Total Used Idle(Expired)        Conflict Disable
 ----------------------------------------------------------------------------
       10.0.23.1          10.0.23.254   253      0          252(0)              0       1
 ----------------------------------------------------------------------------



Flush the existing Vlanif1 address from S2 to allow for dynamic allocation of a new

IP address from the interface GigabitEthernet0/0/2 pool.


[S2]interface Vlanif 1
[S2-Vlanif1]shutdown
[S2-Vlanif1]undo shutdown




                                                      HUAWEI TECHNOLOGIES                   Page100
Enable interface Gigabit Ethernet 0/0/2 to allow the DHCP server to become active
on the network and to begin sending DHCP discover messages.
[R1]interface GigabitEthernet0/0/2
[R1-GigabitEthernet0/0/2]undo shutdown


<R1>display ip pool interface GigabitEthernet0/0/2
  Pool-name          : GigabitEthernet0/0/2
  Pool-No            :1
  Lease              : 1 Days 12 Hours 0 Minutes
  Domain-name        :-
  DNS-server0        : 10.0.23.254
  NBNS-server0       :-
  Netbios-type       :-
  Position           : Interface           Status             : Unlocked
  Gateway-0          : 10.0.23.1
  Network            : 10.0.23.0
  Mask               : 255.255.255.0
  VPN instance       : --
 ----------------------------------------------------------------------------
          Start             End        Total Used Idle(Expired)       Conflict Disable
 ----------------------------------------------------------------------------
      10.0.23.1     10.0.23.254      253       1          251(0)           0               1
 ----------------------------------------------------------------------------
<S2>display ip interface brief
…output omitted…
Interface                              IP Address/Mask          Physical        Protocol
MEth0/0/1                                    unassigned               down             down
NULL0                                        unassigned               up               up(s)
Vlanif1                                      10.0.23.253/24           up               up



The interface Vlanif1 shows to have been allocated an address from the
GigabitEthernet0/0/2 address pool of R1.
Flush the existing Vlanif1 address from S1 to allow for dynamic allocation of a new
IP address from the interface GigabitEther0/0/1 pool.


[S1]interface Vlanif 1
[S1-Vlanif1]shutdown
[S1-Vlanif1]undo shutdown


Enable interface GigabitGigabitEthernet 0/0/1 to allow the DHCP server to become


                                                      HUAWEI TECHNOLOGIES                      Page101
active on the network and to begin sending DHCP discover messages.


[R3]interface GigabitEthernet 0/0/1
[R3-GigabitEthernet0/0/1]undo shutdown



Verify that the new IP address as been allocated from the interface pool.


