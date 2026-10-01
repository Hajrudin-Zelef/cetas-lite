---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-6
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "attention"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [680, 909]
sha256: 01fa39d974f1b43db99214b6a5fb8a395151a9cd9c7454821b6d152a0a303949
---

# Guide Wazuh — SIEM & XDR Open Source en production

```bash
# Méthode propre : relancer la génération via l'assistant
cd /tmp
sudo bash wazuh-install.sh --generate-config-files   # régénère CA + certs
# Puis réinstallez les composants concernés, ou copiez les certs aux bons endroits
# et redémarrez : wazuh-indexer, wazuh-dashboard, filebeat
```

> ⚠️ En production, utilisez des certificats signés par **votre CA interne** plutôt que l'autosigné : déploiement plus propre sur les agents et les navigateurs du SOC.

### 16.3 Vérifier un certificat

```bash
echo | openssl s_client -connect localhost:9200 2>/dev/null \
  | openssl x509 -noout -subject -issuer -dates
```

---

## 17. L'agent Wazuh : rôle et fonctionnement

L'agent est un programme léger installé sur chaque machine surveillée. Il :

1. **Collecte** : logs système (`/var/log/*`, Event Viewer Windows), FIM, inventaire logiciel, état réseau, exécution de commandes (ex. `netstat`), audit SCA.
2. **Envoie** au manager via le port 1514 (UDP par défaut ; chiffrement symétrique avec la clé d'agent).
3. **Exécute** les réponses actives demandées par le manager (ex. bannir une IP via le pare-feu local).

L'agent **n'analyse pas** : il collecte et transmet. Toute l'intelligence est côté manager. Conséquence : un agent compromis ne peut pas falsifier les règles, mais peut cesser d'envoyer → d'où la supervision des agents (section 24).

Systèmes supportés (4.x) : Debian/Ubuntu, RHEL/CentOS/Rocky, SUSE, Windows 10/11 et Server 2016+, macOS. Architectures x86_64 et ARM64.

---

## 18. Installer un agent sur Debian/Ubuntu

### 18.1 Ajouter le dépôt Wazuh

```bash
# Sur la machine AGENT (pas le manager)
sudo apt update && sudo apt install -y curl gnupg lsb-release
curl -s https://packages.wazuh.com/key/GPG-KEY-WAZUH | sudo gpg --dearmor \
  -o /usr/share/keyrings/wazuh.gpg
echo "deb [signed-by=/usr/share/keyrings/wazuh.gpg] https://packages.wazuh.com/4.x/apt/ stable main" \
  | sudo tee /etc/apt/sources.list.d/wazuh.list
sudo apt update
```

### 18.2 Installer en déclarant le manager à l'avance

```bash
# WAZUH_MANAGER : IP ou FQDN du manager (ou du load balancer)
sudo WAZUH_MANAGER='10.0.0.21' WAZUH_AGENT_GROUP='linux-serveurs' \
  apt install -y wazuh-agent
```

Les variables `WAZUH_MANAGER` et `WAZUH_AGENT_GROUP` pré-remplissent `/var/ossec/etc/ossec.conf`. C'est **la méthode la plus fiable** pour du déploiement automatisé (Ansible, cloud-init).

### 18.3 Démarrer et figer

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now wazuh-agent
sudo apt-mark hold wazuh-agent     # mises à jour planifiées uniquement
```

### 18.4 Vérifier côté agent

```bash
# L'agent a-t-il une clé et voit-il le manager ?
sudo grep -E 'MANAGER_IP|GROUP' /var/ossec/etc/ossec.conf | head -5
sudo tail -20 /var/ossec/logs/ossec.log
# Attendu : "Successfully connected to server" / pas d'erreur "Unable to connect"
```

### 18.5 Déploiement automatisé avec Ansible (exemple)

```yaml
# playbook extrait : déploiement agent sur un parc Debian/Ubuntu
- hosts: serveurs_linux
  become: yes
  vars:
    wazuh_manager: "10.0.0.21"
    wazuh_group: "linux-serveurs"
  tasks:
    - name: Ajouter la clé GPG Wazuh
      ansible.builtin.get_url:
        url: https://packages.wazuh.com/key/GPG-KEY-WAZUH
        dest: /usr/share/keyrings/wazuh.gpg
        mode: '0644'
    - name: Ajouter le dépôt
      ansible.builtin.apt_repository:
        repo: "deb [signed-by=/usr/share/keyrings/wazuh.gpg] https://packages.wazuh.com/4.x/apt/ stable main"
        filename: wazuh
    - name: Installer l'agent
      ansible.builtin.apt:
        name: wazuh-agent
        state: present
      environment:
        WAZUH_MANAGER: "{{ wazuh_manager }}"
        WAZUH_AGENT_GROUP: "{{ wazuh_group }}"
    - name: Activer et démarrer
      ansible.builtin.systemd:
        name: wazuh-agent
        enabled: yes
        state: started
```

---

## 19. Installer un agent sur Windows

### 19.1 Installation graphique (un poste isolé)

1. Téléchargez le MSI : `https://packages.wazuh.com/4.x/windows/wazuh-agent-4.12-1.msi` (adaptez la version).
2. Lancez l'installeur → renseignez l'**IP du manager**.
3. Terminez l'assistant, puis ouvrez **Wazuh Agent** → *Manage → Start*.

### 19.2 Installation silencieuse (parc, GPO, script)

```powershell
# PowerShell administrateur
$msi = "$env:TEMP\wazuh-agent.msi"
Invoke-WebRequest -Uri "https://packages.wazuh.com/4.x/windows/wazuh-agent-4.12-1.msi" -OutFile $msi
msiexec /i $msi /q WAZUH_MANAGER="10.0.0.21" WAZUH_AGENT_GROUP="windows-postes" WAZUH_REGISTRATION_SERVER="10.0.0.21"
Start-Service wazuh-agent
```

Paramètres MSI utiles :

| Paramètre | Rôle |
|---|---|
| `WAZUH_MANAGER` | IP/FQDN du manager |
| `WAZUH_AGENT_GROUP` | Groupe initial |
| `WAZUH_REGISTRATION_SERVER` | Serveur d'inscription (1515) |
| `WAZUH_REGISTRATION_PASSWORD` | Mot de passe d'inscription (voir section 20) |

### 19.3 Vérifier côté Windows

```powershell
Get-Service wazuh-agent
Get-Content "C:\Program Files (x86)\ossec-agent\ossec.log" -Tail 20
```

Attendu : `STATUS: Connected`.

---

## 20. Inscription d'un agent : les 4 méthodes

Un agent doit être **inscrit** (clé partagée avec le manager) avant d'envoyer quoi que ce soit. Quatre méthodes :

### Méthode A — Manuelle (`manage_agents`, la classique)

```bash
# Sur le MANAGER
sudo /var/ossec/bin/manage_agents
# A → Add an agent → nom, IP (ou 'any'), ID auto
# E → Extract key → copiez la clé
# Sur l'AGENT
sudo /var/ossec/bin/manage_agents   # I → Import key → collez
sudo systemctl restart wazuh-agent
```

Simple, mais **ne passe pas à l'échelle**.

### Méthode B — Mot de passe d'inscription (recommandée pour les parcs)

```bash
# Sur le MANAGER : définir un mot de passe d'inscription
echo 'MonMotDePasseInscription' | sudo tee /var/ossec/etc/authd.pass
sudo chmod 640 /var/ossec/etc/authd.pass
sudo chown root:wazuh /var/ossec/etc/authd.pass
```

`/var/ossec/etc/ossec.conf` du manager (service `authd`) :

```xml
<auth>
  <disabled>no</disabled>
  <port>1515</port>
  <use_source_ip>yes</use_source_ip>
  <force_insert>yes</force_insert>
  <force_time>0</force_time>
  <purge>yes</purge>
  <use_password>yes</use_password>
  <ciphers>HIGH:!ADH:!EXP:!MD5:!RC4:!3DES:!CAMELLIA:@STRENGTH</ciphers>
  <!-- <ssl_agent_ca> pour du TLS mutuel en environnement exposé -->
</auth>
```

Côté agent :

```bash
# Sur l'AGENT Linux
sudo /var/ossec/bin/agent-auth -m 10.0.0.21 -P 'MonMotDePasseInscription' -G linux-serveurs
sudo systemctl restart wazuh-agent
```

> Le mot de passe d'inscription transite en clair si vous n'activez pas le TLS sur `authd` : **réservez cette méthode au réseau interne** ou activez `ssl_auto_negotiate`.

### Méthode C — Inscription via API (automatisation)

```bash
# Obtenir un token API (manager)
TOKEN=$(curl -sk -u wazuh-wui:'<MDP>' https://localhost:55000/security/user/authenticate | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
# Créer l'agent
curl -sk -X POST https://localhost:55000/agents \
  -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' \
  -d '{"name":"srv-web-01","ip":"10.0.1.15"}'
```

Pratique pour coupler à votre CMDB ou à la création de VM.

### Méthode D — Provisioning cloud-init / image dorée

Incluez l'agent + un script `agent-auth` au premier boot dans vos templates. **Attention** : ne clonez jamais une VM avec une clé déjà inscrite → doublons d'ID. La séquence sûre dans l'image : agent installé mais **non inscrit** ; inscription au premier démarrage via cloud-init.

### Comparatif

| Méthode | Échelle | Sécurité | Cas d'usage |
|---|---|---|---|
| Manuelle | < 10 agents | Bonne | Lab, exceptions |
| Mot de passe | 10 – 1000 | Moyenne (TLS conseillé) | Parc standard |
| API | 100+ | Bonne | Automatisation, CMDB |
| Cloud-init | 100+ | Bonne | Infra as Code |

---

## 21. Clés d'agent : gestion et rotation

