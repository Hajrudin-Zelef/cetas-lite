---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-15
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [3041, 3204]
sha256: 19c97ae7ef097f2292d6f2febe4e1d8b130a2f2ac48815106c510fc404ed685c
---

# Guide ultra-complet — Huawei eKit S310

## 111. Cas n°13 — 802.1X rejette tout le monde

**Symptômes :** après activation, aucun poste ne passe (ou certains seulement).

**Diagnostic :**
1. Le **serveur RADIUS** est-il joignable ? (`ping`, `display radius-server`.)
2. **Clé partagée** identique des deux côtés ? (Erreur n°1.)
3. Les postes ont-ils un **supplicant** configuré (Windows : service « Config
   automatique de réseau câblé ») ?
4. Logs RADIUS côté serveur : que dit le rejet ? (mauvais login, certificat...)

**Solution :** corrige la clé, vérifie le réseau vers le RADIUS, déploie en
**mode pilote** (quelques ports) avant généralisation. **Prévois le bypass
d'urgence** : procédure écrite pour désactiver 802.1X port par port en cas de
panne du RADIUS.

⚠️ Un RADIUS en panne + 802.1X strict = **tout le site coupé**. Soit un RADIUS
redondé, soit un mode « fail-open » documenté et assumé (à vérifier sur la
version exacte).

---

## 112. Cas n°14 — Mise à jour firmware qui échoue

**Symptômes :** le transfert échoue, ou le switch ne boote pas sur la nouvelle
image, ou des fonctions ne marchent plus après.

**Diagnostic :**
1. Image **compatible** avec le modèle exact ? (Notes de release lues ?)
2. Transfert **complet** ? (Taille du fichier en flash = taille attendue.)
3. `display startup` : la **bonne** image est-elle marquée pour le prochain boot ?
4. Assez de **place en flash** ? (`dir flash:/`.)

**Solution :** re-transfère (en binaire, pas en ASCII !), libère de la place,
re-déclare l'image, reboot. Si dysfonctionnement après boot : **rollback**
(section 92). **Jamais de mise à jour sans sauvegarde + image de secours.**

---

## 113. Cas n°15 — Ventilateur bruyant / surchauffe

**Symptômes :** bruit anormal (sifflement, claquement), LED d'alarme
température, ou switch qui reboot tout seul par forte chaleur.

**Diagnostic :**
1. `display temperature` / `display fan` (noms à vérifier sur la version exacte) :
   température, vitesse, état des ventilateurs.
2. Environnement : local à combien ? Dégagements respectés (section 12) ?
   Poussière ?
3. Un ventilateur **bloqué** (poussière, câble qui touche) ?

