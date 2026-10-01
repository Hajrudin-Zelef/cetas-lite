---
id: collect-261001-general-networking/general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026-4
title: "repérez le nom de votre interface, par exemple eth0 ou ens18"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026.md
source_anchor: ""
source_lines: [276, 350]
sha256: afec6092645d8037bd957a92684bc9356bbac197da5153faab42af5a56dc1c15
---

# repérez le nom de votre interface, par exemple eth0 ou ens18

Concrètement, cela implique plusieurs précautions pour une entreprise française. Informez vos salariés, via la charte informatique ou le règlement intérieur, qu’une supervision réseau est en place à des fins de sécurité, conformément au principe de transparence du RGPD. Limitez la durée de conservation des logs à ce qui est strictement nécessaire à la détection et à l’investigation d’incident, en général quelques mois, et documentez cette durée dans votre registre de traitement. Restreignez l’accès aux journaux Suricata et à votre SIEM aux seules personnes habilitées, avec une journalisation des accès lorsque c’est possible. Enfin, si votre organisation dépend du secteur public ou traite des données de santé, vérifiez que votre architecture de supervision réseau s’inscrit bien dans le cadre plus large de vos obligations SecNumCloud ou NIS2, qui imposent déjà une capacité de détection documentée.

## Automatiser le déploiement avec Ansible

Pour une organisation qui gère plusieurs sondes Suricata sur différents sites ou segments réseau, installer et configurer chaque serveur manuellement devient vite ingérable. Un playbook Ansible simple permet de standardiser l’installation, le déploiement des règles locales et le redémarrage du service sur l’ensemble du parc en une seule commande.

```
---
- name: Déployer Suricata 8.0.6 sur les sondes réseau
  hosts: sondes_suricata
  become: true
  tasks:
    - name: Installer Suricata depuis les backports
      apt:
        name: suricata
        state: present
        default_release: bookworm-backports
        update_cache: true
    - name: Copier la configuration HOME_NET personnalisée
      template:
        src: templates/suricata.yaml.j2
        dest: /etc/suricata/suricata.yaml
      notify: redémarrer suricata
    - name: Copier les règles locales
      copy:
        src: files/local.rules
        dest: /etc/suricata/rules/local.rules
      notify: redémarrer suricata
    - name: Mettre à jour les règles publiques
      command: suricata-update
  handlers:
    - name: redémarrer suricata
      systemd:
        name: suricata
        state: restarted
```
Ce playbook peut être exécuté via une intégration continue (GitLab CI, GitHub Actions) à chaque modification du dépôt Git contenant vos règles locales, ce qui garantit que toutes vos sondes appliquent la même politique de détection en quelques minutes, sans intervention manuelle serveur par serveur.

## Erreurs fréquentes à éviter lors du déploiement

Plusieurs pièges reviennent systématiquement lors des premiers déploiements de Suricata, y compris chez des équipes techniques expérimentées.

- **Oublier de définir HOME_NET correctement.** Une valeur par défaut ou trop large fait que de nombreuses règles orientées “trafic sortant” ou “trafic interne” ne se déclenchent jamais, donnant une fausse impression de calme.
- **Laisser les fonctions d’offload réseau activées.** TSO, GSO et GRO peuvent réassembler ou fragmenter les paquets avant que Suricata ne les voie, provoquant des détections manquées difficiles à diagnostiquer.
- **Activer trop de règles d’un coup sans phase de tuning.** Un ruleset complet en mode IPS sans période d’observation en IDS génère un volume de faux positifs qui décourage rapidement les équipes et peut bloquer du trafic métier légitime.
- **Ne jamais mettre à jour les règles.** Un cron mal configuré ou oublié laisse le moteur fonctionner avec des signatures vieilles de plusieurs mois, inefficaces contre les menaces récentes.
- **Sous-dimensionner le CPU par rapport au débit réel.** Un lien à 1 Gbit/s saturé génère des kernel_drops massifs sur un serveur à 2 cœurs, rendant la détection partielle et peu fiable.
- **Confondre mode IDS et mode IPS dès le départ.** Démarrer directement en blocage actif sans phase de validation peut interrompre des flux métier critiques, notamment lors de faux positifs sur du trafic chiffré ou des API internes.
- **Ignorer la rotation des logs.** Le fichier eve.json peut grossir de plusieurs gigaoctets par jour sur un réseau chargé ; sans logrotate configuré, le disque se remplit et le service peut planter.

## Dépannage : 8 problèmes courants et leurs solutions

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Le service ne démarre pas | Erreur de syntaxe dans suricata.yaml | Lancer `suricata -T -c /etc/suricata/suricata.yaml -v` pour localiser l’erreur | 
| Aucune alerte générée | HOME_NET mal configuré ou interface incorrecte | Vérifier la variable HOME_NET et confirmer l’interface avec `ip a` | 
| capture.kernel_drops élevé | CPU insuffisant ou offload réseau actif | Augmenter les threads af-packet et désactiver TSO/GSO/GRO | 
| eve.json ne grossit pas | Module eve-log désactivé dans la config | Vérifier `enabled: yes` sous la section eve-log | 
| suricata-update échoue | Pas d’accès Internet sortant ou proxy non configuré | Tester la connectivité et définir la variable HTTP_PROXY si nécessaire | 
| Trop de faux positifs | Règles génériques activées sans contexte réseau | Désactiver les signatures non pertinentes via un fichier disable.conf | 
| IPS bloque du trafic légitime | Règle en mode “drop” mal calibrée pour l’environnement | Repasser la règle en mode “alert” le temps de l’analyser | 
| Service qui plante après quelques heures | Disque saturé par les logs non tournés | Configurer logrotate sur /var/log/suricata/ et limiter la rétention | 

## Astuces avancées pour aller plus loin

Une fois le déploiement de base stabilisé, plusieurs réglages permettent d’affiner la détection et de réduire la charge opérationnelle. Activez d’abord l’inspection TLS avec extraction de certificats pour repérer les certificats auto-signés suspects ou les domaines nouvellement enregistrés utilisés dans des campagnes de phishing, une technique de plus en plus courante en 2026 face au chiffrement quasi systématique du trafic web.

Exploitez également le module `file-store` de Suricata pour extraire automatiquement les fichiers transitant sur le réseau (exécutables, documents Office, archives) et les soumettre à un antivirus ou à un bac à sable d’analyse. Cette fonctionnalité, combinée à des règles de détection de type YARA importées via `suricata-update`, transforme Suricata en un outil de détection de malware réseau à part entière, et pas seulement en signature matching classique.

Pour les environnements à fort débit, envisagez PF_RING ou DPDK comme méthode de capture alternative à af-packet standard : ces bibliothèques contournent une partie de la pile réseau du noyau Linux et réduisent significativement les pertes de paquets sur des liens à plusieurs Gbit/s. Enfin, si votre organisation dispose de plusieurs sites, pensez à uniformiser vos règles de détection via un dépôt Git central, ce qui facilite l’audit de conformité NIS2 et garantit une politique de détection cohérente sur l’ensemble du parc.

## Projet complet : IDS/IPS Suricata avec supervision Wazuh

Voici l’architecture complète d’un déploiement fonctionnel, du capteur réseau à l’alerte visible dans un tableau de bord, telle qu’elle peut être mise en place sur un serveur unique pour une PME.

