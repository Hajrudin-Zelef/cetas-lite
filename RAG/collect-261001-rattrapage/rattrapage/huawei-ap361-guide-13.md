---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-13
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [2022, 2196]
sha256: ee109e1ae5eebeb1eb71a676e1d2dc84eea83142ee7c8a04d2f3c85217600558
---

# Huawei eKit AP361 — Guide ultra-complet

- Nouvelle **cloison** (placo avec isolant métallisé, vitre blindée) → zone d'ombre.
- Nouvel équipement **perturbateur** (variateur LED, onduleur, moteur).
- Câble **sectionné puis réparé** à la va-vite (domino, scotch) → 100 Mbit/s.
- AP **déplacé** par les peintres et jamais remis (« il gênait »).
- Nouveau **voisin Wi-Fi** agressif sur votre canal.

**Méthode** : « quoi qui a changé ? » → visite sur place → mesure avant/après.
Dans 80 % des cas, la cause est **physique** (câble, emplacement, cloison), pas
logicielle.

## 129. Cas n°9 — Un SSID a disparu

**Symptômes** : un des SSID n'est plus diffusé (les autres OK).

**Diagnostic** :

1. eKit : le SSID est-il toujours **activé** ? (Un collègue l'a-t-il désactivé
   « temporairement » il y a 3 mois ?)
2. Est-il restreint à une **plage horaire** ? (SSID invité coupé la nuit, c'est
   normal.)
3. Est-il lié à une **radio désactivée** ? (SSID 5 GHz only + radio 5 GHz en
   panne = SSID invisible.)
4. L'AP diffuse-t-il les autres SSID ? Non → problème AP global, pas SSID.

## 130. Cas n°10 — L'AP est « en ligne » mais personne ne s'y connecte

**Symptômes** : AP vert dans eKit, zéro client dessus, les voisins sont chargés.

**Causes** :

1. **SSID non appliqué** à cet AP (oubli dans le template / groupe).
2. **Radio désactivée** ou en panne (vérifier l'état radio).
3. **Canal DFS en attente** : l'AP a détecté un radar et attend (jusqu'à 10 min)
   — ou change de canal en boucle près d'un aéroport.
4. **Puissance trop faible** : l'AP chuchote, personne ne l'entend.
5. **VLAN** : les clients s'associent puis repartent faute de DHCP (voir §126) —
   l'AP paraît « vide ».

## 131. Cas n°11 — Débit asymétrique (down OK, up minable)

**Symptômes** : 200 Mbit/s en download, 5 Mbit/s en upload.

**Causes** :

