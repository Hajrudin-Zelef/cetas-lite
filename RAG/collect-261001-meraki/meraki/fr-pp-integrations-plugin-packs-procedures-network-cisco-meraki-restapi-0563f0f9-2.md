---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-2
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: ["China", "Oracle", "United States"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [126, 204]
sha256: 278217e0a4d8a07fff2fe9fa4ba84011b779072d3c9a9a0263f6b4790ecf1abe
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

- Quel que soit le type de la licence (online ou offline), installez le connecteur Cisco Meraki Rest API depuis l'interface web et le menu Configuration > Connecteurs > Connecteurs de supervision.
Plugin
À partir de Centreon 22.04, il est possible de demander le déploiement automatique du plugin lors de l'utilisation d'un connecteur. Si cette fonctionnalité est activée, et que vous ne souhaitez pas découvrir des éléments pour la première fois, alors cette étape n'est pas requise.
Plus d'informations dans la section Installer le plugin.
Utilisez les commandes ci-dessous en fonction du gestionnaire de paquets de votre système d'exploitation :
- Alma / RHEL / Oracle Linux 8
- Alma / RHEL / Oracle Linux 9
- Debian 11 & 12
- CentOS 7
dnf install centreon-plugin-Network-Cisco-Meraki-Restapi
apt install centreon-plugin-network-cisco-meraki-restapi
yum install centreon-plugin-Network-Cisco-Meraki-Restapi
Utiliser le connecteur de supervision
Utiliser un modèle d'hôte issu du connecteur
- Net-Cisco-Meraki-Cloudcontroller-Restapi-custom
- Net-Cisco-Meraki-Device-Restapi-custom
- Net-Cisco-Meraki-Network-Restapi-custom
- Ajoutez un hôte à Centreon depuis la page Configuration > Hôtes.
- Complétez les champs Nom, Alias & IP Address/DNS correspondant à votre ressource.
- Appliquez le modèle d'hôte Net-Cisco-Meraki-Cloudcontroller-Restapi-custom. Une liste de macros apparaît. Les macros vous permettent de définir comment le connecteur se connectera à la ressource, ainsi que de personnaliser le comportement du connecteur.
- Renseignez les macros désirées. Attention, certaines macros sont obligatoires.
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| MERAKIAPIHOSTNAME | Meraki API hostname The default value 'api.meraki.com' will work for most of the world. However, for organizations hosted in the following country dashboards, you need to override this value and specify the respective base URI instead: Canada: https://api.meraki.ca/api/v1 China: https://api.meraki.cn/api/v1 India: https://api.meraki.in/api/v1 United States FedRAMP: https://api.gov-meraki.com/api/v1 Please refer to Meraki API documentation https://developer.cisco.com/meraki/api-v1/getting-started/#base-uri for more details | api.meraki.com | X | 
| MERAKIAPITOKEN | Meraki API token |  | X | 
| MERAKIAPIPROTO | Define the protocol to reach the API | https |  | 
| MERAKIAPIPORT | Define the TCP port to use to reach the API | 443 |  | 
| PROXYURL | Proxy URL. Example: http://my.proxy:3128 |  |  | 
| MERAKIAPIEXTRAOPTIONS | Any extra option you may want to add to every command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
- Déployez la configuration. L'hôte apparaît dans la liste des hôtes supervisés, et dans la page Statut des ressources. La commande envoyée par le connecteur est indiquée dans le panneau de détails de l'hôte : celle-ci montre les valeurs des macros.
- Ajoutez un hôte à Centreon depuis la page Configuration > Hôtes.
- Complétez les champs Nom, Alias & IP Address/DNS correspondant à votre ressource.
- Appliquez le modèle d'hôte Net-Cisco-Meraki-Device-Restapi-custom. Une liste de macros apparaît. Les macros vous permettent de définir comment le connecteur se connectera à la ressource, ainsi que de personnaliser le comportement du connecteur.
- Renseignez les macros désirées. Attention, certaines macros sont obligatoires.
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| MERAKIAPIHOSTNAME | Meraki API hostname The default value 'api.meraki.com' will work for most of the world. However, for organizations hosted in the following country dashboards, you need to override this value and specify the respective base URI instead: Canada: https://api.meraki.ca/api/v1 China: https://api.meraki.cn/api/v1 India: https://api.meraki.in/api/v1 United States FedRAMP: https://api.gov-meraki.com/api/v1 Please refer to Meraki API documentation https://developer.cisco.com/meraki/api-v1/getting-started/#base-uri for more details | api.meraki.com | X | 
| MERAKIAPITOKEN | Meraki API token |  | X | 
| MERAKIAPIPROTO | Define the protocol to reach the API | https |  | 
| MERAKIAPIPORT | Define the TCP port to use to reach the API | 443 |  | 
| MERAKIDEVICENAME | Filter devices by name (can be a regexp) |  |  | 
| PROXYURL | Proxy URL. Example: http://my.proxy:3128 |  |  | 
| MERAKIAPIEXTRAOPTIONS | Any extra option you may want to add to every command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
- Déployez la configuration. L'hôte apparaît dans la liste des hôtes supervisés, et dans la page Statut des ressources. La commande envoyée par le connecteur est indiquée dans le panneau de détails de l'hôte : celle-ci montre les valeurs des macros.
- Ajoutez un hôte à Centreon depuis la page Configuration > Hôtes.
- Complétez les champs Nom, Alias & IP Address/DNS correspondant à votre ressource.
- Appliquez le modèle d'hôte Net-Cisco-Meraki-Network-Restapi-custom. Une liste de macros apparaît. Les macros vous permettent de définir comment le connecteur se connectera à la ressource, ainsi que de personnaliser le comportement du connecteur.
- Renseignez les macros désirées. Attention, certaines macros sont obligatoires.
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| MERAKIAPIHOSTNAME | Meraki API hostname The default value 'api.meraki.com' will work for most of the world. However, for organizations hosted in the following country dashboards, you need to override this value and specify the respective base URI instead: Canada: https://api.meraki.ca/api/v1 China: https://api.meraki.cn/api/v1 India: https://api.meraki.in/api/v1 United States FedRAMP: https://api.gov-meraki.com/api/v1 Please refer to Meraki API documentation https://developer.cisco.com/meraki/api-v1/getting-started/#base-uri for more details | api.meraki.com | X | 
| MERAKIAPITOKEN | Meraki API token |  | X | 
| MERAKIAPIPROTO | Define the protocol to reach the API | https |  | 
| MERAKIAPIPORT | Define the TCP port to use to reach the API | 443 |  | 
| MERAKINETWORKNAME | Filter network name (can be a regexp) |  |  | 
| PROXYURL | Proxy URL. Example: http://my.proxy:3128 |  |  | 
| MERAKIAPIEXTRAOPTIONS | Any extra option you may want to add to every command (a --verbose flag for example). Toutes les options sont listées ici. |  |  | 
- Déployez la configuration. L'hôte apparaît dans la liste des hôtes supervisés, et dans la page Statut des ressources. La commande envoyée par le connecteur est indiquée dans le panneau de détails de l'hôte : celle-ci montre les valeurs des macros.
Utiliser un modèle de service issu du connecteur
- Si vous avez utilisé un modèle d'hôte et coché la case Créer aussi les services liés aux modèles, les services associés au modèle ont été créés automatiquement, avec les modèles de services correspondants. Sinon, créez les services désirés manuellement et appliquez-leur un modèle de service.
- Renseignez les macros désirées (par exemple, ajustez les seuils d'alerte). Les macros indiquées ci-dessous comme requises (Obligatoire) doivent être renseignées.
- Api-Requests
- Cache
- Device
- Devices
- Network
- Networks
- Vpn-Tunnels
| Macro | Description | Valeur par défaut | Obligatoire | 
|---|---|---|---|
| FILTERORGANIZATIONNAME | Filter organization name (can be a regexp) |  |  | 
| WARNINGAPIREQUESTS200 | Threshold |  |  | 
| CRITICALAPIREQUESTS200 | Threshold |  |  | 
| WARNINGAPIREQUESTS404 | Threshold |  |  | 
| CRITICALAPIREQUESTS404 | Threshold |  |  | 
| WARNINGAPIREQUESTS429 | Threshold |  |  | 
| CRITICALAPIREQUESTS429 | Threshold |  |  | 
| EXTRAOPTIONS | Any extra option you may want to add to the command (a --verbose flag for example). Toutes les options sont listées ici. | --verbose |  | 
| Macro | Description | Valeur par défaut | Obligatoire | 
