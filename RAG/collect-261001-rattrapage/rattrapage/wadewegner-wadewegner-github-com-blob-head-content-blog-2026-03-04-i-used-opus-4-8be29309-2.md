---
id: collect-261001-rattrapage/rattrapage/wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309-2
title: "wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309.md
source_anchor: ""
source_lines: [93, 130]
sha256: e47bc33b0efe827b4da778214d8b765299e4b8389d670606eee7c0629e9dd92e
---

# wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309

Every radio at max transmit power. When all your APs are screaming at full volume, clients see strong signals from multiple APs and don't roam efficiently. The LLM lowered 2.4GHz power on most APs to create cleaner cell boundaries, keeping power high only where a distant IoT device needed the extra reach.
An IoT device that disconnected after optimization. This is the kind of thing that makes network tuning scary — you change settings and something breaks. The LLM diagnosed it immediately: the minimum RSSI threshold was kicking the device because lowering TX power had pushed its signal just below the threshold. The fix was surgical: relax the threshold slightly and bump power back up on the one AP closest to that device.
Zero network segmentation. Every device — laptops, phones, NAS, Ring cameras, smart fridge, sauna controller — was on the same flat network. A compromised IoT device could reach everything. The LLM created an IoT VLAN, a dedicated IoT SSID, and firewall rules that block IoT devices from initiating connections to the trusted network while still allowing your phone to control them.
Missing security features. DNS-over-HTTPS was configured but not turned on. DNS filtering was available but disabled. Protected management frames were off. NAT-PMP was still enabled (an attack surface). All easy fixes via the API.
Firmware updates. Every device had pending firmware updates. The LLM triggered them all via the API, and when the controller software itself was updated (which I did manually), the devices all showed as "offline" until the LLM force-provisioned them to reconnect.
Things I learned that'll save you time:
PowerShell and JSON don't mix well. If you're on Windows, don't try to pass JSON directly in curl commands. PowerShell's escaping rules will mangle it. Instead, write your JSON payloads to files and use curl -d @filename.json. The LLM figured this out after the first failed API call and switched strategies automatically.
Some settings silently ignore your values. For example, setting a minimum data rate required also setting minrate_setting_preference to "manual" — otherwise the API returned success but didn't actually change anything. The LLM caught this by verifying the config after each change, saw the value hadn't stuck, and figured out the missing parameter.
"Medium" TX power doesn't always mean medium. On some AP models, setting tx_power_mode to "medium" was accepted but the actual transmit power didn't change. Using "custom" with an explicit dBm value worked. These are the kinds of API quirks that are painful to discover manually but that an LLM can brute-force through quickly.
Hardware limitations are real. We tried to enable WPA3 but the APs were Wave 1 hardware that doesn't support it. We tried to enable IDS but the gateway didn't have the processing power. The LLM identified these limitations from the API responses, logged them, and moved on instead of banging its head against the wall. It also noted them as reasons to consider hardware upgrades down the road.
Firewall rule indices have a specific valid range. The API requires rule indices matching a particular regex pattern, and the "obvious" values (2000, 4000) were actually out of range. The LLM discovered the valid pattern from the error response and adjusted. This is the kind of API archaeology that takes a human 30 minutes of frustrated Googling.
Always verify after firmware upgrades. Firmware updates can reset settings. After upgrading all devices and the controller software, we did a full verification pass to confirm every configuration change had survived. They all did in our case, but it's worth checking.
Here are the prompts that worked well, roughly in order:
- Kickoff: "I'd like your help optimizing my UniFi network. Context is in context.md."
- Analysis: "Please write an analysis.md file with your findings and a recommendations.md file with what you suggest we do."
- Execution: "Let's proceed with batch 1. Please create a log.md and track everything you do."
- Course correction: "What if we just turn off [AP name] and go with three APs? My house isn't that big, and it might help with conflicts." (Don't be afraid to push back or suggest alternatives.)
- Troubleshooting: "I can't get [device] to connect. Can you see what's happening?" (The LLM can pull client stats and diagnose the issue.)
- Impact assessment: "For batch 2, what is the impact to clients and performance?" (Ask before approving risky changes.)
- Verification: "Can you check to confirm the network looks good?"
- Next steps: "Anything else we should do?" (Good for catching things like post-upgrade verification.)
The pattern is: give context, ask for analysis, review recommendations, approve in batches, verify after each batch, and keep a log.
Download a backup after each batch of changes. In the UniFi controller UI, go to Settings → System → Backup and download a copy of your config. Do this before you start, and again after each batch. If something goes sideways days later and you can't remember what changed, you can restore to any checkpoint. The LLM can also trigger a backup via the API (POST /api/s/default/cmd/backup), so you can even bake it into the workflow.
Disable or delete the local admin account when you're done. You created it for API access, and now the work is finished. There's no reason to leave an extra admin account sitting around with no 2FA. Either delete it entirely or at minimum disable it. You can always recreate it if you want to run another optimization session later.
A few reasons this turned out better than I expected:
The LLM is tireless at API exploration. It pulled 15+ endpoints, parsed the JSON, cross-referenced device configs with client stats, and identified patterns across dozens of devices. This would take a human an hour just for the data collection.
It catches things you'd miss. I didn't know my 5GHz channels were at HT40 when they could be VHT80. I didn't know NAT-PMP was enabled. I didn't know the minimum data rate was set to 1 Mbps. These are settings buried deep in the config that you'd only find by reading every field of every API response.
It recovers from errors gracefully. When the PowerShell JSON escaping broke, it switched to file-based payloads. When a setting didn't apply, it investigated why and found the missing parameter. When a device disconnected, it diagnosed the root cause and applied a targeted fix. Each failure became a learning moment that it logged and adapted to.
The approval-gate pattern keeps you in control. At no point did the LLM go rogue. Every change was proposed, explained, and waited for my "go ahead." The log file means you have a complete audit trail. And rollback instructions were included for every batch.
It's reproducible. The context file, analysis, recommendations, and change log together form a complete record. If you need to redo it (new house, new hardware, friend asks for help), you have a template.
A few limitations worth noting:
- 
It can't move devices between SSIDs. WiFi clients store their own credentials. When you create a new IoT SSID, you have to manually reconnect each device via its companion app or settings screen. There's no API for "tell this Ring camera to use a different network."
- 
It can't fix hardware limitations. If your APs don't support WPA3 or your gateway doesn't have IDS capability, no amount of API calls will change that. But it will tell you exactly what your hardware can and can't do, which is valuable for planning upgrades.
- 
It doesn't know your physical layout. It can tell you that two APs are on the same channel and that's bad, but it doesn't know which APs are physically close together. You'll need to provide that context or correct its assumptions. In my case, I told it which AP was closest to a problematic device, and it adjusted its approach accordingly.
- 
