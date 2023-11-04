// Package utils
// @program: fairman-server-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-07 16:54
package utils

import (
	"bytes"
	"cnic/fairman-gin/global"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Get sendGETsend
// url：         Request Address
// retStruct:    Returned structure
// response：    Request return content
func Get(url string, retStruct interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Timeout time：5Timeout time
	client := &http.Client{Timeout: 5 * time.Minute, Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error() + "https 28")
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error() + "https 45")
		return err
	}

	//fmt.Println(string(result))

	err = json.Unmarshal(result, &retStruct)
	return err
}

// Post sendPOSTsend
// url：         Request Address
// data：        POSTRequest submitted data
// contentType： Request Body Format，Request Body Format：application/json
// retStruct: 	 Returned structure
// content：     Request to put back the content
func Post(url string, data interface{}, contentType string, retStruct interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Timeout time：5Timeout time
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
	var jsonStr []byte
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}

	if contentType == "" {
		contentType = "application/json"
	}

	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(result, &retStruct)
	return err
}

func Put(url string, data interface{}, contentType string, retStruct interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var jsonStr []byte
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	if contentType == "" {
		contentType = "application/json"
	}

	req.Header.Add("Content-Type", contentType)
	client := http.DefaultClient
	client.Transport = tr
	res, _ := client.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(result, &retStruct)
	return nil
}

func Delete(url string, data interface{}, contentType string, retStruct interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var jsonStr []byte
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}

	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))

	if contentType == "" {
		contentType = "application/json"
	}

	req.Header.Add("Content-Type", contentType)
	client := http.DefaultClient
	client.Transport = tr
	res, _ := client.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(result, &retStruct)
	return nil
}

//func DownloadFile(url string, localPath string) error {
//	var (
//		buf = make([]byte, 32*1024)
//	)
//
//	tmpFilePath := localPath + ".download"
//
//	client := new(http.Client)
//
//	resp, err := client.Get(url)
//	if err != nil {
//		return err
//	}
//	file, err := os.Create(tmpFilePath)
//	if resp.Body == nil {
//		return errors.New("Failed to create file")
//	}
//	defer resp.Body.Close()
//
//	buffer, err := io.CopyBuffer(file, resp.Body, buf)
//
//	fmt.Println(buffer)
//	if err != nil {
//		return err
//	}
//
//	if err == nil {
//		file.Close()
//		err = os.Rename(tmpFilePath, localPath)
//	}
//	return err
//}

func DownLoad(url string, path string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Create output file
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return nil
}

func DownloadFile2(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func DownLoadFile(strUrl string, strPath string) error {
	fmt.Println("Download files：", strUrl, strPath)
	// JudgestrUrlJudgehttp/httpsJudge，JudgehttpJudge
	if !strings.HasPrefix(strUrl, "http") {
		strUrl = "http://" + strUrl
	}
	_, err := url.ParseRequestURI(strUrl)
	if err != nil {
		return errors.New("URL error")
	}

	client := http.DefaultClient
	//client.Timeout = time.Second * 60 //Set timeout time
	resp, err := client.Get(strUrl)
	if err != nil {
		global.FairLog.Error(fmt.Sprintf("cannot fetch URL %q: %v", strUrl, err))
		return err
	}

	if resp.StatusCode != http.StatusOK {
		global.FairLog.Error(fmt.Sprintf("unexpected http GET status: %s", resp.Status))
		return err
	}

	if resp.ContentLength <= 0 {
		log.Println("[*] Destination server does not support breakpoint download.")
	}

	contentType := resp.Header.Get("Content-Type")
	if resp.StatusCode != http.StatusOK {
		global.FairLog.Error(fmt.Sprintf("unexpected http GET status: %s", resp.Status))
		return err
	}
	if contentType == "application/json; charset=utf-8" {
		formData := make(map[string]interface{})
		err = json.NewDecoder(resp.Body).Decode(&formData)
		if err != nil {
			fmt.Println(fmt.Errorf("cannot decode JSON: %v", err))
		}
		fmt.Println(formData)
		return Errorf("Download failed, Download failed")
	} else {
		//application/octet-stream
		raw := resp.Body
		defer raw.Close()

		// Create the file
		out, err := os.Create(strPath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Writer the body to file
		written, err := io.Copy(out, resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(Errorf(written, " bytes downloaded."))

		//reader := bufio.NewReaderSize(raw, 1024*32)
		//
		//file, err := os.Create(strPath)
		//if err != nil {
		//	panic(err)
		//}
		//writer := bufio.NewWriter(file)
		//
		//buff := make([]byte, 32*1024)
		//written := 0
		//for {
		//	nr, er := reader.Read(buff)
		//	if nr > 0 {
		//		nw, ew := writer.Write(buff[0:nr])
		//		if nw > 0 {
		//			written += nw
		//		}
		//		if ew != nil {
		//			err = ew
		//			break
		//		}
		//		if nr != nw {
		//			err = io.ErrShortWrite
		//			break
		//		}
		//	}
		//	if er != nil {
		//		if er != io.EOF {
		//			err = er
		//		}
		//		break
		//	}
		//}
		//if err != nil {
		//	panic(err)
		//}
	}
	return nil
}

func DataJson(data interface{}, resp interface{}) error {
	arr, err := json.Marshal(data)
	if err != nil {
		return Errorf(err)
	}
	return json.Unmarshal(arr, &resp)
}

func ScanPort(protocol string, port uint32) error {
	//for _, port := range ports {
	conn, err := net.Dial(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		return nil
		conn.Close()
	}
	return errors.New(fmt.Sprintf("port%dport", port))

	//}
	//return nil
}
