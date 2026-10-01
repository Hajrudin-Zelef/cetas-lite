---
id: collect-261001-general-networking/general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026-4
title: "Vérifier l'espace disque disponible (minimum 20 Go recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "exploit", "open source"]
source: docs/RAG/collect-261001-general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026.md
source_anchor: ""
source_lines: [251, 341]
sha256: 50ba7f4927046c8833084d8b578d36530def0dafc021c2d62bc19d68f6eaf9d7
---

# Vérifier l'espace disque disponible (minimum 20 Go recommandé)

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Scan bloqué à 0 % | Feed NVT incomplet ou RAM insuffisante | Vérifier Administration → Feed Status et augmenter la RAM à 8 Go minimum | 
| Interface GSA inaccessible sur le port 9392 | Pare-feu UFW bloque le port | `sudo ufw allow 9392/tcp` | 
| Erreur « Feed not synced » | Synchronisation initiale en cours ou échouée | `docker compose logs -f nvt-feed-sync` pour suivre la progression | 
| Mot de passe admin perdu | Logs de création déjà purgés | Réinitialiser via `gvmd --user=admin --new-password=` | 
| Scan authentifié échoue en SSH | Clé publique non déployée sur la cible | Copier la clé avec `ssh-copy-id` vers le compte de scan dédié | 
| Conteneur pg-gvm redémarre en boucle | Volume de données corrompu | Restaurer depuis la dernière sauvegarde `pg_dumpall` | 
| Rapport vide malgré un scan « terminé » | Filtre de qualité de détection (QoD) trop strict | Réduire `min_qod` dans le filtre du rapport, ex. de 70 à 30 | 
| API GMP refuse la connexion Python | Port 9390 non exposé ou certificat TLS invalide | Vérifier `docker compose ps` et le mapping de port dans`docker-compose.yml` | 
| Synchronisation du feed très lente | Bande passante limitée ou blocage proxy sortant | Vérifier la connectivité vers les serveurs Greenbone et augmenter le délai d'expiration du service | 
| Interface GSA lente à charger les rapports volumineux | Base PostgreSQL non indexée ou trop de résultats affichés | Filtrer par niveau de gravité minimum et paginer les résultats plutôt que tout afficher | 

## Conseils avancés pour une gestion de vulnérabilités mature

Une fois l'installation stabilisée, plusieurs pratiques permettent de passer d'un simple scanner ponctuel à un vrai programme de gestion des vulnérabilités, aligné avec les attentes de la directive NIS2 pour les entités essentielles et importantes.

Segmentez vos cibles par criticité métier plutôt que par simple plage IP : un groupe « Serveurs exposés Internet » scanné quotidiennement, un groupe « Postes de travail internes » scanné hebdomadairement, un groupe « OT/IoT » scanné mensuellement avec un profil léger pour éviter les interruptions de service. Ce découpage évite de noyer les équipes sous des rapports ingérables et concentre l'attention là où le risque est le plus élevé.

Couplez OpenVAS à un SIEM comme Wazuh pour corréler automatiquement une vulnérabilité détectée avec une tentative d'exploitation observée dans les logs, plutôt que de traiter les deux flux séparément. Beaucoup d'équipes découvrent une faille critique via un scan, sans jamais vérifier si elle a déjà été exploitée dans les semaines précédentes faute de corrélation.

Mettez en place un tableau de bord de suivi des délais de remédiation (mean time to remediate), avec un objectif réaliste : 72 heures pour le critique, 7 jours pour l'élevé. Un scanner qui tourne sans suivi de correction produit des rapports que personne ne lit après la troisième semaine.

Enfin, ne vous fiez jamais à un seul outil. OpenVAS excelle sur la détection réseau et système, mais reste complémentaire d'un scanner de code statique, d'un outil de test d'intrusion applicatif et d'une veille CERT active, comme celle publiée par le CERT-FR de l'ANSSI.

Pensez aussi à gérer les faux positifs de façon structurée plutôt qu'au cas par cas. Greenbone permet de créer des « Overrides », des règles qui reclassent automatiquement un résultat donné (par exemple une faille sur un service dont vous savez qu'il est déjà corrigé par un correctif compensatoire) sans supprimer l'information du rapport brut. Documentez systématiquement la justification de chaque override avec une date de réexamen, sans quoi ces exceptions s'accumulent et finissent par masquer de vraies régressions lors d'une mise à jour ultérieure.

