---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-1
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1, 127]
sha256: 93bc070955b17285f84e7d584fc0c03384857cd13dcfe91ee763ea062fc53e4a
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

> **Public** : administrateurs systèmes, chefs de service infrastructures (angle métier "systèmes & énergies" : un plan d'adressage fiable = moins de tickets, moins d'interruptions de production).
> **Versions couvertes** : Windows Server 2019, 2022 et 2025 (le rôle DHCP est stable depuis 2012 R2 ; je signale chaque différence de version).
> **Environnements de référence** : contoso.local, plage 192.0.2.0/24 (TEST-NET-1, RFC 5737) — adapte à ton plan d'adressage réel.
> **Prérequis** : AD DS déployé, DNS interne fonctionnel, serveur membre du domaine, IPv4 (IPv6 traité en section 12bis).

---

## Sommaire

| Bloc | Sections | Contenu |
|------|----------|---------|
| A — Fondations | 1–12 | Rappels DORA, rôle du DHCP, installation, autorisation AD |
| B — Étendues et baux | 13–30 | Création d'étendues, exclusions, durée des baux, réservations |
| C — Options DHCP | 31–44 | Options standard, PXE/WDS, précédence serveur/étendue/réservation |
| D — Scopes avancés | 45–58 | Superscopes, multicast, stratégies, filtres MAC, classes |
| E — Haute disponibilité | 59–68 | Basculement hot standby / load balance, relais DHCP, VLANs |
| F — DNS dynamique & sauvegarde | 69–76 | DNS dynamique, compte dédié, export/import, migration |
| G — Dépannage | 77–91 | 15 cas pratiques commentés, journaux, audit |
| H — Sécurité | 92–100 | Rogue DHCP, DHCP snooping, durcissement |
| I — Bonnes pratiques | 101–108 | Dimensionnement, durées de bail, redondance, documentation |
| Annexes | A1–A6 | Checklists, pense-bête de poche, glossaire, quiz, erreurs classiques, pour aller plus loin |

---

# Bloc A — Fondations

## 1. Pourquoi le DHCP est critique en entreprise

Sans DHCP, chaque poste, imprimante, téléphone IP et capteur IoT devrait être configuré à la main : adresse IP, masque, passerelle, DNS. À 50 machines c'est pénible, à 500 c'est ingérable, à 2000 c'est une source permanente d'incidents.

Le DHCP (Dynamic Host Configuration Protocol, RFC 2131/2132) automatise l'attribution :

- **Adresse IPv4** (et IPv6 via DHCPv6)
- **Masque de sous-réseau**
- **Passerelle par défaut** (option 003)
- **Serveurs DNS** (option 006)
- **Nom de domaine DNS** (option 015)
- Et des dizaines d'options métier : PXE (066/067), NTP (042), WINS, proxy WPAD (252)...

En entreprise, le DHCP est presque toujours couplé à **Active Directory** (autorisation obligatoire, voir section 9) et au **DNS dynamique** (voir section 69). Un DHCP mal dimensionné ou non redondé = des utilisateurs qui ne peuvent plus se connecter au réseau au renouvellement de leur bail. C'est un service d'infrastructure au même niveau que DNS et AD.

## 2. Le cycle DORA : Discover, Offer, Request, Acknowledge

Tout échange DHCP tient en 4 messages. À connaître par cœur pour le dépannage :

```text
Client (0.0.0.0:68)                    Serveur DHCP (*:67)
       |                                         |
       |--- 1. DHCPDISCOVER (broadcast) --------->|
       |     "Qui peut me donner une adresse ?"   |
       |                                         |
       |<-- 2. DHCPOFFER (broadcast/unicast) ----|
       |     "Je propose 192.0.2.50, masque      |
       |      255.255.255.0, bail 8 jours"       |
       |                                         |
       |--- 3. DHCPREQUEST (broadcast) --------->|
       |     "J'accepte 192.0.2.50"               |
       |     (broadcast pour prévenir les        |
       |      autres serveurs qui ont offert)    |
       |                                         |
       |<-- 4. DHCPACK (broadcast/unicast) ------|
       |     "OK, c'est à toi pour 8 jours"      |
       |                                         |
       |=== Le client configure son IP ==========|
```

Détails à retenir :

