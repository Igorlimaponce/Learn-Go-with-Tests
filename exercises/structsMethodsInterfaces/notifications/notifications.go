package main

import "fmt"

type NotificacaoEmail struct {
	DestinatarioEmail string
	Assunto           string
	Corpo             string
}

type NotificacaoSMS struct {
	NumeroTelefone int
	Mensagem       string
}

type NotificacaoPush struct {
	DeviceID int
	Titulo   string
	Mensagem string
}

type Notificator interface {
	Send() error
}

// Aqui eu defini que a interface Notificator tem um metodo Send()

// Agora vou implementar esse Mesmo Metodo Send() em todas as structs que vao usar ele

func (n NotificacaoEmail) Send() error {
	fmt.Printf("(Email para: %s) ", n.DestinatarioEmail)
	return nil
}

func (n NotificacaoSMS) Send() error {
	fmt.Printf("(SMS para: %d) ", n.NumeroTelefone)
	return nil
}

func (n NotificacaoPush) Send() error {
	fmt.Printf("(Push para Device: %d) ", n.DeviceID)
	return nil
}

// Essa funcao acabe se tornando polimorfica pois ela
// Ela não se importa se está recebendo um Email, SMS ou Push; ela só se importa em receber algo que satisfaça o contrato Notificator
func ConcluirCompra(n Notificator) error {
	fmt.Printf("Enviando notificação... ") // Adicionei um print para vermos
	err := n.Send()
	// Esse seria o caso de uso, eu poderia chamar ConcluirCompra utilizando qualquer uma das 3 structs que criei
	if err != nil {
		fmt.Println("Falhou!")
		return err
	}
	fmt.Println("Enviada com sucesso!")
	return nil
}

// Agora, vamos refinar a função que "chama"
func main() {
	// 1. Criamos instâncias REAIS das suas structs com dados
	email := NotificacaoEmail{
		DestinatarioEmail: "cliente@email.com",
		Assunto:           "Seu pedido foi confirmado!",
	}

	sms := NotificacaoSMS{
		NumeroTelefone: 999999999, // Veja a dica abaixo sobre isso!
		Mensagem:       "Seu pedido 123 foi confirmado e está a caminho.",
	}

	push := NotificacaoPush{
		DeviceID: 777,
		Titulo:   "Pedido Confirmado!",
		Mensagem: "Toque para ver os detalhes do seu pedido.",
	}

	// slice da INTERFACE
	// Este slice pode conter QUALQUER struct que implemente "Send()"
	notificacoes := []Notificator{email, sms, push}

	fmt.Println("Iniciando processo de conclusão de compra...")
	for _, n := range notificacoes {
		// Não importa o tipo, apenas chamamos ConcluirCompra
		// O Go sabe qual método Send() chamar (o do email, do sms ou do push)
		err := ConcluirCompra(n)
		if err != nil {
			fmt.Println("Erro ao processar notificação:", err)
		}
	}
}
