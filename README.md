# TeleBoto
- 电报机器人框架
- 目前功能尚不完善，持续开发中
- 使用`gorequest`开发

# 快速上手

1. 创建`TeleBot`对象

   ```go
   t := telegram.New().SetToken("123456:SDASDsafeWADFsdaDWdD").SetChatID("-421133341")
   ```

   或者

   ```go
   t := telegram.Create("123456:SDASDsafeWADFsdaDWdD", "-421133341")
   ```

2. 设置代理

   大陆用户通常需要设置代理才能正常使用电报，方法如下

   ```go
   t.SetProxy("socks5://127.0.0.1:1086")
   ```

   或者使用默认代理"socks5://127.0.0.1:1086"

   ```go
   t.UseDefaultProxy()
   ```

3. 发送文字消息

   返回的内容会保存在 `t.Body`，返回的响应保存在`t.Response`，错误包含在`t.Errors`

   ```go
   t.SendText(chatID, text, parseMode string) // 这个方法可以手动指定chatID, 发送内容, 发送模板,(该方法不推荐使用)
   ```

   发送纯文本

   ```go
   t.SendPlain(text string) // 这个发送方法会调用提前设置的chatID，以后的方法没有明确指定chatID都为使用提前设置的chatID
   ```

   例如

   ```go
   errs := t.SetChatID("-123456").SendPlain().Errors
   if errs != nil{
     panic(errs)
   }
   fmt.Println(string(t.Body))
   fmt.Println(r.Response.StatusCode)
   ```

   其他格式

   ```go
   t.SendMarkdown(text string)
   ```

   ```go
   t.SendHTML(text string)
   ```

4. 获取结构体返回

   默认的`t.Body`是`[]byte`类型，如若需要以结构体返回则使用`Fetch`系列方法

   *需要注意的是，定义的结构体不是和telegram官方的结构体百分百一样，但基本满足需求*

   例子

   ```go
   updates, errs := t.GetUpdates().FetchUpdates()
   ...
   ```

5. 发送消息不触发通知

   ```go
   t.DisableNotification()
   ```

   

   

