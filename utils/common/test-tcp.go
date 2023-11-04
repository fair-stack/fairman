// Package common
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-09-05 15:44
package common

import (
	"net"
	"time"
)

//func TcpGather(ip string, ports []string) map[string]string {
//	// inspect emqx 1883, 8083, 8080, 18083 inspect
//	results := make(map[string]string)
//	for _, port := range ports {
//		address := net.JoinHostPort(ip, port)
//		// 3 seconds until timeout
//		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
//		if err != nil {
//			results[port] = "failed"
//			// todo log handler
//		} else {
//			if conn != nil {
//				results[port] = "success"
//				_ = conn.Close()
//			} else {
//				results[port] = "failed"
//			}
//		}
//	}
//	return results
//}

func TcpGather(ip string, ports []string) bool {
	// inspect emqx 1883, 8083, 8080, 18083 inspect
	//results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 seconds until timeout
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			return false
		} else {
			if conn != nil {
				_ = conn.Close()
			} else {
				return false
			}
		}
	}
	return true
}
