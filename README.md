# XZ-Key
go语言全局按键监听事件绑定（需要root运行）

# 依赖
[keylogger](https://github.com/MarinX/keylogger)

# 案例
``` golang
...
import xzkey "github.com/CxZMoE/XZ-Key"
...

k := xzkey.NewKeyboard()
if k != nil{
    defer k.StopReadEvent()
    k.BindKeyEvent("sayHallo", func() {
			fmt.Println("Hello")
        }, k.Keys[xzkey.LCTRL], k.Keys[xzkey.LALT], k.Keys[xzkey.H])
        // 按键组合数量没有被限制
    }
}

...

```