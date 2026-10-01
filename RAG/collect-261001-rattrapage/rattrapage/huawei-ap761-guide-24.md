---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-24
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["energy", "ethernet", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2406, 2572]
sha256: 39cd610c5183d7e5a8f6acf96c506079163a7388bea543cea583dab75629c603
---

# Guide ultra-complet — Huawei eKit AP761

```
AP761 — L'ESSENTIEL
Wi-Fi 6 (ax), 2×2 2.4+5 GHz, 1.775 Gbps max
Ports : 1× GE RJ45 (PoE-In) + 1× SFP GE — COMBO, optique prioritaire
PoE : 802.3at/af — 17.7 W max — af = FONCTIONS RESTREINTES
Budget : réserver 25 W / AP sur le switch
Antennes : directionnelles 65° — 10 dBi (2.4) / 11 dBi (5)
IP68, -40 à +65 °C, 6 kV surtension
1024 clients max (512/radio), 16 SSID/radio
Dimensions : 200×200×69 mm, 1.91 kg
Réf : 02355VFB — EAN : 6901443451289
Modes : Fat / Fit / Cloud
Canaux : 2.4 GHz = 20 MHz, canaux 1/6/11
         5 GHz = 40/80 MHz, préférer 36-48 (non-DFS)
Puissance : régler pour -67 dBm en limite de zone
Sécurité : WPA2-PSK mini, WPA3-SAE si parc OK, jamais WEP
```

## 110. Pense-bête de poche — page 2 : CLI express

```
--- Lecture (ne modifient rien) ---
display version                    # modèle + firmware
display current-configuration      # config complète
display ip interface brief         # IP des interfaces
display interface brief            # état + débit négocié (100M ? 1G ?)
display wlan vap all               # SSID diffusés
display logbuffer                  # logs récents
display clock                      # heure (NTP ?)
display cpu-usage / display memory # charge

--- Configuration ---
system-view                        # mode config
sysname NOM                        # nommer l'AP
save                               # sauvegarder (TOUJOURS après modif)
reboot                             # redémarrer
?                                  # aide contextuelle (TON MEILLEUR AMI)
[TAB]                              # complétion

--- Diagnostic express ---
ping 192.168.10.1                  # la passerelle répond ?
display arp | include <ip>         # résolution ARP OK ?
display poe interface              # (côté switch) af ou at ?
```

## 111. Pense-bête de poche — page 3 : dépannage express

```
AP éteint          → PoE ? câble ? injecteur ? parafoudre ? budget switch ?
AP non vu (cloud)  → IP ? Internet ? firewall ? NTP ? DNS ?
SSID sans IP       → trunk VLAN ? DHCP ? scope plein ? snooping ?
Débit faible       → bande 2.4/5 ? lien 100M/1G ? largeur canal ?
                     airtime saturé ? client zombie ?
Coupures roaming   → trou couverture ? 802.11r ? même VLAN ?
Interférences      → corréler l'heure, changer de canal, passer en 5 GHz
DFS                → normal ! passer en 36-48 si critique
WPA3 casse clients → mode transition, SSID IoT séparé
Portail invisible  → couper VPN, DNS auto, http://neverssl.com
Surchauffe         → pare-soleil, pas de coffret étanche, vérifier orientation
Rogue              → classifier AVANT de contre-attaquer
```

**La question magique devant tout ticket :**
« UN client ou TOUS les clients ? »
→ Un = problème client. Tous (un AP) = problème AP/réseau.
  Tous (partout) = problème d'infra (DHCP/RADIUS/Internet).

## 112. Glossaire (A–Z)

**802.11a/b/g/n/ac/ax/be** : les générations Wi-Fi. a/b/g = historiques ; n = Wi-Fi 4 ; ac = Wi-Fi 5 ; ax = Wi-Fi 6 ; be = Wi-Fi 7. L'AP761 fait a/b/g/n/ac/ax [constructeur].

**802.11k** : le client reçoit la liste des AP voisins (aide au roaming). Chap. 56.

**802.11r** : roaming rapide (ré-authentification accélérée). Chap. 56.

**802.11v** : l'AP suggère au client de changer d'AP (BSS Transition). Chap. 56.

**802.11w (PMF)** : protection des trames de management contre la désauthentification forgée. Chap. 51.

**802.1Q** : le standard des VLAN taggés sur trunk. Chap. 52.

**802.1X** : authentification réseau par identifiant (EAP), avec RADIUS. Chap. 48.

**802.3af/at/bt** : PoE (15 W), PoE+ (30 W), PoE++ (60/90 W). L'AP761 : af/at, 17.7 W max [constructeur]. Chap. 7.

**ACL** : liste de contrôle d'accès — règles de filtrage. Chap. 55.

**Airtime** : le temps d'occupation du canal radio. Le Wi-Fi partage le temps, pas le débit. Chap. 90.

