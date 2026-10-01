---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-best-practices-for-locking-down-admin-access-to-unifi-controller-4173e953-2
title: "blog-best-practices-for-locking-down-admin-access-to-unifi-controller-4173e953"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-best-practices-for-locking-down-admin-access-to-unifi-controller-4173e953.md
source_anchor: ""
source_lines: [120, 167]
sha256: c8f5d4ecb26ed9e1080b0cac32fda8061fa59ded90c72b949c55d3d62ca01c41
---

# blog-best-practices-for-locking-down-admin-access-to-unifi-controller-4173e953

If you manage multiple controllers/sites:
- Use a single UniFi Identity account
- Give each admin site-level permissions via Identity portal
- This avoids duplicate accounts and lets you revoke access across sites in one hit
Caution: Even with Identity, remember that granting "Administrator" role on any site gives global access. Use Site Administrator roles unless full global access is required.
14. Plan for Disaster Recovery
What happens if worst-case strikes?
- Store controller backups and config regularly
- Save firewall configs, VPN settings, SSL certs
- Rehearse restoration—time matters
- Have an alternate admin process (SSH via console, backup account)
- Critical: Maintain at least one local admin account with documented credentials for emergency access
Common Security Mistakes to Avoid
- Accidentally granting global access: Remember that Administrator role = all sites
- Sharing credentials: Never share accounts, always use individual logins
- No local backup: Always have a non-SSO admin for emergencies
- Not reviewing permissions: Regularly audit who has what access
- Exposing port 8443: Use VPN instead
- Ignoring local accounts: These can bypass SSO 2FA—secure them carefully
Community Tips from Admins
Reddit users and MSPs share what's worked for them:
"Put all the sites in your own cloud controller ... then any sites added, you have admin over automatically."
"We use HostiFi.com ... SSH is enforced via private keys and logins have 2FA capability."
"Don't open anything to the WAN. Just set up a VPN server and have everything available via LAN."
These comments echo the importance of central hosting, secure access, and isolation from the internet.
Related Resources
- How to Create a Global Administrator Account in UniFi Network Controller - Step-by-step guide with screenshots
- Setting up administrators in UniFiOS - Administrators, roles and permissions
- How to restore UniFi Controller - Restoring a controller from backup
Why This Matters
Hungry attackers, misconfigurations, or simple mistakes can all lead to compromised network access. With proper admin lockdown, you:
- Prevent unauthorized changes
- Maintain accountability
- Reduce risk of data exfiltration
- Ensure stability and uptime, even if credentials leak
- Stay compliant with privacy standards
UniFi is powerful, but it's only safe in experienced hands. Understanding the permission system—including the single-site-to-global Administrator behavior—is crucial for maintaining security.
Admin Role Successfully Changed
Final Thoughts
Locking down admin access isn't extra work, it's essential. With unique accounts, proper role assignment (understanding that Administrator = global access), strong passwords, 2FA, restricted access, and active monitoring, you build a secure foundation.
Key Takeaway: The most critical security control is understanding that granting "Administrator" role on any site gives that account access to everything. Use this permission sparingly and only for trusted personnel who truly need global access.
And if you're managing multiple clients or sites, skip self-hosting. Let us at UniHosted handle your controller hosting. We follow all these best practices automatically—including maintaining secure local admin backups—so you can focus on network performance, not servers or security patches.
Need help setting up proper admin access controls? Our team can guide you through creating the right permission structure for your specific needs.
Related guides
Keep reading
- Add a Global Administrator account in UniFi NetworkStep-by-step guide to adding an administrator to all sites in the UniFi Controller. Learn how UniFi admin permissions work and how to create a global admin account. Read guide
- Setting up administrators in UniFiOS Read guide
- How to restore UniFi Controller: A step by step guideStep-by-step guide to restore a UniFi Controller from backup. How to import settings, re-adopt devices, and recover after a server move. Read guide
