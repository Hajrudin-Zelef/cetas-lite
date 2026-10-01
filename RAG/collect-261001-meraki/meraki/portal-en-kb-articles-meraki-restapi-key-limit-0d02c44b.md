---
id: collect-261001-meraki/meraki/portal-en-kb-articles-meraki-restapi-key-limit-0d02c44b
title: "portal-en-kb-articles-meraki-restapi-key-limit-0d02c44b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/portal-en-kb-articles-meraki-restapi-key-limit-0d02c44b.md
source_anchor: ""
source_lines: [1, 9]
sha256: efe9181da844a1fa76519f05e8bbb76f6adafbdd21ad3be40451e8ba69b57205
---

# portal-en-kb-articles-meraki-restapi-key-limit-0d02c44b

Data collection in Site24x7 will be skipped for that particular Meraki organization and its associated devices if the REST API key exceeds the prescribed limit. To prevent this, enable alerting by following the steps below:
 
- Log in to your Site24x7 account.
- Click Network > Meraki > Meraki organization > Inventory > Threshold and Availability.
- Click the pencil icon beside the Threshold and Availability field.
- Once the Add Threshold Profile pop-up displays, under the Threshold Configuration section, toggle the Alert if the REST API key limit is reached for the Meraki organization option to Yes. This is set to Yes by default.
      5. Save your changes.
Once you save your changes, alerts will be triggered whenever the REST API key limit is reached for a Meraki organization. You can also configure alerts based on the thresholds of various performance metrics of your Meraki device.
Updated: 2 years ago
