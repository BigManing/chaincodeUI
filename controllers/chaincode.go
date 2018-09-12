package controllers

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	SVN    string = "SVN"
	INIT   string = "INIT"
	UPDATE string = "UPDATE"
	CLEAR  string = "CLEAR"
)

// ChaincodeController operations for InitChaincode
type ChaincodeController struct {
	beego.Controller
}

func getCurrentChaincodeName() (string, error) {
	fileName := "/tmp/currentChaincodeName"
	var ccName string
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return ccName, err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return ccName, err
	}
	ccName = string(buf)
	return ccName, nil
}
func (c *ChaincodeController) Get() {
	// 1 获取当前chaincode名称
	// 2 读取、写入其版本号

	operation := c.GetString("operation")
	chaincodeName := c.GetString("ccName")
	chaincodeName = "IncrementChaincode"

	var cmd *exec.Cmd
	switch operation {
	case SVN:
		cmd = exec.Command("svn", "update")
	case INIT:
		ccName, err := getCurrentChaincodeName()
		if err != nil {
			c.Ctx.WriteString("获取当前chaincode名称失败：" + err.Error())
		}
		c.Ctx.WriteString("ccName:"+ccName)
		return
	case UPDATE:
		ccName, err := getCurrentChaincodeName()
		if err != nil {
			c.Ctx.WriteString("获取当前chaincode名称失败：" + err.Error())
			return
		}
		if ccName=="" {
			c.Ctx.WriteString("当前chaincode名称为空！")
			return
		}
		fileName := "/tmp/" + ccName
		// 读取文件
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
		if err != nil {
			c.Ctx.WriteString("打开失败：" + err.Error())
			return
		}
		defer file.Close()
		// 读取版本号
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			c.Ctx.WriteString("读取失败：" + err.Error())
			return
		}
		versionStr := string(buf)
		// 字符串转成float
		f, err := strconv.ParseFloat(versionStr, 64)
		if err != nil {
			c.Ctx.WriteString("解析版本号失败：" + err.Error())
			return
		}
		f = f + 0.1
		// 重新更新版本号
		if err := ioutil.WriteFile(fileName, []byte(strconv.FormatFloat(f, 'f', -1, 64)), 0644); err != nil {
			c.Ctx.WriteString("写入失败：" + err.Error())
		}
		// ------------------------- 根据原先的版本号 更新chaincode
		cmd = exec.Command("docker", "exec", "cli", "./scripts/beijinsuo_increment.sh", "upgrade", versionStr)
	case CLEAR:
		cmd = exec.Command("svn", "update")
	default:
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("不支持这个操作！" + operation))
		return
	}

	cmd.Dir = "/home/jiang/project/qianyuan/chaincode"
	out, err := cmd.CombinedOutput()

	resultCode := 200
	buffer := bytes.NewBuffer(out)
	if err != nil {
		resultCode = 500
		buffer.WriteString("\n更新失败！")
	}

	c.Ctx.Output.SetStatus(resultCode)
	buffer.WriteString("\n更新成功！")
	c.Ctx.Output.Body(buffer.Bytes())
	return
}
