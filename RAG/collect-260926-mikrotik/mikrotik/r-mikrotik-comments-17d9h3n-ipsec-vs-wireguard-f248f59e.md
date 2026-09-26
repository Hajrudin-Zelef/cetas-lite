---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-17d9h3n-ipsec-vs-wireguard-f248f59e
title: "r-mikrotik-comments-17d9h3n-ipsec-vs-wireguard-f248f59e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-17d9h3n-ipsec-vs-wireguard-f248f59e.md
source_anchor: ""
source_lines: [1, 130]
sha256: 7a8ed650404c8eca6e1ab5bb36078079a87baa9dc197968d3cf75419b8ca7d69
---

# r-mikrotik-comments-17d9h3n-ipsec-vs-wireguard-f248f59e

IPSEc vs Wireguard 
        
    Hi guy :
Witch one do you prefer for site to site VPN Mikrotilk and why , witch one do you think is more secure and why.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
I like both personally. Wireguard is very simple to set up and is very fast despite not having hardware acceleration. IPSEC is hardware accelerated on most models but is more complicated to setup (more options and settings) which means mistakes that hurt security or efficiency can be made. Unless you are pushing lots of data over a VPN link and need every bit of speed you can get, wireguard should be more than enough.
wireguard because of simplicity. Never figured out how to configure ipsec
Since you don’t have ipsec vti support, it’s easier with wireguard in a situation where your clients have dynamic ip. If you want to be able to run a dynamic routing protocol.
functionally, there's no difference between vti and gre interface
How do you easily build a gre tunnel where one of the sides is on dynamic ip or behind cgnat?
If that helps you decide, WireGuard is a more advanced and easier to use VPN technology than IPsec. It is also safer by default, which might help you too.
WireGuard is the best VPN for most VPN uses, like giving employees access to private infrastructure that is protected by a firewall and letting workstations join so that people can work together.
The IPsec protocol works well in places where rules require old encryption methods or methods that WireGuard does not allow to be used. This could be because of legacy operating systems or Internet of Things devices.
If you are not in any of those groups, WireGuard is a better VPN option for setting up private network connections between companies and employees. This is because IPsec configurations are more complicated, and both the user and the administrator have to do more work to set up and maintain secure VPN connections.
WireGuard is great for streaming, gaming, peer-to-peer (P2P) networking, and giving employees access from afar, but IPSec is better for legal compliance and working with older technologies.
Overall, WireGuard looks like a good VPN protocol. It has easy-to-use interfaces, strong security features, and fast speed. It is a strong competitor in the VPN market because it has a small code base and is easy to check.
In the meantime, IPsec is still a reliable and tried-and-true choice that provides strong security and wide compatibility. Many people still use IPsec as their VPN protocol, even though there are new ones out there.
For more information: https://www.zenarmor.com/docs/network-security-tutorials/ipsec-vs-wireguard
Another vote for Wireguard, I have multiple setups with both PtP and roadwarrior examples and it is by far the easiest.
