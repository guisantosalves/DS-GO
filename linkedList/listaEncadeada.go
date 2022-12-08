package linkedList

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type ListaEncadeada struct {
	inicio *No
	fim    *No
}

func (lista *ListaEncadeada) Add(v int) {
	var novoNo = No{valor: v, proximo: nil}
	if lista.fim == nil {
		lista.inicio = &novoNo
		lista.fim = &novoNo
	} else {
		// colocando o prox como novoNo
		// atualizando o lista.fim ser o novoNo adicionado
		lista.fim.proximo = &novoNo
		lista.fim = &novoNo
	}
}

func (lista *ListaEncadeada) getNo(indice int) *No {
	var no = lista.inicio
	for i := 0; no != nil; i++ {
		if i == indice {
			return no
		}
		no = no.proximo
	}
	return nil
}

func (lista *ListaEncadeada) Get(indice int) int {
	// return value from getNo
	return lista.getNo(indice).valor
}

func (lista *ListaEncadeada) Remove(indice int) {
	if lista.inicio == nil {
		panic("Lista vazia!!")
	}

	if indice == 0 {
		lista.inicio = lista.inicio.proximo
		if lista.inicio == nil {
			lista.fim = nil
		}
	} else {
		// quero retirar x
		// faço o anterior a x apontar para o proxímo de x
		var no = lista.getNo(indice - 1)
		if no == nil || no.proximo == nil {
			panic("indice nao existe")
		}

		no.proximo = no.proximo.proximo
	}

}

func Main() {
	var lista = ListaEncadeada{}

	lista.Add(10)
	lista.Add(20)
	lista.Add(50)

	fmt.Println(lista.Get(2))
	lista.Remove(2)
	fmt.Println(lista.getNo(2))

}
