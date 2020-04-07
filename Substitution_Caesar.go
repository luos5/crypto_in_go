package main

//Caesar only for a-m
func Caesar_encode(file []byte,key int) []byte{
	var output = make([]byte,len(file))
	for i,v := range file{
		if v>=96{
			output[i] = byte((int(v)+key-96)%26+96)
		}else{
			output[i] = byte((int(v)+key-64)%26+64)
		}	
	}
	return output
}
//Caesar only for a-m
func Caesar_decode(file []byte,key int) []byte{
	var output = make([]byte,len(file))
	for i,v := range file{
		if v>=96{
			output[i] = byte((int(v)-key-96)%26+96)
		}else{
			output[i] = byte((int(v)-key-64)%26+64)
		}	
	}
	return output
}
//Caesar for ascii
func Caesar_encode_ascii(file []byte,key int) []byte{
	var output = make([]byte,len(file))
	for i,v := range file{
		output[i] = byte((int(v)+key)%256)
	}
	return output
}
//Caesar for ascii
func Caesar_decode_ascii(file []byte,key int) []byte{
	var output = make([]byte,len(file))
	for i,v := range file{
		output[i] = byte((int(v)-key)%256)
	}
	return output
}
