---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-18
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1797, 1906]
sha256: 895a82ff18fb18f6ff2884e2f7badef6a5f761446e52ba754e88d98056b5b927
---

# Guide ultra-complet — Huawei eKit AP761

**En mode Cloud (le plus simple) :**
1. Console cloud → site → AP → « Mise à jour firmware ».
2. Choisir la version (la console propose les versions validées).
3. Planifier : immédiate (site pilote) ou programmée (nuit).
4. Lancer, **surveiller** : l'AP télécharge, flashe, redémarre, se ré-adopte.
5. Vérifier : version (`display version` ou console), SSID diffusés, clients qui reviennent, logs propres.

**En mode Fat (CLI) :**
```
# 1. Déposer le firmware sur l'AP (SFTP) :
sftp admin@192.168.10.11
sftp> put AP761-V200Rxxx.cc /

# 2. Désigner le firmware de démarrage :
<AP761> startup system-software AP761-V200Rxxx.cc

# 3. Vérifier :
<AP761> display startup
# -> le nouveau firmware doit être « startup system software »

# 4. Sauvegarder la config, puis redémarrer :
<AP761> save
<AP761> reboot
# -> confirmer, attendre 3-5 min, vérifier la version :
<AP761> display version
```

**Checklist post-mise à jour :**
- [ ] Version correcte affichée.
- [ ] Tous les SSID diffusés (vérifier les deux radios).
- [ ] Clients tests connectés sur chaque SSID.
- [ ] Supervision verte (SNMP/cloud).
- [ ] Logs sans erreurs nouvelles (10 min d'observation).
- [ ] **Noter dans le dossier de site** : ancienne → nouvelle version, date, opérateur.

## 76. Rollback : revenir en arrière proprement

**Quand faire un rollback :** régression avérée après mise à jour (clients qui ne s'associent plus, débit effondré, redémarrages en boucle — cas 16). Pas sur une simple impression : **mesurer d'abord** (comparer les indicateurs chap. 67 avant/après).

**Procédure (Fat AP) :**
```
# L'ancien firmware est normalement conservé en backup :
<AP761> display startup
# -> repérer l'ancien firmware (ex. AP761-V200Ryyy.cc)

<AP761> startup system-software AP761-V200Ryyy.cc
<AP761> save
<AP761> reboot
# -> vérifier le retour à l'ancienne version
```

**En mode Cloud :** la console permet généralement de re-flasher une version antérieure (fonction « rollback » ou réinstallation d'une version — à vérifier sur la fiche du modèle exact).

**Règles du rollback :**
1. **Ne pas paniquer :** un rollback précipité sans diagnostic fait perdre l'information. Noter les symptômes AVANT.
2. **La config a peut-être migré** : si le nouveau firmware a converti la config dans un nouveau format, l'ancien firmware peut la refuser partiellement → **restaurer la sauvegarde pré-mise à jour** (chap. 73) après le rollback.
3. **Signaler** : remonter la régression au support / au distributeur avec les logs — ça sert à tout le monde.
4. **Geler** les mises à jour du reste du parc tant que la cause n'est pas comprise.

**Prévention :** garder **toujours** l'avant-dernière version stable sous la main (fichier + procédure testée). Le jour où tu en as besoin, tu n'auras pas le temps de la chercher.

## 77. Dépannage : méthode générale (les 5 couches)

Avant les cas particuliers, LA méthode. 90 % des tickets Wi-Fi se résolvent en suivant cet ordre, sans sauter d'étape :

**Couche 1 — Physique :** L'AP est-il alimenté ? LED ? Le câble est-il bon ? Le PoE arrive-t-il ? Le switch voit-il le lien ?
**Couche 2 — Réseau filaire :** L'AP a-t-il une IP ? Ping OK ? Le VLAN est-il bon ? Le DHCP répond-il ? Le trunk laisse-t-il passer les VLAN ?
**Couche 3 — Radio :** Les SSID sont-ils diffusés ? Le canal est-il correct ? La puissance est-elle suffisante ? Y a-t-il de l'interférence ?
**Couche 4 — Association :** Le client s'associe-t-il ? L'authentification passe-t-elle (PSK ? 802.1X ? Portail ?) ? Obtient-il une IP ?
**Couche 5 — Service :** Le client navigue-t-il ? Le débit est-il correct ? Le DNS répond-il ? Internet est-il OK ?

**Les 4 questions à poser à chaque ticket :**
1. **Quoi** exactement ? (« pas de Wi-Fi » → quoi : pas de SSID ? pas d'IP ? lent ?)
2. **Qui** est impacté ? (un client ? tous ? une zone ? un SSID ?)
3. **Quand** ça a commencé ? (toujours ? depuis quand ? à heures fixes ?)
4. **Qu'est-ce qui a changé ?** (travaux ? nouvel AP voisin ? mise à jour ? orage ?)

**Règle d'or :** « qui est impacté » oriente tout. **Un seul client** = problème client (pilote, config). **Tous les clients d'un AP** = problème AP/réseau. **Tous les clients partout** = problème d'infrastructure (DHCP, RADIUS, Internet). Ne jamais reconfigurer l'infrastructure pour un seul client têtu.

## 78. Cas 1 : l'AP ne s'allume pas / pas de PoE

**Symptômes :** LED éteinte, l'AP n'apparaît nulle part, pas de lien sur le switch.

**Diagnostic ordonné :**
1. **Le switch/injecteur fournit-il du PoE ?** Vérifier le voyant PoE du port, la config du port (`display poe` côté switch Huawei), le budget restant (chap. 25). Un budget épuisé = le dernier AP ne démarre pas.
2. **Le câble est-il bon ?** Tester au testeur (continuité + longueur). Un câble trop long (> 100 m) ou en CCA fait chuter la tension : l'AP peut clignoter sans jamais démarrer (brownout).
3. **L'injecteur est-il gigabit et 802.3at ?** Un injecteur 100 Mbps ou 802.3af seul peut empêcher le démarrage complet (fonctions restreintes, voire boot en boucle).
4. **Le parafoudre est-il mort ?** (chap. 26) : le shunter temporairement pour tester. S'il a pris un coup, il coupe tout.
5. **Tester l'AP sur table** : l'amener au local, le brancher en direct sur un injecteur connu bon avec un cordon court. S'il démarre = le problème est le câble/l'infrastructure. S'il ne démarre pas = l'AP est probablement mort (SAV).

**Piège :** un AP qui **démarre puis s'éteint en boucle** = souvent une alimentation limite (tension qui s'effondre quand les radios s'allument et tirent du courant). Mesurer sous charge si possible, ou simplifier : cordon court + injecteur 30 W.

## 79. Cas 2 : l'AP démarre mais n'apparaît pas dans le cloud

**Symptômes :** LED verte, mais l'AP reste « hors ligne » dans la console cloud / l'app.

**Diagnostic ordonné :**
1. **L'AP a-t-il une IP ?** Le trouver via le DHCP ou en BLE avec l'app. Pas d'IP = problème VLAN/DHCP (couche 2, chap. 77).
2. **L'AP sort-il sur Internet ?** Depuis l'AP (si CLI accessible) : `ping 8.8.8.8`, puis `ping` du nom de la plateforme cloud. Pas de sortie = firewall/proxy qui bloque.
3. **Le firewall laisse-t-il passer ?** La plateforme cloud eKit utilise des ports/URL spécifiques (**à vérifier sur la fiche du modèle exact** — typiquement HTTPS 443 vers des domaines Huawei). Un firewall d'entreprise qui ne laisse sortir que le web via proxy = l'AP ne passe pas.
4. **L'heure est-elle juste ?** (chap. 70) : sans NTP, le TLS échoue. `display clock` puis comparer.
5. **Le DNS résout-il ?** `nslookup`/`ping` du domaine cloud depuis l'AP.
6. **L'AP est-il déjà adopté ailleurs ?** Un AP adopté sur un autre compte/site doit être **libéré** avant ré-adoption.

**Test décisif :** brancher l'AP sur un réseau « simple » (box internet + DHCP, sans firewall) : s'il s'adopte = le problème est ton réseau (firewall/VLAN/DNS). Sinon = problème AP/compte.

## 80. Cas 3 : SSID visible mais pas d'IP (DHCP)

**Symptômes :** le client s'associe au Wi-Fi mais reste en 169.254.x.x (APIPA) ou « connecté, pas d'Internet ».

