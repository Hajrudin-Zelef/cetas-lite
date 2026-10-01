---
id: collect-261001-cisco/cisco/r-cisco-comments-17wpgza-issues-with-ios-xe-1794a-4dacc76a
title: "r-cisco-comments-17wpgza-issues-with-ios-xe-1794a-4dacc76a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-17wpgza-issues-with-ios-xe-1794a-4dacc76a.md
source_anchor: ""
source_lines: [1, 155]
sha256: 72b50975dc6f82be1fcadbdecc957400bf903ae27e0593ee4adb19d2f3146274
---

# r-cisco-comments-17wpgza-issues-with-ios-xe-1794a-4dacc76a

Issues with IOS XE 17.9.4a 
        
        
        
    
    
    We have just upgraded to 17.9.4a last night, and then suddenly, some 9 hours later, nearly all updated switches started malfunctioning and had to be rebooted.
Has anyone else experienced anything bizarre with the 17.9.4a version?
P.S.: We are updated Catalyst 9200s and Catalyst 9300s.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
How did your trunks to the core look? Was the core touched at all? Were you able to SSH into these switches or did you have to console in to reload them?
We ended up needing physical access to the switches as this was a fairly urgent situation since it happened during office hours. SSH and ping were failing towards the affected devices.
How are they after reboot?
They work so far. Hopefully it doesn't mess up again.
What are the issues?
Connection dropped for all ports for some reason. Switch was still on, but no traffic going through.
At least show us the logs or something....
Been using it for 9300L with no issues
In our environment during the upgrade process we had several switches similarly go totally dead even worse than this. We had 6 9300 switches completely die no console output, no boot interrupt with keypress to get into rommon no nothing, we had to RMA the switches. This was going from 17.9.4 to 17.9.4a using DNAC to deploy the software.
We've got 92-9500s running it so far, no issues yet.
I did have a fabric mishap on my 9500 when I upgraded, but not sure if that's just DNA Center stuff, or part of the upgrade from 17.6.3 to .9.4a.
We just did no problem so far. Look for vtp config or allowed vlan config at trunk ports
We use Cisco Prime and there is a bug that can cause 17.9.x to blow up.
When Prime runs "show install summary" on a switch, the bug causes the databases that IOS-XE uses to mis-use some tables and create a memory leak. Switch will crash and reboot once there is no more memory and something pushes it over the limit, like to handle the authentication of a user.
They claim this bug is fixed in 17.12.x, but not yet in 17.9.x.
They say the bug will be fixed in 17.9.6.
I had to drop back to 17.6.x
https://bst.cloudapps.cisco.com/bugsearch/bug/CSCwf23122
We updated a cisco 4331 DMVPN hub router to this and a bunch of our remotes will no longer build to it.
Check isakmp policy encryption method… default changed from 16.12.5 to 17.6.3
Are you using DES?
I have 75 x 9300 on 17.9.4a with Dot1x.
I have not seen this before.
Have twenty 9300u and ux versions running sd-access/ise with no issues.
Most of my offices have the 9500’s in a stackwise virtual configuration. We just took the outage overnight since they don’t support ISSU. I learned that the hard way. If you use DNAC SWiM, DNAC will allow you to configure ISSU on the 9500’s but it does not work.
we suspect a case where switches randomly removes dacl pushed by ISE after some time. Anybody has the same issue? 9200 with code 17.9.4a.
"always code upgrade to the starred release even if you're stable on older code"
17.9.4a is one of the starred releases. I would add a couple of caveats to any firmware update:
Do it for a reason, ie. new feature you need, software bug or security issue
Let the "Starred Release" age a bit. I've seen Cisco add and remove "starred releases" within days of each other
+1
A lot of my customers' running more critical infrastructure. most of them have a lifecycle policy regularly upgrading to "starred" releases, because in case any issue, Cisco TAC would start with that step anyway, before even watching the logs.
We finished upgrades from 17.9.4 to 17.9.4.a without issues. We have a mix of 9500-16x cores and 9300’s.
Had one 9200 that reloaded to rommon, but it was a bug (had to unplug physically for 5-10min before it could boot normally). except that case everything went smooth for other 9200/9500 upgrades to 17.9.4a with DNA. Also did one 9500 that was compatible with ISSU and only 4 pings lost. I’m talking more than 30 switches.
