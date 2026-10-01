---
id: collect-261001-rattrapage/rattrapage/opnsense-cli-guide-2
title: "OPNsense — Guide d'administration en CLI (sans l'interface web)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/opnsense_cli_guide.md
source_anchor: ""
source_lines: [232, 551]
sha256: 9f922e8840d77e4d7bbf34f0ef7c52f3da3cf6c99ac51966eae8285855df48c6
---

# Compter les états (signe de saturation / attaque) :
pfctl -ss | wc -l
pfctl -s memory
```

---

## 10. Pare-feu `pf` avec `pfctl` — action

```shell
# Tuer les états d'un hôte (utile après changement de règle/NAT) :
pfctl -k 192.168.1.50
pfctl -k 192.168.1.50 -k 0.0.0.0/0   # forme complète

# Recharger les règles depuis la conf OPNsense :
configctl filter reload
# (équivaut à : pfctl -f /tmp/rules.debug après régénération)

# Désactiver / réactiver pf — DANGER à distance (vous vous coupez !) :
pfctl -d                        # disable (console physique uniquement !)
pfctl -e                        # enable

# Vider les tables d'un alias :
pfctl -t mon_alias -T flush
# Ajouter une IP à une table :
pfctl -t mon_alias -T add 203.0.113.7
```

---

## 11. Ajouter des règles firewall en CLI : `easyrule`

```shell
# Bloquer une IP sur WAN :
easyrule block wan 203.0.113.7

# Autoriser un flux précis :
easyrule pass lan tcp 192.168.1.50 203.0.113.7 443

# Syntaxe générale :
# easyrule {block|pass} <interface> <proto> <source> <dest> <port>
easyrule --help
```

- `easyrule` écrit dans `config.xml` (<filter>) puis recharge pf :
  c'est la méthode CLI **propre** pour ajouter une règle.
- Pour des règles complexes (schedules, groupes), éditez `<filter>`
  dans config.xml ou utilisez l'API REST.

---

## 12. NAT en CLI

```shell
pfctl -sn                       # voir les règles NAT actives
pfctl -s nat                    # idem

# Traductions en cours :
pfctl -s state | grep -i ">" | head

# Le NAT outbound se configure dans <nat><outbound> de config.xml ;
# le port-forward dans <nat><rule> :
grep -A12 "<nat>" /conf/config.xml | head -40

# Après modif :
configctl filter reload
pfctl -k 192.168.1.50           # tue les vieux états pour forcer le nouveau NAT
```

- Mode outbound : `automatic` (tout le LAN en PAT sur l'IP WAN) par
  défaut ; passez en `hybrid`/`manual` pour des exceptions.
- Un port-forward = règle NAT **+** règle firewall WAN→destination
  (n'oubliez jamais la seconde).

---

## 13. DHCP en CLI

```shell
# Conf générée :
cat /var/dhcpd/etc/dhcpd.conf | head -40

# Baux actifs :
cat /var/dhcpd/var/db/dhcpd.leases

# Baux via la conf OPNsense :
grep -A8 "<dhcpd>" /conf/config.xml | head -30

# Redémarrer après modif :
configctl dhcpd restart

# Voir qui a quelle IP en live :
arp -a | grep 192.168.1.
```

---

## 14. DNS Unbound en CLI

```shell
# Tester la résolution via l'Unbound local :
drill www.google.com @127.0.0.1
dig @127.0.0.1 www.google.com

# Stats et cache :
unbound-control stats
unbound-control dump_cache | head
unbound-control flush_zone example.com

# Conf générée :
grep -v "^#" /var/unbound/unbound.conf | head -40

# Redémarrer :
configctl unbound restart

# Blocklists DNS (blocklist plugin) :
grep -A5 "<unboundplus>" /conf/config.xml | head -20
```

---

## 15. VPN : IPsec (strongSwan)

```shell
ipsec status                    # tunnels et SA actifs
ipsec statusall                 # détail complet
swanctl --list-sas              # SAs IKEv2 (moderne)
swanctl --list-conns            # connexions configurées
setkey -D                       # SAD (associations de sécurité, IKEv1)

# Logs IPsec en direct :
clog -f /var/log/ipsec.log
# ou : tail -f /var/log/charon.log  (selon version)

# Redémarrer :
configctl ipsec restart
ipsec restart

