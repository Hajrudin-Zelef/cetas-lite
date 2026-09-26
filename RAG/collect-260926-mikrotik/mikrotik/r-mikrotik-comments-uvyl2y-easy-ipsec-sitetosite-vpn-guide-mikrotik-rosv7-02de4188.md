---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-uvyl2y-easy-ipsec-sitetosite-vpn-guide-mikrotik-rosv7-02de4188
title: "r-mikrotik-comments-uvyl2y-easy-ipsec-sitetosite-vpn-guide-mikrotik-rosv7-02de4188"
domain: mikrotik
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/lot-mikrotik/forum/ipsec/r-mikrotik-comments-uvyl2y-easy-ipsec-sitetosite-vpn-guide-mikrotik-rosv7-02de4188.md
source_anchor: ""
source_lines: [1, 123]
sha256: cbe26f8af9bba76f1472042ed0041a851055bcbc54d20f3f509282164dbe94e1
---

# r-mikrotik-comments-uvyl2y-easy-ipsec-sitetosite-vpn-guide-mikrotik-rosv7-02de4188

☁️Easy IPSEC Site-To-Site VPN Guide, MikroTik ROSv7☁️
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
I've been using IPSec to set up site-to-site VPNs for like ever, both using RouterOS, and also using other firewall products. But since RouterOS 7 supports Wireguard out of the box, I really don't see a need to fiddle around with IPSec for S2S.
True, I would actually recommend Wireguard wherever possible. Although IPSEC still has many practical uses such as connecting to cloud providers or to other vendor's equipment that does not support wireguard.
Currently, the ciphers in Wireguard are done purely in software; the IPSec ciphers can be hardware accelerated. On devices with hw acceleration, you can get better performance with IPSec than with Wireguard.
Hey people of Reddit, hope you guys are all doing great. I know IPSEC on MikroTik tends to be very confusing at times. So I did my best to update my Site-To-Site VPN guide to ROSv7 and go through all the different steps that we need to take when configuring IPSEC. So we do cover a bit of theory and configure IPSEC on two MikroTiks via Winbox and CLI. Feel free to leave suggestions or if there are other things you guys would love to see on MikroTik.
I am considering also going into some other video types like talking about Network Design vs Network Implementations. Though that is more or less a general network thing and less a MikroTik thing.
I was just trying to figure this out myself last night. Excellent!
Why is it always a video? It's not that difficult to type up a list of instructions. It's like those companies that post PDF's to their web sites for people to download for everything, rather than just posting the text in the page. Or even worse, word documents.
You could think of this as show and tell, some people learn concepts a bit better having a visual aide, if it's not your liking the Mikrotik docs which I use as a reference when making these videos tend to give everything in a pure text format.
The main issue with the documents is that they don't actually give good examples. They're usually 'Here is all the information *splat*'. When I get in a situation like this, I'm trying to just get something up and running, and I'll go into depth later. (Yes, I do read the documentation. I've even rewritten documentation in the past (unrelated))
Man I wish Mikrotik supported IPSec virtual tunnel interfaces. It would make it super simple. As it is I use GRE+IPSec because it's easier to setup but it doesn't work with things like AWS.
