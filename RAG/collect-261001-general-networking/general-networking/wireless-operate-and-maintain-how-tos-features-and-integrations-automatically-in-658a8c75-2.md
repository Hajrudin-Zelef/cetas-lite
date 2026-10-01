---
id: collect-261001-general-networking/general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75-2
title: "wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75.md
source_anchor: ""
source_lines: [291, 387]
sha256: c730ee941848e203d234a7cf0bdacc7757f70ca8dc5353b6d5448d77e1138c66
---

# wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75

    Cryptomining - Allows you to block identities from accessing known cryptomining pools where miners group together and share resources—processing power—to better gather and share cryptocurrencies, and from known web cryptomining source code repositories
Content Categories Definitions
- 
    Adult Themes—Sites that are adult in nature and are not defined in other rating categories
- 
    Adware—Sites that distribute applications that display advertisements without the user's knowledge or choice. It does NOT include sites that serve advertising
- 
    Alcohol—Sites about alcohol use, commercial and otherwise
- 
    Chat—Sites where you can chat in real time with groups of people; includes IRC and video-chat sites
- 
    Classifieds—Sites for buying and selling (or bartering) goods and services; includes sites with real-estate and housing listings
- 
    Dating—Sites for meeting other people
- 
    Drugs—Sites about illegal or recreational drug use
- 
    File Storage—Sites that offer space for hosting, sharing, and backup of digital files
- 
    Forums/Message Boards—Sites with discussions, including bulletin boards, message boards, and forums
- 
    Gambling—Sites that offer gambling or information about gambling
- 
    Games—Sites that offer game-play and information about games (news, tips, cheat codes)
- 
    German Youth Protection—Content deemed harmful to minors. This category helps prevent viewing of youth-endangering content in Germany. Blocked pages for this category will include German text. This list is not controlled by Umbrella and is created to be controlled by the BPjM (Federal Review Board for Media Harmful to Minors) to be compliant with German law. For more information, see http://www.bundespruefstelle.de/bpjm/Service/english.html. 
Note: Cisco Umbrella does not guarantee compliance with German law.
- 
    Hate/Discrimination—Sites that promote intolerance based on gender, age, race, nationality, religion, sexual orientation, or other group identities
- 
    Instant Messaging—Sites that offer access or software to communicate in real time with other individuals
- 
    Internet Watch Foundation (IWF)—Sites that contain child sexual abuse content; for more information about this category, see Internet Watch Foundation
- 
    Lingerie/Bikini—Sites displaying or dedicated to clothing that could be considered adult-only
- 
    Nudity—Sites that provide images or representations of nudity
- 
    P2P/File Sharing—Sites that facilitate the sharing of digital files between individuals, especially through peer-to-peer software, including torrent sites
- 
    Photo Sharing—Sites for sharing photographs, galleries, and albums
- 
    Pornography—Anything relating to pornography, including mild depiction, soft pornography, or hard-core pornography
- 
    Proxy/Anonymizer—Sites providing proxy bypass information or services; also, sites that allow the user to surf the net anonymously or send anonymous emails
- 
    Sexuality—Sites that provide information, images, or implications of bondage, sadism, masochism, fetish, beating, body piercing, or self-mutilation; this category is not intended for LGBT-related sites that do not fall under the aforementioned criteria
- 
    Social Networking—Sites that promote interaction and networking between people
- 
    Tasteless—Sites that contain information on subjects such as mutilation, torture, horror, or the grotesque; includes pro-Anorexia and pro-suicide related sites
- 
    Terrorism—Sites that promote terrorism or are linked with terrorist organizations
- 
    Video Sharing—Sites for sharing video content
- 
    Visual Search Engines—Sites that allow searching for images based on keywords
- 
    Weapons—Sites about weapons, commercial and otherwise
- 
    Webmail—Sites that offer the ability to send or receive email
Note: If Umbrella blocks a website based on configured content filtering policy, you can do a domain lookup on OpenDNS.com to learn more about the categorization.
Note: If you are running into issues with this integration, please contact Meraki Support. A Meraki Support Engineer will escalate with Umbrella Support on your behalf if necessary. There is no need to contact Umbrella Support directly, as they cannot support this Meraki feature.
Provisioning Automatic Umbrella in Co-term and PDL Organizations
Please follow these steps to provision Automatic Umbrella:
- Claim an appropriate quantity of MR Advanced licenses or MR Enterprise + MR Upgrade licenses on the Organization > Configure > License info page. Licenses also need to be assigned to specific APs in a PDL organization.
- Go to the Wireless > Configure > Firewall & Traffic shaping page and click the Enable Umbrella protection button.
3. Select the checkmark next to Automatically provision Umbrella and click the Auto Provision button.
4. Confirm that you want to provision Automatic Umbrella.
Warning: This process is irreversible. Once the Automatic Umbrella is provisioned, using Manual integration with Umbrella in your organizations will not be possible.
5. Wait a few minutes for the provisioning process to complete.
Enabling Automatic Umbrella Integration on an SSID
- Navigate to Wireless > Configure > Firewall & traffic shaping
- Select the SSID from the dropdown on the top of the page.
- Link your Meraki organization to the Umbrella organization we automatically provisioned for you in the previous step to retrieve the list of the predefined Umbrella policies.
4. Select the desired predefined policy and save the changes at the bottom of the page.
Disabling Automatic Umbrella Integration on an SSID
To disable Umbrella protection on an SSID
- Navigate to Wireless > Configure > Firewall & traffic shaping
- Select the SSID from the dropdown on the top of the page.
- Go to Block applications and content categories and click Disable Umbrella protection.
4. Confirm your desired changes by clicking Yes and save changes on the bottom of the page.
Note: Disabling Automatic Umbrella simply "unlinks" your Meraki organization from the automatically created Umbrella organization; however, the API integration between your Meraki organization and the Umbrella organization still exists. You can always link your SSID back to your automatically provisioned Umbrella organization, as shown below.
Therefore, it's not possible to manually link such a Meraki organization to your own Umbrella organization via API. You can confirm that your Meraki organization is still connected to the automatically provisioned Umbrella organization by checking the Network-wide > Configure > General page (Cisco Umbrella account).
DNS Umbrella Exclusions
Warning: Changing an excluded domain (adding or removing) will result in all clients being temporarily disconnected from the SSID.
While it’s impossible to customize predefined Umbrella policies, it’s possible to exclude some domain names from Umbrella protection.
- Navigate to Wireless > Configure > Access control
- Select the desired SSID from the dropdown on the top of the page
- Select Bridge mode: Make clients part of the LAN in the Addressing and traffic section
- Save changes on the bottom of the page.
- Navigate to Wireless > Configure > Firewall & traffic shaping
- Select the SSID from the dropdown on the top of the page.
- Go to Block applications and content categories
- Add the desired domain names that should be excluded from Umbrella protection.
- Save changes on the bottom of the page.
DNS requests for excluded domains are not redirected to Umbrella and are forwarded to the DNS server specified by the client. This behavior is beneficial for preventing DNS requests for local resources from being sent to Umbrella instead of allowing them to reach internal DNS servers to resolve correctly. MRs automatically add the '.local' and 'in-addr.arpa' domains to be excluded from Umbrella redirection by default.
