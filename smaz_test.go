package smaz_test

func TestSmaz(t *testing.T) {
	expect := []byte{254, 84, 76, 56, 172, 62, 173, 152, 62, 195, 70}
	dec, err := Decompress(expect)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	if sdec := string(dec); sdec != `This is a small string` {
		fmt.Println("unexpected decompressed result", sdec)
		return
	}
	out, err := Compress([]byte(`This is a small string`))
	if err != nil {
		fmt.Println("error", err)
		return
	}
	if !bytes.Equal(out, expect) {
		fmt.Println("unexpected compress result", out)
		return
	}
	fmt.Println("insufficient test passed")

}
