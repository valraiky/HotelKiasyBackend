package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(r *http.Request) (string, error) {
	// Vérifier que la méthode est bien POST
	if r.Method != http.MethodPost {
		return "", fmt.Errorf("méthode non autorisée")
	}

	// Récupérer le fichier du formulaire
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}
	defer file.Close()

	// Créer le dossier d'upload s'il n'existe pas
	os.MkdirAll("./uploads", os.ModePerm)

	// Créer un fichier sur le disque
	dst, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création du fichier : %w", err)
	}
	defer dst.Close()

	// Copier le contenu du fichier uploadé vers le fichier de destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la sauvegarde du fichier : %w", err)
	}

	// Retourner le chemin du fichier uploadé
	return "./uploads/" + header.Filename, nil
}
