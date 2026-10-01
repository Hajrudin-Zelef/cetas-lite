---
id: collect-261001-rattrapage/rattrapage/opnsense-cli-guide-1
title: "OPNsense — Guide d'administration en CLI (sans l'interface web)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-rattrapage/opnsense_cli_guide.md
source_anchor: ""
source_lines: [1, 231]
sha256: 4c49526275caf8058bc75591d35334cefc56da5e54213668232eaa85e05b0039
---

# OPNsense — Guide d'administration en CLI (sans l'interface web)

> OPNsense est un firewall/routeur basé sur **FreeBSD** (+ `pf` pour le
> filtrage). Contrairement à Cisco ou Huawei, il n'y a pas de « CLI
> propriétaire » : l'administration en ligne de commande passe par
> **le shell FreeBSD**, l'utilitaire **`configctl`** et **`pfctl`**.
> Ce guide montre comment tout faire au clavier : config, firewall,
> NAT, VPN, diagnostic, mises à jour.
>
> **À jour 2025-2026** : vise OPNsense **24.x / 25.x** (base FreeBSD 14,
> pf, WireGuard/OpenVPN/IPsec via strongSwan, Unbound, Suricata).
> Certaines options du menu console peuvent varier légèrement selon
> la version — la logique reste identique.

---

## 1. Philosophie : comment OPNsense est construit

- **Un seul fichier de vérité** : `/conf/config.xml` contient TOUTE la
  configuration (interfaces, règles, NAT, VPN, utilisateurs…).
- L'interface web n'est qu'un éditeur de ce fichier : chaque « Apply »
  régénère les fichiers de conf des services via **configd**
  (templates) puis recharge les services.
- En CLI, vous avez deux approches :
  1. **Utilitaires dédiés** (`configctl`, `easyrule`, `pfctl`…) — propre ;
  2. **Éditer `/conf/config.xml`** à la main puis recharger — puissant,
     à manier avec précaution (toujours sauvegarder avant).
- Le filtrage est fait par **`pf`** (packet filter de FreeBSD) :
  `pfctl` l'interroge et le pilote en direct.

---

## 2. Accéder au shell

```shell
# En console physique / VGA / série : le menu s'affiche après login root.
# En SSH (une fois activé) : choisissez l'option 8 (Shell).

# Activer SSH (à faire une fois, via web UI ou en éditant config.xml) :
# System > Settings > Administration > Enable Secure Shell.
# En CLI pur, éditez /conf/config.xml, section <system><ssh>... (voir §4).

# Une fois dans le shell :
whoami          # root
uname -a        # FreeBSD 14.x
cat /usr/local/opnsense/version/opnsense.version
```

- Le prompt du menu console propose le choix **8) Shell** pour tomber
  sur un vrai shell `sh`/`csh` en root.
- `exit` depuis le shell **revient au menu console**, pas de déconnexion.

---

## 3. Le menu console (référence)

```
0) Logout                       7) Ping host
1) Assign Interfaces            8) Shell
2) Set interface(s) IP address  9) pfTop (ou pfctl)
3) Reset the root password      10) Firewall Logs (vue live)
4) Reset to factory defaults    11) Reload all services
5) Power off system             12) Firmware update
6) Reboot system                13) Restore a backup
```

- **1)** : réassigner les interfaces physiques (em0, igb0…) aux rôles
  LAN/WAN/OPT. Indispensable après changement de carte réseau.
- **2)** : configurer IP/masque/passerelle/DHCP d'une interface, activer
  le serveur DHCP — le « quick setup » d'une fresh install.
- **3)** : réinitialiser le mot de passe root / admin web.
- **4)** : retour usine (efface `/conf/config.xml` !).
- **12)** : mise à jour du firmware sans web UI (voir §18).
- **13)** : restaurer un backup de config depuis un fichier.

---

## 4. `/conf/config.xml` — le cœur du système

```shell
# Toujours sauvegarder AVANT de toucher :
cp /conf/config.xml /conf/config.xml.bak-$(date +%Y%m%d-%H%M)
cp /conf/config.xml /root/config-backup.xml

# Voir la structure (extraits) :
grep -A5 "<interfaces>" /conf/config.xml | head -30
grep -A3 "<filter>" /conf/config.xml | head -20

# Valider que le XML est bien formé après édition :
xmllint --noout /conf/config.xml && echo "XML OK"
# (xmllint peut ne pas être installé : pkg install -y libxml2)
```

**Workflow d'édition manuelle** :
1. `cp /conf/config.xml /conf/config.xml.bak`
2. Éditez avec `vi` / `ee` (sections `<interfaces>`, `<filter>`,
   `<nat>`, `<ipsec>`, `<openvpn>`, `<unbound>`…).
