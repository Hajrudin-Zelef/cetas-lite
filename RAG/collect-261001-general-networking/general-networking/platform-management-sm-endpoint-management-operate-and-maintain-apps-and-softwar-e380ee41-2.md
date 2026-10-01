---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41-2
title: "platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41.md
source_anchor: ""
source_lines: [113, 146]
sha256: c1d575bacfeeaad2c5c4a73412c23eb3a4d0b46c0afd27835791c338bf9a13a3
---

# platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-e380ee41

For more information on distribution of custom enterprise applications, reference Apple's Distribute proprietary in-house apps to Apple devices documentation.
Specify a link to an externally hosted manifest .plist file
- Navigate to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- Within the Source section drop-down, choose "Specify a manifest URL".
- Provide a link to an externally hosted .plist file. This link must be accessible by dashboard and targeted devices.
- Select a tag group to deploy the app to from the "Scope" field (please consult this Knowledge Base article for information on creating and assigning Tags).
- Save Changes and allow 1 - 2 minutes for changes to take effect.
- When a device with the same tag next checks-in with Systems Manager, the iOS device user will be prompted to install the app. The user will not need to input an Apple ID or password to download the application.
- The app can be updated by first replacing the .plist and .ipa files on the remote server, and then use the select dropdown to select out of date entries then use the push dropdown to trigger an install or update command. The .plist file must contain a link to an updated .ipa file. If the .plist or .ipa file have changed locations, update their links where required.
Host the manifest .plist file within Systems Manager
- Navigate to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- Within the Source section drop-down, choose "Upload a manifest plist".
- Upload a manifest.plist file.
- Select a tag group to deploy the app to from the "Scope" field (please consult this Knowledge Base article for information on creating and assigning Tags).
- Save Changes and allow 1 - 2 minutes for changes to take effect.
- When a device with the same tag next checks-in with Systems Manager, the iOS device user will be prompted to install the app. The user will not need to input an Apple ID or password to download the application.
- The app can be updated by first replacing the .ipa file on the remote server, use the select dropdown to select out of date entries then use the push dropdown to trigger an install or update command. If the URL hosting the .ipa file has changed, update the manifest .plist file and upload the new manifest .plist file to dashboard.
Host the .ipa file in Systems Manager
- Navigate to Systems Manager > Apps > Add App > iOS > Custom (Enterprise) app.
- Within the Source section drop-down, choose "Upload an IPA".
- Upload a .ipa file.
- Select a tag group to deploy the app to from the "Scope" field (please consult this Knowledge Base article for information on creating and assigning Tags).
- Save Changes and allow 1 - 2 minutes for changes to take effect.
- When a device with the same tag next checks-in with Systems Manager, the iOS device user will be prompted to install the app. The user will not need to input an Apple ID or password to download the application.
- The app can be updated by first replacing the .ipa file on dashboard, then use the select dropdown to select out of date entries then use the push dropdown to trigger an install or update command.
Android Custom Apps
Applications can be pushed silently to all Android devices in both BYOD and Device Owner mode for Google Play Store apps.
To install a custom .apk silently, upload the .apk to Google's managed Play for Work Store as a private (or public) app and push it as a Google Play Store application in Systems Manager. For more information view how to manage private Android apps in Google Play. 
Android 15+ requires custom apps to be built with Android SDK version 24+. See Android developer documentation and/or consult the app developer for more information.
Installing Custom Android Apps
To upload custom .apk files to the managed Play Store for silent deployment please follow this Google article. Afterwards your account should be able to add that app name/identifier to the Systems Manager Apps page as a Store app.
Additionally, there is the option to distribute the .apk file directly through the Meraki Dashboard. In Systems manager > Manage > Apps, click Add new > Android > Custom app. Fill in the fields as desired and either link to a URL where your .apk is hosted, or upload it directly to the Meraki Cloud.
See this article for information on app configuration settings.
With custom uploaded .apk apps, end users may need to accept an additional prompt (e.g. "Do you want to install this application?"). The ability to install apps may also require devices to allow unknown sources, which can be enabled from a SM profile or manually on device by enabling developer mode. To avoid these constraints, the recommendation for modern Android (11+) is to use the Google Play Store private or public app method, as this is the supported method per Google. Using this method the apps can be pushed as Google Play Store apps via SM and install silently.
