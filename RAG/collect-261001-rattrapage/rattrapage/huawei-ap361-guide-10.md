---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-10
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "ethernet", "voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1476, 1670]
sha256: f83d66b42434eb45e044ee17824e2f3ac701ce325040595cf523fb039af821fa
---

# Huawei eKit AP361 — Guide ultra-complet

- **Parade** : PMF / 802.11w (§70). Avec PMF required, les trames deauth forgées
  sont rejetées.
- **Détection** : pics de déconnexions simultanées sur un AP = suspect.
- En WPA2 sans PMF : vous êtes vulnérable. C'est un argument de plus pour
  migrer vers WPA3 ou activer PMF.

## 93. Isolation client (client isolation) : quand l'activer

L'isolation empêche les clients d'un même SSID de communiquer **entre eux**
(ils ne voient que la passerelle).

| SSID | Isolation | Pourquoi |
|---|---|---|
| Invité | ✅ **Toujours** | Les visiteurs n'ont rien à se dire |
| IoT | ✅ Recommandé | Une caméra piratée ne doit pas attaquer la suivante |
| Bureau | ⚠️ Selon besoin | Casse le partage local (imprimantes, Chromecast, partage Windows) |

**Piège** : activer l'isolation sur le SSID bureau un vendredi soir = lundi
matin, « l'imprimante réseau ne marche plus » et « je ne vois plus le PC de
Paul ». Testez les usages **avant**.

## 94. PMF : rappel opérationnel

Déjà vu au §70. En résumé d'exploitation :

- WPA3 → PMF required (imposé).
- WPA2 → PMF optional si le parc le supporte, sinon disabled (noté au dossier).
- Un client qui **s'associe puis est immédiatement éjecté** après un changement
  PMF = incompatibilité PMF → revenez en arrière pour ce SSID.

## 95. Sécurité du management : l'AP lui-même est une cible

- L'AP a une IP, un OS, des services : c'est un équipement réseau à part entière.
- VLAN management dédié (§73), pas d'accès depuis les VLANs utilisateurs.
- Comptes nominatifs, mots de passe robustes (§144).
- HTTPS uniquement (§146), SSH restreint (§145).

## 96. Checklist sécurité Wi-Fi

- [ ] Détection rogue activée et qualifiée (pas de « rogue » non traité > 48 h)
- [ ] PMF : required (WPA3) / optional ou disabled documenté (WPA2)
- [ ] Isolation client : active sur Invité et IoT, choix documenté sur Bureau
- [ ] Pas de WEP, pas de WPA-TKIP (obsolètes, cassés)
- [ ] Clés PSK robustes + rotation planifiée
- [ ] Ronde Wi-Fi trimestrielle à l'analyseur (nouveaux SSID suspects ?)

---
---

# K. QOS / WMM

## 97. WMM : les 4 files d'attente

Le WMM (Wi-Fi Multimedia) classe le trafic en 4 catégories d'accès, par priorité
décroissante :

| File (AC) | Usage | Priorité |
|---|---|---|
| AC_VO (Voice) | Voix (téléphonie) | La plus haute |
| AC_VI (Video) | Vidéo (visio) | Haute |
| AC_BE (Best Effort) | Bureautique courante | Normale |
| AC_BK (Background) | Mises à jour, backup | La plus basse |

**C'est actif par défaut** sur la plupart des équipements Wi-Fi : ne le désactivez
jamais sans raison valable.

## 98. Marquer le trafic : DSCP et 802.1p

L'AP mappe les marques du monde filaire vers les files WMM :

| DSCP | 802.1p (CoS) | File WMM | Usage |
|---|---|---|---|
| EF (46) | 5 | AC_VO | Voix |
| AF41 (34), CS4 | 4 | AC_VI | Vidéo |
| 0 (BE) | 0 | AC_BE | Défaut |

**À faire côté filaire** : votre switch/téléphonie doit **marquer** (ou faire
confiance aux marques) en entrée. Si rien n'est marqué, tout tombe en Best
Effort et la QoS Wi-Fi ne sert à rien. Vérifiez la chaîne complète :
téléphone → switch → AP → air.

## 99. Prioriser la voix et la visio : recette

1. **Marquage** : les téléphones marquent EF (souvent par défaut) ; vérifiez.
2. **Confiance** : sur le port du switch vers l'AP, `trust dscp` (ou cos).
3. **SSID voix dédié** (optionnel mais propre) : QoS maximale, 5 GHz, 802.11r.
4. **Limitez le reste** : plafond de débit sur le SSID invité pour qu'un
   streaming 4K d'invité ne mange pas l'airtime de la visio du directeur.
5. **Testez en charge** : appel + speedtest simultanés → la voix doit rester claire.

