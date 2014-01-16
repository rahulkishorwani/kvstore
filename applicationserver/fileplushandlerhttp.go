/* Print Env
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"net/rpc"
	"bufio"
)

type Key string
type Value string

type Args struct {
	K , V *string
}


type BaseOperations string

func isexit() {
	rd:= bufio.NewReader(os.Stdin)
	ip,_:=rd.ReadString('\n')
	ipstr:=string([]byte(ip)[0:])
	ipstr=ipstr[:len(ipstr)-1]
	if ipstr=="EXIT" {
		os.Exit(1)
	}
}

func main() {
	// file handler for most files
	//fileServer := http.FileServer(http.Dir("/var/www"))
	//http.Handle("/", fileServer)

	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printEnv)

	// deliver requests to the handlers
	// That's it!
	go isexit()	

	err := http.ListenAndServe(":8000", nil)
	checkError(err)
	
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	//env := os.Environ()
	//writer.Write([]byte("<html><head><h1>Adding key store</h1></head><body><table><form><tr><td>Enter key value</td><td><input type='text'></td></tr><tr><td><input type='submit'></td></tr></form></table></body></html>"))
	//for _, v := range env {
	//	writer.Write([]byte(v + "\n"))
	//}
	//writer.Write([]byte("</body></html>"))
    	//writer.Write([]byte(req.PostFormValue("key")))
	//writer.Write([]byte(req.PostFormValue("value")))
	//writer.Write([]byte(req.PostFormValue("add")))
	//writer.Write([]byte(req.Form.Get("add")))
	if req.PostFormValue("insert")!="" {
		//writer.Write([]byte(req.PostFormValue("insert")))	
		//writer.Write([]byte(req.PostFormValue("key")))	
		//writer.Write([]byte(req.PostFormValue("value")))
		

		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		//mk:="17"
		mk:=req.PostFormValue("key")
		//mv:="8"
		mv:=req.PostFormValue("value")

		args := Args{&mk, &mv}
		var reply string
		err = client.Call("BaseOperations.Insert", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//fmt.Printf("Reply %s %s %s\n", *args.K, *args.V, reply)
		//writer.Write([]byte("Inserted successfully"))
		
		str:=""
		
		if reply=="Insertion successful" {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td><td><input type='text'  name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td><td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Inserted key</td><td><input type='text' value='"+mk+"'></td>	<td>Inserted value</td><td><input type='text' value='"+mv+"'></td></tr></table></form></body></html>"	
		} else {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td>		<td><input type='text' name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td>		<td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Sorry,insert unsuccessful</td></tr></table></form></body></html>"
		}
		
		
		writer.Write([]byte(str))
	















	
	}
	if req.PostFormValue("removekey")!="" {
		//writer.Write([]byte(req.PostFormValue("removekey")))	
		//writer.Write([]byte(req.PostFormValue("key")))

		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		//mk:="17"
		mk:=req.PostFormValue("key")
		
		

		//args := Args{&mk, &mv}
		var reply string
		err = client.Call("BaseOperations.Removebykey", mk, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//fmt.Printf("Reply %s %s %s\n", *args.K, *args.V, reply)
		//writer.Write([]byte("Removebykey successfully"))

		str:=""
		
		if reply=="Remove by key is successful" {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td><td><input type='text'  name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td><td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Removed key</td><td><input type='text' value='"+mk+"'></td></tr></table></form></body></html>"	
		} else {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td>		<td><input type='text' name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td>		<td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Sorry,delete unsuccessful</td></tr></table></form></body></html>"
		}
		writer.Write([]byte(str))
	
	}
	
	if req.PostFormValue("removevalue")!="" {
		//writer.Write([]byte(req.PostFormValue("removevalue")))	
		//writer.Write([]byte(req.PostFormValue("value")))	


		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		//mv:="8"
		mv:=req.PostFormValue("value")
		
		

		//args := Args{&mk, &mv}
		var reply string
		err = client.Call("BaseOperations.Removebyvalue", mv, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//fmt.Printf("Reply %s %s %s\n", *args.K, *args.V, reply)
		//writer.Write([]byte("Removebyvalue successfully"))

		
		str:=""
		
		if reply=="Remove by value successful" {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td><td><input type='text'  name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td><td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Removed value</td><td><input type='text' value='"+mv+"'></td></tr></table></form></body></html>"	
		} else {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td>		<td><input type='text' name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td>		<td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Sorry,delete unsuccessful</td></tr></table></form></body></html>"
		}

		writer.Write([]byte(str))


	}
	if req.PostFormValue("searchkey")!="" {
		//writer.Write([]byte(req.PostFormValue("searchkey")))	
		//writer.Write([]byte(req.PostFormValue("key")))	

		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		//mk:="17"
		mk:=req.PostFormValue("key")
		
		

		//args := Args{&mk, &mv}
		var reply string
		err = client.Call("BaseOperations.Searchbykey", mk, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//fmt.Printf("Reply %s %s %s\n", *args.K, *args.V, reply)
		//writer.Write([]byte("Searchbykey successfully"))
		str:=""
		
		if reply!="" {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td><td><input type='text'  name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td><td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Searched value</td><td><input type='text' value='"+reply+"'></td></tr></table></form></body></html>"	
		} else {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td>		<td><input type='text' name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td>		<td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Sorry,search unsuccessful</td></tr></table></form></body></html>"
		}
		writer.Write([]byte(str))


	}
	if req.PostFormValue("searchvalue")!="" {
		//writer.Write([]byte(req.PostFormValue("searchvalue")))
		//writer.Write([]byte(req.PostFormValue("value")))


		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		//mv:="8"
		mv:=req.PostFormValue("value")
		
		

		//args := Args{&mk, &mv}
		var reply string
		err = client.Call("BaseOperations.Searchbyvalue", mv, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//fmt.Printf("Reply %s %s %s\n", *args.K, *args.V, reply)
		//writer.Write([]byte("Searchbyvalue successfully"))
		
		str:=""
		
		if reply!="" {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td><td><input type='text'  name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td><td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Searched key</td><td><input type='text' value='"+reply+"'></td></tr></table></form></body></html>"	
		} else {
			str="<html><head><h1>Key-store interface</h1></head><body><form name='input' action='http://localhost:8000/cgi-bin/printenv' method='post'><table><tr><td>Enter Key</td><td><input type='text' name='key'></td></tr><tr><td>Enter Value</td>		<td><input type='text' name='value'></td></tr><tr><td><input type='submit' value='Insert the key value pairs to the store' name='insert'></td><td><input type='submit' value='Remove by the key from the store' name='removekey'></td><td><input type='submit' value='Remove by the value from the store' name='removevalue'></td><td><input type='submit' value='Search by key in the store' name='searchkey'></td>		<td><input type='submit' value='Search by value in the store' name='searchvalue'></td></tr><tr><td>Sorry,search unsuccessful</td></tr></table></form></body></html>"
		}
		writer.Write([]byte(str))


	
	}
	
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

