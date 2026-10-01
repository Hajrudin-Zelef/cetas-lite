---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-12
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["arr", "diffusion"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1052, 1141]
sha256: 01581f784d9aec44998186dec91c0fabe152365d6e3908f1800986bcf699b54c
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

```
1. QUOI ?  -> Symptome exact, perimetre (1 client ? 1 AP ? tout le site ?)
2. QUAND ? -> Depuis quand ? Qu'est-ce qui a change juste avant ?
3. OU ?    -> Localiser : cloud eKit (quel equipement est rouge ?)
4. COUCHE PAR COUCHE (bas -> haut) :
   Physique (cable, PoE, LED) -> Liaison (lien up ?) -> IP (DHCP ?)
   -> Wi-Fi (SSID visible ?) -> Service (Internet ? portail ?)
5. UNE HYPOTHESE, UN TEST : ne change jamais 3 choses a la fois.
6. NOTER : symptome, cause, solution, dans le dossier du site.
```

**Le réflexe n°1 :** ouvrir l'app eKit et regarder **quel est l'état du site**. 50 % des diagnostics se font là, en 2 minutes.

## 92. Cas n°1 — Un équipement ne s'onboard pas (reste invisible dans l'app)

**Symptômes :** l'AP/switch/AR n'apparaît pas dans le site après scan ou déploiement Wi-Fi.
**Diagnostic :**
1. L'équipement est-il **sous tension** ? (LED allumée ?)
2. A-t-il **accès à Internet** ? (pour un AP derrière un switch : le switch a-t-il son uplink ? pour l'AR : le WAN est-il monté ?)
3. Le **SN scanné** est-il le bon ? (erreur de lecture, étiquette d'un autre carton — ça arrive plus souvent qu'on ne croit)
4. L'équipement est-il déjà onboardé sur **un autre site/compte** ? (matériel d'occasion ou de démo)
5. Firmware trop ancien pour le déploiement par Wi-Fi → passer par le **scan de code-barres** (cas prévu par la fiche technique).
**Solution :** corriger la cause (câble, uplink, re-scan du bon SN, suppression de l'ancien site), puis relancer l'onboarding. Si l'équipement vient d'un autre compte : le **retirer proprement** de l'ancien site d'abord (ou reset usine + ré-onboard — comportement exact **à vérifier sur la documentation officielle**).

## 93. Cas n°2 — Le cloud ne voit plus tout un site (tous les équipements hors ligne)

**Symptômes :** dans l'app, tout le site passe au rouge d'un coup. Les utilisateurs sur place disent « Internet ne marche plus » (ou parfois « ça marche encore en local »).
**Diagnostic :**
1. **Coupure électrique ?** Appeler le contact sur place : « est-ce que la baie est allumée ? »
2. **Lien opérateur tombé ?** (box/ONT éteint, travaux dans la rue)
3. **Passerelle AR plantée ?** (plus de ping, LED anormale)
4. Si le réseau **local fonctionne encore** (imprimante OK, fichiers OK) → c'est uniquement l'uplink Internet ou le lien cloud qui est coupé, pas le LAN.
**Solution :** selon la cause — faire réenclencher le disjoncteur, redémarrer la box opérateur **puis** l'AR (dans l'ordre : modem → AR → switch → AP), appeler l'opérateur si le lien reste down. **Ne jamais** reset les équipements eKit dans ce cas : le problème est en amont, pas dans leur config. Au retour : vérifier que tout repasse au vert dans le cloud (délai de quelques minutes normal).

## 94. Cas n°3 — AP adopté mais aucun SSID diffusé

**Symptômes :** l'AP est vert dans l'app, mais aucun SSID n'apparaît sur les téléphones.
**Diagnostic :**
1. Les SSID sont-ils **associés à cet AP** (ou au groupe de cet AP) dans la config du site ? (Erreur classique : SSID créés mais non appliqués aux nouveaux AP.)
2. L'AP a-t-il **reçu sa configuration** ? (forcer une re-synchronisation depuis le cloud — libellé exact **à vérifier** dans l'app.)
3. Les **radios** sont-elles activées ? (un AP en « monitor » ou avec radios désactivées ne diffuse rien.)
4. Conflit de **pays/réglementation** ? (canaux 6 GHz interdits localement sur un AP673H par ex. — vérifier la conformité locale.)
**Solution :** associer les SSID à l'AP, re-pousser la config, vérifier l'état des radios. Tester avec un téléphone à 2 mètres de l'AP (si ça marche à 2 m mais pas à 20 m, c'est un problème de couverture, pas de diffusion).

