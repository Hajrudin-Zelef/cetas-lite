---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-alert-webex-bot-notification-b655813d-2
title: "codeexchange-github-repo-gve-sw-gve-devnet-meraki-alert-webex-bot-notification-b655813d"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "sandbox"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-alert-webex-bot-notification-b655813d.md
source_anchor: ""
source_lines: [73, 111]
sha256: 0574678124dc573a5d005edca590097efd0afd8a6548026c8ce1e0dab07b3c4a
---

# codeexchange-github-repo-gve-sw-gve-devnet-meraki-alert-webex-bot-notification-b655813d

WEBEX_ROOM_ID = "<insert_webex_room_id>"
BEEHIVE_ID = "<insert_beehive_id>" # optional
BEEP_API_TOKEN = "<insert_beep_api_token>"
MERAKI_API_KEY = "<insert_meraki_api_key>" MERAKI_NETWORK_ID = "<insert_meraki_network_id>"
MERAKI_CAMERA_SERIAL_FRONT = "<insert_meraki_camera_serial_front>" MERAKI_CAMERA_SERIAL_SIDE = "<insert_meraki_camera_serial_side>" MERAKI_CAMERA_SERIAL_DISTANCE = "<insert_meraki_camera_serial_distance>" MERAKI_CAMERA_SERIAL_AP = "<insert_meraki_camera_serial_ap>" MERAKI_DEVICE_CLIENTS = ["<insert_list_of_names>"]
Now is the time to launch the application! If you are testing locally, then you need to conduct one more step in between using ngrok.
Start an ngrok tunnel using the following command:
    $ ngrok http https://0.0.0.0:5001
You will obtain two public ip addresses. Copy the https address and paste that into the TEAMS_BOT_URL variable in the .env file.
If you have the application running on a server exposed with a public IP address, then copy the https address and paste that into the TEAMS_BOT_URL variable in the .env file.
Now simply type in the following command in your terminal:
$ python bot.py
And head over to the following url:
https://0.0.0.0:5001
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
Provide links to related white papers:
Provide a link to a related DevNet Sandbox:
Provide links to related Learning Labs or modules on DevNet:
Provide links to related solutions on DevNet Ecosystem Exchange:
Owner
Contributors
Categories
CollaborationIoTNetworkingSecurity
Products
MerakiWebex
Programming Languages
Python
License
Code Exchange Community
Get help, share code, and collaborate with other developers in the Code Exchange community.View Community
Lorsque vous consultez un site Web, celui-ci peut enregistrer ou récupérer des informations présentes dans votre navigateur, le plus souvent sous la forme de témoins. Ces informations peuvent vous concerner, ainsi que vos préférences ou votre périphérique, et servent principalement à faire fonctionner le site selon vos attentes. Généralement, ces informations ne permettent pas de vous identifier directement, mais elles sont utilisées afin de vous offrir une expérience Web personnalisée. Parce que nous respectons votre droit à la vie privée, vous pouvez choisir de ne pas autoriser certains types de témoins. Cliquez sur les différents en-têtes de catégorie pour en apprendre davantage et modifier nos paramètres par défaut. Toutefois, bloquer certains types de témoins peut avoir une incidence sur votre expérience du site et les services que nous proposons.
Ces cookies sont indispensables au bon fonctionnement du site web et ne peuvent pas être désactivés au niveau de nos systèmes. Ils ne résultent généralement que d'actions que vous avez effectuées et qui correspondent à une demande de service, comme lorsque vous définissez vos préférences en matière de confidentialité, que vous vous connectez ou que vous remplissez des formulaires. Vous pouvez configurer votre navigateur pour qu'il bloque ces cookies ou vous avertisse de leur présence, mais certaines parties du site ne seront alors plus opérationnelles. Ces cookies n'enregistrent aucune information d'identification personnelle.
Ces cookies fournissent des indicateurs relatifs aux performances et à la facilité d'utilisation de notre site. Ils visent principalement à récolter des informations sur la façon dont vous interagissez sur le site, notamment le temps de chargement des pages, les temps de réponse, les messages d'erreur et le déroulement des interactions pour chaque visiteur. Nous pouvons ainsi examiner et analyser le comportement des visiteurs pour améliorer l'ergonomie et les fonctionnalités de notre site. Ces cookies nous permettent également de comptabiliser les visites et les sources de trafic afin de mesurer et d'améliorer les performances du site. Ils nous aident à savoir quelles pages sont les plus et les moins populaires, et à voir la manière dont les visiteurs naviguent sur le site. Si vous n'autorisez pas ces cookies, nous ne saurons pas si vous avez déjà visité le site et ne pourrons pas en mesurer les performances.
Ces cookies peuvent être intégrés à notre site par nos partenaires publicitaires. Ils peuvent être utilisés par ces entreprises pour établir un profil de vos centres d'intérêt et vous présenter des publicités pertinentes sur d'autres sites. Si ces cookies n'enregistrent pas directement vos informations personnelles, ils identifient de manière unique votre navigateur et votre terminal d'accès à Internet. Si vous n'autorisez pas ces cookies, des publicités moins ciblées s'afficheront.
Ces cookies permettent d’améliorer et de personnaliser les fonctionnalités du site Web. Ils peuvent être activés par nos équipes, ou par des tiers dont les services sont utilisés sur les pages de notre site Web. Si vous n'acceptez pas ces cookies, une partie ou la totalité de ces services risquent de ne pas fonctionner correctement.
