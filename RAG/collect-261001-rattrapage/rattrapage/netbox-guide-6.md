---
id: collect-261001-rattrapage/rattrapage/netbox-guide-6
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1050, 1262]
sha256: 8eb662972b193fcdea5e08a82d1edb5e402ba9bed36ac3d667390960f0815d69
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

- `VRF-MGMT` : plan d'administration des équipements (10.255.0.0/16), jamais routé
  vers les utilisateurs.
- `VRF-GUEST` : Wi-Fi invités, cloisonné.
- `VRF-CLIENT-X` : chevauchement d'adressage entre deux clients hébergés.

> ⚠️ Un préfixe hors VRF (VRF « Global ») et le même préfixe dans une VRF sont deux
> objets distincts : c'est voulu. En revanche, deux fois le même préfixe dans la
> MÊME VRF = doublon refusé par NetBox.

---

## 32. IPAM : RIR et agrégats (blocs d'adresses)

- **RIR** (*Regional Internet Registry*) : `RIPE-NCC` (Europe), `ARIN`… À créer dans
  IPAM > RIR. Utile si vous gérez vos propres blocs publics.
- **Agrégat** : le bloc de haut niveau que vous possédez ou qui vous est alloué.
  Ex. `193.0.2.0/24` (bloc RIPE alloué), `10.0.0.0/8` (privé, RIR « RFC-1918 »).

Créez un RIR `RFC-1918` et les trois agrégats privés une fois pour toutes :

| Agrégat | Usage conseillé |
|---|---|
| `10.0.0.0/8` | interconnexions, datacenters, gros sites |
| `172.16.0.0/12` | agences, Wi-Fi, vidéosurveillance |
| `192.168.0.0/16` | petits sites, management local, lab |

Depuis un agrégat, NetBox affiche l'**utilisation** : quels sous-préfixes sont
découpés, quel espace reste libre. C'est votre tableau de bord d'adressage.

---

## 33. IPAM : préfixes — le cœur de l'adressage

Le **préfixe** est l'objet IPAM central : `10.10.20.0/24`, `2001:db8:1::/64`.
**IPAM > Préfixes > + Ajouter**.

| Champ | Exemple | Commentaire |
|---|---|---|
| Préfixe | `10.10.20.0/24` | notation CIDR stricte |
| VRF | `VRF-PROD` ou Global | |
| Site | `SIEGE-Paris11` | rattachement géographique |
| VLAN | `VLAN-20` | le VLAN qui porte ce sous-réseau |
| Statut | `Actif` | voir ci-dessous |
| Rôle | `LAN-Utilisateurs` | à créer : `Serveurs`, `Management`, `WiFi`… |
| Tenant | `DSI` | |
| Description | `LAN utilisateurs — bâtiment A, étage 2` | |

**Statuts de préfixe** (à utiliser avec rigueur) :

| Statut | Signification |
|---|---|
| `Container` | bloc parent découpé en sous-préfixes (ex. `10.10.0.0/16`) |
| `Actif` | en production |
| `Réservé` | attribué à un projet, pas encore déployé |
| `Déprécié` | en cours de retrait, ne plus l'utiliser |

> 💡 **Convention** : tout bloc `/16` ou plus grand est `Container`. On ne met JAMAIS
> d'adresses IP directement dans un `Container` : on découpe d'abord en préfixes enfants.

---

## 34. IPAM : découper et visualiser un plan d'adressage

Depuis la fiche d'un préfixe `Container`, le bouton **+ Préfixes enfants** permet de
découper : NetBox propose les sous-préfixes libres.

Exemple : `10.10.0.0/16` (Container, site Siège) découpé en :

| Préfixe | Rôle | VLAN | Usage |
|---|---|---|---|
| `10.10.10.0/24` | `Serveurs` | 10 | serveurs physiques et VM |
| `10.10.20.0/24` | `LAN-Utilisateurs` | 20 | postes, étage 2 |
| `10.10.30.0/24` | `WiFi` | 30 | Wi-Fi corporate |
| `10.10.40.0/24` | `Impression` | 40 | imprimantes/MFP |
| `10.10.254.0/24` | `Management` | 254 | iDRAC, switchs, onduleurs (cartes réseau) |

La vue du préfixe affiche une **barre d'utilisation** : % d'IP utilisées, libres,
et la liste des adresses. C'est l'outil n°1 pour répondre à « il reste des IP sur
le VLAN 20 ? » en 10 secondes.

**Plan d'adressage type pour un site moyen** (à adapter) :

- `.0` réseau, `.1` passerelle (VRRP/HSRP : `.2`/`.3` physiques, `.1` virtuelle),
- `.4`–`.19` : équipements réseau (switchs, AP),
- `.20`–`.49` : serveurs / hyperviseurs,
- `.50`–`.199` : DHCP dynamique,
- `.200`–`.249` : équipements fixes (imprimantes, caméras),
- `.250`–`.254` : management (iDRAC/iLO, cartes onduleurs, PDU).

---

## 35. IPAM : plages IP (IP ranges)

Une **plage IP** réserve un intervalle à l'intérieur d'un préfixe sans créer chaque
adresse : parfait pour les pools DHCP, les VIP, les plages d'impression.

