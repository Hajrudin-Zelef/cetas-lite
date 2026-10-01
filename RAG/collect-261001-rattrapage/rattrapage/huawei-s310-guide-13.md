---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-13
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [2644, 2842]
sha256: 1ab67aabe8cdcd5386b27297837bf16947224261c252df21bd3f2a54c4e0be4c
---

# Guide ultra-complet — Huawei eKit S310

**Méthode bouton RST :** appui long (~5–10 s) jusqu'à changement de la LED SYS
(durée exacte **à vérifier sur le guide du modèle exact**).

⚠️ **Irréversible.** Avant : sauvegarde la config (section 88) — même si tu
penses ne plus en avoir besoin. Après : refais toute la checklist de première
configuration (section 26), **y compris le mot de passe admin**.

---

## 94. iStack : le principe (stacking intelligent)

**iStack** permet de regrouper **jusqu'à 4 switches de la même série** en **un
seul switch logique** (valeur datasheet) :
- **Une seule IP** de management, **une seule config** pour tout le stack.
- Les liens inter-switch servent aussi à l'**agrégation inter-équipements**
  (un Eth-Trunk avec un membre sur chaque switch = redondance matérielle).
- Si un membre tombe, les autres continuent (redondance).

**Prérequis impératifs :**
- Même **modèle** (même série S310), même **version logicielle** sur tous.
- Câblage des ports de stack **avant** la mise sous tension du stack.
- Certains modèles ont **2 ports de stack dédiés** (pas de config nécessaire,
  les uplinks restent libres) — présence **à vérifier sur la fiche du modèle exact**.

---

## 95. iStack : configuration pas à pas

**Topologie :** 2× S310-24P4S en anneau (chaque switch relié aux deux voisins
par les ports de stack — en anneau pour 3–4 membres, en chaîne simple pour 2).

**Sur le futur master (SW-1) :**

```
system-view
stack slot 0 priority 200          # priorité la plus haute = master
quit
save
# Câble les ports de stack vers SW-2 (ports dédiés ou ports business configurés)
```

**Sur SW-2 :**

```
system-view
stack slot 1 priority 100
quit
save
```

Puis **éteins les deux**, câble les ports de stack, **allume le master d'abord**,
puis le second. Le stack se forme automatiquement ; le second adopte la config
du master.

> La syntaxe exacte (`stack slot`, `stack priority`, ports de stack à déclarer
> via `stack-port`) est **à vérifier sur la version logicielle du modèle exact** :
> elle varie selon que les ports de stack sont dédiés ou non.

🔧 Vérifications :

```
display stack
display stack topology
```

Tu dois voir les 2 membres, leurs rôles (master/standby/slave), et les liens
de stack UP.

---

## 96. iStack : vérifications, pannes et bonnes pratiques

**Pannes classiques :**

| Symptôme | Cause probable | Action |
|---|---|---|
| Le 2e switch ne rejoint pas | Versions logicielles différentes | Aligne les versions **avant** |
| Stack instable, bascules | Câble de stack défectueux / port business mal déclaré | Teste les câbles, vérifie `display stack topology` |
| Split-brain (2 masters) | Lien de stack coupé | **Prévois un lien de secours** (anneau, pas chaîne) ou la détection de conflit (MAD/dual-active detect si dispo — à vérifier) |
| Config du membre effacée | Normal : le slave adopte la config du master | Sauvegarde la config du slave **avant** de le stacker |

**Bonnes pratiques :**

1. **Topologie en anneau** dès 3 membres (pas de point unique de défaillance).
2. Même version logicielle **partout**, vérifiée avant chaque ajout.
3. Étiquette les câbles de stack (ils ne doivent **jamais** être débranchés
   « pour tester autre chose »).
4. En cas d'ajout d'un membre : éteins-le, câble, allume — jamais de branchement
   « à chaud » sur un stack en production sans procédure.

---

## 97. Cascading : l'alternative quand on ne stacke pas

Si tu ne veux/peux pas stacker (modèles différents, versions différentes,
simplicité), relie les switches en **cascade** : uplink du switch 2 vers un
port du switch 1, en **trunk** avec RSTP.

```
# Sur SW-1, port vers SW-2 :
interface GigabitEthernet0/0/28
 port link-type trunk
 port trunk allow-pass vlan 10 20 30 50 99
 description CASCADE_VERS_SW2
quit
# RSTP actif des deux côtés (sections 44-45) : la redondance éventuelle est gérée
save
```

