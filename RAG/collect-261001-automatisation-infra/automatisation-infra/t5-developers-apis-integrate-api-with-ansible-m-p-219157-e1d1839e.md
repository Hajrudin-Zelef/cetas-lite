---
id: collect-261001-automatisation-infra/automatisation-infra/t5-developers-apis-integrate-api-with-ansible-m-p-219157-e1d1839e
title: "t5-developers-apis-integrate-api-with-ansible-m-p-219157-e1d1839e"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/t5-developers-apis-integrate-api-with-ansible-m-p-219157-e1d1839e.md
source_anchor: ""
source_lines: [1, 72]
sha256: e0ba9d7bcddc82dd77eb601b58d08cd08e528515e96e22a64c0ad1bde4529817
---

# t5-developers-apis-integrate-api-with-ansible-m-p-219157-e1d1839e

Integrate API with ansible
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-12-2023 02:08 PM
Hello everyone, I am currently finding the Cisco Meraki collection module in Ansible to obtain the IP information of my servers, Network, Name, Model, Mac Address among others, this is in order to obtain all the data, What is the module?
Kind regards
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-12-2023 03:57 PM
To gather information about your servers, such as IP, network, name, model, MAC address, and other data using Ansible, you can use the cisco.meraki.networks module. This module is part of the cisco.meraki collection and allows you to manage operations like creating, updating, and deleting networks within the Meraki environment.
Here’s an example of how you can use the module in an Ansible playbook:
- name: Get Meraki network information
cisco.meraki.networks:
auth_key: your_meraki_auth_key
state: query
org_name: your_organization_name
net_name: your_network_name
delegate_to: localhost
Make sure to replace your_meraki_auth_key, your_organization_name, and your_network_name with your actual data. This playbook will query and return information about the specified network.
For more details and configuration options, you can refer to the official documentation of the cisco.meraki.networks module. Remember that you’ll need to have the cisco.meraki collection installed to use this module. If you haven’t installed it yet, you can do so with the following command:
ansible-galaxy collection install cisco.meraki
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-14-2023 11:31 AM
Hello thank you very much for the information, yes I was just reviewing the following links based on the cisco meraki module:
Thanks to the playbook that you shared with me, I have a question, the auth_key parameter, having the value of the parameter, communicates directly with the cisco meraki platform?
My second question is, running this playbook, it will automatically connect to my cisco meraki?
Best regards
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-15-2023 01:35 AM
Please, if this post was useful, leave your kudos and mark it as solved.
