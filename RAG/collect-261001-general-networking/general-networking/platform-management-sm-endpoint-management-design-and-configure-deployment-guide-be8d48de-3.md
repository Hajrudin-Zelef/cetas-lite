---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de-3
title: "platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de.md
source_anchor: ""
source_lines: [84, 96]
sha256: 0682f385eb63df695c9a62b67bc85010e9d79f7ef251eb84066fda59e2179a3d
---

# platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de

Adding one app to the block list, like in the above example, will allow all other System apps to show on the device(s). Without this, Android Device Owner mode may block all apps. For more information on this, see Controlling System Apps.
Installing Applications
Applications can be pushed to all Android Enterprise enabled devices in either Work Profile and Device Owner mode for publicly listed Google Play Store apps, or custom .apk Android apps. On Work Profile enrolled devices, a notification may appear when the app install command is sent.
Play Store Apps
For more information regarding deploying Android Store Apps please refer to this document.
Scoping
Once applications have been approved for an Android Enterprise organization they need to be scoped to devices to appear in the device app store. Approved applications (which have been scoped to devices via tags) will appear in both the Meraki Systems Manager App under "Managed Apps" as well as in the Play Store. Approved applications for the Work Play Store essentially create an allow list of applications a device can download and use. Once pushed to a device, these applications will silently install.
See this article for information on app configuration settings.
Apps that have been added into Dashboard but not approved may be listed in the Play Store, but will not be available to download until approved.
Custom (Enterprise) Applications
To upload custom .apk files to the Managed Play Store please follow this Google article.
Additionally, there is the option to distribute the .apk file directly through the Meraki Dashboard - for more information on this topic please reference this article.
See this article for information on app configuration settings.
