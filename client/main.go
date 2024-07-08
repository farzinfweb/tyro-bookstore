package main

import (
	"bookstore/protos"
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type SearchReqParams struct {
	SearchTerm string `query:"s"`
	Page       uint32 `query:"page"`
	PerPage    uint32 `query:"per_page"`
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:7080", opts...)
	if err != nil {
		return
	}
	defer conn.Close()

	client := protos.NewBookstoreClient(conn)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/buy", func(c echo.Context) error {
		req := &protos.BuyReq{
			BookId:   "6686b707291845d30bf31a99",
			Quantity: 2,
		}
		buyResp, err := client.Buy(context.TODO(), req)
		if err != nil {
			return err
		}
		return c.JSON(200, buyResp)
	})
	e.GET("/books", func(c echo.Context) error {
		reqParams := new(SearchReqParams)
		err := c.Bind(reqParams)
		if err != nil {
			return err
		}
		searchResp, err := client.Search(context.TODO(), &protos.SearchReq{
			SearchTerm: reqParams.SearchTerm,
			Page:       reqParams.Page,
			PerPage:    reqParams.PerPage,
		})
		if err != nil {
			return err
		}
		return c.JSON(200, searchResp)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
