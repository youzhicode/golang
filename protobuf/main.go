package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	trippb "github.com/youzhicode/golang/protobuf/proto/gen/go"
	"github.com/youzhicode/golang/protobuf/tripservice"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Printf("failed to listen %v\n", err)
	}

	s := grpc.NewServer()

	trippb.RegisterTripServiceServer(s, tripservice.TestService)

	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8888",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("Cannot start grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8887", mux)
	if err != nil {
		log.Fatalf("Listen error %v\n", err)
	}
}

// func main() {
// 	trip := trippb.Trip{
// 		Start:       "abc",
// 		End:         "def",
// 		DurationSec: 3600,
// 		FeeCent:     10000,
// 		StartPos: &trippb.Location{
// 			Latitude:  30,
// 			Longitude: 50,
// 		},
// 		EndPos: &trippb.Location{
// 			Latitude:  40,
// 			Longitude: 70,
// 		},
// 		PathLocations: []*trippb.Location{
// 			{
// 				Latitude:  31,
// 				Longitude: 51,
// 			},
// 			{
// 				Latitude:  35,
// 				Longitude: 58,
// 			},
// 		},
// 		Status: trippb.TripStatus_IN_PROGRESS,
// 	}

// 	fmt.Println(&trip)
// 	b, err := proto.Marshal(&trip)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%X\n", b)

// 	var trip2 trippb.Trip
// 	err = proto.Unmarshal(b, &trip2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(&trip2)

// 	b, err = json.Marshal(&trip2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", b)
// }
