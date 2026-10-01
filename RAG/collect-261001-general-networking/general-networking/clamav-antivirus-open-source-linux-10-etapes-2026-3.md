---
id: collect-261001-general-networking/general-networking/clamav-antivirus-open-source-linux-10-etapes-2026-3
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/clamav-antivirus-open-source-linux-10-etapes-2026.md
source_anchor: ""
source_lines: [182, 316]
sha256: 04412ff79e6836010c4d213a67c2bebb10c8c8ac903ab26d007d694b0626965b
---

# Debian / Ubuntu

```
# Ajouts dans /etc/clamav/clamd.conf
OnAccessIncludePath /srv/documents
OnAccessExcludePath /srv/documents/tmp
OnAccessMaxFileSize 50M
OnAccessPrevention no
OnAccessExtraScanning yes
```
Redémarrez clamd, puis lancez clamonacc en tant que service dédié :

```
sudo systemctl restart clamav-daemon
sudo clamonacc --fdpass --log=/var/log/clamav/clamonacc.log
```
`OnAccessPrevention no` signifie que ClamAV journalise et alerte sur les fichiers infectés sans bloquer leur écriture. Passer cette valeur à `yes` active le blocage préventif, mais attention : cela nécessite le module noyau fanotify avec les permissions adéquates, et une mauvaise configuration peut ralentir sensiblement les écritures sur le volume surveillé. Sur un serveur de production à fort trafic disque, testez d’abord en mode journalisation seule pendant plusieurs jours avant d’activer le blocage.

## Étape 7 : projet complet — sécuriser un serveur mail Postfix avec ClamAV Milter

Voici le cas d’usage le plus répandu de ClamAV en entreprise : filtrer les pièces jointes malveillantes avant qu’elles n’atteignent la boîte de réception. L’intégration passe par le protocole Milter, une interface standard supportée par Postfix pour déléguer le filtrage à un service tiers.

Configurez d’abord clamav-milter dans `/etc/clamav/clamav-milter.conf` :

```
# /etc/clamav/clamav-milter.conf
MilterSocket inet:[email protected]
MilterSocketGroup clamav
MilterSocketMode 660
User clamav
ClamdSocket unix:/var/run/clamav/clamd.ctl
OnInfected Reject
RejectMsg "Message rejeté : contenu malveillant détecté par ClamAV"
LogFile /var/log/clamav/clamav-milter.log
LogTime yes
MaxFileSize 100M
```
Puis déclarez le milter dans la configuration Postfix, en modifiant `/etc/postfix/main.cf` :

```
# Ajouts dans /etc/postfix/main.cf
milter_default_action = accept
milter_protocol = 6
smtpd_milters = inet:127.0.0.1:7357
non_smtpd_milters = inet:127.0.0.1:7357
```
Redémarrez les deux services dans l’ordre : clamav-milter d’abord, puis Postfix, sans quoi Postfix tenterait de se connecter à un socket Milter encore inexistant.

```
sudo systemctl restart clamav-milter
sudo systemctl restart postfix
sudo systemctl status clamav-milter postfix
```
Testez l’ensemble de la chaîne en envoyant un mail contenant le fichier EICAR en pièce jointe vers une boîte de test. Le message doit être rejeté à l’étape SMTP, avec le motif `Message rejeté : contenu malveillant détecté par ClamAV` visible dans les logs Postfix (`/var/log/mail.log`) et dans `/var/log/clamav/clamav-milter.log`. Si le mail passe malgré le fichier de test, vérifiez d’abord que `smtpd_milters` pointe bien vers le bon port, puis que le service clamav-milter écoute effectivement sur ce port avec `ss -tlnp | grep 7357`.

Pour les organisations qui préfèrent mettre en quarantaine plutôt que rejeter purement et simplement, remplacez `OnInfected Reject` par `OnInfected Quarantine` et définissez un dossier de quarantaine dédié via `QuarantineDir`. Ce choix évite les faux positifs bloquants pour l’expéditeur tout en isolant le contenu suspect pour analyse manuelle ultérieure.

## Étape 8 : automatiser les scans planifiés avec systemd timers

Même avec clamonacc actif sur les répertoires critiques, un scan complet planifié reste utile pour détecter les menaces dormantes ou les fichiers déposés avant l’activation de la surveillance temps réel. Plutôt qu’une entrée cron classique, un timer systemd offre un meilleur suivi des exécutions et de leurs échecs.

