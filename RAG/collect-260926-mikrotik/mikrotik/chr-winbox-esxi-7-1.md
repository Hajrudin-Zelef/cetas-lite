---
id: collect-260926-mikrotik/mikrotik/chr-winbox-esxi-7-1
title: "chr-winbox-esxi-7"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel", "license", "licenses", "memory", "throughput"]
source: docs/RAG/lot-mikrotik/tools/chr-winbox-esxi-7.md
source_anchor: ""
source_lines: [1, 237]
sha256: 3d12fb8f5c18e18062105ede2bf35c7fe859efe57de800abeb65204c1f95669e
---

# chr-winbox-esxi-7

Hello.. Apologies if this has already been discussed but couldn’t find anything.

Currently have around 650 users ppp on CCR 1072, with max CPU usage around 15%

Tried setting up R820 server with ESXI 7 and CHR 6.48.6

All users are connecting well and I can access and manage the end users router and also run btest without any problems.

The only problem I’m facing is that the CHR winbox doesn’t load! It signs in but blank, no interfaces or logs visible and disconnects. However users are uninterrupted.

Switch has DAC cable to the server 10g intel NIC. Switch MTU 10218

MTU on VSwitch ESXI is 9000 and interface CHR is 1550, ospf/ibgp 1530 and all works well.

Loopback bandwidth tests also shows good results upto 20G

Hyper threading is disabled on the r820 bios. Any suggestion or feedback?

             
            
           
          
            
            
              At winbox window clear cache before connecting and try again.

             
            
           
          
            
            
              Clearing cache works, but once I close the window and try to login again same problem white windows. Every time I have to clear cache

Any suggestions? Or solutions?

             
            
           
          
            
            
              Update: Also noticed the moment traffic reaches 600-700Mbps I can not access the winbox at all. Approximately 800 ppp active accounts

             
            
           
          
            
            
              I run many virtual CHR ROS routers ( both v6.xx and v7.xx ) on several VmWare ESXi servers.

They work well and they are fast and reliable.

Those CHRs can only hit 10-GIg , because I only have 10-Gig nics on my physical VmWare ESXi servers hosting these CHRs.

I use vmnic3 ethernet interfaces.

North Idaho Tom Jones

             
            
           
          
            
            
              Hi Tom, thank you for your response. I would like to know if you have any CHR configured as ppp server with 500+ users and how is it performing?

I have read in forums that CHR works solid for routing [bgp etc], but I am facing a challenge with ppp services.

My server specs are as follows:

Dell PowerEdge R650xs

CPU = Dual Intel® Xeon® Gold 6330 2G, 28C/56T, 11.2GT/s, 42M Cache, Turbo, HT (205W) DDR4-2933

HDD = SSD

Memory = 4 x 16GB RDIMM, 3200MT/s

Network = Intel X710-T2L Dual Port 10GbE BASE-T Adapter

             
            
           
          
            
            
              
Does VMWare show high CPU/mem for CHR at this load?  That be the most telling if your connection problems are related to load.  I think this is a winbox protocol thing, not related have 800 users on a beefy CHR.  But just a guess.

My initial thought is winbox is using a higher MTU of CHR to send its packets, which a PC at MTU 1500 doesn’t like it.  PMTUD should kick in, but if ping is blocked, it can’t work.  Also depending on the config if ICMP is bogging down/queued/etc, or worse blocked by the firewall, winbox wouldn’t be able to figure out the lower MTU for TCP.

Other thing is I always enable promiscuous mode on the VMWare interfaces to CHR.  Not sure, but if disabled, ESXi might not like some of the L2 features of winbox.

             
            
           
          
            
            
              If you have a VmWare ESXi server that is showing high CPU usage - try this :

- In the BIOS on your physical VmWare ESXi server , disable hyper-threading.
- Limit the total number of CPUs assigned to virtual servers to have a combined CPU usage of one less than the actual of physical CPU cores you have.
- Use Vmnic3 network interfaces where possible.
- Use paravirtual interfaces where possible
- Use thin hard disks
- Use physical 10-Gig ( or 100-Gig ) network cards instead of 1-GIg network cards.

