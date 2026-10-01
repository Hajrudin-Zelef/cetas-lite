---
id: collect-261001-rattrapage/rattrapage/netbox-guide-4
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [673, 850]
sha256: 74b56105b7b72c58fa5fc53f322533e9717bb5123bd7c3fd6ccd297ce669261c
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

- [ ] La page d'accueil s'affiche (tableau de bord, version en bas de page : 4.x)
- [ ] `Admin > Journaux` : aucune erreur critique
- [ ] Créez un **tenant** « DSI » (Organisation > Tenants) — voir section 18
- [ ] Créez une **région** « France » puis un **site** « Siège » (section 19)
- [ ] Créez un **préfixe** `10.0.0.0/8` avec le statut « Container » (section 36)
- [ ] Créez un **VLAN** de test (section 40)
- [ ] Créez un **token API** pour vous (en haut à droite > API Tokens) et testez :
  `curl -H "Authorization: Token VOTRE_TOKEN" https://netbox.lan-entreprise.fr/api/`
- [ ] Vérifiez la sauvegarde : lancez un `pg_dump` manuel (section 63) et contrôlez le fichier
- [ ] Notez l'URL, les comptes et la procédure dans votre documentation d'exploitation

> 💡 **Règle d'équipe** : personne ne travaille en `admin` au quotidien. Créez
> immédiatement un compte nominatif par administrateur (section 46).

---

## 18. Organisation : tenants et groupes de tenants

Le **tenant** (locataire) est l'entité à qui « appartient » un objet : un client, une
filiale, un département, un projet. C'est l'axe de cloisonnement n°1.

**Cas d'usage typiques :**

| Tenant | Objets rattachés |
|---|---|
| `DSI-Interne` | Baies, serveurs, VLAN et préfixes du SI interne |
| `Client-Dupont` | Équipements et IP dédiés à un client hébergé |
| `Projet-Video` | Caméras, préfixe et VLAN du projet vidéosurveillance |

Création : **Organisation > Tenants > + Ajouter**. Renseignez au minimum le nom et
le slug (généré automatiquement). Les **groupes de tenants** permettent de regrouper
(ex. groupe « Clients hébergés » contenant `Client-Dupont`, `Client-Martin`).

**Bonnes pratiques :**

- Un tenant `Infra-Mutualisee` pour le cœur (cœur de réseau, onduleurs, baies partagées).
- Nommage en `MAJUSCULES-avec-tirets`, sans accents ni espaces : `DSI`, `ATELIER`, `MAGASIN`.
- Attribuez un tenant à CHAQUE objet dès sa création (site, baie, équipement, préfixe,
  VLAN, circuit). Un objet « sans tenant » devient vite inclassable.

---

## 19. DCIM : régions, groupes de sites et sites

Hiérarchie géographique : **Région → (Groupe de sites) → Site**.

- **Région** : zone géographique large. Ex. `Europe`, `France`, `IDF`.
  Les régions peuvent être imbriquées (région parente `Europe` → enfant `France`).
- **Groupe de sites** : regroupement logique. Ex. `Agences`, `Datacenters`, `Usines`.
- **Site** : le lieu physique documenté. Ex. `SIEGE-Paris11`, `AGENCE-Lyon`, `DC-Gravelines`.

Fiche site — champs à remplir systématiquement :

| Champ | Exemple | Pourquoi c'est important |
|---|---|---|
| Nom | `SIEGE-Paris11` | identifiant unique, repris partout |
| Slug | `siege-paris11` | utilisé par l'API |
| Statut | Actif | distingue les sites fermés/en projet |
| Région / Groupe | `France` / `Bureaux` | filtres et rapports |
| Tenant | `DSI` | propriété |
| Adresse physique | `14 rue des Forges, 75011 Paris` | intervention sur site |
| Longitude/Latitude | `48.8637, 2.3792` | carte des sites |
| ASN | `64512` | si le site a son AS privé |

> 💡 Renseignez les coordonnées GPS : la vue « carte » de NetBox devient un vrai
> tableau de bord géographique de votre infrastructure.

---

## 20. DCIM : baies (racks), rôles et réservations

La **baie** (rack) est l'unité physique de base : on y installe les équipements par
unités U.

**Rôles de baie** (à créer dans DCIM > Rôles de baies) : `Baie-Reseau`, `Baie-Serveurs`,
`Baie-Brassage`, `Baie-Onduleur`. Le rôle colore la baie dans les vues et sert de filtre.

Création d'une baie : **DCIM > Baies > + Ajouter** :

| Champ | Exemple | Commentaire |
|---|---|---|
| Nom | `SIEGE-Baie-A01` | convention : `SITE-Fonction-Numéro` |
| Site | `SIEGE-Paris11` | rattachement obligatoire |
| Tenant | `DSI` | |
| Statut | Actif | |
| Rôle | `Baie-Reseau` | |
| Type | 19 pouces | |
| Largeur | 19 pouces | |
| Hauteur | `42U` | standard datacenter |
| Unités numérotées | descendant (42→1) | le U1 est en bas : standard |

