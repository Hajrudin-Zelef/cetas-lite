---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-11
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1671, 1849]
sha256: 902638ab851efbf4a009df471c57c7344d4d0a895c7a0f2e92a44ac7ee5cbae6
---

# Huawei eKit AP361 — Guide ultra-complet

- Envoyez les logs vers votre serveur syslog central (celui de votre infra,
  ou un Graylog/Loki — voir vos guides).
- Niveau : `informational` en temps normal, `debugging` ponctuellement sur
  demande (jamais en continu : ça noie le serveur).
- **À logger impérativement** : associations/échecs, changements de config,
  upgrades, reboots, détections rogue.

Exemple indicatif :

```
[AP361] info-center loghost 192.168.10.60
[AP361] info-center source default channel loghost log level informational
```

## 108. Les KPI à suivre (tableau de bord)

| KPI | Cible | Où le voir |
|---|---|---|
| Disponibilité AP | > 99,5 % mensuel | eKit / Zabbix (ping + SNMP) |
| Clients max simultanés / AP | < 40-50 (AP361) | eKit |
| Taux d'échec d'association | < 5 % | eKit |
| Utilisation canal 5 GHz | < 60 % en heures pleines | eKit |
| Débit médian client (5 GHz) | > 50 Mbit/s en bureautique | Tests iperf périodiques |
| Temps de roaming (voix) | < 100 ms | Test d'appel en marchant |

> 📊 Un KPI sans cible, c'est de la déco. Fixez les cibles **avec** la direction
> (ou au moins informez-la) : le jour d'une plainte, vous sortez le graphe.

## 109. Zabbix : intégration type

Vous avez déjà un guide Zabbix : voici le template minimal « AP361 » :

- **Ping** : disponibilité (trigger : 3 échecs → alerte).
- **SNMP** : `sysUpTime` (reboot détecté si reset), nombre de clients par radio,
  état du port GE.
