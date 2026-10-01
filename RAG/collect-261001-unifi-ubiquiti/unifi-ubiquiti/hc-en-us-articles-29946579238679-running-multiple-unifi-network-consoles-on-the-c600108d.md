---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29946579238679-running-multiple-unifi-network-consoles-on-the-c600108d
title: "hc-en-us-articles-29946579238679-running-multiple-unifi-network-consoles-on-the--c600108d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29946579238679-running-multiple-unifi-network-consoles-on-the--c600108d.md
source_anchor: ""
source_lines: [1, 19]
sha256: d70c6574fd7816b19d11266baf8a89be16658ad0b557de16dfe87d36440d3db7
---

# hc-en-us-articles-29946579238679-running-multiple-unifi-network-consoles-on-the--c600108d

Running Multiple UniFi Network Consoles on the Same Site
You should only run only one Cloud Gateway or one Shadow Mode (high availability) pair of Cloud Gateways at a given site. (For a step-by-step guide on configuring Shadow Mode, click here.)
However, in large-scale deployments, you can run more than one console at your site to offload UniFi Protect, Access, Talk, and Connect.
If you are not using Cloud Gateways, you should run only one instance of UniFi Network (e.g., CloudKey or self-hosted Network deployment) at a given site.
Scaling Your Network Efficiently
If your network requires greater scale, consider these options:
- UniFi Cloud Gateways scale seamlessly for growing UniFi Network deployments, up to 5,000 clients on the Enterprise Fortress Gateway. See more here.
- Offload services to dedicated consoles for better performance:
By offloading resource-heavy services, you ensure optimal performance and scalability while keeping your network streamlined.
Site Manager and Vantage Point
If you have multiple UniFi consoles at the same location, use Site Manager to group them for centralized management. This allows you to control settings across consoles without manually switching between them. To group sites:
- Open Site Manager in your browser.
- Click on the Site Group dropdown and select Add Site Group.
- Create your group and click Add.
In addition to Site Groups, Site Manager also supports merging UniFi consoles that are on the same UniFi network. This is particularly useful when you add a second UniFi console to your network, but do not want to manage it as a separate UniFi site. For example, if you are using a UniFi Cloud Gateway to run UniFi Network and want to add a UniFi NVR running Protect to the same site. To merge consoles:
- Open Site Manager in your browser.
- There can only be one instance of each UniFi application running on any given site. Before merging, stop any duplicate applications (Protect, Access, Talk, Connect) on your site's Control Plane.
- Locate the UniFi console you wish to integrate into your other UniFi console and click Merge. It will collapse under the primary console for the site.
For setups with multiple Protect consoles, Vantage Point provides a unified dashboard, making camera management more efficient. To learn more about Vantage Point, click here.
