---
id: collect-261001-general-networking/general-networking/nmap-scanner-un-reseau-en-12-etapes-2026-4
title: "Nmap version 7.991 ( https://nmap.org )"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/nmap-scanner-un-reseau-en-12-etapes-2026.md
source_anchor: ""
source_lines: [221, 317]
sha256: 74360351bd23d750714c86db2e6859eb32652d5101aeec13837e42d13a3ece24
---

# Nmap version 7.991 ( https://nmap.org )

```
#!/usr/bin/env python3
import nmap
import json
import smtplib
from email.mime.text import MIMEText
from pathlib import Path
RESEAU_CIBLE = "192.168.1.0/24"
FICHIER_HISTORIQUE = Path("/var/log/audit_nmap/dernier_scan.json")
EXPEDITEUR = "[email protected]"
DESTINATAIRE = "[email protected]"
def lancer_scan():
    scanner = nmap.PortScanner()
    scanner.scan(hosts=RESEAU_CIBLE, arguments="-sS -sV -T4 --top-ports 200")
    resultat = {}
    for hote in scanner.all_hosts():
        ports_ouverts = []
        if "tcp" in scanner[hote]:
            for port, info in scanner[hote]["tcp"].items():
                if info["state"] == "open":
                    ports_ouverts.append({"port": port, "service": info.get("name", "?")})
        resultat[hote] = ports_ouverts
    return resultat
def comparer_et_alerter(nouveau_scan):
    ancien_scan = {}
    if FICHIER_HISTORIQUE.exists():
        ancien_scan = json.loads(FICHIER_HISTORIQUE.read_text())
    nouveaux_ports = []
    for hote, ports in nouveau_scan.items():
        anciens_ports = {p["port"] for p in ancien_scan.get(hote, [])}
        for p in ports:
            if p["port"] not in anciens_ports:
                nouveaux_ports.append(f"{hote}:{p['port']} ({p['service']})")
    if nouveaux_ports:
        envoyer_alerte(nouveaux_ports)
    FICHIER_HISTORIQUE.parent.mkdir(parents=True, exist_ok=True)
    FICHIER_HISTORIQUE.write_text(json.dumps(nouveau_scan, indent=2))
def envoyer_alerte(nouveaux_ports):
    contenu = "Nouveaux ports détectés depuis le dernier audit :\n\n" + "\n".join(nouveaux_ports)
    message = MIMEText(contenu)
    message["Subject"] = "[Audit Nmap] Nouveaux ports ouverts détectés"
    message["From"] = EXPEDITEUR
    message["To"] = DESTINATAIRE
    with smtplib.SMTP("localhost") as serveur:
        serveur.send_message(message)
if __name__ == "__main__":
    scan_actuel = lancer_scan()
    comparer_et_alerter(scan_actuel)
```
Installez la dépendance et testez le script manuellement avant de l’automatiser :

```
pip install python-nmap
sudo python3 audit_reseau.py
```
Une fois validé, planifiez son exécution nocturne via cron :

```
sudo crontab -e
# Ajouter la ligne suivante pour un scan chaque nuit à 2h15
15 2 * * * /usr/bin/python3 /home/audit/audit_reseau.py
```
Vous disposez maintenant d’un projet complet et fonctionnel : un scan différentiel automatisé qui vous alerte uniquement sur les changements, plutôt que de vous noyer sous un rapport complet chaque matin. C’est exactement la logique qu’appliquent les intégrations Wazuh-Nmap en production, à plus grande échelle.

## Bonnes pratiques après un scan

Un scan Nmap n’a de valeur que si ses résultats déclenchent une action. Une fois le rapport en main, fermez systématiquement les ports non justifiés par un besoin métier documenté, mettez à jour les services dont la version détectée correspond à une CVE connue, et si vous avez utilisé `--script vuln`, traitez en priorité tout résultat marqué `VULNERABLE (Exploitable)`. Conservez vos rapports XML dans un dossier daté : c’est votre seule preuve de conformité en cas d’audit NIS2, un sujet que nous détaillons dans notre guide sur la directive NIS2.

