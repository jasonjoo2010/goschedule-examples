# Task Worker

Task worker is a bigger implementation based on `simple worker`. `Task` as a concept is abstracted
from general purpose workflows.

## Organization Model

Logics are organized and splitted into `select()` and `execute()` which:

`select()` targets the `tasks` should be done or executed later;  
`execute(task)` and `execute([]task)` focuses on how to get it/them done correctly or consistently.

There are two scheduling model for task: `normal` and `stream`. You can get further information [here](https://github.com/jasonjoo2010/goschedule/blob/master/MODELS.md).

And it's a litter complicated we can take a brief tour through several examples following.

## Definition

```text
Strategy:{"Id":"BatchExecutor","IpList":["127.0.0.1"],"MaxOnSingleScheduler":0,"Total":2,"Kind":3,"Bind":"BatchExecutor","Parameter":"","Enabled":false,"CronBegin":"","CronEnd":"","Extra":null}
Strategy:{"Id":"OrderRecycleTask","IpList":["127.0.0.1"],"MaxOnSingleScheduler":0,"Total":1,"Kind":3,"Bind":"OrderRecycleTask","Parameter":"","Enabled":false,"CronBegin":"0 * * * * ?","CronEnd":"","Extra":null}
Strategy:{"Id":"SingleExecutor","IpList":["127.0.0.1"],"MaxOnSingleScheduler":0,"Total":1,"Kind":3,"Bind":"SingleExecutor","Parameter":"","Enabled":false,"CronBegin":"","CronEnd":"","Extra":null}
Strategy:{"Id":"SingleStreamExecutor","IpList":["127.0.0.1"],"MaxOnSingleScheduler":0,"Total":1,"Kind":3,"Bind":"SingleStreamExecutor","Parameter":"","Enabled":false,"CronBegin":"","CronEnd":"","Extra":null}
Task:{"Id":"BatchExecutor","IntervalNoData":5000,"Interval":1000,"FetchCount":20,"BatchCount":5,"ExecutorCount":2,"Model":0,"Parameter":"","Bind":"batchExecutor","Items":[{"Id":"p0","Parameter":""},{"Id":"p1","Parameter":""}],"MaxTaskItems":0,"HeartbeatInterval":5000,"DeathTimeout":60000}
Task:{"Id":"OrderRecycleTask","IntervalNoData":500,"Interval":0,"FetchCount":10,"BatchCount":1,"ExecutorCount":3,"Model":0,"Parameter":"","Bind":"orderRecycleTask","Items":[{"Id":"p0","Parameter":""}],"MaxTaskItems":0,"HeartbeatInterval":5000,"DeathTimeout":60000}
Task:{"Id":"SingleExecutor","IntervalNoData":5000,"Interval":1000,"FetchCount":5,"BatchCount":1,"ExecutorCount":3,"Model":0,"Parameter":"","Bind":"singleExecutor","Items":[{"Id":"p0","Parameter":""}],"MaxTaskItems":0,"HeartbeatInterval":5000,"DeathTimeout":60000}
Task:{"Id":"SingleStreamExecutor","IntervalNoData":500,"Interval":0,"FetchCount":5,"BatchCount":1,"ExecutorCount":3,"Model":1,"Parameter":"","Bind":"singleStreamExecutor","Items":[{"Id":"p0","Parameter":""}],"MaxTaskItems":0,"HeartbeatInterval":5000,"DeathTimeout":60000}
```

## Single / Batch / Stream

First there are three demo tasks: Single executor, batch executor under normal scheduling model and another task under stream scheduling model. You can look deeper to feel the difference.

Further more, `singleStreamExecutor` is defined as a pre-defined instance which means if you want scheduler to use / share the same instance you can follow this.

To demostrate the difference straightly between normal and stream model, you may got following outputs from `SingleExecutor` strategy:

```log
....
fetch for {p0 }
Finish task p0:1
Finish task p0:2
Finish task p0:3
Finish task p0:5
Finish task p0:4
fetch for {p0 }
Finish task p0:3
Finish task p0:1
Finish task p0:2
Finish task p0:5
Finish task p0:4
.....
```

The orders is changing in the same batch (or `select()`) but the phases don't change(...5-5-5-5...). For `SingleStreamExecutor`:

```log
....
Finish task(single instanced) p0:4
Finish task(single instanced) p0:3
Finish task(single instanced) p0:5
fetch for {p0 }
Finish task(single instanced) p0:2
Finish task(single instanced) p0:1
Finish task(single instanced) p0:3
fetch for {p0 }
Finish task(single instanced) p0:1
Finish task(single instanced) p0:4
Finish task(single instanced) p0:5
Finish task(single instanced) p0:4
Finish task(single instanced) p0:3
Finish task(single instanced) p0:2
....
```

The first idle executor will always try to get new tasks (call `select()`) and get back to work. They are working in stream model.

## Order Recycling Task

So let's get to an actual example. Assume we want to "expire" any order that user doesn't pay for it in N minutes from the time he/she created it. We need an order scanner.

After some discussions, we decide:

1. Scan should be triggerred in every minute.
2. Scan should stop if no more order expired.
3. Scan should continue if scanning cannot be done in one minute.
4. Scan should not be triggerred if it was running.

We use `orderRecycleTask` to simulate it. Because we don't need it to be partitioned so we just set it an arbitary task item `p0`. To make it scheduled periodicly we use cron expression. For detailed settings please refer it on you console panel. You may got the output as below:

```log
2020-06-29T16:52:00+08:00 select 8 expired orders
 2020-06-29T16:52:00+08:00 close order 2279094282
 2020-06-29T16:52:00+08:00 close order 4638730130
 .....
 2020-06-29T16:52:00+08:00 close order 757821653
 2020-06-29T16:52:00+08:00 select 7 expired orders
 2020-06-29T16:52:00+08:00 close order 7220128104
 2020-06-29T16:52:01+08:00 close order 6512185088
 2020-06-29T16:52:01+08:00 close order 5470995089
 2020-06-29T16:52:01+08:00 close order 7092216550
 .....
 2020-06-29T16:52:03+08:00 close order 5263969896
 2020-06-29T16:52:03+08:00 close order 9096906095
 2020-06-29T16:52:03+08:00 close order 4980605223
 2020-06-29T16:52:03+08:00 no more
 2020-06-29T16:53:00+08:00 select 1 expired orders
 2020-06-29T16:53:00+08:00 close order 5105923662
 2020-06-29T16:53:00+08:00 select 2 expired orders
 2020-06-29T16:53:00+08:00 close order 6474476365
 2020-06-29T16:53:00+08:00 close order 7425487454
 2020-06-29T16:53:00+08:00 select 2 expired orders
 2020-06-29T16:53:00+08:00 close order 3594053309
 2020-06-29T16:53:00+08:00 close order 7823889182
 2020-06-29T16:53:00+08:00 select 7 expired orders
 2020-06-29T16:53:01+08:00 close order 2045489127
 .....
```
