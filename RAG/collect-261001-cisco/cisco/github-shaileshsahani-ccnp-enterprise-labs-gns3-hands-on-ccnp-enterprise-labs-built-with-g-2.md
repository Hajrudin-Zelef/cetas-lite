---
id: collect-261001-cisco/cisco/github-shaileshsahani-ccnp-enterprise-labs-gns3-hands-on-ccnp-enterprise-labs-built-with-g-2
title: "1. Install GNS3"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2026-08-09"]
keywords: ["copyright", "license", "mit license"]
source: docs/RAG/collect-261001-cisco/github-shaileshsahani-ccnp-enterprise-labs-gns3-hands-on-ccnp-enterprise-labs-built-with-gns3-coveri.md
source_anchor: ""
source_lines: [152, 213]
sha256: f0c716b04eff2ddb4119684277f9f98eb0ace86d22cdb778bbc3599e3a076139
---

# 1. Install GNS3

```
c7200-adventerprisek9-mz.152-4.S.bin
i86bi-linux-l3-adventerprisek9-15.5.2T.bin   (IOSv)
vios_l2-adventerprisek9-m.ssa.high_iron_20190423.bin   (IOSvL2)
asav981-7.qcow2   (ASAv, optional)
```
```
# 1. Install GNS3
sudo add-apt-repository ppa:gns3/ppa
sudo apt update
sudo apt install gns3-gui gns3-server
# 2. Install a hypervisor
sudo apt install qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils virt-manager
# 3. Clone the repo
git clone https://github.com/ShaileshSahani/CCNP-Enterprise-Advanced-Labs.git
cd CCNP-Enterprise-Advanced-Labs
# 4. Import Cisco images in GNS3
# Edit → Preferences → IOS routers → New → follow the import wizard
```
Then: pick a lab → read its README → build the topology in GNS3 → apply configs → verify → intentionally break something → fix it.

| Level | Labs | Routers | Focus | 
|---|---|---|---|
| Intermediate | 1–10 | 3–5 | Core routing, redistribution | 
| Advanced | 11–26 | 4–6 | BGP, VRF, multi-tenancy | 
| Expert | 27–45 | 4–6 | MPLS, WAN, security | 
| Architect | 46–50 | 3–8 | Automation, full integration | 

1. **Physical** — interfaces up/up?
2. **Layer 2** — VLANs, trunking, STP state
3. **Layer 3** — IP addressing, routing table
4. **Layer 4** — ACLs, firewall rules
5. **Application** — protocol-specific debugs

- CCNA-level networking knowledge
- TCP/IP and subnetting fluency
- Comfort with Cisco IOS CLI
- Basic GNS3 experience
- Working knowledge of routing protocol fundamentals

**References:** CCNP Enterprise ENCOR 350-401 & ENARSI 300-410 Official Cert Guides, TCP/IP Illustrated Vol. 1, Cisco Learning Network, GNS3 Documentation

- Lab structure defined (50 labs)
- README published
- Individual lab documentation (in progress)
- GNS3 project files per lab
- Configuration templates
- VXLAN EVPN labs
- SD-WAN concepts
- Segment Routing (SR-MPLS)
- IPv6 enterprise design
- Terraform for network automation
- CML/VIRL compatibility

**Shailesh Sahani**
Network Technician → Network Engineer | CCNP Enterprise (in progress)

If this helped your CCNP journey: star the repo, fork it, or open an issue with feedback.

MIT License — Copyright (c) 2026 Shailesh Sahani. See LICENSE for full text.

*Last updated: August 9, 2026*
