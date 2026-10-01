---
id: collect-261001-general-networking/general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes-5
title: "1. Mise à jour système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["dpo", "open source"]
source: docs/RAG/collect-261001-general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes.md
source_anchor: ""
source_lines: [304, 403]
sha256: a7b35d7227aa1da256c67675cc093bacd75e9fc168f5c219423cfe4d7589dd37
---

# 1. Mise à jour système

La troisième est le délai moyen entre l'envoi de l'email et le premier signalement. Plus ce délai se raccourcit d'une campagne à l'autre, plus cela indique que les employés développent un réflexe de vérification rapide, ce qui réduit la fenêtre d'exposition en cas d'attaque réelle où chaque minute compte pour l'équipe sécurité qui doit bloquer la propagation.

## Projet complet : script de déploiement automatisé

Voici un script bash qui regroupe l'ensemble des étapes d'installation vues plus haut, à adapter à votre environnement avant exécution.

```
#!/bin/bash
set -e
VERSION="v0.12.1"
ARCHIVE="gophish-${VERSION}-linux-64bit.zip"
# 1. Mise à jour système
apt update && apt upgrade -y
apt install -y wget unzip curl ufw
# 2. Création de l'utilisateur dédié
useradd -r -m -d /opt/gophish -s /usr/sbin/nologin gophish || true
mkdir -p /opt/gophish
chown gophish:gophish /opt/gophish
# 3. Téléchargement et installation
cd /opt/gophish
sudo -u gophish wget "https://github.com/gophish/gophish/releases/download/${VERSION}/${ARCHIVE}"
sudo -u gophish unzip -o "${ARCHIVE}"
sudo -u gophish chmod +x gophish
# 4. Service systemd
cat > /etc/systemd/system/gophish.service << 'EOF'
[Unit]
Description=Gophish - plateforme de simulation de phishing
After=network.target
[Service]
Type=simple
User=gophish
Group=gophish
WorkingDirectory=/opt/gophish
ExecStart=/opt/gophish/gophish
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable --now gophish
# 5. Pare-feu minimal
ufw allow 22/tcp
ufw allow 80/tcp
ufw --force enable
echo "Installation terminée. Consultez 'journalctl -u gophish' pour récupérer le mot de passe admin initial."
```
Exemple de sortie attendue après lancement du service, à récupérer via journalctl -u gophish -n 50 :

```
time="2026-09-07T09:12:03Z" level=info msg="Please login with the username admin and the password XjK9-mPz2-Qw7Y"
time="2026-09-07T09:12:03Z" level=info msg="Starting admin server at https://127.0.0.1:3333"
time="2026-09-07T09:12:03Z" level=info msg="Starting phishing server at http://0.0.0.0:80"
```
## Comparer GoPhish aux alternatives

GoPhish n'est pas le seul outil de simulation de phishing sur le marché. Voici comment il se positionne face aux principales alternatives, open source comme commerciales.

| Outil | Licence | Interface | Cas d'usage typique | 
|---|---|---|---|
| GoPhish | Open source (MIT) | Web légère, API REST | Équipes techniques internes, budget limité | 
| KnowBe4 | Commercial SaaS | Web complète avec reporting avancé | Grandes organisations avec budget dédié à la sensibilisation | 
| Proofpoint Security Awareness | Commercial SaaS | Intégrée à la suite email Proofpoint | Entreprises déjà clientes Proofpoint pour la sécurité email | 
| King Phisher | Open source (BSD) | Application desktop | Pentesteurs cherchant des scénarios plus complexes | 

Pour une équipe technique en interne qui veut garder le contrôle total sur ses données et son infrastructure sans payer de licence, GoPhish reste le choix le plus rationnel. Les solutions commerciales apportent en échange des bibliothèques de modèles prêts à l'emploi, un reporting automatisé plus poussé, et un support client, ce qui compense leur coût pour les organisations qui manquent de temps interne à consacrer à la maintenance de l'outil.

## Foire aux questions

### GoPhish est-il légal à utiliser en France ?

Oui, à condition de disposer d'une autorisation écrite de la direction et du DPO, de limiter la campagne au périmètre autorisé, et d'informer les employés en amont qu'une démarche de sensibilisation incluant potentiellement des simulations peut avoir lieu, conformément aux principes du RGPD.

### Faut-il un serveur dédié pour faire tourner GoPhish ?

Non, une petite VM avec 1 vCPU et 1 Go de RAM suffit pour la plupart des campagnes de taille PME. GoPhish, écrit en Go, a une empreinte mémoire très réduite comparée à des solutions Java ou Python équivalentes.

### Peut-on utiliser GoPhish avec Gmail ou Outlook comme relais SMTP ?

Techniquement oui via un mot de passe d'application, mais ce n'est pas recommandé pour une campagne réelle : les fournisseurs grand public limitent le volume d'envoi et suspendent rapidement les comptes qui envoient des emails ressemblant à du phishing. Utilisez un relais SMTP dédié à la simulation.

### Les résultats de GoPhish sont-ils fiables à 100 % ?

Non. Le tracking d'ouverture dépend du chargement d'une image, que beaucoup de clients de messagerie bloquent par défaut, ce qui sous-estime souvent le taux d'ouverture réel. Le taux de clic et de soumission de données reste en revanche fiable, car il repose sur une action explicite de l'utilisateur.

### Que faire des identifiants captés pendant une campagne ?

Ne les stockez jamais en clair au-delà de la durée nécessaire à l'analyse pédagogique. Supprimez-les de la base GoPhish une fois le débriefing effectué, et ne les communiquez jamais individuellement en dehors d'un cadre confidentiel encadré par le RSSI et le DPO.

### GoPhish peut-il remplacer une formation complète à la cybersécurité ?

Non, c'est un complément, pas un substitut. Une simulation révèle un niveau de vigilance à un instant T, mais ne transmet pas les connaissances nécessaires pour progresser. Associez toujours vos campagnes à un module de formation et à un débriefing collectif.

### Existe-t-il une version cloud hébergée de GoPhish ?

Non, GoPhish est distribué uniquement en auto-hébergement via son binaire téléchargeable sur GitHub. Il n'existe pas d'offre SaaS officielle maintenue par le projet lui-même.

### Comment savoir si mon organisation a besoin de ce type de test ?

Si votre entreprise a déjà subi une compromission par email, gère des données sensibles, ou entre dans le périmètre d'obligations réglementaires comme la directive NIS2, une campagne de simulation régulière fait partie des mesures de base recommandées par l'ANSSI pour réduire le risque d'ingénierie sociale.
