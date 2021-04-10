package interview

import "fmt"

type Student struct {
	Name string
}

func main() {

	st := &Student{
		Name: "aa",
	}

	Kill(st)

	fmt.Println(st)

}

func Kill(student *Student) {
	student = nil
}




urls:=[]string{}

//初始化urls



for _,url:=range urls{
download(url)
}

type Canvas struct{

}


func(c *Canvas)Draw(type){

if type == tiger{
...
}

if type == lion{
...
}

}

type Tiger struct{

}

type Lion struct{

}
