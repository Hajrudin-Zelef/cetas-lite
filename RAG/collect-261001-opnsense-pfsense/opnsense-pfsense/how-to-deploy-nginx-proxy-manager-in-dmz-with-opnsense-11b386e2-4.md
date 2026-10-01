---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2-4
title: "how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2.md
source_anchor: ""
source_lines: [152, 182]
sha256: d8038b6044d469a8f47631e652f7b94838639547e5a35aa0df65925881ce949d
---

# how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2

The “Scheme”, “Forward Hostname / IP”, and “Forward Port” are used to enter the information for the hosted service. You can use either the hostname or IP address. The service you are hosting can be either HTTP or HTTPS. It is important to note that it is possible to have HTTPS enabled for the entry point on the reverse proxy but HTTP on the backend service. This allows you to host a service using HTTPS that may not necessarily support HTTPS out of the box. You may decide if you wish to enable HTTPS on the backend or not so the data arrives at the service encrypted. Since the SSL certificates are managed on the reverse proxy, the backend service most likely would be using a self-signed certificate. Some may prefer to have HTTPS encrypted all the way to the service, but it may be debatable how useful a self-signed certificate is on the backend if the certificates are not being authenticated. You can enable the “Block Common Exploits” option because why not? Not sure why you need to leave that disabled unless perhaps it can cause issues with certain services you are running behind the proxy.
You may request a new certificate at the same time you are creating a proxy host which is convenient. Enable “Force SSL” to always use HTTPS for the service and enable “HTTP/2 Support” to improve performance (you may not notice much difference on a home network with low amounts of traffic). If you want extra security enforcement by the web browser, you can enable the “HSTS Enabled” option. Click the “Save” button to add the proxy host.
You should be careful with the HSTS option. I somehow managed to get my root domain using HSTS and all of my subdomains displayed an error blocking access to my services in the browser. I had to manually clear out the HSTS settings out of my browser for my domain name in order to restore access.
Once the proxy host is added, you will see it listed on the “Proxy Hosts” page.
Optional: Setting up a Wildcard Certificate
If you have several subdomains you are using for your apps/services, you may want to set up a wildcard certificate. This will allow you to use the same certificate for all of your hosted services. To set up a wildcard certificate, you will need to use the DNS method of creating a certificate. Depending on the DNS service provider you are using, the steps may vary slightly. In my example, I am going to use Cloudflare since that is the DNS provider I am using. You will need to choose one of the supported providers in the dropdown box and follow the instructions. In general, many of them use an API and require you to enter the API key/token to allow access to your DNS provider.
Enter the domain with the wildcard in the “Domain Names” box. If you want to host something at the room domain, you need to include it in the list. The example shows *.homenetworkguy.com but you can also include the root domain such as homenetworkguy.com as well.
Bonus: Set up Cloudflare
Many users prefer to host their services behind Cloudflare to take advantage of the DDoS protection, caching, and not exposing their WAN IP addresses.
Set up DNS Hostname and Enable Proxy
To set up a Cloudflare proxy, you do not have to have your domain name registered with Cloudflare, but you will need to use the Cloudflare name servers instead of your registrar’s name servers. This allows Cloudflare to intercept the traffic for your domain should you choose to set up a proxy. Once your domain name is set up on Cloudflare (a free account will work just fine for home users), you simply add your hostname and enter your WAN IP address. “A” records are for IPv4 addresses while “AAAA” records are for IPv6 addresses. If your WAN IP is dynamic, you can set up a dynamic DNS client in OPNsense (via a plugin) or on a system you have running on your network.
The difference with using a Cloudflare proxy is shown in the diagram below. As you can see, there is not a radical difference. Instead of directly accessing your external WAN address, external clients will access your services via the Cloudflare proxy. Since we set up split DNS, the internal client will be directed to the Nginx Proxy Manager in the DMZ network rather than be routed through Cloudflare so you can will have access to them even when your Internet is down. You will also gain a slight increase in responsiveness without being routed through Cloudflare.
Set the Cloudflare Encryption Mode
The SSL/TLS Encryption mode in Cloudflare needs to be set to “Full” since Let’s Encrypt certificates are being used by the reverse proxy. To set the mode in Cloudflare, go to the “SSL/TLS” page and clicking the “Full” radio button:
You could use “Full (Strict)” mode and create an origin certificate on Cloudflare which you can load into your reverse proxy to increase the security between Cloudflare and your reverse proxy. However, web browsers may not like accessing the reverse proxy directly with the origin certificate installed since they are designed to be used for authentication by Cloudflare. I had some issues with browsers blocking access (but perhaps I had some HSTS issues going on too). You could try taking that approach, but if you encounter issues with the certificates, using the “Full” mode is a good compromise for allowing both internal and external access to your reverse proxy.
Disable Caching
Depending on the services you are hosting, you will need to disable caching because free accounts are only allowed to cache so much data before being banned. Streaming services such as Plex Media Server will definitely need caching to be disabled. I had issues with transferring files in Nextcloud until I disabled the cache (you also need to be careful about enabling certain optimizations since I had trouble with the Rocket Loader feature of Cloudflare in the past).
To disable the cache, go to your domain in Cloudflare and click on the “Rules” page. Then add a rule with your URL you want to disable the cache. If you want to disable it for everything on your domain, use wildcard characters such as *homenetworkguy.com/*. I have seen some examples that suggest using *.homenetworkguy.com/* but that only matches your subdomains and not the root domain. You need to remove the first period to match everything.
Select the “Cache Level” setting with the value of “Bypass” to disable the cache. Click “Save and Deploy” to enable the cache rule.
Lock down IPs to Cloudflare
Cloudflare publishes a list of IP address ranges which are used to proxy connections, which makes it very easy to create a firewall rule which only allows Cloudflare IPs to connect to your reverse proxy. I wish more services would publish their IPs so that it is easier to tighten down access. Of course, using certificates and other forms of authentication further protects access to your services.
On the “Firewall > Aliases” page, create a new alias by clicking the “+” button. Enter “Cloudflare” or whatever name you prefer for the alias “Name”. For the “Type”, select “URL Table (IPs)” so that you can enter the URLs published by Cloudflare for their IPs. The “Refresh Frequency” can be set to daily by entering “1”, and the hour of the day can be set in the “Hours” box. You probably do not need to refresh it daily, but it does not hurt to do so. It is not a huge list so it will not take long to download the lists. For the content, enter the URLs https://www.cloudflare.com/ips-v4 and https://www.cloudflare.com/ips-v6 for both IPv4 adn IPv6 addresses. Enter a “Description” if you like.
Once the alias has been created, you can use it in your NAT port forward rule as the “Source” addresses:
| Option | Value | 
|---|---|
| Interface | WAN | 
| TCP/IP Version | IPv4+IPv6 (IPv6 is optional) | 
| Protocol | TCP | 
| Source | Cloudflare (the alias you just created) | 
| Source Port | any | 
| Destination | WAN address | 
