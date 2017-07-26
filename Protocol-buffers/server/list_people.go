package main

p := pb.Person{
  Id: 1234,
  Name : "Jihar Al Gifari",
  Email : "jihar.akakom14@gmail.com",
  Phones: []*pb.Person_PhoneNumber{
    {Number: "555-4321", Type: pb.Person_HOME},
  },
}
