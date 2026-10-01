---
id: collect-261001-meraki/meraki/r-meraki-comments-qlh95c-meraki-layer-3-firewall-rules-not-working-7595f03e
title: "r-meraki-comments-qlh95c-meraki-layer-3-firewall-rules-not-working-7595f03e"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["sol"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-qlh95c-meraki-layer-3-firewall-rules-not-working-7595f03e.md
source_anchor: ""
source_lines: [1, 154]
sha256: c72ffe180b119c5fa31f507e6304a72cb5d712f4cbe01cfde45aeb4e4590551a
---

# r-meraki-comments-qlh95c-meraki-layer-3-firewall-rules-not-working-7595f03e

[supprimé]
      Meraki Layer 3 Firewall rules not working 
        
     
     Désolé, cette publication a été supprimée par son auteur.  
         
        
        
        
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Do you have any group policies? Are they applied to any devices or VLAN's? Are the devices whitelisted?
Also, did you set the protocol to "Any", same with the source & destination ports (should be set to any).
So here is something that catches me, a lot of times I apply a group policy and accidentally set the option to use the policies firewall rules, which override the MX firewall rules. I also think globally whitelisting with the group policy will override the MX rules.
Check your devices in clients to make sure there are no policy’s selected, and check your vlans on the Mx to make sure there isn’t a default policy selected.
Is it wired or wireless? The Wireless tab has it’s own firewall rules as well
Wired, and not even using Meraki APs for this site.
Are the devices gateways pointing to the firewall or another L3 switch?
Firewall, they picked up an IP via DHCP from the firewall and then we’re statically assigned.
If an IP isn't allowed to access anything out of its own subnet then it's a deny all rule for that IP in your access list.
It does still need internet access, it just should not be allowed to access any other local subnets in this particular network (which are all defined in the deny rule).
What protocol are you denying? I'm guessing all? Also as a troubleshooting step, might be worth it to create an explicit allow above everything for an example of traffic you want to be blocking then re-test - check the counters, if it goes up you know the traffic is definitely traversing the MX and not some other way
Yes, denying all.
might be worth it to create an explicit allow above everything for an example of traffic you want to be blocking then re-test - check the counters, if it goes up you know the traffic is definitely traversing the MX and not some other way. Alternatively you could do a packet capture
How sure are you that traffic routes through the MX between those VLANs?
Are any of those subnets across the SD-WAN? Firewall rules on MX don't apply to SD-WAN traffic; there's a separate section of ACLs in the SD-WAN page for that.
Positive, they are all local subnets/VLANs that exist only on that MX Firewall. None of the subnets are across the SD-WAN.
I’ve seen this issue before too. It seems ICMP is allowed but all TCP and UDP traffic is blocked. Not sure why that is, I didn’t dig into it after the pets were blocked
ICMP is a different protocol to TCP or UDP, so a block any IP rule won't catch it.
I am going to go out on a limb here, but were you already running the pings? If so, those flows were probably allowed already from your test source and destinations.
Start a new test to a different (new) destination or stop your tests for about 15 minutes and the flows should expire.
Once new flows are established, the rules should apply just fine.
I could be wrong, but I'm fairly sure the Meraki default rule blocks all inter-vlan routing
The default is allow all.
Commentaire supprimé par le membre
do you mind sharing the link
Commentaire supprimé par le membre
Actually this might be what my problem is. I’m going to test pinging IPs on a different subnet that are not an MX IP. Thanks!
