---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsense-vs-pfsense-for-homelabs-2026-complete-comparison-diymediaserver-3
title: "SSH into firewall, check what driver your NIC is using"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google", "Intel"]
dates: ["2026-01-31"]
keywords: ["advisory", "cost", "full-duplex", "incident", "intel", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/opnsense-vs-pfsense-for-homelabs-2026-complete-comparison-diymediaserver.md
source_anchor: ""
source_lines: [330, 480]
sha256: d8b787eae72afc8617349826978a9c3da22147733780a7b3f7d1a1547f514165
---

# SSH into firewall, check what driver your NIC is using

## Common Mistakes That Will Bite You

### Don’t Virtualize on Hardware You’re Using for Other Things

Your firewall VM needs dedicated hardware or a hypervisor that’s always on. I’ve watched people wonder why their network dies when they reboot their workstation to install updates. Because your firewall is on it. Obviously.

Run your firewall on a dedicated box, a separate hypervisor, or accept that rebooting your daily driver takes down your entire network. There’s no middle ground here.

### Don’t Skip Backups Before Updates

Both platforms make this easy:

**OPNsense:** System > Configuration > Backups > Download configuration

**pfSense:** Diagnostics > Backup & Restore > Download configuration as XML

Save it locally with the date in the filename: `firewall-backup-2026-01-31.xml`

Do this before updates. Do this before changing anything important. Do this monthly even if you’re not changing anything. Store it somewhere that’s not the firewall. Your NAS, your workstation, cloud storage, doesn’t matter. Anywhere but the firewall itself.

When (not if) you need to restore, you’ll thank past-you for being paranoid.

### Don’t Enable Every IDS Rule

More rules ≠ more security. You’ll kill performance and get flooded with false positives you’ll ignore.

Start with recommended rulesets. Monitor for a week. Add more only if you need them. I ran with three rulesets for 18 months before adding a fourth. You don’t need 47 different threat feeds.

### Don’t Use Realtek NICs If You Can Avoid It

They work. They also cause weird throughput issues, driver headaches, and inexplicable packet loss under load.

Intel NICs cost $20 more used on eBay. Buy Intel. The i350-T2 and i350-T4 are solid choices. Your future self will appreciate it when you’re not troubleshooting phantom network issues on a deadline.

### Don’t Trust Default Settings for Production

Both platforms ship with sensible defaults for home use. But “sensible defaults” means:

- No firewall rules blocking RFC1918 (private IP addresses) traffic on WAN (fine for home, terrible for dual-WAN or VPS)
- DNS resolver allowing queries from all interfaces (convenient but insecure)

Review the defaults. Adjust for your environment. Lock down access to the web UI. Enable stricter firewall rules. Don’t assume “default” means “secure.”

## Minimum Viable Firewall Setup

Stop overthinking the initial config. Day one, you need:

1. **WAN interface configured** - DHCP from ISP or static IP, whichever your ISP provides
2. **LAN interface with DHCP enabled** - Both do this automatically during install
3. **Default allow rule on LAN** - Both create this automatically
4. **DNS set to upstream resolvers** - 1.1.1.1 and 8.8.8.8, or your preference

That’s it. Everything else is optional.

VLANs? Add them when you need device isolation. VPNs? Add them when you need remote access. IDS? Add it when you want visibility into traffic. Custom dashboards? Add them when the defaults feel limiting.

Start simple. Add complexity only when you have a specific need. Your firewall’s job is routing packets and blocking threats. It does not need to look impressive in screenshots.

## Troubleshooting Common OPNsense and pfSense Issues

### Internet Works, but Throughput Is Terrible

**What you see:** Slow speeds despite fast connection

**The fix:**
Disable hardware offloading first. This is the most common culprit.

**OPNsense:** Interfaces > Settings > Disable “Hardware CRC”, “Hardware TSO”, “Hardware LRO”
**pfSense:** System > Advanced > Networking > Disable all hardware checksum offloading

Reboot. Test again.

If that doesn’t fix it, check IDS rulesets. Too many active rules kills performance. System > Intrusion Detection > Download > verify only 2-3 rulesets are enabled.

Verify your NIC drivers are loaded correctly:

```
# SSH into firewall, check what driver your NIC is using
ifconfig -a | grep -A 4 "em0"
```
If you see “re0” (Realtek), you found your problem. Intel NICs use “em”, “igb”, or “ix” drivers. Realtek uses “re”. Replace the NIC.

If you’re seeing 100 Mbps on gigabit, the NIC negotiated wrong. Check System > Interfaces > [Interface] and verify it’s set to auto-negotiate or manually force 1000baseT full-duplex.

### VPN Is Slow

**What you see:** WireGuard or OpenVPN performing poorly

**The fix:**
Verify WireGuard is using kernel implementation (not userspace).
Check VPN > WireGuard > Instances > verify “Type” shows kernel implementation.

Disable unnecessary logging (verbose logging kills performance). VPN > WireGuard > Advanced > set log level to “error” only.

Test without IDS. Suricata inspecting VPN traffic = very slow. Services > Intrusion Detection > disable temporarily, test VPN speed.

WireGuard should be fast. If it’s not, your config or hardware is wrong.

### Upgrade Broke Something

**What you see:** Features missing or broken after update

**The fix:**
Restore config backup (you made one, right?).
System > Configuration > Backups > restore previous config.

Check package compatibility in changelogs. Before updating, read the release notes. They list package compatibility issues. Review deprecated features list. Sometimes features get removed. Check migration guides.

Test packages individually after core upgrade. Update core first, reboot, then update packages one at a time.

Restore your backup. Try again. Five minutes reading release notes can save an hour troubleshooting. (Yes, I have learned this the hard way.)

## Security Differences That Matter

Both platforms are secure by default. The differences are operational, not architectural.

### OPNsense Security Advantages

**Two-factor authentication built-in.** No plugin needed. System > Access > Users > [Select User] > Generate new secret. Works with Google Authenticator, Authy, or any TOTP app.

**IDS integrated and easier to configure.** Services > Intrusion Detection > Download tab. Select rulesets, enable IDS, done. No separate package installation or config files to manage.

**Faster security patches.** No CE/Plus delay. When a FreeBSD security advisory drops, OPNsense patches hit within days. pfSense CE users wait for Plus to get patched first.

**More frequent security updates.** Bi-yearly major releases plus security patches as needed. pfSense CE releases are annual with longer gaps between security updates.

### pfSense Security Advantages

**More mature IDS rulesets.** Snort has been around longer than Suricata. More documentation, more tuned rules for specific scenarios. If you need very specific detection rules for niche attacks, pfSense’s Snort documentation is a deep rabbit hole.

**More third-party security plugins.** ntopng integration is tighter. More options for anomaly detection and traffic analysis. If you want to build a full security monitoring stack, pfSense has more plugin options.

**More documented attack mitigation examples.** Decade of forum posts, blog tutorials, and Stack Overflow answers. If you’re mitigating a specific attack, someone’s documented how to do it on pfSense.

### What Matters

Neither platform has had a major security incident in recent years. Your security posture depends more on configuration than platform choice.

Common security mistakes I see:

- Web UI exposed to WAN (don’t do this)
- Default admin passwords (change them immediately)
- No firewall rules blocking RFC1918 on WAN (matters for dual-WAN setups)
- Permissive outbound rules on LAN (most people allow all, should be more restrictive)

Fix these regardless of platform. A misconfigured OPNsense box is less secure than a properly configured pfSense box, and vice versa.

## What Surprised Me After Switching

**OPNsense updates are faster.** pfSense updates took 10-15 minutes and always required a reboot. OPNsense updates finish in 2-3 minutes. Most don’t need a reboot. The few that do reboot in under 60 seconds.

