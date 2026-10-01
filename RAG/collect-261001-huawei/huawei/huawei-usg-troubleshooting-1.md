---
id: collect-261001-huawei/huawei/huawei-usg-troubleshooting-1
title: "Huawei USG — Guide de dépannage (troubleshooting récurrent)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-huawei/huawei_usg_troubleshooting.md
source_anchor: ""
source_lines: [1, 206]
sha256: 0a0470b25b064100c6d2ccb934f3e1b136f339abec76a460d199b4a190c04f51
---

# Huawei USG — Guide de dépannage (troubleshooting récurrent)

> Les pannes les plus fréquentes sur USG6000 (V500/V600), avec pour chacune :
> symptômes, causes probables, commandes de diagnostic et corrections.
> Méthode générale : toujours dépanner **de bas en haut** —
> 1. Physique / interfaces, 2. Zones, 3. Security-policy,
> 4. NAT, 5. Routage, 6. UTM / applicatif.
> La majorité des « pannes » USG sont des oublis de **security-policy**
> ou de **NAT** : le pare-feu refuse par défaut.

---

## 1. Réflexes de base avant tout diagnostic

```shell
display version                    # version logicielle (noter le VxxxRxxx)
display device                     # état matériel, cartes
display cpu-usage                  # CPU > 80 % = suspect
display memory-usage               # mémoire saturée = sessions/fuites
display interface brief            # UP/DOWN de toutes les interfaces
display ip interface brief
display logbuffer                  # erreurs récentes
display trapbuffer                 # alarmes
display clock                      # une horloge fausse casse les logs/VPN !
```

- Vérifiez **l'heure** en premier sur les problèmes VPN/certificats/licence.
- `save` **après** chaque correction validée, jamais avant.
- En cas de doute sur une règle : `display security-policy rule all`
  puis regardez l'ordre des règles (la première qui matche gagne).

---

## 2. Pas d'accès Internet depuis le LAN

**Symptômes** : les postes n'ouvrent aucune page, `ping 8.8.8.8` échoue
depuis un PC du LAN.

**Checklist ordonnée** :

```shell
# a) L'interface LAN est-elle UP avec la bonne IP ?
display interface GigabitEthernet 0/0/1
display ip interface brief

# b) Le PC a-t-il une passerelle ? (depuis le USG)
ping -a 192.168.1.1 192.168.1.50
display arp all | include 192.168.1.50

# c) La route par défaut existe-t-elle ?
display ip routing-table 0.0.0.0

# d) Le NAT source est-il configuré ET matche-t-il ?
display nat-policy rule all
display nat session all | include 192.168.1.50

# e) La security-policy autorise-t-elle trust -> untrust ?
display security-policy rule all
display firewall session table source-ip 192.168.1.50

# f) Le DNS fonctionne-t-il ?
nslookup www.google.com
# (depuis le PC : ping 8.8.8.8 OK mais pas de navigation = DNS)
```

**Corrections typiques** :
- Interface DOWN : câble, `undo shutdown` sur l'interface.
- Pas de route par défaut : `ip route-static 0.0.0.0 0.0.0.0 <passerelle>`.
- NAT oublié : ajouter la règle `nat-policy` trust→untrust
  (`action source-nat easy-ip`, voir guide CLI §8).
- Règle security-policy manquante ou **mal ordonnée** (une règle `deny`
  placée avant la règle `permit`).
- Interface WAN pas dans la zone `untrust` (`firewall zone untrust` →
  `add interface ...`).
- PPPoE : `display pppoe-client session summary` — session absente =
  identifiants faux ou ligne KO.

---

## 3. Le NAT ne fonctionne pas

**Symptômes** : `display nat session all` vide alors que le trafic part ;
ou certains protocoles (FTP, SIP, jeux) cassés.

```shell
display nat-policy rule all          # la règle existe ?
display nat address-group            # le pool a des IP libres ?
display nat session all              # sessions traduites visibles ?
display nat session source-ip 192.168.1.50
```

**Causes fréquentes** :
- La règle NAT ne matche pas : `source-zone` / `destination-zone`
  inversées, ou `source-address` trop restrictif.
- Pool épuisé : trop de sessions pour trop peu d'IP publiques
  (élargir le pool ou passer en `easy-ip`).
- Protocoles à ports dynamiques (FTP actif, SIP) : activer l'**ALG**
  correspondant (web UI : Network > ALG, ou vérifier qu'il n'est pas
  désactivé).
