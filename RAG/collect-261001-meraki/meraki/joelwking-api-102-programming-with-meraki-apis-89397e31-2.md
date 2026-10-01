---
id: collect-261001-meraki/meraki/joelwking-api-102-programming-with-meraki-apis-89397e31-2
title: "joelwking-api-102-programming-with-meraki-apis-89397e31"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-meraki/joelwking-api-102-programming-with-meraki-apis-89397e31.md
source_anchor: ""
source_lines: [82, 100]
sha256: 09632252438fdba62847d5b044c68a6c9f9dae639c2ae87f79703a18b1b60acf
---

# joelwking-api-102-programming-with-meraki-apis-89397e31

- 15.
- 16. Phantom Cyber About us Phantomis the first community- powered security automation & orchestration platform. Phantom’s open and extensible architecture helps you work smarter, respond faster, and strengthen your defenses. Visit www.phantom.us/join to get the Free Community version of Phantom.
- 17. Meraki Dashboard Integrated withinPhantom Cyber’s application framework Locate devices usingAPIs
- 18.
- 19. • Simplified WebGUI to the Meraki dashboard • Enforces business rules for deployment by 3rd party ‘feet on the street’ • Deploying 50 locations per week • Technician uses laptop and barcode scanner • Customer supplied CSV file defines street address, GPS locations, network addressing, etc. Meraki Provisioner Asynchrony Labs, creates and deploying custom applications and mobile solutions leveraging Cisco Meraki technology. www.asynchrony.com/
- 20.
- 21. Meraki as amanagement switch in the data center • NexusTop of Rack (ToR) switch mgmt interface connected to Meraki switch / security appliance • Mgmt interface receives IP address via DHCP • API on Service Request System provides MAC address and hostname ofToR switch • Ansible playbook: – Queries Service Request System for Nexus switch MAC and hostname, Meraki networkId andVLAN – Queries Meraki API to determine mgmt IP assigned via DHCP – Creates hostname entry in Infobox – Generates Fixed IP assignments via Meraki API INFOBLOX SERVICE REQUEST SYSTEM DASHBOARD.MERAKI.COM CONTROL NODE MERAKI SWITCH NEXUS TOP OF RACK
- 23. Result Benefits: • Meraki asa management switch forTop of Rack (ToR) data center switches simplifies operations • No static IP configuration required onToR switch interface mgmt0 • Subsequent playbooks reference ToR switch by the Infoblox DNS entry • PowerOn Auto Provisioning (POAP) USB device - for initial (admin userid and password) config Resources: https://github.com/joelwking/code-samples/blob/master/api-server/SimpleHTTPclient.yml https://github.com/joelwking/devnet-create-meraki-api/blob/master/management_switch/meraki_mgmt_sw.yml
- 24. Workshop Project • Ansibleplaybook uses the URI module to update the fixedIpAssignment • Meraki dashboard uses a HTTP 302 redirect • Python Requests module handles redirects Write Ansible module to replace URI module for this function
- 25.
- 26. Meraki Partner Portal www.merakipartners.com MerakiDeveloper Portal developers.meraki.com/tagged/Automation Experimental Python library leveraging theCisco Meraki Provisioning RESTfulAPI github.com/meraki/provisioning-lib Phantom | Meraki App | Blog github.com/joelwking/Phantom-Cyber#meraki-app blog.phantom.us/2016/05/09/community-double-play/ Ansible - 3WaysToTryTower Free www.ansible.com/tower-trial Resources Eligible partner engineers* that attend will receive a free stack of promo equipment and licensing.
- 27. Create synergy betweendevelopers and engineers We learn best by a physical manipulation of our environment – Meraki APIs and gear Use cases and workshops facilitate learning Key take-aways
- #2 API 102: Programming with Meraki APIs.
- #6 Our customer base is migrating from traditional WAN network infrastructure using CLI (command line interface) to cloud managed solutions, like cisco Meraki. While the cloud managed dashboards increase network operations efficiency, API interfaces enable programmatic interfaces to other systems within the enterprise.
- #8 Traditional networking job roles will be evolving to software-enabled network roles.
- #9 (1) A Beginner's Guide to Integrated Development Environments – Mashable mashable.com/2010/10/06/ide-guide/
- #11 http://developers.meraki.com http://dashboard.meraki.com/api_docs If you have any API related issues or need help, please email support@meraki.com If you are looking for an update on the latest new API features please visit developers.meraki.com/news If you are looking for where to get started with the Meraki developer platform, please visit developers.meraki.com/start For sales and escalations, please contact api-team@meraki.com
- #13 Network function virtualization (NFV)
- #18 https://github.com/joelwking/Phantom-Cyber/tree/master/meraki
