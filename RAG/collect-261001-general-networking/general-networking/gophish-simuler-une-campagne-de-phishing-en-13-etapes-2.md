---
id: collect-261001-general-networking/general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes-2
title: "1. Mise à jour système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes.md
source_anchor: ""
source_lines: [46, 187]
sha256: 96d47b931fb170dd6117a1da212c214203846a922b2019445e8cce3394e7f289
---

# 1. Mise à jour système

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y wget unzip curl ufw
```
Créez ensuite un utilisateur système dédié. Faire tourner GoPhish sous root est une mauvaise pratique courante : en cas de faille dans l’application (et GoPhish en a connu plusieurs, voir la section pièges plus bas), un compromis limité à un utilisateur sans privilèges réduit fortement l’impact.

```
sudo useradd -r -m -d /opt/gophish -s /usr/sbin/nologin gophish
sudo mkdir -p /opt/gophish
sudo chown gophish:gophish /opt/gophish
```
## Étape 2 : télécharger et installer GoPhish 0.12.1

Rendez-vous sur la page des releases GitHub de GoPhish pour vérifier le nom exact de l’archive correspondant à votre architecture, puis téléchargez-la sur le serveur.

```
cd /opt/gophish
sudo -u gophish wget https://github.com/gophish/gophish/releases/download/v0.12.1/gophish-v0.12.1-linux-64bit.zip
sudo -u gophish unzip gophish-v0.12.1-linux-64bit.zip
sudo -u gophish chmod +x gophish
```
L’archive contient le binaire gophish, un fichier config.json, un dossier static (assets de l’interface admin) et un dossier db contenant les migrations SQLite. Ne renommez pas ces dossiers : le binaire s’attend à les trouver au même niveau que lui.

## Étape 3 : comprendre et adapter config.json

Le fichier config.json définit deux serveurs distincts : le serveur d’administration (interface web de pilotage) et le serveur de phishing (celui qui sert les pages d’atterrissage aux cibles). Par défaut, l’admin écoute sur le port 3333 en HTTPS, et le serveur de phishing sur le port 80 en HTTP.

```
{
  "admin_server": {
    "listen_url": "127.0.0.1:3333",
    "use_tls": true,
    "cert_path": "gophish_admin.crt",
    "key_path": "gophish_admin.key"
  },
  "phish_server": {
    "listen_url": "0.0.0.0:80",
    "use_tls": false,
    "cert_path": "",
    "key_path": ""
  },
  "db_name": "sqlite3",
  "db_path": "gophish.db",
  "migrations_prefix": "db/db_",
  "contact_address": "",
  "logging": {
    "filename": "",
    "level": ""
  }
}
```
Un point de sécurité important : limitez admin_server.listen_url à 127.0.0.1:3333 plutôt qu’à 0.0.0.0:3333. L’interface d’administration ne doit jamais être exposée directement sur Internet ; passez par un tunnel SSH ou un VPN pour vous y connecter. Seul le serveur de phishing, qui sert les pages piégées aux cibles, a besoin d’être accessible publiquement.

## Étape 4 : premier lancement et récupération du mot de passe admin

Lancez GoPhish une première fois en mode manuel pour vérifier que tout démarre correctement avant de créer le service systemd.

```
cd /opt/gophish
sudo -u gophish ./gophish
```
À la première exécution, la console affiche le mot de passe généré automatiquement pour le compte admin. Copiez-le immédiatement : il ne sera plus jamais réaffiché en clair dans les logs. Si vous le perdez, il faudra réinitialiser le mot de passe via la base SQLite ou recréer la base.

## Étape 5 : créer le service systemd

Pour que GoPhish démarre automatiquement au boot et redémarre en cas de crash, créez une unité systemd dédiée.

```
sudo tee /etc/systemd/system/gophish.service > /dev/null << 'EOF'
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
sudo systemctl daemon-reload
sudo systemctl enable --now gophish
sudo systemctl status gophish
```
Un piège fréquent ici : si WorkingDirectory ne pointe pas exactement vers le dossier contenant config.json et gophish.db, le service échoue au démarrage avec une erreur de type « unable to open database file ». Vérifiez toujours ce chemin en premier en cas d'échec.

## Étape 6 : configurer le pare-feu

N'ouvrez que ce qui est strictement nécessaire. Le port 80 doit être accessible depuis Internet pour que les cibles atteignent la page piégée ; le port 3333 ne doit être accessible que depuis votre poste ou via VPN.

```
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow from VOTRE_IP_ADMIN to any port 3333
sudo ufw enable
```
Si vous préférez ne pas exposer le port 3333 du tout, une tunnellisation SSH est plus sûre : ssh -L 3333:127.0.0.1:3333 utilisateur@serveur, puis ouvrez https://127.0.0.1:3333 dans votre navigateur local.

## Étape 7 : configurer le profil d'envoi SMTP

Connectez-vous à l'interface d'administration, puis rendez-vous dans le menu Sending Profiles pour créer un nouveau profil. C'est ce profil qui définit comment GoPhish enverra les emails de la campagne.

| Champ | Exemple de valeur | Remarque | 
|---|---|---|
| Name | SMTP interne test | Nom interne, non visible par les cibles | 
| From | Support IT <[email protected]> | Utilisez un domaine dédié, jamais le domaine réel de l'entreprise | 
| Host | smtp.simulation-interne.fr:587 | Relais SMTP autorisé, avec SPF/DKIM configurés | 
| Username / Password | Identifiants du compte SMTP | À stocker dans un coffre-fort de secrets, pas en clair | 
| Use TLS | Activé | Obligatoire si le relais impose STARTTLS | 

Utilisez systématiquement le bouton Send Test Email avant de continuer. C'est le moyen le plus rapide de détecter un problème d'authentification SMTP ou un blocage par un filtre antispam avant de le découvrir en pleine campagne.

## Étape 8 : créer le modèle d'email

Dans le menu Email Templates, créez un nouveau modèle. GoPhish permet d'utiliser des variables qui seront remplacées automatiquement par les données de chaque cible : {{.FirstName}}, {{.LastName}}, {{.Email}}. Cette personnalisation augmente le réalisme du test, exactement comme le ferait un attaquant qui a récupéré un annuaire d'entreprise.

```
Objet : Action requise : mise à jour de votre mot de passe avant le {{.Date}}
Bonjour {{.FirstName}},
Notre politique de sécurité impose un renouvellement de mot de passe.
Merci de vous connecter avant expiration de votre compte :
{{.URL}}
Cordialement,
Le support informatique
```
La variable {{.URL}} est essentielle : c'est elle qui génère le lien de tracking unique par destinataire, celui qui permettra à GoPhish de savoir précisément qui a cliqué. Ne la remplacez jamais par une URL statique, sous peine de perdre toute la granularité des résultats.

## Étape 9 : créer la page d'atterrissage (landing page)

Dans Landing Pages, deux options s'offrent à vous : écrire le HTML à la main, ou utiliser la fonction Import Site qui clone automatiquement une page existante (par exemple votre propre portail d'authentification interne) pour maximiser le réalisme du test.

Activez l'option Capture Submitted Data si vous voulez enregistrer ce que les cibles saisissent dans le formulaire, et Capture Passwords si le formulaire contient un champ mot de passe. Soyez conscient que ces données deviennent immédiatement sensibles au sens du RGPD : elles doivent être stockées de façon sécurisée et supprimées après l'exploitation pédagogique de la campagne.

Ajoutez systématiquement une page de redirection post-soumission qui explique clairement à l'utilisateur qu'il vient de participer à un exercice de sensibilisation, avec des conseils pour repérer un vrai email de phishing la prochaine fois. C'est ce qui transforme un piège en moment de formation plutôt qu'en simple sanction déguisée.

## Étape 10 : importer la liste des cibles (Users & Groups)

Dans Users & Groups, créez un nouveau groupe et importez votre fichier CSV. Le format attendu est simple : une colonne Email obligatoire, et des colonnes First Name, Last Name, Position facultatives mais utiles pour la personnalisation.