**Bonnes pratiques d'exploitation :**

- Réservez les U avec la fonction **Réservations** (ex. « U20-U21 réservés pour le
  futur firewall — ticket #1234 ») : fini les installations qui se marchent dessus.
- Documentez la baie AVANT d'y visser quoi que ce soit (voir section 74, la règle d'or).
- Photographiez la baie et joignez l'image à la fiche (pièces jointes).

---

## 21. DCIM : fabricants et types d'équipements

Ne créez jamais un équipement « à la main » sans **type d'équipement** : le type décrit
le modèle physique (dimensions, interfaces, alimentations, console).

1. **Fabricant** (DCIM > Fabricants) : `Cisco`, `HPE`, `Schneider-Electric`, `Eaton`, `APC`.
2. **Type d'équipement** (DCIM > Types d'équipements) : ex. `Cisco Catalyst C9300-48P`,
   `APC Smart-UPS SRT 6000`, `Eaton 9PX 3000`.

Fiche type d'équipement — à remplir avec soin (c'est du travail mutualisé : fait une
fois, réutilisé pour chaque exemplaire) :

| Champ | Exemple |
|---|---|
| Fabricant | `Cisco` |
| Modèle | `Catalyst C9300-48P` |
| Slug | `cisco-catalyst-c9300-48p` |
| Hauteur (U) | `1` |
| Est full-depth | oui/non |
| Commentaires | `48x 1G PoE+, 4x SFP+ — alim. redondante` |

Ensuite, définissez les **composants modèles** du type (templates) : interfaces
(`GigabitEthernet1/0/1` … `1/0/48`), ports d'alimentation, console. À la création de
l'équipement réel, NetBox instancie automatiquement tous ces composants : énorme gain
de temps.

> 💡 Créez les types d'équipements pour vos onduleurs et PDU (ex.
> `Eaton 9PX 3000 RT`, `APC AP8858 PDU`) : vous pourrez câbler l'alimentation
> proprement (section 26) — le lien direct avec votre métier énergies.

---

## 22. DCIM : équipements (devices)

L'**équipement** est l'instance physique : « le switch n°3 de la baie A01 ».
**DCIM > Équipements > + Ajouter** :

| Champ | Exemple | Obligatoire ? |
|---|---|---|
| Nom | `SIEGE-SW-A01-01` | ✅ (unique par site) |
| Type d'équipement | `Cisco Catalyst C9300-48P` | ✅ |
| Rôle d'équipement | `Switch-Acces` | ✅ (à créer : `Switch-Coeur`, `Serveur`, `Onduleur`, `Firewall`…) |
| Site | `SIEGE-Paris11` | ✅ |
| Baie | `SIEGE-Baie-A01` | si monté en baie |
| Position (U) | `12` | ✅ si baie renseignée |
| Face | Avant | |
| Statut | Actif | ✅ |
| Tenant | `DSI` | recommandé |
| Numéro de série | `FOC1234ABCD` | très recommandé (SAV, garantie) |
| Asset tag (inventaire) | `INV-2024-0342` | recommandé (lien avec l'inventaire comptable) |

**Cycle de vie avec les statuts** : `En stock` → `Actif` → `En maintenance` →
`Hors service` → `Retiré`. Utilisez-les : un équipement `Hors service` reste documenté
(utile pour l'historique) sans polluer les vues d'exploitation (filtrez sur `Actif`).

---

## 23. DCIM : interfaces et bonnes pratiques de nommage

Les interfaces sont créées automatiquement depuis le type d'équipement, mais vérifiez
et complétez :

- **Nommage** : gardez le nommage constructeur (`GigabitEthernet1/0/1`, `eth0`, `eno1`).
  N'inventez pas vos propres noms : l'interface NetBox doit correspondre à ce que
  l'OS/CLI affiche, sinon le dépannage devient un cauchemar.
- **Description** : utilisez-la pour la destination : `Vers SRV-HYP-01:eth0`,
  `Uplink vers COEUR-01:Gi1/0/48`, `Imprimante Accueil`.
- **Type** : `1000BASE-T`, `10GBASE-SR`, `SFP+`… (conditionne les câbles compatibles).
- **Mode 802.1Q** : `Access` + VLAN, ou `Tagged` + VLANs (trunk), ou `Tagged All`.
- **MTU** : `1500`, ou `9000` (jumbo) si votre infra le supporte de bout en bout.
- **Activé/désactivé** : reflétez l'état réel du port (un port `shutdown` côté switch
  = désactivé dans NetBox).

**Interface management** : créez une interface dédiée (ex. `mgmt0`, `iDRAC`, `iLO`)
avec son IP : c'est elle qui portera l'adresse de supervision (lien IPAM, section 38).

---

