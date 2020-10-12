package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"grpc_sample/proto/math"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type httpRequest struct {
	A int32
	B int32
}

var (
	c   math.MathClient
	ctx = context.Background()
)

func init() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s",
			os.Getenv("GRPC_HOST"),
			os.Getenv("GRPC_PORT"),
		),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting GRPC client...\n" +
		"\t- GRPC_HOST=localhost GRPC_PORT=50001 PORT=8000 go run main.go")

	c = math.NewMathClient(conn)
}

func sum(w http.ResponseWriter, r *http.Request) {
	var body httpRequest
	json.NewDecoder(r.Body).Decode(&body)
	log.Println(r.Body)

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sum, err := c.Sum(ctxTimeout, &math.Request{
		A: body.A,
		B: body.B,
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sum.String()))
}

func mul(w http.ResponseWriter, r *http.Request) {
	var body httpRequest
	params := r.URL.Query()
	log.Println("params: ", params)
	a, _ := strconv.ParseInt(params.Get("a"), 10, 32)
	body.A = int32(a)
	b, _ := strconv.ParseInt(params.Get("b"), 10, 32)
	body.B = int32(b)

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mul, err := c.Mul(ctxTimeout, &math.Request{
		A: body.A,
		B: body.B,
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mul.String()))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sum", sum).Methods("POST")
	r.HandleFunc("/mul", mul).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
