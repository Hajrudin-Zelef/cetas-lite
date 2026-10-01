---
id: collect-261001-rattrapage/rattrapage/netbox-guide-5
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [851, 1049]
sha256: 81b3726646e5b89524943d3dfb4d33e83937948599c1be5a32aa1057e522673f
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

## 24. DCIM : câblage — câbles et connexions

NetBox modélise les liaisons physiques entre interfaces/ports via des **câbles**.
Deux interfaces reliées = un câble entre elles (ou une chaîne via des panneaux de brassage).

Création : depuis la fiche d'une interface > **Connecter**, ou **DCIM > Câbles**.

| Champ | Exemple |
|---|---|
| Type | `Cat6a`, `Fibre OM4`, `DAC` |
| Longueur / unité | `3`, `mètres` |
| Couleur | `Jaune` (jarretières), `Bleu`… |
| Étiquette | `JAR-A01-014` |

**Code couleur conseillé** (à afficher dans le local technique) :

| Couleur | Usage |
|---|---|
| Bleu | brassage cuivre standard |
| Jaune | uplinks / inter-baies |
| Rouge | liens critiques / cœur |
| Vert | fibre optique |
| Noir | alimentation (ne jamais mélanger avec les jarretières data) |

> ⚠️ Câblez dans NetBox **avant ou pendant** l'intervention, jamais « quand on aura
> le temps ». Un câblage non documenté est un câblage perdu : la section 74 en fait
> une règle d'équipe.

---

## 25. DCIM : chemins de câbles (cable paths)

Quand un lien traverse plusieurs segments (switch → panneau de brassage → panneau
distant → serveur), NetBox calcule automatiquement le **chemin de câble** de bout en bout.

Chaîne type :

```
SW-A01-01:Gi1/0/1 --(jarretière)--> PB-A01:Panneau-24p:Port-01
   --(rocade)--> PB-B02:Panneau-24p:Port-01 --(jarretière)--> SRV-02:eth0
```

Intérêts concrets :

- **Dépannage** : « le port 12 du panneau B02 ne répond plus » → le chemin montre
  tous les segments à tester, dans l'ordre.
- **Changement de jarretière** : on sait exactement quoi débrancher sans couper
  le voisin.
- **Audit** : exportez les chemins pour vérifier la redondance (deux chemins
  physiquement disjoints pour un lien critique ?).

Pour que les chemins fonctionnent, modélisez les **panneaux de brassage** comme des
équipements (rôle `Panneau-Brassage`, type passif) avec leurs ports avant/arrière,
puis câblez chaque segment. C'est du travail initial, mais c'est ce qui transforme
NetBox en véritable jumeau numérique de votre câblage.

---

## 26. DCIM : alimentation — PDU, onduleurs et lien avec le métier énergies

C'est la section qui relie NetBox à votre cœur de métier. NetBox modélise la chaîne
électrique complète : **onduleur → PDU → alimentation de l'équipement**.

### 26.1. Modéliser l'onduleur et les PDU

1. Créez le type d'équipement onduleur (ex. `Eaton 9PX 3000 RT`, 2U) avec ses
   **prises de sortie** (composants « power outlets ») : `Sortie-1` … `Sortie-8`.
2. Créez l'équipement onduleur dans la baie (rôle `Onduleur`), ex. `SIEGE-UPS-A01`.
3. Créez les PDU (rôle `PDU`, type `APC AP8858` 0U ou 1U) avec leurs prises :
   ex. `SIEGE-PDU-A01-A` et `SIEGE-PDU-A01-B` (double alimentation A/B).
4. Câblez : `UPS:Sortie-1` → `PDU-A:Entrée`, `UPS:Sortie-2` → `PDU-B:Entrée`
   (connexions d'alimentation, comme des câbles data mais entre ports électriques).

### 26.2. Raccorder les équipements

Pour chaque serveur/switch à double alimentation :

```
SRV-01:PSU1 --câble C13--> PDU-A:Prise-05
SRV-01:PSU2 --câble C13--> PDU-B:Prise-05
```

Bénéfices immédiats :

- **Plan de délestage** : en cas de maintenance onduleur, la liste des équipements
  branchés sur chaque PDU est exacte — fini les « on coupe et on verra bien ».
- **Équilibrage de charge** : visualisez la répartition A/B prise par prise.
- **Capacité** : additionnez les puissances (champ « puissance max » des alimentations)
  par PDU et comparez au calibre du disjoncteur amont.

### 26.3. Champs utiles à renseigner (onglet Alimentation)

| Champ | Exemple | Usage |
|---|---|---|
| Tension d'alimentation | `230 V` | dimensionnement |
| Puissance max (W) | `800` | bilan de puissance par baie |
| Type de prise | `C13/C14`, `C19/C20` | commande des cordons |

> 💡 **Idée d'exploitation** : créez un **rapport personnalisé** (section 59) « Bilan
> de puissance par baie » qui additionne les puissances des équipements et alerte
> quand on dépasse 80 % du calibre onduleur/PDU. C'est exactement le type d'outil
> qui fait la différence entre « on subit » et « on pilote » son énergie.

---

## 27. Circuits : opérateurs et liaisons FAI

Le module **Circuits** documente vos liens opérateurs (fibre, xDSL, 4G/5G, MPLS).

1. **Fournisseur** (Circuits > Fournisseurs) : `Orange-Business`, `Free-Pro`, `OVH`.
   Renseignez le portail client, le support (téléphone, horaires) et les identifiants
   de compte — en intervention à 2h du matin, c'est précieux.
2. **Type de circuit** : `Fibre FTTO`, `FTTH`, `MPLS`, `xDSL`, `4G-Secours`.
3. **Circuit** : ex. `CT-OBS-2024-001` :
   - Débit montant/descendant (`1 Gb/s` / `1 Gb/s`),
   - Identifiant opérateur (n° de ligne, référence contrat),
   - Dates de mise en service / engagement,
   - **Terminaison A** (votre routeur : `SIEGE-RTR-01:Gi0/0`) et **terminaison Z**
     (équipement opérateur si connu).

**Bonnes pratiques :**

- Nommez les circuits avec l'année et un compteur : `CT-<Fournisseur>-<AAAA>-<NNN>`.
- Renseignez les **dates d'engagement** : NetBox ne relance pas tout seul, mais un
  rapport personnalisé « engagements arrivant à échéance » (section 59) le fait.
- Liez le circuit au **tenant** concerné et au **site**.

---

## 28. Virtualisation : clusters, VM et interfaces virtuelles

NetBox documente aussi le virtuel : **Virtualisation > Clusters**, puis **Machines virtuelles**.

1. **Type de cluster** : `Proxmox-VE`, `VMware-vSphere`, `Hyper-V`.
2. **Cluster** : ex. `PVE-SIEGE` (site = `SIEGE-Paris11`, 3 nœuds).
3. **Machine virtuelle** : ex. `VM-NETBOX-01` :
   - Cluster, statut (`Actif`), rôle (`Infra`, `Applicatif`),
   - vCPU, RAM, disque (Go),
   - Plateforme (`Debian 12`), tenant.

Les **interfaces virtuelles** (`eth0`…) portent les adresses IP comme en physique :
votre IPAM reste cohérent entre monde physique et virtuel.

> 💡 **Cohérence d'inventaire** : la VM `VM-NETBOX-01` qui héberge NetBox lui-même
> doit exister dans NetBox (cluster, ressources, IP). C'est le test ultime :
> « si NetBox ne se documente pas lui-même, personne ne le fera ».

---

## 29. Adresses MAC et découverte

NetBox stocke les adresses MAC au niveau des interfaces (physiques et virtuelles).
Renseignez-les pour :

- les équipements réseau (table ARP, DHCP statique),
- les serveurs (inventaire, Wake-on-LAN),
- les VM (cohérence hyperviseur).

Astuce d'alimentation en masse : exportez la table ARP de votre cœur de réseau,
puis utilisez l'API (section 56) ou un script pour rattacher les MAC aux interfaces.
Ne saisissez jamais des centaines de MAC à la main : **script > clic**.

---

## 30. Pièces jointes, photos et documentation liée

Chaque objet NetBox accepte des **pièces jointes** (photos, plans, PDF) :

- Photo de la baie (avant/après intervention),
- Plan du local technique,
- Fiche technique de l'onduleur (courbe d'autonomie),
- Contrat opérateur (sur la fiche circuit),
- Procédure de redémarrage (sur la fiche équipement critique).

Limitez la taille (`client_max_body_size` nginx, section 15) et **ne stockez jamais
de secrets** (mots de passe, clés privées) en pièce jointe : NetBox n'est pas un
coffre-fort. Pour les mots de passe : gestionnaire dédié (Vault, KeePass d'équipe).

> ⚠️ Les pièces jointes vivent dans `media/` : elles font partie de la sauvegarde
> (section 64). Une sauvegarde BDD sans `media/` = des fiches sans photos.

---

## 31. IPAM : les VRF (tables de routage virtuelles)

Une **VRF** (*Virtual Routing and Forwarding*) isole des plans d'adressage qui se
chevauchent : même préfixe `10.1.0.0/16` utilisable dans deux VRF différentes sans conflit.

**IPAM > VRF > + Ajouter** :

| Champ | Exemple |
|---|---|
| Nom | `VRF-PROD`, `VRF-GUEST`, `VRF-MGMT` |
| RD (Route Distinguisher) | `65000:100` |
| Tenant | `DSI` |
| Description | `VRF de production — routage MPLS` |

Cas d'usage typiques :

