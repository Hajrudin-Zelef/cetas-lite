---
id: collect-261001-meraki/meraki/monkeyfishfrog-meraki-api-projects-44fa6d6a
title: "monkeyfishfrog-meraki-api-projects-44fa6d6a"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/monkeyfishfrog-meraki-api-projects-44fa6d6a.md
source_anchor: ""
source_lines: [1, 28]
sha256: fbc35053840f73b8dd604a1d71b14709175644e88d73b1cdea481efa73b5f48d
---

# monkeyfishfrog-meraki-api-projects-44fa6d6a

This project demonstrates the use of the Cisco Meraki Dashboard API using Postman collections and (optionally) automation scripts.
- ✅ Collection to:
  - List organizations
  - List networks per organization
  - Get connected clients and devices
  - Change switch port configurations
- ✅ Uses collection/environment variables for api_key ,org_id , andnetwork_id
- ✅ Ready for export to CSV or scripting extensions
| File | Description | 
|---|---|
| meraki-api-collection.postman_collection.json | Main Postman collection | 
| postman_environment_template.json | Example Postman environment (API key redacted) | 
| sample-output.json | Example output from API requests | 
| scripts/ | (Optional) Python/Node scripts to run requests | 
- 🔑 Get your Meraki API key
- 🧪 Import the collection into Postman:
  - File → Import → Select meraki-api-collection.postman_collection.json
- File → Import → Select 
- Set your environment variable api_key with your Meraki Dashboard API key.
- Use GET https://api.meraki.com/api/v1/organizations to list your accessible orgs.
- Fill in the org_id from the response, and use it to access networks, clients, and devices.
You can also view and use this collection from Postman’s public workspace (no login required):
- Understand RESTful interaction with Cisco Meraki APIs
- Practice secure API key handling with Postman variables
- Automate network visibility, reporting, and config changes
- Postman
- Cisco Meraki Dashboard API v1
MIT — free to use for learning or portfolio purposes.
