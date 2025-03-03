# Backend de Réservation d'Hôtel

Ce projet est un backend API RESTful développé en Go pour gérer les réservations d'hôtels.

## Fonctionnalités

* Gestion des chambres (création, lecture, mise à jour, suppression)
* Gestion des réservations (création, lecture, annulation)
* Authentification des utilisateurs (clients et administrateurs)
* Recherche d'hôtels par ville, dates, etc.
* Gestion des disponibilités des chambres.

## Prérequis

* Go (version gogo1.24.0 )
* mysql
* Git

## Installation

1.  Clonez le dépôt :

    ```bash
    git clone [https://github.com/valraiky/HotelKiasyBackend](https://github.com/valraiky/HotelKiasyBackend)
    ```

2.  Accédez au répertoire du projet :

    ```bash
    mkdir golang-crud && cd golang-crud
    ```

3.  Installer les packages nécessaires :

    ```bash
    go get github.com/gin-gonic/gin
    go get gorm.io/driver/mysql
    go get gorm.io/gorm
    ```

4.  Configurez la base de données :

    * Créez une base de données Mysql.
    * Modifiez le fichier `config.yaml` avec vos informations de connexion à la base de données.

5.  Exécutez les migrations de la base de données :

    ```bash
    go run cmd/migrate/*.go up ou faite avec automigrate
    ```

## Utilisation

1.  Démarrez le serveur :

    ```bash
    go run cmd/api/*.go
    ```

2.  L'API sera disponible à l'adresse `http://localhost:8080`.

## Points de terminaison de l'API

Voici quelques points de terminaison importants :

* `POST /hotels` : Créer un hôtel (nécessite une authentification administrateur)
* `GET /hotels` : Récupérer la liste des hôtels
* `GET /hotels/{id}` : Récupérer un hôtel par ID
* `POST /rooms` : Créer une chambre (nécessite une authentification administrateur)
* `POST /reservations` : Créer une réservation (nécessite une authentification client)
* `GET /reservations/{id}` : Récupérer une réservation par ID

## Tests

Pour exécuter les tests, utilisez la commande suivante :

```bash
go test ./...