---
id: collect-261001-general-networking/general-networking/r-networking-comments-ha4cmp-need-ospf-troubleshooting-tips-and-sites-e6d325d0
title: "r-networking-comments-ha4cmp-need-ospf-troubleshooting-tips-and-sites-e6d325d0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-ha4cmp-need-ospf-troubleshooting-tips-and-sites-e6d325d0.md
source_anchor: ""
source_lines: [1, 138]
sha256: e7721a08f5892415ae8f80b24ee742d0d1e1ef07b1466ec717957256ad7da2d1
---

# r-networking-comments-ha4cmp-need-ospf-troubleshooting-tips-and-sites-e6d325d0

Need OSPF Troubleshooting Tip's and sites 
        
    My OSPF network has thankfully been very stable over the past years. So my troubleshooting skills have become rusty.
Does anyone have a fav website, whitepaper for troubleshooting OSPF issues they would like to share?
Yes, I have the cisco OSPF stuff, but some blogs and sites have a much better 'take' on things than cisco. Yes, I do have cisco equipment. Thanks!
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
I’ll save you the read. If you have an OSPF problem probably your MTU.
It's always the MTU... Network troubleshooting should be: DNS -> load balancers -> MTU -> traffic selectors / encryption domain -> route filters, in that order.
Would you take the time to explain this reasoning to someone still learning networking? I know OSPF is a link- state IGP, what points to the issue being the mtu so clearly?
If you OSPF with any regularity you'll know its picky about MTU. If there is MTU mismatch you won't establish and its all too common with overhead situations on WAN.
In most situations today OSPF config is basic. I'm sure other situations exist but I use it exclusively as an underlying IGP for BGP routing. IS-IS is an alternative.
Here is one of my Juniper routers OSPF configurations. Fairly basic and only talking to one neighbor.
area 0.0.0.0 { interface lo0.0 { passive; } interface xe-0/1/5.0; interface xe-0/1/0.VLAN; interface irb.VLAN { metric 10; } }
Take a look at Nick Russo’s Cisco Live OSPF troubleshooting presentation, video recordings and github repo: https://github.com/nickrusso42518/ospf_digrst2337
Watch this session from Cisco Live: https://www.ciscolive.com/global/on-demand-library.html?search=Troubleshooting%20OSPF#/session/1542224324396001rr0i
I don't think it get's much better than this one.
INE CCIE RS Workbook.
I always use networklessons.com to brush up on my skills however you do need to subscribe, i think you can get a free trial if that helps.
Otherwise I usually just google any specific questions and there is normally a forum post which helps.
They spent a shitload of money on SEO it seems, their site is literally in the first 5 in google output for ANY network related query and it’s frustrating af. I had to write a greasemonkey script just to hide this damn site from google results.
Just out of curiosity, what is wrong with this website in your humble opinion ?
Thanks.
I also google a few topics to brush up my skill and their website came up. Do you think it is worth it to get their subscription. I am also preparing for CCNP ENCOR will it covers those topics as well. Thanks
I think its worth it but i use it most days studying towards CCIE and previously for my CCNP, the trial is actually 7 days for $1 so you can always have a look before going for the monthly/yearly subscriptions.
They always let anyone read the first parts of the lessons for free and the lessons update frequently with python being the newest lessons available last time i checked.
Depends. I used them for a couple months but I dropped them just because I didn't think they went deep enough for what I am looking for (CCIE).
That being said, they are a great intro to most topics and they give great examples. I think for the CCNP they would be worth it.
