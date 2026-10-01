---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-3
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [247, 443]
sha256: 20f971f446f5215ac96658a59756913dd413290d154070e81b15239400e11ae9
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

```
                    ┌─────────────────────────┐
  PC1 (192.168.10.11/24) ──E0/0/1 │                         │ E0/0/3── PC3 (192.168.20.13/24)
  PC2 (192.168.10.12/24) ──E0/0/2 │      SW1 (S3700)        │ E0/0/4── PC4 (192.168.20.14/24)
                    └─────────────────────────┘
```

- 1 switch **S3700** nommé `SW1`.
- 4 **PC** : PC1, PC2 (VLAN 10, réseau 192.168.10.0/24), PC3, PC4 (VLAN 20, réseau 192.168.20.0/24).
- Câblage : E0/0/1→PC1, E0/0/2→PC2, E0/0/3→PC3, E0/0/4→PC4.
- Configuration IP des PC : double-clic sur le PC > onglet de configuration IPv4 (passerelle : laisser vide pour ce TP, pas de routage).

### Énoncé (à distribuer aux stagiaires)

1. Créer les VLAN 10 (nom : `USERS`) et 20 (nom : `GUESTS`) sur SW1.
2. Affecter E0/0/1 et E0/0/2 en **access VLAN 10**, E0/0/3 et E0/0/4 en **access VLAN 20**.
3. Configurer les IP des 4 PC selon le plan d'adressage.
4. Tester : PC1 ping PC2 (doit réussir), PC1 ping PC3 (doit échouer — pourquoi ?).
5. Vérifications : afficher la table VLAN et l'état des ports.

### Correction pas à pas

**Sur SW1 :**

```
<Huawei>system-view
[Huawei]sysname SW1
[SW1]vlan 10
[SW1-vlan10]description USERS
[SW1-vlan10]quit
[SW1]vlan 20
[SW1-vlan20]description GUESTS
[SW1-vlan20]quit

# Affectation des ports access (méthode par batch, à enseigner)
[SW1]port-group group-member Ethernet 0/0/1 to Ethernet 0/0/2
[SW1-port-group]port link-type access
[SW1-port-group]port default vlan 10
[SW1-port-group]quit

[SW1]port-group group-member Ethernet 0/0/3 to Ethernet 0/0/4
[SW1-port-group]port link-type access
[SW1-port-group]port default vlan 20
[SW1-port-group]quit

[SW1]save
```

**Sur les PC** (double-clic > configuration) :

| PC | IP | Masque |
|---|---|---|
| PC1 | 192.168.10.11 | 255.255.255.0 |
| PC2 | 192.168.10.12 | 255.255.255.0 |
| PC3 | 192.168.20.13 | 255.255.255.0 |
| PC4 | 192.168.20.14 | 255.255.255.0 |

**Tests** (clic droit sur PC > invite de commande, ou double-clic) :

```
PC1> ping 192.168.10.12     # OK : même VLAN, même sous-réseau
PC1> ping 192.168.20.13     # ECHEC : VLAN différents = broadcast domains séparés
PC3> ping 192.168.20.14     # OK
```

### Vérifications

```
[SW1]display vlan
# Affiche les VLAN créés et les ports membres (untagged)

[SW1]display port vlan
# Vue synthétique : Port | Link Type | PVID | VLAN List

[SW1]display mac-address
# Table MAC apprise par VLAN : constater que PC1/PC2 sont en VLAN 10, PC3/PC4 en VLAN 20
```

### Pièges classiques

1. **Oublier `port default vlan`** : le port reste en VLAN 1, les PC ne se voient pas. Diagnostic : `display port vlan` montre PVID=1.
2. **Confondre VLAN et sous-réseau** : mettre PC1 en 192.168.20.x tout en le laissant en VLAN 10 ne fera pas communiquer avec le VLAN 20. Le VLAN isole au niveau 2, l'IP au niveau 3 : il faut les deux alignés.
3. **PC mal configuré** : masque oublié (eNSP met parfois 0.0.0.0) — toujours revérifier la fiche IP du PC.
4. **VLAN créé mais port en `hybrid` par défaut ?** Non : sur VRP les ports sont `hybrid`... en fait sur les S Huawei, le type par défaut est **hybrid** avec PVID 1. Le passer explicitement en `access` évite toute ambiguïté — c'est pourquoi la correction le fait.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| VLAN 10/20 créés et nommés | 4 |
| Ports en access dans les bons VLAN | 6 |
| Adressage IP des PC correct | 4 |
| Tests ping conformes aux attendus + explication de l'échec inter-VLAN | 4 |
| Vérifications `display` produites | 2 |

### Durée estimée
**45 minutes** (dont 10 min de démo animateur).

