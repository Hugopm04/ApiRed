package modelo

import (
	"time"
)

type Clima struct {
	coordenada            Coordenada
	registroMeteorologico map[time.Time]Meteorologia
}

func NuevoClima(coord Coordenada) Clima {
	return Clima{coordenada: coord, registroMeteorologico: make(map[time.Time]Meteorologia)}
}
