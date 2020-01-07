package main

import (
	"google.golang.org/grpc"
	"log"
		"time"
	pb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewClientClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var payment = []*pb.ChannelPayment{}
	payment = append(payment, &pb.ChannelPayment{ChannelId: 4, Amount: 3})
	var payments = pb.ChannelPayments{ChannelPayments: payment}

	r, err := client.AgreementRequest(ctx, &pb.AgreeRequestsMessage{PaymentNumber: 3, ChannelPayments: &payments})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("HALO", r.GetResult())
}
