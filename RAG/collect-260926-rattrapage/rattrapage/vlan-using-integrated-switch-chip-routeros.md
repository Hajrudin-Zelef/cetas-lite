---
id: collect-260926-rattrapage/rattrapage/vlan-using-integrated-switch-chip-routeros
title: "vlan-using-integrated-switch-chip-routeros"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/vlan-using-integrated-switch-chip-routeros.md
source_anchor: ""
source_lines: [1, 55]
sha256: 095a6d6d5e671b3757d1d90a082d7dd9ada96cce704ef6a35874917bf4efa494
---

# vlan-using-integrated-switch-chip-routeros

Recently i had to board this topic in a config and finally i think i understand how this work:

i will try to do this as a tutorial with 3 vlans for integrated switch on hap and rb951 series not for CRS, this tutorial was tested on rb951ui rb951g and hap lite and works ok:

In this case ether1 is the master port for ether2 to ether5

1. On switch add vlan0 with VID 0 as your native vlan, i chose vlan0 but it can be any number you want, add all ports you want to work with that vlan as native vlan and the switch cpu port to ensure management and default gateway functionality of the router for that vlan, in my case i want the native vlan to work on all ports
2. add any other vlan you need and add the ports where you want that vlan to work (tagged) in and the switch cpu port to ensure management and default gateway functionality of the router for that vlan, in my case i added vlan10 with VID 10 and vlan 20 with VID 20 only want ether1 and switch cpu to be in that vlan, ether1 will be like the trunk port
 
 
Continue in next post i cannot add more images

 
                
            
           
          
            
            
              continuing:

my vlan list looks like this:


now configure the ports:


First thing to do is configure native vlan as default vlan on all ports you want, including switch cpu port in this case all ports use vlan0 as native vlan.

Then configure vlan header= always strip on accessports using only one vlan for end devices in this case ether2 to ether5

Configure vlan header= leave as is on trunk ports and switch cpu port, in this case ether1 is a trunk

and finally configure vlan mode=secure to enforce your configuration

Now add your vlan interfaces to configure router ip address using master port of the switch as the physical interface for this vlans:


from now on you can do what you want with your vlans, for example add a vlan to a bridge to another interface like virtual ap to use this vlan on a separate wireless lan, or configure dhcp server etc etc.

i invested several hours trying to understand this, i hope this can help somebody to do vlans quickly and take advantage of this nice functionality.

any doubt, errata or comment please comment i will be happy to answer

             
            
           
          
            
            
              PD note:

BE AWARE RB750GR3 DOES NOT SUPPORT VLANS ON THE SWITCH CHIP

LOOKS LIKE CABLE TEST DOES NOT WORK ON RB750GR3
