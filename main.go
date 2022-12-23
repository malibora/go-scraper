package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("ananas.rs"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.OnHTML("button", func(e *colly.HTMLElement) {
		if e.Text == "Dodaj u korpu" {
			_, disabled := e.DOM.Attr("disabled")

			if !disabled {
				sendEmail("")
			} else {
				fmt.Println("Tastatura je nedostupna")
			}

		}

	})
	c.Visit("https://ananas.rs/proizvod/microsoft-mis-tastatura-sculpt-ergonomic-desktop-crni/257216")

}

func sendEmail(body string) {

	from := mail.NewEmail("Boris Popov", "boris.popov@htecgroup.com")
	subject := "tastatura je dostupna"
	to := mail.NewEmail("Boris Popov", "malibora@gmail.com")
	plainTextContent := "https://ananas.rs/proizvod/microsoft-mis-tastatura-sculpt-ergonomic-desktop-crni/257216"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//<button disabled="" width="100%" class="sc-1vt7mai-0 sc-1pmeklq-0 gZLqBM irGVHj">Dodaj u korpu</button>