**Solution :** dépoussière (bombe d'air sec, switch **éteint**), dégage les
aérations, baisse la température du local. Si un ventilateur est HS : les
ventilateurs du S310 sont **intégrés non remplaçables** (valeur datasheet :
« built-in ») → c'est un **retour SAV**. D'où l'importance du **switch de
secours** (section 116).

⚠️ Un switch qui reboot par chaleur **abîme** le matériel à chaque cycle.
Ne laisse pas traîner : déplace-le ou ventile le local.

---

## 114. Cas n°16 — Intermittences inexpliquées (le pire)

**Symptômes :** coupures brèves aléatoires, sur un ou plusieurs ports, sans
cause évidente. Le cauchemar du dépanneur.

**Méthode (ordonnée) :**
1. **Heure exacte** des coupures (logs + témoignages) : corrélée à quelque
   chose ? (Clim qui démarre, machine industrielle, onduleur qui commute...)
2. `display interface` : **compteurs d'erreurs** sur les ports touchés.
   Erreurs qui montent = couche 1 (câble, interférences).
3. **Alimentation** : micro-coupures ? Le switch est-il sur onduleur ?
   (`display logbuffer` : reboots inexpliqués ?)
4. **Boucle partielle** : un équipement qui renvoie du trafic par intermittence
   (borne Wi-Fi en mode répéteur sauvage, double branchement).
5. **Surchauffe** cyclique (clim coupée la nuit/week-end ?).
6. Isole : déplace l'équipement sur un autre port/switch pour voir si le
   problème **suit l'équipement** ou **reste au port**.

✅ **Documente tout**, même les fausses pistes. Les intermittences se résolvent
par élimination, et l'historique écrit vaut de l'or quand le problème revient
3 mois plus tard.

---

## 115. Cas n°17 — L'AP Wi-Fi ne diffuse plus après un redémarrage du switch

**Symptômes :** après reboot/mise à jour, un ou plusieurs AP restent éteints ou
ne redémarrent pas.

**Diagnostic :**
1. Le PoE est-il revenu ? (`display poe interface`.) Avec le **Perpetual PoE**,
   il n'aurait pas dû couper (section 59) — sauf coupure **électrique** réelle.
2. **Budget** : au redémarrage, tous les AP appellent en même temps (pic) →
   si le budget est tout juste, certains ne montent pas. `display poe
   power-state` juste après boot.
3. L'AP a-t-il perdu sa config ? (Problème côté AP/contrôleur, pas switch.)

**Solution :** `poe power-off` / `poe power-on` ciblé sur les AP récalcitrants,
augmente la marge du budget (section 62), vérifie les priorités (section 61).
Si l'AP a perdu sa config : re-provisionne via l'app eKit.

---

## 116. Cas n°18 — Tempête de broadcast venant d'un équipement défaillant

**Symptômes :** comme une boucle (section 99) mais **sans** boucle physique :
un équipement (caméra HS, carte réseau défaillante, boucle logicielle dans un
AP) inonde le réseau.

**Diagnostic :**
1. `display interface` : **un** port avec un compteur d'output/input anormalement
   élevé par rapport aux autres = le suspect.
2. Miroir (section 85) + Wireshark : quel type de trafic ? (ARP en boucle ?
   Broadcast UDP ?)
3. `display mac-address` : une MAC qui « flappe » entre plusieurs ports ?

**Solution :** `shutdown` le port suspect → le réseau repart → isole et remplace
l'équipement. Puis active le **storm control** (section 73) pour que la prochaine
fois, le port se coupe tout seul.

✅ **Le storm control transforme une panne réseau totale en panne d'un seul
port.** C'est exactement pour ça qu'on l'active en préventif.

---

## 117. Maintenance préventive : le plan annuel

Un switch ne demande presque rien — jusqu'au jour où il demande tout. Le plan
ci-dessous coûte quelques heures par an et évite la plupart des pannes.

| Fréquence | Actions |
|---|---|
| **Mensuel** (10 min, à distance) | `display interface brief` (ports down anormaux), `display poe power-state` (budget), `display logbuffer` (erreurs récentes), `display cpu-usage` / mémoire, vérif alertes eKit |
| **Trimestriel** (30 min) | Sauvegarde config hors site (section 88), vérif NTP/heure, contrôle température locale, revue des ports inutilisés (à désactiver), mise à jour de la doc d'adressage |
| **Semestriel** (1 h) | Dépoussiérage (bombe d'air sec, switch éteint si très poussiéreux), vérif serrage terre et cordons, test d'un port de secours, revue des ACL et comptes (départs ?), test de restauration d'une sauvegarde sur maquette |
| **Annuel** (demi-journée) | Audit complet : firmware à jour ? (section 91), mots de passe renouvelés, inventaire physique vs doc, test du switch de secours (il boote ? sa config est à jour ?), exercice de panne (débranche un uplink : la redondance marche-t-elle vraiment ?), revue du plan avec l'équipe |

✅ **Règle :** chaque visite donne lieu à un **compte-rendu daté** d'une page
(état, actions, anomalies, prochaine échéance). Sans écrit, la maintenance
n'existe pas.

---

## 118. Checklist de visite sur site (à imprimer)

- [ ] État visuel : LED PWR/SYS vertes, pas de LED rouge, bruit de ventilateur normal
- [ ] Température du local : ___ °C (cible < 35 °C)
- [ ] Dépoussiérage fait (date : ___)
- [ ] Terre vérifiée (serrage visuel)
- [ ] `display interface brief` : aucun port down inexpliqué
- [ ] `display poe power-state` : budget dispo ___ W (alerte si < 20 %)
- [ ] `display logbuffer` : erreurs depuis la dernière visite ? ___
- [ ] Heure correcte : `display clock` → ___
- [ ] Sauvegarde config exportée (fichier : ___)
- [ ] Étiquetage ports à jour (nouveaux équipements décrits ?)
- [ ] Switch de secours : présent, config à jour (date : ___)
- [ ] Prochaine visite prévue le : ___

---

## 119. Pièces de rechange : le stock minimum

