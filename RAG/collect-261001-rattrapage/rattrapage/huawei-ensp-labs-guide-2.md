---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-2
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [110, 246]
sha256: 48b206c540f1727af468d70afb6fd4256dfa1835381051b924b904dd82474c03
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

1. Clic droit sur l'icône eNSP > **Exécuter en tant qu'administrateur** (toujours).
2. Menu **Menu > Outils > Options** : vérifier que l'onglet VirtualBox détecte bien la version installée.
3. Créer un nouveau projet : la palette d'équipements à gauche doit afficher routeurs, switches, pare-feu, WLAN, etc.
4. Si un équipement affiche `???` après démarrage : problème d'image VirtualBox — réinstaller eNSP en administrateur, ou vérifier qu'aucun antivirus ne bloque les VM.

### 1.6. Les images : rien à ajouter (c'est inclus)

Contrairement à GNS3, **eNSP embarque déjà les images** des équipements dans son installeur : AR1220/2220/3260, S3700, S5700, USG5500, AC6605, AP... Il n'y a **aucun import d'image à faire**. C'est une des forces d'eNSP pour une formation d'équipe : zéro manipulation de firmware.

### 1.7. Hello World : deux routeurs, un ping

**Objectif** : valider que l'installation fonctionne de bout en bout.

**Topologie :**

```
[R1 : AR2220] GE0/0/0 (192.168.1.1/24) ---- câble Copper ---- GE0/0/0 (192.168.1.2/24) [R2 : AR2220]
```

**Procédure :**

1. Glisser-déposer deux routeurs **AR2220** sur le canevas.
2. Sélectionner l'outil **Copper** (câble), cliquer sur R1 > choisir **GE0/0/0**, cliquer sur R2 > choisir **GE0/0/0**.
3. Bouton **Démarrer** (triangle vert) : les deux routeurs passent au vert après 1 à 2 minutes (le premier boot est long).
4. Clic droit sur R1 > **CLI** : une console s'ouvre.

**Configuration R1 :**

```
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R1
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ip address 192.168.1.1 24
[R1-GigabitEthernet0/0/0]quit
```

**Configuration R2 :**

```
<Huawei>system-view
[Huawei]sysname R2
[R2]interface GigabitEthernet 0/0/0
[R2-GigabitEthernet0/0/0]ip address 192.168.1.2 24
[R2-GigabitEthernet0/0/0]quit
```

**Test :**

```
[R2]ping 192.168.1.1
  PING 192.168.1.1: 56  data bytes, press CTRL_C to break
    Reply from 192.168.1.1: bytes=56 Sequence=1 ttl=255 time=10 ms
    ...
```

> Si le ping échoue : vérifier que les interfaces sont **up** (`display interface brief`), que les câbles sont bien branchés sur les bonnes interfaces, et que les deux équipements sont démarrés (icône verte).

**Sauvegarder** : Menu > **Fichier > Enregistrer** : le fichier `.topo` contient la topologie *et* les configurations (eNSP sauvegarde les configs dans le .topo). Nommer `hello_world.topo`.

---

## 2. Prise en main d'eNSP

### 2.1. L'interface en 5 zones

1. **Palette d'équipements** (gauche) : catégories Routeur, Switch, Pare-feu, WLAN, Serveur, Terminal, etc. Glisser-déposer sur le canevas.
2. **Canevas** (centre) : la topologie. Molette = zoom, clic droit = menu contextuel d'un équipement.
3. **Barre d'outils** (haut) : Nouveau, Ouvrir, Enregistrer, Démarrer/Arrêter tous les équipements, sélection, ajout de texte/étiquettes, câblage (Copper/Serial/Auto), capture.
4. **Panneau de démarrage** : le bouton triangle vert démarre **tous** les équipements ; le carré rouge les arrête. On peut aussi démarrer/arrêter unitairement par clic droit.
5. **Barre de statut** (bas) : état des équipements.

### 2.2. Ajouter des équipements et câbler

- **Ajout** : glisser l'icône depuis la palette. Renommer immédiatement (clic droit > **Renommer** ou double-clic sur le nom) : `R1`, `SW1`, `FW1`... Un lab bien nommé se dépanne deux fois plus vite.
- **Câblage** : cliquer l'outil câble (ou touche rapide), puis cliquer le premier équipement (choisir l'interface dans la liste), puis le second. Types :
  - **Copper** : Ethernet (cas général).
  - **Serial** : liaisons série (WAN, TP Frame-Relay/PPP si besoin).
  - **Auto** : eNSP choisit (éviter en formation : préférer Copper explicite).
