---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-jy8n8j-bridge-nat-vs-ip-firewall-nat-d494b2fd
title: "r-mikrotik-comments-jy8n8j-bridge-nat-vs-ip-firewall-nat-d494b2fd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/r-mikrotik-comments-jy8n8j-bridge-nat-vs-ip-firewall-nat-d494b2fd.md
source_anchor: ""
source_lines: [1, 127]
sha256: db37ce7b3d3e3248efa3f6143731b081fd7d6cf60398b7bad6a92f584f13896e
---

# r-mikrotik-comments-jy8n8j-bridge-nat-vs-ip-firewall-nat-d494b2fd

[supprimé]
Bridge NAT vs IP Firewall NAT
o What is the difference between using Bridge NAT and IP Firewall NAT ?
o Why it is recommended to use IP Firewall NAT then linking public WAN and LAN interfaces ?
Haven't found deep explanation about that in wikis..
      Bridge nat or firewall nat ? - MikroTik
Bridge - RouterOS - MikroTik Documentation
    
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
tl'dr: L2/MAC/Bridging = same network, L3/IP/Routing = different networks. LAN and WAN are different networks.
Bridge NAT is for NATting L2 traffic only. You are actually not translating IP addresses but MAC addresses. That is an action which you rarely need to do except some very special situations (e.g. hijacking traffic)
If you are "linking" public WAN and LAN interfaces, it is called "routing" and these interfaces are in separate L2 zones so it needs to be L3 forward. Your WAN is not aware what IPs are on your LAN, so the source IP of the traffic (e.g. your computer's IP) must be translated to your router's WAN IP. That way your ISP knows, where goes the reply packet. Once it arrives to your router from WAN, it will be translated back to your computer's IP and delivered through LAN.
I understand it may be bit confusing with all these layers, but if you are willing to learn, I would recommend to read this document. It is really well written: https://www.cloudflare.com/en-au/learning/ddos/glossary/open-systems-interconnection-model-osi/
Thanks! Bridge NAT could be labelled as "L2 NAT" in Winbox...
What about that "Use IP firewall" in bride? When do we actually need to enable it?
When you want to filter traffic in the same subnet, or want to limit throughput
