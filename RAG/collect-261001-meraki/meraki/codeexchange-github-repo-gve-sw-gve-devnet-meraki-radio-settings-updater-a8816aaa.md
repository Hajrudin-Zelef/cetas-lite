---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-radio-settings-updater-a8816aaa
title: "codeexchange-github-repo-gve-sw-gve-devnet-meraki-radio-settings-updater-a8816aaa"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-radio-settings-updater-a8816aaa.md
source_anchor: ""
source_lines: [1, 47]
sha256: 7b9fc54bd5cecfdef4f8acc33ee417d1310561a75f0366ab14784011e464efab
---

# codeexchange-github-repo-gve-sw-gve-devnet-meraki-radio-settings-updater-a8816aaa

This repository contains sample code for bulk updating radio settings across multiple networks.
export_rfprofiles.py will collect & export RF profiles for an existing Meraki network, which can then be copied or modified to use with the update script.
update_rfprofiles.py is capable of uploading one or many Meraki RF profiles & applying them to APs if desired. The script will determine whether the profiles are new, or existing ones that need to be updated.
git clone <repo_url>
pip install -r requirements.txt
You may choose to provide the Cisco Meraki API key via the MERAKI_DASHBOARD_API_KEY environment variable.
If the environment variable is not provided, then the script will prompt for the API key.
This step only applies for the update_rfprofiles.py script. If only using the export_rfprofiles.py code, please proceed to the Usage steps.
An example of the CSV format is below:
Network Name,RF Profiles,APs
Network 01,"Profile01, Profile02","AAAA-AAAA-AAAA,BBBB-BBBB-BBBB"
Network 02,Profile01,ALL
Network 03,"Profile01, Profile02",None
Network 04,Profile01,
Network Name must match the name of the network as shown in Meraki Dashboard.
RF Profiles can be a list of one or more RF profiles to upload to this network. The profile name must match the name field within the RF profile config, not the file name.
APs will optionally list the Meraki wireless access points to apply the new RF profiles to. Valid options include:
None. This will update/create profiles only & not apply to any APsALL will provision the assigned RF profile to all access points on the specified network
NOTE: Only 1 RF profile can be assigned to an individual AP. Therefore, the script will reject any CSV entries that include multiple profiles & AP assignments. Please create only 1 CSV entry for each set of profile-to-AP assignments.
To export RF profiles from an existing network, use the following command:
python3 export_rfprofiles.py
The script will prompt to ask which network to export profiles from & where to write the exported data to. All RF profiles exported by this script will be converted to YAML.
To begin creating / assigning RF profiles, this script will need:
Once ready, run the following command:
python3 update_rfprofiles.py
This script will prompt for the location of the CSV file & directory containing RF profiles.
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
Meraki
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
