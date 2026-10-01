---
id: collect-261001-general-networking/general-networking/questions-1123097-routing-issue-on-debian-11-vm-83eb5c65
title: "Routing issue on Debian 11 VM"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1123097-routing-issue-on-debian-11-vm-83eb5c65.md
source_anchor: ""
source_lines: [1, 40]
sha256: e6a32ff69a9d0acbfbdb0b0150811ad4d566bf4899f8a36a6f4bfac068d3276b
---

# Routing issue on Debian 11 VM

*Score : 1 | Source : https://serverfault.com/questions/1123097/routing-issue-on-debian-11-vm*

I updated my OPnsense "router/gateway" to 23.1 a few days ago, and now I am experiencing an issue with one of my machines. I have one machine with OPnsense installed that acts as a basic NAT router for my private subnet 192.168.50.0/24. This machine/router uses the private IP 192.168.50.1, which is the gateway for my 192.168.50.0/24 network.
I have a physical KVM machine with 1x Public NIC and 1x Private NIC. On this machine, I have a VM running also equipped with 1x Public IP and then 1x Private IP falling in the 192.168.50.0/24 subnet as its private network. I am trying to mount an NFS share on this VM. I have two NFS shares and only ONE NFS share mounts successfully while the second refuses.
This is what the NFS mounts look like:
Successful Mount
root@s-145-VM:~# mount -t nfs -vvv 102.165.XXX.YYY:/data/secondary /mnt/SecStorage/test
mount.nfs: timeout set for Thu Feb 16 14:07:58 2023
mount.nfs: trying text-based options 'vers=4.2,addr=102.165.XXX.YYY,**clientaddr=197.189.XXX.YYY**'
root@s-145-VM:~#
Failed Mount:
root@s-145-VM:~# mount -t nfs 102.165.XXX.ZZZ:/data/secondary /mnt/SecStorage/test
mount.nfs: access denied by server while mounting 102.165.XXX.ZZZ:/data/secondary
root@s-145-VM:~# mount -t nfs -vvv 102.165.XXX.ZZZ:/data/secondary /mnt/SecStorage/test
mount.nfs: timeout set for Thu Feb 16 14:07:11 2023
mount.nfs: trying text-based options 'vers=4.2,addr=102.165.XXX.ZZZ,**clientaddr=192.168.50.53**'
mount.nfs: mount(2): Operation not permitted
mount.nfs: trying text-based options 'addr=102.165.XXX.ZZZ'
mount.nfs: prog 100003, trying vers=3, prot=6
mount.nfs: trying 102.165.XXX.ZZZ prog 100003 vers 3 prot TCP port 2049
mount.nfs: prog 100005, trying vers=3, prot=17
mount.nfs: trying 102.165.XXX.ZZZ prog 100005 vers 3 prot UDP port 892
mount.nfs: mount(2): Permission denied
mount.nfs: access denied by server while mounting 102.165.XXX.ZZZ:/data/secondary
As you can see, in the failed mount attempt, the VM passes its private IP address, 192.168.50.53 when it tries to mount the second NFS. And not its public IP 197.189.XXX.YYY. And the reason this is happening is that the VM automatically after each reboot adds the following line to its IP routing table:
102.165.XXX.ZZZ via 192.168.50.1 dev eth1
When I delete this entry from the routing table, the NFS mounts successfully.  Something in the OPnsense router or Linux kernel is causing this behaviour and making it create this entry in the IP route table. And I have no idea what it could be. What puzzles me, it only creates this entry for this IP and not for the other NFS, which is successful and they are basically on the same network space 102.165.XXX
I have double-checked the Opnsense server and rebooted it many times and I cannot pinpoint what is causing this behaviour in the VM.
Any suggestions for this will be extremely appreciated, please.
EDIT:
So after some more investigation, the IP route that is added (102.165.XXX.ZZZ via 192.168.50.1 dev eth1) is 100% correct and expected. When I remove this entry from the IP route, the VM tries to access NFS via the gateway of the public IP (default route), and then it works.
So essentially, it means that somehow my router is not passing the traffic accordingly. But if I run a ping test from my router its successful, see here:
# /sbin/ping -4 -S '192.168.50.1'  -c '3' '102.165.XXX.YYY'
PING 102.165.XXX.YYY (102.165.XXX.YYY) from 192.168.50.1: 56 data bytes
64 bytes from 102.165.XXX.YYY: icmp_seq=0 ttl=60 time=1.140 ms
64 bytes from 102.165.XXX.YYY: icmp_seq=1 ttl=60 time=1.200 ms
64 bytes from 102.165.XXX.YYY: icmp_seq=2 ttl=60 time=1.348 ms
So what am I missing? Should I add some sort of route manually to the KVM host maybe?
