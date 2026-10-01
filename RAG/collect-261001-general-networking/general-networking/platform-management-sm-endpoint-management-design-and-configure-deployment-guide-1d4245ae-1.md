---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae-1
title: "platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae.md
source_anchor: ""
source_lines: [1, 32]
sha256: e0c210a35ba4e81f2a0cd64c00e4abb7492f10b562c1ff4809190854304a2709
---

# platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae

Migrating from another Management Solution to Meraki Systems Manager
Migrating all of your devices from another MDM provider into Meraki Systems Manager can feel like an overwhelming task at first, but do not fear! Switching MDM providers does not have to be daunting, overwhelming, or painful. There are many reasons for switching MDM providers: whether your current MDM provider does not provide the support you expect or you simply desire a simple but rich feature set for managing devices. This guide will assist you in transitioning from another MDM provider into Meraki Systems Manager as seamlessly as possible, with the least disruption for end users.
This article covers Apple device migration (iOS and macOS) as there are several tools to leverage to smooth out the process. Windows, Android, and Chrome OS migration steps will be added at a future date.
Preparing to Make the Switch
There are a few things that should be considered before making the switch from another MDM provider to Meraki Systems Manager. Your organization's timeline requirements, communication to end users, and using first party tools like Apple's Device Enrollment Program, can help outline an action plan for the transition.
Timeline
Timeline requirements by your organization should be determined before jumping in. Planning to migrate during a time where end users can be without their devices can help make a migration easier. If you do not have access to the devices, end users may be required to perform some actions depending on the steps you are using to migrate.
- With Apple's Device Enrollment Program (DEP): The switch from another MDM provider into Meraki Systems Manager can be performed at a faster pace when leveraging DEP. DEP can be used to automatically enroll devices into Meraki Systems Manager with much less end user action required.
- Without Apple's Device Enrollment Program (DEP): Consider implementing DEP moving forward. DEP gives your devices automatic enrollment & supervision into Meraki Systems Manager. Furthermore, as of iOS 11, you can now add non-DEP iOS devices into your existing DEP account, regardless of the device's purchase source. There is no better time to add non-DEP devices to DEP than right now during this migration! Without using DEP, plan for the devices the be away from end users during the transition or for end users to perform some basic steps on the devices.
Communication
If you plan to have end users perform any actions (e.g. go through device's initial setup assistant) it is good practice to detail the exact steps your end users will be expected to perform and communicate it to them prior to the migration. Review the steps below in this guide for the migration, test the deployment, and outline what is needed to your users.
- With Apple's DEP: DEP device setup requires devices to be factory erased to go through the automatic re-enrollment. Devices can stay with end users during this migration, but be sure to notify your end users that the device will be factory erased and to backup any sensitive data.
- Without Apple's DEP: Devices can be manually un-enrolled and re-enrolled without a factory reset being required. This requires some manual steps for end users (covered throughout this guide below). Another option would be to collect all devices and mass enroll them through Apple Configurator, Sentry splash if you are using Meraki MR wireless access points, or manual enrollment.
Backing Up Data from a Previous MDM Provider
Be sure to review the data which your previous MDM provider has stored and export anything that you would like to use within Meraki Systems Manager. While you can always rebuild this data in Meraki Systems Manager later, an export/import action is always a faster experience than rebuilding everything. Consider exporting the following from the previous MDM provider:
- Configuration profiles (.mobileconfig)
- Security certificates
- Device inventory
- Packages, apps, and scripts (.pkg, .dmg, .app, .plist, .ipa files, etc)
Note: Some MDM providers may not allow you to export some of this data. No problem! We can rebuild/reconfigure these in Meraki Systems Manager later.
This is a good opportunity to confirm that end users have backed up any data they do not wish to lose from their managed devices. Photos, app data, and any non-managed device configurations may be lost during the migration process. Have users backup their data so they can restore it after the migration, if applicable.
What about the Apple Push Notification service (APNs) certificate?
The Apple Push Notification service certificate is required for Apple devices to communicate with MDMs. This means an APNs certificate is likely already created for the previous MDM provider. While it's technically possible to use the current APNs certificate and renew it with Meraki's .csr (found in Organization > MDM), it is recommended to generate a totally new APNs certificate to be used exclusively with Meraki Systems Manager. You can create as many APNs certificates as you need (two different APNs certificates in this scenario if you are transitioning from one MDM provider into Systems Manager) in the Apple Push Certificate Portal. We will set up the new Meraki Organization's APNs during the Meraki Systems Manager during the migration steps (below).
What about Apple's Volume Purchase Program (VPP) & licensed applications?
VPP can be used to manage app licenses with the Mac App Store and iOS App Store for your users & devices. VPP allows MDMs to grant/revoke licenses to/from an end user's Apple ID or directly to the device. Login to your previous MDM provider and remove any of your old VPP token(s). We will set up the VPP tokens into Meraki Systems Manager during the migration steps (below) so Systems Manager can control the VPP grant/revoke licensing actions.
Migrating accounts/certificates into Meraki Systems Manager
Apple Push Notification service (APNs) - required
Create a new APNs certificate on Apple Push Certificate Portal and add it to your Systems Manager Dashboard in Organization > MDM. Detailed steps on how to do this can be found here. Be sure to create a new APNs certificate! Do not renew your existing certificate with your previous MDM provider.
Apple's Volume Purchase Program (VPP) - optional
Login to your previous MDM provider and remove/delete any of your previous VPP token(s). Download the VPP token(s) again from your VPP account(s) on Apple VPP Portal (or Apple School Manager Portal). Now login to your Meraki Dashboard and go to Organization > MDM and add the VPP tokens so Systems Manager can control VPP app licensing moving forward. Detailed steps on setting up VPP in Systems Manager can be found here.
Apple's Device Enrollment Program (DEP) - optional
Login to the Apple Business Manager (or Apple School Manager Portal) with your primary Apple ID administrator account. Create a new "MDM Server" and move all the devices from the previous MDM Server (if you had one with your previous MDM provider) into the newly created Meraki MDM Server. Then, link your DEP (or Apple School Manager) account to Meraki Systems Manager by downloading the token from the DEP Portal (or Apple School Manager Portal) and uploading it to your Meraki Dashboard in Organization > MDM. Details steps on linking up a DEP account to Meraki Systems Manager can be found here. Once DEP is connected Meraki Systems Manager, you should be able to see the devices syncing into Systems Manager > Manage > DEP. They are not enrolled at this point, but this is the first step to get them into Meraki Systems Manager. Assign DEP settings to your devices so they all appear as "Assigned" (below).
