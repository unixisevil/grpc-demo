package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/unixisevil/grpc-demo/blog/blogpb"
)

type server struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	collection *mongo.Collection
}

type blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
	s.infoLog.Printf("recv CreateBlog req\n")
	blogData := req.GetBlog()
	data := blog{
		AuthorID: blogData.AuthorId,
		Title:    blogData.Title,
		Content:  blogData.Content,
	}
	res, err := s.collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("not expected oid"),
		)
	}
	return &blogpb.CreateBlogRes{Id: oid.Hex()}, nil
}

func (s *server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogReq) (*blogpb.ReadBlogRes, error) {
	s.infoLog.Printf("recv ReadBlog req\n")
	id := req.GetId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse blog id"),
		)
	}
	blog := &blog{}
	filter := bson.M{"_id": oid}
	if err := s.collection.FindOne(context.Background(), filter).Decode(blog); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("not find blog with req id"),
		)
	}
	return &blogpb.ReadBlogRes{
		Blog: &blogpb.Blog{
			Id:       blog.ID.Hex(),
			AuthorId: blog.AuthorID,
			Title:    blog.Title,
			Content:  blog.Content,
		},
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {
	s.infoLog.Printf("recv UpdateBlog req\n")
	data := req.GetBlog()
	id := data.GetId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse blog id"),
		)
	}
	blog := &blog{}
	filter := bson.M{"_id": oid}
	if err := s.collection.FindOne(context.Background(), filter).Decode(blog); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("not find blog with req id"),
		)
	}

	blog.AuthorID = data.AuthorId
	blog.Title = data.Title
	blog.Content = data.Content

	_, err = s.collection.ReplaceOne(context.Background(), filter, blog)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not update blog: %v", err),
		)
	}
	return &blogpb.UpdateBlogRes{
		Blog: &blogpb.Blog{
			Id:       blog.ID.Hex(),
			AuthorId: blog.AuthorID,
			Title:    blog.Title,
			Content:  blog.Content,
		},
	}, nil
}

func (s *server) DelBlog(ctx context.Context, req *blogpb.DelBlogReq) (*blogpb.DelBlogRes, error) {
	s.infoLog.Printf("recv DelBlog req\n")
	id := req.GetId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse blog id"),
		)
	}
	filter := bson.M{"_id": oid}
	res, err := s.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not del blog with req id"),
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("can not find blog with req id"),
		)
	}
	return &blogpb.DelBlogRes{Id: id}, nil
}

func (s *server) ListBlog(req *blogpb.ListBlogReq, stream blogpb.BlogService_ListBlogServer) error {
	s.infoLog.Printf("recv ListBlog req\n")
	cur, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error while call find: %v", err),
		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		blog := &blog{}
		if err := cur.Decode(blog); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("internal error while call decode: %v", err),
			)
		}
		stream.Send(&blogpb.ListBlogRes{
			Blog: &blogpb.Blog{
				Id:       blog.ID.Hex(),
				AuthorId: blog.AuthorID,
				Title:    blog.Title,
				Content:  blog.Content,
			},
		})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error while call cur.Err(): %v", err),
		)
	}
	return nil
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		errorLog.Fatalf("failed to create mongodb client: %v", err)
	}
	if err = client.Connect(context.Background()); err != nil {
		errorLog.Fatalf("failed to connect mongodb: %v", err)
	}

	infoLog.Println("blog server started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		errorLog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	rpcStub := &server{
		infoLog:    infoLog,
		errorLog:   errorLog,
		collection: client.Database("mydb").Collection("blog"),
	}
	blogpb.RegisterBlogServiceServer(s, rpcStub)

	reflection.Register(s)
	go func() {
		infoLog.Println("starting server")
		if err := s.Serve(lis); err != nil {
			errorLog.Fatalf("failed to serv: %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	infoLog.Println("stopping server")
	s.Stop()
	infoLog.Println("closing listener")
	lis.Close()
}
