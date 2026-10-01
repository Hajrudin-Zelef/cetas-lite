---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-install-and-configure-crowdsec-on-opnsense-a6785fd1-2
title: "how-to-install-and-configure-crowdsec-on-opnsense-a6785fd1"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-install-and-configure-crowdsec-on-opnsense-a6785fd1.md
source_anchor: ""
source_lines: [26, 55]
sha256: 5920afdafbd8abf1c806c85d2e4325a105caf14feeb73943ea8cddd037aff18e
---

# how-to-install-and-configure-crowdsec-on-opnsense-a6785fd1

| Direction | in (needs to be “out” if you select the WAN interface – see note below) | 
| TCP/IP Version | IPv4 (or IPv6) | 
| Protocol | any | 
| Source | any | 
| Source Port | any | 
| Destination | crowdsec_blacklists (or crowdsec6_blacklists for IPv6) | 
| Destination Port | any | 
| Description | Block outgoing connections to IPs on the CrowdSec block list | 
Note
If you select the WAN interface, the direction needs to be set to out because you want to filter traffic leaving the firewall. For all of your internal networks, using in works because all traffic from devices on your network enters into each network interface. Therefore, outgoing traffic can be blocked with the direction of in. Using the direction of out for your internal networks will work, but it is less efficient to process firewall rules on the local interfaces using the out direction.
If you only want to filter on the WAN interface in the floating rule, you could simply create a rule with direction out on the WAN interface itself.
Register for the CrowdSec Console (Optional)
If you wish to take advantage of the free CrowdSec Console, go to the CrowdSec Console registration page. Once you have created an account, you can add your OPNsense instance to the Console by running the command shown on the “Instances” page, which is the default page which opens after logging in. If you already have at least one instance, you can add another instance by clicking on the “Add Instance” button on the “Instances” page.
Once you run the command above (using the shell via SSH) on your OPNsense system, you will be prompted for an enrollment request when you refresh the webpage.
After enrollment is complete, you will see the instance in the CrowdSec Console.
As you can see, it is very simple to add your CrowdSec instance to the CrowdSec Console!
Test that CrowdSec is Operational
You should try testing CrowdSec after everything is set up to ensure proper functioning. One way is to manually add a temporary ban entering an IP address of your choice (if you use the same IP you are currently logged in with, you will lose access to SSH for 1 minute). You should notice that you are temporarily locked out of SSH after running the following command:
sudo cscli decisions add --ip 192.168.1.10 --duration 1m
Since I was running OPNsense in a virtual machine while testing CrowdSec, I actually just tried entering the SSH password incorrectly several times, and I was immediately locked out of SSH for a bruteforce attempt. I believe the ban only lasted maybe 24 hours since I was able to log back in the next day without needing to revert my virtual machine. That was a simple real world test to see CrowdSec in action.
If you registered for the CrowdSec Console, you will see the event under the “Alerts” page.
Take it to the Next Level: Multi-server CrowdSec Installation
A CrowdSec agent can be installed on multiple systems, virtual machines, and containers on your network to monitor for malicious activity and report back to the local API running on OPNsense or another system. CrowdSec considers this a multi-server installation, which is a more advanced and powerful use case officially supported by CrowdSec (even by the CrowdSec Console). This allows you to monitor important services running on your network and take immediate action if anything potentially malicious has been detected.
There are two primary ways you can set up a multi-server CrowdSec environment: running the local API on OPNsense, which is the default installation of CrowdSec or by hosting the LAPI on some other internal server on your network. In either scenario, the bouncer continues to run on the OPNsense firewall to protect your entire network from traffic originating from malicious IP addresses. The main difference between the two options is where all of the CrowdSec agents and bouncers access the local API. If your OPNsense router is constrained on hardware resources or you prefer to minimize the number of services running on your router, running a separate LAPI server would be the ideal option to choose.
To keep this guide focused and a reasonable length, I will only describe the multi-server installation at a high level. It will be a good segue to a more advanced CrowdSec guide, which may be beneficial to readers who expose various services to the Internet (or perhaps simply to increase monitoring within your internal network(s)).
CrowdSec with Local API (LAPI) on OPNsense
As you can see in the diagram below, the default installation of the CrowdSec plugin has the CrowdSec agent, LAPI, and bouncer running on the OPNsense system. CrowdSec agents and bouncers on Server 1 and Server 2 report to the LAPI on OPNsense. The LAPI on OPNsense communicates with the CrowdSec central API (CAPI) to pull updates on the current malicious IP addresses and to report back any useful threat intelligence gathered by the CrowdSec agents.
CrowdSec with Local API (LAPI) on Internal Server
You may also run the CrowdSec LAPI on a dedicated system, VM, container, etc. In this case, the OPNsense CrowdSec agent and the bouncer is configured to communicate with the LAPI located somewhere on your internal network. Server 1 and Server 2 also report to the same LAPI. The LAPI on the internal server communicates with the CrowdSec CAPI rather than OPNsense since it is no longer hosting the LAPI.
I hope this guide helps get you started with the basics of CrowdSec in OPNsense! In the future, I would like to explore a more complex multi-server installation of CrowdSec used to protect self-hosted services that are exposed to the Internet.
