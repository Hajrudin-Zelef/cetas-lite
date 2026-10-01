---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d-2
title: "setup-adguard-home-opnsense-adblocker-c464187d"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d.md
source_anchor: ""
source_lines: [52, 161]
sha256: 8f3f5195d041a9ebf8c7d069a5d785ff485255123333962fd0a13130e87acb07
---

# setup-adguard-home-opnsense-adblocker-c464187d

So back in our browser, we can nagivate to: System > Firmware > Plugins. On this page we can search for adguard or scroll through the list to find it.
Then we just click the plus icon on the right side to install (not shown in the screenshotbelow).
This should install pretty quickly:
***GOT REQUEST TO INSTALL***
Currently running OPNsense 22.7.9 (amd64/OpenSSL) at Sun Dec  4 12:48:38 EST 2022
Updating OPNsense repository catalogue...
OPNsense repository is up to date.
Updating mimugmail repository catalogue...
mimugmail repository is up to date.
All repositories are up to date.
The following 1 package(s) will be affected (of 0 checked):
New packages to be INSTALLED:
	os-adguardhome-maxit: 1.8 [mimugmail]
Number of packages to be installed: 1
The process will require 35 MiB more space.
7 MiB to be downloaded.
[1/1] Fetching os-adguardhome-maxit-1.8.pkg: .......... done
Checking integrity... done (0 conflicting)
[1/1] Installing os-adguardhome-maxit-1.8...
[1/1] Extracting os-adguardhome-maxit-1.8: .......... done
Stopping configd...done
Starting configd.
Migrated OPNsense\Adguardhome\General from 0.0.0 to 0.0.1
Reloading plugin configuration
Configuring system logging...done.
Reloading template OPNsense/Adguardhome: OK
Checking integrity... done (0 conflicting)
Nothing to do.
***DONE***
Now all we have to do is enable the plugin. Please refresh the browser page with F5.
Navigate down to Services > Adguardhome > General. Our only option here will be an Enable checkbox, so we’ll select that & Save.
From Adguard ver 1.9 onwards, on OPNSense 23.1.6<…, Adguard may need to have one more field ticked for ensuring that is the main DNS
The rest of the setup & initial configuration will be done directly from the AdGuard-specific web interface.
Step5: Reconfigure Unbound
In our setup, we will actually use two internal DNS server services. The Adguard DNS requests will be forwarded to Unbound which would act as a validating, recursive, and caching DNS resolver and will encrypt our traffic with DNSSEC.
Since DNS as default is listening on port 53 we also want AdGuard Home to listen on this port to make or life easier. Out of the box OPNsense is already running Unbound on this port. We need to change this so they don’t conflict with each other.
- Navigate to Services > Unbound DNS > General
- Change Listen Port to 5353
- Enable DNSSEC Support
- Enable Register DHCP leases
- Enable Register DHCP static mappings (this will resolve hostnames for us in AdGuard Home)
- Register IPv6 link-local addresses
- Save the settings.
Now we are ready to configure AdGuard Home itself. I will not go in to all configuration here but some things are needed to make this work optimal with OPNsense.
Step6: Initial Setup for AdGuard
By default, the AdGuard Home web interface will run on port 3000 & is not HTTPS-enabled. So if your OPNsense firewall is at https://192.168.1.1, you’ll need to connect to http://192.168.1.1:3000.
As long as that works – we’ll see the initial setup prompt below:
We’ll click on Get Started.
Now we’ll be asked to configure the Admin Web interface (the interface we’re connected to now) and the DNS server interface (which clients will use to resolve domain names).
By default, AdGuard home will try to set both of these to listen on All interfaces – and set the web on port 80 & DNS on port 53.
I would recommend setting the Listen Interface on both of these to only your LAN-side networks. There is no reason to enable them on your WAN, and it can be a security risk to do so.
You may also get warnings that port 80 & 53 may already be in use. For the web interface, we could change 80 to 3000 & just keep what we’re using now.
However, if we change the default DNS port, that will cause some additional problems since client machines will query port 53.
So here’s what my set up looks like so far, with 192.168.60.1 being my LAN side interface:
- AdGuard administration board will use port 3000 and could be only reached from the LAN network
- AdGuard will be the default DNS for LAN and VLANS
On the next page, we’ll be prompted to set up an administrative user & password for logging into AdGuard.
Remember to put a strong password initially.
Next we’ll be given instructions on how to set up client devices. In my lab network, the OPNsense firewall is providing DNS server configuration via DHCP – so we’ll get to that configuration shortly.
For now, we’ll just click Next.
On the last screen, we’ll just get a message saying that setup is complete & a link to open the dashboard:
And now we can log in:
Browse to http://192.168.1.1:3000 or in my example http://192.168.60.1:3000
NOTE: If the AdGuard didn’t come up, please disable and then enable the service from OPNsense Services > Adguard Home. As last resort please reinstall the Adguard (first disable plugin, change the port 5353 in unbound to 53)
Step6: AdGuard additional configuration
After logging in, the first thing we’ll see is a pretty empty dashboard. We don’t have any clients configured to use this yet, so there isn’t anything to report on.
If there’s an update available, please download it first.
Configure AdGuard to use Unbound
When the wizard is complete we can login to AdGuard Home with the credentials you entered.
- In AdGuard Home navigate to Settings -> DNS settings and scroll down toUpstream DNS servers -> Private reverse DNS servers .
- Here we enter the Unbound server we changed earlier in OPNsense settings, 192.168.1.1:5353 , or with other port pointing to you OPNsense instance if you have another one.
If you use a domain name to resolve local hosts by name instead of IP you might need to tweak that in AdGuard Home as well. Let’s say you’ve entered a domain under System: Settings: General that is home.mydomain.xyz, and you want that to take precedence over the public DNS, if that also exists, when you are at home.
- In AdGuard Home navigate to Settings -> DNS settings and go to top section underUpstream DNS servers .
- Add [/home.mydomain.xyz/]192.168.1.1:5353 at the top of that list.
- Now you will resolve local machines when connected to your LAN, and if connecting over the internet the public DNS record will be used instead.
In my example my firewall is on lan address 192.168.60.1 with port 5353 for Unbound
- In AdGuard Home navigate to Settings -> DNS settings and go to top section underPrivate reverse DNS servers
- Here we enter the Unbound server we changed earlier in OPNsense settings, 192.168.1.1:5353 , or with other port pointing to you OPNsense instance if you have another one.
- Leave the Bootstrap DNS servers as default
- In Private reverse DNS servers type your Unbound server once more 192.168.1.1:5353
- Use private reverse DNS resolvers should be enabled
Now Click on Save and then Test upstreams. If the test is successful you will got a prompt.
Step7: Configure the Unbound upstream DNS
Let’s set the upstream Unbound DNS server to use encryption when sending a request to public DNS server.
In OPNsense please go to Services > Unbound DNS > DNS over TLS
We are using the Cloudflare and Quad9 Public DNS as an example. Click on Add and fill the gaps:
- Select enabled
- Domain: (leave this field empty)
- Server IP: 1.1.1.1
- Server port: 853
- Veify CN: cloudflare-dns.com
and Save.
Do the same for 3 more public
- Select enabled
- Domain: (leave this field empty)
- Server IP: 1.1.1.3
- Server port: 853
- Veify CN: cloudflare-dns.com
- Select enabled
- Domain: (leave this field empty)
- Server IP: 149.112.112.112
- Server port: 853
- Veify CN: dns.quad9.net
- Select enabled
- Domain: (leave this field empty)
- Server IP: 9.9.9.9
- Server port: 853
- Veify CN: dns.quad9.net
When all of the DNS servers are selected click on Save:
Please note that if you’ve experience an issue with Unbound TLS forwardings, crashing Unbound with high CPU, please check Firewall > Diagnostics > Live View > port is 853 if none of the requests are being blocked
