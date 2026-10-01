---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-22
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["arr", "benchmarks", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [3521, 3696]
sha256: 527d6f07187b6aa5fb622ae1006239d320be7cb19f3e65462c136630c5e36119
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Les trois questions à poser devant tout devis :
1. Combien de **watts** en charge réelle, pas en TDP ?
2. Combien d'**euros de licences** par watt utile ?
3. Combien de **kWh sur 5 ans**, froid compris ?

Si le commercial ne peut pas répondre aux trois, changez de commercial —
pas de serveur.

## 201. Index rapide des tableaux du guide

- §3 : générations EPYC — §6 : références Turin — §8 : SP5/SP6 —
  §10 : choix EPYC par workload — §12/13/14 : Xeon 6900P/6700E/6700P —
  §15 : sockets Intel — §16 : Xeon 6+ — §17/96 : MRDIMM —
  §19 : choix Xeon par workload — §20/21/22 : comparatifs EPYC vs Xeon —
  §23 : 1S/2S/4S — §30 : BP/cœur — §31/32/33 : configs types —
  §39 : formats de barrettes — §40 : fréquences DDR5 — §44 : capacités —
  §46 : impact canaux vides — §54 : conso DIMM — §56/97 : CXL —
  §62 : loi cubique ventilateurs — §67 : températures —
  §69/73/75 : formats/profondeur/poids — §78 : kW/rack —
  §84 : 80 PLUS — §85 : conso par config — §89 : PDU —
  §93/94 : Venice/Diamond Rapids — §95 : DDR6 — §101/102 : TCO —
  §103 : benchmarks — §127/128 : BOMs — §142 : BP par plateforme —
  §171-174 : cas d'école — §175 : tableau final — §189 : PCIe —
  §190 : réseau.

---

*Fin du guide — 201 sections, rédigé et vérifié le 27/09/2026.*
*Pour Zelef — chiffres, BOMs, énergie. Les faits datés se revérifient,*
*les méthodes restent.*

---

## 202. Retour terrain : la baie qui disjonctait l'été

Cas réel type (anonymisé, schéma classique) : une baie de 8 serveurs 2S
fonctionnait parfaitement 10 mois par an, puis disjonctait 2-3 fois chaque
été. Diagnostic :
1. La climatisation de la salle, sous-dimensionnée, laissait la température
   monter à 32 °C l'après-midi.
2. Les ventilateurs serveurs passaient à plein régime (loi cubique :
   +150 W par serveur).
3. La PDU, dimensionnée « juste » sur la conso nominale, dépassait son
   calibre avec le surplus ventilateurs + le rendement dégradé des
   alimentations à chaud.
4. Déclenchement du disjoncteur PDU → 8 serveurs d'un coup.

Corrections : caches de baie (il en manquait 12U), consigne clim à 25 °C
max avec alarme, PDU 32 A au lieu de 20 A, cTDP -10 % l'été. **Coût total :
moins de 2 000 €. Coût des 3 pannes : bien plus.** La morale : les pannes
d'été sont des pannes de dimensionnement, pas de fatalité.

## 203. Retour terrain : le cluster qui ramait à cause d'une barrette

Symptôme : sur un cluster de 4 nœuds identiques, un nœud affichait des
performances 15 % inférieures en base de données, sans erreur apparente.
Pistes explorées (dans l'ordre) : réseau, disques, OS, BIOS.
Cause réelle : **une barrette remplacée 6 mois plus tôt n'était pas de la
même référence** (même capacité, rang différent). Le contrôleur avait
désactivé l'entrelacement optimal sur 2 canaux → bande passante asymétrique
→ le moteur DB, NUMA-sensible, perdait 15 %.

Leçon : après tout remplacement de barrette, **revérifiez `dmidecode` et
les compteurs EDAC**, et comparez les débits (STREAM) entre nœuds
identiques. Un écart > 5 % entre nœuds « identiques » = enquête.

## 204. Test de charge avant production : protocole complet

Au-delà de la réception unitaire (section 187), le test de charge valide
l'ensemble (serveurs + réseau + électrique + froid) :

| Phase | Durée | Objectif |
|---|---|---|
| Montée en charge | 2 h | vérifier les courbes ventilateurs et la stabilité |
| Palier 100 % (mprime + fio) | 4 h | pic électrique, pic thermique |
| Démarrage simultané | 1 test | appel de courant, tenue onduleur |
| Coupure secteur simulée | 1 test | bascule onduleur, arrêt ordonné NUT |
| Nuit en idle | 12 h | socle de consommation, PUE réel |

Mesures à consigner : kW par PDU (toutes les 5 min), températures
entrée/sortie, PUE instantané, autonomie batterie mesurée vs calculée.
**Un test de charge raté avant production évite une panne en production.**

## 205. Gestion des alertes : du bruit au signal

Une supervision qui alerte trop n'alerte plus. Règles de tri :
- **Critique** (SMS/appel) : panne ventilateur, panne alimentation,
  température CPU > 90 °C, onduleur sur batterie, erreurs ECC non corrigées.
- **Warning** (ticket) : ECC corrigées en hausse, température d'entrée >
  30 °C, PDU > 80 %, firmware obsolète.
- **Info** (rapport hebdo) : statistiques, tendances, top consommateurs.

Anti-patterns : alerter sur chaque pic CPU (normal), ne pas alerter sur la
batterie onduleur (critique), avoir 500 alertes non lues (désensibilisation).
**Relisez vos seuils tous les 6 mois** : une alerte jamais déclenchée est
inutile, une alerte toujours déclenchée est mal réglée.

## 206. Plan de continuité électrique : l'exercice annuel

Une fois par an, en heures creuses, avec toutes les équipes :
1. Coupure secteur simulée (ou réelle si possible) → vérification bascule
   onduleur (0 ms en double conversion).
2. Démarrage groupe électrogène → mesure du temps (objectif < 60 s).
3. Arrêt ordonné d'un périmètre non critique via NUT → vérification des
   logs et des délais.
4. Redémarrage séquencé (pas simultané !) → vérification de l'appel de
   courant par vague.
5. Débrief écrit : écarts, actions correctives, mise à jour des procédures.

**Un plan non testé est un vœu pieux.** L'exercice annuel est aussi le
moment où l'on découvre que la batterie a 4 ans, que le groupe n'a pas
tourné depuis 18 mois, et que personne ne connaît le mot de passe du BMC.

---

## 207. Annexe : commandes Linux indispensables

```bash
# --- CPU ---
lscpu                          # topologie, frequences, flags
cat /proc/cpuinfo | grep "model name" | head -1
cpupower frequency-info         # gouverneur, frequences reelles

# --- Memoire ---
dmidecode -t memory | grep -E "Size|Speed|Manufacturer|Part Number"
free -h
cat /sys/devices/system/edac/mc/mc*/ce_count   # erreurs ECC corrigees
numactl --hardware             # topologie NUMA

# --- Thermique / BMC ---
ipmitool sdr type Temperature  # sondes via IPMI
ipmitool sdr type Fan          # ventilateurs
ipmitool sel list              # journal d'evenements BMC

# --- Disques ---
nvme smart-log /dev/nvme0     # usure, temperature NVMe
lsblk -o NAME,SIZE,TYPE,MODEL

# --- Reseau ---
ip -s link                     # compteurs d'erreurs
ethtool eth0 | grep Speed

# --- Charge ---
stress-ng --cpu 0 --timeout 3600s   # test de charge CPU
fio --name=test --rw=randread --bs=4k --size=10G --numjobs=8  # I/O
```

## 208. Annexe : script de health-check quotidien (exemple)

```bash
#!/bin/bash
# health-check.sh : a executer chaque matin via cron, sortie vers supervision
echo "=== $(hostname) $(date) ==="
echo "--- ECC ---"
for f in /sys/devices/system/edac/mc/mc*/ce_count; do
  echo "$f : $(cat $f)"
done
echo "--- Temperatures CPU ---"
grep . /sys/class/thermal/thermal_zone*/temp 2>/dev/null | head -5
echo "--- NVMe (usure %) ---"
for d in /dev/nvme*n1; do
  echo "$d : $(nvme smart-log $d 2>/dev/null | grep percentage_used)"
done
echo "--- Memoire ---"
free -h | head -2
echo "--- Uptime ---"
uptime
```

Adaptez les seuils à votre supervision (Zabbix : UserParameter, ou
node_exporter + textfile). L'important n'est pas l'outil, c'est la
**régularité** : un check quotidien lit les compteurs avant qu'ils ne
débordent.

## 209. Annexe : template de fiche serveur