If you have multiple physical VmWare ESXi servers ;

- group your virtual servers together as much as possible to avoid traffic from one VmWare ESXi server going to another Physical VmWare ESXi server , then going back.  ( Don’t ping-pong your traffic between VmWare ESXi servers.

Also - If you are strong on networking , take a look at Delayed-Ack in your network settings in your physical and virtual servers.  Sometimes , this can get you a 500+ percent increase in network I/O throughput.

**North Idaho Tom Jones**

             
            
           
          
            
            
              Hello Tom, Hypervisor is ESXI 7.0.3.. I have used the ova template and customized hardware as per the screenshots below. vswitch mtu 9000 connected to Mikrotik Switch L2 MTU SFP+ 10218 using the DAC 10g cable

I have applied all your suggestions that I have read across other posts also. Can you guide or share a post which shows where I can disable the delayed_ack on the ESXI please. Once I manage to do that then I want to schedule maintenance window and try again to migrate the users. VM is off for now.

vswitch screenshot:


ESXI Host specs:


VM harware screenshot:


             
            
           
          
            
            
              Question , is the total CPU count for all virtual machines less than the total cores you physically have ?

Question , what is the license level you have on your CHR ?

Also - try changing your MTU to 1500 for everything ( VmWare ESXi server ) and your network switches and any MTU settings in vm machines.

This is from my notes I have compiled over the years:

In the VmWare ESXi shell , here are some delayed_ack commands:

To enable delayed_ack:

**vsish -e set /net/tcpip/instances/defaultTcpipStack/sysctl/_net_inet_tcp_delayed_ack 1**

To disable delayed_ack ( default ):

vsish -e set /net/tcpip/instances/defaultTcpipStack/sysctl/_net_inet_tcp_delayed_ack 0

The above two commands take place right away - no vmware esxi reboot needed.



To have VmWare ESXi reboot with delayed_ack enabled - you need to vi ( edit a file )

vi this file:	/etc/rc.local.d/local.sh

and place this line inside the local.sh file

vsish -e set /net/tcpip/instances/defaultTcpipStack/sysctl/_net_inet_tcp_delayed_ack 1


vsish -e set /net/tcpip/instances/defaultTcpipStack/sysctl/_net_inet_tcp_delayed_ack 1

-and-

vi this file:	/etc/rc.local.d/local.sh

add this or change the file to contain the line below

vsish -e set /net/tcpip/instances/defaultTcpipStack/sysctl/_net_inet_tcp_delayed_ack 1




Let me know if anything here helped.

**North Idaho Tom Jones**

             
            
           
          
            
            
              Answer = CPU: There is only 1 VM with the specs I shared configured on this host, nothing else. Specs I shared in earlier screenshot of the host and VM.

Question =  Can you suggest CPU allocation from the host specs I shared and cores per socket?

Answer = License: vSphere 7 Enterprise Plus // CHR License = P10

MTU is as follows:

ESXI vSwitch = 9000

Physical eth1 interface on CHR = 1550

OSPF VLAN = 1530

VPLS = 1508 [carrying users ppp from remote sites]

Screenshot of eth1 on CHR


Switch ports connecting the edge and CHR L2MTU = 10218

OSPF vlan edge to CHR = 1530

I got the delayed_ack from your other posts and have applied that.. Appreciate you sharing the same here.

             
            
           
          
            
            
              On my CHRs , I use P-Unlimites licenses.

Some additional things I do:

- I do not use any CHRs for do-everything routers,
- I always have two vmnic3 interfaces to each of my CHRs.  One for the primary input interface and the second interface is normally 802.1q trunk Vlan interfaces.

In my CHR ISP/Wisp/fiber/GPON networks , each of my CHRs perform a dedicated function ( only one function ).

