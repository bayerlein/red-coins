//Pacote principal da aplicação, onde é subido o servidor e as rotas são mapeadas
package main

import (
	"log"
	"net/http"

	"github.com/bayerlein/red-coins/controllers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

//Metodo que configura e expoem as rotas da aplicação
func Routes() *chi.Mux {

	//Cria o objeto que contem as configurações
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	//Cria uma controller de bitcoins, onde podemos acessar as rotas especificas do bitcoin
	bitCoinController := controllers.NewBitCoinController()

	//Cria uma controller de usuarios, onde podemos acessar as rotas especificas de usuarios
	userController := controllers.NewUserController()

	//Define o mapeamento das rotas
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/bitcoin", bitCoinController.Routes())
		r.Mount("/api/user", userController.Routes())
	})

	return router
}

//Metodo principal da aplicação, onde o servidor é efetivamente ativado
func main() {
	router := Routes()

	//A função faz o log das rotas acessadas
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)

		return nil
	}

	//Faz tratamento de erro da função 'walkFunc'
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Erro - %s\n", err.Error())
	}

	//Sobe o servidor na porta definida
	log.Fatal(http.ListenAndServe(":8080", router))
}
