---
id: collect-261001-huawei/huawei/questions-3067-load-balancing-difference-between-dns-and-ip-forwarding-vs-redire-bce3ab10
title: "questions-3067-load-balancing-difference-between-dns-and-ip-forwarding-vs-redire-bce3ab10"
domain: huawei
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["apache", "research"]
source: docs/RAG/collect-261001-huawei/questions-3067-load-balancing-difference-between-dns-and-ip-forwarding-vs-redire-bce3ab10.md
source_anchor: ""
source_lines: [1, 46]
sha256: 40194b2643a13d524cc9023b8f889d46a6bdd0779549a1d5e57526c2db5e36ef
---

# questions-3067-load-balancing-difference-between-dns-and-ip-forwarding-vs-redire-bce3ab10

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have come across a situation that I cannot understand. We have a Fortigate firewall that we have enabled to do load-balancing across two back-end Apache web servers. A DNS name is then mapped to the virtual IP on the Load Balancer.
As expected, when you browse to the DNS name/URL (e.g. www.something.com), the Load Balancer serves up a page from one of the back-end Apache web servers. The URL in the browser stays www.something.com. From what I understand, the Load Balancer in this case is simply forwarding packets between the browser and Apache while always staying in-path.
However, if I browse to the IP address that the DNS is mapped to, then the Load Balancer returns a HTTP 302 Found, with the Location header set to the DNS URL of one of the Apaches. The URL in the browser changes to the back-end server DNS.
Why is the Load Balancer redirecting when queried via IP, but correctly forwarding in-path when queried via DNS name.
I haven't used a Fortigate FW for load balancing, so I will answer some of the questions more generally.
First, as to your problem, the load balancer is doing exactly as it is supposed to do and I think your servers may not be configured correctly to respond to a request on their IP address. If you were to test this behind the load balancing, you could set the domain name in a local client's hosts file, behind the firewall with the server and access it both with domain name and internal IP. You will probably get the same result you are seeing now.
My guess is that you have virtual hosting turned on (to support multiple domains on a single server) and the "default" is not serving the same pages as your domain. You are getting a web page back from the server in both cases. If you need help with configuring your web server, you may want to try on ServerFault.
Second, to get into a bit more detail. A load balancer typically operates at L7 for at least HTTP and HTTPS clusters. This means that they don't just look at the IP address and forward it, nor do they "redirect" the page.
When they receive a request, they actually parse the request and forward it to a server after processing the request. There are many things that they can do at this point, such as rewriting the headers in both directions, potentially adding cookies (for persistence) into the data going back to the client, terminating SSL sessions, matching based on the URL, etc.
I recommend you spend some time fully reading the vendor documents to get a better understanding of how load balancing works (with Fortigate you can read both theirs and Coyote Point - another load balancing company Fortigate acquired). Understanding what it is doing will help you in cases like this and will allow you to unlock capabilities that you didn't realize existed.
After reading the HTTP host-based load balancing in the Fortigate Load Balancing doc, I can see how you can have an atypical load-balancing configuration that could result in what you describe. However, without part of your config, we cannot be certain if this is the case for you.
Fortigate FortiOS allows a Virtual Server to be created that is tied to Real Servers that each have a different host header configuration. Should any requests match the VIP of your Virtual Server, the load balanced requests will only go to Real Servers that match that host header. The most important part which nicely explains your symptoms, is that one of the Real Servers can omit the host header so it matches on any host header.
The Real Server without a host header might have been configured as a sort of "catch-all" that lands on a site that does the redirect.
Using the example below, only the 1st and 2nd rservers handle traffic matching your preferred DNS name via the host header, but the 3rd rserver takes anything matching all other host headers which includes the DNS VIP itself and sends to a site that could do a redirect.
config firewall vip
edit "http-host-ldb"
set type server-load-balance
set extip 192.0.2.1
set extintf "lan"
set server-type http
set ldb-method http-host
set extport 80
config realservers
edit 1
set http-host "www.example.com"
set ip 192.168.2.1
set port 80
next
edit 2
set http-host "www.example.com"
set ip 192.168.2.2
set port 80
next
edit 3
set ip 192.168.2.3
set port 80
next
end
end
I suppose the firewall load-balancing could be doing the redirect itself, but we cannot tell with the limited info provided.
