---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e-3
title: "how-to-install-and-configure-opnsense-db0f659e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e.md
source_anchor: ""
source_lines: [63, 79]
sha256: 2559598c59b612776e055b23cc403815eba9ca75f838660d0d1ff138f0e09fc7
---

# how-to-install-and-configure-opnsense-db0f659e

The WAN interface configuration page has a bunch of settings available since there are various ways to connect to the Internet. If you happen to have an ISP where you can use DHCP, you may simply leave everything at the default setting and click “Next”. However, other ISP configurations may be more complex and OPNsense provides a number of ways you can connect to your ISP.
If you are planning to use your OPNsense router behind your ISP router, you will need to uncheck the “Block RFC1918 Private Networks” and the “Block bogon networks” boxes so that your WAN interface can operate correctly on your local network. Otherwise, all local network traffic will be blocked and you will have trouble accessing the Internet through your ISP provided router.
For the sake of simplicity of this guide, I am going to assume you are using OPNsense as your primary router. Click “Next” once you have entered the appropriate settings.
In comparison to the WAN interface, the LAN interface settings appear to be very simple. The setup wizard does not provide the full set of available configuration options for the LAN interface (possibly due to the fact you could end up losing connection or locking yourself out of the web interface if you are not careful).
Keep in mind if you change the default network addresses for the LAN, you will lose connection at the end of the wizard and will either need to reload your DHCP lease or disconnect/reconnect to your network to obtain a new IP address (that is assuming the wizard also sets up the appropriate DHCP address ranges – I have not personally tested it).
To keep things simple for a basic OPNsense installation, simply click “Next” without making any changes.
If you already changed your root user password during the installation process, simply click “Next” since you do not need to change it again. This would be a great time to change the default password if you did not do so during the installation process.
It would be quite silly to leave the default password unchanged when you are installing a very secure router/firewall OS like OPNsense – do not leave the front door unlocked in an otherwise secure building!
Click “Reload” to apply all of the changes you have made so far. If you changed the hostname/domain name, you may need to enter the new host/domain name to access the web interface again or simply use the IP address of the LAN interface.
You will see a status message of the configuration reloading.
All changes have been applied!
Next Steps
If your basic network is functioning properly with your new OPNsense installation, I would like to say congratulations! You have taken the first step in learning more about securing your home network. You may wish to take some time to get familiar with the configuration options available on the web interface. If you have tinkered with the settings in consumer grade routers in the past, you may come across several settings which look similar. However, there are many more knobs and dials to turn in OPNsense.
It is worth noting that by default all incoming connections to both IPv4 and IPv6 addresses are blocked by default and all outgoing connections are allowed much like a consumer grade router. There is a basic level of protection in place for incoming connections so you do not need to worry about being completely vulnerable and exposed with the default installation.
When you are ready to implement more security measures beyond the default configuration, I have compiled a non-exhaustive list of features you may wish to consider.
This site covers a variety of OPNsense related topics as well as other home networking information such as how to configure your network switches to use VLANs. Whenever you are ready to tackle a new topic, be sure to search this site to see if there is a guide to help you.
Please feel free to contact me about suggestions for new topics, and I will see what I can do to try to cover it. Because I do not do this full time, I operate this site using the few brief moments of time I have available to work on it.
