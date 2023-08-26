package main

// Publisher 发布者接口
type Publisher interface {
	AddSubscriber(subscriber Subscriber)    // 添加订阅者
	RemoveSubscriber(subscriber Subscriber) // 移除订阅者
	NotifySubscribers()                     // 通知订阅者
}

// Subscriber 订阅者接口
type Subscriber interface {
	GetID() string                // 获取订阅者ID
	Notify(contextMessage string) // 通知订阅者
}

// Object 被观察者, 实现 Publisher 接口
type Object struct {
	Subscribers    []Subscriber // 订阅者列表
	ContextMessage string       // 上下文信息
}

func (p *Object) AddSubscriber(s Subscriber) {
	p.Subscribers = append(p.Subscribers, s)
}

func (p *Object) RemoveSubscriber(s Subscriber) {
	for i, subscriber := range p.Subscribers {
		if subscriber.GetID() == s.GetID() {
			p.Subscribers = append(p.Subscribers[:i], p.Subscribers[i+1:]...)
		}
	}
}

func (p *Object) NotifySubscribers() {
	for _, subscriber := range p.Subscribers {
		subscriber.Notify(p.ContextMessage)
	}
}

// UpdateContextMessage 更新上下文信息
func (p *Object) UpdateContextMessage(newMessage string) {
	p.ContextMessage = newMessage
	p.NotifySubscribers()
}

// SubscriberA 订阅者A, 实现 Subscriber 接口
type SubscriberA struct {
	ID string
}

func (s *SubscriberA) GetID() string {
	return s.ID
}

func (s *SubscriberA) Notify(contextMessage string) {
	println("SubscriberA received:", contextMessage)
}

// SubscriberB 订阅者B, 实现 Subscriber 接口
type SubscriberB struct {
	ID string
}

func (s *SubscriberB) GetID() string {
	return s.ID
}

func (s *SubscriberB) Notify(contextMessage string) {
	println("SubscriberB received:", contextMessage)
}

func main() {
	// 创建被观察者
	object := Object{}

	// 创建订阅者
	subscriberA := SubscriberA{ID: "subscriberA"}
	subscriberB := SubscriberB{ID: "subscriberB"}

	// 添加订阅者 A 和 B
	object.AddSubscriber(&subscriberA)
	object.AddSubscriber(&subscriberB)

	// 更新上下文信息, 通知所有订阅者
	object.UpdateContextMessage("Hello World!")

	// 移除订阅者B
	object.RemoveSubscriber(&subscriberB)

	// 更新上下文信息,
	object.UpdateContextMessage("Hello World Again!")
}
