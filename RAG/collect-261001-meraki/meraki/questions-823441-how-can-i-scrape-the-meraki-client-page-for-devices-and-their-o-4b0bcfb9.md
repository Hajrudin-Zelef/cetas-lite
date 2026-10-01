---
id: collect-261001-meraki/meraki/questions-823441-how-can-i-scrape-the-meraki-client-page-for-devices-and-their-o-4b0bcfb9
title: "Browse the Meraki website"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-823441-how-can-i-scrape-the-meraki-client-page-for-devices-and-their-o-4b0bcfb9.md
source_anchor: ""
source_lines: [1, 18]
sha256: 6bbfd73076611cfba2033290b2fa4c9e8fdc7b69e39eb1970c4a8c9543829bc6
---

# Browse the Meraki website

I am trying to automate a weekly process where I download a CSV copy from the Client page on the Meraki website.
Before I lose you I am aware there is an API and it works very well. However there in a device attribute that is not exposed in the API. I got this from their support:
  At this point, the "Owner" field is not accessible through the API, and it is currently in development.
I need the Owner as that is how I associate devices to people in our organization.
Using PowerShell I am able to successfully log into the website and get a 200 response for a basic page. My problem comes when I try to use the same session to get to the data I want I keep getting a page with just breadcrumbs and no actual data.
# Browse the Meraki website
$response = Invoke-WebRequest -Uri "https://n116.meraki.com/login/dashboard_login" -SessionVariable meraki
# Form Credentials
$formFields = @{
    email = "my address"
    password = "my password"
}
# Authenticate to the site
Invoke-WebRequest -Uri $response.Forms.Action -WebSession $meraki -Body $formFields -Method Post
# Get the client list
Invoke-WebRequest -Uri "https://n116.meraki.com/my-company/n/_-tSGb0b/manage/pcc/list" -WebSession $meraki -Method Get
Worst case is I am hoping to get the parsedHtml of the Client List table so I can work with its contents. I can't seem to get any useful data no matter what page. I am not sure if the issue is working with Meraki or my PowerShell approach. I have only done simple scraping up until this point. 
I realize this is a very specific request but I am curious to know if someone that has Meraki is able or has been able to get this information using these means.
