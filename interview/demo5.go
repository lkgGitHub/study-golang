package interview

urls:=[]string{}

//初始化urls

theadChan:=make(chan struct{},100)
wg:=sync.Waitgroup{}

theadChan<-struct{}
wg.add(1)
for _,url:=range urls{


go func(url){
defer func(){
wg.done()
<-theadChan

}

download(url)

} (url)
}

wg.wait()
