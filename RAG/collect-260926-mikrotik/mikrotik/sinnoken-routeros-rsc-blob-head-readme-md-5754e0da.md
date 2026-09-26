---
id: collect-260926-mikrotik/mikrotik/sinnoken-routeros-rsc-blob-head-readme-md-5754e0da
title: "sinnoken-routeros-rsc-blob-head-readme-md-5754e0da"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/sinnoken-routeros-rsc-blob-head-readme-md-5754e0da.md
source_anchor: ""
source_lines: [1, 33]
sha256: f0a04bedf62861a5fee84dc6cd99d0a9fe7c8421009c79c5b9b87df860e49789
---

# sinnoken-routeros-rsc-blob-head-readme-md-5754e0da

This Python utility is designed to fetch threat intelligence data from the IPsum project and generate optimized firewall address list configuration files (.rsc) specifically for MikroTik RouterOS.
- Multi-Source Fetching: Downloads IP address lists from multiple configured URLs.
- Validation & De-duplication: Automatically validates IP formats and removes duplicates for a cleaner list.
- RouterOS Optimization: Generates ready-to-use MikroTik firewall address-list commands.
- High Performance: Utilizes multi-threading to significantly speed up the download and processing workflow.
This script handles the downloading of the blacklist and refreshes the local address list.
/system script add name="downloadBlackList" owner="HybridNetworks" source={
    /tool fetch url="https://github.com/sinnoken/routeros-rsc/raw/refs/heads/main/rsc/STAMPARM-IPSUM-LEVEL-3.rsc" mode=https;
    :delay 5;
    /ip firewall address-list remove [find where list="STAMPARM-IPSUM-LEVEL-3"];
    :delay 5;
    /import file-name=STAMPARM-IPSUM-LEVEL-3.rsc;
    :delay 5;
    /file remove STAMPARM-IPSUM-LEVEL-3.rsc;
}
Set a schedule to automatically update the blacklist every 3 days.
/system scheduler add comment="BlackList" interval=3d \
    name="BlackListUpdate" on-event=downloadBlackList \
    start-date=jan/01/1970 start-time=10:10:10
For optimal performance and lower CPU usage, use IP Firewall RAW to drop traffic.
/ip firewall raw
add action=drop chain=prerouting comment="STAMPARM-IPSUM-LEVEL-3" \
    src-address-list=STAMPARM-IPSUM-LEVEL-3
- Environment: Ensure you have Python 3.x and the requests library installed.
- Clone: Download or clone this repository to your local machine.
- Execute: Run the processor script:
python script_name.py
- Output: The generated .rsc configuration files will be saved in the./rsc/ directory.
- Python 3.x
- requests library
- Compatibility: Verify that your RouterOS version supports the generated commands (compatible with most v6 and v7 builds).
- Customization: You can modify the urls list within the script to target different IPsum threat levels (e.g., level 1 to level 8).
Contributions are welcome! Please feel free to submit Issue reports, Feature requests, or Pull Requests to improve the code.
