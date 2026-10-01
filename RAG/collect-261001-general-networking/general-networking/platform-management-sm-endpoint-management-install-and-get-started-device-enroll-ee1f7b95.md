---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-ee1f7b95
title: "platform-management-sm-endpoint-management-install-and-get-started-device-enroll-ee1f7b95"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-ee1f7b95.md
source_anchor: ""
source_lines: [1, 41]
sha256: 09d9515f7a19cd8a6faf75aa0b9a945bd7f96a3e2b111e54704627cc0587e9e5
---

# platform-management-sm-endpoint-management-install-and-get-started-device-enroll-ee1f7b95

Using the Systems Manager iOS App for Additional Functionality
日本語版はこちら
The Meraki SM iOS App provides additional features, and another method of enrolling iOS devices into a Systems Manager network. The app is compatible with iPhone, iPad, and iPod Touch, and requires iOS 9 or higher. While installing the iOS app is not required to enroll/manage an iOS device, it does provides for additional features not otherwise available. Installing and enabling the SM app currently supports the following additional features:
- GPS tracking: When running, the app can report a device's GPS location. Devices where permission is set to Always and Precise location tracking will be able to update their GPS coordinates as the device moves, even while the SM app is not running.
- Notifications: Messages can be sent to managed iOS devices with the app installed.
- Cellular data usage: While running, SM will track the device's data usage for the month.
- SSID: The last reported SSID connection is displayed on the Monitor > Devices page.
- LAN IP: The last reported LAN IP address of the device is displayed on client's details page as well as the Monitor > Devices page.
- Backpack: Distribute files to your end users.
- Suggested Network Enrollment: When iOS/iPadOS devices are connected to a SSID from a Meraki MR, they are shown a list of "Suggested networks" to make SM enrollment easier.
- Homescreen shortcuts: Tap and hold on the SM MDM app from the homescreen to invoke shortcuts for check-ins and apps list refreshes.
The SM app can either be downloaded/installed after a device is already enrolled in Systems Manager, or as the initial method of enrolling an iOS device. This Knowledge Base article will discuss both options:
The Android SM app provides similar functionality, but is mandatory for enrollment and installed as part of that setup process.
Note: All provided screen shots for Option 1 are as they appear on iOS device itself. Option 2 includes screen shots from both iOS device and Dashboard.
Enrolling into Systems Manager via the iOS App
1. Open iTunes App Store and search for "Meraki."
2. Download/install Meraki Systems Manager.
3. Input the Network ID found on Systems Manager > Manage > Add Devices page and select Enroll This Device.
- If no Network ID is displayed on Systems Manager > Manage > Add devices page (iOS tab), a valid Apple push certificate has not yet been uploaded into Dashboard. Instructions for doing so can be found on the Organization > Settings page.
- If the iOS device is connecting to internet through a Meraki access point, the SM App will scan for any Systems Manager networks within the same Dashboard organization as the Meraki access points. This 'Automatic Network Discovery' feature is enabled by default. To disable this feature, please consult this Knowledge Base article.
- You may also populate the Network ID field by selecting 'Use QR Code' and pointing the device's camera at the QR code found on the 'iOS' tab on the Systems Manager > Manage > 'Add devices' page. This QR code may be printed or otherwise distributed to end users instead of providing the 10-digit Network ID.
4. Install the Meraki SM management profile.
On iOS 12+ the enrollment redirect does not automatically happen due to a security change on iOS. Instead, the Meraki Management profile is downloaded into the Settings app and can be installed from there.
5. When prompted, select OK to allow for the SM app to report the device's location via GPS. If Don't Allow selected, Systems Manager will use an alternate method of approximating the location.
6. Confirm the device is enrolled in the desired Systems Manager network (name will match SM Network name in Dashboard when tapping on the Enrolled entry)
Deploying the iOS App to an already managed device
To add increased functionality, the Meraki SM iOS app can be installed on iOS devices already managed in a SM network. The SM app can either be downloaded manually from each device (follow steps 1-2 from Option 1 above), or deployed via SM like any other iOS app from the Systems Manager > Manage > Apps.
1. Navigate to Systems Manager > Manage > Apps page.
2. Select + Add new; then iOS app.
3. Search for 'Meraki' and add "Meraki Systems Manager" app.
4. Using Tags, configure the scope of target devices in the Targets section of the app details page.
5. Select Save Changes.
6. When a device within the scope of the SM app's targets checks-in with Systems Manager, the device user will be prompted to download/install the Meraki SM iOS app.
7. After downloading the SM app, the app must be opened to sync with the device's current SM network. This sync should be automatic upon launching the app, and the device user should not need to input a Network ID.
8. The Meraki SM app will sync with the SM network tied to the already installed Meraki management profile on the device.
9. The Meraki SM app is now installed and enabled on the iOS device.
Common Troubleshooting Tips
- Similar to any other iOS app, the Meraki SM iOS app must be running in order to report device information, including GPS location, SSID, and LAN IP address. In most cases, the SM app will not be running since iOS unloads apps not in use. Network administrators can prompt the device user to launch the app from the MDM commands section on the Client's details page. The message can be customized to encourage the device user to accept the notification prompt. Doing so will launch the app and allow for the device to update information to Dashboard.
- If deploying the SM app via Option 2, and the device user is not being prompted to install the app even after the device checks-in with Systems Manager, it may be because the app is already installed on the device. If so, an error message will be logged in the Activity Log at the bottom of the Client's details page.
- If the SM app is installed on a device (confirmed either on device itself or from Apps inventory section on Client's details page), but a Notification or GPS location command cannot be sent (notification to install SM iOS app appears instead), it is likely because the SM app has not been launched on the device, and thus has not synced with the device's existing Systems Manager network. Launching and syncing the app will remedy this issue.
The Meraki SM iOS app provides for an additional method of enrolling iOS devices in a Systems Manager network. Additionally, the SM app supports increased functionality and features not otherwise available to SM network administrators. For more comprehensive device monitoring, including more accurate location tracking, installing and enabling the Meraki SM app should be considered.
