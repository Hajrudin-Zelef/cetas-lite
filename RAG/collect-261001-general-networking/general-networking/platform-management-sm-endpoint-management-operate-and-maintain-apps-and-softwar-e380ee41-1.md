---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41-1
title: "platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41.md
source_anchor: ""
source_lines: [1, 112]
sha256: 31d4e9eee59c26a2dc9e2e7cd9a0aef6cbfc11a9fc84bff04dc68e25375c42bb
---

# platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41

How to Install Custom Apps on iOS and Android Devices
Overview
Systems Manager distributes enterprise iOS apps that are developed in-house when organization needs are not met by the public or private Apple app store release channel. Enterprise iOS apps are apps developed using the iOS Enterprise Developer Program, which allows for the development of internally created iOS apps. To deploy these in-house enterprise apps over-the-air to managed iOS clients, Systems Manager distributes a manifest file (.plist) to devices during app installation.
For Android, Systems Manager pushes applications silently to all Android devices in both BYOD and Device Owner mode for Google Play Store apps. You can also distribute a custom .apk file directly through the Meraki Dashboard.
This article explains how to install custom (enterprise) apps on iOS and Android devices using Systems Manager.
Prerequisites
Use only in-house developed applications with the iOS enterprise process. Access applications developed by any third party using the public Apple app store or the Apple Business Manager custom/private app process.
For iOS custom (enterprise) apps:
- Package the enterprise app using the over-the-air (OTA) method.
- Sign the enterprise app using a valid Enterprise distribution certificate.
For Android custom apps:
- 
    Android 15+ requires custom apps to be built with Android SDK version 24+. Refer to the Android developer documentation and/or consult the app developer for more information.
To silently install a custom .apk, upload the .apk to Google's managed Play for Work Store as a private (or public) app and push it as a Google Play Store application in Systems Manager. Refer to how to manage private Android apps in Google Play for more information.
Step-by-step instructions
Installing custom (enterprise) iOS apps
To begin, go to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app. Systems Manager allows for three methods of Enterprise iOS app deployment:
- Specify a link to an externally hosted manifest .plist file.
- Host the manifest .plist file within Systems Manager.
- Host the .ipa file within Systems Manager. Dashboard generates, hosts, and pushes a manifest .plist file.
For more information on distribution of custom enterprise applications, refer to Apple Distribute proprietary in-house apps to Apple devices documentation.
Specify a link to an externally hosted manifest .plist file
- 
    Go to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- 
    Within the Source section drop-down, choose Specify a manifest URL.
- 
    Provide a link to an externally hosted .plist file. Dashboard and targeted devices must be able to access this link.
- 
    Select a tag group to deploy the app to from the Scope field. Refer to this Knowledge Base article for information on creating and assigning Tags.
- 
    Select Save Changes and allow 1–2 minutes for changes to take effect.
Host the manifest .plist file within Systems Manager
- 
    Go to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- 
    Within the Source section drop-down, choose Upload a manifest plist.
- 
    Upload a manifest .plist file.
- 
    Select a tag group to deploy the app to from the Scope field. Refer to this Knowledge Base article for information on creating and assigning Tags.
- 
    Select Save Changes and allow 1–2 minutes for changes to take effect.
Host the .ipa file in Systems Manager
- 
    Go to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- 
    Within the Source section drop-down, choose Upload an IPA.
- 
    Upload a .ipa file.
- 
    Select a tag group to deploy the app to from the Scope field. Refer to this Knowledge Base article for information on creating and assigning Tags.
- 
    Select Save Changes and allow 1–2 minutes for changes to take effect.
Installing custom Android apps
You can deploy custom Android apps using two methods.
Silent deployment through the managed Play Store (recommended for Android 11+)
- 
    Upload your custom .apk files to the managed Play Store by following this Google article.
- 
    Add that app name/identifier to the Systems Manager Apps page as a Store app.
Refer to this article for information on app configuration settings.
Distribute the .apk directly through the Meraki dashboard
- 
    Go to Systems Manager > Manage > Apps.
- 
    Select Add new > Android > Custom app.
- 
    Fill in the fields as desired.
- 
    Either link to a URL where your .apk is hosted, or upload it directly to the Meraki Cloud.
Updating iOS enterprise apps
The update method depends on how the app was originally deployed.
For an externally hosted manifest .plist:
- 
    Replace the .plist and .ipa files on the remote server.
- 
    Use the select dropdown to select out-of-date entries.
- 
    Use the push dropdown to trigger an install or update command.
The .plist file must contain a link to an updated .ipa file. If the .plist or .ipa file changes location, update their links where required.
For a manifest .plist hosted within Systems Manager:
- 
    Replace the .ipa file on the remote server.
- 
    Use the select dropdown to select out-of-date entries.
- 
    Use the push dropdown to trigger an install or update command.
If the URL hosting the .ipa file changes, update the manifest .plist file and upload the new manifest .plist file to dashboard.
For an .ipa file hosted in Systems Manager:
- 
    Replace the .ipa file on dashboard.
- 
    Use the select dropdown to select out-of-date entries.
- 
    Use the push dropdown to trigger an install or update command.
Verification
When a device with the same tag next checks in with Systems Manager, the iOS device user receives a prompt to install the app. The user does not need to enter an Apple ID or password to download the application.
For Android apps distributed as a custom uploaded .apk, end users may need to accept an additional prompt (for example, "Do you want to install this application?"). Installing apps may also require devices to allow unknown sources.
Troubleshooting
For custom uploaded .apk apps, the ability to install apps may require devices to allow unknown sources. Enable this from an SM profile or manually on the device by enabling developer mode.
To avoid these constraints, the recommendation for modern Android (11+) is to use the Google Play Store private or public app method, as this is the supported method per Google. Using this method, the apps push as Google Play Store apps via SM and install silently.
iOS Custom (Enterprise) Apps
Systems Manager allows for the distribution of enterprise iOS apps that are developed in-house and the organization needs are not met by the public/private Apple app store release channel. Enterprise iOS apps are apps developed using the iOS Enterprise Developer Program. This program allows for development of internally created iOS apps. To deploy these in-house enterprise apps over-the-air to managed iOS clients, Systems Manager will distribute a manifest file (.plist) to devices during app installation.
- Only in-house developed applications should be used with this process. Applications developed by any third party should only be accessed using the public Apple app store or Apple Business Manager custom/private app process.
- The enterprise app should be packaged using the over-the-air (OTA) method.
- The enterprise app should be signed using a valid Enterprise distribution certificate.
Installing Custom (Enterprise) iOS Apps
Begin by navigating to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app. Systems Manager allows for three methods of Enterprise iOS app deployment:
- Specify a link to an externally hosted manifest .plist file.
- Host the manifest .plist file within Systems Manager.
- Host the .ipa file within Systems Manager. Dashboard will generate, host, and push a manifest .plist file.
