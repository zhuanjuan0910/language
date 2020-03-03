package main

import "net/http"

func do() error {
	res, err := http.Get("http://xxxxxxxxxx")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	// ..code...

	return nil
}

func main() {
	do()
}
