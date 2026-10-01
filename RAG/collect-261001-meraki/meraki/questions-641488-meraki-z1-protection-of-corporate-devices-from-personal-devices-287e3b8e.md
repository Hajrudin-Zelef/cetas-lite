---
id: collect-261001-meraki/meraki/questions-641488-meraki-z1-protection-of-corporate-devices-from-personal-devices-287e3b8e
title: "questions-641488-meraki-z1-protection-of-corporate-devices-from-personal-devices-287e3b8e"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["compute"]
source: docs/RAG/collect-261001-meraki/questions-641488-meraki-z1-protection-of-corporate-devices-from-personal-devices-287e3b8e.md
source_anchor: ""
source_lines: [1, 8]
sha256: a9babd63aca8380d2ad7b256092d94e3d9426e427383ea965de0425ac1349833
---

# questions-641488-meraki-z1-protection-of-corporate-devices-from-personal-devices-287e3b8e

My company employees about 600-700 remote employees, each of whom is issued a corporate IP phone an a thin client compute endpoint.
In our current deployment, both devices are configured (defaults settings, vpn, etc) at our corporate office and shipped to the end users.
We have a couple issues with this current deployment, chief of among these that our support team has almost no visibility into our employees home network when they need support (making it difficult to determine if the issue is on our side or with the employees ISP.
We are considering swapping over our standard infrastructure deployment to include a managed access point (likely the Cisco Meraki z1) for better visibility into our employees home network experience. One proposal for this roll out includes us asking the employees to connect as many of their home devices to through their Meraki as possible (to better enable us to provide QOS guarantees to their work devices).
This change makes me uneasy in two ways:
1) Currently, both devices issued to our employees manage their own vpn settings, with all other ports of communication locked down, so we feel very confident through our audits that the devices only talk through encrypted channels. The suggested deployment with the Meraki includes a new network link not managed by the VPN (the Meraki itself would be connected over VPN, the devices will not be)
2) Perhaps there is not anything significantly different between connecting the devices to the Meraki AP vs the current deployment where employees have their devices connected to their home routers directly, but I find myself wishing I better understood the stateful firewall settings on the Meraki z1 to seperate the home devices from the corporate devices.
Are there any experienced Meraki users or network security professionals out there who can comment on either of these concerns?
