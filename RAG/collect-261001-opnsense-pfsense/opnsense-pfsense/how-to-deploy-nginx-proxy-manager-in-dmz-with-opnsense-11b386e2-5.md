---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2-5
title: "how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2.md
source_anchor: ""
source_lines: [183, 191]
sha256: 66259da30237d1b8968d6007c11f1e453efb074f46b914599bb0a27ba59d4c21
---

# how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2

| Destination Port | WebServerPorts (an alias for port 80 and 443) | 
| Redirect target IP | 192.168.2.50 (or use an alias which may include the IPv6 address) | 
| Redirect target port | WebServerPorts (an alias for port 80 and 443) | 
| Description | Allow external access to Nginx Proxy Manager | 
| Filter rule association | Add associated filter rule (or Pass) | 
Cloudflare Argo Tunnels
Cloudflare has made Argo Tunnels available on the free tier so anyone can use them. If your ISP blocks port 80, you may want to use an Argo Tunnel so that you can have HTTP to HTTPS redirection. You may also want to use it if you feel it is easier than creating the firewall rules as described in this guide. Cloudflare claims you do not need to set up any firewall rules using Argo Tunnels since you are establishing a direct outbound connection (a tunnel) to Cloudflare. It sounds to me a bit like SSH tunneling. I have not had the chance to experiment with Argo Tunnels, but I wanted to mention that as an option for those who have more restrictive ISPs or for those who like the idea of not having to deal with as many firewall rules.
Conclusion
Reverse proxies can be a very useful tool for hosting multiple services in your network and offer performance and security benefits. The versatility of reverse proxies allows you to use them for both internal or external access to your services. I hope this guide helps increase your understanding of reverse proxies and the benefits of using one. Also, I hope providing the additional context of an example network further aided in that understanding. When I write I strive to include some context because I feel that it can be crucial to gaining a deeper understanding of the topic at hand.
