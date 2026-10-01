---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d-3
title: "setup-adguard-home-opnsense-adblocker-c464187d"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d.md
source_anchor: ""
source_lines: [162, 233]
sha256: de84805456af459fb63b4e20b590324d4c4bb1137307f66c408783377d9521c7
---

# setup-adguard-home-opnsense-adblocker-c464187d

In my case there was an issue happening from time to time with Cloudflare cert exchange, so there was need to create rules for that on WAN port
Action: PASS
Direction: IN
Protocol: TCP/UDP
Source: 1.1.1.1/32
Port: 853
Destination: WAN address
Step8: Make the AdGuard OPNsense a systemwide DNS resolver
Under OPNsense you can navigate to System: Settings: General and add backup DNS servers under Networking -> DNS servers, but if you leave them empty only AdGuard Home will be used.
I usually leave them empty.
Also make sure that:
- Disable Allow DNS server list to be overridden by DHCP/PPP on WAN
Step10: Configure OPNsense DHCP to use AdGuard
By default, if a specific DNS server is not configured for your client DHCP settings, then OPNsense will provide the clients with the same DNS server it uses. This could have been a DNS server that was configured when you set up OPNsense, or it also can use DNS servers that are provided by your internet service provider.
So to update our LAN DHCP configuration, we’ll head back to our OPNsense web interface. From there, we’ll navigate to Services > DHCPv4 > [LAN].
Make sure that the DNS server field is empty.
After these changes all your DNS requests from PC should go through our configured internal DNSes.
Step11: Client Testing
Now we should be all set up! However, it’s important to note that because of the way DHCP works, clients may not pick up the new configuration immediately. When DHCP assigns an IP address, it also tells the client how long it can use that address for. So if a client stays powered-on & connected, it won’t ask for new configuration until that timer expires.
We can speed that up by resetting the network interface on our clients. This can be done in a number of ways including rebooting the client or simply disconnecting from wifi/ethernet & reconnecting.
In Windows please type in Search > View Network Connections. It will show you a list of active adapters for Internet. Right Click on the connected adapter and click on Properties. Choose Internet Protocol Version 4 > Properties and make sure that the option Obtain DNS server address automatically is enabled.
If you’ve changed anything, please reset your Internet connection on PC.
I’m using a Windows computer as my test system, so first I’ll check via the nslookup command from command prompt – which will query our configured DNS server & return the resolved IP addresses.
As we can see, we did get the correct IP address.
Now we can test one of the sites that are blocked by default via Adguard and see the result.
- adblock-one-protection.com
- uvsvlisbartwq.com
- wooden-comfort.com
Now that’s the result we want! By returning the 0.0.0.0 result, our client can no longer resolve that domain. So if this was an advertisement or tracking domain, it’s now blocked from loading.
You can check the query from the Adguard Administration Panel > Query
You can search the request by URL, check the filter result (Blocked) and the IP/hostname of device which was trying to reach the domain.
The filtering service is running excellent!
Step13: DNS leak test
When you’re doing plain text requests to your Public DNS, your hostname/IP and URL can be easily intercepted by all intermediary devices, ISPs, marketing companies that are on the way to reach the Public DNS. They gather such information for your Public DNS to sell you ads, profile you and spy on your traffic.
Therefore we have set up DNS over TLS encryption in the Unbound in previous steps.
We can test if our queries are going in encrypted format by visiting sites:
If none of the results has your Public IP or hostname that means that the encryption works fine.
Step12: Configure blocklists in AdGuard Home
Blocking Domains
First thing we’ll look at is our DNS blocklists. We’ll navigate to Filters > DNS blocklists.
Here is where we can ask AdGuard to query lists of what domains to block. By default, AdGuard does include two – but we can add more if we want:
Photo
If we want to add to the configured blocklists, we can do so by clicking the Add Blocklist button. This will prompt us whether we want to choose from a pre-populated list, or supply our own custom list:
Photo
The easy option will be selecting from the provided lists:
Photo
There are a ton of different curated block lists available depending on what you’re trying to block. If we wanted to use a custom list, a lot can be found on GitHub just by searching for PiHole or Adguard blocklists.
How to pick a blocklist will be up to you. There are blocklists that focus on advertisements, tracking & analytics, parental controls, etc. So it just depends on what areas you want to focus on.
Allowing Domains & Custom Filtering
If we have a list of known services that we want to ensure are never blocked, we can pull those lists via Filters > DNS allowlists. However, it’s more likely you’ll find a handful of domains you want to unblock, rather than a whole list.
For that – we can go to Filters > Custom filtering rules. At the bottom of this page there is a tool to check filtering, where we can enter a domain name & instantly see what the result is.
For example, with the default ruleset I’ll check to see if 0x2142.com is filtered:
Photo
So by default that domain isn’t found anywhere, so it will be permitted. The tool also gives us a convenient button to quickly block a domain.
We can click that button, or add the syntax ||0x2142.com^ to the custom filtering rules at the top of the page (and saving via the Apply button). Now if we check the results again – the filter check will show the domain is blocked:
Photo
And of course, we don’t want to block 0x2142.com!! So let’s add this to our allowlist instead, so that it can never be blocked . We can do that by adding @@||0x2142.com^$important to the custom filtering.
And now we’ll see a green box that shows that the domain is permitted via an allowlist:
Blocking Known Services
The other option worth mentioning is the ability to block certain known services, like WhatsApp, Twitter, Reddit, etc. This can be great if there are certain services you want to block, or for use as parental controls.
This can be found on the Filters > Blocked Services page.
This way we can select a service to block, rather than having to know all of the individual domains that service uses. For example, I’ll go ahead and select YouTube to block – and we’ll check that later on after we configure our clients.
Troubleshooting Blocked Domains
Okay, so now we know our blocking works…. But now someone in our home is trying to access YouTube & it’s not working. How can we tell if that’s our AdGuard service?
Our first stop might be the AdGuard query log. Opening this log, we can filter by domain name or client – or show only blocked queries if we like.
Pretty quickly we can see the issue – we blocked YouTube’s services:
Now we know how to fix the issue, which would be to unblock that service. However, if it was just a specific domain that was blocked, we would likely want to add it to our custom filtering as we showed earlier.
Reporting
Last but not least, we can also check our AdGuard Home dashboard again, which should be much more interesting than before:
Here we can quickly see how many queries have been made & how many were blocked for various reasons. We’ll also see what clients are using our DNS server, and which are making the most queries.
Most interesting (at least to me), is being able to see the top domains that were queried or blocked. Here’s where you might find some interesting information. For example, on my test machine – it’s a fresh installation of Ubuntu & we used FireFox to test. But we can see that even during the brief time it’s been set up, almost all of the highest queried domains belong to Mozilla’s analytics services. So it may be tempting to add those to our custom blocklists.
Additional Info
