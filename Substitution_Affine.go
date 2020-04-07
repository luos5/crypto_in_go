package main
//key include a,b f(x)=ax+b
type key_type struct{
	a int
	b int
}
//Affine for ascii
func Affine_encode(file []byte,key key_type) []byte{
	output := make([]byte,len(file))
	for i,v := range file{
		output[i] = byte((int(v)*key.a+key.b)%256)
	}
	return output
}
//Affine for ascii
func Affine_decode(file []byte,key key_type) []byte{
	output := make([]byte,len(file))
	//get the modular multipulication inverse of key_a
	key_mmi := MMI(key.a,256)
	for i,v := range file{
		output[i] = byte(((int(v)-key.b)*key_mmi)%256)
	}
	return output
}
