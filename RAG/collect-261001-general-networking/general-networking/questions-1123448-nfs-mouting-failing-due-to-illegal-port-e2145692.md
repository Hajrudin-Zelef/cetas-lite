---
id: collect-261001-general-networking/general-networking/questions-1123448-nfs-mouting-failing-due-to-illegal-port-e2145692
title: "NFS Mouting Failing due to illegal port"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1123448-nfs-mouting-failing-due-to-illegal-port-e2145692.md
source_anchor: ""
source_lines: [1, 29]
sha256: bd8891edbae66a9fb63068df4d5085e9218bd960a5caf7512de7329fccde3144
---

# NFS Mouting Failing due to illegal port

*Score : 1 | Source : https://serverfault.com/questions/1123448/nfs-mouting-failing-due-to-illegal-port*

I have a VM machine that has a public IP interface and a private IP interface. The private interface is assigned 192.168.50.78. Then I have a dedicated host that acts as my "router" using private IP 192.168.50.1 and this is, therefore, my gateway for my 192.168.50.0/24 local network. The "router" I have is OPNsense.
That being said, I am trying to mount within my VM to my two remote NFS servers. The mount fails because of an illegal port coming from my VM, which is behind NAT:
[root@NFSHOST ~]# tail -f /var/log/messages | grep "rpc.mountd"
Feb 21 22:31:40 NFSHOS rpc.mountd[28721]: refused mount request from 197.189.XXX.ZZZ for /data/secondary (/data/secondary): illegal port 33744
Feb 21 22:35:22 NFSHOS rpc.mountd[28721]: refused mount request from 197.189.XXX.ZZZ for /data/secondary (/data/secondary): illegal port 40085
And this is due to port translation happening. A solution is to add insecure to my /etc/export file, which I have tested and can confirm does indeed work.
However, as it suggests, it insecure and highly not recommended. I have tried adding a port forwarding rule on OPNsense and for some reason, it remains to fail with illegal ports.
Is there anyway I can solve this issue, PS I cannot use "bridge" mode either I need to set some sort of rule in OPNsense to handle this accordingly and I am not exactly sure.

---

### Reponse (acceptee) — score 0

Managed to finally solve this specific issue. I was told here: OPNsense Link
The following:
IIRC NFS servers expect the source ports coming from clients to be under 1024 to be considered secure. Make sure you check the "Static-port" on your NAT outbound rule to prevent the firewall from modifying the source port on TCP and UDP packets.
I ended up adding a NAT outbound rule on OPNsense for the entire local LAN with static port enabled, and it solved the issue. I can now mount NFS via secure port ranges.

---

### Reponse — score 0

For me the problem was the server did not support older protocol version. You can either force a client to use other version like this:
mount -o vers=3 -t nfs 123:/
This applies to standard linux client of course. If you can not influence a client like this check this post how to have other protocol support in server side:
