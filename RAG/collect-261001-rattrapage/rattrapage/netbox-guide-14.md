---
id: collect-261001-rattrapage/rattrapage/netbox-guide-14
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [2706, 2863]
sha256: 74fad3fe09b4d5d08d76b4cfbdad169456632276e25cf66793eac6084e97dae8
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

```
Région : France
└── Groupe de sites : Agences
    └── Site : AGENCE-Nantes (adresse, GPS, tenant DSI)
Tenant : DSI (+ INFRA-MUTUALISEE pour l'onduleur/la baie si partagés)
```

### Étape 2 — Plan d'adressage (préfixe container + enfants)

| Préfixe | VLAN | Rôle | Usage |
|---|---|---|---|
| `10.30.0.0/16` | — | (Container) | bloc du site |
| `10.30.10.0/24` | 10 | Serveurs | serveurs + NAS |
| `10.30.20.0/24` | 20 | LAN-Utilisateurs | 15 postes (DHCP `.50`–`.150`) |
| `10.30.30.0/24` | 30 | WiFi | AP + invités (VRF-GUEST pour les invités !) |
| `10.30.40.0/24` | 40 | Impression | 2 MFP |
| `10.30.254.0/24` | 254 | Management | switchs, routeur, iDRAC, carte onduleur |

Créez les VLAN 10/20/30/40/254 dans le groupe `VLAN-Nantes`, puis les préfixes
(statut `Actif`, sauf le `/16` en `Container`), puis les plages DHCP.

### Étape 3 — Baie et chaîne électrique

```
Baie : AGENCE-Baie-A01 (42U, rôle Baie-Reseau, tenant INFRA-MUTUALISEE)
U42 : (vide — ventilation haute)
U40 : AGENCE-UPS-A01 (onduleur 3 kVA, 2U → U40-U41)
U38 : AGENCE-PDU-A01-A (0U vertical — noté en commentaire si 0U)
U37 : AGENCE-PDU-A01-B
U30 : AGENCE-RTR-01 (routeur, 1U)
U28 : AGENCE-SW-A01-01 (switch 48p, 1U)
U27 : AGENCE-SW-A01-02 (switch 48p, 1U)
U20 : AGENCE-SRV-HYP-01 (serveur 2U → U20-U21)
U18 : AGENCE-SRV-HYP-02 (serveur 2U → U18-U19)
U10 : AGENCE-NAS-01 (NAS 2U → U10-U11)
U01-U05 : réserve (réservation NetBox : "extension 2027")
```

Câblage électrique (ports d'alimentation) :

```
AGENCE-UPS-A01:Sortie-1 --> AGENCE-PDU-A01-A:Entrée
AGENCE-UPS-A01:Sortie-2 --> AGENCE-PDU-A01-B:Entrée
AGENCE-SRV-HYP-01:PSU1 --> PDU-A:Prise-01   / PSU2 --> PDU-B:Prise-01
AGENCE-SRV-HYP-02:PSU1 --> PDU-A:Prise-02   / PSU2 --> PDU-B:Prise-02
AGENCE-NAS-01:PSU1     --> PDU-A:Prise-03   / PSU2 --> PDU-B:Prise-03
AGENCE-SW-A01-01:PSU1  --> PDU-A:Prise-04   (mono-alim : noter "pas de redondance")
```

> ⚠️ Le switch mono-alimenté est un **risque documenté** : ajoutez le tag
> `mono-alimentation` et un champ personnalisé `criticite=Majeur`. En cas de
> maintenance PDU-A, on sait qu'il faut planifier une coupure.

### Étape 4 — Câblage data et uplink

```
AGENCE-RTR-01:Gi0/0 --fibre OM4 10m--> AGENCE-SW-A01-01:SFP-49 (uplink, trunk)
AGENCE-SW-A01-01:SFP-50 --fibre--> AGENCE-SW-A01-02:SFP-49 (stack/uplink)
AGENCE-SW-A01-01:Gi1/0/1 --Cat6a 3m--> AGENCE-SRV-HYP-01:eth0 (VLAN 10, access)
AGENCE-SW-A01-01:Gi1/0/2 --Cat6a 3m--> AGENCE-SRV-HYP-02:eth0
Patch panel : AGENCE-PB-01 (24 ports) en U25, rocade vers les prises murales.
```

### Étape 5 — Adressage et management

- IP management : `10.30.254.11` (RTR), `.12`/`.13` (SW), `.21`/`.22` (iDRAC),
  `.30` (carte réseau onduleur — supervision SNMP !).
- Passerelles : `.1` sur chaque VLAN (VRRP si deux routeurs).
- DNS : `agence-nantes-sw-01.lan-entreprise.fr`… renseigné dans chaque fiche IP.
- Services : `ssh`/`https`/`snmp` sur les équipements réseau.

### Étape 6 — Vérifications finales (checklist)

- [ ] Chaque équipement a : site, baie+U, n° de série, IP de management, tenant
- [ ] Chaque câble a : type, longueur, étiquette
- [ ] La chaîne électrique est complète (aucun équipement « flottant » sans PDU)
- [ ] Le rapport « bilan de puissance » ne dépasse pas 80 % du calibre onduleur
- [ ] Photos de la baie jointes à la fiche
- [ ] Export CSV archivé dans la documentation du site

---

## 87. Cas pratique n°2 : refonte du plan d'adressage d'un site existant

**Contexte** : le site `SIEGE-Paris11` utilise encore du `192.168.1.0/24` « historique »
partagé entre serveurs, utilisateurs et imprimantes. Objectif : migrer vers un plan
propre sans couper la production.

1. **Inventaire de l'existant** : scan + table ARP + DHCP → importez TOUT dans NetBox
   dans le préfixe `192.168.1.0/24` (statut `Déprécié`, rôle par usage réel constaté).
   Vous avez maintenant la photo exacte du chaos.
2. **Concevez la cible** : nouveau container `10.10.0.0/16`, découpé par rôle
   (section 34). Créez les préfixes en statut `Réservé`.
3. **Plan de migration par VLAN** : un VLAN = une soirée d'intervention.
   Ex. semaine 1 : VLAN 40 (imprimantes — faible risque), semaine 2 : VLAN 10 (serveurs).
4. **Double documentation pendant la migration** : l'IP ancienne (statut `Déprécié`,
   description « migré vers 10.10.x.y le … ») reste 30 jours, puis suppression.
