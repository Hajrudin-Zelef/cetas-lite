---
id: collect-261001-meraki/meraki/questions-55520-cisco-meraki-switch-sfp-compatibility-and-dual-link-a0dfbf5f
title: "questions-55520-cisco-meraki-switch-sfp-compatibility-and-dual-link-a0dfbf5f"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "wavelength"]
source: docs/RAG/collect-261001-meraki/questions-55520-cisco-meraki-switch-sfp-compatibility-and-dual-link-a0dfbf5f.md
source_anchor: ""
source_lines: [1, 5]
sha256: 4cbfbc91ac97d41852402d587f245bb9805bdfbb93b567921fb651b4437362d0
---

# questions-55520-cisco-meraki-switch-sfp-compatibility-and-dual-link-a0dfbf5f

Our company wants to connect two different floors together about 150-200 feet apart. I was thinking fiber would be the best option for this but I don't have much experience so I would like to see if what I am thinking of doing would work. We have a Cisco Meraki MX-84 security appliance and a Cisco SG200-50 switch.
Based on my research it appears that the Meraki MA-SFP-1GB-SX transceiver is compatible with the MX-84 and the Cisco MGBSX1 transceiver is compatible with the SG200-50. I assume that these two transceivers would work together properly based on the fact they both operate at 850 nm wavelength and are both multimode fiber. Can anyone confirm if my assumption is correct?
Also I am wondering if it makes sense to have two links to each switch rather than a single link? I don't know if there is any benefit to running two links other than if one link fails it still has another link to work off of.
Docs for the Meraki SFP accessories https://meraki.cisco.com/lib/pdf/meraki_datasheet_sfp.pdf
Cisco SG200-50 specs https://www.cisco.com/c/en/us/products/collateral/switches/small-business-200-series-smart-switches/data_sheet_c78-634369.html
