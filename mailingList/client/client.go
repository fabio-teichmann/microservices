package main

import (
	"context"
	"log"
	pb "mailingList/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func logRespone(res *pb.EmailResponse, err error) {
	if err != nil {
		log.Fatalf("    error: %v", err)
	}

	if res.EmailEntry == nil {
		log.Printf("    email not found")
	} else {
		log.Printf("    response: %v", res.EmailEntry)
	}
}

func createEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("create email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // timeout after 1 second
	defer cancel()                                                        // frees ressources after timeout

	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	logRespone(res, err)

	return res.EmailEntry
}

func getEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("get email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // timeout after 1 second
	defer cancel()                                                        // frees ressources after timeout

	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	logRespone(res, err)

	return res.EmailEntry
}

func getEmailBatch(client pb.MailingListServiceClient, count, page int) {
	log.Printf("get email batch")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // timeout after 1 second
	defer cancel()                                                        // frees ressources after timeout

	res, err := client.GetBatchEmail(ctx, &pb.GetEmailBatchRequest{Page: int32(page), Count: int32(count)})
	if err != nil {
		log.Fatalf("    error: %v", err)
	}
	log.Println("response:")
	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("    item [%v of %v]: %s", i+1, len(res.EmailEntries), res.EmailEntries[i])
	}
}

func updateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) *pb.EmailEntry {
	log.Printf("update email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // timeout after 1 second
	defer cancel()                                                        // frees ressources after timeout

	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	logRespone(res, err)

	return res.EmailEntry
}

func deleteEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("delete email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // timeout after 1 second
	defer cancel()                                                        // frees ressources after timeout

	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: addr})
	logRespone(res, err)

	return res.EmailEntry
}

var args struct {
	GrpcAddr string `arg:"env:MAILINGLIST_GRPC_ADDR"`
}

func main() {
	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}

	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMailingListServiceClient(conn)
	newEmail := createEmail(client, "444@999.999")
	newEmail.ConfirmedAt = 10000
	updateEmail(client, *newEmail)
	deleteEmail(client, newEmail.Email)

	getEmailBatch(client, 5, 1)
}
