---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-9
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [757, 900]
sha256: 607dd697c7ad87c41f2c7635b1f5c26917140564612d65e261bc26c9d4550924
---

# Guide ultra-complet — Huawei eKit AP761

**Procédure :**
1. Installer l'app HUAWEI eKit, créer un compte (ou se connecter).
2. Créer un **site** : nom explicite (`Client_Dupont_Siege`, `Usine_Nord`) — un site = un bâtiment ou un ensemble géré ensemble.
3. **Ajouter un équipement → scanner le QR code** de l'étiquette de l'AP (d'où l'intérêt de la photo, chap. 5).
4. L'app détecte l'AP en **BLE** à proximité : elle lit son état (démarré ? IP obtenue ? version firmware ?).
5. L'AP est **adopté** dans le site : il télécharge sa configuration et son firmware si besoin, puis redémarre.
6. Vérifier dans l'app : AP « en ligne », vert.

**Si le QR ne passe pas :**
- Saisie manuelle du numéro de série (sur l'étiquette).
- Vérifier que le smartphone a le Bluetooth activé et les autorisations de localisation (le BLE en a besoin sur Android).
- Se placer à moins de quelques mètres de l'AP.

**Bonnes pratiques :**
- Nommer l'AP dès l'adoption : `AP761-Cour-Est`, pas `AP-3F2A`. Un nom qui dit **où** il est.
- Renseigner la **position sur le plan** du site dans l'app si la fonction existe : dans 2 ans, tu sauras quel AP est où sans monter sur le toit.
- **Un seul administrateur « propriétaire »** du site + des comptes techniciens : éviter que 5 personnes se battent sur la même config.

## 30. Mode cloud : créer un site, adopter l'AP

Une fois l'AP adopté (chap. 29), la configuration se fait dans la console cloud (web ou app). Logique eKit : on configure des **modèles** (SSID, radio) appliqués au site, pas AP par AP.

**Structure logique type :**
```
Site « Usine_Nord »
├── Réseau filaire (VLAN, DHCP)
├── Wi-Fi
│   ├── SSID « USINE-PROD » (WPA2-Enterprise, VLAN 20)
│   ├── SSID « USINE-INVITE » (Portal, VLAN 30)
│   └── Paramètres radio (canaux, puissance)
└── Équipements
    ├── AP761-Cour-Est (adopté, en ligne)
    └── AP761-Parking (adopté, en ligne)
```

**Paramètres à régler au niveau du site (avant de brancher les clients) :**
1. **Pays/région réglementaire** : conditionne les canaux et puissances autorisés. **À régler en premier** — un mauvais domaine réglementaire = canaux interdits ou puissance bridée (chap. 41).
2. **SSID** : nom, sécurité (chap. 43–49), VLAN (chap. 52).
3. **Radio** : canaux 2.4/5 GHz, largeurs, puissances (chap. 37–40).
4. **NTP, DNS** : pour les logs et le TLS (chap. 70).

**Le cloud en panne, et le Wi-Fi ?** Question classique : **le Wi-Fi local continue de fonctionner** si la connexion au cloud tombe (l'AP garde sa dernière configuration). Ce qu'on perd : l'administration à distance, les stats temps réel, les alertes. Au retour d'Internet, l'AP se resynchronise. **Mais** : ne pas planifier de mise en service un jour où la fibre du site n'est pas encore active — l'adoption initiale a besoin d'Internet.

## 31. Mode Fat : accès web local et CLI

En mode **Fat**, l'AP761 est autonome : on l'administre directement, en web ou en CLI. C'est le mode du site isolé sans Internet.

**Accès web :**
1. Brancher un PC sur le même réseau que l'AP (ou directement sur l'AP via un switch).
2. Trouver l'IP de l'AP : via le serveur DHCP (chercher le hostname/MAC) ou via l'app eKit en BLE.
3. Ouvrir `https://<ip-ap>` (accepter l'avertissement de certificat autosigné — normal).
4. S'identifier avec le compte admin (identifiants par défaut : **à vérifier sur la fiche du modèle exact** — les changer **immédiatement**, chap. 106).

**Accès CLI (SSH) :**
```
# Depuis un poste du même réseau :
ssh admin@192.168.1.10
# Mot de passe : <celui défini> (jamais le défaut en production)
<AP761> display version          # vérifier modèle + version firmware
<AP761> display ip interface brief
<AP761> system-view              # passer en mode configuration
[AP761] sysname AP761-Cour-Est   # nom explicite, tout de suite
```

**Commandes de lecture indispensables (ne modifient rien) :**
```
<AP761> display current-configuration   # la config complète
<AP761> display wlan ap all             # état des radios/AP
<AP761> display wlan sta                # clients connectés (selon version)
<AP761> display interface brief
<AP761> display logbuffer               # journaux récents
```

> ⚠️ Les noms de commandes `display wlan ...` varient selon la version logicielle : utilise **la touche `?`** et la complétion par tabulation. Exemple : `display wlan ?` liste les options disponibles sur TA version. C'est la méthode la plus fiable, bien plus que n'importe quel guide.

## 32. Mode Fit : raccordement à un AC (WAC)

En mode **Fit**, l'AP761 est une « antenne intelligente » pilotée par un **contrôleur (AC)** — dans l'écosystème eKit, typiquement un **WAC** (ex. AC650 ou passerelle AR700 avec fonction WAC — à vérifier sur la fiche du modèle exact pour la compatibilité).

**Principe :** l'AP démarre, obtient une IP par DHCP, **découvre l'AC** (option DHCP 43, DNS, ou adresse statique), établit un tunnel **CAPWAP**, télécharge sa configuration, et remonte ses clients à l'AC.

**Découverte de l'AC — les 3 méthodes :**
1. **DHCP option 43** (la plus propre) : le serveur DHCP annonce l'IP de l'AC aux AP.
2. **DNS** : l'AP résout un nom prédéfini (ex. `huawei-wac` — à vérifier sur la fiche du modèle exact).
3. **Statique en CLI** :
```
<AP761> system-view
[AP761] capwap source interface vlanif 10
[AP761] controller ip address 192.168.10.5
```

**Vérifications côté AP :**
```
<AP761> display capwap status       # état du tunnel vers l'AC
<AP761> display wlan ap all          # l'AP doit apparaître « normal » côté AC
```

**Quand choisir Fit plutôt que Cloud :**
- Tu as déjà un AC et un parc existant.
- Tu veux l'**analyse de spectre** (seul le mode Fit la supporte sur l'AP761 — chap. 65).
- Le site n'a pas d'Internet fiable mais a un AC local.
- Politique interne : pas de management cloud (données sensibles).

**Inconvénient :** l'AC est un point central — s'il tombe, les AP en Fit perdent leur cerveau (les clients déjà connectés continuent souvent en local selon la config, mais plus de nouvelles associations ni de roaming propre). Prévoir la redondance d'AC sur les sites critiques — à vérifier sur la fiche du modèle exact.

## 33. Changer de mode : Fat ↔ Fit ↔ Cloud

On ne change pas de mode « pour voir » : chaque bascule **redémarre l'AP et réinitialise sa configuration locale**. Procédure avec fenêtre de maintenance.

**Principe général (CLI, depuis le mode Fat) :**
```
<AP761> system-view
[AP761] ap-mode cloud     # vers le cloud eKit
# ou
[AP761] ap-mode fit       # vers contrôleur (CAPWAP)
# ou
[AP761] ap-mode fat       # retour autonome
[AP761] quit
<AP761> save              # sauvegarder avant redémarrage
<AP761> reboot
```

**Points d'attention :**
- **Sauvegarder la config** avant (`save` + export, chap. 72) : la bascule efface.
- Vérifier que le **firmware supporte le mode cible** : certains firmwares eKit sont « cloud-only » par défaut — à vérifier sur la fiche du modèle exact.
- Après bascule vers Fit : l'AP doit trouver l'AC (chap. 32) sinon il boucle en découverte — prévoir l'option DHCP 43 **avant**.
- Après bascule vers Cloud : adoption par QR/BLE (chap. 29).
- **Ne jamais basculer à distance un AP dont on n'a pas d'accès physique de secours** : si la bascule échoue, c'est l'échelle.

## 34. Configuration initiale minimale (recette Fat AP)

La recette « ça marche en 15 minutes » en mode Fat, en CLI. Adapter les noms, VLAN et clés (clés **fictives** ci-dessous — génère les tiennes).

```
<AP761> system-view
[AP761] sysname AP761-Cour-Est

# --- Management ---
[AP761] vlan batch 10 20 30
[AP761] interface vlanif 10
[AP761-Vlanif10] ip address 192.168.10.11 24
[AP761-Vlanif10] quit
[AP761] ip route-static 0.0.0.0 0.0.0.0 192.168.10.1