```
# /etc/systemd/system/clamav-full-scan.service
[Unit]
Description=Scan ClamAV complet planifie
After=clamav-daemon.service
[Service]
Type=oneshot
ExecStart=/usr/bin/clamdscan --multiscan --fdpass -r /srv/documents /home
Nice=15
IOSchedulingClass=idle
# /etc/systemd/system/clamav-full-scan.timer
[Unit]
Description=Execution quotidienne du scan ClamAV complet
[Timer]
OnCalendar=*-*-* 02:30:00
Persistent=true
[Install]
WantedBy=timers.target
```
Activez le timer :

```
sudo systemctl daemon-reload
sudo systemctl enable --now clamav-full-scan.timer
systemctl list-timers | grep clamav
```
Le choix de `clamdscan` plutôt que `clamscan` n’est pas anodin : clamdscan délègue le scan au démon clamd déjà actif en mémoire, ce qui le rend nettement plus rapide sur de gros volumes qu’un clamscan classique qui recharge la base à chaque lancement. Les paramètres `Nice=15` et `IOSchedulingClass=idle` évitent que le scan nocturne ne monopolise les ressources si d’autres tâches de maintenance tournent au même moment.

## Étape 9 : gérer la quarantaine et les faux positifs

Aucun moteur antivirus, ClamAV inclus, n’atteint zéro faux positif. Un document PDF avec une structure inhabituelle ou un script légitime utilisant des techniques d’obscurcissement peut parfois déclencher une alerte. Prévoir un processus de quarantaine et de révision limite l’impact opérationnel de ces cas.

Créez un répertoire de quarantaine dédié avec des permissions restreintes :

```
sudo mkdir -p /var/quarantine/clamav
sudo chown clamav:clamav /var/quarantine/clamav
sudo chmod 700 /var/quarantine/clamav
```
Pour les scans planifiés, remplacez la suppression automatique par un déplacement en quarantaine :

`clamdscan --multiscan --move=/var/quarantine/clamav -r /srv/documents`
Si un fichier légitime se retrouve mis en quarantaine, ne le restaurez jamais aveuglément. Vérifiez d’abord le nom de la signature déclenchée (visible dans les logs, sous la forme `NomDeSignature FOUND`), recherchez ce nom sur la base de connaissances ClamAV ou sur un service comme VirusTotal pour confirmer qu’il s’agit bien d’un faux positif, puis signalez-le à l’équipe ClamAV via leur processus officiel de soumission de faux positifs avant de restaurer le fichier. Cette étape de vérification, souvent sautée par manque de temps, est justement celle qui évite qu’un vrai malware ne soit restauré par excès de confiance.

## Étape 10 : superviser, journaliser et alerter

Un antivirus qui tourne sans supervision est un antivirus dont vous ne saurez jamais s’il fonctionne encore. Configurez au minimum une rotation de logs et une alerte sur les détections.

```
# /etc/logrotate.d/clamav
/var/log/clamav/*.log {
    weekly
    rotate 8
    compress
    missingok
    notifempty
    create 640 clamav adm
}
```
Pour une alerte simple par mail à chaque détection dans les logs clamd, un script surveillé par un timer systemd supplémentaire suffit dans la plupart des PME :

```
#!/bin/bash
# /usr/local/bin/clamav-alert.sh
LOGFILE="/var/log/clamav/clamd.log"
FOUND=$(grep "FOUND" "$LOGFILE" | tail -n 20)
if [ -n "$FOUND" ]; then
    echo "$FOUND" | mail -s "Alerte ClamAV : detection sur $(hostname)" [email protected]
fi
```
Si votre organisation dispose déjà d’un SIEM ou d’un outil de centralisation de logs comme Graylog ou Wazuh, préférez y faire remonter directement les fichiers de log ClamAV plutôt qu’un script mail isolé : la corrélation avec d’autres événements de sécurité (tentatives de connexion, alertes réseau) donne une vision bien plus complète qu’une alerte antivirus prise isolément.

## Cycle de vie des versions ClamAV : LTS, non-LTS et dates de fin de support

ClamAV distingue deux types de releases. Les branches LTS (Long Term Support) bénéficient d’au moins trois ans de support de sécurité à compter de leur publication initiale. Les branches non-LTS reçoivent un support plus court, généralement jusqu’à quatre mois après la publication de la release suivante. Connaître ces échéances évite de faire tourner une version qui ne reçoit plus de correctifs de sécurité, un vrai risque sur un serveur exposé à des pièces jointes venant de l’extérieur.

