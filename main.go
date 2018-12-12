package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"github.com/willtrking/kubecon2018/cart"
	"github.com/willtrking/kubecon2018/merchandise"
	"github.com/willtrking/kubecon2018/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"github.com/willtrking/kubecon2018/account"
)

var RootCmd = &cobra.Command{
	Use: "kubecon2018",
	Run: func(cmd *cobra.Command, args []string) {},
}

var MerchandiseCmd = &cobra.Command{
	Use: "merchandise",
	Run: func(cmd *cobra.Command, args []string) {

		lis, err := net.Listen("tcp", "0.0.0.0:10000")

		if lis != nil {
			defer lis.Close()
		}

		if err != nil {
			panic(err)
		}

		server := grpc.NewServer(grpc.Creds(nil))

		proto.RegisterMerchandiseServer(server, &merchandise.MerchandiseGRPC{})

		fmt.Println("Starting merchandise server...")
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}

	},
}

var AccountCmd = &cobra.Command{
	Use: "account",
	Run: func(cmd *cobra.Command, args []string) {

		lis, err := net.Listen("tcp", "0.0.0.0:10000")

		if lis != nil {
			defer lis.Close()
		}

		if err != nil {
			panic(err)
		}

		server := grpc.NewServer(grpc.Creds(nil))

		proto.RegisterAccountsServer(server, &account.AccountGRPC{})

		fmt.Println("Starting account server...")
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}

	},
}

var CartCmd = &cobra.Command{
	Use: "cart",
	Run: func(cmd *cobra.Command, args []string) {

		lis, err := net.Listen("tcp", "0.0.0.0:10000")

		if lis != nil {
			defer lis.Close()
		}

		if err != nil {
			panic(err)
		}

		server := grpc.NewServer(grpc.Creds(nil))

		merchConn, err := grpc.Dial(os.Getenv("MERCHANDISE_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		accountConn, err := grpc.Dial(os.Getenv("ACCOUNT_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		proto.RegisterCartServer(server, &cart.CartGRPC{
			MerchandiseClient: proto.NewMerchandiseClient(merchConn),
			AccountsClient: proto.NewAccountsClient(accountConn),
		})

		fmt.Println("Starting cart server...")
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}

	},
}

var DemoCmd0 = &cobra.Command{
	Use: "demo-0",
	Run: func(cmd *cobra.Command, args []string) {

		cartConn, err := grpc.Dial(os.Getenv("CART_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		cartClient := proto.NewCartClient(cartConn)

		for {
			demoRun(cartClient, 0, 500)
		}

	},
}

var DemoCmd50 = &cobra.Command{
	Use: "demo-50",
	Run: func(cmd *cobra.Command, args []string) {

		cartConn, err := grpc.Dial(os.Getenv("CART_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		cartClient := proto.NewCartClient(cartConn)

		for {
			demoRun(cartClient, 250, 250)
		}

	},
}

var DemoCmd75 = &cobra.Command{
	Use: "demo-75",
	Run: func(cmd *cobra.Command, args []string) {

		cartConn, err := grpc.Dial(os.Getenv("CART_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		cartClient := proto.NewCartClient(cartConn)

		for {
			demoRun(cartClient, 375, 125)
		}

	},
}

var DemoCmd100 = &cobra.Command{
	Use: "demo-100",
	Run: func(cmd *cobra.Command, args []string) {

		cartConn, err := grpc.Dial(os.Getenv("CART_HOST"), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		cartClient := proto.NewCartClient(cartConn)

		for {
			demoRun(cartClient, 500, 0)
		}

	},
}

func init() {

	RootCmd.AddCommand(CartCmd)
	RootCmd.AddCommand(MerchandiseCmd)
	RootCmd.AddCommand(AccountCmd)
	RootCmd.AddCommand(DemoCmd0)
	RootCmd.AddCommand(DemoCmd50)
	RootCmd.AddCommand(DemoCmd75)
	RootCmd.AddCommand(DemoCmd100)
}

func main() {

	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
