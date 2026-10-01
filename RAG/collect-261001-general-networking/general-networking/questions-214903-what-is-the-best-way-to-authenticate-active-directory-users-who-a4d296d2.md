---
id: collect-261001-general-networking/general-networking/questions-214903-what-is-the-best-way-to-authenticate-active-directory-users-who-a4d296d2
title: "questions-214903-what-is-the-best-way-to-authenticate-active-directory-users-who-a4d296d2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-214903-what-is-the-best-way-to-authenticate-active-directory-users-who-a4d296d2.md
source_anchor: ""
source_lines: [1, 4]
sha256: 48475be5ea4427a36deb3a23ce43ad54a62e7fe511b8a7770c4b0e1d0b90e44a
---

# questions-214903-what-is-the-best-way-to-authenticate-active-directory-users-who-a4d296d2

I know .1x is the protocol that actually does the job, does RADIUS have to be involved in this process, .1x only works with RADIUS (AKA NPS on server-side Windows)?
1 Answer 1
802.1X authentication has three parties: a supplicant, an authenticator, and an authentication server. The RADIUS (or its proposed successor Diameter) is used for the communication between an authenticator (e.g. a WiFi AP) and the authentication server. In this case, you are using Active Directory as the backend authentication server, and you need RADIUS for that. The AP (or the controller for your WiFi network) could also have an internal database and perform the authentication against it, but that's never the case with the Active Directory.
While some authentication methods available might also use e.g. LAPD queries for authentication directly against the Active Directory server, RADIUS has its advantages. One is the possibility to authenticate AD computer accounts instead of AD user accounts: This authentication doesn't involve saving passwords to the wireless network settings, and enables you to only allow Active Directory joined computers to access the network. With user based authentication you either allow practically all BYOD devices, or you are required to manually update a whitelist of allowed devices. With the help of Group Policy you can make all company computers connect to the correct WiFi network without any user actions at all. That profile can be forced to trust only certificates authenticated by a certain CA. This makes RADIUS the most advanced way to authenticate against the Active Directory.
