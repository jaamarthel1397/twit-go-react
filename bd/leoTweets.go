package bd

import (
	"context"
	"log"
	"time"

	"github.com/jaamarthel1397/twit-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options" //permite definir opciones para filtrar
)

/* LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //el limite de lo que vamos a traer
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //la forma en que van a ser ordenados
	//value -1 significa que seran de forma descendente, ultimos primero
	opciones.SetSkip((pagina - 1) * 20) //multiplicado por el limite
	//queremos la pagina 1, donde no se va a saltar nada
	//para la segunda pagina nos saltamos los primeros 20, sucesivamente

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
	//todo crea un contexto vacio
}
