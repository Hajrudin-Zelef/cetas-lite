---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-snapshot-collector-57b99d69
title: "codeexchange-github-repo-gve-sw-gve-devnet-meraki-snapshot-collector-57b99d69"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2024-01-31"]
keywords: ["license", "sandbox"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-snapshot-collector-57b99d69.md
source_anchor: ""
source_lines: [1, 52]
sha256: f0d3ceeb2cfdc3ab8b9a4a59bc4bb9294e252991c2b449c30515193b3ceffe8f
---

# codeexchange-github-repo-gve-sw-gve-devnet-meraki-snapshot-collector-57b99d69

This code repository performs a bulk collection of camera snapshots from Meraki MV cameras.
Using this code, we can:
git clone https://github.com/gve-sw/gve_devnet_meraki_snapshot_collector.git
pip install -r requirements.txt
The script has built-in help to provide information on the available options:
$ python3 app.py --help
 Usage: app.py [OPTIONS]
 Run bulk collection of snapshots from Meraki MV cameras
 If date & time are not specified, then script will collect snapshots for the current date  
 & time
╭─ Options ────────────────────────────────────────────────────────────────────────────────╮
│ *  --apikey  -k      TEXT                              Meraki dashboard API key          │
│                                                        [env var:                         │
│                                                        MERAKI_DASHBOARD_API_KEY]         │
│                                                        [default: None]                   │
│                                                        [required]                        │
│    --time    -t      [%Y-%m-%d|%Y-%m-%dT%H:%M:%S|      Date & time to collect snapshots  │
│                      %Y-%m-%d %H:%M:%S]                for                               │
│                                                        [default: None]                   │
│    --help                                              Show this message and exit.       │
╰──────────────────────────────────────────────────────────────────────────────────────────╯
╭─ Additional Options ─────────────────────────────────────────────────────────────────────╮
│ --outputdir   -o      TEXT  Output directory [default: snapshots]                        │
│ --outputhtml  -h            Output HTML report                                           │
╰──────────────────────────────────────────────────────────────────────────────────────────╯
The only required item is the apikey parameter. If this is not provided, the script will prompt to ask for the Meraki dashboard API key. Alternatively, we can specify the key using the MERAKI_DASHBOARD_API_KEY environment variable.
If no --time parameter is specified, the script will query the cameras for a current snapshot. Otherwise, we can specify time using any of the listed formats in the help message. If using the format with a space between date & time, the input must be wrapped in quotes. For example: python3 app.py --time "2024-01-31 14:30:00"
By default, the script will place all snapshots in a local ./snapshots/ directory. This can be overriden using the --outputdir argument.
In addition, the script can optionally export an HTML report for easier review. To enable this, use the --outputhtml flag. Once generated, the script will create a local index.html page to serve the camera snapshots.
This code can be tested with the Meraki Enterprise Sandbox.
Alternatively, the Meraki Always-On may be used - but may only have mock cameras available, which cannot generate snapshots.
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
Owner
Contributors
Categories
SecurityNetworking
Products
Meraki
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
