---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c-1
title: "wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c.md
source_anchor: ""
source_lines: [1, 52]
sha256: ed2765a1ff586c3e02b4c035dc0fa3426bdbed363bc4a3e671de905c2c368128
---

# wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c

Conducting Site Surveys with MR Access Points
Click 日本語 for Japanese
Learn more with this free online training course on the Meraki Learning Hub:
To best understand the RF environment both pre and post-deployment of wireless infrastructure, it is strongly encouraged to perform site surveys and make changes based on the result. This article outlines some tools and best practices for performing site surveys with an MR wireless deployment.
Site Survey Mode
Cisco Meraki Access Points can be configured to broadcast a dedicated SSID for Site Surveys, without the access point requiring an active Internet connection to the Meraki Dashboard. Clients can then connect to this SSID in order to conduct passive or active site surveys. A MR in site survey mode will be denoted by a blinking green status LED on the unit.
Note: Enabling Site Survey Mode requires access to the local status page. Though the local status page is enabled by default, administrators do have the option to disable the local status page on their devices. Ensure that access is enabled in the dashboard under Network-wide > General > Device configuration. For more information, refer to Using the Cisco Meraki Device Local Status Page.
MR 32.1.X Safe Config Override for Site Surveying
Starting with Firmware MR 32.1.X, Safe Configuration validation is bypassed on the local status page, allowing network admins to immediately convert APs to site survey mode without requiring the APs to be connected to the dashboard for up to 2 hours.
Before MR 32.1.X firmware, to utilize the 5 or 6 GHz radios on an AP for site surveys, the flex radio on your 6 GHz-capable APs must be set to 6GHz, and your AP(s) must be connected to dashboard before enabling survey mode for up to 2 hours.
This would allow the configuration of your APs to be marked as "safe" and let your APs be used in a site survey mode without a persistent connection to the Meraki dashboard.
Configuring a Survey SSID
The following instructions outline how to enable site survey mode on an MR access point:
- Ensure that the AP has come online and checked into the Cisco Meraki Dashboard at least once. It should currently be green in Dashboard.
- Access the local status page of the AP by associating with any SSID broadcast by the device, and opening ap.meraki.com in a web browser.
Note: If the AP does not have a wired connection to the network, it will broadcast an SSID similar to "meraki-scanning" for access to the local status page:
- On the local status page, navigate to the Configure tab.
- Set Survey mode to enable to turn on site survey mode.
- 
    Optional: After the site survey SSID has been enabled, the channel and power usage of the survey SSID can be manually set on a per-radio basis.
Note: The site survey SSID is only broadcasted when the AP is claimed into a network. If the AP has not been claimed, it will only broadcast the SSID on 2.4GHz.
6. After setting the access point for Site Survey mode, the AP starts broadcasting an open SSID, "site_survey-<Mac address of the AP>" as shown below:
Note: While in Site Survey Mode, the access point will no longer be able to serve connectivity for any wireless clients and can be used only for site survey purposes. The access point is now ready to be used at the location where the site survey needs to be performed. The AP can be powered using a PoE injector or a PoE switch and does not need to have wired connectivity to the internet.
The survey SSID being broadcast from the access point can now be used to perform active/passive site surveys using any professional site survey tool. Examples include:
These tools have the coverage patterns for Cisco Wireless access points built-in, enabling predictive survey functionality as well.
MR 32.1.X Site Survey Configuration Enhancements
With MR 32.1.x firmware the channel width for each radio can be configured for on the local status page of an AP configured for site survey mode.
With MR 32.1.x firmware minimum bitrates can now be configured for on the local status page of an AP configured for site survey mode.
AP Name in Beacon, without CCX
MR 32.1.X firmware brings support to allow AP’s to broadcast their hostname attribute in beacon frames without requiring network support to enable Cisco Compatible Extensions (CCX). This feature inclusion addresses operational challenges faced during site surveys and production network troubleshooting, where CCX was disabled by default. Additionally, the MR 32.1.X firmware introduces an increased character limit for the AP name, previously allowing up to 40 characters.
Hardware: Wi-Fi 6, Wi-Fi 6E, Wi-Fi 7
Firmware: MR 32.1.1+
Configuration:
Once the AP is running MR 32.1.1+ the AP name field in Dashboard can be configured up to 40 characters in length. Additionally, when in site survey mode, CCX IE and this new VS IE will be ON by default.
Note: CCX AP name character length is limited to 16 characters while the VS IE can be set up to 32 characters in length.
Note: Contact support to have your APs configured with the AP name in beacon frame for non site survey functionality.
Note: Must be running Wireshark Version 4.5.0rc0-2355-g3d1b47e14b4f (v4.5.0rc0-2355-g3d1b47e14b4f)or later
Configuring 6 GHz Site Survey
Note: By default, flex radios on 6E APs will default to broadcast in the 5 Ghz band. To utilize the 6 GHz band for site surveys on a Cisco Meraki AP (MR). The flex radio must be provisioned to operate in the 6 GHz band via RF profiles in the Meraki dashboard.
1. Within the Meraki dashboard, navigate to Wireless > Configure > Radio settings > RF profiles
2. Click Edit on the desired RF profile
3. Within Flex radio selection of 6E AP, toggle the band preference to utilize 6 GHz
4. Save the configuration
Using a Survey SSID
Using Channel Widths other than 20Mhz
MR access points in survey mode default to 20Mhz channel widths; this is typically the channel width most used for high density deployments. It is possible to change this channel width for your survey; however, this requires connecting the AP to the internet and making a configuration change on the Meraki dashboard before setting the AP in survey mode. Use the RF profile feature to create a new profile e.g 'surveyor' and apply it to the AP.
Go to Wireless > Radio Settings > RF profiles and create a new profile and under the 5 Ghz set the Channel width to manual. Select between the 40MHz and 80 MHz channels.
Once the profile is created, assign it to the AP.
Make sure to clear any channel width overrides
Check on the dashboard that the access point configuration has been updated (this takes around 30 seconds)
Now when the AP is set to survey mode (from the local status page), it will show the desired change made in the 5Ghz channel width.
 
