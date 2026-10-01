---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-2
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "license", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [116, 242]
sha256: 53642927a48f2521b9e8d3a3e440963eee31448ce53f0c75974aba6d92c8292e
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

1. Branche le câble console, ouvre PuTTY/secureCRT : **9600 bauds, 8 bits, pas de parité, 1 stop, pas de contrôle de flux**.
2. Allume le boîtier, attends le prompt `Username:` / `Login authentication`.
3. Identifiants initiaux : **à vérifier sur la fiche du modèle exact et le « HUAWEI Security Products Default Usernames and Passwords »** — historiquement `admin` / `Admin@123` sur beaucoup de modèles, mais les versions récentes (V600R007C20+) **n'ont plus d'admin par défaut** : la première connexion web impose de **créer un compte administrateur**.
4. ⚠️ Le système **force le changement du mot de passe par défaut** à la première connexion — prévois le nouveau mot de passe AVANT d'être devant la baie (politique de complexité : 8+ caractères, majuscules/minuscules/chiffres/caractères spéciaux selon version).

## 7. Capacités : sessions, politiques, VPN (ordres de grandeur)

Valeurs indicatives (varient selon modèle et version — **à vérifier sur la fiche du modèle exact**) :

| Capacité | Petit modèle (63xx) | Moyen (66x0) | Gros (6670/6680) |
|----------|---------------------|--------------|------------------|
| Sessions concurrentes | 500 000 – 1 M | 2 – 4 M | 6 – 10 M+ |
| Nouvelles sessions/s | 30 000 – 60 000 | 150 000 – 300 000 | 500 000+ |
| Politiques de sécurité | 5 000 – 10 000 | 20 000 – 50 000 | 50 000+ |
| Tunnels IPSec | 1 000 – 2 000 | 5 000 – 10 000 | 10 000 – 20 000 |
| Utilisateurs SSL VPN simultanés | 100 – 500 | 1 000 – 2 000 | 2 000 – 4 000+ |
| Firewalls virtuels (vsys) | 4 – 16 | 64 – 128 | 256+ |

🔧 Commandes pour voir TES limites réelles : `display firewall session table verbose` (compteur), `display license` (limites logicielles), `display memory` / `display cpu`.

## 8. Versions logicielles : V500 vs V600, quoi choisir

- **V500R005C20** (et antérieures) : la version historique des USG6000 classiques. Stable, très documentée, CLI complète.
- **V600R006 / V600R007** : versions des USG6000E (et portées sur certains USG6000). Apportent la **New Web UI**, des perfs UTM améliorées, de nouvelles signatures.
- ⚠️ **Avant tout upgrade** : vérifie la **matrice de compatibilité** (le firmware doit correspondre EXACTEMENT au modèle), sauvegarde la config (section 115), et prévois un **rollback** (section 117). Un upgrade qui échoue à 2h du matin sans backup, c'est une nuit blanche.
- Pour connaître ta version : `display version` (CLI) ou System > System Information (web).

## 9. USG6000 vs USG6000E vs USG6000F : faut-il migrer ?

| Critère | USG6000 | USG6000E | USG6000F (HiSecEngine) |
|---------|---------|----------|----------------------|
| Statut | Mature / fin de vente progressive | Génération courante | Dernière génération « AI » |
| Version | V500R005 | V600R006/007 | V600R007+ / R024 |
| Web UI | Classique | Nouvelle | Nouvelle |
| Perfs UTM | Bonnes | Meilleures (moteur optimisé) | Les meilleures (accélération IA) |
| Ton besoin | Parc existant, ça tourne | Renouvellement 2024-2026 | Nouveau projet exigeant |

Conseil terrain : **on ne migre pas un firewall qui tourne bien** juste pour la nouveauté. On migre quand : fin de support, besoin de débit UTM supérieur, ou fonctionnalités manquantes (ex. TLS 1.3 inspection poussée). Et on teste TOUJOURS la nouvelle version en maquette avant.

## 10. Positionnement : où mettre l'USG6000 dans ton architecture

Scénarios classiques :
1. **Passerelle Internet PME/ETI** : entre le routeur FAI et le LAN — NAT + politiques + UTM. Le cas n°1, traité en détail section 161.
2. **Firewall de datacenter** : devant les serveurs (zone DMZ), politiques fines par application.
3. **Concentrateur VPN** : IPSec site-à-site vers les agences + SSL VPN pour les nomades (sections 57-72).
4. **Coupe-feu inter-VLAN** : entre zones de sensibilité différentes (prod/bureautique/invités/industriel).
5. **Paire HA actif/passif** : deux USG6000 en hot standby pour les sites qui ne doivent pas tomber (sections 97-106).

