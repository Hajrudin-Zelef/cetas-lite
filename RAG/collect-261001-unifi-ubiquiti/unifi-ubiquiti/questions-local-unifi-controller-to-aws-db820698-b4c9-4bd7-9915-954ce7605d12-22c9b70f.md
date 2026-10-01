---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-local-unifi-controller-to-aws-db820698-b4c9-4bd7-9915-954ce7605d12-22c9b70f
title: "questions-local-unifi-controller-to-aws-db820698-b4c9-4bd7-9915-954ce7605d12-22c9b70f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-local-unifi-controller-to-aws-db820698-b4c9-4bd7-9915-954ce7605d12-22c9b70f.md
source_anchor: ""
source_lines: [1, 17]
sha256: 556410c9e0280caf56916cb85ee97f5fbee1dd8d165e49893d0035cb3b4993fd
---

# questions-local-unifi-controller-to-aws-db820698-b4c9-4bd7-9915-954ce7605d12-22c9b70f

@UI-Team
I am a audio video installer and have about 40 differnt sites(all are on the same unifi controller) on a laptop that i use to program on site. I recently just set up a AWS ec2 with unifi. I understand these sites that i already set up do not point back to my AWS Machine. But from time to time i end up doing service calls on these sites and i will want to migrate these from my local computer to my AWS server.
Should the Unifi on my local computer be loaded on to my AWS?
What is the best way to get these already setup account over to my new cloud? (I realize i need to be on site)
@actionav =>Actually yes, you'll need to be onsite at least one time per site ;-(
To migrate allmost everything (not the statistics, but they are probably empty as your laptop is a controller)
1) You need to export the sites from the laptop, and import the sites to the new AWS controller
2) You need to prepare the security group of your AWS controller (firewall), to be sure the port 8080 is open and accessible from your sites
3) You need your controller to have an Elastic IP (fix public IP)
4) You need eventually to have your controller IP mapped to a statis DNS name (so, the controller is accessible via a DNS name and not an IP, better for later migration/upgrade)
5) You need in the controller you got on your laptop and on the AWS controller to change the "Settings => Controller => Controller Settings => Controller hostname/IP" set to the new AWS controller IP/Name, and the "Override inform host with Controller Hostname/IP" enabled
6) go onsite with the laptop, and wait a little, each elements will first connect to the laptop controller (as usual), but then will be provisioned with the public IP/DNS name of the AWS controller, and you must then see them in the AWS controller step by step
@actionav => you should ask in the Unifi routing/switching ro Unifi Wireless section, it's better.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
