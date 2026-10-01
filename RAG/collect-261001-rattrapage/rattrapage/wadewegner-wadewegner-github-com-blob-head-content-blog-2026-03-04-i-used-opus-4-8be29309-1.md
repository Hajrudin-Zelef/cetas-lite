---
id: collect-261001-rattrapage/rattrapage/wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309-1
title: "wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-03-04"]
keywords: ["claude", "distribution", "opus 4", "throughput", "tool use"]
source: docs/RAG/collect-261001-rattrapage/wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309.md
source_anchor: ""
source_lines: [1, 92]
sha256: 0ccce361f8d4e5e7b073b5967d7c3c0b2ee0be26350c83deb05ec07b52ac147b
---

# wadewegner-wadewegner-github-com-blob-head-content-blog-2026-03-04-i-used-opus-4-8be29309

| title | I Used Opus 4.6 to Audit and Optimize My UniFi Network (And You Can Too) | 
|---|---|
| date | 2026-03-04 12:00:00 -0800 | 
| description | A step-by-step guide to using Claude Opus 4.6 and the UniFi API to audit your home network, fix radio conflicts, add VLAN segmentation, and harden security — all in about an hour. | 
| draft | false | 
Most of us set up our home networks once and never touch them again. Maybe you ran the UniFi setup wizard, picked an SSID name, and called it a day. Everything works, so why mess with it?
Here's the thing: "works" and "works well" are very different. My network had been running for years with co-channel interference on 2.4GHz, every radio cranked to max power, no VLANs, no intrusion detection, and IoT devices sharing a flat network with my NAS. I knew it could be better but never had the motivation to sit down with the UniFi API docs and figure it all out.
Then I saw this tweet from DHH:
Giving Opus access to your ubiquity interface to debug wifi and network issues is unbelievably effective. It's now correctly identified and fixed problems in two different installations I've had that were plagued for months/years with issues. So good.
One conversation later, my network had clean channel assignments, proper VLAN segmentation, hardened security settings, updated firmware on every device, and a detailed log of every change made. The whole thing took about an hour of wall-clock time, most of which was me reviewing proposed changes and saying "go ahead."
Here's how to do it yourself.
- A UniFi network: I'm pretty sure this works with any UniFi setup: a UDM, UDR, or Cloud Key (where the controller is built into the hardware), or a self-hosted UniFi Network Application on Windows/Linux/Raspberry Pi. The API is the same; the only difference is the base path (UDM/UDR uses /proxy/network/api/s/default/ while self-hosted uses/api/s/default/ ). Don't worry, the LLM will figure it out.
- Cursor with a capable model (I used Claude Opus 4.6, but any model with strong tool use should work), or Claude Code if you prefer the terminal
- About an hour of your time (less if your network is in better shape than mine was)
- A willingness to let an AI poke at your network config (with your approval at every step)
This was the only real tricky part for me. If you log into your UniFi controller with a Ubiquiti SSO account (which most people do), that account uses 2FA and doesn't play nicely with API calls from a script.
The fix: create a local-only admin account specifically for API access.
In the UniFi controller UI:
- Go to Settings → Admins & Users → Add Admin
- Choose Local Access Only (not Ubiquiti account)
- Give it a username and a strong password
- Set the role to Administrator (it needs read/write access)
- Make sure Remote Access is unchecked
Now you have credentials that work with simple username/password authentication against the local API. No 2FA, no SSO token dance.
This is the key to making the LLM actually useful. Don't just say "optimize my network." Give it a structured handoff document. I created a context.md file that included:
Authentication details:
The UniFi controller is a self-hosted UniFi Network Application running at https://127.0.0.1:8443.
A local admin account was created for API use:
- Username: <your-username>
- Password: <your-password>
- Login endpoint: https://127.0.0.1:8443/api/login
How to authenticate:
1. Write credentials to a file to avoid shell escaping issues
2. curl -k -X POST -c cookies.txt -H "Content-Type: application/json" -d @body.json https://127.0.0.1:8443/api/login
3. Use -b cookies.txt on subsequent requests
API endpoint reference — a table of every useful endpoint:
| Endpoint          | Method | Description                          |
|-------------------|--------|--------------------------------------|
| stat/sysinfo      | GET    | Controller version, system info      |
| stat/device       | GET    | All devices — full detail            |
| stat/sta          | GET    | All connected clients                |
| stat/health       | GET    | Dashboard health                     |
| rest/networkconf  | GET    | Network configs (VLANs, subnets)     |
| rest/wlanconf     | GET    | SSID configs                         |
| rest/firewallrule | GET    | Firewall rules                       |
| cmd/devmgr        | POST   | Device management (provision, etc.)  |The plan — what you actually want done:
Phase 1: Pull ALL endpoints and save the JSON responses
Phase 2: Analyze and produce recommendations
Phase 3: Apply changes (only after my explicit approval)
Important context — anything about your setup the LLM needs to know:
- This is a self-hosted controller (not UniFi OS), so API base path 
  is /api/s/default/ with no /proxy/network prefix
