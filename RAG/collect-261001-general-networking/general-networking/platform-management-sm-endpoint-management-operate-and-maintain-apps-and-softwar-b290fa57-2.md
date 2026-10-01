---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57-2
title: "platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57.md
source_anchor: ""
source_lines: [41, 72]
sha256: f686209933ab283cf8b73e644e7c119b3e9350250d384f417f1ad772f8c987b5
---

# platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57

  - Windows: shutdown /r
  - OS X: shutdown -r now
- Visible in SSP: Specifies if the app can be added or removed by end users from the Self Service Portal (SSP).
- Targets: Specifies a scope of devices that will have this application installed. Please refer to our documentation for more information on scoping by device tag.
4. Click Save Changes. Unless auto-install is disabled, this will push out the application to all devices within the specified scope.
Example Configuration
The following example configuration will quietly install FileZilla from a Dropbox URL on all managed Windows clients with the "bginstall" tag. These clients will also reboot after installation has completed.
Monitoring Installation
The status of an installation can be checked via the Monitor > Devices page, by clicking on a particular client and navigating to the Event log or Activity log section.
Removing Applications from Devices for Windows
Certain .msi and .app programs can be remotely uninstalled via the Systems Manager > Monitor > Software page. To remove a specific application, search for the application under Software inventory and click uninstall.
Note: That software inventory "uninstall" button for windows agent installed apps only works if there is an uninstaller file for the application in "c:\windows\program files\". If this is not the case, create your own .msi or .exe that does the uninstall, upload that as a custom app to Dashboard and scope it for the specific windows device.
Note: Auto-uninstall for an app pushed from the dashboard on windows devices doesn't automatically remove an app when that app is removed from the device's scope.
Removing Applications from Devices for MacOS
At this time, the Systems Manager agent does not support uninstallation of custom apps on MacOS.
Troubleshooting
If the software install command is issued and the application does not appear on the client device, check the Dashboard-side logging to see if the app install command errored out. If you see a 'Success' entry, this means the command successfully went through, but may not have completely executed on the client device. At this point, you should reference the agent logs on the device itself, and provide these to Meraki Support for troubleshooting assistance.
App is installed on device but reports as Missing or Not Installed
This can happen for several reasons, but it is generally a mismatch between the actual installed app name and the manually configured SM > Apps name. If the App name on SM > Apps is different than the actual name of the installed app reported to Systems Manager from the device the app will report as "Not Installed". To troubleshoot this, look at the actual name of the app according to the device. Click on the device and find the app in the installed App List:
The name, "Google Chrome", reported from the device must be the same name we give for the app on SM > Apps:
Other related issues:
- App must be installed from SM > App custom app as "Success" at least once.
- Multiple of the same App Name on the SM > Apps page is not supported. If multiple of the same app name are configured on SM > Apps, they may both report as "Not Installed".
- If the app is never reported as an "Installed" on the device, it will always show as "Not Installed" and missing this app.
    
  - For macOS, the app is reported as "installed" if the device sees it installed locally in:
        
    - About This Mac > System Report... > Application
  - The SM Agent collects the above information with the following terminal command:
        
    - system_profiler -xml SPApplicationsDataType
- For macOS, the app is reported as "installed" if the device sees it installed locally in:
