---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-4
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-01-15", "2026-09-27"]
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [663, 869]
sha256: 5b10e0eca8dd23f9629a166cc05cb2d57daac0b9b7f27ff3da814c6ebdc71ae5
---

# VRP — Le système d'exploitation transversal Huawei

Types : `D` = dynamique, `S` = statique, `I` = interface. Variantes :

```vrp
<AR720>display arp interface GigabitEthernet 0/0/0
<AR720>display arp 192.168.1.2
<AR720>display arp statistics
<AR720>display arp all                      # sur certains modèles : table complète
```

Une entrée ARP manquante pour la passerelle = vérifier le lien L2 et le
VLAN avant tout.

## 39. MAC : `display mac-address`

```vrp
<S310>display mac-address
MAC Address    VLAN/VSI/BD   Learned-From        Type
------------------------------------------------------------------------------
00e0-4c12-3456 10            GE0/0/1             dynamic
00e0-4c78-9abc 10            GE0/0/2             dynamic
------------------------------------------------------------------------------
Total items displayed = 2
<S310>display mac-address interface GigabitEthernet 0/0/1
<S310>display mac-address vlan 10
<S310>display mac-address 00e0-4c12-3456
<S310>display mac-address statistics
```

Sur un switch, c'est LA commande pour savoir **où est branché un
équipement** : on cherche sa MAC, on trouve le port. (Sur l'USG6000, voir
aussi `display firewall session table` pour les sessions.)

## 40. Logs : `display logbuffer` et `display trapbuffer`

```vrp
<AR720>display logbuffer
Logging buffer configuration and contents:enabled
Allowed max buffer size : 1024
Actual buffer size : 256
Channel number : 4 , Channel name : logbuffer
Dropped messages : 0
Overwritten messages : 0
Current messages : 87

Jan 15 2026 14:32:10 AR720 %%01IFNET/4/LINK_STATE(l)[3]:The line protocol IP
on the interface GigabitEthernet0/0/1 has entered the DOWN state.
```

```vrp
<AR720>display logbuffer | include DOWN
<AR720>display logbuffer | begin Jan 15
<AR720>display trapbuffer                    # traps SNMP enregistrés
<AR720>display logbuffer size 512            # (syntaxe à vérifier selon version)
```

Chaque message porte un code `%%01IFNET/4/LINK_STATE` : module / sévérité /
sujet. Sévérités : 0 Emergency → 7 Debug. En dépannage, commencez toujours
par lire les 50 dernières lignes du logbuffer.

## 41. CPU et mémoire : `display cpu-usage`, `display memory`

```vrp
<AR720>display cpu-usage
CPU Usage Stat. Cycle: 60 (Second)
CPU Usage            : 12% Max: 45%@Jan 15 2026 14:00:00
CPU Usage Stat. Time : 2026-01-15 14:35:00
<AR720>display cpu-usage history 60m          # historique (selon modèle)
<AR720>display memory
System Total Memory: 2048 M bytes
Total Memory Used:  612 M bytes
Memory Using Percentage: 29%
<AR720>display memory-usage                   # variante selon version
```

Seuils d'alerte terrain : CPU > 80 % durable = investiguer (`display
cpu-defend`, processus gourmand) ; mémoire > 85 % = risque d'instabilité.

## 42. Informations système : `display clock`, `display users`

```vrp
<AR720>display clock
2026-09-27 01:30:00+02:00 Sunday
Time Zone(DefaultZoneName) : CET
<AR720>display users
  User-Intf    Delay    Type   Network Address     AuthenStatus AuthorcmdStatus
  0   CON 0    00:00:00                                                  pass
  1   VTY 0    00:01:12   SSH   192.168.10.50       pass           no
<AR720>display current-configuration | include sysname
```

`display users` montre qui est connecté (CON = console, VTY = distant) —
utile avant un `reboot` ou un upgrade.

## 43. Diagnostic réseau : `ping` et `tracert`

```vrp
<AR720>ping 192.168.1.254
  PING 192.168.1.254: 56  data bytes, press CTRL_C to break
    Reply from 192.168.1.254: bytes=56 Sequence=1 ttl=64 time=1 ms
    ...
  --- 192.168.1.254 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 1/1/2 ms
