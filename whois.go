/*
 * Copyright © 2022, Théo BARRAGUÉ
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the “Software”), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * The Software is provided “as is”, without warranty of any kind, express or
 * implied, including but not limited to the warranties of merchantability, fitness
 * for a particular purpose and noninfringement. In no event shall the authors or
 * copyright holders X be liable for any claim, damages or other liability, whether
 * in an action of contract, tort or otherwise, arising from, out of or in
 * connection with the software or the use or other dealings in the Software.
 *
 * Except as contained in this notice, the name of the Théo BARRAGUÉ shall not be
 * used in advertising or otherwise to promote the sale, use or other dealings in
 * this Software without prior written authorization from the Théo BARRAGUÉ.
 */

package whois

import (
	"errors"
	"io"
	"net"
	"strings"
)

func send(str string, host string) (string, error) {
	conn, err := net.Dial("tcp", host+":43")

	if err != nil {
		return "", err
	}

	_, err = conn.Write([]byte(str))

	if err != nil {
		return "", err
	}

	var result strings.Builder
	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			if err != io.EOF {
				return "", err
			}

			break
		}

		result.WriteString(string(buffer[:n]))
	}

	return result.String(), nil
}

func getRefer(name string) (string, error) {
	result, err := send(name+"\r\n", "whois.iana.org")

	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(result, "\n") {
		if strings.HasPrefix(line, "refer") {
			splitted := strings.Split(line, ":")

			if len(splitted) > 1 {
				refer := strings.TrimSpace(splitted[1])
				return refer, nil
			}

			break
		}
	}

	return "", errors.New("no refer for " + name)
}

func Whois(name string) (string, error) {
	refer, err := getRefer(name)

	if err != nil {
		return "", err
	}

	result, err := send(name+"\r\n", refer)

	if err != nil {
		return "", err
	}

	return result, nil
}
