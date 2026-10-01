---
id: collect-261001-fortinet/fortinet/netconfig-automate-fortigate-static-routes-with-python-and-the-rest-api-4gk2-c59b3aae
title: "netconfig-automate-fortigate-static-routes-with-python-and-the-rest-api-4gk2-c59b3aae"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/netconfig-automate-fortigate-static-routes-with-python-and-the-rest-api-4gk2-c59b3aae.md
source_anchor: ""
source_lines: [1, 78]
sha256: 2dc5cabc19e82733f77f12a024d65cb358f29d69ddf21fa3187cac3b61d0938a
---

# netconfig-automate-fortigate-static-routes-with-python-and-the-rest-api-4gk2-c59b3aae

If you manage FortiGate firewalls and you're still clicking through Network > Static Routes > Create New every time a new subnet comes online — this one's for you.
I recently put together a full walkthrough on automating FortiGate static route provisioning using Python and the FortiOS REST API. Rather than repeat the whole thing here, I'll give you the core of it and link to the full guide at the end.
Why This Actually Matters
One route on one firewall is a two-minute GUI job. Ten routes across twenty firewalls is an afternoon of errors waiting to happen. The FortiOS REST API gives you a clean, auditable, repeatable way to push config — and every change shows up in System Events with a transaction ID, so your audit trail is automatic.
The Endpoint
The CMDB path for static routes is:
POST https://<fortigate-ip>/api/v2/cmdb/router/static
Auth is Bearer token — you create an API admin under System > Administrators > REST API Admin and use the generated token in your request header.
The Gotcha Most People Hit
FortiOS's CMDB API requires your payload attributes to be nested inside a top-level "json" key. Miss this and you'll get a vague 400 with no useful error:
pythonpayload = {
    "json": {                          # <-- this wrapper is required
        "dst": "10.200.10.0 255.255.255.0",
        "gateway": "192.168.70.1",
        "device": "wan1",
        "comment": "Provisioned via Python REST API"
    }
}
Also note: dst takes a space-separated network mask format, not CIDR — the API speaks the underlying CMDB schema, not what the GUI displays.
**The Full Script
import requests
import json
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
FORTIGATE_IP = "192.168.70.86"
API_TOKEN = "your-api-token-here"
DESTINATION_SUBNET = "10.200.10.0 255.255.255.0"
GATEWAY_IP = "192.168.70.1"
OUTBOUND_INTERFACE = "wan1"
url = f"https://{FORTIGATE_IP}/api/v2/cmdb/router/static"
headers = {
    'Content-Type': 'application/json',
    'Authorization': f'Bearer {API_TOKEN}'
}
payload = {
    "json": {
        "dst": DESTINATION_SUBNET,
        "gateway": GATEWAY_IP,
        "device": OUTBOUND_INTERFACE,
        "comment": "Provisioned automatically via Python REST API script"
    }
}
print(f"Sending POST request to add static route to {DESTINATION_SUBNET}...")
response = requests.post(url, headers=headers, json=payload, verify=False)
if response.status_code == 200:
    try:
        output = response.json()
        print("\n[+] Configuration Applied Successfully!")
        print(json.dumps(output, indent=4))
    except json.decoder.JSONDecodeError:
        print("\n[!] Route might be applied, but response body parsing failed.")
else:
    print(f"\n[-] Failed! HTTP Status Code: {response.status_code}")
    print(response.text)
What a Successful Response Looks Like
json{
    "http_method": "POST",
    "revision": "2dce04a4b69aa39900f2e326dbcfdecb",
    "revision_changed": true,
    "mkey": 2,
    "status": "success",
    "http_status": 200,
    "vdom": "root",
    "path": "router",
    "name": "static",
    "version": "v8.0.0"
}
Save that mkey value — you'll need it if you ever want to PUT (update) or DELETE the same route via the API later.
Verifying It Landed
Two quick checks:
GUI: Network > Static Routes — the new entry appears immediately, with your comment field intact. That comment is your visual flag for "automation-managed."
Logs: Log & Report > System Events — look for the API admin's entry with message Add router.static 2. The full log detail shows every attribute you sent in the payload logged verbatim. Instant audit trail, no extra work.
**Want the Full Walkthrough?
**I've written this up in full detail over on my networking blog — covering API admin setup, token scoping, production tips (GET before POST, idempotent PUT updates, token rotation), and the full System Events log breakdown with screenshots:
👉 Add a FortiGate Static Route via Python REST API — netconfig.io
Covers FortiOS 7.x through 8.0, tested on a FortiGate-60F in a live lab environment.
Drop any questions below — happy to help if you hit a weird API error or need to adapt this for multi-VDOM setups.
Top comments (0)
