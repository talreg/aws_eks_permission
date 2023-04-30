package utils

import "os"

type FileTokenRetriever struct {
	TokenFilePath string
}

func (ftr *FileTokenRetriever) GetIdentityToken() ([]byte, error) {
	tokenBytes, err := os.ReadFile(ftr.TokenFilePath)
	if err != nil {
		return nil, err
	}
	return tokenBytes, nil
}
