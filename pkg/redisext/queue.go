package redisext

import (
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"github.com/mohsensamiei/gopher/pkg/errors"
	"time"
)

type Queue struct {
	cli  *Client
	base rmq.Queue
}

type Delivery struct {
	queue *Queue
	base  rmq.Delivery
}

func (d Delivery) Reject() {
	if err := d.base.Reject(); err != nil {
		d.queue.cli.errChan <- err
	}
}

func (d Delivery) Push() {
	if err := d.base.Push(); err != nil {
		d.queue.cli.errChan <- err
	}
}

func (d Delivery) Ack() {
	if err := d.base.Ack(); err != nil {
		d.queue.cli.errChan <- err
	}
}

func (d Delivery) Payload(value any) error {
	return json.Unmarshal([]byte(d.base.Payload()), value)
}

type Consumer interface {
	Tag() string
	Consume(delivery *Delivery) error
}

func (q *Queue) SetConsumer(c Consumer) error {
	if err := q.base.StartConsuming(q.cli.RedisQueuePool, time.Second); err != nil {
		return err
	}
	deliveryChan := make(chan rmq.Delivery)
	if _, err := q.base.AddConsumerFunc(c.Tag(), func(delivery rmq.Delivery) {
		deliveryChan <- delivery
	}); err != nil {
		return err
	}
	for i := int64(0); i < q.cli.RedisQueuePool; i++ {
		go func() {
			for delivery := range deliveryChan {
				if err := c.Consume(&Delivery{
					queue: q,
					base:  delivery,
				}); err != nil {
					q.cli.errChan <- errors.Cast(err).WithDetailF("payload: %v", delivery.Payload())
				}
			}
		}()
	}
	return nil
}

func (q *Queue) Publish(value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return q.base.Publish(string(data))
}