<AR720>ping -a 192.168.1.1 192.168.1.254    # source forcée (-a = source IP)
<AR720>ping -c 100 -s 1400 192.168.1.254    # 100 paquets de 1400 octets
<AR720>tracert 8.8.8.8
<AR720>tracert -a 192.168.1.1 8.8.8.8
```

Options clés : `-a` (adresse source), `-c` (nombre de paquets), `-s`
(taille), `-t` (timeout). Pour tester un MTU / une fragmentation : gros
paquets avec `-s 1500`.

## 44. `debugging` — le debug à la Huawei

```vrp
<AR720>terminal debugging                   # autorise l'affichage des debugs
<AR720>debugging ip icmp
<AR720>debugging ospf packet
PING 192.168.1.2 ...
<AR720>undo debugging all                   # TOUJOURS couper après usage
<AR720>undo terminal debugging
<AR720>display debugging                    # vérifie ce qui est actif
```

Règles d'or :
1. **Fenêtre de maintenance** : un debug peut saturer le CPU.
2. **Cibler** : `debugging ospf packet interface GE0/0/0` plutôt que tout OSPF.
3. **Couper** : `undo debugging all` en fin de session, puis `save` si la
   config `terminal debugging` a été persistée par erreur.
4. Sur console, les debugs s'affichent directement ; en SSH/Telnet il faut
   `terminal monitor` + `terminal debugging`.

## 45. `display diagnostic-information` — le tout-en-un

```vrp
<AR720>display diagnostic-information
```

Génère un **fichier** (ou un long affichage) regroupant version, config,
interfaces, routage, logs, CPU, mémoire... C'est ce que le support Huawei
demande en premier lors d'un ticket. Redirigez vers un fichier pour
l'archiver :

```vrp
<AR720>display diagnostic-information > flash:/diag_2026-09-27.txt
```

(La redirection `>` vers fichier existe sur la plupart des versions VRP —
à vérifier sur votre version exacte.)

## 46. Santé matérielle : alarmes et environnement

```vrp
<AR720>display alarm active                 # alarmes en cours
<AR720>display device alarm hardware        # alarmes matérielles
<AR720>display health                       # état de santé global (selon modèle)
<AR720>display temperature all              # variante : display device temperature all
<AR720>display fan                          # variante : display device fan
<AR720>display power                        # variante : display device power
```

En exploitation, un check hebdomadaire `display alarm active` + température
+ alimentations évite 80 % des pannes « surprises ».

## 47. NTP et horloge

```vrp
<AR720>display ntp-service status
 clock status: synchronized
 clock stratum: 3
 reference clock ID: 192.168.100.1
 nominal frequency: 64.0002 Hz
 ...
<AR720>display ntp-service sessions
```

Sans horloge synchronisée, les logs sont inexploitables en corrélation
d'incidents. Configurez toujours au moins un serveur NTP (voir section 58).

## 48. Sessions et sécurité : qui fait quoi

```vrp
<AR720>display ssh server status            # état du serveur SSH
<AR720>display ssh user-information         # utilisateurs SSH connectés
<AR720>display telnet server status
<AR720>display snmp-agent statistics
<USG6000>display firewall session table     # sessions traversant le pare-feu
<USG6000>display security-policy rule all   # règles effectives
```

## 49. Vérifier la redondance : stacking / cluster

```vrp
<S310>display stack                          # état iStack (S310)
<S310>display stack topology
<S310>display css status                     # CSS sur gros switchs (à vérifier)
<AR720>display vrrp                          # VRRP
<AR720>display vrrp brief
<USG6000>display hrp state                   # HRP (haute dispo USG)
<USG6000>display hrp interface
```

## 50. Tableau récapitulatif des `display` par thème

