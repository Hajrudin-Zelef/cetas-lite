---
id: collect-261001-huawei/huawei/r-networking-comments-15vxvkc-huawei-firewall-usg6525e-vlan-issue-67727ff1
title: "r-networking-comments-15vxvkc-huawei-firewall-usg6525e-vlan-issue-67727ff1"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-15vxvkc-huawei-firewall-usg6525e-vlan-issue-67727ff1.md
source_anchor: ""
source_lines: [1, 16]
sha256: 67b5b13aae578022df60c2401c61d7cda7bf94716755253e09536e76c13da183
---

# r-networking-comments-15vxvkc-huawei-firewall-usg6525e-vlan-issue-67727ff1

Huawei Firewall USG6525E VLAN issue 
        
        
        
    
    
    
      Hey there, I am not really sure if this is the most appropriate mean to ask about this, but the support from the official Huawei page is not that helpful anyway. I am an intern in a data center and I was assign to set up a Huawei Firewall from two different clients. One of these clients have a VLAN assigned to them and the other one has two different VLANs.
I created these three VLANs in the Network > Interface > Add Interface > Type: VLAN.
Therefore for client number one, I assigned its VLAN to a physical interface in an access mode. And for the other client I assigned both VLANs to other interfacces in Trunk mode.
But I can't connect to any of those VLANs. If I ping to a server within those networks I can't reach them from the Firewall. Am I missing something? I don't really have a lot of experience in networking. I've been more a of programming kind of guy, but the people from the office are not quite helpful to me either. I'd appreciate any type of hint you guys could give me on how to solve this.
    
Section des commentaires
You're pinging a vlan without assign an address or something that is pingable?
You're an intern you should talk to your supervisor about this. You're an intern for a reason ask them.about it.
You need to do intervlan routing. This is a Cisco example but the concept is similar.