1. **Liaison asymétrique** : l'AP arrose fort (puissance haute) mais le client
   (smartphone, faible puissance d'émission) ne « remonte » pas bien. →
   **Baisser la puissance de l'AP** (contre-intuitif mais efficace).
2. **Interférences** près du client (pas près de l'AP).
3. **Limite configurée** : rate-limiting asymétrique oublié (§100).
4. **Client** : économie d'énergie agressive qui bride l'émission.

## 132. Cas n°12 — Imprimante / IoT qui ne se connecte plus

**Symptômes** : l'objet connecté marchait, puis plus rien après un changement
(ou sans raison apparente).

**Checklist IoT** :

- [ ] L'objet est-il **2,4 GHz only** ? (La plupart.) Le SSID IoT est-il bien en
      2,4 GHz ? Le band steering ne le bloque-t-il pas ? (§52)
- [ ] Clé : caractères spéciaux non supportés par l'objet ? (Certains n'acceptent
      que l'alphanumérique.)
- [ ] WPA3 : l'objet ne connaît que WPA2 ? → SSID IoT en **WPA2-PSK**.
- [ ] DHCP : l'objet a-t-il une IP ? (Beaucoup d'IoT n'ont pas d'écran : sniffez
      le DHCP.)
- [ ] L'objet s'est-il mis à jour tout seul avec un firmware buggué ? (Ça arrive.)
- [ ] Redémarrage électrique de l'objet (le « have you tried turning it off and
      on again » s'applique cruellement bien aux IoT).

## 133. Cas n°13 — Lenteurs à heures fixes (12h-14h, 18h…)

**Symptômes** : ça rame tous les jours aux mêmes heures.

**Causes** :

1. **Micro-ondes** (12h-14h en 2,4 GHz) — le grand classique.
2. **Pic d'usage** : tout le monde en visio à 14h → saturation de la cellule
   (normal, à dimensionner).
3. **Sauvegardes** : un backup qui part à 18h et sature le lien ou l'airtime.
4. **Voisin** : le restaurant d'à côté allume son Wi-Fi le midi.
5. **Planification** : scans DCA ou upgrades planifiés aux mauvaises heures.

**Méthode** : corrélez avec l'agenda du bâtiment, pas seulement avec les logs.

## 134. Cas n°14 — Après une coupure de courant, le Wi-Fi ne revient pas

**Symptômes** : le courant est revenu, mais pas le Wi-Fi (ou partiellement).

**Diagnostic** :

1. Le **switch PoE** est-il revenu ? (S'il n'est pas sur onduleur et que son
   disjoncteur a sauté, les AP restent morts.)
2. Les AP **bootent-ils** ? (2-4 min par AP, ne paniquez pas à T+30 s.)
3. Le **DHCP/DNS/Internet** sont-ils revenus ? (L'AP sans cloud reste-t-il
   opérationnel ? — voir §43.)
4. Un AP ne revient pas : PoE ? (§124) Firmware corrompu par la coupure
   pendant un upgrade ? (§118).
5. **Ordre de démarrage** : le routeur/firewall doit être up **avant** les AP
   (sinon les AP bootent sans DHCP/DNS et restent dans un état bancal → prévoir
   un redémarrage PoE des AP après le retour du cœur réseau).

> 🔌 **Procédure de réalimentation** : cœur réseau (routeur, DHCP, DNS) → attendre
> 5 min → switch PoE → attendre 5 min → vérifier les AP. Automatisez ça avec des
> prises séquencées si les coupures sont fréquentes.

## 135. Cas n°15 — Le cloud eKit voit l'AP « hors ligne » mais le Wi-Fi marche

**Symptômes** : les utilisateurs sont contents, mais eKit dit l'AP hors ligne.

**Causes** :

1. L'AP a perdu l'accès **Internet** (le cloud) mais pas le LAN → le Wi-Fi local
   continue (selon version), la gestion est aveugle.
2. **DNS** en panne : l'AP ne résout plus les domaines Huawei.
3. **Horloge** : décalage NTP → TLS vers le cloud refusé.
4. **Firewall** : une règle a été ajoutée et bloque les flux sortants de l'AP.

**Danger** : sans supervision cloud, vous ne verrez pas la prochaine vraie panne.
Traitez comme un incident de supervision (priorité haute, pas critique).

## 136. Cas n°16 — Boucle / tempête broadcast via un AP

**Symptômes** : tout le réseau rame, switch saturé, même le filaire.

**Causes** :

1. Un **petit switch** branché sous l'AP (pour « avoir 2 prises ») avec un câble
   en boucle.
2. L'AP branché sur **deux ports** du switch (l'utilisateur a vu 2 prises…).
3. Un **rogue AP** en mode répéteur qui reboucle.

**Action immédiate** : débranchez le port suspect (vous retrouverez le calme en
quelques secondes), puis cherchez la boucle physiquement. **Prévention** :
`port-isolate` sur les ports AP (§75), **STP** actif sur le switch, et interdiction
affichée de brancher des petits switchs sauvages.

## 137. Arbre de décision rapide (à afficher dans la baie)

```
AP LED éteinte ?
├── Oui → PoE ? (testeur) → Non → câble/switch/budget (§124)
│                          → Oui → AP HS ? tester sur injecteur → RMA
└── Non (LED allumée)
    ├── SSID visible ?
    │   ├── Non → SSID activé ? radio OK ? canal DFS ? (§129, §130)
    │   └── Oui → association OK ?
    │       ├── Non → 1 client ou tous ? clé ? WPA3 ? pilote ? (§123)
    │       └── Oui → IP obtenue ?
    │           ├── Non (169.254.x.x) → DHCP/VLAN/trunk (§126)
    │           └── Oui → Internet OK ?
    │               ├── Non → routage/firewall/DNS
    │               └── Oui mais lent → port 100M ? interférences ?
    │                   charge ? (§121, §122)
    └── Wi-Fi OK mais eKit « hors ligne » → Internet/DNS/NTP/firewall (§135)
```

---
---

# O. MAINTENANCE PRÉVENTIVE

## 138. Plan de maintenance annuel (modèle)

| Période | Actions | Charge |
|---|---|---|
| **Mensuel** | Revue eKit (2 min/jour en fait), acquittement alarmes, vérification sauvegardes | 30 min |
| **Trimestriel** | Ronde radio à l'analyseur, dépoussiérage, contrôle fixations, revue des logs, rotation clé invités | ½ journée |
| **Semestriel** | Audit sécurité (rogue, PMF, comptes), test de la procédure « cloud injoignable », revue du plan radio | 1 journée |
| **Annuel** | Upgrade firmware de référence, audit complet, exercice de réalimentation (§134), révision documentaire, inventaire | 2-3 jours |

> 📅 Mettez ces créneaux **au planning d'équipe** en début d'année. La maintenance
> qui n'est pas planifiée n'existe pas.

## 139. Checklist trimestrielle (ronde terrain)

