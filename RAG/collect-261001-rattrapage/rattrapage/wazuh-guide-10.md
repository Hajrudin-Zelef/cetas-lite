---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-10
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "distribution"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [1599, 1821]
sha256: 590e2472a2c860833002c5e65a5300740de68e53d0ab022de602351a6219c5c9
---

# Guide Wazuh — SIEM & XDR Open Source en production

Dashboard → *Vulnerability Detection*. Chaque CVE affiche : score CVSS, sévérité, paquet concerné, version installée, version corrigée, agent(s).

Matrice de priorisation conseillée (petite équipe) :

| CVSS | Exposition | Action | Délai |
|---|---|---|---|
| ≥ 9.0 (Critique) | Internet (DMZ, frontal web) | Patch immédiat ou mitigation | 24–48 h |
| ≥ 9.0 | Interne | Planifier | 7 jours |
| 7.0–8.9 (Haute) | Internet | Planifier prioritaire | 14 jours |
| 7.0–8.9 | Interne | Cycle de patch normal | 30 jours |
| < 7.0 | — | Regrouper par lot | 90 jours |

Rituel hebdomadaire : exportez la liste des CVE critiques/hautes, croisez avec vos fenêtres de maintenance, suivez dans un tableau (ou ticket). Wazuh détecte, **vous** pilotez le patch management.

Faux positifs fréquents : CVE sur des paquets dont la distribution a **backporté** le correctif sans changer le numéro de version (classique Debian/Ubuntu). Vérifiez le changelog du paquet avant de paniquer :

```bash
# Exemple : vérifier si le CVE est corrigé sur Ubuntu
apt changelog <paquet> | grep -i CVE-2026-XXXX
```

---

## 41. SCA — audit de conformité CIS : principes

Le module **SCA** (*Security Configuration Assessment*) exécute des **politiques** (fichiers YAML) : une série de contrôles (ex. « le pare-feu est actif », « le compte root ne se connecte pas en SSH »). Chaque contrôle = *pass*, *fail* ou *not applicable*.

Politiques fournies (4.x) : **CIS** pour Debian, Ubuntu, RHEL, Windows 10/Server 2019+, macOS ; aussi PCI-DSS, NIST 800-53, TSC.

C'est votre outil pour :
- préparer un audit (interne ou client),
- mesurer la dérive de configuration dans le temps,
- prouver la conformité (« 92 % des contrôles CIS passent sur le parc »).

---

## 42. Activer et lire les audits SCA

SCA est actif par défaut sur les agents 4.x. Vérifiez côté agent :

```bash
sudo grep -A5 "<sca>" /var/ossec/etc/ossec.conf
```

```xml
<!-- Extrait ossec.conf agent Linux -->
<sca>
  <enabled>yes</enabled>
  <scan_on_start>yes</scan_on_start>
  <interval>12h</interval>
  <skip_nfs>yes</skip_nfs>
  <policies>
    <policy>/var/ossec/ruleset/sca/cis_debian12.yml</policy>
  </policies>
</sca>
```

Lecture : Dashboard → *Security Configuration Assessment* → choisissez l'agent → score global + détail par contrôle. Chaque contrôle *fail* donne la **remédiation** (la politique inclut `remediation`).

Automatisez le suivi :

```bash
# Via API : récupérer le dernier scan SCA d'un agent
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/sca/007?pretty=true' | grep -E '"pass"|"fail"|"score"'
```

> Les scans SCA consomment du CPU agent pendant quelques minutes : planifiez `interval` hors heures de pointe pour les serveurs chargés.

---

## 43. Écrire une politique SCA personnalisée

Vos exigences internes (ex. « l'antivirus d'entreprise doit tourner », « tel montage doit être chiffré ») n'existent pas dans CIS → écrivez votre politique.

```yaml
# /var/ossec/etc/shared/linux-serveurs/sca_metier.yml
policy:
  id: "soc_metier_linux"
  file: "sca_metier.yml"
  name: "Politique interne - serveurs Linux"
  description: "Contrôles métier du SOC."
  references:
    - "Interne: POL-SEC-001"

requirements:
  title: "Assurer le socle de sécurité interne"
  description: "Contrôles non couverts par CIS."
  condition: all

checks:
  - id: 10001
    title: "L'agent Wazuh est actif"
    description: "L'agent doit tourner en permanence."
    rationale: "Sans agent, aucune supervision."
    remediation: "systemctl enable --now wazuh-agent"
    compliance:
      - cis: ["1.1.1"]
    rules:
      - 'c:systemctl is-active wazuh-agent -> r:active'

  - id: 10002
    title: "Aucun compte avec UID 0 autre que root"
    description: "Un seul UID 0 autorisé."
    remediation: "Supprimer ou corriger le compte en trop."
    rules:
      - 'c:awk -F: ''$3 == 0 {print $1}'' /etc/passwd -> r:root && c:awk -F: ''$3==0'' /etc/passwd | wc -l -> r:^1$'

  - id: 10003
    title: "Le pare-feu est actif"
    description: "UFW ou nftables doit filtrer."
    remediation: "ufw enable"
    condition: any
    rules:
      - 'c:ufw status -> r:active'
      - 'c:nft list ruleset -> r:table'
```