Il est également recommandé de documenter chaque campagne de scan dans un registre interne : date, périmètre exact, personne à l’origine du scan et référence de l’autorisation écrite. Ce registre sert deux objectifs distincts. D’abord, il évite les confusions internes lorsqu’une équipe SOC détecte un scan Nmap dans ses journaux et doit déterminer en quelques minutes s’il s’agit d’un audit légitime ou d’une activité suspecte à traiter en urgence. Ensuite, il constitue une pièce de dossier utile en cas de contrôle réglementaire, notamment dans les secteurs soumis à NIS2 ou à des exigences sectorielles comme celles de la santé ou de la finance.

## Nmap face aux autres scanners de réseau

Nmap n’est pas le seul outil de reconnaissance réseau disponible, même s’il reste le plus polyvalent. Le choix entre ces outils dépend rarement d’une préférence technique isolée : il dépend surtout du livrable attendu. Un rapport d’audit destiné à un comité de direction ou à un organisme de certification a besoin d’une sortie structurée avec scoring de risque, ce que fournissent nativement des scanners de vulnérabilités comme OpenVAS. Un diagnostic rapide en pleine intervention sur incident, en revanche, tolère mal la lourdeur de ces plateformes et privilégie la vitesse d’exécution de Nmap. Voici comment il se positionne face à trois alternatives fréquemment citées en 2026.

| Outil | Type | Point fort | Limite | 
|---|---|---|---|
| Nmap | Scanner de ports + NSE | Gratuit, scriptable, référence du secteur | Ligne de commande, courbe d’apprentissage réelle | 
| Zenmap | Interface graphique pour Nmap | Visualisation topologique, profils prédéfinis | Ne couvre pas toutes les options CLI avancées | 
| OpenVAS / Greenbone | Scanner de vulnérabilités | Base CVE intégrée, rapports de conformité | Plus lourd à déployer, orienté audit planifié | 
| Angry IP Scanner | Scanner d’hôtes basique | Ultra simple, multiplateforme | Pas de détection de service ni de scripting | 

En pratique, beaucoup d’équipes sécurité combinent les deux approches : Nmap pour la découverte rapide et l’automatisation légère, Trivy pour le scan de conteneurs et un scanner de vulnérabilités dédié comme OpenVAS pour les audits de conformité approfondis nécessitant un rapport formel.

## 5 pièges fréquents à éviter avec Nmap

La majorité des incidents liés à Nmap ne viennent pas d’une mauvaise intention, mais d’une commande mal calibrée lancée sans réfléchir à son impact réseau. Voici les erreurs qui reviennent le plus souvent, aussi bien chez les débutants que chez des administrateurs expérimentés pressés par le temps.

- **Lancer un scan SYN sans sudo.** Nmap bascule silencieusement sur un scan Connect (-sT) moins discret et plus lent, sans toujours prévenir clairement l’utilisateur débutant.
- **Scanner une plage IP trop large sans autorisation.** Un`/16` lancé par erreur au lieu d’un`/24` peut toucher des milliers de machines hors périmètre autorisé.
- **Confondre port filtré et port fermé.** Un port « filtered » signifie qu’un pare-feu bloque la réponse, pas que le service est absent : le traiter comme fermé fausse tout l’audit.
- **Oublier -Pn derrière un pare-feu strict.** Si les échos ICMP sont bloqués, Nmap déclare l’hôte injoignable alors qu’il répond parfaitement sur ses ports TCP.
- **Lancer –script vuln en pleine production sans fenêtre de maintenance.** Certains scripts intrusifs peuvent perturber des services fragiles ou déclencher des blocages automatiques côté pare-feu applicatif.

## Dépannage : 8 problèmes courants et leurs solutions

Même avec une installation propre, Nmap produit régulièrement des résultats surprenants pour qui découvre l’outil. La plupart de ces situations s’expliquent par le comportement des pare-feux modernes ou par une mauvaise combinaison d’options, et se résolvent en une seule commande une fois le diagnostic posé.