Règle d'architecture : **un firewall = un point de passage obligé**. Si un flux peut contourner l'USG (ex. une box 4G branchée en douce), ta politique ne vaut rien. Vérifie le câblage AVANT la config.

---
---

# BLOC B — INITIALISATION : PREMIER DÉMARRAGE

## 11. Checklist de déballage et d'installation physique

📋 Avant de brancher quoi que ce soit :
- [ ] Carton : boîtier, cordons d'alimentation, câble console, kit de rails/oreilles rack, documentation.
- [ ] Vérifier la référence du modèle sur l'étiquette = celle du bon de commande.
- [ ] Rack : 1U ou 3U libres, rails montés, **dégagement avant/arrière** pour l'air (section 5).
- [ ] Alimentation : prise(s) ondulée(s) disponible(s), tension compatible (100-240 V AC ou -48 V DC).
- [ ] Câblage : prévoir les jarretières (cuivre/fibre selon ports), étiqueter CHAQUE câble aux deux bouts.
- [ ] Console : PC portable + câble console + PuTTY configuré en 9600-8-N-1.
- [ ] Noter le **numéro de série** (étiquette + `display esn`) — indispensable pour les licences et le support.
- [ ] Photo de la baie avant/après : ça sauve des heures en dépannage à distance.

## 12. Premier démarrage : séquence et timings

1. Branche l'alimentation, interrupteur ON.
2. Le boot prend **2 à 5 minutes** (plus long avec disques durs / gros modèles).
3. Sur la console, tu vois le décompte mémoire, le chargement du VRP, puis le prompt de login.
4. ⚠️ **Ne coupe JAMAIS l'alimentation pendant le boot ou un upgrade** — corruption du flash possible.
5. Si le boot boucle ou échoue : vérifie l'alimentation, puis boot sur le firmware de secours (section 117).

## 13. Mot de passe initial et création du compte admin

```huawei
# Connexion console, puis :
system-view
# Changer le mot de passe admin existant :
[USG] aaa
[USG-aaa] local-user admin password
        # Saisir le nouveau mot de passe (fictif ici) deux fois
[USG-aaa] local-user admin service-type web telnet ssh terminal
[USG-aaa] local-user admin privilege level 15
[USG-aaa] quit
# Créer un compte nominatif (bonne pratique, voir section 133) :
[USG-aaa] local-user chef-reseau password
[USG-aaa] local-user chef-reseau service-type ssh web
[USG-aaa] local-user chef-reseau privilege level 15
[USG-aaa] quit
save
```

⚠️ Sur V600R007C20+, la **première connexion web** impose la création d'un administrateur (pas de compte par défaut). Le compte créé a le rôle d'administrateur système mais **ne peut pas** être l'administrateur de système virtuel `manager-user@@vsys-name`.

## 14. Configuration de l'interface de management

```huawei
system-view
# Interface de management dédiée (exemple GE0/0/0, adresse fictive) :
[USG] interface GigabitEthernet 0/0/0
[USG-GigabitEthernet0/0/0] ip address 192.168.0.1 24
[USG-GigabitEthernet0/0/0] service-manage enable
[USG-GigabitEthernet0/0/0] service-manage https permit
[USG-GigabitEthernet0/0/0] service-manage ssh permit
[USG-GigabitEthernet0/0/0] service-manage ping permit
[USG-GigabitEthernet0/0/0] quit
# Mettre l'interface de management dans une zone dédiée ou trust :
[USG] firewall zone name MGMT
[USG-zone-mgmt] set priority 90
[USG-zone-mgmt] add interface GigabitEthernet 0/0/0
[USG-zone-mgmt] quit
save
```

Bonnes pratiques :
- L'interface de management a par défaut **192.168.0.1/24** sur beaucoup de modèles (à vérifier) — change-la ou isole-la.
- **HTTPS uniquement** pour le web (désactive HTTP, section 133).
- Ne mets JAMAIS l'interface de management dans la zone **untrust**.
- 📋 Web : System > Network > Interface, puis activer les services de management par interface.

## 15. Mise à l'heure : NTP (obligatoire avant les logs)

Sans heure correcte, tes logs sont inexploitables et les certificats VPN échouent.

