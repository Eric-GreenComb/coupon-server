package wechat

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

// BindJSON BindJSON
func BindJSON(data io.Reader, dest interface{}) error {
	value := reflect.ValueOf(dest)

	if value.Kind() != reflect.Ptr {
		return errors.New("BindJSON not a pointer")
	}

	decoder := json.NewDecoder(data)

	if err := decoder.Decode(dest); err != nil {
		return err
	}

	return nil
}

//WxpayCalcSign 微信支付 下单签名
func WxpayCalcSign(mReq map[string]interface{}, key string) string {

	//fmt.Println("========STEP 1, 对key进行升序排序.========")
	//fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	_sortedKeys := make([]string, 0)
	for k := range mReq {
		_sortedKeys = append(_sortedKeys, k)
	}

	sort.Strings(_sortedKeys)

	//fmt.Println("========STEP2, 对key=value的键值对用&连接起来，略过空值========")
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var _buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	for _, _k := range _sortedKeys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		_value := fmt.Sprintf("%v", mReq[_k])
		if _value != "" && _k != "signed" {
			_buffer.WriteString(_k)
			_buffer.WriteString("=")
			_buffer.WriteString(_value)
			_buffer.WriteString("&")
		}
	}

	//fmt.Println("========STEP3, 在键值对的最后加上key=API_KEY========")
	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		_buffer.WriteString("key=")
		_buffer.WriteString(key)
	}

	//fmt.Println("========STEP4, 进行MD5签名并且将所有字符转为大写.========")
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	_md5Ctx := md5.New()
	_md5Ctx.Write(_buffer.Bytes())
	_cipherStr := _md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(_cipherStr))

	return upperSign
}
