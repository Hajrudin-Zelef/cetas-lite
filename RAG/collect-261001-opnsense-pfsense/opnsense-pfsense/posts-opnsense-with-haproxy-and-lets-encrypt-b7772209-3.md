---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209-3
title: "local access only subdomains"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209.md
source_anchor: ""
source_lines: [231, 338]
sha256: 50c8ab7fb4480e7dacf6b752b61ee97a082fefcab94e6c54e8aca444d37c6464
---

# local access only subdomains

The Condition is generally made for the condition type. Again, as mentioned above, Host starts with.
The Rule states if the condition is met to execute a function. The function being, Use specified Backend Pool.
HAProxy Public Services
Next, go to: Services --> HAProxy --> Settings --> Virtual Services --> Public Services
Here we will create the frontends that are listening on our interface IPs and the virtual IP we created earlier.
You will have to place the rules for all of your services that you don’t want to get SSL offloaded in here.
At first we will create our SNI_frontend which will decide whether the traffic is going to be SSL offloaded or not.
Our default backend in this frontend will be the SSL_backend, that redirects all traffic to the virtual SSL_server which is actually the HTTPS_frontend.
Server Name Indicator frontend service
- Name is0_SNI_frontend
- Description should beListening on 0.0.0.0:443, 0.0.0.0:80 because you dont get to see the Listen Addresses for Public Services. It’s a convience thing.
- Listen Addresses need to be typed in, begin with0.0.0.0:443 and then enter0.0.0.0:80
- Type needs to beTCP
- Default Backend Pool is for theSSL_backend
- Save.
- Apply.
HTTP redirect to HTTPS frontend service
Now we will create our HTTP_frontend.
Make sure to place the HTTPtoHTTPS_rule in this frontend!
This frontend is necessary in order to redirect HTTP traffic to HTTPS. But you could also use it to serve non SSL encrypted services on port 80.
- Create a new Public Service and in the upper left hand corner of the create page, findadvanced mode and turn it on (green).
- Name is1_HTTP_frontend
- Listen Addresses is127.1.2.3:80 this combines our localhost address,127.1.2.3 plus the HTTP port:80 .
- Bind option pass-through must be:1 accept-proxy
- Checkbox Enable HTTP/2
- Also, checkbox X-Forwarded-For header
- Select Rules needs to haveHTTPtoHTTPS_rule set by clicking inside the square and selecting it.
- Save.
- Apply.
HTTPS frontend service
Now, the big show, create an “HTTPS_frontend”.
- This will be the primary frontend for traffic.
- It will be doing SSL offloading using the Let’s Encrypt certificate from the begining of this tutorial.
- You will place the “PUBLIC_SUBDOMAINS_rule” and any other rules for services that you want to get SSL offloaded in here.
- Create a new Public Service
- In the upper left hand corner of the edit page, find advanced mode and turn it on (green).
- Name is1_HTTPS_frontend
- Listen Addresses is127.1.2.3:443 this combines our localhost address,127.1.2.3 plus the HTTPS port:443 .
- Bind option pass-through must be:1 accept-proxy
- Checkbox Enable SSL offloading
- A new settings area, ‘SSL Offloading,’ will appear, and you will need to select your Let’s Encrypt certificate under the Certificates section.
- SSL option pass-through should be set to:1 curves secp384r1
- Enable Advanced settings needs to be checked.
  - This opens a NEW section for Cipers
Current Ciphers and Cipher Suites for a 100% A+ SSLLabs rating
Last updated/verified on 20230223 using Mozilla SSL Configuration Generator.
All ciphers with a strength of 128 bit or below have been removed in order to get a 100% A+ rating at SSL Labs.
1
2
3
4
5
Cipher List
ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES256-GCM-SHA384
Cipher Suites
TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256
- Cipher List should read from the code block aboveECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES256-GCM-SHA384
- Cipher Suites should also read from the code block aboveTLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256
- HSTS includeSubDomains needs checked.
- HSTS preload also needs checked.
- HSTS max-age can be increased to63072000
- Bind options need to haveprefer-client-ciphers removed, and addno-sslv3 ,no-tlsv10 ,no-tlsv11 ,no-tls-tickets
- Checkbox Enable HTTP/2
- Also, checkbox X-Forwarded-For header
- Select Rules needs to havePUBLIC_SUBDOMAINS_rule set by clicking inside the square and selecting it.
- Save.
- Apply.
Update Map file? Restart Service!
Anytime you want to add a new service and subdomain here are the steps:
- Create a Real Server that points to the IP address and PORT of the service that is running (example: I’m running Proxmox and want to verified SSL that service. I’d find the IP address of the machine, and the port number of the service you’re trying to connect to. In Proxmox’s case this is:  [IP - 192.168.1.12 with a PORT 8006] and, as Proxmox is encrypted by default, we must check theSSL checkbox to tell to HAProxy to expect an SSL connection. MAKE SURE to UNCHECKVerify SSL Certificate .)
  - That’s it! There are a lot of options, but all the Real Server needs is IP and PORT and if that port is SSL or not.
- That’s it! There are a lot of options, but all the 
- Creating a new Backend Server . Enter YOURSERVICENAMEHERE_backend for the name of this backend server. Then, select the name of the Real Server created above in theServers section.
  - This is where your pool would live if you had multiple servers doing the same thing.
- Modifing the Rules file. This is why it’s important to be consistant with names, we now must enter the subdomain to match with a backend. Naming our backends consistantly helps ensure when we need to add/edit the map file, it’s seemless.
  - The map file matches the entered subdomain with the backend service.
- Restart HAProxy.
  - That’s right nothing will work unless HAProxy loads the new map file.
Life without a map file: part 2
Setting a certificate for applications’ backend pool… without using a map file… if you so desire
If you didnt use a map file and are not going to point your traffic to a Virtual IP, then these instructions should get you up and running.
- Name set it to anyting you like, example:Public_Faceing_Pool .
- Listen Address should be your Public IP Address with port 443 on the end, example:123.45.67.89:443 .
- Type isHTTP/HTTPS (SSL offloading) [default] .
- Default Backend Pool can be set to any backend pool you made earlier – for testing of course, as ALL subdomains will connect to this one backend that then finds the server.
- Enable SSL offloading check this box. It allows us to specify a certificate.
- Certificates if you click in the box, you should see the certificate name you used above.
- Default certificate this is a drop down, choose.
- Select Rules use the rule that connects the condition to the backend pool.
- Save.
Test your new certificate
Access from external networks should now already be working.
Just try to access your URL “nextcloud.test.house.lan” from any device that is not connected to your local network, a good test device is your smartphone on cellular data.
You can now test your SSL settings at: https://www.ssllabs.com/ssltest/index.html
Access to your services, directly from internal networks
Choosing a DNS solution
If you try to access your URL “nextcloud.test.house.lan” from a device in your internal network, it should fail. Not just because that’s not a real TLD, but because your search domain is in your local network.
There are two ways of fixing this.
- With (Option A) being the better of the two. * Option A - Split DNS
- With (Option B) you lose the ability to track originating source IP in HAProxy when going through NAT.
- Option B - NAT Reflection
Option A - Split DNS (DNS Overrides)
Option A presumes you can create DNS entries on your local network’s DNS.
Unbound DNS will easily set up DNS overrides.
To set this up, go to: Services --> Unbound DNS --> Overrides
- Here you will need to create “Host Overrides” for each of your services.
- Yes, every subdomain that you want to seperate from the upstream public DNS.
The IP address can be any LAN (or VLAN) interface IP of your OPNsense.
