---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-14dwdb7-show-hidden-policies-in-gui-5f781a08
title: "r-fortinet-comments-14dwdb7-show-hidden-policies-in-gui-5f781a08"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-14dwdb7-show-hidden-policies-in-gui-5f781a08.md
source_anchor: ""
source_lines: [1, 159]
sha256: ab0fcb3cf7e0d5b9c1c6836d7660ff7a610befbf2d1b5790e85a03275210dd53
---

# r-fortinet-comments-14dwdb7-show-hidden-policies-in-gui-5f781a08

Show hidden policies in GUI? 
        
    Hi everyone,
Just doing an audit of our policies and I have found a handful of policies that only show in the CLI. My question is, is there a CLI command to unhide hidden policies in the GUI?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Have you tried to reset the columns in the GUI while on the policy page?
Yep, no filters, etc. I started looking at unused policies as the security rating service pointed out a handful of policies that were not used in the last 90 days and found a few that only showed in the CLI and not the GUI.
Any commonality you can find? I know they certain objects were/are GUI hidden. For example an ISDB group used to be hidden from the gui
Yes, I just did. I starting to wonder if they are left-over orphaned policies. As mentioned in this thread, the whole reason I am looking into these is that the security rating service has shown a few policies that have not been used in over 90 days, and those lower numbered policies (I assume created around the time of initial deployment) fall into this category.
Sorry if this has been answered already, but am I correct in saying that certain things appear in the policy list in the CLI that are built in, but don't show in the GUI?
I've never heard of firewall policies that were hidden on a fortigate. If they are proxy policies, there is a different screen for those... (same with ZTNA polices and local-in policies)
I've had central SNAT rules not show in the GUI (I believe they were corrupted / invalid and idk whether they worked) so I do believe it.
For me it happened using the API, somehow a policy wasn't updated or created right on 6.4.x firmware
Have you try "diagnose debug config-error-log read" If there are errors in the configuration, they will be displayed there In one version I had policy errors, after a reboot everything was fine again.Unfortunately, I can't remember exactly which version it was.
Can you show us a snipit of the cli. Obviously block the actual address but showing us the general output would help to understand.
Sure. Here is an example of one. The "hidden" policies are lower in the number range. I have changed values to keep it a bit more non-descriptive
config firewall policy
edit x
set uuid xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx
set srcintf "vw2"
set dstintf "vw1"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set service "ALL"
set utm-status enable
set ssl-ssh-profile "my-certificate-name"
set av-profile "my-av"
set webfilter-profile "my-filter"
set ips-sensor "my-ips-sensor"
set application-list "my-application-list"
next
end
Is this a virtual wire pair policy? Those are in a different menu in the GUI
Thanks!
Which FortiOS?
If you filter in a column and the remove that column, the filter remains. Add every column and you’ll find the filter
Thanks for the tip. I tried that and the lower police numbers I am looking for are not appearing. Good to know though :-)
If these are regular firewall policies, then it's a behavior I have never seen. From what I know, interface policies or local-in policies are managed through CLI, but I don't think regular firewall policies not appearing in the GUI is an expected behavior. Maybe you have applied a filter that's preventing you to view all the policies.
There is no CLI command that hide regular firewall policies. A few explanations that other ppl already said:
Virtual wire-pair policies are in the same CLI table (config firewall policy) but shown in a different page/view
You may have invisible filter set - try to reset the table via the setting icon in top left corner (hover over the table menu)
Could be a bug
See if you add a new policy with exact settings but different name and of it shows up
Try changing the name or source interface and see if it helps
I think you’re referring to “local-in” policies. Google that.
I have seen instances where if you built the config file and then restored it, but referenced an object that wasn't defined, it would not show the policy in the GUI. Make sure all referenced objects and security profiles exist exactly as named in the policy.
