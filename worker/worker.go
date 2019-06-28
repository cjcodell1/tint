package worker

import (
    "sync"
    "log"

    "github.com/cjcodell1/tint/tm"
)

type Job struct {
    Id int
    Input string
    TM tm.TuringMachine
}

type Result struct {
    Id int
    Input string
    TM tm.TuringMachine
    Configs []tm.Config // in chronological order
}


func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    for j := range jobs {
        var conf tm.Config
        var err error
        var confs = make([]tm.Config, 0)
        var result Result

        for conf = j.TM.Start(j.Input); !(j.TM.IsAccept(conf) || j.TM.IsReject(conf)); conf, err = j.TM.Step(conf) {
            if err != nil {
                log.Fatalln(err.Error())
            }

            confs = append(confs, conf)

        }

        // need to append the Accept or Reject config to the slice
        confs = append(confs, conf)
        result = Result{j.Id, j.Input, j.TM, confs}
        results <- result
    }

    wg.Done()
}

func TestAll(inputs []string, tm tm.TuringMachine) []Result {
    var wg sync.WaitGroup
    jobs := make(chan Job, 100)
    results := make(chan Result, 100)

    // create worker pool
    for w := 1; w <= 5; w++ {
        wg.Add(1)
        go worker(jobs, results, &wg)
    }

    // hand out jobs
    for index, input := range inputs {
        turing := tm
        job := Job{index, input, turing}
        jobs <- job
    }
    close(jobs)

    // wait for workers to finish
    wg.Wait()

    close(results) // workers are done sending to results

    // convert channel to slice
    toReturn := make([]Result, 0)
    for r := range results {
        toReturn = append(toReturn, r)
    }

    return toReturn
}