Sur le plan organisationnel, évitez de confier la lecture des rapports à une seule personne. Un tableau de répartition simple, par exemple par équipe applicative ou par périmètre technique, accélère la remédiation et évite le goulot d'étranglement classique où un seul administrateur croule sous des centaines d'alertes chaque semaine sans pouvoir toutes les traiter à temps.

## Projet complet : script d'automatisation hebdomadaire avec alerte email

Voici un exemple de projet fonctionnel qui lance un scan complet chaque semaine, exporte le rapport en PDF et envoie une alerte email si des failles critiques sont détectées. Adaptez les identifiants et l'UUID de tâche à votre environnement.

```
#!/usr/bin/env python3
import smtplib
import time
from email.mime.text import MIMEText
from gvm.connections import TLSConnection
from gvm.protocols.gmp import Gmp
from gvm.transforms import EtreeCheckCommandTransform
TASK_ID = "11111111-2222-3333-4444-555555555555"
GVM_HOST = "192.168.1.10"
ADMIN_USER = "admin"
ADMIN_PASS = "VotreMotDePasseFort123!"
ALERT_EMAIL = "[email protected]"
def envoyer_alerte(nb_critiques):
    msg = MIMEText(f"{nb_critiques} vulnérabilité(s) critique(s) détectée(s) lors du scan hebdomadaire OpenVAS.")
    msg["Subject"] = "[ALERTE] Scan OpenVAS - Vulnérabilités critiques"
    msg["From"] = "[email protected]"
    msg["To"] = ALERT_EMAIL
    with smtplib.SMTP("localhost") as s:
        s.send_message(msg)
def main():
    connection = TLSConnection(hostname=GVM_HOST, port=9390)
    transform = EtreeCheckCommandTransform()
    with Gmp(connection=connection, transform=transform) as gmp:
        gmp.authenticate(ADMIN_USER, ADMIN_PASS)
        gmp.start_task(TASK_ID)
        # Attendre la fin du scan (vérification toutes les 5 minutes)
        while True:
            status = gmp.get_task(TASK_ID).find(".//status").text
            if status == "Done":
                break
            time.sleep(300)
        report = gmp.get_reports(filter_string=f"task_id={TASK_ID}")
        nb_critiques = report.xpath("count(//result[threat='High'])")
        if int(nb_critiques) > 0:
            envoyer_alerte(int(nb_critiques))
if __name__ == "__main__":
    main()
```
Planifiez ce script avec le cron déjà configuré à l'étape 7, et conservez un historique des rapports dans un dossier daté pour constituer votre dossier de conformité NIS2 ou de certification ISO 27001.

## OpenVAS face aux autres scanners de vulnérabilités en 2026

Pour situer Greenbone Community Edition, voici un comparatif rapide avec les alternatives les plus citées par les équipes sécurité européennes.

| Outil | Modèle | Tests de vulnérabilités | Cas d'usage privilégié | 
|---|---|---|---|
| OpenVAS / Greenbone CE | Open source, gratuit | plus de 160 000 (feed communautaire) | PME, collectivités, labos, conformité NIS2 | 
| Nessus Essentials | Gratuit jusqu'à 16 IP | environ 90 000 | Petits environnements, formation | 
| Qualys VMDR | SaaS commercial | propriétaire, mis à jour en continu | Grands groupes, cloud multi-comptes | 
| Trivy | Open source, gratuit | bases CVE pour conteneurs/IaC | Pipelines CI/CD, images Docker | 

Le choix dépend surtout du périmètre : OpenVAS reste la meilleure option gratuite pour un scan réseau et système large, tandis que Trivy se concentre sur les conteneurs et Qualys cible les environnements cloud à grande échelle avec un budget dédié.

## Coût réel : OpenVAS gratuit face aux solutions commerciales

Le principal argument en faveur d'OpenVAS reste financier. Là où un abonnement Qualys VMDR ou Tenable Nessus Professional se facture par nombre d'actifs scannés, souvent plusieurs milliers d'euros par an dès quelques centaines de machines, Greenbone Community Edition ne coûte rien en licence. Le coût réel se déplace vers l'infrastructure (un serveur dédié avec suffisamment de RAM) et le temps d'administration nécessaire pour maintenir l'outil, lire les rapports et suivre la remédiation.