Déclarez-la dans l'`agent.conf` du groupe :

```xml
<agent_config>
  <sca>
    <policies>
      <policy>/var/ossec/etc/shared/linux-serveurs/sca_metier.yml</policy>
    </policies>
  </sca>
</agent_config>
```

> Testez chaque check **manuellement** sur une machine avant de déployer : un check mal écrit = 100 % de *fail* et un SOC qui n'y croit plus.

---

## 44. Active Response : le blocage automatique

L'**Active Response** exécute un script (sur l'agent, ou sur le manager) quand une règle se déclenche. Cas d'usage roi : **bannir automatiquement l'IP** qui fait du force brute SSH.

Chaîne :

```
Règle 5763 (force brute SSH, niveau 10)
   │  <active-response> associé dans ossec.conf
   ▼
Manager → ordre à l'agent : exécute firewall-drop.sh add <ip>
   │  L'agent ajoute une règle iptables/nftables DROP
   ▼  Après <timeout> (ex. 10 min), l'agent retire la règle (delete)
```

Deux modes :

| Mode | Où s'exécute | Usage |
|---|---|---|
| Agent (défaut) | Sur l'agent qui a levé l'alerte | Bloquer l'attaquant au plus près |
| Manager (`location: server`) | Sur le manager | Centraliser (ex. alimenter un pare-feu central via script custom) |

> ⚠️ L'active response est une **arme** : un faux positif = vous vous bannissez vous-même (ou un client). Toujours : allowlist (section 45), timeout court au début, tests en lab.

---

## 45. Configurer le blocage automatique d'attaque IP (firewall-drop)

### Étape 1 — Activer côté manager

`/var/ossec/etc/ossec.conf` du **manager** :

```xml
<command>
  <name>firewall-drop</name>
  <executable>firewall-drop.sh</executable>
  <timeout_allowed>yes</timeout_allowed>
</command>

<active-response>
  <command>firewall-drop</command>
  <location>local</location>              <!-- sur l'agent source -->
  <rules_id>5763</rules_id>               <!-- force brute SSH (exemple) -->
  <timeout>600</timeout>                  <!-- 10 minutes -->
</active-response>
```

Règles candidates au blocage auto (à adapter) :

| Règle | Description | Timeout conseillé |
|---|---|---|
| 5763 | Force brute SSH (niveau 10) | 600 s |
| 31151 | Force brute web (selon ruleset) | 600 s |
| 100100+ | Vos règles métier critiques | 1800 s |

### Étape 2 — Allowlist : ne jamais se bannir soi-même

`/var/ossec/etc/ossec.conf` du manager :

```xml
<global>
  <white_list>127.0.0.1</white_list>
  <white_list>10.0.0.0/8</white_list>       <!-- votre LAN admin -->
  <white_list>198.51.100.0/24</white_list>  <!-- IP publiques du SOC / VPN -->
</global>
```

> Mettez-y **toutes** vos IP de management, votre VPN, votre prestataire infogéré. Le jour où vous testez un scan depuis votre poste, vous serez content de l'avoir fait.

### Étape 3 — Vérifier que ça marche (lab !)

```bash
# Depuis une machine de test NON allowlistée, bourrinez SSH :
for i in $(seq 1 12); do ssh -o ConnectTimeout=2 test@<cible> <<< "mauvais"; done
# Sur la cible :
sudo iptables -L -n | grep DROP     # ou : sudo nft list set inet filter ...
sudo tail -5 /var/ossec/logs/active-responses.log
```

Vous devez voir l'IP bannie puis **débannie** après le timeout. Si le débannissement n'a pas lieu : vérifiez `timeout_allowed` et les droits du script.

---

## 46. Active Response avancé : scripts personnalisés

Les scripts vivent dans `/var/ossec/active-response/bin/` (manager **et** agents : le script doit exister **où il s'exécute**).

Exemple : script qui **désactive un compte local** après 5 échecs (Linux) :

