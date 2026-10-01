---
id: collect-261001-general-networking/general-networking/stbanmc-communityblocklists-0d7c4252
title: "stbanmc-communityblocklists-0d7c4252"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["attribution", "cost", "latency", "license", "open source"]
source: docs/RAG/collect-261001-general-networking/stbanmc-communityblocklists-0d7c4252.md
source_anchor: ""
source_lines: [1, 125]
sha256: 89127ddb8c72d47130b131719f4d895b4725f0013e0acd0b1d46e8a44e6c312f
---

# stbanmc-communityblocklists-0d7c4252

Community-driven URL/domain blocklists organized by category. Compatible with pfBlockerNG, Pi-hole, AdGuard Home, OPNsense, Unbound, dnsmasq, SquidGuard, and any DNS-based filtering solution.
4,000,000+ domains across 74 categories, updated weekly from 5 open sources. Programmatic catalogue: manifest.json.
Add as Adlist (hosts format):
https://raw.githubusercontent.com/StbanMc/CommunityBlocklists/main/exports/hosts/<category>.txt
Add as DNS blocklist:
https://raw.githubusercontent.com/StbanMc/CommunityBlocklists/main/exports/adguard/<category>.txt
- Firewall → pfBlockerNG → DNSBL → DNSBL Groups → + Add
- Source URL:
https://raw.githubusercontent.com/StbanMc/CommunityBlocklists/main/exports/domains/<category>.txt
- Format: Auto
- State: ON
- Frequency: Weekly (matches our update cycle)
- Save → go to pfBlockerNG → Update → Reload → Force
Tip: subscribe to specific categories instead of combined/all-domains.txt. Easier to revert if a single category blocks something legitimate.
https://raw.githubusercontent.com/StbanMc/CommunityBlocklists/main/exports/unbound/<category>.conf
https://raw.githubusercontent.com/StbanMc/CommunityBlocklists/main/exports/dnsmasq/<category>.conf
exports/combined/all-domains.txt   # Plain domains
exports/combined/all-hosts.txt     # Hosts format
exports/combined/all-adguard.txt   # AdGuard format
SquidGuard expects a single .tar.gz containing a top-level directory with
<category>/domains (and optional <category>/urls) entries. A
pre-packaged, reproducible tarball is published on every weekly run to
GitHub Pages as a direct download (no redirects — important for SquidGuard).
- Services → SquidGuard Proxy Filter → Blacklist tab
- Blacklist URL:
https://stbanmc.github.io/CommunityBlocklists/squidguard.tar.gz
- Click Download and wait for the percentage to reach 100% (the tarball is ~30 MB — first download takes 1-2 minutes on a typical pfSense box).
- General settings tab → make sure Enable is on → Apply.
- Target categories tab → the categories you want to act on will appear
automatically as [squidguard] <category> . Set them to deny,
allow, or whitelist as needed.
- Common ACL (or Groups ACL) tab → wire categories into your access rules → Apply.
- SquidGuard Proxy Filter → General settings → Apply at the bottom to recompile the BerkeleyDB indexes (this takes 2-5 minutes for the full set — log shows progress).
Integrity check (optional but recommended): download
squidguard.tar.gz.sha256 from the same Pages URL and verify before
applying in production.
Why Pages and not the GitHub Releases URL? GitHub Releases serve
assets through a signed-redirect chain that the pfSense SquidGuard
downloader does not follow correctly (it produces Bad content. Terminate.).
The Pages URL is a single HTTP 200 with no redirects. The Releases asset is
still published in parallel for manual download / versioned audit.
GitHub raw is fine for most users. If you are far from US edges or need lower latency, jsDelivr automatically mirrors this repo:
https://cdn.jsdelivr.net/gh/StbanMc/CommunityBlocklists@main/exports/domains/<category>.txt
https://cdn.jsdelivr.net/gh/StbanMc/CommunityBlocklists@main/exports/hosts/<category>.txt
https://cdn.jsdelivr.net/gh/StbanMc/CommunityBlocklists@main/exports/adguard/<category>.txt
To pin to a specific release tag instead of @main (recommended for production firewalls), replace @main with the tag, e.g. @v2026.05.10.
The full list of categories with sizes, formats and source attribution lives at manifest.json — useful for tools, dashboards or wrappers that want to discover what is available without scraping the README.
| Format | Directory | Use case | 
|---|---|---|
| Plain domains | exports/domains/ | pfBlockerNG, generic | 
| Hosts file | exports/hosts/ | Pi-hole, /etc/hosts | 
| AdGuard | exports/adguard/ | AdGuard Home | 
| dnsmasq | exports/dnsmasq/ | dnsmasq, OPNsense | 
| Unbound | exports/unbound/ | Unbound DNS resolver | 
| Combined | exports/combined/ | All categories merged | 
| SquidGuard tarball | squidguard.tar.gz on Pages + Release | SquidGuard (pfSense, OPNsense) | 
| SquidGuard source tree | ListFilter-Squid/BL/ | Local build / inspection | 
| Category | File name | Description | 
|---|---|---|
| Phishing/Scam | hacking | Phishing, scam, hacking tools | 
| Malware/Spyware | spyware | Malware, spyware, tracking, ransomware, cryptojacking | 
| Fraud | costtraps | Fraud, cost traps | 
| Aggressive | aggressive | Aggressive/hostile content | 
| Violence | violence | Violence, dangerous material | 
| Category | File name | Description | 
|---|---|---|
| Advertising | adv | Ads, trackers | 
| VPN/Proxy | anonvpn | Anonymous VPN, proxies, DoH | 
| Redirectors | redirector | URL redirectors | 
| URL shorteners | urlshortener | URL shortening services | 
| Remote control | remotecontrol | Remote access tools | 
| Dynamic DNS | dynamic | Dynamic DNS services | 
| Category | File name | Description | 
|---|---|---|
| Pornography | porn | Pornography (1.1M+ domains) | 
| Lingerie | sex-lingerie | Lingerie | 
| Dating | dating | Dating sites | 
| Category | File name | Description | 
|---|---|---|
| Gambling | gamble | Gambling, betting (238K+ domains) | 
| Online games | hobby-games-online | Online games | 
| Movies | movies | Movie streaming | 
| Music | music | Music streaming | 
| Social networks | socialnet | Social networks | 
| Chat | chat | Chat platforms | 
| Webmail | webmail | Webmail services | 
| Category | File name | Description | 
|---|---|---|
| Shopping | shopping | Online shopping | 
| Banking | finance-banking | Banking | 
| Trading/Crypto | finance-trading | Trading, cryptocurrency | 
| Category | File name | Description | 
|---|---|---|
| News | news | News, press | 
| Warez/Piracy | warez | Piracy, warez | 
| Drugs | tracker-drugs | Drug-related | 
| Job search | jobsearch | Job search sites | 
See all 74 categories in ListFilter-Squid/BL/.
| Source | Updated | What it provides | 
|---|---|---|
| UT1 Blacklists (Univ. Toulouse) | Daily | 40+ categorized lists | 
| BlockListProject | Regular | ads, malware, phishing, gambling, porn, drugs, piracy, scam, tracking | 
| Hagezi DNS Blocklists | Regular | gambling, DoH | 
| StevenBlack/hosts | Regular | unified ads/malware | 
| Phishing Army | Daily | phishing domains | 
GitHub Actions updates the lists every Monday automatically:
- Downloads latest data from all upstream sources
- Merges with existing lists (additive, no domains lost)
- Deduplicates and sorts
- Exports to all formats
- Auto-commits
Trigger manually from the Actions tab anytime.
bash scripts/update_lists.sh      # Update from sources
bash scripts/export_formats.sh    # Export to all formats
See CONTRIBUTING.md for details.
- Add domains to categories via Pull Request
- Report false positives via Issues
- Suggest new categories or data sources
- Improve scripts or add export formats
- 4,000,000+ domains
- 74 categories
- 6 export formats
- 5 upstream sources
- Weekly automated updates
Open source. See LICENSE.
