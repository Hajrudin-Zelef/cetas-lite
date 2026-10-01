---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-tkfxj4-vlans-on-ubiquiti-and-vmware-esxi-493a7490
title: "r-ubiquiti-comments-tkfxj4-vlans-on-ubiquiti-and-vmware-esxi-493a7490"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-tkfxj4-vlans-on-ubiquiti-and-vmware-esxi-493a7490.md
source_anchor: ""
source_lines: [1, 145]
sha256: 3fdd9dd01549a7177ec216e4c803935fb40b7ae800368d31f4c66cdfb7ff0ff3
---

# r-ubiquiti-comments-tkfxj4-vlans-on-ubiquiti-and-vmware-esxi-493a7490

VLANs on Ubiquiti and VMware ESXI 
        
        
        
    
    
    Hello folks. I am planning to start building a home lab with ESXI and the UDM Pro, however, as much as I understand how to create VLANs on the UDM, I still do not understand how to link those VLANs to ESXI’s vSwitch, and ESXI‘s VLANs. Please help.
 
     Publication archivée. Impossible de voter et de publier de nouveaux commentaires.  
        
        
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
There are multiple ways to get this done. Single vlan port groups and vswitches that match 1-1 on Ubiquiti or using a trunk port group (vlan 4095) on ESXi and allowing all vlans out from the UniFi switch that connects to the ESXi host. I’m using both styles throughout my network and both work equally as easy. 1-1 is simple and basic but trunk port groups are great for advanced config.
This is good info. I am hoping to fully utilize the 4x1Gbit network ports on the servers, however, using trunk seems to make more sense, especially since I will have a switch between the UDM and the servers.
What version of ESXi do you have? Only the higher end versions support VLANs. If you don't have the higher end versions, then you take NIC1 on the server and plug it into a port. Assign the VLAN to that port in the UDM. The port on the server, belongs in a vswitch labelled after that VLAN.
For example, let's say you have a host with 4 NICs. When physically looking at the server NIC0, 1,2, and 3. They should be labelled. NIC0 and 1 should be vSwitch0 for management, NIC2 could be vSwitch1 (IOT) or label it how ever you want, NIC3 could be vSwitch2 (LAN). on the dream machine, what ever port NIC0 and 1 are in, should be your default VLAN1. Port 2 should be assigned the IOT VLAN, and port 3 should be assigned LAN VLAN.
I am hoping to experiment with VLANs in the 60 day trial of ESXI 7. So I was not sure how the hypervisor VLANs are configured to work properly with Ubiquiti VLANs.
From what you’re saying I guess I can setup tagged or untagged VLANs and bind them to a NIC, whereas on the Ubiquiti, I bind the VLAN to a port on the switch.
Nice!
The way I describe it is how you do it if you don't have ESXi licensing. If you do, there's a field with a VLAN you just enter and the switch VLAN needs to match for the ports that vswitch is plugged into.
u/whitedragon551 the free version of ESXi DOES support VLAN tagging. VLAN tagging is a function of the standard switch. The standard switch is included with free ESXi.
u/AnonymusChief You don't tag nics in ESXi. You apply VLAN tags to Virtual Machine port groups and/or vmkernel ports.
Thanks u/cerealkillerzz. This is good info.