3. Validez le XML (`xmllint`).
4. Appliquez : rechargez le service concerné via `configctl`
   (voir §5) ou `reboot` si vous avez touché aux interfaces.
5. En cas de problème : restaurez le `.bak` et rechargez.

> Les backups automatiques sont dans `/conf/backup/` (rotatifs).
> L'option 13 du menu console restaure l'un d'eux.

---

## 5. `configctl` — le couteau suisse

`configctl` pilote configd : il liste et exécute les actions des services.

```shell
configctl                       # liste TOUTES les actions disponibles
configctl interface list        # interfaces connues du système
configctl interface reconfigure # réapplique la conf réseau
configctl filter reload         # recharge les règles pf
configctl webgui restart        # redémarre l'interface web
configctl unbound restart       # redémarre le DNS Unbound
configctl dhcpd restart         # redémarre le serveur DHCP
configctl ipsec status          # état des tunnels IPsec
configctl ipsec start           # démarre strongSwan
configctl openvpn status
configctl wireguard status
configctl suricata status
configctl firmware status       # état des mises à jour
```

- Astuce : `configctl <service>` sans argument liste souvent les
  sous-actions du service.
- Après édition manuelle de `config.xml`, c'est `configctl` qui
  « applique » sans reboot.

---

## 6. Interfaces réseau en CLI

```shell
ifconfig                        # toutes les interfaces (état, IP, MAC)
ifconfig em0                    # détail d'une interface
ifconfig -a | grep -A1 "status: active"

# Monter / descendre une interface :
ifconfig em1 up
ifconfig em1 down

# IP temporaire (perdue au reboot si pas dans config.xml !) :
ifconfig em1 inet 192.168.2.1 netmask 255.255.255.0

# Statistiques et erreurs :
netstat -i
netstat -I em0 -w 1             # débit live d'une interface

# Renommer / réassigner durablement : option 1 du menu console,
# ou éditer <interfaces> dans config.xml puis :
configctl interface reconfigure
```

- Noms FreeBSD : `em`, `igb`, `ix` (Intel), `re` (Realtek), `vtnet`
  (virtuel/KVM), `vmx` (VMware).
- Le rôle logique (lan/wan/opt1) est mappé dans `<interfaces>` ;
  `ifconfig` montre les noms physiques.

---

## 7. VLAN, LAGG (agrégat), bridges

```shell
# VLAN 10 sur em1 (temporaire ; durable = config.xml) :
ifconfig em1.10 create vlan 10 vlandev em1
ifconfig em1.10 inet 192.168.10.1 netmask 255.255.255.0 up

# LAGG (équivalent EtherChannel) :
ifconfig lagg0 create
ifconfig lagg0 laggproto lacp laggport em2 laggport em3
ifconfig lagg0 up

# Bridge :
ifconfig bridge0 create
ifconfig bridge0 addm em1 addm em2 up

# Voir :
ifconfig lagg0
ifconfig bridge0
```

- En production, créez-les via `<vlans>`, `<laggs>`, `<bridges>` dans
  `config.xml` (ou la web UI une fois), puis `configctl interface
  reconfigure`.

---

## 8. Routage

```shell
netstat -rn                     # table de routage complète
netstat -rn | grep default      # passerelle par défaut
route -n get 8.8.8.8            # quelle route pour cette destination ?

# Route statique temporaire :
route add -net 192.168.5.0/24 192.168.1.254
route delete -net 192.168.5.0/24 192.168.1.254

# Passerelles configurées (multi-WAN) :
grep -A10 "<gateways>" /conf/config.xml

# Groupes de passerelles (failover / load-balancing) :
grep -A15 "<gateway_groups>" /conf/config.xml
```

- Les routes durables vont dans `<staticroutes>` de config.xml.
- Le routage dynamique (OSPF/BGP) passe par le plugin **FRR**
  (`vtysh` pour la CLI Quagga/FRR une fois installé).

---

## 9. Pare-feu `pf` avec `pfctl` — lecture

```shell
pfctl -sr                       # règles de filtrage actives
pfctl -sn                       # règles NAT actives
pfctl -s state                  # alias : -ss, états/connexions actives
pfctl -ss | head -20
pfctl -s info                   # compteurs globaux pf
pfctl -s Tables                 # tables (alias) chargées
pfctl -t bogons -T show | head  # contenu d'une table
pfctl -s queue                  # files de traffic shaping
pfctl -s labels                 # règles avec labels (utile au debug)

# États pour une IP précise :
pfctl -ss | grep 192.168.1.50

