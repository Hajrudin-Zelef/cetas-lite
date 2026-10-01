---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d-1
title: "setup-adguard-home-opnsense-adblocker-c464187d"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["revenue"]
source: docs/RAG/collect-261001-opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d.md
source_anchor: ""
source_lines: [1, 51]
sha256: 7eae19e25351c9b735fe4de7e9aed51b5680660f1b86d5f4ddf082075484f8f0
---

# setup-adguard-home-opnsense-adblocker-c464187d

Why do I need AdGuard Home as my internal DNS resolver?
AdGuard Home is a free and open-source software that can block ads and tracking on all devices connected to your home Wi-Fi network. Unlike traditional ad blockers that only work on specific devices or browsers, AdGuard Home covers all devices without the need to install any client-side software. Additionally, AdGuard Home offers a range of features beyond ad-blocking and tracking prevention, including traffic encryption, making it a comprehensive solution for controlling and securing your home network.
The majority of websites we browse nowadays come equipped with supplementary components for advertising, analytics, and engagement tracking. While these tools can be useful for website owners to generate revenue and gain insight into their audience’s preferences,
Wondering how AdGuard Home works? In essence, it serves as a DNS server that reroutes ad and tracking domains to a “black hole,” effectively preventing your devices from establishing connections to those servers. This is the same software that powers our reliable and tested public AdGuard DNS servers, and AdGuard Home shares many similarities in code. As a result, AdGuard Home can manage traffic for virtually any internet-connected device, including smart TVs, refrigerators, and even light bulbs.
In simple words, the Domain Name System (DNS) is a set of databases that convert hostnames into IP addresses. DNS is like a phone book for the internet, where easy-to-recall hostnames such as www.google.com are transformed into IP addresses like 216.58.217.46. It occurs seamlessly behind the scenes once you enter a URL into your web browser’s address bar.
To accomplish this, your computer contacts its designated DNS server and provides the website name, such as google.com. The DNS server searches for the IP address associated with that domain and supplies your computer with the result, such as 203.0.113.52. Following this, your computer can connect to the designated address and load the website content.
We can derive several benefits from using DNS-level blockers, including enhanced privacy and reduced advertising clutter when browsing the web. Additionally, by blocking unwanted content at an early stage, we can minimize bandwidth usage and data costs. This may also lead to slight performance improvements, as there is less content to load for each site. As an example it reduces the bandwidth use if you’re connected via mobile network.
Another significant advantage is security, as several DNS blocklists continuously update with the latest suspicious or malicious domains. Blocking and preventing clients from connecting to these domains quickly can enhance our security.
However, there are a few downsides to using DNS-level blockers. For instance, many of these blockers draw from various website blocklists, which may not always be entirely accurate. This means that some advertisements or tracking may still appear. Additionally, legitimate parts of some sites may also get blocked, and several websites nowadays require the loading of third-party components to work correctly. Although most things work correctly, users should be aware that some troubleshooting and manual unblocking of website components may be necessary at times.
STEP1: Decide on the AdGuard DNS implementation
There are plenty of ways to implement the AdGuard in your network. The Adblocker on all major platforms Windows, Linux, Docker and could be implemented into the network devices such as firewalls and routers. Please check the official GitHub repository.
If you already running OPNsense, instead of linking the firewall to the AdGuard running on small server you can embed it into the firewall system. This way you’ll get rid of additional failure point for your Internet, if you’d like to resolve all the traffic by internal DNS.
However, if you prefer to set up AdGuard (or Pi-hole, or others) on other separate machine, it is also fine. Following this approach you’ll need to update your client network’s DHCP options to use the desired, internal DNS servers.
Please keep in mind that the AdGuard DNS may not be suitable in Enterprise scenarios with a need of complex traffic filtering for thousands of machines. For that kind of purpose is better to stick to commercial ZenArmor plug-in for OPNSense or hardware Proxy like Zscaler or Symantec EdgeSWG.
NOTE: Adguard will only filter DNS requests with the domain name, it won’t prevent machine to connect directly to Public IP address!
STEP2: AdGuard + OPNsense topology
In the original scenario your computer uses external DNS provided by the DHCP and sends a request via your firewall to the public DNS and receives response on the default route to reach the specific website.
This is how the classic DNS request is handled:
Once you implement the AdGuard internal DNS on your firewall, it will act as a mediator between public DNS and will filter traffic from unwanted requests and make the requests sent to the public DNS significantly more secure.
Here is the request flow via internal DNS to public DNS resolver:
STEP3: Add the Community Repository to OPNsense
By default, AdGuard Home is not included in the available plugins to download/install in OPNsense. In order to extend the list of plugins, we need to add the community plugin repository that includes a list of additional packages.
Before we can install the AdGuard Home plugin, we will need to setup & install that community repository.
To do this, we’ll need direct SSH or console access to our OPNsense appliance from our LAN.
SSH is disabled by default, but we can enable it quickly by navigating to System > Settings > Administration and then scrolling down to the Secure Shell section.
We’ll need to check the box for Enable Secure Shell and Permit Password Login. If you’re logging into OPNsense with the root account, you’ll also need to select Permit root user login.
Then scroll down to the bottom of the page & click Save.
By default OPNsense will set SSH Listen Interface set to All. It is recommend to enable access only LAN interface.
If you don’t need SSH to be accessible all the time, please remember to disable this service once you’re finished setting this up.
Okay, now that’s enabled – we can connect to our OPNsense appliance using your preferred SSH client (like PuTTY).
If you’re using the root account, you’ll likely be dropped into the OPNsense shell – but you can select option 8 here to access the underlying FreeBSD shell.
In order to install the community repository, we’ll pull down the repository config file using the following command:
fetch -o /usr/local/etc/pkg/repos/mimugmail.conf https://www.routerperformance.net/mimugmail.conf
Then, we’ll need to ask OPNsense to update it’s local cache with the new repo – so it knows what packages are hosted there:
pkg update
If everything is successful, you’ll see output similar to below – which lists the mimugmailrepository now:
root@0xOPNsense:/home/matt # pkg update
Updating OPNsense repository catalogue...
Fetching meta.conf: 100%    163 B   0.2kB/s    00:01
Fetching packagesite.pkg: 100%  229 KiB 234.3kB/s    00:01
Processing entries: 100%
OPNsense repository update completed. 822 packages processed.
Updating mimugmail repository catalogue...
Fetching meta.conf: 100%    163 B   0.2kB/s    00:01
Fetching packagesite.pkg: 100%   54 KiB  54.8kB/s    00:01
Processing entries: 100%
mimugmail repository update completed. 177 packages processed.
All repositories are up to date.
After these steps, you no longer need opened SSH connection, so please close the port 22 by going to System > Settings > Administration and then scrolling down to the Secure Shell section.
Step4: Installing the AdGuard Home Package
Now that the additional package repository is set up, we can download & install the AdGuard Home plugin via the OPNsense web interface.
