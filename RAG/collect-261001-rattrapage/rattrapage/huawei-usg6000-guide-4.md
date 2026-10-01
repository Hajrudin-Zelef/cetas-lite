---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-4
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [409, 577]
sha256: b6a6d1788ffbb488ca593e7761ecef846806504ac57c2e71ba919b4a3642b8a6
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

- La priorité (0–100) exprime la **confiance** : local (100) > trust (85) > dmz (50) > untrust (5).
- Elle intervient dans : l'**ordre d'évaluation implicite**, certaines fonctions de **détection d'attaque** (seuils par zone), et la **lisibilité** de la politique.
- ⚠️ La priorité **ne remplace pas** une politique explicite : avec une politique par défaut en `deny`, même trust→untrust est bloqué sans règle. Ne compte jamais sur la priorité pour « autoriser par défaut ».

## 26. Cas particulier : plusieurs interfaces dans la même zone

```huawei
[USG] firewall zone trust
[USG-zone-trust] add interface GigabitEthernet 1/0/3
[USG-zone-trust] add interface GigabitEthernet 1/0/4
[USG-zone-trust] add interface Vlanif 10
[USG-zone-trust] quit
```

- Le trafic **intra-zone** (entre deux interfaces de la même zone) : par défaut, il **passe** sans politique interzone sur beaucoup de versions — vérifie avec `display firewall interzone` / teste. Si tu veux le filtrer (ex. isoler deux départements), crée des zones distinctes ou active le filtrage intra-zone.
- Utile pour : agréger plusieurs liens LAN, ou mettre un port de **bypass/miroir**.

## 27. Sous-interfaces et VLAN : zones par VLAN

```huawei
system-view
[USG] interface GigabitEthernet 1/0/3.10
[USG-GigabitEthernet1/0/3.10] vlan-type dot1q 10
[USG-GigabitEthernet1/0/3.10] ip address 192.168.10.1 24
[USG-GigabitEthernet1/0/3.10] quit
[USG] interface GigabitEthernet 1/0/3.20
[USG-GigabitEthernet1/0/3.20] vlan-type dot1q 20
[USG-GigabitEthernet1/0/3.20] ip address 192.168.20.1 24
[USG-GigabitEthernet1/0/3.20] quit
[USG] firewall zone trust
[USG-zone-trust] add interface GigabitEthernet 1/0/3.10
[USG-zone-trust] quit
[USG] firewall zone name GUEST
[USG-zone-guest] set priority 20
[USG-zone-guest] add interface GigabitEthernet 1/0/3.20
[USG-zone-guest] quit
save
```

⚠️ La sous-interface **hérite** de la zone ? **Non** : chaque sous-interface doit être ajoutée explicitement à une zone. Oubli classique après création d'un VLAN.

## 28. Web : gérer les zones en interface graphique

Chemin (New Web UI) : **Network > Zone**. On y voit les zones, leurs priorités, les interfaces membres. Bouton « Add » pour créer une zone, glisser-déposer les interfaces.
Avantage du web : **vue d'ensemble immédiate** — en CLI, `display zone` reste le plus rapide pour vérifier.

---
---

# BLOC D — POLITIQUES DE SÉCURITÉ (INTERZONE POLICY)

## 29. Principes : comment l'USG évalue une politique

