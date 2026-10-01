---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-4
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [352, 423]
sha256: a0eee6114efbcb76006609e2dc56a5732546eae4b7fa8df20f11cd7bbc82ac17
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
- Déployez la configuration. Le service apparaît dans la liste des services supervisés, et dans la page Statut des ressources. La commande envoyée par le connecteur est indiquée dans le panneau de détails du service : celle-ci montre les valeurs des macros.
Comment puis-je tester le plugin et que signifient les options des commandes ?
Une fois le plugin installé, vous pouvez tester celui-ci directement en ligne
de commande depuis votre collecteur Centreon en vous connectant avec
l'utilisateur centreon-engine (su - centreon-engine). Vous pouvez tester
que le connecteur arrive bien à superviser une ressource en utilisant une commande
telle que celle-ci (remplacez les valeurs d'exemple par les vôtres) :
/usr/lib/centreon/plugins/centreon_cisco_meraki_restapi.pl \
  --plugin='network::cisco::meraki::cloudcontroller::restapi::plugin' \
  --mode='devices' \
  --hostname='api.meraki.com' \
  --api-token='12345abcd6789efgh0123abcd4567efgh8901abcd' \
  --proxyurl='http://proxy.mycompany:8080' \
  --filter-device-name='^.*$' \
  --verbose
La commande devrait retourner un message de sortie similaire à :
OK: Device 'centreon-par-training-ap' status: online - connection success: 0 - traffic in: 51.66 b/s, out: 515.86 b/s - link 'WAN 1' status: active | 
'devices.total.online.count'=0;;;0;1 'devices.total.offline.count'=0;;;0;1 'devices.total.alerting.count'=0;;;0;1 
'centreon-par-training-ap#device.connections.success.count'=0;;;0; 'centreon-par-training-ap#device.connections.auth.count'=0;;;0; 
'centreon-par-training-ap#device.connections.assoc.count'=0;;;0; 'centreon-par-training-ap#device.connections.dhcp.count'=0;;;0; 
'centreon-par-training-ap#device.connections.dns.count'=0;;;0; 'centreon-par-training-ap#device.traffic.in.bitspersecond'=51.6626907073509b/s;;;0; 
'centreon-par-training-ap#device.traffic.out.bitspersecond'=515.864632454924b/s;;;0;
checking device 'centreon-par-training-ap'
    status: online
    connection success: 0
    traffic in: 51.66 b/s, out: 515.86 b/s
    link 'WAN 1' status: active
Diagnostic des erreurs communes
Rendez-vous sur la documentation dédiée des plugins basés sur HTTP/API.
Modes disponibles
Dans la plupart des cas, un mode correspond à un modèle de service. Le mode est renseigné dans la commande d'exécution du connecteur. Dans l'interface de Centreon, il n'est pas nécessaire de les spécifier explicitement, leur utilisation est implicite dès lors que vous utilisez un modèle de service. En revanche, vous devrez spécifier le mode correspondant à ce modèle si vous voulez tester la commande d'exécution du connecteur dans votre terminal.
Tous les modes disponibles peuvent être affichés en ajoutant le paramètre
--list-mode à la commande :
/usr/lib/centreon/plugins/centreon_cisco_meraki_restapi.pl \
	--plugin=network::cisco::meraki::cloudcontroller::restapi::plugin \
	--list-mode
Le plugin apporte les modes suivants :
| Mode | Modèle de service associé | 
|---|---|
| api-requests [code] | Net-Cisco-Meraki-Cloudcontroller-Api-Requests-Restapi-custom | 
| cache [code] | Net-Cisco-Meraki-Cloudcontroller-Cache-Restapi-custom | 
| devices [code] | Net-Cisco-Meraki-Cloudcontroller-Device-Restapi-custom Net-Cisco-Meraki-Cloudcontroller-Devices-Restapi-custom | 
| discovery [code] | Utilisé pour la découverte d'hôtes | 
| list-devices [code] | Utilisé pour la découverte de services | 
| list-tags [code] | Utilisé pour la découverte de services | 
| list-vpn-tunnels [code] | Utilisé pour la découverte de services | 
| networks [code] | Net-Cisco-Meraki-Cloudcontroller-Network-Restapi-custom Net-Cisco-Meraki-Cloudcontroller-Networks-Restapi-custom | 
| vpn-tunnels [code] | Net-Cisco-Meraki-Cloudcontroller-Vpn-Tunnels-Restapi-custom | 
Options disponibles
Options génériques
Les options génériques sont listées ci-dessous :
| Option | Description | 
|---|---|
| --mode | Define the mode in which you want the plugin to be executed (see --list-mode). | 
| --dyn-mode | Specify a mode with the module's path (advanced). | 
| --list-mode | List all available modes. | 
| --mode-version | Check minimal version of mode. If not, unknown error. | 
| --version | Return the version of the plugin. | 
| --custommode | When a plugin offers several ways (CLI, library, etc.) to get information the desired one must be defined with this option. | 
| --list-custommode | List all available custom modes. | 
| --multiple | Multiple custom mode objects. This may be required by some specific modes (advanced). | 
| --pass-manager | Define the password manager you want to use. Supported managers are: environment, file, keepass, hashicorpvault and teampass. | 
| --verbose | Display extended status information (long output). | 
| --debug | Display debug messages. | 
| --show-password | By default, sensitive information in command lines is hidden in debug output and replaced with *** (however, debug logs may still display sensitive information). Using the C option will display the passwords in plain text. | 
| --filter-perfdata | Filter perfdata that match the regexp. Example: adding --filter-perfdata='avg' will remove all metrics that do not contain 'avg' from performance data. | 
| --filter-perfdata-adv | Filter perfdata based on a "if" condition using the following variables: label, value, unit, warning, critical, min, max. Variables must be written either %{variable} or %(variable). Example: adding --filter-perfdata-adv='not (%(value) == 0 and %(max) eq "")' will remove all metrics whose value equals 0 and that don't have a maximum value. | 
| --explode-perfdata-max | Create a new metric for each metric that comes with a maximum limit. The new metric will be named identically with a '_max' suffix. Example: it will split 'used_prct'=26.93%;0:80;0:90;0;100 into 'used_prct'=26.93%;0:80;0:90;0;100 'used_prct_max'=100%;;;; | 
| --change-perfdata --extend-perfdata | Change or extend perfdata. Syntax: --extend-perfdata=searchlabel,newlabel,target[,[<new-unit-of-mesure>],[min],[max]] Common examples: onvert storage free perfdata into used: --change-perfdata='free,used,invert()' Convert storage free perfdata into used: --change-perfdata='used,free,invert()' Scale traffic values automatically: --change-perfdata='traffic,,scale(auto)' Scale traffic values in Mbps: --change-perfdata='traffic_in,,scale(Mbps),mbps' Change traffic values in percent: --change-perfdata='traffic_in,,percent()' =back | 
| --change-perfdata | Change or extend perfdata. Syntax: --extend-perfdata=searchlabel,newlabel,target[,[<new-unit-of-mesure>],[min],[max]] Common examples: onvert storage free perfdata into used: --change-perfdata='free,used,invert()' Convert storage free perfdata into used: --change-perfdata='used,free,invert()' Scale traffic values automatically: --change-perfdata='traffic,,scale(auto)' Scale traffic values in Mbps: --change-perfdata='traffic_in,,scale(Mbps),mbps' Change traffic values in percent: --change-perfdata='traffic_in,,percent()' =back | 
| --extend-perfdata | Change or extend perfdata. Syntax: --extend-perfdata=searchlabel,newlabel,target[,[<new-unit-of-mesure>],[min],[max]] Common examples: onvert storage free perfdata into used: --change-perfdata='free,used,invert()' Convert storage free perfdata into used: --change-perfdata='used,free,invert()' Scale traffic values automatically: --change-perfdata='traffic,,scale(auto)' Scale traffic values in Mbps: --change-perfdata='traffic_in,,scale(Mbps),mbps' Change traffic values in percent: --change-perfdata='traffic_in,,percent()' =back | 