- **Syslog** : corrélation (reboot → chercher l'upgrade ou la coupure PoE).
- **Carte** : un « host » par AP, groupés par site, avec la topologie.

## 110. Rapport mensuel Wi-Fi (modèle)

Fournissez à la direction (ou gardez pour vous) un rapport d'une page :

1. Disponibilité du parc (%)
2. Top 3 des AP les plus chargés (+ action prévue si saturation)
3. Incidents du mois (date, durée, cause, résolution)
4. Firmwares : version déployée, homogénéité
5. Actions préventives du mois suivant

> 💼 Ce rapport, c'est votre assurance : le jour où « le Wi-Fi ne marche jamais »,
> vous répondez avec des chiffres, pas avec des impressions.

## 111. Checklist supervision

- [ ] Tous les AP dans eKit + dans Zabbix (double rattachement)
- [ ] Alertes configurées (tableau §105 adapté à l'outil)
- [ ] Syslog centralisé, niveau informational
- [ ] SNMP v3 restreint au VLAN management (si supporté)
- [ ] Rapport mensuel produit et archivé

---
---

# M. SAUVEGARDE, RESTAURATION, FIRMWARE

## 112. Sauvegarder la configuration : méthode

**En mode cloud** : la config vit dans le cloud eKit. Complétez par :

- Une **documentation versionnée** du site : SSID, clés (référence au coffre,
  jamais en clair), VLAN, plan radio, versions firmware.
- Des **captures d'écran** des pages de config critiques (ou export si dispo — 🔎).
- Le **tableau d'adressage** (AP, IP, MAC, SN, emplacement).

**En mode Fat** : sauvegarde classique :

```
<AP361> save                       # sauvegarde en flash (vrpcfg.cfg)
<AP361> backup configuration to 192.168.10.60  ap361-etage2.cfg
```

(🔎 Commande exacte selon version : `save`, puis transfert TFTP/FTP/SFTP vers
le serveur de sauvegarde.)

**Règle** : toute modification de config = sauvegarde + note au journal
(qui, quand, pourquoi). Ça prend 2 minutes et ça sauve des carrières.

## 113. Restaurer : procédure

1. Identifiez la **bonne** sauvegarde (date, version firmware compatible).
2. En Fat : transférez le fichier puis `startup saved-configuration <fichier>`
   + reboot (🔎 syntaxe selon version).
3. En cloud : réappliquez le template / restaurez depuis l'historique du site
   si la fonction existe.
4. **Vérifiez** : SSID diffusés, VLAN, test client complet (§31).
5. Si ça ne revient pas : reset d'usine (§26) + reconfiguration depuis la doc
   (d'où l'importance de la doc versionnée, pas seulement du fichier binaire).

## 114. Mise à jour firmware : procédure complète

**Avant** :

- [ ] Lire les **release notes** (correctifs sécu ? régressions connues ? prérequis ?)
- [ ] Sauvegarder la config (§112)
- [ ] Choisir la fenêtre : **hors heures ouvrées**, avec marge (prévoir 2x le temps estimé)
- [ ] Prévenir les utilisateurs (« maintenance Wi-Fi de 20 h à 22 h »)
- [ ] Vérifier l'espace flash et la stabilité PoE (pas d'upgrade pendant un orage)

**Pendant** :

1. Upgrader **1 AP pilote**, attendre le reboot complet.
2. Valider 24 h : associations, roaming, débit, logs (pas d'erreur nouvelle).
3. Généraliser **par vagues** (ex. un étage par soir), jamais tout le parc d'un coup.
4. Surveiller chaque vague (un AP qui ne revient pas en 15 min = intervention).

**Après** :

- [ ] Tous les AP sur la même version (vérifié un par un)
- [ ] Tests clients OK sur chaque zone
- [ ] Note au journal : version, date, AP concernés, incidents

## 115. Rollback : revenir en arrière

Si la nouvelle version pose problème (bug, incompatibilité client) :

1. **Ne paniquez pas** : la plupart des équipements Huawei conservent l'ancienne
   image en secours (🔎 vérifiez : `display version`, partition backup).
2. Procédure type (indicatif) : redémarrer sur l'image précédente via la
   commande de boot ou via l'app eKit (fonction « revenir à la version
   précédente » si elle existe — 🔎).
3. Si pas d'image de secours : réinstallez l'ancienne version (fichier conservé
   sur votre serveur — **gardez toujours les N-1**).
4. **Règle d'or** : ne supprimez l'ancienne image qu'après **2 semaines** de
   production stable sur la nouvelle.

## 116. Gérer les versions sur un parc multi-sites

| Principe | Détail |
|---|---|
| Une version de référence | La même partout (ex. `VxxxRxxxCxxSPCxxx`) |
| Fenêtre de validation | 2 semaines sur site pilote avant généralisation |
| Gel des upgrades | Pas d'upgrade en période critique (inventaires, examens, clôtures) |
| Registre | Tableau : site / AP / version / date d'upgrade / validé par |

## 117. Ce qu'il ne faut jamais faire

- ⛔ Upgrader pendant les heures ouvrées « parce que c'est rapide ».
- ⛔ Couper le PoE pendant l'écriture flash.
- ⛔ Upgrader sans avoir lu les release notes.
- ⛔ Laisser cohabiter 3 versions différentes « parce que ça marche ».
- ⛔ Supprimer l'image N-1 le jour même.

## 118. AP « briquée » après upgrade interrompu : sauvetage

Symptômes : LED rouge/éteinte anormale, pas d'IP, reset sans effet.

1. Laissez-la **10 minutes** (certains bootloaders tentent une récupération auto).
2. Tentez le reset long (§26).
3. Si l'AP expose un mode de secours (BootROM via console — 🔎 selon modèle,
   port console non confirmé sur AP361) : réinjection du firmware via TFTP.
4. Sinon : **RMA** (retour garantie). D'où l'importance d'avoir **1 AP de spare**
   en stock (voir §140).

## 119. Checklist maintenance logicielle

- [ ] Registre des versions tenu à jour
- [ ] Config sauvegardée avant chaque changement
- [ ] Images N et N-1 conservées sur le serveur
- [ ] 1 AP de spare en stock, sur la version de référence
- [ ] Fenêtres de maintenance planifiées au trimestre

---
---

# N. DÉPANNAGE TERRAIN — 16 CAS

## 120. Méthode : le modèle « Symptôme → Hypothèses → Test → Solution »

Avant les cas particuliers, la méthode qui marche à tous les coups :