Exemple indicatif côté switch Huawei (trust DSCP sur port AP) :

```
[SW-GigabitEthernet0/0/5] trust dscp
```

## 100. Limitation de débit (rate limiting) par SSID / par client

Utile pour :

- **Invités** : 10 Mbit/s down / 5 Mbit/s up par client (suffit pour mail/web,
  décourage le streaming massif).
- **IoT** : 2-5 Mbit/s (une caméra n'a pas besoin de plus en général — sauf
  4K, à dimensionner).
- **Bureau** : généralement pas de limite, ou limite haute anti-abus
  (ex. 100 Mbit/s).

🔎 Vérifiez les granularités offertes par votre version eKit (par SSID, par
client, montants/descendants séparés).

## 101. Admission control : le videur de la boîte

L'admission control refuse de nouveaux clients voix/vidéo quand la cellule est
saturée, plutôt que de dégrader tout le monde. C'est un réglage avancé :
à n'activer que si vous avez un vrai besoin voix dense (centre d'appels sur Wi-Fi)
et après mesure.

## 102. Erreurs QoS classiques

| Erreur | Symptôme |
|---|---|
| Rien n'est marqué côté filaire | La QoS Wi-Fi ne change rien (tout en BE) |
| Trust boundary mal placé | Les marques des clients sont effacées au switch |
| Pas de limite sur l'invité | Un invité en streaming 4K sature la cellule |
| WMM désactivé « pour tester » | Tout le monde en BE, voix hachée — et on oublie de le réactiver |

## 103. Checklist QoS

- [ ] WMM actif (par défaut — vérifié, pas supposé)
- [ ] Chaîne de marquage vérifiée de bout en bout (téléphone → AP)
- [ ] `trust dscp` (ou équivalent) sur les ports AP du switch
- [ ] SSID invité : débit plafonné
- [ ] Test en charge voix + données : OK

---
---

# L. SUPERVISION

## 104. Superviser via eKit : le quotidien

Le cloud/l'app eKit est votre écran de contrôle principal :

**Quotidien (2 minutes)** :

- Tous les AP verts ? (Si un est gris : voir §124.)
- Alarmes non acquittées ? (Les traiter ou les qualifier.)
- Un AP avec un nombre de clients anormalement haut/bas par rapport à d'habitude ?

**Hebdomadaire (15 minutes)** :

- Courbe de clients par AP : un AP saturé régulièrement = redécoupage de couverture.
- Taux d'échec d'association : en hausse = interférences ou problème de clé/serveur.
- Versions firmware : homogènes ?
- Espace disque / état du cloud : RAS.

## 105. Alertes : lesquelles configurer

| Alerte | Seuil indicatif | Action |
|---|---|---|
| AP hors ligne | Immédiat | Intervention (PoE/réseau) |
| Taux d'échec d'asso > 10 % | 15 min glissantes | Diagnostic interférences/auth |
| Utilisation canal > 70 % | 1 h | Rééquilibrage canaux/puissance |
| Rogue AP détecté | Immédiat | Qualification (§90) |
| Firmware non homogène | Hebdo | Planifier upgrade |

🔎 Les seuils configurables dépendent de la version eKit : adaptez ce tableau à
ce que l'outil permet réellement.

## 106. SNMP : la supervision « à l'ancienne » qui dépanne

Si l'AP361 expose SNMP (🔎 à vérifier selon version — souvent SNMP v2c/v3 en
lecture seule sur les AP Huawei) :

1. Activez SNMP **v3** (auth + priv) — jamais v1/v2c en production sauf
   contrainte (communautés en clair).
2. Restreignez l'accès SNMP au VLAN management et à l'IP du superviseur.
3. Intégrez à votre supervision existante (**Zabbix** — voir votre guide Zabbix) :
   - `sysUpTime`, état des radios, nombre de clients associés.
   - État du port Ethernet (up/down, vitesse négociée).
4. MIB Huawei : `HUAWEI-WLAN-*` (🔎 vérifiez les MIB supportées par le firmware).

Exemple indicatif (syntaxe VRP classique si CLI dispo) :

```
[AP361] snmp-agent sys-info version v3
[AP361] snmp-agent group v3 MONITEUR authentication
[AP361] snmp-agent usm-user v3 UserZabbix group MONITEUR auth sha Fictif-Auth-2026! priv aes Fictif-Priv-2026!
[AP361] snmp-agent target-host trap address udp-domain 192.168.10.50 params securityname UserZabbix v3 authentication
```

> ⚠️ Mots de passe **fictifs** ci-dessus. En vrai : 20+ caractères au coffre.

## 107. Syslog : centraliser les journaux

