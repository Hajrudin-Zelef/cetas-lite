---
id: collect-261001-general-networking/general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75-1
title: "wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2022-05", "2025-04-26"]
keywords: ["license", "licenses", "research", "training"]
source: docs/RAG/collect-261001-general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75.md
source_anchor: ""
source_lines: [1, 290]
sha256: 3492bb9782a3f1ecd87ddae87b048b5a0ba922dd730eaea23600c73e8f5eaef1
---

# wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75

Automatically Integrating Cisco Umbrella with Meraki Networks
Learn more with these free online training courses on the Meraki Learning Hub:
Overview
Warning: Cisco Meraki deprecated the ability for customers to automatically integrate Cisco Umbrella (Umbrella) with their Meraki organizations on April 26, 2025. After the deprecation date, however, customers may still manually integrate Umbrella with their new Meraki organizations by following the instructions in our documentation.
Automatic Umbrella integration allows Meraki administrators to link their Meraki dashboard with the Cisco Umbrella dashboard effortlessly and easily assign predefined Umbrella content filtering and security policies to protect wireless clients from malicious content or unwanted web resources.
Once you assign desired Umbrella policies, MR access points will intercept all DNS requests from wireless clients and redirect them to Cisco's Umbrella for evaluation. Based on the disposition from Umbrella, the client’s request will either be allowed or blocked.
Warning: If Automatic Umbrella is explicitly provisioned, using Manual integration with Umbrella in such organizations will not be possible. This process is irreversible.
Prerequisites
This integration is available in both per-device (PDL) and co-termination (Co-Term) licensing models.
Note: Before May 2022, adding an MR Advanced or MR Upgrade license to a co-term organization would initiate a PDL conversion process; however, this is no longer the case. Conversions from co-term to PDL are no longer supported, as referenced in the Per-Device Licensing article.
Note: MR Upgrade licenses are intended to be used as an add-on for MR Enterprise licenses and are not standalone licenses.
Requirements for PDL Organizations
- 
    All MR access points in the desired Meraki network must have MR Advanced or MR Enterprise + MR Upgrade licenses.
- 
    Network(s) where you wish to enable this integration must run to MR 26.1 or more recent firmware.
Note: Due to the nature of the PDL model, it’s possible to have this integration enabled in some (but not all) networks within a single Meraki organization.
Requirements for Co-Term Organizations
- 
    All MR access points in the Meraki organization must have MR Advanced or MR Enterprise + MR Upgrade licenses. Once your Meraki organization meets this condition, its MR Product Edition will change to Advanced Enterprise.
- 
    Network(s) where you wish to enable this integration must run to MR 26.1 or more recent firmware.
Note: Due to the nature of the co-term model, it’s not possible to have this integration enabled in some (but not all) networks within a single Meraki organization. Please refer to the Meraki MR Licensing Guide to learn more about MR license types.
Predefined Umbrella Policies in the Meraki Dashboard
There are seven predefined Umbrella policies, which consist of different combinations of security settings and content filtering.
Security & Appropriate Use Filtering
- 
    Security & Full Appropriate Use Filtering
- 
    Security & Basic Appropriate Use Filtering
- 
    Security Filtering Only
- 
    Security & Moderate Appropriate Use Filtering (Default)
Appropriate Use Filtering
- 
    Full Appropriate Use Filtering
This policy is targeting a corporate SSID use case with the most restrictive content policies. Corporate employees should not have access to categories like Alcohol, Chat, Dating, Drugs, Gambling, Games, Instant Messaging, Lingerie/Bikini, Nudity, Photo Sharing, Pornography, Social Networking, Video Sharing, and others.
- 
    Moderate Appropriate Use Filtering
This policy is meant for the guest SSID use case. Content settings for this category will allow users to access common chat apps (e.g. Facebook Messenger, WhatsApp), file storage platforms (e.g. box.com, dropbox.com), photo sharing (e.g. instagram.com), social networking (e.g. facebook.com, twitter.com), and video sharing (e.g. youtube.com) while being blocked from visiting Drugs, Gambling, Hate/Discrimination, Lingerie/Bikini, Nudity, Pornography, Terrorism, Weapons, and other content categories that should not be accessed on a typical guest wireless network.
- 
    Basic Appropriate Use Filtering
