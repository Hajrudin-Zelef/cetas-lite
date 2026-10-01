---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d-3
title: "art-of-wifi-unifi-api-client-91bb094d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d.md
source_anchor: ""
source_lines: [235, 253]
sha256: 1a9458d94261a2baa065a9b5bf1ab66a5c1ce5b78bfc5dc0b74a62c97398ae55
---

# art-of-wifi-unifi-api-client-91bb094d

With versions 1.x.x of the API client, the entire client was contained within a single file which can be useful in specific cases. This has changed with version 2.0.0 where the code is now split across multiple files and inclusion in your project is managed using composer.
If you are looking for the version 1.x.x code, you can tell composer to install that version by using the following
syntax in your composer.json file:
{
    "require": {
        "art-of-wifi/unifi-api-client": "^1.1"
    }
}
Alternatively, you can download the latest 1.x.x code from the releases page.
Whenever necessary, we will make sure to update the version_1 branch with the latest 1.x.x code.
This API Client class is based on the initial work by the following developers:
- domwo: https://community.ui.com/questions/little-php-class-for-unifi-api/933d3fb3-b401-4499-993a-f9af079a4a3a
- fbagnol: https://github.com/fbagnol/class.unifi.php
and the API as published by Ubiquiti:
A big thanks to all the contributors who have helped with this project!
If you would like to contribute to this project, please open an issue and include your suggestions or code there or else create a pull request.
Art of WiFi develops software and tools that enhance the capabilities of UniFi networks. From captive portals and reporting solutions to device search utilities, our goal is to make UniFi deployments more powerful and easier to manage.
If you're looking for a specific solution or just want to see what else we offer, feel free to explore our website:
Many of the functions in this API client class are not officially supported by Ubiquiti and as such, may not be supported in future versions of the UniFi Controller API.
