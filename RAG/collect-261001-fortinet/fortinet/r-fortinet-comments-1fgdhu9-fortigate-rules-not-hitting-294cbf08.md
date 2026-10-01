---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1fgdhu9-fortigate-rules-not-hitting-294cbf08
title: "r-fortinet-comments-1fgdhu9-fortigate-rules-not-hitting-294cbf08"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1fgdhu9-fortigate-rules-not-hitting-294cbf08.md
source_anchor: ""
source_lines: [1, 129]
sha256: 772e2cc3f0ba8b4efd366dcc12c45f67dcdc18c8642f45393adbeee10adabe05
---

# r-fortinet-comments-1fgdhu9-fortigate-rules-not-hitting-294cbf08

Fortigate rules not hitting  
        
    Hi guys. So I’m new to firewall management and had a question. I’ve put some deny rules the firewall and have added some source ips and some destination ips. The thing is, if the rules are not being hit even after the policy has been pushed. The destination ips are NATed, so I need to know, do I put the IPs the real IPs are mapped to ( from the NAT pool)? Or the real IPs in the rule?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
First, look up the diagnose debug flow commands. They can show you all of the steps the firewall is taking to make a decision for each session and it's one of the go-to troubleshooting tools.
Second, since you mentioned IP Pools, so SNAT - the firewall allow/block decision takes effect before NAT happens. So your addresses in the policy should be the real IP's. If it was DNAT, your destination address would be a VIP instead. On older firmware versions, you'd need to make sure "set match-vip enable" is configured on a blocking policy containing a VIP - this is set by default on newer firmware versions.
The destination should be the actual IP of the destination. Only specify a NAT ip for the destination if the next hop after the Fortigate, but before the destination is handling the NAT. Also if your deny rule is intended to block traffic to a VIP, then you have to edit the policy in cli after saving it. You'll need to add 'set vip-match enable'.
With fortigate deny policy don't apply to vip with this ip. You need to create a specific deny policy for the vip or to use the parameter match-vip in cli.
https://community.fortinet.com/t5/FortiGate/Technical-Tip-DENY-Policy-for-Virtual-IP-Firewall-Policy/ta-p/192456
i think that might be the issue. I'll look into that thank you
Had this recently with a 7.0.15 Fortigate (NVA) The link above resolved the issue and I began getting hits on the deny rule on VIPs.
Set match vip enable or w/e via cli. Just be sure to add it to the correct policy ID
That depends on the FortiOS version. One of the 7.2.x releases changed it so the FG defaults “set match-vip” to enable for new policies.
I'd start with this... where's the policy in the rule list? New policies are added to the bottom of the list, right above the implicit 'deny all' rule. A firewall looks at policies from the top down until it finds a match. If you haven't dragged your deny policies to the top of the list, it's likely that the traffic is hitting a permissive rule first.
It's on the top
Just observe the logs by going to forward traffic under Log & Report section check for policy ID through which it is allowed.
I think that's what I'll have to do
Are you logging the traffic or is it set to UTM only.
