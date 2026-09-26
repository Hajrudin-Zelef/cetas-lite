---
id: collect-260926-mikrotik/mikrotik/questions-1727243-how-do-i-reset-the-default-configuration-on-a-mikrotik-chateau-427bb092
title: "questions-1727243-how-do-i-reset-the-default-configuration-on-a-mikrotik-chateau-427bb092"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/troubleshooting/questions-1727243-how-do-i-reset-the-default-configuration-on-a-mikrotik-chateau-427bb092.md
source_anchor: ""
source_lines: [1, 3]
sha256: ac3738a6b6a24139ad92beabbf72a63488837285a7a2e7669f1f8591b884c731
---

# questions-1727243-how-do-i-reset-the-default-configuration-on-a-mikrotik-chateau-427bb092

This should be simple, but it's proving to be quite difficult. I recently opened a brand new MikroTik Chateau LTE12. The default configuration has pretty typical with WiFi and LTE enabled. It was installed with a beta firmware version so I upgraded to the latest version. After rebooting, the device has lost all configuration. It's in a "no configuration" state and not even the LEDs work. Once I enable some configurations, I can get the network LED to work, but the wireless, LTE and power LEDs do not. The easiest way to fix this seems to be to restore the factory configuration that provides a working state, however I cannot seem to figure out how. I've tried doing a reset via Winbox, CLI, and Netinstall with a clean install. Nothing restores a working configuration. I also tried holding the reset button and for 3-5 seconds until the power LED flashes. This didn't work either.
1 Answer 1
Well, apparently this is by design. On the Chateau devices, the power LED is not turned on unless the LTE interface is up and connected despite the fact that the devices has dedicated LTE LEDs.
