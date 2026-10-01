---
id: collect-261001-general-networking/general-networking/how-to-use-custom-domain-name-in-internal-network-b6a0530b
title: "how-to-use-custom-domain-name-in-internal-network-b6a0530b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/how-to-use-custom-domain-name-in-internal-network-b6a0530b.md
source_anchor: ""
source_lines: [1, 19]
sha256: 341164e3e75a14c1bb0725243da02afd52729bcc5d0463007ec9ee1b4d7549e9
---

# how-to-use-custom-domain-name-in-internal-network-b6a0530b

Use a Custom Domain Name on Your Local Network Using OPNsense
Table of Contents
One nice thing about utilizing more advanced router features is that you can do neat things such as using your own domain name in your home network. This allows you to refer to devices on your network by using the device name/hostname followed by the domain name. For instance, server.homenetworkguy.com could refer to a host named “server” that exists on the network using the domain name “homenetworkguy.com”.
Of course, depending on your configuration, you could just use the hostname “server” to access various services on your device(s) such as a fileshare, which will be shorter than including the domain name. However, if you plan to run a web server or run software which has a web interface, you may need to use the FQDN (Fully-Qualified Domain Name) rather than just the hostname since it could help facilitate the same-origin policy for improved website security. Also, if you are running multiple websites on the same server using different domain names using SNI, you will need to use the full domain name so the server knows which page to display.
In this how-to, I will be using the OPNsense router software to configure a domain name for a network.
Setting up the Global Domain Name for your Network
To set up the global domain name for your home network, go to the “System > Settings > General” page. Then enter a hostname for your OPNsense router and the default domain name for your entire network. In the example below, I would be able to access the OPNsense web administration page by visiting to https://router.homenetworkguy.com.
Note: It is not necessary to have a real domain name registered if you are just using the domain name internally on your network. However, if you do have a real domain name (regardless if you are hosting anything public on your home network or not), you may also use it internally for devices on your network. If you use a real domain name, your local DNS server on your router will know how to resolve names for devices on your local network even if they those hostnames are not included the public DNS records. You only need to add/update the IP addresses of domains/subdomains of public DNS records if you plan to host publicly accessible services on your network.
Even though you do not have to use a real domain name for your local devices, you cannot simply use any domain name you want. The help message says to avoid using “local” as the domain name since it interferes with the Bonjour protocol, for instance. You should also avoid using domain names of real websites that you plan to visit. I imagine you will not be able to reach such websites since the local DNS server on your router will be trying to resolve that domain name to devices on your local network.
On the same General Settings page, make sure you have the option “Do not use the local DNS service as a nameserver for this system” unchecked because you will exclude the local DNS service from resolving names on your network. This essentially forces DNS lookups to be performed upstream (with your ISP DNS or whatever external DNS servers you have defined) instead of locally.
Register DHCP Hostnames in Unbound DNS
Now go to “Services > Unbound DNS > General”. Ensure you have checked the “DHCP Registration” and “DHCP Static Mappings” options (you only need the latter option if you are using static DHCP mappings). The “DHCP Registration” option allows hostnames from all DHCP clients to be resolvable. It will use the hostname that is set up in each device’s configuration. If you have not configured a hostname for your devices, the devices will typically have a default hostname. Not all devices will have a hostname set so you may see blank hostnames on the DHCPv4 Leases page (“Services > DHCPv4 > Leases”). You will not be able to access the device by its hostname until you set one (either on the device itself or via a static DHCP mapping).
The “DHCP Static Mappings” option will allow hostnames to be resolved when creating a static DHCP mapping for a particular device. If you set a hostname in a static DHCP mapping with this option enabled, you will essentially be using the hostname specified in the static DHCP mapping rather than the hostname configured on the device itself. Keep that in mind if you are expecting the configured hostnames to match. As mentioned above, you can create a static DHCP mapping for any device that does not have a hostname in order to provide it with one (but you will need to give it a static IP address as well).
Create DHCP Static Mappings (Optional)
If you have any static IPs you would like to set for some of your devices such as your server(s), you have the option to specify a hostname. This hostname will be used instead of the hostname set on the device itself if the “DHCP Static Mappings” option is set.
You also have the ability to set a different domain name for the static mapping for the device if you so desire. It is likely unnecessary in most home network environments, but perhaps it is useful if you are hosting multiple domains on your network.
Viewing a List of All Hostnames
The easiest way to see a list of hostnames for all of your network devices is to go to “Leases” page by navigating to “Services > DHCPv4 > Leases”. It will show the MAC addresses, hostnames, IP addresses, online/offline status, and if it is statically mapped. The page is fairly self-explanatory. I am not including a screenshot since I would have to blot out most of my device information (to satisfy my security-minded paranoia).
Now you should be able to refer to devices in your home network by their hostname and domain name (the FQDN)!