- **Règles de câblage à enseigner dès le début** :
  - Sur un routeur AR2220 d'eNSP : `GE0/0/0` = port de management/diag souvent, préférer `GE0/0/1` et `GE0/0/2` pour les labs (comme dans les TP ci-dessous).
  - Sur un S3700 : `Ethernet0/0/1` à `Ethernet0/0/22` + 2 ports `GE0/0/1-2`.
  - Sur un S5700 : `GE0/0/1` à `GE0/0/24`.
- **Supprimer un câble** : clic droit sur le câble > Supprimer. **Déplacer** : les câbles suivent les équipements.

### 2.3. La console (CLI)

- Clic droit sur un équipement démarré > **CLI** : ouvre la console série.
- Touches utiles : `?` (aide contextuelle), `Tab` (complétion), `Ctrl+Z` (retour en vue utilisateur depuis n'importe quelle vue), flèches haut/bas (historique).
- **Modes de vue** :
  - `<R1>` : user view — commandes `display`, `ping`, `save`.
  - `[R1]` : system view (via `system-view`) — toute la configuration.
  - `[R1-GigabitEthernet0/0/1]` : interface view.
- **Enregistrer la config** : `save` (répondre `y`). Sans `save`, la config est perdue à l'arrêt (eNSP peut aussi l'embarquer dans le .topo selon l'option, mais `save` reste le réflexe à enseigner).

### 2.4. Capture Wireshark sur un lien

1. L'équipement doit être démarré.
2. Clic droit sur le **câble** > **Démarrer la capture** (ou bouton capture dans la barre d'outils puis clic sur le câble).
3. Wireshark s'ouvre et capture en direct : idéal pour voir les hello OSPF, les DTP/CDP (absents chez Huawei), les BPDU STP, les échanges DHCP (Discover/Offer/Request/ACK), les paquets IKE/IPSec.
4. Arrêter : clic droit sur le câble > **Arrêter la capture**.

> **Astuce d'animateur** : la capture est l'outil pédagogique le plus puissant d'eNSP. Pour chaque TP, prévoir un moment "regardons les paquets" : BPDU (TP3), DHCP (TP4), Hello OSPF (TP6), IKE (TP9).

### 2.5. Sauvegarde : le fichier .topo

- **Fichier > Enregistrer** (`.topo`) : contient le placement, le câblage, les noms **et** les configurations des équipements.
- **Bonne pratique d'équipe** : un dossier par TP (`TP01_VLAN/`, `TP02_Trunk/`...), avec `sujet.md` (énoncé) + `topologie.topo` (topologie de départ, équipements vierges ou pré-câblés) + `correction.md`.
- **Partage** : un `.topo` s'ouvre sur n'importe quel poste eNSP de même version. Pour distribuer un TP "à compléter", enregistrer le .topo **après** avoir fait `reset saved-configuration` + `reboot` sur chaque équipement (ou simplement ne jamais faire `save` sur la topologie de départ).

### 2.6. L'objet "Cloud" : pont vers le réel

L'équipement **Cloud** permet de relier eNSP au réseau de la machine hôte (carte physique ou interface loopback) : utile pour, par exemple, faire sortir un ping du lab vers le LAN du formateur, ou brancher un vrai PC. Configuration : clic droit sur le Cloud > **Paramètres**, lier une interface UDP ou une carte réseau à un port du Cloud, puis câbler comme un équipement normal. (Non utilisé dans les TP ci-dessous, mais à connaître.)

### 2.7. Réflexes à faire acquérir dès cette séance

1. Toujours nommer les équipements (`sysname` + renommage graphique).
2. Toujours vérifier l'état des interfaces (`display interface brief`) avant de configurer.
3. `save` en fin de TP.
4. Capturer un lien pour "voir" un protocole au moins une fois par TP.
5. Documenter : copier la config finale (`display current-configuration`) dans un fichier texte.

**Durée estimée de la séance** : 1 h 30 (installation déjà faite en amont).
**Barème** : non noté — séance de découverte, validation par le "hello world" fonctionnel.

---

## 3. TP1 — VLAN de base sur S3700

### Objectif
Comprendre la notion de VLAN : segmentation logique d'un switch, ports access, isolation entre VLAN, communication intra-VLAN.

### Prérequis
Séance de prise en main (section 2). Savoir ouvrir une console et passer en `system-view`.

### Topologie

