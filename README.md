# Hands-On Software Architecture with Golang

<a href="https://www.packtpub.com/application-development/hands-software-architecture-golang?utm_source=github&utm_medium=repository&utm_campaign=9781788622592 "><img src="https://d255esdrn735hr.cloudfront.net/sites/default/files/imagecache/ppv4_main_book_cover/B09392_NEW.png" alt="Hands-On Software Architecture with Golang" height="256px" align="right"></a>

This is the code repository for [Hands-On Software Architecture with Golang](https://www.packtpub.com/application-development/hands-software-architecture-golang?utm_source=github&utm_medium=repository&utm_campaign=9781788622592), published by Packt.

**Design and architect highly scalable and robust applications using Go**

## What is this book about?
Building software requires careful planning and architectural considerations; Golang was developed with a fresh perspective on building next-generation applications on the cloud with distributed and concurrent computing concerns.

This book covers the following exciting features:
* Understand architectural paradigms and deep dive into Microservices 
* Design parallelism/concurrency patterns and learn object-oriented design patterns in Go 
* Explore API-driven systems architecture with introduction to REST and GraphQL standards 
* Build event-driven architectures and make your architectures anti-fragile 
* Engineer scalability and learn how to migrate to Go from other languages 
* Get to grips with deployment considerations with CICD pipeline, cloud deployments, and so on 
* Build an end-to-end e-commerce (travel) application backend in Go 

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1788622596) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
type Args struct {
    A, B int
}

type MuliplyService struct{}

func (t *Arith) Do(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}
```

**Following is what you need for this book:**
Hands-On Software Architecture with Golang is for software developers, architects, and CTOs looking to use Go in their software architecture to build enterprise-grade applications. Programming knowledge of Golang is assumed.

With the following software and hardware list you can run all code files present in the book (Chapter 1-12).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1-12 | CURL, Go (preferably v 0.9), Git, Cassandra, Kafka, Redis, and NSQ | Windows, Mac OS X, and Linux (Any) |

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](http://www.packtpub.com/sites/default/files/downloads/9781788622592_ColorImages.pdf).

### Related products
* Hands-On Serverless Applications with Go [[Packt]](https://www.packtpub.com/application-development/hands-serverless-applications-go?utm_source=github&utm_medium=repository&utm_campaign=9781789134612) [[Amazon]](https://www.amazon.com/dp/B07DT9DD4V)

* Mastering Go [[Packt]](https://www.packtpub.com/networking-and-servers/mastering-go?utm_source=github&utm_medium=repository&utm_campaign=9781788626545) [[Amazon]](https://www.amazon.com/dp/1788626540)

## Get to Know the Author
**Jyotiswarup Raiturkar**
has architected products ranging from high-volume e-commerce sites to core infrastructure products. Notable products include the Walmart Labs Ecommerce Fulfillment Platform, Intuit Mint, SellerApp, Goibibo, Microsoft Virtual Server, and ShiftPixy, to name a few. Nowadays, he codes in Golang, Python, and Java.

### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.


