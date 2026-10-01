---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-lanswitch-17-1-1-b-lans-swit-03793f34
title: "c-en-us-td-docs-routers-ncs4200-configuration-guide-lanswitch-17-1-1-b-lans-swit-03793f34"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-lanswitch-17-1-1-b-lans-swit-03793f34.md
source_anchor: ""
source_lines: [1, 52]
sha256: 565cc3431b4782f213452c229b1fc38329e25fbd0d0ee0d8b16327e9537f9424
---

# c-en-us-td-docs-routers-ncs4200-configuration-guide-lanswitch-17-1-1-b-lans-swit-03793f34

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
The Multiple Spanning
Tree Protocol (MSTP) is an STP variant that allows multiple and independent
spanning trees to be created over the same physical network. The parameters for
each spanning tree can be configured separately, so as to cause a different
network devices to be selected as the root bridge or different paths to be
selected to form the loop-free topology. Consequently, a given physical
interface can be blocked for some of the spanning trees and unblocked for
others.
Having set up multiple
spanning trees, the set of VLANs in use can be partitioned among them; for
example, VLANs 1 - 100 can be assigned to spanning tree 1, VLANs 101 - 200 can
be assigned to spanning tree 2, VLANs 201 - 300 can be assigned to spanning
tree 3, and so on. Since each spanning tree has a different active topology
with different active links, this has the effect of dividing the data traffic
among the available redundant links based on the VLAN - a form of load
balancing.
By default, MSTP is
disabled on all interfaces. MSTP need not be enabled explicitly on each
interfaces. By turning the global configuration on, it is enabled on all
interfaces.
Configures the encapsulation. Defines the matching criteria that maps the ingress dot1q or untagged frames on an interface
for the appropriate service instance.
Step 8
l2protocol peer stp
Example:
Router (config-if-srv)# l2protocol peer stp
Configures STP to peer with a neighbor on a port that has an EFP
service instance.
Step 9
end
Example:
Device(config-mstp-if)# end
Returns to
privileged EXEC mode.
Configuration
Example
This example shows how to
configure STP to peer with a neighbor on a service instance.
interface GigabitEthernet0/0/0
no ip address
negotiation auto
service instance trunk 10 ethernet
encapsulation dot1q 10-20
bridge-domain from-encapsulation
!
service instance 1024 ethernet
encapsulation untagged
l2protocol peer stp
bridge-domain 1024
!
end
