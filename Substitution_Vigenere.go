package main
//vigenere polyalphabetic
func Vigenere_encode(file []byte,key []byte) []byte{
	output := make([]byte,len(file))
	len_key := len(key)
	for i,v := range file{
		if v>=97{
			output[i] = byte((int(v+key[i%len_key])-96)%26+96)
		}else{
			output[i] = byte((int(v+key[i%len_key])-64)%26+64)
		}
	}
	return output
}
//vigenere polyalphabetic
func Vigenere_decode(file []byte,key []byte) []byte{
	output := make([]byte,len(file))
	len_key := len(key)
	for i,v := range file{
		if v>=97{
			output[i] = byte((int(v-key[i%len_key])-96)%26+96)
		}else{
			output[i] = byte((int(v-key[i%len_key])-64)%26+64)
		}
	}
	return output
}

