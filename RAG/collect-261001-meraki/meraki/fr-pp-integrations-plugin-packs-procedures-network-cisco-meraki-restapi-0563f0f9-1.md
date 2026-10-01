---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-1
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: ["Oracle"]
dates: []
keywords: ["distribution", "latency"]
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [1, 125]
sha256: 5fbdbdffbe58e41b389f7e9119a86f9396f389918f10185a48c8fa2e50f20777
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

Cisco Meraki Rest API
Dépendances du connecteur de supervision
Les connecteurs de supervision suivants sont automatiquement installés lors de l'installation du connecteur Cisco Meraki Rest API depuis la page Configuration > Connecteurs > Connecteurs de supervision :
Contenu du pack
Modèles
Le connecteur de supervision Cisco Meraki Rest API apporte 3 modèles d'hôte :
- Net-Cisco-Meraki-Cloudcontroller-Restapi-custom
- Net-Cisco-Meraki-Device-Restapi-custom
- Net-Cisco-Meraki-Network-Restapi-custom
Le connecteur apporte les modèles de service suivants (classés selon le modèle d'hôte auquel ils sont rattachés) :
- Net-Cisco-Meraki-Cloudcontroller-Restapi-custom
- Net-Cisco-Meraki-Device-Restapi-custom
- Net-Cisco-Meraki-Network-Restapi-custom
- Non rattachés à un modèle d'hôte
| Alias | Modèle de service | Description | 
|---|---|---|
| Api-Requests | Net-Cisco-Meraki-Cloudcontroller-Api-Requests-Restapi-custom | Contrôle l'utilisation de l'API Cisco Meraki | 
Les services listés ci-dessus sont créés automatiquement lorsque le modèle d'hôte Net-Cisco-Meraki-Cloudcontroller-Restapi-custom est utilisé.
| Alias | Modèle de service | Description | 
|---|---|---|
| Device | Net-Cisco-Meraki-Cloudcontroller-Device-Restapi-custom | Contrôle l'utilisation des équipements | 
Les services listés ci-dessus sont créés automatiquement lorsque le modèle d'hôte Net-Cisco-Meraki-Device-Restapi-custom est utilisé.
| Alias | Modèle de service | Description | 
|---|---|---|
| Network | Net-Cisco-Meraki-Cloudcontroller-Network-Restapi-custom | Contrôle l'utilisation des réseaux | 
Les services listés ci-dessus sont créés automatiquement lorsque le modèle d'hôte Net-Cisco-Meraki-Network-Restapi-custom est utilisé.
| Alias | Modèle de service | Description | Découverte | 
|---|---|---|---|
| Cache | Net-Cisco-Meraki-Cloudcontroller-Cache-Restapi-custom | Service permettant de générer les fichiers de cache |  | 
| Devices | Net-Cisco-Meraki-Cloudcontroller-Devices-Restapi-custom | Contrôle l'utilisation des équipements | X | 
| Networks | Net-Cisco-Meraki-Cloudcontroller-Networks-Restapi-custom | Contrôle l'utilisation des réseaux |  | 
| Vpn-Tunnels | Net-Cisco-Meraki-Cloudcontroller-Vpn-Tunnels-Restapi-custom | Contrôle l'utilisation des tunnels VPN | X | 
Les services listés ci-dessus ne sont pas créés automatiquement lorsqu'un modèle d'hôte est appliqué. Pour les utiliser, créez un service manuellement et appliquez le modèle de service souhaité.
Si la case Découverte est cochée, cela signifie qu'une règle de découverte de service existe pour ce service.
Règles de découverte
Découverte d'hôtes
| Nom de la règle | Description | 
|---|---|
| Cisco Meraki Devices | Découvre les appareils Cisco Meraki via RestAPI | 
| Cisco Meraki Networks | Découvre les réseaux Cisco Meraki via RestAPI | 
Rendez-vous sur la documentation dédiée pour en savoir plus sur la découverte automatique d'hôtes.
Découverte de services
| Nom de la règle | Description | 
|---|---|
| Net-Cisco-Meraki-Cloudcontroller-Restapi-Vpn-Tunnels-Network-Name | Découvre les tunnels VPN en se basant sur les noms des réseaux | 
| Net-Cisco-Meraki-RestAPI-Device | Découvre les équipements et supervise le statut | 
| Net-Cisco-Meraki-RestAPI-Tag | Découvre les tags | 
Rendez-vous sur la documentation dédiée pour en savoir plus sur la découverte automatique de services et sa planification.
Métriques & statuts collectés
Voici le tableau des services pour ce connecteur, détaillant les métriques et statuts rattachés à chaque service.
- Api-Requests
- Cache
- Device
- Devices
- Network
- Networks
- Vpn-Tunnels
| Nom | Unité | 
|---|---|
| organizations#organization.api.requests.200.count | count | 
| organizations#organization.api.requests.404.count | count | 
| organizations#organization.api.requests.429.count | count | 
Pas de métrique pour ce service.
| Nom | Unité | 
|---|---|
| devices.total.online.count | count | 
| devices.total.online.percentage | % | 
| devices.total.offline.count | count | 
| devices.total.offline.percentage | % | 
| devices.total.alerting.count | count | 
| status | N/A | 
| devices~device.load.count | count | 
| devices~device.connections.success.count | count | 
| devices~device.connections.auth.count | count | 
| devices~device.connections.assoc.count | count | 
| devices~device.connections.dhcp.count | count | 
| devices~device.connections.dns.count | count | 
| devices~device.traffic.in.bitspersecond | b/s | 
| devices~device.traffic.out.bitspersecond | b/s | 
| devices~device.links.ineffective.count | count | 
| link-status | N/A | 
| devices~device_links#device.link.latency.milliseconds | ms | 
| devices~device_links#device.link.loss.percentage | % | 
| port-status | N/A | 
| devices~device_ports#device.port.traffic.in.bitspersecond | b/s | 
| devices~device_ports#device.port.traffic.out.bitspersecond | b/s | 
| Nom | Unité | 
|---|---|
| networks#network.connections.success.count | count | 
| networks#network.connections.auth.count | count | 
| networks#network.connections.assoc.count | count | 
| networks#network.connections.dhcp.count | count | 
| networks#network.connections.dns.count | count | 
| networks#network.traffic.in.bitspersecond | b/s | 
| networks#network.traffic.out.bitspersecond | b/s | 
| Nom | Unité | 
|---|---|
| vpn.tunnels.unreachable.count | count | 
| device-status | N/A | 
| vpn-status | N/A | 
Prérequis
Plus d'informations à propos de l'API de Cisco Meraki sont disponibles sur la documentation officielle : https://documentation.meraki.com/zGeneral_Administration/Other_Topics/The_Cisco_Meraki_Dashboard_API
Afin de pouvoir utiliser l'API Cisco Meraki, activez tout d'abord celle-ci pour votre organisation sur le portail Cisco Meraki à l'aide du menu Organization > Settings > Dashboard API access.
Une fois l'API activée, allez dans le menu my profile pour générer une API Key. Celle-ci sera associé à votre compte administrateur Cisco Meraki Dashboard.
Vous pouvez générer, révoquer et regénérer une API Key pour votre profil.
Sauvegardez votre API Key en lieu sûr puisqu'elle contient des informations d'authentification pour toute votre organisation. Il est possible de regénérer l'API Key à tout moment, cela révoquera la clé existante.
Informations sur l'URI de base
La valeur par défaut de l'URI de base api.meraki.com utilisée dans la macro MERAKIAPIHOSTNAME est correcte pour la plupart des régions.
Cependant, pour les organisations hébergées dans les pays suivants, il faut remplacer cette valeur par l'URI correspondante :
- Canada: https://api.meraki.ca/api/v1
- Chine: https://api.meraki.cn/api/v1
- Inde: https://api.meraki.in/api/v1
- États-Unis (FedRAMP): https://api.gov-meraki.com/api/v1
L'utilisation d'une mauvaise URI peut empêcher ce connecteur de fonctionner correctement. Veuillez consulter la documentation de l'API Meraki à l'adresse https://developer.cisco.com/meraki/api-v1/getting-started/#base-uri pour plus de détails.
Installer le connecteur de supervision
Pack
La procédure d'installation des connecteurs de supervision diffère légèrement suivant que votre licence est offline ou online.
- Si la plateforme est configurée avec une licence online, l'installation d'un paquet n'est pas requise pour voir apparaître le connecteur dans le menu Configuration > Connecteurs > Connecteurs de supervision. Au contraire, si la plateforme utilise une licence offline, installez le paquet sur le serveur central via la commande correspondant au gestionnaire de paquets associé à sa distribution :
- Alma / RHEL / Oracle Linux 8
- Alma / RHEL / Oracle Linux 9
- Debian 11 & 12
- CentOS 7
dnf install centreon-pack-network-cisco-meraki-restapi
apt install centreon-pack-network-cisco-meraki-restapi
yum install centreon-pack-network-cisco-meraki-restapi
