---
id: collect-261001-meraki/meraki/r-meraki-comments-snlxoh-suddenly-sites-are-blocked-on-meraki-firewall-dcaaa953
title: "r-meraki-comments-snlxoh-suddenly-sites-are-blocked-on-meraki-firewall-dcaaa953"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-snlxoh-suddenly-sites-are-blocked-on-meraki-firewall-dcaaa953.md
source_anchor: ""
source_lines: [1, 152]
sha256: bf03296a0233cd27ad8577a75a2955ad98f47ca93cbb52348b809e52f09fd0bb
---

# r-meraki-comments-snlxoh-suddenly-sites-are-blocked-on-meraki-firewall-dcaaa953

Suddenly sites are blocked on Meraki Firewall 
        
    Hi guys,
Something strange is happening at my company, suddenly some sites are not accessible for the end-users for example discord.
And no this site is not blocked within the rules of the meraki Firewall.
Is there anybody who knows what the cause of this could be?
Kind regards
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Check your content filtering settings. Check the organization change log to see if someone changed something. Call meraki support.
Are you certain they're being blocked? Not accessible != blocked.
What have you checked/tried already?
I have checked the content filtering settings, searched for any changes on meraki but I can't find anything....
It's discord so I doubt it's not accessible
I can access discord when using a Proxy site... Something strange is happening
Discord is working fine for me, and we are all Meraki - MX, Switches, WAPs. You may be having some other issue or someone added some sites to your block list perhaps? Are you the only admin for your environment?
Not sure if you fixed the issue yet, but we decided to reboot our MX during off hours and everything worked afterwards.
I will try that after the working hours.
If you have updated your firmware recently, they switched their content filtering to use Cisco Talos. You'll want to to go talosintelligence.com/and search for the specific sites that are being blocked to see what they classify them as. It could be something that is causing the issue.
I know that most ppl will figure it out without this, but there is a typo in the link up there. It should be https://talosintelligence.com/
Guys, I fixed the issue:
when checking out my wifi settings I saw something strange, the DNS was set to 127.0.0.1 . After I set it back to automatic (DHCP) everything is working fine again. It was indeed a DNS problem.But it doesn't explain why some people in my company have their DNS set to 127.0.0.1 suddenly....
Are you using Cisco umbrella roaming client? That acts as a dns proxy, changing your dns to 127.0.0.1 and then will forward the requests out to umbrella out encrypted over 443. If you don’t have it installed, I believe Cisco Anyconnect has this functionality as well but not sure if it handles it in the same way.
In services.msc you can check for Umbrella Remote Client.
Yeah I can see umbrella in my services.
Now I understand why the DNS was set 127.0.0.1 thanks for the information!
I actually recently had something similar happen at a couple of different client sites. The clients started complaining about the firewall blocking some sites they need access to, double checked and Meraki wasn't configured to block those sites. It turned out they had recently renewed their Comcast internet contract, and SecurityEdge was added to the contract. SecurityEdge includes a bunch of content filtering.
Don't know if that applies to you, but might be worth checking.
If you upgraded to the latest beta, it changes the data sourced used for category blocking.
I’m having the same is here. A bunch of websites became unreachable as of this morning. Still trying to figure out why. Let me know if you get answers
DNS?
long capable subtract cautious existence forgetful divide grandiose run fear
This post was mass deleted and anonymized with Redact
This is the way!
Sorry, too tempting 🙈
If it is being blocked it will show up in event log or security center where you can whitelist or change settings. Otherwise your trouble may be upstream. Check Downdetector?
I’m having the same issue as OP. Nothing is being logged in either event logs and security center. Just waiting on support to answer.
Did you check your IDS in addition to content filtering?
Event log isn't going to show you the layer 7 hits for some dumb reason so maybe look into the list of countries you're blocking.
It happened to me when I updated to RC16.15
Facebook network (Facebook, Instagram ecc) wasn't solving their IP addresses and it wasn't my DNS. After rolling back to the stable version, it started to work again
We recently had to readd newer defintions for our URL/Content filtering after Meraki did an update.....So looking forward to ripping it all out for something else...