**Airtime fairness** : allocation équitable du temps de parole entre clients (au lieu d'un débit égal qui pénalise les rapides).

**AP** : point d'accès.

**Association** : le fait pour un client de se connecter à un AP (après authentification).

**Atténuation** : perte de signal (distance, murs, végétation, pluie).

**Balise (beacon)** : trame émise ~10×/s par l'AP pour annoncer le SSID. Chaque SSID actif = des beacons = de l'airtime consommé.

**Band steering** : encourager les clients bi-bande à utiliser le 5 GHz plutôt que le 2.4 GHz. Chap. 42.

**BLE** : Bluetooth Low Energy — ici en 5.2 sur l'AP761, pour l'onboarding et la maintenance de proximité [constructeur].

**BSS coloring** : « couleur » du réseau en Wi-Fi 6 : permet de distinguer son réseau de celui des voisins et de réduire l'attente. Chap. 12.

**CAC (Call Admission Control)** : limite le nombre d'appels voix simultanés pour garantir la qualité. Chap. 54.

**Canal** : la fréquence centrale d'émission. En 2.4 GHz : 1/6/11. En 5 GHz : 36–48 (non-DFS), 52–140 (DFS).

**CAPWAP** : protocole entre un AP en mode Fit et son contrôleur (AC). Chap. 32.

**Client / STA** : l'équipement qui se connecte (smartphone, PC, caméra).

**Combo (port)** : deux ports physiques (ici RJ45 + SFP) dont un seul est actif à la fois. Chap. 6.

**Désauthentification (attaque)** : forgeage de trames pour éjecter les clients — contré par le PMF/802.11w. Chap. 51.

**DFS** : détection radar obligatoire sur les canaux 5 GHz 52–140 ; l'AP doit quitter le canal si radar. Chap. 41.

**DHCP snooping** : ne laisse passer que les réponses DHCP des serveurs de confiance. Chap. 55.

**Downtilt** : inclinaison de l'antenne vers le bas pour couvrir une zone en contrebas.

**EAP** : protocole d'authentification (PEAP, EAP-TLS…) utilisé avec 802.1X. Chap. 48.

**Evil twin** : AP pirate usurpant un SSID légitime. Chap. 64.

**Fat / Fit / Cloud** : les trois modes de l'AP761 — autonome, piloté par contrôleur, piloté par le cloud [constructeur]. Chap. 11.

**Forward-mode** : direct (le trafic client sort en local sur l'AP) vs tunnel (remonte au contrôleur). Chap. 34.

**Goulot (bottleneck)** : le maillon qui limite le débit (souvent : le client, le câble à 100M, ou l'airtime).

**IP68** : étanche à la poussière (6) et à l'immersion temporaire (8) [constructeur]. Chap. 5.

**Isolation client (STA isolation)** : les clients d'un même SSID ne se voient pas. Chap. 53.

**LLDP** : protocole de découverte de voisinage (sert aussi à négocier la classe PoE).

**MLO** : Multi-Link Operation — Wi-Fi 7 : utiliser plusieurs bandes simultanément. **Non supporté par l'AP761.** Chap. 15.

**MRC** : Maximum Ratio Combining — combinaison des signaux reçus sur les antennes pour améliorer la réception [constructeur].

**MU-MIMO** : l'AP parle à plusieurs clients en même temps grâce à ses flux spatiaux [constructeur]. Chap. 12.

**Multi-RU** : Wi-Fi 7 : attribuer plusieurs blocs de sous-porteuses à un même client. **Non supporté par l'AP761.** Chap. 18.

**NAT** : traduction d'adresses (supportée par l'AP761 [constructeur], utile en mode routeur de poche).

**Noise floor (bruit)** : le niveau de bruit ambiant ; le SNR = signal − bruit. Chap. 61.

**OFDMA** : découpage du canal en sous-porteuses pour servir plusieurs clients en même temps — le cœur du Wi-Fi 6 [constructeur]. Chap. 12.

**PIRE** : puissance isotrope rayonnée équivalente = puissance conduite + gain d'antenne. C'est elle qui est limitée par la réglementation.

**PMF** : voir 802.11w.

**PoE** : Power over Ethernet — alimenter l'AP par le câble réseau. Chap. 7, 25.

**Portail captif** : page d'authentification pour les invités. Chap. 49.

**Preamble puncturing** : Wi-Fi 7 : utiliser les morceaux propres d'un canal partiellement brouillé. **Non supporté par l'AP761.** Chap. 18.

**PSK** : clé pré-partagée (WPA2-PSK, WPA3-SAE). Chap. 45–46.

**QAM** : modulation — 1024-QAM (Wi-Fi 6, max de l'AP761 [constructeur]), 4096-QAM (Wi-Fi 7). Chap. 17.

