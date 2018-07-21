package stream_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

func EncryptToStream(w io.Writer, iv []byte, pass string) (io.Writer, error) {
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("len(iv)[%d] != aes.BlockSize[%d]", len(iv), aes.BlockSize)
	}
	cryptBlock, err := aes.NewCipher([]byte(pass))
	if err != nil {
		return nil, err
	}
	streamEncrypt := cipher.NewCFBEncrypter(cryptBlock, iv[:])
	sw := cipher.StreamWriter{
		S: streamEncrypt,
		W: w,
	}
	return sw, nil
}

func DecryptFromStream(r io.Reader, iv []byte, pass string) (io.Reader, error) {
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("len(iv)[%d] != aes.BlockSize[%d]", len(iv), aes.BlockSize)
	}
	cryptBlock, err := aes.NewCipher([]byte(pass))
	if err != nil {
		return nil, err
	}
	streamDecrypt := cipher.NewCFBDecrypter(cryptBlock, iv[:])
	sr := cipher.StreamReader{
		S: streamDecrypt,
		R: r,
	}
	return sr, nil
}
