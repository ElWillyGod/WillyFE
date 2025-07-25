package core

import (
	"strings"
)

/*
* implementación simplificada de un buscador de archivos utilizando el B+ Tree.
* solo dos tipos de búsqueda:
* 1. búsqueda exacta: se busca el nombre del archivo completo.
* 2. búsqueda parcial: se busca el nombre del archivo que contenga una subcadena.
 */

type Buscador struct {
	tree *BPlusTree
}

func NuevoBuscador(tree *BPlusTree) *Buscador {
	return &Buscador{
		tree: tree,
	}
}

func (b *Buscador) BuscarExacto(nombreArchivo string) ([]string, string) {

	clave := strings.ToLower(nombreArchivo)
	nodo := b.tree.EncontrarHoja(clave)

	rutas := b.tree.buscarEnHoja(nodo, clave)
	if rutas != nil {
		return rutas, clave
	}

	return nil, ""
}

func (b *Buscador) BuscarParcial(subcadena string) []string {
	var rutas []string
	subcadena = strings.ToLower(subcadena)

	nodoActual := b.tree.EncontrarPrimeraHoja()
	for nodoActual != nil {
		for _, entrada := range nodoActual.Entradas {
			if strings.Contains(entrada.Clave, subcadena) {
				rutas = append(rutas, entrada.Rutas...)
			}
		}
		nodoActual = nodoActual.Siguiente
	}

	return rutas
}
