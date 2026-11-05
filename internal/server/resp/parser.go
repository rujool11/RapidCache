package respserver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
 RESP (REdis Serialization Protocol) Quick Reference:

 Every data type is prefixed by a type byte, followed by the content and \r\n.

 Types:
   + : Simple Strings (e.g., "+PONG\r\n")
   - : Errors         (e.g., "-ERR unknown command\r\n")
   : : Integers       (e.g., ":1000\r\n")
   $ : Bulk Strings   (e.g., "$5\r\nhello\r\n" -> length followed by data)
   * : Arrays         (e.g., "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n" -> count followed by elements)

 Client commands (like GET, SET, PING) are always sent as RESP Arrays of Bulk Strings.

 Eg.

 *3\r\n
 $3\r\nSET\r\n
 $5\r\nmykey\r\n
 $7\r\nmyvalue\r\n

*/

const MAX_NUM_ARGS = 1024

// returns slice of arguments
// only accepts arrays of bulk strings as per standard RESP commands
func parseCommand(reader *bufio.Reader) ([]string, error) {
	// io.EOF is thrown when no input available
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if prefix != '*' {
		return nil, fmt.Errorf("Expected array prefix '*', got %c", prefix)
	}

	numArgsString, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// trim trailing \r or \n
	numArgsString = strings.TrimSpace(numArgsString)

	numArgs, err := strconv.Atoi(numArgsString)
	if err != nil {
		return nil, err
	}

	if numArgs <= 0 || numArgs >= MAX_NUM_ARGS {
		return nil, fmt.Errorf("Expected number of arguments between 1 - %d, received %d", MAX_NUM_ARGS, numArgs)
	}

	args, err := parseBulkStrings(numArgs, reader)
	if err != nil {
		return nil, err
	}

	return args, nil
}

// eg. $5\r\nhello\r\n
func parseBulkStrings(n int, reader *bufio.Reader) ([]string, error) {
	args := make([]string, 0, n)
	for range n {
		prefix, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}

		// 1. bulk strings must start with "$"
		if prefix != '$' {
			return nil, fmt.Errorf("Expected '$', got %c", prefix)
		}

		// 2. read length line eg. "3\r\n"
		lenStr, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Error reading length %s", err)
		}

		lenStr = strings.TrimSpace(lenStr)
		length, err := strconv.Atoi(lenStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid bulk string length: %s", err)
		}

		// 3. read exact number of bytes as length
		buf := make([]byte, length) // buffer of exact data bytes length
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return nil, err
		}

		// 4. consume training CRLF ("\r\n\")
		trailing := make([]byte, 2)
		_, err = io.ReadFull(reader, trailing)
		if err != nil {
			return nil, err
		}

		if trailing[0] != '\r' && trailing[1] != '\n' {
			return nil, fmt.Errorf("Missing trailing CRLF in bulk string")
		}

		args = append(args, string(buf))

	}

	return args, nil
}
