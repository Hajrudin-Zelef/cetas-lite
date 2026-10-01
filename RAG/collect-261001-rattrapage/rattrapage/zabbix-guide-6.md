---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-6
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache", "memory"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [844, 1013]
sha256: 24407c29fdce54d0a2c9e3655e7610694ee4e9a4d24a543c25dacd14df4f6ecb
---

# Guide Zabbix complet — Supervision d'infrastructure en production

Pour 50 serveurs, oubliez le clic : utilisez l'**API** (exemple minimal, à adapter) :

```bash
# Obtenir un token : Administration → General → API tokens (Zabbix 6.4+/7.x : Bearer token)
TOKEN="votre_token_api_ici"

curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" \
  -d '{
    "jsonrpc": "2.0", "method": "host.create", "id": 1,
    "params": {
      "host": "srv-fichiers-02",
      "name": "Serveur fichiers 02",
      "groups": [{"groupid": "12"}],
      "interfaces": [{"type": 1, "main": 1, "useip": 1, "ip": "192.168.10.62", "dns": "", "port": "10050"}],
      "templates": [{"templateid": "10001"}]
    }
  }' http://192.168.10.50/api_jsonrpc.php
```

> 💡 Récupérez les `groupid`/`templateid` via `hostgroup.get` / `template.get`, ou lisez-les dans l'URL du frontend.

## 27. Templates : principes et bonnes pratiques

Un template = **items + triggers + graphes + dashboards + règles LLD** réutilisables. Principes :

1. **Ne modifiez jamais les templates officiels directement** : clonez-les (`Clone`) puis modifiez la copie. Sinon la prochaine mise à jour écrase vos changements.
2. **Un template = une fonction** : `Linux by Zabbix agent`, `Onduleur SNMP générique`, `Switch Cisco SNMP`. Empilez-les sur les hôtes (un hôte peut avoir N templates).
3. **Macros de template** `{$SEUIL_CPU_CRIT}` : définissez les seuils en macros au niveau template, surchargez-les au niveau hôte. Changer un seuil = changer une macro, pas 50 triggers.
4. **Versionnez** : exportez vos templates en YAML (*Configuration → Templates → Export*) et commitez dans Git.

```
Template "Onduleur SNMP générique"
├── Macros : {$UPS_LOAD_WARN}="80", {$UPS_LOAD_CRIT}="95", {$UPS_BATT_WARN}="30"
├── Items  : ups.load, ups.battery.charge, ups.input.voltage ...
├── Triggers : "Charge onduleur > {$UPS_LOAD_CRIT}"
└── Dashboards : vue onduleur
         │
         ├── Hôte ups-salle-01  (macros surchargées si besoin)
         └── Hôte ups-agence-sud
```

## 28. Templates officiels essentiels (Linux, Windows, SNMP, réseau)

Livrés avec Zabbix (*Configuration → Templates*, groupe *Templates*) :

| Template | Usage | Prérequis |
|---|---|---|
| `Linux by Zabbix agent` | Serveurs Linux | Agent 2 |
| `Windows by Zabbix agent` | Serveurs Windows | Agent 2 Windows |
| `Generic SNMP` | Base pour tout équipement SNMP | SNMP v2c/v3 |
| `Cisco IOS SNMP` / `Cisco Catalyst...` | Switchs/routeurs Cisco | SNMP |
| `UPS SNMP` (selon version) | Onduleurs génériques | SNMP, MIB UPS |
| `VMware Hypervisor / Guest` | ESXi, vCenter | Compte vSphere |
| `Apache by Zabbix agent`, `Nginx…`, `PostgreSQL…`, `MySQL…` | Applicatifs | Plugin Agent 2 |
| `ICMP Ping` | Disponibilité réseau pure | fping |

> 💡 Le template officiel est un **point de départ**, pas une fin : ajustez intervalles et seuils à votre contexte (un ping toutes les 30 s vers un site VSAT = bruit ; passez à 5 min).

## 29. Créer et lier un template personnalisé

Exemple : template `Energie — Salle serveur` regroupant sondes de température.

1. *Configuration → Templates → Create template* : nom `Salle serveur générique`, groupe `Templates/Energie`.
2. Onglet **Macros** (héritées) :

```
{$TEMP_WARN} = 27
{$TEMP_CRIT} = 32
{$HUMI_WARN} = 70
```

3. Créez les items (ici via SNMP, section 49), les triggers référençant les macros, un graphe.
4. Liez aux hôtes : *Configuration → Hosts → [hôte] → Templates → Link new templates*.
5. **Testez sur un hôte pilote** 48 h avant déploiement massif.

## 30. Items : actifs vs passifs, trapper, dépendants

