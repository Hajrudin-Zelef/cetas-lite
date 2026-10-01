---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360058776614-manage-unifi-talk-subscriptions-3020be85
title: "hc-en-us-articles-360058776614-manage-unifi-talk-subscriptions-3020be85"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "United States"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360058776614-manage-unifi-talk-subscriptions-3020be85.md
source_anchor: ""
source_lines: [1, 62]
sha256: c1f5ed43451f626373e82108e18eaac50de006fa4a5b0f3e6f0a073c3fa4af58
---

# hc-en-us-articles-360058776614-manage-unifi-talk-subscriptions-3020be85

Manage UniFi Talk Subscriptions
With a UniFi Talk subscription, you can easily access local phone numbers that support emergency calling. Each UniFi Talk phone number offers 3,000 monthly minutes for worldwide inbound calls and outbound calls to the US, Canada, and Mexico.
While setting up your UniFi Talk system, you will have the option to activate a subscription with a dedicated number for each UniFi Talk phone connected to the system. If you have multiple phones connected during the setup, you can choose to set up additional phones with extensions only and add them to a group to share a number after setup.
Each UniFi Talk phone number supports unlimited concurrent calls, with a limit of one outgoing call per second.
Notes:
- Currently, UniFi Talk subscription service is only available in the US, the UK, and Canada, but we are working on bringing the service to other territories and countries.
  - Outside of these supported regions, unlocked UniFi Talk phones can be used with third-party SIP providers. Instructions for configuring third-party SIPs can be found here.
- Monthly minutes are shared by all active numbers managed by your UniFi Talk application. Local extension calling is free and does not require a UniFi Talk subscription.
- $0.02/min beyond the first 3,000 minutes. International calls start at $0.01/min. View our international calling rates.
Purchasing and Assigning New Numbers After Setup
Purchase additional number(s) in the UniFi Talk application
- Navigate to UniFi Talk > Settings > Numbers & Subscription.
- Click the Purchase Numbers button.
- Search for or select an area code from the Area Code drop-down menu.
- Select your desired number from the available numbers listed in the Number drop-down menu, or type your preferred number into the field to see if it is available to purchase.
- Repeat this process to add more numbers. Once you have selected all desired numbers, click Next and complete your purchase.
Note: While we work hard to maintain a wide variety of numbers in our inventory, we can't guarantee your preferred number(s) will be in stock at the time of your purchase.
Ported Numbers and Subscriptions
Porting a number into UniFi Talk is free. Once the port is complete, you’ll receive an email notification and the number will appear as Ready for activation in the Phone Lines section. Activation includes purchasing a subscription for that number—existing subscriptions do not apply to newly ported numbers.
Assign newly purchased number(s) to a user
- Navigate to UniFi Talk > Assignments > Users.
- Click on the user you wish to assign the number to.
- Select the Settings tab and locate the Manage section.
- Expand the Number drop-down menu and assign the number to this user.
Upgrading to the UniFi Talk Pro Plan
The UniFi Talk Pro Plan is currently only available in the United States and Canada. It offers higher usage allowances as well as caller ID name (CNAM) lookups for incoming calls (USA only) and support for softphone with the license-free UniFi Identity mobile app for iOS and Android. Learn more about UniFi Talk subscription plans here.
To upgrade to the Pro Plan:
- Navigate to Settings > Numbers & Subscription.
- Click on Upgrade to Pro located at the top of the page or click Upgrade next to a number listed within Talk Numbers.
- Follow the prompts to add or select a payment method to upgrade the number successfully.
Looking up Caller ID Name (CNAM)
Caller ID Name (CNAM) is a service available in the United States public telephone network that displays a personal or business name associated with the calling party’s number. The UniFi Talk Pro Plan includes 3,000 monthly CNAM lookups for incoming calls.
Requirements:
- UniFi Talk Application version 3.1.8 and later
- At least one UniFi Talk Pro Plan subscription
- CNAM is only available in the United States
To enable CNAM lookups for incoming calls:
- Navigate to Settings > Call Settings > Additional
- Select the checkbox next to CNAM Lookup
Using the UniFi Talk Softphone
With the UniFi Talk Softphone, users can make and receive calls and access voicemail while on the go—all from the license-free UniFi Identity mobile app which also offers seamless access and control features such as One-Click WiFi, One-Click VPN, and Door Access. At the moment UniFi Talk softphones are not available for UID Enterprise users.
Requirements:
- UniFi Talk Application version 3.1.8 and later
- At least one UniFi Talk Pro Plan subscription
- Enabled license-free UniFi Identity
To provision a softphone to a UniFi Talk user:
- First, enable Talk Softphone for your UniFi Console at OS Settings > Admins & Users > Identity Settings. From this page, you will also see the number of softphones available to assign to users.
- Then, navigate to OS Settings > Admins & Users and select the user to assign a softphone to. Select the Settings tab of the user’s management side panel and select Softphone.
- Select the Overview tab of the user’s management side panel and click Send to invite the user to UniFi Identity by email.
- Have the user follow the instructions within the invitation email to download the UniFi Identity mobile app and load their credential. Once loaded, the Talk Softphone will be ready for use within the UniFi Identity mobile app.
Removing Numbers and Canceling a UniFi Talk Subscription
A UniFi Talk subscription can be modified or canceled at any time.
If you have access to the UniFi Talk application:
- Navigate to UniFi Talk > Settings > Numbers & Subscription.
- Select the checkbox of each number you wish to delete. In order to delete a number, it must first be unassigned from the Smart Attendant, if applicable.
- Click the Remove and then Delete button to confirm.
Note: This method will immediately remove the number from your UniFi Talk subscription. You will not be refunded for the current billing cycle. Delete all numbers to cancel your UniFi Talk subscription entirely.
Alternatively, you can cancel your UniFi Talk subscription and remove all associated numbers from your Ubiquiti Account portal:
- Access your Ubiquiti Account by logging in at account.ui.com.
- Navigate to Payments & Subscriptions and locate your UniFi Talk subscription under Active Subscriptions.
- Click the three dots that display to the right of your subscription to expand the menu, and click Cancel Subscription.
Note: This method will immediately cancel your UniFi Talk subscription and remove all UniFi Talk numbers associated with this subscription from your account. You will not be refunded for the current billing cycle.
