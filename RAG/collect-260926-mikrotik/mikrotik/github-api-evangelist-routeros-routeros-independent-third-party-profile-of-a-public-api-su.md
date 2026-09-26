---
id: collect-260926-mikrotik/mikrotik/github-api-evangelist-routeros-routeros-independent-third-party-profile-of-a-public-api-su
title: "github-api-evangelist-routeros-routeros-independent-third-party-profile-of-a-public-api-surface-by-a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-11-07", "2026-05-19"]
keywords: ["agent", "pricing"]
source: docs/RAG/lot-mikrotik/RouterOS/github-api-evangelist-routeros-routeros-independent-third-party-profile-of-a-public-api-surface-by-a.md
source_anchor: ""
source_lines: [1, 117]
sha256: 048945e50db6c42beba6aa2538ffff9011c033aaadc23415ef819b64d4e7db25
---

# github-api-evangelist-routeros-routeros-independent-third-party-profile-of-a-public-api-surface-by-a

**This is not our API.** This repository is an independent, third-party profile of a company's
**publicly available** API surface, maintained by API Evangelist.
API Evangelist does not operate, host, resell, or support this company's APIs, and is not
affiliated with or endorsed by the company unless stated on the profile.

**Where the information came from.** Everything here is assembled from material a member of the
public can reach with a browser and no credentials — the company's own website, developer portal
and documentation, the specifications it publishes for public use (OpenAPI, AsyncAPI, JSON Schema,
`apis.json`, `llms.txt` and similar), its public repositories, and its public status, pricing and
changelog pages. **Nothing here is obtained by breaching a system, defeating an access control, or
using credentials of any kind.**

**The rating is an independent assessment.** The Kin Score and Agent Readiness rating are
independently calculated scores of a company's *public* API artifacts, produced by API Evangelist
against a published rubric. They are not certifications, endorsements, security assessments, or
audits, and they score published artifacts — not the quality, safety, or security of the software.

**Corrections, re-scores, and removal are free.** No partnership, contract, or purchase is
required, and you do not need to justify the request.


**Something wrong?** Open an issue on this repository, or email
info@apievangelist.com.
**Published something new?** Ask for a re-score and we will re-run the rating.
**Want the listing taken down?** Say so and we will honor it. The profile is reduced to your
company name, a factual description, and a link to your own site, and the company is recorded as
**unrated** — never scored zero for having asked.
**Response times.** Acknowledgement within **one business day**; removal or restriction within
**two business days**; corrections and re-scores within **five business days**.

**Not from the company, and here with a question?** You are welcome here — we would rather be the
front line and point you the right way than have a good report go nowhere. What this repository
can answer is narrow, though, so it is worth knowing who you are actually looking for:


**A question about how the API works, an account, billing, or a bug in the service** — that is
the company's own support, not us. We profile this API; we do not operate it and cannot see
your account.
**A bug in an open-source project we only catalog** — file it on that project's own repository.
This has happened with a real and correct bug report that reached us instead of the people who
could fix it, which helped nobody.
**Anything about this listing itself** — the description, the tags, the rating, a missing or
wrong artifact — is ours. Open an issue here.
**Not sure, or something general about API Evangelist or APIs.io** — open an issue on the
APIs.io Inbox and we will route it.
**This repository contains no software, and we will never ask you to download anything.** There is
no build, release, installer, or binary here — only text and machine-readable API descriptions, so
there is nothing here that can be "corrupt" or need "repairing". Any issue, comment, or email
claiming otherwise and offering a download link is not from us and is hostile. Do not follow the
link; it is a lure. Report it to GitHub and, if you like, tell us at
info@apievangelist.com so we can take it down.

**On a security or compliance team?** Email
info@apievangelist.com with *security* in the subject line and
you will get a person, not a form. We will tell you exactly which public URLs this profile was
built from so your team can see the same surface we did, and we will take the listing down on
request while you work through it.

Full detail: **Where this data comes from**


RouterOS is MikroTik's powerful network operating system designed for managing routers, switches, access points, and other network devices. It provides a comprehensive REST API (v7.1+) and a TCP-based binary API for programmatic management of IP addresses, interfaces, firewall rules, routing, VPN configurations, DHCP, DNS, and system resources. RouterOS powers MikroTik hardware and can also be deployed as a virtual machine (CHR).

**APIs.json:** https://raw.githubusercontent.com/api-evangelist/routeros/refs/heads/main/apis.yml

- **Type:** Index

- Networking
- Routers
- Network Management
- Firewall
- MikroTik

- **Created:** 2024-11-07
- **Modified:** 2026-05-19

The RouterOS REST API is a JSON wrapper over the RouterOS console API, available from RouterOS v7.1beta4+. It enables create, read, update, and delete operations on all RouterOS configuration menus via standard HTTP methods (GET, PUT, PATCH, DELETE, POST). Authentication uses HTTP Basic Auth with console credentials. Supports filtering, property selection (.proplist), and complex queries. Accessible at https://{router-ip}/rest.

- **Human URL:** https://help.mikrotik.com/docs/spaces/ROS/pages/47579162/REST+API
- **Base URL:**`https://{router-ip}/rest`

- Networking
- Router Management
- REST API
- Network Configuration
- Firewall
- DHCP
- DNS

- Documentation
- Documentation
- OpenAPI — OpenAPI Specification
- Spectral Rules
- JSON Schema — JSON Schema
- JSON Schema — JSON Schema
- JSON Schema — JSON Schema
- JSON Structure
- J S O N L D Context
- Example
- Example
- Example
- Vocabulary
- Postman Collection — Postman Collection 2.1
- Open Collection — Open Collection 1.0

The RouterOS TCP API is the native binary protocol for RouterOS, running on TCP port 8728 (standard) and TCP port 8729 (SSL/TLS). It uses a sentence-based word protocol with variable-length encoding, supporting tagged concurrent commands, streaming changes via /listen, and cancellation. Used by most RouterOS client libraries (Python, PHP, Java, Go, etc.).

- **Human URL:** https://help.mikrotik.com/docs/spaces/ROS/pages/47579160/API
- **Base URL:**`tcp://{router-ip}:8728`

- Networking
- Router Management
- TCP API
- Binary Protocol

**FN:** Kin Lane
**Email:** kin@apievangelist.com
