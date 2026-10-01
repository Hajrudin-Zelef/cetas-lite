---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209-1
title: "local access only subdomains"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209.md
source_anchor: ""
source_lines: [1, 109]
sha256: 67ee78d7282f90e4851f353009191f42429c93a9916eea928c05f853e7b5c2ce
---

# local access only subdomains

Setting up HAProxy and Let’s Encrypt on OPNsense
If you’re reading this, wondering why my blog came up before the official documentation - they’ve removed the original documentation on account of thier enterprise-only Reverse Proxy and Webserver.
You can find a copy of the original documentation at archive.org
HAProxy uses ACME Let’s Encrypt for SSL authentication
OPNSense’s HAProxy package can use ACME for certificates.
First, we must install those two packages.
Installing required software
- In your OPNsense, go to: System --> Firmware --> Updates and install all updates.
- Then, head to: System --> Firmware --> Plugins and install the following plugins:os-acme-client ,os-haproxy
Reverse Proxy OPNsense Preperation
There are several changes we have to make to the defaults of OPNsense before we can intake traffic to our router.
System preparation
In OPNsense go to: System --> Settings --> Administration
You will need to checkbox the Disable web GUI redirect rule and change the Web GUI TCP port to a number you can remember, example: 4443.
This change is to allow your router to reply to requests on the default ports for HAProxy’s traffic (80/443). By moving the port number OPNsense uses, you’re able to free up those port numbers for HAProxy.
 Again, be sure you have change the  Web GUI TCP port to ** a number you can remember. **
Then, reconnect to OPNsense’s web interface using the port you entered, example: https://192.168.1.1:4443
Setting a Virtual IP
I like to use a virtual IP instead of just pointing all traffic back to the localhost.
It helps me see traffic a little bit better. If you want to create a virtual IP open: Interfaces --> Virtual IPs --> Settings
Creating a Virtual IP
When thinking about this new IP address for your “SSL Offloading Server” you would want to chose an IP that belongs to another network than any you’re using.
The localhost subnet should, generally, be available to us to help avoid IP conflicts in your local network.
The chosen IP/Subnet will be the IP on which the HTTP_frontend and HTTPS_frontend will be listening on.
We can use anything from 127.0.0.0-127.255.255.255. For this, we will choose to use: 127.1.2.3/32
- Mode should beIP Alias
- Interface can remainLAN
- Network Address is the address we picked above127.1.2.3/32
- Save.
- Apply.
Port Alias for HAProxy’s Ports
To make things easier on us, we should create an alias for all the ports we need HAProxy to use.
Go to: Firewall --> Aliases
You will likely see some other aliases here.
Now, we are going to create an alias for the ports that HAProxy will be listening on. Click the add button to begin.
- Name anything you like, example:HAProxy_Ports .
- Type should bePort(s)
- Content needs to be80 and443 . You can type the number and hit enter to save the field.
- Now, both ports are set and we have a recognizable name for the alias, click save.
Let the traffic start following
Modify the firewall: Firewall --> Rules --> WAN
Make a firewall rule to allow any inbound traffic on the WAN interface connecting to the HAProxy_Ports alias.
- Add a new firewall rule.
- Set the action toPass
- Interface isWAN
- Direction isIN
- TCP/IP version4
- Protocol needsTCP
- Source can beAny
- Destination needs to beThis Firewall
- Destination port range should be set to ourHAProxy_Ports alias.
- Save.
Appeasing the OCSP Must Staple
Head to settings: System --> Settings --> Cron to create a new cron job.
Because our certificate has the OCSP Must Staple extension we need to update HAProxy’s OCSP data regularly.
If this isnt done, clients connecting to HAProxy will get a security warning and won’t be able to connect.
The fix is to create a cron job
I have had issues with OCSP Staples so I set my cron job to run everyday, every hour. This is more than likly not nessasary and may be a part of this tutorial that I update… but until then…
- Add a new cron job, set any number forMinutes and the other fields should be an* .
- For command chooseUpdate HAProxy OCSP data .
- Save.
It doesn’t matter if this job runs before or after a certificate renewal as the OCSP data gets updated anyway after installing new certificates in HAProxy. This is because the ACME plugin restarts HAProxy after installing the new certificates.
Let’s Encrypt (ACME Client)
Enable and Update Schedule
Moving on, go to: Services --> ACME Client --> Settings
- Uncheck Show introduction pages
- Check Enable Plugin
- We don’t need the HAProxy integration . Leave this one alone. It is for ‘HTTP-01’ and this tutorial is using the ‘DNS-01’ challenge.
Update Schedule
Click on the Update Schedule tab next to the Settings tab in the ACME Client services section.
This is located in: Services --> ACME Client --> Settings --> Update Schedule
- The Update Schedule tab will allows us to configure at which time of the day our certificates are renewed.
Why change time of day?
If everyone renews their certs at the same time, let’s say, on the :00 of the hour – there will be a heavy load of renewals on the CA.
You want your renewal request to happen at a time of the day where there is not much load on your services as well. The ACME plugin restarts HAProxy so it can use your new certificates which results in a very short downtime of HAProxy.
- Pick from the numbers below for the minutes of the Renew ACME certificate command, the hour is up to your use case.
2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59
Restart services when there are changes
Once you complete the certificate renewal, you have to make sure the service restarts so it uses the new file.
To automate our service restart, visit: Services --> ACME Client --> Automations
Inorder to restart HAProxy after the cert updates, create a new automation.
- Name can be anything to identify the automation, example:RestartHAproxy
- Run command should select the commandRestart HAProxy (OPNsense plugin)
Account and CA
Next, go to: Services --> ACME Client --> Accounts
- On this page, enter a Name – for reference, I use the domain name (ssl.test.house.lan).
- E-Mail Address is the most important part. This is what is tied to the certificate for renewals. Too many renewals you’ll hit your limit and have to wait a week to try again.
Note: You can use the staging environment, Let's Encrypt Test CA, instead of the default ACME CA Let's Encrypt
ACME DNS-01 challenge
To request a certificate, we need to issue a challenge.
Head to: Services --> ACME Client --> Challenge Types
To get a wildcard certificate we need to use a DNS challenge. This tells Let’s Encrypt we own the entire domain and can therefore issue certificates to the subdomains beneath it.
- For Name I would put the assocaited domain and the challenge type (acmedns.test.house.lan).
- Challenge Type should be set toDNS-01
- DNS Service is up to your provider.
  - I use CloudFlare, as anyone can use them by changing Nameservers. Create a new API Token (with correct credentials).
- Save and double check you got the Account ID, API Token, and all credential’s permissions set correctly.
Issuing a new certificate
Issue a certificate: Services --> ACME Client --> Certificates
- Common Name is the URL of the certificate you’re requesting. We want a wildcard for a subdomain (*.test.house.lan).
- Select the ACME Account we created earlier (example: ssl.test.house.lan).
- Choose the Challenge Type that you named above (example - acmedns.test.house.lan).
- Key Length is a preference. I useec-384 . Generally, the higher number the better.
- Be sure you check the OSCP Must Staple checkbox. This is a modern requirement.
- Automations should have our restart policy we made earlier, select that now.
- Save
Depending on what Certificate Authority you chose earlier (Test CA or Not) will issue your certificate.
To issue your certificate, you need to hit the circular arrow button, it’s alongside all the other buttons under commands
Log file view of the action
