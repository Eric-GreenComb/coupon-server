package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/coupon-server/bean"
	"github.com/Eric-GreenComb/coupon-server/persist"
	"github.com/Eric-GreenComb/coupon-server/regexp"
)

// CreateAdminUser CreateAdminUser
func CreateAdminUser(c *gin.Context) {

	var _userFarams bean.UserParams
	if c.BindJSON(&_userFarams) != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "bind json error."})
		return
	}

	if _userFarams.UserID == "" || _userFarams.Name == "" || _userFarams.Pwd == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	if !regexp.IsMobile(_userFarams.UserID) {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "UserID must phone number."})
		return
	}

	sum := sha256.Sum256([]byte(_userFarams.Pwd))
	_Passwd := fmt.Sprintf("%x", sum)

	var _user bean.AdminUsers
	_user.UserID = _userFarams.UserID
	_user.Name = _userFarams.Name
	_user.Passwd = _Passwd
	_user.Email = _userFarams.Email

	err := persist.GetPersist().CreateAdminUser(_user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _user})
	}
}

// AdminUserInfo AdminUserInfo
func AdminUserInfo(c *gin.Context) {

	_userid := c.Params.ByName("userid")

	if _userid == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	user, err := persist.GetPersist().AdminUserInfo(_userid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": user})
}

// UpdateAdminUserPasswd UpdateAdminUserPasswd
func UpdateAdminUserPasswd(c *gin.Context) {

	var _userFarams bean.UserParams
	if c.BindJSON(&_userFarams) != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "bind json error."})
		return
	}

	_userid := _userFarams.UserID
	_old := _userFarams.OldPwd
	_new := _userFarams.Pwd

	if _userid == "" || _old == "" || _new == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	if _old == _new {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "new pwd is eq old pwd"})
		return
	}

	sumOld := sha256.Sum256([]byte(_old))
	oldPasswd := fmt.Sprintf("%x", sumOld)

	_, err := persist.GetPersist().AdminLogin(_userid, oldPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	sum := sha256.Sum256([]byte(_new))
	newPasswd := fmt.Sprintf("%x", sum)

	err = persist.GetPersist().UpdateAdminUserPasswd(_userid, newPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "success"})
}
