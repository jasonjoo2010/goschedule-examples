# Func Worker Example

Func worker is a very simple, triggerred based design worker model. We can focus
on the logic need to be executed when triggerring. So we could optimize what we have
done in `simple worker` example.

## Definition

```text
Strategy:{"Id":"HotSellingRefresher","IpList":["127.0.0.1"],"MaxOnSingleScheduler":0,"Total":1,"Kind":2,"Bind":"HotSellingRefresher","Parameter":"","Enabled":true,"CronBegin":"0/30 * * * * ?","CronEnd":"","Extra":null}
```

## Cron

This time we use the cron feature of strategy to make it be called every 30 seconds. It follows general cron pattern with
seconds in precisition.

## Test

You can observe it having the output as below:

```text
2020-06-23T16:11:00+08:00 refreshed
2020-06-23T16:11:30+08:00 refreshed
2020-06-23T16:12:00+08:00 refreshed
2020-06-23T16:12:30+08:00 refreshed
2020-06-23T16:13:00+08:00 refreshed
```

Again you can stop any of the instances to see the scheduling.

## Other Features

`Strategy` has other features and you can see the [documentation](https://github.com/jasonjoo2010/goschedule/blob/master/CRON.md),
console panel for further information:

* Parameter
* Same worker can be binded with different strategies(different strategyId when invoking)
* Cron
* Multiple instance / limit on single instance
* Targets allowed to be scheduled on
