---
id: collect-261001-cisco/cisco/r-networking-comments-ifme0c-cisco-9500-stackwise-virtual-720e218d
title: "r-networking-comments-ifme0c-cisco-9500-stackwise-virtual-720e218d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/r-networking-comments-ifme0c-cisco-9500-stackwise-virtual-720e218d.md
source_anchor: ""
source_lines: [1, 47]
sha256: 1cb6f38e38758dd29ec4bcbbc99a1019305fffce4459675c454cc72c991f82c4
---

# r-networking-comments-ifme0c-cisco-9500-stackwise-virtual-720e218d

Cisco 9500 Stackwise virtual 
        
    Hello all,
In short I need to set up stackwise virtual on a pair of 9500 switches which I've racked, but not turned on, configured, or stacked yet.
I've checked release notes and software versions, and all is well regarding that side of things, and the switches are compatible.
Now I've never done this before, and I've wrote a script for this based on the cisco configuration template and the SVL is an active QSFP+ cable, and I was wonderring if anyone can give me a sanity check before I go and do this?
My mental logic after reading the notes are;
Stack the switches with the factory default configuraiton. Then assign the managememnt ip and port information, as you would with a normal set of switches which had stacking cables and stacking modules.
Here is the script which I will run on both switches;
######## Stackwise Virtual Configuration to be run on both switches ########
en
switch 1 renumber 2
switch 1 priority 5
conf t
stackwise-virtual
domain 2 												************** Change This if multiple stackwise virtual domains in deployment
end
show stackwise-virtual									************** verify stackwise-virtual configuration
write memory
reload
############ Configuration of Stackwise Virtual Link ############
en
conf t
interface FortyGigabitEthernetX/X/X 				   ************** Change This to the 40GB interface
stackwise-virtual link 1
end
write memory
reload
Section des commentaires
The first reload, after configuring the domain, is not really needed. You ca still reload the switches, but i don't think it will do anything. Also, it's best practice to configure a keepalive link as well. And lastly, VSW support varies with exact switch model and IOS version. (Only the latest versions support the high performance models and usinf the network module ports for VSW).
Let me knwo if you want me to look for the script i am using for this.
If you have a script I can have a look at that would be great!
Regarding the keep alive link, would that just be another point to point link between the switches similar to a VPC keep alive on a Nexus?
As for switch and version it's a C9500-48YC and I was going to upgrade it to Gibraltar-16.12.3a on to it as that it is the latest gold star release for this model of switch.
Sorry, I should've said dual-active detection link. Pretty similar to keepalive but it doesn't need IP's.
That IOS sounds good, we plan to deploy the same in one of the sites.
I use the below script on both switches to configure the virtual stack pairs. I then configure priorities and numbers (there's probably a better approach).
Whatever you do, with these C9500s, don’t take one out of the stack when powered off and then try and reconfigure it. Fantastic way to brick them.
Really? Brick them? Christ. ok. Noted hahaha thanks for letting me know
You don't really need the renumbering bit to be honest.
Configure stackwise domain and Stackwise-virtual links on both devices, reload the one you want to be switch #1, wait until it boots and reload the one which you want to be switch #2.
Plus don't reload until you have both domain and Stackwise-virtual links configured.
DAD links can be configured after the stack is formed (reload will be needed on both devices for it to take effect).
I know this is an old post but if you have production networks on a VSL setup and do DAD afterward is there a downtime? I know you have to reboot the switches but if VSL is working you should be able to fail over to other VSL “member” and keep things up right? I made this mistake and just realized I don’t have DAD setup. We did a code upgrade before moving networks over to this 9606 VSL pair and didn’t have any downtime so I’m assuming rebooting one of the 9606 wouldn’t cause a down time if the other takes over?
Commentaire supprimé par un membre de l’équipe de modération
That's something to check out.
Stackwise-virtual on 9500 High Performance models is only supported from 16.10.1 onwards, so if your device comes from factory with a lower version you need to upgrade it first before configuring Stackwise-virtual.
