---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-9
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1433, 1635]
sha256: ec7bebf74a42a4011e2f24c172cbac1094f6297960e477150bff5a0bafdf238f
---

# Guide Zabbix complet — Supervision d'infrastructure en production

result = {"data": []}
try:
    out = subprocess.run(["upsc", "-l"], capture_output=True, text=True, timeout=10)
    for name in out.stdout.split():
        name = name.strip()
        if name:
            result["data"].append({"{#UPSNAME}": name})
except Exception:
    pass
print(json.dumps(result))
```

Test :

```bash
zabbix_get -s 127.0.0.1 -k custom.ups.discovery
# {"data": [{"{#UPSNAME}": "ups-salle-01"}, {"{#UPSNAME}": "ups-agence"}]}
```

Côté template : règle de découverte de type *Zabbix agent*, clé `custom.ups.discovery`, puis prototypes :

```
Item prototype : upsc[{#UPSNAME},battery.charge]
Trigger prototype : "Batterie {#UPSNAME} < {$UPS_BATT_WARN}%"
```

> 💡 La LLD maison + `zabbix_sender`/`upsc` : c'est comme ça qu'on supervise proprement ce que Zabbix ne connaît pas nativement.

## 49. Supervision SNMP : principes (v2c vs v3)

SNMP = le langage commun des équipements réseau et énergie. Deux versions à connaître :

| Critère | SNMP v2c | SNMP v3 |
|---|---|---|
| Authentification | Communauté (mot de passe en clair, ex. `public`) | Utilisateur + auth (SHA) + priv (AES) |
| Chiffrement | Non | Oui (authPriv) |
| Simplicité | Très simple | Plus complexe |
| Recommandation | Labo / réseau de gestion isolé | ✅ **Production** |

**OID** : adresse d'une donnée, ex. `1.3.6.1.2.1.1.3.0` (uptime). **MIB** : dictionnaire qui traduit les OID en noms (`SNMPv2-MIB::sysUpTime.0`).

Outils de diagnostic (paquet `snmp`) :

```bash
apt install -y snmp snmp-mibs-downloader
# Tester un équipement v2c
snmpwalk -v2c -c public 192.168.10.2 1.3.6.1.2.1.1.1.0
# Tester v3
snmpwalk -v3 -l authPriv -u zabbix-mon -a SHA -A 'mot_de_passe_ici' \
  -x AES -X 'mot_de_passe_ici' 192.168.10.90 1.3.6.1.2.1.33
```

> 💡 **Toujours valider avec `snmpwalk` en ligne de commande AVANT de créer l'item Zabbix.** Si `snmpwalk` ne répond pas, Zabbix ne répondra pas non plus — inutile de chercher côté Zabbix.

Dans Zabbix : interface **SNMP** sur l'hôte (port 161/UDP), items de type **SNMP agent**, clé = OID (`1.3.6.1.2.1.1.3.0`) ou nom MIB (`sysUpTime.0` si les MIB sont chargées).

## 50. Superviser un switch via SNMP pas à pas

Exemple : `sw-coeur-01` (192.168.10.2), SNMP v2c communauté `communaute_lecture_ici`.

1. Créez l'hôte : groupe `Réseau/Switchs`, interface **SNMP** `192.168.10.2:161`.
2. Onglet **Macros** (héritées) : `{$SNMP_COMMUNITY}` = `communaute_lecture_ici`.
3. Liez le template `Cisco IOS SNMP` (ou `Generic SNMP` + LLD maison).
4. Items essentiels à vérifier/créer :

| Item | OID / clé | Prétraitement |
|---|---|---|
| sysDescr | `1.3.6.1.2.1.1.1.0` | — |
| sysUpTime | `1.3.6.1.2.1.1.3.0` | diviser par 100 (centisecondes) |
| CPU 5 min | `1.3.6.1.4.1.9.9.109.1.1.1.1.8.1` (Cisco) | — |
| ifInOctets (LLD) | `IF-MIB::ifInOctets[{#IFINDEX}]` | change per second × 8 |
| ifOperStatus (LLD) | `IF-MIB::ifOperStatus[{#IFINDEX}]` | value mapping |

5. Triggers :

```
# Interface down (en excluant les ports administrativement down)
{sw:ifOperStatus[{#IFINDEX}].last()}=2 and {sw:ifAdminStatus[{#IFINDEX}].last()}=1
# CPU switch > 80 %
{sw:cpu.5min.avg(5m)}>80
# Reboot du switch
{sw:sysUpTime.change()}<0
```

> 💡 Sur les LLD d'interfaces : filtrez `{#IFOPERSTATUS}` et excluez les ports non connectés durablement, sinon chaque brassage génère une alerte.

## 51. Superviser un onduleur via SNMP (cas métier — le cœur de ce guide)

C'est **votre** cas d'usage prioritaire : un onduleur non supervisé = une coupure subie au lieu d'une coupure gérée.

### 51.1. La MIB UPS standard (RFC 1628) — `1.3.6.1.2.1.33`

La plupart des onduleurs (APC, Eaton, Vertiv, Schneider, Riello…) exposent la **UPS-MIB** standard. OID clés :

| Donnée | OID (suffixe de 1.3.6.1.2.1.33) | Nom MIB |
|---|---|---|
| Identité / modèle | `1.1.1.1.2.0` | `upsIdentModel.0` |
| État batterie | `1.2.1.1.0` | `upsBatteryStatus.0` (1=unknown, 2=normal, 3=low, 4=depleted) |
| Charge restante batterie | `1.2.3.1.1.2.1` | `upsEstimatedChargeRemaining.1` (%) |
| Tension batterie | `1.2.2.7.0` | `upsBatteryVoltage.0` (dixièmes de V) |
| Température batterie | `1.2.7.0` | `upsBatteryTemperature.0` (°C) |
| Source d'alimentation | `1.1.4.1.1.2.1` | `upsOutputSource.1` (3=normal/mains, 5=battery) |
| Charge en sortie | `1.4.4.1.1.5.1` | `upsOutputPercentLoad.1` (%) |
| Tension d'entrée | `1.3.3.1.1.3.1` | `upsInputVoltage.1` (V) |
| Fréquence d'entrée | `1.3.3.1.1.2.1` | `upsInputFrequency.1` (dixièmes de Hz) |
| Dernier test batterie | `1.2.6.4.0` | `upsTestResultsSummary.0` |
| Autonomie estimée | `1.2.3.1.1.3.1` | `upsEstimatedMinutesRemaining.1` (min) |

Vérifiez d'abord en CLI :

```bash
# L'onduleur parle-t-il la MIB standard ?
snmpwalk -v2c -c communaute_lecture_ici 192.168.10.90 1.3.6.1.2.1.33.1.2.3.1.1.2.1
# Réponse attendue : INTEGER: 100  (batterie à 100 %)
snmpwalk -v2c -c communaute_lecture_ici 192.168.10.90 1.3.6.1.2.1.33.1.4.1.1.2.1
# Réponse attendue : INTEGER: 3  (3 = sur secteur)
```

### 51.2. Valeurs pièges à connaître

- `upsOutputSource` : **3** = secteur normal, **5** = sur batterie, 4 = bypass. Le trigger critique porte sur `=5`.
- `upsBatteryStatus` : **2** = normale, **3** = basse, **4** = épuisée.
- Tensions/fréquences souvent en **dixièmes** : prétraitement `× 0.1`.
- Certains constructeurs (Eaton, Huawei) ont leurs **MIB propriétaires** en plus : chargez-les pour le détail (courants par phase, etc.), mais la RFC 1628 suffit pour l'essentiel.

### 51.3. SNMP traps de l'onduleur (temps réel)

En plus du polling, configurez l'onduleur pour envoyer des **traps** vers le server Zabbix (port 162/UDP). Côté Zabbix :

```bash
# /etc/zabbix/zabbix_server.conf
StartSNMPTrapper=1
SNMPTrapperFile=/var/log/zabbix/zabbix_traps.log
```

```bash
# Réception des traps (snmptrapd)
apt install -y snmptrapd
```

`/etc/snmp/snmptrapd.conf` :

```
authCommunity log,execute,net public
perl do "/usr/share/zabbix/bin/zabbix_trap_receiver.pl";
```

Item de type **SNMP trap** sur l'hôte onduleur, clé `snmptrap["linkDown"]` ou `snmptrap.fallback` pour tout intercepter.

> 💡 Polling (toutes les 30-60 s) + traps (instantané) = le meilleur des deux mondes pour un onduleur.

## 52. Template onduleur : items, triggers et dashboard dédiés

Construisez (ou clonez/adaptez) un template `Onduleur SNMP générique (RFC1628)` :

**Macros** :

```
{$UPS_LOAD_WARN} = 75
{$UPS_LOAD_CRIT} = 90
{$UPS_BATT_WARN} = 50
{$UPS_BATT_CRIT} = 30
{$UPS_BATT_CRIT_MIN} = 15
{$UPS_TEMP_CRIT} = 40
{$SNMP_COMMUNITY} = communaute_lecture_ici
```

**Items** (type SNMP agent, intervalle 1m ; 30s pour les critiques) :

| Nom | Clé (OID) | Prétraitement |
|---|---|---|
| Charge sortie % | `1.3.6.1.2.1.33.1.4.4.1.1.5.1` | — |
| Source alimentation | `1.3.6.1.2.1.33.1.1.4.1.1.2.1` | value mapping |
| Batterie % | `1.3.6.1.2.1.33.1.2.3.1.1.2.1` | — |
| Autonomie restante (min) | `1.3.6.1.2.1.33.1.2.3.1.1.3.1` | — |
| Tension batterie | `1.3.6.1.2.1.33.1.2.2.7.0` | × 0.1 |
| Température batterie | `1.3.6.1.2.1.33.1.2.7.0` | — |
| Tension entrée | `1.3.6.1.2.1.33.1.3.3.1.1.3.1` | — |
| État batterie | `1.3.6.1.2.1.33.1.2.1.1.0` | value mapping |

**Value mappings** (*Administration → General → Value mapping*) :

```
UPS output source : 1→other, 2→none, 3→normal (secteur), 4→bypass, 5→BATTERIE, 6→booster, 7→reducer
UPS battery status : 1→unknown, 2→normale, 3→basse, 4→épuisée
```

**Triggers** (les plus importants de votre supervision) :

```
# 🔴 DISASTER — Onduleur sur batterie
{ups:upsOutputSource.last()}=5
  Nom : "COUPURE SECTEUR : {HOST.NAME} sur batterie"

# 🔴 DISASTER — Batterie critique sur batterie (< 15 %)
{ups:upsOutputSource.last()}=5 and {ups:battery.charge.last()}<{$UPS_BATT_CRIT_MIN}

# 🟠 HIGH — Batterie faible
{ups:battery.charge.last()}<{$UPS_BATT_CRIT}

# 🟠 HIGH — Charge > 90 %
{ups:output.load.avg(2m)}>{$UPS_LOAD_CRIT}

