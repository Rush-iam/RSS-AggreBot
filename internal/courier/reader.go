package courier

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/db_client"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/rss_feed"
	"log"
	"time"
)

const readPeriod = time.Second * 30

func (c *courier) RunReader() {
	for {
		repeatTime := time.Now().Add(readPeriod)
		sources, err := c.db.GetActiveSources()
		if err != nil {
			log.Printf("Reader: %v", err)
		} else {
			for _, source := range sources {
				_, isBusy := c.sourcesInWork.LoadOrStore(source.Id, true)
				if !isBusy {
					go c.readerRoutine(source)
				}
			}
		}
		time.Sleep(time.Until(repeatTime))
	}
}

func (c *courier) readerRoutine(source *db_client.CourierSource) {
	newJob := job{source: source}

	feed, err := rss_feed.Fetch(source.Url)
	if err != nil {
		newJob.wasReadError = true
		log.Printf("readerRoutine: %+v, %v", source, err)
	} else {
		newJob.entries = feed.Items
	}
	c.jobQueue <- newJob
}
