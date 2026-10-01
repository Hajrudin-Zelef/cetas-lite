---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-17
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2424, 2548]
sha256: d2406eea3f57d0c39e6f519d3cd56ccaf29f82bb1addc0406c387516633657c9
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

- **Symptômes** : CPU à 100 %, réseau saturé, tout rame — juste après un brassage.
- **Diagnostic** : `display cpu`, `display interface` (compteurs d'erreurs qui explosent), vérifier physiquement : **deux câbles** entre les mêmes switchs ? Un câble en **boucle** sur le même équipement ?
- **Solution** : débrancher le câble fautif, vérifier le spanning-tree des switchs. 📋 **Étiqueter les câbles** (section 11) — ça évite 80 % de ces incidents.

## 159. Cas 16 : authentification AD qui échoue en masse

- **Symptômes** : plus personne ne s'authentifie (VPN, portail).
- **Diagnostic** : `ping` vers l'AD, port **389/636** ouvert ? **Compte de bind** : mot de passe expiré (le classique !) ? **Heure** : désynchro > 5 min = échec Kerberos. `display aaa online-fail-record`.
- **Solution** : renouveler le mot de passe du compte de service (avec **rappel calendaire** avant la prochaine expiration), resynchroniser NTP, basculer temporairement sur la **base locale de secours** si besoin.

## 160. Cas 17 : logs qui n'arrivent plus au syslog

- **Symptômes** : le serveur syslog ne reçoit plus rien (découvert pendant une enquête, évidemment).
- **Diagnostic** : `ping` vers le syslog depuis l'USG ; politique **local→untrust/trust** : UDP 514 autorisé ? Le **disque** du syslog est-il plein ? La config `info-center loghost` est-elle toujours là (`display current-configuration | include loghost`) ?
- **Solution** : corriger réseau/politique/espace disque. 📋 **Superviser la réception** des logs (alerte si 0 log reçu en 15 min) — sinon tu ne le sauras jamais.

## 161. Cas 18 : lenteurs intermittentes (le pire)

- **Symptômes** : « parfois c'est lent », non reproductible à la demande.
- **Méthode** : ne pas bricoler — **mesurer d'abord** :
  1. Historique supervision : CPU/mémoire/sessions aux heures des plaintes.
  2. `display firewall session table` en pointe : **saturation** (table pleine ?) ?
  3. Corrélations : backup à 14h ? MAJ signatures ? Scan antivirus des postes ?
  4. Tester **en bypass** (si possible) pour isoler le firewall du reste.
- **Causes typiques** : saturation de sessions (un poste vérolé qui spamme), lien FAI saturé (pas le firewall !), UTM en pointe, **boucle partielle**.
- **Solution** : traiter la cause mesurée, pas le symptôme. Documenter.

## 162. Kit de survie : les 10 commandes à connaître par cœur

```huawei
display version                          # version logicielle
display current-configuration            # config complète
display ip routing-table                 # routage
display zone                             # zones et interfaces
display security-policy rule all         # politiques dans l'ordre
display nat-policy rule all              # NAT dans l'ordre
display firewall session table           # sessions actives
display cpu / display memory             # ressources
display ike sa / display ipsec sa        # VPN
display hrp state verbose                # HA
display logbuffer                        # logs récents
display alarm                            # alarmes matérielles
debugging packet-filter ... / undo debugging all   # trace ciblée (à couper !)
```

---
---

# BLOC O — CAS PRATIQUES COMPLETS COMMENTÉS

## 163. Cas pratique 1 : PME 80 postes, 1 lien fibre — configuration complète

**Contexte** : siège PME, 80 postes (192.168.10.0/24), 1 serveur web vitrine + 1 serveur mail en DMZ (172.16.1.0/24), lien fibre 1 Gbit/s (IP publique fixe 203.0.113.1), 10 nomades en SSL VPN. Modèle : **USG6380** (6 Gbit/s firewall — dimensionné pour ~1 Gbit/s UTM, voir section 73).

```huawei
system-view
# ===== Interfaces et zones =====
[USG] interface GigabitEthernet 1/0/1
[USG-GigabitEthernet1/0/1] description LIEN-FIBRE-FAI
[USG-GigabitEthernet1/0/1] ip address 203.0.113.1 30
[USG-GigabitEthernet1/0/1] quit
[USG] interface GigabitEthernet 1/0/2
[USG-GigabitEthernet1/0/2] description DMZ-SERVERS
[USG-GigabitEthernet1/0/2] ip address 172.16.1.1 24
[USG-GigabitEthernet1/0/2] quit
[USG] interface GigabitEthernet 1/0/3
[USG-GigabitEthernet1/0/3] description LAN-BUREAUTIQUE
[USG-GigabitEthernet1/0/3] ip address 192.168.10.1 24
[USG-GigabitEthernet1/0/3] quit
[USG] firewall zone untrust
[USG-zone-untrust] set priority 5
[USG-zone-untrust] add interface GigabitEthernet 1/0/1
[USG-zone-untrust] quit
[USG] firewall zone dmz
[USG-zone-dmz] set priority 50
[USG-zone-dmz] add interface GigabitEthernet 1/0/2
[USG-zone-dmz] quit
[USG] firewall zone trust
[USG-zone-trust] set priority 85
[USG-zone-trust] add interface GigabitEthernet 1/0/3
[USG-zone-trust] quit
# ===== Routage =====
[USG] ip route-static 0.0.0.0 0.0.0.0 203.0.113.2    # passerelle FAI (fictive)
# ===== NAT source (Easy IP) =====
[USG] nat-policy
[USG-policy-nat] rule name nat-lan
[USG-policy-nat-rule-nat-lan] source-zone trust
[USG-policy-nat-rule-nat-lan] destination-zone untrust
[USG-policy-nat-rule-nat-lan] source-address 192.168.10.0 24
[USG-policy-nat-rule-nat-lan] action nat easy-ip
[USG-policy-nat-rule-nat-lan] quit
[USG-policy-nat] quit
# ===== NAT destination (serveurs publiés) =====
[USG] nat-policy
[USG-policy-nat] rule name dnat-web
[USG-policy-nat-rule-dnat-web] source-zone untrust
[USG-policy-nat-rule-dnat-web] destination-zone dmz
[USG-policy-nat-rule-dnat-web] destination-address 203.0.113.10 32
[USG-policy-nat-rule-dnat-web] service http
[USG-policy-nat-rule-dnat-web] action nat static
[USG-policy-nat-rule-dnat-web] static-to 172.16.1.10
[USG-policy-nat-rule-dnat-web] quit
[USG-policy-nat] rule name dnat-mail
[USG-policy-nat-rule-dnat-mail] source-zone untrust
[USG-policy-nat-rule-dnat-mail] destination-zone dmz
[USG-policy-nat-rule-dnat-mail] destination-address 203.0.113.11 32
[USG-policy-nat-rule-dnat-mail] service smtp
[USG-policy-nat-rule-dnat-mail] action nat static
[USG-policy-nat-rule-dnat-mail] static-to 172.16.1.25
[USG-policy-nat-rule-dnat-mail] quit
[USG-policy-nat] quit
# ===== Politiques (extraits, voir sections 35-40 pour le détail) =====
[USG] security-policy
[USG-policy-security] default action deny
[USG-policy-security] quit
# ... ajouter les règles LAN→Internet, untrust→DMZ, protections local (sections 35-40) ...
save
```

📋 **Recette de mise en service** : ping LAN→firewall, DNS, web, mail entrant/sortant, EICAR (section 85), test depuis un mobile en 4G (vue « extérieure »), backup config.

## 164. Cas pratique 2 : site industriel avec 2 liens WAN (principal + secours 4G)

**Contexte** : site de production, coupure Internet = arrêt de la supervision. Lien principal fibre + **secours 4G** (clé USB/routeur 4G).