**Stack vs cascade :**

|  | iStack | Cascade |
|---|---|---|
| Management | 1 IP, 1 config | 1 IP par switch |
| Redondance inter-switch | Oui (trunk inter-membres) | Oui si double lien + RSTP |
| Complexité | Plus élevée (versions, câblage dédié) | Faible |
| Quand | Baie unique, besoin de redondance matérielle | Sites simples, modèles mixtes |

✅ **En PME standard : cascade + RSTP suffit largement.** Le stack se justifie
quand tu veux de l'agrégation inter-châssis ou une gestion unifiée poussée.

---

## 98. Dépannage : la méthode (à appliquer avant chaque cas ci-dessous)

1. **Définis le symptôme précisément :** qui, quoi, depuis quand, quoi qui a changé.
   « Ça ne marche plus » n'est pas un symptôme. « Le PC du bureau 8 n'a plus
   d'IP depuis ce matin 8h, après le brassage d'hier soir » en est un.
2. **Console ou SSH d'abord :** `display interface brief`, `display logbuffer`,
   `display clock` (l'heure est-elle bonne ?).
3. **Isole la couche :** physique (LED, câble) → lien (négociation) → VLAN →
   IP/DHCP → routage → applicatif. Ne saute pas d'étape.
4. **Un changement à la fois**, et note ce que tu fais. Deux changements
   simultanés = tu ne sauras jamais lequel a réparé (ou cassé).
5. **Vérifie ce qui a changé** : config récente (`display logbuffer`,
   historique), travaux, coupure électrique, nouvel équipement branché.
6. Si tu sèches après 30 minutes : **escalade** avec les infos collectées
   (logs, `display`...), pas avec « ça marche pas ».

🔧 **Le kit de survie du dépanneur :** câble console + adaptateur USB, 2
jarretières testées bonnes, un PC portable avec PuTTY/Wireshark, la doc
d'adressage du site, et ce guide.

---

## 99. Cas n°1 — Boucle réseau : tout le réseau est à genoux

**Symptômes :** plus rien ne répond (ou par intermittence), toutes les LED des
ports clignotent frénétiquement en même temps, le switch est injoignable ou
très lent.

**Diagnostic :**
1. En console : `display interface` → compteurs qui explosent, CPU à 100 %
   (`display cpu-usage`).
2. `display logbuffer` → tempête de changements d'état STP si STP est actif.
3. Cherche le coupable : **quoi de neuf ?** Un petit switch de bureau branché
   avec 2 câbles ? Un câble qui boucle deux prises murales ? Des travaux ?

**Solution :**
1. **Débranche les liens suspects un par un** jusqu'à ce que le réseau
   reparte (commence par les ajouts récents).
2. Une fois stable : active **RSTP** (section 44) + **BPDU guard** (47) +
   **storm control** (73) + **edge ports** (46) sur tous les ports utilisateurs.
3. Retrouve la boucle physique et supprime-la (ou transforme-la en vraie
   redondance avec RSTP, section 49).
4. Documente l'incident : cause, port, équipement, heure.

✅ **Prévention :** cette panne ne doit arriver **qu'une fois**. Après, les
protections ci-dessus sont obligatoires sur tous les sites.

---

## 100. Cas n°2 — Le VLAN ne passe pas sur le trunk

**Symptômes :** les PC du VLAN 20 ne se voient pas d'un switch à l'autre (ou ne
joignent pas la passerelle), alors que le VLAN 10 fonctionne sur le même lien.

**Diagnostic :**
1. Des deux côtés : `display port vlan <interface-trunk>` → le VLAN 20 est-il
   dans `allow-pass` **des deux côtés** ?
2. `display vlan 20` → le VLAN existe-t-il **des deux côtés** ?
3. PVID identique des deux côtés ?

**Solution :** ajoute le VLAN manquant :
```
interface GigabitEthernet0/0/28
 port trunk allow-pass vlan 10 20 30 50 99
```
**des deux côtés**, puis `save`.

⚠️ C'est **la** panne VLAN la plus fréquente. Le réflexe `display port vlan`
des deux côtés la résout en 2 minutes.

---

## 101. Cas n°3 — Le PoE ne monte pas

**Symptômes :** AP/caméra/téléphone éteint, LED PoE éteinte ou orange.