**IPAM > Plages IP > + Ajouter** :

| Champ | Exemple |
|---|---|
| Plage | `10.10.20.50` → `10.10.20.199` |
| VRF / préfixe parent | déduit automatiquement |
| Statut | `Actif` |
| Rôle | `DHCP` |
| Description | `Pool DHCP — LAN utilisateurs étage 2` |

Cas d'usage :

- `DHCP` : pools des serveurs DHCP (à croiser avec la conf. réelle du serveur !),
- `VIP` : adresses virtuelles (VRRP, load-balancers),
- `Réservé` : plage gelée pour un projet futur.

> 💡 Une plage `DHCP` documentée + un export API vers votre serveur DHCP (section 56)
> = des pools toujours synchronisés avec la documentation. C'est l'automatisation
> qui commence.

---

## 36. IPAM : adresses IP et assignation aux interfaces

L'**adresse IP** est l'objet le plus créé au quotidien : `10.10.10.15/24`.
**IPAM > Adresses IP > + Ajouter**, ou directement depuis une interface
(**Assigner une IP**).

| Champ | Exemple |
|---|---|
| Adresse | `10.10.10.15/24` (le masque est obligatoire) |
| VRF | héritée du préfixe |
| Statut | `Actif` (`Réservé`, `DHCP`) |
| Rôle | `VIP`, `Loopback`, `Anycast`… |
| DNS | `srv-fichiers-01.lan-entreprise.fr` |
| Description | `Interface de production` |

**Assignation** : depuis la fiche IP > **Assigner** > choisir l'objet :
interface d'équipement (`SIEGE-SW-A01-01:mgmt0`), interface de VM, ou plage.
Une IP assignée apparaît sur la fiche de l'interface et inversement : la liaison
IPAM ↔ DCIM est faite.

**Règles d'hygiène :**

- Une IP = un masque correct (celui du préfixe). NetBox alerte si l'IP n'appartient
  à aucun préfixe connu (« hors préfixe » = anomalie à traiter).
- Statut `DHCP` pour les adresses dynamiques connues (réservations), `Réservé`
  pour les futures attributions.
- Renseignez le **DNS** : NetBox devient aussi votre aide-mémoire de nommage.

---

## 37. IPAM : rôles de préfixes et d'adresses

Les **rôles** qualifient l'usage. Créez votre référentiel une fois (IPAM > Rôles),
réutilisez partout :

| Rôle | Couleur suggérée | Usage |
|---|---|---|
| `Loopback` | gris | /32 des routeurs |
| `Interconnexion` | orange | /30, /31 entre équipements |
| `Serveurs` | bleu | sous-réseaux serveurs |
| `LAN-Utilisateurs` | vert | postes |
| `WiFi` | violet | SSID corporate/invités |
| `Management` | rouge | administration (OOB) |
| `Impression` | jaune | imprimantes/MFP |
| `Videosurveillance` | noir | caméras |
| `DMZ` | rouge foncé | zone exposée |
| `Secours-4G` | cyan | liens de backup |

Les couleurs apparaissent dans les barres d'utilisation : d'un coup d'œil, on voit
la nature de chaque préfixe. **Ne multipliez pas les rôles** : 10–12 rôles stables
valent mieux que 40 rôles approximatifs.

---

## 38. VLAN : groupes de VLAN et VLAN

Hiérarchie : **Groupe de VLAN → VLAN**. Le groupe porte la plage d'ID
(ex. `100–199 : Serveurs`), le VLAN est l'instance (`VLAN 110 : SRV-Production`).

**IPAM > Groupes de VLAN > + Ajouter** : nom (`VLAN-Siege`), site, plage d'ID.
**IPAM > VLAN > + Ajouter** :

| Champ | Exemple |
|---|---|
| Nom | `SRV-Production` |
| ID (VID) | `110` |
| Groupe | `VLAN-Siege` |
| Site | `SIEGE-Paris11` |
| Statut / Rôle | `Actif` / `Serveurs` |
| Tenant | `DSI` |

**Liaison VLAN ↔ préfixe** : depuis le préfixe, renseignez le VLAN (et inversement).
NetBox vérifie la cohérence : un préfixe `10.10.10.0/24` marqué VLAN 110 dont les
interfaces en mode Access portent le VLAN 110 = documentation cohérente.

> ⚠️ Les VLAN 1, 1002–1005 (Cisco) et les plages réservées : documentez-les comme
> `Réservé` pour éviter qu'un collègue ne les réutilise par mégarde.

---

## 39. Services (ports TCP/UDP par équipement/VM)

Le module **IPAM > Services** documente « quoi écoute où » : utile pour les règles
de pare-feu et le dépannage.

Exemple sur `VM-NETBOX-01` :

| Nom | Protocole | Ports | Description |
|---|---|---|---|
| `https` | TCP | `443` | interface web NetBox |
| `ssh` | TCP | `22` | administration |

Exemple sur un serveur de fichiers : `smb` TCP 445, sur un onduleur avec carte
réseau : `https` TCP 443 + `snmp` UDP 161.