- Always use -k flag (self-signed cert)
- Port 8443 for management API, port 8080 is device inform only
The context file is doing the heavy lifting here. It turns a vague request into a structured task with clear boundaries. The LLM knows how to authenticate, what to look at, and what the rules of engagement are.
And if you're wondering how I built this context.md file, that's right, I did it with Opus 4.6.
Here's roughly how I kicked it off:
"I'd like your help optimizing my UniFi network. Context is in context.md."
That's it. The LLM read the context file, authenticated against the API, and started pulling data from every endpoint. Within a few minutes it had a complete picture of my network: hardware inventory, firmware versions, SSID configuration, radio settings, client distribution, RF environment, security posture, everything.
From there I asked it to produce two deliverables:
"Please write an analysis.md file with your findings and a recommendations.md file with what you suggest we do."
The analysis came back as a structured audit covering topology, RF environment, client health, and security gaps. The recommendations were organized into batches by risk level, each with specific API payloads ready to go.
This is where it gets fun. Rather than applying everything at once, we worked through changes in batches:
Batch 1: WiFi Radio Optimization (low risk, high impact)
- Fix channel conflicts
- Widen 5GHz channels
- Adjust transmit power
- Enable fast roaming and multicast enhancement
- Set minimum RSSI thresholds
Batch 2: Security Hardening (medium risk)
- Enable protected management frames
- Turn on DNS-over-HTTPS
- Enable DNS filtering
Batch 3: VLAN Segmentation (higher risk, critical impact)
- Create an IoT network on its own VLAN
- Create a dedicated IoT SSID
- Set up firewall rules to isolate IoT from trusted devices
Batch 4: Minor Optimizations (low risk)
- Enable IGMP snooping
- Disable NAT-PMP
For each batch, the LLM would:
- Show me exactly what it planned to change and why
- Wait for my explicit go-ahead
- Execute the API calls
- Verify the results
- Log everything to a change log file
The key prompt pattern was simple:
"Let's proceed with batch 1. Please create a log.md and track everything you do."
That log.md instruction was important. It meant every API call, payload, response, and verification step was recorded. If something went wrong, I could trace exactly what happened and roll it back.
Here's a taste of what the LLM found and fixed, generalized so it applies to most UniFi setups:
Co-channel interference on 2.4GHz. Two access points were both on channel 1. There are only three non-overlapping 2.4GHz channels (1, 6, 11), and with the controller set to "auto," it had made a bad choice. The LLM assigned manual channels to eliminate overlap and even suggested reducing to three APs since four APs on three channels means at least one conflict.
5GHz channels stuck at 40MHz width. The APs supported 80MHz channels (VHT80) but were configured for HT40. Switching to 80MHz nearly doubled available throughput. The LLM also assigned non-overlapping 80MHz blocks across all APs, including DFS channels where appropriate.
