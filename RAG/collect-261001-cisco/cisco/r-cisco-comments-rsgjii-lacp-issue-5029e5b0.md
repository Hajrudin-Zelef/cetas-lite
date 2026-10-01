---
id: collect-261001-cisco/cisco/r-cisco-comments-rsgjii-lacp-issue-5029e5b0
title: "r-cisco-comments-rsgjii-lacp-issue-5029e5b0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "parameters"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-rsgjii-lacp-issue-5029e5b0.md
source_anchor: ""
source_lines: [1, 130]
sha256: 8298b74738d52c02dd64227a869b2306ec74eb7c517fa013cb08a86498f3170e
---

# r-cisco-comments-rsgjii-lacp-issue-5029e5b0

LACP issue 
        
    I have server 2016 that has teamed network interfaces 1 gb each connected to cisco nexus what are commands that i can run on switch side to see why LACP stoped working I see that server shows lacp protocol error.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
This will give you a lot of information.
Has the port channel stopped or is LACP reporting an error? LACP is not the same as a port channel.
I see lacp error on server side but not sure nexus side. Not sure if it is possible to debug on os side there is setting for fast or slow handshake
sh logsandsh etherchannel [PO NUMBER] summaryare the two commands I would use first.
It is very important, during troubleshooting, to assign "dead" or "inactive" (locally significant) VLANs to etherchannel ports so no wayward traffic can cause issues. Only when layer 1/layer 2 traffic is working do I swap the dead VLANs with normal ones.
Nexus doesn't use the etherchannel command.
Me bad.
I did not know this was a Nexus platform.
port-channels are Etherchannels
show port-channel summary to start. check physical cables, check if switches are communicating if you are connected to two different ones. check if switch config changed on ports.
this may help.. https://www.cisco.com/c/en/us/support/docs/switches/nexus-9000-series-switches/118851-technote-lacp-00.html#anc4
silly question, does MTU match on both ends?
Also, https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus9000/sw/6-x/interfaces/configuration/guide/b_Cisco_Nexus_9000_Series_NX-OS_Interfaces_Configuration_Guide/b_Cisco_Nexus_9000_Series_NX-OS_Interfaces_Configuration_Guide_chapter_0110.html#concept_98946DAEC5AA41B085E5D04709CDCB4B
I would try all the the things previously stated, but after I would look at the documentation on server 2016 and your hardware to be sure there arent parameters that need to be set for LACP ie active active or active passive etc. Also be sure you know which side needs to be brought online first it shouldnt matter but sometimes it can
Are you sure you're not using FEX with the Nexus? That changes the design as to what it allowed from LAG ports.
