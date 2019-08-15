package main

import (
	"First-go/src/MgoDemos/sysinit"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name string
	Age  int
}

//初始化连接

var collection *mgo.Collection

func init() {
	sysinit.InitMongo()
	session := sysinit.MongoSession.Copy()
	if session == nil {
		log.Println("获取session失败")
	}
	collection = session.DB("").C("student")
}

//增加一条数据
func InsertOne(person Person) {
	err := collection.Insert(person)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("insert  succeed", person)
	}
}

//查找一条数据
func FindOne(SQL bson.M) Person {
	result := Person{}
	collection.Find(SQL).One(&result)
	return result
}

func Find(SQL bson.M) []Person {
	result := []Person{}
	err := collection.Find(SQL).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}