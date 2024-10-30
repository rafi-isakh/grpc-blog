package main

import (
	pb "github.com/rafi-isakh/grpc-blog/proto"
)

type BlogItem struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
	AuthorId string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id: 		data.ID, 
		AuthorId: 	data.AuthorId,
		Title: 		data.Title,
		Content: 	data.Content,
	}
}