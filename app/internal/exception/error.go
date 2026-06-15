package exception

import "errors"

var ErrFileNotFound error = errors.New("Arquivo não encontrado")
