---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-nlo0rp-need-help-firewall-nat-rules-32c19484
title: "r-mikrotik-comments-nlo0rp-need-help-firewall-nat-rules-32c19484"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-nlo0rp-need-help-firewall-nat-rules-32c19484.md
source_anchor: ""
source_lines: [1, 132]
sha256: 6f15ddf28aef2817eb23a6078535c5cc63cef3307015284f013fb116bca67482
---

# r-mikrotik-comments-nlo0rp-need-help-firewall-nat-rules-32c19484

Need Help - Firewall // NAT Rules 
        
    Hello All,
I am trying to set up a site-to-site from an RB4011 to HAP aC.
I followed this post. https://www.reddit.com/r/mikrotik/comments/kn9t9y/mikrotik_site_to_site_ipsec_tunnel_both_dynamic_up/
But cannot figure out the firewall and NAT rules to make this work. I have it set up on Dynamic IP so I cannot allow any specific public IP.
Thank you
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
I found out that one of the routers was missing a default rule that was making me chase my tail. I got it working.
My configuration for a static ip uses this:
IP -> Firewall -> Filter Rules: +New
General -> Chain: input
General -> Src Address: 87.65.43.21
In Interface: ether1-gateway
Action -> Action: accept
You should be able to change the src to a url and use the cloud dynamic dns.
I tried putting the URL in and it says it must be an IP :(
Pretty sure they all have to be IPs. If you just need to resolve the IP at creation time, then
/ip fir fil add src-address=[:resolve google.com] ...will get the first DNS result forgoogle.comand use that.
There's several promising articles here:
https://www.google.com/search?q=mikrotik+vpn+dynamic+ip&oq=Mikrotik+VPN+dy&aqs=chrome.1.69i57j0j0i22i30l3.36642j0j9&client=ms-android-att-us-revc&sourceid=chrome-mobile&ie=UTF-8
