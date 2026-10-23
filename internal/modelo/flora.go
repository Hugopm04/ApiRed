package modelo

type Flora struct {
	flores map[string]Flor
}

func NuevaFlora() *Flora { return &Flora{flores: make(map[string]Flor)} }