1. Le firewall reçoit un paquet, détermine **zone source** (interface d'entrée) et **zone destination** (routage).
2. Il parcourt les règles **dans l'ordre** (de haut en bas) pour ce couple de zones.
3. **Première règle qui matche → appliquée** (permit/deny). Les suivantes sont ignorées.
4. Si **aucune règle ne matche** → **action par défaut** (configurable : `permit` ou `deny`).

```huawei
system-view
[USG] security-policy
[USG-policy-security] default action deny      # RECOMMANDÉ : tout ce qui n'est pas autorisé est bloqué
[USG-policy-security] quit
save
```

⚠️ **Le choix le plus important de tout le firewall.** `default action permit` = passoire (pratique en migration, mortel en prod). `default action deny` = il faut TOUT autoriser explicitement, mais tu sais exactement ce qui passe. **En prod : deny.**

## 30. Anatomie d'une règle : les 8 critères

```huawei
[USG-policy-security] rule name LAN-vers-Internet
[USG-policy-security-rule-LAN-vers-Internet] source-zone trust          # 1. zone source
[USG-policy-security-rule-LAN-vers-Internet] destination-zone untrust  # 2. zone destination
[USG-policy-security-rule-LAN-vers-Internet] source-address 192.168.10.0 24   # 3. adresses sources
[USG-policy-security-rule-LAN-vers-Internet] destination-address any   # 4. adresses destination
[USG-policy-security-rule-LAN-vers-Internet] service http https dns    # 5. services (ports)
[USG-policy-security-rule-LAN-vers-Internet] application any           # 6. application (UTM)
[USG-policy-security-rule-LAN-vers-Internet] user any                 # 7. utilisateur
[USG-policy-security-rule-LAN-vers-Internet] time-range any            # 8. plage horaire
[USG-policy-security-rule-LAN-vers-Internet] action permit             # action
[USG-policy-security-rule-LAN-vers-Internet] quit
```

Critères non renseignés = `any` (tout). **Plus la règle est précise, plus elle est sûre** — mais plus elle est longue à maintenir. L'équilibre, c'est le métier.

## 31. L'ordre des règles : le piège n°1

```
❌ MAUVAIS ORDRE :
  rule 1 : deny  trust→untrust  service p2p        ← jamais atteinte...
  rule 2 : permit trust→untrust service any        ← ...car celle-ci matche d'abord !

✅ BON ORDRE :
  rule 1 : deny  trust→untrust  service p2p        ← spécifique d'abord
  rule 2 : permit trust→untrust service http https ← général ensuite
```

**Règle d'or : du plus spécifique au plus général.** En CLI, insère une règle à la bonne position :
```huawei
[USG-policy-security] rule name bloquer-p2p
# ... critères ...
[USG-policy-security] move rule bloquer-p2p before LAN-vers-Internet
```

🔧 Voir l'ordre réel : `display security-policy rule all` (l'ordre d'affichage = l'ordre d'évaluation).

## 32. Objets d'adresse : ne jamais mettre d'IP en dur dans les règles

```huawei
system-view
[USG] ip address-set serveurs-dmz type object
[USG-object-group-address-serveurs-dmz] address 0 172.16.1.10 32
[USG-object-group-address-serveurs-dmz] address 1 172.16.1.11 32
[USG-object-group-address-serveurs-dmz] address 2 172.16.1.12 32
[USG-object-group-address-serveurs-dmz] quit
[USG] ip address-set lan-bureautique type range
[USG-object-group-address-lan-bureautique] address 0 range 192.168.10.1 192.168.10.254
[USG-object-group-address-lan-bureautique] quit
# Utilisation dans une règle :
[USG] security-policy
[USG-policy-security] rule name DMZ-vers-LAN-deny
[USG-policy-security-rule-DMZ-vers-LAN-deny] source-zone dmz
[USG-policy-security-rule-DMZ-vers-LAN-deny] destination-zone trust
[USG-policy-security-rule-DMZ-vers-LAN-deny] source-address address-set serveurs-dmz
[USG-policy-security-rule-DMZ-vers-LAN-deny] action deny
[USG-policy-security-rule-DMZ-vers-LAN-deny] quit
[USG-policy-security] quit
save
```

Avantage : quand une IP de serveur change, tu modifies **l'objet**, pas les 12 règles qui l'utilisent.

## 33. Services personnalisés (ports non standards)

```huawei
system-view
[USG] ip service-set appli-metier
[USG-object-group-service-appli-metier] service 0 protocol tcp destination-port 8443
[USG-object-group-service-appli-metier] service 1 protocol tcp destination-port 9443
[USG-object-group-service-appli-metier] service 2 protocol udp destination-port 1194
[USG-object-group-service-appli-metier] quit
save
```

Services prédéfinis utiles : `http`, `https`, `dns`, `ftp`, `smtp`, `pop3`, `ssh`, `telnet`, `ping` (icmp), `snmp`. Pour une appli métier sur port exotique, crée ton service-set — ça rend la règle **lisible** (« appli-metier » parle mieux que « tcp/8443 »).

## 34. Plages horaires : restreindre l'accès dans le temps

```huawei
system-view
[USG] time-range heures-bureau
[USG-time-range-heures-bureau] period-range 08:00:00 to 18:00:00 working-day
[USG-time-range-heures-bureau] quit
[USG] security-policy
[USG-policy-security] rule name guest-horaire
[USG-policy-security-rule-guest-horaire] source-zone GUEST
[USG-policy-security-rule-guest-horaire] destination-zone untrust
[USG-policy-security-rule-guest-horaire] time-range heures-bureau
[USG-policy-security-rule-guest-horaire] action permit
[USG-policy-security-rule-guest-horaire] quit
[USG-policy-security] quit
save
```

Cas d'usage : Wi-Fi invités coupé la nuit, accès VPN nomades limités aux horaires de travail, maintenance le week-end.

## 35. Règle LAN → Internet complète et commentée