| Message | Émetteur | Destination | Port src/dst | Contenu clé |
|---------|----------|-------------|--------------|--------------|
| DHCPDISCOVER | Client | Broadcast 255.255.255.255 | 68 → 67 | Transaction ID, MAC client |
| DHCPOFFER | Serveur | Broadcast ou unicast | 67 → 68 | IP proposée, masque, durée du bail, options |
| DHCPREQUEST | Client | Broadcast | 68 → 67 | IP acceptée (ou demande de renouvellement) |
| DHCPACK | Serveur | Broadcast ou unicast | 67 → 68 | Confirmation + options finales |

> **Astuce dépannage** : dans Wireshark, filtre `bootp`. Si tu vois le DISCOVER mais pas le OFFER, le serveur ne répond pas (service arrêté, étendue épuisée, filtre MAC, ou le broadcast ne traverse pas — relais DHCP manquant, voir section 64).

Variantes du cycle à connaître :

- **Renouvellement (T1 = 50 % du bail)** : le client envoie un DHCPREQUEST en **unicast** directement au serveur qui lui a donné le bail. Pas de DISCOVER.
- **Rebinding (T2 = 87,5 % du bail)** : si le serveur ne répond pas, le client rebroadcast un DHCPREQUEST à n'importe quel serveur.
- **Expiration** : le client recommence un DORA complet et doit arrêter d'utiliser l'adresse.
- **DHCPNAK** : le serveur refuse (ex. : le client demande une adresse d'un autre sous-réseau après un déplacement). Le client doit redémarrer un DORA.
- **DHCPRELEASE** : le client libère volontairement son bail (`ipconfig /release`).
- **DHCPINFORM** : le client a déjà une IP (configurée manuellement) et demande **uniquement les options** (DNS, domaine...). Le serveur répond par un ACK sans attribuer d'adresse.

## 3. Ports, protocoles et flux réseau du DHCP

```text
UDP 67  : serveur DHCP (écoute)
UDP 68  : client DHCP (écoute)
```

Règles pare-feu Windows (créées automatiquement à l'installation du rôle) :

```powershell
# Vérifier les règles pare-feu DHCP créées par le rôle
Get-NetFirewallRule -DisplayGroup "DHCP Server" | 
    Select-Object DisplayName, Enabled, Direction, Action
```

À retenir pour les **pare-feu intermédiaires** (entre VLANs) :

- Autoriser UDP 67/68 **dans les deux sens** entre le serveur DHCP et les agents de relais.
- Le trafic client→serveur initial est en broadcast : il ne traverse pas les routeurs sans agent de relais (section 64).

## 4. DHCP vs IP statique : que mettre en DHCP en entreprise

| Type d'équipement | Recommandation | Méthode |
|-------------------|----------------|---------|
| Postes utilisateurs (fixe/portable) | DHCP dynamique | Étendue standard |
| Postes en mobilité (Wi-Fi) | DHCP, bail court | Étendue dédiée, bail 8–24 h |
| Invités (Wi-Fi guest) | DHCP isolé, bail très court | VLAN + étendue dédiée, bail 1–4 h |
| Imprimantes / MFP (copieurs) | **Réservation DHCP** | Réservation par MAC (section 22) |
| Serveurs | IP statique **ou** réservation | Statique hors plage DHCP de préférence |
| Contrôleurs de domaine / DNS | **IP statique obligatoire** | Hors plage DHCP, documenté |
| Serveur DHCP lui-même | **IP statique obligatoire** | Évidemment |
| Téléphones IP | DHCP + option 150/156 (selon constructeur) | Étendue voix dédiée (VLAN voix) |
| Switchs / points d'accès (management) | Réservation ou statique | Réservation conseillée |
| Caméras / IoT / GTB | DHCP, VLAN dédié, bail long | Étendue IoT, options spécifiques |

> **Lien métier copieurs** : les MFP (multifonctions) doivent avoir une **adresse IP fixe** car les pilotes d'impression, les carnets d'adresses LDAP, les dossiers de numérisation SMB et les flux de fax pointent vers cette IP. Si l'IP du copieur change, tout le service impression tombe. La **réservation DHCP** (section 22) est la méthode recommandée : IP fixe gérée centralement, sans aller configurer chaque copieur à la main.

## 5. Architecture type d'un DHCP d'entreprise