### Fiche animateur — points à insister
- Un VLAN = un domaine de broadcast = un sous-réseau IP (en pratique d'entreprise).
- Faire le parallèle avec le terrain : sur vos S310, les VLAN séparent typiquement `DATA`, `VOIX` (téléphonie), `WIFI`, `GESTION`.
- Montrer `display mac-address` : la table MAC est **par VLAN**, c'est la preuve concrète de la segmentation.
- Question piège à poser : "PC1 en 192.168.10.11/24 et PC3 en 192.168.10.13/24 mais PC3 branché sur un port VLAN 20 : ping ?" → Échec, car le tag VLAN bloque au niveau 2 avant même l'IP.

---

## 4. TP2 — Trunk 802.1Q entre deux switches

### Objectif
Transporter plusieurs VLAN sur un lien inter-switch (trunk 802.1Q), comprendre le taggage, diagnostiquer un trunk mal configuré.

### Prérequis
TP1 validé (VLAN, ports access).

### Topologie

```
PC1 (VLAN 10, .11) ──E0/0/1
PC2 (VLAN 20, .13) ──E0/0/2    SW1 (S3700) ═══ trunk GE0/0/1 ═══ GE0/0/1 SW2 (S3700)    E0/0/1── PC3 (VLAN 10, .12)
                                                                                      E0/0/2── PC4 (VLAN 20, .14)
```

- 2 switches **S3700** : SW1, SW2.
- Lien trunk : **GE0/0/1 ↔ GE0/0/1** (câble Copper).
- Ports access : comme indiqué, VLAN 10 = 192.168.10.0/24, VLAN 20 = 192.168.20.0/24.

### Énoncé

1. Sur SW1 et SW2, créer VLAN 10 (`USERS`) et VLAN 20 (`GUESTS`).
2. Configurer les ports access (E0/0/1 → VLAN 10, E0/0/2 → VLAN 20 sur chaque switch).
3. Configurer le lien GE0/0/1 en **trunk** autorisant les VLAN 10 et 20, des deux côtés.
4. Tester tous les pings croisés (PC1↔PC3 OK, PC2↔PC4 OK, PC1↔PC2 KO).
5. **Dépannage** : l'animateur casse volontairement un côté du trunk (ex. `port trunk allow-pass vlan 10` seul sur SW2). Le stagiaire doit diagnostiquer avec les `display` et réparer.

### Correction pas à pas

**SW1 :**

```
<Huawei>system-view
[Huawei]sysname SW1
[SW1]vlan batch 10 20
[SW1-vlan10]description USERS
[SW1-vlan20]description GUESTS
[SW1]quit  # (quitter les vues vlan)

[SW1]interface Ethernet 0/0/1
[SW1-Ethernet0/0/1]port link-type access
[SW1-Ethernet0/0/1]port default vlan 10
[SW1-Ethernet0/0/1]quit
[SW1]interface Ethernet 0/0/2
[SW1-Ethernet0/0/2]port link-type access
[SW1-Ethernet0/0/2]port default vlan 20
[SW1-Ethernet0/0/2]quit

# Le trunk
[SW1]interface GigabitEthernet 0/0/1
[SW1-GigabitEthernet0/0/1]port link-type trunk
[SW1-GigabitEthernet0/0/1]port trunk allow-pass vlan 10 20
[SW1-GigabitEthernet0/0/1]quit
[SW1]save
```

**SW2 :** identique (mêmes VLAN, mêmes ports access, même trunk sur GE0/0/1).

**Adressage PC :**

| PC | IP | VLAN |
|---|---|---|
| PC1 | 192.168.10.11/24 | 10 |
| PC3 | 192.168.10.12/24 | 10 |
| PC2 | 192.168.20.13/24 | 20 |
| PC4 | 192.168.20.14/24 | 20 |

**Tests attendus :**

```
PC1> ping 192.168.10.12   # OK (VLAN 10 à travers le trunk)
PC2> ping 192.168.20.14   # OK (VLAN 20 à travers le trunk)
PC1> ping 192.168.20.13   # KO (VLAN différents)
```

### Vérifications

```
[SW1]display port vlan
# GE0/0/1 doit apparaître en trunk, PVID 1, VLAN List incluant 10 et 20 (tagged)

[SW1]display vlan 10
[SW1]display vlan 20
# Vérifier que GE0/0/1 est membre tagged des deux VLAN
```

**Capture Wireshark sur le lien trunk** (moment pédagogique fort) : lancer un ping PC1→PC3 et observer les trames Ethernet avec le champ **802.1Q VLAN ID = 10**. Puis ping PC2→PC4 : VLAN ID = 20. Le tag est la preuve visible du trunk.

### Pièges classiques

