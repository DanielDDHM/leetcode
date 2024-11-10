package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Nó sentinela para facilitar a construção da lista
	current := dummy     // Ponteiro para o nó atual da lista resultante
	carry := 0           // Variável para armazenar o transporte (carry) durante a adição

	for l1 != nil || l2 != nil || carry != 0 {
		// Calcula a soma dos valores e do transporte
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10                        // Atualiza o transporte
		current.Next = &ListNode{Val: sum % 10} // Cria um novo nó com o valor da unidade da soma
		current = current.Next                  // Avança para o próximo nó
	}

	return dummy.Next // Retorna o próximo nó após o nó sentinela, que é o início da lista resultante
}
