---
id: collect-261001-meraki/meraki/questions-797523-azure-hub-multiple-spoke-vpn-using-meraki-mx-security-appliance-3900f7f1
title: "questions-797523-azure-hub-multiple-spoke-vpn-using-meraki-mx-security-appliance-3900f7f1"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "research"]
source: docs/RAG/collect-261001-meraki/questions-797523-azure-hub-multiple-spoke-vpn-using-meraki-mx-security-appliance-3900f7f1.md
source_anchor: ""
source_lines: [1, 12]
sha256: a3852ac9c6ef6ac7021cac2cad8ad1603485d1480393670accdc3ddc6339e9f4
---

# questions-797523-azure-hub-multiple-spoke-vpn-using-meraki-mx-security-appliance-3900f7f1

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I want to set up various infrastructure in MS Azure that will then be available to multiple locations that are equipped with Cisco Meraki MX Security Appliances. Unfortunately, the MXs don't yet support route based VPNs, and Azure only supports multiple site to site networks when using route based VPN. I think similar challenges may exist with AWS and other cloud service providers.
I think I may be able to work around this limitation using a virtual firewall, such as Cisco ASAv, but I haven't been able to find any documentation or marketing material that makes it clear this is suitable. I know I have done hub/spoke VPN with physical ASAs in the past, but I have no experience with ASAv.
Has anyone got any experience doing cloud provider hub with ASAv (or any other virtual firewall) and branch office spoke using firewalls that don't support IKEv2 or route based VPNs, such as Meraki MX, Cisco ASA etc?
As mentioned above, we were able to accomplish this by standing up a Cisco CSR in Azure. We have 50 MX60W's and a few MX100 all connecting into the Azure CSR which then allows a direct connection to our Azure virtual servers.
Of course the best solution would be standing up a virtual MX in Azure. Our Meraki sales rep keeps promising that this is coming but no news yet. He mentioned recently that they are in beta with a virtual MX in AWS. With all focus on setting up cloud-based hosting environments (i.e., Azure, AWS), I think Meraki is missing out on how many companies want to connect all of their locations seamlessly.
You'll need a static IP on the CSR, but can use the Meraki dynamic DNS names. The Meraki VPN is setup in the Organization wide VPN section, and distributed to the MXs based on tags. The Phase 1 and 2 and pre-shared key all have to match exactly on both sides.
