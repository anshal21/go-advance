package usage

func FileWriter(fileName string) {
	f, _ := os.Create(fileName) 
    defer f.Close()
   
	// Do the processing here
} 



