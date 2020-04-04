package main
import(
	"fmt"	
	"os"
	"io"
//	"io/ioutil"
	"bufio"
	"strings"
)
var (
	sbox = make([]int,256)
)
func main(){
	inputReader := bufio.NewReader(os.Stdin)
	//choose encode/decode
	fmt.Printf("RC4...\nPlease select the fuction:\nA.encode\tB.decode\n")
	chose, chose_err := inputReader.ReadString('\n')
	chose = strings.Replace(chose,"\n","",-1)
	//possible reading error
	if chose_err != nil{
		fmt.Println("There were errors reading, exiting program.")
		return
	}	

	//get file name
	fmt.Println("Please enter the path of the file:")
	file_name, file_name_err := inputReader.ReadString('\n')
	file_name = strings.Replace(file_name,"\n","",-1)
	if file_name_err != nil{
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	//get key
	fmt.Println("Please enter the key(<=256 words):")
	key, key_err := inputReader.ReadString('\n')
	key = strings.Replace(key,"\n","",-1)
	var key_byte []byte = []byte(key)
	if key_err != nil{
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	//read file
	file_dir, file_dir_err := os.Open(file_name)
	if file_dir_err != nil{
		fmt.Println("An error occurred on opening the inputfile\n" +"Does the file exist?\n" +"Have you got acces to it?\n")
        	return
	}
	defer file_dir.Close()
	//inputReader = bufio.NewReader(file_dir)
	file_slice := make([]byte,0)
	buf := make([]byte,1)
	for {
		n, err := file_dir.Read(buf)
		if err != nil && err != io.EOF{
			panic(err)
		}
		if 0 == n {
			break
		}
		file_slice = append(file_slice, buf[:n]...)
	}
	//set sbox
	for i:=0;i<=255;i++{
		sbox[i]=i
	}
	//mix_sbox(key)
	mix_sbox(key_byte)
	
	switch chose{
	case "A":	encode(file_slice,file_name)	
	case "B":	decode(file_slice,file_name)
	default:	fmt.Println("Illegal input!")
	}
}
func encode(file []byte,name string){
	len_file := len(file)
	tmp := make([]byte,len_file)
	a := 0
	b := 0
	for i:=0;i<len_file;i++{
		a = (a+1)%256
		b = (b+sbox[a])%256
		sbox[a],sbox[b] = sbox[b],sbox[a]
		tmp[i] = byte(int(file[i])^((sbox[a]+sbox[b])%256))
	}
	//output,write in a new file
	output_file,output_file_err := os.OpenFile(name+".encode",os.O_WRONLY|os.O_CREATE,0666)
	if output_file_err != nil{
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer output_file.Close()
	output_file.Write(tmp)
	return
}
func decode(file []byte,name string){
	len_file := len(file)
	tmp := make([]byte,len_file)
	a := 0
	b := 0
	for i:=0;i<len_file;i++{
		a = (a+1)%256
		b = (b+sbox[a])%256
		sbox[a],sbox[b] = sbox[b],sbox[a]
		tmp[i] = byte(int(file[i])^((sbox[a]+sbox[b])%256))
	}
	//output,write in a new file
	output_file,output_file_err := os.OpenFile(name+".decode",os.O_WRONLY|os.O_CREATE,0666)
	if output_file_err != nil{
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer output_file.Close()
	output_file.Write(tmp)
	return
}
//func mix_sbox(key string){
//	len_key := len(key)
func mix_sbox(key_byte []byte){
	len_key := len(key_byte)
	j := 0
	for i:=0;i<=255;i++{
//		j = (j+sbox[i]+int(key[i%len_key]))%256
		j = (j+sbox[i]+int(key_byte[i%len_key]))%256
		sbox[i],sbox[j] = sbox[j],sbox[i]
	}
	return
}