- Double NAT en amont (box opérateur) : le USG reçoit déjà une IP
  privée → préférer le mode bridge/DMZ de la box, ou accepter le
  double NAT (casse IPsec sans NAT-T).
- Pour tester sans la security-policy : créez une règle `permit` large
  temporaire, validez le NAT, puis resserrez.

---

## 4. Une règle security-policy bloque (ou laisse passer) à tort

```shell
# Voir les sessions réellement établies pour une IP
display firewall session table source-ip 192.168.1.50
display firewall session table destination-port 443

# Voir l'ordre et le contenu des règles
display security-policy rule all

# Statistiques globales
display firewall statistics
```

**Méthode** :
1. Reproduisez le flux, puis regardez `display firewall session table` :
   session absente = bloqué par une politique (ou pas de route/NAT).
2. Remontez les règles **dans l'ordre** : la première règle dont
   (source-zone, destination-zone, adresses, service, schedule)
   correspondent décide.
3. Pièges classiques :
   - règle `deny any any` placée **avant** les permits ;
   - `destination-zone local` oubliée pour administrer le USG ;
   - `schedule time-range` expiré (la règle ne s'applique plus !) ;
   - objet `address-set` vide ou mal renseigné ;
   - trafic retour bloqué : normalement inutile (stateful), mais si la
     session a expiré (`display firewall session table verbose`),
     augmentez le timeout du service.

**Fix** : réordonner (`move rule`), corriger les objets, ou ajouter une
règle de log temporaire (`action permit` + session-logging) pour voir.

---

## 5. Serveur publié inaccessible depuis Internet (NAT server)

**Symptômes** : `nat server` configuré, mais pas de connexion entrante.

```shell
display nat server
display security-policy rule all     # règle untrust -> dmz/trust ?
display firewall session table destination-ip 192.168.50.10
```

**Checklist** :
1. La règle security-policy **untrust → dmz** (ou trust) existe avec le
   bon `service` (port) et matche l'IP **privée** destination
   (le NAT server s'applique avant la politique : la règle voit l'IP
   interne !).
2. Le serveur a bien le USG comme **passerelle par défaut** (sinon le
   retour part ailleurs → session asymétrique, paquets jetés).
3. Pas de conflit : le port public n'est pas déjà utilisé par un autre
   `nat server` ou par un service du USG lui-même.
4. FAI qui bloque le port (25, 80, 443 courants) : testez depuis
   l'extérieur avec un autre port.
5. ARP : si l'IP publique n'est pas celle de l'interface, le USG répond
   en proxy-ARP par défaut ; sinon déclarez-la.
6. Depuis le LAN, l'accès via l'IP **publique** ne marche pas sans
   hairpin : utilisez le DNS interne ou l'IP privée en interne.

---

## 6. VPN IPsec : le tunnel ne monte pas

**Symptômes** : `display ike sa` vide.

```shell
display ike sa
display ike proposal
display ipsec proposal
# Côté réseau :
ping <ip-publique-distante>
display ip routing-table <ip-publique-distante>
```

**Phase 1 (IKE) — checklist** :
1. Connectivité : les deux pairs se pinguent en public ?
2. `remote-address` correct des deux côtés (attention aux IP
   dynamiques : utilisez alors l'identification par FQDN ou aggressive
   mode côté dynamique).
3. **Pre-shared-key identique** (erreur n°1, invisible dans les logs).
4. Propositions IKE compatibles : encryption / DH group / auth-algo
   doivent avoir une intersection (même `ike proposal` des 2 côtés =
   le plus simple).
5. UDP 500/4500 et protocole ESP (50) **autorisés** dans la
   security-policy entre `local` et `untrust` dans les deux sens.
6. Horloge : un décalage > quelques minutes peut faire échouer
   l'authentification (surtout avec certificats).
7. NAT-T : si un NAT est sur le chemin, les deux côtés doivent
   l'accepter (automatique en général, UDP 4500).

**Phase 2 (IPsec) — `ike sa` OK mais `ipsec sa` vide** :
1. Propositions IPsec compatibles (ESP auth/encryption).
2. **ACL miroir** : `permit ip source LAN_A destination LAN_B` d'un
   côté, `permit ip source LAN_B destination LAN_A` de l'autre.
   C'est l'erreur n°1 de la phase 2.
3. PFS (Perfect Forward Secrecy) : même groupe DH des deux côtés
   (ou désactivé des deux côtés).

