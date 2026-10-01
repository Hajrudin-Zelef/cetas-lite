---
id: collect-261001-meraki/meraki/t5-developers-apis-sending-alerts-to-webex-directly-from-the-meraki-dashboard-m-66b1a348
title: "t5-developers-apis-sending-alerts-to-webex-directly-from-the-meraki-dashboard-m--66b1a348"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-developers-apis-sending-alerts-to-webex-directly-from-the-meraki-dashboard-m--66b1a348.md
source_anchor: ""
source_lines: [1, 115]
sha256: c672a150ff2667ca7c517c94b0fe3832a288f8845def42f64934c12f37113b65
---

# t5-developers-apis-sending-alerts-to-webex-directly-from-the-meraki-dashboard-m--66b1a348

Sending alerts to Webex directly from the Meraki Dashboard ?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-15-2023 09:56 AM
Sending alerts to Webex directly from the Meraki Dashboard ?
Meraki Dashboard comes with built-in payloads for webhooks (Webex, Slack, Teams, …).
Those payloads are a ready to use solution for interacting with existing applications. But if you are using the given Webex payload (Network-wide > Alerts > Webhooks > Add or edit webhook templates)...
{
 "markdown": "**{{alertType}}**\nAlert ID: {{alertId}}\nAlert level: {{alertLevel}}\nOccurred at: {{occurredAt}}\nOrganization name: [{{organizationName}}]({{organizationUrl}})\nOrganization ID: {{organizationId}}\nNetwork name: [{{networkName}}]({{networkUrl}})\nNetwork ID: {{networkId}}\nNetwork Tags: {{networkTags | join: ', '}}\nDevice name: {{deviceName}}\nDevice serial: [{{deviceSerial}}]({{deviceUrl}})\nDevice MAC: {{deviceMac}}\nDevice tags: {{deviceTags | join: ', '}}\nDevice model: {{deviceModel}}\n**Alert Data** {{ alertData | markdown_with_additional_formatting_for_webex }}\n"
}
...you will need to host a script listening for Meraki webhooks, and then posting a message to Webex API (via your Webex Bot).
This workflow is correct, however for the Custom app, you will need to deal with:
- Scripting (Python, JavaScript, …)
- Networking/Security (NAT rules if you are hosting the application behind your ISP router, firewall rules if you are using a cloud VM with public IP, …)
- Getting computing resources (a Raspberry Pi at your place, a Cloud VM, Zapier, ...)
- Billing (hosting, cloud services, …)
Setting up a self-hosted / cloud-hosted application can thus be time/money consuming.
Workaround
In this article, I am presenting a workaround, allowing you to post Webex messages directly from the Meraki Dashboard using a custom webhook payload:
Same result, but one less hop!
What are custom webhook payload templates ?
Custom payload templates allow you to define the structure of webhooks sent from Meraki Dashboard. By using it, you will be able to decide which data you want to send, and how those data are formatted.
The structure will be written in Liquid, both for the header and the body of the HTTP POST request.
Create a new webhook payload template
A new webhook payload template can be created from the Meraki Dashboard (Network-wide > Alerts > Webhooks > Add or edit webhook templates).
Or via Meraki APIs:
https://developer.cisco.com/meraki/api-latest/#!create-network-webhooks-payload-template
Custom webhook payload template requirements
The custom webhook payload template needs to:
- Be sent to Webex API (https://webexapis.com/v1/messages)
- Authenticate to the Webex API with the bot token
- Include expected formated data
- (Optional) Provide an Adaptive Card for better display
For the given requirements, I am going to use the Meraki Dashboard.
… Be sent to Webex API
Network-wide > Alerts > Webhooks
… Authenticate to the Webex API with the bot token
Network-wide > Alerts > Webhooks > Add or edit webhook templates
… Include expected formated data
Network-wide > Alerts > Webhooks > Add or edit webhook templates
Then, do not forget to apply the custom webhook template to your alerts (Network-wide > Alerts):
… (Optional) Provide an Adaptive Card for better display
Network-wide > Alerts > Webhooks > Add or edit webhook templates
In the Liquid Body, you can keep the markdown format of the native Webex payload template, or you can specify an Adaptive Card for better display.
An example of card is available on the following GitHub repo : https://github.com/xaviervalette/meraki-webex-custom-webhooks/blob/main/liquidBody.json
To create your own custom Adaptive Card, you can use the Button and Card Designer Tool of Webex.
Here is the result! I hope it helps!
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-15-2023 11:52 AM
That is much better. No one in this space really wants to be hosting their own code.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-17-2023 06:55 AM
Just for completion.
You don't have to host something. Webex is also supporting an integration for incoming webhooks to post something into a room
https://apphub.webex.com/applications/incoming-webhooks-cisco-systems-38054-23307
I'm using this and the basic integrated template.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-29-2023 02:48 PM
Hello Guys,
The Adaptive Card is possible used with incoming-webhooks?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-30-2023 05:41 PM
Hi @xvalette , as @Greenberet mentioned, this statement is inaccurate:
@xvalette wrote:
But if you are using the given Webex payload (Network-wide > Alerts > Webhooks > Add or edit webhook templates)...
...you will need to host a script listening for Meraki webhooks, and then posting a message to Webex API (via your Webex Bot).
Webex supports incoming webhooks out of the box. Link: https://apphub.webex.com/applications/incoming-webhooks-cisco-systems-38054-23307-75252
