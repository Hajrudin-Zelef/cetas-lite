---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-28692158912279-adding-admins-in-unifi-f0aec64f
title: "hc-en-us-articles-28692158912279-adding-admins-in-unifi-f0aec64f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-28692158912279-adding-admins-in-unifi-f0aec64f.md
source_anchor: ""
source_lines: [1, 29]
sha256: 454166df5a9699038e22c20c28fb9b8f6f882cfd65bb49b6609fced75ca87aae
---

# hc-en-us-articles-28692158912279-adding-admins-in-unifi-f0aec64f

Adding Admins in UniFi
UniFi uses a role-based access control (RBAC) model to manage administrator access. How Admins are added and managed depends on whether you are using UniFi Fabrics for centralized, multi-site administration or managing individual sites independently.
To learn more about the difference between Admins and Users, click here.
Admin Management with UniFi Fabrics
When using UniFi Fabrics, admin access is managed centrally at the Fabric level, allowing roles to be defined once and applied consistently across multiple sites.
Fabric-based admin management supports:
- Centralized role definitions with granular permissions
- Assigning multiple roles to a single admin
- Mapping admin roles to Identity Provider (IdP) user groups for automated onboarding and offboarding
When an IdP is integrated, admin access can be managed using group membership instead of manual assignments, keeping permissions aligned with organizational changes.
For an overview of Fabrics, check out this UniFi Academy Topic.
To learn how roles and permissions work in Fabrics, see UniFi Fabrics: Managing People, Roles, and Permissions.
Adding Admins in Site Manager (Non-Fabric)
Note: Admins must already have a UI Account, which they can create at account.ui.com.
- Log in to Site Manager.
- Select People from the left navigation bar.
- Click the New Admin button.
- Select one or multiple sites to add the Admin to.
- Assign permissions for each application. You can:
  - Apply the same permissions across all sites.
  - Check Site Specific to configure different permissions for each site.
- Confirm and save your settings.
Local-Only Management
If remote management is disabled, you can also add local admins by following these steps:
- In your Network Application, go to Admins in the left navigation bar.
- Select the “+” symbol in the top right corner.
- Add your admin information.
  - Optionally choose whether to allow Remote Management.
- Click Invite.
