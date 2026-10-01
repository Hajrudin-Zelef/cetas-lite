---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1196065-how-to-adopt-ubiquiti-gear-into-unifi-os-http-on-8080-not-work-25bfa468
title: "questions-1196065-how-to-adopt-ubiquiti-gear-into-unifi-os-http-on-8080-not-work-25bfa468"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1196065-how-to-adopt-ubiquiti-gear-into-unifi-os-http-on-8080-not-work-25bfa468.md
source_anchor: ""
source_lines: [1, 40]
sha256: 51239fa4dc219af1e56007d5a3bf658820d5e8862d02af3074525bb793d05495
---

# questions-1196065-how-to-adopt-ubiquiti-gear-into-unifi-os-http-on-8080-not-work-25bfa468

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm running UniFI OS Server inside a qemu VM to manage the UniFI APs and Ubiquiti switches on my local network, but cannot get it to "adopt" them. I restored a backup from the UniFI OS running on my laptop which does work, but I don't want to run it there because I want this as an always-on service that is accessible over the management network.
In the UniFI OS Network Console, devices show as "Adopting" and then indicate they failed to connect and display "Click to Resolve"
Everything is on the same physical network/VLAN (VLAN 1)
Router: 10.0.2.1/24
UniFi OS Server VM: 10.0.2.3
UniFi switches/APs: using DHCP from 10.0.2.100 - 10.0.2.200 range
All devices can ping each other. I am able to SSH into the devices from the UniFI OS server and the router. Over SSH, I can make HTTP requests on port 80 to the UniFI server (curl -v http://10.0.2.3/ works).
However the inform URL on port 8080 (curl -v http://10.0.2.3:8080/inform) never responds. I see the requests in tcpdump, but the UniFI OS server never responds.
To make things more confusing, the exact same request from any other routed network (e.g., 10.2.0.0/24, 10.2.201.0/23) works — I get the expected HTTP 400 from the controller.
Inside UniFi OS, lsof shows port 8080 being handled by slirp4netns, which likely is the issue, because nginx (which I installed as a test) on port 80 works just fine...
Why would local-subnet traffic fail while routed traffic works?
After literal months of debugging I found the root cause: the 10.0.2.0/24 network is used internally by podman/slirp4netns which are components of UniFI OS.
The simplest fix was to move my management network to a different CIDR range 🙃
Specifically:
slirp4netns defaults to using 10.0.2.0/24 as its internal network for rootless containers.
UniFi OS runs its controller containers through podman + slirp4netns, which means inside the namespace it brings up interfaces like:
10.0.2.2 (gateway)
10.0.2.x for the container network
This is the same subnet I was using on my physical network for VLAN 1 (10.0.2.0/24).
When traffic arrives on 10.0.2.3:8080 from anywhere outside that subnet, slirp handles the NAT/port-forward normally and the container replies.
When traffic arrives from a host inside 10.0.2.0/24, slirp gets confused because the source IP collides with its internal network. The packet reaches the VM, but slirp refuses to NAT/forward it.
I did find this post on the UniFI forum which indicated that the range can be changed, but I didn't want to mess with that...
If your network or any of your remote site2site networks use 10.0.2.0/24, you're going to have trouble connecting.
The reason is that podman uses a tool called slirp4netns which basically helps set up rootless networking and by default is uses 10.0.2.0/24 as the internal network for containers.
I used the GlennR script to install unifi-os-server, ran uosserver stop, created /home/uosserver/.config/containers/containers.conf with the below contents. Customize the cidr address to something that you're not already using.
network_cmd_options=["cidr=10.30.0.0/24"]
Then I ran uosserver start and then I was able to connect as expected.
I installed this on Debian 13 Trixie, so these references helped me
TL;DR: Don't use 10.0.2.0/24 for your management network when using UniFI OS!
This command showed me that port 8080 was being used and provided a corresponding PID:
command
netstat -ano -p tcp
To see what is using that PID, open Task Manager, click the Processes tab, then More Details, and finally the PID column. If you don't see the PID column, right-click and select it.
In my case, the PID corresponded to a file called AgentService.exe. I Googled that name and discovered that it was associated with the Minitools program, Shadowmaker (for partition management), which I had recently installed. I uninstalled the program and was able to run the Unifi Network Server program.
