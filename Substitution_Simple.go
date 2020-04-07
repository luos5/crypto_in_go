package main

//simple substitution ciphier 
func SSC_encode(file []byte,key map[byte]byte) []byte{
	output := make([]byte,len(file))
	for i,v := range file{
		value,exist:=key[v]
		if exist==true{
			output[i] = value
		}
	}
	return output
}
//simple substitution ciphier 
func SSC_decode(file []byte,reverse_key map[byte]byte) []byte{
	output := make([]byte,len(file))
	for i,v := range file{
		value,exist:=reverse_key[v]
		if exist==true{
			output[i] = value
		}
	}
	return output
}
//reverse the key table/map
func reverse_key_map(key map[byte]byte)(reverse_key map[byte]byte){
	reverse_key = make(map[byte]byte,len(key))
	for i,v := range key{
		reverse_key[v] = i
	}
	return reverse_key
}
//read and build key map
func key_map(plain[]byte,cipher[]byte)(key map[byte]byte){
	key = make(map[byte]byte,len(plain))
	for i,v := range plain{
		key[v] = cipher[i]
	}
	return key
}