5. **Bascule** : changez les statuts (`Réservé` → `Actif`, ancien → `Déprécié`),
   mettez à jour DHCP/DNS via les scripts (section 56), vérifiez avec le rapport
   « IP sans DNS » et « équipements sans IP de management » (section 59).
6. **Clôture** : quand le vieux préfixe est vide, supprimez-le. Fêtez ça. 🎉

> 💡 Pendant toute la migration, le **changelog** (section 45) est votre meilleur ami :
> chaque changement est horodaté et attribué. En cas de rollback, on sait exactement
> quoi défaire.

---

## 88. Cas pratique n°3 : migrer un tableur d'inventaire vers NetBox

**Contexte** : 300 lignes Excel « inventaire réseau » (nom, IP, modèle, baie, U).
Objectif : import propre, sans doublons, avec traçabilité.

1. **Nettoyez le tableur** : une colonne = un champ NetBox ; normalisez les noms
   selon la convention (section 82) ; vérifiez l'unicité des noms et des IP
   (`=NB.SI` / suppressions de doublons).
2. **Préparez le référentiel** : fabricants, types d'équipements, rôles, site, baie,
   VLAN, préfixes — l'import des équipements ÉCHOUE si le type ou le site n'existe pas.
3. **Importez par couches**, dans l'ordre des dépendances :
   `sites → baies → types → équipements → interfaces → IP → câbles`.
4. **Testez sur 5 lignes** dans le tenant `BAC-A-SABLE`, vérifiez, supprimez, puis
   importez le lot réel par paquets de 50 (section 81).
5. **Réconciliez** : rapport « équipements sans IP » + « IP sans équipement » ;
   traitez les écarts (le tableur mentait probablement sur 5–10 % des lignes —
   c'est normal, c'est pour ça qu'on migre).
6. **Archivez le tableur** en pièce jointe d'un ticket « migration inventaire »,
   puis **supprimez-le des partages** : deux sources de vérité = zéro source de vérité.

---

## 89. Cas pratique n°4 : préparer une maintenance onduleur sans surprise

**Contexte** : remplacement des batteries de `SIEGE-UPS-A01` (bypass + arrêt).
C'est LE cas qui justifie la modélisation électrique (section 26).

1. **Liste d'impact** : depuis la fiche de l'onduleur > onglet **Alimentations connectées**,
   exportez tous les équipements alimentés (via PDU-A et PDU-B).
2. **Tri par criticité** : champ personnalisé `criticite` (section 47) → les
   `Critique` (cœur, firewalls) exigent un créneau validé par la direction.
3. **Vérifiez la redondance** : pour chaque équipement, contrôlez qu'il a bien ses
   deux PSU câblées (A ET B). Les mono-alimentés (tag `mono-alimentation`) seront
   coupés : planifiez leur arrêt propre AVANT.
4. **Ordre d'arrêt/démarrage** : documentez-le dans la description de l'intervention
   (VM d'abord, hyperviseurs ensuite, stockage en dernier à l'arrêt ; inverse au
   redémarrage — les dépendances sont visibles via les câbles et les rôles).
5. **Le jour J** : passez les équipements en statut `En maintenance` dans NetBox
   (ils disparaissent des vues « Actif », la supervision peut être mise en pause
   proprement via l'export Zabbix, section 62).
6. **Retour** : repassez en `Actif`, joignez le rapport d'intervention et la photo
   des batteries neuves à la fiche de l'onduleur, mettez à jour `date_fin_garantie`.

> 💡 Ce cas pratique, mené une fois « pour de vrai », convertit les derniers
> sceptiques : « on a su EXACTEMENT ce qui allait s'éteindre » vaut tous les discours.

---

## 90. Cas pratique n°5 : audit de sécurité express depuis NetBox

