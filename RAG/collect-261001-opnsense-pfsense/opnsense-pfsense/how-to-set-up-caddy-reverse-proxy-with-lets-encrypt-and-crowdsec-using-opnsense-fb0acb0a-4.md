---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense-fb0acb0a-4
title: "caddy.service"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense--fb0acb0a.md
source_anchor: ""
source_lines: [258, 304]
sha256: eb411e0aa0a36d5db9e9aa58af003c782b0deb8f3b0d4992657aa254ae7add74
---

# caddy.service

 localhost                                          192.168.1.1      2024-03-11T14:11:50Z   ✔️        v1.6.0-freebsd-4b8e6cd7   password    21s            
 caddyDmz   192.168.10.10    2024-03-11T13:55:29Z   🚫                                  password    ⚠️ 16m42s       
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
To validate the machine, you can enter the following command.
sudo cscli machines validate caddyDmz
Add Collection(s) to CrowdSec Agent on the Caddy Server
To add extra Caddy-specific parsers, you can add the following collection to your CrowdSec installation on your Caddy server. The Caddy collection includes a Caddy log parser and basic HTTP protections.
While the Caddy log parser may only be beneficial if you are using Caddy as a web server instead of a reverse proxy, the included basic HTTP protections should be helpful to protect web apps that are behind the reverse proxy.
sudo cscli collections install crowdsecurity/caddy
If you are hosting a public service (which great care must be taken to address security), it may not be a bad idea to also include the HTTP DoS collection to help detect denial of service attacks.
Of course, you should test this does not interfere with the normal operation of your app/service. I also do not know if this would be beneficial if you are using a Cloudflare proxy which includes DDoS protection.
sudo cscli collections install crowdsecurity/http-dos
You may now finally restart the CrowdSec agent on the Caddy server.
sudo systemctl reload crowdsec
After reloading CrowdSec, you should check if it is running properly.
sudo systemctl status crowdsec
Test the Configuration
Assuming all was configured properly, you should be able to access the service that is behind the reverse proxy by visiting https://homepage.homenetworkguy.com.
Note
Keep in mind that certificate issuance by the Caddy reverse proxy may take a short while. I noticed such a delay in issuance when testing this guide on a virtual machine environment behind my main network. It seemed to take longer than when I deployed the reverse proxy on my primary network.
I am not sure why there seemed to be a delay in certificate issuance when trying to solve the DNS challenge in my virtualized lab environment, but I wanted to make you aware in case it happens to you as well.
For CrowdSec, you can test if it is blocking connections by manually initiating a ban. The CrowdSec bouncer on the Caddy installation should be aware of banned IPs from polling the LAPI periodically or on each HTTP invocation depending on how you configured the bouncer for Caddy.
Change the IP address in the following command to a system you wish to test a ban. The nice thing about this ban is that it is temporary (5 minutes in the example below) so you will not need to worry about getting locked out permanently. You may enter this command on the Caddy server:
sudo cscli decisions add --ip 192.168.10.50 --duration 5m --type ban
If you look at the OPNsense CrowdSec interface, you will notice the ban was sent to the LAPI on OPNsense. Any machines that are using the LAPI on OPNsense will apply the same ban. This is great because if something is bruteforcing the OPNsense web UI or SSH and gets banned, it will also be immediately banned on the Caddy reverse proxy and vice-versa!
Although this process was a bit lengthy, you only need to do this entire process once. You will be able to rebuild new versions of Caddy using xcaddy (stop the Caddy service first), and you may add more services behind the reverse proxy by updating the Caddyfile (and restarting the Caddy service).
I hope this guide saves you some time figuring out how to put all of the pieces together if you are trying to accomplish something similar on your home network!
Upgrading Caddy
When you want to upgrade to a new Caddy version, simply run the same xcaddy command as described earlier in this guide to build a new version.
xcaddy build \
    --with github.com/caddy-dns/cloudflare \
    --with github.com/hslatman/caddy-crowdsec-bouncer/http
Tip
I recommend making a backup before upgrading in case you encounter issues upgrading.
Stop the Caddy service using the following command:
sudo systemctl stop caddy
Delete the existing Caddy installation.
rm -rf /usr/bin/caddy
Or you could rename the folder so you can revert back if you have trouble.
mv /usr/bin/caddy /usr/bin/caddy-old
Then move the newly build Caddy application to the /usr/bin folder.
mv caddy /usr/bin/
Start the Caddy service again.
sudo systemctl start caddy
You may check the status to verify you do not see any error messages.
sudo systemctl status caddy
The final test of a successful upgrade will be to check if you can still access the sites you have set up behind the reverse proxy.
