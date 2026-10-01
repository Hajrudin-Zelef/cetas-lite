---
id: collect-261001-general-networking/general-networking/questions-1158744-hetzner-ubuntu-dns-resolution-not-working-8d08f1e3
title: "hetzner ubuntu dns resolution not working"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1158744-hetzner-ubuntu-dns-resolution-not-working-8d08f1e3.md
source_anchor: ""
source_lines: [1, 40]
sha256: ad146b2219361c584644a5896c560cc8381fa73a4b927fec41c113b9f25048b1
---

# hetzner ubuntu dns resolution not working

*Score : -1 | Source : https://serverfault.com/questions/1158744/hetzner-ubuntu-dns-resolution-not-working*

I created a natgateway using opnsense
I created a VM in a private network which is behind the natgateway so that all internet traffic will pass through opensense into the client VM
I can ping google ip (142.250.190.142) through opnsense which works fine
but I can't ping the domain name www.google.com and also I can't update OS using apt update
I get error this
temporary failure in name resolution
I am using hetzner cloud provider for VMs
the archecture of hetzner is that de default gateway take the first ip address since servers are not connected directly. The flow is like this (client server > gateway > opnsense)
Interface config on the client VM
/etc/network/interfaces
auto lo
iface lo inet loopback
auto enp7s0
iface enp7s0 inet dhcp
   post-up ip route add default via 10.0.0.1
   dns-nameservers 185.12.64.2 185.12.64.1
IP address
opnsense router 10.0.0.2
VM 10.0.0.3
gateway 10.0.0.1

---

### Reponse (acceptee) — score 0

I edited this file /etc/systemd/resolved.conf and added dns

---

### Reponse — score 0

I edited this file /etc/systemd/resolved.conf and added dns
You don't want to do that.
If you want to stick with the debian networking stuff (i.e., /etc/network/interfaces; I never really know what it is called; ifup scripts?), delete the current /etc/resolv.conf (which should be a symlink), and use touch to recreate a empty regular file at the same path. Then restart this networking stuff, or maybe just reboot if you don't know how. After the reboot you should be able to see the nameservers being added to /etc/resolv.conf by this debian networking stuff.
Otherwise, stop using and disable this stuff and switch to systemd-networkd (and systemd-resolved).
Make sure you revert the change to /etc/systemd/resolved.conf in either case.
