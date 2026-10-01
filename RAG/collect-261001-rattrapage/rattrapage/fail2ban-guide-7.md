---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-7
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["arr", "attention", "datacenter", "incident"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [1444, 1712]
sha256: 448acd2e30e3df97c37da4de44185fb1e5e81014c5d6e211771796ea05a19497
---

# Guide fail2ban — Le bouclier anti-brute-force

- [ ] Votre **IP publique actuelle** (si elle est fixe ; sinon, voir section 25)
- [ ] Le **bastion** depuis lequel vous administrez
- [ ] Les **sondes de supervision** (Zabbix, Prometheus, Nagios, Uptime Kuma) — elles ouvrent des connexions en boucle et se font bannir par les jails agressifs (nginx-badbot, sshd mode agressif)
- [ ] Les **reverse proxy** internes (si fail2ban voit l'IP du proxy au lieu du client — voir section 52)
- [ ] Le **loopback** (127.0.0.1/8, ::1) — toujours
- [ ] Les **IP des collègues** en télétravail fixe

### Syntaxe acceptée

```
ignoreip = 127.0.0.1/8 ::1 192.168.1.14 10.0.0.0/8 2001:db8::/32
```

IP uniques, CIDR v4 et v6, séparés par des espaces. On peut aussi utiliser des noms DNS (résolus au démarrage — **à éviter** : si le DNS est down au boot, fail2ban échoue).

### Vérifier qu'une IP est ignorée

```bash
sudo fail2ban-client get sshd ignoreip
# 127.0.0.1/8 ::1 192.168.10.0/24 ...
```

### Le cas des IP dynamiques (télétravail, 4G)

Si votre IP publique change, `ignoreip` statique ne suffit pas. Options :

1. **VPN d'administration** : vous vous connectez toujours via le VPN, et c'est le **réseau VPN** qui est en `ignoreip`. C'est la bonne solution.
2. **Ne jamais exposer SSH** sans VPN/bastion (recommandé).
3. En dernier recours : `bantime` court sur sshd (30m) pour limiter la casse en cas d'auto-ban.

---

## 25. Ne pas se bannir soi-même : 7 techniques

L'auto-bannissement est **l'erreur n°1** (section 39). Voici la défense en profondeur.

### Technique 1 : ignoreip complet (base)

Voir section 24. C'est non négociable.

### Technique 2 : accès console de secours

Avant toute modification de fail2ban sur un serveur distant :

- Ouvrez une **console hors-bande** (console Proxmox/VMware, IPMI/iDRAC, console du provider cloud).
- Ou gardez une **session SSH déjà ouverte** : un ban fail2ban bloque les NOUVELLES connexions, pas toujours les sessions établies (selon l'action : REJECT sur INPUT bloque aussi les paquets des sessions existantes — ne comptez pas dessus).

> Règle : **ne touchez jamais à fail2ban en SSH sans console de secours**. Le jour où vous en aurez besoin, vous comprendrez.

### Technique 3 : tester avec un bantime court

```bash
# Pendant les tests, bantime très court :
sudo fail2ban-client set sshd bantime 60
# ... vos tests ...
# Puis remettez la vraie valeur et rechargez jail.local
```

Si vous vous bannissez pendant les tests, 60 secondes plus tard c'est fini.

### Technique 4 : le « dead man's switch »

Avant de recharger une config risquée, programmez un garde-fou :

```bash
# Dans 5 minutes, fail2ban s'arrête (et les règles de ban tombent avec actionstop)
echo "systemctl stop fail2ban" | at now + 5 minutes
# Si tout va bien, annulez :
atrm $(atq | cut -f1)
```

Variante sans `at` :

```bash
(sleep 300 && systemctl stop fail2ban) &
# puis kill %1 si tout va bien
```

### Technique 5 : jail de test séparé

Testez vos nouveaux filters sur un jail dédié avec des seuils inoffensifs, pas directement sur `sshd` :

```ini
[sshd-test]
enabled  = true
filter   = sshd
logpath  = /var/log/auth.log
port     = 22222
maxretry = 100
bantime  = 60
action   = iptables[name=sshd-test, port=22222]
```

### Technique 6 : vérifier AVANT de recharger

```bash
sudo fail2ban-client --test   # valide la syntaxe de toute la config
sudo fail2ban-regex /var/log/auth.log /etc/fail2ban/filter.d/sshd.conf | tail -3
```

### Technique 7 : procédure d'urgence écrite

Affichez dans votre documentation d'équipe :

```
URGENCE AUTO-BAN — serveur <nom>
1. Console Proxmox : https://pve.exemple.fr:8006 (datacenter > srv-web > Console)
2. Login root (mot de passe dans le coffre)
3. fail2ban-client set sshd unbanip <VOTRE_IP>
   ou : fail2ban-client unban --all
4. Vérifier : fail2ban-client status sshd
```

---

## 26. fail2ban-client : usage quotidien

### Statuts

```bash
# Vue d'ensemble : jails actifs
sudo fail2ban-client status

# Détail d'un jail : compteurs + IP bannies
sudo fail2ban-client status sshd

# Détail de TOUS les jails en une commande
sudo fail2ban-client status --all
```

### Bannir / débannir manuellement

```bash
# Bannir une IP à la main (utile en incident)
sudo fail2ban-client set sshd banip 203.0.113.45

# Débannir
sudo fail2ban-client set sshd unbanip 203.0.113.45

# Tout débannir d'un jail
sudo fail2ban-client set sshd unban --all
```

### Paramètres à chaud (sans reload)

```bash
# Changer le bantime d'un jail immédiatement
sudo fail2ban-client set sshd bantime 7200

# Changer maxretry / findtime
sudo fail2ban-client set sshd maxretry 3
sudo fail2ban-client set sshd findtime 600

# Ajouter une IP à ignoreip sans toucher au fichier
sudo fail2ban-client set sshd addignoreip 198.51.100.99
sudo fail2ban-client get sshd ignoreip
```

> Attention : les changements `set` sont **perdus au prochain reload/restart**. Pour les rendre durables, modifiez `jail.local`.

### Recharger la configuration

```bash
# Recharge en douceur (recharge les jails modifiés, garde les bans)
sudo fail2ban-client reload

# Recharge un seul jail
sudo fail2ban-client reload sshd

# Redémarrage complet (coupe tous les bans puis les restaure via la db)
sudo fail2ban-client restart
```

### Ping et diagnostic

```bash
# Le serveur répond-il ?
sudo fail2ban-client ping
# Server replied: pong

# Version
sudo fail2ban-client --version
```

---

## 27. Débannir : le classique du lundi

Lundi 8h : « je n'arrive plus à me connecter au serveur ». Dans 90 % des cas : l'utilisateur (ou vous) s'est fait bannir le week-end.

### Procédure standard

```bash
# 1. Identifier le jail concerné
sudo fail2ban-client status | grep "Jail list"

# 2. Vérifier que l'IP est bien bannie
sudo fail2ban-client status sshd | grep -A5 "Banned IP list"

# 3. Débannir
sudo fail2ban-client set sshd unbanip 198.51.100.23

# 4. Vérifier dans les logs
sudo grep "Unban 198.51.100.23" /var/log/fail2ban.log
```

### Si on ne sait pas quel jail a banni

```bash
# Chercher dans le log fail2ban
sudo grep "Ban 198.51.100.23" /var/log/fail2ban.log
# 2026-09-26 08:15:01,234 fail2ban.actions [1234]: NOTICE [sshd] Ban 198.51.100.23
# → c'était le jail sshd

# Ou scanner tous les jails
for j in $(sudo fail2ban-client status | grep "Jail list" | sed 's/.*://;s/,//g'); do
  echo "== $j =="
  sudo fail2ban-client status "$j" | grep -A3 "Banned IP list"
done
```

### Après le débannissement : traiter la cause

Un débannissement sans correction = le même appel mardi.

| Cause probable | Correction |
|---|---|
| Mot de passe expiré / changé | mettre à jour le client mail, le script, la conf |
| Smartphone en boucle | corriger le mot de passe sur le téléphone |
| Script de supervision mal configuré | corriger le script OU mettre sa sonde en `ignoreip` |
| Vrai attaquant | ne pas débannir, analyser (section 33) |
| C'est vous, en test | ajouter votre IP/réseau en `ignoreip` |

### Le cas du ban persistant (recidive)

Si l'IP est aussi dans `recidive`, il faut la débannir des **deux** jails :

```bash
sudo fail2ban-client set sshd unbanip 198.51.100.23
sudo fail2ban-client set recidive unbanip 198.51.100.23
```

---

## 28. Bannissement temporaire vs permanent

### Temporaire (défaut) : le bon choix 95 % du temps

```ini
[sshd]
bantime = 1h
```

Avantages : auto-nettoyant, pas de risque d'accumulation, pardonne les faux positifs.

### Permanent : quand et comment

```ini
# Ban permanent pour un jail donné
[nginx-botsearch]
bantime = -1
```

`bantime = -1` = banni jusqu'au redémarrage de fail2ban (ou unban manuel). Avec la base sqlite persistante (section 29), le ban survit même au redémarrage.

**Quand c'est justifié :**
- Jail `nginx-botsearch` / `nginx-badbot` : aucun faux positif possible, attaquant certain.
- IP identifiée comme source d'attaque ciblée (après analyse AbuseIPDB, section 33).

