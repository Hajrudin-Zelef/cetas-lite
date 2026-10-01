---
id: collect-261001-general-networking/general-networking/brians-tech-corner-blog-blob-head-content-blog-2026-01-01-why-i-switched-to-unif-775bae17-2
title: "brians-tech-corner-blog-blob-head-content-blog-2026-01-01-why-i-switched-to-unif-775bae17"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-general-networking/brians-tech-corner-blog-blob-head-content-blog-2026-01-01-why-i-switched-to-unif-775bae17.md
source_anchor: ""
source_lines: [171, 253]
sha256: cf514729cdd1cd912315141061204dd8a0bbd2975a46592044d4ef95d1a3c0d3
---

# brians-tech-corner-blog-blob-head-content-blog-2026-01-01-why-i-switched-to-unif-775bae17

UniFi handled:
- Inter-VLAN routing
- DHCP per network
- Firewall enforcement
This also explained something that confused me early on:
- Some services worked from Wi-Fi
- The same services failed from wired clients
Different networks, different rules and now I could see that.
One of the most practical benefits of moving to UniFi and VLANs was finally solving the IoT security problem.
On a flat network, every IoT device had full access to everything:
- Smart plugs could reach my laptop
- Security cameras had access to file servers
- Random gadgets could talk to infrastructure
This isn't theoretical paranoia it's terrible security posture.
With UniFi, I created a dedicated IoT network:
- VLAN ID: 40
- Subnet: 192.168.40.0/24
- Default behavior: No access to anything else
Then I added explicit firewall rules for what IoT devices actually need:
Allow: IoT VLAN → Internet (outbound only)
Allow: Default LAN → IoT VLAN (so you can manage devices)
Block: IoT VLAN → Default LAN
Block: IoT VLAN → Kubernetes VLAN
Suddenly, IoT devices were:
- Reachable from my phone/laptop (when I need to configure them)
- Able to reach their cloud services
- Completely isolated from everything else
Here's where it gets interesting: Home Assistant bridges the gap securely.
I run Home Assistant on a Raspberry Pi on the Default LAN, but with a critical exception:
- A firewall rule allows the Pi's MAC address to communicate with the IoT VLAN
- This is the only device on the Default LAN with that privilege
- IoT devices cannot reach the Default LAN or Kubernetes VLAN directly
- My phone/laptop talks to Home Assistant (on the LAN), which then controls IoT devices
The flow looks like this:
Phone (Default LAN) → Home Assistant (Default LAN) → IoT Devices (IoT VLAN)
                      [Firewall exception by MAC address]
Home Assistant becomes the controlled access point, not just a dashboard.
With firewall rules in place, I can:
- Control lights, plugs, and sensors through Home Assistant
- Block IoT devices from scanning the network
- Let IoT devices reach their cloud APIs (if needed)
- Monitor everything from the UniFi controller
If a device gets compromised, it's trapped on VLAN 40 with no way to reach anything else.
If you're running infrastructure at home, IoT isolation isn't optional:
- Cheap IoT devices have terrible security
- Your Kubernetes cluster is a high-value target
- Mixing trusted and untrusted devices on the same network is asking for trouble
VLANs let you keep the convenience without the risk.
Home Assistant works just as well (better, even), and your network is no longer a free-for-all.
Once things were configured correctly, I stopped guessing and started validating.
Some checks I used constantly:
- Can I ping a Kubernetes node from my desktop?
- Can I reach a NodePort from outside the VLAN?
- Does ingress respond when I force the Host header?
Example:
curl http://192.168.30.120:31929 -H "Host: argocd.k8s.local"
If that worked, the network was correct.
If it didn’t, Kubernetes wasn’t the problem yet.
That single check saved me hours of debugging in higher layers.
After the migration, I had:
- Clear network boundaries
- Predictable routing
- VLAN-aware switching everywhere
- The ability to debug ingress and NodePorts properly
- A foundation that behaved like real infrastructure
Most importantly, I stopped blaming Kubernetes for network problems that weren't Kubernetes problems.
- Consumer networking gear hides problems
- VLANs force clarity (and that's a good thing)
- Managed switches are non-negotiable for labs
- If networking is flaky, everything above it is noise
- UniFi doesn't prevent mistakes it makes them visible
For reference, my setup currently includes:
- UniFi Gateway (routing, firewall, VLANs)
- UniFi managed switches (core and access)
- UniFi access points for Wi-Fi coverage
The specific models aren’t critical, what matters is that every hop in the path is VLAN-aware and centrally managed.
Think of this post as the network foundation.
Everything that follows: Proxmox bridges, VLAN trunks, Kubernetes ingress, and storage only works cleanly because the network is no longer guessing.
With UniFi in place, I could finally move forward confidently:
- Proxmox networking stopped being mysterious
- Kubernetes nodes behaved predictably
- Ingress, storage, and GitOps started to make sense
Check out our Kubernetes Homelab series of posts if interested in setting up your own homelab.