## 95. Cas n°4 — Clients connectés au Wi-Fi mais « pas d'Internet »

**Symptômes :** le Wi-Fi s'associe, mais pas de navigation. Le plus fréquent de tous les tickets.
**Diagnostic (dans l'ordre) :**
1. Le client a-t-il une **IP** ? (169.254.x.x = pas de DHCP → problème DHCP ou VLAN)
2. Le client ping-t-il la **passerelle** ? (non → problème L2/VLAN/SSID ; oui → problème routage/DNS/uplink)
3. Le **DNS** résout-il ? (`nslookup` vers un nom connu — si l'IP passe mais pas le DNS, c'est le DNS)
4. Le **portail captif** s'affiche-t-il ? (sur SSID invités : si le portail ne se déclenche pas, tester l'URL de détection captive du téléphone)
5. Un seul client ou tous ? (un seul → oublier le réseau sur le client et reconnecter ; tous → infra)
**Solution :** selon l'étage du problème — relancer le DHCP (scope plein ? bail trop long ?), corriger le VLAN du SSID, corriger les DNS sur l'AR, vérifier l'uplink Internet.

## 96. Cas n°5 — Wi-Fi « lent » (le ticket le plus flou)

**Symptômes :** « c'est lent » — sans plus de précision. Ne jamais partir en chasse sans mesurer.
**Diagnostic :**
1. **Mesurer** : speedtest à côté de l'AP, puis à l'endroit de la plainte. Noter les chiffres.
2. **Combien de clients** sur l'AP concerné ? (vue clients du cloud — 40 clients sur un AP361 = saturation, pas panne)
3. **Le lien filaire** de l'AP est-il en 1 Gbit/s ? (un AP Wi-Fi 6 bridé en 100 Mbit/s à cause d'un mauvais câble = débit plafond à ~90 Mbit/s)
4. **Interférences** : l'AP change-t-il de canal sans arrêt ? Un four à micro-ondes / un vieux AP voisin en 2,4 GHz ?
5. **L'uplink Internet** est-il saturé ? (tester en filaire direct sur l'AR pour isoler Wi-Fi vs Internet)
**Solution :** ajouter un AP si densité, remplacer le câble si lien 100 Mbit/s, figer les canaux si instabilité, limiter le débit invités si c'est eux qui saturent, augmenter le débit opérateur si l'uplink est le goulot.

## 97. Cas n°6 — Le portail captif ne s'affiche pas

**Symptômes :** connexion au SSID invités OK, mais pas de page d'accueil ; Internet « ne marche pas ».
**Diagnostic :**
1. Tester sur **2 téléphones différents** (iPhone + Android — les mécanismes de détection captive diffèrent).
2. Le client a-t-il une IP du VLAN invités ? (non → DHCP/VLAN)
3. Le DNS du VLAN invités fonctionne-t-il ? (la détection captive repose sur DNS + HTTP)
4. Un **VPN** actif sur le téléphone ? (un VPN bloque souvent la détection du portail — le couper pour le test)
5. Date/heure du téléphone correctes ? (un certificat HTTPS avec une date fausse = page bloquée)
**Solution :** corriger DHCP/DNS du VLAN invités, communiquer le **PSK invités de secours** (section 40) en attendant, et afficher en réception une petite affichette « Wi-Fi : si la page ne s'ouvre pas, ouvrez votre navigateur sur [adresse] » (adresse fictive à définir par site).

## 98. Cas n°7 — Un AP redémarre en boucle

**Symptômes :** l'AP apparaît/disparaît du cloud, les clients se déconnectent par vagues.
**Diagnostic :**
1. **PoE insuffisant ?** (cas n°1 : AP761 sur port 802.3af au lieu de at → fonctions limitées / redémarrages ; vérifier la classe PoE négociée)
2. **Budget PoE du switch dépassé ?** (ajout récent d'une caméra ? le switch coupe les ports les moins prioritaires)
3. **Câble défectueux** (un brin qui fait des micro-coupures = renégociations PoE en boucle → tester au testeur, remplacer le câble)
4. **Surchauffe** (AP dans un faux plafond à 50 °C ?)
5. **Firmware corrompu** (après une MAJ interrompue → re-flasher)
**Solution :** passer l'AP sur un port PoE+ / at, alléger le budget PoE, changer le câble, ventiler, re-flasher. **Ne jamais** laisser un AP « qui redémarre de temps en temps » : c'est toujours un symptôme qui empire.

## 99. Cas n°8 — Après une mise à jour, plus rien ne marche