| Type | Qui initie ? | Cas d'usage |
|---|---|---|
| **Zabbix agent (passif)** | Server → Agent | Parc simple, server dimensionné |
| **Zabbix agent (actif)** | Agent → Server | NAT, pare-feu, gros parcs ✅ recommandé |
| **SNMP agent** | Server → équipement | Switchs, onduleurs, imprimantes |
| **SNMP trap** | Équipement → Server | Alertes temps réel (défaut onduleur) |
| **Trapper** | Script externe → Server | `zabbix_sender` depuis vos scripts |
| **Calculé** | Server (formule) | Ratios, différences |
| **Agrégé** | Server (agrégation) | Moyenne d'un groupe, somme |
| **Dépendant** | Préprocessing d'un item maître | Découper un JSON/XML en plusieurs items |
| **HTTP agent** | Server → URL | API REST, sondes web |
| **JMX / IPMI** | Server → équipement | Java, hardware |

Anatomie d'un item : **clé** (`system.cpu.load[all,avg1]`), **intervalle** (`1m`), **période de rétention history/trends** (surcharge possible par item), **type d'information** (numérique, texte, log…), **prétraitement**.

### Envoyer des données avec zabbix_sender (trapper)

```bash
# Depuis un script de sauvegarde, un onduleur sans SNMP, etc.
zabbix_sender -z 192.168.10.50 -s "srv-fichiers-01" \
  -k backup.status -o "OK - 42 Go"
zabbix_sender -z 192.168.10.50 -s "srv-fichiers-01" \
  -k backup.size -o 45056
```

Côté frontend, créez un item de type **Zabbix trapper** avec la clé `backup.status` sur l'hôte `srv-fichiers-01`.

> 💡 `zabbix_sender` est votre **couteau suisse** : tout script maison (sauvegarde, relevé de compteur, sonde DIY) peut alimenter Zabbix en 1 ligne.

## 31. Les clés d'items de l'agent : les indispensables

🖨️ **Référence rapide** (Agent 2, testez avec `zabbix_get -s <ip> -k <clé>`) :

| Clé | Ce qu'elle mesure |
|---|---|
| `agent.hostname` | Nom déclaré par l'agent |
| `agent.ping` | 1 si l'agent répond |
| `system.cpu.num` | Nombre de vCPU |
| `system.cpu.load[all,avg1]` | Load average 1 min |
| `system.cpu.util` | % CPU total |
| `vm.memory.size[total]` / `[available]` | RAM totale / disponible |
| `vfs.fs.size[/,total]` / `[pused]` | Espace disque / % utilisé |
| `net.if.in[eth0]` / `net.if.out[eth0]` | Octets entrants/sortants |
| `net.if.discovery` | LLD des interfaces |
| `vfs.dev.discovery` | LLD des disques |
| `system.uptime` | Uptime en secondes |
| `proc.num[]` | Nombre de processus |
| `systemd.unit.is_active[ssh]` | État d'un service systemd (Agent 2) |
| `system.boottime` | Timestamp du boot |

### UserParameter : vos propres clés

Dans `/etc/zabbix/zabbix_agent2.d/custom.conf` :

```ini
# Température CPU (nécessite lm-sensors)
UserParameter=cpu.temp,sensors -u 2>/dev/null | awk '/Package id 0/{print $4}' | tr -d '+°C' | head -1
# Taille d'un dossier applicatif en Mo
UserParameter=app.dir.size[/var/lib/app],du -sm /var/lib/app 2>/dev/null | cut -f1
```

```bash
systemctl restart zabbix-agent2
zabbix_get -s 127.0.0.1 -k cpu.temp
```

> ⚠️ Les `UserParameter` s'exécutent avec les droits de l'utilisateur `zabbix` : prévoyez les droits sudo ciblés si besoin (visudo, `NOPASSWD` sur la commande exacte).

## 32. Items calculés et agrégés

**Item calculé** (formule sur le même hôte ou d'autres) :

```
# % de RAM utilisée (plus lisible que "available")
100 * (last(//vm.memory.size[total]) - last(//vm.memory.size[available])) / last(//vm.memory.size[total])
```

Syntaxe : `last(/hôte/clé)` — `/` seul = hôte courant. Fonctions dispo : `last`, `min`, `max`, `avg`.

**Item agrégé** (across hosts) :

```
# Charge CPU moyenne de tous les serveurs du groupe "Serveurs/Linux"
grpsum["Serveurs/Linux",system.cpu.load[all,avg1],last,]
grpavg["Serveurs/Linux",vm.memory.size[pavailable],last,]
```

> 💡 Les agrégés sont parfaits pour les **vues direction** : "charge moyenne du parc", "nombre d'onduleurs sur batterie".

## 33. Prétraitement des valeurs (preprocessing)

Transformez la donnée **avant stockage** (onglet *Preprocessing* de l'item). Étapes exécutées dans l'ordre :

