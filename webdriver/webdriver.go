package webdriver

import (
	"os"
	"time"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var WebDriver selenium.WebDriver

//goland:noinspection GoUnhandledErrorResult,SpellCheckingInspection
func init() {
	chatgptProxyServer := os.Getenv("CHATGPT_PROXY_SERVER")
	if chatgptProxyServer == "" {
		return
	}

	chromeArgs := []string{
		"--no-sandbox",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--disable-blink-features=AutomationControlled",
		"--headless=new",
	}

	networkProxyServer := os.Getenv("NETWORK_PROXY_SERVER")
	if networkProxyServer != "" {
		chromeArgs = append(chromeArgs, "--proxy-server="+networkProxyServer)
	}

	WebDriver, _ = selenium.NewRemote(selenium.Capabilities{
		"chromeOptions": chrome.Capabilities{
			Args:            chromeArgs,
			ExcludeSwitches: []string{"enable-automation"},
		},
	}, chatgptProxyServer)

	WebDriver.Get(api.ChatGPTUrl)
	if !isAccessDenied(WebDriver) && !isAtCapacity(WebDriver) {
		HandleCaptcha(WebDriver)
		WebDriver.SetAsyncScriptTimeout(time.Second * api.ScriptExecutionTimeout)
	}
}