# La conf vient de <ipsec> dans config.xml :
grep -A10 "<ipsec>" /conf/config.xml | head -30
```

---

## 16. VPN : OpenVPN et WireGuard

```shell
# OpenVPN :
configctl openvpn status
cat /var/etc/openvpn/*.conf 2>/dev/null | head -30
# Logs :
clog -f /var/log/openvpn.log

# WireGuard :
wg show                         # pairs, handshakes, transfert
wg show wg0 peers
wg show wg0 latest-handshakes
# Un handshake > 3 min = pair injoignable (réseau ou clé fausse)

# Redémarrer :
configctl wireguard restart
configctl openvpn restart
```

- WireGuard : vérifiez `Endpoint`, `AllowedIPs` et l'horloge (dérive >
  quelques minutes = échec d'authentification).
- OpenVPN : `verb 3` dans la conf pour des logs exploitables.

---

## 17. IDS/IPS Suricata en CLI

```shell
configctl suricata status
suricata --version
# Règles chargées :
ls /usr/local/etc/suricata/rules/
# Logs d'alertes :
clog -f /var/log/suricata/eve.json | grep alert
# ou : tail -f /var/log/suricata/fast.log

# Recharger après changement de règles :
configctl suricata reload
```

---

## 18. Logs : tout voir en CLI

```shell
# Logs circulaires (clog) — les principaux :
clog /var/log/system.log | tail -50       # système général
clog -f /var/log/filter.log               # firewall en direct (live view)
clog /var/log/filter.log | grep 192.168.1.50
clog /var/log/dhcpd.log | tail -30
clog /var/log/unbound.log | tail -30
clog /var/log/ipsec.log | tail -50
clog /var/log/openvpn.log | tail -30

# Boot et matériel :
dmesg | tail -40
dmesg | grep -i "error\|fail"

# Tous les logs disponibles :
ls -lh /var/log/
```

- `clog -f` = `tail -f` pour les logs circulaires (ne jamais `cat`
  bêtement un fichier clog : utilisez `clog` pour le lire).
- Filtrez par IP/port pour isoler un flux bloqué : c'est l'équivalent
  du « live view » de la web UI.

---

## 19. Monitoring et performance

```shell
top -b | head -20               # CPU / mémoire (mode batch)
systat -vmstat 1                # vue système live (Ctrl+C pour quitter)
vmstat 1 5
netstat -m                      # mémoire mbuf (réseau)
pfctl -s memory                 # mémoire utilisée par pf
pfctl -ss | wc -l               # nombre d'états (table pleine ?)
df -h                           # espace disque
dmesg | grep -i temperature     # températures (selon matériel)
```

- États pf par défaut : limite haute (plusieurs centaines de milliers)
  ; si `pfctl -s info` montre des `state-table` insertions failures =
  table pleine → augmentez dans config.xml ou traquez l'attaque.

---

## 20. Mises à jour firmware en CLI

```shell
# Vérifier les mises à jour :
configctl firmware status
opnsense-update -c              # check only

# Mise à jour (équivalent option 12 du menu) :
opnsense-update -t opnsense

# Voir l'historique :
opnsense-update -l | head

# Kernel / base séparément si besoin :
# opnsense-update -t kernel
# opnsense-update -t base
```

- Toujours un **backup de config** avant (`cp /conf/config.xml …`).
- Après une mise à jour majeure, `reboot` et vérifiez
  `cat /usr/local/opnsense/version/opnsense.version`.

---

## 21. Sauvegarde / restauration en CLI

```shell
# Backup manuel :
cp /conf/config.xml /root/opnsense-backup-$(date +%Y%m%d).xml

# Backups automatiques rotatifs :
ls -lt /conf/backup/ | head

# Exporter hors de la machine :
scp /conf/config.xml admin@192.168.1.100:/backups/

# Restaurer :
cp /root/opnsense-backup-20260101.xml /conf/config.xml
configctl configd reload
# ou option 13 du menu console, ou reboot.

# Chiffrement du backup : la web UI propose un mot de passe ;
# en CLI, chiffrez vous-même :
# openssl enc -aes-256-cbc -in /conf/config.xml -out backup.enc
```

---

## 22. Utilisateurs et mots de passe en CLI

```shell
# Changer le mot de passe root :
passwd

# L'admin web = utilisateur « root » par défaut (section <system>
# de config.xml, <user>). Ajouter un utilisateur en CLI = éditer
# <system><user> dans config.xml (mot de passe hashé bcrypt).

# Mot de passe perdu : option 3 du menu console
# (réinitialise le mot de passe root).

# Voir les utilisateurs configurés :
grep -A6 "<user>" /conf/config.xml | head -30

# Clés SSH autorisées :
cat /root/.ssh/authorized_keys
```

---

## 23. HA / CARP en CLI

```shell
# Interfaces CARP :
ifconfig | grep -A3 carp

# État CARP (MASTER / BACKUP) :
ifconfig carp0 | grep carp:

# Paramètres sysctl :
sysctl net.inet.carp.preempt
sysctl net.inet.carp.allow

# Forcer le basculement (maintenance) :
ifconfig carp0 advskew 254    # se rend moins prioritaire
# puis remettre : ifconfig carp0 advskew 0

