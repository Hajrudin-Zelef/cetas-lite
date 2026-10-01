---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-hniv2m-multi-site-management-3738076f
title: "r-ubiquiti-comments-hniv2m-multi-site-management-3738076f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-hniv2m-multi-site-management-3738076f.md
source_anchor: ""
source_lines: [1, 16]
sha256: 9ae889970add069ee21694049483e903dc5e366e69d5d7c998edd274fcb7e507
---

# r-ubiquiti-comments-hniv2m-multi-site-management-3738076f

Multi Site Management 
        
    I have used Ubiquiti a little on a single site just looking for some advise on how to manage a Multisite Install.
I have 4 locations in different locations, I want to install Unifi AP in all of these sites is there a away to manage these all from a single UDM Pro at HQ and create some sort of VPN to manage the devices and route the Internet traffic out of a local VDSL connection.
Or is it either a UDM-Pro or Enterprise Gateway on each site.
Any thoughts and advise would be appreciated
Thanks
Section des commentaires
Your going to need a gateway device at each physical site. With the UDMP you have to use the integrated controller, so if you can expose the UDMP to the internet and use a USG at the other sites you may be able to make it work. Best suggestion is a USG-PRO-4 at each site and a cloud hosted controller, or wait for the upcoming UXG to do this once it's released later this year.
If you did a UDMP at each site you can make it work too
The original MultiSite model was as follows:
Create a single, central controller. This could be a cloud-hosted one, a custom-built one on-prem, or a CloudKey.
On that controller, you would create multiple sites for multiple locations, and using L3 adoption, adopt devices from each site to the single controller.
Because the UDM-Pro (and UDM) require you to use its built-in controller, that not possible. The older USG and USG-Pro as well as the upcoming UXG-Pro (and likely a smaller version) do not have controllers, so they fit the multisite model.
If you do not plan to use UniFi switches or Gateways at the remote sites, technically you could probably get VPN's in place so all of the remote AP's can be adapted to the UDM's controller. That's more advanced and not likely what you are looking for.
VPN or publicly accessible routers.
