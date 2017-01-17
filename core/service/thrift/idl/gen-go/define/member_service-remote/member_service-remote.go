// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"define"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Result Login(string user, string pwd, bool update)")
	fmt.Fprintln(os.Stderr, "  Member GetMember(i32 id)")
	fmt.Fprintln(os.Stderr, "  Member GetMemberByUser(string user)")
	fmt.Fprintln(os.Stderr, "  Profile GetProfile(i32 id)")
	fmt.Fprintln(os.Stderr, "  MemberSummary Summary(i32 memberId)")
	fmt.Fprintln(os.Stderr, "  string GetToken(i32 memberId, bool reset)")
	fmt.Fprintln(os.Stderr, "  bool CheckToken(i32 memberId, string token)")
	fmt.Fprintln(os.Stderr, "  void RemoveToken(i32 memberId)")
	fmt.Fprintln(os.Stderr, "  Address GetAddress(i32 memberId, i32 addrId)")
	fmt.Fprintln(os.Stderr, "  Account GetAccount(i32 memberId)")
	fmt.Fprintln(os.Stderr, "   InviterArray(i32 memberId, i32 depth)")
	fmt.Fprintln(os.Stderr, "  Result ChargeAccount(i32 memberId, i32 account, i32 kind, string title, string outerNo, double amount, i32 relateUser)")
	fmt.Fprintln(os.Stderr, "  Result DiscountAccount(i32 memberId, i32 account, string title, string outerNo, double amount, i32 relateUser, bool mustLargeZero)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := define.NewMemberServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "Login":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "Login requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.Login(value0, value1, value2))
		fmt.Print("\n")
		break
	case "GetMember":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMember requires 1 args")
			flag.Usage()
		}
		tmp0, err37 := (strconv.Atoi(flag.Arg(1)))
		if err37 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetMember(value0))
		fmt.Print("\n")
		break
	case "GetMemberByUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMemberByUser requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetMemberByUser(value0))
		fmt.Print("\n")
		break
	case "GetProfile":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetProfile requires 1 args")
			flag.Usage()
		}
		tmp0, err39 := (strconv.Atoi(flag.Arg(1)))
		if err39 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetProfile(value0))
		fmt.Print("\n")
		break
	case "Summary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Summary requires 1 args")
			flag.Usage()
		}
		tmp0, err40 := (strconv.Atoi(flag.Arg(1)))
		if err40 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.Summary(value0))
		fmt.Print("\n")
		break
	case "GetToken":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetToken requires 2 args")
			flag.Usage()
		}
		tmp0, err41 := (strconv.Atoi(flag.Arg(1)))
		if err41 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetToken(value0, value1))
		fmt.Print("\n")
		break
	case "CheckToken":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CheckToken requires 2 args")
			flag.Usage()
		}
		tmp0, err43 := (strconv.Atoi(flag.Arg(1)))
		if err43 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.CheckToken(value0, value1))
		fmt.Print("\n")
		break
	case "RemoveToken":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveToken requires 1 args")
			flag.Usage()
		}
		tmp0, err45 := (strconv.Atoi(flag.Arg(1)))
		if err45 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.RemoveToken(value0))
		fmt.Print("\n")
		break
	case "GetAddress":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetAddress requires 2 args")
			flag.Usage()
		}
		tmp0, err46 := (strconv.Atoi(flag.Arg(1)))
		if err46 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err47 := (strconv.Atoi(flag.Arg(2)))
		if err47 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GetAddress(value0, value1))
		fmt.Print("\n")
		break
	case "GetAccount":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAccount requires 1 args")
			flag.Usage()
		}
		tmp0, err48 := (strconv.Atoi(flag.Arg(1)))
		if err48 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetAccount(value0))
		fmt.Print("\n")
		break
	case "InviterArray":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "InviterArray requires 2 args")
			flag.Usage()
		}
		tmp0, err49 := (strconv.Atoi(flag.Arg(1)))
		if err49 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err50 := (strconv.Atoi(flag.Arg(2)))
		if err50 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.InviterArray(value0, value1))
		fmt.Print("\n")
		break
	case "ChargeAccount":
		if flag.NArg()-1 != 7 {
			fmt.Fprintln(os.Stderr, "ChargeAccount requires 7 args")
			flag.Usage()
		}
		tmp0, err51 := (strconv.Atoi(flag.Arg(1)))
		if err51 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err52 := (strconv.Atoi(flag.Arg(2)))
		if err52 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		tmp2, err53 := (strconv.Atoi(flag.Arg(3)))
		if err53 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5, err56 := (strconv.ParseFloat(flag.Arg(6), 64))
		if err56 != nil {
			Usage()
			return
		}
		value5 := argvalue5
		tmp6, err57 := (strconv.Atoi(flag.Arg(7)))
		if err57 != nil {
			Usage()
			return
		}
		argvalue6 := int32(tmp6)
		value6 := argvalue6
		fmt.Print(client.ChargeAccount(value0, value1, value2, value3, value4, value5, value6))
		fmt.Print("\n")
		break
	case "DiscountAccount":
		if flag.NArg()-1 != 7 {
			fmt.Fprintln(os.Stderr, "DiscountAccount requires 7 args")
			flag.Usage()
		}
		tmp0, err58 := (strconv.Atoi(flag.Arg(1)))
		if err58 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err59 := (strconv.Atoi(flag.Arg(2)))
		if err59 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4, err62 := (strconv.ParseFloat(flag.Arg(5), 64))
		if err62 != nil {
			Usage()
			return
		}
		value4 := argvalue4
		tmp5, err63 := (strconv.Atoi(flag.Arg(6)))
		if err63 != nil {
			Usage()
			return
		}
		argvalue5 := int32(tmp5)
		value5 := argvalue5
		argvalue6 := flag.Arg(7) == "true"
		value6 := argvalue6
		fmt.Print(client.DiscountAccount(value0, value1, value2, value3, value4, value5, value6))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
