---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-15oyf6u-firewall-rules-to-block-everything-except-allowed-38f9dc09
title: "r-ubiquiti-comments-15oyf6u-firewall-rules-to-block-everything-except-allowed-38f9dc09"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-15oyf6u-firewall-rules-to-block-everything-except-allowed-38f9dc09.md
source_anchor: ""
source_lines: [1, 164]
sha256: e92056e7d433702649a7819da468876e63ab414c4b7560f0e2e97f16b463b82b
---

# r-ubiquiti-comments-15oyf6u-firewall-rules-to-block-everything-except-allowed-38f9dc09

[supprimé]
      [deleted by user] 
        
     
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
The firewall comes like this anyway. You just need to add rules for anything you want to allow. And then internal stuff of course. P.S. theyve changed all this now with traffic rules
Firewall rules usually process in a top to bottom routine.
The general pre-defined bottom rule is "deny any any" - so if a rule doesn't already exist to allow the traffic, it will block it.
Unfi udm is probably not the best platform to start to learn on.... Maybe EdgeOS instead
So there's really no point in messing with Internet in rules because everything I need is already there in the form of the predefined rules?
Hmm, then what about LAN in? By default, it seems to allow anything. Would my rules make sense if I put them there? Or is that also not really needed?
Then what kind of rules do people even add if the default rules completely suffice?
You could use LAN I. Rules to block transit from one vlan or subnet -coming IN from interface vlan10 for instance - transiting to port 22 on the firewall.
That's a bit simplified but the basis of it
Unifi doesn't block inter-vlan traffic by default (in most cases) so that would be a use case
I’ll take a wack at this. So usually by default for most gateway/security devices they filter all ports from the internet anyway so a block all inbound is redundant most of the time. The only time this is not true is if you have does any port forwarding or intentionally exposed a port to the internet such as VoIP or VPN.
The only time a port opens is when a device from the inside creates a new connection to a host on the internet. The router will make note of that connection and allow or not allow the traffic biased on the rules set then set the connection as established. In your rule for allowed ports you don’t have New connections checked so while your allowing established connections your hosts can’t make new connections to the internet.
Your better off blocking the ports you want going outbound to the internet matching all states.
Thanks. That information was helpful.
A couple of notes.
Inbound ports are blocked by default
For example if the packet is sourced from outside of the network, it won’t make it through unless you have a rule that specifically allows it
Outbound ports are allowed by default
for example a user opens a browser and connects to a website on 80 or 443, this is allowed. You could go in and block LAN to WAN traffic on 80 and 443 and you’d effectively block Web access
There are a huge number of ports that get used behind the scenes so by doing a blanket statement of block all, you’re going to have a hell of a time getting things to work.
25 is SMTP, are you actually using SMTP? That’s really only for a server not a mail client.
You blocked port 80 which is standard HTTP traffic.
Lastly, I know you said “more control” but what does that mean to you? What are you actually trying to accomplish?
Thanks for your post.What I'm after is mainly two things: Learn about networking and secure my network so open ports can't be used to gain access to my network by a malicious third party.
What I mean by more control is the very limited range of functions and features you get from routers from your ISP (at least in my country). You can't even subnet with those devices or configure any firewall rules. That's why I bought my own router.
Gotcha. Ok, if I were you, I’d start reading about networking, look at CompTIA Network + it’s a pretty basic course and covers basic day to day networking.
Then or if you don’t want to do that and just be hands on. Find a port you want to block and block just the one port. This way you can troubleshoot; “ I enabled this rule and everything breaks” that’s tough to work through.
Remember, from the outside, where malicious content comes from is blocked by default unless you or a user “let” it in by requesting the packet from the LAN. Understanding the source of the packet is really important.
There are several good channels and guides on youtube. I would start here - https://www.youtube.com/watch?v=PyDpLJIKg0M - which is a discussion of the UDR but has very good top level coverage of VLANs and firewall rules to limit traffic between VLANs and things like IOT devices and the internet.
There are other channels like "The Hookup" and "Crosstalk Solutions" that dive into Ubiquiti. Also McCann Tech - https://evanmccann.net/ - has very good tutorial information.
Lots of good info out there. Enjoy
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Like others have mentioned, the default behaviour on incoming traffic is to block everything.
A fun exercise is to set up egress filtering.
Block everything outgoing, then open up ports/services for the different hosts on your network.
