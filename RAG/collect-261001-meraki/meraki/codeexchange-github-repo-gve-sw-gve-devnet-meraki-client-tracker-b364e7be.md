---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-client-tracker-b364e7be
title: "codeexchange-github-repo-gve-sw-gve-devnet-meraki-client-tracker-b364e7be"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-client-tracker-b364e7be.md
source_anchor: ""
source_lines: [1, 48]
sha256: 906cf251e074152f7f77b494b719df4b3e7ac9739f4e97ebcd036e0f3667055d
---

# codeexchange-github-repo-gve-sw-gve-devnet-meraki-client-tracker-b364e7be

This flask app gathers Client Details (MAC, IP, CDP Neighbors, etc.) across Catalyst Switches and across the Meraki networks in an organization. Information is gathered based on MAC address and time period (optionally: IP Address).
The app helps track a client across a mixed network, identifying the most recent connections on both Catalyst Switches and Meraki Networks.
Note:
In order to use the Meraki API, you need to enable the API for your organization first. After enabling API access, you can generate an API key. Follow these instructions to enable API access and generate an API key:
Organization > Settings > Dashboard API accessEnable access to the Cisco Meraki Dashboard APIMy Profile > API accessGenerate API key
For more information on how to generate an API key, please click here.
Note: You can add your account as Full Organization Admin to your organizations by following the instructions here.
git clone [repository name].env_sample file to .env. Rename config_sample.py to config.py.
MERAKI_API_KEY = ""
ORG_NAME = ""
SWITCH_INFO = [
     {
        "device_type": "cisco_ios",
        "ip": "X.X.X.X",
        "username": "",
        "password": "",
        "secret": ""
    }   
]
Note: Please ensure SSH is configured on each Catalyst Switch, as this is required for the Netmiko SSH connection. For more information on Netmiko, consult this guide.
pip3 install -r requirements.txt
To run the program, use the command:
flask run
Navigate to the Flask URL, and the main page will be displayed:
To get client information, provide a MAC Address (or IP), and a Time Period (or Custom Interval). The script will first query all Catalyst Switches provided in the config.py list, then it will query the Meraki Cloud.
The 3 Main Pieces of Client Data Are:
Each piece of data can be downloaded as an Excel file as well.
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
Owner
Contributors
Categories
Networking
Products
MerakiCatalyst Switches
Programming Languages
CSS
License
Code Exchange Community
Get help, share code, and collaborate with other developers in the Code Exchange community.View Community
Lorsque vous consultez un site Web, celui-ci peut enregistrer ou récupérer des informations présentes dans votre navigateur, le plus souvent sous la forme de témoins. Ces informations peuvent vous concerner, ainsi que vos préférences ou votre périphérique, et servent principalement à faire fonctionner le site selon vos attentes. Généralement, ces informations ne permettent pas de vous identifier directement, mais elles sont utilisées afin de vous offrir une expérience Web personnalisée. Parce que nous respectons votre droit à la vie privée, vous pouvez choisir de ne pas autoriser certains types de témoins. Cliquez sur les différents en-têtes de catégorie pour en apprendre davantage et modifier nos paramètres par défaut. Toutefois, bloquer certains types de témoins peut avoir une incidence sur votre expérience du site et les services que nous proposons.
Ces cookies sont indispensables au bon fonctionnement du site web et ne peuvent pas être désactivés au niveau de nos systèmes. Ils ne résultent généralement que d'actions que vous avez effectuées et qui correspondent à une demande de service, comme lorsque vous définissez vos préférences en matière de confidentialité, que vous vous connectez ou que vous remplissez des formulaires. Vous pouvez configurer votre navigateur pour qu'il bloque ces cookies ou vous avertisse de leur présence, mais certaines parties du site ne seront alors plus opérationnelles. Ces cookies n'enregistrent aucune information d'identification personnelle.
Ces cookies fournissent des indicateurs relatifs aux performances et à la facilité d'utilisation de notre site. Ils visent principalement à récolter des informations sur la façon dont vous interagissez sur le site, notamment le temps de chargement des pages, les temps de réponse, les messages d'erreur et le déroulement des interactions pour chaque visiteur. Nous pouvons ainsi examiner et analyser le comportement des visiteurs pour améliorer l'ergonomie et les fonctionnalités de notre site. Ces cookies nous permettent également de comptabiliser les visites et les sources de trafic afin de mesurer et d'améliorer les performances du site. Ils nous aident à savoir quelles pages sont les plus et les moins populaires, et à voir la manière dont les visiteurs naviguent sur le site. Si vous n'autorisez pas ces cookies, nous ne saurons pas si vous avez déjà visité le site et ne pourrons pas en mesurer les performances.
Ces cookies peuvent être intégrés à notre site par nos partenaires publicitaires. Ils peuvent être utilisés par ces entreprises pour établir un profil de vos centres d'intérêt et vous présenter des publicités pertinentes sur d'autres sites. Si ces cookies n'enregistrent pas directement vos informations personnelles, ils identifient de manière unique votre navigateur et votre terminal d'accès à Internet. Si vous n'autorisez pas ces cookies, des publicités moins ciblées s'afficheront.
Ces cookies permettent d’améliorer et de personnaliser les fonctionnalités du site Web. Ils peuvent être activés par nos équipes, ou par des tiers dont les services sont utilisés sur les pages de notre site Web. Si vous n'acceptez pas ces cookies, une partie ou la totalité de ces services risquent de ne pas fonctionner correctement.
