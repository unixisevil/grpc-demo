package main

import (
	"context"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/unixisevil/grpc-demo/blog/blogpb"
)

type App struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	blog     blogpb.BlogServiceClient
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("blog client started")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		errorLog.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := blogpb.NewBlogServiceClient(conn)

	app := &App{
		infoLog:  infoLog,
		errorLog: errorLog,
		blog:     c,
	}

	bid := app.createBlog()
	app.readBlog(bid)
	app.updateBlog(bid)
	app.delBlog(bid)
	app.readBlog("12345")
	app.readBlog(bid)
	for i := 0; i < 5; i++ {
		app.createBlog()
	}
	app.listBlog()
}

func (app *App) createBlog() string {
	blog := &blogpb.Blog{
		AuthorId: "jian yu",
		Title:    "grpc demo",
		Content:  "grpc is great",
	}
	res, err := app.blog.CreateBlog(context.Background(), &blogpb.CreateBlogReq{Blog: blog})
	if err != nil {
		app.errorLog.Printf("error while calling CreateBlog rpc: %v", err)
		return ""
	}
	app.infoLog.Printf("response from CreateBlog: %s\n", res.Id)
	return res.Id
}

func (app *App) readBlog(id string) {
	req := &blogpb.ReadBlogReq{Id: id}
	res, err := app.blog.ReadBlog(context.Background(), req)
	if err != nil {
		app.errorLog.Printf("error while calling ReadBlog rpc: %v", err)
		return
	}
	app.infoLog.Printf("got blog entry: %+v\n", res.Blog)
}

func (app *App) updateBlog(id string) {
	req := &blogpb.UpdateBlogReq{
		Blog: &blogpb.Blog{
			Id:       id,
			AuthorId: "new author",
			Title:    "new title",
			Content:  "new content",
		},
	}
	res, err := app.blog.UpdateBlog(context.Background(), req)
	if err != nil {
		app.errorLog.Printf("error while calling UpdateBlog rpc: %v", err)
		return
	}
	app.infoLog.Printf("got blog entry: %+v\n", res.Blog)
}

func (app *App) delBlog(id string) {
	req := &blogpb.DelBlogReq{
		Id: id,
	}
	res, err := app.blog.DelBlog(context.Background(), req)
	if err != nil {
		app.errorLog.Printf("error while calling DelBlog rpc: %v", err)
		return
	}
	app.infoLog.Printf("delete blog entry: %s successfully\n", res.Id)
}

func (app *App) listBlog() {
	stream, err := app.blog.ListBlog(context.Background(), &blogpb.ListBlogReq{})
	if err != nil {
		app.errorLog.Printf("error while calling ListBlog rpc: %v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			app.errorLog.Printf("error while recv server stream: %v", err)
			return
		}
		app.infoLog.Printf("recv one blog entry: %+v\n", res.Blog)
	}
}
