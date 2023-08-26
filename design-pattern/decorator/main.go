package main

// Notifier 通知器接口
type Notifier interface {
	Notify()
}

type BasicNotifier struct {
}

func (b *BasicNotifier) Notify() {
	println("send basic message")
}

type SMSDecorator struct {
	// 被装饰的对象
	Notifier
}

func (s *SMSDecorator) Notify() {
	// 注意：此处发送短信的逻辑是在调用被装饰对象的 Notify 方法之后，实际上这里可以根据需求自由调整
	s.Notifier.Notify()
	println("send SMS")
}

type EmailDecorator struct {
	// 被装饰的对象
	Notifier
}

func (e *EmailDecorator) Notify() {
	e.Notifier.Notify()
	println("send email")
}

type WeChatDecorator struct {
	// 被装饰的对象
	Notifier
}

func (w *WeChatDecorator) Notify() {
	w.Notifier.Notify()
	println("send WeChat")
}

func main() {
	weChatDecorator := &WeChatDecorator{
		Notifier: &EmailDecorator{
			Notifier: &SMSDecorator{
				Notifier: &BasicNotifier{},
			},
		},
	}
	// send basic message
	// send SMS
	// send email
	// send WeChat
	weChatDecorator.Notify()
}