This policy is meant for school environments. Students will be protected from viewing inappropriate content, while still being allowed to do the necessary research for their classwork or homework.
Predefined Umbrella Policies Breakdown
Security & Full Appropriate Use Filtering
Security settings:
- Malware
- C&C Callbacks
- Phishing Attacks
- Cryptomining
Content filtering settings:
- 
    Adult Themes
- 
    Adware
- 
    Alcohol
- 
    Chat
- 
    Classifieds
- 
    Dating
- 
    Drugs
- 
    File Storage
- 
    Forums/Message Boards
- 
    Gambling
- 
    Games
- 
    German Youth Protection
- 
    Hate/Discrimination
- 
    Instant Messaging
- 
    Internet Watch Foundation
- 
    Lingerie/Bikini
- 
    Nudity
- 
    P2P/File Sharing
- 
    Photo Sharing
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Social Networking
- 
    Tasteless
- 
    Terrorism
- 
    Video Sharing
- 
    Visual Search Engines
- 
    Weapons
- 
    Webmail
Security & Basic Appropriate Use Filtering
Security settings:
- Malware
- C&C Callbacks
- Phishing Attacks
- Cryptomining
Content filtering settings:
- 
    German Youth Protection
- 
    Internet Watch Foundation
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Tasteless
Security Filtering Only
Security settings:
- Malware
- C&C Callbacks
- Phishing Attacks
- Cryptomining
Content filtering settings - none.
Security & Moderate Appropriate Use Filtering (Default)
Security settings:
- Malware
- C&C Callbacks
- Phishing Attacks
- Cryptomining
Content filtering settings:
- 
    Adware
- 
    Alcohol
- 
    Dating
- 
    Drugs
- 
    Gambling
- 
    German Youth Protection
- 
    Hate/Discrimination
- 
    Internet Watch Foundation
- 
    Lingerie/Bikini
- 
    Nudity
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Tasteless
- 
    Terrorism
Full Appropriate Use Filtering
Security settings - none
Content filtering settings:
- 
    Adult Themes
- 
    Adware
- 
    Alcohol
- 
    Chat
- 
    Classifieds
- 
    Dating
- 
    Drugs
- 
    File Storage
- 
    Forums/Message Boards
- 
    Gambling
- 
    Games
- 
    German Youth Protection
- 
    Hate/Discrimination
- 
    Instant Messaging
- 
    Internet Watch Foundation
- 
    Lingerie/Bikini
- 
    Nudity
- 
    P2P/File Sharing
- 
    Photo Sharing
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Social Networking
- 
    Tasteless
- 
    Terrorism
- 
    Video Sharing
- 
    Visual Search Engines
- 
    Weapons
- 
    Webmail
Moderate Appropriate Use Filtering
Security settings - none
Content filtering settings:
- 
    Adware
- 
    Alcohol
- 
    Dating
- 
    Drugs
- 
    Gambling
- 
    German Youth Protection
- 
    Hate/Discrimination
- 
    Internet Watch Foundation
- 
    Lingerie/Bikini
- 
    Nudity
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Tasteless
- 
    Terrorism
Basic Appropriate Use Filtering
Security settings - none
Content filtering settings - block the following categories:
- 
    German Youth Protection
- 
    Internet Watch Foundation
- 
    Pornography
- 
    Proxy/Anonymizer
- 
    Sexuality
- 
    Tasteless
Security Categories Definitions
- 
    Malware - Block requests to access servers hosting malware and compromised websites through any application, protocol, or port.
- 
    Command Control Callbacks - Prevent compromised devices from communicating with hackers' command and control servers through any application, protocol, or port, and help identify potentially infected machines on your network
- 
    Phishing Attacks - Protect users from fraudulent hoax websites designed to steal personal information
- 
